package httputils

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"inwpuun/proxy_assignment/pkg/errorutil"
	"net/http"
	"net/url"
	"time"

	"github.com/joomcode/errorx"
	"github.com/valyala/fasthttp"
)

type RequestOption struct {
	Req        *fasthttp.Request
	RetryCount int
	RetryDelay time.Duration
}

type HTTPOption interface {
	Apply(*RequestOption) error
}

type HTTPHeader struct {
	Key   string
	Value string
}

type HTTPArgs struct {
	*fasthttp.Args
}

type RetryOption struct {
	Count int
	Delay time.Duration
}

func (r RetryOption) Apply(option *RequestOption) error {
	option.RetryCount = r.Count
	option.RetryDelay = r.Delay
	return nil
}

func (h HTTPArgs) Apply(option *RequestOption) error {
	option.Req.SetBody(h.QueryString())
	option.Req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return nil
}

func (h HTTPHeader) Apply(option *RequestOption) error {
	option.Req.Header.Set(h.Key, h.Value)
	return nil
}

type HTTPQuery url.Values

func (q HTTPQuery) Apply(option *RequestOption) error {
	for key, values := range url.Values(q) {
		for _, value := range values {
			option.Req.URI().QueryArgs().Add(key, value)
		}
	}
	return nil
}

func declareReqResp() (req *fasthttp.Request, resp *fasthttp.Response) {
	return fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
}

func getClient() *fasthttp.Client {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	client := &fasthttp.Client{
		TLSConfig: tlsConfig,
	}
	return client
}

func WithRequestDetail(err *errorx.Error, path string, requestQuery string, requestBody string) *errorx.Error {
	err = err.WithProperty(RequestPathProperty, errorutil.ErrorMeta[string]{
		Type: "request_path",
		Data: path,
	})
	err = err.WithProperty(RequestQueryProperty, errorutil.ErrorMeta[string]{
		Type: "request_query",
		Data: requestQuery,
	})
	err = err.WithProperty(RequestBodyProperty, errorutil.ErrorMeta[string]{
		Type: "request_body",
		Data: requestBody,
	})
	return err
}

func Request[ReqBody any](method string, path string, body ReqBody) (resp *fasthttp.Response, err error) {
	req, resp := declareReqResp()

	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(path)
	if method != http.MethodGet {
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, HttpUtilsMarshalError.Wrap(err, "Error marshaling request body")
		}
		req.SetBody(reqBody)
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.SetMethod(method)

	reqOption := &RequestOption{
		RetryCount: 0,
		RetryDelay: 0,
	}

	client := getClient()

	err = errors.New("Request is not called due to invalid retry count")
	for range reqOption.RetryCount + 1 {
		err = client.Do(req, resp)
		if resp.StatusCode()/100 != 2 && reqOption.RetryCount > 0 {
			time.Sleep(reqOption.RetryDelay)
			continue
		}
		break
	}
	if err != nil {
		query := req.URI().QueryArgs().String()
		httpPath := fmt.Sprintf("%s %s", method, path)
		reqBody := string(req.Body())
		err := HttpUtilsRequestError.Wrap(err, "Error sending request")
		err = WithRequestDetail(err, httpPath, query, reqBody)

		return nil, err
	}

	return resp, nil
}

func RequestWithOnlyResponseBody[ReqBody, RespBody any](method string, path string, body ReqBody) (RespBody, error) {
	var defaultRespBody RespBody
	resp, err := Request(method, path, body)
	if err != nil {
		return defaultRespBody, errorutil.WithStack(err)
	}

	var result RespBody
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		err := HttpUtilsUnmarshalError.Wrap(err, "Error unmarshalling response")
		err = err.WithProperty(ResponseBodyProperty, errorutil.ErrorMeta[string]{
			Type: "response_body",
			Data: string(resp.Body()),
		})
		return defaultRespBody, err
	}

	return result, nil
}

func GetWithOnlyResponseBody[RespBody any](path string, options ...HTTPOption) (RespBody, error) {
	return RequestWithOnlyResponseBody[any, RespBody](http.MethodGet, path, nil)
}

func PostWithOnlyResponseBody[ReqBody, RespBody any](path string, body ReqBody, options ...HTTPOption) (RespBody, error) {
	return RequestWithOnlyResponseBody[ReqBody, RespBody](http.MethodPost, path, body)
}

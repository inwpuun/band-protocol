package httputils

import "github.com/joomcode/errorx"

var (
	HttpUtilsError        = errorx.NewNamespace("httputils")
	HttpUtilsRequestError = HttpUtilsError.NewType("request_error")

	HttpUtilsUnmarshalError = HttpUtilsError.NewType("unmarshal_error")
	HttpUtilsMarshalError   = HttpUtilsError.NewType("marshal_error")

	HttpUtilsApplyOptionError = HttpUtilsError.NewType("apply_option_error")
)

var (
	RequestPathProperty  = errorx.RegisterProperty("request_path")
	RequestQueryProperty = errorx.RegisterProperty("request_query")
	RequestBodyProperty  = errorx.RegisterProperty("request_body")
	ResponseBodyProperty = errorx.RegisterProperty("response_body")
)

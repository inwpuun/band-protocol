## Getting Started

first, Install Dependencies
`make tidy`

you can run this application using
`make start`

you can test Broadcast Transaction with cURL
`curl --location 'http://localhost:8080/v1/crypto/broadcast' \
--header 'Content-Type: application/json' \
--form 'symbol="ETH"' \
--form 'price="4500"' \
--form 'timestamp="1678912345"'`

or you can test Transaction Status Monitoring with cURL
`curl --location 'http://localhost:8080/v1/crypto/check/57e9dc07f96d63b621552f6a6817d7b8f5183e58b73dc8ae95d626de82b7dab7'`

this is example code to call Broadcast Transaction 
```
resp, err := httputils.PostWithOnlyResponseBody[models.BroadcastTransactionRequestDto, models.BroadcastTransactionResponseDto](url, in)
if err != nil {
    return models.BroadcastTransactionResponseDto{}, fmt.Errorf("can't broadcast transaction in url %s from crypto repository: %w", url, err)
}
```
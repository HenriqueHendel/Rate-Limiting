# Rate-Limiting
Token Bucket rate limiting algorithm build in Golang

There are two endpoints: 'createClient', where you can create a client to do some requests. So you define the name, max tokens and fill rate from a client. Fill rate is the rate from fill the client bucket. The other endpoint is to test the bucket algorithm, where the client use 2 tokens from he's bucket per request.

Link to POSTMAN API:
https://www.postman.com/henriquesantoslopes/workspace/henrique-lopes/example/15797996-ebb4a62e-5015-4dec-91ca-7bf9e78fbf60

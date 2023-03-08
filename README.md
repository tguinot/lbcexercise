# Le Bon Coin exercise

To install dependencies:
> $ go mod tidy

To build:
> $ go build main.go

To run:
> $ ./main -p <PORT>

Endpoint URL (POST): http://localhost:<PORT>/fizzbuzz
Request format:
{
	"int1": 4,
	"int2": 3,
	"limit": 12,
	"str1": "fizz",
	"str2": "buzz"
}

To query metrics:
> $ curl http://localhost:<PORT>/metrics

To run go tests:
> go test -v ./controllers

To run hurl tests (requires hurl -> https://hurl.dev):
> $ hurl --variable port=<PORT> hurl_tests/all.hurl

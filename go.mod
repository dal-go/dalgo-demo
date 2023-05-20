module github.com/dal-go/dalgo-demo

go 1.20

require (
	github.com/dal-go/dalgo v0.2.30
	github.com/dal-go/mocks4dalgo v0.1.8
	github.com/golang/mock v1.6.0
)

//replace github.com/dal-go/mocks4dalgo => ../mocks4dalgo

require github.com/strongo/random v0.0.1 // indirect

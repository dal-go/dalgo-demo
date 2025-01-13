module github.com/dal-go/dalgo-demo

go 1.23

toolchain go1.23.4

require (
	github.com/dal-go/dalgo v0.14.2
	github.com/dal-go/mocks4dalgo v0.1.27
	github.com/golang/mock v1.6.0
)

//replace github.com/dal-go/mocks4dalgo => ../mocks4dalgo

require github.com/strongo/random v0.0.1 // indirect

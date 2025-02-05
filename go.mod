module github.com/dal-go/dalgo-demo

go 1.23

toolchain go1.23.6

require (
	github.com/dal-go/dalgo v0.16.1
	github.com/dal-go/mocks4dalgo v0.2.0
	go.uber.org/mock v0.5.0
)

//replace github.com/dal-go/mocks4dalgo => ../mocks4dalgo

require github.com/strongo/random v0.0.1 // indirect

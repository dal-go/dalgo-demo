module github.com/dal-go/dalgo-demo

go 1.23.0

toolchain go1.25.5

require (
	github.com/dal-go/dalgo v0.26.5
	github.com/dal-go/mocks4dalgo v0.3.10
	go.uber.org/mock v0.6.0
)

//replace github.com/dal-go/mocks4dalgo => ../mocks4dalgo

require github.com/strongo/random v0.0.1 // indirect

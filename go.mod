module github.com/dal-go/dalgo-demo

go 1.23

toolchain go1.24.2

require (
	github.com/dal-go/dalgo v0.18.2
	github.com/dal-go/mocks4dalgo v0.2.6
	go.uber.org/mock v0.5.2
)

//replace github.com/dal-go/mocks4dalgo => ../mocks4dalgo

require github.com/strongo/random v0.0.1 // indirect

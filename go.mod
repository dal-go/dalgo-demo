module github.com/dal-go/dalgo-demo

go 1.24.0

toolchain go1.25.5

require (
	github.com/dal-go/dalgo v0.41.1
	go.uber.org/mock v0.6.0
)

//replace github.com/dal-go/mocks4dalgo => ../mocks4dalgo

require (
	github.com/RoaringBitmap/roaring v1.9.4 // indirect
	github.com/RoaringBitmap/roaring/v2 v2.14.4 // indirect
	github.com/bits-and-blooms/bitset v1.24.4 // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/strongo/random v0.0.1 // indirect
)

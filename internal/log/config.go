package log

type Config struct {
	Segment struct {
		MaxStoreBytes uint64
		MaxIndexByte  uint64
		InitialOffset uint64
	}
}

package reader

type IReader interface {
	ReadInt() int
	ReadDouble() float64
	ReadString() string
}

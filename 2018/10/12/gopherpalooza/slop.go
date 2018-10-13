package main

type deviceRegister struct {
	i1     int
	i2     int64
	i3     int
	i4     uint64
	s      string
	bitmap [256]int
	i5     int32
	i6     int
	f1     [2]float64
	ch1    [512]byte
	ch0    [180]byte
	i7     int64
	i8     int
}

func main() {
	go func(_ deviceRegister) {}(deviceRegister{})
}

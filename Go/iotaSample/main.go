package main

import "fmt"

type ByteSize float64
type ByteSlice []byte

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b/KB)
	}
	return fmt.Sprintf("%.2fB", b)
}

func (p *ByteSlice) Append(data []byte) {
	slice := *p
	*p = slice
}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	*p = slice
	return len(data), nil
}

func main(){
	a := ByteSize(1e3)
	fmt.Println(a.String())

	a = ByteSize(1e6)
	fmt.Println(a.String())

	a = ByteSize(1e9)
	fmt.Println(a.String())

	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days.", 7)
}

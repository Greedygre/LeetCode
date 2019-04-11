//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//)
//
//const (
//	_        = iota // ignore first value by assigning to blank identifier
//	KB int64 = 1 << (4 * iota)
//	MB
//	GB
//	TB
//	PB
//	EB
//)
//
//type a interface {
//	MarshalJSON() ([]byte, error)
//}
//type Counter struct {
//	n int
//}
//
//func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	ctr.n++
//	fmt.Fprintf(w, "counter = %d\n", ctr.n)
//}
//func main() {
//	picture := make([][]uint8, 5) // One row per unit of y.
//	// Allocate one large slice to hold all the pixels.
//	pixels := make([]uint8, 20) // Has type []uint8 even though picture is [][]uint8.
//	// Loop over the rows, slicing each row from the front of the remaining pixels slice.
//	pixels = []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
//	for i := range picture {
//		picture[i], pixels = pixels[:4], pixels[4:]
//	}
//	fmt.Println(picture)
//	fmt.Println(KB, MB,
//		GB,
//		TB,
//		PB,
//		EB,
//	)
//	var val a
//	if _, ok := val.(json.Marshaler); ok {
//		fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
//	}
//}
//
//// 256=16*4*4,2^4

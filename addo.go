package addo

import (
	"fmt"
	"flag"
	
)
func add(numone, numtwo int) int {
	numthree := numone + numtwo
	return numthree
}


func main() {
	numone := flag.Int("one", 0, "The First Number To Add")
	numtwo := flag.Int("two", 0, "The Second Number To Add")
	flag.Parse()
	fmt.Println(add(*numone, *numtwo))
}
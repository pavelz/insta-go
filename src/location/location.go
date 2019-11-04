package location

import "fmt"

type Location struct {
	lat float64
	lng float64
}

func (l Location) Save(){
	fmt.Printf("save!\n")
}

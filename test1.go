package main

import (
	"fmt"
	"math/rand"
)

func main(){
	i := 0
	var myslice = [100]int{}

	for i <=99 {
		x := rand.Intn(100)
		if x % 50 != 0 {
			
			i++
			continue
		}
		myslice[i] = x
		i++
	}

		


	//fmt.Printf("%v, %T\n", x,x)
	fmt.Printf("%v, %T, %v\n", myslice, myslice, len(myslice))
}
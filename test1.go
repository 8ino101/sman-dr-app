package main

import (
	"fmt"
	"math/rand"
)

type mymap map[bool]string

func main(){

	mappa := mymap{

		true: "top",
		false: "flop",
	}

	fmt.Println("mappa:", mappa)
	i := 0
	var myslice = [100]int{}

	for i <=99 {
		var y bool 
		x := rand.Intn(100)
		toss := rand.Intn(2)
		if toss == 1 {
			y = true
		}	else {
			y = false
		}
		fmt.Println("--->", mappa[y])
		myslice[i] = x
		i++
	}

	//fmt.Printf("%v, %T\n", x,x)
	fmt.Printf("%v, %T, %v\n", myslice, myslice, len(myslice))
}
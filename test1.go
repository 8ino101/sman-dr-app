package main

import (
	"fmt"
	"math/rand"
)

type mymap map[bool]string

const (

	dimRand = 100
)

func main(){

	mappa := mymap{

		true: "top",
		false: "flop",
	}

	fmt.Println("mappa:", mappa)
	i := 0
	var myslice = [dimRand]int{}

	for i <=99 {
		var y bool 
		x := rand.Intn(dimRand)
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
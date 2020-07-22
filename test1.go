package main

import (
	"fmt"
	"math/rand"
	"log"
	"os"
	)


type mymap map[bool]string

const (

	dimRand = 10
)

const (

	_ = iota
	a = 1 << (1*iota)
	b
	c
	d
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func init() {
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	InfoLogger.Println("main started")
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	
	mappa := mymap{

		true: "top",
		false: "flop",
	}

	fmt.Println("mappa:", mappa)

	var myslice = [dimRand]int{}

	for i := 0; i < dimRand; i++ {
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
	}

	//fmt.Printf("%v, %T\n", x,x)
	fmt.Printf("%v, %T, %v\n", myslice, myslice, len(myslice))
	//WarningLogger.Println("warning")
	InfoLogger.Println("end of prog")
}
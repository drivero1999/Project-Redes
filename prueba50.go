package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MyRand(min,max int) int{
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min)+min
}

func MyRand2(min,max int) int{
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min)+min
}

func main(){
	ranx := MyRand(0,4)
	ranx2 := MyRand2(0,4)
	fmt.Print(ranx)
	fmt.Print(ranx2)
}
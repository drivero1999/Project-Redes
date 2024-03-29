package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

func main() {	
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()
	x:=0
	fmt.Println("Press ESC to quit")
	for {
		x+=1
		fmt.Println(x)
		char, key, err := keyboard.GetKey()
		if (err != nil) {
			panic(err)
		} else if (key == keyboard.KeyEsc) {
			break
		}
		fmt.Printf("You pressed: %q\r\n", char)
	}	
}

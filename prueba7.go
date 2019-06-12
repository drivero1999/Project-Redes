package main 

import (
    "fmt" 
    "os" 
    "os/exec" 
    "runtime" 
    "time" 
) 

var clear map[string]func() //create a map for storing clear funcs 

func init() { 
    clear = make(map[string]func()) //Initialize it 
    clear["linux"] = func() { 
     cmd := exec.Command("clear") //Linux example, its tested 
     cmd.Stdout = os.Stdout 
     cmd.Run() 
    } 
    clear["windows"] = func() { 
     cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
     cmd.Stdout = os.Stdout 
     cmd.Run() 
    } 
} 

func CallClear() { 
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc. 
    if ok { //if we defined a clear func for that platform: 
     value() //we execute it 
    } else { //unsupported platform 
     panic("Your platform is unsupported! I can't clear terminal screen :(") 
    } 
} 

func main() { 
    x:=0
    fmt.Println("I will clean the screen in 2 seconds!") 
    time.Sleep(2 * time.Second) 
    for {
	x+=1
        fmt.Println(x)
    	CallClear() 
	time.Sleep(100*time.Millisecond)
	}
    fmt.Println("I'm alone...") 
}

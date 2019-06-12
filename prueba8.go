package main




import (
	
	"fmt"
	"os"
	"os/exec" 
	"time"
	
	 

)

func main(){
	x:=0

	for{
	x+=1
        fmt.Println(x)
	fmt.Println("goaaaa")
	time.Sleep(1 * time.Millisecond)
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
	}

}

package main

import "github.com/julienroland/keyboard-termbox"
import "fmt"
import term "github.com/nsf/termbox-go"
import "os/exec"
import "os"
import "time"
 
func main() {

    for {
    running := true
    err := term.Init()
    if err != nil {
        panic(err)
    }

    defer term.Close()

    kb := termbox.New()
    kb.Bind(func() {
        fmt.Println("pressed space!")
        running = false
    }, "/")
    x:=0

    for running {
        x+=1
        fmt.Println(x)
        go func(msg string){

            kb.Poll(term.PollEvent())
        }("gola")
        //time.Sleep((1) * time.Second)
        time.Sleep(1000*time.Millisecond)
        fmt.Println("AAs")
        c := exec.Command("clear")
        c.Stdout = os.Stdout
        c.Run()
        
    }
    fmt.Println("exito")
    time.Sleep(2*time.Second)
}


}
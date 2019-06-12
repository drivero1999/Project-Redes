/*
A very simple TCP client written in Go.

This is a toy project that I used to learn the fundamentals of writing
Go code and doing some really basic network stuff.

Maybe it will be fun for you to read. It's not meant to be
particularly idiomatic, or well-written for that matter.
*/
package main




import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"time"
	//"github.com/julienroland/keyboard-termbox"
	//"os/exec" 
	 //term "github.com/nsf/termbox-go"
	term "github.com/nsf/termbox-go"

)

var host = flag.String("host", "localhost", "The hostname or IP to connect to; defaults to \"localhost\".")
var port = flag.Int("port", 8000, "The port to connect to; defaults to 8000.")
var dest1 = 1

var conn net.Conn

func reset() {
         term.Sync() // cosmestic purpose
 }
func main() {
	flag.Parse()
	reconnect:

		//var destinos [2]net.Conn
		dest1=1
		//dest := *host + ":" + strconv.Itoa(*port)
		dest := "192.168.121.9" + ":" + strconv.Itoa(*port)
		dest3 := "192.168.121.23" + ":" + strconv.Itoa(*port)
		dest2 := "168.121.9" + ":" + strconv.Itoa(*port)
		dest4 := "197.168.7.42" + ":" + strconv.Itoa(*port)
		//dest4 := "192.168.121.20" + ":" + strconv.Itoa(*port)
		dest5 := "192.168.121.16" + ":" + strconv.Itoa(*port)
		fmt.Printf("Connecting to %s...\n", dest)
        //conn
		//conn2, err := net.Dial("tcp", dest)
		conn1,err:=net.Dial("tcp",dest2)
		//conn3,err:=net.Dial("tcp",dest3)

		//fmt.Println(conn,conn1)
		
		//if conn==conn1{
			fmt.Println("AS")
		//}
		connect:=0
		//for j:=0; j<2;j++{

			if  connect==0{
				conn2, _ := net.Dial("tcp", dest)
				if conn2!=conn1{
					conn=conn2
					connect+=1
					fmt.Println("BB")		
				}
			
			fmt.Println("DD")
			fmt.Println(connect)
			}
			if connect==0{
				fmt.Println("MM")

				conn3,_:=net.Dial("tcp",dest3)
				fmt.Println("conexion")
				fmt.Println(conn3,conn1)
				if conn3!=conn1{
					conn=conn3
					fmt.Println("HH")
					//os.Exit(1)
					connect+=1
				}
				fmt.Println("VV")
				
				//os.Exit(1)
			}
			if connect==0{
				fmt.Println("MM")

				conn4,_:=net.Dial("tcp",dest4)
				fmt.Println("conexion")
				fmt.Println(conn4,conn1)
				if conn4!=conn1{
					conn=conn4
					fmt.Println("HH")
					//os.Exit(1)
					connect+=1
				}
				fmt.Println("server4")
				//os.Exit(1)
				
				//os.Exit(1)
			}
			if connect==0{
				fmt.Println("MM")

				conn5,_:=net.Dial("tcp",dest5)
				fmt.Println("conexion")
				fmt.Println(conn5,conn1)
				if conn5!=conn1{
					conn=conn5
					fmt.Println("HH")
					//os.Exit(1)
					connect+=1
				}
				fmt.Println("server4")
				//os.Exit(1)
				
				//os.Exit(1)
			}

		


		//}

		//os.Exit(2)

		//os.Exit(2)

		if err != nil {

			if _, t := err.(*net.OpError); t {
				fmt.Println("Some problem connecting.")
			} else {
				fmt.Println("Unknown error: " + err.Error())
			}
			fmt.Println("ASASAS")
			//os.Exit(1)
		}

	go readConnection(conn)

	for {
		
   	
     	//fmt.Println("gdhfh")
     	
     	if dest1 == 0{
     		fmt.Println("RECONECTANDO")
     		//os.Exit(1)
     		goto reconnect
     	}
     	// <-time.After(2*time.Second)



		
		//reader := bufio.NewReader(os.Stdin)

		//fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
		//text, _ := reader.ReadString('\n')
		////////////////


		 err := term.Init()
         if err != nil {
         		fmt.Println("HYY")
                 panic(err)
         }

         defer term.Close()

		keyPressListenerLoop:
         for {

         		if dest1==0{
				break
				}
                 switch ev := term.PollEvent(); ev.Type {
                 case term.EventKey:
                         switch ev.Key {

                         case term.KeyEsc:
                                 break keyPressListenerLoop
                         case term.KeyArrowUp:
                                 reset()
                                 fmt.Println("Arrow Up pressed")

                                 _, err = conn.Write([]byte("/u \n"))
                         case term.KeyArrowDown:
                                 reset()
                                 //fmt.Println("Arrow Down pressed")
                                 _, err = conn.Write([]byte("/d \n"))
                         case term.KeyArrowLeft:
                                 reset()
                                 //fmt.Println("Arrow Left pressed")
                                 _, err = conn.Write([]byte("/l \n"))
                         case term.KeyArrowRight:
                                 reset()
                                 //fmt.Println("Arrow Right pressed")
                                 _, err = conn.Write([]byte("/r \n"))
                         case term.KeySpace:
                                 reset()
                                 //fmt.Println("Space pressed")
                         case term.KeyBackspace:
                                 reset()
                                 //fmt.Println("Backspace pressed")
                         case term.KeyEnter:
                                 reset()
                                 //fmt.Println("Enter pressed")
                         case term.KeyTab:
                                 reset()
                                 //fmt.Println("Tab pressed")

                         default:
                                 // we only want to read a single character or one key pressed event
                                 reset()
                                 //fmt.Println("ASCII : ", ev.Ch)

                         }
                 case term.EventError:
                         panic(ev.Err)
                 }
         }



		///////////////
		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		//_, err = conn.Write([]byte(text))
		
		if err != nil {
			fmt.Println("Error writing to stream.")
			goto reconnect
		}
		
	}
}
/*
func tiempos(conn net.Conn){
	
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

    for running {
    	//fmt.Println("hola") si entra
        go func(msg string){

            kb.Poll(term.PollEvent())
            //time.Sleep(1000*time.Millisecond)
            //conn.Write([]byte("Time"+"\n"))
        }("gola")
        //fmt.Println("hola") si entra
        //time.Sleep((1) * time.Second)
        time.Sleep(1000*time.Millisecond)

		
        //fmt.Println("hola")
        conn.Write([]byte("Time"+"\n"))
        fmt.Println("Hola2")
        /*
        //c := exec.Command("clear")
        //c.Stdout = os.Stdout
        //c.Run()
    } 

	
}*/

func readConnection(conn net.Conn) {
	for {
		if dest1==0{
			break
		}
		scanner := bufio.NewScanner(conn)

		for {
			ok := scanner.Scan()
			text := scanner.Text()

			command := handleCommands(text)
			if !command && len(text)!=15 && len(text)!=0{
				
				//fmt.Printf("\b\b** %s\n> ", text)
				fmt.Println("\n")
				//fmt.Println("ASAS")
				//fmt.Println(text)
				fmt.Println("SAASA")
				fmt.Println(len(text))
				s := ""
				fmt.Printf("%c", []rune(text)[0])
				//s+=string([]rune(text)[0])
				
				for i := 1; i < 920; i++ {
					
					s+=string([]rune(text)[i])
					//fmt.Println(i)
					
					
					
					//fmt.Println("\n")
					if i%40 ==0 || i==919 {
					   fmt.Println(s)
					   s=""
					   //fmt.Println(i)
					   //fmt.Println(len(text))
					   s+=string([]rune(text)[i])
					   //i=i+1
					   //fmt.Println(i)	
					   //fmt.Printf("%c", []rune(text)[i+1])

					}
					//fmt.Println(s)
				}
			
			}else{
				//var host = flag.String("host", "localhost", "The hostname or IP to connect to; defaults to \"localhost\".")
				//var port = flag.Int("port", 8000, "The port to connect to; defaults to 8000.")
				//port=flag.Int("port", 8001, "The port to connect to; defaults to 8000.")
				fmt.Println("CHanging")

				*port=8000
				dest1=0
				break
				//goto reconnect

			}

			if !ok {
				fmt.Println("Reached EOF on server connection.")
				break
			}
		}
	}
}

func handleCommands(text string) bool {
	r, err := regexp.Compile("^%.*%$")
	if err == nil {
		if r.MatchString(text) {

			switch {
			case text == "%quit%":
				fmt.Println("\b\bServer is leaving. Hanging up.")
				os.Exit(0)
			}

			return true
		}
	}

	return false
}

/*
A very simple TCP server written in Go.

This is a toy project that I used to learn the fundamentals of writing
Go code and doing some really basic network stuff.

Maybe it will be fun for you to read. It's not meant to be
particularly idiomatic, or well-written for that matter.
*/
package main


//import "github.com/MauricioCDZ/REDES1/snake"
import (
	"bufio"
	"flag"
	"io/ioutil"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
	//"os/exec" 
	//"net/http"
	//"reflect"
	"strings"
	"math/rand"
	
	
)

var addr = flag.String("addr", "", "The address to listen to; default is \"\" (all interfaces).")
var port = flag.Int("port", 8000, "The port to listen on; default is 8000.")
/*
type MatString string
type IntMat [23][40]string
var ss = IntMat {
	[40]string{"╔","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═",
	" ","║"," ","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═","╗" }, 
	
	[40]string{"║",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","." ,".",
	".","║",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","." ,"║" },
	[40]string{"║",".","╔","═","═","═","═","═","═","╗",".","╔","═","═","═","═","═","═","═" ,"╗",
	".","║",".","╔","═","═","═","═","═","═","═","╗",".","╔","═","═","═","╗","." ,"║" },
	[40]string{"║",".","║","▓","▓","▓","▓","▓","▓","║",".","║","▓","▓","▓","▓","▓","▓","▓" ,"║",
	".","║",".","║","▓","▓","▓","▓","▓","▓","▓","║",".","║","▓","▓","▓","║","." ,"║" },
	[40]string{"║",".","╚","═","═","═","═","═","═","╝",".","╚","═","═","═","═","═","═","═" ,"╝",
	".","║",".","╚","═","═","═","═","═","═","═","╝",".","╚","═","═","═","╝","." ,"║" },
	[40]string{"║",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","." ,".",
	".","║",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","." ,"║" },
	[40]string{"║",".","═","═","═","═","═","═","═","═",".","║",".","═","═","═","═","═","═" ,"═",
	"═","╬","═","═","═","═","═","═","═","═",".","║",".","═","═","═","═","═","." ,"║" },
	[40]string{"║",".",".",".",".",".",".",".",".",".",".","║",".",".",".",".",".",".","." ,".",
	".","║",".",".",".",".",".",".",".",".",".","║",".",".",".",".",".",".","." ,"║" },
	[40]string{"╚","═","═","═","═","═","═","═","═","╗",".","╠","═","═","═","═","═","═","═" ,"═",
	".","║",".","═","═","═","═","═","═","═","═","╣",".","╔","═","═","═","═","═" ,"╝" },
	[40]string{" "," "," "," "," "," "," "," "," ","║",".","║",".",".",".",".",".",".","." ,".",
	".",".",".",".",".",".",".",".",".",".",".","║",".","║"," "," "," "," "," " ," " },
	[40]string{"═","═","═","═","═","═","═","═","═","╝",".","║",".","╔","═","═","═","═","═" ,"═",
	"═"," ","═","═","═","═","═","═","═","╗",".","║",".","╚","═","═","═","═","═" ,"═" },
	[40]string{".",".",".",".",".",".",".",".",".",".",".",".",".","║",".",".",".",".","." ,".",
	".",".",".",".",".",".",".",".",".","║",".",".",".",".",".",".",".",".","." ,"." },
	[40]string{"═","═","═","═","═","═","═","═","═","╗",".","║",".","╚","═","═","═","═","═" ,"═",
	"═"," ","═","═","═","═","═","═","═","╝",".","║",".","╔","═","═","═","═","═" ,"═" },
	[40]string{" "," "," "," "," "," "," "," "," ","║",".","║",".",".",".",".",".",".","." ,".",
	".",".",".",".",".",".",".",".",".",".",".","║",".","║"," "," "," "," "," " ," " },
	[40]string{"╔","═","═","═","═","═","═","═","═","╝",".","║",".","═","═","═","═","═","═" ,"═",
	"═","╦","═","═","═","═","═","═","═","═",".","║",".","╚","═","═","═","═","═" ,"╗" },
	[40]string{"║",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","." ,".",
	".","║",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","." ,"║" },
	[40]string{"║",".","═","═","═","═","═","═","═","╗",".","═","═","═","═","═","═","═","═" ,"═",
	".","║",".","═","═","═","═","═","═","═","═","═",".","╔","═","═","═","═","." ,"║" },
	[40]string{"║",".",".",".",".",".",".",".",".","║",".",".",".",".",".",".",".",".","." ,".",
	".",".",".",".",".",".",".",".",".",".",".",".",".","║",".",".",".",".","." ,"║" },
	[40]string{"╠","═","═","═","═","═","═","═",".","║",".","║",".","═","═","═","═","═","═" ,"═",
	"═","═","═","═","═","═","═","═","═","═",".","║",".","║",".","═","═","═","═" ,"╣" },
	[40]string{"║",".",".",".",".",".",".",".",".",".",".","║",".",".",".",".",".",".","." ,".",
	".",".",".",".",".",".",".",".",".",".",".","║",".",".",".",".",".",".","." ,"║" },
	[40]string{"║",".","═","═","═","═","═","═","═","═","═","╩","═","═","═","═","═","═","═" ,"═",
	".","║",".","═","═","═","═","═","═","═","═","╩","═","═","═","═","═","═","." ,"║" },
	[40]string{"║",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","." ,".",
	".","║",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".",".","." ,"║" },
	[40]string{"╚","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═" ,"═",
	" ","║"," ","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═","═" ,"╝" }, 
	
	}
*/
type player struct {
    posx int 
    posy  int
    symbol string 
    usado bool
    huella string
}
type fantasma struct{
	posx int
	posy int
	symbol string

}

 var player1=player{posx:1,posy:1,symbol:"A",usado:false,huella:"&"}
 var player2=player{posx:1,posy:38,symbol:"B",usado:false,huella:"#"}
 var player3=player{posx:21,posy:1,symbol:"C",usado:false,huella:"*"}
 var player4=player{posx:21,posy:38,symbol:"D",usado:false,huella:"$"}
 //var fantasma1=fantasma{posx:11,posy: 15,symbol: "@"} 
 var fantasma1=fantasma{posx: 9,posy: 21,symbol: "O"} 
 var fantasma2=fantasma{posx: 13,posy: 21,symbol: "K"} 

 //var diccionario=make(map[string]*player)
 type diccionario1 map[string]*player
 var diccionario=make(diccionario1)
 //var leido=1

 var lista [4]player
 var clientes[4]string

 var actual=""
 var tumbarServer=false
 var folas=1


//RECONECTAR EL SERVIDOR
var ss = [23][40]string{}

func check(e error) {
    if e != nil {
        panic(e)
    }
}



func main() {




	

	/////////////////////////////////////////////
	flag.Parse()
	

	fmt.Println("Starting server...")

	src := *addr + ":" + strconv.Itoa(*port)

	listener, _ := net.Listen("tcp", src)
	fmt.Printf("Listening on %s.\n", src)

	defer listener.Close()


	
	//fmt.Println(Matrix2String(ss))
	
	//player1:=player{posx:2,posy:1,symbol:"A"}
	//movePlayer(player1)
	//movePlayer(player2)
	/*
	doGame1()
		lista[0]=player1
 		lista[1]=player2
 		*/
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}
		
		
		go startPolling2(conn)
		go handleConnection(conn)
		
		
		//
		//conn.Write([]byte(Matrix2String(ss)+"\n"))
	}
	//fmt.Println(person{"Bob", 20})
	//fmt.Println(s.name)
	//s := person{name: "Sean", age: 50}
	//sp.age = 51


	
}
//matrix [23][40]string
func Matrix2String(t [23][40]string) string {
	s := ""

	for i := range t {
		for _, n := range t[i] {
			s += fmt.Sprintf("%s", n)
			
		}
		
		//fmt.Println(s)
			
		//s = ""
		
		
	}
	//fmt.Println(s)
	return s
}

func movePlayer(t player){
	ss[t.posx][t.posy]=t.symbol
	//coloca en el tablero el jugador
}

func moveDown(t *player)int{
	var num=0
	var ascii=string(ss[t.posx+1][t.posy])
	var cont=0
	if  t.posx==21 && t.posy==20{
		ss[t.posx][t.posy]=t.huella
		t.posx=1
		t.posy=20
		ss[t.posx][t.posy]=t.symbol
		cont=1
		num-=20
	} 
	if  t.posx==21 && t.posy==22{
		ss[t.posx][t.posy]=t.huella
		t.posx=1
		t.posy=22
		ss[t.posx][t.posy]=t.symbol
		cont=1
		num-=20
	}
	if ascii != "═" && ascii!= "║" && ascii != "╔" && ascii != "╗" && ascii != "╝" && ascii != "╚" && cont==0{
		//fmt.Println("entre",ss[t.posx+1][t.posy])
		ss[t.posx][t.posy]=t.huella
		t.posx+=1
		ss[t.posx][t.posy]=t.symbol
		//fmt.Println("contador")
		num+=1
	}
	//fmt.Println("num es: ")
	//fmt.Println(num)
	return int(num)
}

func moveUp(t *player)int{
	var num=0
	var ascii=string(ss[t.posx-1][t.posy])
	var cont=0
	if  t.posx==1 && t.posy==20{
		ss[t.posx][t.posy]=t.huella
		t.posx=21
		t.posy=20
		ss[t.posx][t.posy]=t.symbol
		cont=1
		num+=20
	} 
	if  t.posx==1 && t.posy==22{
		ss[t.posx][t.posy]=t.huella
		t.posx=21
		t.posy=22
		ss[t.posx][t.posy]=t.symbol
		cont=1
		num+=20
	}
	if ascii != "═" && ascii!= "║" && ascii != "╔" && ascii != "╗" && ascii != "╝" && ascii != "╚" && cont==0{
		ss[t.posx][t.posy]=t.huella
		t.posx-=1
		ss[t.posx][t.posy]=t.symbol
		num-=1
	}
	//fmt.Println("num es: ")
	//fmt.Println(num)	
	return int(num)
}

func moveRight(t *player)int{
	var num=0
	var ascii=string(ss[t.posx][t.posy+1])
	var cont=0
	if  t.posx==11 && t.posy==38{
		ss[t.posx][t.posy]=t.huella
		t.posx=11
		t.posy=1
		ss[t.posx][t.posy]=t.symbol
		num-=37
		cont=1
	}
	if ascii != "═" && ascii!= "║" && ascii != "╔" && ascii != "╗" && ascii != "╝" && ascii != "╚" && cont==0{
		ss[t.posx][t.posy]=t.huella
		t.posy+=1
		ss[t.posx][t.posy]=t.symbol
		num+=1

	}	
	//fmt.Println("num es: ")
	//fmt.Println(num)
	return int(num)
}

func moveLeft(t *player)int{
	var num=0
	var ascii=string(ss[t.posx][t.posy-1])
	var cont=0
	if  t.posx==11 && t.posy==1{
		ss[t.posx][t.posy]=t.huella
		t.posx=11
		t.posy=38
		ss[t.posx][t.posy]=t.symbol
		num+=37
		cont=1
	}
	if ascii != "═" && ascii!= "║" && ascii != "╔" && ascii != "╗" && ascii != "╝" && ascii != "╚" && cont==0{
		ss[t.posx][t.posy]=t.huella
		t.posy-=1
		ss[t.posx][t.posy]=t.symbol
		num-=1
	}
	//fmt.Println("num es: ")
	//fmt.Println(num)
	return int(num)
}

func moveFantasma(t fantasma){
	ss[t.posx][t.posy]=t.symbol
	//coloca en el tablero el fantasma
}

func moveDown2(t fantasma)int{
	//fmt.Println("entro a funcion down2")
	var num=0
	var ascii=string(ss[t.posx+1][t.posy])
	var cont=0
	if  t.posx==21 && t.posy==20{
		ss[t.posx][t.posy]="."
		t.posx=1
		t.posy=20
		ss[t.posx][t.posy]=t.symbol
		cont=1
		num-=20
	} 
	if  t.posx==21 && t.posy==22{
		ss[t.posx][t.posy]="."
		t.posx=1
		t.posy=22
		ss[t.posx][t.posy]=t.symbol
		cont=1
		num-=20
	}
	if ascii != "═" && ascii!= "║" && ascii != "╔" && ascii != "╗" && ascii != "╝" && ascii != "╚" && cont==0{
		//fmt.Println("entre",ss[t.posx+1][t.posy])
		ss[t.posx][t.posy]="."
		t.posx+=1
		ss[t.posx][t.posy]=t.symbol
		//fmt.Println("contador")
		num+=1
	}
	//fmt.Println("num es: ")
	//fmt.Println(num)
	return int(num)
}

func moveUp2(t fantasma)int{
	//fmt.Println("entro a funcion up2")
	var num=0
	var ascii=string(ss[t.posx-1][t.posy])
	var cont=0
	if  t.posx==1 && t.posy==20{
		ss[t.posx][t.posy]="."
		t.posx=21
		t.posy=20
		ss[t.posx][t.posy]=t.symbol
		cont=1
		num+=20
	} 
	if  t.posx==1 && t.posy==22{
		ss[t.posx][t.posy]="."
		t.posx=21
		t.posy=22
		ss[t.posx][t.posy]=t.symbol
		cont=1
		num+=20
	}
	if ascii != "═" && ascii!= "║" && ascii != "╔" && ascii != "╗" && ascii != "╝" && ascii != "╚" && cont==0{
		ss[t.posx][t.posy]="."
		t.posx-=1
		ss[t.posx][t.posy]=t.symbol
		num-=1
	}
	//fmt.Println("num es: ")
	//fmt.Println(num)	
	return int(num)
}

func moveRight2(t fantasma)int{
	//fmt.Println("entro a funcion right2")
	var num=0
	var ascii=string(ss[t.posx][t.posy+1])
	var cont=0
	if  t.posx==11 && t.posy==38{
		ss[t.posx][t.posy]="."
		t.posx=11
		t.posy=1
		ss[t.posx][t.posy]=t.symbol
		num-=37
		cont=1
	}
	if ascii != "═" && ascii!= "║" && ascii != "╔" && ascii != "╗" && ascii != "╝" && ascii != "╚" && cont==0{
		ss[t.posx][t.posy]="."
		t.posy+=1
		ss[t.posx][t.posy]=t.symbol
		num+=1

	}	
	//fmt.Println("num es: ")
	//fmt.Println(num)
	return int(num)
}
func moveLeft2(t fantasma)int{
	//fmt.Println("entro a funcion left2")
	var num=0
	var ascii=string(ss[t.posx][t.posy-1])
	var cont=0
	if  t.posx==11 && t.posy==1{
		ss[t.posx][t.posy]="."
		t.posx=11
		t.posy=38
		ss[t.posx][t.posy]=t.symbol
		num+=37
		cont=1
	}
	if ascii != "═" && ascii!= "║" && ascii != "╔" && ascii != "╗" && ascii != "╝" && ascii != "╚" && cont==0{
		ss[t.posx][t.posy]="."
		t.posy-=1
		ss[t.posx][t.posy]=t.symbol
		num-=1
	}
	//fmt.Println("num es: ")
	//fmt.Println(num)
	return int(num)
}

func MyRand(min,max int) int{
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min)+min
}


func startPolling2(conn net.Conn) {

  for {
  	
    //<-timife.After(2 * time.Second)
    <-time.After(2*time.Second)
    conn.Write([]byte(Matrix2String(ss)+"\n"))
    //go doSomething("from polling 2")
    //c := exec.Command("clear")
      //  c.Stdout = os.Stdout
       // c.Run()

	for i:=0; i<5; i++{
	    //FANTASMAS=================================================================================
		tipoMove:= MyRand(0,4)
		//fmt.Println("el numero random es: " , tipoMove)
		if tipoMove==0{
			//fmt.Println("entro if 0")
		    var cord=moveUp2(fantasma1)
		    fantasma1.posx+=cord

		    var cord2=moveLeft2(fantasma2)
		    fantasma2.posy+=cord2
		    //fmt.Println("coordenadas fantasma1: " , fantasma1.posx , fantasma1.posy)
		    //conn.Write([]byte(Matrix2String(ss)+"\n"))
			d1 := []byte(Matrix2String(ss)+strconv.Itoa(player1.posx)+"*"+strconv.Itoa(player1.posy)+"|"+strconv.Itoa(player2.posx)+"*"+strconv.Itoa(player2.posy)+"|"+strconv.Itoa(player3.posx)+"*"+strconv.Itoa(player3.posy)+"|"+strconv.Itoa(player4.posx)+"*"+strconv.Itoa(player4.posy)+"|"+strconv.Itoa(fantasma1.posx)+"*"+strconv.Itoa(fantasma1.posy)+"|"+strconv.Itoa(fantasma2.posx)+"*"+strconv.Itoa(fantasma2.posy)+"|"+"                    .")

    		ioutil.WriteFile("C:/Users/david/Desktop/proyectoRedes/file2.txt", d1, 0644)		    
		}
		if tipoMove==1{
			//fmt.Println("entro if 1")
			var cord=moveDown2(fantasma1)
			fantasma1.posx+=cord

			var cord2=moveUp2(fantasma2)
		    fantasma2.posx+=cord2
			//fmt.Println("coordenadas fantasma1: " , fantasma1.posx , fantasma1.posy)
			//conn.Write([]byte(Matrix2String(ss)+"\n"))
			d1 := []byte(Matrix2String(ss)+strconv.Itoa(player1.posx)+"*"+strconv.Itoa(player1.posy)+"|"+strconv.Itoa(player2.posx)+"*"+strconv.Itoa(player2.posy)+"|"+strconv.Itoa(player3.posx)+"*"+strconv.Itoa(player3.posy)+"|"+strconv.Itoa(player4.posx)+"*"+strconv.Itoa(player4.posy)+"|"+strconv.Itoa(fantasma1.posx)+"*"+strconv.Itoa(fantasma1.posy)+"|"+strconv.Itoa(fantasma2.posx)+"*"+strconv.Itoa(fantasma2.posy)+"|"+"                    .")

    		ioutil.WriteFile("C:/Users/david/Desktop/proyectoRedes/file2.txt", d1, 0644)
		}
		if tipoMove==2{
			//fmt.Println("entro if 2")
		    var cord=moveRight2(fantasma1)
		    fantasma1.posy+=cord
		    //fmt.Println("coordenadas fantasma1: " , fantasma1.posx , fantasma1.posy)
		    //conn.Write([]byte(Matrix2String(ss)+"\n"))
		    var cord2=moveUp2(fantasma2)
		    fantasma2.posx+=cord2
			d1 := []byte(Matrix2String(ss)+strconv.Itoa(player1.posx)+"*"+strconv.Itoa(player1.posy)+"|"+strconv.Itoa(player2.posx)+"*"+strconv.Itoa(player2.posy)+"|"+strconv.Itoa(player3.posx)+"*"+strconv.Itoa(player3.posy)+"|"+strconv.Itoa(player4.posx)+"*"+strconv.Itoa(player4.posy)+"|"+strconv.Itoa(fantasma1.posx)+"*"+strconv.Itoa(fantasma1.posy)+"|"+strconv.Itoa(fantasma2.posx)+"*"+strconv.Itoa(fantasma2.posy)+"|"+"                    .")

    		ioutil.WriteFile("C:/Users/david/Desktop/proyectoRedes/file2.txt", d1, 0644)
		}
		if tipoMove==3{
			//fmt.Println("entro if 3")
		    var cord=moveLeft2(fantasma1)
		    fantasma1.posy+=cord
		    //fmt.Println("coordenadas fantasma1: " , fantasma1.posx , fantasma1.posy)
		    //conn.Write([]byte(Matrix2String(ss)+"\n"))
		    var cord2=moveRight2(fantasma2)
		    fantasma2.posy+=cord2
			d1 := []byte(Matrix2String(ss)+strconv.Itoa(player1.posx)+"*"+strconv.Itoa(player1.posy)+"|"+strconv.Itoa(player2.posx)+"*"+strconv.Itoa(player2.posy)+"|"+strconv.Itoa(player3.posx)+"*"+strconv.Itoa(player3.posy)+"|"+strconv.Itoa(player4.posx)+"*"+strconv.Itoa(player4.posy)+"|"+strconv.Itoa(fantasma1.posx)+"*"+strconv.Itoa(fantasma1.posy)+"|"+strconv.Itoa(fantasma2.posx)+"*"+strconv.Itoa(fantasma2.posy)+"|"+"                    .")

    		ioutil.WriteFile("C:/Users/david/Desktop/proyectoRedes/file2.txt", d1, 0644)
		}

		//CHOUQE=========================================================
		if player1.posx == fantasma1.posx && player1.posy == fantasma1.posy {
			conn.Close()
			fmt.Println("jugador 1 murio")
		}
		if player1.posx == fantasma2.posx && player1.posy == fantasma2.posy {
			fmt.Println("jugador 1 murio")
			conn.Close()
		}
		if player2.posx == fantasma1.posx && player2.posy == fantasma1.posy {
			fmt.Println("jugador 2 murio")
			conn.Close()
		}
		if player2.posx == fantasma2.posx && player2.posy == fantasma2.posy {
			fmt.Println("jugador 2 murio")
			conn.Close()
		}
		if player3.posx == fantasma1.posx && player3.posy == fantasma1.posy {
			conn.Close()
			fmt.Println("jugador 3 murio")
		}
		if player3.posx == fantasma2.posx && player3.posy == fantasma2.posy {
			conn.Close()
			fmt.Println("jugador 3 murio")
		}
		if player4.posx == fantasma1.posx && player4.posy == fantasma1.posy {
			conn.Close()
			fmt.Println("jugador 4 murio")
		}
		if player4.posx == fantasma2.posx && player4.posy == fantasma2.posy {
			conn.Close()
			fmt.Println("jugador 4 murio")
		}
	}

    if tumbarServer==true{
    	//debe enviarle que hacer al cliente para que se conecte al otro servidor
    	//strconv.Itoa(-42)
    	//ss[player1.posx][player1.posx]=","
    	d1 := []byte(Matrix2String(ss)+strconv.Itoa(diccionario[actual].posx)+strconv.Itoa(diccionario[actual].posy)+strconv.Itoa(fantasma1.posx)+strconv.Itoa(fantasma1.posy)+strconv.Itoa(fantasma2.posx)+strconv.Itoa(fantasma2.posy))
    	err := ioutil.WriteFile("C:/Users/david/Desktop/proyectoRedes/file2.txt", d1, 0644)
    	check(err)
    	//
    	conn.Write([]byte(conn.RemoteAddr().String()))
    	conn.Close()	
    }
   
    
  }
}

var pr=0

func handleConnection(conn net.Conn) {


	
	remoteAddr := conn.RemoteAddr().String()


	fmt.Println("Client connected from " + remoteAddr)
	if pr==0{
		doGame1()
		fmt.Println("Coordendas p1:",player1.posx,player1.posy)
		lista[0]=player1
 		lista[1]=player2
 		lista[2]=player3
 		lista[3]=player4


		pr=1
	}

	/////////////////////////////////////
	


	/////////////////////////////////////





	gg:= conn.RemoteAddr() //SO IMPORTANT
	fmt.Println(len(gg.String()))

	scanner := bufio.NewScanner(conn)
	

	//fmt.Println("direccion x del conn",scanner)
	

	//conn.Write([]byte(Matrix2String(ss)+"\n"))
	for {
		

		ok := scanner.Scan()
		//fmt.Println("search",scanner)
		//var valScaner= Println("search",scanner)

		var valScaner=fmt.Sprintf("foo: %s", scanner)
		//fmt.Println(scanner)
		//i:=strings.Index(valScaner,"0xc")
		//fmt.Println("Index: ",i)

		fmt.Println("after: ",string(valScaner[29:39]))
		//fmt.Println("dicci",reflect.TypeOf(diccionario["ss"]))
		if diccionario["ss"]!=diccionario[string(valScaner[29:39])] {
			fmt.Println("")

		} else{

			for i:=0; i < 4;i++{
				if lista[i].usado==false{
					lista[i].usado=true
					diccionario[string(valScaner[29:39])]=&lista[i]
					//fmt.Println("has")
					break
				}
			}
		}
		actual=valScaner[29:39]
		/*
		if leido==1{ //llamo al hacer el tablero
			fmt.Println("ahablame")
			doGame()

		} */



		//fmt.Println("map :",diccionario)
		if !ok {
			break

		}
		folas=0
		handleMessage(scanner.Text(), conn)
		
	}

	fmt.Println("Client at " + remoteAddr + " disconnected.")
}
func doGame1(){
		dat, err := ioutil.ReadFile("C:/Users/david/Desktop/proyectoRedes/file2.txt")
    check(err)
    mapa:=string(dat)
    s:=bufio.NewScanner(strings.NewReader(mapa))
    //fmt.Printf("%s", placeOfInterest)
    ok:=s.Scan()
    if !ok {
                fmt.Println("Reached EOF on server connection.")
                
            }
    text:=s.Text()

    k:=0
    j:=0
    fmt.Println()
    ss[0][0]=string([]rune(text)[0])
    for i := 1; i < 920; i++ {
                    
                    //l+=string([]rune(text)[i])

                    //fmt.Println(string([]rune(text)[i]))
                    //fmt.Println(j)
                    //fmt.Println(k)

                    ss[k][j] = string([]rune(text)[i-1])
                    //fmt.Println(k)
                    j++
                    
                    //fmt.Println("\n")
                    if i%40 ==0 || i==919 {
                       k++
                       j=0
                       //fmt.Println("Entre",k," i:",i)
                       //fmt.Println(string([]rune(text)[i-1]))
                       

                    }
                    //fmt.Println(s)
                }
     ss[22][39]=string([]rune(text)[919])

     dat, err = ioutil.ReadFile("C:/Users/david/Desktop/proyectoRedes/file2.txt")
    check(err)
    mapa=string(dat)
    s=bufio.NewScanner(strings.NewReader(mapa))
    //fmt.Printf("%s", placeOfInterest)
    ok=s.Scan()
    if !ok {
                fmt.Println("Reached EOF on server connection.")
                
            }
    text=s.Text()

 
    fmt.Println()
    
    fmt.Println(strconv.Atoi(string([]rune(text)[920])))
    // fmt.Println(strconv.Atoi(string([]rune(text)[921])))

    //c, err := strconv.Atoi(string([]rune(text)[920]))
    //f,err:=  strconv.Atoi(string([]rune(text)[921])) 
    c:=""
    f:=""
    asterisco:=0
    coord:=0 //El separador de coordenadas
    for i := 920; i < 961; i++ {
    	xx:=string([]rune(text)[i])
    	xs:=strings.Replace(xx," ","",-1)
    	xl:=strings.Replace(xs,"*","",-1)
    	if coord==0 && string([]rune(text)[i])!="*" {
    		c+=xl
    		fmt.Println(c)

    	}
    	if string([]rune(text)[i])=="*"{
    		asterisco+=1
    		coord=1
    		//continue 
    	}
    	if coord==1 && string([]rune(text)[i])!="*"{
    		fmt.Println("soy xl",xl)
    		f+=xl
    		fmt.Println(f)
    	}
    	if string([]rune(text)[i])=="|" {
    		vb:=strings.Replace(c,"|","",-1)
    		vbl:=strings.Replace(f,"|","",-1)

    		c1,_:=strconv.Atoi(vb)
    		f1,_:=strconv.Atoi(vbl)
    		fmt.Println("AS",c1,"DA",f1)
    		//check(err)
    		if asterisco==1{
    			player1.posx=c1
    			player1.posy=f1
    			fmt.Println("c1:",c1,"f1:",f1)
    			fmt.Println("mod player1:",player1.posx,player1.posy)

    		}
    		if asterisco==2{
    			player2.posx=c1
    			player2.posy=f1
    			fmt.Println("c1:",c1,"f1:",f1)
    			fmt.Println("mod player2:",player2.posx,player2.posy)
    		}
    		if asterisco==3{
    			player3.posx=c1
    			player3.posy=f1
    			fmt.Println("c1:",c1,"f1:",f1)
    			fmt.Println("mod player3:",player3.posx,player3.posy)
    		}
    		if asterisco==4{
    			player4.posx=c1
    			player4.posy=f1
    			fmt.Println("c1:",c1,"f1:",f1)
    			fmt.Println("mod player4:",player4.posx,player4.posy)
    		}
    		if asterisco==5{
    			fantasma1.posx=c1
    			fantasma1.posy=f1
    			fmt.Println("c1:",c1,"f1:",f1)
    			fmt.Println("mod fantasma1:",fantasma1.posx,fantasma1.posy)
    		}
    		if asterisco==6{
    			fantasma2.posx=c1
    			fantasma2.posy=f1
    			fmt.Println("c1:",c1,"f1:",f1)
    			fmt.Println("mod fantasma2:",fantasma2.posx,fantasma2.posy)
    		}
    		c=""
    		f=""
    		coord=0
    	}
    }
    fmt.Println(c,f)    
    ss[22][39]=string([]rune(text)[919])
    //c1,err:=strconv.Atoi(c)
    //f1,err:=strconv.Atoi(f)
    //diccionario[actual].posx=c1
    //diccionario[actual].posy=f1


    fmt.Println("POSICIONES:",player1.posx,player1.posy)
    ss[player1.posx][player1.posy]=player1.symbol
    ss[player2.posx][player2.posy]=player2.symbol
    ss[player3.posx][player3.posy]=player3.symbol
    ss[player4.posx][player4.posy]=player4.symbol

    ss[fantasma1.posx][fantasma1.posy]=fantasma1.symbol
    ss[fantasma2.posx][fantasma2.posy]=fantasma2.symbol

    


}
/*
func doGame(){

		dat, err := ioutil.ReadFile("C:/Users/david/Desktop/proyectoRedes/file2.txt")
    check(err)
    mapa:=string(dat)
    s:=bufio.NewScanner(strings.NewReader(mapa))
    //fmt.Printf("%s", placeOfInterest)
    ok:=s.Scan()
    if !ok {
                fmt.Println("Reached EOF on server connection.")
                
            }
    text:=s.Text()

 
    fmt.Println()
    
    fmt.Println(strconv.Atoi(string([]rune(text)[920])))
     fmt.Println(strconv.Atoi(string([]rune(text)[921])))

    //c, err := strconv.Atoi(string([]rune(text)[920]))
    //f,err:=  strconv.Atoi(string([]rune(text)[921])) 
    c:=""
    f:=""
    asterisco:=0
    coord:=0 //El separador de coordenadas
    for i := 920; i < 930; i++ {
    	xx:=string([]rune(text)[i])
    	xs:=strings.Replace(xx," ","",-1)
    	xl:=strings.Replace(xs,"*","",-1)
    	if coord==0{
    		c+=xl
    		fmt.Println(c)

    	}
    	if string([]rune(text)[i])=="*"{
    		asterisco+=1
    		coord=1

    	}
    	if coord==1{
    		f+=xl
    		fmt.Println(f)
    	}
    	if string([]rune(text)[i])=="|"{
    		c1,err:=strconv.Atoi(c)
    		f1,err:=strconv.Atoi(f)

    		if asterisco==1{
    			player1.posx=c1
    			player1.posy=f1
    		}
    		if asterisco==2{
    			player2.posx=c1
    			player2.posy=f1
    		}
    		c=""
    		f=""
    		coord=0
    	}

    }
    fmt.Println(c,f)    
    ss[22][39]=string([]rune(text)[919])
    //c1,err:=strconv.Atoi(c)
    //f1,err:=strconv.Atoi(f)
    //diccionario[actual].posx=c1
    //diccionario[actual].posy=f1


    //fmt.Println("POSICIONES:",diccionario[actual].posx,diccionario[actual].posy)
    ss[player1.posx][player1.posy]=player1.symbol
    ss[player2.posx][player2.posy]=player2.symbol

    

	


}
*/
func handleMessage(message string, conn net.Conn) {
	fmt.Println("> " + message)

	if len(message) > 0 && message[0] == '/' {
		switch {
		case message == "/time":
			//resp := "It is " + time.Now().String() + "\n"
			resp := "It is " + time.Now().String()+"\n"

			fmt.Print("< " + resp)
			conn.Write([]byte(resp))

		case message == "/quit":
			fmt.Println("Quitting.")
			conn.Write([]byte("I'm shutting down now.\n"))
			fmt.Println("< " + "%quit%")
			conn.Write([]byte("%quit%\n"))
			os.Exit(0)
		case message == "/iniciarjuego":
			//snake.NewGame().Start()
			var tt= Matrix2String(ss)
			
			//resp := "\n"
			//conn.Write([]byte(resp))
			conn.Write([]byte(tt+"\n"))
		case message[1] =='d': //jugador abajo
			   
		

			var x=moveDown(diccionario[actual])
			//diccionario[actual].posx+=x
			if diccionario[actual].symbol=="A" {
				player1.posx+=x
			}
			if diccionario[actual].symbol=="B" {
				player2.posx+=x
			}
			/////
			d1 := []byte(Matrix2String(ss)+strconv.Itoa(player1.posx)+"*"+strconv.Itoa(player1.posy)+"|"+strconv.Itoa(player2.posx)+"*"+strconv.Itoa(player2.posy)+"|"+strconv.Itoa(player3.posx)+"*"+strconv.Itoa(player3.posy)+"|"+strconv.Itoa(player4.posx)+"*"+strconv.Itoa(player4.posy)+"|"+strconv.Itoa(fantasma1.posx)+"*"+strconv.Itoa(fantasma1.posy)+"|"+strconv.Itoa(fantasma2.posx)+"*"+strconv.Itoa(fantasma2.posy)+"|"+"                    .")

    		ioutil.WriteFile("C:/Users/david/Desktop/proyectoRedes/file2.txt", d1, 0644)
    		
    	////
			fmt.Println(diccionario[actual].posx,diccionario[actual].posy)	
			conn.Write([]byte(Matrix2String(ss)+"\n"))
		case message[1] == 'u':
			var x=moveUp(diccionario[actual])
			//diccionario[actual].posx+=x

			if diccionario[actual].symbol=="A" {

				player1.posx+=x
				fmt.Println("el valor de x es:",x)
				fmt.Println(player1.posy)
			}
			if diccionario[actual].symbol=="B" {
				player2.posx+=x

			}
			fmt.Println("antes de d1",player1.posy)
			d1 := []byte(Matrix2String(ss)+strconv.Itoa(player1.posx)+"*"+strconv.Itoa(player1.posy)+"|"+strconv.Itoa(player2.posx)+"*"+strconv.Itoa(player2.posy)+"|"+strconv.Itoa(player3.posx)+"*"+strconv.Itoa(player3.posy)+"|"+strconv.Itoa(player4.posx)+"*"+strconv.Itoa(player4.posy)+"|"+strconv.Itoa(fantasma1.posx)+"*"+strconv.Itoa(fantasma1.posy)+"|"+strconv.Itoa(fantasma2.posx)+"*"+strconv.Itoa(fantasma2.posy)+"|"+"                    .")


    		ioutil.WriteFile("C:/Users/david/Desktop/proyectoRedes/file2.txt", d1, 0644)
    		//check(err)
			///
			fmt.Println(diccionario[actual].posx,diccionario[actual].posy)
			conn.Write([]byte(Matrix2String(ss)+"\n"))
		case message[1] == 'r':
			var x=moveRight(diccionario[actual])
			//diccionario[actual].posy+=x
			if diccionario[actual].symbol=="A" {
				player1.posy+=x
			}
			if diccionario[actual].symbol=="B" {
				player2.posy+=x
			}
			///
			d1 := []byte(Matrix2String(ss)+strconv.Itoa(player1.posx)+"*"+strconv.Itoa(player1.posy)+"|"+strconv.Itoa(player2.posx)+"*"+strconv.Itoa(player2.posy)+"|"+strconv.Itoa(player3.posx)+"*"+strconv.Itoa(player3.posy)+"|"+strconv.Itoa(player4.posx)+"*"+strconv.Itoa(player4.posy)+"|"+strconv.Itoa(fantasma1.posx)+"*"+strconv.Itoa(fantasma1.posy)+"|"+strconv.Itoa(fantasma2.posx)+"*"+strconv.Itoa(fantasma2.posy)+"|"+"                    .")

    		ioutil.WriteFile("C:/Users/david/Desktop/proyectoRedes/file2.txt", d1, 0644)
    		//check(err)
			///
			fmt.Println(diccionario[actual].posx,diccionario[actual].posy)
			conn.Write([]byte(Matrix2String(ss)+"\n"))

		case message[1]== 'l':
			var x=moveLeft(diccionario[actual])
			//diccionario[actual].posx+=x
			if diccionario[actual].symbol=="A" {
				player1.posy+=x
			}
			if diccionario[actual].symbol=="B" {
				player2.posy+=x
			}
			///
			d1 := []byte(Matrix2String(ss)+strconv.Itoa(player1.posx)+"*"+strconv.Itoa(player1.posy)+"|"+strconv.Itoa(player2.posx)+"*"+strconv.Itoa(player2.posy)+"|"+strconv.Itoa(player3.posx)+"*"+strconv.Itoa(player3.posy)+"|"+strconv.Itoa(player4.posx)+"*"+strconv.Itoa(player4.posy)+"|"+strconv.Itoa(fantasma1.posx)+"*"+strconv.Itoa(fantasma1.posy)+"|"+strconv.Itoa(fantasma2.posx)+"*"+strconv.Itoa(fantasma2.posy)+"|"+"                    .")

    		ioutil.WriteFile("C:/Users/david/Desktop/proyectoRedes/file2.txt", d1, 0644)
    		//check(err)
			///
			fmt.Println(diccionario[actual].posx,diccionario[actual].posy)
			conn.Write([]byte(Matrix2String(ss)+"\n"))

			
		case message== "/t":
			fmt.Println("HEMOS DETECTADO FALLAS EN EL SERVIDOR, SERA DIRECCIONADO A OTRO")
			 time.Sleep(2*time.Second)
			 tumbarServer=true
			 //os.Exit(3)
		default:
			//conn.Write([]byte("Unrecognized command.\n"))
			conn.Write([]byte(Matrix2String(ss)+"\n"))

		}

	}else{
		fmt.Println("TTIme")
		conn.Write([]byte(Matrix2String(ss)+"\n"))

	}

}

// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "strings"
    //"os"
)

var matrix1 = [23][40]string{}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func printMatrix(matrix [23][40]string ){
    fmt.Println("SS")
    
    for i:=0; i < 23; i++{
        s:=""
        for j:=0; j<40;j++{
            s+= matrix[i][j]
        }
        fmt.Println(s)
        s=""


    }
    fmt.Println(matrix[11][39])

}
func main() {
    /*
    // To start, here's how to dump a string (or just
    // bytes) into a file.
    d1 := []byte("hellogo\n")
    err := ioutil.WriteFile("/home/mauricio/go/src/github.com/redes/file2.txt", d1, 0644)
    check(err) */
    


    //REead

    // Perhaps the most basic file reading task is
    // slurping a file's entire contents into memory.
    dat, err := ioutil.ReadFile("/home/mauricio/go/src/github.com/redes/file2.txt")
    check(err)
    //fmt.Print(string(dat))
    mapa:=string(dat)

    /*func main() {
    in := `{"sample":"json string"}`

    s := bufio.NewScanner(strings.NewReader(in))
    s.Split(bufio.ScanRunes)

    for s.Scan() {
        fmt.Println(s.Text())
    }
    }*/
    s:=bufio.NewScanner(strings.NewReader(mapa))
    /*for s.Scan(){
        fmt.Println(s.Text())
    }*/
    //[]rune(text)[i]
    //fmt.Printf("%s", placeOfInterest)
    ok:=s.Scan()
    if !ok {
                fmt.Println("Reached EOF on server connection.")
                
            }
    text:=s.Text()
    fmt.Println(len(text))
    
    //l := ""
    //fmt.Printf("%c", []rune(text)[0])
    //fmt.Println(string([]rune(text)[430]))
     //matrix1[i][j] = valInt
    k:=0
    j:=0
    //j++
    matrix1[0][0]=string([]rune(text)[0])
    for i := 1; i < 920; i++ {
                    
                    //l+=string([]rune(text)[i])

                    //fmt.Println(string([]rune(text)[i]))
                    //fmt.Println(j)
                    //fmt.Println(k)

                    matrix1[k][j] = string([]rune(text)[i-1])
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
    matrix1[22][39]=string([]rune(text)[919])
    printMatrix(matrix1)
    

    // You'll often want more control over how and what
    // parts of a file are read. For these tasks, start
    // by `Open`ing a file to obtain an `os.File` value.
    /*
    f, err := os.Open("/home/mauricio/go/src/github.com/redes/file2.txt")
    check(err)

    // Read some bytes from the beginning of the file.
    // Allow up to 5 to be read but also note how many
    // actually were read.
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))
    */

}
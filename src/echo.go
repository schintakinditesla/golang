/** Implementing Echo to print its command-line arguments make sure input is entered during go run <filename> <input> just like in main(int argc, char** argv) command line input.**/

package main

import(
	"fmt"
	"os"
)

func main(){
	var s, sep string
	for i := 1; i < len(os.Args); i++{
	    s += sep + os.Args[i]
	    sep = " "
    }
    fmt.Println(s)
}

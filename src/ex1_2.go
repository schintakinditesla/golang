package main
import(
	"fmt"
	"os"
)

func main(){
	for i, arg := range os.Args[1:]{
		fmt.Println(i, arg)
	}
}

// exercise 1.2 solution

package main
import(
	"fmt"
	"os"
	"strings"
)

func main(){
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//Join is a keyword used when it is costly to append strings i.e. when the data is large. This is more or less avoiding too many lines of code

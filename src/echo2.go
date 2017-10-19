package main
import(
	"fmt"
	"os"
)

func main(){
	s,sep := "",""    
	/* other ways of declaring variables are 
	1. s := ""  //short variable declaration only used within function not for packages	
	2. var s string	  //default initialization to zero for strings
	3. var s = ""	 //used while declaring multiple variables
	4. var s string = "" //explicit about variables type.
	*/

	for _,arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

// range is a keyword that produces pair of values 1. index and 2. the value of element at that inded.
// underscore or blank identifier (_) is used to when a syntax requires a variable but program logic does not.

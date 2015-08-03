package cmd

//list command
/**
Possibilities:
	-a : do not ignore entries starting with .

*/
import (
	"fmt"
	"os"
)

func Execute() {
	fmt.Println("Execute ls command")
	parse(os.Args[1:], LS_SUPPORT_OPTS)
}

package ls

//list command
/**
Possibilities:
	-a : do not ignore entries starting with .

*/
import (
	"fmt"
	"os"
)

func exploreFolder(dirPath string) {
	files, _ := ioutil.ReadDir(dirPath)
	fmt.Println(dirPath, files)
	for i := range files {
		file := files[i]

		if file.IsDir() {
			fmt.Println("Dir :: ", file.Name())
			exploreFolder(dirPath + "/" + file.Name())
		} else {
			fmt.Println("\t", file.Name())
		}
	}
}

func Execute() {
	fmt.Println("Execute ls command")
	parse(os.Args[1:], LS_SUPPORT_OPTS)
}

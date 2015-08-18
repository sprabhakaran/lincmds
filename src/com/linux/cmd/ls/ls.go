package ls

//list command
/**
Possibilities:
	-a : do not ignore entries starting with .

*/
import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	opts_arr, err := parse(os.Args[1:], LS_SUPPORT_OPTS)
	if err != nil {
		errors.New(err.Error())
	}

	fmt.Println(opts_arr)
}

type ListOptions struct {
}

func parse(opts []string, supported_opts map[string]string) ([]string, error) {
	//get_opts(opts)
	return opts, nil
}

func get_opts(opts []string) []string {

	for idx := range opts {
		opt := opts[idx]
		if strings.HasPrefix(opt, "-") {
			split_char(opt[1:])
		}
	}
	return opts
}

func split_char(str string) {

}

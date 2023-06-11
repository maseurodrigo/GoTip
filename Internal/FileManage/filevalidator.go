package FileManage

import (
	"fmt"
	"os"
)

func CheckFile(file string) (bool, string) {

	// Check if file exists
	if _, errFileStruct := os.Stat(file); errFileStruct == nil {
		return true, ""
	} else if os.IsNotExist(errFileStruct) {
		return false, fmt.Sprintf("File %s doesn't exist\n", file)
	} else {
		return false, fmt.Sprintf("Error checking file %s: %v\n", file, errFileStruct)
	}
}

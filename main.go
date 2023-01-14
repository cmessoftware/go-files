package files

import (
	"bufio"
	"fmt"
	"os"
)

//Keep in mind that the visibility in go package is defined in the object name itself
//If we want define exported object (public keyword in other languaje such as Python, Java or C#)
//we must use Capital letter.

func ReadFileLines(filePath string) []string {

	// read the whole content of file and pass it to file variable, in case of error pass it to err variable
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}

	return fileLines
}

package standardio

import (
	"fmt"
	"os"
)

func PrintOSArgs() {

	fmt.Println("Printing with Separation")
	printWithSeparation()
	fmt.Println(os.Args)

}

func printWithSeparation() {

	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[1])
	}
}

package mangeo

import (
	"fmt"
	"os/exec"
  "os"
	"runtime"
)

version := "1.0.0"

func checkw() bool {
  if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}

func execute(string command) {

	// here we perform the pwd command.
	// we can store the output of this in our out variable
	// and catch any errors in err
	out, err := exec.Command(command).Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Printf("%s\n", output)
}

func main() {
	fmt.Println("Mangeo Build System %s\n\n Running on %s", version, runtime.GOOS)
  
}
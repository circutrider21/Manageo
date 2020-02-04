package manageo

import (
	"fmt"
	"os/exec"
  "os"
	"runtime"
)
func checkw() bool {
  if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}

func execute(command string) {
	out, err := exec.Command(command).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Printf("%s\n", output)
}

func gobuild(filename string) {
  bqg := "go build " + filename
  execute(bqg)
}

func main() {
  version := "1.0.0"
	fmt.Printf("Manageo Build System %s\n\n Running on %s\n", version, runtime.GOOS)
  switch arga := os.Args[1]; arga {
	case "version":
		fmt.Println(version)
	case "go":
		gobuild(os.Args[2])
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Println("No Actions Have Been Specified")
	}
}
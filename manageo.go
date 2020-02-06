package manageo

import (
	"fmt"
	"os/exec"
  "os"
	"runtime"
  "io/ioutil"
  "encoding/json"
)
func checkw() bool {
  if runtime.GOOS == "windows" {
		return true
	} else {
		return false
	}
}
func helptext() {
  fmt.Println("Usage: Manageo <language> <action> <filename>")
  fmt.Print("\n")
  fmt.Println("Suppoted Languages and examples:")
  fmt.Println("\t Go: Manageo go <filename>")
  fmt.Println("\t Ruby: Manageo ruby <windows | console> <filename>")
}

func fjhan(command string) {
  dat, err := ioutil.ReadFile("mango.json")
  if err != nil {
    panic(err)
  }
  bparse := string(dat)
  var result map[string]interface{}
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
  version := "1.1.0"
	fmt.Printf("Manageo Build System %s\n\n Running on %s\n", version, runtime.GOOS)
  switch arga := os.Args[1]; arga {
	case "version":
		fmt.Println(version)
	case "go":
		gobuild(os.Args[2])
  case "help":
    helptext()
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fjhan(os.Args[2])
	}
}
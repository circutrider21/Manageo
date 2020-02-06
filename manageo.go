package manageo

import (
	"fmt"
	"os/exec"
  "os"
	"runtime"
  "io/ioutil"
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
  fmt.Println("\t Go")

}

func fjhan(command string) {
  dat, err := ioutil.ReadFile("mango.json")
  if e != nil {
    panic(e)
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

func rbuild(filename, setting string) {
  switch setting {
    case "window":
      huj := "ocra --windows " + filename
    case "console":
      huj := "ocra --console " + filename
    default:
      huj := "ocra --console " + filename
  }
  execute(huj)
}

func main() {
  version := "1.0.0"
	fmt.Printf("Manageo Build System %s\n\n Running on %s\n", version, runtime.GOOS)
  switch arga := os.Args[1]; arga {
	case "version":
		fmt.Println(version)
	case "go":
		gobuild(os.Args[2])
  case "help":
    helptext()
  case "ruby":
    rbuild(os.Args[3], os.Args[2])
  case "file":
    fjhan()
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Println("No Actions Have Been Specified")
	}
}
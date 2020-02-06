package manageo

import "testing"

/*func TestAdd(t *testing.T) {
    total := add(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}*/
func testos(t *testing.T) {
  osan := checkw()
  req := false
  if osan != false {
    t.Errorf("OS Check failed. Required %b got %b", req, osan)
  }
}
package iteration

import (
  "testing"
  "fmt"
)

func TestRepeat(t *testing.T) {
  repeated := Repeat("a", 5)
  expected := "aaaaa"

  if repeated != expected {
    t.Errorf("expected '%s' but got '%s'", expected, repeated)
  }

}

func ExampleRepeat() {
  result := Repeat("a", 5)
  fmt.Println(result)
  // Output "aaaaa"
}

func BenchmarkRepeat(b *testing.B) {
  for i:= 0; i < b.N; i++ {
    Repeat("a", 5)
  }

}

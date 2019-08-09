package thangs

import (
	"fmt"
	"testing"
)

type FakeThang struct {
	name string
	status bool
}

func (f *FakeThang) DoThangs(thang Thangs) {
	fmt.Println("fake thangs dooo")
}

func (f *FakeThang) PrintOutThur() {
	fmt.Println("Print out thurr")
}

func TestThangs_DoThangs(t *testing.T) {
	fakeThing := FakeThang{}

	fakeThing.DoThangs(fakeThing)
}

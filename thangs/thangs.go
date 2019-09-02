package thangs

type DoThangs interface {
	DoThangs(thang Thangs)
}

type Thangs struct {
	name string
	status bool
}

func GetBool() bool {
	return true
}

func (t *Thangs) DoThangs(thang Thangs) {
	// do stuff
}

func (t *Thangs) NewThangs() *Thangs {
	return &Thangs{}
}

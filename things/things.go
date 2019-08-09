package things

type Things struct {
	status bool
	name string
}

func GiveBool() bool {
	return false
}

func NewThings() *Things {
	return &Things{}
}

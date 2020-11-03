package parser

import "fmt"

type Fetch struct {
	text string
}

func NewFetch() *Fetch {
	return &Fetch{}
}

func (f *Fetch) Parse() (*Path, error) {
	return &Path{}, fmt.Errorf("fetch format has not been implemented")
}

func (f *Fetch) SetText(text string) {
	f.text = text
}

func (f *Fetch) RequiredField() func(input string) error {
	return func(input string) error {
		return fmt.Errorf("fetch format has not been implemented")
	}
}


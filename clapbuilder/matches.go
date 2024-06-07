package clapbuilder

type Matches struct{}

func (m *Matches) GetOneString(name string) (string, bool) {
	return "", false
}

func (m *Matches) GetOneU8(name string) (uint8, bool) {
	return 0, false
}

func (m *Matches) SubcommandMatches(name string) (*Matches, bool) {
	return nil, false
}

func (m *Matches) GetFlag(name string) bool {
	return false
}

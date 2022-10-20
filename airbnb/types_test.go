package airbnb

import "testing"

func Test_randomName(t *testing.T) {
	got := randomName(nameElements)
	t.Log(got)
	got = randomName(nameElements)
	t.Log(got)
	got = randomName(nameElements)
	t.Log(got)
	got = randomName(nameElements)
	t.Log(got)
}

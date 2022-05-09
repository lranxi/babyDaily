package random

import (
	"testing"
	"unicode/utf8"
)

func TestStr(t *testing.T) {
	t.Log(len("世界A"), utf8.RuneCountInString("世界A"))
}

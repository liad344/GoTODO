package Parser

import "testing"

func TestGOFUNCRegex(t *testing.T) {
	got := regex("func(str *string ) realName (lots, of, arguments int)", GOFUNCRegex)
	if got <= 0 {
		t.Errorf("Should be a function non func")
	}
}
func BenchmarkGOFUNCRegex(b *testing.B) {
	got := regex("func(str *string ) realName (lots, of, arguments int)", GOFUNCRegex)
	if got <= 0 {
		b.Errorf("Should be a function")
	}
}

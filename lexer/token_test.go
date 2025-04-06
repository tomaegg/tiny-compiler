package lexer

import "testing"

func TestInitToken(t *testing.T) {
	for pos, e := range TokenDescs() {
		if e == nil {
			t.Errorf("find nil at <%d>", pos)
		}
		t.Log(e)
	}
}

package lexer

import (
	"testing"
)

func TestKwTable(t *testing.T) {
	for pos, e := range TokenDescs() {
		if e == nil {
			t.Errorf("find nil at <%d>", pos)
		}
		t.Log(e)
	}

	kw := GetKWTable()

	cases := []struct {
		token string
		exist bool
	}{
		{"i32", true},
		{"let", true},
		{"if", true},
		{"else", true},
		{"return", true},
		{"mut", true},
		{"fn", true},
		{"loop", true},
		{"for", true},
		{"in", true},
		{"break", true},
		{"continue", true},
		{"fuck", false},
		{"bitch", false},
		{"xupeng", false},
	}

	for _, c := range cases {
		tk := kw.Exist(c.token)
		exists := tk != nil
		if exists != c.exist {
			t.Errorf("([%s] in keyword) should be %v, but %v",
				c.token,
				c.exist,
				exists)
		} else {
			t.Logf("token[%s] --> %v", c.token, c.exist)
		}
	}
}

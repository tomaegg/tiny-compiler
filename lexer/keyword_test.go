package lexer

import (
	"testing"
)

func TestKwTable(t *testing.T) {
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
		exist := kw.Exist(c.token) != nil
		if exist != c.exist {
			t.Errorf("(token[%s] in keyword) should be %v, but %v",
				c.token,
				c.exist,
				exist)
		} else {
			t.Logf("token[%s] --> %v", c.token, c.exist)
		}
	}
}

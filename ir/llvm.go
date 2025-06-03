package ir

import (
	"fmt"
	"tj-compiler/symtable"

	"tinygo.org/x/go-llvm"
)

type symLLVMValueKeyType struct{}

var symLLVMValueKey = symLLVMValueKeyType{}

func SetValue(s symtable.Symbol, t llvm.Value) {
	s.Set(symLLVMValueKey, t)
}

func GetValue(s symtable.Symbol) llvm.Value {
	val := s.Get(symLLVMValueKey)
	if val == nil {
		panic(fmt.Sprintf("should not get nil llvm.Value in <%s>", s))
	}
	return val.(llvm.Value)
}

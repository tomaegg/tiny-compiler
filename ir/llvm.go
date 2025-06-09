package ir

import (
	"fmt"
	"tj-compiler/symtable"

	"tinygo.org/x/go-llvm"
)

type (
	llvmValueKeyType struct{}
	llvmTypeKeyType  struct{}
)

var (
	ctxLLVMValueKey = llvmValueKeyType{}
	ctxLLVMTypeKey  = llvmTypeKeyType{}
)

func SetValue(s symtable.Symbol, t llvm.Value) {
	s.Set(ctxLLVMValueKey, t)
}

func GetValue(s symtable.Symbol) llvm.Value {
	val := s.Get(ctxLLVMValueKey)
	if val == nil {
		panic(fmt.Sprintf("should not get nil llvm.Value in <%s>", s))
	}
	return val.(llvm.Value)
}

func SetType(s symtable.Symbol, t llvm.Type) {
	s.Set(ctxLLVMTypeKey, t)
}

func GetType(s symtable.Symbol) llvm.Type {
	val := s.Get(ctxLLVMTypeKey)
	if val == nil {
		panic(fmt.Sprintf("should not get nil llvm.Type in <%s>", s))
	}
	return val.(llvm.Type)
}

func IsATerminator(instr llvm.Value) bool {
	if instr.IsNil() {
		return false
	}
	op := instr.InstructionOpcode()
	switch op {
	// TODO: 目前仅仅涉及这两种
	case llvm.Br, llvm.Ret:
		return true
	default:
		return false
	}
}

.PHONY: all parser symtable ir

BINS=parser symtable ir

all: $(BINS)
	strip build/*

parser:
	@mkdir -p build
	go build -o build/parser cmd/parser/main.go

symtable:
	@mkdir -p build
	go build -o build/symtable cmd/symtable/main.go

ir:
	@mkdir -p build
	go build -tags=llvm19 -o build/ir cmd/ir/main.go

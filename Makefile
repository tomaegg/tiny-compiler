.PHONY: generate pack build run dot fmt antlr4

PWD:=$(shell realpath .)
SUBMIT_ZIP=submit.zip
SUBMIT_DIR=submit
GEN_DIR=g4/parser
OUT_DIR=build
TARGET_FILE=cmd/$(TARGET)/main.go
ANTLR4_LINK=https://www.antlr.org/download/antlr-4.13.2-complete.jar
ANTLR4=java -jar $(TOOL_DIR)/antlr-4.13.2-complete.jar
TOOL_DIR=$(PWD)/build_tools

define build-target
    mkdir -p $(OUT_DIR)
    GOOS=darwin GOARCH=amd64 go build -o $(OUT_DIR)/$(1)-macos-amd64 cmd/$(1)/main.go
    GOOS=darwin GOARCH=arm64 go build -o $(OUT_DIR)/$(1)-macos-arm64 cmd/$(1)/main.go
    GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/$(1)-linux-amd64 cmd/$(1)/main.go
    GOOS=linux GOARCH=arm64 go build -o $(OUT_DIR)/$(1)-linux-arm64 cmd/$(1)/main.go
    GOOS=windows GOARCH=amd64 go build -o $(OUT_DIR)/$(1)-windows-amd64.exe cmd/$(1)/main.go
endef


antlr4:
	mkdir -p $(TOOL_DIR)
	wget  $(ANTLR4_LINK) -O $(TOOL_DIR)/antlr-4.13.2-complete.jar

run: 
	@go run -v cmd/$(TARGET)/main.go $(ARGS)

pack: generate build
	rm -rf $(SUBMIT_DIR) $(SUBMIT_ZIP) && mkdir -p $(SUBMIT_DIR)
	rsync -avm --include='*/' --include='*.go' \
	--include="README.md" \
	--include='go.mod' --include='go.sum' \
	--include='build/*' --include="example/*" \
	--exclude='*' . $(SUBMIT_DIR)
	sed -i '/^$$/d; /^\/\//d' $(SUBMIT_DIR)/$(GEN_DIR)/*.go
	zip -r $(SUBMIT_ZIP) $(SUBMIT_DIR)
	rm -rf $(SUBMIT_DIR)

dot: 
	go run -v cmd/symtable/main.go $(ARGS) > symtable.gv
	dot -Tsvg symtable.gv -o symtable.svg
	dot -Tpng symtable.gv -o symtable.png
	dot -Tpdf symtable.gv -o symtable.pdf

generate: 
	cd g4 && \
	$(ANTLR4) -Dlanguage=Go -o parser RustLikeLexer.g4 && \
	$(ANTLR4) -Dlanguage=Go -visitor -no-listener -o parser RustLikeParser.g4 
	rm $(GEN_DIR)/*.tokens $(GEN_DIR)/*.interp
	go fmt ./g4/parser/

build:
	$(call build-target,symtable)
	$(call build-target,parser)

fmt:
	go fmt ./...

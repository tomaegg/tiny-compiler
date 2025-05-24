.PHONY: generate pack build

SUBMIT_ZIP=submit.zip
SUBMIT_DIR=submit
GEN_DIR=g4/parser
OUT_DIR=build
TARGET_FILE=cmd/$(TARGET)/main.go

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
	dot -Tpng symtable.gv -o symtable.png

generate: 
	cd g4 && \
	antlr4 -Dlanguage=Go -o parser RustLikeLexer.g4 && \
	antlr4 -Dlanguage=Go -o parser RustLikeParser.g4 
	rm $(GEN_DIR)/*.tokens $(GEN_DIR)/*.interp
	go fmt ./g4/parser/

build:
	mkdir -p $(OUT_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(OUT_DIR)/$(TARGET)-macos-amd64 $(TARGET_FILE)
	GOOS=darwin GOARCH=arm64 go build -o $(OUT_DIR)/$(TARGET)-macos-arm64 $(TARGET_FILE)
	GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/$(TARGET)-linux-amd64 $(TARGET_FILE)
	GOOS=linux GOARCH=arm64 go build -o $(OUT_DIR)/$(TARGET)-linux-arm64 $(TARGET_FILE)
	GOOS=windows GOARCH=amd64 go build -o $(OUT_DIR)/$(TARGET)-windows-amd64.exe $(TARGET_FILE)

fmt:
	go fmt ./...

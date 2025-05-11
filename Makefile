.PHONY: generate pack

SUBMIT_DIR=submit
SUBMIT_ZIP=submit.zip
GEN_DIR=g4/parser
PARSER_OUT=parser
PARSER_TARGET_FILE=cmd/parser/main.go
OUT_DIR=build

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

generate: 
	cd g4 && \
	antlr4 -Dlanguage=Go -o parser RustLikeLexer.g4 && \
	antlr4 -Dlanguage=Go -o parser RustLikeParser.g4 
	rm $(GEN_DIR)/*.tokens $(GEN_DIR)/*.interp
	go fmt ./...

build-parser:
	GOOS=darwin GOARCH=amd64 go build -o $(OUT_DIR)/$(PARSER_OUT)-macos-amd64 $(PARSER_TARGET_FILE) 
	GOOS=darwin GOARCH=arm64 go build -o $(OUT_DIR)/$(PARSER_OUT)-macos-arm64 $(PARSER_TARGET_FILE) 
	GOOS=linux GOARCH=amd64 go build -o $(OUT_DIR)/$(PARSER_OUT)-linux-amd64 $(PARSER_TARGET_FILE) 
	GOOS=linux GOARCH=arm64 go build -o $(OUT_DIR)/$(PARSER_OUT)-linux-arm64 $(PARSER_TARGET_FILE) 
	GOOS=windows GOARCH=amd64 go build -o $(OUT_DIR)/$(PARSER_OUT)-windows-amd64.exe $(PARSER_TARGET_FILE) 

build: build-parser
	mkdir -p $(OUT_DIR)

.PHONY: generate pack build run dot fmt antlr4 ir symtable parser

############################################
# FOR DEV ONLY
############################################

PWD:=$(shell realpath .)
SUBMIT_ZIP=submit.zip
SUBMIT_DIR=submit
GEN_DIR=g4/parser
OUT_DIR=build
TARGET_FILE=cmd/$(TARGET)/main.go
ANTLR4_LINK=https://www.antlr.org/download/antlr-4.13.2-complete.jar
ANTLR4=java -jar $(TOOL_DIR)/antlr-4.13.2-complete.jar
TOOL_DIR=$(PWD)/build_tools

antlr4:
	mkdir -p $(TOOL_DIR)
	wget  $(ANTLR4_LINK) -O $(TOOL_DIR)/antlr-4.13.2-complete.jar

parser:
	@go run -v cmd/parser/main.go $(ARGS)

symtable:
	@go run -v cmd/symtable/main.go $(ARGS)

ir: 
	@go run -tags=llvm19 -v cmd/ir/main.go $(ARGS)

dot: 
	go run -v cmd/symtable/main.go $(ARGS) > symtable.gv
	dot -Tsvg symtable.gv -o symtable.svg
	dot -Tpng symtable.gv -o symtable.png
	dot -Tpdf symtable.gv -o symtable.pdf

pack: build
	rm -rf $(SUBMIT_DIR) $(SUBMIT_ZIP) && mkdir -p $(SUBMIT_DIR)
	git archive --format=tar --output=src.tar HEAD
	tar xvf src.tar --directory=$(SUBMIT_DIR) --exclude='*.g4' --exclude='Makefile' && rm src.tar
	sed -i '/^$$/d; /^\/\//d' $(SUBMIT_DIR)/$(GEN_DIR)/*.go
	./run.sh export && mv tiny-compiler-latest.tar $(SUBMIT_DIR)/ # 导出image
	zip -r $(SUBMIT_ZIP) $(SUBMIT_DIR)
	rm -rf $(SUBMIT_DIR)

generate: 
	cd g4 && \
	$(ANTLR4) -Dlanguage=Go -o parser RustLikeLexer.g4 && \
	$(ANTLR4) -Dlanguage=Go -visitor -no-listener -o parser RustLikeParser.g4 
	rm $(GEN_DIR)/*.tokens $(GEN_DIR)/*.interp
	go fmt ./g4/parser/

build: fmt
	./run.sh build

fmt:
	go fmt ./...

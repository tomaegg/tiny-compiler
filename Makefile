.PHONY: generate pack

SUBMIT_DIR=submit
SUBMIT_ZIP=submit.zip
GEN_DIR=g4/parser

pack: generate
	rm -rf $(SUBMIT_DIR) $(SUBMIT_ZIP) && mkdir -p $(SUBMIT_DIR)
	rsync -avm --include='*/' --include='*.go' \
	--include='go.mod' --include='go.sum' \
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

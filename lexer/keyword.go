package lexer

import "sync"

type KeywordTable interface {
	Exist(string) Token
}

type keywordTable struct {
	table  map[string]Token
	tokens []Token
}

var kwTable keywordTable = keywordTable{
	table:  make(map[string]Token),
	tokens: ConstTokens(),
}

var onceKTable sync.Once

// 返回keywordtable
func GetKWTable() KeywordTable {
	onceKTable.Do(initKWTable)
	return &kwTable
}

func initKWTable() {
	kwTable.put(TokenINT32)
	kwTable.put(TokenLET)
	kwTable.put(TokenIF)
	kwTable.put(TokenELSE)
	kwTable.put(TokenRETURN)
	kwTable.put(TokenMUT)
	kwTable.put(TokenFN)
	kwTable.put(TokenFOR)
	kwTable.put(TokenIN)
	kwTable.put(TokenLOOP)
	kwTable.put(TokenBREAK)
	kwTable.put(TokenCONTINUE)
}

func (kw *keywordTable) put(tkType TokenType) {
	tk := kw.tokens[tkType]
	kwTable.table[tk.Literal()] = tk
}

func (kw *keywordTable) Exist(text string) Token {
	return kw.table[text]
}

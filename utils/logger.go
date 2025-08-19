package utils

import (
	"github.com/antlr4-go/antlr/v4"
	log "github.com/sirupsen/logrus"
)

func PosLogger(token antlr.Token) *log.Entry {
	line, col := token.GetLine(), token.GetColumn()
	return log.WithFields(log.Fields{
		"line": line,
		"col":  col,
	})
}

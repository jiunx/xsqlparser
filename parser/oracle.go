package parser

import (
	"context"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	perseverant "github.com/jiunx/xsqlparser/dialect/oracle/gen"
	"github.com/jiunx/xsqlparser/listener"
	"github.com/jiunx/xsqlparser/model"
)

var (
	oracleListener = listener.OracleInstance
)

type oracleParser struct{}

// Parse parses the sql and returns a Statement .
func (p *oracleParser) Parse(ctx context.Context, sql string) (stmt *model.Statement, err error) {
	is := antlr.NewInputStream(sql)
	lexer := perseverant.NewOracleStatementLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	sqlStatementParser := perseverant.NewOracleStatementParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(oracleListener, sqlStatementParser.Execute())
	oracleListener.Stmt.Sql = sql
	return &oracleListener.Stmt, nil
}

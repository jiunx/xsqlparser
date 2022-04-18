package parser

import (
	"context"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	perseverant "github.com/jiunx/xsqlparser/dialect/sql92/gen"
	"github.com/jiunx/xsqlparser/listener"
	"github.com/jiunx/xsqlparser/model"
)

var (
	sql92Listener = listener.Sql92Instance
)

type sql92Parser struct{}

//Parse parses the sql and returns a Statement .
func (p *sql92Parser) Parse(ctx context.Context, sql string) (stmt *model.Statement, err error) {
	is := antlr.NewInputStream(sql)
	lexer := perseverant.NewSQL92StatementLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	sqlStatementParser := perseverant.NewSQL92StatementParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(sql92Listener, sqlStatementParser.Execute())
	sql92Listener.Stmt.Sql = sql
	return &sql92Listener.Stmt, nil
}

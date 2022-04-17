package parser

import (
	"context"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	perseverant "github.com/jiunx/xsqlparser/dialect/sqlserver/gen"
	"github.com/jiunx/xsqlparser/listener"
	"github.com/jiunx/xsqlparser/model"
)

var (
	sqlserverListener = listener.SqlserverInstance
)

type sqlserverParser struct{}

//Parse parses the sql and returns a Statement .
func (p *sqlserverParser) Parse(ctx context.Context, sql string) (stmt *model.Statement, err error) {
	is := antlr.NewInputStream(sql)
	lexer := perseverant.NewSQLServerStatementLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	sqlStatementParser := perseverant.NewSQLServerStatementParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(sqlserverListener, sqlStatementParser.Execute())
	sqlserverListener.Stmt.Sql = sql
	return &sqlserverListener.Stmt, nil
}

package parser

import (
	"context"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	perseverant "github.com/jiunx/xsqlparser/dialect/mysql/gen"
	"github.com/jiunx/xsqlparser/listener"
	"github.com/jiunx/xsqlparser/model"
)

var (
	mysqlListener = listener.MysqlInstance
)

type mysqlParser struct{}

//Parse parses the sql and returns a Statement .
func (p *mysqlParser) Parse(ctx context.Context, sql string) (stmt *model.Statement, err error) {
	is := antlr.NewInputStream(sql)
	lexer := perseverant.NewMySQLStatementLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	sqlStatementParser := perseverant.NewMySQLStatementParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(mysqlListener, sqlStatementParser.Execute())
	mysqlListener.Stmt.Sql = sql
	return &mysqlListener.Stmt, nil
}

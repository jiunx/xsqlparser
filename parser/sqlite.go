package parser

import (
	"context"

	"github.com/antlr4-go/antlr/v4"
	perseverant "github.com/jiunx/xsqlparser/dialect/sqlite/gen"
	"github.com/jiunx/xsqlparser/listener"
	"github.com/jiunx/xsqlparser/model"
)

var (
	sqlitelListener = listener.SqliteInstance
)

type sqliteParser struct{}

// Parse parses the sql and returns a Statement .
func (p *sqliteParser) Parse(ctx context.Context, sql string) (stmt *model.Statement, err error) {
	is := antlr.NewInputStream(sql)
	lexer := perseverant.NewSQLiteLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	sqlStatementParser := perseverant.NewSQLiteParser(stream)
	sqlStatementParser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := sqlStatementParser.Parse()

	antlr.ParseTreeWalkerDefault.Walk(sqlitelListener, tree)
	sqlitelListener.Stmt.Sql = sql
	return &sqlitelListener.Stmt, nil
}

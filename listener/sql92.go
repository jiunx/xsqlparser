package listener

import (
	parser "github.com/jiunx/xsqlparser/dialect/sql92/gen"
	"github.com/jiunx/xsqlparser/model"
	"github.com/jiunx/xsqlparser/util"
)

var Sql92Instance = &Sql92Listener{}

type Sql92Listener struct {
	*parser.BaseSQL92StatementListener
	Stmt model.Statement
}

func (l *Sql92Listener) EnterExecute(ctx *parser.ExecuteContext) {
	stmt := model.Statement{}
	sqlType := ctx.GetStart().GetText()
	stmt.Ast = ctx
	stmt.Language = util.Sql92
	stmt.SqlType = sqlType
	l.Stmt = stmt
}

func (l *Sql92Listener) EnterProjections(ctx *parser.ProjectionsContext) {
	if ctx.UnqualifiedShorthand() != nil {
		l.Stmt.Results = append(l.Stmt.Results, "*")
	}
}

func (l *Sql92Listener) EnterProjection(ctx *parser.ProjectionContext) {
	l.Stmt.Results = append(l.Stmt.Results, ctx.GetText())
}

func (l *Sql92Listener) EnterColumnName(ctx *parser.ColumnNameContext) {
	if ctx != nil {
		l.Stmt.Columns = append(l.Stmt.Columns, ctx.GetStop().GetText())
	}
}

func (l *Sql92Listener) EnterTableName(ctx *parser.TableNameContext) {
	if ctx != nil {
		l.Stmt.Tables = append(l.Stmt.Tables, ctx.Name().GetText())
	}
}

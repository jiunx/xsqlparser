package listener

import (
	parser "github.com/yhao521/xsqlparser/dialect/sqlserver/gen"
	"github.com/yhao521/xsqlparser/model"
	"github.com/yhao521/xsqlparser/util"
)

var SqlserverInstance = &sqlserverListener{}

type sqlserverListener struct {
	*parser.BaseSQLServerStatementListener
	Stmt model.Statement
}

func (l *sqlserverListener) EnterExecute(ctx *parser.ExecuteContext) {
	stmt := model.Statement{}
	sqlType := ctx.GetStart().GetText()
	stmt.Ast = ctx
	stmt.Language = util.SqlServer
	stmt.SqlType = sqlType
	l.Stmt = stmt
}

func (l *sqlserverListener) EnterProjections(ctx *parser.ProjectionsContext) {
	if ctx.UnqualifiedShorthand() != nil {
		l.Stmt.Results = append(l.Stmt.Results, "*")
	}
}

func (l *sqlserverListener) EnterProjection(ctx *parser.ProjectionContext) {
	l.Stmt.Results = append(l.Stmt.Results, ctx.GetText())
}

func (l *sqlserverListener) EnterColumnName(ctx *parser.ColumnNameContext) {
	if ctx != nil {
		l.Stmt.Columns = append(l.Stmt.Columns, ctx.GetStop().GetText())
	}
}

func (l *sqlserverListener) EnterTableName(ctx *parser.TableNameContext) {
	if ctx != nil {
		l.Stmt.Tables = append(l.Stmt.Tables, ctx.Name().GetText())
	}
}

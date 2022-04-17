package listener

import (
	parser "github.com/jiunx/xsqlparser/dialect/oracle/gen"
	"github.com/jiunx/xsqlparser/model"
	"github.com/jiunx/xsqlparser/util"
)

var OracleInstance = NewOracleListener()

type OracleListener struct {
	*parser.BaseOracleStatementListener
	Stmt model.Statement
}

func NewOracleListener() *OracleListener {
	return &OracleListener{}
}

func (l *OracleListener) EnterExecute(ctx *parser.ExecuteContext) {
	stmt := model.Statement{}
	sqlType := ctx.GetStart().GetText()
	stmt.Ast = ctx
	stmt.Language = util.Oracle
	stmt.SqlType = sqlType
	l.Stmt = stmt
}

func (l *OracleListener) EnterSelectList(ctx *parser.SelectListContext) {
	if ctx.UnqualifiedShorthand() != nil {
		l.Stmt.Results = append(l.Stmt.Results, "*")
	}
}

func (l *OracleListener) EnterSelectProjection(ctx *parser.SelectProjectionContext) {
	if ctx != nil {
		l.Stmt.Results = append(l.Stmt.Results, ctx.GetText())
	}
}

func (l *OracleListener) EnterColumnName(ctx *parser.ColumnNameContext) {
	if ctx != nil {
		l.Stmt.Columns = append(l.Stmt.Columns, ctx.Name().GetText())
	}
}

func (l *OracleListener) EnterTableName(ctx *parser.TableNameContext) {
	if ctx != nil {
		l.Stmt.Tables = append(l.Stmt.Tables, ctx.Name().GetText())
	}
}

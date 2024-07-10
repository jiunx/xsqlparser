package listener

import (
	parser "github.com/yhao521/xsqlparser/dialect/mysql/gen"
	"github.com/yhao521/xsqlparser/model"
	"github.com/yhao521/xsqlparser/util"
)

var MysqlInstance = &MySQLListener{}

type MySQLListener struct {
	*parser.BaseMySQLStatementListener
	Stmt model.Statement
}

func (l *MySQLListener) EnterExecute(ctx *parser.ExecuteContext) {
	stmt := model.Statement{}
	sqlType := ctx.GetStart().GetText()
	stmt.Ast = ctx
	stmt.Language = util.Mysql
	stmt.SqlType = sqlType
	l.Stmt = stmt
}

func (l *MySQLListener) EnterProjections(ctx *parser.ProjectionsContext) {
	if ctx.UnqualifiedShorthand() != nil {
		l.Stmt.Results = append(l.Stmt.Results, "*")
	}
}

func (l *MySQLListener) EnterProjection(ctx *parser.ProjectionContext) {
	l.Stmt.Results = append(l.Stmt.Results, ctx.GetText())
}

func (l *MySQLListener) EnterColumnRef(ctx *parser.ColumnRefContext) {
	if ctx != nil {
		l.Stmt.Columns = append(l.Stmt.Columns, ctx.GetStop().GetText())
	}
}

func (l *MySQLListener) EnterTableName(ctx *parser.TableNameContext) {
	if ctx != nil {
		l.Stmt.Tables = append(l.Stmt.Tables, ctx.Name().GetText())
	}
}

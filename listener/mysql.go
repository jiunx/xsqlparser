package listener

import (
	parser "github.com/jiunx/xsqlparser/dialect/mysql/gen"
	"github.com/jiunx/xsqlparser/model"
	"github.com/jiunx/xsqlparser/util"
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

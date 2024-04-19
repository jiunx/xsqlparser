package listener

import (
	parser "github.com/jiunx/xsqlparser/dialect/sqlite/gen"
	"github.com/jiunx/xsqlparser/model"
)

var SqliteInstance = &SQLiteListener{}

type SQLiteListener struct {
	*parser.BaseSQLiteParserListener
	Stmt model.Statement
}

// func (l *SQLiteListener) EnterExecute(ctx *parser.Begin_stmtContext) {
// 	stmt := model.Statement{}
// 	sqlType := ctx.GetStart().GetText()
// 	stmt.Ast = ctx
// 	stmt.Language = util.Sqlite
// 	stmt.SqlType = sqlType
// 	l.Stmt = stmt
// }

// func (l *SQLiteListener) EnterProjections(ctx *parser.AliasContext) {
// 	if ctx.UnqualifiedShorthand() != nil {
// 		l.Stmt.Results = append(l.Stmt.Results, "*")
// 	}
// }

// func (l *SQLiteListener) EnterProjection(ctx *parser.Any_nameContext) {
// 	l.Stmt.Results = append(l.Stmt.Results, ctx.GetText())
// }

func (l *SQLiteListener) EnterColumnRef(ctx *parser.Column_defContext) {
	if ctx != nil {
		l.Stmt.Columns = append(l.Stmt.Columns, ctx.GetStop().GetText())
	}
}

func (l *SQLiteListener) EnterTableName(ctx *parser.Table_nameContext) {
	if ctx != nil {
		l.Stmt.Tables = append(l.Stmt.Tables, ctx.Any_name().GetText())
	}
}

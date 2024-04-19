package listener

import (
	"fmt"

	parser "github.com/jiunx/xsqlparser/dialect/sqlite/gen"
	"github.com/jiunx/xsqlparser/model"
	"github.com/jiunx/xsqlparser/util"
)

var SqliteInstance = &SQLiteListener{}

type SQLiteListener struct {
	*parser.BaseSQLiteParserListener
	Stmt model.Statement
}

func (l *SQLiteListener) EnterParse(ctx *parser.ParseContext) {
	stmt := model.Statement{}
	// fmt.Println("EnterExecute", ctx)
	// fmt.Println("EnterExecute", util.Sqlite)
	sqlType := ctx.GetStart().GetText()
	stmt.Ast2 = ctx
	stmt.Language = util.Sqlite
	stmt.SqlType = sqlType
	l.Stmt = stmt
}

// func (l *SQLiteListener) EnterProjections(ctx *parser.AliasContext) {
// 	if ctx.UnqualifiedShorthand() != nil {
// 		l.Stmt.Results = append(l.Stmt.Results, "*")
// 	}
// }

func (l *SQLiteListener) EnterType_name(ctx *parser.Type_nameContext) {
	// fmt.Println("EnterProjection", ctx.GetText())
	l.Stmt.ColumnsDef = append(l.Stmt.ColumnsDef, ctx.GetText())
}

func (l *SQLiteListener) EnterColumn_def(ctx *parser.Column_defContext) {
	if ctx != nil {
		// fmt.Println("EnterColumnRef", ctx.GetText())
		l.Stmt.Columns = append(l.Stmt.Columns, ctx.Column_name().Any_name().GetText())
		// l.Stmt.ColumnsDef = append(l.Stmt.ColumnsDef, ctx.Type_name().GetText())
		// l.Stmt.ColumnsType = append(l.Stmt.ColumnsType, ctx().Any_name().Keyword().GetText())
	}
}

func (l *SQLiteListener) EnterTable_name(ctx *parser.Table_nameContext) {
	if ctx != nil {
		// fmt.Println("EnterTableName", ctx.GetText())
		l.Stmt.Tables = append(l.Stmt.Tables, ctx.Any_name().GetText())
	}
}

func (l *SQLiteListener) EnterTable_constraint(ctx *parser.Table_constraintContext) {
	// fmt.Println(ctx.GetText())
	if ctx != nil {
		fmt.Println("EnterColumn_constraint", ctx.Name().Any_name().Keyword())
		l.Stmt.Results = append(l.Stmt.Results, ctx.Column_name(0).GetText())
	}
}

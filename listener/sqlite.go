package listener

import (
	"fmt"

	parser "github.com/yhao521/xsqlparser/dialect/sqlite/gen"
	"github.com/yhao521/xsqlparser/model"
	"github.com/yhao521/xsqlparser/util"
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

func (l *SQLiteListener) EnterType_name(ctx *parser.Type_nameContext) {
	// fmt.Println("EnterProjection", ctx.GetText())
	// tableList := model.Table{}
	// tableList
	l.Stmt.TableList.ColumnsType = append(l.Stmt.TableList.ColumnsType, ctx.GetText())
}

func (l *SQLiteListener) EnterColumn_def(ctx *parser.Column_defContext) {
	if ctx != nil {
		l.Stmt.Columns = append(l.Stmt.Columns, ctx.Column_name().Any_name().GetText())

		// tableList := model.Table{}
		// tableList.Columns = l.Stmt.Columns
		l.Stmt.TableList.Columns = l.Stmt.Columns
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

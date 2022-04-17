package parser

import (
	"context"
	"github.com/jiunx/xsqlparser/util"
	"testing"
)

func TestMysqlParse(t *testing.T) {
	sql := "select id,name,address from t"
	parserFactory := GetParserFactoryInstance()
	statement, err := parserFactory.Parse(context.Background(), util.Mysql, sql)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", statement)
}

func TestOracleParse(t *testing.T) {
	sql := "select id,name,address from t"
	parserFactory := GetParserFactoryInstance()
	statement, err := parserFactory.Parse(context.Background(), util.Oracle, sql)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", statement)
}

func TestPostGreParse(t *testing.T) {
	sql := "select id,name,address from t"
	parserFactory := GetParserFactoryInstance()
	statement, err := parserFactory.Parse(context.Background(), util.PostGre, sql)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", statement)
}

func TestSqlServerParse(t *testing.T) {
	sql := "select id,name,address from t"
	parserFactory := GetParserFactoryInstance()
	statement, err := parserFactory.Parse(context.Background(), util.SqlServer, sql)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", statement)
}

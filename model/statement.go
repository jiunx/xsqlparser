package model

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	antlr2 "github.com/antlr4-go/antlr/v4"
)

// Statement is a representation of a statement in the database.
type Statement struct {
	Sql       string      `json:"sql"`
	SqlType   string      `json:"sqlType"`
	Language  string      `json:"language"`
	Tables    []string    `json:"tables"`
	Columns   []string    `json:"columns"`
	Results   []string    `json:"results"`
	TableList Table       `json:"tableList"`
	Ast       antlr.Tree  `json:"ast"`
	Ast2      antlr2.Tree `json:"ast2"`
}

type Table struct {
	TableName   string   `json:"tableName"`
	Columns     []string `json:"columns"`
	ColumnsDef  []string `json:"columnsDef"`
	ColumnsType []string `json:"columnsType"`
}

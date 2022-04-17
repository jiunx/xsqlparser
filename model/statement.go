package model

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Statement is a representation of a statement in the database.
type Statement struct {
	Sql      string     `json:"sql"`
	SqlType  string     `json:"sqlType"`
	Language string     `json:"language"`
	Tables   []string   `json:"tables"`
	Columns  []string   `json:"columns"`
	Results  []string   `json:"columns"`
	Ast      antlr.Tree `json:"ast"`
}

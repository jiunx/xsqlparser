package parser

import (
	"context"
	"errors"
	"github.com/jiunx/xsqlparser/model"
	"github.com/jiunx/xsqlparser/util"
)

var (
	parserFactoryInstance = NewParserFactory()
)

func init() {
	parserFactoryInstance.Register(util.Mysql, &mysqlParser{})
	parserFactoryInstance.Register(util.Oracle, &oracleParser{})
	parserFactoryInstance.Register(util.SqlServer, &sqlserverParser{})
	//parserFactoryInstance.Register(Opens, &OpensParser{})
	//parserFactoryInstance.Register(sql92, &Sql92Parser{})
	//parserFactoryInstance.Register(Postgre, &PostgreParser{})
}

type parser interface {
	Parse(context.Context, string) (*model.Statement, error)
}

// parserFactory is a factory for creating parsers.
type parserFactory struct {
	parsers map[string]parser
}

// NewParserFactory creates a new parserFactory.
func NewParserFactory() *parserFactory {
	return &parserFactory{
		parsers: make(map[string]parser),
	}
}

// GetParserFactoryInstance returns the singleton instance of parserFactory.
func GetParserFactoryInstance() *parserFactory {
	return parserFactoryInstance
}

// Register registers a parser for a given language.
func (pf *parserFactory) Register(language string, parser parser) {
	pf.parsers[language] = parser
}

// GetParser returns a parser for a given language.
func (pf *parserFactory) GetParser(language string) (parser, error) {
	if parser, ok := pf.parsers[language]; ok {
		return parser, nil
	}
	ErrParserNotSupportDB := errors.New("parser not support db")
	return nil, ErrParserNotSupportDB
}

// Parse parses a given string using a parser for a given language.
func (pf *parserFactory) Parse(ctx context.Context, language string, s string) (*model.Statement, error) {
	parser, err := pf.GetParser(language)
	if err != nil {
		return nil, err
	}
	return parser.Parse(ctx, s)
}

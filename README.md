# xSqlParser

Parse SQL statements, support mysql, oracle, sqlserver...

The first open source project of a golang novice, welcome to star, fork, if you have any needs, comments and suggestions, please mention them in the issue, I can also use antlr's TokenStreamRewrite to make logical changes to specific nodes of sql, and the project is still temporarily This feature was not added, I can add it if needed.

### DOCUMENTATION📜

<hr>

[![EN doc](https://img.shields.io/badge/document-English-blue.svg)](https://github.com/jiunx/xsqlparser/blob/master/README.md)
[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](https://github.com/jiunx/xsqlparser/blob/master/README-cn.md)
## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.


### Installing

Get it

```
go get github.com/jiunx/xsqlparser
```



## Running the tests

You can use the test code in parser_test.go to test that the parser parses the sql statement correctly.

```
    sql := "select id,name,address from t"
    parserFactory := parser.GetParserFactoryInstance()
    mysqlParser, err := parserFactory.GetParser(util.Mysql)
    if err != nil {
        fmt.Println(err)
        return
    }
    statement, err2 := mysqlParser.Parse(context.Background(), sql)
    if err2 != nil {
        fmt.Println(err2)
    }
    fmt.Printf("%+v\n", statement)
```
You will get output like this
```

    &{Sql:select id,name,address from t SqlType:select Language:mysql Tables:[t] Columns:[id name address] Results:[id name address] Ast:0xc000246c00}

```
The following describes the structure of the statement
```

statement.SqlType is the sql type, such as select, insert, update, delete, create, drop, etc.

statement.Tables is the table name, such as t.

statement.Columns is the column name, such as id, name, address.

statement.Results is the resultset, such as id, name, address.

statement.Ast This sql AST tree, if you are familiar with ANTLR or tree operations, you can also use listener to parse sql.
.
```

## Contributing
jiunx,  I am the author of this project.

[antlr](https://github.com/antlr/antlr4),  mainly rely on technology

[shardingsphere](https://github.com/apache/shardingsphere),the grammer file of antlr is obtained from here
## Versioning

For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Authors

jiunx

## Acknowledgments

* Antlr4


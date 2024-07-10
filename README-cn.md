# xSqlParser

解析SQL语句，支持mysql、oracle、sqlserver,sql92

一名golang新手的第一个开源项目，欢迎star, fork, 有什么需求、意见和建议尽管在issue里提,
我还可以通过antlr的TokenStreamRewrite来对sql的特定节点进行逻辑修改,项目内暂时还没有添加这项功能,如果需要的话我可以加上。


### DOCUMENTATION📜

<hr>

[![EN doc](https://img.shields.io/badge/document-English-blue.svg)](https://github.com/yhao521/xsqlparser/blob/master/README.md)
[![CN doc](https://img.shields.io/badge/文档-中文版-blue.svg)](https://github.com/yhao521/xsqlparser/blob/master/README-cn.md)
## 快速入门



### 安装

获取最新版本


```
go get github.com/yhao521/xsqlparser
```




## 运行测试

您可以使用 parser_test.go 中的测试代码来测试解析器是否正确解析了 sql 语句。
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
你会得到这样的输出
```

    &{Sql:select id,name,address from t SqlType:select Language:mysql Tables:[t] Columns:[id name address] Results:[id name address] Ast:0xc000246c00}

```
下面介绍statement的结构
```

statement.SqlType sql的类型, 如 select, insert, update, delete, create, drop, etc.

statement.Tables  sql涉及到的表名, 如 t.

statement.Columns sql涉及到的字段名, 如 id, name, address.

statement.Results select语句的结果集, 如 id, name, address.

statement.Ast     这条sql的AST树，如果你熟悉ANTLR或者树操作的话你也可以通过listener来解析sql.
```

## Contributing
jiunx（作者）.

[antlr](https://github.com/antlr/antlr4),  主要依赖的技术

[shardingsphere](https://github.com/apache/shardingsphere),从这里获取的antlr语法文件

## Authors

jiunx

## Acknowledgments

* Antlr4


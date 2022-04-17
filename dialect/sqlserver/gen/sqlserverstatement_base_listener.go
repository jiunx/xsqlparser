// Code generated from E:/GoProject/src/github.com/2212442929/xsqlparser/dialect/sqlserver/grammer\SQLServerStatement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQLServerStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSQLServerStatementListener is a complete listener for a parse tree produced by SQLServerStatementParser.
type BaseSQLServerStatementListener struct{}

var _ SQLServerStatementListener = &BaseSQLServerStatementListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSQLServerStatementListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSQLServerStatementListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSQLServerStatementListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSQLServerStatementListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExecute is called when production execute is entered.
func (s *BaseSQLServerStatementListener) EnterExecute(ctx *ExecuteContext) {}

// ExitExecute is called when production execute is exited.
func (s *BaseSQLServerStatementListener) ExitExecute(ctx *ExecuteContext) {}

// EnterInsert is called when production insert is entered.
func (s *BaseSQLServerStatementListener) EnterInsert(ctx *InsertContext) {}

// ExitInsert is called when production insert is exited.
func (s *BaseSQLServerStatementListener) ExitInsert(ctx *InsertContext) {}

// EnterInsertDefaultValue is called when production insertDefaultValue is entered.
func (s *BaseSQLServerStatementListener) EnterInsertDefaultValue(ctx *InsertDefaultValueContext) {}

// ExitInsertDefaultValue is called when production insertDefaultValue is exited.
func (s *BaseSQLServerStatementListener) ExitInsertDefaultValue(ctx *InsertDefaultValueContext) {}

// EnterInsertValuesClause is called when production insertValuesClause is entered.
func (s *BaseSQLServerStatementListener) EnterInsertValuesClause(ctx *InsertValuesClauseContext) {}

// ExitInsertValuesClause is called when production insertValuesClause is exited.
func (s *BaseSQLServerStatementListener) ExitInsertValuesClause(ctx *InsertValuesClauseContext) {}

// EnterInsertSelectClause is called when production insertSelectClause is entered.
func (s *BaseSQLServerStatementListener) EnterInsertSelectClause(ctx *InsertSelectClauseContext) {}

// ExitInsertSelectClause is called when production insertSelectClause is exited.
func (s *BaseSQLServerStatementListener) ExitInsertSelectClause(ctx *InsertSelectClauseContext) {}

// EnterUpdate is called when production update is entered.
func (s *BaseSQLServerStatementListener) EnterUpdate(ctx *UpdateContext) {}

// ExitUpdate is called when production update is exited.
func (s *BaseSQLServerStatementListener) ExitUpdate(ctx *UpdateContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseSQLServerStatementListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseSQLServerStatementListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterSetAssignmentsClause is called when production setAssignmentsClause is entered.
func (s *BaseSQLServerStatementListener) EnterSetAssignmentsClause(ctx *SetAssignmentsClauseContext) {
}

// ExitSetAssignmentsClause is called when production setAssignmentsClause is exited.
func (s *BaseSQLServerStatementListener) ExitSetAssignmentsClause(ctx *SetAssignmentsClauseContext) {}

// EnterAssignmentValues is called when production assignmentValues is entered.
func (s *BaseSQLServerStatementListener) EnterAssignmentValues(ctx *AssignmentValuesContext) {}

// ExitAssignmentValues is called when production assignmentValues is exited.
func (s *BaseSQLServerStatementListener) ExitAssignmentValues(ctx *AssignmentValuesContext) {}

// EnterAssignmentValue is called when production assignmentValue is entered.
func (s *BaseSQLServerStatementListener) EnterAssignmentValue(ctx *AssignmentValueContext) {}

// ExitAssignmentValue is called when production assignmentValue is exited.
func (s *BaseSQLServerStatementListener) ExitAssignmentValue(ctx *AssignmentValueContext) {}

// EnterDelete is called when production delete is entered.
func (s *BaseSQLServerStatementListener) EnterDelete(ctx *DeleteContext) {}

// ExitDelete is called when production delete is exited.
func (s *BaseSQLServerStatementListener) ExitDelete(ctx *DeleteContext) {}

// EnterSingleTableClause is called when production singleTableClause is entered.
func (s *BaseSQLServerStatementListener) EnterSingleTableClause(ctx *SingleTableClauseContext) {}

// ExitSingleTableClause is called when production singleTableClause is exited.
func (s *BaseSQLServerStatementListener) ExitSingleTableClause(ctx *SingleTableClauseContext) {}

// EnterMultipleTablesClause is called when production multipleTablesClause is entered.
func (s *BaseSQLServerStatementListener) EnterMultipleTablesClause(ctx *MultipleTablesClauseContext) {
}

// ExitMultipleTablesClause is called when production multipleTablesClause is exited.
func (s *BaseSQLServerStatementListener) ExitMultipleTablesClause(ctx *MultipleTablesClauseContext) {}

// EnterMultipleTableNames is called when production multipleTableNames is entered.
func (s *BaseSQLServerStatementListener) EnterMultipleTableNames(ctx *MultipleTableNamesContext) {}

// ExitMultipleTableNames is called when production multipleTableNames is exited.
func (s *BaseSQLServerStatementListener) ExitMultipleTableNames(ctx *MultipleTableNamesContext) {}

// EnterSelect is called when production select is entered.
func (s *BaseSQLServerStatementListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production select is exited.
func (s *BaseSQLServerStatementListener) ExitSelect(ctx *SelectContext) {}

// EnterAggregationClause is called when production aggregationClause is entered.
func (s *BaseSQLServerStatementListener) EnterAggregationClause(ctx *AggregationClauseContext) {}

// ExitAggregationClause is called when production aggregationClause is exited.
func (s *BaseSQLServerStatementListener) ExitAggregationClause(ctx *AggregationClauseContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BaseSQLServerStatementListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BaseSQLServerStatementListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterDuplicateSpecification is called when production duplicateSpecification is entered.
func (s *BaseSQLServerStatementListener) EnterDuplicateSpecification(ctx *DuplicateSpecificationContext) {
}

// ExitDuplicateSpecification is called when production duplicateSpecification is exited.
func (s *BaseSQLServerStatementListener) ExitDuplicateSpecification(ctx *DuplicateSpecificationContext) {
}

// EnterProjections is called when production projections is entered.
func (s *BaseSQLServerStatementListener) EnterProjections(ctx *ProjectionsContext) {}

// ExitProjections is called when production projections is exited.
func (s *BaseSQLServerStatementListener) ExitProjections(ctx *ProjectionsContext) {}

// EnterProjection is called when production projection is entered.
func (s *BaseSQLServerStatementListener) EnterProjection(ctx *ProjectionContext) {}

// ExitProjection is called when production projection is exited.
func (s *BaseSQLServerStatementListener) ExitProjection(ctx *ProjectionContext) {}

// EnterTop is called when production top is entered.
func (s *BaseSQLServerStatementListener) EnterTop(ctx *TopContext) {}

// ExitTop is called when production top is exited.
func (s *BaseSQLServerStatementListener) ExitTop(ctx *TopContext) {}

// EnterTopNum is called when production topNum is entered.
func (s *BaseSQLServerStatementListener) EnterTopNum(ctx *TopNumContext) {}

// ExitTopNum is called when production topNum is exited.
func (s *BaseSQLServerStatementListener) ExitTopNum(ctx *TopNumContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseSQLServerStatementListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseSQLServerStatementListener) ExitAlias(ctx *AliasContext) {}

// EnterUnqualifiedShorthand is called when production unqualifiedShorthand is entered.
func (s *BaseSQLServerStatementListener) EnterUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) {
}

// ExitUnqualifiedShorthand is called when production unqualifiedShorthand is exited.
func (s *BaseSQLServerStatementListener) ExitUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) {}

// EnterQualifiedShorthand is called when production qualifiedShorthand is entered.
func (s *BaseSQLServerStatementListener) EnterQualifiedShorthand(ctx *QualifiedShorthandContext) {}

// ExitQualifiedShorthand is called when production qualifiedShorthand is exited.
func (s *BaseSQLServerStatementListener) ExitQualifiedShorthand(ctx *QualifiedShorthandContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseSQLServerStatementListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseSQLServerStatementListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterTableReferences is called when production tableReferences is entered.
func (s *BaseSQLServerStatementListener) EnterTableReferences(ctx *TableReferencesContext) {}

// ExitTableReferences is called when production tableReferences is exited.
func (s *BaseSQLServerStatementListener) ExitTableReferences(ctx *TableReferencesContext) {}

// EnterTableReference is called when production tableReference is entered.
func (s *BaseSQLServerStatementListener) EnterTableReference(ctx *TableReferenceContext) {}

// ExitTableReference is called when production tableReference is exited.
func (s *BaseSQLServerStatementListener) ExitTableReference(ctx *TableReferenceContext) {}

// EnterTableFactor is called when production tableFactor is entered.
func (s *BaseSQLServerStatementListener) EnterTableFactor(ctx *TableFactorContext) {}

// ExitTableFactor is called when production tableFactor is exited.
func (s *BaseSQLServerStatementListener) ExitTableFactor(ctx *TableFactorContext) {}

// EnterJoinedTable is called when production joinedTable is entered.
func (s *BaseSQLServerStatementListener) EnterJoinedTable(ctx *JoinedTableContext) {}

// ExitJoinedTable is called when production joinedTable is exited.
func (s *BaseSQLServerStatementListener) ExitJoinedTable(ctx *JoinedTableContext) {}

// EnterJoinSpecification is called when production joinSpecification is entered.
func (s *BaseSQLServerStatementListener) EnterJoinSpecification(ctx *JoinSpecificationContext) {}

// ExitJoinSpecification is called when production joinSpecification is exited.
func (s *BaseSQLServerStatementListener) ExitJoinSpecification(ctx *JoinSpecificationContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseSQLServerStatementListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseSQLServerStatementListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseSQLServerStatementListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseSQLServerStatementListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseSQLServerStatementListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseSQLServerStatementListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseSQLServerStatementListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseSQLServerStatementListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseSQLServerStatementListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseSQLServerStatementListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterCteClause is called when production cteClause is entered.
func (s *BaseSQLServerStatementListener) EnterCteClause(ctx *CteClauseContext) {}

// ExitCteClause is called when production cteClause is exited.
func (s *BaseSQLServerStatementListener) ExitCteClause(ctx *CteClauseContext) {}

// EnterOutputClause is called when production outputClause is entered.
func (s *BaseSQLServerStatementListener) EnterOutputClause(ctx *OutputClauseContext) {}

// ExitOutputClause is called when production outputClause is exited.
func (s *BaseSQLServerStatementListener) ExitOutputClause(ctx *OutputClauseContext) {}

// EnterOutputWithColumns is called when production outputWithColumns is entered.
func (s *BaseSQLServerStatementListener) EnterOutputWithColumns(ctx *OutputWithColumnsContext) {}

// ExitOutputWithColumns is called when production outputWithColumns is exited.
func (s *BaseSQLServerStatementListener) ExitOutputWithColumns(ctx *OutputWithColumnsContext) {}

// EnterOutputWithColumn is called when production outputWithColumn is entered.
func (s *BaseSQLServerStatementListener) EnterOutputWithColumn(ctx *OutputWithColumnContext) {}

// ExitOutputWithColumn is called when production outputWithColumn is exited.
func (s *BaseSQLServerStatementListener) ExitOutputWithColumn(ctx *OutputWithColumnContext) {}

// EnterOutputWithAaterisk is called when production outputWithAaterisk is entered.
func (s *BaseSQLServerStatementListener) EnterOutputWithAaterisk(ctx *OutputWithAateriskContext) {}

// ExitOutputWithAaterisk is called when production outputWithAaterisk is exited.
func (s *BaseSQLServerStatementListener) ExitOutputWithAaterisk(ctx *OutputWithAateriskContext) {}

// EnterOutputTableName is called when production outputTableName is entered.
func (s *BaseSQLServerStatementListener) EnterOutputTableName(ctx *OutputTableNameContext) {}

// ExitOutputTableName is called when production outputTableName is exited.
func (s *BaseSQLServerStatementListener) ExitOutputTableName(ctx *OutputTableNameContext) {}

// EnterQueryHint is called when production queryHint is entered.
func (s *BaseSQLServerStatementListener) EnterQueryHint(ctx *QueryHintContext) {}

// ExitQueryHint is called when production queryHint is exited.
func (s *BaseSQLServerStatementListener) ExitQueryHint(ctx *QueryHintContext) {}

// EnterUseHitName is called when production useHitName is entered.
func (s *BaseSQLServerStatementListener) EnterUseHitName(ctx *UseHitNameContext) {}

// ExitUseHitName is called when production useHitName is exited.
func (s *BaseSQLServerStatementListener) ExitUseHitName(ctx *UseHitNameContext) {}

// EnterParameterMarker is called when production parameterMarker is entered.
func (s *BaseSQLServerStatementListener) EnterParameterMarker(ctx *ParameterMarkerContext) {}

// ExitParameterMarker is called when production parameterMarker is exited.
func (s *BaseSQLServerStatementListener) ExitParameterMarker(ctx *ParameterMarkerContext) {}

// EnterLiterals is called when production literals is entered.
func (s *BaseSQLServerStatementListener) EnterLiterals(ctx *LiteralsContext) {}

// ExitLiterals is called when production literals is exited.
func (s *BaseSQLServerStatementListener) ExitLiterals(ctx *LiteralsContext) {}

// EnterStringLiterals is called when production stringLiterals is entered.
func (s *BaseSQLServerStatementListener) EnterStringLiterals(ctx *StringLiteralsContext) {}

// ExitStringLiterals is called when production stringLiterals is exited.
func (s *BaseSQLServerStatementListener) ExitStringLiterals(ctx *StringLiteralsContext) {}

// EnterNumberLiterals is called when production numberLiterals is entered.
func (s *BaseSQLServerStatementListener) EnterNumberLiterals(ctx *NumberLiteralsContext) {}

// ExitNumberLiterals is called when production numberLiterals is exited.
func (s *BaseSQLServerStatementListener) ExitNumberLiterals(ctx *NumberLiteralsContext) {}

// EnterDateTimeLiterals is called when production dateTimeLiterals is entered.
func (s *BaseSQLServerStatementListener) EnterDateTimeLiterals(ctx *DateTimeLiteralsContext) {}

// ExitDateTimeLiterals is called when production dateTimeLiterals is exited.
func (s *BaseSQLServerStatementListener) ExitDateTimeLiterals(ctx *DateTimeLiteralsContext) {}

// EnterHexadecimalLiterals is called when production hexadecimalLiterals is entered.
func (s *BaseSQLServerStatementListener) EnterHexadecimalLiterals(ctx *HexadecimalLiteralsContext) {}

// ExitHexadecimalLiterals is called when production hexadecimalLiterals is exited.
func (s *BaseSQLServerStatementListener) ExitHexadecimalLiterals(ctx *HexadecimalLiteralsContext) {}

// EnterBitValueLiterals is called when production bitValueLiterals is entered.
func (s *BaseSQLServerStatementListener) EnterBitValueLiterals(ctx *BitValueLiteralsContext) {}

// ExitBitValueLiterals is called when production bitValueLiterals is exited.
func (s *BaseSQLServerStatementListener) ExitBitValueLiterals(ctx *BitValueLiteralsContext) {}

// EnterBooleanLiterals is called when production booleanLiterals is entered.
func (s *BaseSQLServerStatementListener) EnterBooleanLiterals(ctx *BooleanLiteralsContext) {}

// ExitBooleanLiterals is called when production booleanLiterals is exited.
func (s *BaseSQLServerStatementListener) ExitBooleanLiterals(ctx *BooleanLiteralsContext) {}

// EnterNullValueLiterals is called when production nullValueLiterals is entered.
func (s *BaseSQLServerStatementListener) EnterNullValueLiterals(ctx *NullValueLiteralsContext) {}

// ExitNullValueLiterals is called when production nullValueLiterals is exited.
func (s *BaseSQLServerStatementListener) ExitNullValueLiterals(ctx *NullValueLiteralsContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSQLServerStatementListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSQLServerStatementListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterUnreservedWord is called when production unreservedWord is entered.
func (s *BaseSQLServerStatementListener) EnterUnreservedWord(ctx *UnreservedWordContext) {}

// ExitUnreservedWord is called when production unreservedWord is exited.
func (s *BaseSQLServerStatementListener) ExitUnreservedWord(ctx *UnreservedWordContext) {}

// EnterDatabaseName is called when production databaseName is entered.
func (s *BaseSQLServerStatementListener) EnterDatabaseName(ctx *DatabaseNameContext) {}

// ExitDatabaseName is called when production databaseName is exited.
func (s *BaseSQLServerStatementListener) ExitDatabaseName(ctx *DatabaseNameContext) {}

// EnterSchemaName is called when production schemaName is entered.
func (s *BaseSQLServerStatementListener) EnterSchemaName(ctx *SchemaNameContext) {}

// ExitSchemaName is called when production schemaName is exited.
func (s *BaseSQLServerStatementListener) ExitSchemaName(ctx *SchemaNameContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseSQLServerStatementListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseSQLServerStatementListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterProcedureName is called when production procedureName is entered.
func (s *BaseSQLServerStatementListener) EnterProcedureName(ctx *ProcedureNameContext) {}

// ExitProcedureName is called when production procedureName is exited.
func (s *BaseSQLServerStatementListener) ExitProcedureName(ctx *ProcedureNameContext) {}

// EnterViewName is called when production viewName is entered.
func (s *BaseSQLServerStatementListener) EnterViewName(ctx *ViewNameContext) {}

// ExitViewName is called when production viewName is exited.
func (s *BaseSQLServerStatementListener) ExitViewName(ctx *ViewNameContext) {}

// EnterTriggerName is called when production triggerName is entered.
func (s *BaseSQLServerStatementListener) EnterTriggerName(ctx *TriggerNameContext) {}

// ExitTriggerName is called when production triggerName is exited.
func (s *BaseSQLServerStatementListener) ExitTriggerName(ctx *TriggerNameContext) {}

// EnterSequenceName is called when production sequenceName is entered.
func (s *BaseSQLServerStatementListener) EnterSequenceName(ctx *SequenceNameContext) {}

// ExitSequenceName is called when production sequenceName is exited.
func (s *BaseSQLServerStatementListener) ExitSequenceName(ctx *SequenceNameContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseSQLServerStatementListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseSQLServerStatementListener) ExitTableName(ctx *TableNameContext) {}

// EnterQueueName is called when production queueName is entered.
func (s *BaseSQLServerStatementListener) EnterQueueName(ctx *QueueNameContext) {}

// ExitQueueName is called when production queueName is exited.
func (s *BaseSQLServerStatementListener) ExitQueueName(ctx *QueueNameContext) {}

// EnterContractName is called when production contractName is entered.
func (s *BaseSQLServerStatementListener) EnterContractName(ctx *ContractNameContext) {}

// ExitContractName is called when production contractName is exited.
func (s *BaseSQLServerStatementListener) ExitContractName(ctx *ContractNameContext) {}

// EnterServiceName is called when production serviceName is entered.
func (s *BaseSQLServerStatementListener) EnterServiceName(ctx *ServiceNameContext) {}

// ExitServiceName is called when production serviceName is exited.
func (s *BaseSQLServerStatementListener) ExitServiceName(ctx *ServiceNameContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseSQLServerStatementListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseSQLServerStatementListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterOwner is called when production owner is entered.
func (s *BaseSQLServerStatementListener) EnterOwner(ctx *OwnerContext) {}

// ExitOwner is called when production owner is exited.
func (s *BaseSQLServerStatementListener) ExitOwner(ctx *OwnerContext) {}

// EnterName is called when production name is entered.
func (s *BaseSQLServerStatementListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseSQLServerStatementListener) ExitName(ctx *NameContext) {}

// EnterColumnNames is called when production columnNames is entered.
func (s *BaseSQLServerStatementListener) EnterColumnNames(ctx *ColumnNamesContext) {}

// ExitColumnNames is called when production columnNames is exited.
func (s *BaseSQLServerStatementListener) ExitColumnNames(ctx *ColumnNamesContext) {}

// EnterColumnNamesWithSort is called when production columnNamesWithSort is entered.
func (s *BaseSQLServerStatementListener) EnterColumnNamesWithSort(ctx *ColumnNamesWithSortContext) {}

// ExitColumnNamesWithSort is called when production columnNamesWithSort is exited.
func (s *BaseSQLServerStatementListener) ExitColumnNamesWithSort(ctx *ColumnNamesWithSortContext) {}

// EnterTableNames is called when production tableNames is entered.
func (s *BaseSQLServerStatementListener) EnterTableNames(ctx *TableNamesContext) {}

// ExitTableNames is called when production tableNames is exited.
func (s *BaseSQLServerStatementListener) ExitTableNames(ctx *TableNamesContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseSQLServerStatementListener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseSQLServerStatementListener) ExitIndexName(ctx *IndexNameContext) {}

// EnterConstraintName is called when production constraintName is entered.
func (s *BaseSQLServerStatementListener) EnterConstraintName(ctx *ConstraintNameContext) {}

// ExitConstraintName is called when production constraintName is exited.
func (s *BaseSQLServerStatementListener) ExitConstraintName(ctx *ConstraintNameContext) {}

// EnterCollationName is called when production collationName is entered.
func (s *BaseSQLServerStatementListener) EnterCollationName(ctx *CollationNameContext) {}

// ExitCollationName is called when production collationName is exited.
func (s *BaseSQLServerStatementListener) ExitCollationName(ctx *CollationNameContext) {}

// EnterDataTypeLength is called when production dataTypeLength is entered.
func (s *BaseSQLServerStatementListener) EnterDataTypeLength(ctx *DataTypeLengthContext) {}

// ExitDataTypeLength is called when production dataTypeLength is exited.
func (s *BaseSQLServerStatementListener) ExitDataTypeLength(ctx *DataTypeLengthContext) {}

// EnterPrimaryKey is called when production primaryKey is entered.
func (s *BaseSQLServerStatementListener) EnterPrimaryKey(ctx *PrimaryKeyContext) {}

// ExitPrimaryKey is called when production primaryKey is exited.
func (s *BaseSQLServerStatementListener) ExitPrimaryKey(ctx *PrimaryKeyContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSQLServerStatementListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSQLServerStatementListener) ExitExpr(ctx *ExprContext) {}

// EnterAndOperator is called when production andOperator is entered.
func (s *BaseSQLServerStatementListener) EnterAndOperator(ctx *AndOperatorContext) {}

// ExitAndOperator is called when production andOperator is exited.
func (s *BaseSQLServerStatementListener) ExitAndOperator(ctx *AndOperatorContext) {}

// EnterOrOperator is called when production orOperator is entered.
func (s *BaseSQLServerStatementListener) EnterOrOperator(ctx *OrOperatorContext) {}

// ExitOrOperator is called when production orOperator is exited.
func (s *BaseSQLServerStatementListener) ExitOrOperator(ctx *OrOperatorContext) {}

// EnterNotOperator is called when production notOperator is entered.
func (s *BaseSQLServerStatementListener) EnterNotOperator(ctx *NotOperatorContext) {}

// ExitNotOperator is called when production notOperator is exited.
func (s *BaseSQLServerStatementListener) ExitNotOperator(ctx *NotOperatorContext) {}

// EnterBooleanPrimary is called when production booleanPrimary is entered.
func (s *BaseSQLServerStatementListener) EnterBooleanPrimary(ctx *BooleanPrimaryContext) {}

// ExitBooleanPrimary is called when production booleanPrimary is exited.
func (s *BaseSQLServerStatementListener) ExitBooleanPrimary(ctx *BooleanPrimaryContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseSQLServerStatementListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseSQLServerStatementListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseSQLServerStatementListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseSQLServerStatementListener) ExitPredicate(ctx *PredicateContext) {}

// EnterBitExpr is called when production bitExpr is entered.
func (s *BaseSQLServerStatementListener) EnterBitExpr(ctx *BitExprContext) {}

// ExitBitExpr is called when production bitExpr is exited.
func (s *BaseSQLServerStatementListener) ExitBitExpr(ctx *BitExprContext) {}

// EnterSimpleExpr is called when production simpleExpr is entered.
func (s *BaseSQLServerStatementListener) EnterSimpleExpr(ctx *SimpleExprContext) {}

// ExitSimpleExpr is called when production simpleExpr is exited.
func (s *BaseSQLServerStatementListener) ExitSimpleExpr(ctx *SimpleExprContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseSQLServerStatementListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseSQLServerStatementListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterAggregationFunction is called when production aggregationFunction is entered.
func (s *BaseSQLServerStatementListener) EnterAggregationFunction(ctx *AggregationFunctionContext) {}

// ExitAggregationFunction is called when production aggregationFunction is exited.
func (s *BaseSQLServerStatementListener) ExitAggregationFunction(ctx *AggregationFunctionContext) {}

// EnterAggregationFunctionName is called when production aggregationFunctionName is entered.
func (s *BaseSQLServerStatementListener) EnterAggregationFunctionName(ctx *AggregationFunctionNameContext) {
}

// ExitAggregationFunctionName is called when production aggregationFunctionName is exited.
func (s *BaseSQLServerStatementListener) ExitAggregationFunctionName(ctx *AggregationFunctionNameContext) {
}

// EnterDistinct is called when production distinct is entered.
func (s *BaseSQLServerStatementListener) EnterDistinct(ctx *DistinctContext) {}

// ExitDistinct is called when production distinct is exited.
func (s *BaseSQLServerStatementListener) ExitDistinct(ctx *DistinctContext) {}

// EnterSpecialFunction is called when production specialFunction is entered.
func (s *BaseSQLServerStatementListener) EnterSpecialFunction(ctx *SpecialFunctionContext) {}

// ExitSpecialFunction is called when production specialFunction is exited.
func (s *BaseSQLServerStatementListener) ExitSpecialFunction(ctx *SpecialFunctionContext) {}

// EnterCastFunction is called when production castFunction is entered.
func (s *BaseSQLServerStatementListener) EnterCastFunction(ctx *CastFunctionContext) {}

// ExitCastFunction is called when production castFunction is exited.
func (s *BaseSQLServerStatementListener) ExitCastFunction(ctx *CastFunctionContext) {}

// EnterCharFunction is called when production charFunction is entered.
func (s *BaseSQLServerStatementListener) EnterCharFunction(ctx *CharFunctionContext) {}

// ExitCharFunction is called when production charFunction is exited.
func (s *BaseSQLServerStatementListener) ExitCharFunction(ctx *CharFunctionContext) {}

// EnterRegularFunction is called when production regularFunction is entered.
func (s *BaseSQLServerStatementListener) EnterRegularFunction(ctx *RegularFunctionContext) {}

// ExitRegularFunction is called when production regularFunction is exited.
func (s *BaseSQLServerStatementListener) ExitRegularFunction(ctx *RegularFunctionContext) {}

// EnterRegularFunctionName is called when production regularFunctionName is entered.
func (s *BaseSQLServerStatementListener) EnterRegularFunctionName(ctx *RegularFunctionNameContext) {}

// ExitRegularFunctionName is called when production regularFunctionName is exited.
func (s *BaseSQLServerStatementListener) ExitRegularFunctionName(ctx *RegularFunctionNameContext) {}

// EnterCaseExpression is called when production caseExpression is entered.
func (s *BaseSQLServerStatementListener) EnterCaseExpression(ctx *CaseExpressionContext) {}

// ExitCaseExpression is called when production caseExpression is exited.
func (s *BaseSQLServerStatementListener) ExitCaseExpression(ctx *CaseExpressionContext) {}

// EnterCaseWhen is called when production caseWhen is entered.
func (s *BaseSQLServerStatementListener) EnterCaseWhen(ctx *CaseWhenContext) {}

// ExitCaseWhen is called when production caseWhen is exited.
func (s *BaseSQLServerStatementListener) ExitCaseWhen(ctx *CaseWhenContext) {}

// EnterCaseElse is called when production caseElse is entered.
func (s *BaseSQLServerStatementListener) EnterCaseElse(ctx *CaseElseContext) {}

// ExitCaseElse is called when production caseElse is exited.
func (s *BaseSQLServerStatementListener) ExitCaseElse(ctx *CaseElseContext) {}

// EnterPrivateExprOfDb is called when production privateExprOfDb is entered.
func (s *BaseSQLServerStatementListener) EnterPrivateExprOfDb(ctx *PrivateExprOfDbContext) {}

// ExitPrivateExprOfDb is called when production privateExprOfDb is exited.
func (s *BaseSQLServerStatementListener) ExitPrivateExprOfDb(ctx *PrivateExprOfDbContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseSQLServerStatementListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseSQLServerStatementListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderByItem is called when production orderByItem is entered.
func (s *BaseSQLServerStatementListener) EnterOrderByItem(ctx *OrderByItemContext) {}

// ExitOrderByItem is called when production orderByItem is exited.
func (s *BaseSQLServerStatementListener) ExitOrderByItem(ctx *OrderByItemContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseSQLServerStatementListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseSQLServerStatementListener) ExitDataType(ctx *DataTypeContext) {}

// EnterDataTypeName is called when production dataTypeName is entered.
func (s *BaseSQLServerStatementListener) EnterDataTypeName(ctx *DataTypeNameContext) {}

// ExitDataTypeName is called when production dataTypeName is exited.
func (s *BaseSQLServerStatementListener) ExitDataTypeName(ctx *DataTypeNameContext) {}

// EnterAtTimeZoneExpr is called when production atTimeZoneExpr is entered.
func (s *BaseSQLServerStatementListener) EnterAtTimeZoneExpr(ctx *AtTimeZoneExprContext) {}

// ExitAtTimeZoneExpr is called when production atTimeZoneExpr is exited.
func (s *BaseSQLServerStatementListener) ExitAtTimeZoneExpr(ctx *AtTimeZoneExprContext) {}

// EnterCastExpr is called when production castExpr is entered.
func (s *BaseSQLServerStatementListener) EnterCastExpr(ctx *CastExprContext) {}

// ExitCastExpr is called when production castExpr is exited.
func (s *BaseSQLServerStatementListener) ExitCastExpr(ctx *CastExprContext) {}

// EnterConvertExpr is called when production convertExpr is entered.
func (s *BaseSQLServerStatementListener) EnterConvertExpr(ctx *ConvertExprContext) {}

// ExitConvertExpr is called when production convertExpr is exited.
func (s *BaseSQLServerStatementListener) ExitConvertExpr(ctx *ConvertExprContext) {}

// EnterWindowedFunction is called when production windowedFunction is entered.
func (s *BaseSQLServerStatementListener) EnterWindowedFunction(ctx *WindowedFunctionContext) {}

// ExitWindowedFunction is called when production windowedFunction is exited.
func (s *BaseSQLServerStatementListener) ExitWindowedFunction(ctx *WindowedFunctionContext) {}

// EnterOverClause is called when production overClause is entered.
func (s *BaseSQLServerStatementListener) EnterOverClause(ctx *OverClauseContext) {}

// ExitOverClause is called when production overClause is exited.
func (s *BaseSQLServerStatementListener) ExitOverClause(ctx *OverClauseContext) {}

// EnterPartitionByClause is called when production partitionByClause is entered.
func (s *BaseSQLServerStatementListener) EnterPartitionByClause(ctx *PartitionByClauseContext) {}

// ExitPartitionByClause is called when production partitionByClause is exited.
func (s *BaseSQLServerStatementListener) ExitPartitionByClause(ctx *PartitionByClauseContext) {}

// EnterRowRangeClause is called when production rowRangeClause is entered.
func (s *BaseSQLServerStatementListener) EnterRowRangeClause(ctx *RowRangeClauseContext) {}

// ExitRowRangeClause is called when production rowRangeClause is exited.
func (s *BaseSQLServerStatementListener) ExitRowRangeClause(ctx *RowRangeClauseContext) {}

// EnterWindowFrameExtent is called when production windowFrameExtent is entered.
func (s *BaseSQLServerStatementListener) EnterWindowFrameExtent(ctx *WindowFrameExtentContext) {}

// ExitWindowFrameExtent is called when production windowFrameExtent is exited.
func (s *BaseSQLServerStatementListener) ExitWindowFrameExtent(ctx *WindowFrameExtentContext) {}

// EnterWindowFrameBetween is called when production windowFrameBetween is entered.
func (s *BaseSQLServerStatementListener) EnterWindowFrameBetween(ctx *WindowFrameBetweenContext) {}

// ExitWindowFrameBetween is called when production windowFrameBetween is exited.
func (s *BaseSQLServerStatementListener) ExitWindowFrameBetween(ctx *WindowFrameBetweenContext) {}

// EnterWindowFrameBound is called when production windowFrameBound is entered.
func (s *BaseSQLServerStatementListener) EnterWindowFrameBound(ctx *WindowFrameBoundContext) {}

// ExitWindowFrameBound is called when production windowFrameBound is exited.
func (s *BaseSQLServerStatementListener) ExitWindowFrameBound(ctx *WindowFrameBoundContext) {}

// EnterWindowFramePreceding is called when production windowFramePreceding is entered.
func (s *BaseSQLServerStatementListener) EnterWindowFramePreceding(ctx *WindowFramePrecedingContext) {
}

// ExitWindowFramePreceding is called when production windowFramePreceding is exited.
func (s *BaseSQLServerStatementListener) ExitWindowFramePreceding(ctx *WindowFramePrecedingContext) {}

// EnterWindowFrameFollowing is called when production windowFrameFollowing is entered.
func (s *BaseSQLServerStatementListener) EnterWindowFrameFollowing(ctx *WindowFrameFollowingContext) {
}

// ExitWindowFrameFollowing is called when production windowFrameFollowing is exited.
func (s *BaseSQLServerStatementListener) ExitWindowFrameFollowing(ctx *WindowFrameFollowingContext) {}

// EnterColumnNameWithSort is called when production columnNameWithSort is entered.
func (s *BaseSQLServerStatementListener) EnterColumnNameWithSort(ctx *ColumnNameWithSortContext) {}

// ExitColumnNameWithSort is called when production columnNameWithSort is exited.
func (s *BaseSQLServerStatementListener) ExitColumnNameWithSort(ctx *ColumnNameWithSortContext) {}

// EnterIndexOption is called when production indexOption is entered.
func (s *BaseSQLServerStatementListener) EnterIndexOption(ctx *IndexOptionContext) {}

// ExitIndexOption is called when production indexOption is exited.
func (s *BaseSQLServerStatementListener) ExitIndexOption(ctx *IndexOptionContext) {}

// EnterCompressionOption is called when production compressionOption is entered.
func (s *BaseSQLServerStatementListener) EnterCompressionOption(ctx *CompressionOptionContext) {}

// ExitCompressionOption is called when production compressionOption is exited.
func (s *BaseSQLServerStatementListener) ExitCompressionOption(ctx *CompressionOptionContext) {}

// EnterEqTime is called when production eqTime is entered.
func (s *BaseSQLServerStatementListener) EnterEqTime(ctx *EqTimeContext) {}

// ExitEqTime is called when production eqTime is exited.
func (s *BaseSQLServerStatementListener) ExitEqTime(ctx *EqTimeContext) {}

// EnterEqOnOffOption is called when production eqOnOffOption is entered.
func (s *BaseSQLServerStatementListener) EnterEqOnOffOption(ctx *EqOnOffOptionContext) {}

// ExitEqOnOffOption is called when production eqOnOffOption is exited.
func (s *BaseSQLServerStatementListener) ExitEqOnOffOption(ctx *EqOnOffOptionContext) {}

// EnterEqKey is called when production eqKey is entered.
func (s *BaseSQLServerStatementListener) EnterEqKey(ctx *EqKeyContext) {}

// ExitEqKey is called when production eqKey is exited.
func (s *BaseSQLServerStatementListener) ExitEqKey(ctx *EqKeyContext) {}

// EnterEqOnOff is called when production eqOnOff is entered.
func (s *BaseSQLServerStatementListener) EnterEqOnOff(ctx *EqOnOffContext) {}

// ExitEqOnOff is called when production eqOnOff is exited.
func (s *BaseSQLServerStatementListener) ExitEqOnOff(ctx *EqOnOffContext) {}

// EnterOnPartitionClause is called when production onPartitionClause is entered.
func (s *BaseSQLServerStatementListener) EnterOnPartitionClause(ctx *OnPartitionClauseContext) {}

// ExitOnPartitionClause is called when production onPartitionClause is exited.
func (s *BaseSQLServerStatementListener) ExitOnPartitionClause(ctx *OnPartitionClauseContext) {}

// EnterPartitionExpressions is called when production partitionExpressions is entered.
func (s *BaseSQLServerStatementListener) EnterPartitionExpressions(ctx *PartitionExpressionsContext) {
}

// ExitPartitionExpressions is called when production partitionExpressions is exited.
func (s *BaseSQLServerStatementListener) ExitPartitionExpressions(ctx *PartitionExpressionsContext) {}

// EnterPartitionExpression is called when production partitionExpression is entered.
func (s *BaseSQLServerStatementListener) EnterPartitionExpression(ctx *PartitionExpressionContext) {}

// ExitPartitionExpression is called when production partitionExpression is exited.
func (s *BaseSQLServerStatementListener) ExitPartitionExpression(ctx *PartitionExpressionContext) {}

// EnterNumberRange is called when production numberRange is entered.
func (s *BaseSQLServerStatementListener) EnterNumberRange(ctx *NumberRangeContext) {}

// ExitNumberRange is called when production numberRange is exited.
func (s *BaseSQLServerStatementListener) ExitNumberRange(ctx *NumberRangeContext) {}

// EnterLowPriorityLockWait is called when production lowPriorityLockWait is entered.
func (s *BaseSQLServerStatementListener) EnterLowPriorityLockWait(ctx *LowPriorityLockWaitContext) {}

// ExitLowPriorityLockWait is called when production lowPriorityLockWait is exited.
func (s *BaseSQLServerStatementListener) ExitLowPriorityLockWait(ctx *LowPriorityLockWaitContext) {}

// EnterOnLowPriorLockWait is called when production onLowPriorLockWait is entered.
func (s *BaseSQLServerStatementListener) EnterOnLowPriorLockWait(ctx *OnLowPriorLockWaitContext) {}

// ExitOnLowPriorLockWait is called when production onLowPriorLockWait is exited.
func (s *BaseSQLServerStatementListener) ExitOnLowPriorLockWait(ctx *OnLowPriorLockWaitContext) {}

// EnterIgnoredIdentifier is called when production ignoredIdentifier is entered.
func (s *BaseSQLServerStatementListener) EnterIgnoredIdentifier(ctx *IgnoredIdentifierContext) {}

// ExitIgnoredIdentifier is called when production ignoredIdentifier is exited.
func (s *BaseSQLServerStatementListener) ExitIgnoredIdentifier(ctx *IgnoredIdentifierContext) {}

// EnterIgnoredIdentifiers is called when production ignoredIdentifiers is entered.
func (s *BaseSQLServerStatementListener) EnterIgnoredIdentifiers(ctx *IgnoredIdentifiersContext) {}

// ExitIgnoredIdentifiers is called when production ignoredIdentifiers is exited.
func (s *BaseSQLServerStatementListener) ExitIgnoredIdentifiers(ctx *IgnoredIdentifiersContext) {}

// EnterMatchNone is called when production matchNone is entered.
func (s *BaseSQLServerStatementListener) EnterMatchNone(ctx *MatchNoneContext) {}

// ExitMatchNone is called when production matchNone is exited.
func (s *BaseSQLServerStatementListener) ExitMatchNone(ctx *MatchNoneContext) {}

// EnterVariableName is called when production variableName is entered.
func (s *BaseSQLServerStatementListener) EnterVariableName(ctx *VariableNameContext) {}

// ExitVariableName is called when production variableName is exited.
func (s *BaseSQLServerStatementListener) ExitVariableName(ctx *VariableNameContext) {}

// EnterExecuteAsClause is called when production executeAsClause is entered.
func (s *BaseSQLServerStatementListener) EnterExecuteAsClause(ctx *ExecuteAsClauseContext) {}

// ExitExecuteAsClause is called when production executeAsClause is exited.
func (s *BaseSQLServerStatementListener) ExitExecuteAsClause(ctx *ExecuteAsClauseContext) {}

// EnterTransactionName is called when production transactionName is entered.
func (s *BaseSQLServerStatementListener) EnterTransactionName(ctx *TransactionNameContext) {}

// ExitTransactionName is called when production transactionName is exited.
func (s *BaseSQLServerStatementListener) ExitTransactionName(ctx *TransactionNameContext) {}

// EnterTransactionVariableName is called when production transactionVariableName is entered.
func (s *BaseSQLServerStatementListener) EnterTransactionVariableName(ctx *TransactionVariableNameContext) {
}

// ExitTransactionVariableName is called when production transactionVariableName is exited.
func (s *BaseSQLServerStatementListener) ExitTransactionVariableName(ctx *TransactionVariableNameContext) {
}

// EnterSavepointName is called when production savepointName is entered.
func (s *BaseSQLServerStatementListener) EnterSavepointName(ctx *SavepointNameContext) {}

// ExitSavepointName is called when production savepointName is exited.
func (s *BaseSQLServerStatementListener) ExitSavepointName(ctx *SavepointNameContext) {}

// EnterSavepointVariableName is called when production savepointVariableName is entered.
func (s *BaseSQLServerStatementListener) EnterSavepointVariableName(ctx *SavepointVariableNameContext) {
}

// ExitSavepointVariableName is called when production savepointVariableName is exited.
func (s *BaseSQLServerStatementListener) ExitSavepointVariableName(ctx *SavepointVariableNameContext) {
}

// EnterEntityType is called when production entityType is entered.
func (s *BaseSQLServerStatementListener) EnterEntityType(ctx *EntityTypeContext) {}

// ExitEntityType is called when production entityType is exited.
func (s *BaseSQLServerStatementListener) ExitEntityType(ctx *EntityTypeContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseSQLServerStatementListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseSQLServerStatementListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterCreateTableClause is called when production createTableClause is entered.
func (s *BaseSQLServerStatementListener) EnterCreateTableClause(ctx *CreateTableClauseContext) {}

// ExitCreateTableClause is called when production createTableClause is exited.
func (s *BaseSQLServerStatementListener) ExitCreateTableClause(ctx *CreateTableClauseContext) {}

// EnterCreateIndex is called when production createIndex is entered.
func (s *BaseSQLServerStatementListener) EnterCreateIndex(ctx *CreateIndexContext) {}

// ExitCreateIndex is called when production createIndex is exited.
func (s *BaseSQLServerStatementListener) ExitCreateIndex(ctx *CreateIndexContext) {}

// EnterCreateDatabase is called when production createDatabase is entered.
func (s *BaseSQLServerStatementListener) EnterCreateDatabase(ctx *CreateDatabaseContext) {}

// ExitCreateDatabase is called when production createDatabase is exited.
func (s *BaseSQLServerStatementListener) ExitCreateDatabase(ctx *CreateDatabaseContext) {}

// EnterCreateFunction is called when production createFunction is entered.
func (s *BaseSQLServerStatementListener) EnterCreateFunction(ctx *CreateFunctionContext) {}

// ExitCreateFunction is called when production createFunction is exited.
func (s *BaseSQLServerStatementListener) ExitCreateFunction(ctx *CreateFunctionContext) {}

// EnterCreateProcedure is called when production createProcedure is entered.
func (s *BaseSQLServerStatementListener) EnterCreateProcedure(ctx *CreateProcedureContext) {}

// ExitCreateProcedure is called when production createProcedure is exited.
func (s *BaseSQLServerStatementListener) ExitCreateProcedure(ctx *CreateProcedureContext) {}

// EnterCreateView is called when production createView is entered.
func (s *BaseSQLServerStatementListener) EnterCreateView(ctx *CreateViewContext) {}

// ExitCreateView is called when production createView is exited.
func (s *BaseSQLServerStatementListener) ExitCreateView(ctx *CreateViewContext) {}

// EnterCreateTrigger is called when production createTrigger is entered.
func (s *BaseSQLServerStatementListener) EnterCreateTrigger(ctx *CreateTriggerContext) {}

// ExitCreateTrigger is called when production createTrigger is exited.
func (s *BaseSQLServerStatementListener) ExitCreateTrigger(ctx *CreateTriggerContext) {}

// EnterCreateSequence is called when production createSequence is entered.
func (s *BaseSQLServerStatementListener) EnterCreateSequence(ctx *CreateSequenceContext) {}

// ExitCreateSequence is called when production createSequence is exited.
func (s *BaseSQLServerStatementListener) ExitCreateSequence(ctx *CreateSequenceContext) {}

// EnterCreateService is called when production createService is entered.
func (s *BaseSQLServerStatementListener) EnterCreateService(ctx *CreateServiceContext) {}

// ExitCreateService is called when production createService is exited.
func (s *BaseSQLServerStatementListener) ExitCreateService(ctx *CreateServiceContext) {}

// EnterCreateSchema is called when production createSchema is entered.
func (s *BaseSQLServerStatementListener) EnterCreateSchema(ctx *CreateSchemaContext) {}

// ExitCreateSchema is called when production createSchema is exited.
func (s *BaseSQLServerStatementListener) ExitCreateSchema(ctx *CreateSchemaContext) {}

// EnterAlterTable is called when production alterTable is entered.
func (s *BaseSQLServerStatementListener) EnterAlterTable(ctx *AlterTableContext) {}

// ExitAlterTable is called when production alterTable is exited.
func (s *BaseSQLServerStatementListener) ExitAlterTable(ctx *AlterTableContext) {}

// EnterAlterIndex is called when production alterIndex is entered.
func (s *BaseSQLServerStatementListener) EnterAlterIndex(ctx *AlterIndexContext) {}

// ExitAlterIndex is called when production alterIndex is exited.
func (s *BaseSQLServerStatementListener) ExitAlterIndex(ctx *AlterIndexContext) {}

// EnterAlterDatabase is called when production alterDatabase is entered.
func (s *BaseSQLServerStatementListener) EnterAlterDatabase(ctx *AlterDatabaseContext) {}

// ExitAlterDatabase is called when production alterDatabase is exited.
func (s *BaseSQLServerStatementListener) ExitAlterDatabase(ctx *AlterDatabaseContext) {}

// EnterAlterProcedure is called when production alterProcedure is entered.
func (s *BaseSQLServerStatementListener) EnterAlterProcedure(ctx *AlterProcedureContext) {}

// ExitAlterProcedure is called when production alterProcedure is exited.
func (s *BaseSQLServerStatementListener) ExitAlterProcedure(ctx *AlterProcedureContext) {}

// EnterAlterFunction is called when production alterFunction is entered.
func (s *BaseSQLServerStatementListener) EnterAlterFunction(ctx *AlterFunctionContext) {}

// ExitAlterFunction is called when production alterFunction is exited.
func (s *BaseSQLServerStatementListener) ExitAlterFunction(ctx *AlterFunctionContext) {}

// EnterAlterView is called when production alterView is entered.
func (s *BaseSQLServerStatementListener) EnterAlterView(ctx *AlterViewContext) {}

// ExitAlterView is called when production alterView is exited.
func (s *BaseSQLServerStatementListener) ExitAlterView(ctx *AlterViewContext) {}

// EnterAlterTrigger is called when production alterTrigger is entered.
func (s *BaseSQLServerStatementListener) EnterAlterTrigger(ctx *AlterTriggerContext) {}

// ExitAlterTrigger is called when production alterTrigger is exited.
func (s *BaseSQLServerStatementListener) ExitAlterTrigger(ctx *AlterTriggerContext) {}

// EnterAlterSequence is called when production alterSequence is entered.
func (s *BaseSQLServerStatementListener) EnterAlterSequence(ctx *AlterSequenceContext) {}

// ExitAlterSequence is called when production alterSequence is exited.
func (s *BaseSQLServerStatementListener) ExitAlterSequence(ctx *AlterSequenceContext) {}

// EnterAlterService is called when production alterService is entered.
func (s *BaseSQLServerStatementListener) EnterAlterService(ctx *AlterServiceContext) {}

// ExitAlterService is called when production alterService is exited.
func (s *BaseSQLServerStatementListener) ExitAlterService(ctx *AlterServiceContext) {}

// EnterAlterSchema is called when production alterSchema is entered.
func (s *BaseSQLServerStatementListener) EnterAlterSchema(ctx *AlterSchemaContext) {}

// ExitAlterSchema is called when production alterSchema is exited.
func (s *BaseSQLServerStatementListener) ExitAlterSchema(ctx *AlterSchemaContext) {}

// EnterDropTable is called when production dropTable is entered.
func (s *BaseSQLServerStatementListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production dropTable is exited.
func (s *BaseSQLServerStatementListener) ExitDropTable(ctx *DropTableContext) {}

// EnterDropIndex is called when production dropIndex is entered.
func (s *BaseSQLServerStatementListener) EnterDropIndex(ctx *DropIndexContext) {}

// ExitDropIndex is called when production dropIndex is exited.
func (s *BaseSQLServerStatementListener) ExitDropIndex(ctx *DropIndexContext) {}

// EnterDropDatabase is called when production dropDatabase is entered.
func (s *BaseSQLServerStatementListener) EnterDropDatabase(ctx *DropDatabaseContext) {}

// ExitDropDatabase is called when production dropDatabase is exited.
func (s *BaseSQLServerStatementListener) ExitDropDatabase(ctx *DropDatabaseContext) {}

// EnterDropFunction is called when production dropFunction is entered.
func (s *BaseSQLServerStatementListener) EnterDropFunction(ctx *DropFunctionContext) {}

// ExitDropFunction is called when production dropFunction is exited.
func (s *BaseSQLServerStatementListener) ExitDropFunction(ctx *DropFunctionContext) {}

// EnterDropProcedure is called when production dropProcedure is entered.
func (s *BaseSQLServerStatementListener) EnterDropProcedure(ctx *DropProcedureContext) {}

// ExitDropProcedure is called when production dropProcedure is exited.
func (s *BaseSQLServerStatementListener) ExitDropProcedure(ctx *DropProcedureContext) {}

// EnterDropView is called when production dropView is entered.
func (s *BaseSQLServerStatementListener) EnterDropView(ctx *DropViewContext) {}

// ExitDropView is called when production dropView is exited.
func (s *BaseSQLServerStatementListener) ExitDropView(ctx *DropViewContext) {}

// EnterDropTrigger is called when production dropTrigger is entered.
func (s *BaseSQLServerStatementListener) EnterDropTrigger(ctx *DropTriggerContext) {}

// ExitDropTrigger is called when production dropTrigger is exited.
func (s *BaseSQLServerStatementListener) ExitDropTrigger(ctx *DropTriggerContext) {}

// EnterDropSequence is called when production dropSequence is entered.
func (s *BaseSQLServerStatementListener) EnterDropSequence(ctx *DropSequenceContext) {}

// ExitDropSequence is called when production dropSequence is exited.
func (s *BaseSQLServerStatementListener) ExitDropSequence(ctx *DropSequenceContext) {}

// EnterDropService is called when production dropService is entered.
func (s *BaseSQLServerStatementListener) EnterDropService(ctx *DropServiceContext) {}

// ExitDropService is called when production dropService is exited.
func (s *BaseSQLServerStatementListener) ExitDropService(ctx *DropServiceContext) {}

// EnterDropSchema is called when production dropSchema is entered.
func (s *BaseSQLServerStatementListener) EnterDropSchema(ctx *DropSchemaContext) {}

// ExitDropSchema is called when production dropSchema is exited.
func (s *BaseSQLServerStatementListener) ExitDropSchema(ctx *DropSchemaContext) {}

// EnterTruncateTable is called when production truncateTable is entered.
func (s *BaseSQLServerStatementListener) EnterTruncateTable(ctx *TruncateTableContext) {}

// ExitTruncateTable is called when production truncateTable is exited.
func (s *BaseSQLServerStatementListener) ExitTruncateTable(ctx *TruncateTableContext) {}

// EnterFileTableClause is called when production fileTableClause is entered.
func (s *BaseSQLServerStatementListener) EnterFileTableClause(ctx *FileTableClauseContext) {}

// ExitFileTableClause is called when production fileTableClause is exited.
func (s *BaseSQLServerStatementListener) ExitFileTableClause(ctx *FileTableClauseContext) {}

// EnterCreateDefinitionClause is called when production createDefinitionClause is entered.
func (s *BaseSQLServerStatementListener) EnterCreateDefinitionClause(ctx *CreateDefinitionClauseContext) {
}

// ExitCreateDefinitionClause is called when production createDefinitionClause is exited.
func (s *BaseSQLServerStatementListener) ExitCreateDefinitionClause(ctx *CreateDefinitionClauseContext) {
}

// EnterCreateTableDefinitions is called when production createTableDefinitions is entered.
func (s *BaseSQLServerStatementListener) EnterCreateTableDefinitions(ctx *CreateTableDefinitionsContext) {
}

// ExitCreateTableDefinitions is called when production createTableDefinitions is exited.
func (s *BaseSQLServerStatementListener) ExitCreateTableDefinitions(ctx *CreateTableDefinitionsContext) {
}

// EnterCreateTableDefinition is called when production createTableDefinition is entered.
func (s *BaseSQLServerStatementListener) EnterCreateTableDefinition(ctx *CreateTableDefinitionContext) {
}

// ExitCreateTableDefinition is called when production createTableDefinition is exited.
func (s *BaseSQLServerStatementListener) ExitCreateTableDefinition(ctx *CreateTableDefinitionContext) {
}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseSQLServerStatementListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseSQLServerStatementListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterColumnDefinitionOption is called when production columnDefinitionOption is entered.
func (s *BaseSQLServerStatementListener) EnterColumnDefinitionOption(ctx *ColumnDefinitionOptionContext) {
}

// ExitColumnDefinitionOption is called when production columnDefinitionOption is exited.
func (s *BaseSQLServerStatementListener) ExitColumnDefinitionOption(ctx *ColumnDefinitionOptionContext) {
}

// EnterEncryptedOptions is called when production encryptedOptions is entered.
func (s *BaseSQLServerStatementListener) EnterEncryptedOptions(ctx *EncryptedOptionsContext) {}

// ExitEncryptedOptions is called when production encryptedOptions is exited.
func (s *BaseSQLServerStatementListener) ExitEncryptedOptions(ctx *EncryptedOptionsContext) {}

// EnterColumnConstraint is called when production columnConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterColumnConstraint(ctx *ColumnConstraintContext) {}

// ExitColumnConstraint is called when production columnConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitColumnConstraint(ctx *ColumnConstraintContext) {}

// EnterComputedColumnConstraint is called when production computedColumnConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterComputedColumnConstraint(ctx *ComputedColumnConstraintContext) {
}

// ExitComputedColumnConstraint is called when production computedColumnConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitComputedColumnConstraint(ctx *ComputedColumnConstraintContext) {
}

// EnterComputedColumnForeignKeyConstraint is called when production computedColumnForeignKeyConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterComputedColumnForeignKeyConstraint(ctx *ComputedColumnForeignKeyConstraintContext) {
}

// ExitComputedColumnForeignKeyConstraint is called when production computedColumnForeignKeyConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitComputedColumnForeignKeyConstraint(ctx *ComputedColumnForeignKeyConstraintContext) {
}

// EnterComputedColumnForeignKeyOnAction is called when production computedColumnForeignKeyOnAction is entered.
func (s *BaseSQLServerStatementListener) EnterComputedColumnForeignKeyOnAction(ctx *ComputedColumnForeignKeyOnActionContext) {
}

// ExitComputedColumnForeignKeyOnAction is called when production computedColumnForeignKeyOnAction is exited.
func (s *BaseSQLServerStatementListener) ExitComputedColumnForeignKeyOnAction(ctx *ComputedColumnForeignKeyOnActionContext) {
}

// EnterPrimaryKeyConstraint is called when production primaryKeyConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterPrimaryKeyConstraint(ctx *PrimaryKeyConstraintContext) {
}

// ExitPrimaryKeyConstraint is called when production primaryKeyConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitPrimaryKeyConstraint(ctx *PrimaryKeyConstraintContext) {}

// EnterDiskTablePrimaryKeyConstraintOption is called when production diskTablePrimaryKeyConstraintOption is entered.
func (s *BaseSQLServerStatementListener) EnterDiskTablePrimaryKeyConstraintOption(ctx *DiskTablePrimaryKeyConstraintOptionContext) {
}

// ExitDiskTablePrimaryKeyConstraintOption is called when production diskTablePrimaryKeyConstraintOption is exited.
func (s *BaseSQLServerStatementListener) ExitDiskTablePrimaryKeyConstraintOption(ctx *DiskTablePrimaryKeyConstraintOptionContext) {
}

// EnterClusterOption is called when production clusterOption is entered.
func (s *BaseSQLServerStatementListener) EnterClusterOption(ctx *ClusterOptionContext) {}

// ExitClusterOption is called when production clusterOption is exited.
func (s *BaseSQLServerStatementListener) ExitClusterOption(ctx *ClusterOptionContext) {}

// EnterPrimaryKeyWithClause is called when production primaryKeyWithClause is entered.
func (s *BaseSQLServerStatementListener) EnterPrimaryKeyWithClause(ctx *PrimaryKeyWithClauseContext) {
}

// ExitPrimaryKeyWithClause is called when production primaryKeyWithClause is exited.
func (s *BaseSQLServerStatementListener) ExitPrimaryKeyWithClause(ctx *PrimaryKeyWithClauseContext) {}

// EnterPrimaryKeyOnClause is called when production primaryKeyOnClause is entered.
func (s *BaseSQLServerStatementListener) EnterPrimaryKeyOnClause(ctx *PrimaryKeyOnClauseContext) {}

// ExitPrimaryKeyOnClause is called when production primaryKeyOnClause is exited.
func (s *BaseSQLServerStatementListener) ExitPrimaryKeyOnClause(ctx *PrimaryKeyOnClauseContext) {}

// EnterOnSchemaColumn is called when production onSchemaColumn is entered.
func (s *BaseSQLServerStatementListener) EnterOnSchemaColumn(ctx *OnSchemaColumnContext) {}

// ExitOnSchemaColumn is called when production onSchemaColumn is exited.
func (s *BaseSQLServerStatementListener) ExitOnSchemaColumn(ctx *OnSchemaColumnContext) {}

// EnterOnFileGroup is called when production onFileGroup is entered.
func (s *BaseSQLServerStatementListener) EnterOnFileGroup(ctx *OnFileGroupContext) {}

// ExitOnFileGroup is called when production onFileGroup is exited.
func (s *BaseSQLServerStatementListener) ExitOnFileGroup(ctx *OnFileGroupContext) {}

// EnterOnString is called when production onString is entered.
func (s *BaseSQLServerStatementListener) EnterOnString(ctx *OnStringContext) {}

// ExitOnString is called when production onString is exited.
func (s *BaseSQLServerStatementListener) ExitOnString(ctx *OnStringContext) {}

// EnterMemoryTablePrimaryKeyConstraintOption is called when production memoryTablePrimaryKeyConstraintOption is entered.
func (s *BaseSQLServerStatementListener) EnterMemoryTablePrimaryKeyConstraintOption(ctx *MemoryTablePrimaryKeyConstraintOptionContext) {
}

// ExitMemoryTablePrimaryKeyConstraintOption is called when production memoryTablePrimaryKeyConstraintOption is exited.
func (s *BaseSQLServerStatementListener) ExitMemoryTablePrimaryKeyConstraintOption(ctx *MemoryTablePrimaryKeyConstraintOptionContext) {
}

// EnterWithBucket is called when production withBucket is entered.
func (s *BaseSQLServerStatementListener) EnterWithBucket(ctx *WithBucketContext) {}

// ExitWithBucket is called when production withBucket is exited.
func (s *BaseSQLServerStatementListener) ExitWithBucket(ctx *WithBucketContext) {}

// EnterColumnForeignKeyConstraint is called when production columnForeignKeyConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterColumnForeignKeyConstraint(ctx *ColumnForeignKeyConstraintContext) {
}

// ExitColumnForeignKeyConstraint is called when production columnForeignKeyConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitColumnForeignKeyConstraint(ctx *ColumnForeignKeyConstraintContext) {
}

// EnterForeignKeyOnAction is called when production foreignKeyOnAction is entered.
func (s *BaseSQLServerStatementListener) EnterForeignKeyOnAction(ctx *ForeignKeyOnActionContext) {}

// ExitForeignKeyOnAction is called when production foreignKeyOnAction is exited.
func (s *BaseSQLServerStatementListener) ExitForeignKeyOnAction(ctx *ForeignKeyOnActionContext) {}

// EnterForeignKeyOn is called when production foreignKeyOn is entered.
func (s *BaseSQLServerStatementListener) EnterForeignKeyOn(ctx *ForeignKeyOnContext) {}

// ExitForeignKeyOn is called when production foreignKeyOn is exited.
func (s *BaseSQLServerStatementListener) ExitForeignKeyOn(ctx *ForeignKeyOnContext) {}

// EnterCheckConstraint is called when production checkConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterCheckConstraint(ctx *CheckConstraintContext) {}

// ExitCheckConstraint is called when production checkConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitCheckConstraint(ctx *CheckConstraintContext) {}

// EnterColumnIndex is called when production columnIndex is entered.
func (s *BaseSQLServerStatementListener) EnterColumnIndex(ctx *ColumnIndexContext) {}

// ExitColumnIndex is called when production columnIndex is exited.
func (s *BaseSQLServerStatementListener) ExitColumnIndex(ctx *ColumnIndexContext) {}

// EnterWithIndexOption is called when production withIndexOption is entered.
func (s *BaseSQLServerStatementListener) EnterWithIndexOption(ctx *WithIndexOptionContext) {}

// ExitWithIndexOption is called when production withIndexOption is exited.
func (s *BaseSQLServerStatementListener) ExitWithIndexOption(ctx *WithIndexOptionContext) {}

// EnterIndexOnClause is called when production indexOnClause is entered.
func (s *BaseSQLServerStatementListener) EnterIndexOnClause(ctx *IndexOnClauseContext) {}

// ExitIndexOnClause is called when production indexOnClause is exited.
func (s *BaseSQLServerStatementListener) ExitIndexOnClause(ctx *IndexOnClauseContext) {}

// EnterOnDefault is called when production onDefault is entered.
func (s *BaseSQLServerStatementListener) EnterOnDefault(ctx *OnDefaultContext) {}

// ExitOnDefault is called when production onDefault is exited.
func (s *BaseSQLServerStatementListener) ExitOnDefault(ctx *OnDefaultContext) {}

// EnterFileStreamOn is called when production fileStreamOn is entered.
func (s *BaseSQLServerStatementListener) EnterFileStreamOn(ctx *FileStreamOnContext) {}

// ExitFileStreamOn is called when production fileStreamOn is exited.
func (s *BaseSQLServerStatementListener) ExitFileStreamOn(ctx *FileStreamOnContext) {}

// EnterColumnConstraints is called when production columnConstraints is entered.
func (s *BaseSQLServerStatementListener) EnterColumnConstraints(ctx *ColumnConstraintsContext) {}

// ExitColumnConstraints is called when production columnConstraints is exited.
func (s *BaseSQLServerStatementListener) ExitColumnConstraints(ctx *ColumnConstraintsContext) {}

// EnterComputedColumnDefinition is called when production computedColumnDefinition is entered.
func (s *BaseSQLServerStatementListener) EnterComputedColumnDefinition(ctx *ComputedColumnDefinitionContext) {
}

// ExitComputedColumnDefinition is called when production computedColumnDefinition is exited.
func (s *BaseSQLServerStatementListener) ExitComputedColumnDefinition(ctx *ComputedColumnDefinitionContext) {
}

// EnterColumnSetDefinition is called when production columnSetDefinition is entered.
func (s *BaseSQLServerStatementListener) EnterColumnSetDefinition(ctx *ColumnSetDefinitionContext) {}

// ExitColumnSetDefinition is called when production columnSetDefinition is exited.
func (s *BaseSQLServerStatementListener) ExitColumnSetDefinition(ctx *ColumnSetDefinitionContext) {}

// EnterTableConstraint is called when production tableConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterTableConstraint(ctx *TableConstraintContext) {}

// ExitTableConstraint is called when production tableConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitTableConstraint(ctx *TableConstraintContext) {}

// EnterTablePrimaryConstraint is called when production tablePrimaryConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterTablePrimaryConstraint(ctx *TablePrimaryConstraintContext) {
}

// ExitTablePrimaryConstraint is called when production tablePrimaryConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitTablePrimaryConstraint(ctx *TablePrimaryConstraintContext) {
}

// EnterPrimaryKeyUnique is called when production primaryKeyUnique is entered.
func (s *BaseSQLServerStatementListener) EnterPrimaryKeyUnique(ctx *PrimaryKeyUniqueContext) {}

// ExitPrimaryKeyUnique is called when production primaryKeyUnique is exited.
func (s *BaseSQLServerStatementListener) ExitPrimaryKeyUnique(ctx *PrimaryKeyUniqueContext) {}

// EnterDiskTablePrimaryConstraintOption is called when production diskTablePrimaryConstraintOption is entered.
func (s *BaseSQLServerStatementListener) EnterDiskTablePrimaryConstraintOption(ctx *DiskTablePrimaryConstraintOptionContext) {
}

// ExitDiskTablePrimaryConstraintOption is called when production diskTablePrimaryConstraintOption is exited.
func (s *BaseSQLServerStatementListener) ExitDiskTablePrimaryConstraintOption(ctx *DiskTablePrimaryConstraintOptionContext) {
}

// EnterMemoryTablePrimaryConstraintOption is called when production memoryTablePrimaryConstraintOption is entered.
func (s *BaseSQLServerStatementListener) EnterMemoryTablePrimaryConstraintOption(ctx *MemoryTablePrimaryConstraintOptionContext) {
}

// ExitMemoryTablePrimaryConstraintOption is called when production memoryTablePrimaryConstraintOption is exited.
func (s *BaseSQLServerStatementListener) ExitMemoryTablePrimaryConstraintOption(ctx *MemoryTablePrimaryConstraintOptionContext) {
}

// EnterHashWithBucket is called when production hashWithBucket is entered.
func (s *BaseSQLServerStatementListener) EnterHashWithBucket(ctx *HashWithBucketContext) {}

// ExitHashWithBucket is called when production hashWithBucket is exited.
func (s *BaseSQLServerStatementListener) ExitHashWithBucket(ctx *HashWithBucketContext) {}

// EnterTableForeignKeyConstraint is called when production tableForeignKeyConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterTableForeignKeyConstraint(ctx *TableForeignKeyConstraintContext) {
}

// ExitTableForeignKeyConstraint is called when production tableForeignKeyConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitTableForeignKeyConstraint(ctx *TableForeignKeyConstraintContext) {
}

// EnterTableIndex is called when production tableIndex is entered.
func (s *BaseSQLServerStatementListener) EnterTableIndex(ctx *TableIndexContext) {}

// ExitTableIndex is called when production tableIndex is exited.
func (s *BaseSQLServerStatementListener) ExitTableIndex(ctx *TableIndexContext) {}

// EnterIndexNameOption is called when production indexNameOption is entered.
func (s *BaseSQLServerStatementListener) EnterIndexNameOption(ctx *IndexNameOptionContext) {}

// ExitIndexNameOption is called when production indexNameOption is exited.
func (s *BaseSQLServerStatementListener) ExitIndexNameOption(ctx *IndexNameOptionContext) {}

// EnterIndexOptions is called when production indexOptions is entered.
func (s *BaseSQLServerStatementListener) EnterIndexOptions(ctx *IndexOptionsContext) {}

// ExitIndexOptions is called when production indexOptions is exited.
func (s *BaseSQLServerStatementListener) ExitIndexOptions(ctx *IndexOptionsContext) {}

// EnterPeriodClause is called when production periodClause is entered.
func (s *BaseSQLServerStatementListener) EnterPeriodClause(ctx *PeriodClauseContext) {}

// ExitPeriodClause is called when production periodClause is exited.
func (s *BaseSQLServerStatementListener) ExitPeriodClause(ctx *PeriodClauseContext) {}

// EnterPartitionScheme is called when production partitionScheme is entered.
func (s *BaseSQLServerStatementListener) EnterPartitionScheme(ctx *PartitionSchemeContext) {}

// ExitPartitionScheme is called when production partitionScheme is exited.
func (s *BaseSQLServerStatementListener) ExitPartitionScheme(ctx *PartitionSchemeContext) {}

// EnterFileGroup is called when production fileGroup is entered.
func (s *BaseSQLServerStatementListener) EnterFileGroup(ctx *FileGroupContext) {}

// ExitFileGroup is called when production fileGroup is exited.
func (s *BaseSQLServerStatementListener) ExitFileGroup(ctx *FileGroupContext) {}

// EnterTableOptions is called when production tableOptions is entered.
func (s *BaseSQLServerStatementListener) EnterTableOptions(ctx *TableOptionsContext) {}

// ExitTableOptions is called when production tableOptions is exited.
func (s *BaseSQLServerStatementListener) ExitTableOptions(ctx *TableOptionsContext) {}

// EnterTableOption is called when production tableOption is entered.
func (s *BaseSQLServerStatementListener) EnterTableOption(ctx *TableOptionContext) {}

// ExitTableOption is called when production tableOption is exited.
func (s *BaseSQLServerStatementListener) ExitTableOption(ctx *TableOptionContext) {}

// EnterDataDelectionOption is called when production dataDelectionOption is entered.
func (s *BaseSQLServerStatementListener) EnterDataDelectionOption(ctx *DataDelectionOptionContext) {}

// ExitDataDelectionOption is called when production dataDelectionOption is exited.
func (s *BaseSQLServerStatementListener) ExitDataDelectionOption(ctx *DataDelectionOptionContext) {}

// EnterTableStretchOptions is called when production tableStretchOptions is entered.
func (s *BaseSQLServerStatementListener) EnterTableStretchOptions(ctx *TableStretchOptionsContext) {}

// ExitTableStretchOptions is called when production tableStretchOptions is exited.
func (s *BaseSQLServerStatementListener) ExitTableStretchOptions(ctx *TableStretchOptionsContext) {}

// EnterTableStretchOption is called when production tableStretchOption is entered.
func (s *BaseSQLServerStatementListener) EnterTableStretchOption(ctx *TableStretchOptionContext) {}

// ExitTableStretchOption is called when production tableStretchOption is exited.
func (s *BaseSQLServerStatementListener) ExitTableStretchOption(ctx *TableStretchOptionContext) {}

// EnterMigrationState_ is called when production migrationState_ is entered.
func (s *BaseSQLServerStatementListener) EnterMigrationState_(ctx *MigrationState_Context) {}

// ExitMigrationState_ is called when production migrationState_ is exited.
func (s *BaseSQLServerStatementListener) ExitMigrationState_(ctx *MigrationState_Context) {}

// EnterTableOperationOption is called when production tableOperationOption is entered.
func (s *BaseSQLServerStatementListener) EnterTableOperationOption(ctx *TableOperationOptionContext) {
}

// ExitTableOperationOption is called when production tableOperationOption is exited.
func (s *BaseSQLServerStatementListener) ExitTableOperationOption(ctx *TableOperationOptionContext) {}

// EnterDistributionOption is called when production distributionOption is entered.
func (s *BaseSQLServerStatementListener) EnterDistributionOption(ctx *DistributionOptionContext) {}

// ExitDistributionOption is called when production distributionOption is exited.
func (s *BaseSQLServerStatementListener) ExitDistributionOption(ctx *DistributionOptionContext) {}

// EnterDataWareHouseTableOption is called when production dataWareHouseTableOption is entered.
func (s *BaseSQLServerStatementListener) EnterDataWareHouseTableOption(ctx *DataWareHouseTableOptionContext) {
}

// ExitDataWareHouseTableOption is called when production dataWareHouseTableOption is exited.
func (s *BaseSQLServerStatementListener) ExitDataWareHouseTableOption(ctx *DataWareHouseTableOptionContext) {
}

// EnterDataWareHousePartitionOption is called when production dataWareHousePartitionOption is entered.
func (s *BaseSQLServerStatementListener) EnterDataWareHousePartitionOption(ctx *DataWareHousePartitionOptionContext) {
}

// ExitDataWareHousePartitionOption is called when production dataWareHousePartitionOption is exited.
func (s *BaseSQLServerStatementListener) ExitDataWareHousePartitionOption(ctx *DataWareHousePartitionOptionContext) {
}

// EnterCreateIndexSpecification is called when production createIndexSpecification is entered.
func (s *BaseSQLServerStatementListener) EnterCreateIndexSpecification(ctx *CreateIndexSpecificationContext) {
}

// ExitCreateIndexSpecification is called when production createIndexSpecification is exited.
func (s *BaseSQLServerStatementListener) ExitCreateIndexSpecification(ctx *CreateIndexSpecificationContext) {
}

// EnterAlterDefinitionClause is called when production alterDefinitionClause is entered.
func (s *BaseSQLServerStatementListener) EnterAlterDefinitionClause(ctx *AlterDefinitionClauseContext) {
}

// ExitAlterDefinitionClause is called when production alterDefinitionClause is exited.
func (s *BaseSQLServerStatementListener) ExitAlterDefinitionClause(ctx *AlterDefinitionClauseContext) {
}

// EnterAddColumnSpecification is called when production addColumnSpecification is entered.
func (s *BaseSQLServerStatementListener) EnterAddColumnSpecification(ctx *AddColumnSpecificationContext) {
}

// ExitAddColumnSpecification is called when production addColumnSpecification is exited.
func (s *BaseSQLServerStatementListener) ExitAddColumnSpecification(ctx *AddColumnSpecificationContext) {
}

// EnterModifyColumnSpecification is called when production modifyColumnSpecification is entered.
func (s *BaseSQLServerStatementListener) EnterModifyColumnSpecification(ctx *ModifyColumnSpecificationContext) {
}

// ExitModifyColumnSpecification is called when production modifyColumnSpecification is exited.
func (s *BaseSQLServerStatementListener) ExitModifyColumnSpecification(ctx *ModifyColumnSpecificationContext) {
}

// EnterAlterColumnOperation is called when production alterColumnOperation is entered.
func (s *BaseSQLServerStatementListener) EnterAlterColumnOperation(ctx *AlterColumnOperationContext) {
}

// ExitAlterColumnOperation is called when production alterColumnOperation is exited.
func (s *BaseSQLServerStatementListener) ExitAlterColumnOperation(ctx *AlterColumnOperationContext) {}

// EnterAlterColumnAddOptions is called when production alterColumnAddOptions is entered.
func (s *BaseSQLServerStatementListener) EnterAlterColumnAddOptions(ctx *AlterColumnAddOptionsContext) {
}

// ExitAlterColumnAddOptions is called when production alterColumnAddOptions is exited.
func (s *BaseSQLServerStatementListener) ExitAlterColumnAddOptions(ctx *AlterColumnAddOptionsContext) {
}

// EnterAlterColumnAddOption is called when production alterColumnAddOption is entered.
func (s *BaseSQLServerStatementListener) EnterAlterColumnAddOption(ctx *AlterColumnAddOptionContext) {
}

// ExitAlterColumnAddOption is called when production alterColumnAddOption is exited.
func (s *BaseSQLServerStatementListener) ExitAlterColumnAddOption(ctx *AlterColumnAddOptionContext) {}

// EnterConstraintForColumn is called when production constraintForColumn is entered.
func (s *BaseSQLServerStatementListener) EnterConstraintForColumn(ctx *ConstraintForColumnContext) {}

// ExitConstraintForColumn is called when production constraintForColumn is exited.
func (s *BaseSQLServerStatementListener) ExitConstraintForColumn(ctx *ConstraintForColumnContext) {}

// EnterGeneratedColumnNamesClause is called when production generatedColumnNamesClause is entered.
func (s *BaseSQLServerStatementListener) EnterGeneratedColumnNamesClause(ctx *GeneratedColumnNamesClauseContext) {
}

// ExitGeneratedColumnNamesClause is called when production generatedColumnNamesClause is exited.
func (s *BaseSQLServerStatementListener) ExitGeneratedColumnNamesClause(ctx *GeneratedColumnNamesClauseContext) {
}

// EnterGeneratedColumnNameClause is called when production generatedColumnNameClause is entered.
func (s *BaseSQLServerStatementListener) EnterGeneratedColumnNameClause(ctx *GeneratedColumnNameClauseContext) {
}

// ExitGeneratedColumnNameClause is called when production generatedColumnNameClause is exited.
func (s *BaseSQLServerStatementListener) ExitGeneratedColumnNameClause(ctx *GeneratedColumnNameClauseContext) {
}

// EnterGeneratedColumnName is called when production generatedColumnName is entered.
func (s *BaseSQLServerStatementListener) EnterGeneratedColumnName(ctx *GeneratedColumnNameContext) {}

// ExitGeneratedColumnName is called when production generatedColumnName is exited.
func (s *BaseSQLServerStatementListener) ExitGeneratedColumnName(ctx *GeneratedColumnNameContext) {}

// EnterAlterDrop is called when production alterDrop is entered.
func (s *BaseSQLServerStatementListener) EnterAlterDrop(ctx *AlterDropContext) {}

// ExitAlterDrop is called when production alterDrop is exited.
func (s *BaseSQLServerStatementListener) ExitAlterDrop(ctx *AlterDropContext) {}

// EnterAlterTableDropConstraint is called when production alterTableDropConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterAlterTableDropConstraint(ctx *AlterTableDropConstraintContext) {
}

// ExitAlterTableDropConstraint is called when production alterTableDropConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitAlterTableDropConstraint(ctx *AlterTableDropConstraintContext) {
}

// EnterDropConstraintName is called when production dropConstraintName is entered.
func (s *BaseSQLServerStatementListener) EnterDropConstraintName(ctx *DropConstraintNameContext) {}

// ExitDropConstraintName is called when production dropConstraintName is exited.
func (s *BaseSQLServerStatementListener) ExitDropConstraintName(ctx *DropConstraintNameContext) {}

// EnterDropConstraintWithClause is called when production dropConstraintWithClause is entered.
func (s *BaseSQLServerStatementListener) EnterDropConstraintWithClause(ctx *DropConstraintWithClauseContext) {
}

// ExitDropConstraintWithClause is called when production dropConstraintWithClause is exited.
func (s *BaseSQLServerStatementListener) ExitDropConstraintWithClause(ctx *DropConstraintWithClauseContext) {
}

// EnterDropConstraintOption is called when production dropConstraintOption is entered.
func (s *BaseSQLServerStatementListener) EnterDropConstraintOption(ctx *DropConstraintOptionContext) {
}

// ExitDropConstraintOption is called when production dropConstraintOption is exited.
func (s *BaseSQLServerStatementListener) ExitDropConstraintOption(ctx *DropConstraintOptionContext) {}

// EnterOnOffOption is called when production onOffOption is entered.
func (s *BaseSQLServerStatementListener) EnterOnOffOption(ctx *OnOffOptionContext) {}

// ExitOnOffOption is called when production onOffOption is exited.
func (s *BaseSQLServerStatementListener) ExitOnOffOption(ctx *OnOffOptionContext) {}

// EnterDropColumnSpecification is called when production dropColumnSpecification is entered.
func (s *BaseSQLServerStatementListener) EnterDropColumnSpecification(ctx *DropColumnSpecificationContext) {
}

// ExitDropColumnSpecification is called when production dropColumnSpecification is exited.
func (s *BaseSQLServerStatementListener) ExitDropColumnSpecification(ctx *DropColumnSpecificationContext) {
}

// EnterDropIndexSpecification is called when production dropIndexSpecification is entered.
func (s *BaseSQLServerStatementListener) EnterDropIndexSpecification(ctx *DropIndexSpecificationContext) {
}

// ExitDropIndexSpecification is called when production dropIndexSpecification is exited.
func (s *BaseSQLServerStatementListener) ExitDropIndexSpecification(ctx *DropIndexSpecificationContext) {
}

// EnterAlterCheckConstraint is called when production alterCheckConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterAlterCheckConstraint(ctx *AlterCheckConstraintContext) {
}

// ExitAlterCheckConstraint is called when production alterCheckConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitAlterCheckConstraint(ctx *AlterCheckConstraintContext) {}

// EnterAlterTableTrigger is called when production alterTableTrigger is entered.
func (s *BaseSQLServerStatementListener) EnterAlterTableTrigger(ctx *AlterTableTriggerContext) {}

// ExitAlterTableTrigger is called when production alterTableTrigger is exited.
func (s *BaseSQLServerStatementListener) ExitAlterTableTrigger(ctx *AlterTableTriggerContext) {}

// EnterAlterSwitch is called when production alterSwitch is entered.
func (s *BaseSQLServerStatementListener) EnterAlterSwitch(ctx *AlterSwitchContext) {}

// ExitAlterSwitch is called when production alterSwitch is exited.
func (s *BaseSQLServerStatementListener) ExitAlterSwitch(ctx *AlterSwitchContext) {}

// EnterAlterSet is called when production alterSet is entered.
func (s *BaseSQLServerStatementListener) EnterAlterSet(ctx *AlterSetContext) {}

// ExitAlterSet is called when production alterSet is exited.
func (s *BaseSQLServerStatementListener) ExitAlterSet(ctx *AlterSetContext) {}

// EnterSetFileStreamClause is called when production setFileStreamClause is entered.
func (s *BaseSQLServerStatementListener) EnterSetFileStreamClause(ctx *SetFileStreamClauseContext) {}

// ExitSetFileStreamClause is called when production setFileStreamClause is exited.
func (s *BaseSQLServerStatementListener) ExitSetFileStreamClause(ctx *SetFileStreamClauseContext) {}

// EnterSetSystemVersionClause is called when production setSystemVersionClause is entered.
func (s *BaseSQLServerStatementListener) EnterSetSystemVersionClause(ctx *SetSystemVersionClauseContext) {
}

// ExitSetSystemVersionClause is called when production setSystemVersionClause is exited.
func (s *BaseSQLServerStatementListener) ExitSetSystemVersionClause(ctx *SetSystemVersionClauseContext) {
}

// EnterAlterSetOnClause is called when production alterSetOnClause is entered.
func (s *BaseSQLServerStatementListener) EnterAlterSetOnClause(ctx *AlterSetOnClauseContext) {}

// ExitAlterSetOnClause is called when production alterSetOnClause is exited.
func (s *BaseSQLServerStatementListener) ExitAlterSetOnClause(ctx *AlterSetOnClauseContext) {}

// EnterDataConsistencyCheckClause is called when production dataConsistencyCheckClause is entered.
func (s *BaseSQLServerStatementListener) EnterDataConsistencyCheckClause(ctx *DataConsistencyCheckClauseContext) {
}

// ExitDataConsistencyCheckClause is called when production dataConsistencyCheckClause is exited.
func (s *BaseSQLServerStatementListener) ExitDataConsistencyCheckClause(ctx *DataConsistencyCheckClauseContext) {
}

// EnterHistoryRetentionPeriodClause is called when production historyRetentionPeriodClause is entered.
func (s *BaseSQLServerStatementListener) EnterHistoryRetentionPeriodClause(ctx *HistoryRetentionPeriodClauseContext) {
}

// ExitHistoryRetentionPeriodClause is called when production historyRetentionPeriodClause is exited.
func (s *BaseSQLServerStatementListener) ExitHistoryRetentionPeriodClause(ctx *HistoryRetentionPeriodClauseContext) {
}

// EnterHistoryRetentionPeriod is called when production historyRetentionPeriod is entered.
func (s *BaseSQLServerStatementListener) EnterHistoryRetentionPeriod(ctx *HistoryRetentionPeriodContext) {
}

// ExitHistoryRetentionPeriod is called when production historyRetentionPeriod is exited.
func (s *BaseSQLServerStatementListener) ExitHistoryRetentionPeriod(ctx *HistoryRetentionPeriodContext) {
}

// EnterAlterTableTableIndex is called when production alterTableTableIndex is entered.
func (s *BaseSQLServerStatementListener) EnterAlterTableTableIndex(ctx *AlterTableTableIndexContext) {
}

// ExitAlterTableTableIndex is called when production alterTableTableIndex is exited.
func (s *BaseSQLServerStatementListener) ExitAlterTableTableIndex(ctx *AlterTableTableIndexContext) {}

// EnterIndexWithName is called when production indexWithName is entered.
func (s *BaseSQLServerStatementListener) EnterIndexWithName(ctx *IndexWithNameContext) {}

// ExitIndexWithName is called when production indexWithName is exited.
func (s *BaseSQLServerStatementListener) ExitIndexWithName(ctx *IndexWithNameContext) {}

// EnterIndexNonClusterClause is called when production indexNonClusterClause is entered.
func (s *BaseSQLServerStatementListener) EnterIndexNonClusterClause(ctx *IndexNonClusterClauseContext) {
}

// ExitIndexNonClusterClause is called when production indexNonClusterClause is exited.
func (s *BaseSQLServerStatementListener) ExitIndexNonClusterClause(ctx *IndexNonClusterClauseContext) {
}

// EnterAlterTableIndexOnClause is called when production alterTableIndexOnClause is entered.
func (s *BaseSQLServerStatementListener) EnterAlterTableIndexOnClause(ctx *AlterTableIndexOnClauseContext) {
}

// ExitAlterTableIndexOnClause is called when production alterTableIndexOnClause is exited.
func (s *BaseSQLServerStatementListener) ExitAlterTableIndexOnClause(ctx *AlterTableIndexOnClauseContext) {
}

// EnterIndexClusterClause is called when production indexClusterClause is entered.
func (s *BaseSQLServerStatementListener) EnterIndexClusterClause(ctx *IndexClusterClauseContext) {}

// ExitIndexClusterClause is called when production indexClusterClause is exited.
func (s *BaseSQLServerStatementListener) ExitIndexClusterClause(ctx *IndexClusterClauseContext) {}

// EnterAlterTableOption is called when production alterTableOption is entered.
func (s *BaseSQLServerStatementListener) EnterAlterTableOption(ctx *AlterTableOptionContext) {}

// ExitAlterTableOption is called when production alterTableOption is exited.
func (s *BaseSQLServerStatementListener) ExitAlterTableOption(ctx *AlterTableOptionContext) {}

// EnterOnHistoryTableClause is called when production onHistoryTableClause is entered.
func (s *BaseSQLServerStatementListener) EnterOnHistoryTableClause(ctx *OnHistoryTableClauseContext) {
}

// ExitOnHistoryTableClause is called when production onHistoryTableClause is exited.
func (s *BaseSQLServerStatementListener) ExitOnHistoryTableClause(ctx *OnHistoryTableClauseContext) {}

// EnterIfExist is called when production ifExist is entered.
func (s *BaseSQLServerStatementListener) EnterIfExist(ctx *IfExistContext) {}

// ExitIfExist is called when production ifExist is exited.
func (s *BaseSQLServerStatementListener) ExitIfExist(ctx *IfExistContext) {}

// EnterCreateDatabaseClause is called when production createDatabaseClause is entered.
func (s *BaseSQLServerStatementListener) EnterCreateDatabaseClause(ctx *CreateDatabaseClauseContext) {
}

// ExitCreateDatabaseClause is called when production createDatabaseClause is exited.
func (s *BaseSQLServerStatementListener) ExitCreateDatabaseClause(ctx *CreateDatabaseClauseContext) {}

// EnterFileDefinitionClause is called when production fileDefinitionClause is entered.
func (s *BaseSQLServerStatementListener) EnterFileDefinitionClause(ctx *FileDefinitionClauseContext) {
}

// ExitFileDefinitionClause is called when production fileDefinitionClause is exited.
func (s *BaseSQLServerStatementListener) ExitFileDefinitionClause(ctx *FileDefinitionClauseContext) {}

// EnterDatabaseOption is called when production databaseOption is entered.
func (s *BaseSQLServerStatementListener) EnterDatabaseOption(ctx *DatabaseOptionContext) {}

// ExitDatabaseOption is called when production databaseOption is exited.
func (s *BaseSQLServerStatementListener) ExitDatabaseOption(ctx *DatabaseOptionContext) {}

// EnterFileStreamOption is called when production fileStreamOption is entered.
func (s *BaseSQLServerStatementListener) EnterFileStreamOption(ctx *FileStreamOptionContext) {}

// ExitFileStreamOption is called when production fileStreamOption is exited.
func (s *BaseSQLServerStatementListener) ExitFileStreamOption(ctx *FileStreamOptionContext) {}

// EnterFileSpec is called when production fileSpec is entered.
func (s *BaseSQLServerStatementListener) EnterFileSpec(ctx *FileSpecContext) {}

// ExitFileSpec is called when production fileSpec is exited.
func (s *BaseSQLServerStatementListener) ExitFileSpec(ctx *FileSpecContext) {}

// EnterDatabaseFileSpecOption is called when production databaseFileSpecOption is entered.
func (s *BaseSQLServerStatementListener) EnterDatabaseFileSpecOption(ctx *DatabaseFileSpecOptionContext) {
}

// ExitDatabaseFileSpecOption is called when production databaseFileSpecOption is exited.
func (s *BaseSQLServerStatementListener) ExitDatabaseFileSpecOption(ctx *DatabaseFileSpecOptionContext) {
}

// EnterDatabaseFileGroup is called when production databaseFileGroup is entered.
func (s *BaseSQLServerStatementListener) EnterDatabaseFileGroup(ctx *DatabaseFileGroupContext) {}

// ExitDatabaseFileGroup is called when production databaseFileGroup is exited.
func (s *BaseSQLServerStatementListener) ExitDatabaseFileGroup(ctx *DatabaseFileGroupContext) {}

// EnterDatabaseFileGroupContains is called when production databaseFileGroupContains is entered.
func (s *BaseSQLServerStatementListener) EnterDatabaseFileGroupContains(ctx *DatabaseFileGroupContainsContext) {
}

// ExitDatabaseFileGroupContains is called when production databaseFileGroupContains is exited.
func (s *BaseSQLServerStatementListener) ExitDatabaseFileGroupContains(ctx *DatabaseFileGroupContainsContext) {
}

// EnterDatabaseLogOns is called when production databaseLogOns is entered.
func (s *BaseSQLServerStatementListener) EnterDatabaseLogOns(ctx *DatabaseLogOnsContext) {}

// ExitDatabaseLogOns is called when production databaseLogOns is exited.
func (s *BaseSQLServerStatementListener) ExitDatabaseLogOns(ctx *DatabaseLogOnsContext) {}

// EnterDeclareVariable is called when production declareVariable is entered.
func (s *BaseSQLServerStatementListener) EnterDeclareVariable(ctx *DeclareVariableContext) {}

// ExitDeclareVariable is called when production declareVariable is exited.
func (s *BaseSQLServerStatementListener) ExitDeclareVariable(ctx *DeclareVariableContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseSQLServerStatementListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseSQLServerStatementListener) ExitVariable(ctx *VariableContext) {}

// EnterTableVariable is called when production tableVariable is entered.
func (s *BaseSQLServerStatementListener) EnterTableVariable(ctx *TableVariableContext) {}

// ExitTableVariable is called when production tableVariable is exited.
func (s *BaseSQLServerStatementListener) ExitTableVariable(ctx *TableVariableContext) {}

// EnterVariTableTypeDefinition is called when production variTableTypeDefinition is entered.
func (s *BaseSQLServerStatementListener) EnterVariTableTypeDefinition(ctx *VariTableTypeDefinitionContext) {
}

// ExitVariTableTypeDefinition is called when production variTableTypeDefinition is exited.
func (s *BaseSQLServerStatementListener) ExitVariTableTypeDefinition(ctx *VariTableTypeDefinitionContext) {
}

// EnterTableVariableClause is called when production tableVariableClause is entered.
func (s *BaseSQLServerStatementListener) EnterTableVariableClause(ctx *TableVariableClauseContext) {}

// ExitTableVariableClause is called when production tableVariableClause is exited.
func (s *BaseSQLServerStatementListener) ExitTableVariableClause(ctx *TableVariableClauseContext) {}

// EnterVariableTableColumnDefinition is called when production variableTableColumnDefinition is entered.
func (s *BaseSQLServerStatementListener) EnterVariableTableColumnDefinition(ctx *VariableTableColumnDefinitionContext) {
}

// ExitVariableTableColumnDefinition is called when production variableTableColumnDefinition is exited.
func (s *BaseSQLServerStatementListener) ExitVariableTableColumnDefinition(ctx *VariableTableColumnDefinitionContext) {
}

// EnterVariableTableColumnConstraint is called when production variableTableColumnConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterVariableTableColumnConstraint(ctx *VariableTableColumnConstraintContext) {
}

// ExitVariableTableColumnConstraint is called when production variableTableColumnConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitVariableTableColumnConstraint(ctx *VariableTableColumnConstraintContext) {
}

// EnterVariableTableConstraint is called when production variableTableConstraint is entered.
func (s *BaseSQLServerStatementListener) EnterVariableTableConstraint(ctx *VariableTableConstraintContext) {
}

// ExitVariableTableConstraint is called when production variableTableConstraint is exited.
func (s *BaseSQLServerStatementListener) ExitVariableTableConstraint(ctx *VariableTableConstraintContext) {
}

// EnterSetVariable is called when production setVariable is entered.
func (s *BaseSQLServerStatementListener) EnterSetVariable(ctx *SetVariableContext) {}

// ExitSetVariable is called when production setVariable is exited.
func (s *BaseSQLServerStatementListener) ExitSetVariable(ctx *SetVariableContext) {}

// EnterSetVariableClause is called when production setVariableClause is entered.
func (s *BaseSQLServerStatementListener) EnterSetVariableClause(ctx *SetVariableClauseContext) {}

// ExitSetVariableClause is called when production setVariableClause is exited.
func (s *BaseSQLServerStatementListener) ExitSetVariableClause(ctx *SetVariableClauseContext) {}

// EnterCursorVariable is called when production cursorVariable is entered.
func (s *BaseSQLServerStatementListener) EnterCursorVariable(ctx *CursorVariableContext) {}

// ExitCursorVariable is called when production cursorVariable is exited.
func (s *BaseSQLServerStatementListener) ExitCursorVariable(ctx *CursorVariableContext) {}

// EnterCursorClause is called when production cursorClause is entered.
func (s *BaseSQLServerStatementListener) EnterCursorClause(ctx *CursorClauseContext) {}

// ExitCursorClause is called when production cursorClause is exited.
func (s *BaseSQLServerStatementListener) ExitCursorClause(ctx *CursorClauseContext) {}

// EnterCompoundOperation is called when production compoundOperation is entered.
func (s *BaseSQLServerStatementListener) EnterCompoundOperation(ctx *CompoundOperationContext) {}

// ExitCompoundOperation is called when production compoundOperation is exited.
func (s *BaseSQLServerStatementListener) ExitCompoundOperation(ctx *CompoundOperationContext) {}

// EnterFuncParameters is called when production funcParameters is entered.
func (s *BaseSQLServerStatementListener) EnterFuncParameters(ctx *FuncParametersContext) {}

// ExitFuncParameters is called when production funcParameters is exited.
func (s *BaseSQLServerStatementListener) ExitFuncParameters(ctx *FuncParametersContext) {}

// EnterFuncReturns is called when production funcReturns is entered.
func (s *BaseSQLServerStatementListener) EnterFuncReturns(ctx *FuncReturnsContext) {}

// ExitFuncReturns is called when production funcReturns is exited.
func (s *BaseSQLServerStatementListener) ExitFuncReturns(ctx *FuncReturnsContext) {}

// EnterFuncMutiReturn is called when production funcMutiReturn is entered.
func (s *BaseSQLServerStatementListener) EnterFuncMutiReturn(ctx *FuncMutiReturnContext) {}

// ExitFuncMutiReturn is called when production funcMutiReturn is exited.
func (s *BaseSQLServerStatementListener) ExitFuncMutiReturn(ctx *FuncMutiReturnContext) {}

// EnterFuncInlineReturn is called when production funcInlineReturn is entered.
func (s *BaseSQLServerStatementListener) EnterFuncInlineReturn(ctx *FuncInlineReturnContext) {}

// ExitFuncInlineReturn is called when production funcInlineReturn is exited.
func (s *BaseSQLServerStatementListener) ExitFuncInlineReturn(ctx *FuncInlineReturnContext) {}

// EnterFuncScalarReturn is called when production funcScalarReturn is entered.
func (s *BaseSQLServerStatementListener) EnterFuncScalarReturn(ctx *FuncScalarReturnContext) {}

// ExitFuncScalarReturn is called when production funcScalarReturn is exited.
func (s *BaseSQLServerStatementListener) ExitFuncScalarReturn(ctx *FuncScalarReturnContext) {}

// EnterTableTypeDefinition is called when production tableTypeDefinition is entered.
func (s *BaseSQLServerStatementListener) EnterTableTypeDefinition(ctx *TableTypeDefinitionContext) {}

// ExitTableTypeDefinition is called when production tableTypeDefinition is exited.
func (s *BaseSQLServerStatementListener) ExitTableTypeDefinition(ctx *TableTypeDefinitionContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseSQLServerStatementListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseSQLServerStatementListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterFunctionOption is called when production functionOption is entered.
func (s *BaseSQLServerStatementListener) EnterFunctionOption(ctx *FunctionOptionContext) {}

// ExitFunctionOption is called when production functionOption is exited.
func (s *BaseSQLServerStatementListener) ExitFunctionOption(ctx *FunctionOptionContext) {}

// EnterValidStatement is called when production validStatement is entered.
func (s *BaseSQLServerStatementListener) EnterValidStatement(ctx *ValidStatementContext) {}

// ExitValidStatement is called when production validStatement is exited.
func (s *BaseSQLServerStatementListener) ExitValidStatement(ctx *ValidStatementContext) {}

// EnterProcParameters is called when production procParameters is entered.
func (s *BaseSQLServerStatementListener) EnterProcParameters(ctx *ProcParametersContext) {}

// ExitProcParameters is called when production procParameters is exited.
func (s *BaseSQLServerStatementListener) ExitProcParameters(ctx *ProcParametersContext) {}

// EnterProcParameter is called when production procParameter is entered.
func (s *BaseSQLServerStatementListener) EnterProcParameter(ctx *ProcParameterContext) {}

// ExitProcParameter is called when production procParameter is exited.
func (s *BaseSQLServerStatementListener) ExitProcParameter(ctx *ProcParameterContext) {}

// EnterCreateOrAlterProcClause is called when production createOrAlterProcClause is entered.
func (s *BaseSQLServerStatementListener) EnterCreateOrAlterProcClause(ctx *CreateOrAlterProcClauseContext) {
}

// ExitCreateOrAlterProcClause is called when production createOrAlterProcClause is exited.
func (s *BaseSQLServerStatementListener) ExitCreateOrAlterProcClause(ctx *CreateOrAlterProcClauseContext) {
}

// EnterWithCreateProcOption is called when production withCreateProcOption is entered.
func (s *BaseSQLServerStatementListener) EnterWithCreateProcOption(ctx *WithCreateProcOptionContext) {
}

// ExitWithCreateProcOption is called when production withCreateProcOption is exited.
func (s *BaseSQLServerStatementListener) ExitWithCreateProcOption(ctx *WithCreateProcOptionContext) {}

// EnterProcOption is called when production procOption is entered.
func (s *BaseSQLServerStatementListener) EnterProcOption(ctx *ProcOptionContext) {}

// ExitProcOption is called when production procOption is exited.
func (s *BaseSQLServerStatementListener) ExitProcOption(ctx *ProcOptionContext) {}

// EnterProcAsClause is called when production procAsClause is entered.
func (s *BaseSQLServerStatementListener) EnterProcAsClause(ctx *ProcAsClauseContext) {}

// ExitProcAsClause is called when production procAsClause is exited.
func (s *BaseSQLServerStatementListener) ExitProcAsClause(ctx *ProcAsClauseContext) {}

// EnterProcSetOption is called when production procSetOption is entered.
func (s *BaseSQLServerStatementListener) EnterProcSetOption(ctx *ProcSetOptionContext) {}

// ExitProcSetOption is called when production procSetOption is exited.
func (s *BaseSQLServerStatementListener) ExitProcSetOption(ctx *ProcSetOptionContext) {}

// EnterCreateOrAlterViewClause is called when production createOrAlterViewClause is entered.
func (s *BaseSQLServerStatementListener) EnterCreateOrAlterViewClause(ctx *CreateOrAlterViewClauseContext) {
}

// ExitCreateOrAlterViewClause is called when production createOrAlterViewClause is exited.
func (s *BaseSQLServerStatementListener) ExitCreateOrAlterViewClause(ctx *CreateOrAlterViewClauseContext) {
}

// EnterViewAttribute is called when production viewAttribute is entered.
func (s *BaseSQLServerStatementListener) EnterViewAttribute(ctx *ViewAttributeContext) {}

// ExitViewAttribute is called when production viewAttribute is exited.
func (s *BaseSQLServerStatementListener) ExitViewAttribute(ctx *ViewAttributeContext) {}

// EnterWithCommonTableExpr is called when production withCommonTableExpr is entered.
func (s *BaseSQLServerStatementListener) EnterWithCommonTableExpr(ctx *WithCommonTableExprContext) {}

// ExitWithCommonTableExpr is called when production withCommonTableExpr is exited.
func (s *BaseSQLServerStatementListener) ExitWithCommonTableExpr(ctx *WithCommonTableExprContext) {}

// EnterCommonTableExpr is called when production commonTableExpr is entered.
func (s *BaseSQLServerStatementListener) EnterCommonTableExpr(ctx *CommonTableExprContext) {}

// ExitCommonTableExpr is called when production commonTableExpr is exited.
func (s *BaseSQLServerStatementListener) ExitCommonTableExpr(ctx *CommonTableExprContext) {}

// EnterCreateTriggerClause is called when production createTriggerClause is entered.
func (s *BaseSQLServerStatementListener) EnterCreateTriggerClause(ctx *CreateTriggerClauseContext) {}

// ExitCreateTriggerClause is called when production createTriggerClause is exited.
func (s *BaseSQLServerStatementListener) ExitCreateTriggerClause(ctx *CreateTriggerClauseContext) {}

// EnterDmlTriggerOption is called when production dmlTriggerOption is entered.
func (s *BaseSQLServerStatementListener) EnterDmlTriggerOption(ctx *DmlTriggerOptionContext) {}

// ExitDmlTriggerOption is called when production dmlTriggerOption is exited.
func (s *BaseSQLServerStatementListener) ExitDmlTriggerOption(ctx *DmlTriggerOptionContext) {}

// EnterMethodSpecifier is called when production methodSpecifier is entered.
func (s *BaseSQLServerStatementListener) EnterMethodSpecifier(ctx *MethodSpecifierContext) {}

// ExitMethodSpecifier is called when production methodSpecifier is exited.
func (s *BaseSQLServerStatementListener) ExitMethodSpecifier(ctx *MethodSpecifierContext) {}

// EnterTriggerTarget is called when production triggerTarget is entered.
func (s *BaseSQLServerStatementListener) EnterTriggerTarget(ctx *TriggerTargetContext) {}

// ExitTriggerTarget is called when production triggerTarget is exited.
func (s *BaseSQLServerStatementListener) ExitTriggerTarget(ctx *TriggerTargetContext) {}

// EnterCreateOrAlterSequenceClause is called when production createOrAlterSequenceClause is entered.
func (s *BaseSQLServerStatementListener) EnterCreateOrAlterSequenceClause(ctx *CreateOrAlterSequenceClauseContext) {
}

// ExitCreateOrAlterSequenceClause is called when production createOrAlterSequenceClause is exited.
func (s *BaseSQLServerStatementListener) ExitCreateOrAlterSequenceClause(ctx *CreateOrAlterSequenceClauseContext) {
}

// EnterCreateIndexClause is called when production createIndexClause is entered.
func (s *BaseSQLServerStatementListener) EnterCreateIndexClause(ctx *CreateIndexClauseContext) {}

// ExitCreateIndexClause is called when production createIndexClause is exited.
func (s *BaseSQLServerStatementListener) ExitCreateIndexClause(ctx *CreateIndexClauseContext) {}

// EnterFilterPredicate is called when production filterPredicate is entered.
func (s *BaseSQLServerStatementListener) EnterFilterPredicate(ctx *FilterPredicateContext) {}

// ExitFilterPredicate is called when production filterPredicate is exited.
func (s *BaseSQLServerStatementListener) ExitFilterPredicate(ctx *FilterPredicateContext) {}

// EnterConjunct is called when production conjunct is entered.
func (s *BaseSQLServerStatementListener) EnterConjunct(ctx *ConjunctContext) {}

// ExitConjunct is called when production conjunct is exited.
func (s *BaseSQLServerStatementListener) ExitConjunct(ctx *ConjunctContext) {}

// EnterAlterIndexClause is called when production alterIndexClause is entered.
func (s *BaseSQLServerStatementListener) EnterAlterIndexClause(ctx *AlterIndexClauseContext) {}

// ExitAlterIndexClause is called when production alterIndexClause is exited.
func (s *BaseSQLServerStatementListener) ExitAlterIndexClause(ctx *AlterIndexClauseContext) {}

// EnterRelationalIndexOption is called when production relationalIndexOption is entered.
func (s *BaseSQLServerStatementListener) EnterRelationalIndexOption(ctx *RelationalIndexOptionContext) {
}

// ExitRelationalIndexOption is called when production relationalIndexOption is exited.
func (s *BaseSQLServerStatementListener) ExitRelationalIndexOption(ctx *RelationalIndexOptionContext) {
}

// EnterPartitionNumberRange is called when production partitionNumberRange is entered.
func (s *BaseSQLServerStatementListener) EnterPartitionNumberRange(ctx *PartitionNumberRangeContext) {
}

// ExitPartitionNumberRange is called when production partitionNumberRange is exited.
func (s *BaseSQLServerStatementListener) ExitPartitionNumberRange(ctx *PartitionNumberRangeContext) {}

// EnterReorganizeOption is called when production reorganizeOption is entered.
func (s *BaseSQLServerStatementListener) EnterReorganizeOption(ctx *ReorganizeOptionContext) {}

// ExitReorganizeOption is called when production reorganizeOption is exited.
func (s *BaseSQLServerStatementListener) ExitReorganizeOption(ctx *ReorganizeOptionContext) {}

// EnterSetIndexOption is called when production setIndexOption is entered.
func (s *BaseSQLServerStatementListener) EnterSetIndexOption(ctx *SetIndexOptionContext) {}

// ExitSetIndexOption is called when production setIndexOption is exited.
func (s *BaseSQLServerStatementListener) ExitSetIndexOption(ctx *SetIndexOptionContext) {}

// EnterResumableIndexOptions is called when production resumableIndexOptions is entered.
func (s *BaseSQLServerStatementListener) EnterResumableIndexOptions(ctx *ResumableIndexOptionsContext) {
}

// ExitResumableIndexOptions is called when production resumableIndexOptions is exited.
func (s *BaseSQLServerStatementListener) ExitResumableIndexOptions(ctx *ResumableIndexOptionsContext) {
}

// EnterAlterDatabaseClause is called when production alterDatabaseClause is entered.
func (s *BaseSQLServerStatementListener) EnterAlterDatabaseClause(ctx *AlterDatabaseClauseContext) {}

// ExitAlterDatabaseClause is called when production alterDatabaseClause is exited.
func (s *BaseSQLServerStatementListener) ExitAlterDatabaseClause(ctx *AlterDatabaseClauseContext) {}

// EnterAddSecondaryOption is called when production addSecondaryOption is entered.
func (s *BaseSQLServerStatementListener) EnterAddSecondaryOption(ctx *AddSecondaryOptionContext) {}

// ExitAddSecondaryOption is called when production addSecondaryOption is exited.
func (s *BaseSQLServerStatementListener) ExitAddSecondaryOption(ctx *AddSecondaryOptionContext) {}

// EnterEditionOptions is called when production editionOptions is entered.
func (s *BaseSQLServerStatementListener) EnterEditionOptions(ctx *EditionOptionsContext) {}

// ExitEditionOptions is called when production editionOptions is exited.
func (s *BaseSQLServerStatementListener) ExitEditionOptions(ctx *EditionOptionsContext) {}

// EnterServiceObjective is called when production serviceObjective is entered.
func (s *BaseSQLServerStatementListener) EnterServiceObjective(ctx *ServiceObjectiveContext) {}

// ExitServiceObjective is called when production serviceObjective is exited.
func (s *BaseSQLServerStatementListener) ExitServiceObjective(ctx *ServiceObjectiveContext) {}

// EnterAlterDatabaseOptionSpec is called when production alterDatabaseOptionSpec is entered.
func (s *BaseSQLServerStatementListener) EnterAlterDatabaseOptionSpec(ctx *AlterDatabaseOptionSpecContext) {
}

// ExitAlterDatabaseOptionSpec is called when production alterDatabaseOptionSpec is exited.
func (s *BaseSQLServerStatementListener) ExitAlterDatabaseOptionSpec(ctx *AlterDatabaseOptionSpecContext) {
}

// EnterFileAndFilegroupOptions is called when production fileAndFilegroupOptions is entered.
func (s *BaseSQLServerStatementListener) EnterFileAndFilegroupOptions(ctx *FileAndFilegroupOptionsContext) {
}

// ExitFileAndFilegroupOptions is called when production fileAndFilegroupOptions is exited.
func (s *BaseSQLServerStatementListener) ExitFileAndFilegroupOptions(ctx *FileAndFilegroupOptionsContext) {
}

// EnterAddOrModifyFilegroups is called when production addOrModifyFilegroups is entered.
func (s *BaseSQLServerStatementListener) EnterAddOrModifyFilegroups(ctx *AddOrModifyFilegroupsContext) {
}

// ExitAddOrModifyFilegroups is called when production addOrModifyFilegroups is exited.
func (s *BaseSQLServerStatementListener) ExitAddOrModifyFilegroups(ctx *AddOrModifyFilegroupsContext) {
}

// EnterFilegroupUpdatabilityOption is called when production filegroupUpdatabilityOption is entered.
func (s *BaseSQLServerStatementListener) EnterFilegroupUpdatabilityOption(ctx *FilegroupUpdatabilityOptionContext) {
}

// ExitFilegroupUpdatabilityOption is called when production filegroupUpdatabilityOption is exited.
func (s *BaseSQLServerStatementListener) ExitFilegroupUpdatabilityOption(ctx *FilegroupUpdatabilityOptionContext) {
}

// EnterAddOrModifyFiles is called when production addOrModifyFiles is entered.
func (s *BaseSQLServerStatementListener) EnterAddOrModifyFiles(ctx *AddOrModifyFilesContext) {}

// ExitAddOrModifyFiles is called when production addOrModifyFiles is exited.
func (s *BaseSQLServerStatementListener) ExitAddOrModifyFiles(ctx *AddOrModifyFilesContext) {}

// EnterAcceleratedDatabaseRecovery is called when production acceleratedDatabaseRecovery is entered.
func (s *BaseSQLServerStatementListener) EnterAcceleratedDatabaseRecovery(ctx *AcceleratedDatabaseRecoveryContext) {
}

// ExitAcceleratedDatabaseRecovery is called when production acceleratedDatabaseRecovery is exited.
func (s *BaseSQLServerStatementListener) ExitAcceleratedDatabaseRecovery(ctx *AcceleratedDatabaseRecoveryContext) {
}

// EnterAutoOption is called when production autoOption is entered.
func (s *BaseSQLServerStatementListener) EnterAutoOption(ctx *AutoOptionContext) {}

// ExitAutoOption is called when production autoOption is exited.
func (s *BaseSQLServerStatementListener) ExitAutoOption(ctx *AutoOptionContext) {}

// EnterAutomaticTuningOption is called when production automaticTuningOption is entered.
func (s *BaseSQLServerStatementListener) EnterAutomaticTuningOption(ctx *AutomaticTuningOptionContext) {
}

// ExitAutomaticTuningOption is called when production automaticTuningOption is exited.
func (s *BaseSQLServerStatementListener) ExitAutomaticTuningOption(ctx *AutomaticTuningOptionContext) {
}

// EnterChangeTrackingOption is called when production changeTrackingOption is entered.
func (s *BaseSQLServerStatementListener) EnterChangeTrackingOption(ctx *ChangeTrackingOptionContext) {
}

// ExitChangeTrackingOption is called when production changeTrackingOption is exited.
func (s *BaseSQLServerStatementListener) ExitChangeTrackingOption(ctx *ChangeTrackingOptionContext) {}

// EnterChangeTrackingOptionList is called when production changeTrackingOptionList is entered.
func (s *BaseSQLServerStatementListener) EnterChangeTrackingOptionList(ctx *ChangeTrackingOptionListContext) {
}

// ExitChangeTrackingOptionList is called when production changeTrackingOptionList is exited.
func (s *BaseSQLServerStatementListener) ExitChangeTrackingOptionList(ctx *ChangeTrackingOptionListContext) {
}

// EnterCursorOption is called when production cursorOption is entered.
func (s *BaseSQLServerStatementListener) EnterCursorOption(ctx *CursorOptionContext) {}

// ExitCursorOption is called when production cursorOption is exited.
func (s *BaseSQLServerStatementListener) ExitCursorOption(ctx *CursorOptionContext) {}

// EnterExternalAccessOption is called when production externalAccessOption is entered.
func (s *BaseSQLServerStatementListener) EnterExternalAccessOption(ctx *ExternalAccessOptionContext) {
}

// ExitExternalAccessOption is called when production externalAccessOption is exited.
func (s *BaseSQLServerStatementListener) ExitExternalAccessOption(ctx *ExternalAccessOptionContext) {}

// EnterQueryStoreOptions is called when production queryStoreOptions is entered.
func (s *BaseSQLServerStatementListener) EnterQueryStoreOptions(ctx *QueryStoreOptionsContext) {}

// ExitQueryStoreOptions is called when production queryStoreOptions is exited.
func (s *BaseSQLServerStatementListener) ExitQueryStoreOptions(ctx *QueryStoreOptionsContext) {}

// EnterQueryStoreOptionList is called when production queryStoreOptionList is entered.
func (s *BaseSQLServerStatementListener) EnterQueryStoreOptionList(ctx *QueryStoreOptionListContext) {
}

// ExitQueryStoreOptionList is called when production queryStoreOptionList is exited.
func (s *BaseSQLServerStatementListener) ExitQueryStoreOptionList(ctx *QueryStoreOptionListContext) {}

// EnterQueryCapturePolicyOptionList is called when production queryCapturePolicyOptionList is entered.
func (s *BaseSQLServerStatementListener) EnterQueryCapturePolicyOptionList(ctx *QueryCapturePolicyOptionListContext) {
}

// ExitQueryCapturePolicyOptionList is called when production queryCapturePolicyOptionList is exited.
func (s *BaseSQLServerStatementListener) ExitQueryCapturePolicyOptionList(ctx *QueryCapturePolicyOptionListContext) {
}

// EnterRecoveryOption is called when production recoveryOption is entered.
func (s *BaseSQLServerStatementListener) EnterRecoveryOption(ctx *RecoveryOptionContext) {}

// ExitRecoveryOption is called when production recoveryOption is exited.
func (s *BaseSQLServerStatementListener) ExitRecoveryOption(ctx *RecoveryOptionContext) {}

// EnterSqlOption is called when production sqlOption is entered.
func (s *BaseSQLServerStatementListener) EnterSqlOption(ctx *SqlOptionContext) {}

// ExitSqlOption is called when production sqlOption is exited.
func (s *BaseSQLServerStatementListener) ExitSqlOption(ctx *SqlOptionContext) {}

// EnterSnapshotOption is called when production snapshotOption is entered.
func (s *BaseSQLServerStatementListener) EnterSnapshotOption(ctx *SnapshotOptionContext) {}

// ExitSnapshotOption is called when production snapshotOption is exited.
func (s *BaseSQLServerStatementListener) ExitSnapshotOption(ctx *SnapshotOptionContext) {}

// EnterServiceBrokerOption is called when production serviceBrokerOption is entered.
func (s *BaseSQLServerStatementListener) EnterServiceBrokerOption(ctx *ServiceBrokerOptionContext) {}

// ExitServiceBrokerOption is called when production serviceBrokerOption is exited.
func (s *BaseSQLServerStatementListener) ExitServiceBrokerOption(ctx *ServiceBrokerOptionContext) {}

// EnterTargetRecoveryTimeOption is called when production targetRecoveryTimeOption is entered.
func (s *BaseSQLServerStatementListener) EnterTargetRecoveryTimeOption(ctx *TargetRecoveryTimeOptionContext) {
}

// ExitTargetRecoveryTimeOption is called when production targetRecoveryTimeOption is exited.
func (s *BaseSQLServerStatementListener) ExitTargetRecoveryTimeOption(ctx *TargetRecoveryTimeOptionContext) {
}

// EnterTermination is called when production termination is entered.
func (s *BaseSQLServerStatementListener) EnterTermination(ctx *TerminationContext) {}

// ExitTermination is called when production termination is exited.
func (s *BaseSQLServerStatementListener) ExitTermination(ctx *TerminationContext) {}

// EnterCreateServiceClause is called when production createServiceClause is entered.
func (s *BaseSQLServerStatementListener) EnterCreateServiceClause(ctx *CreateServiceClauseContext) {}

// ExitCreateServiceClause is called when production createServiceClause is exited.
func (s *BaseSQLServerStatementListener) ExitCreateServiceClause(ctx *CreateServiceClauseContext) {}

// EnterAlterServiceClause is called when production alterServiceClause is entered.
func (s *BaseSQLServerStatementListener) EnterAlterServiceClause(ctx *AlterServiceClauseContext) {}

// ExitAlterServiceClause is called when production alterServiceClause is exited.
func (s *BaseSQLServerStatementListener) ExitAlterServiceClause(ctx *AlterServiceClauseContext) {}

// EnterAlterServiceOptArg is called when production alterServiceOptArg is entered.
func (s *BaseSQLServerStatementListener) EnterAlterServiceOptArg(ctx *AlterServiceOptArgContext) {}

// ExitAlterServiceOptArg is called when production alterServiceOptArg is exited.
func (s *BaseSQLServerStatementListener) ExitAlterServiceOptArg(ctx *AlterServiceOptArgContext) {}

// EnterSchemaNameClause is called when production schemaNameClause is entered.
func (s *BaseSQLServerStatementListener) EnterSchemaNameClause(ctx *SchemaNameClauseContext) {}

// ExitSchemaNameClause is called when production schemaNameClause is exited.
func (s *BaseSQLServerStatementListener) ExitSchemaNameClause(ctx *SchemaNameClauseContext) {}

// EnterSchemaElement is called when production schemaElement is entered.
func (s *BaseSQLServerStatementListener) EnterSchemaElement(ctx *SchemaElementContext) {}

// ExitSchemaElement is called when production schemaElement is exited.
func (s *BaseSQLServerStatementListener) ExitSchemaElement(ctx *SchemaElementContext) {}

// EnterCreateTableAsSelectClause is called when production createTableAsSelectClause is entered.
func (s *BaseSQLServerStatementListener) EnterCreateTableAsSelectClause(ctx *CreateTableAsSelectClauseContext) {
}

// ExitCreateTableAsSelectClause is called when production createTableAsSelectClause is exited.
func (s *BaseSQLServerStatementListener) ExitCreateTableAsSelectClause(ctx *CreateTableAsSelectClauseContext) {
}

// EnterCreateTableAsSelect is called when production createTableAsSelect is entered.
func (s *BaseSQLServerStatementListener) EnterCreateTableAsSelect(ctx *CreateTableAsSelectContext) {}

// ExitCreateTableAsSelect is called when production createTableAsSelect is exited.
func (s *BaseSQLServerStatementListener) ExitCreateTableAsSelect(ctx *CreateTableAsSelectContext) {}

// EnterCreateRemoteTableAsSelect is called when production createRemoteTableAsSelect is entered.
func (s *BaseSQLServerStatementListener) EnterCreateRemoteTableAsSelect(ctx *CreateRemoteTableAsSelectContext) {
}

// ExitCreateRemoteTableAsSelect is called when production createRemoteTableAsSelect is exited.
func (s *BaseSQLServerStatementListener) ExitCreateRemoteTableAsSelect(ctx *CreateRemoteTableAsSelectContext) {
}

// EnterWithDistributionOption is called when production withDistributionOption is entered.
func (s *BaseSQLServerStatementListener) EnterWithDistributionOption(ctx *WithDistributionOptionContext) {
}

// ExitWithDistributionOption is called when production withDistributionOption is exited.
func (s *BaseSQLServerStatementListener) ExitWithDistributionOption(ctx *WithDistributionOptionContext) {
}

// EnterOptionQueryHintClause is called when production optionQueryHintClause is entered.
func (s *BaseSQLServerStatementListener) EnterOptionQueryHintClause(ctx *OptionQueryHintClauseContext) {
}

// ExitOptionQueryHintClause is called when production optionQueryHintClause is exited.
func (s *BaseSQLServerStatementListener) ExitOptionQueryHintClause(ctx *OptionQueryHintClauseContext) {
}

// EnterGrant is called when production grant is entered.
func (s *BaseSQLServerStatementListener) EnterGrant(ctx *GrantContext) {}

// ExitGrant is called when production grant is exited.
func (s *BaseSQLServerStatementListener) ExitGrant(ctx *GrantContext) {}

// EnterRevoke is called when production revoke is entered.
func (s *BaseSQLServerStatementListener) EnterRevoke(ctx *RevokeContext) {}

// ExitRevoke is called when production revoke is exited.
func (s *BaseSQLServerStatementListener) ExitRevoke(ctx *RevokeContext) {}

// EnterDeny is called when production deny is entered.
func (s *BaseSQLServerStatementListener) EnterDeny(ctx *DenyContext) {}

// ExitDeny is called when production deny is exited.
func (s *BaseSQLServerStatementListener) ExitDeny(ctx *DenyContext) {}

// EnterClassPrivilegesClause is called when production classPrivilegesClause is entered.
func (s *BaseSQLServerStatementListener) EnterClassPrivilegesClause(ctx *ClassPrivilegesClauseContext) {
}

// ExitClassPrivilegesClause is called when production classPrivilegesClause is exited.
func (s *BaseSQLServerStatementListener) ExitClassPrivilegesClause(ctx *ClassPrivilegesClauseContext) {
}

// EnterClassTypePrivilegesClause is called when production classTypePrivilegesClause is entered.
func (s *BaseSQLServerStatementListener) EnterClassTypePrivilegesClause(ctx *ClassTypePrivilegesClauseContext) {
}

// ExitClassTypePrivilegesClause is called when production classTypePrivilegesClause is exited.
func (s *BaseSQLServerStatementListener) ExitClassTypePrivilegesClause(ctx *ClassTypePrivilegesClauseContext) {
}

// EnterOptionForClause is called when production optionForClause is entered.
func (s *BaseSQLServerStatementListener) EnterOptionForClause(ctx *OptionForClauseContext) {}

// ExitOptionForClause is called when production optionForClause is exited.
func (s *BaseSQLServerStatementListener) ExitOptionForClause(ctx *OptionForClauseContext) {}

// EnterClassPrivileges is called when production classPrivileges is entered.
func (s *BaseSQLServerStatementListener) EnterClassPrivileges(ctx *ClassPrivilegesContext) {}

// ExitClassPrivileges is called when production classPrivileges is exited.
func (s *BaseSQLServerStatementListener) ExitClassPrivileges(ctx *ClassPrivilegesContext) {}

// EnterOnClassClause is called when production onClassClause is entered.
func (s *BaseSQLServerStatementListener) EnterOnClassClause(ctx *OnClassClauseContext) {}

// ExitOnClassClause is called when production onClassClause is exited.
func (s *BaseSQLServerStatementListener) ExitOnClassClause(ctx *OnClassClauseContext) {}

// EnterClassTypePrivileges is called when production classTypePrivileges is entered.
func (s *BaseSQLServerStatementListener) EnterClassTypePrivileges(ctx *ClassTypePrivilegesContext) {}

// ExitClassTypePrivileges is called when production classTypePrivileges is exited.
func (s *BaseSQLServerStatementListener) ExitClassTypePrivileges(ctx *ClassTypePrivilegesContext) {}

// EnterOnClassTypeClause is called when production onClassTypeClause is entered.
func (s *BaseSQLServerStatementListener) EnterOnClassTypeClause(ctx *OnClassTypeClauseContext) {}

// ExitOnClassTypeClause is called when production onClassTypeClause is exited.
func (s *BaseSQLServerStatementListener) ExitOnClassTypeClause(ctx *OnClassTypeClauseContext) {}

// EnterPrivilegeType is called when production privilegeType is entered.
func (s *BaseSQLServerStatementListener) EnterPrivilegeType(ctx *PrivilegeTypeContext) {}

// ExitPrivilegeType is called when production privilegeType is exited.
func (s *BaseSQLServerStatementListener) ExitPrivilegeType(ctx *PrivilegeTypeContext) {}

// EnterBasicPermission is called when production basicPermission is entered.
func (s *BaseSQLServerStatementListener) EnterBasicPermission(ctx *BasicPermissionContext) {}

// ExitBasicPermission is called when production basicPermission is exited.
func (s *BaseSQLServerStatementListener) ExitBasicPermission(ctx *BasicPermissionContext) {}

// EnterObjectPermission is called when production objectPermission is entered.
func (s *BaseSQLServerStatementListener) EnterObjectPermission(ctx *ObjectPermissionContext) {}

// ExitObjectPermission is called when production objectPermission is exited.
func (s *BaseSQLServerStatementListener) ExitObjectPermission(ctx *ObjectPermissionContext) {}

// EnterServerPermission is called when production serverPermission is entered.
func (s *BaseSQLServerStatementListener) EnterServerPermission(ctx *ServerPermissionContext) {}

// ExitServerPermission is called when production serverPermission is exited.
func (s *BaseSQLServerStatementListener) ExitServerPermission(ctx *ServerPermissionContext) {}

// EnterServerPrincipalPermission is called when production serverPrincipalPermission is entered.
func (s *BaseSQLServerStatementListener) EnterServerPrincipalPermission(ctx *ServerPrincipalPermissionContext) {
}

// ExitServerPrincipalPermission is called when production serverPrincipalPermission is exited.
func (s *BaseSQLServerStatementListener) ExitServerPrincipalPermission(ctx *ServerPrincipalPermissionContext) {
}

// EnterDatabasePermission is called when production databasePermission is entered.
func (s *BaseSQLServerStatementListener) EnterDatabasePermission(ctx *DatabasePermissionContext) {}

// ExitDatabasePermission is called when production databasePermission is exited.
func (s *BaseSQLServerStatementListener) ExitDatabasePermission(ctx *DatabasePermissionContext) {}

// EnterDatabasePrincipalPermission is called when production databasePrincipalPermission is entered.
func (s *BaseSQLServerStatementListener) EnterDatabasePrincipalPermission(ctx *DatabasePrincipalPermissionContext) {
}

// ExitDatabasePrincipalPermission is called when production databasePrincipalPermission is exited.
func (s *BaseSQLServerStatementListener) ExitDatabasePrincipalPermission(ctx *DatabasePrincipalPermissionContext) {
}

// EnterSchemaPermission is called when production schemaPermission is entered.
func (s *BaseSQLServerStatementListener) EnterSchemaPermission(ctx *SchemaPermissionContext) {}

// ExitSchemaPermission is called when production schemaPermission is exited.
func (s *BaseSQLServerStatementListener) ExitSchemaPermission(ctx *SchemaPermissionContext) {}

// EnterServiceBrokerPermission is called when production serviceBrokerPermission is entered.
func (s *BaseSQLServerStatementListener) EnterServiceBrokerPermission(ctx *ServiceBrokerPermissionContext) {
}

// ExitServiceBrokerPermission is called when production serviceBrokerPermission is exited.
func (s *BaseSQLServerStatementListener) ExitServiceBrokerPermission(ctx *ServiceBrokerPermissionContext) {
}

// EnterEndpointPermission is called when production endpointPermission is entered.
func (s *BaseSQLServerStatementListener) EnterEndpointPermission(ctx *EndpointPermissionContext) {}

// ExitEndpointPermission is called when production endpointPermission is exited.
func (s *BaseSQLServerStatementListener) ExitEndpointPermission(ctx *EndpointPermissionContext) {}

// EnterCertificatePermission is called when production certificatePermission is entered.
func (s *BaseSQLServerStatementListener) EnterCertificatePermission(ctx *CertificatePermissionContext) {
}

// ExitCertificatePermission is called when production certificatePermission is exited.
func (s *BaseSQLServerStatementListener) ExitCertificatePermission(ctx *CertificatePermissionContext) {
}

// EnterSymmetricKeyPermission is called when production symmetricKeyPermission is entered.
func (s *BaseSQLServerStatementListener) EnterSymmetricKeyPermission(ctx *SymmetricKeyPermissionContext) {
}

// ExitSymmetricKeyPermission is called when production symmetricKeyPermission is exited.
func (s *BaseSQLServerStatementListener) ExitSymmetricKeyPermission(ctx *SymmetricKeyPermissionContext) {
}

// EnterAsymmetricKeyPermission is called when production asymmetricKeyPermission is entered.
func (s *BaseSQLServerStatementListener) EnterAsymmetricKeyPermission(ctx *AsymmetricKeyPermissionContext) {
}

// ExitAsymmetricKeyPermission is called when production asymmetricKeyPermission is exited.
func (s *BaseSQLServerStatementListener) ExitAsymmetricKeyPermission(ctx *AsymmetricKeyPermissionContext) {
}

// EnterAssemblyPermission is called when production assemblyPermission is entered.
func (s *BaseSQLServerStatementListener) EnterAssemblyPermission(ctx *AssemblyPermissionContext) {}

// ExitAssemblyPermission is called when production assemblyPermission is exited.
func (s *BaseSQLServerStatementListener) ExitAssemblyPermission(ctx *AssemblyPermissionContext) {}

// EnterAvailabilityGroupPermission is called when production availabilityGroupPermission is entered.
func (s *BaseSQLServerStatementListener) EnterAvailabilityGroupPermission(ctx *AvailabilityGroupPermissionContext) {
}

// ExitAvailabilityGroupPermission is called when production availabilityGroupPermission is exited.
func (s *BaseSQLServerStatementListener) ExitAvailabilityGroupPermission(ctx *AvailabilityGroupPermissionContext) {
}

// EnterFullTextPermission is called when production fullTextPermission is entered.
func (s *BaseSQLServerStatementListener) EnterFullTextPermission(ctx *FullTextPermissionContext) {}

// ExitFullTextPermission is called when production fullTextPermission is exited.
func (s *BaseSQLServerStatementListener) ExitFullTextPermission(ctx *FullTextPermissionContext) {}

// EnterClass_ is called when production class_ is entered.
func (s *BaseSQLServerStatementListener) EnterClass_(ctx *Class_Context) {}

// ExitClass_ is called when production class_ is exited.
func (s *BaseSQLServerStatementListener) ExitClass_(ctx *Class_Context) {}

// EnterClassType is called when production classType is entered.
func (s *BaseSQLServerStatementListener) EnterClassType(ctx *ClassTypeContext) {}

// ExitClassType is called when production classType is exited.
func (s *BaseSQLServerStatementListener) ExitClassType(ctx *ClassTypeContext) {}

// EnterRoleClause is called when production roleClause is entered.
func (s *BaseSQLServerStatementListener) EnterRoleClause(ctx *RoleClauseContext) {}

// ExitRoleClause is called when production roleClause is exited.
func (s *BaseSQLServerStatementListener) ExitRoleClause(ctx *RoleClauseContext) {}

// EnterCreateUser is called when production createUser is entered.
func (s *BaseSQLServerStatementListener) EnterCreateUser(ctx *CreateUserContext) {}

// ExitCreateUser is called when production createUser is exited.
func (s *BaseSQLServerStatementListener) ExitCreateUser(ctx *CreateUserContext) {}

// EnterDropUser is called when production dropUser is entered.
func (s *BaseSQLServerStatementListener) EnterDropUser(ctx *DropUserContext) {}

// ExitDropUser is called when production dropUser is exited.
func (s *BaseSQLServerStatementListener) ExitDropUser(ctx *DropUserContext) {}

// EnterAlterUser is called when production alterUser is entered.
func (s *BaseSQLServerStatementListener) EnterAlterUser(ctx *AlterUserContext) {}

// ExitAlterUser is called when production alterUser is exited.
func (s *BaseSQLServerStatementListener) ExitAlterUser(ctx *AlterUserContext) {}

// EnterCreateRole is called when production createRole is entered.
func (s *BaseSQLServerStatementListener) EnterCreateRole(ctx *CreateRoleContext) {}

// ExitCreateRole is called when production createRole is exited.
func (s *BaseSQLServerStatementListener) ExitCreateRole(ctx *CreateRoleContext) {}

// EnterDropRole is called when production dropRole is entered.
func (s *BaseSQLServerStatementListener) EnterDropRole(ctx *DropRoleContext) {}

// ExitDropRole is called when production dropRole is exited.
func (s *BaseSQLServerStatementListener) ExitDropRole(ctx *DropRoleContext) {}

// EnterAlterRole is called when production alterRole is entered.
func (s *BaseSQLServerStatementListener) EnterAlterRole(ctx *AlterRoleContext) {}

// ExitAlterRole is called when production alterRole is exited.
func (s *BaseSQLServerStatementListener) ExitAlterRole(ctx *AlterRoleContext) {}

// EnterCreateLogin is called when production createLogin is entered.
func (s *BaseSQLServerStatementListener) EnterCreateLogin(ctx *CreateLoginContext) {}

// ExitCreateLogin is called when production createLogin is exited.
func (s *BaseSQLServerStatementListener) ExitCreateLogin(ctx *CreateLoginContext) {}

// EnterDropLogin is called when production dropLogin is entered.
func (s *BaseSQLServerStatementListener) EnterDropLogin(ctx *DropLoginContext) {}

// ExitDropLogin is called when production dropLogin is exited.
func (s *BaseSQLServerStatementListener) ExitDropLogin(ctx *DropLoginContext) {}

// EnterAlterLogin is called when production alterLogin is entered.
func (s *BaseSQLServerStatementListener) EnterAlterLogin(ctx *AlterLoginContext) {}

// ExitAlterLogin is called when production alterLogin is exited.
func (s *BaseSQLServerStatementListener) ExitAlterLogin(ctx *AlterLoginContext) {}

// EnterSetTransaction is called when production setTransaction is entered.
func (s *BaseSQLServerStatementListener) EnterSetTransaction(ctx *SetTransactionContext) {}

// ExitSetTransaction is called when production setTransaction is exited.
func (s *BaseSQLServerStatementListener) ExitSetTransaction(ctx *SetTransactionContext) {}

// EnterSetImplicitTransactions is called when production setImplicitTransactions is entered.
func (s *BaseSQLServerStatementListener) EnterSetImplicitTransactions(ctx *SetImplicitTransactionsContext) {
}

// ExitSetImplicitTransactions is called when production setImplicitTransactions is exited.
func (s *BaseSQLServerStatementListener) ExitSetImplicitTransactions(ctx *SetImplicitTransactionsContext) {
}

// EnterImplicitTransactionsValue is called when production implicitTransactionsValue is entered.
func (s *BaseSQLServerStatementListener) EnterImplicitTransactionsValue(ctx *ImplicitTransactionsValueContext) {
}

// ExitImplicitTransactionsValue is called when production implicitTransactionsValue is exited.
func (s *BaseSQLServerStatementListener) ExitImplicitTransactionsValue(ctx *ImplicitTransactionsValueContext) {
}

// EnterBeginTransaction is called when production beginTransaction is entered.
func (s *BaseSQLServerStatementListener) EnterBeginTransaction(ctx *BeginTransactionContext) {}

// ExitBeginTransaction is called when production beginTransaction is exited.
func (s *BaseSQLServerStatementListener) ExitBeginTransaction(ctx *BeginTransactionContext) {}

// EnterBeginDistributedTransaction is called when production beginDistributedTransaction is entered.
func (s *BaseSQLServerStatementListener) EnterBeginDistributedTransaction(ctx *BeginDistributedTransactionContext) {
}

// ExitBeginDistributedTransaction is called when production beginDistributedTransaction is exited.
func (s *BaseSQLServerStatementListener) ExitBeginDistributedTransaction(ctx *BeginDistributedTransactionContext) {
}

// EnterCommit is called when production commit is entered.
func (s *BaseSQLServerStatementListener) EnterCommit(ctx *CommitContext) {}

// ExitCommit is called when production commit is exited.
func (s *BaseSQLServerStatementListener) ExitCommit(ctx *CommitContext) {}

// EnterCommitWork is called when production commitWork is entered.
func (s *BaseSQLServerStatementListener) EnterCommitWork(ctx *CommitWorkContext) {}

// ExitCommitWork is called when production commitWork is exited.
func (s *BaseSQLServerStatementListener) ExitCommitWork(ctx *CommitWorkContext) {}

// EnterRollback is called when production rollback is entered.
func (s *BaseSQLServerStatementListener) EnterRollback(ctx *RollbackContext) {}

// ExitRollback is called when production rollback is exited.
func (s *BaseSQLServerStatementListener) ExitRollback(ctx *RollbackContext) {}

// EnterRollbackWork is called when production rollbackWork is entered.
func (s *BaseSQLServerStatementListener) EnterRollbackWork(ctx *RollbackWorkContext) {}

// ExitRollbackWork is called when production rollbackWork is exited.
func (s *BaseSQLServerStatementListener) ExitRollbackWork(ctx *RollbackWorkContext) {}

// EnterSavepoint is called when production savepoint is entered.
func (s *BaseSQLServerStatementListener) EnterSavepoint(ctx *SavepointContext) {}

// ExitSavepoint is called when production savepoint is exited.
func (s *BaseSQLServerStatementListener) ExitSavepoint(ctx *SavepointContext) {}

// EnterCall is called when production call is entered.
func (s *BaseSQLServerStatementListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseSQLServerStatementListener) ExitCall(ctx *CallContext) {}

// EnterExplain is called when production explain is entered.
func (s *BaseSQLServerStatementListener) EnterExplain(ctx *ExplainContext) {}

// ExitExplain is called when production explain is exited.
func (s *BaseSQLServerStatementListener) ExitExplain(ctx *ExplainContext) {}

// EnterExplainableStatement is called when production explainableStatement is entered.
func (s *BaseSQLServerStatementListener) EnterExplainableStatement(ctx *ExplainableStatementContext) {
}

// ExitExplainableStatement is called when production explainableStatement is exited.
func (s *BaseSQLServerStatementListener) ExitExplainableStatement(ctx *ExplainableStatementContext) {}

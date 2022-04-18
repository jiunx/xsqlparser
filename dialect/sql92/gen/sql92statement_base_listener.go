// Code generated from E:/GoProject/src/github.com/jiunx/xsqlparser/dialect/sql92/grammer\SQL92Statement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQL92Statement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSQL92StatementListener is a complete listener for a parse tree produced by SQL92StatementParser.
type BaseSQL92StatementListener struct{}

var _ SQL92StatementListener = &BaseSQL92StatementListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSQL92StatementListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSQL92StatementListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSQL92StatementListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSQL92StatementListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExecute is called when production execute is entered.
func (s *BaseSQL92StatementListener) EnterExecute(ctx *ExecuteContext) {}

// ExitExecute is called when production execute is exited.
func (s *BaseSQL92StatementListener) ExitExecute(ctx *ExecuteContext) {}

// EnterInsert is called when production insert is entered.
func (s *BaseSQL92StatementListener) EnterInsert(ctx *InsertContext) {}

// ExitInsert is called when production insert is exited.
func (s *BaseSQL92StatementListener) ExitInsert(ctx *InsertContext) {}

// EnterInsertValuesClause is called when production insertValuesClause is entered.
func (s *BaseSQL92StatementListener) EnterInsertValuesClause(ctx *InsertValuesClauseContext) {}

// ExitInsertValuesClause is called when production insertValuesClause is exited.
func (s *BaseSQL92StatementListener) ExitInsertValuesClause(ctx *InsertValuesClauseContext) {}

// EnterInsertSelectClause is called when production insertSelectClause is entered.
func (s *BaseSQL92StatementListener) EnterInsertSelectClause(ctx *InsertSelectClauseContext) {}

// ExitInsertSelectClause is called when production insertSelectClause is exited.
func (s *BaseSQL92StatementListener) ExitInsertSelectClause(ctx *InsertSelectClauseContext) {}

// EnterUpdate is called when production update is entered.
func (s *BaseSQL92StatementListener) EnterUpdate(ctx *UpdateContext) {}

// ExitUpdate is called when production update is exited.
func (s *BaseSQL92StatementListener) ExitUpdate(ctx *UpdateContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseSQL92StatementListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseSQL92StatementListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterSetAssignmentsClause is called when production setAssignmentsClause is entered.
func (s *BaseSQL92StatementListener) EnterSetAssignmentsClause(ctx *SetAssignmentsClauseContext) {}

// ExitSetAssignmentsClause is called when production setAssignmentsClause is exited.
func (s *BaseSQL92StatementListener) ExitSetAssignmentsClause(ctx *SetAssignmentsClauseContext) {}

// EnterAssignmentValues is called when production assignmentValues is entered.
func (s *BaseSQL92StatementListener) EnterAssignmentValues(ctx *AssignmentValuesContext) {}

// ExitAssignmentValues is called when production assignmentValues is exited.
func (s *BaseSQL92StatementListener) ExitAssignmentValues(ctx *AssignmentValuesContext) {}

// EnterAssignmentValue is called when production assignmentValue is entered.
func (s *BaseSQL92StatementListener) EnterAssignmentValue(ctx *AssignmentValueContext) {}

// ExitAssignmentValue is called when production assignmentValue is exited.
func (s *BaseSQL92StatementListener) ExitAssignmentValue(ctx *AssignmentValueContext) {}

// EnterBlobValue is called when production blobValue is entered.
func (s *BaseSQL92StatementListener) EnterBlobValue(ctx *BlobValueContext) {}

// ExitBlobValue is called when production blobValue is exited.
func (s *BaseSQL92StatementListener) ExitBlobValue(ctx *BlobValueContext) {}

// EnterDelete is called when production delete is entered.
func (s *BaseSQL92StatementListener) EnterDelete(ctx *DeleteContext) {}

// ExitDelete is called when production delete is exited.
func (s *BaseSQL92StatementListener) ExitDelete(ctx *DeleteContext) {}

// EnterSingleTableClause is called when production singleTableClause is entered.
func (s *BaseSQL92StatementListener) EnterSingleTableClause(ctx *SingleTableClauseContext) {}

// ExitSingleTableClause is called when production singleTableClause is exited.
func (s *BaseSQL92StatementListener) ExitSingleTableClause(ctx *SingleTableClauseContext) {}

// EnterSelect is called when production select is entered.
func (s *BaseSQL92StatementListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production select is exited.
func (s *BaseSQL92StatementListener) ExitSelect(ctx *SelectContext) {}

// EnterUnionClause is called when production unionClause is entered.
func (s *BaseSQL92StatementListener) EnterUnionClause(ctx *UnionClauseContext) {}

// ExitUnionClause is called when production unionClause is exited.
func (s *BaseSQL92StatementListener) ExitUnionClause(ctx *UnionClauseContext) {}

// EnterSelectClause is called when production selectClause is entered.
func (s *BaseSQL92StatementListener) EnterSelectClause(ctx *SelectClauseContext) {}

// ExitSelectClause is called when production selectClause is exited.
func (s *BaseSQL92StatementListener) ExitSelectClause(ctx *SelectClauseContext) {}

// EnterSelectSpecification is called when production selectSpecification is entered.
func (s *BaseSQL92StatementListener) EnterSelectSpecification(ctx *SelectSpecificationContext) {}

// ExitSelectSpecification is called when production selectSpecification is exited.
func (s *BaseSQL92StatementListener) ExitSelectSpecification(ctx *SelectSpecificationContext) {}

// EnterDuplicateSpecification is called when production duplicateSpecification is entered.
func (s *BaseSQL92StatementListener) EnterDuplicateSpecification(ctx *DuplicateSpecificationContext) {
}

// ExitDuplicateSpecification is called when production duplicateSpecification is exited.
func (s *BaseSQL92StatementListener) ExitDuplicateSpecification(ctx *DuplicateSpecificationContext) {}

// EnterProjections is called when production projections is entered.
func (s *BaseSQL92StatementListener) EnterProjections(ctx *ProjectionsContext) {}

// ExitProjections is called when production projections is exited.
func (s *BaseSQL92StatementListener) ExitProjections(ctx *ProjectionsContext) {}

// EnterProjection is called when production projection is entered.
func (s *BaseSQL92StatementListener) EnterProjection(ctx *ProjectionContext) {}

// ExitProjection is called when production projection is exited.
func (s *BaseSQL92StatementListener) ExitProjection(ctx *ProjectionContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseSQL92StatementListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseSQL92StatementListener) ExitAlias(ctx *AliasContext) {}

// EnterUnqualifiedShorthand is called when production unqualifiedShorthand is entered.
func (s *BaseSQL92StatementListener) EnterUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) {}

// ExitUnqualifiedShorthand is called when production unqualifiedShorthand is exited.
func (s *BaseSQL92StatementListener) ExitUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) {}

// EnterQualifiedShorthand is called when production qualifiedShorthand is entered.
func (s *BaseSQL92StatementListener) EnterQualifiedShorthand(ctx *QualifiedShorthandContext) {}

// ExitQualifiedShorthand is called when production qualifiedShorthand is exited.
func (s *BaseSQL92StatementListener) ExitQualifiedShorthand(ctx *QualifiedShorthandContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseSQL92StatementListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseSQL92StatementListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterTableReferences is called when production tableReferences is entered.
func (s *BaseSQL92StatementListener) EnterTableReferences(ctx *TableReferencesContext) {}

// ExitTableReferences is called when production tableReferences is exited.
func (s *BaseSQL92StatementListener) ExitTableReferences(ctx *TableReferencesContext) {}

// EnterEscapedTableReference is called when production escapedTableReference is entered.
func (s *BaseSQL92StatementListener) EnterEscapedTableReference(ctx *EscapedTableReferenceContext) {}

// ExitEscapedTableReference is called when production escapedTableReference is exited.
func (s *BaseSQL92StatementListener) ExitEscapedTableReference(ctx *EscapedTableReferenceContext) {}

// EnterTableReference is called when production tableReference is entered.
func (s *BaseSQL92StatementListener) EnterTableReference(ctx *TableReferenceContext) {}

// ExitTableReference is called when production tableReference is exited.
func (s *BaseSQL92StatementListener) ExitTableReference(ctx *TableReferenceContext) {}

// EnterTableFactor is called when production tableFactor is entered.
func (s *BaseSQL92StatementListener) EnterTableFactor(ctx *TableFactorContext) {}

// ExitTableFactor is called when production tableFactor is exited.
func (s *BaseSQL92StatementListener) ExitTableFactor(ctx *TableFactorContext) {}

// EnterJoinedTable is called when production joinedTable is entered.
func (s *BaseSQL92StatementListener) EnterJoinedTable(ctx *JoinedTableContext) {}

// ExitJoinedTable is called when production joinedTable is exited.
func (s *BaseSQL92StatementListener) ExitJoinedTable(ctx *JoinedTableContext) {}

// EnterJoinSpecification is called when production joinSpecification is entered.
func (s *BaseSQL92StatementListener) EnterJoinSpecification(ctx *JoinSpecificationContext) {}

// ExitJoinSpecification is called when production joinSpecification is exited.
func (s *BaseSQL92StatementListener) ExitJoinSpecification(ctx *JoinSpecificationContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseSQL92StatementListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseSQL92StatementListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseSQL92StatementListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseSQL92StatementListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseSQL92StatementListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseSQL92StatementListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseSQL92StatementListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseSQL92StatementListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterLimitRowCount is called when production limitRowCount is entered.
func (s *BaseSQL92StatementListener) EnterLimitRowCount(ctx *LimitRowCountContext) {}

// ExitLimitRowCount is called when production limitRowCount is exited.
func (s *BaseSQL92StatementListener) ExitLimitRowCount(ctx *LimitRowCountContext) {}

// EnterLimitOffset is called when production limitOffset is entered.
func (s *BaseSQL92StatementListener) EnterLimitOffset(ctx *LimitOffsetContext) {}

// ExitLimitOffset is called when production limitOffset is exited.
func (s *BaseSQL92StatementListener) ExitLimitOffset(ctx *LimitOffsetContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseSQL92StatementListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseSQL92StatementListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterParameterMarker is called when production parameterMarker is entered.
func (s *BaseSQL92StatementListener) EnterParameterMarker(ctx *ParameterMarkerContext) {}

// ExitParameterMarker is called when production parameterMarker is exited.
func (s *BaseSQL92StatementListener) ExitParameterMarker(ctx *ParameterMarkerContext) {}

// EnterLiterals is called when production literals is entered.
func (s *BaseSQL92StatementListener) EnterLiterals(ctx *LiteralsContext) {}

// ExitLiterals is called when production literals is exited.
func (s *BaseSQL92StatementListener) ExitLiterals(ctx *LiteralsContext) {}

// EnterStringLiterals is called when production stringLiterals is entered.
func (s *BaseSQL92StatementListener) EnterStringLiterals(ctx *StringLiteralsContext) {}

// ExitStringLiterals is called when production stringLiterals is exited.
func (s *BaseSQL92StatementListener) ExitStringLiterals(ctx *StringLiteralsContext) {}

// EnterNumberLiterals is called when production numberLiterals is entered.
func (s *BaseSQL92StatementListener) EnterNumberLiterals(ctx *NumberLiteralsContext) {}

// ExitNumberLiterals is called when production numberLiterals is exited.
func (s *BaseSQL92StatementListener) ExitNumberLiterals(ctx *NumberLiteralsContext) {}

// EnterDateTimeLiterals is called when production dateTimeLiterals is entered.
func (s *BaseSQL92StatementListener) EnterDateTimeLiterals(ctx *DateTimeLiteralsContext) {}

// ExitDateTimeLiterals is called when production dateTimeLiterals is exited.
func (s *BaseSQL92StatementListener) ExitDateTimeLiterals(ctx *DateTimeLiteralsContext) {}

// EnterHexadecimalLiterals is called when production hexadecimalLiterals is entered.
func (s *BaseSQL92StatementListener) EnterHexadecimalLiterals(ctx *HexadecimalLiteralsContext) {}

// ExitHexadecimalLiterals is called when production hexadecimalLiterals is exited.
func (s *BaseSQL92StatementListener) ExitHexadecimalLiterals(ctx *HexadecimalLiteralsContext) {}

// EnterBitValueLiterals is called when production bitValueLiterals is entered.
func (s *BaseSQL92StatementListener) EnterBitValueLiterals(ctx *BitValueLiteralsContext) {}

// ExitBitValueLiterals is called when production bitValueLiterals is exited.
func (s *BaseSQL92StatementListener) ExitBitValueLiterals(ctx *BitValueLiteralsContext) {}

// EnterBooleanLiterals is called when production booleanLiterals is entered.
func (s *BaseSQL92StatementListener) EnterBooleanLiterals(ctx *BooleanLiteralsContext) {}

// ExitBooleanLiterals is called when production booleanLiterals is exited.
func (s *BaseSQL92StatementListener) ExitBooleanLiterals(ctx *BooleanLiteralsContext) {}

// EnterNullValueLiterals is called when production nullValueLiterals is entered.
func (s *BaseSQL92StatementListener) EnterNullValueLiterals(ctx *NullValueLiteralsContext) {}

// ExitNullValueLiterals is called when production nullValueLiterals is exited.
func (s *BaseSQL92StatementListener) ExitNullValueLiterals(ctx *NullValueLiteralsContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSQL92StatementListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSQL92StatementListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterUnreservedWord is called when production unreservedWord is entered.
func (s *BaseSQL92StatementListener) EnterUnreservedWord(ctx *UnreservedWordContext) {}

// ExitUnreservedWord is called when production unreservedWord is exited.
func (s *BaseSQL92StatementListener) ExitUnreservedWord(ctx *UnreservedWordContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseSQL92StatementListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseSQL92StatementListener) ExitVariable(ctx *VariableContext) {}

// EnterSchemaName is called when production schemaName is entered.
func (s *BaseSQL92StatementListener) EnterSchemaName(ctx *SchemaNameContext) {}

// ExitSchemaName is called when production schemaName is exited.
func (s *BaseSQL92StatementListener) ExitSchemaName(ctx *SchemaNameContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseSQL92StatementListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseSQL92StatementListener) ExitTableName(ctx *TableNameContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseSQL92StatementListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseSQL92StatementListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterViewName is called when production viewName is entered.
func (s *BaseSQL92StatementListener) EnterViewName(ctx *ViewNameContext) {}

// ExitViewName is called when production viewName is exited.
func (s *BaseSQL92StatementListener) ExitViewName(ctx *ViewNameContext) {}

// EnterOwner is called when production owner is entered.
func (s *BaseSQL92StatementListener) EnterOwner(ctx *OwnerContext) {}

// ExitOwner is called when production owner is exited.
func (s *BaseSQL92StatementListener) ExitOwner(ctx *OwnerContext) {}

// EnterName is called when production name is entered.
func (s *BaseSQL92StatementListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseSQL92StatementListener) ExitName(ctx *NameContext) {}

// EnterConstraintName is called when production constraintName is entered.
func (s *BaseSQL92StatementListener) EnterConstraintName(ctx *ConstraintNameContext) {}

// ExitConstraintName is called when production constraintName is exited.
func (s *BaseSQL92StatementListener) ExitConstraintName(ctx *ConstraintNameContext) {}

// EnterColumnNames is called when production columnNames is entered.
func (s *BaseSQL92StatementListener) EnterColumnNames(ctx *ColumnNamesContext) {}

// ExitColumnNames is called when production columnNames is exited.
func (s *BaseSQL92StatementListener) ExitColumnNames(ctx *ColumnNamesContext) {}

// EnterTableNames is called when production tableNames is entered.
func (s *BaseSQL92StatementListener) EnterTableNames(ctx *TableNamesContext) {}

// ExitTableNames is called when production tableNames is exited.
func (s *BaseSQL92StatementListener) ExitTableNames(ctx *TableNamesContext) {}

// EnterCharacterSetName is called when production characterSetName is entered.
func (s *BaseSQL92StatementListener) EnterCharacterSetName(ctx *CharacterSetNameContext) {}

// ExitCharacterSetName is called when production characterSetName is exited.
func (s *BaseSQL92StatementListener) ExitCharacterSetName(ctx *CharacterSetNameContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSQL92StatementListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSQL92StatementListener) ExitExpr(ctx *ExprContext) {}

// EnterAndOperator is called when production andOperator is entered.
func (s *BaseSQL92StatementListener) EnterAndOperator(ctx *AndOperatorContext) {}

// ExitAndOperator is called when production andOperator is exited.
func (s *BaseSQL92StatementListener) ExitAndOperator(ctx *AndOperatorContext) {}

// EnterOrOperator is called when production orOperator is entered.
func (s *BaseSQL92StatementListener) EnterOrOperator(ctx *OrOperatorContext) {}

// ExitOrOperator is called when production orOperator is exited.
func (s *BaseSQL92StatementListener) ExitOrOperator(ctx *OrOperatorContext) {}

// EnterNotOperator is called when production notOperator is entered.
func (s *BaseSQL92StatementListener) EnterNotOperator(ctx *NotOperatorContext) {}

// ExitNotOperator is called when production notOperator is exited.
func (s *BaseSQL92StatementListener) ExitNotOperator(ctx *NotOperatorContext) {}

// EnterBooleanPrimary is called when production booleanPrimary is entered.
func (s *BaseSQL92StatementListener) EnterBooleanPrimary(ctx *BooleanPrimaryContext) {}

// ExitBooleanPrimary is called when production booleanPrimary is exited.
func (s *BaseSQL92StatementListener) ExitBooleanPrimary(ctx *BooleanPrimaryContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseSQL92StatementListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseSQL92StatementListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseSQL92StatementListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseSQL92StatementListener) ExitPredicate(ctx *PredicateContext) {}

// EnterBitExpr is called when production bitExpr is entered.
func (s *BaseSQL92StatementListener) EnterBitExpr(ctx *BitExprContext) {}

// ExitBitExpr is called when production bitExpr is exited.
func (s *BaseSQL92StatementListener) ExitBitExpr(ctx *BitExprContext) {}

// EnterSimpleExpr is called when production simpleExpr is entered.
func (s *BaseSQL92StatementListener) EnterSimpleExpr(ctx *SimpleExprContext) {}

// ExitSimpleExpr is called when production simpleExpr is exited.
func (s *BaseSQL92StatementListener) ExitSimpleExpr(ctx *SimpleExprContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseSQL92StatementListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseSQL92StatementListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterAggregationFunction is called when production aggregationFunction is entered.
func (s *BaseSQL92StatementListener) EnterAggregationFunction(ctx *AggregationFunctionContext) {}

// ExitAggregationFunction is called when production aggregationFunction is exited.
func (s *BaseSQL92StatementListener) ExitAggregationFunction(ctx *AggregationFunctionContext) {}

// EnterAggregationFunctionName is called when production aggregationFunctionName is entered.
func (s *BaseSQL92StatementListener) EnterAggregationFunctionName(ctx *AggregationFunctionNameContext) {
}

// ExitAggregationFunctionName is called when production aggregationFunctionName is exited.
func (s *BaseSQL92StatementListener) ExitAggregationFunctionName(ctx *AggregationFunctionNameContext) {
}

// EnterDistinct is called when production distinct is entered.
func (s *BaseSQL92StatementListener) EnterDistinct(ctx *DistinctContext) {}

// ExitDistinct is called when production distinct is exited.
func (s *BaseSQL92StatementListener) ExitDistinct(ctx *DistinctContext) {}

// EnterSpecialFunction is called when production specialFunction is entered.
func (s *BaseSQL92StatementListener) EnterSpecialFunction(ctx *SpecialFunctionContext) {}

// ExitSpecialFunction is called when production specialFunction is exited.
func (s *BaseSQL92StatementListener) ExitSpecialFunction(ctx *SpecialFunctionContext) {}

// EnterCastFunction is called when production castFunction is entered.
func (s *BaseSQL92StatementListener) EnterCastFunction(ctx *CastFunctionContext) {}

// ExitCastFunction is called when production castFunction is exited.
func (s *BaseSQL92StatementListener) ExitCastFunction(ctx *CastFunctionContext) {}

// EnterConvertFunction is called when production convertFunction is entered.
func (s *BaseSQL92StatementListener) EnterConvertFunction(ctx *ConvertFunctionContext) {}

// ExitConvertFunction is called when production convertFunction is exited.
func (s *BaseSQL92StatementListener) ExitConvertFunction(ctx *ConvertFunctionContext) {}

// EnterPositionFunction is called when production positionFunction is entered.
func (s *BaseSQL92StatementListener) EnterPositionFunction(ctx *PositionFunctionContext) {}

// ExitPositionFunction is called when production positionFunction is exited.
func (s *BaseSQL92StatementListener) ExitPositionFunction(ctx *PositionFunctionContext) {}

// EnterSubstringFunction is called when production substringFunction is entered.
func (s *BaseSQL92StatementListener) EnterSubstringFunction(ctx *SubstringFunctionContext) {}

// ExitSubstringFunction is called when production substringFunction is exited.
func (s *BaseSQL92StatementListener) ExitSubstringFunction(ctx *SubstringFunctionContext) {}

// EnterExtractFunction is called when production extractFunction is entered.
func (s *BaseSQL92StatementListener) EnterExtractFunction(ctx *ExtractFunctionContext) {}

// ExitExtractFunction is called when production extractFunction is exited.
func (s *BaseSQL92StatementListener) ExitExtractFunction(ctx *ExtractFunctionContext) {}

// EnterTrimFunction is called when production trimFunction is entered.
func (s *BaseSQL92StatementListener) EnterTrimFunction(ctx *TrimFunctionContext) {}

// ExitTrimFunction is called when production trimFunction is exited.
func (s *BaseSQL92StatementListener) ExitTrimFunction(ctx *TrimFunctionContext) {}

// EnterRegularFunction is called when production regularFunction is entered.
func (s *BaseSQL92StatementListener) EnterRegularFunction(ctx *RegularFunctionContext) {}

// ExitRegularFunction is called when production regularFunction is exited.
func (s *BaseSQL92StatementListener) ExitRegularFunction(ctx *RegularFunctionContext) {}

// EnterRegularFunctionName is called when production regularFunctionName is entered.
func (s *BaseSQL92StatementListener) EnterRegularFunctionName(ctx *RegularFunctionNameContext) {}

// ExitRegularFunctionName is called when production regularFunctionName is exited.
func (s *BaseSQL92StatementListener) ExitRegularFunctionName(ctx *RegularFunctionNameContext) {}

// EnterMatchExpression is called when production matchExpression is entered.
func (s *BaseSQL92StatementListener) EnterMatchExpression(ctx *MatchExpressionContext) {}

// ExitMatchExpression is called when production matchExpression is exited.
func (s *BaseSQL92StatementListener) ExitMatchExpression(ctx *MatchExpressionContext) {}

// EnterCaseExpression is called when production caseExpression is entered.
func (s *BaseSQL92StatementListener) EnterCaseExpression(ctx *CaseExpressionContext) {}

// ExitCaseExpression is called when production caseExpression is exited.
func (s *BaseSQL92StatementListener) ExitCaseExpression(ctx *CaseExpressionContext) {}

// EnterCaseWhen is called when production caseWhen is entered.
func (s *BaseSQL92StatementListener) EnterCaseWhen(ctx *CaseWhenContext) {}

// ExitCaseWhen is called when production caseWhen is exited.
func (s *BaseSQL92StatementListener) ExitCaseWhen(ctx *CaseWhenContext) {}

// EnterCaseElse is called when production caseElse is entered.
func (s *BaseSQL92StatementListener) EnterCaseElse(ctx *CaseElseContext) {}

// ExitCaseElse is called when production caseElse is exited.
func (s *BaseSQL92StatementListener) ExitCaseElse(ctx *CaseElseContext) {}

// EnterIntervalExpression is called when production intervalExpression is entered.
func (s *BaseSQL92StatementListener) EnterIntervalExpression(ctx *IntervalExpressionContext) {}

// ExitIntervalExpression is called when production intervalExpression is exited.
func (s *BaseSQL92StatementListener) ExitIntervalExpression(ctx *IntervalExpressionContext) {}

// EnterIntervalUnit is called when production intervalUnit is entered.
func (s *BaseSQL92StatementListener) EnterIntervalUnit(ctx *IntervalUnitContext) {}

// ExitIntervalUnit is called when production intervalUnit is exited.
func (s *BaseSQL92StatementListener) ExitIntervalUnit(ctx *IntervalUnitContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseSQL92StatementListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseSQL92StatementListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderByItem is called when production orderByItem is entered.
func (s *BaseSQL92StatementListener) EnterOrderByItem(ctx *OrderByItemContext) {}

// ExitOrderByItem is called when production orderByItem is exited.
func (s *BaseSQL92StatementListener) ExitOrderByItem(ctx *OrderByItemContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseSQL92StatementListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseSQL92StatementListener) ExitDataType(ctx *DataTypeContext) {}

// EnterDataTypeName is called when production dataTypeName is entered.
func (s *BaseSQL92StatementListener) EnterDataTypeName(ctx *DataTypeNameContext) {}

// ExitDataTypeName is called when production dataTypeName is exited.
func (s *BaseSQL92StatementListener) ExitDataTypeName(ctx *DataTypeNameContext) {}

// EnterDataTypeLength is called when production dataTypeLength is entered.
func (s *BaseSQL92StatementListener) EnterDataTypeLength(ctx *DataTypeLengthContext) {}

// ExitDataTypeLength is called when production dataTypeLength is exited.
func (s *BaseSQL92StatementListener) ExitDataTypeLength(ctx *DataTypeLengthContext) {}

// EnterCharacterSet is called when production characterSet is entered.
func (s *BaseSQL92StatementListener) EnterCharacterSet(ctx *CharacterSetContext) {}

// ExitCharacterSet is called when production characterSet is exited.
func (s *BaseSQL92StatementListener) ExitCharacterSet(ctx *CharacterSetContext) {}

// EnterCollateClause is called when production collateClause is entered.
func (s *BaseSQL92StatementListener) EnterCollateClause(ctx *CollateClauseContext) {}

// ExitCollateClause is called when production collateClause is exited.
func (s *BaseSQL92StatementListener) ExitCollateClause(ctx *CollateClauseContext) {}

// EnterIgnoredIdentifier is called when production ignoredIdentifier is entered.
func (s *BaseSQL92StatementListener) EnterIgnoredIdentifier(ctx *IgnoredIdentifierContext) {}

// ExitIgnoredIdentifier is called when production ignoredIdentifier is exited.
func (s *BaseSQL92StatementListener) ExitIgnoredIdentifier(ctx *IgnoredIdentifierContext) {}

// EnterDropBehaviour is called when production dropBehaviour is entered.
func (s *BaseSQL92StatementListener) EnterDropBehaviour(ctx *DropBehaviourContext) {}

// ExitDropBehaviour is called when production dropBehaviour is exited.
func (s *BaseSQL92StatementListener) ExitDropBehaviour(ctx *DropBehaviourContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseSQL92StatementListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseSQL92StatementListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterAlterTable is called when production alterTable is entered.
func (s *BaseSQL92StatementListener) EnterAlterTable(ctx *AlterTableContext) {}

// ExitAlterTable is called when production alterTable is exited.
func (s *BaseSQL92StatementListener) ExitAlterTable(ctx *AlterTableContext) {}

// EnterDropTable is called when production dropTable is entered.
func (s *BaseSQL92StatementListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production dropTable is exited.
func (s *BaseSQL92StatementListener) ExitDropTable(ctx *DropTableContext) {}

// EnterCreateDatabase is called when production createDatabase is entered.
func (s *BaseSQL92StatementListener) EnterCreateDatabase(ctx *CreateDatabaseContext) {}

// ExitCreateDatabase is called when production createDatabase is exited.
func (s *BaseSQL92StatementListener) ExitCreateDatabase(ctx *CreateDatabaseContext) {}

// EnterDropDatabase is called when production dropDatabase is entered.
func (s *BaseSQL92StatementListener) EnterDropDatabase(ctx *DropDatabaseContext) {}

// ExitDropDatabase is called when production dropDatabase is exited.
func (s *BaseSQL92StatementListener) ExitDropDatabase(ctx *DropDatabaseContext) {}

// EnterCreateView is called when production createView is entered.
func (s *BaseSQL92StatementListener) EnterCreateView(ctx *CreateViewContext) {}

// ExitCreateView is called when production createView is exited.
func (s *BaseSQL92StatementListener) ExitCreateView(ctx *CreateViewContext) {}

// EnterDropView is called when production dropView is entered.
func (s *BaseSQL92StatementListener) EnterDropView(ctx *DropViewContext) {}

// ExitDropView is called when production dropView is exited.
func (s *BaseSQL92StatementListener) ExitDropView(ctx *DropViewContext) {}

// EnterCreateTableSpecification is called when production createTableSpecification is entered.
func (s *BaseSQL92StatementListener) EnterCreateTableSpecification(ctx *CreateTableSpecificationContext) {
}

// ExitCreateTableSpecification is called when production createTableSpecification is exited.
func (s *BaseSQL92StatementListener) ExitCreateTableSpecification(ctx *CreateTableSpecificationContext) {
}

// EnterCreateDefinitionClause is called when production createDefinitionClause is entered.
func (s *BaseSQL92StatementListener) EnterCreateDefinitionClause(ctx *CreateDefinitionClauseContext) {
}

// ExitCreateDefinitionClause is called when production createDefinitionClause is exited.
func (s *BaseSQL92StatementListener) ExitCreateDefinitionClause(ctx *CreateDefinitionClauseContext) {}

// EnterCreateDatabaseSpecification_ is called when production createDatabaseSpecification_ is entered.
func (s *BaseSQL92StatementListener) EnterCreateDatabaseSpecification_(ctx *CreateDatabaseSpecification_Context) {
}

// ExitCreateDatabaseSpecification_ is called when production createDatabaseSpecification_ is exited.
func (s *BaseSQL92StatementListener) ExitCreateDatabaseSpecification_(ctx *CreateDatabaseSpecification_Context) {
}

// EnterCreateDefinition is called when production createDefinition is entered.
func (s *BaseSQL92StatementListener) EnterCreateDefinition(ctx *CreateDefinitionContext) {}

// ExitCreateDefinition is called when production createDefinition is exited.
func (s *BaseSQL92StatementListener) ExitCreateDefinition(ctx *CreateDefinitionContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseSQL92StatementListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseSQL92StatementListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterDataTypeOption is called when production dataTypeOption is entered.
func (s *BaseSQL92StatementListener) EnterDataTypeOption(ctx *DataTypeOptionContext) {}

// ExitDataTypeOption is called when production dataTypeOption is exited.
func (s *BaseSQL92StatementListener) ExitDataTypeOption(ctx *DataTypeOptionContext) {}

// EnterCheckConstraintDefinition is called when production checkConstraintDefinition is entered.
func (s *BaseSQL92StatementListener) EnterCheckConstraintDefinition(ctx *CheckConstraintDefinitionContext) {
}

// ExitCheckConstraintDefinition is called when production checkConstraintDefinition is exited.
func (s *BaseSQL92StatementListener) ExitCheckConstraintDefinition(ctx *CheckConstraintDefinitionContext) {
}

// EnterReferenceDefinition is called when production referenceDefinition is entered.
func (s *BaseSQL92StatementListener) EnterReferenceDefinition(ctx *ReferenceDefinitionContext) {}

// ExitReferenceDefinition is called when production referenceDefinition is exited.
func (s *BaseSQL92StatementListener) ExitReferenceDefinition(ctx *ReferenceDefinitionContext) {}

// EnterReferenceOption is called when production referenceOption is entered.
func (s *BaseSQL92StatementListener) EnterReferenceOption(ctx *ReferenceOptionContext) {}

// ExitReferenceOption is called when production referenceOption is exited.
func (s *BaseSQL92StatementListener) ExitReferenceOption(ctx *ReferenceOptionContext) {}

// EnterKeyParts is called when production keyParts is entered.
func (s *BaseSQL92StatementListener) EnterKeyParts(ctx *KeyPartsContext) {}

// ExitKeyParts is called when production keyParts is exited.
func (s *BaseSQL92StatementListener) ExitKeyParts(ctx *KeyPartsContext) {}

// EnterKeyPart is called when production keyPart is entered.
func (s *BaseSQL92StatementListener) EnterKeyPart(ctx *KeyPartContext) {}

// ExitKeyPart is called when production keyPart is exited.
func (s *BaseSQL92StatementListener) ExitKeyPart(ctx *KeyPartContext) {}

// EnterConstraintDefinition is called when production constraintDefinition is entered.
func (s *BaseSQL92StatementListener) EnterConstraintDefinition(ctx *ConstraintDefinitionContext) {}

// ExitConstraintDefinition is called when production constraintDefinition is exited.
func (s *BaseSQL92StatementListener) ExitConstraintDefinition(ctx *ConstraintDefinitionContext) {}

// EnterPrimaryKeyOption is called when production primaryKeyOption is entered.
func (s *BaseSQL92StatementListener) EnterPrimaryKeyOption(ctx *PrimaryKeyOptionContext) {}

// ExitPrimaryKeyOption is called when production primaryKeyOption is exited.
func (s *BaseSQL92StatementListener) ExitPrimaryKeyOption(ctx *PrimaryKeyOptionContext) {}

// EnterPrimaryKey is called when production primaryKey is entered.
func (s *BaseSQL92StatementListener) EnterPrimaryKey(ctx *PrimaryKeyContext) {}

// ExitPrimaryKey is called when production primaryKey is exited.
func (s *BaseSQL92StatementListener) ExitPrimaryKey(ctx *PrimaryKeyContext) {}

// EnterUniqueOption is called when production uniqueOption is entered.
func (s *BaseSQL92StatementListener) EnterUniqueOption(ctx *UniqueOptionContext) {}

// ExitUniqueOption is called when production uniqueOption is exited.
func (s *BaseSQL92StatementListener) ExitUniqueOption(ctx *UniqueOptionContext) {}

// EnterForeignKeyOption is called when production foreignKeyOption is entered.
func (s *BaseSQL92StatementListener) EnterForeignKeyOption(ctx *ForeignKeyOptionContext) {}

// ExitForeignKeyOption is called when production foreignKeyOption is exited.
func (s *BaseSQL92StatementListener) ExitForeignKeyOption(ctx *ForeignKeyOptionContext) {}

// EnterCreateLikeClause is called when production createLikeClause is entered.
func (s *BaseSQL92StatementListener) EnterCreateLikeClause(ctx *CreateLikeClauseContext) {}

// ExitCreateLikeClause is called when production createLikeClause is exited.
func (s *BaseSQL92StatementListener) ExitCreateLikeClause(ctx *CreateLikeClauseContext) {}

// EnterAlterDefinitionClause is called when production alterDefinitionClause is entered.
func (s *BaseSQL92StatementListener) EnterAlterDefinitionClause(ctx *AlterDefinitionClauseContext) {}

// ExitAlterDefinitionClause is called when production alterDefinitionClause is exited.
func (s *BaseSQL92StatementListener) ExitAlterDefinitionClause(ctx *AlterDefinitionClauseContext) {}

// EnterAddColumnSpecification is called when production addColumnSpecification is entered.
func (s *BaseSQL92StatementListener) EnterAddColumnSpecification(ctx *AddColumnSpecificationContext) {
}

// ExitAddColumnSpecification is called when production addColumnSpecification is exited.
func (s *BaseSQL92StatementListener) ExitAddColumnSpecification(ctx *AddColumnSpecificationContext) {}

// EnterModifyColumnSpecification is called when production modifyColumnSpecification is entered.
func (s *BaseSQL92StatementListener) EnterModifyColumnSpecification(ctx *ModifyColumnSpecificationContext) {
}

// ExitModifyColumnSpecification is called when production modifyColumnSpecification is exited.
func (s *BaseSQL92StatementListener) ExitModifyColumnSpecification(ctx *ModifyColumnSpecificationContext) {
}

// EnterDropColumnSpecification is called when production dropColumnSpecification is entered.
func (s *BaseSQL92StatementListener) EnterDropColumnSpecification(ctx *DropColumnSpecificationContext) {
}

// ExitDropColumnSpecification is called when production dropColumnSpecification is exited.
func (s *BaseSQL92StatementListener) ExitDropColumnSpecification(ctx *DropColumnSpecificationContext) {
}

// EnterAddConstraintSpecification is called when production addConstraintSpecification is entered.
func (s *BaseSQL92StatementListener) EnterAddConstraintSpecification(ctx *AddConstraintSpecificationContext) {
}

// ExitAddConstraintSpecification is called when production addConstraintSpecification is exited.
func (s *BaseSQL92StatementListener) ExitAddConstraintSpecification(ctx *AddConstraintSpecificationContext) {
}

// EnterDropConstraintSpecification is called when production dropConstraintSpecification is entered.
func (s *BaseSQL92StatementListener) EnterDropConstraintSpecification(ctx *DropConstraintSpecificationContext) {
}

// ExitDropConstraintSpecification is called when production dropConstraintSpecification is exited.
func (s *BaseSQL92StatementListener) ExitDropConstraintSpecification(ctx *DropConstraintSpecificationContext) {
}

// EnterSetTransaction is called when production setTransaction is entered.
func (s *BaseSQL92StatementListener) EnterSetTransaction(ctx *SetTransactionContext) {}

// ExitSetTransaction is called when production setTransaction is exited.
func (s *BaseSQL92StatementListener) ExitSetTransaction(ctx *SetTransactionContext) {}

// EnterCommit is called when production commit is entered.
func (s *BaseSQL92StatementListener) EnterCommit(ctx *CommitContext) {}

// ExitCommit is called when production commit is exited.
func (s *BaseSQL92StatementListener) ExitCommit(ctx *CommitContext) {}

// EnterRollback is called when production rollback is entered.
func (s *BaseSQL92StatementListener) EnterRollback(ctx *RollbackContext) {}

// ExitRollback is called when production rollback is exited.
func (s *BaseSQL92StatementListener) ExitRollback(ctx *RollbackContext) {}

// EnterLevelOfIsolation is called when production levelOfIsolation is entered.
func (s *BaseSQL92StatementListener) EnterLevelOfIsolation(ctx *LevelOfIsolationContext) {}

// ExitLevelOfIsolation is called when production levelOfIsolation is exited.
func (s *BaseSQL92StatementListener) ExitLevelOfIsolation(ctx *LevelOfIsolationContext) {}

// EnterGrant is called when production grant is entered.
func (s *BaseSQL92StatementListener) EnterGrant(ctx *GrantContext) {}

// ExitGrant is called when production grant is exited.
func (s *BaseSQL92StatementListener) ExitGrant(ctx *GrantContext) {}

// EnterRevoke is called when production revoke is entered.
func (s *BaseSQL92StatementListener) EnterRevoke(ctx *RevokeContext) {}

// ExitRevoke is called when production revoke is exited.
func (s *BaseSQL92StatementListener) ExitRevoke(ctx *RevokeContext) {}

// EnterPrivilegeClause is called when production privilegeClause is entered.
func (s *BaseSQL92StatementListener) EnterPrivilegeClause(ctx *PrivilegeClauseContext) {}

// ExitPrivilegeClause is called when production privilegeClause is exited.
func (s *BaseSQL92StatementListener) ExitPrivilegeClause(ctx *PrivilegeClauseContext) {}

// EnterPrivileges is called when production privileges is entered.
func (s *BaseSQL92StatementListener) EnterPrivileges(ctx *PrivilegesContext) {}

// ExitPrivileges is called when production privileges is exited.
func (s *BaseSQL92StatementListener) ExitPrivileges(ctx *PrivilegesContext) {}

// EnterPrivilegeType is called when production privilegeType is entered.
func (s *BaseSQL92StatementListener) EnterPrivilegeType(ctx *PrivilegeTypeContext) {}

// ExitPrivilegeType is called when production privilegeType is exited.
func (s *BaseSQL92StatementListener) ExitPrivilegeType(ctx *PrivilegeTypeContext) {}

// EnterGrantee is called when production grantee is entered.
func (s *BaseSQL92StatementListener) EnterGrantee(ctx *GranteeContext) {}

// ExitGrantee is called when production grantee is exited.
func (s *BaseSQL92StatementListener) ExitGrantee(ctx *GranteeContext) {}

// EnterOnObjectClause is called when production onObjectClause is entered.
func (s *BaseSQL92StatementListener) EnterOnObjectClause(ctx *OnObjectClauseContext) {}

// ExitOnObjectClause is called when production onObjectClause is exited.
func (s *BaseSQL92StatementListener) ExitOnObjectClause(ctx *OnObjectClauseContext) {}

// EnterObjectType is called when production objectType is entered.
func (s *BaseSQL92StatementListener) EnterObjectType(ctx *ObjectTypeContext) {}

// ExitObjectType is called when production objectType is exited.
func (s *BaseSQL92StatementListener) ExitObjectType(ctx *ObjectTypeContext) {}

// EnterPrivilegeLevel is called when production privilegeLevel is entered.
func (s *BaseSQL92StatementListener) EnterPrivilegeLevel(ctx *PrivilegeLevelContext) {}

// ExitPrivilegeLevel is called when production privilegeLevel is exited.
func (s *BaseSQL92StatementListener) ExitPrivilegeLevel(ctx *PrivilegeLevelContext) {}

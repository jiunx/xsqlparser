// Code generated from E:/GoProject/src/github.com/2212442929/xsqlparser/dialect/mysql/grammer\MySQLStatement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // MySQLStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMySQLStatementListener is a complete listener for a parse tree produced by MySQLStatementParser.
type BaseMySQLStatementListener struct{}

var _ MySQLStatementListener = &BaseMySQLStatementListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMySQLStatementListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMySQLStatementListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMySQLStatementListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMySQLStatementListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExecute is called when production execute is entered.
func (s *BaseMySQLStatementListener) EnterExecute(ctx *ExecuteContext) {}

// ExitExecute is called when production execute is exited.
func (s *BaseMySQLStatementListener) ExitExecute(ctx *ExecuteContext) {}

// EnterInsert is called when production insert is entered.
func (s *BaseMySQLStatementListener) EnterInsert(ctx *InsertContext) {}

// ExitInsert is called when production insert is exited.
func (s *BaseMySQLStatementListener) ExitInsert(ctx *InsertContext) {}

// EnterInsertSpecification is called when production insertSpecification is entered.
func (s *BaseMySQLStatementListener) EnterInsertSpecification(ctx *InsertSpecificationContext) {}

// ExitInsertSpecification is called when production insertSpecification is exited.
func (s *BaseMySQLStatementListener) ExitInsertSpecification(ctx *InsertSpecificationContext) {}

// EnterInsertValuesClause is called when production insertValuesClause is entered.
func (s *BaseMySQLStatementListener) EnterInsertValuesClause(ctx *InsertValuesClauseContext) {}

// ExitInsertValuesClause is called when production insertValuesClause is exited.
func (s *BaseMySQLStatementListener) ExitInsertValuesClause(ctx *InsertValuesClauseContext) {}

// EnterFields is called when production fields is entered.
func (s *BaseMySQLStatementListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseMySQLStatementListener) ExitFields(ctx *FieldsContext) {}

// EnterInsertIdentifier is called when production insertIdentifier is entered.
func (s *BaseMySQLStatementListener) EnterInsertIdentifier(ctx *InsertIdentifierContext) {}

// ExitInsertIdentifier is called when production insertIdentifier is exited.
func (s *BaseMySQLStatementListener) ExitInsertIdentifier(ctx *InsertIdentifierContext) {}

// EnterTableWild is called when production tableWild is entered.
func (s *BaseMySQLStatementListener) EnterTableWild(ctx *TableWildContext) {}

// ExitTableWild is called when production tableWild is exited.
func (s *BaseMySQLStatementListener) ExitTableWild(ctx *TableWildContext) {}

// EnterInsertSelectClause is called when production insertSelectClause is entered.
func (s *BaseMySQLStatementListener) EnterInsertSelectClause(ctx *InsertSelectClauseContext) {}

// ExitInsertSelectClause is called when production insertSelectClause is exited.
func (s *BaseMySQLStatementListener) ExitInsertSelectClause(ctx *InsertSelectClauseContext) {}

// EnterOnDuplicateKeyClause is called when production onDuplicateKeyClause is entered.
func (s *BaseMySQLStatementListener) EnterOnDuplicateKeyClause(ctx *OnDuplicateKeyClauseContext) {}

// ExitOnDuplicateKeyClause is called when production onDuplicateKeyClause is exited.
func (s *BaseMySQLStatementListener) ExitOnDuplicateKeyClause(ctx *OnDuplicateKeyClauseContext) {}

// EnterValueReference is called when production valueReference is entered.
func (s *BaseMySQLStatementListener) EnterValueReference(ctx *ValueReferenceContext) {}

// ExitValueReference is called when production valueReference is exited.
func (s *BaseMySQLStatementListener) ExitValueReference(ctx *ValueReferenceContext) {}

// EnterDerivedColumns is called when production derivedColumns is entered.
func (s *BaseMySQLStatementListener) EnterDerivedColumns(ctx *DerivedColumnsContext) {}

// ExitDerivedColumns is called when production derivedColumns is exited.
func (s *BaseMySQLStatementListener) ExitDerivedColumns(ctx *DerivedColumnsContext) {}

// EnterReplace is called when production replace is entered.
func (s *BaseMySQLStatementListener) EnterReplace(ctx *ReplaceContext) {}

// ExitReplace is called when production replace is exited.
func (s *BaseMySQLStatementListener) ExitReplace(ctx *ReplaceContext) {}

// EnterReplaceSpecification is called when production replaceSpecification is entered.
func (s *BaseMySQLStatementListener) EnterReplaceSpecification(ctx *ReplaceSpecificationContext) {}

// ExitReplaceSpecification is called when production replaceSpecification is exited.
func (s *BaseMySQLStatementListener) ExitReplaceSpecification(ctx *ReplaceSpecificationContext) {}

// EnterReplaceValuesClause is called when production replaceValuesClause is entered.
func (s *BaseMySQLStatementListener) EnterReplaceValuesClause(ctx *ReplaceValuesClauseContext) {}

// ExitReplaceValuesClause is called when production replaceValuesClause is exited.
func (s *BaseMySQLStatementListener) ExitReplaceValuesClause(ctx *ReplaceValuesClauseContext) {}

// EnterReplaceSelectClause is called when production replaceSelectClause is entered.
func (s *BaseMySQLStatementListener) EnterReplaceSelectClause(ctx *ReplaceSelectClauseContext) {}

// ExitReplaceSelectClause is called when production replaceSelectClause is exited.
func (s *BaseMySQLStatementListener) ExitReplaceSelectClause(ctx *ReplaceSelectClauseContext) {}

// EnterUpdate is called when production update is entered.
func (s *BaseMySQLStatementListener) EnterUpdate(ctx *UpdateContext) {}

// ExitUpdate is called when production update is exited.
func (s *BaseMySQLStatementListener) ExitUpdate(ctx *UpdateContext) {}

// EnterUpdateSpecification_ is called when production updateSpecification_ is entered.
func (s *BaseMySQLStatementListener) EnterUpdateSpecification_(ctx *UpdateSpecification_Context) {}

// ExitUpdateSpecification_ is called when production updateSpecification_ is exited.
func (s *BaseMySQLStatementListener) ExitUpdateSpecification_(ctx *UpdateSpecification_Context) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseMySQLStatementListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseMySQLStatementListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterSetAssignmentsClause is called when production setAssignmentsClause is entered.
func (s *BaseMySQLStatementListener) EnterSetAssignmentsClause(ctx *SetAssignmentsClauseContext) {}

// ExitSetAssignmentsClause is called when production setAssignmentsClause is exited.
func (s *BaseMySQLStatementListener) ExitSetAssignmentsClause(ctx *SetAssignmentsClauseContext) {}

// EnterAssignmentValues is called when production assignmentValues is entered.
func (s *BaseMySQLStatementListener) EnterAssignmentValues(ctx *AssignmentValuesContext) {}

// ExitAssignmentValues is called when production assignmentValues is exited.
func (s *BaseMySQLStatementListener) ExitAssignmentValues(ctx *AssignmentValuesContext) {}

// EnterAssignmentValue is called when production assignmentValue is entered.
func (s *BaseMySQLStatementListener) EnterAssignmentValue(ctx *AssignmentValueContext) {}

// ExitAssignmentValue is called when production assignmentValue is exited.
func (s *BaseMySQLStatementListener) ExitAssignmentValue(ctx *AssignmentValueContext) {}

// EnterBlobValue is called when production blobValue is entered.
func (s *BaseMySQLStatementListener) EnterBlobValue(ctx *BlobValueContext) {}

// ExitBlobValue is called when production blobValue is exited.
func (s *BaseMySQLStatementListener) ExitBlobValue(ctx *BlobValueContext) {}

// EnterDelete is called when production delete is entered.
func (s *BaseMySQLStatementListener) EnterDelete(ctx *DeleteContext) {}

// ExitDelete is called when production delete is exited.
func (s *BaseMySQLStatementListener) ExitDelete(ctx *DeleteContext) {}

// EnterDeleteSpecification is called when production deleteSpecification is entered.
func (s *BaseMySQLStatementListener) EnterDeleteSpecification(ctx *DeleteSpecificationContext) {}

// ExitDeleteSpecification is called when production deleteSpecification is exited.
func (s *BaseMySQLStatementListener) ExitDeleteSpecification(ctx *DeleteSpecificationContext) {}

// EnterSingleTableClause is called when production singleTableClause is entered.
func (s *BaseMySQLStatementListener) EnterSingleTableClause(ctx *SingleTableClauseContext) {}

// ExitSingleTableClause is called when production singleTableClause is exited.
func (s *BaseMySQLStatementListener) ExitSingleTableClause(ctx *SingleTableClauseContext) {}

// EnterMultipleTablesClause is called when production multipleTablesClause is entered.
func (s *BaseMySQLStatementListener) EnterMultipleTablesClause(ctx *MultipleTablesClauseContext) {}

// ExitMultipleTablesClause is called when production multipleTablesClause is exited.
func (s *BaseMySQLStatementListener) ExitMultipleTablesClause(ctx *MultipleTablesClauseContext) {}

// EnterSelect is called when production select is entered.
func (s *BaseMySQLStatementListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production select is exited.
func (s *BaseMySQLStatementListener) ExitSelect(ctx *SelectContext) {}

// EnterSelectWithInto is called when production selectWithInto is entered.
func (s *BaseMySQLStatementListener) EnterSelectWithInto(ctx *SelectWithIntoContext) {}

// ExitSelectWithInto is called when production selectWithInto is exited.
func (s *BaseMySQLStatementListener) ExitSelectWithInto(ctx *SelectWithIntoContext) {}

// EnterQueryExpression is called when production queryExpression is entered.
func (s *BaseMySQLStatementListener) EnterQueryExpression(ctx *QueryExpressionContext) {}

// ExitQueryExpression is called when production queryExpression is exited.
func (s *BaseMySQLStatementListener) ExitQueryExpression(ctx *QueryExpressionContext) {}

// EnterQueryExpressionBody is called when production queryExpressionBody is entered.
func (s *BaseMySQLStatementListener) EnterQueryExpressionBody(ctx *QueryExpressionBodyContext) {}

// ExitQueryExpressionBody is called when production queryExpressionBody is exited.
func (s *BaseMySQLStatementListener) ExitQueryExpressionBody(ctx *QueryExpressionBodyContext) {}

// EnterUnionClause is called when production unionClause is entered.
func (s *BaseMySQLStatementListener) EnterUnionClause(ctx *UnionClauseContext) {}

// ExitUnionClause is called when production unionClause is exited.
func (s *BaseMySQLStatementListener) ExitUnionClause(ctx *UnionClauseContext) {}

// EnterQueryExpressionParens is called when production queryExpressionParens is entered.
func (s *BaseMySQLStatementListener) EnterQueryExpressionParens(ctx *QueryExpressionParensContext) {}

// ExitQueryExpressionParens is called when production queryExpressionParens is exited.
func (s *BaseMySQLStatementListener) ExitQueryExpressionParens(ctx *QueryExpressionParensContext) {}

// EnterQueryPrimary is called when production queryPrimary is entered.
func (s *BaseMySQLStatementListener) EnterQueryPrimary(ctx *QueryPrimaryContext) {}

// ExitQueryPrimary is called when production queryPrimary is exited.
func (s *BaseMySQLStatementListener) ExitQueryPrimary(ctx *QueryPrimaryContext) {}

// EnterQuerySpecification is called when production querySpecification is entered.
func (s *BaseMySQLStatementListener) EnterQuerySpecification(ctx *QuerySpecificationContext) {}

// ExitQuerySpecification is called when production querySpecification is exited.
func (s *BaseMySQLStatementListener) ExitQuerySpecification(ctx *QuerySpecificationContext) {}

// EnterCall is called when production call is entered.
func (s *BaseMySQLStatementListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseMySQLStatementListener) ExitCall(ctx *CallContext) {}

// EnterDoStatement is called when production doStatement is entered.
func (s *BaseMySQLStatementListener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production doStatement is exited.
func (s *BaseMySQLStatementListener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterHandlerStatement is called when production handlerStatement is entered.
func (s *BaseMySQLStatementListener) EnterHandlerStatement(ctx *HandlerStatementContext) {}

// ExitHandlerStatement is called when production handlerStatement is exited.
func (s *BaseMySQLStatementListener) ExitHandlerStatement(ctx *HandlerStatementContext) {}

// EnterHandlerOpenStatement is called when production handlerOpenStatement is entered.
func (s *BaseMySQLStatementListener) EnterHandlerOpenStatement(ctx *HandlerOpenStatementContext) {}

// ExitHandlerOpenStatement is called when production handlerOpenStatement is exited.
func (s *BaseMySQLStatementListener) ExitHandlerOpenStatement(ctx *HandlerOpenStatementContext) {}

// EnterHandlerReadIndexStatement is called when production handlerReadIndexStatement is entered.
func (s *BaseMySQLStatementListener) EnterHandlerReadIndexStatement(ctx *HandlerReadIndexStatementContext) {
}

// ExitHandlerReadIndexStatement is called when production handlerReadIndexStatement is exited.
func (s *BaseMySQLStatementListener) ExitHandlerReadIndexStatement(ctx *HandlerReadIndexStatementContext) {
}

// EnterHandlerReadStatement is called when production handlerReadStatement is entered.
func (s *BaseMySQLStatementListener) EnterHandlerReadStatement(ctx *HandlerReadStatementContext) {}

// ExitHandlerReadStatement is called when production handlerReadStatement is exited.
func (s *BaseMySQLStatementListener) ExitHandlerReadStatement(ctx *HandlerReadStatementContext) {}

// EnterHandlerCloseStatement is called when production handlerCloseStatement is entered.
func (s *BaseMySQLStatementListener) EnterHandlerCloseStatement(ctx *HandlerCloseStatementContext) {}

// ExitHandlerCloseStatement is called when production handlerCloseStatement is exited.
func (s *BaseMySQLStatementListener) ExitHandlerCloseStatement(ctx *HandlerCloseStatementContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseMySQLStatementListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseMySQLStatementListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterLoadStatement is called when production loadStatement is entered.
func (s *BaseMySQLStatementListener) EnterLoadStatement(ctx *LoadStatementContext) {}

// ExitLoadStatement is called when production loadStatement is exited.
func (s *BaseMySQLStatementListener) ExitLoadStatement(ctx *LoadStatementContext) {}

// EnterLoadDataStatement is called when production loadDataStatement is entered.
func (s *BaseMySQLStatementListener) EnterLoadDataStatement(ctx *LoadDataStatementContext) {}

// ExitLoadDataStatement is called when production loadDataStatement is exited.
func (s *BaseMySQLStatementListener) ExitLoadDataStatement(ctx *LoadDataStatementContext) {}

// EnterLoadXmlStatement is called when production loadXmlStatement is entered.
func (s *BaseMySQLStatementListener) EnterLoadXmlStatement(ctx *LoadXmlStatementContext) {}

// ExitLoadXmlStatement is called when production loadXmlStatement is exited.
func (s *BaseMySQLStatementListener) ExitLoadXmlStatement(ctx *LoadXmlStatementContext) {}

// EnterExplicitTable is called when production explicitTable is entered.
func (s *BaseMySQLStatementListener) EnterExplicitTable(ctx *ExplicitTableContext) {}

// ExitExplicitTable is called when production explicitTable is exited.
func (s *BaseMySQLStatementListener) ExitExplicitTable(ctx *ExplicitTableContext) {}

// EnterTableValueConstructor is called when production tableValueConstructor is entered.
func (s *BaseMySQLStatementListener) EnterTableValueConstructor(ctx *TableValueConstructorContext) {}

// ExitTableValueConstructor is called when production tableValueConstructor is exited.
func (s *BaseMySQLStatementListener) ExitTableValueConstructor(ctx *TableValueConstructorContext) {}

// EnterRowConstructorList is called when production rowConstructorList is entered.
func (s *BaseMySQLStatementListener) EnterRowConstructorList(ctx *RowConstructorListContext) {}

// ExitRowConstructorList is called when production rowConstructorList is exited.
func (s *BaseMySQLStatementListener) ExitRowConstructorList(ctx *RowConstructorListContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseMySQLStatementListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseMySQLStatementListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterCteClause is called when production cteClause is entered.
func (s *BaseMySQLStatementListener) EnterCteClause(ctx *CteClauseContext) {}

// ExitCteClause is called when production cteClause is exited.
func (s *BaseMySQLStatementListener) ExitCteClause(ctx *CteClauseContext) {}

// EnterSelectSpecification is called when production selectSpecification is entered.
func (s *BaseMySQLStatementListener) EnterSelectSpecification(ctx *SelectSpecificationContext) {}

// ExitSelectSpecification is called when production selectSpecification is exited.
func (s *BaseMySQLStatementListener) ExitSelectSpecification(ctx *SelectSpecificationContext) {}

// EnterDuplicateSpecification is called when production duplicateSpecification is entered.
func (s *BaseMySQLStatementListener) EnterDuplicateSpecification(ctx *DuplicateSpecificationContext) {
}

// ExitDuplicateSpecification is called when production duplicateSpecification is exited.
func (s *BaseMySQLStatementListener) ExitDuplicateSpecification(ctx *DuplicateSpecificationContext) {}

// EnterProjections is called when production projections is entered.
func (s *BaseMySQLStatementListener) EnterProjections(ctx *ProjectionsContext) {}

// ExitProjections is called when production projections is exited.
func (s *BaseMySQLStatementListener) ExitProjections(ctx *ProjectionsContext) {}

// EnterProjection is called when production projection is entered.
func (s *BaseMySQLStatementListener) EnterProjection(ctx *ProjectionContext) {}

// ExitProjection is called when production projection is exited.
func (s *BaseMySQLStatementListener) ExitProjection(ctx *ProjectionContext) {}

// EnterUnqualifiedShorthand is called when production unqualifiedShorthand is entered.
func (s *BaseMySQLStatementListener) EnterUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) {}

// ExitUnqualifiedShorthand is called when production unqualifiedShorthand is exited.
func (s *BaseMySQLStatementListener) ExitUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) {}

// EnterQualifiedShorthand is called when production qualifiedShorthand is entered.
func (s *BaseMySQLStatementListener) EnterQualifiedShorthand(ctx *QualifiedShorthandContext) {}

// ExitQualifiedShorthand is called when production qualifiedShorthand is exited.
func (s *BaseMySQLStatementListener) ExitQualifiedShorthand(ctx *QualifiedShorthandContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseMySQLStatementListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseMySQLStatementListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterTableReferences is called when production tableReferences is entered.
func (s *BaseMySQLStatementListener) EnterTableReferences(ctx *TableReferencesContext) {}

// ExitTableReferences is called when production tableReferences is exited.
func (s *BaseMySQLStatementListener) ExitTableReferences(ctx *TableReferencesContext) {}

// EnterEscapedTableReference is called when production escapedTableReference is entered.
func (s *BaseMySQLStatementListener) EnterEscapedTableReference(ctx *EscapedTableReferenceContext) {}

// ExitEscapedTableReference is called when production escapedTableReference is exited.
func (s *BaseMySQLStatementListener) ExitEscapedTableReference(ctx *EscapedTableReferenceContext) {}

// EnterTableReference is called when production tableReference is entered.
func (s *BaseMySQLStatementListener) EnterTableReference(ctx *TableReferenceContext) {}

// ExitTableReference is called when production tableReference is exited.
func (s *BaseMySQLStatementListener) ExitTableReference(ctx *TableReferenceContext) {}

// EnterTableFactor is called when production tableFactor is entered.
func (s *BaseMySQLStatementListener) EnterTableFactor(ctx *TableFactorContext) {}

// ExitTableFactor is called when production tableFactor is exited.
func (s *BaseMySQLStatementListener) ExitTableFactor(ctx *TableFactorContext) {}

// EnterPartitionNames is called when production partitionNames is entered.
func (s *BaseMySQLStatementListener) EnterPartitionNames(ctx *PartitionNamesContext) {}

// ExitPartitionNames is called when production partitionNames is exited.
func (s *BaseMySQLStatementListener) ExitPartitionNames(ctx *PartitionNamesContext) {}

// EnterIndexHintList is called when production indexHintList is entered.
func (s *BaseMySQLStatementListener) EnterIndexHintList(ctx *IndexHintListContext) {}

// ExitIndexHintList is called when production indexHintList is exited.
func (s *BaseMySQLStatementListener) ExitIndexHintList(ctx *IndexHintListContext) {}

// EnterIndexHint is called when production indexHint is entered.
func (s *BaseMySQLStatementListener) EnterIndexHint(ctx *IndexHintContext) {}

// ExitIndexHint is called when production indexHint is exited.
func (s *BaseMySQLStatementListener) ExitIndexHint(ctx *IndexHintContext) {}

// EnterJoinedTable is called when production joinedTable is entered.
func (s *BaseMySQLStatementListener) EnterJoinedTable(ctx *JoinedTableContext) {}

// ExitJoinedTable is called when production joinedTable is exited.
func (s *BaseMySQLStatementListener) ExitJoinedTable(ctx *JoinedTableContext) {}

// EnterInnerJoinType is called when production innerJoinType is entered.
func (s *BaseMySQLStatementListener) EnterInnerJoinType(ctx *InnerJoinTypeContext) {}

// ExitInnerJoinType is called when production innerJoinType is exited.
func (s *BaseMySQLStatementListener) ExitInnerJoinType(ctx *InnerJoinTypeContext) {}

// EnterOuterJoinType is called when production outerJoinType is entered.
func (s *BaseMySQLStatementListener) EnterOuterJoinType(ctx *OuterJoinTypeContext) {}

// ExitOuterJoinType is called when production outerJoinType is exited.
func (s *BaseMySQLStatementListener) ExitOuterJoinType(ctx *OuterJoinTypeContext) {}

// EnterNaturalJoinType is called when production naturalJoinType is entered.
func (s *BaseMySQLStatementListener) EnterNaturalJoinType(ctx *NaturalJoinTypeContext) {}

// ExitNaturalJoinType is called when production naturalJoinType is exited.
func (s *BaseMySQLStatementListener) ExitNaturalJoinType(ctx *NaturalJoinTypeContext) {}

// EnterJoinSpecification is called when production joinSpecification is entered.
func (s *BaseMySQLStatementListener) EnterJoinSpecification(ctx *JoinSpecificationContext) {}

// ExitJoinSpecification is called when production joinSpecification is exited.
func (s *BaseMySQLStatementListener) ExitJoinSpecification(ctx *JoinSpecificationContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseMySQLStatementListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseMySQLStatementListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseMySQLStatementListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseMySQLStatementListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseMySQLStatementListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseMySQLStatementListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseMySQLStatementListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseMySQLStatementListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterLimitRowCount is called when production limitRowCount is entered.
func (s *BaseMySQLStatementListener) EnterLimitRowCount(ctx *LimitRowCountContext) {}

// ExitLimitRowCount is called when production limitRowCount is exited.
func (s *BaseMySQLStatementListener) ExitLimitRowCount(ctx *LimitRowCountContext) {}

// EnterLimitOffset is called when production limitOffset is entered.
func (s *BaseMySQLStatementListener) EnterLimitOffset(ctx *LimitOffsetContext) {}

// ExitLimitOffset is called when production limitOffset is exited.
func (s *BaseMySQLStatementListener) ExitLimitOffset(ctx *LimitOffsetContext) {}

// EnterWindowClause is called when production windowClause is entered.
func (s *BaseMySQLStatementListener) EnterWindowClause(ctx *WindowClauseContext) {}

// ExitWindowClause is called when production windowClause is exited.
func (s *BaseMySQLStatementListener) ExitWindowClause(ctx *WindowClauseContext) {}

// EnterWindowItem is called when production windowItem is entered.
func (s *BaseMySQLStatementListener) EnterWindowItem(ctx *WindowItemContext) {}

// ExitWindowItem is called when production windowItem is exited.
func (s *BaseMySQLStatementListener) ExitWindowItem(ctx *WindowItemContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseMySQLStatementListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseMySQLStatementListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterSelectLinesInto is called when production selectLinesInto is entered.
func (s *BaseMySQLStatementListener) EnterSelectLinesInto(ctx *SelectLinesIntoContext) {}

// ExitSelectLinesInto is called when production selectLinesInto is exited.
func (s *BaseMySQLStatementListener) ExitSelectLinesInto(ctx *SelectLinesIntoContext) {}

// EnterSelectFieldsInto is called when production selectFieldsInto is entered.
func (s *BaseMySQLStatementListener) EnterSelectFieldsInto(ctx *SelectFieldsIntoContext) {}

// ExitSelectFieldsInto is called when production selectFieldsInto is exited.
func (s *BaseMySQLStatementListener) ExitSelectFieldsInto(ctx *SelectFieldsIntoContext) {}

// EnterSelectIntoExpression is called when production selectIntoExpression is entered.
func (s *BaseMySQLStatementListener) EnterSelectIntoExpression(ctx *SelectIntoExpressionContext) {}

// ExitSelectIntoExpression is called when production selectIntoExpression is exited.
func (s *BaseMySQLStatementListener) ExitSelectIntoExpression(ctx *SelectIntoExpressionContext) {}

// EnterLockClause is called when production lockClause is entered.
func (s *BaseMySQLStatementListener) EnterLockClause(ctx *LockClauseContext) {}

// ExitLockClause is called when production lockClause is exited.
func (s *BaseMySQLStatementListener) ExitLockClause(ctx *LockClauseContext) {}

// EnterLockClauseList is called when production lockClauseList is entered.
func (s *BaseMySQLStatementListener) EnterLockClauseList(ctx *LockClauseListContext) {}

// ExitLockClauseList is called when production lockClauseList is exited.
func (s *BaseMySQLStatementListener) ExitLockClauseList(ctx *LockClauseListContext) {}

// EnterLockStrength is called when production lockStrength is entered.
func (s *BaseMySQLStatementListener) EnterLockStrength(ctx *LockStrengthContext) {}

// ExitLockStrength is called when production lockStrength is exited.
func (s *BaseMySQLStatementListener) ExitLockStrength(ctx *LockStrengthContext) {}

// EnterLockedRowAction is called when production lockedRowAction is entered.
func (s *BaseMySQLStatementListener) EnterLockedRowAction(ctx *LockedRowActionContext) {}

// ExitLockedRowAction is called when production lockedRowAction is exited.
func (s *BaseMySQLStatementListener) ExitLockedRowAction(ctx *LockedRowActionContext) {}

// EnterTableLockingList is called when production tableLockingList is entered.
func (s *BaseMySQLStatementListener) EnterTableLockingList(ctx *TableLockingListContext) {}

// ExitTableLockingList is called when production tableLockingList is exited.
func (s *BaseMySQLStatementListener) ExitTableLockingList(ctx *TableLockingListContext) {}

// EnterTableIdentOptWild is called when production tableIdentOptWild is entered.
func (s *BaseMySQLStatementListener) EnterTableIdentOptWild(ctx *TableIdentOptWildContext) {}

// ExitTableIdentOptWild is called when production tableIdentOptWild is exited.
func (s *BaseMySQLStatementListener) ExitTableIdentOptWild(ctx *TableIdentOptWildContext) {}

// EnterTableAliasRefList is called when production tableAliasRefList is entered.
func (s *BaseMySQLStatementListener) EnterTableAliasRefList(ctx *TableAliasRefListContext) {}

// ExitTableAliasRefList is called when production tableAliasRefList is exited.
func (s *BaseMySQLStatementListener) ExitTableAliasRefList(ctx *TableAliasRefListContext) {}

// EnterParameterMarker is called when production parameterMarker is entered.
func (s *BaseMySQLStatementListener) EnterParameterMarker(ctx *ParameterMarkerContext) {}

// ExitParameterMarker is called when production parameterMarker is exited.
func (s *BaseMySQLStatementListener) ExitParameterMarker(ctx *ParameterMarkerContext) {}

// EnterCustomKeyword is called when production customKeyword is entered.
func (s *BaseMySQLStatementListener) EnterCustomKeyword(ctx *CustomKeywordContext) {}

// ExitCustomKeyword is called when production customKeyword is exited.
func (s *BaseMySQLStatementListener) ExitCustomKeyword(ctx *CustomKeywordContext) {}

// EnterLiterals is called when production literals is entered.
func (s *BaseMySQLStatementListener) EnterLiterals(ctx *LiteralsContext) {}

// ExitLiterals is called when production literals is exited.
func (s *BaseMySQLStatementListener) ExitLiterals(ctx *LiteralsContext) {}

// EnterString_ is called when production string_ is entered.
func (s *BaseMySQLStatementListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BaseMySQLStatementListener) ExitString_(ctx *String_Context) {}

// EnterStringLiterals is called when production stringLiterals is entered.
func (s *BaseMySQLStatementListener) EnterStringLiterals(ctx *StringLiteralsContext) {}

// ExitStringLiterals is called when production stringLiterals is exited.
func (s *BaseMySQLStatementListener) ExitStringLiterals(ctx *StringLiteralsContext) {}

// EnterNumberLiterals is called when production numberLiterals is entered.
func (s *BaseMySQLStatementListener) EnterNumberLiterals(ctx *NumberLiteralsContext) {}

// ExitNumberLiterals is called when production numberLiterals is exited.
func (s *BaseMySQLStatementListener) ExitNumberLiterals(ctx *NumberLiteralsContext) {}

// EnterTemporalLiterals is called when production temporalLiterals is entered.
func (s *BaseMySQLStatementListener) EnterTemporalLiterals(ctx *TemporalLiteralsContext) {}

// ExitTemporalLiterals is called when production temporalLiterals is exited.
func (s *BaseMySQLStatementListener) ExitTemporalLiterals(ctx *TemporalLiteralsContext) {}

// EnterHexadecimalLiterals is called when production hexadecimalLiterals is entered.
func (s *BaseMySQLStatementListener) EnterHexadecimalLiterals(ctx *HexadecimalLiteralsContext) {}

// ExitHexadecimalLiterals is called when production hexadecimalLiterals is exited.
func (s *BaseMySQLStatementListener) ExitHexadecimalLiterals(ctx *HexadecimalLiteralsContext) {}

// EnterBitValueLiterals is called when production bitValueLiterals is entered.
func (s *BaseMySQLStatementListener) EnterBitValueLiterals(ctx *BitValueLiteralsContext) {}

// ExitBitValueLiterals is called when production bitValueLiterals is exited.
func (s *BaseMySQLStatementListener) ExitBitValueLiterals(ctx *BitValueLiteralsContext) {}

// EnterBooleanLiterals is called when production booleanLiterals is entered.
func (s *BaseMySQLStatementListener) EnterBooleanLiterals(ctx *BooleanLiteralsContext) {}

// ExitBooleanLiterals is called when production booleanLiterals is exited.
func (s *BaseMySQLStatementListener) ExitBooleanLiterals(ctx *BooleanLiteralsContext) {}

// EnterNullValueLiterals is called when production nullValueLiterals is entered.
func (s *BaseMySQLStatementListener) EnterNullValueLiterals(ctx *NullValueLiteralsContext) {}

// ExitNullValueLiterals is called when production nullValueLiterals is exited.
func (s *BaseMySQLStatementListener) ExitNullValueLiterals(ctx *NullValueLiteralsContext) {}

// EnterCollationName is called when production collationName is entered.
func (s *BaseMySQLStatementListener) EnterCollationName(ctx *CollationNameContext) {}

// ExitCollationName is called when production collationName is exited.
func (s *BaseMySQLStatementListener) ExitCollationName(ctx *CollationNameContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseMySQLStatementListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseMySQLStatementListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifierKeywordsUnambiguous is called when production identifierKeywordsUnambiguous is entered.
func (s *BaseMySQLStatementListener) EnterIdentifierKeywordsUnambiguous(ctx *IdentifierKeywordsUnambiguousContext) {
}

// ExitIdentifierKeywordsUnambiguous is called when production identifierKeywordsUnambiguous is exited.
func (s *BaseMySQLStatementListener) ExitIdentifierKeywordsUnambiguous(ctx *IdentifierKeywordsUnambiguousContext) {
}

// EnterIdentifierKeywordsAmbiguous1RolesAndLabels is called when production identifierKeywordsAmbiguous1RolesAndLabels is entered.
func (s *BaseMySQLStatementListener) EnterIdentifierKeywordsAmbiguous1RolesAndLabels(ctx *IdentifierKeywordsAmbiguous1RolesAndLabelsContext) {
}

// ExitIdentifierKeywordsAmbiguous1RolesAndLabels is called when production identifierKeywordsAmbiguous1RolesAndLabels is exited.
func (s *BaseMySQLStatementListener) ExitIdentifierKeywordsAmbiguous1RolesAndLabels(ctx *IdentifierKeywordsAmbiguous1RolesAndLabelsContext) {
}

// EnterIdentifierKeywordsAmbiguous2Labels is called when production identifierKeywordsAmbiguous2Labels is entered.
func (s *BaseMySQLStatementListener) EnterIdentifierKeywordsAmbiguous2Labels(ctx *IdentifierKeywordsAmbiguous2LabelsContext) {
}

// ExitIdentifierKeywordsAmbiguous2Labels is called when production identifierKeywordsAmbiguous2Labels is exited.
func (s *BaseMySQLStatementListener) ExitIdentifierKeywordsAmbiguous2Labels(ctx *IdentifierKeywordsAmbiguous2LabelsContext) {
}

// EnterIdentifierKeywordsAmbiguous3Roles is called when production identifierKeywordsAmbiguous3Roles is entered.
func (s *BaseMySQLStatementListener) EnterIdentifierKeywordsAmbiguous3Roles(ctx *IdentifierKeywordsAmbiguous3RolesContext) {
}

// ExitIdentifierKeywordsAmbiguous3Roles is called when production identifierKeywordsAmbiguous3Roles is exited.
func (s *BaseMySQLStatementListener) ExitIdentifierKeywordsAmbiguous3Roles(ctx *IdentifierKeywordsAmbiguous3RolesContext) {
}

// EnterIdentifierKeywordsAmbiguous4SystemVariables is called when production identifierKeywordsAmbiguous4SystemVariables is entered.
func (s *BaseMySQLStatementListener) EnterIdentifierKeywordsAmbiguous4SystemVariables(ctx *IdentifierKeywordsAmbiguous4SystemVariablesContext) {
}

// ExitIdentifierKeywordsAmbiguous4SystemVariables is called when production identifierKeywordsAmbiguous4SystemVariables is exited.
func (s *BaseMySQLStatementListener) ExitIdentifierKeywordsAmbiguous4SystemVariables(ctx *IdentifierKeywordsAmbiguous4SystemVariablesContext) {
}

// EnterTextOrIdentifier is called when production textOrIdentifier is entered.
func (s *BaseMySQLStatementListener) EnterTextOrIdentifier(ctx *TextOrIdentifierContext) {}

// ExitTextOrIdentifier is called when production textOrIdentifier is exited.
func (s *BaseMySQLStatementListener) ExitTextOrIdentifier(ctx *TextOrIdentifierContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseMySQLStatementListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseMySQLStatementListener) ExitVariable(ctx *VariableContext) {}

// EnterUserVariable is called when production userVariable is entered.
func (s *BaseMySQLStatementListener) EnterUserVariable(ctx *UserVariableContext) {}

// ExitUserVariable is called when production userVariable is exited.
func (s *BaseMySQLStatementListener) ExitUserVariable(ctx *UserVariableContext) {}

// EnterSystemVariable is called when production systemVariable is entered.
func (s *BaseMySQLStatementListener) EnterSystemVariable(ctx *SystemVariableContext) {}

// ExitSystemVariable is called when production systemVariable is exited.
func (s *BaseMySQLStatementListener) ExitSystemVariable(ctx *SystemVariableContext) {}

// EnterSetSystemVariable is called when production setSystemVariable is entered.
func (s *BaseMySQLStatementListener) EnterSetSystemVariable(ctx *SetSystemVariableContext) {}

// ExitSetSystemVariable is called when production setSystemVariable is exited.
func (s *BaseMySQLStatementListener) ExitSetSystemVariable(ctx *SetSystemVariableContext) {}

// EnterOptionType is called when production optionType is entered.
func (s *BaseMySQLStatementListener) EnterOptionType(ctx *OptionTypeContext) {}

// ExitOptionType is called when production optionType is exited.
func (s *BaseMySQLStatementListener) ExitOptionType(ctx *OptionTypeContext) {}

// EnterInternalVariableName is called when production internalVariableName is entered.
func (s *BaseMySQLStatementListener) EnterInternalVariableName(ctx *InternalVariableNameContext) {}

// ExitInternalVariableName is called when production internalVariableName is exited.
func (s *BaseMySQLStatementListener) ExitInternalVariableName(ctx *InternalVariableNameContext) {}

// EnterSetExprOrDefault is called when production setExprOrDefault is entered.
func (s *BaseMySQLStatementListener) EnterSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// ExitSetExprOrDefault is called when production setExprOrDefault is exited.
func (s *BaseMySQLStatementListener) ExitSetExprOrDefault(ctx *SetExprOrDefaultContext) {}

// EnterTransactionCharacteristics is called when production transactionCharacteristics is entered.
func (s *BaseMySQLStatementListener) EnterTransactionCharacteristics(ctx *TransactionCharacteristicsContext) {
}

// ExitTransactionCharacteristics is called when production transactionCharacteristics is exited.
func (s *BaseMySQLStatementListener) ExitTransactionCharacteristics(ctx *TransactionCharacteristicsContext) {
}

// EnterIsolationLevel is called when production isolationLevel is entered.
func (s *BaseMySQLStatementListener) EnterIsolationLevel(ctx *IsolationLevelContext) {}

// ExitIsolationLevel is called when production isolationLevel is exited.
func (s *BaseMySQLStatementListener) ExitIsolationLevel(ctx *IsolationLevelContext) {}

// EnterIsolationTypes is called when production isolationTypes is entered.
func (s *BaseMySQLStatementListener) EnterIsolationTypes(ctx *IsolationTypesContext) {}

// ExitIsolationTypes is called when production isolationTypes is exited.
func (s *BaseMySQLStatementListener) ExitIsolationTypes(ctx *IsolationTypesContext) {}

// EnterTransactionAccessMode is called when production transactionAccessMode is entered.
func (s *BaseMySQLStatementListener) EnterTransactionAccessMode(ctx *TransactionAccessModeContext) {}

// ExitTransactionAccessMode is called when production transactionAccessMode is exited.
func (s *BaseMySQLStatementListener) ExitTransactionAccessMode(ctx *TransactionAccessModeContext) {}

// EnterSchemaName is called when production schemaName is entered.
func (s *BaseMySQLStatementListener) EnterSchemaName(ctx *SchemaNameContext) {}

// ExitSchemaName is called when production schemaName is exited.
func (s *BaseMySQLStatementListener) ExitSchemaName(ctx *SchemaNameContext) {}

// EnterSchemaNames is called when production schemaNames is entered.
func (s *BaseMySQLStatementListener) EnterSchemaNames(ctx *SchemaNamesContext) {}

// ExitSchemaNames is called when production schemaNames is exited.
func (s *BaseMySQLStatementListener) ExitSchemaNames(ctx *SchemaNamesContext) {}

// EnterCharsetName is called when production charsetName is entered.
func (s *BaseMySQLStatementListener) EnterCharsetName(ctx *CharsetNameContext) {}

// ExitCharsetName is called when production charsetName is exited.
func (s *BaseMySQLStatementListener) ExitCharsetName(ctx *CharsetNameContext) {}

// EnterSchemaPairs is called when production schemaPairs is entered.
func (s *BaseMySQLStatementListener) EnterSchemaPairs(ctx *SchemaPairsContext) {}

// ExitSchemaPairs is called when production schemaPairs is exited.
func (s *BaseMySQLStatementListener) ExitSchemaPairs(ctx *SchemaPairsContext) {}

// EnterSchemaPair is called when production schemaPair is entered.
func (s *BaseMySQLStatementListener) EnterSchemaPair(ctx *SchemaPairContext) {}

// ExitSchemaPair is called when production schemaPair is exited.
func (s *BaseMySQLStatementListener) ExitSchemaPair(ctx *SchemaPairContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseMySQLStatementListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseMySQLStatementListener) ExitTableName(ctx *TableNameContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseMySQLStatementListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseMySQLStatementListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseMySQLStatementListener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseMySQLStatementListener) ExitIndexName(ctx *IndexNameContext) {}

// EnterConstraintName is called when production constraintName is entered.
func (s *BaseMySQLStatementListener) EnterConstraintName(ctx *ConstraintNameContext) {}

// ExitConstraintName is called when production constraintName is exited.
func (s *BaseMySQLStatementListener) ExitConstraintName(ctx *ConstraintNameContext) {}

// EnterUserIdentifierOrText is called when production userIdentifierOrText is entered.
func (s *BaseMySQLStatementListener) EnterUserIdentifierOrText(ctx *UserIdentifierOrTextContext) {}

// ExitUserIdentifierOrText is called when production userIdentifierOrText is exited.
func (s *BaseMySQLStatementListener) ExitUserIdentifierOrText(ctx *UserIdentifierOrTextContext) {}

// EnterUserName is called when production userName is entered.
func (s *BaseMySQLStatementListener) EnterUserName(ctx *UserNameContext) {}

// ExitUserName is called when production userName is exited.
func (s *BaseMySQLStatementListener) ExitUserName(ctx *UserNameContext) {}

// EnterEventName is called when production eventName is entered.
func (s *BaseMySQLStatementListener) EnterEventName(ctx *EventNameContext) {}

// ExitEventName is called when production eventName is exited.
func (s *BaseMySQLStatementListener) ExitEventName(ctx *EventNameContext) {}

// EnterServerName is called when production serverName is entered.
func (s *BaseMySQLStatementListener) EnterServerName(ctx *ServerNameContext) {}

// ExitServerName is called when production serverName is exited.
func (s *BaseMySQLStatementListener) ExitServerName(ctx *ServerNameContext) {}

// EnterWrapperName is called when production wrapperName is entered.
func (s *BaseMySQLStatementListener) EnterWrapperName(ctx *WrapperNameContext) {}

// ExitWrapperName is called when production wrapperName is exited.
func (s *BaseMySQLStatementListener) ExitWrapperName(ctx *WrapperNameContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseMySQLStatementListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseMySQLStatementListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterViewName is called when production viewName is entered.
func (s *BaseMySQLStatementListener) EnterViewName(ctx *ViewNameContext) {}

// ExitViewName is called when production viewName is exited.
func (s *BaseMySQLStatementListener) ExitViewName(ctx *ViewNameContext) {}

// EnterOwner is called when production owner is entered.
func (s *BaseMySQLStatementListener) EnterOwner(ctx *OwnerContext) {}

// ExitOwner is called when production owner is exited.
func (s *BaseMySQLStatementListener) ExitOwner(ctx *OwnerContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseMySQLStatementListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseMySQLStatementListener) ExitAlias(ctx *AliasContext) {}

// EnterName is called when production name is entered.
func (s *BaseMySQLStatementListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseMySQLStatementListener) ExitName(ctx *NameContext) {}

// EnterTableList is called when production tableList is entered.
func (s *BaseMySQLStatementListener) EnterTableList(ctx *TableListContext) {}

// ExitTableList is called when production tableList is exited.
func (s *BaseMySQLStatementListener) ExitTableList(ctx *TableListContext) {}

// EnterViewNames is called when production viewNames is entered.
func (s *BaseMySQLStatementListener) EnterViewNames(ctx *ViewNamesContext) {}

// ExitViewNames is called when production viewNames is exited.
func (s *BaseMySQLStatementListener) ExitViewNames(ctx *ViewNamesContext) {}

// EnterColumnNames is called when production columnNames is entered.
func (s *BaseMySQLStatementListener) EnterColumnNames(ctx *ColumnNamesContext) {}

// ExitColumnNames is called when production columnNames is exited.
func (s *BaseMySQLStatementListener) ExitColumnNames(ctx *ColumnNamesContext) {}

// EnterGroupName is called when production groupName is entered.
func (s *BaseMySQLStatementListener) EnterGroupName(ctx *GroupNameContext) {}

// ExitGroupName is called when production groupName is exited.
func (s *BaseMySQLStatementListener) ExitGroupName(ctx *GroupNameContext) {}

// EnterRoutineName is called when production routineName is entered.
func (s *BaseMySQLStatementListener) EnterRoutineName(ctx *RoutineNameContext) {}

// ExitRoutineName is called when production routineName is exited.
func (s *BaseMySQLStatementListener) ExitRoutineName(ctx *RoutineNameContext) {}

// EnterShardLibraryName is called when production shardLibraryName is entered.
func (s *BaseMySQLStatementListener) EnterShardLibraryName(ctx *ShardLibraryNameContext) {}

// ExitShardLibraryName is called when production shardLibraryName is exited.
func (s *BaseMySQLStatementListener) ExitShardLibraryName(ctx *ShardLibraryNameContext) {}

// EnterComponentName is called when production componentName is entered.
func (s *BaseMySQLStatementListener) EnterComponentName(ctx *ComponentNameContext) {}

// ExitComponentName is called when production componentName is exited.
func (s *BaseMySQLStatementListener) ExitComponentName(ctx *ComponentNameContext) {}

// EnterPluginName is called when production pluginName is entered.
func (s *BaseMySQLStatementListener) EnterPluginName(ctx *PluginNameContext) {}

// ExitPluginName is called when production pluginName is exited.
func (s *BaseMySQLStatementListener) ExitPluginName(ctx *PluginNameContext) {}

// EnterHostName is called when production hostName is entered.
func (s *BaseMySQLStatementListener) EnterHostName(ctx *HostNameContext) {}

// ExitHostName is called when production hostName is exited.
func (s *BaseMySQLStatementListener) ExitHostName(ctx *HostNameContext) {}

// EnterPort is called when production port is entered.
func (s *BaseMySQLStatementListener) EnterPort(ctx *PortContext) {}

// ExitPort is called when production port is exited.
func (s *BaseMySQLStatementListener) ExitPort(ctx *PortContext) {}

// EnterCloneInstance is called when production cloneInstance is entered.
func (s *BaseMySQLStatementListener) EnterCloneInstance(ctx *CloneInstanceContext) {}

// ExitCloneInstance is called when production cloneInstance is exited.
func (s *BaseMySQLStatementListener) ExitCloneInstance(ctx *CloneInstanceContext) {}

// EnterCloneDir is called when production cloneDir is entered.
func (s *BaseMySQLStatementListener) EnterCloneDir(ctx *CloneDirContext) {}

// ExitCloneDir is called when production cloneDir is exited.
func (s *BaseMySQLStatementListener) ExitCloneDir(ctx *CloneDirContext) {}

// EnterChannelName is called when production channelName is entered.
func (s *BaseMySQLStatementListener) EnterChannelName(ctx *ChannelNameContext) {}

// ExitChannelName is called when production channelName is exited.
func (s *BaseMySQLStatementListener) ExitChannelName(ctx *ChannelNameContext) {}

// EnterLogName is called when production logName is entered.
func (s *BaseMySQLStatementListener) EnterLogName(ctx *LogNameContext) {}

// ExitLogName is called when production logName is exited.
func (s *BaseMySQLStatementListener) ExitLogName(ctx *LogNameContext) {}

// EnterRoleName is called when production roleName is entered.
func (s *BaseMySQLStatementListener) EnterRoleName(ctx *RoleNameContext) {}

// ExitRoleName is called when production roleName is exited.
func (s *BaseMySQLStatementListener) ExitRoleName(ctx *RoleNameContext) {}

// EnterRoleIdentifierOrText is called when production roleIdentifierOrText is entered.
func (s *BaseMySQLStatementListener) EnterRoleIdentifierOrText(ctx *RoleIdentifierOrTextContext) {}

// ExitRoleIdentifierOrText is called when production roleIdentifierOrText is exited.
func (s *BaseMySQLStatementListener) ExitRoleIdentifierOrText(ctx *RoleIdentifierOrTextContext) {}

// EnterEngineRef is called when production engineRef is entered.
func (s *BaseMySQLStatementListener) EnterEngineRef(ctx *EngineRefContext) {}

// ExitEngineRef is called when production engineRef is exited.
func (s *BaseMySQLStatementListener) ExitEngineRef(ctx *EngineRefContext) {}

// EnterTriggerName is called when production triggerName is entered.
func (s *BaseMySQLStatementListener) EnterTriggerName(ctx *TriggerNameContext) {}

// ExitTriggerName is called when production triggerName is exited.
func (s *BaseMySQLStatementListener) ExitTriggerName(ctx *TriggerNameContext) {}

// EnterTriggerTime is called when production triggerTime is entered.
func (s *BaseMySQLStatementListener) EnterTriggerTime(ctx *TriggerTimeContext) {}

// ExitTriggerTime is called when production triggerTime is exited.
func (s *BaseMySQLStatementListener) ExitTriggerTime(ctx *TriggerTimeContext) {}

// EnterTableOrTables is called when production tableOrTables is entered.
func (s *BaseMySQLStatementListener) EnterTableOrTables(ctx *TableOrTablesContext) {}

// ExitTableOrTables is called when production tableOrTables is exited.
func (s *BaseMySQLStatementListener) ExitTableOrTables(ctx *TableOrTablesContext) {}

// EnterUserOrRole is called when production userOrRole is entered.
func (s *BaseMySQLStatementListener) EnterUserOrRole(ctx *UserOrRoleContext) {}

// ExitUserOrRole is called when production userOrRole is exited.
func (s *BaseMySQLStatementListener) ExitUserOrRole(ctx *UserOrRoleContext) {}

// EnterPartitionName is called when production partitionName is entered.
func (s *BaseMySQLStatementListener) EnterPartitionName(ctx *PartitionNameContext) {}

// ExitPartitionName is called when production partitionName is exited.
func (s *BaseMySQLStatementListener) ExitPartitionName(ctx *PartitionNameContext) {}

// EnterIdentifierList is called when production identifierList is entered.
func (s *BaseMySQLStatementListener) EnterIdentifierList(ctx *IdentifierListContext) {}

// ExitIdentifierList is called when production identifierList is exited.
func (s *BaseMySQLStatementListener) ExitIdentifierList(ctx *IdentifierListContext) {}

// EnterAllOrPartitionNameList is called when production allOrPartitionNameList is entered.
func (s *BaseMySQLStatementListener) EnterAllOrPartitionNameList(ctx *AllOrPartitionNameListContext) {
}

// ExitAllOrPartitionNameList is called when production allOrPartitionNameList is exited.
func (s *BaseMySQLStatementListener) ExitAllOrPartitionNameList(ctx *AllOrPartitionNameListContext) {}

// EnterTriggerEvent is called when production triggerEvent is entered.
func (s *BaseMySQLStatementListener) EnterTriggerEvent(ctx *TriggerEventContext) {}

// ExitTriggerEvent is called when production triggerEvent is exited.
func (s *BaseMySQLStatementListener) ExitTriggerEvent(ctx *TriggerEventContext) {}

// EnterTriggerOrder is called when production triggerOrder is entered.
func (s *BaseMySQLStatementListener) EnterTriggerOrder(ctx *TriggerOrderContext) {}

// ExitTriggerOrder is called when production triggerOrder is exited.
func (s *BaseMySQLStatementListener) ExitTriggerOrder(ctx *TriggerOrderContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseMySQLStatementListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseMySQLStatementListener) ExitExpr(ctx *ExprContext) {}

// EnterAndOperator is called when production andOperator is entered.
func (s *BaseMySQLStatementListener) EnterAndOperator(ctx *AndOperatorContext) {}

// ExitAndOperator is called when production andOperator is exited.
func (s *BaseMySQLStatementListener) ExitAndOperator(ctx *AndOperatorContext) {}

// EnterOrOperator is called when production orOperator is entered.
func (s *BaseMySQLStatementListener) EnterOrOperator(ctx *OrOperatorContext) {}

// ExitOrOperator is called when production orOperator is exited.
func (s *BaseMySQLStatementListener) ExitOrOperator(ctx *OrOperatorContext) {}

// EnterNotOperator is called when production notOperator is entered.
func (s *BaseMySQLStatementListener) EnterNotOperator(ctx *NotOperatorContext) {}

// ExitNotOperator is called when production notOperator is exited.
func (s *BaseMySQLStatementListener) ExitNotOperator(ctx *NotOperatorContext) {}

// EnterBooleanPrimary is called when production booleanPrimary is entered.
func (s *BaseMySQLStatementListener) EnterBooleanPrimary(ctx *BooleanPrimaryContext) {}

// ExitBooleanPrimary is called when production booleanPrimary is exited.
func (s *BaseMySQLStatementListener) ExitBooleanPrimary(ctx *BooleanPrimaryContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseMySQLStatementListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseMySQLStatementListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseMySQLStatementListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseMySQLStatementListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseMySQLStatementListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseMySQLStatementListener) ExitPredicate(ctx *PredicateContext) {}

// EnterBitExpr is called when production bitExpr is entered.
func (s *BaseMySQLStatementListener) EnterBitExpr(ctx *BitExprContext) {}

// ExitBitExpr is called when production bitExpr is exited.
func (s *BaseMySQLStatementListener) ExitBitExpr(ctx *BitExprContext) {}

// EnterSimpleExpr is called when production simpleExpr is entered.
func (s *BaseMySQLStatementListener) EnterSimpleExpr(ctx *SimpleExprContext) {}

// ExitSimpleExpr is called when production simpleExpr is exited.
func (s *BaseMySQLStatementListener) ExitSimpleExpr(ctx *SimpleExprContext) {}

// EnterColumnRef is called when production columnRef is entered.
func (s *BaseMySQLStatementListener) EnterColumnRef(ctx *ColumnRefContext) {}

// ExitColumnRef is called when production columnRef is exited.
func (s *BaseMySQLStatementListener) ExitColumnRef(ctx *ColumnRefContext) {}

// EnterColumnRefList is called when production columnRefList is entered.
func (s *BaseMySQLStatementListener) EnterColumnRefList(ctx *ColumnRefListContext) {}

// ExitColumnRefList is called when production columnRefList is exited.
func (s *BaseMySQLStatementListener) ExitColumnRefList(ctx *ColumnRefListContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseMySQLStatementListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseMySQLStatementListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterAggregationFunction is called when production aggregationFunction is entered.
func (s *BaseMySQLStatementListener) EnterAggregationFunction(ctx *AggregationFunctionContext) {}

// ExitAggregationFunction is called when production aggregationFunction is exited.
func (s *BaseMySQLStatementListener) ExitAggregationFunction(ctx *AggregationFunctionContext) {}

// EnterAggregationFunctionName is called when production aggregationFunctionName is entered.
func (s *BaseMySQLStatementListener) EnterAggregationFunctionName(ctx *AggregationFunctionNameContext) {
}

// ExitAggregationFunctionName is called when production aggregationFunctionName is exited.
func (s *BaseMySQLStatementListener) ExitAggregationFunctionName(ctx *AggregationFunctionNameContext) {
}

// EnterDistinct is called when production distinct is entered.
func (s *BaseMySQLStatementListener) EnterDistinct(ctx *DistinctContext) {}

// ExitDistinct is called when production distinct is exited.
func (s *BaseMySQLStatementListener) ExitDistinct(ctx *DistinctContext) {}

// EnterOverClause is called when production overClause is entered.
func (s *BaseMySQLStatementListener) EnterOverClause(ctx *OverClauseContext) {}

// ExitOverClause is called when production overClause is exited.
func (s *BaseMySQLStatementListener) ExitOverClause(ctx *OverClauseContext) {}

// EnterWindowSpecification is called when production windowSpecification is entered.
func (s *BaseMySQLStatementListener) EnterWindowSpecification(ctx *WindowSpecificationContext) {}

// ExitWindowSpecification is called when production windowSpecification is exited.
func (s *BaseMySQLStatementListener) ExitWindowSpecification(ctx *WindowSpecificationContext) {}

// EnterFrameClause is called when production frameClause is entered.
func (s *BaseMySQLStatementListener) EnterFrameClause(ctx *FrameClauseContext) {}

// ExitFrameClause is called when production frameClause is exited.
func (s *BaseMySQLStatementListener) ExitFrameClause(ctx *FrameClauseContext) {}

// EnterFrameStart is called when production frameStart is entered.
func (s *BaseMySQLStatementListener) EnterFrameStart(ctx *FrameStartContext) {}

// ExitFrameStart is called when production frameStart is exited.
func (s *BaseMySQLStatementListener) ExitFrameStart(ctx *FrameStartContext) {}

// EnterFrameEnd is called when production frameEnd is entered.
func (s *BaseMySQLStatementListener) EnterFrameEnd(ctx *FrameEndContext) {}

// ExitFrameEnd is called when production frameEnd is exited.
func (s *BaseMySQLStatementListener) ExitFrameEnd(ctx *FrameEndContext) {}

// EnterFrameBetween is called when production frameBetween is entered.
func (s *BaseMySQLStatementListener) EnterFrameBetween(ctx *FrameBetweenContext) {}

// ExitFrameBetween is called when production frameBetween is exited.
func (s *BaseMySQLStatementListener) ExitFrameBetween(ctx *FrameBetweenContext) {}

// EnterSpecialFunction is called when production specialFunction is entered.
func (s *BaseMySQLStatementListener) EnterSpecialFunction(ctx *SpecialFunctionContext) {}

// ExitSpecialFunction is called when production specialFunction is exited.
func (s *BaseMySQLStatementListener) ExitSpecialFunction(ctx *SpecialFunctionContext) {}

// EnterCurrentUserFunction is called when production currentUserFunction is entered.
func (s *BaseMySQLStatementListener) EnterCurrentUserFunction(ctx *CurrentUserFunctionContext) {}

// ExitCurrentUserFunction is called when production currentUserFunction is exited.
func (s *BaseMySQLStatementListener) ExitCurrentUserFunction(ctx *CurrentUserFunctionContext) {}

// EnterGroupConcatFunction is called when production groupConcatFunction is entered.
func (s *BaseMySQLStatementListener) EnterGroupConcatFunction(ctx *GroupConcatFunctionContext) {}

// ExitGroupConcatFunction is called when production groupConcatFunction is exited.
func (s *BaseMySQLStatementListener) ExitGroupConcatFunction(ctx *GroupConcatFunctionContext) {}

// EnterWindowFunction is called when production windowFunction is entered.
func (s *BaseMySQLStatementListener) EnterWindowFunction(ctx *WindowFunctionContext) {}

// ExitWindowFunction is called when production windowFunction is exited.
func (s *BaseMySQLStatementListener) ExitWindowFunction(ctx *WindowFunctionContext) {}

// EnterWindowingClause is called when production windowingClause is entered.
func (s *BaseMySQLStatementListener) EnterWindowingClause(ctx *WindowingClauseContext) {}

// ExitWindowingClause is called when production windowingClause is exited.
func (s *BaseMySQLStatementListener) ExitWindowingClause(ctx *WindowingClauseContext) {}

// EnterLeadLagInfo is called when production leadLagInfo is entered.
func (s *BaseMySQLStatementListener) EnterLeadLagInfo(ctx *LeadLagInfoContext) {}

// ExitLeadLagInfo is called when production leadLagInfo is exited.
func (s *BaseMySQLStatementListener) ExitLeadLagInfo(ctx *LeadLagInfoContext) {}

// EnterNullTreatment is called when production nullTreatment is entered.
func (s *BaseMySQLStatementListener) EnterNullTreatment(ctx *NullTreatmentContext) {}

// ExitNullTreatment is called when production nullTreatment is exited.
func (s *BaseMySQLStatementListener) ExitNullTreatment(ctx *NullTreatmentContext) {}

// EnterCheckType is called when production checkType is entered.
func (s *BaseMySQLStatementListener) EnterCheckType(ctx *CheckTypeContext) {}

// ExitCheckType is called when production checkType is exited.
func (s *BaseMySQLStatementListener) ExitCheckType(ctx *CheckTypeContext) {}

// EnterRepairType is called when production repairType is entered.
func (s *BaseMySQLStatementListener) EnterRepairType(ctx *RepairTypeContext) {}

// ExitRepairType is called when production repairType is exited.
func (s *BaseMySQLStatementListener) ExitRepairType(ctx *RepairTypeContext) {}

// EnterCastFunction is called when production castFunction is entered.
func (s *BaseMySQLStatementListener) EnterCastFunction(ctx *CastFunctionContext) {}

// ExitCastFunction is called when production castFunction is exited.
func (s *BaseMySQLStatementListener) ExitCastFunction(ctx *CastFunctionContext) {}

// EnterConvertFunction is called when production convertFunction is entered.
func (s *BaseMySQLStatementListener) EnterConvertFunction(ctx *ConvertFunctionContext) {}

// ExitConvertFunction is called when production convertFunction is exited.
func (s *BaseMySQLStatementListener) ExitConvertFunction(ctx *ConvertFunctionContext) {}

// EnterCastType is called when production castType is entered.
func (s *BaseMySQLStatementListener) EnterCastType(ctx *CastTypeContext) {}

// ExitCastType is called when production castType is exited.
func (s *BaseMySQLStatementListener) ExitCastType(ctx *CastTypeContext) {}

// EnterNchar is called when production nchar is entered.
func (s *BaseMySQLStatementListener) EnterNchar(ctx *NcharContext) {}

// ExitNchar is called when production nchar is exited.
func (s *BaseMySQLStatementListener) ExitNchar(ctx *NcharContext) {}

// EnterPositionFunction is called when production positionFunction is entered.
func (s *BaseMySQLStatementListener) EnterPositionFunction(ctx *PositionFunctionContext) {}

// ExitPositionFunction is called when production positionFunction is exited.
func (s *BaseMySQLStatementListener) ExitPositionFunction(ctx *PositionFunctionContext) {}

// EnterSubstringFunction is called when production substringFunction is entered.
func (s *BaseMySQLStatementListener) EnterSubstringFunction(ctx *SubstringFunctionContext) {}

// ExitSubstringFunction is called when production substringFunction is exited.
func (s *BaseMySQLStatementListener) ExitSubstringFunction(ctx *SubstringFunctionContext) {}

// EnterExtractFunction is called when production extractFunction is entered.
func (s *BaseMySQLStatementListener) EnterExtractFunction(ctx *ExtractFunctionContext) {}

// ExitExtractFunction is called when production extractFunction is exited.
func (s *BaseMySQLStatementListener) ExitExtractFunction(ctx *ExtractFunctionContext) {}

// EnterCharFunction is called when production charFunction is entered.
func (s *BaseMySQLStatementListener) EnterCharFunction(ctx *CharFunctionContext) {}

// ExitCharFunction is called when production charFunction is exited.
func (s *BaseMySQLStatementListener) ExitCharFunction(ctx *CharFunctionContext) {}

// EnterTrimFunction is called when production trimFunction is entered.
func (s *BaseMySQLStatementListener) EnterTrimFunction(ctx *TrimFunctionContext) {}

// ExitTrimFunction is called when production trimFunction is exited.
func (s *BaseMySQLStatementListener) ExitTrimFunction(ctx *TrimFunctionContext) {}

// EnterValuesFunction is called when production valuesFunction is entered.
func (s *BaseMySQLStatementListener) EnterValuesFunction(ctx *ValuesFunctionContext) {}

// ExitValuesFunction is called when production valuesFunction is exited.
func (s *BaseMySQLStatementListener) ExitValuesFunction(ctx *ValuesFunctionContext) {}

// EnterWeightStringFunction is called when production weightStringFunction is entered.
func (s *BaseMySQLStatementListener) EnterWeightStringFunction(ctx *WeightStringFunctionContext) {}

// ExitWeightStringFunction is called when production weightStringFunction is exited.
func (s *BaseMySQLStatementListener) ExitWeightStringFunction(ctx *WeightStringFunctionContext) {}

// EnterLevelClause is called when production levelClause is entered.
func (s *BaseMySQLStatementListener) EnterLevelClause(ctx *LevelClauseContext) {}

// ExitLevelClause is called when production levelClause is exited.
func (s *BaseMySQLStatementListener) ExitLevelClause(ctx *LevelClauseContext) {}

// EnterLevelInWeightListElement is called when production levelInWeightListElement is entered.
func (s *BaseMySQLStatementListener) EnterLevelInWeightListElement(ctx *LevelInWeightListElementContext) {
}

// ExitLevelInWeightListElement is called when production levelInWeightListElement is exited.
func (s *BaseMySQLStatementListener) ExitLevelInWeightListElement(ctx *LevelInWeightListElementContext) {
}

// EnterRegularFunction is called when production regularFunction is entered.
func (s *BaseMySQLStatementListener) EnterRegularFunction(ctx *RegularFunctionContext) {}

// ExitRegularFunction is called when production regularFunction is exited.
func (s *BaseMySQLStatementListener) ExitRegularFunction(ctx *RegularFunctionContext) {}

// EnterShorthandRegularFunction is called when production shorthandRegularFunction is entered.
func (s *BaseMySQLStatementListener) EnterShorthandRegularFunction(ctx *ShorthandRegularFunctionContext) {
}

// ExitShorthandRegularFunction is called when production shorthandRegularFunction is exited.
func (s *BaseMySQLStatementListener) ExitShorthandRegularFunction(ctx *ShorthandRegularFunctionContext) {
}

// EnterCompleteRegularFunction is called when production completeRegularFunction is entered.
func (s *BaseMySQLStatementListener) EnterCompleteRegularFunction(ctx *CompleteRegularFunctionContext) {
}

// ExitCompleteRegularFunction is called when production completeRegularFunction is exited.
func (s *BaseMySQLStatementListener) ExitCompleteRegularFunction(ctx *CompleteRegularFunctionContext) {
}

// EnterRegularFunctionName is called when production regularFunctionName is entered.
func (s *BaseMySQLStatementListener) EnterRegularFunctionName(ctx *RegularFunctionNameContext) {}

// ExitRegularFunctionName is called when production regularFunctionName is exited.
func (s *BaseMySQLStatementListener) ExitRegularFunctionName(ctx *RegularFunctionNameContext) {}

// EnterMatchExpression is called when production matchExpression is entered.
func (s *BaseMySQLStatementListener) EnterMatchExpression(ctx *MatchExpressionContext) {}

// ExitMatchExpression is called when production matchExpression is exited.
func (s *BaseMySQLStatementListener) ExitMatchExpression(ctx *MatchExpressionContext) {}

// EnterMatchSearchModifier is called when production matchSearchModifier is entered.
func (s *BaseMySQLStatementListener) EnterMatchSearchModifier(ctx *MatchSearchModifierContext) {}

// ExitMatchSearchModifier is called when production matchSearchModifier is exited.
func (s *BaseMySQLStatementListener) ExitMatchSearchModifier(ctx *MatchSearchModifierContext) {}

// EnterCaseExpression is called when production caseExpression is entered.
func (s *BaseMySQLStatementListener) EnterCaseExpression(ctx *CaseExpressionContext) {}

// ExitCaseExpression is called when production caseExpression is exited.
func (s *BaseMySQLStatementListener) ExitCaseExpression(ctx *CaseExpressionContext) {}

// EnterDatetimeExpr is called when production datetimeExpr is entered.
func (s *BaseMySQLStatementListener) EnterDatetimeExpr(ctx *DatetimeExprContext) {}

// ExitDatetimeExpr is called when production datetimeExpr is exited.
func (s *BaseMySQLStatementListener) ExitDatetimeExpr(ctx *DatetimeExprContext) {}

// EnterBinaryLogFileIndexNumber is called when production binaryLogFileIndexNumber is entered.
func (s *BaseMySQLStatementListener) EnterBinaryLogFileIndexNumber(ctx *BinaryLogFileIndexNumberContext) {
}

// ExitBinaryLogFileIndexNumber is called when production binaryLogFileIndexNumber is exited.
func (s *BaseMySQLStatementListener) ExitBinaryLogFileIndexNumber(ctx *BinaryLogFileIndexNumberContext) {
}

// EnterCaseWhen is called when production caseWhen is entered.
func (s *BaseMySQLStatementListener) EnterCaseWhen(ctx *CaseWhenContext) {}

// ExitCaseWhen is called when production caseWhen is exited.
func (s *BaseMySQLStatementListener) ExitCaseWhen(ctx *CaseWhenContext) {}

// EnterCaseElse is called when production caseElse is entered.
func (s *BaseMySQLStatementListener) EnterCaseElse(ctx *CaseElseContext) {}

// ExitCaseElse is called when production caseElse is exited.
func (s *BaseMySQLStatementListener) ExitCaseElse(ctx *CaseElseContext) {}

// EnterIntervalExpression is called when production intervalExpression is entered.
func (s *BaseMySQLStatementListener) EnterIntervalExpression(ctx *IntervalExpressionContext) {}

// ExitIntervalExpression is called when production intervalExpression is exited.
func (s *BaseMySQLStatementListener) ExitIntervalExpression(ctx *IntervalExpressionContext) {}

// EnterIntervalValue is called when production intervalValue is entered.
func (s *BaseMySQLStatementListener) EnterIntervalValue(ctx *IntervalValueContext) {}

// ExitIntervalValue is called when production intervalValue is exited.
func (s *BaseMySQLStatementListener) ExitIntervalValue(ctx *IntervalValueContext) {}

// EnterIntervalUnit is called when production intervalUnit is entered.
func (s *BaseMySQLStatementListener) EnterIntervalUnit(ctx *IntervalUnitContext) {}

// ExitIntervalUnit is called when production intervalUnit is exited.
func (s *BaseMySQLStatementListener) ExitIntervalUnit(ctx *IntervalUnitContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseMySQLStatementListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseMySQLStatementListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderByItem is called when production orderByItem is entered.
func (s *BaseMySQLStatementListener) EnterOrderByItem(ctx *OrderByItemContext) {}

// ExitOrderByItem is called when production orderByItem is exited.
func (s *BaseMySQLStatementListener) ExitOrderByItem(ctx *OrderByItemContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseMySQLStatementListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseMySQLStatementListener) ExitDataType(ctx *DataTypeContext) {}

// EnterStringList is called when production stringList is entered.
func (s *BaseMySQLStatementListener) EnterStringList(ctx *StringListContext) {}

// ExitStringList is called when production stringList is exited.
func (s *BaseMySQLStatementListener) ExitStringList(ctx *StringListContext) {}

// EnterTextString is called when production textString is entered.
func (s *BaseMySQLStatementListener) EnterTextString(ctx *TextStringContext) {}

// ExitTextString is called when production textString is exited.
func (s *BaseMySQLStatementListener) ExitTextString(ctx *TextStringContext) {}

// EnterTextStringHash is called when production textStringHash is entered.
func (s *BaseMySQLStatementListener) EnterTextStringHash(ctx *TextStringHashContext) {}

// ExitTextStringHash is called when production textStringHash is exited.
func (s *BaseMySQLStatementListener) ExitTextStringHash(ctx *TextStringHashContext) {}

// EnterFieldOptions is called when production fieldOptions is entered.
func (s *BaseMySQLStatementListener) EnterFieldOptions(ctx *FieldOptionsContext) {}

// ExitFieldOptions is called when production fieldOptions is exited.
func (s *BaseMySQLStatementListener) ExitFieldOptions(ctx *FieldOptionsContext) {}

// EnterPrecision is called when production precision is entered.
func (s *BaseMySQLStatementListener) EnterPrecision(ctx *PrecisionContext) {}

// ExitPrecision is called when production precision is exited.
func (s *BaseMySQLStatementListener) ExitPrecision(ctx *PrecisionContext) {}

// EnterTypeDatetimePrecision is called when production typeDatetimePrecision is entered.
func (s *BaseMySQLStatementListener) EnterTypeDatetimePrecision(ctx *TypeDatetimePrecisionContext) {}

// ExitTypeDatetimePrecision is called when production typeDatetimePrecision is exited.
func (s *BaseMySQLStatementListener) ExitTypeDatetimePrecision(ctx *TypeDatetimePrecisionContext) {}

// EnterCharsetWithOptBinary is called when production charsetWithOptBinary is entered.
func (s *BaseMySQLStatementListener) EnterCharsetWithOptBinary(ctx *CharsetWithOptBinaryContext) {}

// ExitCharsetWithOptBinary is called when production charsetWithOptBinary is exited.
func (s *BaseMySQLStatementListener) ExitCharsetWithOptBinary(ctx *CharsetWithOptBinaryContext) {}

// EnterAscii is called when production ascii is entered.
func (s *BaseMySQLStatementListener) EnterAscii(ctx *AsciiContext) {}

// ExitAscii is called when production ascii is exited.
func (s *BaseMySQLStatementListener) ExitAscii(ctx *AsciiContext) {}

// EnterUnicode is called when production unicode is entered.
func (s *BaseMySQLStatementListener) EnterUnicode(ctx *UnicodeContext) {}

// ExitUnicode is called when production unicode is exited.
func (s *BaseMySQLStatementListener) ExitUnicode(ctx *UnicodeContext) {}

// EnterCharset is called when production charset is entered.
func (s *BaseMySQLStatementListener) EnterCharset(ctx *CharsetContext) {}

// ExitCharset is called when production charset is exited.
func (s *BaseMySQLStatementListener) ExitCharset(ctx *CharsetContext) {}

// EnterDefaultCollation is called when production defaultCollation is entered.
func (s *BaseMySQLStatementListener) EnterDefaultCollation(ctx *DefaultCollationContext) {}

// ExitDefaultCollation is called when production defaultCollation is exited.
func (s *BaseMySQLStatementListener) ExitDefaultCollation(ctx *DefaultCollationContext) {}

// EnterDefaultEncryption is called when production defaultEncryption is entered.
func (s *BaseMySQLStatementListener) EnterDefaultEncryption(ctx *DefaultEncryptionContext) {}

// ExitDefaultEncryption is called when production defaultEncryption is exited.
func (s *BaseMySQLStatementListener) ExitDefaultEncryption(ctx *DefaultEncryptionContext) {}

// EnterDefaultCharset is called when production defaultCharset is entered.
func (s *BaseMySQLStatementListener) EnterDefaultCharset(ctx *DefaultCharsetContext) {}

// ExitDefaultCharset is called when production defaultCharset is exited.
func (s *BaseMySQLStatementListener) ExitDefaultCharset(ctx *DefaultCharsetContext) {}

// EnterSignedLiteral is called when production signedLiteral is entered.
func (s *BaseMySQLStatementListener) EnterSignedLiteral(ctx *SignedLiteralContext) {}

// ExitSignedLiteral is called when production signedLiteral is exited.
func (s *BaseMySQLStatementListener) ExitSignedLiteral(ctx *SignedLiteralContext) {}

// EnterNow is called when production now is entered.
func (s *BaseMySQLStatementListener) EnterNow(ctx *NowContext) {}

// ExitNow is called when production now is exited.
func (s *BaseMySQLStatementListener) ExitNow(ctx *NowContext) {}

// EnterColumnFormat is called when production columnFormat is entered.
func (s *BaseMySQLStatementListener) EnterColumnFormat(ctx *ColumnFormatContext) {}

// ExitColumnFormat is called when production columnFormat is exited.
func (s *BaseMySQLStatementListener) ExitColumnFormat(ctx *ColumnFormatContext) {}

// EnterStorageMedia is called when production storageMedia is entered.
func (s *BaseMySQLStatementListener) EnterStorageMedia(ctx *StorageMediaContext) {}

// ExitStorageMedia is called when production storageMedia is exited.
func (s *BaseMySQLStatementListener) ExitStorageMedia(ctx *StorageMediaContext) {}

// EnterDirection is called when production direction is entered.
func (s *BaseMySQLStatementListener) EnterDirection(ctx *DirectionContext) {}

// ExitDirection is called when production direction is exited.
func (s *BaseMySQLStatementListener) ExitDirection(ctx *DirectionContext) {}

// EnterKeyOrIndex is called when production keyOrIndex is entered.
func (s *BaseMySQLStatementListener) EnterKeyOrIndex(ctx *KeyOrIndexContext) {}

// ExitKeyOrIndex is called when production keyOrIndex is exited.
func (s *BaseMySQLStatementListener) ExitKeyOrIndex(ctx *KeyOrIndexContext) {}

// EnterFieldLength is called when production fieldLength is entered.
func (s *BaseMySQLStatementListener) EnterFieldLength(ctx *FieldLengthContext) {}

// ExitFieldLength is called when production fieldLength is exited.
func (s *BaseMySQLStatementListener) ExitFieldLength(ctx *FieldLengthContext) {}

// EnterCharacterSet is called when production characterSet is entered.
func (s *BaseMySQLStatementListener) EnterCharacterSet(ctx *CharacterSetContext) {}

// ExitCharacterSet is called when production characterSet is exited.
func (s *BaseMySQLStatementListener) ExitCharacterSet(ctx *CharacterSetContext) {}

// EnterCollateClause is called when production collateClause is entered.
func (s *BaseMySQLStatementListener) EnterCollateClause(ctx *CollateClauseContext) {}

// ExitCollateClause is called when production collateClause is exited.
func (s *BaseMySQLStatementListener) ExitCollateClause(ctx *CollateClauseContext) {}

// EnterFieldOrVarSpec is called when production fieldOrVarSpec is entered.
func (s *BaseMySQLStatementListener) EnterFieldOrVarSpec(ctx *FieldOrVarSpecContext) {}

// ExitFieldOrVarSpec is called when production fieldOrVarSpec is exited.
func (s *BaseMySQLStatementListener) ExitFieldOrVarSpec(ctx *FieldOrVarSpecContext) {}

// EnterNotExistClause is called when production notExistClause is entered.
func (s *BaseMySQLStatementListener) EnterNotExistClause(ctx *NotExistClauseContext) {}

// ExitNotExistClause is called when production notExistClause is exited.
func (s *BaseMySQLStatementListener) ExitNotExistClause(ctx *NotExistClauseContext) {}

// EnterExistClause is called when production existClause is entered.
func (s *BaseMySQLStatementListener) EnterExistClause(ctx *ExistClauseContext) {}

// ExitExistClause is called when production existClause is exited.
func (s *BaseMySQLStatementListener) ExitExistClause(ctx *ExistClauseContext) {}

// EnterConnectionId is called when production connectionId is entered.
func (s *BaseMySQLStatementListener) EnterConnectionId(ctx *ConnectionIdContext) {}

// ExitConnectionId is called when production connectionId is exited.
func (s *BaseMySQLStatementListener) ExitConnectionId(ctx *ConnectionIdContext) {}

// EnterLabelName is called when production labelName is entered.
func (s *BaseMySQLStatementListener) EnterLabelName(ctx *LabelNameContext) {}

// ExitLabelName is called when production labelName is exited.
func (s *BaseMySQLStatementListener) ExitLabelName(ctx *LabelNameContext) {}

// EnterCursorName is called when production cursorName is entered.
func (s *BaseMySQLStatementListener) EnterCursorName(ctx *CursorNameContext) {}

// ExitCursorName is called when production cursorName is exited.
func (s *BaseMySQLStatementListener) ExitCursorName(ctx *CursorNameContext) {}

// EnterConditionName is called when production conditionName is entered.
func (s *BaseMySQLStatementListener) EnterConditionName(ctx *ConditionNameContext) {}

// ExitConditionName is called when production conditionName is exited.
func (s *BaseMySQLStatementListener) ExitConditionName(ctx *ConditionNameContext) {}

// EnterUnionOption is called when production unionOption is entered.
func (s *BaseMySQLStatementListener) EnterUnionOption(ctx *UnionOptionContext) {}

// ExitUnionOption is called when production unionOption is exited.
func (s *BaseMySQLStatementListener) ExitUnionOption(ctx *UnionOptionContext) {}

// EnterNoWriteToBinLog is called when production noWriteToBinLog is entered.
func (s *BaseMySQLStatementListener) EnterNoWriteToBinLog(ctx *NoWriteToBinLogContext) {}

// ExitNoWriteToBinLog is called when production noWriteToBinLog is exited.
func (s *BaseMySQLStatementListener) ExitNoWriteToBinLog(ctx *NoWriteToBinLogContext) {}

// EnterChannelOption is called when production channelOption is entered.
func (s *BaseMySQLStatementListener) EnterChannelOption(ctx *ChannelOptionContext) {}

// ExitChannelOption is called when production channelOption is exited.
func (s *BaseMySQLStatementListener) ExitChannelOption(ctx *ChannelOptionContext) {}

// EnterPreparedStatement is called when production preparedStatement is entered.
func (s *BaseMySQLStatementListener) EnterPreparedStatement(ctx *PreparedStatementContext) {}

// ExitPreparedStatement is called when production preparedStatement is exited.
func (s *BaseMySQLStatementListener) ExitPreparedStatement(ctx *PreparedStatementContext) {}

// EnterExecuteStatement is called when production executeStatement is entered.
func (s *BaseMySQLStatementListener) EnterExecuteStatement(ctx *ExecuteStatementContext) {}

// ExitExecuteStatement is called when production executeStatement is exited.
func (s *BaseMySQLStatementListener) ExitExecuteStatement(ctx *ExecuteStatementContext) {}

// EnterExecuteVarList is called when production executeVarList is entered.
func (s *BaseMySQLStatementListener) EnterExecuteVarList(ctx *ExecuteVarListContext) {}

// ExitExecuteVarList is called when production executeVarList is exited.
func (s *BaseMySQLStatementListener) ExitExecuteVarList(ctx *ExecuteVarListContext) {}

// EnterAlterStatement is called when production alterStatement is entered.
func (s *BaseMySQLStatementListener) EnterAlterStatement(ctx *AlterStatementContext) {}

// ExitAlterStatement is called when production alterStatement is exited.
func (s *BaseMySQLStatementListener) ExitAlterStatement(ctx *AlterStatementContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseMySQLStatementListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseMySQLStatementListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterPartitionClause is called when production partitionClause is entered.
func (s *BaseMySQLStatementListener) EnterPartitionClause(ctx *PartitionClauseContext) {}

// ExitPartitionClause is called when production partitionClause is exited.
func (s *BaseMySQLStatementListener) ExitPartitionClause(ctx *PartitionClauseContext) {}

// EnterPartitionTypeDef is called when production partitionTypeDef is entered.
func (s *BaseMySQLStatementListener) EnterPartitionTypeDef(ctx *PartitionTypeDefContext) {}

// ExitPartitionTypeDef is called when production partitionTypeDef is exited.
func (s *BaseMySQLStatementListener) ExitPartitionTypeDef(ctx *PartitionTypeDefContext) {}

// EnterSubPartitions is called when production subPartitions is entered.
func (s *BaseMySQLStatementListener) EnterSubPartitions(ctx *SubPartitionsContext) {}

// ExitSubPartitions is called when production subPartitions is exited.
func (s *BaseMySQLStatementListener) ExitSubPartitions(ctx *SubPartitionsContext) {}

// EnterPartitionKeyAlgorithm is called when production partitionKeyAlgorithm is entered.
func (s *BaseMySQLStatementListener) EnterPartitionKeyAlgorithm(ctx *PartitionKeyAlgorithmContext) {}

// ExitPartitionKeyAlgorithm is called when production partitionKeyAlgorithm is exited.
func (s *BaseMySQLStatementListener) ExitPartitionKeyAlgorithm(ctx *PartitionKeyAlgorithmContext) {}

// EnterDuplicateAsQueryExpression is called when production duplicateAsQueryExpression is entered.
func (s *BaseMySQLStatementListener) EnterDuplicateAsQueryExpression(ctx *DuplicateAsQueryExpressionContext) {
}

// ExitDuplicateAsQueryExpression is called when production duplicateAsQueryExpression is exited.
func (s *BaseMySQLStatementListener) ExitDuplicateAsQueryExpression(ctx *DuplicateAsQueryExpressionContext) {
}

// EnterAlterTable is called when production alterTable is entered.
func (s *BaseMySQLStatementListener) EnterAlterTable(ctx *AlterTableContext) {}

// ExitAlterTable is called when production alterTable is exited.
func (s *BaseMySQLStatementListener) ExitAlterTable(ctx *AlterTableContext) {}

// EnterStandaloneAlterTableAction is called when production standaloneAlterTableAction is entered.
func (s *BaseMySQLStatementListener) EnterStandaloneAlterTableAction(ctx *StandaloneAlterTableActionContext) {
}

// ExitStandaloneAlterTableAction is called when production standaloneAlterTableAction is exited.
func (s *BaseMySQLStatementListener) ExitStandaloneAlterTableAction(ctx *StandaloneAlterTableActionContext) {
}

// EnterAlterTableActions is called when production alterTableActions is entered.
func (s *BaseMySQLStatementListener) EnterAlterTableActions(ctx *AlterTableActionsContext) {}

// ExitAlterTableActions is called when production alterTableActions is exited.
func (s *BaseMySQLStatementListener) ExitAlterTableActions(ctx *AlterTableActionsContext) {}

// EnterAlterTablePartitionOptions is called when production alterTablePartitionOptions is entered.
func (s *BaseMySQLStatementListener) EnterAlterTablePartitionOptions(ctx *AlterTablePartitionOptionsContext) {
}

// ExitAlterTablePartitionOptions is called when production alterTablePartitionOptions is exited.
func (s *BaseMySQLStatementListener) ExitAlterTablePartitionOptions(ctx *AlterTablePartitionOptionsContext) {
}

// EnterAlterCommandList is called when production alterCommandList is entered.
func (s *BaseMySQLStatementListener) EnterAlterCommandList(ctx *AlterCommandListContext) {}

// ExitAlterCommandList is called when production alterCommandList is exited.
func (s *BaseMySQLStatementListener) ExitAlterCommandList(ctx *AlterCommandListContext) {}

// EnterAlterList is called when production alterList is entered.
func (s *BaseMySQLStatementListener) EnterAlterList(ctx *AlterListContext) {}

// ExitAlterList is called when production alterList is exited.
func (s *BaseMySQLStatementListener) ExitAlterList(ctx *AlterListContext) {}

// EnterCreateTableOptionsSpaceSeparated is called when production createTableOptionsSpaceSeparated is entered.
func (s *BaseMySQLStatementListener) EnterCreateTableOptionsSpaceSeparated(ctx *CreateTableOptionsSpaceSeparatedContext) {
}

// ExitCreateTableOptionsSpaceSeparated is called when production createTableOptionsSpaceSeparated is exited.
func (s *BaseMySQLStatementListener) ExitCreateTableOptionsSpaceSeparated(ctx *CreateTableOptionsSpaceSeparatedContext) {
}

// EnterAddColumn is called when production addColumn is entered.
func (s *BaseMySQLStatementListener) EnterAddColumn(ctx *AddColumnContext) {}

// ExitAddColumn is called when production addColumn is exited.
func (s *BaseMySQLStatementListener) ExitAddColumn(ctx *AddColumnContext) {}

// EnterAddTableConstraint is called when production addTableConstraint is entered.
func (s *BaseMySQLStatementListener) EnterAddTableConstraint(ctx *AddTableConstraintContext) {}

// ExitAddTableConstraint is called when production addTableConstraint is exited.
func (s *BaseMySQLStatementListener) ExitAddTableConstraint(ctx *AddTableConstraintContext) {}

// EnterChangeColumn is called when production changeColumn is entered.
func (s *BaseMySQLStatementListener) EnterChangeColumn(ctx *ChangeColumnContext) {}

// ExitChangeColumn is called when production changeColumn is exited.
func (s *BaseMySQLStatementListener) ExitChangeColumn(ctx *ChangeColumnContext) {}

// EnterModifyColumn is called when production modifyColumn is entered.
func (s *BaseMySQLStatementListener) EnterModifyColumn(ctx *ModifyColumnContext) {}

// ExitModifyColumn is called when production modifyColumn is exited.
func (s *BaseMySQLStatementListener) ExitModifyColumn(ctx *ModifyColumnContext) {}

// EnterAlterTableDrop is called when production alterTableDrop is entered.
func (s *BaseMySQLStatementListener) EnterAlterTableDrop(ctx *AlterTableDropContext) {}

// ExitAlterTableDrop is called when production alterTableDrop is exited.
func (s *BaseMySQLStatementListener) ExitAlterTableDrop(ctx *AlterTableDropContext) {}

// EnterDisableKeys is called when production disableKeys is entered.
func (s *BaseMySQLStatementListener) EnterDisableKeys(ctx *DisableKeysContext) {}

// ExitDisableKeys is called when production disableKeys is exited.
func (s *BaseMySQLStatementListener) ExitDisableKeys(ctx *DisableKeysContext) {}

// EnterEnableKeys is called when production enableKeys is entered.
func (s *BaseMySQLStatementListener) EnterEnableKeys(ctx *EnableKeysContext) {}

// ExitEnableKeys is called when production enableKeys is exited.
func (s *BaseMySQLStatementListener) ExitEnableKeys(ctx *EnableKeysContext) {}

// EnterAlterColumn is called when production alterColumn is entered.
func (s *BaseMySQLStatementListener) EnterAlterColumn(ctx *AlterColumnContext) {}

// ExitAlterColumn is called when production alterColumn is exited.
func (s *BaseMySQLStatementListener) ExitAlterColumn(ctx *AlterColumnContext) {}

// EnterAlterIndex is called when production alterIndex is entered.
func (s *BaseMySQLStatementListener) EnterAlterIndex(ctx *AlterIndexContext) {}

// ExitAlterIndex is called when production alterIndex is exited.
func (s *BaseMySQLStatementListener) ExitAlterIndex(ctx *AlterIndexContext) {}

// EnterAlterCheck is called when production alterCheck is entered.
func (s *BaseMySQLStatementListener) EnterAlterCheck(ctx *AlterCheckContext) {}

// ExitAlterCheck is called when production alterCheck is exited.
func (s *BaseMySQLStatementListener) ExitAlterCheck(ctx *AlterCheckContext) {}

// EnterAlterConstraint is called when production alterConstraint is entered.
func (s *BaseMySQLStatementListener) EnterAlterConstraint(ctx *AlterConstraintContext) {}

// ExitAlterConstraint is called when production alterConstraint is exited.
func (s *BaseMySQLStatementListener) ExitAlterConstraint(ctx *AlterConstraintContext) {}

// EnterRenameColumn is called when production renameColumn is entered.
func (s *BaseMySQLStatementListener) EnterRenameColumn(ctx *RenameColumnContext) {}

// ExitRenameColumn is called when production renameColumn is exited.
func (s *BaseMySQLStatementListener) ExitRenameColumn(ctx *RenameColumnContext) {}

// EnterAlterRenameTable is called when production alterRenameTable is entered.
func (s *BaseMySQLStatementListener) EnterAlterRenameTable(ctx *AlterRenameTableContext) {}

// ExitAlterRenameTable is called when production alterRenameTable is exited.
func (s *BaseMySQLStatementListener) ExitAlterRenameTable(ctx *AlterRenameTableContext) {}

// EnterRenameIndex is called when production renameIndex is entered.
func (s *BaseMySQLStatementListener) EnterRenameIndex(ctx *RenameIndexContext) {}

// ExitRenameIndex is called when production renameIndex is exited.
func (s *BaseMySQLStatementListener) ExitRenameIndex(ctx *RenameIndexContext) {}

// EnterAlterConvert is called when production alterConvert is entered.
func (s *BaseMySQLStatementListener) EnterAlterConvert(ctx *AlterConvertContext) {}

// ExitAlterConvert is called when production alterConvert is exited.
func (s *BaseMySQLStatementListener) ExitAlterConvert(ctx *AlterConvertContext) {}

// EnterAlterTableForce is called when production alterTableForce is entered.
func (s *BaseMySQLStatementListener) EnterAlterTableForce(ctx *AlterTableForceContext) {}

// ExitAlterTableForce is called when production alterTableForce is exited.
func (s *BaseMySQLStatementListener) ExitAlterTableForce(ctx *AlterTableForceContext) {}

// EnterAlterTableOrder is called when production alterTableOrder is entered.
func (s *BaseMySQLStatementListener) EnterAlterTableOrder(ctx *AlterTableOrderContext) {}

// ExitAlterTableOrder is called when production alterTableOrder is exited.
func (s *BaseMySQLStatementListener) ExitAlterTableOrder(ctx *AlterTableOrderContext) {}

// EnterAlterOrderList is called when production alterOrderList is entered.
func (s *BaseMySQLStatementListener) EnterAlterOrderList(ctx *AlterOrderListContext) {}

// ExitAlterOrderList is called when production alterOrderList is exited.
func (s *BaseMySQLStatementListener) ExitAlterOrderList(ctx *AlterOrderListContext) {}

// EnterTableConstraintDef is called when production tableConstraintDef is entered.
func (s *BaseMySQLStatementListener) EnterTableConstraintDef(ctx *TableConstraintDefContext) {}

// ExitTableConstraintDef is called when production tableConstraintDef is exited.
func (s *BaseMySQLStatementListener) ExitTableConstraintDef(ctx *TableConstraintDefContext) {}

// EnterAlterCommandsModifierList is called when production alterCommandsModifierList is entered.
func (s *BaseMySQLStatementListener) EnterAlterCommandsModifierList(ctx *AlterCommandsModifierListContext) {
}

// ExitAlterCommandsModifierList is called when production alterCommandsModifierList is exited.
func (s *BaseMySQLStatementListener) ExitAlterCommandsModifierList(ctx *AlterCommandsModifierListContext) {
}

// EnterAlterCommandsModifier is called when production alterCommandsModifier is entered.
func (s *BaseMySQLStatementListener) EnterAlterCommandsModifier(ctx *AlterCommandsModifierContext) {}

// ExitAlterCommandsModifier is called when production alterCommandsModifier is exited.
func (s *BaseMySQLStatementListener) ExitAlterCommandsModifier(ctx *AlterCommandsModifierContext) {}

// EnterWithValidation is called when production withValidation is entered.
func (s *BaseMySQLStatementListener) EnterWithValidation(ctx *WithValidationContext) {}

// ExitWithValidation is called when production withValidation is exited.
func (s *BaseMySQLStatementListener) ExitWithValidation(ctx *WithValidationContext) {}

// EnterStandaloneAlterCommands is called when production standaloneAlterCommands is entered.
func (s *BaseMySQLStatementListener) EnterStandaloneAlterCommands(ctx *StandaloneAlterCommandsContext) {
}

// ExitStandaloneAlterCommands is called when production standaloneAlterCommands is exited.
func (s *BaseMySQLStatementListener) ExitStandaloneAlterCommands(ctx *StandaloneAlterCommandsContext) {
}

// EnterAlterPartition is called when production alterPartition is entered.
func (s *BaseMySQLStatementListener) EnterAlterPartition(ctx *AlterPartitionContext) {}

// ExitAlterPartition is called when production alterPartition is exited.
func (s *BaseMySQLStatementListener) ExitAlterPartition(ctx *AlterPartitionContext) {}

// EnterConstraintClause is called when production constraintClause is entered.
func (s *BaseMySQLStatementListener) EnterConstraintClause(ctx *ConstraintClauseContext) {}

// ExitConstraintClause is called when production constraintClause is exited.
func (s *BaseMySQLStatementListener) ExitConstraintClause(ctx *ConstraintClauseContext) {}

// EnterTableElementList is called when production tableElementList is entered.
func (s *BaseMySQLStatementListener) EnterTableElementList(ctx *TableElementListContext) {}

// ExitTableElementList is called when production tableElementList is exited.
func (s *BaseMySQLStatementListener) ExitTableElementList(ctx *TableElementListContext) {}

// EnterTableElement is called when production tableElement is entered.
func (s *BaseMySQLStatementListener) EnterTableElement(ctx *TableElementContext) {}

// ExitTableElement is called when production tableElement is exited.
func (s *BaseMySQLStatementListener) ExitTableElement(ctx *TableElementContext) {}

// EnterRestrict is called when production restrict is entered.
func (s *BaseMySQLStatementListener) EnterRestrict(ctx *RestrictContext) {}

// ExitRestrict is called when production restrict is exited.
func (s *BaseMySQLStatementListener) ExitRestrict(ctx *RestrictContext) {}

// EnterFulltextIndexOption is called when production fulltextIndexOption is entered.
func (s *BaseMySQLStatementListener) EnterFulltextIndexOption(ctx *FulltextIndexOptionContext) {}

// ExitFulltextIndexOption is called when production fulltextIndexOption is exited.
func (s *BaseMySQLStatementListener) ExitFulltextIndexOption(ctx *FulltextIndexOptionContext) {}

// EnterDropTable is called when production dropTable is entered.
func (s *BaseMySQLStatementListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production dropTable is exited.
func (s *BaseMySQLStatementListener) ExitDropTable(ctx *DropTableContext) {}

// EnterDropIndex is called when production dropIndex is entered.
func (s *BaseMySQLStatementListener) EnterDropIndex(ctx *DropIndexContext) {}

// ExitDropIndex is called when production dropIndex is exited.
func (s *BaseMySQLStatementListener) ExitDropIndex(ctx *DropIndexContext) {}

// EnterAlterAlgorithmOption is called when production alterAlgorithmOption is entered.
func (s *BaseMySQLStatementListener) EnterAlterAlgorithmOption(ctx *AlterAlgorithmOptionContext) {}

// ExitAlterAlgorithmOption is called when production alterAlgorithmOption is exited.
func (s *BaseMySQLStatementListener) ExitAlterAlgorithmOption(ctx *AlterAlgorithmOptionContext) {}

// EnterAlterLockOption is called when production alterLockOption is entered.
func (s *BaseMySQLStatementListener) EnterAlterLockOption(ctx *AlterLockOptionContext) {}

// ExitAlterLockOption is called when production alterLockOption is exited.
func (s *BaseMySQLStatementListener) ExitAlterLockOption(ctx *AlterLockOptionContext) {}

// EnterTruncateTable is called when production truncateTable is entered.
func (s *BaseMySQLStatementListener) EnterTruncateTable(ctx *TruncateTableContext) {}

// ExitTruncateTable is called when production truncateTable is exited.
func (s *BaseMySQLStatementListener) ExitTruncateTable(ctx *TruncateTableContext) {}

// EnterCreateIndex is called when production createIndex is entered.
func (s *BaseMySQLStatementListener) EnterCreateIndex(ctx *CreateIndexContext) {}

// ExitCreateIndex is called when production createIndex is exited.
func (s *BaseMySQLStatementListener) ExitCreateIndex(ctx *CreateIndexContext) {}

// EnterCreateDatabase is called when production createDatabase is entered.
func (s *BaseMySQLStatementListener) EnterCreateDatabase(ctx *CreateDatabaseContext) {}

// ExitCreateDatabase is called when production createDatabase is exited.
func (s *BaseMySQLStatementListener) ExitCreateDatabase(ctx *CreateDatabaseContext) {}

// EnterAlterDatabase is called when production alterDatabase is entered.
func (s *BaseMySQLStatementListener) EnterAlterDatabase(ctx *AlterDatabaseContext) {}

// ExitAlterDatabase is called when production alterDatabase is exited.
func (s *BaseMySQLStatementListener) ExitAlterDatabase(ctx *AlterDatabaseContext) {}

// EnterCreateDatabaseSpecification_ is called when production createDatabaseSpecification_ is entered.
func (s *BaseMySQLStatementListener) EnterCreateDatabaseSpecification_(ctx *CreateDatabaseSpecification_Context) {
}

// ExitCreateDatabaseSpecification_ is called when production createDatabaseSpecification_ is exited.
func (s *BaseMySQLStatementListener) ExitCreateDatabaseSpecification_(ctx *CreateDatabaseSpecification_Context) {
}

// EnterAlterDatabaseSpecification_ is called when production alterDatabaseSpecification_ is entered.
func (s *BaseMySQLStatementListener) EnterAlterDatabaseSpecification_(ctx *AlterDatabaseSpecification_Context) {
}

// ExitAlterDatabaseSpecification_ is called when production alterDatabaseSpecification_ is exited.
func (s *BaseMySQLStatementListener) ExitAlterDatabaseSpecification_(ctx *AlterDatabaseSpecification_Context) {
}

// EnterDropDatabase is called when production dropDatabase is entered.
func (s *BaseMySQLStatementListener) EnterDropDatabase(ctx *DropDatabaseContext) {}

// ExitDropDatabase is called when production dropDatabase is exited.
func (s *BaseMySQLStatementListener) ExitDropDatabase(ctx *DropDatabaseContext) {}

// EnterAlterInstance is called when production alterInstance is entered.
func (s *BaseMySQLStatementListener) EnterAlterInstance(ctx *AlterInstanceContext) {}

// ExitAlterInstance is called when production alterInstance is exited.
func (s *BaseMySQLStatementListener) ExitAlterInstance(ctx *AlterInstanceContext) {}

// EnterInstanceAction is called when production instanceAction is entered.
func (s *BaseMySQLStatementListener) EnterInstanceAction(ctx *InstanceActionContext) {}

// ExitInstanceAction is called when production instanceAction is exited.
func (s *BaseMySQLStatementListener) ExitInstanceAction(ctx *InstanceActionContext) {}

// EnterChannel is called when production channel is entered.
func (s *BaseMySQLStatementListener) EnterChannel(ctx *ChannelContext) {}

// ExitChannel is called when production channel is exited.
func (s *BaseMySQLStatementListener) ExitChannel(ctx *ChannelContext) {}

// EnterCreateEvent is called when production createEvent is entered.
func (s *BaseMySQLStatementListener) EnterCreateEvent(ctx *CreateEventContext) {}

// ExitCreateEvent is called when production createEvent is exited.
func (s *BaseMySQLStatementListener) ExitCreateEvent(ctx *CreateEventContext) {}

// EnterAlterEvent is called when production alterEvent is entered.
func (s *BaseMySQLStatementListener) EnterAlterEvent(ctx *AlterEventContext) {}

// ExitAlterEvent is called when production alterEvent is exited.
func (s *BaseMySQLStatementListener) ExitAlterEvent(ctx *AlterEventContext) {}

// EnterDropEvent is called when production dropEvent is entered.
func (s *BaseMySQLStatementListener) EnterDropEvent(ctx *DropEventContext) {}

// ExitDropEvent is called when production dropEvent is exited.
func (s *BaseMySQLStatementListener) ExitDropEvent(ctx *DropEventContext) {}

// EnterCreateFunction is called when production createFunction is entered.
func (s *BaseMySQLStatementListener) EnterCreateFunction(ctx *CreateFunctionContext) {}

// ExitCreateFunction is called when production createFunction is exited.
func (s *BaseMySQLStatementListener) ExitCreateFunction(ctx *CreateFunctionContext) {}

// EnterAlterFunction is called when production alterFunction is entered.
func (s *BaseMySQLStatementListener) EnterAlterFunction(ctx *AlterFunctionContext) {}

// ExitAlterFunction is called when production alterFunction is exited.
func (s *BaseMySQLStatementListener) ExitAlterFunction(ctx *AlterFunctionContext) {}

// EnterDropFunction is called when production dropFunction is entered.
func (s *BaseMySQLStatementListener) EnterDropFunction(ctx *DropFunctionContext) {}

// ExitDropFunction is called when production dropFunction is exited.
func (s *BaseMySQLStatementListener) ExitDropFunction(ctx *DropFunctionContext) {}

// EnterCreateProcedure is called when production createProcedure is entered.
func (s *BaseMySQLStatementListener) EnterCreateProcedure(ctx *CreateProcedureContext) {}

// ExitCreateProcedure is called when production createProcedure is exited.
func (s *BaseMySQLStatementListener) ExitCreateProcedure(ctx *CreateProcedureContext) {}

// EnterAlterProcedure is called when production alterProcedure is entered.
func (s *BaseMySQLStatementListener) EnterAlterProcedure(ctx *AlterProcedureContext) {}

// ExitAlterProcedure is called when production alterProcedure is exited.
func (s *BaseMySQLStatementListener) ExitAlterProcedure(ctx *AlterProcedureContext) {}

// EnterDropProcedure is called when production dropProcedure is entered.
func (s *BaseMySQLStatementListener) EnterDropProcedure(ctx *DropProcedureContext) {}

// ExitDropProcedure is called when production dropProcedure is exited.
func (s *BaseMySQLStatementListener) ExitDropProcedure(ctx *DropProcedureContext) {}

// EnterCreateServer is called when production createServer is entered.
func (s *BaseMySQLStatementListener) EnterCreateServer(ctx *CreateServerContext) {}

// ExitCreateServer is called when production createServer is exited.
func (s *BaseMySQLStatementListener) ExitCreateServer(ctx *CreateServerContext) {}

// EnterAlterServer is called when production alterServer is entered.
func (s *BaseMySQLStatementListener) EnterAlterServer(ctx *AlterServerContext) {}

// ExitAlterServer is called when production alterServer is exited.
func (s *BaseMySQLStatementListener) ExitAlterServer(ctx *AlterServerContext) {}

// EnterDropServer is called when production dropServer is entered.
func (s *BaseMySQLStatementListener) EnterDropServer(ctx *DropServerContext) {}

// ExitDropServer is called when production dropServer is exited.
func (s *BaseMySQLStatementListener) ExitDropServer(ctx *DropServerContext) {}

// EnterCreateView is called when production createView is entered.
func (s *BaseMySQLStatementListener) EnterCreateView(ctx *CreateViewContext) {}

// ExitCreateView is called when production createView is exited.
func (s *BaseMySQLStatementListener) ExitCreateView(ctx *CreateViewContext) {}

// EnterAlterView is called when production alterView is entered.
func (s *BaseMySQLStatementListener) EnterAlterView(ctx *AlterViewContext) {}

// ExitAlterView is called when production alterView is exited.
func (s *BaseMySQLStatementListener) ExitAlterView(ctx *AlterViewContext) {}

// EnterDropView is called when production dropView is entered.
func (s *BaseMySQLStatementListener) EnterDropView(ctx *DropViewContext) {}

// ExitDropView is called when production dropView is exited.
func (s *BaseMySQLStatementListener) ExitDropView(ctx *DropViewContext) {}

// EnterCreateTablespace is called when production createTablespace is entered.
func (s *BaseMySQLStatementListener) EnterCreateTablespace(ctx *CreateTablespaceContext) {}

// ExitCreateTablespace is called when production createTablespace is exited.
func (s *BaseMySQLStatementListener) ExitCreateTablespace(ctx *CreateTablespaceContext) {}

// EnterCreateTablespaceInnodb is called when production createTablespaceInnodb is entered.
func (s *BaseMySQLStatementListener) EnterCreateTablespaceInnodb(ctx *CreateTablespaceInnodbContext) {
}

// ExitCreateTablespaceInnodb is called when production createTablespaceInnodb is exited.
func (s *BaseMySQLStatementListener) ExitCreateTablespaceInnodb(ctx *CreateTablespaceInnodbContext) {}

// EnterCreateTablespaceNdb is called when production createTablespaceNdb is entered.
func (s *BaseMySQLStatementListener) EnterCreateTablespaceNdb(ctx *CreateTablespaceNdbContext) {}

// ExitCreateTablespaceNdb is called when production createTablespaceNdb is exited.
func (s *BaseMySQLStatementListener) ExitCreateTablespaceNdb(ctx *CreateTablespaceNdbContext) {}

// EnterAlterTablespace is called when production alterTablespace is entered.
func (s *BaseMySQLStatementListener) EnterAlterTablespace(ctx *AlterTablespaceContext) {}

// ExitAlterTablespace is called when production alterTablespace is exited.
func (s *BaseMySQLStatementListener) ExitAlterTablespace(ctx *AlterTablespaceContext) {}

// EnterAlterTablespaceNdb is called when production alterTablespaceNdb is entered.
func (s *BaseMySQLStatementListener) EnterAlterTablespaceNdb(ctx *AlterTablespaceNdbContext) {}

// ExitAlterTablespaceNdb is called when production alterTablespaceNdb is exited.
func (s *BaseMySQLStatementListener) ExitAlterTablespaceNdb(ctx *AlterTablespaceNdbContext) {}

// EnterAlterTablespaceInnodb is called when production alterTablespaceInnodb is entered.
func (s *BaseMySQLStatementListener) EnterAlterTablespaceInnodb(ctx *AlterTablespaceInnodbContext) {}

// ExitAlterTablespaceInnodb is called when production alterTablespaceInnodb is exited.
func (s *BaseMySQLStatementListener) ExitAlterTablespaceInnodb(ctx *AlterTablespaceInnodbContext) {}

// EnterDropTablespace is called when production dropTablespace is entered.
func (s *BaseMySQLStatementListener) EnterDropTablespace(ctx *DropTablespaceContext) {}

// ExitDropTablespace is called when production dropTablespace is exited.
func (s *BaseMySQLStatementListener) ExitDropTablespace(ctx *DropTablespaceContext) {}

// EnterCreateLogfileGroup is called when production createLogfileGroup is entered.
func (s *BaseMySQLStatementListener) EnterCreateLogfileGroup(ctx *CreateLogfileGroupContext) {}

// ExitCreateLogfileGroup is called when production createLogfileGroup is exited.
func (s *BaseMySQLStatementListener) ExitCreateLogfileGroup(ctx *CreateLogfileGroupContext) {}

// EnterAlterLogfileGroup is called when production alterLogfileGroup is entered.
func (s *BaseMySQLStatementListener) EnterAlterLogfileGroup(ctx *AlterLogfileGroupContext) {}

// ExitAlterLogfileGroup is called when production alterLogfileGroup is exited.
func (s *BaseMySQLStatementListener) ExitAlterLogfileGroup(ctx *AlterLogfileGroupContext) {}

// EnterDropLogfileGroup is called when production dropLogfileGroup is entered.
func (s *BaseMySQLStatementListener) EnterDropLogfileGroup(ctx *DropLogfileGroupContext) {}

// ExitDropLogfileGroup is called when production dropLogfileGroup is exited.
func (s *BaseMySQLStatementListener) ExitDropLogfileGroup(ctx *DropLogfileGroupContext) {}

// EnterCreateTrigger is called when production createTrigger is entered.
func (s *BaseMySQLStatementListener) EnterCreateTrigger(ctx *CreateTriggerContext) {}

// ExitCreateTrigger is called when production createTrigger is exited.
func (s *BaseMySQLStatementListener) ExitCreateTrigger(ctx *CreateTriggerContext) {}

// EnterDropTrigger is called when production dropTrigger is entered.
func (s *BaseMySQLStatementListener) EnterDropTrigger(ctx *DropTriggerContext) {}

// ExitDropTrigger is called when production dropTrigger is exited.
func (s *BaseMySQLStatementListener) ExitDropTrigger(ctx *DropTriggerContext) {}

// EnterRenameTable is called when production renameTable is entered.
func (s *BaseMySQLStatementListener) EnterRenameTable(ctx *RenameTableContext) {}

// ExitRenameTable is called when production renameTable is exited.
func (s *BaseMySQLStatementListener) ExitRenameTable(ctx *RenameTableContext) {}

// EnterCreateDefinitionClause is called when production createDefinitionClause is entered.
func (s *BaseMySQLStatementListener) EnterCreateDefinitionClause(ctx *CreateDefinitionClauseContext) {
}

// ExitCreateDefinitionClause is called when production createDefinitionClause is exited.
func (s *BaseMySQLStatementListener) ExitCreateDefinitionClause(ctx *CreateDefinitionClauseContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseMySQLStatementListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseMySQLStatementListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterFieldDefinition is called when production fieldDefinition is entered.
func (s *BaseMySQLStatementListener) EnterFieldDefinition(ctx *FieldDefinitionContext) {}

// ExitFieldDefinition is called when production fieldDefinition is exited.
func (s *BaseMySQLStatementListener) ExitFieldDefinition(ctx *FieldDefinitionContext) {}

// EnterColumnAttribute is called when production columnAttribute is entered.
func (s *BaseMySQLStatementListener) EnterColumnAttribute(ctx *ColumnAttributeContext) {}

// ExitColumnAttribute is called when production columnAttribute is exited.
func (s *BaseMySQLStatementListener) ExitColumnAttribute(ctx *ColumnAttributeContext) {}

// EnterCheckConstraint is called when production checkConstraint is entered.
func (s *BaseMySQLStatementListener) EnterCheckConstraint(ctx *CheckConstraintContext) {}

// ExitCheckConstraint is called when production checkConstraint is exited.
func (s *BaseMySQLStatementListener) ExitCheckConstraint(ctx *CheckConstraintContext) {}

// EnterConstraintEnforcement is called when production constraintEnforcement is entered.
func (s *BaseMySQLStatementListener) EnterConstraintEnforcement(ctx *ConstraintEnforcementContext) {}

// ExitConstraintEnforcement is called when production constraintEnforcement is exited.
func (s *BaseMySQLStatementListener) ExitConstraintEnforcement(ctx *ConstraintEnforcementContext) {}

// EnterGeneratedOption is called when production generatedOption is entered.
func (s *BaseMySQLStatementListener) EnterGeneratedOption(ctx *GeneratedOptionContext) {}

// ExitGeneratedOption is called when production generatedOption is exited.
func (s *BaseMySQLStatementListener) ExitGeneratedOption(ctx *GeneratedOptionContext) {}

// EnterReferenceDefinition is called when production referenceDefinition is entered.
func (s *BaseMySQLStatementListener) EnterReferenceDefinition(ctx *ReferenceDefinitionContext) {}

// ExitReferenceDefinition is called when production referenceDefinition is exited.
func (s *BaseMySQLStatementListener) ExitReferenceDefinition(ctx *ReferenceDefinitionContext) {}

// EnterOnUpdateDelete is called when production onUpdateDelete is entered.
func (s *BaseMySQLStatementListener) EnterOnUpdateDelete(ctx *OnUpdateDeleteContext) {}

// ExitOnUpdateDelete is called when production onUpdateDelete is exited.
func (s *BaseMySQLStatementListener) ExitOnUpdateDelete(ctx *OnUpdateDeleteContext) {}

// EnterReferenceOption is called when production referenceOption is entered.
func (s *BaseMySQLStatementListener) EnterReferenceOption(ctx *ReferenceOptionContext) {}

// ExitReferenceOption is called when production referenceOption is exited.
func (s *BaseMySQLStatementListener) ExitReferenceOption(ctx *ReferenceOptionContext) {}

// EnterIndexNameAndType is called when production indexNameAndType is entered.
func (s *BaseMySQLStatementListener) EnterIndexNameAndType(ctx *IndexNameAndTypeContext) {}

// ExitIndexNameAndType is called when production indexNameAndType is exited.
func (s *BaseMySQLStatementListener) ExitIndexNameAndType(ctx *IndexNameAndTypeContext) {}

// EnterIndexType is called when production indexType is entered.
func (s *BaseMySQLStatementListener) EnterIndexType(ctx *IndexTypeContext) {}

// ExitIndexType is called when production indexType is exited.
func (s *BaseMySQLStatementListener) ExitIndexType(ctx *IndexTypeContext) {}

// EnterIndexTypeClause is called when production indexTypeClause is entered.
func (s *BaseMySQLStatementListener) EnterIndexTypeClause(ctx *IndexTypeClauseContext) {}

// ExitIndexTypeClause is called when production indexTypeClause is exited.
func (s *BaseMySQLStatementListener) ExitIndexTypeClause(ctx *IndexTypeClauseContext) {}

// EnterKeyParts is called when production keyParts is entered.
func (s *BaseMySQLStatementListener) EnterKeyParts(ctx *KeyPartsContext) {}

// ExitKeyParts is called when production keyParts is exited.
func (s *BaseMySQLStatementListener) ExitKeyParts(ctx *KeyPartsContext) {}

// EnterKeyPart is called when production keyPart is entered.
func (s *BaseMySQLStatementListener) EnterKeyPart(ctx *KeyPartContext) {}

// ExitKeyPart is called when production keyPart is exited.
func (s *BaseMySQLStatementListener) ExitKeyPart(ctx *KeyPartContext) {}

// EnterKeyPartWithExpression is called when production keyPartWithExpression is entered.
func (s *BaseMySQLStatementListener) EnterKeyPartWithExpression(ctx *KeyPartWithExpressionContext) {}

// ExitKeyPartWithExpression is called when production keyPartWithExpression is exited.
func (s *BaseMySQLStatementListener) ExitKeyPartWithExpression(ctx *KeyPartWithExpressionContext) {}

// EnterKeyListWithExpression is called when production keyListWithExpression is entered.
func (s *BaseMySQLStatementListener) EnterKeyListWithExpression(ctx *KeyListWithExpressionContext) {}

// ExitKeyListWithExpression is called when production keyListWithExpression is exited.
func (s *BaseMySQLStatementListener) ExitKeyListWithExpression(ctx *KeyListWithExpressionContext) {}

// EnterIndexOption is called when production indexOption is entered.
func (s *BaseMySQLStatementListener) EnterIndexOption(ctx *IndexOptionContext) {}

// ExitIndexOption is called when production indexOption is exited.
func (s *BaseMySQLStatementListener) ExitIndexOption(ctx *IndexOptionContext) {}

// EnterCommonIndexOption is called when production commonIndexOption is entered.
func (s *BaseMySQLStatementListener) EnterCommonIndexOption(ctx *CommonIndexOptionContext) {}

// ExitCommonIndexOption is called when production commonIndexOption is exited.
func (s *BaseMySQLStatementListener) ExitCommonIndexOption(ctx *CommonIndexOptionContext) {}

// EnterVisibility is called when production visibility is entered.
func (s *BaseMySQLStatementListener) EnterVisibility(ctx *VisibilityContext) {}

// ExitVisibility is called when production visibility is exited.
func (s *BaseMySQLStatementListener) ExitVisibility(ctx *VisibilityContext) {}

// EnterCreateLikeClause is called when production createLikeClause is entered.
func (s *BaseMySQLStatementListener) EnterCreateLikeClause(ctx *CreateLikeClauseContext) {}

// ExitCreateLikeClause is called when production createLikeClause is exited.
func (s *BaseMySQLStatementListener) ExitCreateLikeClause(ctx *CreateLikeClauseContext) {}

// EnterCreateIndexSpecification is called when production createIndexSpecification is entered.
func (s *BaseMySQLStatementListener) EnterCreateIndexSpecification(ctx *CreateIndexSpecificationContext) {
}

// ExitCreateIndexSpecification is called when production createIndexSpecification is exited.
func (s *BaseMySQLStatementListener) ExitCreateIndexSpecification(ctx *CreateIndexSpecificationContext) {
}

// EnterCreateTableOptions is called when production createTableOptions is entered.
func (s *BaseMySQLStatementListener) EnterCreateTableOptions(ctx *CreateTableOptionsContext) {}

// ExitCreateTableOptions is called when production createTableOptions is exited.
func (s *BaseMySQLStatementListener) ExitCreateTableOptions(ctx *CreateTableOptionsContext) {}

// EnterCreateTableOption is called when production createTableOption is entered.
func (s *BaseMySQLStatementListener) EnterCreateTableOption(ctx *CreateTableOptionContext) {}

// ExitCreateTableOption is called when production createTableOption is exited.
func (s *BaseMySQLStatementListener) ExitCreateTableOption(ctx *CreateTableOptionContext) {}

// EnterCreateSRSStatement is called when production createSRSStatement is entered.
func (s *BaseMySQLStatementListener) EnterCreateSRSStatement(ctx *CreateSRSStatementContext) {}

// ExitCreateSRSStatement is called when production createSRSStatement is exited.
func (s *BaseMySQLStatementListener) ExitCreateSRSStatement(ctx *CreateSRSStatementContext) {}

// EnterDropSRSStatement is called when production dropSRSStatement is entered.
func (s *BaseMySQLStatementListener) EnterDropSRSStatement(ctx *DropSRSStatementContext) {}

// ExitDropSRSStatement is called when production dropSRSStatement is exited.
func (s *BaseMySQLStatementListener) ExitDropSRSStatement(ctx *DropSRSStatementContext) {}

// EnterSrsAttribute is called when production srsAttribute is entered.
func (s *BaseMySQLStatementListener) EnterSrsAttribute(ctx *SrsAttributeContext) {}

// ExitSrsAttribute is called when production srsAttribute is exited.
func (s *BaseMySQLStatementListener) ExitSrsAttribute(ctx *SrsAttributeContext) {}

// EnterPlace is called when production place is entered.
func (s *BaseMySQLStatementListener) EnterPlace(ctx *PlaceContext) {}

// ExitPlace is called when production place is exited.
func (s *BaseMySQLStatementListener) ExitPlace(ctx *PlaceContext) {}

// EnterPartitionDefinitions is called when production partitionDefinitions is entered.
func (s *BaseMySQLStatementListener) EnterPartitionDefinitions(ctx *PartitionDefinitionsContext) {}

// ExitPartitionDefinitions is called when production partitionDefinitions is exited.
func (s *BaseMySQLStatementListener) ExitPartitionDefinitions(ctx *PartitionDefinitionsContext) {}

// EnterPartitionDefinition is called when production partitionDefinition is entered.
func (s *BaseMySQLStatementListener) EnterPartitionDefinition(ctx *PartitionDefinitionContext) {}

// ExitPartitionDefinition is called when production partitionDefinition is exited.
func (s *BaseMySQLStatementListener) ExitPartitionDefinition(ctx *PartitionDefinitionContext) {}

// EnterPartitionLessThanValue is called when production partitionLessThanValue is entered.
func (s *BaseMySQLStatementListener) EnterPartitionLessThanValue(ctx *PartitionLessThanValueContext) {
}

// ExitPartitionLessThanValue is called when production partitionLessThanValue is exited.
func (s *BaseMySQLStatementListener) ExitPartitionLessThanValue(ctx *PartitionLessThanValueContext) {}

// EnterPartitionValueList is called when production partitionValueList is entered.
func (s *BaseMySQLStatementListener) EnterPartitionValueList(ctx *PartitionValueListContext) {}

// ExitPartitionValueList is called when production partitionValueList is exited.
func (s *BaseMySQLStatementListener) ExitPartitionValueList(ctx *PartitionValueListContext) {}

// EnterPartitionDefinitionOption is called when production partitionDefinitionOption is entered.
func (s *BaseMySQLStatementListener) EnterPartitionDefinitionOption(ctx *PartitionDefinitionOptionContext) {
}

// ExitPartitionDefinitionOption is called when production partitionDefinitionOption is exited.
func (s *BaseMySQLStatementListener) ExitPartitionDefinitionOption(ctx *PartitionDefinitionOptionContext) {
}

// EnterSubpartitionDefinition is called when production subpartitionDefinition is entered.
func (s *BaseMySQLStatementListener) EnterSubpartitionDefinition(ctx *SubpartitionDefinitionContext) {
}

// ExitSubpartitionDefinition is called when production subpartitionDefinition is exited.
func (s *BaseMySQLStatementListener) ExitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) {}

// EnterOwnerStatement is called when production ownerStatement is entered.
func (s *BaseMySQLStatementListener) EnterOwnerStatement(ctx *OwnerStatementContext) {}

// ExitOwnerStatement is called when production ownerStatement is exited.
func (s *BaseMySQLStatementListener) ExitOwnerStatement(ctx *OwnerStatementContext) {}

// EnterScheduleExpression is called when production scheduleExpression is entered.
func (s *BaseMySQLStatementListener) EnterScheduleExpression(ctx *ScheduleExpressionContext) {}

// ExitScheduleExpression is called when production scheduleExpression is exited.
func (s *BaseMySQLStatementListener) ExitScheduleExpression(ctx *ScheduleExpressionContext) {}

// EnterTimestampValue is called when production timestampValue is entered.
func (s *BaseMySQLStatementListener) EnterTimestampValue(ctx *TimestampValueContext) {}

// ExitTimestampValue is called when production timestampValue is exited.
func (s *BaseMySQLStatementListener) ExitTimestampValue(ctx *TimestampValueContext) {}

// EnterRoutineBody is called when production routineBody is entered.
func (s *BaseMySQLStatementListener) EnterRoutineBody(ctx *RoutineBodyContext) {}

// ExitRoutineBody is called when production routineBody is exited.
func (s *BaseMySQLStatementListener) ExitRoutineBody(ctx *RoutineBodyContext) {}

// EnterServerOption is called when production serverOption is entered.
func (s *BaseMySQLStatementListener) EnterServerOption(ctx *ServerOptionContext) {}

// ExitServerOption is called when production serverOption is exited.
func (s *BaseMySQLStatementListener) ExitServerOption(ctx *ServerOptionContext) {}

// EnterRoutineOption is called when production routineOption is entered.
func (s *BaseMySQLStatementListener) EnterRoutineOption(ctx *RoutineOptionContext) {}

// ExitRoutineOption is called when production routineOption is exited.
func (s *BaseMySQLStatementListener) ExitRoutineOption(ctx *RoutineOptionContext) {}

// EnterProcedureParameter is called when production procedureParameter is entered.
func (s *BaseMySQLStatementListener) EnterProcedureParameter(ctx *ProcedureParameterContext) {}

// ExitProcedureParameter is called when production procedureParameter is exited.
func (s *BaseMySQLStatementListener) ExitProcedureParameter(ctx *ProcedureParameterContext) {}

// EnterFileSizeLiteral is called when production fileSizeLiteral is entered.
func (s *BaseMySQLStatementListener) EnterFileSizeLiteral(ctx *FileSizeLiteralContext) {}

// ExitFileSizeLiteral is called when production fileSizeLiteral is exited.
func (s *BaseMySQLStatementListener) ExitFileSizeLiteral(ctx *FileSizeLiteralContext) {}

// EnterSimpleStatement is called when production simpleStatement is entered.
func (s *BaseMySQLStatementListener) EnterSimpleStatement(ctx *SimpleStatementContext) {}

// ExitSimpleStatement is called when production simpleStatement is exited.
func (s *BaseMySQLStatementListener) ExitSimpleStatement(ctx *SimpleStatementContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseMySQLStatementListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseMySQLStatementListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterValidStatement is called when production validStatement is entered.
func (s *BaseMySQLStatementListener) EnterValidStatement(ctx *ValidStatementContext) {}

// ExitValidStatement is called when production validStatement is exited.
func (s *BaseMySQLStatementListener) ExitValidStatement(ctx *ValidStatementContext) {}

// EnterBeginStatement is called when production beginStatement is entered.
func (s *BaseMySQLStatementListener) EnterBeginStatement(ctx *BeginStatementContext) {}

// ExitBeginStatement is called when production beginStatement is exited.
func (s *BaseMySQLStatementListener) ExitBeginStatement(ctx *BeginStatementContext) {}

// EnterDeclareStatement is called when production declareStatement is entered.
func (s *BaseMySQLStatementListener) EnterDeclareStatement(ctx *DeclareStatementContext) {}

// ExitDeclareStatement is called when production declareStatement is exited.
func (s *BaseMySQLStatementListener) ExitDeclareStatement(ctx *DeclareStatementContext) {}

// EnterFlowControlStatement is called when production flowControlStatement is entered.
func (s *BaseMySQLStatementListener) EnterFlowControlStatement(ctx *FlowControlStatementContext) {}

// ExitFlowControlStatement is called when production flowControlStatement is exited.
func (s *BaseMySQLStatementListener) ExitFlowControlStatement(ctx *FlowControlStatementContext) {}

// EnterCaseStatement is called when production caseStatement is entered.
func (s *BaseMySQLStatementListener) EnterCaseStatement(ctx *CaseStatementContext) {}

// ExitCaseStatement is called when production caseStatement is exited.
func (s *BaseMySQLStatementListener) ExitCaseStatement(ctx *CaseStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseMySQLStatementListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseMySQLStatementListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterIterateStatement is called when production iterateStatement is entered.
func (s *BaseMySQLStatementListener) EnterIterateStatement(ctx *IterateStatementContext) {}

// ExitIterateStatement is called when production iterateStatement is exited.
func (s *BaseMySQLStatementListener) ExitIterateStatement(ctx *IterateStatementContext) {}

// EnterLeaveStatement is called when production leaveStatement is entered.
func (s *BaseMySQLStatementListener) EnterLeaveStatement(ctx *LeaveStatementContext) {}

// ExitLeaveStatement is called when production leaveStatement is exited.
func (s *BaseMySQLStatementListener) ExitLeaveStatement(ctx *LeaveStatementContext) {}

// EnterLoopStatement is called when production loopStatement is entered.
func (s *BaseMySQLStatementListener) EnterLoopStatement(ctx *LoopStatementContext) {}

// ExitLoopStatement is called when production loopStatement is exited.
func (s *BaseMySQLStatementListener) ExitLoopStatement(ctx *LoopStatementContext) {}

// EnterRepeatStatement is called when production repeatStatement is entered.
func (s *BaseMySQLStatementListener) EnterRepeatStatement(ctx *RepeatStatementContext) {}

// ExitRepeatStatement is called when production repeatStatement is exited.
func (s *BaseMySQLStatementListener) ExitRepeatStatement(ctx *RepeatStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseMySQLStatementListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseMySQLStatementListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseMySQLStatementListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseMySQLStatementListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterCursorStatement is called when production cursorStatement is entered.
func (s *BaseMySQLStatementListener) EnterCursorStatement(ctx *CursorStatementContext) {}

// ExitCursorStatement is called when production cursorStatement is exited.
func (s *BaseMySQLStatementListener) ExitCursorStatement(ctx *CursorStatementContext) {}

// EnterCursorCloseStatement is called when production cursorCloseStatement is entered.
func (s *BaseMySQLStatementListener) EnterCursorCloseStatement(ctx *CursorCloseStatementContext) {}

// ExitCursorCloseStatement is called when production cursorCloseStatement is exited.
func (s *BaseMySQLStatementListener) ExitCursorCloseStatement(ctx *CursorCloseStatementContext) {}

// EnterCursorDeclareStatement is called when production cursorDeclareStatement is entered.
func (s *BaseMySQLStatementListener) EnterCursorDeclareStatement(ctx *CursorDeclareStatementContext) {
}

// ExitCursorDeclareStatement is called when production cursorDeclareStatement is exited.
func (s *BaseMySQLStatementListener) ExitCursorDeclareStatement(ctx *CursorDeclareStatementContext) {}

// EnterCursorFetchStatement is called when production cursorFetchStatement is entered.
func (s *BaseMySQLStatementListener) EnterCursorFetchStatement(ctx *CursorFetchStatementContext) {}

// ExitCursorFetchStatement is called when production cursorFetchStatement is exited.
func (s *BaseMySQLStatementListener) ExitCursorFetchStatement(ctx *CursorFetchStatementContext) {}

// EnterCursorOpenStatement is called when production cursorOpenStatement is entered.
func (s *BaseMySQLStatementListener) EnterCursorOpenStatement(ctx *CursorOpenStatementContext) {}

// ExitCursorOpenStatement is called when production cursorOpenStatement is exited.
func (s *BaseMySQLStatementListener) ExitCursorOpenStatement(ctx *CursorOpenStatementContext) {}

// EnterConditionHandlingStatement is called when production conditionHandlingStatement is entered.
func (s *BaseMySQLStatementListener) EnterConditionHandlingStatement(ctx *ConditionHandlingStatementContext) {
}

// ExitConditionHandlingStatement is called when production conditionHandlingStatement is exited.
func (s *BaseMySQLStatementListener) ExitConditionHandlingStatement(ctx *ConditionHandlingStatementContext) {
}

// EnterDeclareConditionStatement is called when production declareConditionStatement is entered.
func (s *BaseMySQLStatementListener) EnterDeclareConditionStatement(ctx *DeclareConditionStatementContext) {
}

// ExitDeclareConditionStatement is called when production declareConditionStatement is exited.
func (s *BaseMySQLStatementListener) ExitDeclareConditionStatement(ctx *DeclareConditionStatementContext) {
}

// EnterDeclareHandlerStatement is called when production declareHandlerStatement is entered.
func (s *BaseMySQLStatementListener) EnterDeclareHandlerStatement(ctx *DeclareHandlerStatementContext) {
}

// ExitDeclareHandlerStatement is called when production declareHandlerStatement is exited.
func (s *BaseMySQLStatementListener) ExitDeclareHandlerStatement(ctx *DeclareHandlerStatementContext) {
}

// EnterGetDiagnosticsStatement is called when production getDiagnosticsStatement is entered.
func (s *BaseMySQLStatementListener) EnterGetDiagnosticsStatement(ctx *GetDiagnosticsStatementContext) {
}

// ExitGetDiagnosticsStatement is called when production getDiagnosticsStatement is exited.
func (s *BaseMySQLStatementListener) ExitGetDiagnosticsStatement(ctx *GetDiagnosticsStatementContext) {
}

// EnterStatementInformationItem is called when production statementInformationItem is entered.
func (s *BaseMySQLStatementListener) EnterStatementInformationItem(ctx *StatementInformationItemContext) {
}

// ExitStatementInformationItem is called when production statementInformationItem is exited.
func (s *BaseMySQLStatementListener) ExitStatementInformationItem(ctx *StatementInformationItemContext) {
}

// EnterConditionInformationItem is called when production conditionInformationItem is entered.
func (s *BaseMySQLStatementListener) EnterConditionInformationItem(ctx *ConditionInformationItemContext) {
}

// ExitConditionInformationItem is called when production conditionInformationItem is exited.
func (s *BaseMySQLStatementListener) ExitConditionInformationItem(ctx *ConditionInformationItemContext) {
}

// EnterConditionNumber is called when production conditionNumber is entered.
func (s *BaseMySQLStatementListener) EnterConditionNumber(ctx *ConditionNumberContext) {}

// ExitConditionNumber is called when production conditionNumber is exited.
func (s *BaseMySQLStatementListener) ExitConditionNumber(ctx *ConditionNumberContext) {}

// EnterStatementInformationItemName is called when production statementInformationItemName is entered.
func (s *BaseMySQLStatementListener) EnterStatementInformationItemName(ctx *StatementInformationItemNameContext) {
}

// ExitStatementInformationItemName is called when production statementInformationItemName is exited.
func (s *BaseMySQLStatementListener) ExitStatementInformationItemName(ctx *StatementInformationItemNameContext) {
}

// EnterConditionInformationItemName is called when production conditionInformationItemName is entered.
func (s *BaseMySQLStatementListener) EnterConditionInformationItemName(ctx *ConditionInformationItemNameContext) {
}

// ExitConditionInformationItemName is called when production conditionInformationItemName is exited.
func (s *BaseMySQLStatementListener) ExitConditionInformationItemName(ctx *ConditionInformationItemNameContext) {
}

// EnterHandlerAction is called when production handlerAction is entered.
func (s *BaseMySQLStatementListener) EnterHandlerAction(ctx *HandlerActionContext) {}

// ExitHandlerAction is called when production handlerAction is exited.
func (s *BaseMySQLStatementListener) ExitHandlerAction(ctx *HandlerActionContext) {}

// EnterConditionValue is called when production conditionValue is entered.
func (s *BaseMySQLStatementListener) EnterConditionValue(ctx *ConditionValueContext) {}

// ExitConditionValue is called when production conditionValue is exited.
func (s *BaseMySQLStatementListener) ExitConditionValue(ctx *ConditionValueContext) {}

// EnterResignalStatement is called when production resignalStatement is entered.
func (s *BaseMySQLStatementListener) EnterResignalStatement(ctx *ResignalStatementContext) {}

// ExitResignalStatement is called when production resignalStatement is exited.
func (s *BaseMySQLStatementListener) ExitResignalStatement(ctx *ResignalStatementContext) {}

// EnterSignalStatement is called when production signalStatement is entered.
func (s *BaseMySQLStatementListener) EnterSignalStatement(ctx *SignalStatementContext) {}

// ExitSignalStatement is called when production signalStatement is exited.
func (s *BaseMySQLStatementListener) ExitSignalStatement(ctx *SignalStatementContext) {}

// EnterSignalInformationItem is called when production signalInformationItem is entered.
func (s *BaseMySQLStatementListener) EnterSignalInformationItem(ctx *SignalInformationItemContext) {}

// ExitSignalInformationItem is called when production signalInformationItem is exited.
func (s *BaseMySQLStatementListener) ExitSignalInformationItem(ctx *SignalInformationItemContext) {}

// EnterUse is called when production use is entered.
func (s *BaseMySQLStatementListener) EnterUse(ctx *UseContext) {}

// ExitUse is called when production use is exited.
func (s *BaseMySQLStatementListener) ExitUse(ctx *UseContext) {}

// EnterHelp is called when production help is entered.
func (s *BaseMySQLStatementListener) EnterHelp(ctx *HelpContext) {}

// ExitHelp is called when production help is exited.
func (s *BaseMySQLStatementListener) ExitHelp(ctx *HelpContext) {}

// EnterExplain is called when production explain is entered.
func (s *BaseMySQLStatementListener) EnterExplain(ctx *ExplainContext) {}

// ExitExplain is called when production explain is exited.
func (s *BaseMySQLStatementListener) ExitExplain(ctx *ExplainContext) {}

// EnterShowDatabases is called when production showDatabases is entered.
func (s *BaseMySQLStatementListener) EnterShowDatabases(ctx *ShowDatabasesContext) {}

// ExitShowDatabases is called when production showDatabases is exited.
func (s *BaseMySQLStatementListener) ExitShowDatabases(ctx *ShowDatabasesContext) {}

// EnterShowTables is called when production showTables is entered.
func (s *BaseMySQLStatementListener) EnterShowTables(ctx *ShowTablesContext) {}

// ExitShowTables is called when production showTables is exited.
func (s *BaseMySQLStatementListener) ExitShowTables(ctx *ShowTablesContext) {}

// EnterShowTableStatus is called when production showTableStatus is entered.
func (s *BaseMySQLStatementListener) EnterShowTableStatus(ctx *ShowTableStatusContext) {}

// ExitShowTableStatus is called when production showTableStatus is exited.
func (s *BaseMySQLStatementListener) ExitShowTableStatus(ctx *ShowTableStatusContext) {}

// EnterShowColumns is called when production showColumns is entered.
func (s *BaseMySQLStatementListener) EnterShowColumns(ctx *ShowColumnsContext) {}

// ExitShowColumns is called when production showColumns is exited.
func (s *BaseMySQLStatementListener) ExitShowColumns(ctx *ShowColumnsContext) {}

// EnterShowIndex is called when production showIndex is entered.
func (s *BaseMySQLStatementListener) EnterShowIndex(ctx *ShowIndexContext) {}

// ExitShowIndex is called when production showIndex is exited.
func (s *BaseMySQLStatementListener) ExitShowIndex(ctx *ShowIndexContext) {}

// EnterShowCreateTable is called when production showCreateTable is entered.
func (s *BaseMySQLStatementListener) EnterShowCreateTable(ctx *ShowCreateTableContext) {}

// ExitShowCreateTable is called when production showCreateTable is exited.
func (s *BaseMySQLStatementListener) ExitShowCreateTable(ctx *ShowCreateTableContext) {}

// EnterFromSchema is called when production fromSchema is entered.
func (s *BaseMySQLStatementListener) EnterFromSchema(ctx *FromSchemaContext) {}

// ExitFromSchema is called when production fromSchema is exited.
func (s *BaseMySQLStatementListener) ExitFromSchema(ctx *FromSchemaContext) {}

// EnterFromTable is called when production fromTable is entered.
func (s *BaseMySQLStatementListener) EnterFromTable(ctx *FromTableContext) {}

// ExitFromTable is called when production fromTable is exited.
func (s *BaseMySQLStatementListener) ExitFromTable(ctx *FromTableContext) {}

// EnterShowLike is called when production showLike is entered.
func (s *BaseMySQLStatementListener) EnterShowLike(ctx *ShowLikeContext) {}

// ExitShowLike is called when production showLike is exited.
func (s *BaseMySQLStatementListener) ExitShowLike(ctx *ShowLikeContext) {}

// EnterShowWhereClause is called when production showWhereClause is entered.
func (s *BaseMySQLStatementListener) EnterShowWhereClause(ctx *ShowWhereClauseContext) {}

// ExitShowWhereClause is called when production showWhereClause is exited.
func (s *BaseMySQLStatementListener) ExitShowWhereClause(ctx *ShowWhereClauseContext) {}

// EnterShowFilter is called when production showFilter is entered.
func (s *BaseMySQLStatementListener) EnterShowFilter(ctx *ShowFilterContext) {}

// ExitShowFilter is called when production showFilter is exited.
func (s *BaseMySQLStatementListener) ExitShowFilter(ctx *ShowFilterContext) {}

// EnterShowProfileType is called when production showProfileType is entered.
func (s *BaseMySQLStatementListener) EnterShowProfileType(ctx *ShowProfileTypeContext) {}

// ExitShowProfileType is called when production showProfileType is exited.
func (s *BaseMySQLStatementListener) ExitShowProfileType(ctx *ShowProfileTypeContext) {}

// EnterSetVariable is called when production setVariable is entered.
func (s *BaseMySQLStatementListener) EnterSetVariable(ctx *SetVariableContext) {}

// ExitSetVariable is called when production setVariable is exited.
func (s *BaseMySQLStatementListener) ExitSetVariable(ctx *SetVariableContext) {}

// EnterOptionValueList is called when production optionValueList is entered.
func (s *BaseMySQLStatementListener) EnterOptionValueList(ctx *OptionValueListContext) {}

// ExitOptionValueList is called when production optionValueList is exited.
func (s *BaseMySQLStatementListener) ExitOptionValueList(ctx *OptionValueListContext) {}

// EnterOptionValueNoOptionType is called when production optionValueNoOptionType is entered.
func (s *BaseMySQLStatementListener) EnterOptionValueNoOptionType(ctx *OptionValueNoOptionTypeContext) {
}

// ExitOptionValueNoOptionType is called when production optionValueNoOptionType is exited.
func (s *BaseMySQLStatementListener) ExitOptionValueNoOptionType(ctx *OptionValueNoOptionTypeContext) {
}

// EnterOptionValue is called when production optionValue is entered.
func (s *BaseMySQLStatementListener) EnterOptionValue(ctx *OptionValueContext) {}

// ExitOptionValue is called when production optionValue is exited.
func (s *BaseMySQLStatementListener) ExitOptionValue(ctx *OptionValueContext) {}

// EnterShowBinaryLogs is called when production showBinaryLogs is entered.
func (s *BaseMySQLStatementListener) EnterShowBinaryLogs(ctx *ShowBinaryLogsContext) {}

// ExitShowBinaryLogs is called when production showBinaryLogs is exited.
func (s *BaseMySQLStatementListener) ExitShowBinaryLogs(ctx *ShowBinaryLogsContext) {}

// EnterShowBinlogEvents is called when production showBinlogEvents is entered.
func (s *BaseMySQLStatementListener) EnterShowBinlogEvents(ctx *ShowBinlogEventsContext) {}

// ExitShowBinlogEvents is called when production showBinlogEvents is exited.
func (s *BaseMySQLStatementListener) ExitShowBinlogEvents(ctx *ShowBinlogEventsContext) {}

// EnterShowCharacterSet is called when production showCharacterSet is entered.
func (s *BaseMySQLStatementListener) EnterShowCharacterSet(ctx *ShowCharacterSetContext) {}

// ExitShowCharacterSet is called when production showCharacterSet is exited.
func (s *BaseMySQLStatementListener) ExitShowCharacterSet(ctx *ShowCharacterSetContext) {}

// EnterShowCollation is called when production showCollation is entered.
func (s *BaseMySQLStatementListener) EnterShowCollation(ctx *ShowCollationContext) {}

// ExitShowCollation is called when production showCollation is exited.
func (s *BaseMySQLStatementListener) ExitShowCollation(ctx *ShowCollationContext) {}

// EnterShowCreateDatabase is called when production showCreateDatabase is entered.
func (s *BaseMySQLStatementListener) EnterShowCreateDatabase(ctx *ShowCreateDatabaseContext) {}

// ExitShowCreateDatabase is called when production showCreateDatabase is exited.
func (s *BaseMySQLStatementListener) ExitShowCreateDatabase(ctx *ShowCreateDatabaseContext) {}

// EnterShowCreateEvent is called when production showCreateEvent is entered.
func (s *BaseMySQLStatementListener) EnterShowCreateEvent(ctx *ShowCreateEventContext) {}

// ExitShowCreateEvent is called when production showCreateEvent is exited.
func (s *BaseMySQLStatementListener) ExitShowCreateEvent(ctx *ShowCreateEventContext) {}

// EnterShowCreateFunction is called when production showCreateFunction is entered.
func (s *BaseMySQLStatementListener) EnterShowCreateFunction(ctx *ShowCreateFunctionContext) {}

// ExitShowCreateFunction is called when production showCreateFunction is exited.
func (s *BaseMySQLStatementListener) ExitShowCreateFunction(ctx *ShowCreateFunctionContext) {}

// EnterShowCreateProcedure is called when production showCreateProcedure is entered.
func (s *BaseMySQLStatementListener) EnterShowCreateProcedure(ctx *ShowCreateProcedureContext) {}

// ExitShowCreateProcedure is called when production showCreateProcedure is exited.
func (s *BaseMySQLStatementListener) ExitShowCreateProcedure(ctx *ShowCreateProcedureContext) {}

// EnterShowCreateTrigger is called when production showCreateTrigger is entered.
func (s *BaseMySQLStatementListener) EnterShowCreateTrigger(ctx *ShowCreateTriggerContext) {}

// ExitShowCreateTrigger is called when production showCreateTrigger is exited.
func (s *BaseMySQLStatementListener) ExitShowCreateTrigger(ctx *ShowCreateTriggerContext) {}

// EnterShowCreateUser is called when production showCreateUser is entered.
func (s *BaseMySQLStatementListener) EnterShowCreateUser(ctx *ShowCreateUserContext) {}

// ExitShowCreateUser is called when production showCreateUser is exited.
func (s *BaseMySQLStatementListener) ExitShowCreateUser(ctx *ShowCreateUserContext) {}

// EnterShowCreateView is called when production showCreateView is entered.
func (s *BaseMySQLStatementListener) EnterShowCreateView(ctx *ShowCreateViewContext) {}

// ExitShowCreateView is called when production showCreateView is exited.
func (s *BaseMySQLStatementListener) ExitShowCreateView(ctx *ShowCreateViewContext) {}

// EnterShowEngine is called when production showEngine is entered.
func (s *BaseMySQLStatementListener) EnterShowEngine(ctx *ShowEngineContext) {}

// ExitShowEngine is called when production showEngine is exited.
func (s *BaseMySQLStatementListener) ExitShowEngine(ctx *ShowEngineContext) {}

// EnterShowEngines is called when production showEngines is entered.
func (s *BaseMySQLStatementListener) EnterShowEngines(ctx *ShowEnginesContext) {}

// ExitShowEngines is called when production showEngines is exited.
func (s *BaseMySQLStatementListener) ExitShowEngines(ctx *ShowEnginesContext) {}

// EnterShowCharset is called when production showCharset is entered.
func (s *BaseMySQLStatementListener) EnterShowCharset(ctx *ShowCharsetContext) {}

// ExitShowCharset is called when production showCharset is exited.
func (s *BaseMySQLStatementListener) ExitShowCharset(ctx *ShowCharsetContext) {}

// EnterShowErrors is called when production showErrors is entered.
func (s *BaseMySQLStatementListener) EnterShowErrors(ctx *ShowErrorsContext) {}

// ExitShowErrors is called when production showErrors is exited.
func (s *BaseMySQLStatementListener) ExitShowErrors(ctx *ShowErrorsContext) {}

// EnterShowEvents is called when production showEvents is entered.
func (s *BaseMySQLStatementListener) EnterShowEvents(ctx *ShowEventsContext) {}

// ExitShowEvents is called when production showEvents is exited.
func (s *BaseMySQLStatementListener) ExitShowEvents(ctx *ShowEventsContext) {}

// EnterShowFunctionCode is called when production showFunctionCode is entered.
func (s *BaseMySQLStatementListener) EnterShowFunctionCode(ctx *ShowFunctionCodeContext) {}

// ExitShowFunctionCode is called when production showFunctionCode is exited.
func (s *BaseMySQLStatementListener) ExitShowFunctionCode(ctx *ShowFunctionCodeContext) {}

// EnterShowFunctionStatus is called when production showFunctionStatus is entered.
func (s *BaseMySQLStatementListener) EnterShowFunctionStatus(ctx *ShowFunctionStatusContext) {}

// ExitShowFunctionStatus is called when production showFunctionStatus is exited.
func (s *BaseMySQLStatementListener) ExitShowFunctionStatus(ctx *ShowFunctionStatusContext) {}

// EnterShowGrant is called when production showGrant is entered.
func (s *BaseMySQLStatementListener) EnterShowGrant(ctx *ShowGrantContext) {}

// ExitShowGrant is called when production showGrant is exited.
func (s *BaseMySQLStatementListener) ExitShowGrant(ctx *ShowGrantContext) {}

// EnterShowMasterStatus is called when production showMasterStatus is entered.
func (s *BaseMySQLStatementListener) EnterShowMasterStatus(ctx *ShowMasterStatusContext) {}

// ExitShowMasterStatus is called when production showMasterStatus is exited.
func (s *BaseMySQLStatementListener) ExitShowMasterStatus(ctx *ShowMasterStatusContext) {}

// EnterShowOpenTables is called when production showOpenTables is entered.
func (s *BaseMySQLStatementListener) EnterShowOpenTables(ctx *ShowOpenTablesContext) {}

// ExitShowOpenTables is called when production showOpenTables is exited.
func (s *BaseMySQLStatementListener) ExitShowOpenTables(ctx *ShowOpenTablesContext) {}

// EnterShowPlugins is called when production showPlugins is entered.
func (s *BaseMySQLStatementListener) EnterShowPlugins(ctx *ShowPluginsContext) {}

// ExitShowPlugins is called when production showPlugins is exited.
func (s *BaseMySQLStatementListener) ExitShowPlugins(ctx *ShowPluginsContext) {}

// EnterShowPrivileges is called when production showPrivileges is entered.
func (s *BaseMySQLStatementListener) EnterShowPrivileges(ctx *ShowPrivilegesContext) {}

// ExitShowPrivileges is called when production showPrivileges is exited.
func (s *BaseMySQLStatementListener) ExitShowPrivileges(ctx *ShowPrivilegesContext) {}

// EnterShowProcedureCode is called when production showProcedureCode is entered.
func (s *BaseMySQLStatementListener) EnterShowProcedureCode(ctx *ShowProcedureCodeContext) {}

// ExitShowProcedureCode is called when production showProcedureCode is exited.
func (s *BaseMySQLStatementListener) ExitShowProcedureCode(ctx *ShowProcedureCodeContext) {}

// EnterShowProcedureStatus is called when production showProcedureStatus is entered.
func (s *BaseMySQLStatementListener) EnterShowProcedureStatus(ctx *ShowProcedureStatusContext) {}

// ExitShowProcedureStatus is called when production showProcedureStatus is exited.
func (s *BaseMySQLStatementListener) ExitShowProcedureStatus(ctx *ShowProcedureStatusContext) {}

// EnterShowProcesslist is called when production showProcesslist is entered.
func (s *BaseMySQLStatementListener) EnterShowProcesslist(ctx *ShowProcesslistContext) {}

// ExitShowProcesslist is called when production showProcesslist is exited.
func (s *BaseMySQLStatementListener) ExitShowProcesslist(ctx *ShowProcesslistContext) {}

// EnterShowProfile is called when production showProfile is entered.
func (s *BaseMySQLStatementListener) EnterShowProfile(ctx *ShowProfileContext) {}

// ExitShowProfile is called when production showProfile is exited.
func (s *BaseMySQLStatementListener) ExitShowProfile(ctx *ShowProfileContext) {}

// EnterShowProfiles is called when production showProfiles is entered.
func (s *BaseMySQLStatementListener) EnterShowProfiles(ctx *ShowProfilesContext) {}

// ExitShowProfiles is called when production showProfiles is exited.
func (s *BaseMySQLStatementListener) ExitShowProfiles(ctx *ShowProfilesContext) {}

// EnterShowRelaylogEvent is called when production showRelaylogEvent is entered.
func (s *BaseMySQLStatementListener) EnterShowRelaylogEvent(ctx *ShowRelaylogEventContext) {}

// ExitShowRelaylogEvent is called when production showRelaylogEvent is exited.
func (s *BaseMySQLStatementListener) ExitShowRelaylogEvent(ctx *ShowRelaylogEventContext) {}

// EnterShowSlavehost is called when production showSlavehost is entered.
func (s *BaseMySQLStatementListener) EnterShowSlavehost(ctx *ShowSlavehostContext) {}

// ExitShowSlavehost is called when production showSlavehost is exited.
func (s *BaseMySQLStatementListener) ExitShowSlavehost(ctx *ShowSlavehostContext) {}

// EnterShowSlaveStatus is called when production showSlaveStatus is entered.
func (s *BaseMySQLStatementListener) EnterShowSlaveStatus(ctx *ShowSlaveStatusContext) {}

// ExitShowSlaveStatus is called when production showSlaveStatus is exited.
func (s *BaseMySQLStatementListener) ExitShowSlaveStatus(ctx *ShowSlaveStatusContext) {}

// EnterShowStatus is called when production showStatus is entered.
func (s *BaseMySQLStatementListener) EnterShowStatus(ctx *ShowStatusContext) {}

// ExitShowStatus is called when production showStatus is exited.
func (s *BaseMySQLStatementListener) ExitShowStatus(ctx *ShowStatusContext) {}

// EnterShowTriggers is called when production showTriggers is entered.
func (s *BaseMySQLStatementListener) EnterShowTriggers(ctx *ShowTriggersContext) {}

// ExitShowTriggers is called when production showTriggers is exited.
func (s *BaseMySQLStatementListener) ExitShowTriggers(ctx *ShowTriggersContext) {}

// EnterShowVariables is called when production showVariables is entered.
func (s *BaseMySQLStatementListener) EnterShowVariables(ctx *ShowVariablesContext) {}

// ExitShowVariables is called when production showVariables is exited.
func (s *BaseMySQLStatementListener) ExitShowVariables(ctx *ShowVariablesContext) {}

// EnterShowWarnings is called when production showWarnings is entered.
func (s *BaseMySQLStatementListener) EnterShowWarnings(ctx *ShowWarningsContext) {}

// ExitShowWarnings is called when production showWarnings is exited.
func (s *BaseMySQLStatementListener) ExitShowWarnings(ctx *ShowWarningsContext) {}

// EnterShowReplicas is called when production showReplicas is entered.
func (s *BaseMySQLStatementListener) EnterShowReplicas(ctx *ShowReplicasContext) {}

// ExitShowReplicas is called when production showReplicas is exited.
func (s *BaseMySQLStatementListener) ExitShowReplicas(ctx *ShowReplicasContext) {}

// EnterShowReplicaStatus is called when production showReplicaStatus is entered.
func (s *BaseMySQLStatementListener) EnterShowReplicaStatus(ctx *ShowReplicaStatusContext) {}

// ExitShowReplicaStatus is called when production showReplicaStatus is exited.
func (s *BaseMySQLStatementListener) ExitShowReplicaStatus(ctx *ShowReplicaStatusContext) {}

// EnterSetCharacter is called when production setCharacter is entered.
func (s *BaseMySQLStatementListener) EnterSetCharacter(ctx *SetCharacterContext) {}

// ExitSetCharacter is called when production setCharacter is exited.
func (s *BaseMySQLStatementListener) ExitSetCharacter(ctx *SetCharacterContext) {}

// EnterClone is called when production clone is entered.
func (s *BaseMySQLStatementListener) EnterClone(ctx *CloneContext) {}

// ExitClone is called when production clone is exited.
func (s *BaseMySQLStatementListener) ExitClone(ctx *CloneContext) {}

// EnterCloneAction is called when production cloneAction is entered.
func (s *BaseMySQLStatementListener) EnterCloneAction(ctx *CloneActionContext) {}

// ExitCloneAction is called when production cloneAction is exited.
func (s *BaseMySQLStatementListener) ExitCloneAction(ctx *CloneActionContext) {}

// EnterCreateLoadableFunction is called when production createLoadableFunction is entered.
func (s *BaseMySQLStatementListener) EnterCreateLoadableFunction(ctx *CreateLoadableFunctionContext) {
}

// ExitCreateLoadableFunction is called when production createLoadableFunction is exited.
func (s *BaseMySQLStatementListener) ExitCreateLoadableFunction(ctx *CreateLoadableFunctionContext) {}

// EnterInstall is called when production install is entered.
func (s *BaseMySQLStatementListener) EnterInstall(ctx *InstallContext) {}

// ExitInstall is called when production install is exited.
func (s *BaseMySQLStatementListener) ExitInstall(ctx *InstallContext) {}

// EnterUninstall is called when production uninstall is entered.
func (s *BaseMySQLStatementListener) EnterUninstall(ctx *UninstallContext) {}

// ExitUninstall is called when production uninstall is exited.
func (s *BaseMySQLStatementListener) ExitUninstall(ctx *UninstallContext) {}

// EnterInstallComponent is called when production installComponent is entered.
func (s *BaseMySQLStatementListener) EnterInstallComponent(ctx *InstallComponentContext) {}

// ExitInstallComponent is called when production installComponent is exited.
func (s *BaseMySQLStatementListener) ExitInstallComponent(ctx *InstallComponentContext) {}

// EnterInstallPlugin is called when production installPlugin is entered.
func (s *BaseMySQLStatementListener) EnterInstallPlugin(ctx *InstallPluginContext) {}

// ExitInstallPlugin is called when production installPlugin is exited.
func (s *BaseMySQLStatementListener) ExitInstallPlugin(ctx *InstallPluginContext) {}

// EnterUninstallComponent is called when production uninstallComponent is entered.
func (s *BaseMySQLStatementListener) EnterUninstallComponent(ctx *UninstallComponentContext) {}

// ExitUninstallComponent is called when production uninstallComponent is exited.
func (s *BaseMySQLStatementListener) ExitUninstallComponent(ctx *UninstallComponentContext) {}

// EnterUninstallPlugin is called when production uninstallPlugin is entered.
func (s *BaseMySQLStatementListener) EnterUninstallPlugin(ctx *UninstallPluginContext) {}

// ExitUninstallPlugin is called when production uninstallPlugin is exited.
func (s *BaseMySQLStatementListener) ExitUninstallPlugin(ctx *UninstallPluginContext) {}

// EnterAnalyzeTable is called when production analyzeTable is entered.
func (s *BaseMySQLStatementListener) EnterAnalyzeTable(ctx *AnalyzeTableContext) {}

// ExitAnalyzeTable is called when production analyzeTable is exited.
func (s *BaseMySQLStatementListener) ExitAnalyzeTable(ctx *AnalyzeTableContext) {}

// EnterHistogram is called when production histogram is entered.
func (s *BaseMySQLStatementListener) EnterHistogram(ctx *HistogramContext) {}

// ExitHistogram is called when production histogram is exited.
func (s *BaseMySQLStatementListener) ExitHistogram(ctx *HistogramContext) {}

// EnterCheckTable is called when production checkTable is entered.
func (s *BaseMySQLStatementListener) EnterCheckTable(ctx *CheckTableContext) {}

// ExitCheckTable is called when production checkTable is exited.
func (s *BaseMySQLStatementListener) ExitCheckTable(ctx *CheckTableContext) {}

// EnterCheckTableOption is called when production checkTableOption is entered.
func (s *BaseMySQLStatementListener) EnterCheckTableOption(ctx *CheckTableOptionContext) {}

// ExitCheckTableOption is called when production checkTableOption is exited.
func (s *BaseMySQLStatementListener) ExitCheckTableOption(ctx *CheckTableOptionContext) {}

// EnterChecksumTable is called when production checksumTable is entered.
func (s *BaseMySQLStatementListener) EnterChecksumTable(ctx *ChecksumTableContext) {}

// ExitChecksumTable is called when production checksumTable is exited.
func (s *BaseMySQLStatementListener) ExitChecksumTable(ctx *ChecksumTableContext) {}

// EnterOptimizeTable is called when production optimizeTable is entered.
func (s *BaseMySQLStatementListener) EnterOptimizeTable(ctx *OptimizeTableContext) {}

// ExitOptimizeTable is called when production optimizeTable is exited.
func (s *BaseMySQLStatementListener) ExitOptimizeTable(ctx *OptimizeTableContext) {}

// EnterRepairTable is called when production repairTable is entered.
func (s *BaseMySQLStatementListener) EnterRepairTable(ctx *RepairTableContext) {}

// ExitRepairTable is called when production repairTable is exited.
func (s *BaseMySQLStatementListener) ExitRepairTable(ctx *RepairTableContext) {}

// EnterAlterResourceGroup is called when production alterResourceGroup is entered.
func (s *BaseMySQLStatementListener) EnterAlterResourceGroup(ctx *AlterResourceGroupContext) {}

// ExitAlterResourceGroup is called when production alterResourceGroup is exited.
func (s *BaseMySQLStatementListener) ExitAlterResourceGroup(ctx *AlterResourceGroupContext) {}

// EnterVcpuSpec is called when production vcpuSpec is entered.
func (s *BaseMySQLStatementListener) EnterVcpuSpec(ctx *VcpuSpecContext) {}

// ExitVcpuSpec is called when production vcpuSpec is exited.
func (s *BaseMySQLStatementListener) ExitVcpuSpec(ctx *VcpuSpecContext) {}

// EnterCreateResourceGroup is called when production createResourceGroup is entered.
func (s *BaseMySQLStatementListener) EnterCreateResourceGroup(ctx *CreateResourceGroupContext) {}

// ExitCreateResourceGroup is called when production createResourceGroup is exited.
func (s *BaseMySQLStatementListener) ExitCreateResourceGroup(ctx *CreateResourceGroupContext) {}

// EnterDropResourceGroup is called when production dropResourceGroup is entered.
func (s *BaseMySQLStatementListener) EnterDropResourceGroup(ctx *DropResourceGroupContext) {}

// ExitDropResourceGroup is called when production dropResourceGroup is exited.
func (s *BaseMySQLStatementListener) ExitDropResourceGroup(ctx *DropResourceGroupContext) {}

// EnterSetResourceGroup is called when production setResourceGroup is entered.
func (s *BaseMySQLStatementListener) EnterSetResourceGroup(ctx *SetResourceGroupContext) {}

// ExitSetResourceGroup is called when production setResourceGroup is exited.
func (s *BaseMySQLStatementListener) ExitSetResourceGroup(ctx *SetResourceGroupContext) {}

// EnterBinlog is called when production binlog is entered.
func (s *BaseMySQLStatementListener) EnterBinlog(ctx *BinlogContext) {}

// ExitBinlog is called when production binlog is exited.
func (s *BaseMySQLStatementListener) ExitBinlog(ctx *BinlogContext) {}

// EnterCacheIndex is called when production cacheIndex is entered.
func (s *BaseMySQLStatementListener) EnterCacheIndex(ctx *CacheIndexContext) {}

// ExitCacheIndex is called when production cacheIndex is exited.
func (s *BaseMySQLStatementListener) ExitCacheIndex(ctx *CacheIndexContext) {}

// EnterCacheTableIndexList is called when production cacheTableIndexList is entered.
func (s *BaseMySQLStatementListener) EnterCacheTableIndexList(ctx *CacheTableIndexListContext) {}

// ExitCacheTableIndexList is called when production cacheTableIndexList is exited.
func (s *BaseMySQLStatementListener) ExitCacheTableIndexList(ctx *CacheTableIndexListContext) {}

// EnterPartitionList is called when production partitionList is entered.
func (s *BaseMySQLStatementListener) EnterPartitionList(ctx *PartitionListContext) {}

// ExitPartitionList is called when production partitionList is exited.
func (s *BaseMySQLStatementListener) ExitPartitionList(ctx *PartitionListContext) {}

// EnterFlush is called when production flush is entered.
func (s *BaseMySQLStatementListener) EnterFlush(ctx *FlushContext) {}

// ExitFlush is called when production flush is exited.
func (s *BaseMySQLStatementListener) ExitFlush(ctx *FlushContext) {}

// EnterFlushOption is called when production flushOption is entered.
func (s *BaseMySQLStatementListener) EnterFlushOption(ctx *FlushOptionContext) {}

// ExitFlushOption is called when production flushOption is exited.
func (s *BaseMySQLStatementListener) ExitFlushOption(ctx *FlushOptionContext) {}

// EnterTablesOption is called when production tablesOption is entered.
func (s *BaseMySQLStatementListener) EnterTablesOption(ctx *TablesOptionContext) {}

// ExitTablesOption is called when production tablesOption is exited.
func (s *BaseMySQLStatementListener) ExitTablesOption(ctx *TablesOptionContext) {}

// EnterKill is called when production kill is entered.
func (s *BaseMySQLStatementListener) EnterKill(ctx *KillContext) {}

// ExitKill is called when production kill is exited.
func (s *BaseMySQLStatementListener) ExitKill(ctx *KillContext) {}

// EnterLoadIndexInfo is called when production loadIndexInfo is entered.
func (s *BaseMySQLStatementListener) EnterLoadIndexInfo(ctx *LoadIndexInfoContext) {}

// ExitLoadIndexInfo is called when production loadIndexInfo is exited.
func (s *BaseMySQLStatementListener) ExitLoadIndexInfo(ctx *LoadIndexInfoContext) {}

// EnterLoadTableIndexList is called when production loadTableIndexList is entered.
func (s *BaseMySQLStatementListener) EnterLoadTableIndexList(ctx *LoadTableIndexListContext) {}

// ExitLoadTableIndexList is called when production loadTableIndexList is exited.
func (s *BaseMySQLStatementListener) ExitLoadTableIndexList(ctx *LoadTableIndexListContext) {}

// EnterResetStatement is called when production resetStatement is entered.
func (s *BaseMySQLStatementListener) EnterResetStatement(ctx *ResetStatementContext) {}

// ExitResetStatement is called when production resetStatement is exited.
func (s *BaseMySQLStatementListener) ExitResetStatement(ctx *ResetStatementContext) {}

// EnterResetOption is called when production resetOption is entered.
func (s *BaseMySQLStatementListener) EnterResetOption(ctx *ResetOptionContext) {}

// ExitResetOption is called when production resetOption is exited.
func (s *BaseMySQLStatementListener) ExitResetOption(ctx *ResetOptionContext) {}

// EnterResetPersist is called when production resetPersist is entered.
func (s *BaseMySQLStatementListener) EnterResetPersist(ctx *ResetPersistContext) {}

// ExitResetPersist is called when production resetPersist is exited.
func (s *BaseMySQLStatementListener) ExitResetPersist(ctx *ResetPersistContext) {}

// EnterRestart is called when production restart is entered.
func (s *BaseMySQLStatementListener) EnterRestart(ctx *RestartContext) {}

// ExitRestart is called when production restart is exited.
func (s *BaseMySQLStatementListener) ExitRestart(ctx *RestartContext) {}

// EnterShutdown is called when production shutdown is entered.
func (s *BaseMySQLStatementListener) EnterShutdown(ctx *ShutdownContext) {}

// ExitShutdown is called when production shutdown is exited.
func (s *BaseMySQLStatementListener) ExitShutdown(ctx *ShutdownContext) {}

// EnterExplainType is called when production explainType is entered.
func (s *BaseMySQLStatementListener) EnterExplainType(ctx *ExplainTypeContext) {}

// ExitExplainType is called when production explainType is exited.
func (s *BaseMySQLStatementListener) ExitExplainType(ctx *ExplainTypeContext) {}

// EnterExplainableStatement is called when production explainableStatement is entered.
func (s *BaseMySQLStatementListener) EnterExplainableStatement(ctx *ExplainableStatementContext) {}

// ExitExplainableStatement is called when production explainableStatement is exited.
func (s *BaseMySQLStatementListener) ExitExplainableStatement(ctx *ExplainableStatementContext) {}

// EnterFormatName is called when production formatName is entered.
func (s *BaseMySQLStatementListener) EnterFormatName(ctx *FormatNameContext) {}

// ExitFormatName is called when production formatName is exited.
func (s *BaseMySQLStatementListener) ExitFormatName(ctx *FormatNameContext) {}

// EnterShow is called when production show is entered.
func (s *BaseMySQLStatementListener) EnterShow(ctx *ShowContext) {}

// ExitShow is called when production show is exited.
func (s *BaseMySQLStatementListener) ExitShow(ctx *ShowContext) {}

// EnterSetTransaction is called when production setTransaction is entered.
func (s *BaseMySQLStatementListener) EnterSetTransaction(ctx *SetTransactionContext) {}

// ExitSetTransaction is called when production setTransaction is exited.
func (s *BaseMySQLStatementListener) ExitSetTransaction(ctx *SetTransactionContext) {}

// EnterSetAutoCommit is called when production setAutoCommit is entered.
func (s *BaseMySQLStatementListener) EnterSetAutoCommit(ctx *SetAutoCommitContext) {}

// ExitSetAutoCommit is called when production setAutoCommit is exited.
func (s *BaseMySQLStatementListener) ExitSetAutoCommit(ctx *SetAutoCommitContext) {}

// EnterBeginTransaction is called when production beginTransaction is entered.
func (s *BaseMySQLStatementListener) EnterBeginTransaction(ctx *BeginTransactionContext) {}

// ExitBeginTransaction is called when production beginTransaction is exited.
func (s *BaseMySQLStatementListener) ExitBeginTransaction(ctx *BeginTransactionContext) {}

// EnterTransactionCharacteristic is called when production transactionCharacteristic is entered.
func (s *BaseMySQLStatementListener) EnterTransactionCharacteristic(ctx *TransactionCharacteristicContext) {
}

// ExitTransactionCharacteristic is called when production transactionCharacteristic is exited.
func (s *BaseMySQLStatementListener) ExitTransactionCharacteristic(ctx *TransactionCharacteristicContext) {
}

// EnterCommit is called when production commit is entered.
func (s *BaseMySQLStatementListener) EnterCommit(ctx *CommitContext) {}

// ExitCommit is called when production commit is exited.
func (s *BaseMySQLStatementListener) ExitCommit(ctx *CommitContext) {}

// EnterRollback is called when production rollback is entered.
func (s *BaseMySQLStatementListener) EnterRollback(ctx *RollbackContext) {}

// ExitRollback is called when production rollback is exited.
func (s *BaseMySQLStatementListener) ExitRollback(ctx *RollbackContext) {}

// EnterSavepoint is called when production savepoint is entered.
func (s *BaseMySQLStatementListener) EnterSavepoint(ctx *SavepointContext) {}

// ExitSavepoint is called when production savepoint is exited.
func (s *BaseMySQLStatementListener) ExitSavepoint(ctx *SavepointContext) {}

// EnterBegin is called when production begin is entered.
func (s *BaseMySQLStatementListener) EnterBegin(ctx *BeginContext) {}

// ExitBegin is called when production begin is exited.
func (s *BaseMySQLStatementListener) ExitBegin(ctx *BeginContext) {}

// EnterLock is called when production lock is entered.
func (s *BaseMySQLStatementListener) EnterLock(ctx *LockContext) {}

// ExitLock is called when production lock is exited.
func (s *BaseMySQLStatementListener) ExitLock(ctx *LockContext) {}

// EnterUnlock is called when production unlock is entered.
func (s *BaseMySQLStatementListener) EnterUnlock(ctx *UnlockContext) {}

// ExitUnlock is called when production unlock is exited.
func (s *BaseMySQLStatementListener) ExitUnlock(ctx *UnlockContext) {}

// EnterReleaseSavepoint is called when production releaseSavepoint is entered.
func (s *BaseMySQLStatementListener) EnterReleaseSavepoint(ctx *ReleaseSavepointContext) {}

// ExitReleaseSavepoint is called when production releaseSavepoint is exited.
func (s *BaseMySQLStatementListener) ExitReleaseSavepoint(ctx *ReleaseSavepointContext) {}

// EnterXa is called when production xa is entered.
func (s *BaseMySQLStatementListener) EnterXa(ctx *XaContext) {}

// ExitXa is called when production xa is exited.
func (s *BaseMySQLStatementListener) ExitXa(ctx *XaContext) {}

// EnterOptionChain is called when production optionChain is entered.
func (s *BaseMySQLStatementListener) EnterOptionChain(ctx *OptionChainContext) {}

// ExitOptionChain is called when production optionChain is exited.
func (s *BaseMySQLStatementListener) ExitOptionChain(ctx *OptionChainContext) {}

// EnterOptionRelease is called when production optionRelease is entered.
func (s *BaseMySQLStatementListener) EnterOptionRelease(ctx *OptionReleaseContext) {}

// ExitOptionRelease is called when production optionRelease is exited.
func (s *BaseMySQLStatementListener) ExitOptionRelease(ctx *OptionReleaseContext) {}

// EnterTableLock is called when production tableLock is entered.
func (s *BaseMySQLStatementListener) EnterTableLock(ctx *TableLockContext) {}

// ExitTableLock is called when production tableLock is exited.
func (s *BaseMySQLStatementListener) ExitTableLock(ctx *TableLockContext) {}

// EnterLockOption is called when production lockOption is entered.
func (s *BaseMySQLStatementListener) EnterLockOption(ctx *LockOptionContext) {}

// ExitLockOption is called when production lockOption is exited.
func (s *BaseMySQLStatementListener) ExitLockOption(ctx *LockOptionContext) {}

// EnterXid is called when production xid is entered.
func (s *BaseMySQLStatementListener) EnterXid(ctx *XidContext) {}

// ExitXid is called when production xid is exited.
func (s *BaseMySQLStatementListener) ExitXid(ctx *XidContext) {}

// EnterGrantRoleOrPrivilegeTo is called when production grantRoleOrPrivilegeTo is entered.
func (s *BaseMySQLStatementListener) EnterGrantRoleOrPrivilegeTo(ctx *GrantRoleOrPrivilegeToContext) {
}

// ExitGrantRoleOrPrivilegeTo is called when production grantRoleOrPrivilegeTo is exited.
func (s *BaseMySQLStatementListener) ExitGrantRoleOrPrivilegeTo(ctx *GrantRoleOrPrivilegeToContext) {}

// EnterGrantRoleOrPrivilegeOnTo is called when production grantRoleOrPrivilegeOnTo is entered.
func (s *BaseMySQLStatementListener) EnterGrantRoleOrPrivilegeOnTo(ctx *GrantRoleOrPrivilegeOnToContext) {
}

// ExitGrantRoleOrPrivilegeOnTo is called when production grantRoleOrPrivilegeOnTo is exited.
func (s *BaseMySQLStatementListener) ExitGrantRoleOrPrivilegeOnTo(ctx *GrantRoleOrPrivilegeOnToContext) {
}

// EnterGrantProxy is called when production grantProxy is entered.
func (s *BaseMySQLStatementListener) EnterGrantProxy(ctx *GrantProxyContext) {}

// ExitGrantProxy is called when production grantProxy is exited.
func (s *BaseMySQLStatementListener) ExitGrantProxy(ctx *GrantProxyContext) {}

// EnterRevokeFrom is called when production revokeFrom is entered.
func (s *BaseMySQLStatementListener) EnterRevokeFrom(ctx *RevokeFromContext) {}

// ExitRevokeFrom is called when production revokeFrom is exited.
func (s *BaseMySQLStatementListener) ExitRevokeFrom(ctx *RevokeFromContext) {}

// EnterRevokeOnFrom is called when production revokeOnFrom is entered.
func (s *BaseMySQLStatementListener) EnterRevokeOnFrom(ctx *RevokeOnFromContext) {}

// ExitRevokeOnFrom is called when production revokeOnFrom is exited.
func (s *BaseMySQLStatementListener) ExitRevokeOnFrom(ctx *RevokeOnFromContext) {}

// EnterUserList is called when production userList is entered.
func (s *BaseMySQLStatementListener) EnterUserList(ctx *UserListContext) {}

// ExitUserList is called when production userList is exited.
func (s *BaseMySQLStatementListener) ExitUserList(ctx *UserListContext) {}

// EnterRoleOrPrivileges is called when production roleOrPrivileges is entered.
func (s *BaseMySQLStatementListener) EnterRoleOrPrivileges(ctx *RoleOrPrivilegesContext) {}

// ExitRoleOrPrivileges is called when production roleOrPrivileges is exited.
func (s *BaseMySQLStatementListener) ExitRoleOrPrivileges(ctx *RoleOrPrivilegesContext) {}

// EnterRoleOrDynamicPrivilege is called when production roleOrDynamicPrivilege is entered.
func (s *BaseMySQLStatementListener) EnterRoleOrDynamicPrivilege(ctx *RoleOrDynamicPrivilegeContext) {
}

// ExitRoleOrDynamicPrivilege is called when production roleOrDynamicPrivilege is exited.
func (s *BaseMySQLStatementListener) ExitRoleOrDynamicPrivilege(ctx *RoleOrDynamicPrivilegeContext) {}

// EnterRoleAtHost is called when production roleAtHost is entered.
func (s *BaseMySQLStatementListener) EnterRoleAtHost(ctx *RoleAtHostContext) {}

// ExitRoleAtHost is called when production roleAtHost is exited.
func (s *BaseMySQLStatementListener) ExitRoleAtHost(ctx *RoleAtHostContext) {}

// EnterStaticPrivilegeSelect is called when production staticPrivilegeSelect is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeSelect(ctx *StaticPrivilegeSelectContext) {}

// ExitStaticPrivilegeSelect is called when production staticPrivilegeSelect is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeSelect(ctx *StaticPrivilegeSelectContext) {}

// EnterStaticPrivilegeInsert is called when production staticPrivilegeInsert is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeInsert(ctx *StaticPrivilegeInsertContext) {}

// ExitStaticPrivilegeInsert is called when production staticPrivilegeInsert is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeInsert(ctx *StaticPrivilegeInsertContext) {}

// EnterStaticPrivilegeUpdate is called when production staticPrivilegeUpdate is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeUpdate(ctx *StaticPrivilegeUpdateContext) {}

// ExitStaticPrivilegeUpdate is called when production staticPrivilegeUpdate is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeUpdate(ctx *StaticPrivilegeUpdateContext) {}

// EnterStaticPrivilegeReferences is called when production staticPrivilegeReferences is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeReferences(ctx *StaticPrivilegeReferencesContext) {
}

// ExitStaticPrivilegeReferences is called when production staticPrivilegeReferences is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeReferences(ctx *StaticPrivilegeReferencesContext) {
}

// EnterStaticPrivilegeDelete is called when production staticPrivilegeDelete is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeDelete(ctx *StaticPrivilegeDeleteContext) {}

// ExitStaticPrivilegeDelete is called when production staticPrivilegeDelete is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeDelete(ctx *StaticPrivilegeDeleteContext) {}

// EnterStaticPrivilegeUsage is called when production staticPrivilegeUsage is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeUsage(ctx *StaticPrivilegeUsageContext) {}

// ExitStaticPrivilegeUsage is called when production staticPrivilegeUsage is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeUsage(ctx *StaticPrivilegeUsageContext) {}

// EnterStaticPrivilegeIndex is called when production staticPrivilegeIndex is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeIndex(ctx *StaticPrivilegeIndexContext) {}

// ExitStaticPrivilegeIndex is called when production staticPrivilegeIndex is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeIndex(ctx *StaticPrivilegeIndexContext) {}

// EnterStaticPrivilegeAlter is called when production staticPrivilegeAlter is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeAlter(ctx *StaticPrivilegeAlterContext) {}

// ExitStaticPrivilegeAlter is called when production staticPrivilegeAlter is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeAlter(ctx *StaticPrivilegeAlterContext) {}

// EnterStaticPrivilegeCreate is called when production staticPrivilegeCreate is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeCreate(ctx *StaticPrivilegeCreateContext) {}

// ExitStaticPrivilegeCreate is called when production staticPrivilegeCreate is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeCreate(ctx *StaticPrivilegeCreateContext) {}

// EnterStaticPrivilegeDrop is called when production staticPrivilegeDrop is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeDrop(ctx *StaticPrivilegeDropContext) {}

// ExitStaticPrivilegeDrop is called when production staticPrivilegeDrop is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeDrop(ctx *StaticPrivilegeDropContext) {}

// EnterStaticPrivilegeExecute is called when production staticPrivilegeExecute is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeExecute(ctx *StaticPrivilegeExecuteContext) {
}

// ExitStaticPrivilegeExecute is called when production staticPrivilegeExecute is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeExecute(ctx *StaticPrivilegeExecuteContext) {}

// EnterStaticPrivilegeReload is called when production staticPrivilegeReload is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeReload(ctx *StaticPrivilegeReloadContext) {}

// ExitStaticPrivilegeReload is called when production staticPrivilegeReload is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeReload(ctx *StaticPrivilegeReloadContext) {}

// EnterStaticPrivilegeShutdown is called when production staticPrivilegeShutdown is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeShutdown(ctx *StaticPrivilegeShutdownContext) {
}

// ExitStaticPrivilegeShutdown is called when production staticPrivilegeShutdown is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeShutdown(ctx *StaticPrivilegeShutdownContext) {
}

// EnterStaticPrivilegeProcess is called when production staticPrivilegeProcess is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeProcess(ctx *StaticPrivilegeProcessContext) {
}

// ExitStaticPrivilegeProcess is called when production staticPrivilegeProcess is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeProcess(ctx *StaticPrivilegeProcessContext) {}

// EnterStaticPrivilegeFile is called when production staticPrivilegeFile is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeFile(ctx *StaticPrivilegeFileContext) {}

// ExitStaticPrivilegeFile is called when production staticPrivilegeFile is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeFile(ctx *StaticPrivilegeFileContext) {}

// EnterStaticPrivilegeGrant is called when production staticPrivilegeGrant is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeGrant(ctx *StaticPrivilegeGrantContext) {}

// ExitStaticPrivilegeGrant is called when production staticPrivilegeGrant is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeGrant(ctx *StaticPrivilegeGrantContext) {}

// EnterStaticPrivilegeShowDatabases is called when production staticPrivilegeShowDatabases is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeShowDatabases(ctx *StaticPrivilegeShowDatabasesContext) {
}

// ExitStaticPrivilegeShowDatabases is called when production staticPrivilegeShowDatabases is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeShowDatabases(ctx *StaticPrivilegeShowDatabasesContext) {
}

// EnterStaticPrivilegeSuper is called when production staticPrivilegeSuper is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeSuper(ctx *StaticPrivilegeSuperContext) {}

// ExitStaticPrivilegeSuper is called when production staticPrivilegeSuper is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeSuper(ctx *StaticPrivilegeSuperContext) {}

// EnterStaticPrivilegeCreateTemporaryTables is called when production staticPrivilegeCreateTemporaryTables is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeCreateTemporaryTables(ctx *StaticPrivilegeCreateTemporaryTablesContext) {
}

// ExitStaticPrivilegeCreateTemporaryTables is called when production staticPrivilegeCreateTemporaryTables is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeCreateTemporaryTables(ctx *StaticPrivilegeCreateTemporaryTablesContext) {
}

// EnterStaticPrivilegeLockTables is called when production staticPrivilegeLockTables is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeLockTables(ctx *StaticPrivilegeLockTablesContext) {
}

// ExitStaticPrivilegeLockTables is called when production staticPrivilegeLockTables is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeLockTables(ctx *StaticPrivilegeLockTablesContext) {
}

// EnterStaticPrivilegeReplicationSlave is called when production staticPrivilegeReplicationSlave is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeReplicationSlave(ctx *StaticPrivilegeReplicationSlaveContext) {
}

// ExitStaticPrivilegeReplicationSlave is called when production staticPrivilegeReplicationSlave is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeReplicationSlave(ctx *StaticPrivilegeReplicationSlaveContext) {
}

// EnterStaticPrivilegeReplicationClient is called when production staticPrivilegeReplicationClient is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeReplicationClient(ctx *StaticPrivilegeReplicationClientContext) {
}

// ExitStaticPrivilegeReplicationClient is called when production staticPrivilegeReplicationClient is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeReplicationClient(ctx *StaticPrivilegeReplicationClientContext) {
}

// EnterStaticPrivilegeCreateView is called when production staticPrivilegeCreateView is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeCreateView(ctx *StaticPrivilegeCreateViewContext) {
}

// ExitStaticPrivilegeCreateView is called when production staticPrivilegeCreateView is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeCreateView(ctx *StaticPrivilegeCreateViewContext) {
}

// EnterStaticPrivilegeShowView is called when production staticPrivilegeShowView is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeShowView(ctx *StaticPrivilegeShowViewContext) {
}

// ExitStaticPrivilegeShowView is called when production staticPrivilegeShowView is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeShowView(ctx *StaticPrivilegeShowViewContext) {
}

// EnterStaticPrivilegeCreateRoutine is called when production staticPrivilegeCreateRoutine is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeCreateRoutine(ctx *StaticPrivilegeCreateRoutineContext) {
}

// ExitStaticPrivilegeCreateRoutine is called when production staticPrivilegeCreateRoutine is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeCreateRoutine(ctx *StaticPrivilegeCreateRoutineContext) {
}

// EnterStaticPrivilegeAlterRoutine is called when production staticPrivilegeAlterRoutine is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeAlterRoutine(ctx *StaticPrivilegeAlterRoutineContext) {
}

// ExitStaticPrivilegeAlterRoutine is called when production staticPrivilegeAlterRoutine is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeAlterRoutine(ctx *StaticPrivilegeAlterRoutineContext) {
}

// EnterStaticPrivilegeCreateUser is called when production staticPrivilegeCreateUser is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeCreateUser(ctx *StaticPrivilegeCreateUserContext) {
}

// ExitStaticPrivilegeCreateUser is called when production staticPrivilegeCreateUser is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeCreateUser(ctx *StaticPrivilegeCreateUserContext) {
}

// EnterStaticPrivilegeEvent is called when production staticPrivilegeEvent is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeEvent(ctx *StaticPrivilegeEventContext) {}

// ExitStaticPrivilegeEvent is called when production staticPrivilegeEvent is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeEvent(ctx *StaticPrivilegeEventContext) {}

// EnterStaticPrivilegeTrigger is called when production staticPrivilegeTrigger is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeTrigger(ctx *StaticPrivilegeTriggerContext) {
}

// ExitStaticPrivilegeTrigger is called when production staticPrivilegeTrigger is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeTrigger(ctx *StaticPrivilegeTriggerContext) {}

// EnterStaticPrivilegeCreateTablespace is called when production staticPrivilegeCreateTablespace is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeCreateTablespace(ctx *StaticPrivilegeCreateTablespaceContext) {
}

// ExitStaticPrivilegeCreateTablespace is called when production staticPrivilegeCreateTablespace is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeCreateTablespace(ctx *StaticPrivilegeCreateTablespaceContext) {
}

// EnterStaticPrivilegeCreateRole is called when production staticPrivilegeCreateRole is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeCreateRole(ctx *StaticPrivilegeCreateRoleContext) {
}

// ExitStaticPrivilegeCreateRole is called when production staticPrivilegeCreateRole is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeCreateRole(ctx *StaticPrivilegeCreateRoleContext) {
}

// EnterStaticPrivilegeDropRole is called when production staticPrivilegeDropRole is entered.
func (s *BaseMySQLStatementListener) EnterStaticPrivilegeDropRole(ctx *StaticPrivilegeDropRoleContext) {
}

// ExitStaticPrivilegeDropRole is called when production staticPrivilegeDropRole is exited.
func (s *BaseMySQLStatementListener) ExitStaticPrivilegeDropRole(ctx *StaticPrivilegeDropRoleContext) {
}

// EnterAclType is called when production aclType is entered.
func (s *BaseMySQLStatementListener) EnterAclType(ctx *AclTypeContext) {}

// ExitAclType is called when production aclType is exited.
func (s *BaseMySQLStatementListener) ExitAclType(ctx *AclTypeContext) {}

// EnterGrantLevelGlobal is called when production grantLevelGlobal is entered.
func (s *BaseMySQLStatementListener) EnterGrantLevelGlobal(ctx *GrantLevelGlobalContext) {}

// ExitGrantLevelGlobal is called when production grantLevelGlobal is exited.
func (s *BaseMySQLStatementListener) ExitGrantLevelGlobal(ctx *GrantLevelGlobalContext) {}

// EnterGrantLevelSchemaGlobal is called when production grantLevelSchemaGlobal is entered.
func (s *BaseMySQLStatementListener) EnterGrantLevelSchemaGlobal(ctx *GrantLevelSchemaGlobalContext) {
}

// ExitGrantLevelSchemaGlobal is called when production grantLevelSchemaGlobal is exited.
func (s *BaseMySQLStatementListener) ExitGrantLevelSchemaGlobal(ctx *GrantLevelSchemaGlobalContext) {}

// EnterGrantLevelTable is called when production grantLevelTable is entered.
func (s *BaseMySQLStatementListener) EnterGrantLevelTable(ctx *GrantLevelTableContext) {}

// ExitGrantLevelTable is called when production grantLevelTable is exited.
func (s *BaseMySQLStatementListener) ExitGrantLevelTable(ctx *GrantLevelTableContext) {}

// EnterCreateUser is called when production createUser is entered.
func (s *BaseMySQLStatementListener) EnterCreateUser(ctx *CreateUserContext) {}

// ExitCreateUser is called when production createUser is exited.
func (s *BaseMySQLStatementListener) ExitCreateUser(ctx *CreateUserContext) {}

// EnterCreateUserEntryNoOption is called when production createUserEntryNoOption is entered.
func (s *BaseMySQLStatementListener) EnterCreateUserEntryNoOption(ctx *CreateUserEntryNoOptionContext) {
}

// ExitCreateUserEntryNoOption is called when production createUserEntryNoOption is exited.
func (s *BaseMySQLStatementListener) ExitCreateUserEntryNoOption(ctx *CreateUserEntryNoOptionContext) {
}

// EnterCreateUserEntryIdentifiedBy is called when production createUserEntryIdentifiedBy is entered.
func (s *BaseMySQLStatementListener) EnterCreateUserEntryIdentifiedBy(ctx *CreateUserEntryIdentifiedByContext) {
}

// ExitCreateUserEntryIdentifiedBy is called when production createUserEntryIdentifiedBy is exited.
func (s *BaseMySQLStatementListener) ExitCreateUserEntryIdentifiedBy(ctx *CreateUserEntryIdentifiedByContext) {
}

// EnterCreateUserEntryIdentifiedWith is called when production createUserEntryIdentifiedWith is entered.
func (s *BaseMySQLStatementListener) EnterCreateUserEntryIdentifiedWith(ctx *CreateUserEntryIdentifiedWithContext) {
}

// ExitCreateUserEntryIdentifiedWith is called when production createUserEntryIdentifiedWith is exited.
func (s *BaseMySQLStatementListener) ExitCreateUserEntryIdentifiedWith(ctx *CreateUserEntryIdentifiedWithContext) {
}

// EnterCreateUserList is called when production createUserList is entered.
func (s *BaseMySQLStatementListener) EnterCreateUserList(ctx *CreateUserListContext) {}

// ExitCreateUserList is called when production createUserList is exited.
func (s *BaseMySQLStatementListener) ExitCreateUserList(ctx *CreateUserListContext) {}

// EnterDefaultRoleClause is called when production defaultRoleClause is entered.
func (s *BaseMySQLStatementListener) EnterDefaultRoleClause(ctx *DefaultRoleClauseContext) {}

// ExitDefaultRoleClause is called when production defaultRoleClause is exited.
func (s *BaseMySQLStatementListener) ExitDefaultRoleClause(ctx *DefaultRoleClauseContext) {}

// EnterRequireClause is called when production requireClause is entered.
func (s *BaseMySQLStatementListener) EnterRequireClause(ctx *RequireClauseContext) {}

// ExitRequireClause is called when production requireClause is exited.
func (s *BaseMySQLStatementListener) ExitRequireClause(ctx *RequireClauseContext) {}

// EnterConnectOptions is called when production connectOptions is entered.
func (s *BaseMySQLStatementListener) EnterConnectOptions(ctx *ConnectOptionsContext) {}

// ExitConnectOptions is called when production connectOptions is exited.
func (s *BaseMySQLStatementListener) ExitConnectOptions(ctx *ConnectOptionsContext) {}

// EnterAccountLockPasswordExpireOptions is called when production accountLockPasswordExpireOptions is entered.
func (s *BaseMySQLStatementListener) EnterAccountLockPasswordExpireOptions(ctx *AccountLockPasswordExpireOptionsContext) {
}

// ExitAccountLockPasswordExpireOptions is called when production accountLockPasswordExpireOptions is exited.
func (s *BaseMySQLStatementListener) ExitAccountLockPasswordExpireOptions(ctx *AccountLockPasswordExpireOptionsContext) {
}

// EnterAccountLockPasswordExpireOption is called when production accountLockPasswordExpireOption is entered.
func (s *BaseMySQLStatementListener) EnterAccountLockPasswordExpireOption(ctx *AccountLockPasswordExpireOptionContext) {
}

// ExitAccountLockPasswordExpireOption is called when production accountLockPasswordExpireOption is exited.
func (s *BaseMySQLStatementListener) ExitAccountLockPasswordExpireOption(ctx *AccountLockPasswordExpireOptionContext) {
}

// EnterAlterUser is called when production alterUser is entered.
func (s *BaseMySQLStatementListener) EnterAlterUser(ctx *AlterUserContext) {}

// ExitAlterUser is called when production alterUser is exited.
func (s *BaseMySQLStatementListener) ExitAlterUser(ctx *AlterUserContext) {}

// EnterAlterUserEntry is called when production alterUserEntry is entered.
func (s *BaseMySQLStatementListener) EnterAlterUserEntry(ctx *AlterUserEntryContext) {}

// ExitAlterUserEntry is called when production alterUserEntry is exited.
func (s *BaseMySQLStatementListener) ExitAlterUserEntry(ctx *AlterUserEntryContext) {}

// EnterAlterUserList is called when production alterUserList is entered.
func (s *BaseMySQLStatementListener) EnterAlterUserList(ctx *AlterUserListContext) {}

// ExitAlterUserList is called when production alterUserList is exited.
func (s *BaseMySQLStatementListener) ExitAlterUserList(ctx *AlterUserListContext) {}

// EnterDropUser is called when production dropUser is entered.
func (s *BaseMySQLStatementListener) EnterDropUser(ctx *DropUserContext) {}

// ExitDropUser is called when production dropUser is exited.
func (s *BaseMySQLStatementListener) ExitDropUser(ctx *DropUserContext) {}

// EnterCreateRole is called when production createRole is entered.
func (s *BaseMySQLStatementListener) EnterCreateRole(ctx *CreateRoleContext) {}

// ExitCreateRole is called when production createRole is exited.
func (s *BaseMySQLStatementListener) ExitCreateRole(ctx *CreateRoleContext) {}

// EnterDropRole is called when production dropRole is entered.
func (s *BaseMySQLStatementListener) EnterDropRole(ctx *DropRoleContext) {}

// ExitDropRole is called when production dropRole is exited.
func (s *BaseMySQLStatementListener) ExitDropRole(ctx *DropRoleContext) {}

// EnterRenameUser is called when production renameUser is entered.
func (s *BaseMySQLStatementListener) EnterRenameUser(ctx *RenameUserContext) {}

// ExitRenameUser is called when production renameUser is exited.
func (s *BaseMySQLStatementListener) ExitRenameUser(ctx *RenameUserContext) {}

// EnterSetDefaultRole is called when production setDefaultRole is entered.
func (s *BaseMySQLStatementListener) EnterSetDefaultRole(ctx *SetDefaultRoleContext) {}

// ExitSetDefaultRole is called when production setDefaultRole is exited.
func (s *BaseMySQLStatementListener) ExitSetDefaultRole(ctx *SetDefaultRoleContext) {}

// EnterSetRole is called when production setRole is entered.
func (s *BaseMySQLStatementListener) EnterSetRole(ctx *SetRoleContext) {}

// ExitSetRole is called when production setRole is exited.
func (s *BaseMySQLStatementListener) ExitSetRole(ctx *SetRoleContext) {}

// EnterSetPassword is called when production setPassword is entered.
func (s *BaseMySQLStatementListener) EnterSetPassword(ctx *SetPasswordContext) {}

// ExitSetPassword is called when production setPassword is exited.
func (s *BaseMySQLStatementListener) ExitSetPassword(ctx *SetPasswordContext) {}

// EnterAuthOption is called when production authOption is entered.
func (s *BaseMySQLStatementListener) EnterAuthOption(ctx *AuthOptionContext) {}

// ExitAuthOption is called when production authOption is exited.
func (s *BaseMySQLStatementListener) ExitAuthOption(ctx *AuthOptionContext) {}

// EnterWithGrantOption is called when production withGrantOption is entered.
func (s *BaseMySQLStatementListener) EnterWithGrantOption(ctx *WithGrantOptionContext) {}

// ExitWithGrantOption is called when production withGrantOption is exited.
func (s *BaseMySQLStatementListener) ExitWithGrantOption(ctx *WithGrantOptionContext) {}

// EnterUserOrRoles is called when production userOrRoles is entered.
func (s *BaseMySQLStatementListener) EnterUserOrRoles(ctx *UserOrRolesContext) {}

// ExitUserOrRoles is called when production userOrRoles is exited.
func (s *BaseMySQLStatementListener) ExitUserOrRoles(ctx *UserOrRolesContext) {}

// EnterRoles is called when production roles is entered.
func (s *BaseMySQLStatementListener) EnterRoles(ctx *RolesContext) {}

// ExitRoles is called when production roles is exited.
func (s *BaseMySQLStatementListener) ExitRoles(ctx *RolesContext) {}

// EnterGrantAs is called when production grantAs is entered.
func (s *BaseMySQLStatementListener) EnterGrantAs(ctx *GrantAsContext) {}

// ExitGrantAs is called when production grantAs is exited.
func (s *BaseMySQLStatementListener) ExitGrantAs(ctx *GrantAsContext) {}

// EnterWithRoles is called when production withRoles is entered.
func (s *BaseMySQLStatementListener) EnterWithRoles(ctx *WithRolesContext) {}

// ExitWithRoles is called when production withRoles is exited.
func (s *BaseMySQLStatementListener) ExitWithRoles(ctx *WithRolesContext) {}

// EnterUserAuthOption is called when production userAuthOption is entered.
func (s *BaseMySQLStatementListener) EnterUserAuthOption(ctx *UserAuthOptionContext) {}

// ExitUserAuthOption is called when production userAuthOption is exited.
func (s *BaseMySQLStatementListener) ExitUserAuthOption(ctx *UserAuthOptionContext) {}

// EnterIdentifiedBy is called when production identifiedBy is entered.
func (s *BaseMySQLStatementListener) EnterIdentifiedBy(ctx *IdentifiedByContext) {}

// ExitIdentifiedBy is called when production identifiedBy is exited.
func (s *BaseMySQLStatementListener) ExitIdentifiedBy(ctx *IdentifiedByContext) {}

// EnterIdentifiedWith is called when production identifiedWith is entered.
func (s *BaseMySQLStatementListener) EnterIdentifiedWith(ctx *IdentifiedWithContext) {}

// ExitIdentifiedWith is called when production identifiedWith is exited.
func (s *BaseMySQLStatementListener) ExitIdentifiedWith(ctx *IdentifiedWithContext) {}

// EnterConnectOption is called when production connectOption is entered.
func (s *BaseMySQLStatementListener) EnterConnectOption(ctx *ConnectOptionContext) {}

// ExitConnectOption is called when production connectOption is exited.
func (s *BaseMySQLStatementListener) ExitConnectOption(ctx *ConnectOptionContext) {}

// EnterTlsOption is called when production tlsOption is entered.
func (s *BaseMySQLStatementListener) EnterTlsOption(ctx *TlsOptionContext) {}

// ExitTlsOption is called when production tlsOption is exited.
func (s *BaseMySQLStatementListener) ExitTlsOption(ctx *TlsOptionContext) {}

// EnterUserFuncAuthOption is called when production userFuncAuthOption is entered.
func (s *BaseMySQLStatementListener) EnterUserFuncAuthOption(ctx *UserFuncAuthOptionContext) {}

// ExitUserFuncAuthOption is called when production userFuncAuthOption is exited.
func (s *BaseMySQLStatementListener) ExitUserFuncAuthOption(ctx *UserFuncAuthOptionContext) {}

// EnterChange is called when production change is entered.
func (s *BaseMySQLStatementListener) EnterChange(ctx *ChangeContext) {}

// ExitChange is called when production change is exited.
func (s *BaseMySQLStatementListener) ExitChange(ctx *ChangeContext) {}

// EnterChangeMasterTo is called when production changeMasterTo is entered.
func (s *BaseMySQLStatementListener) EnterChangeMasterTo(ctx *ChangeMasterToContext) {}

// ExitChangeMasterTo is called when production changeMasterTo is exited.
func (s *BaseMySQLStatementListener) ExitChangeMasterTo(ctx *ChangeMasterToContext) {}

// EnterChangeReplicationFilter is called when production changeReplicationFilter is entered.
func (s *BaseMySQLStatementListener) EnterChangeReplicationFilter(ctx *ChangeReplicationFilterContext) {
}

// ExitChangeReplicationFilter is called when production changeReplicationFilter is exited.
func (s *BaseMySQLStatementListener) ExitChangeReplicationFilter(ctx *ChangeReplicationFilterContext) {
}

// EnterStartSlave is called when production startSlave is entered.
func (s *BaseMySQLStatementListener) EnterStartSlave(ctx *StartSlaveContext) {}

// ExitStartSlave is called when production startSlave is exited.
func (s *BaseMySQLStatementListener) ExitStartSlave(ctx *StartSlaveContext) {}

// EnterStopSlave is called when production stopSlave is entered.
func (s *BaseMySQLStatementListener) EnterStopSlave(ctx *StopSlaveContext) {}

// ExitStopSlave is called when production stopSlave is exited.
func (s *BaseMySQLStatementListener) ExitStopSlave(ctx *StopSlaveContext) {}

// EnterGroupReplication is called when production groupReplication is entered.
func (s *BaseMySQLStatementListener) EnterGroupReplication(ctx *GroupReplicationContext) {}

// ExitGroupReplication is called when production groupReplication is exited.
func (s *BaseMySQLStatementListener) ExitGroupReplication(ctx *GroupReplicationContext) {}

// EnterStartGroupReplication is called when production startGroupReplication is entered.
func (s *BaseMySQLStatementListener) EnterStartGroupReplication(ctx *StartGroupReplicationContext) {}

// ExitStartGroupReplication is called when production startGroupReplication is exited.
func (s *BaseMySQLStatementListener) ExitStartGroupReplication(ctx *StartGroupReplicationContext) {}

// EnterStopGroupReplication is called when production stopGroupReplication is entered.
func (s *BaseMySQLStatementListener) EnterStopGroupReplication(ctx *StopGroupReplicationContext) {}

// ExitStopGroupReplication is called when production stopGroupReplication is exited.
func (s *BaseMySQLStatementListener) ExitStopGroupReplication(ctx *StopGroupReplicationContext) {}

// EnterPurgeBinaryLog is called when production purgeBinaryLog is entered.
func (s *BaseMySQLStatementListener) EnterPurgeBinaryLog(ctx *PurgeBinaryLogContext) {}

// ExitPurgeBinaryLog is called when production purgeBinaryLog is exited.
func (s *BaseMySQLStatementListener) ExitPurgeBinaryLog(ctx *PurgeBinaryLogContext) {}

// EnterThreadTypes is called when production threadTypes is entered.
func (s *BaseMySQLStatementListener) EnterThreadTypes(ctx *ThreadTypesContext) {}

// ExitThreadTypes is called when production threadTypes is exited.
func (s *BaseMySQLStatementListener) ExitThreadTypes(ctx *ThreadTypesContext) {}

// EnterThreadType is called when production threadType is entered.
func (s *BaseMySQLStatementListener) EnterThreadType(ctx *ThreadTypeContext) {}

// ExitThreadType is called when production threadType is exited.
func (s *BaseMySQLStatementListener) ExitThreadType(ctx *ThreadTypeContext) {}

// EnterUtilOption is called when production utilOption is entered.
func (s *BaseMySQLStatementListener) EnterUtilOption(ctx *UtilOptionContext) {}

// ExitUtilOption is called when production utilOption is exited.
func (s *BaseMySQLStatementListener) ExitUtilOption(ctx *UtilOptionContext) {}

// EnterConnectionOptions is called when production connectionOptions is entered.
func (s *BaseMySQLStatementListener) EnterConnectionOptions(ctx *ConnectionOptionsContext) {}

// ExitConnectionOptions is called when production connectionOptions is exited.
func (s *BaseMySQLStatementListener) ExitConnectionOptions(ctx *ConnectionOptionsContext) {}

// EnterMasterDefs is called when production masterDefs is entered.
func (s *BaseMySQLStatementListener) EnterMasterDefs(ctx *MasterDefsContext) {}

// ExitMasterDefs is called when production masterDefs is exited.
func (s *BaseMySQLStatementListener) ExitMasterDefs(ctx *MasterDefsContext) {}

// EnterMasterDef is called when production masterDef is entered.
func (s *BaseMySQLStatementListener) EnterMasterDef(ctx *MasterDefContext) {}

// ExitMasterDef is called when production masterDef is exited.
func (s *BaseMySQLStatementListener) ExitMasterDef(ctx *MasterDefContext) {}

// EnterIgnoreServerIds is called when production ignoreServerIds is entered.
func (s *BaseMySQLStatementListener) EnterIgnoreServerIds(ctx *IgnoreServerIdsContext) {}

// ExitIgnoreServerIds is called when production ignoreServerIds is exited.
func (s *BaseMySQLStatementListener) ExitIgnoreServerIds(ctx *IgnoreServerIdsContext) {}

// EnterIgnoreServerId is called when production ignoreServerId is entered.
func (s *BaseMySQLStatementListener) EnterIgnoreServerId(ctx *IgnoreServerIdContext) {}

// ExitIgnoreServerId is called when production ignoreServerId is exited.
func (s *BaseMySQLStatementListener) ExitIgnoreServerId(ctx *IgnoreServerIdContext) {}

// EnterFilterDefs is called when production filterDefs is entered.
func (s *BaseMySQLStatementListener) EnterFilterDefs(ctx *FilterDefsContext) {}

// ExitFilterDefs is called when production filterDefs is exited.
func (s *BaseMySQLStatementListener) ExitFilterDefs(ctx *FilterDefsContext) {}

// EnterFilterDef is called when production filterDef is entered.
func (s *BaseMySQLStatementListener) EnterFilterDef(ctx *FilterDefContext) {}

// ExitFilterDef is called when production filterDef is exited.
func (s *BaseMySQLStatementListener) ExitFilterDef(ctx *FilterDefContext) {}

// EnterWildTables is called when production wildTables is entered.
func (s *BaseMySQLStatementListener) EnterWildTables(ctx *WildTablesContext) {}

// ExitWildTables is called when production wildTables is exited.
func (s *BaseMySQLStatementListener) ExitWildTables(ctx *WildTablesContext) {}

// EnterWildTable is called when production wildTable is entered.
func (s *BaseMySQLStatementListener) EnterWildTable(ctx *WildTableContext) {}

// ExitWildTable is called when production wildTable is exited.
func (s *BaseMySQLStatementListener) ExitWildTable(ctx *WildTableContext) {}

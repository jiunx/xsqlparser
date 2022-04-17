// Code generated from E:/GoProject/src/github.com/2212442929/xsqlparser/dialect/oracle/grammer\OracleStatement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // OracleStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOracleStatementListener is a complete listener for a parse tree produced by OracleStatementParser.
type BaseOracleStatementListener struct{}

var _ OracleStatementListener = &BaseOracleStatementListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOracleStatementListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOracleStatementListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOracleStatementListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOracleStatementListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExecute is called when production execute is entered.
func (s *BaseOracleStatementListener) EnterExecute(ctx *ExecuteContext) {}

// ExitExecute is called when production execute is exited.
func (s *BaseOracleStatementListener) ExitExecute(ctx *ExecuteContext) {}

// EnterInsert is called when production insert is entered.
func (s *BaseOracleStatementListener) EnterInsert(ctx *InsertContext) {}

// ExitInsert is called when production insert is exited.
func (s *BaseOracleStatementListener) ExitInsert(ctx *InsertContext) {}

// EnterInsertSingleTable is called when production insertSingleTable is entered.
func (s *BaseOracleStatementListener) EnterInsertSingleTable(ctx *InsertSingleTableContext) {}

// ExitInsertSingleTable is called when production insertSingleTable is exited.
func (s *BaseOracleStatementListener) ExitInsertSingleTable(ctx *InsertSingleTableContext) {}

// EnterInsertMultiTable is called when production insertMultiTable is entered.
func (s *BaseOracleStatementListener) EnterInsertMultiTable(ctx *InsertMultiTableContext) {}

// ExitInsertMultiTable is called when production insertMultiTable is exited.
func (s *BaseOracleStatementListener) ExitInsertMultiTable(ctx *InsertMultiTableContext) {}

// EnterMultiTableElement is called when production multiTableElement is entered.
func (s *BaseOracleStatementListener) EnterMultiTableElement(ctx *MultiTableElementContext) {}

// ExitMultiTableElement is called when production multiTableElement is exited.
func (s *BaseOracleStatementListener) ExitMultiTableElement(ctx *MultiTableElementContext) {}

// EnterConditionalInsertClause is called when production conditionalInsertClause is entered.
func (s *BaseOracleStatementListener) EnterConditionalInsertClause(ctx *ConditionalInsertClauseContext) {
}

// ExitConditionalInsertClause is called when production conditionalInsertClause is exited.
func (s *BaseOracleStatementListener) ExitConditionalInsertClause(ctx *ConditionalInsertClauseContext) {
}

// EnterConditionalInsertWhenPart is called when production conditionalInsertWhenPart is entered.
func (s *BaseOracleStatementListener) EnterConditionalInsertWhenPart(ctx *ConditionalInsertWhenPartContext) {
}

// ExitConditionalInsertWhenPart is called when production conditionalInsertWhenPart is exited.
func (s *BaseOracleStatementListener) ExitConditionalInsertWhenPart(ctx *ConditionalInsertWhenPartContext) {
}

// EnterConditionalInsertElsePart is called when production conditionalInsertElsePart is entered.
func (s *BaseOracleStatementListener) EnterConditionalInsertElsePart(ctx *ConditionalInsertElsePartContext) {
}

// ExitConditionalInsertElsePart is called when production conditionalInsertElsePart is exited.
func (s *BaseOracleStatementListener) ExitConditionalInsertElsePart(ctx *ConditionalInsertElsePartContext) {
}

// EnterInsertIntoClause is called when production insertIntoClause is entered.
func (s *BaseOracleStatementListener) EnterInsertIntoClause(ctx *InsertIntoClauseContext) {}

// ExitInsertIntoClause is called when production insertIntoClause is exited.
func (s *BaseOracleStatementListener) ExitInsertIntoClause(ctx *InsertIntoClauseContext) {}

// EnterInsertValuesClause is called when production insertValuesClause is entered.
func (s *BaseOracleStatementListener) EnterInsertValuesClause(ctx *InsertValuesClauseContext) {}

// ExitInsertValuesClause is called when production insertValuesClause is exited.
func (s *BaseOracleStatementListener) ExitInsertValuesClause(ctx *InsertValuesClauseContext) {}

// EnterReturningClause is called when production returningClause is entered.
func (s *BaseOracleStatementListener) EnterReturningClause(ctx *ReturningClauseContext) {}

// ExitReturningClause is called when production returningClause is exited.
func (s *BaseOracleStatementListener) ExitReturningClause(ctx *ReturningClauseContext) {}

// EnterDmlTableExprClause is called when production dmlTableExprClause is entered.
func (s *BaseOracleStatementListener) EnterDmlTableExprClause(ctx *DmlTableExprClauseContext) {}

// ExitDmlTableExprClause is called when production dmlTableExprClause is exited.
func (s *BaseOracleStatementListener) ExitDmlTableExprClause(ctx *DmlTableExprClauseContext) {}

// EnterDmlTableClause is called when production dmlTableClause is entered.
func (s *BaseOracleStatementListener) EnterDmlTableClause(ctx *DmlTableClauseContext) {}

// ExitDmlTableClause is called when production dmlTableClause is exited.
func (s *BaseOracleStatementListener) ExitDmlTableClause(ctx *DmlTableClauseContext) {}

// EnterPartitionExtClause is called when production partitionExtClause is entered.
func (s *BaseOracleStatementListener) EnterPartitionExtClause(ctx *PartitionExtClauseContext) {}

// ExitPartitionExtClause is called when production partitionExtClause is exited.
func (s *BaseOracleStatementListener) ExitPartitionExtClause(ctx *PartitionExtClauseContext) {}

// EnterDmlSubqueryClause is called when production dmlSubqueryClause is entered.
func (s *BaseOracleStatementListener) EnterDmlSubqueryClause(ctx *DmlSubqueryClauseContext) {}

// ExitDmlSubqueryClause is called when production dmlSubqueryClause is exited.
func (s *BaseOracleStatementListener) ExitDmlSubqueryClause(ctx *DmlSubqueryClauseContext) {}

// EnterSubqueryRestrictionClause is called when production subqueryRestrictionClause is entered.
func (s *BaseOracleStatementListener) EnterSubqueryRestrictionClause(ctx *SubqueryRestrictionClauseContext) {
}

// ExitSubqueryRestrictionClause is called when production subqueryRestrictionClause is exited.
func (s *BaseOracleStatementListener) ExitSubqueryRestrictionClause(ctx *SubqueryRestrictionClauseContext) {
}

// EnterTableCollectionExpr is called when production tableCollectionExpr is entered.
func (s *BaseOracleStatementListener) EnterTableCollectionExpr(ctx *TableCollectionExprContext) {}

// ExitTableCollectionExpr is called when production tableCollectionExpr is exited.
func (s *BaseOracleStatementListener) ExitTableCollectionExpr(ctx *TableCollectionExprContext) {}

// EnterCollectionExpr is called when production collectionExpr is entered.
func (s *BaseOracleStatementListener) EnterCollectionExpr(ctx *CollectionExprContext) {}

// ExitCollectionExpr is called when production collectionExpr is exited.
func (s *BaseOracleStatementListener) ExitCollectionExpr(ctx *CollectionExprContext) {}

// EnterUpdate is called when production update is entered.
func (s *BaseOracleStatementListener) EnterUpdate(ctx *UpdateContext) {}

// ExitUpdate is called when production update is exited.
func (s *BaseOracleStatementListener) ExitUpdate(ctx *UpdateContext) {}

// EnterUpdateSpecification is called when production updateSpecification is entered.
func (s *BaseOracleStatementListener) EnterUpdateSpecification(ctx *UpdateSpecificationContext) {}

// ExitUpdateSpecification is called when production updateSpecification is exited.
func (s *BaseOracleStatementListener) ExitUpdateSpecification(ctx *UpdateSpecificationContext) {}

// EnterUpdateSetClause is called when production updateSetClause is entered.
func (s *BaseOracleStatementListener) EnterUpdateSetClause(ctx *UpdateSetClauseContext) {}

// ExitUpdateSetClause is called when production updateSetClause is exited.
func (s *BaseOracleStatementListener) ExitUpdateSetClause(ctx *UpdateSetClauseContext) {}

// EnterUpdateSetColumnList is called when production updateSetColumnList is entered.
func (s *BaseOracleStatementListener) EnterUpdateSetColumnList(ctx *UpdateSetColumnListContext) {}

// ExitUpdateSetColumnList is called when production updateSetColumnList is exited.
func (s *BaseOracleStatementListener) ExitUpdateSetColumnList(ctx *UpdateSetColumnListContext) {}

// EnterUpdateSetColumnClause is called when production updateSetColumnClause is entered.
func (s *BaseOracleStatementListener) EnterUpdateSetColumnClause(ctx *UpdateSetColumnClauseContext) {}

// ExitUpdateSetColumnClause is called when production updateSetColumnClause is exited.
func (s *BaseOracleStatementListener) ExitUpdateSetColumnClause(ctx *UpdateSetColumnClauseContext) {}

// EnterUpdateSetValueClause is called when production updateSetValueClause is entered.
func (s *BaseOracleStatementListener) EnterUpdateSetValueClause(ctx *UpdateSetValueClauseContext) {}

// ExitUpdateSetValueClause is called when production updateSetValueClause is exited.
func (s *BaseOracleStatementListener) ExitUpdateSetValueClause(ctx *UpdateSetValueClauseContext) {}

// EnterAssignmentValues is called when production assignmentValues is entered.
func (s *BaseOracleStatementListener) EnterAssignmentValues(ctx *AssignmentValuesContext) {}

// ExitAssignmentValues is called when production assignmentValues is exited.
func (s *BaseOracleStatementListener) ExitAssignmentValues(ctx *AssignmentValuesContext) {}

// EnterAssignmentValue is called when production assignmentValue is entered.
func (s *BaseOracleStatementListener) EnterAssignmentValue(ctx *AssignmentValueContext) {}

// ExitAssignmentValue is called when production assignmentValue is exited.
func (s *BaseOracleStatementListener) ExitAssignmentValue(ctx *AssignmentValueContext) {}

// EnterDelete is called when production delete is entered.
func (s *BaseOracleStatementListener) EnterDelete(ctx *DeleteContext) {}

// ExitDelete is called when production delete is exited.
func (s *BaseOracleStatementListener) ExitDelete(ctx *DeleteContext) {}

// EnterDeleteSpecification is called when production deleteSpecification is entered.
func (s *BaseOracleStatementListener) EnterDeleteSpecification(ctx *DeleteSpecificationContext) {}

// ExitDeleteSpecification is called when production deleteSpecification is exited.
func (s *BaseOracleStatementListener) ExitDeleteSpecification(ctx *DeleteSpecificationContext) {}

// EnterSelect is called when production select is entered.
func (s *BaseOracleStatementListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production select is exited.
func (s *BaseOracleStatementListener) ExitSelect(ctx *SelectContext) {}

// EnterSelectSubquery is called when production selectSubquery is entered.
func (s *BaseOracleStatementListener) EnterSelectSubquery(ctx *SelectSubqueryContext) {}

// ExitSelectSubquery is called when production selectSubquery is exited.
func (s *BaseOracleStatementListener) ExitSelectSubquery(ctx *SelectSubqueryContext) {}

// EnterSelectUnionClause is called when production selectUnionClause is entered.
func (s *BaseOracleStatementListener) EnterSelectUnionClause(ctx *SelectUnionClauseContext) {}

// ExitSelectUnionClause is called when production selectUnionClause is exited.
func (s *BaseOracleStatementListener) ExitSelectUnionClause(ctx *SelectUnionClauseContext) {}

// EnterParenthesisSelectSubquery is called when production parenthesisSelectSubquery is entered.
func (s *BaseOracleStatementListener) EnterParenthesisSelectSubquery(ctx *ParenthesisSelectSubqueryContext) {
}

// ExitParenthesisSelectSubquery is called when production parenthesisSelectSubquery is exited.
func (s *BaseOracleStatementListener) ExitParenthesisSelectSubquery(ctx *ParenthesisSelectSubqueryContext) {
}

// EnterQueryBlock is called when production queryBlock is entered.
func (s *BaseOracleStatementListener) EnterQueryBlock(ctx *QueryBlockContext) {}

// ExitQueryBlock is called when production queryBlock is exited.
func (s *BaseOracleStatementListener) ExitQueryBlock(ctx *QueryBlockContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseOracleStatementListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseOracleStatementListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterPlsqlDeclarations is called when production plsqlDeclarations is entered.
func (s *BaseOracleStatementListener) EnterPlsqlDeclarations(ctx *PlsqlDeclarationsContext) {}

// ExitPlsqlDeclarations is called when production plsqlDeclarations is exited.
func (s *BaseOracleStatementListener) ExitPlsqlDeclarations(ctx *PlsqlDeclarationsContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseOracleStatementListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseOracleStatementListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionHeading is called when production functionHeading is entered.
func (s *BaseOracleStatementListener) EnterFunctionHeading(ctx *FunctionHeadingContext) {}

// ExitFunctionHeading is called when production functionHeading is exited.
func (s *BaseOracleStatementListener) ExitFunctionHeading(ctx *FunctionHeadingContext) {}

// EnterParameterDeclaration is called when production parameterDeclaration is entered.
func (s *BaseOracleStatementListener) EnterParameterDeclaration(ctx *ParameterDeclarationContext) {}

// ExitParameterDeclaration is called when production parameterDeclaration is exited.
func (s *BaseOracleStatementListener) ExitParameterDeclaration(ctx *ParameterDeclarationContext) {}

// EnterProcedureDeclaration is called when production procedureDeclaration is entered.
func (s *BaseOracleStatementListener) EnterProcedureDeclaration(ctx *ProcedureDeclarationContext) {}

// ExitProcedureDeclaration is called when production procedureDeclaration is exited.
func (s *BaseOracleStatementListener) ExitProcedureDeclaration(ctx *ProcedureDeclarationContext) {}

// EnterProcedureHeading is called when production procedureHeading is entered.
func (s *BaseOracleStatementListener) EnterProcedureHeading(ctx *ProcedureHeadingContext) {}

// ExitProcedureHeading is called when production procedureHeading is exited.
func (s *BaseOracleStatementListener) ExitProcedureHeading(ctx *ProcedureHeadingContext) {}

// EnterProcedureProperties is called when production procedureProperties is entered.
func (s *BaseOracleStatementListener) EnterProcedureProperties(ctx *ProcedurePropertiesContext) {}

// ExitProcedureProperties is called when production procedureProperties is exited.
func (s *BaseOracleStatementListener) ExitProcedureProperties(ctx *ProcedurePropertiesContext) {}

// EnterAccessibleByClause is called when production accessibleByClause is entered.
func (s *BaseOracleStatementListener) EnterAccessibleByClause(ctx *AccessibleByClauseContext) {}

// ExitAccessibleByClause is called when production accessibleByClause is exited.
func (s *BaseOracleStatementListener) ExitAccessibleByClause(ctx *AccessibleByClauseContext) {}

// EnterAccessor is called when production accessor is entered.
func (s *BaseOracleStatementListener) EnterAccessor(ctx *AccessorContext) {}

// ExitAccessor is called when production accessor is exited.
func (s *BaseOracleStatementListener) ExitAccessor(ctx *AccessorContext) {}

// EnterUnitKind is called when production unitKind is entered.
func (s *BaseOracleStatementListener) EnterUnitKind(ctx *UnitKindContext) {}

// ExitUnitKind is called when production unitKind is exited.
func (s *BaseOracleStatementListener) ExitUnitKind(ctx *UnitKindContext) {}

// EnterDefaultCollationClause is called when production defaultCollationClause is entered.
func (s *BaseOracleStatementListener) EnterDefaultCollationClause(ctx *DefaultCollationClauseContext) {
}

// ExitDefaultCollationClause is called when production defaultCollationClause is exited.
func (s *BaseOracleStatementListener) ExitDefaultCollationClause(ctx *DefaultCollationClauseContext) {
}

// EnterCollationOption is called when production collationOption is entered.
func (s *BaseOracleStatementListener) EnterCollationOption(ctx *CollationOptionContext) {}

// ExitCollationOption is called when production collationOption is exited.
func (s *BaseOracleStatementListener) ExitCollationOption(ctx *CollationOptionContext) {}

// EnterInvokerRightsClause is called when production invokerRightsClause is entered.
func (s *BaseOracleStatementListener) EnterInvokerRightsClause(ctx *InvokerRightsClauseContext) {}

// ExitInvokerRightsClause is called when production invokerRightsClause is exited.
func (s *BaseOracleStatementListener) ExitInvokerRightsClause(ctx *InvokerRightsClauseContext) {}

// EnterSubqueryFactoringClause is called when production subqueryFactoringClause is entered.
func (s *BaseOracleStatementListener) EnterSubqueryFactoringClause(ctx *SubqueryFactoringClauseContext) {
}

// ExitSubqueryFactoringClause is called when production subqueryFactoringClause is exited.
func (s *BaseOracleStatementListener) ExitSubqueryFactoringClause(ctx *SubqueryFactoringClauseContext) {
}

// EnterSearchClause is called when production searchClause is entered.
func (s *BaseOracleStatementListener) EnterSearchClause(ctx *SearchClauseContext) {}

// ExitSearchClause is called when production searchClause is exited.
func (s *BaseOracleStatementListener) ExitSearchClause(ctx *SearchClauseContext) {}

// EnterCycleClause is called when production cycleClause is entered.
func (s *BaseOracleStatementListener) EnterCycleClause(ctx *CycleClauseContext) {}

// ExitCycleClause is called when production cycleClause is exited.
func (s *BaseOracleStatementListener) ExitCycleClause(ctx *CycleClauseContext) {}

// EnterSubavFactoringClause is called when production subavFactoringClause is entered.
func (s *BaseOracleStatementListener) EnterSubavFactoringClause(ctx *SubavFactoringClauseContext) {}

// ExitSubavFactoringClause is called when production subavFactoringClause is exited.
func (s *BaseOracleStatementListener) ExitSubavFactoringClause(ctx *SubavFactoringClauseContext) {}

// EnterSubavClause is called when production subavClause is entered.
func (s *BaseOracleStatementListener) EnterSubavClause(ctx *SubavClauseContext) {}

// ExitSubavClause is called when production subavClause is exited.
func (s *BaseOracleStatementListener) ExitSubavClause(ctx *SubavClauseContext) {}

// EnterHierarchiesClause is called when production hierarchiesClause is entered.
func (s *BaseOracleStatementListener) EnterHierarchiesClause(ctx *HierarchiesClauseContext) {}

// ExitHierarchiesClause is called when production hierarchiesClause is exited.
func (s *BaseOracleStatementListener) ExitHierarchiesClause(ctx *HierarchiesClauseContext) {}

// EnterFilterClauses is called when production filterClauses is entered.
func (s *BaseOracleStatementListener) EnterFilterClauses(ctx *FilterClausesContext) {}

// ExitFilterClauses is called when production filterClauses is exited.
func (s *BaseOracleStatementListener) ExitFilterClauses(ctx *FilterClausesContext) {}

// EnterFilterClause is called when production filterClause is entered.
func (s *BaseOracleStatementListener) EnterFilterClause(ctx *FilterClauseContext) {}

// ExitFilterClause is called when production filterClause is exited.
func (s *BaseOracleStatementListener) ExitFilterClause(ctx *FilterClauseContext) {}

// EnterAddCalcsClause is called when production addCalcsClause is entered.
func (s *BaseOracleStatementListener) EnterAddCalcsClause(ctx *AddCalcsClauseContext) {}

// ExitAddCalcsClause is called when production addCalcsClause is exited.
func (s *BaseOracleStatementListener) ExitAddCalcsClause(ctx *AddCalcsClauseContext) {}

// EnterCalcMeasClause is called when production calcMeasClause is entered.
func (s *BaseOracleStatementListener) EnterCalcMeasClause(ctx *CalcMeasClauseContext) {}

// ExitCalcMeasClause is called when production calcMeasClause is exited.
func (s *BaseOracleStatementListener) ExitCalcMeasClause(ctx *CalcMeasClauseContext) {}

// EnterCalcMeasExpression is called when production calcMeasExpression is entered.
func (s *BaseOracleStatementListener) EnterCalcMeasExpression(ctx *CalcMeasExpressionContext) {}

// ExitCalcMeasExpression is called when production calcMeasExpression is exited.
func (s *BaseOracleStatementListener) ExitCalcMeasExpression(ctx *CalcMeasExpressionContext) {}

// EnterAvExpression is called when production avExpression is entered.
func (s *BaseOracleStatementListener) EnterAvExpression(ctx *AvExpressionContext) {}

// ExitAvExpression is called when production avExpression is exited.
func (s *BaseOracleStatementListener) ExitAvExpression(ctx *AvExpressionContext) {}

// EnterAvMeasExpression is called when production avMeasExpression is entered.
func (s *BaseOracleStatementListener) EnterAvMeasExpression(ctx *AvMeasExpressionContext) {}

// ExitAvMeasExpression is called when production avMeasExpression is exited.
func (s *BaseOracleStatementListener) ExitAvMeasExpression(ctx *AvMeasExpressionContext) {}

// EnterLeadLagExpression is called when production leadLagExpression is entered.
func (s *BaseOracleStatementListener) EnterLeadLagExpression(ctx *LeadLagExpressionContext) {}

// ExitLeadLagExpression is called when production leadLagExpression is exited.
func (s *BaseOracleStatementListener) ExitLeadLagExpression(ctx *LeadLagExpressionContext) {}

// EnterLeadLagFunctionName is called when production leadLagFunctionName is entered.
func (s *BaseOracleStatementListener) EnterLeadLagFunctionName(ctx *LeadLagFunctionNameContext) {}

// ExitLeadLagFunctionName is called when production leadLagFunctionName is exited.
func (s *BaseOracleStatementListener) ExitLeadLagFunctionName(ctx *LeadLagFunctionNameContext) {}

// EnterLeadLagClause is called when production leadLagClause is entered.
func (s *BaseOracleStatementListener) EnterLeadLagClause(ctx *LeadLagClauseContext) {}

// ExitLeadLagClause is called when production leadLagClause is exited.
func (s *BaseOracleStatementListener) ExitLeadLagClause(ctx *LeadLagClauseContext) {}

// EnterHierarchyRef is called when production hierarchyRef is entered.
func (s *BaseOracleStatementListener) EnterHierarchyRef(ctx *HierarchyRefContext) {}

// ExitHierarchyRef is called when production hierarchyRef is exited.
func (s *BaseOracleStatementListener) ExitHierarchyRef(ctx *HierarchyRefContext) {}

// EnterWindowExpression is called when production windowExpression is entered.
func (s *BaseOracleStatementListener) EnterWindowExpression(ctx *WindowExpressionContext) {}

// ExitWindowExpression is called when production windowExpression is exited.
func (s *BaseOracleStatementListener) ExitWindowExpression(ctx *WindowExpressionContext) {}

// EnterWindowClause is called when production windowClause is entered.
func (s *BaseOracleStatementListener) EnterWindowClause(ctx *WindowClauseContext) {}

// ExitWindowClause is called when production windowClause is exited.
func (s *BaseOracleStatementListener) ExitWindowClause(ctx *WindowClauseContext) {}

// EnterPrecedingBoundary is called when production precedingBoundary is entered.
func (s *BaseOracleStatementListener) EnterPrecedingBoundary(ctx *PrecedingBoundaryContext) {}

// ExitPrecedingBoundary is called when production precedingBoundary is exited.
func (s *BaseOracleStatementListener) ExitPrecedingBoundary(ctx *PrecedingBoundaryContext) {}

// EnterFollowingBoundary is called when production followingBoundary is entered.
func (s *BaseOracleStatementListener) EnterFollowingBoundary(ctx *FollowingBoundaryContext) {}

// ExitFollowingBoundary is called when production followingBoundary is exited.
func (s *BaseOracleStatementListener) ExitFollowingBoundary(ctx *FollowingBoundaryContext) {}

// EnterRankExpression is called when production rankExpression is entered.
func (s *BaseOracleStatementListener) EnterRankExpression(ctx *RankExpressionContext) {}

// ExitRankExpression is called when production rankExpression is exited.
func (s *BaseOracleStatementListener) ExitRankExpression(ctx *RankExpressionContext) {}

// EnterRankFunctionName is called when production rankFunctionName is entered.
func (s *BaseOracleStatementListener) EnterRankFunctionName(ctx *RankFunctionNameContext) {}

// ExitRankFunctionName is called when production rankFunctionName is exited.
func (s *BaseOracleStatementListener) ExitRankFunctionName(ctx *RankFunctionNameContext) {}

// EnterRankClause is called when production rankClause is entered.
func (s *BaseOracleStatementListener) EnterRankClause(ctx *RankClauseContext) {}

// ExitRankClause is called when production rankClause is exited.
func (s *BaseOracleStatementListener) ExitRankClause(ctx *RankClauseContext) {}

// EnterCalcMeasOrderByClause is called when production calcMeasOrderByClause is entered.
func (s *BaseOracleStatementListener) EnterCalcMeasOrderByClause(ctx *CalcMeasOrderByClauseContext) {}

// ExitCalcMeasOrderByClause is called when production calcMeasOrderByClause is exited.
func (s *BaseOracleStatementListener) ExitCalcMeasOrderByClause(ctx *CalcMeasOrderByClauseContext) {}

// EnterShareOfExpression is called when production shareOfExpression is entered.
func (s *BaseOracleStatementListener) EnterShareOfExpression(ctx *ShareOfExpressionContext) {}

// ExitShareOfExpression is called when production shareOfExpression is exited.
func (s *BaseOracleStatementListener) ExitShareOfExpression(ctx *ShareOfExpressionContext) {}

// EnterShareClause is called when production shareClause is entered.
func (s *BaseOracleStatementListener) EnterShareClause(ctx *ShareClauseContext) {}

// ExitShareClause is called when production shareClause is exited.
func (s *BaseOracleStatementListener) ExitShareClause(ctx *ShareClauseContext) {}

// EnterMemberExpression is called when production memberExpression is entered.
func (s *BaseOracleStatementListener) EnterMemberExpression(ctx *MemberExpressionContext) {}

// ExitMemberExpression is called when production memberExpression is exited.
func (s *BaseOracleStatementListener) ExitMemberExpression(ctx *MemberExpressionContext) {}

// EnterLevelMemberLiteral is called when production levelMemberLiteral is entered.
func (s *BaseOracleStatementListener) EnterLevelMemberLiteral(ctx *LevelMemberLiteralContext) {}

// ExitLevelMemberLiteral is called when production levelMemberLiteral is exited.
func (s *BaseOracleStatementListener) ExitLevelMemberLiteral(ctx *LevelMemberLiteralContext) {}

// EnterPosMemberKeys is called when production posMemberKeys is entered.
func (s *BaseOracleStatementListener) EnterPosMemberKeys(ctx *PosMemberKeysContext) {}

// ExitPosMemberKeys is called when production posMemberKeys is exited.
func (s *BaseOracleStatementListener) ExitPosMemberKeys(ctx *PosMemberKeysContext) {}

// EnterNamedMemberKeys is called when production namedMemberKeys is entered.
func (s *BaseOracleStatementListener) EnterNamedMemberKeys(ctx *NamedMemberKeysContext) {}

// ExitNamedMemberKeys is called when production namedMemberKeys is exited.
func (s *BaseOracleStatementListener) ExitNamedMemberKeys(ctx *NamedMemberKeysContext) {}

// EnterHierNavigationExpression is called when production hierNavigationExpression is entered.
func (s *BaseOracleStatementListener) EnterHierNavigationExpression(ctx *HierNavigationExpressionContext) {
}

// ExitHierNavigationExpression is called when production hierNavigationExpression is exited.
func (s *BaseOracleStatementListener) ExitHierNavigationExpression(ctx *HierNavigationExpressionContext) {
}

// EnterHierAncestorExpression is called when production hierAncestorExpression is entered.
func (s *BaseOracleStatementListener) EnterHierAncestorExpression(ctx *HierAncestorExpressionContext) {
}

// ExitHierAncestorExpression is called when production hierAncestorExpression is exited.
func (s *BaseOracleStatementListener) ExitHierAncestorExpression(ctx *HierAncestorExpressionContext) {
}

// EnterHierParentExpression is called when production hierParentExpression is entered.
func (s *BaseOracleStatementListener) EnterHierParentExpression(ctx *HierParentExpressionContext) {}

// ExitHierParentExpression is called when production hierParentExpression is exited.
func (s *BaseOracleStatementListener) ExitHierParentExpression(ctx *HierParentExpressionContext) {}

// EnterHierLeadLagExpression is called when production hierLeadLagExpression is entered.
func (s *BaseOracleStatementListener) EnterHierLeadLagExpression(ctx *HierLeadLagExpressionContext) {}

// ExitHierLeadLagExpression is called when production hierLeadLagExpression is exited.
func (s *BaseOracleStatementListener) ExitHierLeadLagExpression(ctx *HierLeadLagExpressionContext) {}

// EnterHierLeadLagClause is called when production hierLeadLagClause is entered.
func (s *BaseOracleStatementListener) EnterHierLeadLagClause(ctx *HierLeadLagClauseContext) {}

// ExitHierLeadLagClause is called when production hierLeadLagClause is exited.
func (s *BaseOracleStatementListener) ExitHierLeadLagClause(ctx *HierLeadLagClauseContext) {}

// EnterQdrExpression is called when production qdrExpression is entered.
func (s *BaseOracleStatementListener) EnterQdrExpression(ctx *QdrExpressionContext) {}

// ExitQdrExpression is called when production qdrExpression is exited.
func (s *BaseOracleStatementListener) ExitQdrExpression(ctx *QdrExpressionContext) {}

// EnterQualifier is called when production qualifier is entered.
func (s *BaseOracleStatementListener) EnterQualifier(ctx *QualifierContext) {}

// ExitQualifier is called when production qualifier is exited.
func (s *BaseOracleStatementListener) ExitQualifier(ctx *QualifierContext) {}

// EnterAvHierExpression is called when production avHierExpression is entered.
func (s *BaseOracleStatementListener) EnterAvHierExpression(ctx *AvHierExpressionContext) {}

// ExitAvHierExpression is called when production avHierExpression is exited.
func (s *BaseOracleStatementListener) ExitAvHierExpression(ctx *AvHierExpressionContext) {}

// EnterHierFunctionName is called when production hierFunctionName is entered.
func (s *BaseOracleStatementListener) EnterHierFunctionName(ctx *HierFunctionNameContext) {}

// ExitHierFunctionName is called when production hierFunctionName is exited.
func (s *BaseOracleStatementListener) ExitHierFunctionName(ctx *HierFunctionNameContext) {}

// EnterDuplicateSpecification is called when production duplicateSpecification is entered.
func (s *BaseOracleStatementListener) EnterDuplicateSpecification(ctx *DuplicateSpecificationContext) {
}

// ExitDuplicateSpecification is called when production duplicateSpecification is exited.
func (s *BaseOracleStatementListener) ExitDuplicateSpecification(ctx *DuplicateSpecificationContext) {
}

// EnterUnqualifiedShorthand is called when production unqualifiedShorthand is entered.
func (s *BaseOracleStatementListener) EnterUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) {}

// ExitUnqualifiedShorthand is called when production unqualifiedShorthand is exited.
func (s *BaseOracleStatementListener) ExitUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) {}

// EnterSelectList is called when production selectList is entered.
func (s *BaseOracleStatementListener) EnterSelectList(ctx *SelectListContext) {}

// ExitSelectList is called when production selectList is exited.
func (s *BaseOracleStatementListener) ExitSelectList(ctx *SelectListContext) {}

// EnterSelectProjection is called when production selectProjection is entered.
func (s *BaseOracleStatementListener) EnterSelectProjection(ctx *SelectProjectionContext) {}

// ExitSelectProjection is called when production selectProjection is exited.
func (s *BaseOracleStatementListener) ExitSelectProjection(ctx *SelectProjectionContext) {}

// EnterSelectProjectionExprClause is called when production selectProjectionExprClause is entered.
func (s *BaseOracleStatementListener) EnterSelectProjectionExprClause(ctx *SelectProjectionExprClauseContext) {
}

// ExitSelectProjectionExprClause is called when production selectProjectionExprClause is exited.
func (s *BaseOracleStatementListener) ExitSelectProjectionExprClause(ctx *SelectProjectionExprClauseContext) {
}

// EnterSelectFromClause is called when production selectFromClause is entered.
func (s *BaseOracleStatementListener) EnterSelectFromClause(ctx *SelectFromClauseContext) {}

// ExitSelectFromClause is called when production selectFromClause is exited.
func (s *BaseOracleStatementListener) ExitSelectFromClause(ctx *SelectFromClauseContext) {}

// EnterFromClauseList is called when production fromClauseList is entered.
func (s *BaseOracleStatementListener) EnterFromClauseList(ctx *FromClauseListContext) {}

// ExitFromClauseList is called when production fromClauseList is exited.
func (s *BaseOracleStatementListener) ExitFromClauseList(ctx *FromClauseListContext) {}

// EnterFromClauseOption is called when production fromClauseOption is entered.
func (s *BaseOracleStatementListener) EnterFromClauseOption(ctx *FromClauseOptionContext) {}

// ExitFromClauseOption is called when production fromClauseOption is exited.
func (s *BaseOracleStatementListener) ExitFromClauseOption(ctx *FromClauseOptionContext) {}

// EnterSelectTableReference is called when production selectTableReference is entered.
func (s *BaseOracleStatementListener) EnterSelectTableReference(ctx *SelectTableReferenceContext) {}

// ExitSelectTableReference is called when production selectTableReference is exited.
func (s *BaseOracleStatementListener) ExitSelectTableReference(ctx *SelectTableReferenceContext) {}

// EnterQueryTableExprClause is called when production queryTableExprClause is entered.
func (s *BaseOracleStatementListener) EnterQueryTableExprClause(ctx *QueryTableExprClauseContext) {}

// ExitQueryTableExprClause is called when production queryTableExprClause is exited.
func (s *BaseOracleStatementListener) ExitQueryTableExprClause(ctx *QueryTableExprClauseContext) {}

// EnterFlashbackQueryClause is called when production flashbackQueryClause is entered.
func (s *BaseOracleStatementListener) EnterFlashbackQueryClause(ctx *FlashbackQueryClauseContext) {}

// ExitFlashbackQueryClause is called when production flashbackQueryClause is exited.
func (s *BaseOracleStatementListener) ExitFlashbackQueryClause(ctx *FlashbackQueryClauseContext) {}

// EnterQueryTableExpr is called when production queryTableExpr is entered.
func (s *BaseOracleStatementListener) EnterQueryTableExpr(ctx *QueryTableExprContext) {}

// ExitQueryTableExpr is called when production queryTableExpr is exited.
func (s *BaseOracleStatementListener) ExitQueryTableExpr(ctx *QueryTableExprContext) {}

// EnterLateralClause is called when production lateralClause is entered.
func (s *BaseOracleStatementListener) EnterLateralClause(ctx *LateralClauseContext) {}

// ExitLateralClause is called when production lateralClause is exited.
func (s *BaseOracleStatementListener) ExitLateralClause(ctx *LateralClauseContext) {}

// EnterQueryTableExprSampleClause is called when production queryTableExprSampleClause is entered.
func (s *BaseOracleStatementListener) EnterQueryTableExprSampleClause(ctx *QueryTableExprSampleClauseContext) {
}

// ExitQueryTableExprSampleClause is called when production queryTableExprSampleClause is exited.
func (s *BaseOracleStatementListener) ExitQueryTableExprSampleClause(ctx *QueryTableExprSampleClauseContext) {
}

// EnterQueryTableExprTableClause is called when production queryTableExprTableClause is entered.
func (s *BaseOracleStatementListener) EnterQueryTableExprTableClause(ctx *QueryTableExprTableClauseContext) {
}

// ExitQueryTableExprTableClause is called when production queryTableExprTableClause is exited.
func (s *BaseOracleStatementListener) ExitQueryTableExprTableClause(ctx *QueryTableExprTableClauseContext) {
}

// EnterQueryTableExprViewClause is called when production queryTableExprViewClause is entered.
func (s *BaseOracleStatementListener) EnterQueryTableExprViewClause(ctx *QueryTableExprViewClauseContext) {
}

// ExitQueryTableExprViewClause is called when production queryTableExprViewClause is exited.
func (s *BaseOracleStatementListener) ExitQueryTableExprViewClause(ctx *QueryTableExprViewClauseContext) {
}

// EnterQueryTableExprAnalyticClause is called when production queryTableExprAnalyticClause is entered.
func (s *BaseOracleStatementListener) EnterQueryTableExprAnalyticClause(ctx *QueryTableExprAnalyticClauseContext) {
}

// ExitQueryTableExprAnalyticClause is called when production queryTableExprAnalyticClause is exited.
func (s *BaseOracleStatementListener) ExitQueryTableExprAnalyticClause(ctx *QueryTableExprAnalyticClauseContext) {
}

// EnterInlineExternalTable is called when production inlineExternalTable is entered.
func (s *BaseOracleStatementListener) EnterInlineExternalTable(ctx *InlineExternalTableContext) {}

// ExitInlineExternalTable is called when production inlineExternalTable is exited.
func (s *BaseOracleStatementListener) ExitInlineExternalTable(ctx *InlineExternalTableContext) {}

// EnterInlineExternalTableProperties is called when production inlineExternalTableProperties is entered.
func (s *BaseOracleStatementListener) EnterInlineExternalTableProperties(ctx *InlineExternalTablePropertiesContext) {
}

// ExitInlineExternalTableProperties is called when production inlineExternalTableProperties is exited.
func (s *BaseOracleStatementListener) ExitInlineExternalTableProperties(ctx *InlineExternalTablePropertiesContext) {
}

// EnterExternalTableDataProperties is called when production externalTableDataProperties is entered.
func (s *BaseOracleStatementListener) EnterExternalTableDataProperties(ctx *ExternalTableDataPropertiesContext) {
}

// ExitExternalTableDataProperties is called when production externalTableDataProperties is exited.
func (s *BaseOracleStatementListener) ExitExternalTableDataProperties(ctx *ExternalTableDataPropertiesContext) {
}

// EnterMofifiedExternalTable is called when production mofifiedExternalTable is entered.
func (s *BaseOracleStatementListener) EnterMofifiedExternalTable(ctx *MofifiedExternalTableContext) {}

// ExitMofifiedExternalTable is called when production mofifiedExternalTable is exited.
func (s *BaseOracleStatementListener) ExitMofifiedExternalTable(ctx *MofifiedExternalTableContext) {}

// EnterModifyExternalTableProperties is called when production modifyExternalTableProperties is entered.
func (s *BaseOracleStatementListener) EnterModifyExternalTableProperties(ctx *ModifyExternalTablePropertiesContext) {
}

// ExitModifyExternalTableProperties is called when production modifyExternalTableProperties is exited.
func (s *BaseOracleStatementListener) ExitModifyExternalTableProperties(ctx *ModifyExternalTablePropertiesContext) {
}

// EnterPivotClause is called when production pivotClause is entered.
func (s *BaseOracleStatementListener) EnterPivotClause(ctx *PivotClauseContext) {}

// ExitPivotClause is called when production pivotClause is exited.
func (s *BaseOracleStatementListener) ExitPivotClause(ctx *PivotClauseContext) {}

// EnterPivotForClause is called when production pivotForClause is entered.
func (s *BaseOracleStatementListener) EnterPivotForClause(ctx *PivotForClauseContext) {}

// ExitPivotForClause is called when production pivotForClause is exited.
func (s *BaseOracleStatementListener) ExitPivotForClause(ctx *PivotForClauseContext) {}

// EnterPivotInClause is called when production pivotInClause is entered.
func (s *BaseOracleStatementListener) EnterPivotInClause(ctx *PivotInClauseContext) {}

// ExitPivotInClause is called when production pivotInClause is exited.
func (s *BaseOracleStatementListener) ExitPivotInClause(ctx *PivotInClauseContext) {}

// EnterUnpivotClause is called when production unpivotClause is entered.
func (s *BaseOracleStatementListener) EnterUnpivotClause(ctx *UnpivotClauseContext) {}

// ExitUnpivotClause is called when production unpivotClause is exited.
func (s *BaseOracleStatementListener) ExitUnpivotClause(ctx *UnpivotClauseContext) {}

// EnterUnpivotInClause is called when production unpivotInClause is entered.
func (s *BaseOracleStatementListener) EnterUnpivotInClause(ctx *UnpivotInClauseContext) {}

// ExitUnpivotInClause is called when production unpivotInClause is exited.
func (s *BaseOracleStatementListener) ExitUnpivotInClause(ctx *UnpivotInClauseContext) {}

// EnterSampleClause is called when production sampleClause is entered.
func (s *BaseOracleStatementListener) EnterSampleClause(ctx *SampleClauseContext) {}

// ExitSampleClause is called when production sampleClause is exited.
func (s *BaseOracleStatementListener) ExitSampleClause(ctx *SampleClauseContext) {}

// EnterContainersClause is called when production containersClause is entered.
func (s *BaseOracleStatementListener) EnterContainersClause(ctx *ContainersClauseContext) {}

// ExitContainersClause is called when production containersClause is exited.
func (s *BaseOracleStatementListener) ExitContainersClause(ctx *ContainersClauseContext) {}

// EnterShardsClause is called when production shardsClause is entered.
func (s *BaseOracleStatementListener) EnterShardsClause(ctx *ShardsClauseContext) {}

// ExitShardsClause is called when production shardsClause is exited.
func (s *BaseOracleStatementListener) ExitShardsClause(ctx *ShardsClauseContext) {}

// EnterJoinClause is called when production joinClause is entered.
func (s *BaseOracleStatementListener) EnterJoinClause(ctx *JoinClauseContext) {}

// ExitJoinClause is called when production joinClause is exited.
func (s *BaseOracleStatementListener) ExitJoinClause(ctx *JoinClauseContext) {}

// EnterSelectJoinOption is called when production selectJoinOption is entered.
func (s *BaseOracleStatementListener) EnterSelectJoinOption(ctx *SelectJoinOptionContext) {}

// ExitSelectJoinOption is called when production selectJoinOption is exited.
func (s *BaseOracleStatementListener) ExitSelectJoinOption(ctx *SelectJoinOptionContext) {}

// EnterInnerCrossJoinClause is called when production innerCrossJoinClause is entered.
func (s *BaseOracleStatementListener) EnterInnerCrossJoinClause(ctx *InnerCrossJoinClauseContext) {}

// ExitInnerCrossJoinClause is called when production innerCrossJoinClause is exited.
func (s *BaseOracleStatementListener) ExitInnerCrossJoinClause(ctx *InnerCrossJoinClauseContext) {}

// EnterSelectJoinSpecification is called when production selectJoinSpecification is entered.
func (s *BaseOracleStatementListener) EnterSelectJoinSpecification(ctx *SelectJoinSpecificationContext) {
}

// ExitSelectJoinSpecification is called when production selectJoinSpecification is exited.
func (s *BaseOracleStatementListener) ExitSelectJoinSpecification(ctx *SelectJoinSpecificationContext) {
}

// EnterOuterJoinClause is called when production outerJoinClause is entered.
func (s *BaseOracleStatementListener) EnterOuterJoinClause(ctx *OuterJoinClauseContext) {}

// ExitOuterJoinClause is called when production outerJoinClause is exited.
func (s *BaseOracleStatementListener) ExitOuterJoinClause(ctx *OuterJoinClauseContext) {}

// EnterQueryPartitionClause is called when production queryPartitionClause is entered.
func (s *BaseOracleStatementListener) EnterQueryPartitionClause(ctx *QueryPartitionClauseContext) {}

// ExitQueryPartitionClause is called when production queryPartitionClause is exited.
func (s *BaseOracleStatementListener) ExitQueryPartitionClause(ctx *QueryPartitionClauseContext) {}

// EnterOuterJoinType is called when production outerJoinType is entered.
func (s *BaseOracleStatementListener) EnterOuterJoinType(ctx *OuterJoinTypeContext) {}

// ExitOuterJoinType is called when production outerJoinType is exited.
func (s *BaseOracleStatementListener) ExitOuterJoinType(ctx *OuterJoinTypeContext) {}

// EnterCrossOuterApplyClause is called when production crossOuterApplyClause is entered.
func (s *BaseOracleStatementListener) EnterCrossOuterApplyClause(ctx *CrossOuterApplyClauseContext) {}

// ExitCrossOuterApplyClause is called when production crossOuterApplyClause is exited.
func (s *BaseOracleStatementListener) ExitCrossOuterApplyClause(ctx *CrossOuterApplyClauseContext) {}

// EnterInlineAnalyticView is called when production inlineAnalyticView is entered.
func (s *BaseOracleStatementListener) EnterInlineAnalyticView(ctx *InlineAnalyticViewContext) {}

// ExitInlineAnalyticView is called when production inlineAnalyticView is exited.
func (s *BaseOracleStatementListener) ExitInlineAnalyticView(ctx *InlineAnalyticViewContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseOracleStatementListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseOracleStatementListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterHierarchicalQueryClause is called when production hierarchicalQueryClause is entered.
func (s *BaseOracleStatementListener) EnterHierarchicalQueryClause(ctx *HierarchicalQueryClauseContext) {
}

// ExitHierarchicalQueryClause is called when production hierarchicalQueryClause is exited.
func (s *BaseOracleStatementListener) ExitHierarchicalQueryClause(ctx *HierarchicalQueryClauseContext) {
}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseOracleStatementListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseOracleStatementListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterGroupByItem is called when production groupByItem is entered.
func (s *BaseOracleStatementListener) EnterGroupByItem(ctx *GroupByItemContext) {}

// ExitGroupByItem is called when production groupByItem is exited.
func (s *BaseOracleStatementListener) ExitGroupByItem(ctx *GroupByItemContext) {}

// EnterRollupCubeClause is called when production rollupCubeClause is entered.
func (s *BaseOracleStatementListener) EnterRollupCubeClause(ctx *RollupCubeClauseContext) {}

// ExitRollupCubeClause is called when production rollupCubeClause is exited.
func (s *BaseOracleStatementListener) ExitRollupCubeClause(ctx *RollupCubeClauseContext) {}

// EnterGroupingSetsClause is called when production groupingSetsClause is entered.
func (s *BaseOracleStatementListener) EnterGroupingSetsClause(ctx *GroupingSetsClauseContext) {}

// ExitGroupingSetsClause is called when production groupingSetsClause is exited.
func (s *BaseOracleStatementListener) ExitGroupingSetsClause(ctx *GroupingSetsClauseContext) {}

// EnterGroupingExprList is called when production groupingExprList is entered.
func (s *BaseOracleStatementListener) EnterGroupingExprList(ctx *GroupingExprListContext) {}

// ExitGroupingExprList is called when production groupingExprList is exited.
func (s *BaseOracleStatementListener) ExitGroupingExprList(ctx *GroupingExprListContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseOracleStatementListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseOracleStatementListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseOracleStatementListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseOracleStatementListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterModelClause is called when production modelClause is entered.
func (s *BaseOracleStatementListener) EnterModelClause(ctx *ModelClauseContext) {}

// ExitModelClause is called when production modelClause is exited.
func (s *BaseOracleStatementListener) ExitModelClause(ctx *ModelClauseContext) {}

// EnterCellReferenceOptions is called when production cellReferenceOptions is entered.
func (s *BaseOracleStatementListener) EnterCellReferenceOptions(ctx *CellReferenceOptionsContext) {}

// ExitCellReferenceOptions is called when production cellReferenceOptions is exited.
func (s *BaseOracleStatementListener) ExitCellReferenceOptions(ctx *CellReferenceOptionsContext) {}

// EnterReturnRowsClause is called when production returnRowsClause is entered.
func (s *BaseOracleStatementListener) EnterReturnRowsClause(ctx *ReturnRowsClauseContext) {}

// ExitReturnRowsClause is called when production returnRowsClause is exited.
func (s *BaseOracleStatementListener) ExitReturnRowsClause(ctx *ReturnRowsClauseContext) {}

// EnterReferenceModel is called when production referenceModel is entered.
func (s *BaseOracleStatementListener) EnterReferenceModel(ctx *ReferenceModelContext) {}

// ExitReferenceModel is called when production referenceModel is exited.
func (s *BaseOracleStatementListener) ExitReferenceModel(ctx *ReferenceModelContext) {}

// EnterMainModel is called when production mainModel is entered.
func (s *BaseOracleStatementListener) EnterMainModel(ctx *MainModelContext) {}

// ExitMainModel is called when production mainModel is exited.
func (s *BaseOracleStatementListener) ExitMainModel(ctx *MainModelContext) {}

// EnterModelColumnClauses is called when production modelColumnClauses is entered.
func (s *BaseOracleStatementListener) EnterModelColumnClauses(ctx *ModelColumnClausesContext) {}

// ExitModelColumnClauses is called when production modelColumnClauses is exited.
func (s *BaseOracleStatementListener) ExitModelColumnClauses(ctx *ModelColumnClausesContext) {}

// EnterModelRulesClause is called when production modelRulesClause is entered.
func (s *BaseOracleStatementListener) EnterModelRulesClause(ctx *ModelRulesClauseContext) {}

// ExitModelRulesClause is called when production modelRulesClause is exited.
func (s *BaseOracleStatementListener) ExitModelRulesClause(ctx *ModelRulesClauseContext) {}

// EnterModelIterateClause is called when production modelIterateClause is entered.
func (s *BaseOracleStatementListener) EnterModelIterateClause(ctx *ModelIterateClauseContext) {}

// ExitModelIterateClause is called when production modelIterateClause is exited.
func (s *BaseOracleStatementListener) ExitModelIterateClause(ctx *ModelIterateClauseContext) {}

// EnterCellAssignment is called when production cellAssignment is entered.
func (s *BaseOracleStatementListener) EnterCellAssignment(ctx *CellAssignmentContext) {}

// ExitCellAssignment is called when production cellAssignment is exited.
func (s *BaseOracleStatementListener) ExitCellAssignment(ctx *CellAssignmentContext) {}

// EnterSingleColumnForLoop is called when production singleColumnForLoop is entered.
func (s *BaseOracleStatementListener) EnterSingleColumnForLoop(ctx *SingleColumnForLoopContext) {}

// ExitSingleColumnForLoop is called when production singleColumnForLoop is exited.
func (s *BaseOracleStatementListener) ExitSingleColumnForLoop(ctx *SingleColumnForLoopContext) {}

// EnterMultiColumnForLoop is called when production multiColumnForLoop is entered.
func (s *BaseOracleStatementListener) EnterMultiColumnForLoop(ctx *MultiColumnForLoopContext) {}

// ExitMultiColumnForLoop is called when production multiColumnForLoop is exited.
func (s *BaseOracleStatementListener) ExitMultiColumnForLoop(ctx *MultiColumnForLoopContext) {}

// EnterSubquery is called when production subquery is entered.
func (s *BaseOracleStatementListener) EnterSubquery(ctx *SubqueryContext) {}

// ExitSubquery is called when production subquery is exited.
func (s *BaseOracleStatementListener) ExitSubquery(ctx *SubqueryContext) {}

// EnterModelExpr is called when production modelExpr is entered.
func (s *BaseOracleStatementListener) EnterModelExpr(ctx *ModelExprContext) {}

// ExitModelExpr is called when production modelExpr is exited.
func (s *BaseOracleStatementListener) ExitModelExpr(ctx *ModelExprContext) {}

// EnterAnalyticFunction is called when production analyticFunction is entered.
func (s *BaseOracleStatementListener) EnterAnalyticFunction(ctx *AnalyticFunctionContext) {}

// ExitAnalyticFunction is called when production analyticFunction is exited.
func (s *BaseOracleStatementListener) ExitAnalyticFunction(ctx *AnalyticFunctionContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseOracleStatementListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseOracleStatementListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterForUpdateClause is called when production forUpdateClause is entered.
func (s *BaseOracleStatementListener) EnterForUpdateClause(ctx *ForUpdateClauseContext) {}

// ExitForUpdateClause is called when production forUpdateClause is exited.
func (s *BaseOracleStatementListener) ExitForUpdateClause(ctx *ForUpdateClauseContext) {}

// EnterForUpdateClauseList is called when production forUpdateClauseList is entered.
func (s *BaseOracleStatementListener) EnterForUpdateClauseList(ctx *ForUpdateClauseListContext) {}

// ExitForUpdateClauseList is called when production forUpdateClauseList is exited.
func (s *BaseOracleStatementListener) ExitForUpdateClauseList(ctx *ForUpdateClauseListContext) {}

// EnterForUpdateClauseOption is called when production forUpdateClauseOption is entered.
func (s *BaseOracleStatementListener) EnterForUpdateClauseOption(ctx *ForUpdateClauseOptionContext) {}

// ExitForUpdateClauseOption is called when production forUpdateClauseOption is exited.
func (s *BaseOracleStatementListener) ExitForUpdateClauseOption(ctx *ForUpdateClauseOptionContext) {}

// EnterRowLimitingClause is called when production rowLimitingClause is entered.
func (s *BaseOracleStatementListener) EnterRowLimitingClause(ctx *RowLimitingClauseContext) {}

// ExitRowLimitingClause is called when production rowLimitingClause is exited.
func (s *BaseOracleStatementListener) ExitRowLimitingClause(ctx *RowLimitingClauseContext) {}

// EnterMerge is called when production merge is entered.
func (s *BaseOracleStatementListener) EnterMerge(ctx *MergeContext) {}

// ExitMerge is called when production merge is exited.
func (s *BaseOracleStatementListener) ExitMerge(ctx *MergeContext) {}

// EnterHint is called when production hint is entered.
func (s *BaseOracleStatementListener) EnterHint(ctx *HintContext) {}

// ExitHint is called when production hint is exited.
func (s *BaseOracleStatementListener) ExitHint(ctx *HintContext) {}

// EnterIntoClause is called when production intoClause is entered.
func (s *BaseOracleStatementListener) EnterIntoClause(ctx *IntoClauseContext) {}

// ExitIntoClause is called when production intoClause is exited.
func (s *BaseOracleStatementListener) ExitIntoClause(ctx *IntoClauseContext) {}

// EnterUsingClause is called when production usingClause is entered.
func (s *BaseOracleStatementListener) EnterUsingClause(ctx *UsingClauseContext) {}

// ExitUsingClause is called when production usingClause is exited.
func (s *BaseOracleStatementListener) ExitUsingClause(ctx *UsingClauseContext) {}

// EnterMergeUpdateClause is called when production mergeUpdateClause is entered.
func (s *BaseOracleStatementListener) EnterMergeUpdateClause(ctx *MergeUpdateClauseContext) {}

// ExitMergeUpdateClause is called when production mergeUpdateClause is exited.
func (s *BaseOracleStatementListener) ExitMergeUpdateClause(ctx *MergeUpdateClauseContext) {}

// EnterMergeSetAssignmentsClause is called when production mergeSetAssignmentsClause is entered.
func (s *BaseOracleStatementListener) EnterMergeSetAssignmentsClause(ctx *MergeSetAssignmentsClauseContext) {
}

// ExitMergeSetAssignmentsClause is called when production mergeSetAssignmentsClause is exited.
func (s *BaseOracleStatementListener) ExitMergeSetAssignmentsClause(ctx *MergeSetAssignmentsClauseContext) {
}

// EnterMergeAssignment is called when production mergeAssignment is entered.
func (s *BaseOracleStatementListener) EnterMergeAssignment(ctx *MergeAssignmentContext) {}

// ExitMergeAssignment is called when production mergeAssignment is exited.
func (s *BaseOracleStatementListener) ExitMergeAssignment(ctx *MergeAssignmentContext) {}

// EnterMergeAssignmentValue is called when production mergeAssignmentValue is entered.
func (s *BaseOracleStatementListener) EnterMergeAssignmentValue(ctx *MergeAssignmentValueContext) {}

// ExitMergeAssignmentValue is called when production mergeAssignmentValue is exited.
func (s *BaseOracleStatementListener) ExitMergeAssignmentValue(ctx *MergeAssignmentValueContext) {}

// EnterDeleteWhereClause is called when production deleteWhereClause is entered.
func (s *BaseOracleStatementListener) EnterDeleteWhereClause(ctx *DeleteWhereClauseContext) {}

// ExitDeleteWhereClause is called when production deleteWhereClause is exited.
func (s *BaseOracleStatementListener) ExitDeleteWhereClause(ctx *DeleteWhereClauseContext) {}

// EnterMergeInsertClause is called when production mergeInsertClause is entered.
func (s *BaseOracleStatementListener) EnterMergeInsertClause(ctx *MergeInsertClauseContext) {}

// ExitMergeInsertClause is called when production mergeInsertClause is exited.
func (s *BaseOracleStatementListener) ExitMergeInsertClause(ctx *MergeInsertClauseContext) {}

// EnterMergeInsertColumn is called when production mergeInsertColumn is entered.
func (s *BaseOracleStatementListener) EnterMergeInsertColumn(ctx *MergeInsertColumnContext) {}

// ExitMergeInsertColumn is called when production mergeInsertColumn is exited.
func (s *BaseOracleStatementListener) ExitMergeInsertColumn(ctx *MergeInsertColumnContext) {}

// EnterMergeColumnValue is called when production mergeColumnValue is entered.
func (s *BaseOracleStatementListener) EnterMergeColumnValue(ctx *MergeColumnValueContext) {}

// ExitMergeColumnValue is called when production mergeColumnValue is exited.
func (s *BaseOracleStatementListener) ExitMergeColumnValue(ctx *MergeColumnValueContext) {}

// EnterErrorLoggingClause is called when production errorLoggingClause is entered.
func (s *BaseOracleStatementListener) EnterErrorLoggingClause(ctx *ErrorLoggingClauseContext) {}

// ExitErrorLoggingClause is called when production errorLoggingClause is exited.
func (s *BaseOracleStatementListener) ExitErrorLoggingClause(ctx *ErrorLoggingClauseContext) {}

// EnterRowPatternClause is called when production rowPatternClause is entered.
func (s *BaseOracleStatementListener) EnterRowPatternClause(ctx *RowPatternClauseContext) {}

// ExitRowPatternClause is called when production rowPatternClause is exited.
func (s *BaseOracleStatementListener) ExitRowPatternClause(ctx *RowPatternClauseContext) {}

// EnterRowPatternPartitionBy is called when production rowPatternPartitionBy is entered.
func (s *BaseOracleStatementListener) EnterRowPatternPartitionBy(ctx *RowPatternPartitionByContext) {}

// ExitRowPatternPartitionBy is called when production rowPatternPartitionBy is exited.
func (s *BaseOracleStatementListener) ExitRowPatternPartitionBy(ctx *RowPatternPartitionByContext) {}

// EnterRowPatternOrderBy is called when production rowPatternOrderBy is entered.
func (s *BaseOracleStatementListener) EnterRowPatternOrderBy(ctx *RowPatternOrderByContext) {}

// ExitRowPatternOrderBy is called when production rowPatternOrderBy is exited.
func (s *BaseOracleStatementListener) ExitRowPatternOrderBy(ctx *RowPatternOrderByContext) {}

// EnterRowPatternMeasures is called when production rowPatternMeasures is entered.
func (s *BaseOracleStatementListener) EnterRowPatternMeasures(ctx *RowPatternMeasuresContext) {}

// ExitRowPatternMeasures is called when production rowPatternMeasures is exited.
func (s *BaseOracleStatementListener) ExitRowPatternMeasures(ctx *RowPatternMeasuresContext) {}

// EnterRowPatternMeasureColumn is called when production rowPatternMeasureColumn is entered.
func (s *BaseOracleStatementListener) EnterRowPatternMeasureColumn(ctx *RowPatternMeasureColumnContext) {
}

// ExitRowPatternMeasureColumn is called when production rowPatternMeasureColumn is exited.
func (s *BaseOracleStatementListener) ExitRowPatternMeasureColumn(ctx *RowPatternMeasureColumnContext) {
}

// EnterRowPatternRowsPerMatch is called when production rowPatternRowsPerMatch is entered.
func (s *BaseOracleStatementListener) EnterRowPatternRowsPerMatch(ctx *RowPatternRowsPerMatchContext) {
}

// ExitRowPatternRowsPerMatch is called when production rowPatternRowsPerMatch is exited.
func (s *BaseOracleStatementListener) ExitRowPatternRowsPerMatch(ctx *RowPatternRowsPerMatchContext) {
}

// EnterRowPatternSkipTo is called when production rowPatternSkipTo is entered.
func (s *BaseOracleStatementListener) EnterRowPatternSkipTo(ctx *RowPatternSkipToContext) {}

// ExitRowPatternSkipTo is called when production rowPatternSkipTo is exited.
func (s *BaseOracleStatementListener) ExitRowPatternSkipTo(ctx *RowPatternSkipToContext) {}

// EnterRowPattern is called when production rowPattern is entered.
func (s *BaseOracleStatementListener) EnterRowPattern(ctx *RowPatternContext) {}

// ExitRowPattern is called when production rowPattern is exited.
func (s *BaseOracleStatementListener) ExitRowPattern(ctx *RowPatternContext) {}

// EnterRowPatternTerm is called when production rowPatternTerm is entered.
func (s *BaseOracleStatementListener) EnterRowPatternTerm(ctx *RowPatternTermContext) {}

// ExitRowPatternTerm is called when production rowPatternTerm is exited.
func (s *BaseOracleStatementListener) ExitRowPatternTerm(ctx *RowPatternTermContext) {}

// EnterRowPatternFactor is called when production rowPatternFactor is entered.
func (s *BaseOracleStatementListener) EnterRowPatternFactor(ctx *RowPatternFactorContext) {}

// ExitRowPatternFactor is called when production rowPatternFactor is exited.
func (s *BaseOracleStatementListener) ExitRowPatternFactor(ctx *RowPatternFactorContext) {}

// EnterRowPatternPrimary is called when production rowPatternPrimary is entered.
func (s *BaseOracleStatementListener) EnterRowPatternPrimary(ctx *RowPatternPrimaryContext) {}

// ExitRowPatternPrimary is called when production rowPatternPrimary is exited.
func (s *BaseOracleStatementListener) ExitRowPatternPrimary(ctx *RowPatternPrimaryContext) {}

// EnterRowPatternPermute is called when production rowPatternPermute is entered.
func (s *BaseOracleStatementListener) EnterRowPatternPermute(ctx *RowPatternPermuteContext) {}

// ExitRowPatternPermute is called when production rowPatternPermute is exited.
func (s *BaseOracleStatementListener) ExitRowPatternPermute(ctx *RowPatternPermuteContext) {}

// EnterRowPatternQuantifier is called when production rowPatternQuantifier is entered.
func (s *BaseOracleStatementListener) EnterRowPatternQuantifier(ctx *RowPatternQuantifierContext) {}

// ExitRowPatternQuantifier is called when production rowPatternQuantifier is exited.
func (s *BaseOracleStatementListener) ExitRowPatternQuantifier(ctx *RowPatternQuantifierContext) {}

// EnterRowPatternSubsetClause is called when production rowPatternSubsetClause is entered.
func (s *BaseOracleStatementListener) EnterRowPatternSubsetClause(ctx *RowPatternSubsetClauseContext) {
}

// ExitRowPatternSubsetClause is called when production rowPatternSubsetClause is exited.
func (s *BaseOracleStatementListener) ExitRowPatternSubsetClause(ctx *RowPatternSubsetClauseContext) {
}

// EnterRowPatternSubsetItem is called when production rowPatternSubsetItem is entered.
func (s *BaseOracleStatementListener) EnterRowPatternSubsetItem(ctx *RowPatternSubsetItemContext) {}

// ExitRowPatternSubsetItem is called when production rowPatternSubsetItem is exited.
func (s *BaseOracleStatementListener) ExitRowPatternSubsetItem(ctx *RowPatternSubsetItemContext) {}

// EnterRowPatternDefinitionList is called when production rowPatternDefinitionList is entered.
func (s *BaseOracleStatementListener) EnterRowPatternDefinitionList(ctx *RowPatternDefinitionListContext) {
}

// ExitRowPatternDefinitionList is called when production rowPatternDefinitionList is exited.
func (s *BaseOracleStatementListener) ExitRowPatternDefinitionList(ctx *RowPatternDefinitionListContext) {
}

// EnterRowPatternDefinition is called when production rowPatternDefinition is entered.
func (s *BaseOracleStatementListener) EnterRowPatternDefinition(ctx *RowPatternDefinitionContext) {}

// ExitRowPatternDefinition is called when production rowPatternDefinition is exited.
func (s *BaseOracleStatementListener) ExitRowPatternDefinition(ctx *RowPatternDefinitionContext) {}

// EnterRowPatternRecFunc is called when production rowPatternRecFunc is entered.
func (s *BaseOracleStatementListener) EnterRowPatternRecFunc(ctx *RowPatternRecFuncContext) {}

// ExitRowPatternRecFunc is called when production rowPatternRecFunc is exited.
func (s *BaseOracleStatementListener) ExitRowPatternRecFunc(ctx *RowPatternRecFuncContext) {}

// EnterPatternMeasExpression is called when production patternMeasExpression is entered.
func (s *BaseOracleStatementListener) EnterPatternMeasExpression(ctx *PatternMeasExpressionContext) {}

// ExitPatternMeasExpression is called when production patternMeasExpression is exited.
func (s *BaseOracleStatementListener) ExitPatternMeasExpression(ctx *PatternMeasExpressionContext) {}

// EnterRowPatternClassifierFunc is called when production rowPatternClassifierFunc is entered.
func (s *BaseOracleStatementListener) EnterRowPatternClassifierFunc(ctx *RowPatternClassifierFuncContext) {
}

// ExitRowPatternClassifierFunc is called when production rowPatternClassifierFunc is exited.
func (s *BaseOracleStatementListener) ExitRowPatternClassifierFunc(ctx *RowPatternClassifierFuncContext) {
}

// EnterRowPatternMatchNumFunc is called when production rowPatternMatchNumFunc is entered.
func (s *BaseOracleStatementListener) EnterRowPatternMatchNumFunc(ctx *RowPatternMatchNumFuncContext) {
}

// ExitRowPatternMatchNumFunc is called when production rowPatternMatchNumFunc is exited.
func (s *BaseOracleStatementListener) ExitRowPatternMatchNumFunc(ctx *RowPatternMatchNumFuncContext) {
}

// EnterRowPatternNavigationFunc is called when production rowPatternNavigationFunc is entered.
func (s *BaseOracleStatementListener) EnterRowPatternNavigationFunc(ctx *RowPatternNavigationFuncContext) {
}

// ExitRowPatternNavigationFunc is called when production rowPatternNavigationFunc is exited.
func (s *BaseOracleStatementListener) ExitRowPatternNavigationFunc(ctx *RowPatternNavigationFuncContext) {
}

// EnterRowPatternNavLogical is called when production rowPatternNavLogical is entered.
func (s *BaseOracleStatementListener) EnterRowPatternNavLogical(ctx *RowPatternNavLogicalContext) {}

// ExitRowPatternNavLogical is called when production rowPatternNavLogical is exited.
func (s *BaseOracleStatementListener) ExitRowPatternNavLogical(ctx *RowPatternNavLogicalContext) {}

// EnterRowPatternNavPhysical is called when production rowPatternNavPhysical is entered.
func (s *BaseOracleStatementListener) EnterRowPatternNavPhysical(ctx *RowPatternNavPhysicalContext) {}

// ExitRowPatternNavPhysical is called when production rowPatternNavPhysical is exited.
func (s *BaseOracleStatementListener) ExitRowPatternNavPhysical(ctx *RowPatternNavPhysicalContext) {}

// EnterRowPatternNavCompound is called when production rowPatternNavCompound is entered.
func (s *BaseOracleStatementListener) EnterRowPatternNavCompound(ctx *RowPatternNavCompoundContext) {}

// ExitRowPatternNavCompound is called when production rowPatternNavCompound is exited.
func (s *BaseOracleStatementListener) ExitRowPatternNavCompound(ctx *RowPatternNavCompoundContext) {}

// EnterRowPatternAggregateFunc is called when production rowPatternAggregateFunc is entered.
func (s *BaseOracleStatementListener) EnterRowPatternAggregateFunc(ctx *RowPatternAggregateFuncContext) {
}

// ExitRowPatternAggregateFunc is called when production rowPatternAggregateFunc is exited.
func (s *BaseOracleStatementListener) ExitRowPatternAggregateFunc(ctx *RowPatternAggregateFuncContext) {
}

// EnterParameterMarker is called when production parameterMarker is entered.
func (s *BaseOracleStatementListener) EnterParameterMarker(ctx *ParameterMarkerContext) {}

// ExitParameterMarker is called when production parameterMarker is exited.
func (s *BaseOracleStatementListener) ExitParameterMarker(ctx *ParameterMarkerContext) {}

// EnterLiterals is called when production literals is entered.
func (s *BaseOracleStatementListener) EnterLiterals(ctx *LiteralsContext) {}

// ExitLiterals is called when production literals is exited.
func (s *BaseOracleStatementListener) ExitLiterals(ctx *LiteralsContext) {}

// EnterStringLiterals is called when production stringLiterals is entered.
func (s *BaseOracleStatementListener) EnterStringLiterals(ctx *StringLiteralsContext) {}

// ExitStringLiterals is called when production stringLiterals is exited.
func (s *BaseOracleStatementListener) ExitStringLiterals(ctx *StringLiteralsContext) {}

// EnterNumberLiterals is called when production numberLiterals is entered.
func (s *BaseOracleStatementListener) EnterNumberLiterals(ctx *NumberLiteralsContext) {}

// ExitNumberLiterals is called when production numberLiterals is exited.
func (s *BaseOracleStatementListener) ExitNumberLiterals(ctx *NumberLiteralsContext) {}

// EnterDateTimeLiterals is called when production dateTimeLiterals is entered.
func (s *BaseOracleStatementListener) EnterDateTimeLiterals(ctx *DateTimeLiteralsContext) {}

// ExitDateTimeLiterals is called when production dateTimeLiterals is exited.
func (s *BaseOracleStatementListener) ExitDateTimeLiterals(ctx *DateTimeLiteralsContext) {}

// EnterHexadecimalLiterals is called when production hexadecimalLiterals is entered.
func (s *BaseOracleStatementListener) EnterHexadecimalLiterals(ctx *HexadecimalLiteralsContext) {}

// ExitHexadecimalLiterals is called when production hexadecimalLiterals is exited.
func (s *BaseOracleStatementListener) ExitHexadecimalLiterals(ctx *HexadecimalLiteralsContext) {}

// EnterBitValueLiterals is called when production bitValueLiterals is entered.
func (s *BaseOracleStatementListener) EnterBitValueLiterals(ctx *BitValueLiteralsContext) {}

// ExitBitValueLiterals is called when production bitValueLiterals is exited.
func (s *BaseOracleStatementListener) ExitBitValueLiterals(ctx *BitValueLiteralsContext) {}

// EnterBooleanLiterals is called when production booleanLiterals is entered.
func (s *BaseOracleStatementListener) EnterBooleanLiterals(ctx *BooleanLiteralsContext) {}

// ExitBooleanLiterals is called when production booleanLiterals is exited.
func (s *BaseOracleStatementListener) ExitBooleanLiterals(ctx *BooleanLiteralsContext) {}

// EnterNullValueLiterals is called when production nullValueLiterals is entered.
func (s *BaseOracleStatementListener) EnterNullValueLiterals(ctx *NullValueLiteralsContext) {}

// ExitNullValueLiterals is called when production nullValueLiterals is exited.
func (s *BaseOracleStatementListener) ExitNullValueLiterals(ctx *NullValueLiteralsContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseOracleStatementListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseOracleStatementListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterUnreservedWord is called when production unreservedWord is entered.
func (s *BaseOracleStatementListener) EnterUnreservedWord(ctx *UnreservedWordContext) {}

// ExitUnreservedWord is called when production unreservedWord is exited.
func (s *BaseOracleStatementListener) ExitUnreservedWord(ctx *UnreservedWordContext) {}

// EnterSchemaName is called when production schemaName is entered.
func (s *BaseOracleStatementListener) EnterSchemaName(ctx *SchemaNameContext) {}

// ExitSchemaName is called when production schemaName is exited.
func (s *BaseOracleStatementListener) ExitSchemaName(ctx *SchemaNameContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseOracleStatementListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseOracleStatementListener) ExitTableName(ctx *TableNameContext) {}

// EnterViewName is called when production viewName is entered.
func (s *BaseOracleStatementListener) EnterViewName(ctx *ViewNameContext) {}

// ExitViewName is called when production viewName is exited.
func (s *BaseOracleStatementListener) ExitViewName(ctx *ViewNameContext) {}

// EnterMaterializedViewName is called when production materializedViewName is entered.
func (s *BaseOracleStatementListener) EnterMaterializedViewName(ctx *MaterializedViewNameContext) {}

// ExitMaterializedViewName is called when production materializedViewName is exited.
func (s *BaseOracleStatementListener) ExitMaterializedViewName(ctx *MaterializedViewNameContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseOracleStatementListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseOracleStatementListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterObjectName is called when production objectName is entered.
func (s *BaseOracleStatementListener) EnterObjectName(ctx *ObjectNameContext) {}

// ExitObjectName is called when production objectName is exited.
func (s *BaseOracleStatementListener) ExitObjectName(ctx *ObjectNameContext) {}

// EnterClusterName is called when production clusterName is entered.
func (s *BaseOracleStatementListener) EnterClusterName(ctx *ClusterNameContext) {}

// ExitClusterName is called when production clusterName is exited.
func (s *BaseOracleStatementListener) ExitClusterName(ctx *ClusterNameContext) {}

// EnterIndexName is called when production indexName is entered.
func (s *BaseOracleStatementListener) EnterIndexName(ctx *IndexNameContext) {}

// ExitIndexName is called when production indexName is exited.
func (s *BaseOracleStatementListener) ExitIndexName(ctx *IndexNameContext) {}

// EnterStatisticsTypeName is called when production statisticsTypeName is entered.
func (s *BaseOracleStatementListener) EnterStatisticsTypeName(ctx *StatisticsTypeNameContext) {}

// ExitStatisticsTypeName is called when production statisticsTypeName is exited.
func (s *BaseOracleStatementListener) ExitStatisticsTypeName(ctx *StatisticsTypeNameContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseOracleStatementListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseOracleStatementListener) ExitFunction(ctx *FunctionContext) {}

// EnterPackageName is called when production packageName is entered.
func (s *BaseOracleStatementListener) EnterPackageName(ctx *PackageNameContext) {}

// ExitPackageName is called when production packageName is exited.
func (s *BaseOracleStatementListener) ExitPackageName(ctx *PackageNameContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseOracleStatementListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseOracleStatementListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterIndextypeName is called when production indextypeName is entered.
func (s *BaseOracleStatementListener) EnterIndextypeName(ctx *IndextypeNameContext) {}

// ExitIndextypeName is called when production indextypeName is exited.
func (s *BaseOracleStatementListener) ExitIndextypeName(ctx *IndextypeNameContext) {}

// EnterModelName is called when production modelName is entered.
func (s *BaseOracleStatementListener) EnterModelName(ctx *ModelNameContext) {}

// ExitModelName is called when production modelName is exited.
func (s *BaseOracleStatementListener) ExitModelName(ctx *ModelNameContext) {}

// EnterOperatorName is called when production operatorName is entered.
func (s *BaseOracleStatementListener) EnterOperatorName(ctx *OperatorNameContext) {}

// ExitOperatorName is called when production operatorName is exited.
func (s *BaseOracleStatementListener) ExitOperatorName(ctx *OperatorNameContext) {}

// EnterConstraintName is called when production constraintName is entered.
func (s *BaseOracleStatementListener) EnterConstraintName(ctx *ConstraintNameContext) {}

// ExitConstraintName is called when production constraintName is exited.
func (s *BaseOracleStatementListener) ExitConstraintName(ctx *ConstraintNameContext) {}

// EnterSavepointName is called when production savepointName is entered.
func (s *BaseOracleStatementListener) EnterSavepointName(ctx *SavepointNameContext) {}

// ExitSavepointName is called when production savepointName is exited.
func (s *BaseOracleStatementListener) ExitSavepointName(ctx *SavepointNameContext) {}

// EnterSynonymName is called when production synonymName is entered.
func (s *BaseOracleStatementListener) EnterSynonymName(ctx *SynonymNameContext) {}

// ExitSynonymName is called when production synonymName is exited.
func (s *BaseOracleStatementListener) ExitSynonymName(ctx *SynonymNameContext) {}

// EnterOwner is called when production owner is entered.
func (s *BaseOracleStatementListener) EnterOwner(ctx *OwnerContext) {}

// ExitOwner is called when production owner is exited.
func (s *BaseOracleStatementListener) ExitOwner(ctx *OwnerContext) {}

// EnterName is called when production name is entered.
func (s *BaseOracleStatementListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseOracleStatementListener) ExitName(ctx *NameContext) {}

// EnterTablespaceName is called when production tablespaceName is entered.
func (s *BaseOracleStatementListener) EnterTablespaceName(ctx *TablespaceNameContext) {}

// ExitTablespaceName is called when production tablespaceName is exited.
func (s *BaseOracleStatementListener) ExitTablespaceName(ctx *TablespaceNameContext) {}

// EnterTablespaceSetName is called when production tablespaceSetName is entered.
func (s *BaseOracleStatementListener) EnterTablespaceSetName(ctx *TablespaceSetNameContext) {}

// ExitTablespaceSetName is called when production tablespaceSetName is exited.
func (s *BaseOracleStatementListener) ExitTablespaceSetName(ctx *TablespaceSetNameContext) {}

// EnterServiceName is called when production serviceName is entered.
func (s *BaseOracleStatementListener) EnterServiceName(ctx *ServiceNameContext) {}

// ExitServiceName is called when production serviceName is exited.
func (s *BaseOracleStatementListener) ExitServiceName(ctx *ServiceNameContext) {}

// EnterIlmPolicyName is called when production ilmPolicyName is entered.
func (s *BaseOracleStatementListener) EnterIlmPolicyName(ctx *IlmPolicyNameContext) {}

// ExitIlmPolicyName is called when production ilmPolicyName is exited.
func (s *BaseOracleStatementListener) ExitIlmPolicyName(ctx *IlmPolicyNameContext) {}

// EnterPolicyName is called when production policyName is entered.
func (s *BaseOracleStatementListener) EnterPolicyName(ctx *PolicyNameContext) {}

// ExitPolicyName is called when production policyName is exited.
func (s *BaseOracleStatementListener) ExitPolicyName(ctx *PolicyNameContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseOracleStatementListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseOracleStatementListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterDbLink is called when production dbLink is entered.
func (s *BaseOracleStatementListener) EnterDbLink(ctx *DbLinkContext) {}

// ExitDbLink is called when production dbLink is exited.
func (s *BaseOracleStatementListener) ExitDbLink(ctx *DbLinkContext) {}

// EnterParameterValue is called when production parameterValue is entered.
func (s *BaseOracleStatementListener) EnterParameterValue(ctx *ParameterValueContext) {}

// ExitParameterValue is called when production parameterValue is exited.
func (s *BaseOracleStatementListener) ExitParameterValue(ctx *ParameterValueContext) {}

// EnterDirectoryName is called when production directoryName is entered.
func (s *BaseOracleStatementListener) EnterDirectoryName(ctx *DirectoryNameContext) {}

// ExitDirectoryName is called when production directoryName is exited.
func (s *BaseOracleStatementListener) ExitDirectoryName(ctx *DirectoryNameContext) {}

// EnterDispatcherName is called when production dispatcherName is entered.
func (s *BaseOracleStatementListener) EnterDispatcherName(ctx *DispatcherNameContext) {}

// ExitDispatcherName is called when production dispatcherName is exited.
func (s *BaseOracleStatementListener) ExitDispatcherName(ctx *DispatcherNameContext) {}

// EnterClientId is called when production clientId is entered.
func (s *BaseOracleStatementListener) EnterClientId(ctx *ClientIdContext) {}

// ExitClientId is called when production clientId is exited.
func (s *BaseOracleStatementListener) ExitClientId(ctx *ClientIdContext) {}

// EnterOpaqueFormatSpec is called when production opaqueFormatSpec is entered.
func (s *BaseOracleStatementListener) EnterOpaqueFormatSpec(ctx *OpaqueFormatSpecContext) {}

// ExitOpaqueFormatSpec is called when production opaqueFormatSpec is exited.
func (s *BaseOracleStatementListener) ExitOpaqueFormatSpec(ctx *OpaqueFormatSpecContext) {}

// EnterAccessDriverType is called when production accessDriverType is entered.
func (s *BaseOracleStatementListener) EnterAccessDriverType(ctx *AccessDriverTypeContext) {}

// ExitAccessDriverType is called when production accessDriverType is exited.
func (s *BaseOracleStatementListener) ExitAccessDriverType(ctx *AccessDriverTypeContext) {}

// EnterVarrayItem is called when production varrayItem is entered.
func (s *BaseOracleStatementListener) EnterVarrayItem(ctx *VarrayItemContext) {}

// ExitVarrayItem is called when production varrayItem is exited.
func (s *BaseOracleStatementListener) ExitVarrayItem(ctx *VarrayItemContext) {}

// EnterNestedItem is called when production nestedItem is entered.
func (s *BaseOracleStatementListener) EnterNestedItem(ctx *NestedItemContext) {}

// ExitNestedItem is called when production nestedItem is exited.
func (s *BaseOracleStatementListener) ExitNestedItem(ctx *NestedItemContext) {}

// EnterStorageTable is called when production storageTable is entered.
func (s *BaseOracleStatementListener) EnterStorageTable(ctx *StorageTableContext) {}

// ExitStorageTable is called when production storageTable is exited.
func (s *BaseOracleStatementListener) ExitStorageTable(ctx *StorageTableContext) {}

// EnterLobSegname is called when production lobSegname is entered.
func (s *BaseOracleStatementListener) EnterLobSegname(ctx *LobSegnameContext) {}

// ExitLobSegname is called when production lobSegname is exited.
func (s *BaseOracleStatementListener) ExitLobSegname(ctx *LobSegnameContext) {}

// EnterLocationSpecifier is called when production locationSpecifier is entered.
func (s *BaseOracleStatementListener) EnterLocationSpecifier(ctx *LocationSpecifierContext) {}

// ExitLocationSpecifier is called when production locationSpecifier is exited.
func (s *BaseOracleStatementListener) ExitLocationSpecifier(ctx *LocationSpecifierContext) {}

// EnterXmlSchemaURLName is called when production xmlSchemaURLName is entered.
func (s *BaseOracleStatementListener) EnterXmlSchemaURLName(ctx *XmlSchemaURLNameContext) {}

// ExitXmlSchemaURLName is called when production xmlSchemaURLName is exited.
func (s *BaseOracleStatementListener) ExitXmlSchemaURLName(ctx *XmlSchemaURLNameContext) {}

// EnterElementName is called when production elementName is entered.
func (s *BaseOracleStatementListener) EnterElementName(ctx *ElementNameContext) {}

// ExitElementName is called when production elementName is exited.
func (s *BaseOracleStatementListener) ExitElementName(ctx *ElementNameContext) {}

// EnterSubpartitionName is called when production subpartitionName is entered.
func (s *BaseOracleStatementListener) EnterSubpartitionName(ctx *SubpartitionNameContext) {}

// ExitSubpartitionName is called when production subpartitionName is exited.
func (s *BaseOracleStatementListener) ExitSubpartitionName(ctx *SubpartitionNameContext) {}

// EnterParameterName is called when production parameterName is entered.
func (s *BaseOracleStatementListener) EnterParameterName(ctx *ParameterNameContext) {}

// ExitParameterName is called when production parameterName is exited.
func (s *BaseOracleStatementListener) ExitParameterName(ctx *ParameterNameContext) {}

// EnterEditionName is called when production editionName is entered.
func (s *BaseOracleStatementListener) EnterEditionName(ctx *EditionNameContext) {}

// ExitEditionName is called when production editionName is exited.
func (s *BaseOracleStatementListener) ExitEditionName(ctx *EditionNameContext) {}

// EnterContainerName is called when production containerName is entered.
func (s *BaseOracleStatementListener) EnterContainerName(ctx *ContainerNameContext) {}

// ExitContainerName is called when production containerName is exited.
func (s *BaseOracleStatementListener) ExitContainerName(ctx *ContainerNameContext) {}

// EnterPartitionName is called when production partitionName is entered.
func (s *BaseOracleStatementListener) EnterPartitionName(ctx *PartitionNameContext) {}

// ExitPartitionName is called when production partitionName is exited.
func (s *BaseOracleStatementListener) ExitPartitionName(ctx *PartitionNameContext) {}

// EnterPartitionSetName is called when production partitionSetName is entered.
func (s *BaseOracleStatementListener) EnterPartitionSetName(ctx *PartitionSetNameContext) {}

// ExitPartitionSetName is called when production partitionSetName is exited.
func (s *BaseOracleStatementListener) ExitPartitionSetName(ctx *PartitionSetNameContext) {}

// EnterPartitionKeyValue is called when production partitionKeyValue is entered.
func (s *BaseOracleStatementListener) EnterPartitionKeyValue(ctx *PartitionKeyValueContext) {}

// ExitPartitionKeyValue is called when production partitionKeyValue is exited.
func (s *BaseOracleStatementListener) ExitPartitionKeyValue(ctx *PartitionKeyValueContext) {}

// EnterSubpartitionKeyValue is called when production subpartitionKeyValue is entered.
func (s *BaseOracleStatementListener) EnterSubpartitionKeyValue(ctx *SubpartitionKeyValueContext) {}

// ExitSubpartitionKeyValue is called when production subpartitionKeyValue is exited.
func (s *BaseOracleStatementListener) ExitSubpartitionKeyValue(ctx *SubpartitionKeyValueContext) {}

// EnterZonemapName is called when production zonemapName is entered.
func (s *BaseOracleStatementListener) EnterZonemapName(ctx *ZonemapNameContext) {}

// ExitZonemapName is called when production zonemapName is exited.
func (s *BaseOracleStatementListener) ExitZonemapName(ctx *ZonemapNameContext) {}

// EnterFlashbackArchiveName is called when production flashbackArchiveName is entered.
func (s *BaseOracleStatementListener) EnterFlashbackArchiveName(ctx *FlashbackArchiveNameContext) {}

// ExitFlashbackArchiveName is called when production flashbackArchiveName is exited.
func (s *BaseOracleStatementListener) ExitFlashbackArchiveName(ctx *FlashbackArchiveNameContext) {}

// EnterRoleName is called when production roleName is entered.
func (s *BaseOracleStatementListener) EnterRoleName(ctx *RoleNameContext) {}

// ExitRoleName is called when production roleName is exited.
func (s *BaseOracleStatementListener) ExitRoleName(ctx *RoleNameContext) {}

// EnterUserName is called when production userName is entered.
func (s *BaseOracleStatementListener) EnterUserName(ctx *UserNameContext) {}

// ExitUserName is called when production userName is exited.
func (s *BaseOracleStatementListener) ExitUserName(ctx *UserNameContext) {}

// EnterPassword is called when production password is entered.
func (s *BaseOracleStatementListener) EnterPassword(ctx *PasswordContext) {}

// ExitPassword is called when production password is exited.
func (s *BaseOracleStatementListener) ExitPassword(ctx *PasswordContext) {}

// EnterLogGroupName is called when production logGroupName is entered.
func (s *BaseOracleStatementListener) EnterLogGroupName(ctx *LogGroupNameContext) {}

// ExitLogGroupName is called when production logGroupName is exited.
func (s *BaseOracleStatementListener) ExitLogGroupName(ctx *LogGroupNameContext) {}

// EnterColumnNames is called when production columnNames is entered.
func (s *BaseOracleStatementListener) EnterColumnNames(ctx *ColumnNamesContext) {}

// ExitColumnNames is called when production columnNames is exited.
func (s *BaseOracleStatementListener) ExitColumnNames(ctx *ColumnNamesContext) {}

// EnterTableNames is called when production tableNames is entered.
func (s *BaseOracleStatementListener) EnterTableNames(ctx *TableNamesContext) {}

// ExitTableNames is called when production tableNames is exited.
func (s *BaseOracleStatementListener) ExitTableNames(ctx *TableNamesContext) {}

// EnterOracleId is called when production oracleId is entered.
func (s *BaseOracleStatementListener) EnterOracleId(ctx *OracleIdContext) {}

// ExitOracleId is called when production oracleId is exited.
func (s *BaseOracleStatementListener) ExitOracleId(ctx *OracleIdContext) {}

// EnterCollationName is called when production collationName is entered.
func (s *BaseOracleStatementListener) EnterCollationName(ctx *CollationNameContext) {}

// ExitCollationName is called when production collationName is exited.
func (s *BaseOracleStatementListener) ExitCollationName(ctx *CollationNameContext) {}

// EnterColumnCollationName is called when production columnCollationName is entered.
func (s *BaseOracleStatementListener) EnterColumnCollationName(ctx *ColumnCollationNameContext) {}

// ExitColumnCollationName is called when production columnCollationName is exited.
func (s *BaseOracleStatementListener) ExitColumnCollationName(ctx *ColumnCollationNameContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseOracleStatementListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseOracleStatementListener) ExitAlias(ctx *AliasContext) {}

// EnterDataTypeLength is called when production dataTypeLength is entered.
func (s *BaseOracleStatementListener) EnterDataTypeLength(ctx *DataTypeLengthContext) {}

// ExitDataTypeLength is called when production dataTypeLength is exited.
func (s *BaseOracleStatementListener) ExitDataTypeLength(ctx *DataTypeLengthContext) {}

// EnterPrimaryKey is called when production primaryKey is entered.
func (s *BaseOracleStatementListener) EnterPrimaryKey(ctx *PrimaryKeyContext) {}

// ExitPrimaryKey is called when production primaryKey is exited.
func (s *BaseOracleStatementListener) ExitPrimaryKey(ctx *PrimaryKeyContext) {}

// EnterExprs is called when production exprs is entered.
func (s *BaseOracleStatementListener) EnterExprs(ctx *ExprsContext) {}

// ExitExprs is called when production exprs is exited.
func (s *BaseOracleStatementListener) ExitExprs(ctx *ExprsContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseOracleStatementListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseOracleStatementListener) ExitExprList(ctx *ExprListContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseOracleStatementListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseOracleStatementListener) ExitExpr(ctx *ExprContext) {}

// EnterAndOperator is called when production andOperator is entered.
func (s *BaseOracleStatementListener) EnterAndOperator(ctx *AndOperatorContext) {}

// ExitAndOperator is called when production andOperator is exited.
func (s *BaseOracleStatementListener) ExitAndOperator(ctx *AndOperatorContext) {}

// EnterOrOperator is called when production orOperator is entered.
func (s *BaseOracleStatementListener) EnterOrOperator(ctx *OrOperatorContext) {}

// ExitOrOperator is called when production orOperator is exited.
func (s *BaseOracleStatementListener) ExitOrOperator(ctx *OrOperatorContext) {}

// EnterNotOperator is called when production notOperator is entered.
func (s *BaseOracleStatementListener) EnterNotOperator(ctx *NotOperatorContext) {}

// ExitNotOperator is called when production notOperator is exited.
func (s *BaseOracleStatementListener) ExitNotOperator(ctx *NotOperatorContext) {}

// EnterBooleanPrimary is called when production booleanPrimary is entered.
func (s *BaseOracleStatementListener) EnterBooleanPrimary(ctx *BooleanPrimaryContext) {}

// ExitBooleanPrimary is called when production booleanPrimary is exited.
func (s *BaseOracleStatementListener) ExitBooleanPrimary(ctx *BooleanPrimaryContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseOracleStatementListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseOracleStatementListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseOracleStatementListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseOracleStatementListener) ExitPredicate(ctx *PredicateContext) {}

// EnterBitExpr is called when production bitExpr is entered.
func (s *BaseOracleStatementListener) EnterBitExpr(ctx *BitExprContext) {}

// ExitBitExpr is called when production bitExpr is exited.
func (s *BaseOracleStatementListener) ExitBitExpr(ctx *BitExprContext) {}

// EnterSimpleExpr is called when production simpleExpr is entered.
func (s *BaseOracleStatementListener) EnterSimpleExpr(ctx *SimpleExprContext) {}

// ExitSimpleExpr is called when production simpleExpr is exited.
func (s *BaseOracleStatementListener) ExitSimpleExpr(ctx *SimpleExprContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseOracleStatementListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseOracleStatementListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterAggregationFunction is called when production aggregationFunction is entered.
func (s *BaseOracleStatementListener) EnterAggregationFunction(ctx *AggregationFunctionContext) {}

// ExitAggregationFunction is called when production aggregationFunction is exited.
func (s *BaseOracleStatementListener) ExitAggregationFunction(ctx *AggregationFunctionContext) {}

// EnterAggregationFunctionName is called when production aggregationFunctionName is entered.
func (s *BaseOracleStatementListener) EnterAggregationFunctionName(ctx *AggregationFunctionNameContext) {
}

// ExitAggregationFunctionName is called when production aggregationFunctionName is exited.
func (s *BaseOracleStatementListener) ExitAggregationFunctionName(ctx *AggregationFunctionNameContext) {
}

// EnterAnalyticClause is called when production analyticClause is entered.
func (s *BaseOracleStatementListener) EnterAnalyticClause(ctx *AnalyticClauseContext) {}

// ExitAnalyticClause is called when production analyticClause is exited.
func (s *BaseOracleStatementListener) ExitAnalyticClause(ctx *AnalyticClauseContext) {}

// EnterWindowingClause is called when production windowingClause is entered.
func (s *BaseOracleStatementListener) EnterWindowingClause(ctx *WindowingClauseContext) {}

// ExitWindowingClause is called when production windowingClause is exited.
func (s *BaseOracleStatementListener) ExitWindowingClause(ctx *WindowingClauseContext) {}

// EnterSpecialFunction is called when production specialFunction is entered.
func (s *BaseOracleStatementListener) EnterSpecialFunction(ctx *SpecialFunctionContext) {}

// ExitSpecialFunction is called when production specialFunction is exited.
func (s *BaseOracleStatementListener) ExitSpecialFunction(ctx *SpecialFunctionContext) {}

// EnterCastFunction is called when production castFunction is entered.
func (s *BaseOracleStatementListener) EnterCastFunction(ctx *CastFunctionContext) {}

// ExitCastFunction is called when production castFunction is exited.
func (s *BaseOracleStatementListener) ExitCastFunction(ctx *CastFunctionContext) {}

// EnterCharFunction is called when production charFunction is entered.
func (s *BaseOracleStatementListener) EnterCharFunction(ctx *CharFunctionContext) {}

// ExitCharFunction is called when production charFunction is exited.
func (s *BaseOracleStatementListener) ExitCharFunction(ctx *CharFunctionContext) {}

// EnterRegularFunction is called when production regularFunction is entered.
func (s *BaseOracleStatementListener) EnterRegularFunction(ctx *RegularFunctionContext) {}

// ExitRegularFunction is called when production regularFunction is exited.
func (s *BaseOracleStatementListener) ExitRegularFunction(ctx *RegularFunctionContext) {}

// EnterRegularFunctionName is called when production regularFunctionName is entered.
func (s *BaseOracleStatementListener) EnterRegularFunctionName(ctx *RegularFunctionNameContext) {}

// ExitRegularFunctionName is called when production regularFunctionName is exited.
func (s *BaseOracleStatementListener) ExitRegularFunctionName(ctx *RegularFunctionNameContext) {}

// EnterCaseExpression is called when production caseExpression is entered.
func (s *BaseOracleStatementListener) EnterCaseExpression(ctx *CaseExpressionContext) {}

// ExitCaseExpression is called when production caseExpression is exited.
func (s *BaseOracleStatementListener) ExitCaseExpression(ctx *CaseExpressionContext) {}

// EnterCaseWhen is called when production caseWhen is entered.
func (s *BaseOracleStatementListener) EnterCaseWhen(ctx *CaseWhenContext) {}

// ExitCaseWhen is called when production caseWhen is exited.
func (s *BaseOracleStatementListener) ExitCaseWhen(ctx *CaseWhenContext) {}

// EnterCaseElse is called when production caseElse is entered.
func (s *BaseOracleStatementListener) EnterCaseElse(ctx *CaseElseContext) {}

// ExitCaseElse is called when production caseElse is exited.
func (s *BaseOracleStatementListener) ExitCaseElse(ctx *CaseElseContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseOracleStatementListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseOracleStatementListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderByItem is called when production orderByItem is entered.
func (s *BaseOracleStatementListener) EnterOrderByItem(ctx *OrderByItemContext) {}

// ExitOrderByItem is called when production orderByItem is exited.
func (s *BaseOracleStatementListener) ExitOrderByItem(ctx *OrderByItemContext) {}

// EnterAttributeName is called when production attributeName is entered.
func (s *BaseOracleStatementListener) EnterAttributeName(ctx *AttributeNameContext) {}

// ExitAttributeName is called when production attributeName is exited.
func (s *BaseOracleStatementListener) ExitAttributeName(ctx *AttributeNameContext) {}

// EnterSimpleExprs is called when production simpleExprs is entered.
func (s *BaseOracleStatementListener) EnterSimpleExprs(ctx *SimpleExprsContext) {}

// ExitSimpleExprs is called when production simpleExprs is exited.
func (s *BaseOracleStatementListener) ExitSimpleExprs(ctx *SimpleExprsContext) {}

// EnterLobItem is called when production lobItem is entered.
func (s *BaseOracleStatementListener) EnterLobItem(ctx *LobItemContext) {}

// ExitLobItem is called when production lobItem is exited.
func (s *BaseOracleStatementListener) ExitLobItem(ctx *LobItemContext) {}

// EnterLobItems is called when production lobItems is entered.
func (s *BaseOracleStatementListener) EnterLobItems(ctx *LobItemsContext) {}

// ExitLobItems is called when production lobItems is exited.
func (s *BaseOracleStatementListener) ExitLobItems(ctx *LobItemsContext) {}

// EnterLobItemList is called when production lobItemList is entered.
func (s *BaseOracleStatementListener) EnterLobItemList(ctx *LobItemListContext) {}

// ExitLobItemList is called when production lobItemList is exited.
func (s *BaseOracleStatementListener) ExitLobItemList(ctx *LobItemListContext) {}

// EnterDataType is called when production dataType is entered.
func (s *BaseOracleStatementListener) EnterDataType(ctx *DataTypeContext) {}

// ExitDataType is called when production dataType is exited.
func (s *BaseOracleStatementListener) ExitDataType(ctx *DataTypeContext) {}

// EnterSpecialDatatype is called when production specialDatatype is entered.
func (s *BaseOracleStatementListener) EnterSpecialDatatype(ctx *SpecialDatatypeContext) {}

// ExitSpecialDatatype is called when production specialDatatype is exited.
func (s *BaseOracleStatementListener) ExitSpecialDatatype(ctx *SpecialDatatypeContext) {}

// EnterDataTypeName is called when production dataTypeName is entered.
func (s *BaseOracleStatementListener) EnterDataTypeName(ctx *DataTypeNameContext) {}

// ExitDataTypeName is called when production dataTypeName is exited.
func (s *BaseOracleStatementListener) ExitDataTypeName(ctx *DataTypeNameContext) {}

// EnterDatetimeTypeSuffix is called when production datetimeTypeSuffix is entered.
func (s *BaseOracleStatementListener) EnterDatetimeTypeSuffix(ctx *DatetimeTypeSuffixContext) {}

// ExitDatetimeTypeSuffix is called when production datetimeTypeSuffix is exited.
func (s *BaseOracleStatementListener) ExitDatetimeTypeSuffix(ctx *DatetimeTypeSuffixContext) {}

// EnterTreatFunction is called when production treatFunction is entered.
func (s *BaseOracleStatementListener) EnterTreatFunction(ctx *TreatFunctionContext) {}

// ExitTreatFunction is called when production treatFunction is exited.
func (s *BaseOracleStatementListener) ExitTreatFunction(ctx *TreatFunctionContext) {}

// EnterPrivateExprOfDb is called when production privateExprOfDb is entered.
func (s *BaseOracleStatementListener) EnterPrivateExprOfDb(ctx *PrivateExprOfDbContext) {}

// ExitPrivateExprOfDb is called when production privateExprOfDb is exited.
func (s *BaseOracleStatementListener) ExitPrivateExprOfDb(ctx *PrivateExprOfDbContext) {}

// EnterCaseExpr is called when production caseExpr is entered.
func (s *BaseOracleStatementListener) EnterCaseExpr(ctx *CaseExprContext) {}

// ExitCaseExpr is called when production caseExpr is exited.
func (s *BaseOracleStatementListener) ExitCaseExpr(ctx *CaseExprContext) {}

// EnterSimpleCaseExpr is called when production simpleCaseExpr is entered.
func (s *BaseOracleStatementListener) EnterSimpleCaseExpr(ctx *SimpleCaseExprContext) {}

// ExitSimpleCaseExpr is called when production simpleCaseExpr is exited.
func (s *BaseOracleStatementListener) ExitSimpleCaseExpr(ctx *SimpleCaseExprContext) {}

// EnterSearchedCaseExpr is called when production searchedCaseExpr is entered.
func (s *BaseOracleStatementListener) EnterSearchedCaseExpr(ctx *SearchedCaseExprContext) {}

// ExitSearchedCaseExpr is called when production searchedCaseExpr is exited.
func (s *BaseOracleStatementListener) ExitSearchedCaseExpr(ctx *SearchedCaseExprContext) {}

// EnterElseClause is called when production elseClause is entered.
func (s *BaseOracleStatementListener) EnterElseClause(ctx *ElseClauseContext) {}

// ExitElseClause is called when production elseClause is exited.
func (s *BaseOracleStatementListener) ExitElseClause(ctx *ElseClauseContext) {}

// EnterIntervalExpression is called when production intervalExpression is entered.
func (s *BaseOracleStatementListener) EnterIntervalExpression(ctx *IntervalExpressionContext) {}

// ExitIntervalExpression is called when production intervalExpression is exited.
func (s *BaseOracleStatementListener) ExitIntervalExpression(ctx *IntervalExpressionContext) {}

// EnterObjectAccessExpression is called when production objectAccessExpression is entered.
func (s *BaseOracleStatementListener) EnterObjectAccessExpression(ctx *ObjectAccessExpressionContext) {
}

// ExitObjectAccessExpression is called when production objectAccessExpression is exited.
func (s *BaseOracleStatementListener) ExitObjectAccessExpression(ctx *ObjectAccessExpressionContext) {
}

// EnterConstructorExpr is called when production constructorExpr is entered.
func (s *BaseOracleStatementListener) EnterConstructorExpr(ctx *ConstructorExprContext) {}

// ExitConstructorExpr is called when production constructorExpr is exited.
func (s *BaseOracleStatementListener) ExitConstructorExpr(ctx *ConstructorExprContext) {}

// EnterIgnoredIdentifier is called when production ignoredIdentifier is entered.
func (s *BaseOracleStatementListener) EnterIgnoredIdentifier(ctx *IgnoredIdentifierContext) {}

// ExitIgnoredIdentifier is called when production ignoredIdentifier is exited.
func (s *BaseOracleStatementListener) ExitIgnoredIdentifier(ctx *IgnoredIdentifierContext) {}

// EnterIgnoredIdentifiers is called when production ignoredIdentifiers is entered.
func (s *BaseOracleStatementListener) EnterIgnoredIdentifiers(ctx *IgnoredIdentifiersContext) {}

// ExitIgnoredIdentifiers is called when production ignoredIdentifiers is exited.
func (s *BaseOracleStatementListener) ExitIgnoredIdentifiers(ctx *IgnoredIdentifiersContext) {}

// EnterMatchNone is called when production matchNone is entered.
func (s *BaseOracleStatementListener) EnterMatchNone(ctx *MatchNoneContext) {}

// ExitMatchNone is called when production matchNone is exited.
func (s *BaseOracleStatementListener) ExitMatchNone(ctx *MatchNoneContext) {}

// EnterHashSubpartitionQuantity is called when production hashSubpartitionQuantity is entered.
func (s *BaseOracleStatementListener) EnterHashSubpartitionQuantity(ctx *HashSubpartitionQuantityContext) {
}

// ExitHashSubpartitionQuantity is called when production hashSubpartitionQuantity is exited.
func (s *BaseOracleStatementListener) ExitHashSubpartitionQuantity(ctx *HashSubpartitionQuantityContext) {
}

// EnterOdciParameters is called when production odciParameters is entered.
func (s *BaseOracleStatementListener) EnterOdciParameters(ctx *OdciParametersContext) {}

// ExitOdciParameters is called when production odciParameters is exited.
func (s *BaseOracleStatementListener) ExitOdciParameters(ctx *OdciParametersContext) {}

// EnterDatabaseName is called when production databaseName is entered.
func (s *BaseOracleStatementListener) EnterDatabaseName(ctx *DatabaseNameContext) {}

// ExitDatabaseName is called when production databaseName is exited.
func (s *BaseOracleStatementListener) ExitDatabaseName(ctx *DatabaseNameContext) {}

// EnterLocationName is called when production locationName is entered.
func (s *BaseOracleStatementListener) EnterLocationName(ctx *LocationNameContext) {}

// ExitLocationName is called when production locationName is exited.
func (s *BaseOracleStatementListener) ExitLocationName(ctx *LocationNameContext) {}

// EnterFileName is called when production fileName is entered.
func (s *BaseOracleStatementListener) EnterFileName(ctx *FileNameContext) {}

// ExitFileName is called when production fileName is exited.
func (s *BaseOracleStatementListener) ExitFileName(ctx *FileNameContext) {}

// EnterAsmFileName is called when production asmFileName is entered.
func (s *BaseOracleStatementListener) EnterAsmFileName(ctx *AsmFileNameContext) {}

// ExitAsmFileName is called when production asmFileName is exited.
func (s *BaseOracleStatementListener) ExitAsmFileName(ctx *AsmFileNameContext) {}

// EnterFileNumber is called when production fileNumber is entered.
func (s *BaseOracleStatementListener) EnterFileNumber(ctx *FileNumberContext) {}

// ExitFileNumber is called when production fileNumber is exited.
func (s *BaseOracleStatementListener) ExitFileNumber(ctx *FileNumberContext) {}

// EnterInstanceName is called when production instanceName is entered.
func (s *BaseOracleStatementListener) EnterInstanceName(ctx *InstanceNameContext) {}

// ExitInstanceName is called when production instanceName is exited.
func (s *BaseOracleStatementListener) ExitInstanceName(ctx *InstanceNameContext) {}

// EnterLogminerSessionName is called when production logminerSessionName is entered.
func (s *BaseOracleStatementListener) EnterLogminerSessionName(ctx *LogminerSessionNameContext) {}

// ExitLogminerSessionName is called when production logminerSessionName is exited.
func (s *BaseOracleStatementListener) ExitLogminerSessionName(ctx *LogminerSessionNameContext) {}

// EnterTablespaceGroupName is called when production tablespaceGroupName is entered.
func (s *BaseOracleStatementListener) EnterTablespaceGroupName(ctx *TablespaceGroupNameContext) {}

// ExitTablespaceGroupName is called when production tablespaceGroupName is exited.
func (s *BaseOracleStatementListener) ExitTablespaceGroupName(ctx *TablespaceGroupNameContext) {}

// EnterCopyName is called when production copyName is entered.
func (s *BaseOracleStatementListener) EnterCopyName(ctx *CopyNameContext) {}

// ExitCopyName is called when production copyName is exited.
func (s *BaseOracleStatementListener) ExitCopyName(ctx *CopyNameContext) {}

// EnterMirrorName is called when production mirrorName is entered.
func (s *BaseOracleStatementListener) EnterMirrorName(ctx *MirrorNameContext) {}

// ExitMirrorName is called when production mirrorName is exited.
func (s *BaseOracleStatementListener) ExitMirrorName(ctx *MirrorNameContext) {}

// EnterUriString is called when production uriString is entered.
func (s *BaseOracleStatementListener) EnterUriString(ctx *UriStringContext) {}

// ExitUriString is called when production uriString is exited.
func (s *BaseOracleStatementListener) ExitUriString(ctx *UriStringContext) {}

// EnterQualifiedCredentialName is called when production qualifiedCredentialName is entered.
func (s *BaseOracleStatementListener) EnterQualifiedCredentialName(ctx *QualifiedCredentialNameContext) {
}

// ExitQualifiedCredentialName is called when production qualifiedCredentialName is exited.
func (s *BaseOracleStatementListener) ExitQualifiedCredentialName(ctx *QualifiedCredentialNameContext) {
}

// EnterPdbName is called when production pdbName is entered.
func (s *BaseOracleStatementListener) EnterPdbName(ctx *PdbNameContext) {}

// ExitPdbName is called when production pdbName is exited.
func (s *BaseOracleStatementListener) ExitPdbName(ctx *PdbNameContext) {}

// EnterDiskgroupName is called when production diskgroupName is entered.
func (s *BaseOracleStatementListener) EnterDiskgroupName(ctx *DiskgroupNameContext) {}

// ExitDiskgroupName is called when production diskgroupName is exited.
func (s *BaseOracleStatementListener) ExitDiskgroupName(ctx *DiskgroupNameContext) {}

// EnterTemplateName is called when production templateName is entered.
func (s *BaseOracleStatementListener) EnterTemplateName(ctx *TemplateNameContext) {}

// ExitTemplateName is called when production templateName is exited.
func (s *BaseOracleStatementListener) ExitTemplateName(ctx *TemplateNameContext) {}

// EnterAliasName is called when production aliasName is entered.
func (s *BaseOracleStatementListener) EnterAliasName(ctx *AliasNameContext) {}

// ExitAliasName is called when production aliasName is exited.
func (s *BaseOracleStatementListener) ExitAliasName(ctx *AliasNameContext) {}

// EnterDomain is called when production domain is entered.
func (s *BaseOracleStatementListener) EnterDomain(ctx *DomainContext) {}

// ExitDomain is called when production domain is exited.
func (s *BaseOracleStatementListener) ExitDomain(ctx *DomainContext) {}

// EnterDateValue is called when production dateValue is entered.
func (s *BaseOracleStatementListener) EnterDateValue(ctx *DateValueContext) {}

// ExitDateValue is called when production dateValue is exited.
func (s *BaseOracleStatementListener) ExitDateValue(ctx *DateValueContext) {}

// EnterSessionId is called when production sessionId is entered.
func (s *BaseOracleStatementListener) EnterSessionId(ctx *SessionIdContext) {}

// ExitSessionId is called when production sessionId is exited.
func (s *BaseOracleStatementListener) ExitSessionId(ctx *SessionIdContext) {}

// EnterSerialNumber is called when production serialNumber is entered.
func (s *BaseOracleStatementListener) EnterSerialNumber(ctx *SerialNumberContext) {}

// ExitSerialNumber is called when production serialNumber is exited.
func (s *BaseOracleStatementListener) ExitSerialNumber(ctx *SerialNumberContext) {}

// EnterInstanceId is called when production instanceId is entered.
func (s *BaseOracleStatementListener) EnterInstanceId(ctx *InstanceIdContext) {}

// ExitInstanceId is called when production instanceId is exited.
func (s *BaseOracleStatementListener) ExitInstanceId(ctx *InstanceIdContext) {}

// EnterSqlId is called when production sqlId is entered.
func (s *BaseOracleStatementListener) EnterSqlId(ctx *SqlIdContext) {}

// ExitSqlId is called when production sqlId is exited.
func (s *BaseOracleStatementListener) ExitSqlId(ctx *SqlIdContext) {}

// EnterLogFileName is called when production logFileName is entered.
func (s *BaseOracleStatementListener) EnterLogFileName(ctx *LogFileNameContext) {}

// ExitLogFileName is called when production logFileName is exited.
func (s *BaseOracleStatementListener) ExitLogFileName(ctx *LogFileNameContext) {}

// EnterLogFileGroupsArchivedLocationName is called when production logFileGroupsArchivedLocationName is entered.
func (s *BaseOracleStatementListener) EnterLogFileGroupsArchivedLocationName(ctx *LogFileGroupsArchivedLocationNameContext) {
}

// ExitLogFileGroupsArchivedLocationName is called when production logFileGroupsArchivedLocationName is exited.
func (s *BaseOracleStatementListener) ExitLogFileGroupsArchivedLocationName(ctx *LogFileGroupsArchivedLocationNameContext) {
}

// EnterAsmVersion is called when production asmVersion is entered.
func (s *BaseOracleStatementListener) EnterAsmVersion(ctx *AsmVersionContext) {}

// ExitAsmVersion is called when production asmVersion is exited.
func (s *BaseOracleStatementListener) ExitAsmVersion(ctx *AsmVersionContext) {}

// EnterWalletPassword is called when production walletPassword is entered.
func (s *BaseOracleStatementListener) EnterWalletPassword(ctx *WalletPasswordContext) {}

// ExitWalletPassword is called when production walletPassword is exited.
func (s *BaseOracleStatementListener) ExitWalletPassword(ctx *WalletPasswordContext) {}

// EnterHsmAuthString is called when production hsmAuthString is entered.
func (s *BaseOracleStatementListener) EnterHsmAuthString(ctx *HsmAuthStringContext) {}

// ExitHsmAuthString is called when production hsmAuthString is exited.
func (s *BaseOracleStatementListener) ExitHsmAuthString(ctx *HsmAuthStringContext) {}

// EnterTargetDbName is called when production targetDbName is entered.
func (s *BaseOracleStatementListener) EnterTargetDbName(ctx *TargetDbNameContext) {}

// ExitTargetDbName is called when production targetDbName is exited.
func (s *BaseOracleStatementListener) ExitTargetDbName(ctx *TargetDbNameContext) {}

// EnterCertificateId is called when production certificateId is entered.
func (s *BaseOracleStatementListener) EnterCertificateId(ctx *CertificateIdContext) {}

// ExitCertificateId is called when production certificateId is exited.
func (s *BaseOracleStatementListener) ExitCertificateId(ctx *CertificateIdContext) {}

// EnterCategoryName is called when production categoryName is entered.
func (s *BaseOracleStatementListener) EnterCategoryName(ctx *CategoryNameContext) {}

// ExitCategoryName is called when production categoryName is exited.
func (s *BaseOracleStatementListener) ExitCategoryName(ctx *CategoryNameContext) {}

// EnterOffset is called when production offset is entered.
func (s *BaseOracleStatementListener) EnterOffset(ctx *OffsetContext) {}

// ExitOffset is called when production offset is exited.
func (s *BaseOracleStatementListener) ExitOffset(ctx *OffsetContext) {}

// EnterRowcount is called when production rowcount is entered.
func (s *BaseOracleStatementListener) EnterRowcount(ctx *RowcountContext) {}

// ExitRowcount is called when production rowcount is exited.
func (s *BaseOracleStatementListener) ExitRowcount(ctx *RowcountContext) {}

// EnterPercent is called when production percent is entered.
func (s *BaseOracleStatementListener) EnterPercent(ctx *PercentContext) {}

// ExitPercent is called when production percent is exited.
func (s *BaseOracleStatementListener) ExitPercent(ctx *PercentContext) {}

// EnterRollbackSegment is called when production rollbackSegment is entered.
func (s *BaseOracleStatementListener) EnterRollbackSegment(ctx *RollbackSegmentContext) {}

// ExitRollbackSegment is called when production rollbackSegment is exited.
func (s *BaseOracleStatementListener) ExitRollbackSegment(ctx *RollbackSegmentContext) {}

// EnterQueryName is called when production queryName is entered.
func (s *BaseOracleStatementListener) EnterQueryName(ctx *QueryNameContext) {}

// ExitQueryName is called when production queryName is exited.
func (s *BaseOracleStatementListener) ExitQueryName(ctx *QueryNameContext) {}

// EnterCycleValue is called when production cycleValue is entered.
func (s *BaseOracleStatementListener) EnterCycleValue(ctx *CycleValueContext) {}

// ExitCycleValue is called when production cycleValue is exited.
func (s *BaseOracleStatementListener) ExitCycleValue(ctx *CycleValueContext) {}

// EnterNoCycleValue is called when production noCycleValue is entered.
func (s *BaseOracleStatementListener) EnterNoCycleValue(ctx *NoCycleValueContext) {}

// ExitNoCycleValue is called when production noCycleValue is exited.
func (s *BaseOracleStatementListener) ExitNoCycleValue(ctx *NoCycleValueContext) {}

// EnterOrderingColumn is called when production orderingColumn is entered.
func (s *BaseOracleStatementListener) EnterOrderingColumn(ctx *OrderingColumnContext) {}

// ExitOrderingColumn is called when production orderingColumn is exited.
func (s *BaseOracleStatementListener) ExitOrderingColumn(ctx *OrderingColumnContext) {}

// EnterSubavName is called when production subavName is entered.
func (s *BaseOracleStatementListener) EnterSubavName(ctx *SubavNameContext) {}

// ExitSubavName is called when production subavName is exited.
func (s *BaseOracleStatementListener) ExitSubavName(ctx *SubavNameContext) {}

// EnterBaseAvName is called when production baseAvName is entered.
func (s *BaseOracleStatementListener) EnterBaseAvName(ctx *BaseAvNameContext) {}

// ExitBaseAvName is called when production baseAvName is exited.
func (s *BaseOracleStatementListener) ExitBaseAvName(ctx *BaseAvNameContext) {}

// EnterMeasName is called when production measName is entered.
func (s *BaseOracleStatementListener) EnterMeasName(ctx *MeasNameContext) {}

// ExitMeasName is called when production measName is exited.
func (s *BaseOracleStatementListener) ExitMeasName(ctx *MeasNameContext) {}

// EnterLevelRef is called when production levelRef is entered.
func (s *BaseOracleStatementListener) EnterLevelRef(ctx *LevelRefContext) {}

// ExitLevelRef is called when production levelRef is exited.
func (s *BaseOracleStatementListener) ExitLevelRef(ctx *LevelRefContext) {}

// EnterOffsetExpr is called when production offsetExpr is entered.
func (s *BaseOracleStatementListener) EnterOffsetExpr(ctx *OffsetExprContext) {}

// ExitOffsetExpr is called when production offsetExpr is exited.
func (s *BaseOracleStatementListener) ExitOffsetExpr(ctx *OffsetExprContext) {}

// EnterMemberKeyExpr is called when production memberKeyExpr is entered.
func (s *BaseOracleStatementListener) EnterMemberKeyExpr(ctx *MemberKeyExprContext) {}

// ExitMemberKeyExpr is called when production memberKeyExpr is exited.
func (s *BaseOracleStatementListener) ExitMemberKeyExpr(ctx *MemberKeyExprContext) {}

// EnterDepthExpression is called when production depthExpression is entered.
func (s *BaseOracleStatementListener) EnterDepthExpression(ctx *DepthExpressionContext) {}

// ExitDepthExpression is called when production depthExpression is exited.
func (s *BaseOracleStatementListener) ExitDepthExpression(ctx *DepthExpressionContext) {}

// EnterUnitName is called when production unitName is entered.
func (s *BaseOracleStatementListener) EnterUnitName(ctx *UnitNameContext) {}

// ExitUnitName is called when production unitName is exited.
func (s *BaseOracleStatementListener) ExitUnitName(ctx *UnitNameContext) {}

// EnterProcedureName is called when production procedureName is entered.
func (s *BaseOracleStatementListener) EnterProcedureName(ctx *ProcedureNameContext) {}

// ExitProcedureName is called when production procedureName is exited.
func (s *BaseOracleStatementListener) ExitProcedureName(ctx *ProcedureNameContext) {}

// EnterCpuCost is called when production cpuCost is entered.
func (s *BaseOracleStatementListener) EnterCpuCost(ctx *CpuCostContext) {}

// ExitCpuCost is called when production cpuCost is exited.
func (s *BaseOracleStatementListener) ExitCpuCost(ctx *CpuCostContext) {}

// EnterIoCost is called when production ioCost is entered.
func (s *BaseOracleStatementListener) EnterIoCost(ctx *IoCostContext) {}

// ExitIoCost is called when production ioCost is exited.
func (s *BaseOracleStatementListener) ExitIoCost(ctx *IoCostContext) {}

// EnterNetworkCost is called when production networkCost is entered.
func (s *BaseOracleStatementListener) EnterNetworkCost(ctx *NetworkCostContext) {}

// ExitNetworkCost is called when production networkCost is exited.
func (s *BaseOracleStatementListener) ExitNetworkCost(ctx *NetworkCostContext) {}

// EnterDefaultSelectivity is called when production defaultSelectivity is entered.
func (s *BaseOracleStatementListener) EnterDefaultSelectivity(ctx *DefaultSelectivityContext) {}

// ExitDefaultSelectivity is called when production defaultSelectivity is exited.
func (s *BaseOracleStatementListener) ExitDefaultSelectivity(ctx *DefaultSelectivityContext) {}

// EnterDataItem is called when production dataItem is entered.
func (s *BaseOracleStatementListener) EnterDataItem(ctx *DataItemContext) {}

// ExitDataItem is called when production dataItem is exited.
func (s *BaseOracleStatementListener) ExitDataItem(ctx *DataItemContext) {}

// EnterVariableName is called when production variableName is entered.
func (s *BaseOracleStatementListener) EnterVariableName(ctx *VariableNameContext) {}

// ExitVariableName is called when production variableName is exited.
func (s *BaseOracleStatementListener) ExitVariableName(ctx *VariableNameContext) {}

// EnterValidTimeColumn is called when production validTimeColumn is entered.
func (s *BaseOracleStatementListener) EnterValidTimeColumn(ctx *ValidTimeColumnContext) {}

// ExitValidTimeColumn is called when production validTimeColumn is exited.
func (s *BaseOracleStatementListener) ExitValidTimeColumn(ctx *ValidTimeColumnContext) {}

// EnterAttrDim is called when production attrDim is entered.
func (s *BaseOracleStatementListener) EnterAttrDim(ctx *AttrDimContext) {}

// ExitAttrDim is called when production attrDim is exited.
func (s *BaseOracleStatementListener) ExitAttrDim(ctx *AttrDimContext) {}

// EnterHierarchyName is called when production hierarchyName is entered.
func (s *BaseOracleStatementListener) EnterHierarchyName(ctx *HierarchyNameContext) {}

// ExitHierarchyName is called when production hierarchyName is exited.
func (s *BaseOracleStatementListener) ExitHierarchyName(ctx *HierarchyNameContext) {}

// EnterAnalyticViewName is called when production analyticViewName is entered.
func (s *BaseOracleStatementListener) EnterAnalyticViewName(ctx *AnalyticViewNameContext) {}

// ExitAnalyticViewName is called when production analyticViewName is exited.
func (s *BaseOracleStatementListener) ExitAnalyticViewName(ctx *AnalyticViewNameContext) {}

// EnterSamplePercent is called when production samplePercent is entered.
func (s *BaseOracleStatementListener) EnterSamplePercent(ctx *SamplePercentContext) {}

// ExitSamplePercent is called when production samplePercent is exited.
func (s *BaseOracleStatementListener) ExitSamplePercent(ctx *SamplePercentContext) {}

// EnterSeedValue is called when production seedValue is entered.
func (s *BaseOracleStatementListener) EnterSeedValue(ctx *SeedValueContext) {}

// ExitSeedValue is called when production seedValue is exited.
func (s *BaseOracleStatementListener) ExitSeedValue(ctx *SeedValueContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseOracleStatementListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseOracleStatementListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterRestorePoint is called when production restorePoint is entered.
func (s *BaseOracleStatementListener) EnterRestorePoint(ctx *RestorePointContext) {}

// ExitRestorePoint is called when production restorePoint is exited.
func (s *BaseOracleStatementListener) ExitRestorePoint(ctx *RestorePointContext) {}

// EnterScnValue is called when production scnValue is entered.
func (s *BaseOracleStatementListener) EnterScnValue(ctx *ScnValueContext) {}

// ExitScnValue is called when production scnValue is exited.
func (s *BaseOracleStatementListener) ExitScnValue(ctx *ScnValueContext) {}

// EnterTimestampValue is called when production timestampValue is entered.
func (s *BaseOracleStatementListener) EnterTimestampValue(ctx *TimestampValueContext) {}

// ExitTimestampValue is called when production timestampValue is exited.
func (s *BaseOracleStatementListener) ExitTimestampValue(ctx *TimestampValueContext) {}

// EnterScnTimestampExpr is called when production scnTimestampExpr is entered.
func (s *BaseOracleStatementListener) EnterScnTimestampExpr(ctx *ScnTimestampExprContext) {}

// ExitScnTimestampExpr is called when production scnTimestampExpr is exited.
func (s *BaseOracleStatementListener) ExitScnTimestampExpr(ctx *ScnTimestampExprContext) {}

// EnterReferenceModelName is called when production referenceModelName is entered.
func (s *BaseOracleStatementListener) EnterReferenceModelName(ctx *ReferenceModelNameContext) {}

// ExitReferenceModelName is called when production referenceModelName is exited.
func (s *BaseOracleStatementListener) ExitReferenceModelName(ctx *ReferenceModelNameContext) {}

// EnterMainModelName is called when production mainModelName is entered.
func (s *BaseOracleStatementListener) EnterMainModelName(ctx *MainModelNameContext) {}

// ExitMainModelName is called when production mainModelName is exited.
func (s *BaseOracleStatementListener) ExitMainModelName(ctx *MainModelNameContext) {}

// EnterMeasureColumn is called when production measureColumn is entered.
func (s *BaseOracleStatementListener) EnterMeasureColumn(ctx *MeasureColumnContext) {}

// ExitMeasureColumn is called when production measureColumn is exited.
func (s *BaseOracleStatementListener) ExitMeasureColumn(ctx *MeasureColumnContext) {}

// EnterDimensionColumn is called when production dimensionColumn is entered.
func (s *BaseOracleStatementListener) EnterDimensionColumn(ctx *DimensionColumnContext) {}

// ExitDimensionColumn is called when production dimensionColumn is exited.
func (s *BaseOracleStatementListener) ExitDimensionColumn(ctx *DimensionColumnContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseOracleStatementListener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseOracleStatementListener) ExitPattern(ctx *PatternContext) {}

// EnterAnalyticFunctionName is called when production analyticFunctionName is entered.
func (s *BaseOracleStatementListener) EnterAnalyticFunctionName(ctx *AnalyticFunctionNameContext) {}

// ExitAnalyticFunctionName is called when production analyticFunctionName is exited.
func (s *BaseOracleStatementListener) ExitAnalyticFunctionName(ctx *AnalyticFunctionNameContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseOracleStatementListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseOracleStatementListener) ExitCondition(ctx *ConditionContext) {}

// EnterComparisonCondition is called when production comparisonCondition is entered.
func (s *BaseOracleStatementListener) EnterComparisonCondition(ctx *ComparisonConditionContext) {}

// ExitComparisonCondition is called when production comparisonCondition is exited.
func (s *BaseOracleStatementListener) ExitComparisonCondition(ctx *ComparisonConditionContext) {}

// EnterSimpleComparisonCondition is called when production simpleComparisonCondition is entered.
func (s *BaseOracleStatementListener) EnterSimpleComparisonCondition(ctx *SimpleComparisonConditionContext) {
}

// ExitSimpleComparisonCondition is called when production simpleComparisonCondition is exited.
func (s *BaseOracleStatementListener) ExitSimpleComparisonCondition(ctx *SimpleComparisonConditionContext) {
}

// EnterGroupComparisonCondition is called when production groupComparisonCondition is entered.
func (s *BaseOracleStatementListener) EnterGroupComparisonCondition(ctx *GroupComparisonConditionContext) {
}

// ExitGroupComparisonCondition is called when production groupComparisonCondition is exited.
func (s *BaseOracleStatementListener) ExitGroupComparisonCondition(ctx *GroupComparisonConditionContext) {
}

// EnterFloatingPointCondition is called when production floatingPointCondition is entered.
func (s *BaseOracleStatementListener) EnterFloatingPointCondition(ctx *FloatingPointConditionContext) {
}

// ExitFloatingPointCondition is called when production floatingPointCondition is exited.
func (s *BaseOracleStatementListener) ExitFloatingPointCondition(ctx *FloatingPointConditionContext) {
}

// EnterLogicalCondition is called when production logicalCondition is entered.
func (s *BaseOracleStatementListener) EnterLogicalCondition(ctx *LogicalConditionContext) {}

// ExitLogicalCondition is called when production logicalCondition is exited.
func (s *BaseOracleStatementListener) ExitLogicalCondition(ctx *LogicalConditionContext) {}

// EnterModelCondition is called when production modelCondition is entered.
func (s *BaseOracleStatementListener) EnterModelCondition(ctx *ModelConditionContext) {}

// ExitModelCondition is called when production modelCondition is exited.
func (s *BaseOracleStatementListener) ExitModelCondition(ctx *ModelConditionContext) {}

// EnterIsAnyCondition is called when production isAnyCondition is entered.
func (s *BaseOracleStatementListener) EnterIsAnyCondition(ctx *IsAnyConditionContext) {}

// ExitIsAnyCondition is called when production isAnyCondition is exited.
func (s *BaseOracleStatementListener) ExitIsAnyCondition(ctx *IsAnyConditionContext) {}

// EnterIsPresentCondition is called when production isPresentCondition is entered.
func (s *BaseOracleStatementListener) EnterIsPresentCondition(ctx *IsPresentConditionContext) {}

// ExitIsPresentCondition is called when production isPresentCondition is exited.
func (s *BaseOracleStatementListener) ExitIsPresentCondition(ctx *IsPresentConditionContext) {}

// EnterCellReference is called when production cellReference is entered.
func (s *BaseOracleStatementListener) EnterCellReference(ctx *CellReferenceContext) {}

// ExitCellReference is called when production cellReference is exited.
func (s *BaseOracleStatementListener) ExitCellReference(ctx *CellReferenceContext) {}

// EnterMultisetCondition is called when production multisetCondition is entered.
func (s *BaseOracleStatementListener) EnterMultisetCondition(ctx *MultisetConditionContext) {}

// ExitMultisetCondition is called when production multisetCondition is exited.
func (s *BaseOracleStatementListener) ExitMultisetCondition(ctx *MultisetConditionContext) {}

// EnterIsASetCondition is called when production isASetCondition is entered.
func (s *BaseOracleStatementListener) EnterIsASetCondition(ctx *IsASetConditionContext) {}

// ExitIsASetCondition is called when production isASetCondition is exited.
func (s *BaseOracleStatementListener) ExitIsASetCondition(ctx *IsASetConditionContext) {}

// EnterIsEmptyCondition is called when production isEmptyCondition is entered.
func (s *BaseOracleStatementListener) EnterIsEmptyCondition(ctx *IsEmptyConditionContext) {}

// ExitIsEmptyCondition is called when production isEmptyCondition is exited.
func (s *BaseOracleStatementListener) ExitIsEmptyCondition(ctx *IsEmptyConditionContext) {}

// EnterMemberCondition is called when production memberCondition is entered.
func (s *BaseOracleStatementListener) EnterMemberCondition(ctx *MemberConditionContext) {}

// ExitMemberCondition is called when production memberCondition is exited.
func (s *BaseOracleStatementListener) ExitMemberCondition(ctx *MemberConditionContext) {}

// EnterSubmultisetCondition is called when production submultisetCondition is entered.
func (s *BaseOracleStatementListener) EnterSubmultisetCondition(ctx *SubmultisetConditionContext) {}

// ExitSubmultisetCondition is called when production submultisetCondition is exited.
func (s *BaseOracleStatementListener) ExitSubmultisetCondition(ctx *SubmultisetConditionContext) {}

// EnterPatternMatchingCondition is called when production patternMatchingCondition is entered.
func (s *BaseOracleStatementListener) EnterPatternMatchingCondition(ctx *PatternMatchingConditionContext) {
}

// ExitPatternMatchingCondition is called when production patternMatchingCondition is exited.
func (s *BaseOracleStatementListener) ExitPatternMatchingCondition(ctx *PatternMatchingConditionContext) {
}

// EnterLikeCondition is called when production likeCondition is entered.
func (s *BaseOracleStatementListener) EnterLikeCondition(ctx *LikeConditionContext) {}

// ExitLikeCondition is called when production likeCondition is exited.
func (s *BaseOracleStatementListener) ExitLikeCondition(ctx *LikeConditionContext) {}

// EnterSearchValue is called when production searchValue is entered.
func (s *BaseOracleStatementListener) EnterSearchValue(ctx *SearchValueContext) {}

// ExitSearchValue is called when production searchValue is exited.
func (s *BaseOracleStatementListener) ExitSearchValue(ctx *SearchValueContext) {}

// EnterEscapeChar is called when production escapeChar is entered.
func (s *BaseOracleStatementListener) EnterEscapeChar(ctx *EscapeCharContext) {}

// ExitEscapeChar is called when production escapeChar is exited.
func (s *BaseOracleStatementListener) ExitEscapeChar(ctx *EscapeCharContext) {}

// EnterRegexpLikeCondition is called when production regexpLikeCondition is entered.
func (s *BaseOracleStatementListener) EnterRegexpLikeCondition(ctx *RegexpLikeConditionContext) {}

// ExitRegexpLikeCondition is called when production regexpLikeCondition is exited.
func (s *BaseOracleStatementListener) ExitRegexpLikeCondition(ctx *RegexpLikeConditionContext) {}

// EnterMatchParam is called when production matchParam is entered.
func (s *BaseOracleStatementListener) EnterMatchParam(ctx *MatchParamContext) {}

// ExitMatchParam is called when production matchParam is exited.
func (s *BaseOracleStatementListener) ExitMatchParam(ctx *MatchParamContext) {}

// EnterRangeCondition is called when production rangeCondition is entered.
func (s *BaseOracleStatementListener) EnterRangeCondition(ctx *RangeConditionContext) {}

// ExitRangeCondition is called when production rangeCondition is exited.
func (s *BaseOracleStatementListener) ExitRangeCondition(ctx *RangeConditionContext) {}

// EnterNullCondition is called when production nullCondition is entered.
func (s *BaseOracleStatementListener) EnterNullCondition(ctx *NullConditionContext) {}

// ExitNullCondition is called when production nullCondition is exited.
func (s *BaseOracleStatementListener) ExitNullCondition(ctx *NullConditionContext) {}

// EnterXmlCondition is called when production xmlCondition is entered.
func (s *BaseOracleStatementListener) EnterXmlCondition(ctx *XmlConditionContext) {}

// ExitXmlCondition is called when production xmlCondition is exited.
func (s *BaseOracleStatementListener) ExitXmlCondition(ctx *XmlConditionContext) {}

// EnterEqualsPathCondition is called when production equalsPathCondition is entered.
func (s *BaseOracleStatementListener) EnterEqualsPathCondition(ctx *EqualsPathConditionContext) {}

// ExitEqualsPathCondition is called when production equalsPathCondition is exited.
func (s *BaseOracleStatementListener) ExitEqualsPathCondition(ctx *EqualsPathConditionContext) {}

// EnterPathString is called when production pathString is entered.
func (s *BaseOracleStatementListener) EnterPathString(ctx *PathStringContext) {}

// ExitPathString is called when production pathString is exited.
func (s *BaseOracleStatementListener) ExitPathString(ctx *PathStringContext) {}

// EnterCorrelationInteger is called when production correlationInteger is entered.
func (s *BaseOracleStatementListener) EnterCorrelationInteger(ctx *CorrelationIntegerContext) {}

// ExitCorrelationInteger is called when production correlationInteger is exited.
func (s *BaseOracleStatementListener) ExitCorrelationInteger(ctx *CorrelationIntegerContext) {}

// EnterUnderPathCondition is called when production underPathCondition is entered.
func (s *BaseOracleStatementListener) EnterUnderPathCondition(ctx *UnderPathConditionContext) {}

// ExitUnderPathCondition is called when production underPathCondition is exited.
func (s *BaseOracleStatementListener) ExitUnderPathCondition(ctx *UnderPathConditionContext) {}

// EnterLevels is called when production levels is entered.
func (s *BaseOracleStatementListener) EnterLevels(ctx *LevelsContext) {}

// ExitLevels is called when production levels is exited.
func (s *BaseOracleStatementListener) ExitLevels(ctx *LevelsContext) {}

// EnterJsonCondition is called when production jsonCondition is entered.
func (s *BaseOracleStatementListener) EnterJsonCondition(ctx *JsonConditionContext) {}

// ExitJsonCondition is called when production jsonCondition is exited.
func (s *BaseOracleStatementListener) ExitJsonCondition(ctx *JsonConditionContext) {}

// EnterIsJsonCondition is called when production isJsonCondition is entered.
func (s *BaseOracleStatementListener) EnterIsJsonCondition(ctx *IsJsonConditionContext) {}

// ExitIsJsonCondition is called when production isJsonCondition is exited.
func (s *BaseOracleStatementListener) ExitIsJsonCondition(ctx *IsJsonConditionContext) {}

// EnterJsonEqualCondition is called when production jsonEqualCondition is entered.
func (s *BaseOracleStatementListener) EnterJsonEqualCondition(ctx *JsonEqualConditionContext) {}

// ExitJsonEqualCondition is called when production jsonEqualCondition is exited.
func (s *BaseOracleStatementListener) ExitJsonEqualCondition(ctx *JsonEqualConditionContext) {}

// EnterJsonExistsCondition is called when production jsonExistsCondition is entered.
func (s *BaseOracleStatementListener) EnterJsonExistsCondition(ctx *JsonExistsConditionContext) {}

// ExitJsonExistsCondition is called when production jsonExistsCondition is exited.
func (s *BaseOracleStatementListener) ExitJsonExistsCondition(ctx *JsonExistsConditionContext) {}

// EnterJsonPassingClause is called when production jsonPassingClause is entered.
func (s *BaseOracleStatementListener) EnterJsonPassingClause(ctx *JsonPassingClauseContext) {}

// ExitJsonPassingClause is called when production jsonPassingClause is exited.
func (s *BaseOracleStatementListener) ExitJsonPassingClause(ctx *JsonPassingClauseContext) {}

// EnterJsonExistsOnErrorClause is called when production jsonExistsOnErrorClause is entered.
func (s *BaseOracleStatementListener) EnterJsonExistsOnErrorClause(ctx *JsonExistsOnErrorClauseContext) {
}

// ExitJsonExistsOnErrorClause is called when production jsonExistsOnErrorClause is exited.
func (s *BaseOracleStatementListener) ExitJsonExistsOnErrorClause(ctx *JsonExistsOnErrorClauseContext) {
}

// EnterJsonExistsOnEmptyClause is called when production jsonExistsOnEmptyClause is entered.
func (s *BaseOracleStatementListener) EnterJsonExistsOnEmptyClause(ctx *JsonExistsOnEmptyClauseContext) {
}

// ExitJsonExistsOnEmptyClause is called when production jsonExistsOnEmptyClause is exited.
func (s *BaseOracleStatementListener) ExitJsonExistsOnEmptyClause(ctx *JsonExistsOnEmptyClauseContext) {
}

// EnterJsonTextcontainsCondition is called when production jsonTextcontainsCondition is entered.
func (s *BaseOracleStatementListener) EnterJsonTextcontainsCondition(ctx *JsonTextcontainsConditionContext) {
}

// ExitJsonTextcontainsCondition is called when production jsonTextcontainsCondition is exited.
func (s *BaseOracleStatementListener) ExitJsonTextcontainsCondition(ctx *JsonTextcontainsConditionContext) {
}

// EnterJsonBasicPathExpr is called when production jsonBasicPathExpr is entered.
func (s *BaseOracleStatementListener) EnterJsonBasicPathExpr(ctx *JsonBasicPathExprContext) {}

// ExitJsonBasicPathExpr is called when production jsonBasicPathExpr is exited.
func (s *BaseOracleStatementListener) ExitJsonBasicPathExpr(ctx *JsonBasicPathExprContext) {}

// EnterJsonAbsolutePathExpr is called when production jsonAbsolutePathExpr is entered.
func (s *BaseOracleStatementListener) EnterJsonAbsolutePathExpr(ctx *JsonAbsolutePathExprContext) {}

// ExitJsonAbsolutePathExpr is called when production jsonAbsolutePathExpr is exited.
func (s *BaseOracleStatementListener) ExitJsonAbsolutePathExpr(ctx *JsonAbsolutePathExprContext) {}

// EnterJsonNonfunctionSteps is called when production jsonNonfunctionSteps is entered.
func (s *BaseOracleStatementListener) EnterJsonNonfunctionSteps(ctx *JsonNonfunctionStepsContext) {}

// ExitJsonNonfunctionSteps is called when production jsonNonfunctionSteps is exited.
func (s *BaseOracleStatementListener) ExitJsonNonfunctionSteps(ctx *JsonNonfunctionStepsContext) {}

// EnterJsonObjectStep is called when production jsonObjectStep is entered.
func (s *BaseOracleStatementListener) EnterJsonObjectStep(ctx *JsonObjectStepContext) {}

// ExitJsonObjectStep is called when production jsonObjectStep is exited.
func (s *BaseOracleStatementListener) ExitJsonObjectStep(ctx *JsonObjectStepContext) {}

// EnterJsonFieldName is called when production jsonFieldName is entered.
func (s *BaseOracleStatementListener) EnterJsonFieldName(ctx *JsonFieldNameContext) {}

// ExitJsonFieldName is called when production jsonFieldName is exited.
func (s *BaseOracleStatementListener) ExitJsonFieldName(ctx *JsonFieldNameContext) {}

// EnterLetter is called when production letter is entered.
func (s *BaseOracleStatementListener) EnterLetter(ctx *LetterContext) {}

// ExitLetter is called when production letter is exited.
func (s *BaseOracleStatementListener) ExitLetter(ctx *LetterContext) {}

// EnterDigit is called when production digit is entered.
func (s *BaseOracleStatementListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production digit is exited.
func (s *BaseOracleStatementListener) ExitDigit(ctx *DigitContext) {}

// EnterJsonArrayStep is called when production jsonArrayStep is entered.
func (s *BaseOracleStatementListener) EnterJsonArrayStep(ctx *JsonArrayStepContext) {}

// ExitJsonArrayStep is called when production jsonArrayStep is exited.
func (s *BaseOracleStatementListener) ExitJsonArrayStep(ctx *JsonArrayStepContext) {}

// EnterJsonDescendentStep is called when production jsonDescendentStep is entered.
func (s *BaseOracleStatementListener) EnterJsonDescendentStep(ctx *JsonDescendentStepContext) {}

// ExitJsonDescendentStep is called when production jsonDescendentStep is exited.
func (s *BaseOracleStatementListener) ExitJsonDescendentStep(ctx *JsonDescendentStepContext) {}

// EnterJsonFunctionStep is called when production jsonFunctionStep is entered.
func (s *BaseOracleStatementListener) EnterJsonFunctionStep(ctx *JsonFunctionStepContext) {}

// ExitJsonFunctionStep is called when production jsonFunctionStep is exited.
func (s *BaseOracleStatementListener) ExitJsonFunctionStep(ctx *JsonFunctionStepContext) {}

// EnterJsonItemMethod is called when production jsonItemMethod is entered.
func (s *BaseOracleStatementListener) EnterJsonItemMethod(ctx *JsonItemMethodContext) {}

// ExitJsonItemMethod is called when production jsonItemMethod is exited.
func (s *BaseOracleStatementListener) ExitJsonItemMethod(ctx *JsonItemMethodContext) {}

// EnterJsonFilterExpr is called when production jsonFilterExpr is entered.
func (s *BaseOracleStatementListener) EnterJsonFilterExpr(ctx *JsonFilterExprContext) {}

// ExitJsonFilterExpr is called when production jsonFilterExpr is exited.
func (s *BaseOracleStatementListener) ExitJsonFilterExpr(ctx *JsonFilterExprContext) {}

// EnterJsonCond is called when production jsonCond is entered.
func (s *BaseOracleStatementListener) EnterJsonCond(ctx *JsonCondContext) {}

// ExitJsonCond is called when production jsonCond is exited.
func (s *BaseOracleStatementListener) ExitJsonCond(ctx *JsonCondContext) {}

// EnterJsonDisjunction is called when production jsonDisjunction is entered.
func (s *BaseOracleStatementListener) EnterJsonDisjunction(ctx *JsonDisjunctionContext) {}

// ExitJsonDisjunction is called when production jsonDisjunction is exited.
func (s *BaseOracleStatementListener) ExitJsonDisjunction(ctx *JsonDisjunctionContext) {}

// EnterJsonConjunction is called when production jsonConjunction is entered.
func (s *BaseOracleStatementListener) EnterJsonConjunction(ctx *JsonConjunctionContext) {}

// ExitJsonConjunction is called when production jsonConjunction is exited.
func (s *BaseOracleStatementListener) ExitJsonConjunction(ctx *JsonConjunctionContext) {}

// EnterJsonNegation is called when production jsonNegation is entered.
func (s *BaseOracleStatementListener) EnterJsonNegation(ctx *JsonNegationContext) {}

// ExitJsonNegation is called when production jsonNegation is exited.
func (s *BaseOracleStatementListener) ExitJsonNegation(ctx *JsonNegationContext) {}

// EnterJsonExistsCond is called when production jsonExistsCond is entered.
func (s *BaseOracleStatementListener) EnterJsonExistsCond(ctx *JsonExistsCondContext) {}

// ExitJsonExistsCond is called when production jsonExistsCond is exited.
func (s *BaseOracleStatementListener) ExitJsonExistsCond(ctx *JsonExistsCondContext) {}

// EnterJsonHasSubstringCond is called when production jsonHasSubstringCond is entered.
func (s *BaseOracleStatementListener) EnterJsonHasSubstringCond(ctx *JsonHasSubstringCondContext) {}

// ExitJsonHasSubstringCond is called when production jsonHasSubstringCond is exited.
func (s *BaseOracleStatementListener) ExitJsonHasSubstringCond(ctx *JsonHasSubstringCondContext) {}

// EnterJsonStartsWithCond is called when production jsonStartsWithCond is entered.
func (s *BaseOracleStatementListener) EnterJsonStartsWithCond(ctx *JsonStartsWithCondContext) {}

// ExitJsonStartsWithCond is called when production jsonStartsWithCond is exited.
func (s *BaseOracleStatementListener) ExitJsonStartsWithCond(ctx *JsonStartsWithCondContext) {}

// EnterJsonLikeCond is called when production jsonLikeCond is entered.
func (s *BaseOracleStatementListener) EnterJsonLikeCond(ctx *JsonLikeCondContext) {}

// ExitJsonLikeCond is called when production jsonLikeCond is exited.
func (s *BaseOracleStatementListener) ExitJsonLikeCond(ctx *JsonLikeCondContext) {}

// EnterJsonLikeRegexCond is called when production jsonLikeRegexCond is entered.
func (s *BaseOracleStatementListener) EnterJsonLikeRegexCond(ctx *JsonLikeRegexCondContext) {}

// ExitJsonLikeRegexCond is called when production jsonLikeRegexCond is exited.
func (s *BaseOracleStatementListener) ExitJsonLikeRegexCond(ctx *JsonLikeRegexCondContext) {}

// EnterJsonEqRegexCond is called when production jsonEqRegexCond is entered.
func (s *BaseOracleStatementListener) EnterJsonEqRegexCond(ctx *JsonEqRegexCondContext) {}

// ExitJsonEqRegexCond is called when production jsonEqRegexCond is exited.
func (s *BaseOracleStatementListener) ExitJsonEqRegexCond(ctx *JsonEqRegexCondContext) {}

// EnterJsonInCond is called when production jsonInCond is entered.
func (s *BaseOracleStatementListener) EnterJsonInCond(ctx *JsonInCondContext) {}

// ExitJsonInCond is called when production jsonInCond is exited.
func (s *BaseOracleStatementListener) ExitJsonInCond(ctx *JsonInCondContext) {}

// EnterValueList is called when production valueList is entered.
func (s *BaseOracleStatementListener) EnterValueList(ctx *ValueListContext) {}

// ExitValueList is called when production valueList is exited.
func (s *BaseOracleStatementListener) ExitValueList(ctx *ValueListContext) {}

// EnterJsonComparison is called when production jsonComparison is entered.
func (s *BaseOracleStatementListener) EnterJsonComparison(ctx *JsonComparisonContext) {}

// ExitJsonComparison is called when production jsonComparison is exited.
func (s *BaseOracleStatementListener) ExitJsonComparison(ctx *JsonComparisonContext) {}

// EnterJsonRelativePathExpr is called when production jsonRelativePathExpr is entered.
func (s *BaseOracleStatementListener) EnterJsonRelativePathExpr(ctx *JsonRelativePathExprContext) {}

// ExitJsonRelativePathExpr is called when production jsonRelativePathExpr is exited.
func (s *BaseOracleStatementListener) ExitJsonRelativePathExpr(ctx *JsonRelativePathExprContext) {}

// EnterJsonComparePred is called when production jsonComparePred is entered.
func (s *BaseOracleStatementListener) EnterJsonComparePred(ctx *JsonComparePredContext) {}

// ExitJsonComparePred is called when production jsonComparePred is exited.
func (s *BaseOracleStatementListener) ExitJsonComparePred(ctx *JsonComparePredContext) {}

// EnterJsonVar is called when production jsonVar is entered.
func (s *BaseOracleStatementListener) EnterJsonVar(ctx *JsonVarContext) {}

// ExitJsonVar is called when production jsonVar is exited.
func (s *BaseOracleStatementListener) ExitJsonVar(ctx *JsonVarContext) {}

// EnterJsonScalar is called when production jsonScalar is entered.
func (s *BaseOracleStatementListener) EnterJsonScalar(ctx *JsonScalarContext) {}

// ExitJsonScalar is called when production jsonScalar is exited.
func (s *BaseOracleStatementListener) ExitJsonScalar(ctx *JsonScalarContext) {}

// EnterJsonNumber is called when production jsonNumber is entered.
func (s *BaseOracleStatementListener) EnterJsonNumber(ctx *JsonNumberContext) {}

// ExitJsonNumber is called when production jsonNumber is exited.
func (s *BaseOracleStatementListener) ExitJsonNumber(ctx *JsonNumberContext) {}

// EnterJsonString is called when production jsonString is entered.
func (s *BaseOracleStatementListener) EnterJsonString(ctx *JsonStringContext) {}

// ExitJsonString is called when production jsonString is exited.
func (s *BaseOracleStatementListener) ExitJsonString(ctx *JsonStringContext) {}

// EnterCompoundCondition is called when production compoundCondition is entered.
func (s *BaseOracleStatementListener) EnterCompoundCondition(ctx *CompoundConditionContext) {}

// ExitCompoundCondition is called when production compoundCondition is exited.
func (s *BaseOracleStatementListener) ExitCompoundCondition(ctx *CompoundConditionContext) {}

// EnterExistsCondition is called when production existsCondition is entered.
func (s *BaseOracleStatementListener) EnterExistsCondition(ctx *ExistsConditionContext) {}

// ExitExistsCondition is called when production existsCondition is exited.
func (s *BaseOracleStatementListener) ExitExistsCondition(ctx *ExistsConditionContext) {}

// EnterInCondition is called when production inCondition is entered.
func (s *BaseOracleStatementListener) EnterInCondition(ctx *InConditionContext) {}

// ExitInCondition is called when production inCondition is exited.
func (s *BaseOracleStatementListener) ExitInCondition(ctx *InConditionContext) {}

// EnterIsOfTypeCondition is called when production isOfTypeCondition is entered.
func (s *BaseOracleStatementListener) EnterIsOfTypeCondition(ctx *IsOfTypeConditionContext) {}

// ExitIsOfTypeCondition is called when production isOfTypeCondition is exited.
func (s *BaseOracleStatementListener) ExitIsOfTypeCondition(ctx *IsOfTypeConditionContext) {}

// EnterDatabaseCharset is called when production databaseCharset is entered.
func (s *BaseOracleStatementListener) EnterDatabaseCharset(ctx *DatabaseCharsetContext) {}

// ExitDatabaseCharset is called when production databaseCharset is exited.
func (s *BaseOracleStatementListener) ExitDatabaseCharset(ctx *DatabaseCharsetContext) {}

// EnterNationalCharset is called when production nationalCharset is entered.
func (s *BaseOracleStatementListener) EnterNationalCharset(ctx *NationalCharsetContext) {}

// ExitNationalCharset is called when production nationalCharset is exited.
func (s *BaseOracleStatementListener) ExitNationalCharset(ctx *NationalCharsetContext) {}

// EnterFilenamePattern is called when production filenamePattern is entered.
func (s *BaseOracleStatementListener) EnterFilenamePattern(ctx *FilenamePatternContext) {}

// ExitFilenamePattern is called when production filenamePattern is exited.
func (s *BaseOracleStatementListener) ExitFilenamePattern(ctx *FilenamePatternContext) {}

// EnterConnectString is called when production connectString is entered.
func (s *BaseOracleStatementListener) EnterConnectString(ctx *ConnectStringContext) {}

// ExitConnectString is called when production connectString is exited.
func (s *BaseOracleStatementListener) ExitConnectString(ctx *ConnectStringContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseOracleStatementListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseOracleStatementListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterCreateIndex is called when production createIndex is entered.
func (s *BaseOracleStatementListener) EnterCreateIndex(ctx *CreateIndexContext) {}

// ExitCreateIndex is called when production createIndex is exited.
func (s *BaseOracleStatementListener) ExitCreateIndex(ctx *CreateIndexContext) {}

// EnterAlterTable is called when production alterTable is entered.
func (s *BaseOracleStatementListener) EnterAlterTable(ctx *AlterTableContext) {}

// ExitAlterTable is called when production alterTable is exited.
func (s *BaseOracleStatementListener) ExitAlterTable(ctx *AlterTableContext) {}

// EnterAlterIndex is called when production alterIndex is entered.
func (s *BaseOracleStatementListener) EnterAlterIndex(ctx *AlterIndexContext) {}

// ExitAlterIndex is called when production alterIndex is exited.
func (s *BaseOracleStatementListener) ExitAlterIndex(ctx *AlterIndexContext) {}

// EnterDropTable is called when production dropTable is entered.
func (s *BaseOracleStatementListener) EnterDropTable(ctx *DropTableContext) {}

// ExitDropTable is called when production dropTable is exited.
func (s *BaseOracleStatementListener) ExitDropTable(ctx *DropTableContext) {}

// EnterDropIndex is called when production dropIndex is entered.
func (s *BaseOracleStatementListener) EnterDropIndex(ctx *DropIndexContext) {}

// ExitDropIndex is called when production dropIndex is exited.
func (s *BaseOracleStatementListener) ExitDropIndex(ctx *DropIndexContext) {}

// EnterTruncateTable is called when production truncateTable is entered.
func (s *BaseOracleStatementListener) EnterTruncateTable(ctx *TruncateTableContext) {}

// ExitTruncateTable is called when production truncateTable is exited.
func (s *BaseOracleStatementListener) ExitTruncateTable(ctx *TruncateTableContext) {}

// EnterCreateTableSpecification is called when production createTableSpecification is entered.
func (s *BaseOracleStatementListener) EnterCreateTableSpecification(ctx *CreateTableSpecificationContext) {
}

// ExitCreateTableSpecification is called when production createTableSpecification is exited.
func (s *BaseOracleStatementListener) ExitCreateTableSpecification(ctx *CreateTableSpecificationContext) {
}

// EnterTablespaceClauseWithParen is called when production tablespaceClauseWithParen is entered.
func (s *BaseOracleStatementListener) EnterTablespaceClauseWithParen(ctx *TablespaceClauseWithParenContext) {
}

// ExitTablespaceClauseWithParen is called when production tablespaceClauseWithParen is exited.
func (s *BaseOracleStatementListener) ExitTablespaceClauseWithParen(ctx *TablespaceClauseWithParenContext) {
}

// EnterTablespaceClause is called when production tablespaceClause is entered.
func (s *BaseOracleStatementListener) EnterTablespaceClause(ctx *TablespaceClauseContext) {}

// ExitTablespaceClause is called when production tablespaceClause is exited.
func (s *BaseOracleStatementListener) ExitTablespaceClause(ctx *TablespaceClauseContext) {}

// EnterCreateSharingClause is called when production createSharingClause is entered.
func (s *BaseOracleStatementListener) EnterCreateSharingClause(ctx *CreateSharingClauseContext) {}

// ExitCreateSharingClause is called when production createSharingClause is exited.
func (s *BaseOracleStatementListener) ExitCreateSharingClause(ctx *CreateSharingClauseContext) {}

// EnterCreateDefinitionClause is called when production createDefinitionClause is entered.
func (s *BaseOracleStatementListener) EnterCreateDefinitionClause(ctx *CreateDefinitionClauseContext) {
}

// ExitCreateDefinitionClause is called when production createDefinitionClause is exited.
func (s *BaseOracleStatementListener) ExitCreateDefinitionClause(ctx *CreateDefinitionClauseContext) {
}

// EnterCreateXMLTypeTableClause is called when production createXMLTypeTableClause is entered.
func (s *BaseOracleStatementListener) EnterCreateXMLTypeTableClause(ctx *CreateXMLTypeTableClauseContext) {
}

// ExitCreateXMLTypeTableClause is called when production createXMLTypeTableClause is exited.
func (s *BaseOracleStatementListener) ExitCreateXMLTypeTableClause(ctx *CreateXMLTypeTableClauseContext) {
}

// EnterXmlTypeStorageClause is called when production xmlTypeStorageClause is entered.
func (s *BaseOracleStatementListener) EnterXmlTypeStorageClause(ctx *XmlTypeStorageClauseContext) {}

// ExitXmlTypeStorageClause is called when production xmlTypeStorageClause is exited.
func (s *BaseOracleStatementListener) ExitXmlTypeStorageClause(ctx *XmlTypeStorageClauseContext) {}

// EnterXmlSchemaSpecClause is called when production xmlSchemaSpecClause is entered.
func (s *BaseOracleStatementListener) EnterXmlSchemaSpecClause(ctx *XmlSchemaSpecClauseContext) {}

// ExitXmlSchemaSpecClause is called when production xmlSchemaSpecClause is exited.
func (s *BaseOracleStatementListener) ExitXmlSchemaSpecClause(ctx *XmlSchemaSpecClauseContext) {}

// EnterXmlTypeVirtualColumnsClause is called when production xmlTypeVirtualColumnsClause is entered.
func (s *BaseOracleStatementListener) EnterXmlTypeVirtualColumnsClause(ctx *XmlTypeVirtualColumnsClauseContext) {
}

// ExitXmlTypeVirtualColumnsClause is called when production xmlTypeVirtualColumnsClause is exited.
func (s *BaseOracleStatementListener) ExitXmlTypeVirtualColumnsClause(ctx *XmlTypeVirtualColumnsClauseContext) {
}

// EnterOidClause is called when production oidClause is entered.
func (s *BaseOracleStatementListener) EnterOidClause(ctx *OidClauseContext) {}

// ExitOidClause is called when production oidClause is exited.
func (s *BaseOracleStatementListener) ExitOidClause(ctx *OidClauseContext) {}

// EnterOidIndexClause is called when production oidIndexClause is entered.
func (s *BaseOracleStatementListener) EnterOidIndexClause(ctx *OidIndexClauseContext) {}

// ExitOidIndexClause is called when production oidIndexClause is exited.
func (s *BaseOracleStatementListener) ExitOidIndexClause(ctx *OidIndexClauseContext) {}

// EnterCreateRelationalTableClause is called when production createRelationalTableClause is entered.
func (s *BaseOracleStatementListener) EnterCreateRelationalTableClause(ctx *CreateRelationalTableClauseContext) {
}

// ExitCreateRelationalTableClause is called when production createRelationalTableClause is exited.
func (s *BaseOracleStatementListener) ExitCreateRelationalTableClause(ctx *CreateRelationalTableClauseContext) {
}

// EnterCreateMemOptimizeClause is called when production createMemOptimizeClause is entered.
func (s *BaseOracleStatementListener) EnterCreateMemOptimizeClause(ctx *CreateMemOptimizeClauseContext) {
}

// ExitCreateMemOptimizeClause is called when production createMemOptimizeClause is exited.
func (s *BaseOracleStatementListener) ExitCreateMemOptimizeClause(ctx *CreateMemOptimizeClauseContext) {
}

// EnterCreateParentClause is called when production createParentClause is entered.
func (s *BaseOracleStatementListener) EnterCreateParentClause(ctx *CreateParentClauseContext) {}

// ExitCreateParentClause is called when production createParentClause is exited.
func (s *BaseOracleStatementListener) ExitCreateParentClause(ctx *CreateParentClauseContext) {}

// EnterCreateObjectTableClause is called when production createObjectTableClause is entered.
func (s *BaseOracleStatementListener) EnterCreateObjectTableClause(ctx *CreateObjectTableClauseContext) {
}

// ExitCreateObjectTableClause is called when production createObjectTableClause is exited.
func (s *BaseOracleStatementListener) ExitCreateObjectTableClause(ctx *CreateObjectTableClauseContext) {
}

// EnterRelationalProperties is called when production relationalProperties is entered.
func (s *BaseOracleStatementListener) EnterRelationalProperties(ctx *RelationalPropertiesContext) {}

// ExitRelationalProperties is called when production relationalProperties is exited.
func (s *BaseOracleStatementListener) ExitRelationalProperties(ctx *RelationalPropertiesContext) {}

// EnterRelationalProperty is called when production relationalProperty is entered.
func (s *BaseOracleStatementListener) EnterRelationalProperty(ctx *RelationalPropertyContext) {}

// ExitRelationalProperty is called when production relationalProperty is exited.
func (s *BaseOracleStatementListener) ExitRelationalProperty(ctx *RelationalPropertyContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseOracleStatementListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseOracleStatementListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterVisibleClause is called when production visibleClause is entered.
func (s *BaseOracleStatementListener) EnterVisibleClause(ctx *VisibleClauseContext) {}

// ExitVisibleClause is called when production visibleClause is exited.
func (s *BaseOracleStatementListener) ExitVisibleClause(ctx *VisibleClauseContext) {}

// EnterDefaultNullClause is called when production defaultNullClause is entered.
func (s *BaseOracleStatementListener) EnterDefaultNullClause(ctx *DefaultNullClauseContext) {}

// ExitDefaultNullClause is called when production defaultNullClause is exited.
func (s *BaseOracleStatementListener) ExitDefaultNullClause(ctx *DefaultNullClauseContext) {}

// EnterIdentityClause is called when production identityClause is entered.
func (s *BaseOracleStatementListener) EnterIdentityClause(ctx *IdentityClauseContext) {}

// ExitIdentityClause is called when production identityClause is exited.
func (s *BaseOracleStatementListener) ExitIdentityClause(ctx *IdentityClauseContext) {}

// EnterIdentifyOptions is called when production identifyOptions is entered.
func (s *BaseOracleStatementListener) EnterIdentifyOptions(ctx *IdentifyOptionsContext) {}

// ExitIdentifyOptions is called when production identifyOptions is exited.
func (s *BaseOracleStatementListener) ExitIdentifyOptions(ctx *IdentifyOptionsContext) {}

// EnterIdentityOption is called when production identityOption is entered.
func (s *BaseOracleStatementListener) EnterIdentityOption(ctx *IdentityOptionContext) {}

// ExitIdentityOption is called when production identityOption is exited.
func (s *BaseOracleStatementListener) ExitIdentityOption(ctx *IdentityOptionContext) {}

// EnterEncryptionSpecification is called when production encryptionSpecification is entered.
func (s *BaseOracleStatementListener) EnterEncryptionSpecification(ctx *EncryptionSpecificationContext) {
}

// ExitEncryptionSpecification is called when production encryptionSpecification is exited.
func (s *BaseOracleStatementListener) ExitEncryptionSpecification(ctx *EncryptionSpecificationContext) {
}

// EnterInlineConstraint is called when production inlineConstraint is entered.
func (s *BaseOracleStatementListener) EnterInlineConstraint(ctx *InlineConstraintContext) {}

// ExitInlineConstraint is called when production inlineConstraint is exited.
func (s *BaseOracleStatementListener) ExitInlineConstraint(ctx *InlineConstraintContext) {}

// EnterReferencesClause is called when production referencesClause is entered.
func (s *BaseOracleStatementListener) EnterReferencesClause(ctx *ReferencesClauseContext) {}

// ExitReferencesClause is called when production referencesClause is exited.
func (s *BaseOracleStatementListener) ExitReferencesClause(ctx *ReferencesClauseContext) {}

// EnterConstraintState is called when production constraintState is entered.
func (s *BaseOracleStatementListener) EnterConstraintState(ctx *ConstraintStateContext) {}

// ExitConstraintState is called when production constraintState is exited.
func (s *BaseOracleStatementListener) ExitConstraintState(ctx *ConstraintStateContext) {}

// EnterNotDeferrable is called when production notDeferrable is entered.
func (s *BaseOracleStatementListener) EnterNotDeferrable(ctx *NotDeferrableContext) {}

// ExitNotDeferrable is called when production notDeferrable is exited.
func (s *BaseOracleStatementListener) ExitNotDeferrable(ctx *NotDeferrableContext) {}

// EnterInitiallyClause is called when production initiallyClause is entered.
func (s *BaseOracleStatementListener) EnterInitiallyClause(ctx *InitiallyClauseContext) {}

// ExitInitiallyClause is called when production initiallyClause is exited.
func (s *BaseOracleStatementListener) ExitInitiallyClause(ctx *InitiallyClauseContext) {}

// EnterExceptionsClause is called when production exceptionsClause is entered.
func (s *BaseOracleStatementListener) EnterExceptionsClause(ctx *ExceptionsClauseContext) {}

// ExitExceptionsClause is called when production exceptionsClause is exited.
func (s *BaseOracleStatementListener) ExitExceptionsClause(ctx *ExceptionsClauseContext) {}

// EnterUsingIndexClause is called when production usingIndexClause is entered.
func (s *BaseOracleStatementListener) EnterUsingIndexClause(ctx *UsingIndexClauseContext) {}

// ExitUsingIndexClause is called when production usingIndexClause is exited.
func (s *BaseOracleStatementListener) ExitUsingIndexClause(ctx *UsingIndexClauseContext) {}

// EnterCreateIndexClause is called when production createIndexClause is entered.
func (s *BaseOracleStatementListener) EnterCreateIndexClause(ctx *CreateIndexClauseContext) {}

// ExitCreateIndexClause is called when production createIndexClause is exited.
func (s *BaseOracleStatementListener) ExitCreateIndexClause(ctx *CreateIndexClauseContext) {}

// EnterInlineRefConstraint is called when production inlineRefConstraint is entered.
func (s *BaseOracleStatementListener) EnterInlineRefConstraint(ctx *InlineRefConstraintContext) {}

// ExitInlineRefConstraint is called when production inlineRefConstraint is exited.
func (s *BaseOracleStatementListener) ExitInlineRefConstraint(ctx *InlineRefConstraintContext) {}

// EnterVirtualColumnDefinition is called when production virtualColumnDefinition is entered.
func (s *BaseOracleStatementListener) EnterVirtualColumnDefinition(ctx *VirtualColumnDefinitionContext) {
}

// ExitVirtualColumnDefinition is called when production virtualColumnDefinition is exited.
func (s *BaseOracleStatementListener) ExitVirtualColumnDefinition(ctx *VirtualColumnDefinitionContext) {
}

// EnterOutOfLineConstraint is called when production outOfLineConstraint is entered.
func (s *BaseOracleStatementListener) EnterOutOfLineConstraint(ctx *OutOfLineConstraintContext) {}

// ExitOutOfLineConstraint is called when production outOfLineConstraint is exited.
func (s *BaseOracleStatementListener) ExitOutOfLineConstraint(ctx *OutOfLineConstraintContext) {}

// EnterOutOfLineRefConstraint is called when production outOfLineRefConstraint is entered.
func (s *BaseOracleStatementListener) EnterOutOfLineRefConstraint(ctx *OutOfLineRefConstraintContext) {
}

// ExitOutOfLineRefConstraint is called when production outOfLineRefConstraint is exited.
func (s *BaseOracleStatementListener) ExitOutOfLineRefConstraint(ctx *OutOfLineRefConstraintContext) {
}

// EnterCreateIndexSpecification is called when production createIndexSpecification is entered.
func (s *BaseOracleStatementListener) EnterCreateIndexSpecification(ctx *CreateIndexSpecificationContext) {
}

// ExitCreateIndexSpecification is called when production createIndexSpecification is exited.
func (s *BaseOracleStatementListener) ExitCreateIndexSpecification(ctx *CreateIndexSpecificationContext) {
}

// EnterClusterIndexClause is called when production clusterIndexClause is entered.
func (s *BaseOracleStatementListener) EnterClusterIndexClause(ctx *ClusterIndexClauseContext) {}

// ExitClusterIndexClause is called when production clusterIndexClause is exited.
func (s *BaseOracleStatementListener) ExitClusterIndexClause(ctx *ClusterIndexClauseContext) {}

// EnterIndexAttributes is called when production indexAttributes is entered.
func (s *BaseOracleStatementListener) EnterIndexAttributes(ctx *IndexAttributesContext) {}

// ExitIndexAttributes is called when production indexAttributes is exited.
func (s *BaseOracleStatementListener) ExitIndexAttributes(ctx *IndexAttributesContext) {}

// EnterTableIndexClause is called when production tableIndexClause is entered.
func (s *BaseOracleStatementListener) EnterTableIndexClause(ctx *TableIndexClauseContext) {}

// ExitTableIndexClause is called when production tableIndexClause is exited.
func (s *BaseOracleStatementListener) ExitTableIndexClause(ctx *TableIndexClauseContext) {}

// EnterIndexExpressions is called when production indexExpressions is entered.
func (s *BaseOracleStatementListener) EnterIndexExpressions(ctx *IndexExpressionsContext) {}

// ExitIndexExpressions is called when production indexExpressions is exited.
func (s *BaseOracleStatementListener) ExitIndexExpressions(ctx *IndexExpressionsContext) {}

// EnterIndexExpression is called when production indexExpression is entered.
func (s *BaseOracleStatementListener) EnterIndexExpression(ctx *IndexExpressionContext) {}

// ExitIndexExpression is called when production indexExpression is exited.
func (s *BaseOracleStatementListener) ExitIndexExpression(ctx *IndexExpressionContext) {}

// EnterBitmapJoinIndexClause is called when production bitmapJoinIndexClause is entered.
func (s *BaseOracleStatementListener) EnterBitmapJoinIndexClause(ctx *BitmapJoinIndexClauseContext) {}

// ExitBitmapJoinIndexClause is called when production bitmapJoinIndexClause is exited.
func (s *BaseOracleStatementListener) ExitBitmapJoinIndexClause(ctx *BitmapJoinIndexClauseContext) {}

// EnterColumnSortsClause_ is called when production columnSortsClause_ is entered.
func (s *BaseOracleStatementListener) EnterColumnSortsClause_(ctx *ColumnSortsClause_Context) {}

// ExitColumnSortsClause_ is called when production columnSortsClause_ is exited.
func (s *BaseOracleStatementListener) ExitColumnSortsClause_(ctx *ColumnSortsClause_Context) {}

// EnterColumnSortClause_ is called when production columnSortClause_ is entered.
func (s *BaseOracleStatementListener) EnterColumnSortClause_(ctx *ColumnSortClause_Context) {}

// ExitColumnSortClause_ is called when production columnSortClause_ is exited.
func (s *BaseOracleStatementListener) ExitColumnSortClause_(ctx *ColumnSortClause_Context) {}

// EnterCreateIndexDefinitionClause is called when production createIndexDefinitionClause is entered.
func (s *BaseOracleStatementListener) EnterCreateIndexDefinitionClause(ctx *CreateIndexDefinitionClauseContext) {
}

// ExitCreateIndexDefinitionClause is called when production createIndexDefinitionClause is exited.
func (s *BaseOracleStatementListener) ExitCreateIndexDefinitionClause(ctx *CreateIndexDefinitionClauseContext) {
}

// EnterTableAlias is called when production tableAlias is entered.
func (s *BaseOracleStatementListener) EnterTableAlias(ctx *TableAliasContext) {}

// ExitTableAlias is called when production tableAlias is exited.
func (s *BaseOracleStatementListener) ExitTableAlias(ctx *TableAliasContext) {}

// EnterAlterDefinitionClause is called when production alterDefinitionClause is entered.
func (s *BaseOracleStatementListener) EnterAlterDefinitionClause(ctx *AlterDefinitionClauseContext) {}

// ExitAlterDefinitionClause is called when production alterDefinitionClause is exited.
func (s *BaseOracleStatementListener) ExitAlterDefinitionClause(ctx *AlterDefinitionClauseContext) {}

// EnterAlterTableProperties is called when production alterTableProperties is entered.
func (s *BaseOracleStatementListener) EnterAlterTableProperties(ctx *AlterTablePropertiesContext) {}

// ExitAlterTableProperties is called when production alterTableProperties is exited.
func (s *BaseOracleStatementListener) ExitAlterTableProperties(ctx *AlterTablePropertiesContext) {}

// EnterRenameTableSpecification is called when production renameTableSpecification is entered.
func (s *BaseOracleStatementListener) EnterRenameTableSpecification(ctx *RenameTableSpecificationContext) {
}

// ExitRenameTableSpecification is called when production renameTableSpecification is exited.
func (s *BaseOracleStatementListener) ExitRenameTableSpecification(ctx *RenameTableSpecificationContext) {
}

// EnterColumnClauses is called when production columnClauses is entered.
func (s *BaseOracleStatementListener) EnterColumnClauses(ctx *ColumnClausesContext) {}

// ExitColumnClauses is called when production columnClauses is exited.
func (s *BaseOracleStatementListener) ExitColumnClauses(ctx *ColumnClausesContext) {}

// EnterOperateColumnClause is called when production operateColumnClause is entered.
func (s *BaseOracleStatementListener) EnterOperateColumnClause(ctx *OperateColumnClauseContext) {}

// ExitOperateColumnClause is called when production operateColumnClause is exited.
func (s *BaseOracleStatementListener) ExitOperateColumnClause(ctx *OperateColumnClauseContext) {}

// EnterAddColumnSpecification is called when production addColumnSpecification is entered.
func (s *BaseOracleStatementListener) EnterAddColumnSpecification(ctx *AddColumnSpecificationContext) {
}

// ExitAddColumnSpecification is called when production addColumnSpecification is exited.
func (s *BaseOracleStatementListener) ExitAddColumnSpecification(ctx *AddColumnSpecificationContext) {
}

// EnterColumnOrVirtualDefinitions is called when production columnOrVirtualDefinitions is entered.
func (s *BaseOracleStatementListener) EnterColumnOrVirtualDefinitions(ctx *ColumnOrVirtualDefinitionsContext) {
}

// ExitColumnOrVirtualDefinitions is called when production columnOrVirtualDefinitions is exited.
func (s *BaseOracleStatementListener) ExitColumnOrVirtualDefinitions(ctx *ColumnOrVirtualDefinitionsContext) {
}

// EnterColumnOrVirtualDefinition is called when production columnOrVirtualDefinition is entered.
func (s *BaseOracleStatementListener) EnterColumnOrVirtualDefinition(ctx *ColumnOrVirtualDefinitionContext) {
}

// ExitColumnOrVirtualDefinition is called when production columnOrVirtualDefinition is exited.
func (s *BaseOracleStatementListener) ExitColumnOrVirtualDefinition(ctx *ColumnOrVirtualDefinitionContext) {
}

// EnterColumnProperties is called when production columnProperties is entered.
func (s *BaseOracleStatementListener) EnterColumnProperties(ctx *ColumnPropertiesContext) {}

// ExitColumnProperties is called when production columnProperties is exited.
func (s *BaseOracleStatementListener) ExitColumnProperties(ctx *ColumnPropertiesContext) {}

// EnterColumnProperty is called when production columnProperty is entered.
func (s *BaseOracleStatementListener) EnterColumnProperty(ctx *ColumnPropertyContext) {}

// ExitColumnProperty is called when production columnProperty is exited.
func (s *BaseOracleStatementListener) ExitColumnProperty(ctx *ColumnPropertyContext) {}

// EnterObjectTypeColProperties is called when production objectTypeColProperties is entered.
func (s *BaseOracleStatementListener) EnterObjectTypeColProperties(ctx *ObjectTypeColPropertiesContext) {
}

// ExitObjectTypeColProperties is called when production objectTypeColProperties is exited.
func (s *BaseOracleStatementListener) ExitObjectTypeColProperties(ctx *ObjectTypeColPropertiesContext) {
}

// EnterSubstitutableColumnClause is called when production substitutableColumnClause is entered.
func (s *BaseOracleStatementListener) EnterSubstitutableColumnClause(ctx *SubstitutableColumnClauseContext) {
}

// ExitSubstitutableColumnClause is called when production substitutableColumnClause is exited.
func (s *BaseOracleStatementListener) ExitSubstitutableColumnClause(ctx *SubstitutableColumnClauseContext) {
}

// EnterModifyColumnSpecification is called when production modifyColumnSpecification is entered.
func (s *BaseOracleStatementListener) EnterModifyColumnSpecification(ctx *ModifyColumnSpecificationContext) {
}

// ExitModifyColumnSpecification is called when production modifyColumnSpecification is exited.
func (s *BaseOracleStatementListener) ExitModifyColumnSpecification(ctx *ModifyColumnSpecificationContext) {
}

// EnterModifyColProperties is called when production modifyColProperties is entered.
func (s *BaseOracleStatementListener) EnterModifyColProperties(ctx *ModifyColPropertiesContext) {}

// ExitModifyColProperties is called when production modifyColProperties is exited.
func (s *BaseOracleStatementListener) ExitModifyColProperties(ctx *ModifyColPropertiesContext) {}

// EnterModifyColSubstitutable is called when production modifyColSubstitutable is entered.
func (s *BaseOracleStatementListener) EnterModifyColSubstitutable(ctx *ModifyColSubstitutableContext) {
}

// ExitModifyColSubstitutable is called when production modifyColSubstitutable is exited.
func (s *BaseOracleStatementListener) ExitModifyColSubstitutable(ctx *ModifyColSubstitutableContext) {
}

// EnterDropColumnClause is called when production dropColumnClause is entered.
func (s *BaseOracleStatementListener) EnterDropColumnClause(ctx *DropColumnClauseContext) {}

// ExitDropColumnClause is called when production dropColumnClause is exited.
func (s *BaseOracleStatementListener) ExitDropColumnClause(ctx *DropColumnClauseContext) {}

// EnterDropColumnSpecification is called when production dropColumnSpecification is entered.
func (s *BaseOracleStatementListener) EnterDropColumnSpecification(ctx *DropColumnSpecificationContext) {
}

// ExitDropColumnSpecification is called when production dropColumnSpecification is exited.
func (s *BaseOracleStatementListener) ExitDropColumnSpecification(ctx *DropColumnSpecificationContext) {
}

// EnterColumnOrColumnList is called when production columnOrColumnList is entered.
func (s *BaseOracleStatementListener) EnterColumnOrColumnList(ctx *ColumnOrColumnListContext) {}

// ExitColumnOrColumnList is called when production columnOrColumnList is exited.
func (s *BaseOracleStatementListener) ExitColumnOrColumnList(ctx *ColumnOrColumnListContext) {}

// EnterCascadeOrInvalidate is called when production cascadeOrInvalidate is entered.
func (s *BaseOracleStatementListener) EnterCascadeOrInvalidate(ctx *CascadeOrInvalidateContext) {}

// ExitCascadeOrInvalidate is called when production cascadeOrInvalidate is exited.
func (s *BaseOracleStatementListener) ExitCascadeOrInvalidate(ctx *CascadeOrInvalidateContext) {}

// EnterCheckpointNumber is called when production checkpointNumber is entered.
func (s *BaseOracleStatementListener) EnterCheckpointNumber(ctx *CheckpointNumberContext) {}

// ExitCheckpointNumber is called when production checkpointNumber is exited.
func (s *BaseOracleStatementListener) ExitCheckpointNumber(ctx *CheckpointNumberContext) {}

// EnterRenameColumnClause is called when production renameColumnClause is entered.
func (s *BaseOracleStatementListener) EnterRenameColumnClause(ctx *RenameColumnClauseContext) {}

// ExitRenameColumnClause is called when production renameColumnClause is exited.
func (s *BaseOracleStatementListener) ExitRenameColumnClause(ctx *RenameColumnClauseContext) {}

// EnterConstraintClauses is called when production constraintClauses is entered.
func (s *BaseOracleStatementListener) EnterConstraintClauses(ctx *ConstraintClausesContext) {}

// ExitConstraintClauses is called when production constraintClauses is exited.
func (s *BaseOracleStatementListener) ExitConstraintClauses(ctx *ConstraintClausesContext) {}

// EnterAddConstraintSpecification is called when production addConstraintSpecification is entered.
func (s *BaseOracleStatementListener) EnterAddConstraintSpecification(ctx *AddConstraintSpecificationContext) {
}

// ExitAddConstraintSpecification is called when production addConstraintSpecification is exited.
func (s *BaseOracleStatementListener) ExitAddConstraintSpecification(ctx *AddConstraintSpecificationContext) {
}

// EnterModifyConstraintClause is called when production modifyConstraintClause is entered.
func (s *BaseOracleStatementListener) EnterModifyConstraintClause(ctx *ModifyConstraintClauseContext) {
}

// ExitModifyConstraintClause is called when production modifyConstraintClause is exited.
func (s *BaseOracleStatementListener) ExitModifyConstraintClause(ctx *ModifyConstraintClauseContext) {
}

// EnterConstraintWithName is called when production constraintWithName is entered.
func (s *BaseOracleStatementListener) EnterConstraintWithName(ctx *ConstraintWithNameContext) {}

// ExitConstraintWithName is called when production constraintWithName is exited.
func (s *BaseOracleStatementListener) ExitConstraintWithName(ctx *ConstraintWithNameContext) {}

// EnterConstraintOption is called when production constraintOption is entered.
func (s *BaseOracleStatementListener) EnterConstraintOption(ctx *ConstraintOptionContext) {}

// ExitConstraintOption is called when production constraintOption is exited.
func (s *BaseOracleStatementListener) ExitConstraintOption(ctx *ConstraintOptionContext) {}

// EnterConstraintPrimaryOrUnique is called when production constraintPrimaryOrUnique is entered.
func (s *BaseOracleStatementListener) EnterConstraintPrimaryOrUnique(ctx *ConstraintPrimaryOrUniqueContext) {
}

// ExitConstraintPrimaryOrUnique is called when production constraintPrimaryOrUnique is exited.
func (s *BaseOracleStatementListener) ExitConstraintPrimaryOrUnique(ctx *ConstraintPrimaryOrUniqueContext) {
}

// EnterRenameConstraintClause is called when production renameConstraintClause is entered.
func (s *BaseOracleStatementListener) EnterRenameConstraintClause(ctx *RenameConstraintClauseContext) {
}

// ExitRenameConstraintClause is called when production renameConstraintClause is exited.
func (s *BaseOracleStatementListener) ExitRenameConstraintClause(ctx *RenameConstraintClauseContext) {
}

// EnterDropConstraintClause is called when production dropConstraintClause is entered.
func (s *BaseOracleStatementListener) EnterDropConstraintClause(ctx *DropConstraintClauseContext) {}

// ExitDropConstraintClause is called when production dropConstraintClause is exited.
func (s *BaseOracleStatementListener) ExitDropConstraintClause(ctx *DropConstraintClauseContext) {}

// EnterAlterExternalTable is called when production alterExternalTable is entered.
func (s *BaseOracleStatementListener) EnterAlterExternalTable(ctx *AlterExternalTableContext) {}

// ExitAlterExternalTable is called when production alterExternalTable is exited.
func (s *BaseOracleStatementListener) ExitAlterExternalTable(ctx *AlterExternalTableContext) {}

// EnterObjectProperties is called when production objectProperties is entered.
func (s *BaseOracleStatementListener) EnterObjectProperties(ctx *ObjectPropertiesContext) {}

// ExitObjectProperties is called when production objectProperties is exited.
func (s *BaseOracleStatementListener) ExitObjectProperties(ctx *ObjectPropertiesContext) {}

// EnterAlterIndexInformationClause is called when production alterIndexInformationClause is entered.
func (s *BaseOracleStatementListener) EnterAlterIndexInformationClause(ctx *AlterIndexInformationClauseContext) {
}

// ExitAlterIndexInformationClause is called when production alterIndexInformationClause is exited.
func (s *BaseOracleStatementListener) ExitAlterIndexInformationClause(ctx *AlterIndexInformationClauseContext) {
}

// EnterRenameIndexClause is called when production renameIndexClause is entered.
func (s *BaseOracleStatementListener) EnterRenameIndexClause(ctx *RenameIndexClauseContext) {}

// ExitRenameIndexClause is called when production renameIndexClause is exited.
func (s *BaseOracleStatementListener) ExitRenameIndexClause(ctx *RenameIndexClauseContext) {}

// EnterObjectTableSubstitution is called when production objectTableSubstitution is entered.
func (s *BaseOracleStatementListener) EnterObjectTableSubstitution(ctx *ObjectTableSubstitutionContext) {
}

// ExitObjectTableSubstitution is called when production objectTableSubstitution is exited.
func (s *BaseOracleStatementListener) ExitObjectTableSubstitution(ctx *ObjectTableSubstitutionContext) {
}

// EnterMemOptimizeClause is called when production memOptimizeClause is entered.
func (s *BaseOracleStatementListener) EnterMemOptimizeClause(ctx *MemOptimizeClauseContext) {}

// ExitMemOptimizeClause is called when production memOptimizeClause is exited.
func (s *BaseOracleStatementListener) ExitMemOptimizeClause(ctx *MemOptimizeClauseContext) {}

// EnterMemOptimizeReadClause is called when production memOptimizeReadClause is entered.
func (s *BaseOracleStatementListener) EnterMemOptimizeReadClause(ctx *MemOptimizeReadClauseContext) {}

// ExitMemOptimizeReadClause is called when production memOptimizeReadClause is exited.
func (s *BaseOracleStatementListener) ExitMemOptimizeReadClause(ctx *MemOptimizeReadClauseContext) {}

// EnterMemOptimizeWriteClause is called when production memOptimizeWriteClause is entered.
func (s *BaseOracleStatementListener) EnterMemOptimizeWriteClause(ctx *MemOptimizeWriteClauseContext) {
}

// ExitMemOptimizeWriteClause is called when production memOptimizeWriteClause is exited.
func (s *BaseOracleStatementListener) ExitMemOptimizeWriteClause(ctx *MemOptimizeWriteClauseContext) {
}

// EnterEnableDisableClauses is called when production enableDisableClauses is entered.
func (s *BaseOracleStatementListener) EnterEnableDisableClauses(ctx *EnableDisableClausesContext) {}

// ExitEnableDisableClauses is called when production enableDisableClauses is exited.
func (s *BaseOracleStatementListener) ExitEnableDisableClauses(ctx *EnableDisableClausesContext) {}

// EnterEnableDisableClause is called when production enableDisableClause is entered.
func (s *BaseOracleStatementListener) EnterEnableDisableClause(ctx *EnableDisableClauseContext) {}

// ExitEnableDisableClause is called when production enableDisableClause is exited.
func (s *BaseOracleStatementListener) ExitEnableDisableClause(ctx *EnableDisableClauseContext) {}

// EnterEnableDisableOthers is called when production enableDisableOthers is entered.
func (s *BaseOracleStatementListener) EnterEnableDisableOthers(ctx *EnableDisableOthersContext) {}

// ExitEnableDisableOthers is called when production enableDisableOthers is exited.
func (s *BaseOracleStatementListener) ExitEnableDisableOthers(ctx *EnableDisableOthersContext) {}

// EnterRebuildClause is called when production rebuildClause is entered.
func (s *BaseOracleStatementListener) EnterRebuildClause(ctx *RebuildClauseContext) {}

// ExitRebuildClause is called when production rebuildClause is exited.
func (s *BaseOracleStatementListener) ExitRebuildClause(ctx *RebuildClauseContext) {}

// EnterParallelClause is called when production parallelClause is entered.
func (s *BaseOracleStatementListener) EnterParallelClause(ctx *ParallelClauseContext) {}

// ExitParallelClause is called when production parallelClause is exited.
func (s *BaseOracleStatementListener) ExitParallelClause(ctx *ParallelClauseContext) {}

// EnterUsableSpecification is called when production usableSpecification is entered.
func (s *BaseOracleStatementListener) EnterUsableSpecification(ctx *UsableSpecificationContext) {}

// ExitUsableSpecification is called when production usableSpecification is exited.
func (s *BaseOracleStatementListener) ExitUsableSpecification(ctx *UsableSpecificationContext) {}

// EnterInvalidationSpecification is called when production invalidationSpecification is entered.
func (s *BaseOracleStatementListener) EnterInvalidationSpecification(ctx *InvalidationSpecificationContext) {
}

// ExitInvalidationSpecification is called when production invalidationSpecification is exited.
func (s *BaseOracleStatementListener) ExitInvalidationSpecification(ctx *InvalidationSpecificationContext) {
}

// EnterMaterializedViewLogClause is called when production materializedViewLogClause is entered.
func (s *BaseOracleStatementListener) EnterMaterializedViewLogClause(ctx *MaterializedViewLogClauseContext) {
}

// ExitMaterializedViewLogClause is called when production materializedViewLogClause is exited.
func (s *BaseOracleStatementListener) ExitMaterializedViewLogClause(ctx *MaterializedViewLogClauseContext) {
}

// EnterDropReuseClause is called when production dropReuseClause is entered.
func (s *BaseOracleStatementListener) EnterDropReuseClause(ctx *DropReuseClauseContext) {}

// ExitDropReuseClause is called when production dropReuseClause is exited.
func (s *BaseOracleStatementListener) ExitDropReuseClause(ctx *DropReuseClauseContext) {}

// EnterCollationClause is called when production collationClause is entered.
func (s *BaseOracleStatementListener) EnterCollationClause(ctx *CollationClauseContext) {}

// ExitCollationClause is called when production collationClause is exited.
func (s *BaseOracleStatementListener) ExitCollationClause(ctx *CollationClauseContext) {}

// EnterCommitClause is called when production commitClause is entered.
func (s *BaseOracleStatementListener) EnterCommitClause(ctx *CommitClauseContext) {}

// ExitCommitClause is called when production commitClause is exited.
func (s *BaseOracleStatementListener) ExitCommitClause(ctx *CommitClauseContext) {}

// EnterPhysicalProperties is called when production physicalProperties is entered.
func (s *BaseOracleStatementListener) EnterPhysicalProperties(ctx *PhysicalPropertiesContext) {}

// ExitPhysicalProperties is called when production physicalProperties is exited.
func (s *BaseOracleStatementListener) ExitPhysicalProperties(ctx *PhysicalPropertiesContext) {}

// EnterDeferredSegmentCreation is called when production deferredSegmentCreation is entered.
func (s *BaseOracleStatementListener) EnterDeferredSegmentCreation(ctx *DeferredSegmentCreationContext) {
}

// ExitDeferredSegmentCreation is called when production deferredSegmentCreation is exited.
func (s *BaseOracleStatementListener) ExitDeferredSegmentCreation(ctx *DeferredSegmentCreationContext) {
}

// EnterSegmentAttributesClause is called when production segmentAttributesClause is entered.
func (s *BaseOracleStatementListener) EnterSegmentAttributesClause(ctx *SegmentAttributesClauseContext) {
}

// ExitSegmentAttributesClause is called when production segmentAttributesClause is exited.
func (s *BaseOracleStatementListener) ExitSegmentAttributesClause(ctx *SegmentAttributesClauseContext) {
}

// EnterPhysicalAttributesClause is called when production physicalAttributesClause is entered.
func (s *BaseOracleStatementListener) EnterPhysicalAttributesClause(ctx *PhysicalAttributesClauseContext) {
}

// ExitPhysicalAttributesClause is called when production physicalAttributesClause is exited.
func (s *BaseOracleStatementListener) ExitPhysicalAttributesClause(ctx *PhysicalAttributesClauseContext) {
}

// EnterLoggingClause is called when production loggingClause is entered.
func (s *BaseOracleStatementListener) EnterLoggingClause(ctx *LoggingClauseContext) {}

// ExitLoggingClause is called when production loggingClause is exited.
func (s *BaseOracleStatementListener) ExitLoggingClause(ctx *LoggingClauseContext) {}

// EnterStorageClause is called when production storageClause is entered.
func (s *BaseOracleStatementListener) EnterStorageClause(ctx *StorageClauseContext) {}

// ExitStorageClause is called when production storageClause is exited.
func (s *BaseOracleStatementListener) ExitStorageClause(ctx *StorageClauseContext) {}

// EnterSizeClause is called when production sizeClause is entered.
func (s *BaseOracleStatementListener) EnterSizeClause(ctx *SizeClauseContext) {}

// ExitSizeClause is called when production sizeClause is exited.
func (s *BaseOracleStatementListener) ExitSizeClause(ctx *SizeClauseContext) {}

// EnterMaxsizeClause is called when production maxsizeClause is entered.
func (s *BaseOracleStatementListener) EnterMaxsizeClause(ctx *MaxsizeClauseContext) {}

// ExitMaxsizeClause is called when production maxsizeClause is exited.
func (s *BaseOracleStatementListener) ExitMaxsizeClause(ctx *MaxsizeClauseContext) {}

// EnterTableCompression is called when production tableCompression is entered.
func (s *BaseOracleStatementListener) EnterTableCompression(ctx *TableCompressionContext) {}

// ExitTableCompression is called when production tableCompression is exited.
func (s *BaseOracleStatementListener) ExitTableCompression(ctx *TableCompressionContext) {}

// EnterInmemoryTableClause is called when production inmemoryTableClause is entered.
func (s *BaseOracleStatementListener) EnterInmemoryTableClause(ctx *InmemoryTableClauseContext) {}

// ExitInmemoryTableClause is called when production inmemoryTableClause is exited.
func (s *BaseOracleStatementListener) ExitInmemoryTableClause(ctx *InmemoryTableClauseContext) {}

// EnterInmemoryAttributes is called when production inmemoryAttributes is entered.
func (s *BaseOracleStatementListener) EnterInmemoryAttributes(ctx *InmemoryAttributesContext) {}

// ExitInmemoryAttributes is called when production inmemoryAttributes is exited.
func (s *BaseOracleStatementListener) ExitInmemoryAttributes(ctx *InmemoryAttributesContext) {}

// EnterInmemoryColumnClause is called when production inmemoryColumnClause is entered.
func (s *BaseOracleStatementListener) EnterInmemoryColumnClause(ctx *InmemoryColumnClauseContext) {}

// ExitInmemoryColumnClause is called when production inmemoryColumnClause is exited.
func (s *BaseOracleStatementListener) ExitInmemoryColumnClause(ctx *InmemoryColumnClauseContext) {}

// EnterInmemoryMemcompress is called when production inmemoryMemcompress is entered.
func (s *BaseOracleStatementListener) EnterInmemoryMemcompress(ctx *InmemoryMemcompressContext) {}

// ExitInmemoryMemcompress is called when production inmemoryMemcompress is exited.
func (s *BaseOracleStatementListener) ExitInmemoryMemcompress(ctx *InmemoryMemcompressContext) {}

// EnterInmemoryPriority is called when production inmemoryPriority is entered.
func (s *BaseOracleStatementListener) EnterInmemoryPriority(ctx *InmemoryPriorityContext) {}

// ExitInmemoryPriority is called when production inmemoryPriority is exited.
func (s *BaseOracleStatementListener) ExitInmemoryPriority(ctx *InmemoryPriorityContext) {}

// EnterInmemoryDistribute is called when production inmemoryDistribute is entered.
func (s *BaseOracleStatementListener) EnterInmemoryDistribute(ctx *InmemoryDistributeContext) {}

// ExitInmemoryDistribute is called when production inmemoryDistribute is exited.
func (s *BaseOracleStatementListener) ExitInmemoryDistribute(ctx *InmemoryDistributeContext) {}

// EnterInmemoryDuplicate is called when production inmemoryDuplicate is entered.
func (s *BaseOracleStatementListener) EnterInmemoryDuplicate(ctx *InmemoryDuplicateContext) {}

// ExitInmemoryDuplicate is called when production inmemoryDuplicate is exited.
func (s *BaseOracleStatementListener) ExitInmemoryDuplicate(ctx *InmemoryDuplicateContext) {}

// EnterIlmClause is called when production ilmClause is entered.
func (s *BaseOracleStatementListener) EnterIlmClause(ctx *IlmClauseContext) {}

// ExitIlmClause is called when production ilmClause is exited.
func (s *BaseOracleStatementListener) ExitIlmClause(ctx *IlmClauseContext) {}

// EnterIlmPolicyClause is called when production ilmPolicyClause is entered.
func (s *BaseOracleStatementListener) EnterIlmPolicyClause(ctx *IlmPolicyClauseContext) {}

// ExitIlmPolicyClause is called when production ilmPolicyClause is exited.
func (s *BaseOracleStatementListener) ExitIlmPolicyClause(ctx *IlmPolicyClauseContext) {}

// EnterIlmCompressionPolicy is called when production ilmCompressionPolicy is entered.
func (s *BaseOracleStatementListener) EnterIlmCompressionPolicy(ctx *IlmCompressionPolicyContext) {}

// ExitIlmCompressionPolicy is called when production ilmCompressionPolicy is exited.
func (s *BaseOracleStatementListener) ExitIlmCompressionPolicy(ctx *IlmCompressionPolicyContext) {}

// EnterIlmTimePeriod is called when production ilmTimePeriod is entered.
func (s *BaseOracleStatementListener) EnterIlmTimePeriod(ctx *IlmTimePeriodContext) {}

// ExitIlmTimePeriod is called when production ilmTimePeriod is exited.
func (s *BaseOracleStatementListener) ExitIlmTimePeriod(ctx *IlmTimePeriodContext) {}

// EnterIlmTieringPolicy is called when production ilmTieringPolicy is entered.
func (s *BaseOracleStatementListener) EnterIlmTieringPolicy(ctx *IlmTieringPolicyContext) {}

// ExitIlmTieringPolicy is called when production ilmTieringPolicy is exited.
func (s *BaseOracleStatementListener) ExitIlmTieringPolicy(ctx *IlmTieringPolicyContext) {}

// EnterIlmInmemoryPolicy is called when production ilmInmemoryPolicy is entered.
func (s *BaseOracleStatementListener) EnterIlmInmemoryPolicy(ctx *IlmInmemoryPolicyContext) {}

// ExitIlmInmemoryPolicy is called when production ilmInmemoryPolicy is exited.
func (s *BaseOracleStatementListener) ExitIlmInmemoryPolicy(ctx *IlmInmemoryPolicyContext) {}

// EnterOrganizationClause is called when production organizationClause is entered.
func (s *BaseOracleStatementListener) EnterOrganizationClause(ctx *OrganizationClauseContext) {}

// ExitOrganizationClause is called when production organizationClause is exited.
func (s *BaseOracleStatementListener) ExitOrganizationClause(ctx *OrganizationClauseContext) {}

// EnterHeapOrgTableClause is called when production heapOrgTableClause is entered.
func (s *BaseOracleStatementListener) EnterHeapOrgTableClause(ctx *HeapOrgTableClauseContext) {}

// ExitHeapOrgTableClause is called when production heapOrgTableClause is exited.
func (s *BaseOracleStatementListener) ExitHeapOrgTableClause(ctx *HeapOrgTableClauseContext) {}

// EnterIndexOrgTableClause is called when production indexOrgTableClause is entered.
func (s *BaseOracleStatementListener) EnterIndexOrgTableClause(ctx *IndexOrgTableClauseContext) {}

// ExitIndexOrgTableClause is called when production indexOrgTableClause is exited.
func (s *BaseOracleStatementListener) ExitIndexOrgTableClause(ctx *IndexOrgTableClauseContext) {}

// EnterExternalTableClause is called when production externalTableClause is entered.
func (s *BaseOracleStatementListener) EnterExternalTableClause(ctx *ExternalTableClauseContext) {}

// ExitExternalTableClause is called when production externalTableClause is exited.
func (s *BaseOracleStatementListener) ExitExternalTableClause(ctx *ExternalTableClauseContext) {}

// EnterExternalTableDataProps is called when production externalTableDataProps is entered.
func (s *BaseOracleStatementListener) EnterExternalTableDataProps(ctx *ExternalTableDataPropsContext) {
}

// ExitExternalTableDataProps is called when production externalTableDataProps is exited.
func (s *BaseOracleStatementListener) ExitExternalTableDataProps(ctx *ExternalTableDataPropsContext) {
}

// EnterMappingTableClause is called when production mappingTableClause is entered.
func (s *BaseOracleStatementListener) EnterMappingTableClause(ctx *MappingTableClauseContext) {}

// ExitMappingTableClause is called when production mappingTableClause is exited.
func (s *BaseOracleStatementListener) ExitMappingTableClause(ctx *MappingTableClauseContext) {}

// EnterPrefixCompression is called when production prefixCompression is entered.
func (s *BaseOracleStatementListener) EnterPrefixCompression(ctx *PrefixCompressionContext) {}

// ExitPrefixCompression is called when production prefixCompression is exited.
func (s *BaseOracleStatementListener) ExitPrefixCompression(ctx *PrefixCompressionContext) {}

// EnterIndexOrgOverflowClause is called when production indexOrgOverflowClause is entered.
func (s *BaseOracleStatementListener) EnterIndexOrgOverflowClause(ctx *IndexOrgOverflowClauseContext) {
}

// ExitIndexOrgOverflowClause is called when production indexOrgOverflowClause is exited.
func (s *BaseOracleStatementListener) ExitIndexOrgOverflowClause(ctx *IndexOrgOverflowClauseContext) {
}

// EnterExternalPartitionClause is called when production externalPartitionClause is entered.
func (s *BaseOracleStatementListener) EnterExternalPartitionClause(ctx *ExternalPartitionClauseContext) {
}

// ExitExternalPartitionClause is called when production externalPartitionClause is exited.
func (s *BaseOracleStatementListener) ExitExternalPartitionClause(ctx *ExternalPartitionClauseContext) {
}

// EnterClusterRelatedClause is called when production clusterRelatedClause is entered.
func (s *BaseOracleStatementListener) EnterClusterRelatedClause(ctx *ClusterRelatedClauseContext) {}

// ExitClusterRelatedClause is called when production clusterRelatedClause is exited.
func (s *BaseOracleStatementListener) ExitClusterRelatedClause(ctx *ClusterRelatedClauseContext) {}

// EnterTableProperties is called when production tableProperties is entered.
func (s *BaseOracleStatementListener) EnterTableProperties(ctx *TablePropertiesContext) {}

// ExitTableProperties is called when production tableProperties is exited.
func (s *BaseOracleStatementListener) ExitTableProperties(ctx *TablePropertiesContext) {}

// EnterReadOnlyClause is called when production readOnlyClause is entered.
func (s *BaseOracleStatementListener) EnterReadOnlyClause(ctx *ReadOnlyClauseContext) {}

// ExitReadOnlyClause is called when production readOnlyClause is exited.
func (s *BaseOracleStatementListener) ExitReadOnlyClause(ctx *ReadOnlyClauseContext) {}

// EnterIndexingClause is called when production indexingClause is entered.
func (s *BaseOracleStatementListener) EnterIndexingClause(ctx *IndexingClauseContext) {}

// ExitIndexingClause is called when production indexingClause is exited.
func (s *BaseOracleStatementListener) ExitIndexingClause(ctx *IndexingClauseContext) {}

// EnterTablePartitioningClauses is called when production tablePartitioningClauses is entered.
func (s *BaseOracleStatementListener) EnterTablePartitioningClauses(ctx *TablePartitioningClausesContext) {
}

// ExitTablePartitioningClauses is called when production tablePartitioningClauses is exited.
func (s *BaseOracleStatementListener) ExitTablePartitioningClauses(ctx *TablePartitioningClausesContext) {
}

// EnterRangePartitions is called when production rangePartitions is entered.
func (s *BaseOracleStatementListener) EnterRangePartitions(ctx *RangePartitionsContext) {}

// ExitRangePartitions is called when production rangePartitions is exited.
func (s *BaseOracleStatementListener) ExitRangePartitions(ctx *RangePartitionsContext) {}

// EnterRangeValuesClause is called when production rangeValuesClause is entered.
func (s *BaseOracleStatementListener) EnterRangeValuesClause(ctx *RangeValuesClauseContext) {}

// ExitRangeValuesClause is called when production rangeValuesClause is exited.
func (s *BaseOracleStatementListener) ExitRangeValuesClause(ctx *RangeValuesClauseContext) {}

// EnterTablePartitionDescription is called when production tablePartitionDescription is entered.
func (s *BaseOracleStatementListener) EnterTablePartitionDescription(ctx *TablePartitionDescriptionContext) {
}

// ExitTablePartitionDescription is called when production tablePartitionDescription is exited.
func (s *BaseOracleStatementListener) ExitTablePartitionDescription(ctx *TablePartitionDescriptionContext) {
}

// EnterInmemoryClause is called when production inmemoryClause is entered.
func (s *BaseOracleStatementListener) EnterInmemoryClause(ctx *InmemoryClauseContext) {}

// ExitInmemoryClause is called when production inmemoryClause is exited.
func (s *BaseOracleStatementListener) ExitInmemoryClause(ctx *InmemoryClauseContext) {}

// EnterVarrayColProperties is called when production varrayColProperties is entered.
func (s *BaseOracleStatementListener) EnterVarrayColProperties(ctx *VarrayColPropertiesContext) {}

// ExitVarrayColProperties is called when production varrayColProperties is exited.
func (s *BaseOracleStatementListener) ExitVarrayColProperties(ctx *VarrayColPropertiesContext) {}

// EnterNestedTableColProperties is called when production nestedTableColProperties is entered.
func (s *BaseOracleStatementListener) EnterNestedTableColProperties(ctx *NestedTableColPropertiesContext) {
}

// ExitNestedTableColProperties is called when production nestedTableColProperties is exited.
func (s *BaseOracleStatementListener) ExitNestedTableColProperties(ctx *NestedTableColPropertiesContext) {
}

// EnterLobStorageClause is called when production lobStorageClause is entered.
func (s *BaseOracleStatementListener) EnterLobStorageClause(ctx *LobStorageClauseContext) {}

// ExitLobStorageClause is called when production lobStorageClause is exited.
func (s *BaseOracleStatementListener) ExitLobStorageClause(ctx *LobStorageClauseContext) {}

// EnterVarrayStorageClause is called when production varrayStorageClause is entered.
func (s *BaseOracleStatementListener) EnterVarrayStorageClause(ctx *VarrayStorageClauseContext) {}

// ExitVarrayStorageClause is called when production varrayStorageClause is exited.
func (s *BaseOracleStatementListener) ExitVarrayStorageClause(ctx *VarrayStorageClauseContext) {}

// EnterLobStorageParameters is called when production lobStorageParameters is entered.
func (s *BaseOracleStatementListener) EnterLobStorageParameters(ctx *LobStorageParametersContext) {}

// ExitLobStorageParameters is called when production lobStorageParameters is exited.
func (s *BaseOracleStatementListener) ExitLobStorageParameters(ctx *LobStorageParametersContext) {}

// EnterLobParameters is called when production lobParameters is entered.
func (s *BaseOracleStatementListener) EnterLobParameters(ctx *LobParametersContext) {}

// ExitLobParameters is called when production lobParameters is exited.
func (s *BaseOracleStatementListener) ExitLobParameters(ctx *LobParametersContext) {}

// EnterLobRetentionClause is called when production lobRetentionClause is entered.
func (s *BaseOracleStatementListener) EnterLobRetentionClause(ctx *LobRetentionClauseContext) {}

// ExitLobRetentionClause is called when production lobRetentionClause is exited.
func (s *BaseOracleStatementListener) ExitLobRetentionClause(ctx *LobRetentionClauseContext) {}

// EnterLobDeduplicateClause is called when production lobDeduplicateClause is entered.
func (s *BaseOracleStatementListener) EnterLobDeduplicateClause(ctx *LobDeduplicateClauseContext) {}

// ExitLobDeduplicateClause is called when production lobDeduplicateClause is exited.
func (s *BaseOracleStatementListener) ExitLobDeduplicateClause(ctx *LobDeduplicateClauseContext) {}

// EnterLobCompressionClause is called when production lobCompressionClause is entered.
func (s *BaseOracleStatementListener) EnterLobCompressionClause(ctx *LobCompressionClauseContext) {}

// ExitLobCompressionClause is called when production lobCompressionClause is exited.
func (s *BaseOracleStatementListener) ExitLobCompressionClause(ctx *LobCompressionClauseContext) {}

// EnterExternalPartSubpartDataProps is called when production externalPartSubpartDataProps is entered.
func (s *BaseOracleStatementListener) EnterExternalPartSubpartDataProps(ctx *ExternalPartSubpartDataPropsContext) {
}

// ExitExternalPartSubpartDataProps is called when production externalPartSubpartDataProps is exited.
func (s *BaseOracleStatementListener) ExitExternalPartSubpartDataProps(ctx *ExternalPartSubpartDataPropsContext) {
}

// EnterListPartitions is called when production listPartitions is entered.
func (s *BaseOracleStatementListener) EnterListPartitions(ctx *ListPartitionsContext) {}

// ExitListPartitions is called when production listPartitions is exited.
func (s *BaseOracleStatementListener) ExitListPartitions(ctx *ListPartitionsContext) {}

// EnterListValuesClause is called when production listValuesClause is entered.
func (s *BaseOracleStatementListener) EnterListValuesClause(ctx *ListValuesClauseContext) {}

// ExitListValuesClause is called when production listValuesClause is exited.
func (s *BaseOracleStatementListener) ExitListValuesClause(ctx *ListValuesClauseContext) {}

// EnterListValues is called when production listValues is entered.
func (s *BaseOracleStatementListener) EnterListValues(ctx *ListValuesContext) {}

// ExitListValues is called when production listValues is exited.
func (s *BaseOracleStatementListener) ExitListValues(ctx *ListValuesContext) {}

// EnterHashPartitions is called when production hashPartitions is entered.
func (s *BaseOracleStatementListener) EnterHashPartitions(ctx *HashPartitionsContext) {}

// ExitHashPartitions is called when production hashPartitions is exited.
func (s *BaseOracleStatementListener) ExitHashPartitions(ctx *HashPartitionsContext) {}

// EnterHashPartitionsByQuantity is called when production hashPartitionsByQuantity is entered.
func (s *BaseOracleStatementListener) EnterHashPartitionsByQuantity(ctx *HashPartitionsByQuantityContext) {
}

// ExitHashPartitionsByQuantity is called when production hashPartitionsByQuantity is exited.
func (s *BaseOracleStatementListener) ExitHashPartitionsByQuantity(ctx *HashPartitionsByQuantityContext) {
}

// EnterIndexCompression is called when production indexCompression is entered.
func (s *BaseOracleStatementListener) EnterIndexCompression(ctx *IndexCompressionContext) {}

// ExitIndexCompression is called when production indexCompression is exited.
func (s *BaseOracleStatementListener) ExitIndexCompression(ctx *IndexCompressionContext) {}

// EnterAdvancedIndexCompression is called when production advancedIndexCompression is entered.
func (s *BaseOracleStatementListener) EnterAdvancedIndexCompression(ctx *AdvancedIndexCompressionContext) {
}

// ExitAdvancedIndexCompression is called when production advancedIndexCompression is exited.
func (s *BaseOracleStatementListener) ExitAdvancedIndexCompression(ctx *AdvancedIndexCompressionContext) {
}

// EnterIndividualHashPartitions is called when production individualHashPartitions is entered.
func (s *BaseOracleStatementListener) EnterIndividualHashPartitions(ctx *IndividualHashPartitionsContext) {
}

// ExitIndividualHashPartitions is called when production individualHashPartitions is exited.
func (s *BaseOracleStatementListener) ExitIndividualHashPartitions(ctx *IndividualHashPartitionsContext) {
}

// EnterPartitioningStorageClause is called when production partitioningStorageClause is entered.
func (s *BaseOracleStatementListener) EnterPartitioningStorageClause(ctx *PartitioningStorageClauseContext) {
}

// ExitPartitioningStorageClause is called when production partitioningStorageClause is exited.
func (s *BaseOracleStatementListener) ExitPartitioningStorageClause(ctx *PartitioningStorageClauseContext) {
}

// EnterLobPartitioningStorage is called when production lobPartitioningStorage is entered.
func (s *BaseOracleStatementListener) EnterLobPartitioningStorage(ctx *LobPartitioningStorageContext) {
}

// ExitLobPartitioningStorage is called when production lobPartitioningStorage is exited.
func (s *BaseOracleStatementListener) ExitLobPartitioningStorage(ctx *LobPartitioningStorageContext) {
}

// EnterCompositeRangePartitions is called when production compositeRangePartitions is entered.
func (s *BaseOracleStatementListener) EnterCompositeRangePartitions(ctx *CompositeRangePartitionsContext) {
}

// ExitCompositeRangePartitions is called when production compositeRangePartitions is exited.
func (s *BaseOracleStatementListener) ExitCompositeRangePartitions(ctx *CompositeRangePartitionsContext) {
}

// EnterSubpartitionByRange is called when production subpartitionByRange is entered.
func (s *BaseOracleStatementListener) EnterSubpartitionByRange(ctx *SubpartitionByRangeContext) {}

// ExitSubpartitionByRange is called when production subpartitionByRange is exited.
func (s *BaseOracleStatementListener) ExitSubpartitionByRange(ctx *SubpartitionByRangeContext) {}

// EnterSubpartitionByList is called when production subpartitionByList is entered.
func (s *BaseOracleStatementListener) EnterSubpartitionByList(ctx *SubpartitionByListContext) {}

// ExitSubpartitionByList is called when production subpartitionByList is exited.
func (s *BaseOracleStatementListener) ExitSubpartitionByList(ctx *SubpartitionByListContext) {}

// EnterSubpartitionByHash is called when production subpartitionByHash is entered.
func (s *BaseOracleStatementListener) EnterSubpartitionByHash(ctx *SubpartitionByHashContext) {}

// ExitSubpartitionByHash is called when production subpartitionByHash is exited.
func (s *BaseOracleStatementListener) ExitSubpartitionByHash(ctx *SubpartitionByHashContext) {}

// EnterSubpartitionTemplate is called when production subpartitionTemplate is entered.
func (s *BaseOracleStatementListener) EnterSubpartitionTemplate(ctx *SubpartitionTemplateContext) {}

// ExitSubpartitionTemplate is called when production subpartitionTemplate is exited.
func (s *BaseOracleStatementListener) ExitSubpartitionTemplate(ctx *SubpartitionTemplateContext) {}

// EnterRangeSubpartitionDesc is called when production rangeSubpartitionDesc is entered.
func (s *BaseOracleStatementListener) EnterRangeSubpartitionDesc(ctx *RangeSubpartitionDescContext) {}

// ExitRangeSubpartitionDesc is called when production rangeSubpartitionDesc is exited.
func (s *BaseOracleStatementListener) ExitRangeSubpartitionDesc(ctx *RangeSubpartitionDescContext) {}

// EnterListSubpartitionDesc is called when production listSubpartitionDesc is entered.
func (s *BaseOracleStatementListener) EnterListSubpartitionDesc(ctx *ListSubpartitionDescContext) {}

// ExitListSubpartitionDesc is called when production listSubpartitionDesc is exited.
func (s *BaseOracleStatementListener) ExitListSubpartitionDesc(ctx *ListSubpartitionDescContext) {}

// EnterIndividualHashSubparts is called when production individualHashSubparts is entered.
func (s *BaseOracleStatementListener) EnterIndividualHashSubparts(ctx *IndividualHashSubpartsContext) {
}

// ExitIndividualHashSubparts is called when production individualHashSubparts is exited.
func (s *BaseOracleStatementListener) ExitIndividualHashSubparts(ctx *IndividualHashSubpartsContext) {
}

// EnterRangePartitionDesc is called when production rangePartitionDesc is entered.
func (s *BaseOracleStatementListener) EnterRangePartitionDesc(ctx *RangePartitionDescContext) {}

// ExitRangePartitionDesc is called when production rangePartitionDesc is exited.
func (s *BaseOracleStatementListener) ExitRangePartitionDesc(ctx *RangePartitionDescContext) {}

// EnterCompositeListPartitions is called when production compositeListPartitions is entered.
func (s *BaseOracleStatementListener) EnterCompositeListPartitions(ctx *CompositeListPartitionsContext) {
}

// ExitCompositeListPartitions is called when production compositeListPartitions is exited.
func (s *BaseOracleStatementListener) ExitCompositeListPartitions(ctx *CompositeListPartitionsContext) {
}

// EnterListPartitionDesc is called when production listPartitionDesc is entered.
func (s *BaseOracleStatementListener) EnterListPartitionDesc(ctx *ListPartitionDescContext) {}

// ExitListPartitionDesc is called when production listPartitionDesc is exited.
func (s *BaseOracleStatementListener) ExitListPartitionDesc(ctx *ListPartitionDescContext) {}

// EnterCompositeHashPartitions is called when production compositeHashPartitions is entered.
func (s *BaseOracleStatementListener) EnterCompositeHashPartitions(ctx *CompositeHashPartitionsContext) {
}

// ExitCompositeHashPartitions is called when production compositeHashPartitions is exited.
func (s *BaseOracleStatementListener) ExitCompositeHashPartitions(ctx *CompositeHashPartitionsContext) {
}

// EnterReferencePartitioning is called when production referencePartitioning is entered.
func (s *BaseOracleStatementListener) EnterReferencePartitioning(ctx *ReferencePartitioningContext) {}

// ExitReferencePartitioning is called when production referencePartitioning is exited.
func (s *BaseOracleStatementListener) ExitReferencePartitioning(ctx *ReferencePartitioningContext) {}

// EnterReferencePartitionDesc is called when production referencePartitionDesc is entered.
func (s *BaseOracleStatementListener) EnterReferencePartitionDesc(ctx *ReferencePartitionDescContext) {
}

// ExitReferencePartitionDesc is called when production referencePartitionDesc is exited.
func (s *BaseOracleStatementListener) ExitReferencePartitionDesc(ctx *ReferencePartitionDescContext) {
}

// EnterConstraint is called when production constraint is entered.
func (s *BaseOracleStatementListener) EnterConstraint(ctx *ConstraintContext) {}

// ExitConstraint is called when production constraint is exited.
func (s *BaseOracleStatementListener) ExitConstraint(ctx *ConstraintContext) {}

// EnterSystemPartitioning is called when production systemPartitioning is entered.
func (s *BaseOracleStatementListener) EnterSystemPartitioning(ctx *SystemPartitioningContext) {}

// ExitSystemPartitioning is called when production systemPartitioning is exited.
func (s *BaseOracleStatementListener) ExitSystemPartitioning(ctx *SystemPartitioningContext) {}

// EnterConsistentHashPartitions is called when production consistentHashPartitions is entered.
func (s *BaseOracleStatementListener) EnterConsistentHashPartitions(ctx *ConsistentHashPartitionsContext) {
}

// ExitConsistentHashPartitions is called when production consistentHashPartitions is exited.
func (s *BaseOracleStatementListener) ExitConsistentHashPartitions(ctx *ConsistentHashPartitionsContext) {
}

// EnterConsistentHashWithSubpartitions is called when production consistentHashWithSubpartitions is entered.
func (s *BaseOracleStatementListener) EnterConsistentHashWithSubpartitions(ctx *ConsistentHashWithSubpartitionsContext) {
}

// ExitConsistentHashWithSubpartitions is called when production consistentHashWithSubpartitions is exited.
func (s *BaseOracleStatementListener) ExitConsistentHashWithSubpartitions(ctx *ConsistentHashWithSubpartitionsContext) {
}

// EnterPartitionsetClauses is called when production partitionsetClauses is entered.
func (s *BaseOracleStatementListener) EnterPartitionsetClauses(ctx *PartitionsetClausesContext) {}

// ExitPartitionsetClauses is called when production partitionsetClauses is exited.
func (s *BaseOracleStatementListener) ExitPartitionsetClauses(ctx *PartitionsetClausesContext) {}

// EnterRangePartitionsetClause is called when production rangePartitionsetClause is entered.
func (s *BaseOracleStatementListener) EnterRangePartitionsetClause(ctx *RangePartitionsetClauseContext) {
}

// ExitRangePartitionsetClause is called when production rangePartitionsetClause is exited.
func (s *BaseOracleStatementListener) ExitRangePartitionsetClause(ctx *RangePartitionsetClauseContext) {
}

// EnterRangePartitionsetDesc is called when production rangePartitionsetDesc is entered.
func (s *BaseOracleStatementListener) EnterRangePartitionsetDesc(ctx *RangePartitionsetDescContext) {}

// ExitRangePartitionsetDesc is called when production rangePartitionsetDesc is exited.
func (s *BaseOracleStatementListener) ExitRangePartitionsetDesc(ctx *RangePartitionsetDescContext) {}

// EnterListPartitionsetClause is called when production listPartitionsetClause is entered.
func (s *BaseOracleStatementListener) EnterListPartitionsetClause(ctx *ListPartitionsetClauseContext) {
}

// ExitListPartitionsetClause is called when production listPartitionsetClause is exited.
func (s *BaseOracleStatementListener) ExitListPartitionsetClause(ctx *ListPartitionsetClauseContext) {
}

// EnterAttributeClusteringClause is called when production attributeClusteringClause is entered.
func (s *BaseOracleStatementListener) EnterAttributeClusteringClause(ctx *AttributeClusteringClauseContext) {
}

// ExitAttributeClusteringClause is called when production attributeClusteringClause is exited.
func (s *BaseOracleStatementListener) ExitAttributeClusteringClause(ctx *AttributeClusteringClauseContext) {
}

// EnterClusteringJoin is called when production clusteringJoin is entered.
func (s *BaseOracleStatementListener) EnterClusteringJoin(ctx *ClusteringJoinContext) {}

// ExitClusteringJoin is called when production clusteringJoin is exited.
func (s *BaseOracleStatementListener) ExitClusteringJoin(ctx *ClusteringJoinContext) {}

// EnterClusterClause is called when production clusterClause is entered.
func (s *BaseOracleStatementListener) EnterClusterClause(ctx *ClusterClauseContext) {}

// ExitClusterClause is called when production clusterClause is exited.
func (s *BaseOracleStatementListener) ExitClusterClause(ctx *ClusterClauseContext) {}

// EnterClusteringColumns is called when production clusteringColumns is entered.
func (s *BaseOracleStatementListener) EnterClusteringColumns(ctx *ClusteringColumnsContext) {}

// ExitClusteringColumns is called when production clusteringColumns is exited.
func (s *BaseOracleStatementListener) ExitClusteringColumns(ctx *ClusteringColumnsContext) {}

// EnterClusteringColumnGroup is called when production clusteringColumnGroup is entered.
func (s *BaseOracleStatementListener) EnterClusteringColumnGroup(ctx *ClusteringColumnGroupContext) {}

// ExitClusteringColumnGroup is called when production clusteringColumnGroup is exited.
func (s *BaseOracleStatementListener) ExitClusteringColumnGroup(ctx *ClusteringColumnGroupContext) {}

// EnterClusteringWhen is called when production clusteringWhen is entered.
func (s *BaseOracleStatementListener) EnterClusteringWhen(ctx *ClusteringWhenContext) {}

// ExitClusteringWhen is called when production clusteringWhen is exited.
func (s *BaseOracleStatementListener) ExitClusteringWhen(ctx *ClusteringWhenContext) {}

// EnterZonemapClause is called when production zonemapClause is entered.
func (s *BaseOracleStatementListener) EnterZonemapClause(ctx *ZonemapClauseContext) {}

// ExitZonemapClause is called when production zonemapClause is exited.
func (s *BaseOracleStatementListener) ExitZonemapClause(ctx *ZonemapClauseContext) {}

// EnterRowMovementClause is called when production rowMovementClause is entered.
func (s *BaseOracleStatementListener) EnterRowMovementClause(ctx *RowMovementClauseContext) {}

// ExitRowMovementClause is called when production rowMovementClause is exited.
func (s *BaseOracleStatementListener) ExitRowMovementClause(ctx *RowMovementClauseContext) {}

// EnterFlashbackArchiveClause is called when production flashbackArchiveClause is entered.
func (s *BaseOracleStatementListener) EnterFlashbackArchiveClause(ctx *FlashbackArchiveClauseContext) {
}

// ExitFlashbackArchiveClause is called when production flashbackArchiveClause is exited.
func (s *BaseOracleStatementListener) ExitFlashbackArchiveClause(ctx *FlashbackArchiveClauseContext) {
}

// EnterAlterSynonym is called when production alterSynonym is entered.
func (s *BaseOracleStatementListener) EnterAlterSynonym(ctx *AlterSynonymContext) {}

// ExitAlterSynonym is called when production alterSynonym is exited.
func (s *BaseOracleStatementListener) ExitAlterSynonym(ctx *AlterSynonymContext) {}

// EnterAlterTablePartitioning is called when production alterTablePartitioning is entered.
func (s *BaseOracleStatementListener) EnterAlterTablePartitioning(ctx *AlterTablePartitioningContext) {
}

// ExitAlterTablePartitioning is called when production alterTablePartitioning is exited.
func (s *BaseOracleStatementListener) ExitAlterTablePartitioning(ctx *AlterTablePartitioningContext) {
}

// EnterModifyTablePartition is called when production modifyTablePartition is entered.
func (s *BaseOracleStatementListener) EnterModifyTablePartition(ctx *ModifyTablePartitionContext) {}

// ExitModifyTablePartition is called when production modifyTablePartition is exited.
func (s *BaseOracleStatementListener) ExitModifyTablePartition(ctx *ModifyTablePartitionContext) {}

// EnterModifyRangePartition is called when production modifyRangePartition is entered.
func (s *BaseOracleStatementListener) EnterModifyRangePartition(ctx *ModifyRangePartitionContext) {}

// ExitModifyRangePartition is called when production modifyRangePartition is exited.
func (s *BaseOracleStatementListener) ExitModifyRangePartition(ctx *ModifyRangePartitionContext) {}

// EnterModifyHashPartition is called when production modifyHashPartition is entered.
func (s *BaseOracleStatementListener) EnterModifyHashPartition(ctx *ModifyHashPartitionContext) {}

// ExitModifyHashPartition is called when production modifyHashPartition is exited.
func (s *BaseOracleStatementListener) ExitModifyHashPartition(ctx *ModifyHashPartitionContext) {}

// EnterModifyListPartition is called when production modifyListPartition is entered.
func (s *BaseOracleStatementListener) EnterModifyListPartition(ctx *ModifyListPartitionContext) {}

// ExitModifyListPartition is called when production modifyListPartition is exited.
func (s *BaseOracleStatementListener) ExitModifyListPartition(ctx *ModifyListPartitionContext) {}

// EnterPartitionExtendedName is called when production partitionExtendedName is entered.
func (s *BaseOracleStatementListener) EnterPartitionExtendedName(ctx *PartitionExtendedNameContext) {}

// ExitPartitionExtendedName is called when production partitionExtendedName is exited.
func (s *BaseOracleStatementListener) ExitPartitionExtendedName(ctx *PartitionExtendedNameContext) {}

// EnterAddRangeSubpartition is called when production addRangeSubpartition is entered.
func (s *BaseOracleStatementListener) EnterAddRangeSubpartition(ctx *AddRangeSubpartitionContext) {}

// ExitAddRangeSubpartition is called when production addRangeSubpartition is exited.
func (s *BaseOracleStatementListener) ExitAddRangeSubpartition(ctx *AddRangeSubpartitionContext) {}

// EnterDependentTablesClause is called when production dependentTablesClause is entered.
func (s *BaseOracleStatementListener) EnterDependentTablesClause(ctx *DependentTablesClauseContext) {}

// ExitDependentTablesClause is called when production dependentTablesClause is exited.
func (s *BaseOracleStatementListener) ExitDependentTablesClause(ctx *DependentTablesClauseContext) {}

// EnterAddHashSubpartition is called when production addHashSubpartition is entered.
func (s *BaseOracleStatementListener) EnterAddHashSubpartition(ctx *AddHashSubpartitionContext) {}

// ExitAddHashSubpartition is called when production addHashSubpartition is exited.
func (s *BaseOracleStatementListener) ExitAddHashSubpartition(ctx *AddHashSubpartitionContext) {}

// EnterAddListSubpartition is called when production addListSubpartition is entered.
func (s *BaseOracleStatementListener) EnterAddListSubpartition(ctx *AddListSubpartitionContext) {}

// ExitAddListSubpartition is called when production addListSubpartition is exited.
func (s *BaseOracleStatementListener) ExitAddListSubpartition(ctx *AddListSubpartitionContext) {}

// EnterCoalesceTableSubpartition is called when production coalesceTableSubpartition is entered.
func (s *BaseOracleStatementListener) EnterCoalesceTableSubpartition(ctx *CoalesceTableSubpartitionContext) {
}

// ExitCoalesceTableSubpartition is called when production coalesceTableSubpartition is exited.
func (s *BaseOracleStatementListener) ExitCoalesceTableSubpartition(ctx *CoalesceTableSubpartitionContext) {
}

// EnterAllowDisallowClustering is called when production allowDisallowClustering is entered.
func (s *BaseOracleStatementListener) EnterAllowDisallowClustering(ctx *AllowDisallowClusteringContext) {
}

// ExitAllowDisallowClustering is called when production allowDisallowClustering is exited.
func (s *BaseOracleStatementListener) ExitAllowDisallowClustering(ctx *AllowDisallowClusteringContext) {
}

// EnterAlterMappingTableClauses is called when production alterMappingTableClauses is entered.
func (s *BaseOracleStatementListener) EnterAlterMappingTableClauses(ctx *AlterMappingTableClausesContext) {
}

// ExitAlterMappingTableClauses is called when production alterMappingTableClauses is exited.
func (s *BaseOracleStatementListener) ExitAlterMappingTableClauses(ctx *AlterMappingTableClausesContext) {
}

// EnterDeallocateUnusedClause is called when production deallocateUnusedClause is entered.
func (s *BaseOracleStatementListener) EnterDeallocateUnusedClause(ctx *DeallocateUnusedClauseContext) {
}

// ExitDeallocateUnusedClause is called when production deallocateUnusedClause is exited.
func (s *BaseOracleStatementListener) ExitDeallocateUnusedClause(ctx *DeallocateUnusedClauseContext) {
}

// EnterAllocateExtentClause is called when production allocateExtentClause is entered.
func (s *BaseOracleStatementListener) EnterAllocateExtentClause(ctx *AllocateExtentClauseContext) {}

// ExitAllocateExtentClause is called when production allocateExtentClause is exited.
func (s *BaseOracleStatementListener) ExitAllocateExtentClause(ctx *AllocateExtentClauseContext) {}

// EnterPartitionSpec is called when production partitionSpec is entered.
func (s *BaseOracleStatementListener) EnterPartitionSpec(ctx *PartitionSpecContext) {}

// ExitPartitionSpec is called when production partitionSpec is exited.
func (s *BaseOracleStatementListener) ExitPartitionSpec(ctx *PartitionSpecContext) {}

// EnterPartitionAttributes is called when production partitionAttributes is entered.
func (s *BaseOracleStatementListener) EnterPartitionAttributes(ctx *PartitionAttributesContext) {}

// ExitPartitionAttributes is called when production partitionAttributes is exited.
func (s *BaseOracleStatementListener) ExitPartitionAttributes(ctx *PartitionAttributesContext) {}

// EnterShrinkClause is called when production shrinkClause is entered.
func (s *BaseOracleStatementListener) EnterShrinkClause(ctx *ShrinkClauseContext) {}

// ExitShrinkClause is called when production shrinkClause is exited.
func (s *BaseOracleStatementListener) ExitShrinkClause(ctx *ShrinkClauseContext) {}

// EnterMoveTablePartition is called when production moveTablePartition is entered.
func (s *BaseOracleStatementListener) EnterMoveTablePartition(ctx *MoveTablePartitionContext) {}

// ExitMoveTablePartition is called when production moveTablePartition is exited.
func (s *BaseOracleStatementListener) ExitMoveTablePartition(ctx *MoveTablePartitionContext) {}

// EnterFilterCondition is called when production filterCondition is entered.
func (s *BaseOracleStatementListener) EnterFilterCondition(ctx *FilterConditionContext) {}

// ExitFilterCondition is called when production filterCondition is exited.
func (s *BaseOracleStatementListener) ExitFilterCondition(ctx *FilterConditionContext) {}

// EnterCoalesceTablePartition is called when production coalesceTablePartition is entered.
func (s *BaseOracleStatementListener) EnterCoalesceTablePartition(ctx *CoalesceTablePartitionContext) {
}

// ExitCoalesceTablePartition is called when production coalesceTablePartition is exited.
func (s *BaseOracleStatementListener) ExitCoalesceTablePartition(ctx *CoalesceTablePartitionContext) {
}

// EnterAddTablePartition is called when production addTablePartition is entered.
func (s *BaseOracleStatementListener) EnterAddTablePartition(ctx *AddTablePartitionContext) {}

// ExitAddTablePartition is called when production addTablePartition is exited.
func (s *BaseOracleStatementListener) ExitAddTablePartition(ctx *AddTablePartitionContext) {}

// EnterAddRangePartitionClause is called when production addRangePartitionClause is entered.
func (s *BaseOracleStatementListener) EnterAddRangePartitionClause(ctx *AddRangePartitionClauseContext) {
}

// ExitAddRangePartitionClause is called when production addRangePartitionClause is exited.
func (s *BaseOracleStatementListener) ExitAddRangePartitionClause(ctx *AddRangePartitionClauseContext) {
}

// EnterAddListPartitionClause is called when production addListPartitionClause is entered.
func (s *BaseOracleStatementListener) EnterAddListPartitionClause(ctx *AddListPartitionClauseContext) {
}

// ExitAddListPartitionClause is called when production addListPartitionClause is exited.
func (s *BaseOracleStatementListener) ExitAddListPartitionClause(ctx *AddListPartitionClauseContext) {
}

// EnterHashSubpartsByQuantity is called when production hashSubpartsByQuantity is entered.
func (s *BaseOracleStatementListener) EnterHashSubpartsByQuantity(ctx *HashSubpartsByQuantityContext) {
}

// ExitHashSubpartsByQuantity is called when production hashSubpartsByQuantity is exited.
func (s *BaseOracleStatementListener) ExitHashSubpartsByQuantity(ctx *HashSubpartsByQuantityContext) {
}

// EnterAddSystemPartitionClause is called when production addSystemPartitionClause is entered.
func (s *BaseOracleStatementListener) EnterAddSystemPartitionClause(ctx *AddSystemPartitionClauseContext) {
}

// ExitAddSystemPartitionClause is called when production addSystemPartitionClause is exited.
func (s *BaseOracleStatementListener) ExitAddSystemPartitionClause(ctx *AddSystemPartitionClauseContext) {
}

// EnterAddHashPartitionClause is called when production addHashPartitionClause is entered.
func (s *BaseOracleStatementListener) EnterAddHashPartitionClause(ctx *AddHashPartitionClauseContext) {
}

// ExitAddHashPartitionClause is called when production addHashPartitionClause is exited.
func (s *BaseOracleStatementListener) ExitAddHashPartitionClause(ctx *AddHashPartitionClauseContext) {
}

// EnterDropTablePartition is called when production dropTablePartition is entered.
func (s *BaseOracleStatementListener) EnterDropTablePartition(ctx *DropTablePartitionContext) {}

// ExitDropTablePartition is called when production dropTablePartition is exited.
func (s *BaseOracleStatementListener) ExitDropTablePartition(ctx *DropTablePartitionContext) {}

// EnterPartitionExtendedNames is called when production partitionExtendedNames is entered.
func (s *BaseOracleStatementListener) EnterPartitionExtendedNames(ctx *PartitionExtendedNamesContext) {
}

// ExitPartitionExtendedNames is called when production partitionExtendedNames is exited.
func (s *BaseOracleStatementListener) ExitPartitionExtendedNames(ctx *PartitionExtendedNamesContext) {
}

// EnterPartitionForClauses is called when production partitionForClauses is entered.
func (s *BaseOracleStatementListener) EnterPartitionForClauses(ctx *PartitionForClausesContext) {}

// ExitPartitionForClauses is called when production partitionForClauses is exited.
func (s *BaseOracleStatementListener) ExitPartitionForClauses(ctx *PartitionForClausesContext) {}

// EnterUpdateIndexClauses is called when production updateIndexClauses is entered.
func (s *BaseOracleStatementListener) EnterUpdateIndexClauses(ctx *UpdateIndexClausesContext) {}

// ExitUpdateIndexClauses is called when production updateIndexClauses is exited.
func (s *BaseOracleStatementListener) ExitUpdateIndexClauses(ctx *UpdateIndexClausesContext) {}

// EnterUpdateGlobalIndexClause is called when production updateGlobalIndexClause is entered.
func (s *BaseOracleStatementListener) EnterUpdateGlobalIndexClause(ctx *UpdateGlobalIndexClauseContext) {
}

// ExitUpdateGlobalIndexClause is called when production updateGlobalIndexClause is exited.
func (s *BaseOracleStatementListener) ExitUpdateGlobalIndexClause(ctx *UpdateGlobalIndexClauseContext) {
}

// EnterUpdateAllIndexesClause is called when production updateAllIndexesClause is entered.
func (s *BaseOracleStatementListener) EnterUpdateAllIndexesClause(ctx *UpdateAllIndexesClauseContext) {
}

// ExitUpdateAllIndexesClause is called when production updateAllIndexesClause is exited.
func (s *BaseOracleStatementListener) ExitUpdateAllIndexesClause(ctx *UpdateAllIndexesClauseContext) {
}

// EnterUpdateIndexPartition is called when production updateIndexPartition is entered.
func (s *BaseOracleStatementListener) EnterUpdateIndexPartition(ctx *UpdateIndexPartitionContext) {}

// ExitUpdateIndexPartition is called when production updateIndexPartition is exited.
func (s *BaseOracleStatementListener) ExitUpdateIndexPartition(ctx *UpdateIndexPartitionContext) {}

// EnterIndexPartitionDesc is called when production indexPartitionDesc is entered.
func (s *BaseOracleStatementListener) EnterIndexPartitionDesc(ctx *IndexPartitionDescContext) {}

// ExitIndexPartitionDesc is called when production indexPartitionDesc is exited.
func (s *BaseOracleStatementListener) ExitIndexPartitionDesc(ctx *IndexPartitionDescContext) {}

// EnterIndexSubpartitionClause is called when production indexSubpartitionClause is entered.
func (s *BaseOracleStatementListener) EnterIndexSubpartitionClause(ctx *IndexSubpartitionClauseContext) {
}

// ExitIndexSubpartitionClause is called when production indexSubpartitionClause is exited.
func (s *BaseOracleStatementListener) ExitIndexSubpartitionClause(ctx *IndexSubpartitionClauseContext) {
}

// EnterUpdateIndexSubpartition is called when production updateIndexSubpartition is entered.
func (s *BaseOracleStatementListener) EnterUpdateIndexSubpartition(ctx *UpdateIndexSubpartitionContext) {
}

// ExitUpdateIndexSubpartition is called when production updateIndexSubpartition is exited.
func (s *BaseOracleStatementListener) ExitUpdateIndexSubpartition(ctx *UpdateIndexSubpartitionContext) {
}

// EnterSupplementalLoggingProps is called when production supplementalLoggingProps is entered.
func (s *BaseOracleStatementListener) EnterSupplementalLoggingProps(ctx *SupplementalLoggingPropsContext) {
}

// ExitSupplementalLoggingProps is called when production supplementalLoggingProps is exited.
func (s *BaseOracleStatementListener) ExitSupplementalLoggingProps(ctx *SupplementalLoggingPropsContext) {
}

// EnterSupplementalLogGrpClause is called when production supplementalLogGrpClause is entered.
func (s *BaseOracleStatementListener) EnterSupplementalLogGrpClause(ctx *SupplementalLogGrpClauseContext) {
}

// ExitSupplementalLogGrpClause is called when production supplementalLogGrpClause is exited.
func (s *BaseOracleStatementListener) ExitSupplementalLogGrpClause(ctx *SupplementalLogGrpClauseContext) {
}

// EnterSupplementalIdKeyClause is called when production supplementalIdKeyClause is entered.
func (s *BaseOracleStatementListener) EnterSupplementalIdKeyClause(ctx *SupplementalIdKeyClauseContext) {
}

// ExitSupplementalIdKeyClause is called when production supplementalIdKeyClause is exited.
func (s *BaseOracleStatementListener) ExitSupplementalIdKeyClause(ctx *SupplementalIdKeyClauseContext) {
}

// EnterAlterSession is called when production alterSession is entered.
func (s *BaseOracleStatementListener) EnterAlterSession(ctx *AlterSessionContext) {}

// ExitAlterSession is called when production alterSession is exited.
func (s *BaseOracleStatementListener) ExitAlterSession(ctx *AlterSessionContext) {}

// EnterAlterSessionOption is called when production alterSessionOption is entered.
func (s *BaseOracleStatementListener) EnterAlterSessionOption(ctx *AlterSessionOptionContext) {}

// ExitAlterSessionOption is called when production alterSessionOption is exited.
func (s *BaseOracleStatementListener) ExitAlterSessionOption(ctx *AlterSessionOptionContext) {}

// EnterAdviseClause is called when production adviseClause is entered.
func (s *BaseOracleStatementListener) EnterAdviseClause(ctx *AdviseClauseContext) {}

// ExitAdviseClause is called when production adviseClause is exited.
func (s *BaseOracleStatementListener) ExitAdviseClause(ctx *AdviseClauseContext) {}

// EnterCloseDatabaseLinkClause is called when production closeDatabaseLinkClause is entered.
func (s *BaseOracleStatementListener) EnterCloseDatabaseLinkClause(ctx *CloseDatabaseLinkClauseContext) {
}

// ExitCloseDatabaseLinkClause is called when production closeDatabaseLinkClause is exited.
func (s *BaseOracleStatementListener) ExitCloseDatabaseLinkClause(ctx *CloseDatabaseLinkClauseContext) {
}

// EnterCommitInProcedureClause is called when production commitInProcedureClause is entered.
func (s *BaseOracleStatementListener) EnterCommitInProcedureClause(ctx *CommitInProcedureClauseContext) {
}

// ExitCommitInProcedureClause is called when production commitInProcedureClause is exited.
func (s *BaseOracleStatementListener) ExitCommitInProcedureClause(ctx *CommitInProcedureClauseContext) {
}

// EnterSecuriyClause is called when production securiyClause is entered.
func (s *BaseOracleStatementListener) EnterSecuriyClause(ctx *SecuriyClauseContext) {}

// ExitSecuriyClause is called when production securiyClause is exited.
func (s *BaseOracleStatementListener) ExitSecuriyClause(ctx *SecuriyClauseContext) {}

// EnterParallelExecutionClause is called when production parallelExecutionClause is entered.
func (s *BaseOracleStatementListener) EnterParallelExecutionClause(ctx *ParallelExecutionClauseContext) {
}

// ExitParallelExecutionClause is called when production parallelExecutionClause is exited.
func (s *BaseOracleStatementListener) ExitParallelExecutionClause(ctx *ParallelExecutionClauseContext) {
}

// EnterResumableClause is called when production resumableClause is entered.
func (s *BaseOracleStatementListener) EnterResumableClause(ctx *ResumableClauseContext) {}

// ExitResumableClause is called when production resumableClause is exited.
func (s *BaseOracleStatementListener) ExitResumableClause(ctx *ResumableClauseContext) {}

// EnterEnableResumableClause is called when production enableResumableClause is entered.
func (s *BaseOracleStatementListener) EnterEnableResumableClause(ctx *EnableResumableClauseContext) {}

// ExitEnableResumableClause is called when production enableResumableClause is exited.
func (s *BaseOracleStatementListener) ExitEnableResumableClause(ctx *EnableResumableClauseContext) {}

// EnterDisableResumableClause is called when production disableResumableClause is entered.
func (s *BaseOracleStatementListener) EnterDisableResumableClause(ctx *DisableResumableClauseContext) {
}

// ExitDisableResumableClause is called when production disableResumableClause is exited.
func (s *BaseOracleStatementListener) ExitDisableResumableClause(ctx *DisableResumableClauseContext) {
}

// EnterShardDdlClause is called when production shardDdlClause is entered.
func (s *BaseOracleStatementListener) EnterShardDdlClause(ctx *ShardDdlClauseContext) {}

// ExitShardDdlClause is called when production shardDdlClause is exited.
func (s *BaseOracleStatementListener) ExitShardDdlClause(ctx *ShardDdlClauseContext) {}

// EnterSyncWithPrimaryClause is called when production syncWithPrimaryClause is entered.
func (s *BaseOracleStatementListener) EnterSyncWithPrimaryClause(ctx *SyncWithPrimaryClauseContext) {}

// ExitSyncWithPrimaryClause is called when production syncWithPrimaryClause is exited.
func (s *BaseOracleStatementListener) ExitSyncWithPrimaryClause(ctx *SyncWithPrimaryClauseContext) {}

// EnterAlterSessionSetClause is called when production alterSessionSetClause is entered.
func (s *BaseOracleStatementListener) EnterAlterSessionSetClause(ctx *AlterSessionSetClauseContext) {}

// ExitAlterSessionSetClause is called when production alterSessionSetClause is exited.
func (s *BaseOracleStatementListener) ExitAlterSessionSetClause(ctx *AlterSessionSetClauseContext) {}

// EnterAlterSessionSetClauseOption is called when production alterSessionSetClauseOption is entered.
func (s *BaseOracleStatementListener) EnterAlterSessionSetClauseOption(ctx *AlterSessionSetClauseOptionContext) {
}

// ExitAlterSessionSetClauseOption is called when production alterSessionSetClauseOption is exited.
func (s *BaseOracleStatementListener) ExitAlterSessionSetClauseOption(ctx *AlterSessionSetClauseOptionContext) {
}

// EnterParameterClause is called when production parameterClause is entered.
func (s *BaseOracleStatementListener) EnterParameterClause(ctx *ParameterClauseContext) {}

// ExitParameterClause is called when production parameterClause is exited.
func (s *BaseOracleStatementListener) ExitParameterClause(ctx *ParameterClauseContext) {}

// EnterEditionClause is called when production editionClause is entered.
func (s *BaseOracleStatementListener) EnterEditionClause(ctx *EditionClauseContext) {}

// ExitEditionClause is called when production editionClause is exited.
func (s *BaseOracleStatementListener) ExitEditionClause(ctx *EditionClauseContext) {}

// EnterContainerClause is called when production containerClause is entered.
func (s *BaseOracleStatementListener) EnterContainerClause(ctx *ContainerClauseContext) {}

// ExitContainerClause is called when production containerClause is exited.
func (s *BaseOracleStatementListener) ExitContainerClause(ctx *ContainerClauseContext) {}

// EnterRowArchivalVisibilityClause is called when production rowArchivalVisibilityClause is entered.
func (s *BaseOracleStatementListener) EnterRowArchivalVisibilityClause(ctx *RowArchivalVisibilityClauseContext) {
}

// ExitRowArchivalVisibilityClause is called when production rowArchivalVisibilityClause is exited.
func (s *BaseOracleStatementListener) ExitRowArchivalVisibilityClause(ctx *RowArchivalVisibilityClauseContext) {
}

// EnterAlterDatabase is called when production alterDatabase is entered.
func (s *BaseOracleStatementListener) EnterAlterDatabase(ctx *AlterDatabaseContext) {}

// ExitAlterDatabase is called when production alterDatabase is exited.
func (s *BaseOracleStatementListener) ExitAlterDatabase(ctx *AlterDatabaseContext) {}

// EnterDatabaseClauses is called when production databaseClauses is entered.
func (s *BaseOracleStatementListener) EnterDatabaseClauses(ctx *DatabaseClausesContext) {}

// ExitDatabaseClauses is called when production databaseClauses is exited.
func (s *BaseOracleStatementListener) ExitDatabaseClauses(ctx *DatabaseClausesContext) {}

// EnterStartupClauses is called when production startupClauses is entered.
func (s *BaseOracleStatementListener) EnterStartupClauses(ctx *StartupClausesContext) {}

// ExitStartupClauses is called when production startupClauses is exited.
func (s *BaseOracleStatementListener) ExitStartupClauses(ctx *StartupClausesContext) {}

// EnterRecoveryClauses is called when production recoveryClauses is entered.
func (s *BaseOracleStatementListener) EnterRecoveryClauses(ctx *RecoveryClausesContext) {}

// ExitRecoveryClauses is called when production recoveryClauses is exited.
func (s *BaseOracleStatementListener) ExitRecoveryClauses(ctx *RecoveryClausesContext) {}

// EnterGeneralRecovery is called when production generalRecovery is entered.
func (s *BaseOracleStatementListener) EnterGeneralRecovery(ctx *GeneralRecoveryContext) {}

// ExitGeneralRecovery is called when production generalRecovery is exited.
func (s *BaseOracleStatementListener) ExitGeneralRecovery(ctx *GeneralRecoveryContext) {}

// EnterFullDatabaseRecovery is called when production fullDatabaseRecovery is entered.
func (s *BaseOracleStatementListener) EnterFullDatabaseRecovery(ctx *FullDatabaseRecoveryContext) {}

// ExitFullDatabaseRecovery is called when production fullDatabaseRecovery is exited.
func (s *BaseOracleStatementListener) ExitFullDatabaseRecovery(ctx *FullDatabaseRecoveryContext) {}

// EnterPartialDatabaseRecovery is called when production partialDatabaseRecovery is entered.
func (s *BaseOracleStatementListener) EnterPartialDatabaseRecovery(ctx *PartialDatabaseRecoveryContext) {
}

// ExitPartialDatabaseRecovery is called when production partialDatabaseRecovery is exited.
func (s *BaseOracleStatementListener) ExitPartialDatabaseRecovery(ctx *PartialDatabaseRecoveryContext) {
}

// EnterManagedStandbyRecovery is called when production managedStandbyRecovery is entered.
func (s *BaseOracleStatementListener) EnterManagedStandbyRecovery(ctx *ManagedStandbyRecoveryContext) {
}

// ExitManagedStandbyRecovery is called when production managedStandbyRecovery is exited.
func (s *BaseOracleStatementListener) ExitManagedStandbyRecovery(ctx *ManagedStandbyRecoveryContext) {
}

// EnterDatabaseFileClauses is called when production databaseFileClauses is entered.
func (s *BaseOracleStatementListener) EnterDatabaseFileClauses(ctx *DatabaseFileClausesContext) {}

// ExitDatabaseFileClauses is called when production databaseFileClauses is exited.
func (s *BaseOracleStatementListener) ExitDatabaseFileClauses(ctx *DatabaseFileClausesContext) {}

// EnterCreateDatafileClause is called when production createDatafileClause is entered.
func (s *BaseOracleStatementListener) EnterCreateDatafileClause(ctx *CreateDatafileClauseContext) {}

// ExitCreateDatafileClause is called when production createDatafileClause is exited.
func (s *BaseOracleStatementListener) ExitCreateDatafileClause(ctx *CreateDatafileClauseContext) {}

// EnterFileSpecifications is called when production fileSpecifications is entered.
func (s *BaseOracleStatementListener) EnterFileSpecifications(ctx *FileSpecificationsContext) {}

// ExitFileSpecifications is called when production fileSpecifications is exited.
func (s *BaseOracleStatementListener) ExitFileSpecifications(ctx *FileSpecificationsContext) {}

// EnterFileSpecification is called when production fileSpecification is entered.
func (s *BaseOracleStatementListener) EnterFileSpecification(ctx *FileSpecificationContext) {}

// ExitFileSpecification is called when production fileSpecification is exited.
func (s *BaseOracleStatementListener) ExitFileSpecification(ctx *FileSpecificationContext) {}

// EnterDatafileTempfileSpec is called when production datafileTempfileSpec is entered.
func (s *BaseOracleStatementListener) EnterDatafileTempfileSpec(ctx *DatafileTempfileSpecContext) {}

// ExitDatafileTempfileSpec is called when production datafileTempfileSpec is exited.
func (s *BaseOracleStatementListener) ExitDatafileTempfileSpec(ctx *DatafileTempfileSpecContext) {}

// EnterAutoextendClause is called when production autoextendClause is entered.
func (s *BaseOracleStatementListener) EnterAutoextendClause(ctx *AutoextendClauseContext) {}

// ExitAutoextendClause is called when production autoextendClause is exited.
func (s *BaseOracleStatementListener) ExitAutoextendClause(ctx *AutoextendClauseContext) {}

// EnterRedoLogFileSpec is called when production redoLogFileSpec is entered.
func (s *BaseOracleStatementListener) EnterRedoLogFileSpec(ctx *RedoLogFileSpecContext) {}

// ExitRedoLogFileSpec is called when production redoLogFileSpec is exited.
func (s *BaseOracleStatementListener) ExitRedoLogFileSpec(ctx *RedoLogFileSpecContext) {}

// EnterAlterDatafileClause is called when production alterDatafileClause is entered.
func (s *BaseOracleStatementListener) EnterAlterDatafileClause(ctx *AlterDatafileClauseContext) {}

// ExitAlterDatafileClause is called when production alterDatafileClause is exited.
func (s *BaseOracleStatementListener) ExitAlterDatafileClause(ctx *AlterDatafileClauseContext) {}

// EnterAlterTempfileClause is called when production alterTempfileClause is entered.
func (s *BaseOracleStatementListener) EnterAlterTempfileClause(ctx *AlterTempfileClauseContext) {}

// ExitAlterTempfileClause is called when production alterTempfileClause is exited.
func (s *BaseOracleStatementListener) ExitAlterTempfileClause(ctx *AlterTempfileClauseContext) {}

// EnterLogfileClauses is called when production logfileClauses is entered.
func (s *BaseOracleStatementListener) EnterLogfileClauses(ctx *LogfileClausesContext) {}

// ExitLogfileClauses is called when production logfileClauses is exited.
func (s *BaseOracleStatementListener) ExitLogfileClauses(ctx *LogfileClausesContext) {}

// EnterLogfileDescriptor is called when production logfileDescriptor is entered.
func (s *BaseOracleStatementListener) EnterLogfileDescriptor(ctx *LogfileDescriptorContext) {}

// ExitLogfileDescriptor is called when production logfileDescriptor is exited.
func (s *BaseOracleStatementListener) ExitLogfileDescriptor(ctx *LogfileDescriptorContext) {}

// EnterAddLogfileClauses is called when production addLogfileClauses is entered.
func (s *BaseOracleStatementListener) EnterAddLogfileClauses(ctx *AddLogfileClausesContext) {}

// ExitAddLogfileClauses is called when production addLogfileClauses is exited.
func (s *BaseOracleStatementListener) ExitAddLogfileClauses(ctx *AddLogfileClausesContext) {}

// EnterControlfileClauses is called when production controlfileClauses is entered.
func (s *BaseOracleStatementListener) EnterControlfileClauses(ctx *ControlfileClausesContext) {}

// ExitControlfileClauses is called when production controlfileClauses is exited.
func (s *BaseOracleStatementListener) ExitControlfileClauses(ctx *ControlfileClausesContext) {}

// EnterTraceFileClause is called when production traceFileClause is entered.
func (s *BaseOracleStatementListener) EnterTraceFileClause(ctx *TraceFileClauseContext) {}

// ExitTraceFileClause is called when production traceFileClause is exited.
func (s *BaseOracleStatementListener) ExitTraceFileClause(ctx *TraceFileClauseContext) {}

// EnterDropLogfileClauses is called when production dropLogfileClauses is entered.
func (s *BaseOracleStatementListener) EnterDropLogfileClauses(ctx *DropLogfileClausesContext) {}

// ExitDropLogfileClauses is called when production dropLogfileClauses is exited.
func (s *BaseOracleStatementListener) ExitDropLogfileClauses(ctx *DropLogfileClausesContext) {}

// EnterSwitchLogfileClause is called when production switchLogfileClause is entered.
func (s *BaseOracleStatementListener) EnterSwitchLogfileClause(ctx *SwitchLogfileClauseContext) {}

// ExitSwitchLogfileClause is called when production switchLogfileClause is exited.
func (s *BaseOracleStatementListener) ExitSwitchLogfileClause(ctx *SwitchLogfileClauseContext) {}

// EnterSupplementalDbLogging is called when production supplementalDbLogging is entered.
func (s *BaseOracleStatementListener) EnterSupplementalDbLogging(ctx *SupplementalDbLoggingContext) {}

// ExitSupplementalDbLogging is called when production supplementalDbLogging is exited.
func (s *BaseOracleStatementListener) ExitSupplementalDbLogging(ctx *SupplementalDbLoggingContext) {}

// EnterSupplementalPlsqlClause is called when production supplementalPlsqlClause is entered.
func (s *BaseOracleStatementListener) EnterSupplementalPlsqlClause(ctx *SupplementalPlsqlClauseContext) {
}

// ExitSupplementalPlsqlClause is called when production supplementalPlsqlClause is exited.
func (s *BaseOracleStatementListener) ExitSupplementalPlsqlClause(ctx *SupplementalPlsqlClauseContext) {
}

// EnterSupplementalSubsetReplicationClause is called when production supplementalSubsetReplicationClause is entered.
func (s *BaseOracleStatementListener) EnterSupplementalSubsetReplicationClause(ctx *SupplementalSubsetReplicationClauseContext) {
}

// ExitSupplementalSubsetReplicationClause is called when production supplementalSubsetReplicationClause is exited.
func (s *BaseOracleStatementListener) ExitSupplementalSubsetReplicationClause(ctx *SupplementalSubsetReplicationClauseContext) {
}

// EnterStandbyDatabaseClauses is called when production standbyDatabaseClauses is entered.
func (s *BaseOracleStatementListener) EnterStandbyDatabaseClauses(ctx *StandbyDatabaseClausesContext) {
}

// ExitStandbyDatabaseClauses is called when production standbyDatabaseClauses is exited.
func (s *BaseOracleStatementListener) ExitStandbyDatabaseClauses(ctx *StandbyDatabaseClausesContext) {
}

// EnterActivateStandbyDbClause is called when production activateStandbyDbClause is entered.
func (s *BaseOracleStatementListener) EnterActivateStandbyDbClause(ctx *ActivateStandbyDbClauseContext) {
}

// ExitActivateStandbyDbClause is called when production activateStandbyDbClause is exited.
func (s *BaseOracleStatementListener) ExitActivateStandbyDbClause(ctx *ActivateStandbyDbClauseContext) {
}

// EnterMaximizeStandbyDbClause is called when production maximizeStandbyDbClause is entered.
func (s *BaseOracleStatementListener) EnterMaximizeStandbyDbClause(ctx *MaximizeStandbyDbClauseContext) {
}

// ExitMaximizeStandbyDbClause is called when production maximizeStandbyDbClause is exited.
func (s *BaseOracleStatementListener) ExitMaximizeStandbyDbClause(ctx *MaximizeStandbyDbClauseContext) {
}

// EnterRegisterLogfileClause is called when production registerLogfileClause is entered.
func (s *BaseOracleStatementListener) EnterRegisterLogfileClause(ctx *RegisterLogfileClauseContext) {}

// ExitRegisterLogfileClause is called when production registerLogfileClause is exited.
func (s *BaseOracleStatementListener) ExitRegisterLogfileClause(ctx *RegisterLogfileClauseContext) {}

// EnterCommitSwitchoverClause is called when production commitSwitchoverClause is entered.
func (s *BaseOracleStatementListener) EnterCommitSwitchoverClause(ctx *CommitSwitchoverClauseContext) {
}

// ExitCommitSwitchoverClause is called when production commitSwitchoverClause is exited.
func (s *BaseOracleStatementListener) ExitCommitSwitchoverClause(ctx *CommitSwitchoverClauseContext) {
}

// EnterStartStandbyClause is called when production startStandbyClause is entered.
func (s *BaseOracleStatementListener) EnterStartStandbyClause(ctx *StartStandbyClauseContext) {}

// ExitStartStandbyClause is called when production startStandbyClause is exited.
func (s *BaseOracleStatementListener) ExitStartStandbyClause(ctx *StartStandbyClauseContext) {}

// EnterStopStandbyClause is called when production stopStandbyClause is entered.
func (s *BaseOracleStatementListener) EnterStopStandbyClause(ctx *StopStandbyClauseContext) {}

// ExitStopStandbyClause is called when production stopStandbyClause is exited.
func (s *BaseOracleStatementListener) ExitStopStandbyClause(ctx *StopStandbyClauseContext) {}

// EnterSwitchoverClause is called when production switchoverClause is entered.
func (s *BaseOracleStatementListener) EnterSwitchoverClause(ctx *SwitchoverClauseContext) {}

// ExitSwitchoverClause is called when production switchoverClause is exited.
func (s *BaseOracleStatementListener) ExitSwitchoverClause(ctx *SwitchoverClauseContext) {}

// EnterConvertDatabaseClause is called when production convertDatabaseClause is entered.
func (s *BaseOracleStatementListener) EnterConvertDatabaseClause(ctx *ConvertDatabaseClauseContext) {}

// ExitConvertDatabaseClause is called when production convertDatabaseClause is exited.
func (s *BaseOracleStatementListener) ExitConvertDatabaseClause(ctx *ConvertDatabaseClauseContext) {}

// EnterFailoverClause is called when production failoverClause is entered.
func (s *BaseOracleStatementListener) EnterFailoverClause(ctx *FailoverClauseContext) {}

// ExitFailoverClause is called when production failoverClause is exited.
func (s *BaseOracleStatementListener) ExitFailoverClause(ctx *FailoverClauseContext) {}

// EnterDefaultSettingsClauses is called when production defaultSettingsClauses is entered.
func (s *BaseOracleStatementListener) EnterDefaultSettingsClauses(ctx *DefaultSettingsClausesContext) {
}

// ExitDefaultSettingsClauses is called when production defaultSettingsClauses is exited.
func (s *BaseOracleStatementListener) ExitDefaultSettingsClauses(ctx *DefaultSettingsClausesContext) {
}

// EnterSetTimeZoneClause is called when production setTimeZoneClause is entered.
func (s *BaseOracleStatementListener) EnterSetTimeZoneClause(ctx *SetTimeZoneClauseContext) {}

// ExitSetTimeZoneClause is called when production setTimeZoneClause is exited.
func (s *BaseOracleStatementListener) ExitSetTimeZoneClause(ctx *SetTimeZoneClauseContext) {}

// EnterTimeZoneRegion is called when production timeZoneRegion is entered.
func (s *BaseOracleStatementListener) EnterTimeZoneRegion(ctx *TimeZoneRegionContext) {}

// ExitTimeZoneRegion is called when production timeZoneRegion is exited.
func (s *BaseOracleStatementListener) ExitTimeZoneRegion(ctx *TimeZoneRegionContext) {}

// EnterFlashbackModeClause is called when production flashbackModeClause is entered.
func (s *BaseOracleStatementListener) EnterFlashbackModeClause(ctx *FlashbackModeClauseContext) {}

// ExitFlashbackModeClause is called when production flashbackModeClause is exited.
func (s *BaseOracleStatementListener) ExitFlashbackModeClause(ctx *FlashbackModeClauseContext) {}

// EnterUndoModeClause is called when production undoModeClause is entered.
func (s *BaseOracleStatementListener) EnterUndoModeClause(ctx *UndoModeClauseContext) {}

// ExitUndoModeClause is called when production undoModeClause is exited.
func (s *BaseOracleStatementListener) ExitUndoModeClause(ctx *UndoModeClauseContext) {}

// EnterMoveDatafileClause is called when production moveDatafileClause is entered.
func (s *BaseOracleStatementListener) EnterMoveDatafileClause(ctx *MoveDatafileClauseContext) {}

// ExitMoveDatafileClause is called when production moveDatafileClause is exited.
func (s *BaseOracleStatementListener) ExitMoveDatafileClause(ctx *MoveDatafileClauseContext) {}

// EnterInstanceClauses is called when production instanceClauses is entered.
func (s *BaseOracleStatementListener) EnterInstanceClauses(ctx *InstanceClausesContext) {}

// ExitInstanceClauses is called when production instanceClauses is exited.
func (s *BaseOracleStatementListener) ExitInstanceClauses(ctx *InstanceClausesContext) {}

// EnterSecurityClause is called when production securityClause is entered.
func (s *BaseOracleStatementListener) EnterSecurityClause(ctx *SecurityClauseContext) {}

// ExitSecurityClause is called when production securityClause is exited.
func (s *BaseOracleStatementListener) ExitSecurityClause(ctx *SecurityClauseContext) {}

// EnterPrepareClause is called when production prepareClause is entered.
func (s *BaseOracleStatementListener) EnterPrepareClause(ctx *PrepareClauseContext) {}

// ExitPrepareClause is called when production prepareClause is exited.
func (s *BaseOracleStatementListener) ExitPrepareClause(ctx *PrepareClauseContext) {}

// EnterDropMirrorCopy is called when production dropMirrorCopy is entered.
func (s *BaseOracleStatementListener) EnterDropMirrorCopy(ctx *DropMirrorCopyContext) {}

// ExitDropMirrorCopy is called when production dropMirrorCopy is exited.
func (s *BaseOracleStatementListener) ExitDropMirrorCopy(ctx *DropMirrorCopyContext) {}

// EnterLostWriteProtection is called when production lostWriteProtection is entered.
func (s *BaseOracleStatementListener) EnterLostWriteProtection(ctx *LostWriteProtectionContext) {}

// ExitLostWriteProtection is called when production lostWriteProtection is exited.
func (s *BaseOracleStatementListener) ExitLostWriteProtection(ctx *LostWriteProtectionContext) {}

// EnterCdbFleetClauses is called when production cdbFleetClauses is entered.
func (s *BaseOracleStatementListener) EnterCdbFleetClauses(ctx *CdbFleetClausesContext) {}

// ExitCdbFleetClauses is called when production cdbFleetClauses is exited.
func (s *BaseOracleStatementListener) ExitCdbFleetClauses(ctx *CdbFleetClausesContext) {}

// EnterLeadCdbClause is called when production leadCdbClause is entered.
func (s *BaseOracleStatementListener) EnterLeadCdbClause(ctx *LeadCdbClauseContext) {}

// ExitLeadCdbClause is called when production leadCdbClause is exited.
func (s *BaseOracleStatementListener) ExitLeadCdbClause(ctx *LeadCdbClauseContext) {}

// EnterLeadCdbUriClause is called when production leadCdbUriClause is entered.
func (s *BaseOracleStatementListener) EnterLeadCdbUriClause(ctx *LeadCdbUriClauseContext) {}

// ExitLeadCdbUriClause is called when production leadCdbUriClause is exited.
func (s *BaseOracleStatementListener) ExitLeadCdbUriClause(ctx *LeadCdbUriClauseContext) {}

// EnterPropertyClause is called when production propertyClause is entered.
func (s *BaseOracleStatementListener) EnterPropertyClause(ctx *PropertyClauseContext) {}

// ExitPropertyClause is called when production propertyClause is exited.
func (s *BaseOracleStatementListener) ExitPropertyClause(ctx *PropertyClauseContext) {}

// EnterAlterSystem is called when production alterSystem is entered.
func (s *BaseOracleStatementListener) EnterAlterSystem(ctx *AlterSystemContext) {}

// ExitAlterSystem is called when production alterSystem is exited.
func (s *BaseOracleStatementListener) ExitAlterSystem(ctx *AlterSystemContext) {}

// EnterAlterSystemOption is called when production alterSystemOption is entered.
func (s *BaseOracleStatementListener) EnterAlterSystemOption(ctx *AlterSystemOptionContext) {}

// ExitAlterSystemOption is called when production alterSystemOption is exited.
func (s *BaseOracleStatementListener) ExitAlterSystemOption(ctx *AlterSystemOptionContext) {}

// EnterArchiveLogClause is called when production archiveLogClause is entered.
func (s *BaseOracleStatementListener) EnterArchiveLogClause(ctx *ArchiveLogClauseContext) {}

// ExitArchiveLogClause is called when production archiveLogClause is exited.
func (s *BaseOracleStatementListener) ExitArchiveLogClause(ctx *ArchiveLogClauseContext) {}

// EnterCheckpointClause is called when production checkpointClause is entered.
func (s *BaseOracleStatementListener) EnterCheckpointClause(ctx *CheckpointClauseContext) {}

// ExitCheckpointClause is called when production checkpointClause is exited.
func (s *BaseOracleStatementListener) ExitCheckpointClause(ctx *CheckpointClauseContext) {}

// EnterCheckDatafilesClause is called when production checkDatafilesClause is entered.
func (s *BaseOracleStatementListener) EnterCheckDatafilesClause(ctx *CheckDatafilesClauseContext) {}

// ExitCheckDatafilesClause is called when production checkDatafilesClause is exited.
func (s *BaseOracleStatementListener) ExitCheckDatafilesClause(ctx *CheckDatafilesClauseContext) {}

// EnterDistributedRecovClauses is called when production distributedRecovClauses is entered.
func (s *BaseOracleStatementListener) EnterDistributedRecovClauses(ctx *DistributedRecovClausesContext) {
}

// ExitDistributedRecovClauses is called when production distributedRecovClauses is exited.
func (s *BaseOracleStatementListener) ExitDistributedRecovClauses(ctx *DistributedRecovClausesContext) {
}

// EnterFlushClause is called when production flushClause is entered.
func (s *BaseOracleStatementListener) EnterFlushClause(ctx *FlushClauseContext) {}

// ExitFlushClause is called when production flushClause is exited.
func (s *BaseOracleStatementListener) ExitFlushClause(ctx *FlushClauseContext) {}

// EnterEndSessionClauses is called when production endSessionClauses is entered.
func (s *BaseOracleStatementListener) EnterEndSessionClauses(ctx *EndSessionClausesContext) {}

// ExitEndSessionClauses is called when production endSessionClauses is exited.
func (s *BaseOracleStatementListener) ExitEndSessionClauses(ctx *EndSessionClausesContext) {}

// EnterAlterSystemSwitchLogfileClause is called when production alterSystemSwitchLogfileClause is entered.
func (s *BaseOracleStatementListener) EnterAlterSystemSwitchLogfileClause(ctx *AlterSystemSwitchLogfileClauseContext) {
}

// ExitAlterSystemSwitchLogfileClause is called when production alterSystemSwitchLogfileClause is exited.
func (s *BaseOracleStatementListener) ExitAlterSystemSwitchLogfileClause(ctx *AlterSystemSwitchLogfileClauseContext) {
}

// EnterSuspendResumeClause is called when production suspendResumeClause is entered.
func (s *BaseOracleStatementListener) EnterSuspendResumeClause(ctx *SuspendResumeClauseContext) {}

// ExitSuspendResumeClause is called when production suspendResumeClause is exited.
func (s *BaseOracleStatementListener) ExitSuspendResumeClause(ctx *SuspendResumeClauseContext) {}

// EnterQuiesceClauses is called when production quiesceClauses is entered.
func (s *BaseOracleStatementListener) EnterQuiesceClauses(ctx *QuiesceClausesContext) {}

// ExitQuiesceClauses is called when production quiesceClauses is exited.
func (s *BaseOracleStatementListener) ExitQuiesceClauses(ctx *QuiesceClausesContext) {}

// EnterRollingMigrationClauses is called when production rollingMigrationClauses is entered.
func (s *BaseOracleStatementListener) EnterRollingMigrationClauses(ctx *RollingMigrationClausesContext) {
}

// ExitRollingMigrationClauses is called when production rollingMigrationClauses is exited.
func (s *BaseOracleStatementListener) ExitRollingMigrationClauses(ctx *RollingMigrationClausesContext) {
}

// EnterRollingPatchClauses is called when production rollingPatchClauses is entered.
func (s *BaseOracleStatementListener) EnterRollingPatchClauses(ctx *RollingPatchClausesContext) {}

// ExitRollingPatchClauses is called when production rollingPatchClauses is exited.
func (s *BaseOracleStatementListener) ExitRollingPatchClauses(ctx *RollingPatchClausesContext) {}

// EnterAlterSystemSecurityClauses is called when production alterSystemSecurityClauses is entered.
func (s *BaseOracleStatementListener) EnterAlterSystemSecurityClauses(ctx *AlterSystemSecurityClausesContext) {
}

// ExitAlterSystemSecurityClauses is called when production alterSystemSecurityClauses is exited.
func (s *BaseOracleStatementListener) ExitAlterSystemSecurityClauses(ctx *AlterSystemSecurityClausesContext) {
}

// EnterAffinityClauses is called when production affinityClauses is entered.
func (s *BaseOracleStatementListener) EnterAffinityClauses(ctx *AffinityClausesContext) {}

// ExitAffinityClauses is called when production affinityClauses is exited.
func (s *BaseOracleStatementListener) ExitAffinityClauses(ctx *AffinityClausesContext) {}

// EnterShutdownDispatcherClause is called when production shutdownDispatcherClause is entered.
func (s *BaseOracleStatementListener) EnterShutdownDispatcherClause(ctx *ShutdownDispatcherClauseContext) {
}

// ExitShutdownDispatcherClause is called when production shutdownDispatcherClause is exited.
func (s *BaseOracleStatementListener) ExitShutdownDispatcherClause(ctx *ShutdownDispatcherClauseContext) {
}

// EnterRegisterClause is called when production registerClause is entered.
func (s *BaseOracleStatementListener) EnterRegisterClause(ctx *RegisterClauseContext) {}

// ExitRegisterClause is called when production registerClause is exited.
func (s *BaseOracleStatementListener) ExitRegisterClause(ctx *RegisterClauseContext) {}

// EnterSetClause is called when production setClause is entered.
func (s *BaseOracleStatementListener) EnterSetClause(ctx *SetClauseContext) {}

// ExitSetClause is called when production setClause is exited.
func (s *BaseOracleStatementListener) ExitSetClause(ctx *SetClauseContext) {}

// EnterResetClause is called when production resetClause is entered.
func (s *BaseOracleStatementListener) EnterResetClause(ctx *ResetClauseContext) {}

// ExitResetClause is called when production resetClause is exited.
func (s *BaseOracleStatementListener) ExitResetClause(ctx *ResetClauseContext) {}

// EnterRelocateClientClause is called when production relocateClientClause is entered.
func (s *BaseOracleStatementListener) EnterRelocateClientClause(ctx *RelocateClientClauseContext) {}

// ExitRelocateClientClause is called when production relocateClientClause is exited.
func (s *BaseOracleStatementListener) ExitRelocateClientClause(ctx *RelocateClientClauseContext) {}

// EnterCancelSqlClause is called when production cancelSqlClause is entered.
func (s *BaseOracleStatementListener) EnterCancelSqlClause(ctx *CancelSqlClauseContext) {}

// ExitCancelSqlClause is called when production cancelSqlClause is exited.
func (s *BaseOracleStatementListener) ExitCancelSqlClause(ctx *CancelSqlClauseContext) {}

// EnterFlushPasswordfileMetadataCacheClause is called when production flushPasswordfileMetadataCacheClause is entered.
func (s *BaseOracleStatementListener) EnterFlushPasswordfileMetadataCacheClause(ctx *FlushPasswordfileMetadataCacheClauseContext) {
}

// ExitFlushPasswordfileMetadataCacheClause is called when production flushPasswordfileMetadataCacheClause is exited.
func (s *BaseOracleStatementListener) ExitFlushPasswordfileMetadataCacheClause(ctx *FlushPasswordfileMetadataCacheClauseContext) {
}

// EnterInstanceClause is called when production instanceClause is entered.
func (s *BaseOracleStatementListener) EnterInstanceClause(ctx *InstanceClauseContext) {}

// ExitInstanceClause is called when production instanceClause is exited.
func (s *BaseOracleStatementListener) ExitInstanceClause(ctx *InstanceClauseContext) {}

// EnterSequenceClause is called when production sequenceClause is entered.
func (s *BaseOracleStatementListener) EnterSequenceClause(ctx *SequenceClauseContext) {}

// ExitSequenceClause is called when production sequenceClause is exited.
func (s *BaseOracleStatementListener) ExitSequenceClause(ctx *SequenceClauseContext) {}

// EnterChangeClause is called when production changeClause is entered.
func (s *BaseOracleStatementListener) EnterChangeClause(ctx *ChangeClauseContext) {}

// ExitChangeClause is called when production changeClause is exited.
func (s *BaseOracleStatementListener) ExitChangeClause(ctx *ChangeClauseContext) {}

// EnterCurrentClause is called when production currentClause is entered.
func (s *BaseOracleStatementListener) EnterCurrentClause(ctx *CurrentClauseContext) {}

// ExitCurrentClause is called when production currentClause is exited.
func (s *BaseOracleStatementListener) ExitCurrentClause(ctx *CurrentClauseContext) {}

// EnterGroupClause is called when production groupClause is entered.
func (s *BaseOracleStatementListener) EnterGroupClause(ctx *GroupClauseContext) {}

// ExitGroupClause is called when production groupClause is exited.
func (s *BaseOracleStatementListener) ExitGroupClause(ctx *GroupClauseContext) {}

// EnterLogfileClause is called when production logfileClause is entered.
func (s *BaseOracleStatementListener) EnterLogfileClause(ctx *LogfileClauseContext) {}

// ExitLogfileClause is called when production logfileClause is exited.
func (s *BaseOracleStatementListener) ExitLogfileClause(ctx *LogfileClauseContext) {}

// EnterNextClause is called when production nextClause is entered.
func (s *BaseOracleStatementListener) EnterNextClause(ctx *NextClauseContext) {}

// ExitNextClause is called when production nextClause is exited.
func (s *BaseOracleStatementListener) ExitNextClause(ctx *NextClauseContext) {}

// EnterAllClause is called when production allClause is entered.
func (s *BaseOracleStatementListener) EnterAllClause(ctx *AllClauseContext) {}

// ExitAllClause is called when production allClause is exited.
func (s *BaseOracleStatementListener) ExitAllClause(ctx *AllClauseContext) {}

// EnterToLocationClause is called when production toLocationClause is entered.
func (s *BaseOracleStatementListener) EnterToLocationClause(ctx *ToLocationClauseContext) {}

// ExitToLocationClause is called when production toLocationClause is exited.
func (s *BaseOracleStatementListener) ExitToLocationClause(ctx *ToLocationClauseContext) {}

// EnterFlushClauseOption is called when production flushClauseOption is entered.
func (s *BaseOracleStatementListener) EnterFlushClauseOption(ctx *FlushClauseOptionContext) {}

// ExitFlushClauseOption is called when production flushClauseOption is exited.
func (s *BaseOracleStatementListener) ExitFlushClauseOption(ctx *FlushClauseOptionContext) {}

// EnterDisconnectSessionClause is called when production disconnectSessionClause is entered.
func (s *BaseOracleStatementListener) EnterDisconnectSessionClause(ctx *DisconnectSessionClauseContext) {
}

// ExitDisconnectSessionClause is called when production disconnectSessionClause is exited.
func (s *BaseOracleStatementListener) ExitDisconnectSessionClause(ctx *DisconnectSessionClauseContext) {
}

// EnterKillSessionClause is called when production killSessionClause is entered.
func (s *BaseOracleStatementListener) EnterKillSessionClause(ctx *KillSessionClauseContext) {}

// ExitKillSessionClause is called when production killSessionClause is exited.
func (s *BaseOracleStatementListener) ExitKillSessionClause(ctx *KillSessionClauseContext) {}

// EnterStartRollingMigrationClause is called when production startRollingMigrationClause is entered.
func (s *BaseOracleStatementListener) EnterStartRollingMigrationClause(ctx *StartRollingMigrationClauseContext) {
}

// ExitStartRollingMigrationClause is called when production startRollingMigrationClause is exited.
func (s *BaseOracleStatementListener) ExitStartRollingMigrationClause(ctx *StartRollingMigrationClauseContext) {
}

// EnterStopRollingMigrationClause is called when production stopRollingMigrationClause is entered.
func (s *BaseOracleStatementListener) EnterStopRollingMigrationClause(ctx *StopRollingMigrationClauseContext) {
}

// ExitStopRollingMigrationClause is called when production stopRollingMigrationClause is exited.
func (s *BaseOracleStatementListener) ExitStopRollingMigrationClause(ctx *StopRollingMigrationClauseContext) {
}

// EnterStartRollingPatchClause is called when production startRollingPatchClause is entered.
func (s *BaseOracleStatementListener) EnterStartRollingPatchClause(ctx *StartRollingPatchClauseContext) {
}

// ExitStartRollingPatchClause is called when production startRollingPatchClause is exited.
func (s *BaseOracleStatementListener) ExitStartRollingPatchClause(ctx *StartRollingPatchClauseContext) {
}

// EnterStopRollingPatchClause is called when production stopRollingPatchClause is entered.
func (s *BaseOracleStatementListener) EnterStopRollingPatchClause(ctx *StopRollingPatchClauseContext) {
}

// ExitStopRollingPatchClause is called when production stopRollingPatchClause is exited.
func (s *BaseOracleStatementListener) ExitStopRollingPatchClause(ctx *StopRollingPatchClauseContext) {
}

// EnterRestrictedSessionClause is called when production restrictedSessionClause is entered.
func (s *BaseOracleStatementListener) EnterRestrictedSessionClause(ctx *RestrictedSessionClauseContext) {
}

// ExitRestrictedSessionClause is called when production restrictedSessionClause is exited.
func (s *BaseOracleStatementListener) ExitRestrictedSessionClause(ctx *RestrictedSessionClauseContext) {
}

// EnterSetEncryptionWalletOpenClause is called when production setEncryptionWalletOpenClause is entered.
func (s *BaseOracleStatementListener) EnterSetEncryptionWalletOpenClause(ctx *SetEncryptionWalletOpenClauseContext) {
}

// ExitSetEncryptionWalletOpenClause is called when production setEncryptionWalletOpenClause is exited.
func (s *BaseOracleStatementListener) ExitSetEncryptionWalletOpenClause(ctx *SetEncryptionWalletOpenClauseContext) {
}

// EnterSetEncryptionWalletCloseClause is called when production setEncryptionWalletCloseClause is entered.
func (s *BaseOracleStatementListener) EnterSetEncryptionWalletCloseClause(ctx *SetEncryptionWalletCloseClauseContext) {
}

// ExitSetEncryptionWalletCloseClause is called when production setEncryptionWalletCloseClause is exited.
func (s *BaseOracleStatementListener) ExitSetEncryptionWalletCloseClause(ctx *SetEncryptionWalletCloseClauseContext) {
}

// EnterSetEncryptionKeyClause is called when production setEncryptionKeyClause is entered.
func (s *BaseOracleStatementListener) EnterSetEncryptionKeyClause(ctx *SetEncryptionKeyClauseContext) {
}

// ExitSetEncryptionKeyClause is called when production setEncryptionKeyClause is exited.
func (s *BaseOracleStatementListener) ExitSetEncryptionKeyClause(ctx *SetEncryptionKeyClauseContext) {
}

// EnterEnableAffinityClause is called when production enableAffinityClause is entered.
func (s *BaseOracleStatementListener) EnterEnableAffinityClause(ctx *EnableAffinityClauseContext) {}

// ExitEnableAffinityClause is called when production enableAffinityClause is exited.
func (s *BaseOracleStatementListener) ExitEnableAffinityClause(ctx *EnableAffinityClauseContext) {}

// EnterDisableAffinityClause is called when production disableAffinityClause is entered.
func (s *BaseOracleStatementListener) EnterDisableAffinityClause(ctx *DisableAffinityClauseContext) {}

// ExitDisableAffinityClause is called when production disableAffinityClause is exited.
func (s *BaseOracleStatementListener) ExitDisableAffinityClause(ctx *DisableAffinityClauseContext) {}

// EnterAlterSystemSetClause is called when production alterSystemSetClause is entered.
func (s *BaseOracleStatementListener) EnterAlterSystemSetClause(ctx *AlterSystemSetClauseContext) {}

// ExitAlterSystemSetClause is called when production alterSystemSetClause is exited.
func (s *BaseOracleStatementListener) ExitAlterSystemSetClause(ctx *AlterSystemSetClauseContext) {}

// EnterAlterSystemResetClause is called when production alterSystemResetClause is entered.
func (s *BaseOracleStatementListener) EnterAlterSystemResetClause(ctx *AlterSystemResetClauseContext) {
}

// ExitAlterSystemResetClause is called when production alterSystemResetClause is exited.
func (s *BaseOracleStatementListener) ExitAlterSystemResetClause(ctx *AlterSystemResetClauseContext) {
}

// EnterSharedPoolClause is called when production sharedPoolClause is entered.
func (s *BaseOracleStatementListener) EnterSharedPoolClause(ctx *SharedPoolClauseContext) {}

// ExitSharedPoolClause is called when production sharedPoolClause is exited.
func (s *BaseOracleStatementListener) ExitSharedPoolClause(ctx *SharedPoolClauseContext) {}

// EnterGlobalContextClause is called when production globalContextClause is entered.
func (s *BaseOracleStatementListener) EnterGlobalContextClause(ctx *GlobalContextClauseContext) {}

// ExitGlobalContextClause is called when production globalContextClause is exited.
func (s *BaseOracleStatementListener) ExitGlobalContextClause(ctx *GlobalContextClauseContext) {}

// EnterBufferCacheClause is called when production bufferCacheClause is entered.
func (s *BaseOracleStatementListener) EnterBufferCacheClause(ctx *BufferCacheClauseContext) {}

// ExitBufferCacheClause is called when production bufferCacheClause is exited.
func (s *BaseOracleStatementListener) ExitBufferCacheClause(ctx *BufferCacheClauseContext) {}

// EnterFlashCacheClause is called when production flashCacheClause is entered.
func (s *BaseOracleStatementListener) EnterFlashCacheClause(ctx *FlashCacheClauseContext) {}

// ExitFlashCacheClause is called when production flashCacheClause is exited.
func (s *BaseOracleStatementListener) ExitFlashCacheClause(ctx *FlashCacheClauseContext) {}

// EnterRedoToClause is called when production redoToClause is entered.
func (s *BaseOracleStatementListener) EnterRedoToClause(ctx *RedoToClauseContext) {}

// ExitRedoToClause is called when production redoToClause is exited.
func (s *BaseOracleStatementListener) ExitRedoToClause(ctx *RedoToClauseContext) {}

// EnterIdentifiedByWalletPassword is called when production identifiedByWalletPassword is entered.
func (s *BaseOracleStatementListener) EnterIdentifiedByWalletPassword(ctx *IdentifiedByWalletPasswordContext) {
}

// ExitIdentifiedByWalletPassword is called when production identifiedByWalletPassword is exited.
func (s *BaseOracleStatementListener) ExitIdentifiedByWalletPassword(ctx *IdentifiedByWalletPasswordContext) {
}

// EnterIdentifiedByHsmAuthString is called when production identifiedByHsmAuthString is entered.
func (s *BaseOracleStatementListener) EnterIdentifiedByHsmAuthString(ctx *IdentifiedByHsmAuthStringContext) {
}

// ExitIdentifiedByHsmAuthString is called when production identifiedByHsmAuthString is exited.
func (s *BaseOracleStatementListener) ExitIdentifiedByHsmAuthString(ctx *IdentifiedByHsmAuthStringContext) {
}

// EnterSetParameterClause is called when production setParameterClause is entered.
func (s *BaseOracleStatementListener) EnterSetParameterClause(ctx *SetParameterClauseContext) {}

// ExitSetParameterClause is called when production setParameterClause is exited.
func (s *BaseOracleStatementListener) ExitSetParameterClause(ctx *SetParameterClauseContext) {}

// EnterUseStoredOutlinesClause is called when production useStoredOutlinesClause is entered.
func (s *BaseOracleStatementListener) EnterUseStoredOutlinesClause(ctx *UseStoredOutlinesClauseContext) {
}

// ExitUseStoredOutlinesClause is called when production useStoredOutlinesClause is exited.
func (s *BaseOracleStatementListener) ExitUseStoredOutlinesClause(ctx *UseStoredOutlinesClauseContext) {
}

// EnterGlobalTopicEnabledClause is called when production globalTopicEnabledClause is entered.
func (s *BaseOracleStatementListener) EnterGlobalTopicEnabledClause(ctx *GlobalTopicEnabledClauseContext) {
}

// ExitGlobalTopicEnabledClause is called when production globalTopicEnabledClause is exited.
func (s *BaseOracleStatementListener) ExitGlobalTopicEnabledClause(ctx *GlobalTopicEnabledClauseContext) {
}

// EnterAlterSystemCommentClause is called when production alterSystemCommentClause is entered.
func (s *BaseOracleStatementListener) EnterAlterSystemCommentClause(ctx *AlterSystemCommentClauseContext) {
}

// ExitAlterSystemCommentClause is called when production alterSystemCommentClause is exited.
func (s *BaseOracleStatementListener) ExitAlterSystemCommentClause(ctx *AlterSystemCommentClauseContext) {
}

// EnterContainerCurrentAllClause is called when production containerCurrentAllClause is entered.
func (s *BaseOracleStatementListener) EnterContainerCurrentAllClause(ctx *ContainerCurrentAllClauseContext) {
}

// ExitContainerCurrentAllClause is called when production containerCurrentAllClause is exited.
func (s *BaseOracleStatementListener) ExitContainerCurrentAllClause(ctx *ContainerCurrentAllClauseContext) {
}

// EnterScopeClause is called when production scopeClause is entered.
func (s *BaseOracleStatementListener) EnterScopeClause(ctx *ScopeClauseContext) {}

// ExitScopeClause is called when production scopeClause is exited.
func (s *BaseOracleStatementListener) ExitScopeClause(ctx *ScopeClauseContext) {}

// EnterAnalyze is called when production analyze is entered.
func (s *BaseOracleStatementListener) EnterAnalyze(ctx *AnalyzeContext) {}

// ExitAnalyze is called when production analyze is exited.
func (s *BaseOracleStatementListener) ExitAnalyze(ctx *AnalyzeContext) {}

// EnterPartitionExtensionClause is called when production partitionExtensionClause is entered.
func (s *BaseOracleStatementListener) EnterPartitionExtensionClause(ctx *PartitionExtensionClauseContext) {
}

// ExitPartitionExtensionClause is called when production partitionExtensionClause is exited.
func (s *BaseOracleStatementListener) ExitPartitionExtensionClause(ctx *PartitionExtensionClauseContext) {
}

// EnterValidationClauses is called when production validationClauses is entered.
func (s *BaseOracleStatementListener) EnterValidationClauses(ctx *ValidationClausesContext) {}

// ExitValidationClauses is called when production validationClauses is exited.
func (s *BaseOracleStatementListener) ExitValidationClauses(ctx *ValidationClausesContext) {}

// EnterAssociateStatistics is called when production associateStatistics is entered.
func (s *BaseOracleStatementListener) EnterAssociateStatistics(ctx *AssociateStatisticsContext) {}

// ExitAssociateStatistics is called when production associateStatistics is exited.
func (s *BaseOracleStatementListener) ExitAssociateStatistics(ctx *AssociateStatisticsContext) {}

// EnterColumnAssociation is called when production columnAssociation is entered.
func (s *BaseOracleStatementListener) EnterColumnAssociation(ctx *ColumnAssociationContext) {}

// ExitColumnAssociation is called when production columnAssociation is exited.
func (s *BaseOracleStatementListener) ExitColumnAssociation(ctx *ColumnAssociationContext) {}

// EnterFunctionAssociation is called when production functionAssociation is entered.
func (s *BaseOracleStatementListener) EnterFunctionAssociation(ctx *FunctionAssociationContext) {}

// ExitFunctionAssociation is called when production functionAssociation is exited.
func (s *BaseOracleStatementListener) ExitFunctionAssociation(ctx *FunctionAssociationContext) {}

// EnterStorageTableClause is called when production storageTableClause is entered.
func (s *BaseOracleStatementListener) EnterStorageTableClause(ctx *StorageTableClauseContext) {}

// ExitStorageTableClause is called when production storageTableClause is exited.
func (s *BaseOracleStatementListener) ExitStorageTableClause(ctx *StorageTableClauseContext) {}

// EnterUsingStatisticsType is called when production usingStatisticsType is entered.
func (s *BaseOracleStatementListener) EnterUsingStatisticsType(ctx *UsingStatisticsTypeContext) {}

// ExitUsingStatisticsType is called when production usingStatisticsType is exited.
func (s *BaseOracleStatementListener) ExitUsingStatisticsType(ctx *UsingStatisticsTypeContext) {}

// EnterDefaultCostClause is called when production defaultCostClause is entered.
func (s *BaseOracleStatementListener) EnterDefaultCostClause(ctx *DefaultCostClauseContext) {}

// ExitDefaultCostClause is called when production defaultCostClause is exited.
func (s *BaseOracleStatementListener) ExitDefaultCostClause(ctx *DefaultCostClauseContext) {}

// EnterDefaultSelectivityClause is called when production defaultSelectivityClause is entered.
func (s *BaseOracleStatementListener) EnterDefaultSelectivityClause(ctx *DefaultSelectivityClauseContext) {
}

// ExitDefaultSelectivityClause is called when production defaultSelectivityClause is exited.
func (s *BaseOracleStatementListener) ExitDefaultSelectivityClause(ctx *DefaultSelectivityClauseContext) {
}

// EnterDisassociateStatistics is called when production disassociateStatistics is entered.
func (s *BaseOracleStatementListener) EnterDisassociateStatistics(ctx *DisassociateStatisticsContext) {
}

// ExitDisassociateStatistics is called when production disassociateStatistics is exited.
func (s *BaseOracleStatementListener) ExitDisassociateStatistics(ctx *DisassociateStatisticsContext) {
}

// EnterAudit is called when production audit is entered.
func (s *BaseOracleStatementListener) EnterAudit(ctx *AuditContext) {}

// ExitAudit is called when production audit is exited.
func (s *BaseOracleStatementListener) ExitAudit(ctx *AuditContext) {}

// EnterNoAudit is called when production noAudit is entered.
func (s *BaseOracleStatementListener) EnterNoAudit(ctx *NoAuditContext) {}

// ExitNoAudit is called when production noAudit is exited.
func (s *BaseOracleStatementListener) ExitNoAudit(ctx *NoAuditContext) {}

// EnterAuditPolicyClause is called when production auditPolicyClause is entered.
func (s *BaseOracleStatementListener) EnterAuditPolicyClause(ctx *AuditPolicyClauseContext) {}

// ExitAuditPolicyClause is called when production auditPolicyClause is exited.
func (s *BaseOracleStatementListener) ExitAuditPolicyClause(ctx *AuditPolicyClauseContext) {}

// EnterNoAuditPolicyClause is called when production noAuditPolicyClause is entered.
func (s *BaseOracleStatementListener) EnterNoAuditPolicyClause(ctx *NoAuditPolicyClauseContext) {}

// ExitNoAuditPolicyClause is called when production noAuditPolicyClause is exited.
func (s *BaseOracleStatementListener) ExitNoAuditPolicyClause(ctx *NoAuditPolicyClauseContext) {}

// EnterByUsersWithRoles is called when production byUsersWithRoles is entered.
func (s *BaseOracleStatementListener) EnterByUsersWithRoles(ctx *ByUsersWithRolesContext) {}

// ExitByUsersWithRoles is called when production byUsersWithRoles is exited.
func (s *BaseOracleStatementListener) ExitByUsersWithRoles(ctx *ByUsersWithRolesContext) {}

// EnterContextClause is called when production contextClause is entered.
func (s *BaseOracleStatementListener) EnterContextClause(ctx *ContextClauseContext) {}

// ExitContextClause is called when production contextClause is exited.
func (s *BaseOracleStatementListener) ExitContextClause(ctx *ContextClauseContext) {}

// EnterContextNamespaceAttributesClause is called when production contextNamespaceAttributesClause is entered.
func (s *BaseOracleStatementListener) EnterContextNamespaceAttributesClause(ctx *ContextNamespaceAttributesClauseContext) {
}

// ExitContextNamespaceAttributesClause is called when production contextNamespaceAttributesClause is exited.
func (s *BaseOracleStatementListener) ExitContextNamespaceAttributesClause(ctx *ContextNamespaceAttributesClauseContext) {
}

// EnterComment is called when production comment is entered.
func (s *BaseOracleStatementListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseOracleStatementListener) ExitComment(ctx *CommentContext) {}

// EnterFlashbackDatabase is called when production flashbackDatabase is entered.
func (s *BaseOracleStatementListener) EnterFlashbackDatabase(ctx *FlashbackDatabaseContext) {}

// ExitFlashbackDatabase is called when production flashbackDatabase is exited.
func (s *BaseOracleStatementListener) ExitFlashbackDatabase(ctx *FlashbackDatabaseContext) {}

// EnterScnTimestampClause is called when production scnTimestampClause is entered.
func (s *BaseOracleStatementListener) EnterScnTimestampClause(ctx *ScnTimestampClauseContext) {}

// ExitScnTimestampClause is called when production scnTimestampClause is exited.
func (s *BaseOracleStatementListener) ExitScnTimestampClause(ctx *ScnTimestampClauseContext) {}

// EnterRestorePointClause is called when production restorePointClause is entered.
func (s *BaseOracleStatementListener) EnterRestorePointClause(ctx *RestorePointClauseContext) {}

// ExitRestorePointClause is called when production restorePointClause is exited.
func (s *BaseOracleStatementListener) ExitRestorePointClause(ctx *RestorePointClauseContext) {}

// EnterFlashbackTable is called when production flashbackTable is entered.
func (s *BaseOracleStatementListener) EnterFlashbackTable(ctx *FlashbackTableContext) {}

// ExitFlashbackTable is called when production flashbackTable is exited.
func (s *BaseOracleStatementListener) ExitFlashbackTable(ctx *FlashbackTableContext) {}

// EnterRenameToTable is called when production renameToTable is entered.
func (s *BaseOracleStatementListener) EnterRenameToTable(ctx *RenameToTableContext) {}

// ExitRenameToTable is called when production renameToTable is exited.
func (s *BaseOracleStatementListener) ExitRenameToTable(ctx *RenameToTableContext) {}

// EnterPurge is called when production purge is entered.
func (s *BaseOracleStatementListener) EnterPurge(ctx *PurgeContext) {}

// ExitPurge is called when production purge is exited.
func (s *BaseOracleStatementListener) ExitPurge(ctx *PurgeContext) {}

// EnterRename is called when production rename is entered.
func (s *BaseOracleStatementListener) EnterRename(ctx *RenameContext) {}

// ExitRename is called when production rename is exited.
func (s *BaseOracleStatementListener) ExitRename(ctx *RenameContext) {}

// EnterCreateDatabase is called when production createDatabase is entered.
func (s *BaseOracleStatementListener) EnterCreateDatabase(ctx *CreateDatabaseContext) {}

// ExitCreateDatabase is called when production createDatabase is exited.
func (s *BaseOracleStatementListener) ExitCreateDatabase(ctx *CreateDatabaseContext) {}

// EnterCreateDatabaseClauses is called when production createDatabaseClauses is entered.
func (s *BaseOracleStatementListener) EnterCreateDatabaseClauses(ctx *CreateDatabaseClausesContext) {}

// ExitCreateDatabaseClauses is called when production createDatabaseClauses is exited.
func (s *BaseOracleStatementListener) ExitCreateDatabaseClauses(ctx *CreateDatabaseClausesContext) {}

// EnterDatabaseLoggingClauses is called when production databaseLoggingClauses is entered.
func (s *BaseOracleStatementListener) EnterDatabaseLoggingClauses(ctx *DatabaseLoggingClausesContext) {
}

// ExitDatabaseLoggingClauses is called when production databaseLoggingClauses is exited.
func (s *BaseOracleStatementListener) ExitDatabaseLoggingClauses(ctx *DatabaseLoggingClausesContext) {
}

// EnterTablespaceClauses is called when production tablespaceClauses is entered.
func (s *BaseOracleStatementListener) EnterTablespaceClauses(ctx *TablespaceClausesContext) {}

// ExitTablespaceClauses is called when production tablespaceClauses is exited.
func (s *BaseOracleStatementListener) ExitTablespaceClauses(ctx *TablespaceClausesContext) {}

// EnterDefaultTablespace is called when production defaultTablespace is entered.
func (s *BaseOracleStatementListener) EnterDefaultTablespace(ctx *DefaultTablespaceContext) {}

// ExitDefaultTablespace is called when production defaultTablespace is exited.
func (s *BaseOracleStatementListener) ExitDefaultTablespace(ctx *DefaultTablespaceContext) {}

// EnterDefaultTempTablespace is called when production defaultTempTablespace is entered.
func (s *BaseOracleStatementListener) EnterDefaultTempTablespace(ctx *DefaultTempTablespaceContext) {}

// ExitDefaultTempTablespace is called when production defaultTempTablespace is exited.
func (s *BaseOracleStatementListener) ExitDefaultTempTablespace(ctx *DefaultTempTablespaceContext) {}

// EnterUndoTablespace is called when production undoTablespace is entered.
func (s *BaseOracleStatementListener) EnterUndoTablespace(ctx *UndoTablespaceContext) {}

// ExitUndoTablespace is called when production undoTablespace is exited.
func (s *BaseOracleStatementListener) ExitUndoTablespace(ctx *UndoTablespaceContext) {}

// EnterBigOrSmallFiles is called when production bigOrSmallFiles is entered.
func (s *BaseOracleStatementListener) EnterBigOrSmallFiles(ctx *BigOrSmallFilesContext) {}

// ExitBigOrSmallFiles is called when production bigOrSmallFiles is exited.
func (s *BaseOracleStatementListener) ExitBigOrSmallFiles(ctx *BigOrSmallFilesContext) {}

// EnterExtentManagementClause is called when production extentManagementClause is entered.
func (s *BaseOracleStatementListener) EnterExtentManagementClause(ctx *ExtentManagementClauseContext) {
}

// ExitExtentManagementClause is called when production extentManagementClause is exited.
func (s *BaseOracleStatementListener) ExitExtentManagementClause(ctx *ExtentManagementClauseContext) {
}

// EnterEnablePluggableDatabase is called when production enablePluggableDatabase is entered.
func (s *BaseOracleStatementListener) EnterEnablePluggableDatabase(ctx *EnablePluggableDatabaseContext) {
}

// ExitEnablePluggableDatabase is called when production enablePluggableDatabase is exited.
func (s *BaseOracleStatementListener) ExitEnablePluggableDatabase(ctx *EnablePluggableDatabaseContext) {
}

// EnterFileNameConvert is called when production fileNameConvert is entered.
func (s *BaseOracleStatementListener) EnterFileNameConvert(ctx *FileNameConvertContext) {}

// ExitFileNameConvert is called when production fileNameConvert is exited.
func (s *BaseOracleStatementListener) ExitFileNameConvert(ctx *FileNameConvertContext) {}

// EnterReplaceFileNamePattern is called when production replaceFileNamePattern is entered.
func (s *BaseOracleStatementListener) EnterReplaceFileNamePattern(ctx *ReplaceFileNamePatternContext) {
}

// ExitReplaceFileNamePattern is called when production replaceFileNamePattern is exited.
func (s *BaseOracleStatementListener) ExitReplaceFileNamePattern(ctx *ReplaceFileNamePatternContext) {
}

// EnterTablespaceDatafileClauses is called when production tablespaceDatafileClauses is entered.
func (s *BaseOracleStatementListener) EnterTablespaceDatafileClauses(ctx *TablespaceDatafileClausesContext) {
}

// ExitTablespaceDatafileClauses is called when production tablespaceDatafileClauses is exited.
func (s *BaseOracleStatementListener) ExitTablespaceDatafileClauses(ctx *TablespaceDatafileClausesContext) {
}

// EnterCreateDatabaseLink is called when production createDatabaseLink is entered.
func (s *BaseOracleStatementListener) EnterCreateDatabaseLink(ctx *CreateDatabaseLinkContext) {}

// ExitCreateDatabaseLink is called when production createDatabaseLink is exited.
func (s *BaseOracleStatementListener) ExitCreateDatabaseLink(ctx *CreateDatabaseLinkContext) {}

// EnterConnectToClause is called when production connectToClause is entered.
func (s *BaseOracleStatementListener) EnterConnectToClause(ctx *ConnectToClauseContext) {}

// ExitConnectToClause is called when production connectToClause is exited.
func (s *BaseOracleStatementListener) ExitConnectToClause(ctx *ConnectToClauseContext) {}

// EnterDbLinkAuthentication is called when production dbLinkAuthentication is entered.
func (s *BaseOracleStatementListener) EnterDbLinkAuthentication(ctx *DbLinkAuthenticationContext) {}

// ExitDbLinkAuthentication is called when production dbLinkAuthentication is exited.
func (s *BaseOracleStatementListener) ExitDbLinkAuthentication(ctx *DbLinkAuthenticationContext) {}

// EnterSetTransaction is called when production setTransaction is entered.
func (s *BaseOracleStatementListener) EnterSetTransaction(ctx *SetTransactionContext) {}

// ExitSetTransaction is called when production setTransaction is exited.
func (s *BaseOracleStatementListener) ExitSetTransaction(ctx *SetTransactionContext) {}

// EnterCommit is called when production commit is entered.
func (s *BaseOracleStatementListener) EnterCommit(ctx *CommitContext) {}

// ExitCommit is called when production commit is exited.
func (s *BaseOracleStatementListener) ExitCommit(ctx *CommitContext) {}

// EnterCommentClause is called when production commentClause is entered.
func (s *BaseOracleStatementListener) EnterCommentClause(ctx *CommentClauseContext) {}

// ExitCommentClause is called when production commentClause is exited.
func (s *BaseOracleStatementListener) ExitCommentClause(ctx *CommentClauseContext) {}

// EnterWriteClause is called when production writeClause is entered.
func (s *BaseOracleStatementListener) EnterWriteClause(ctx *WriteClauseContext) {}

// ExitWriteClause is called when production writeClause is exited.
func (s *BaseOracleStatementListener) ExitWriteClause(ctx *WriteClauseContext) {}

// EnterForceClause is called when production forceClause is entered.
func (s *BaseOracleStatementListener) EnterForceClause(ctx *ForceClauseContext) {}

// ExitForceClause is called when production forceClause is exited.
func (s *BaseOracleStatementListener) ExitForceClause(ctx *ForceClauseContext) {}

// EnterRollback is called when production rollback is entered.
func (s *BaseOracleStatementListener) EnterRollback(ctx *RollbackContext) {}

// ExitRollback is called when production rollback is exited.
func (s *BaseOracleStatementListener) ExitRollback(ctx *RollbackContext) {}

// EnterSavepointClause is called when production savepointClause is entered.
func (s *BaseOracleStatementListener) EnterSavepointClause(ctx *SavepointClauseContext) {}

// ExitSavepointClause is called when production savepointClause is exited.
func (s *BaseOracleStatementListener) ExitSavepointClause(ctx *SavepointClauseContext) {}

// EnterSavepoint is called when production savepoint is entered.
func (s *BaseOracleStatementListener) EnterSavepoint(ctx *SavepointContext) {}

// ExitSavepoint is called when production savepoint is exited.
func (s *BaseOracleStatementListener) ExitSavepoint(ctx *SavepointContext) {}

// EnterSetConstraints is called when production setConstraints is entered.
func (s *BaseOracleStatementListener) EnterSetConstraints(ctx *SetConstraintsContext) {}

// ExitSetConstraints is called when production setConstraints is exited.
func (s *BaseOracleStatementListener) ExitSetConstraints(ctx *SetConstraintsContext) {}

// EnterGrant is called when production grant is entered.
func (s *BaseOracleStatementListener) EnterGrant(ctx *GrantContext) {}

// ExitGrant is called when production grant is exited.
func (s *BaseOracleStatementListener) ExitGrant(ctx *GrantContext) {}

// EnterRevoke is called when production revoke is entered.
func (s *BaseOracleStatementListener) EnterRevoke(ctx *RevokeContext) {}

// ExitRevoke is called when production revoke is exited.
func (s *BaseOracleStatementListener) ExitRevoke(ctx *RevokeContext) {}

// EnterObjectPrivilegeClause is called when production objectPrivilegeClause is entered.
func (s *BaseOracleStatementListener) EnterObjectPrivilegeClause(ctx *ObjectPrivilegeClauseContext) {}

// ExitObjectPrivilegeClause is called when production objectPrivilegeClause is exited.
func (s *BaseOracleStatementListener) ExitObjectPrivilegeClause(ctx *ObjectPrivilegeClauseContext) {}

// EnterSystemPrivilegeClause is called when production systemPrivilegeClause is entered.
func (s *BaseOracleStatementListener) EnterSystemPrivilegeClause(ctx *SystemPrivilegeClauseContext) {}

// ExitSystemPrivilegeClause is called when production systemPrivilegeClause is exited.
func (s *BaseOracleStatementListener) ExitSystemPrivilegeClause(ctx *SystemPrivilegeClauseContext) {}

// EnterRoleClause is called when production roleClause is entered.
func (s *BaseOracleStatementListener) EnterRoleClause(ctx *RoleClauseContext) {}

// ExitRoleClause is called when production roleClause is exited.
func (s *BaseOracleStatementListener) ExitRoleClause(ctx *RoleClauseContext) {}

// EnterObjectPrivileges is called when production objectPrivileges is entered.
func (s *BaseOracleStatementListener) EnterObjectPrivileges(ctx *ObjectPrivilegesContext) {}

// ExitObjectPrivileges is called when production objectPrivileges is exited.
func (s *BaseOracleStatementListener) ExitObjectPrivileges(ctx *ObjectPrivilegesContext) {}

// EnterObjectPrivilegeType is called when production objectPrivilegeType is entered.
func (s *BaseOracleStatementListener) EnterObjectPrivilegeType(ctx *ObjectPrivilegeTypeContext) {}

// ExitObjectPrivilegeType is called when production objectPrivilegeType is exited.
func (s *BaseOracleStatementListener) ExitObjectPrivilegeType(ctx *ObjectPrivilegeTypeContext) {}

// EnterOnObjectClause is called when production onObjectClause is entered.
func (s *BaseOracleStatementListener) EnterOnObjectClause(ctx *OnObjectClauseContext) {}

// ExitOnObjectClause is called when production onObjectClause is exited.
func (s *BaseOracleStatementListener) ExitOnObjectClause(ctx *OnObjectClauseContext) {}

// EnterSystemPrivilege is called when production systemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterSystemPrivilege(ctx *SystemPrivilegeContext) {}

// ExitSystemPrivilege is called when production systemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitSystemPrivilege(ctx *SystemPrivilegeContext) {}

// EnterSystemPrivilegeOperation is called when production systemPrivilegeOperation is entered.
func (s *BaseOracleStatementListener) EnterSystemPrivilegeOperation(ctx *SystemPrivilegeOperationContext) {
}

// ExitSystemPrivilegeOperation is called when production systemPrivilegeOperation is exited.
func (s *BaseOracleStatementListener) ExitSystemPrivilegeOperation(ctx *SystemPrivilegeOperationContext) {
}

// EnterAdvisorFrameworkSystemPrivilege is called when production advisorFrameworkSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterAdvisorFrameworkSystemPrivilege(ctx *AdvisorFrameworkSystemPrivilegeContext) {
}

// ExitAdvisorFrameworkSystemPrivilege is called when production advisorFrameworkSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitAdvisorFrameworkSystemPrivilege(ctx *AdvisorFrameworkSystemPrivilegeContext) {
}

// EnterClustersSystemPrivilege is called when production clustersSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterClustersSystemPrivilege(ctx *ClustersSystemPrivilegeContext) {
}

// ExitClustersSystemPrivilege is called when production clustersSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitClustersSystemPrivilege(ctx *ClustersSystemPrivilegeContext) {
}

// EnterContextsSystemPrivilege is called when production contextsSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterContextsSystemPrivilege(ctx *ContextsSystemPrivilegeContext) {
}

// ExitContextsSystemPrivilege is called when production contextsSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitContextsSystemPrivilege(ctx *ContextsSystemPrivilegeContext) {
}

// EnterDataRedactionSystemPrivilege is called when production dataRedactionSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterDataRedactionSystemPrivilege(ctx *DataRedactionSystemPrivilegeContext) {
}

// ExitDataRedactionSystemPrivilege is called when production dataRedactionSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitDataRedactionSystemPrivilege(ctx *DataRedactionSystemPrivilegeContext) {
}

// EnterDatabaseSystemPrivilege is called when production databaseSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterDatabaseSystemPrivilege(ctx *DatabaseSystemPrivilegeContext) {
}

// ExitDatabaseSystemPrivilege is called when production databaseSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitDatabaseSystemPrivilege(ctx *DatabaseSystemPrivilegeContext) {
}

// EnterDatabaseLinksSystemPrivilege is called when production databaseLinksSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterDatabaseLinksSystemPrivilege(ctx *DatabaseLinksSystemPrivilegeContext) {
}

// ExitDatabaseLinksSystemPrivilege is called when production databaseLinksSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitDatabaseLinksSystemPrivilege(ctx *DatabaseLinksSystemPrivilegeContext) {
}

// EnterDebuggingSystemPrivilege is called when production debuggingSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterDebuggingSystemPrivilege(ctx *DebuggingSystemPrivilegeContext) {
}

// ExitDebuggingSystemPrivilege is called when production debuggingSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitDebuggingSystemPrivilege(ctx *DebuggingSystemPrivilegeContext) {
}

// EnterDictionariesSystemPrivilege is called when production dictionariesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterDictionariesSystemPrivilege(ctx *DictionariesSystemPrivilegeContext) {
}

// ExitDictionariesSystemPrivilege is called when production dictionariesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitDictionariesSystemPrivilege(ctx *DictionariesSystemPrivilegeContext) {
}

// EnterDimensionsSystemPrivilege is called when production dimensionsSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterDimensionsSystemPrivilege(ctx *DimensionsSystemPrivilegeContext) {
}

// ExitDimensionsSystemPrivilege is called when production dimensionsSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitDimensionsSystemPrivilege(ctx *DimensionsSystemPrivilegeContext) {
}

// EnterDirectoriesSystemPrivilege is called when production directoriesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterDirectoriesSystemPrivilege(ctx *DirectoriesSystemPrivilegeContext) {
}

// ExitDirectoriesSystemPrivilege is called when production directoriesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitDirectoriesSystemPrivilege(ctx *DirectoriesSystemPrivilegeContext) {
}

// EnterEditionsSystemPrivilege is called when production editionsSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterEditionsSystemPrivilege(ctx *EditionsSystemPrivilegeContext) {
}

// ExitEditionsSystemPrivilege is called when production editionsSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitEditionsSystemPrivilege(ctx *EditionsSystemPrivilegeContext) {
}

// EnterFlashbackDataArchivesPrivilege is called when production flashbackDataArchivesPrivilege is entered.
func (s *BaseOracleStatementListener) EnterFlashbackDataArchivesPrivilege(ctx *FlashbackDataArchivesPrivilegeContext) {
}

// ExitFlashbackDataArchivesPrivilege is called when production flashbackDataArchivesPrivilege is exited.
func (s *BaseOracleStatementListener) ExitFlashbackDataArchivesPrivilege(ctx *FlashbackDataArchivesPrivilegeContext) {
}

// EnterIndexesSystemPrivilege is called when production indexesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterIndexesSystemPrivilege(ctx *IndexesSystemPrivilegeContext) {
}

// ExitIndexesSystemPrivilege is called when production indexesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitIndexesSystemPrivilege(ctx *IndexesSystemPrivilegeContext) {
}

// EnterIndexTypesSystemPrivilege is called when production indexTypesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterIndexTypesSystemPrivilege(ctx *IndexTypesSystemPrivilegeContext) {
}

// ExitIndexTypesSystemPrivilege is called when production indexTypesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitIndexTypesSystemPrivilege(ctx *IndexTypesSystemPrivilegeContext) {
}

// EnterJobSchedulerObjectsSystemPrivilege is called when production jobSchedulerObjectsSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterJobSchedulerObjectsSystemPrivilege(ctx *JobSchedulerObjectsSystemPrivilegeContext) {
}

// ExitJobSchedulerObjectsSystemPrivilege is called when production jobSchedulerObjectsSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitJobSchedulerObjectsSystemPrivilege(ctx *JobSchedulerObjectsSystemPrivilegeContext) {
}

// EnterKeyManagementFrameworkSystemPrivilege is called when production keyManagementFrameworkSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterKeyManagementFrameworkSystemPrivilege(ctx *KeyManagementFrameworkSystemPrivilegeContext) {
}

// ExitKeyManagementFrameworkSystemPrivilege is called when production keyManagementFrameworkSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitKeyManagementFrameworkSystemPrivilege(ctx *KeyManagementFrameworkSystemPrivilegeContext) {
}

// EnterLibrariesFrameworkSystemPrivilege is called when production librariesFrameworkSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterLibrariesFrameworkSystemPrivilege(ctx *LibrariesFrameworkSystemPrivilegeContext) {
}

// ExitLibrariesFrameworkSystemPrivilege is called when production librariesFrameworkSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitLibrariesFrameworkSystemPrivilege(ctx *LibrariesFrameworkSystemPrivilegeContext) {
}

// EnterLogminerFrameworkSystemPrivilege is called when production logminerFrameworkSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterLogminerFrameworkSystemPrivilege(ctx *LogminerFrameworkSystemPrivilegeContext) {
}

// ExitLogminerFrameworkSystemPrivilege is called when production logminerFrameworkSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitLogminerFrameworkSystemPrivilege(ctx *LogminerFrameworkSystemPrivilegeContext) {
}

// EnterMaterizlizedViewsSystemPrivilege is called when production materizlizedViewsSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterMaterizlizedViewsSystemPrivilege(ctx *MaterizlizedViewsSystemPrivilegeContext) {
}

// ExitMaterizlizedViewsSystemPrivilege is called when production materizlizedViewsSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitMaterizlizedViewsSystemPrivilege(ctx *MaterizlizedViewsSystemPrivilegeContext) {
}

// EnterMiningModelsSystemPrivilege is called when production miningModelsSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterMiningModelsSystemPrivilege(ctx *MiningModelsSystemPrivilegeContext) {
}

// ExitMiningModelsSystemPrivilege is called when production miningModelsSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitMiningModelsSystemPrivilege(ctx *MiningModelsSystemPrivilegeContext) {
}

// EnterOlapCubesSystemPrivilege is called when production olapCubesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterOlapCubesSystemPrivilege(ctx *OlapCubesSystemPrivilegeContext) {
}

// ExitOlapCubesSystemPrivilege is called when production olapCubesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitOlapCubesSystemPrivilege(ctx *OlapCubesSystemPrivilegeContext) {
}

// EnterOlapCubeMeasureFoldersSystemPrivilege is called when production olapCubeMeasureFoldersSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterOlapCubeMeasureFoldersSystemPrivilege(ctx *OlapCubeMeasureFoldersSystemPrivilegeContext) {
}

// ExitOlapCubeMeasureFoldersSystemPrivilege is called when production olapCubeMeasureFoldersSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitOlapCubeMeasureFoldersSystemPrivilege(ctx *OlapCubeMeasureFoldersSystemPrivilegeContext) {
}

// EnterOlapCubeDiminsionsSystemPrivilege is called when production olapCubeDiminsionsSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterOlapCubeDiminsionsSystemPrivilege(ctx *OlapCubeDiminsionsSystemPrivilegeContext) {
}

// ExitOlapCubeDiminsionsSystemPrivilege is called when production olapCubeDiminsionsSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitOlapCubeDiminsionsSystemPrivilege(ctx *OlapCubeDiminsionsSystemPrivilegeContext) {
}

// EnterOlapCubeBuildProcessesSystemPrivilege is called when production olapCubeBuildProcessesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterOlapCubeBuildProcessesSystemPrivilege(ctx *OlapCubeBuildProcessesSystemPrivilegeContext) {
}

// ExitOlapCubeBuildProcessesSystemPrivilege is called when production olapCubeBuildProcessesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitOlapCubeBuildProcessesSystemPrivilege(ctx *OlapCubeBuildProcessesSystemPrivilegeContext) {
}

// EnterOperatorsSystemPrivilege is called when production operatorsSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterOperatorsSystemPrivilege(ctx *OperatorsSystemPrivilegeContext) {
}

// ExitOperatorsSystemPrivilege is called when production operatorsSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitOperatorsSystemPrivilege(ctx *OperatorsSystemPrivilegeContext) {
}

// EnterOutlinesSystemPrivilege is called when production outlinesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterOutlinesSystemPrivilege(ctx *OutlinesSystemPrivilegeContext) {
}

// ExitOutlinesSystemPrivilege is called when production outlinesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitOutlinesSystemPrivilege(ctx *OutlinesSystemPrivilegeContext) {
}

// EnterPlanManagementSystemPrivilege is called when production planManagementSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterPlanManagementSystemPrivilege(ctx *PlanManagementSystemPrivilegeContext) {
}

// ExitPlanManagementSystemPrivilege is called when production planManagementSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitPlanManagementSystemPrivilege(ctx *PlanManagementSystemPrivilegeContext) {
}

// EnterPluggableDatabasesSystemPrivilege is called when production pluggableDatabasesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterPluggableDatabasesSystemPrivilege(ctx *PluggableDatabasesSystemPrivilegeContext) {
}

// ExitPluggableDatabasesSystemPrivilege is called when production pluggableDatabasesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitPluggableDatabasesSystemPrivilege(ctx *PluggableDatabasesSystemPrivilegeContext) {
}

// EnterProceduresSystemPrivilege is called when production proceduresSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterProceduresSystemPrivilege(ctx *ProceduresSystemPrivilegeContext) {
}

// ExitProceduresSystemPrivilege is called when production proceduresSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitProceduresSystemPrivilege(ctx *ProceduresSystemPrivilegeContext) {
}

// EnterProfilesSystemPrivilege is called when production profilesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterProfilesSystemPrivilege(ctx *ProfilesSystemPrivilegeContext) {
}

// ExitProfilesSystemPrivilege is called when production profilesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitProfilesSystemPrivilege(ctx *ProfilesSystemPrivilegeContext) {
}

// EnterRolesSystemPrivilege is called when production rolesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterRolesSystemPrivilege(ctx *RolesSystemPrivilegeContext) {}

// ExitRolesSystemPrivilege is called when production rolesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitRolesSystemPrivilege(ctx *RolesSystemPrivilegeContext) {}

// EnterRollbackSegmentsSystemPrivilege is called when production rollbackSegmentsSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterRollbackSegmentsSystemPrivilege(ctx *RollbackSegmentsSystemPrivilegeContext) {
}

// ExitRollbackSegmentsSystemPrivilege is called when production rollbackSegmentsSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitRollbackSegmentsSystemPrivilege(ctx *RollbackSegmentsSystemPrivilegeContext) {
}

// EnterSequencesSystemPrivilege is called when production sequencesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterSequencesSystemPrivilege(ctx *SequencesSystemPrivilegeContext) {
}

// ExitSequencesSystemPrivilege is called when production sequencesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitSequencesSystemPrivilege(ctx *SequencesSystemPrivilegeContext) {
}

// EnterSessionsSystemPrivilege is called when production sessionsSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterSessionsSystemPrivilege(ctx *SessionsSystemPrivilegeContext) {
}

// ExitSessionsSystemPrivilege is called when production sessionsSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitSessionsSystemPrivilege(ctx *SessionsSystemPrivilegeContext) {
}

// EnterSqlTranslationProfilesSystemPrivilege is called when production sqlTranslationProfilesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterSqlTranslationProfilesSystemPrivilege(ctx *SqlTranslationProfilesSystemPrivilegeContext) {
}

// ExitSqlTranslationProfilesSystemPrivilege is called when production sqlTranslationProfilesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitSqlTranslationProfilesSystemPrivilege(ctx *SqlTranslationProfilesSystemPrivilegeContext) {
}

// EnterSynonymsSystemPrivilege is called when production synonymsSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterSynonymsSystemPrivilege(ctx *SynonymsSystemPrivilegeContext) {
}

// ExitSynonymsSystemPrivilege is called when production synonymsSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitSynonymsSystemPrivilege(ctx *SynonymsSystemPrivilegeContext) {
}

// EnterTablesSystemPrivilege is called when production tablesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterTablesSystemPrivilege(ctx *TablesSystemPrivilegeContext) {}

// ExitTablesSystemPrivilege is called when production tablesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitTablesSystemPrivilege(ctx *TablesSystemPrivilegeContext) {}

// EnterTablespacesSystemPrivilege is called when production tablespacesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterTablespacesSystemPrivilege(ctx *TablespacesSystemPrivilegeContext) {
}

// ExitTablespacesSystemPrivilege is called when production tablespacesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitTablespacesSystemPrivilege(ctx *TablespacesSystemPrivilegeContext) {
}

// EnterTriggersSystemPrivilege is called when production triggersSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterTriggersSystemPrivilege(ctx *TriggersSystemPrivilegeContext) {
}

// ExitTriggersSystemPrivilege is called when production triggersSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitTriggersSystemPrivilege(ctx *TriggersSystemPrivilegeContext) {
}

// EnterTypesSystemPrivilege is called when production typesSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterTypesSystemPrivilege(ctx *TypesSystemPrivilegeContext) {}

// ExitTypesSystemPrivilege is called when production typesSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitTypesSystemPrivilege(ctx *TypesSystemPrivilegeContext) {}

// EnterUsersSystemPrivilege is called when production usersSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterUsersSystemPrivilege(ctx *UsersSystemPrivilegeContext) {}

// ExitUsersSystemPrivilege is called when production usersSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitUsersSystemPrivilege(ctx *UsersSystemPrivilegeContext) {}

// EnterViewsSystemPrivilege is called when production viewsSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterViewsSystemPrivilege(ctx *ViewsSystemPrivilegeContext) {}

// ExitViewsSystemPrivilege is called when production viewsSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitViewsSystemPrivilege(ctx *ViewsSystemPrivilegeContext) {}

// EnterMiscellaneousSystemPrivilege is called when production miscellaneousSystemPrivilege is entered.
func (s *BaseOracleStatementListener) EnterMiscellaneousSystemPrivilege(ctx *MiscellaneousSystemPrivilegeContext) {
}

// ExitMiscellaneousSystemPrivilege is called when production miscellaneousSystemPrivilege is exited.
func (s *BaseOracleStatementListener) ExitMiscellaneousSystemPrivilege(ctx *MiscellaneousSystemPrivilegeContext) {
}

// EnterCreateUser is called when production createUser is entered.
func (s *BaseOracleStatementListener) EnterCreateUser(ctx *CreateUserContext) {}

// ExitCreateUser is called when production createUser is exited.
func (s *BaseOracleStatementListener) ExitCreateUser(ctx *CreateUserContext) {}

// EnterDropUser is called when production dropUser is entered.
func (s *BaseOracleStatementListener) EnterDropUser(ctx *DropUserContext) {}

// ExitDropUser is called when production dropUser is exited.
func (s *BaseOracleStatementListener) ExitDropUser(ctx *DropUserContext) {}

// EnterAlterUser is called when production alterUser is entered.
func (s *BaseOracleStatementListener) EnterAlterUser(ctx *AlterUserContext) {}

// ExitAlterUser is called when production alterUser is exited.
func (s *BaseOracleStatementListener) ExitAlterUser(ctx *AlterUserContext) {}

// EnterCreateRole is called when production createRole is entered.
func (s *BaseOracleStatementListener) EnterCreateRole(ctx *CreateRoleContext) {}

// ExitCreateRole is called when production createRole is exited.
func (s *BaseOracleStatementListener) ExitCreateRole(ctx *CreateRoleContext) {}

// EnterDropRole is called when production dropRole is entered.
func (s *BaseOracleStatementListener) EnterDropRole(ctx *DropRoleContext) {}

// ExitDropRole is called when production dropRole is exited.
func (s *BaseOracleStatementListener) ExitDropRole(ctx *DropRoleContext) {}

// EnterAlterRole is called when production alterRole is entered.
func (s *BaseOracleStatementListener) EnterAlterRole(ctx *AlterRoleContext) {}

// ExitAlterRole is called when production alterRole is exited.
func (s *BaseOracleStatementListener) ExitAlterRole(ctx *AlterRoleContext) {}

// EnterSetRole is called when production setRole is entered.
func (s *BaseOracleStatementListener) EnterSetRole(ctx *SetRoleContext) {}

// ExitSetRole is called when production setRole is exited.
func (s *BaseOracleStatementListener) ExitSetRole(ctx *SetRoleContext) {}

// EnterRoleAssignment is called when production roleAssignment is entered.
func (s *BaseOracleStatementListener) EnterRoleAssignment(ctx *RoleAssignmentContext) {}

// ExitRoleAssignment is called when production roleAssignment is exited.
func (s *BaseOracleStatementListener) ExitRoleAssignment(ctx *RoleAssignmentContext) {}

// EnterCall is called when production call is entered.
func (s *BaseOracleStatementListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseOracleStatementListener) ExitCall(ctx *CallContext) {}

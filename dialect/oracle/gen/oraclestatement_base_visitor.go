// Code generated from E:/GoProject/src/github.com/2212442929/xsqlparser/dialect/oracle/grammer\OracleStatement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // OracleStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseOracleStatementVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseOracleStatementVisitor) VisitExecute(ctx *ExecuteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInsert(ctx *InsertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInsertSingleTable(ctx *InsertSingleTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInsertMultiTable(ctx *InsertMultiTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMultiTableElement(ctx *MultiTableElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConditionalInsertClause(ctx *ConditionalInsertClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConditionalInsertWhenPart(ctx *ConditionalInsertWhenPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConditionalInsertElsePart(ctx *ConditionalInsertElsePartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInsertIntoClause(ctx *InsertIntoClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInsertValuesClause(ctx *InsertValuesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitReturningClause(ctx *ReturningClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDmlTableExprClause(ctx *DmlTableExprClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDmlTableClause(ctx *DmlTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartitionExtClause(ctx *PartitionExtClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDmlSubqueryClause(ctx *DmlSubqueryClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubqueryRestrictionClause(ctx *SubqueryRestrictionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTableCollectionExpr(ctx *TableCollectionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCollectionExpr(ctx *CollectionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUpdate(ctx *UpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUpdateSpecification(ctx *UpdateSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUpdateSetClause(ctx *UpdateSetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUpdateSetColumnList(ctx *UpdateSetColumnListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUpdateSetColumnClause(ctx *UpdateSetColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUpdateSetValueClause(ctx *UpdateSetValueClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAssignmentValues(ctx *AssignmentValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAssignmentValue(ctx *AssignmentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDelete(ctx *DeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDeleteSpecification(ctx *DeleteSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSelect(ctx *SelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSelectSubquery(ctx *SelectSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSelectUnionClause(ctx *SelectUnionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitParenthesisSelectSubquery(ctx *ParenthesisSelectSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQueryBlock(ctx *QueryBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPlsqlDeclarations(ctx *PlsqlDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFunctionHeading(ctx *FunctionHeadingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitProcedureDeclaration(ctx *ProcedureDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitProcedureHeading(ctx *ProcedureHeadingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitProcedureProperties(ctx *ProcedurePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAccessibleByClause(ctx *AccessibleByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAccessor(ctx *AccessorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUnitKind(ctx *UnitKindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDefaultCollationClause(ctx *DefaultCollationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCollationOption(ctx *CollationOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInvokerRightsClause(ctx *InvokerRightsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubqueryFactoringClause(ctx *SubqueryFactoringClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSearchClause(ctx *SearchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCycleClause(ctx *CycleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubavFactoringClause(ctx *SubavFactoringClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubavClause(ctx *SubavClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHierarchiesClause(ctx *HierarchiesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFilterClauses(ctx *FilterClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFilterClause(ctx *FilterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAddCalcsClause(ctx *AddCalcsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCalcMeasClause(ctx *CalcMeasClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCalcMeasExpression(ctx *CalcMeasExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAvExpression(ctx *AvExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAvMeasExpression(ctx *AvMeasExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLeadLagExpression(ctx *LeadLagExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLeadLagFunctionName(ctx *LeadLagFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLeadLagClause(ctx *LeadLagClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHierarchyRef(ctx *HierarchyRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitWindowExpression(ctx *WindowExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitWindowClause(ctx *WindowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPrecedingBoundary(ctx *PrecedingBoundaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFollowingBoundary(ctx *FollowingBoundaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRankExpression(ctx *RankExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRankFunctionName(ctx *RankFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRankClause(ctx *RankClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCalcMeasOrderByClause(ctx *CalcMeasOrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitShareOfExpression(ctx *ShareOfExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitShareClause(ctx *ShareClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMemberExpression(ctx *MemberExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLevelMemberLiteral(ctx *LevelMemberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPosMemberKeys(ctx *PosMemberKeysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNamedMemberKeys(ctx *NamedMemberKeysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHierNavigationExpression(ctx *HierNavigationExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHierAncestorExpression(ctx *HierAncestorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHierParentExpression(ctx *HierParentExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHierLeadLagExpression(ctx *HierLeadLagExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHierLeadLagClause(ctx *HierLeadLagClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQdrExpression(ctx *QdrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQualifier(ctx *QualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAvHierExpression(ctx *AvHierExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHierFunctionName(ctx *HierFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDuplicateSpecification(ctx *DuplicateSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSelectList(ctx *SelectListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSelectProjection(ctx *SelectProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSelectProjectionExprClause(ctx *SelectProjectionExprClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSelectFromClause(ctx *SelectFromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFromClauseList(ctx *FromClauseListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFromClauseOption(ctx *FromClauseOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSelectTableReference(ctx *SelectTableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQueryTableExprClause(ctx *QueryTableExprClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFlashbackQueryClause(ctx *FlashbackQueryClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQueryTableExpr(ctx *QueryTableExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLateralClause(ctx *LateralClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQueryTableExprSampleClause(ctx *QueryTableExprSampleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQueryTableExprTableClause(ctx *QueryTableExprTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQueryTableExprViewClause(ctx *QueryTableExprViewClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQueryTableExprAnalyticClause(ctx *QueryTableExprAnalyticClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInlineExternalTable(ctx *InlineExternalTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInlineExternalTableProperties(ctx *InlineExternalTablePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitExternalTableDataProperties(ctx *ExternalTableDataPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMofifiedExternalTable(ctx *MofifiedExternalTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModifyExternalTableProperties(ctx *ModifyExternalTablePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPivotClause(ctx *PivotClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPivotForClause(ctx *PivotForClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPivotInClause(ctx *PivotInClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUnpivotClause(ctx *UnpivotClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUnpivotInClause(ctx *UnpivotInClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSampleClause(ctx *SampleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitContainersClause(ctx *ContainersClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitShardsClause(ctx *ShardsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJoinClause(ctx *JoinClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSelectJoinOption(ctx *SelectJoinOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInnerCrossJoinClause(ctx *InnerCrossJoinClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSelectJoinSpecification(ctx *SelectJoinSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOuterJoinClause(ctx *OuterJoinClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQueryPartitionClause(ctx *QueryPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOuterJoinType(ctx *OuterJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCrossOuterApplyClause(ctx *CrossOuterApplyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInlineAnalyticView(ctx *InlineAnalyticViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHierarchicalQueryClause(ctx *HierarchicalQueryClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitGroupByItem(ctx *GroupByItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRollupCubeClause(ctx *RollupCubeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitGroupingSetsClause(ctx *GroupingSetsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitGroupingExprList(ctx *GroupingExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModelClause(ctx *ModelClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCellReferenceOptions(ctx *CellReferenceOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitReturnRowsClause(ctx *ReturnRowsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitReferenceModel(ctx *ReferenceModelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMainModel(ctx *MainModelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModelColumnClauses(ctx *ModelColumnClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModelRulesClause(ctx *ModelRulesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModelIterateClause(ctx *ModelIterateClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCellAssignment(ctx *CellAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSingleColumnForLoop(ctx *SingleColumnForLoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMultiColumnForLoop(ctx *MultiColumnForLoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModelExpr(ctx *ModelExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAnalyticFunction(ctx *AnalyticFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitForUpdateClause(ctx *ForUpdateClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitForUpdateClauseList(ctx *ForUpdateClauseListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitForUpdateClauseOption(ctx *ForUpdateClauseOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowLimitingClause(ctx *RowLimitingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMerge(ctx *MergeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHint(ctx *HintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIntoClause(ctx *IntoClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUsingClause(ctx *UsingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMergeUpdateClause(ctx *MergeUpdateClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMergeSetAssignmentsClause(ctx *MergeSetAssignmentsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMergeAssignment(ctx *MergeAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMergeAssignmentValue(ctx *MergeAssignmentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDeleteWhereClause(ctx *DeleteWhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMergeInsertClause(ctx *MergeInsertClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMergeInsertColumn(ctx *MergeInsertColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMergeColumnValue(ctx *MergeColumnValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitErrorLoggingClause(ctx *ErrorLoggingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternClause(ctx *RowPatternClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternPartitionBy(ctx *RowPatternPartitionByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternOrderBy(ctx *RowPatternOrderByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternMeasures(ctx *RowPatternMeasuresContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternMeasureColumn(ctx *RowPatternMeasureColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternRowsPerMatch(ctx *RowPatternRowsPerMatchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternSkipTo(ctx *RowPatternSkipToContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPattern(ctx *RowPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternTerm(ctx *RowPatternTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternFactor(ctx *RowPatternFactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternPrimary(ctx *RowPatternPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternPermute(ctx *RowPatternPermuteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternQuantifier(ctx *RowPatternQuantifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternSubsetClause(ctx *RowPatternSubsetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternSubsetItem(ctx *RowPatternSubsetItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternDefinitionList(ctx *RowPatternDefinitionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternDefinition(ctx *RowPatternDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternRecFunc(ctx *RowPatternRecFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPatternMeasExpression(ctx *PatternMeasExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternClassifierFunc(ctx *RowPatternClassifierFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternMatchNumFunc(ctx *RowPatternMatchNumFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternNavigationFunc(ctx *RowPatternNavigationFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternNavLogical(ctx *RowPatternNavLogicalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternNavPhysical(ctx *RowPatternNavPhysicalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternNavCompound(ctx *RowPatternNavCompoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowPatternAggregateFunc(ctx *RowPatternAggregateFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitParameterMarker(ctx *ParameterMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLiterals(ctx *LiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStringLiterals(ctx *StringLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNumberLiterals(ctx *NumberLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDateTimeLiterals(ctx *DateTimeLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHexadecimalLiterals(ctx *HexadecimalLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitBitValueLiterals(ctx *BitValueLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitBooleanLiterals(ctx *BooleanLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNullValueLiterals(ctx *NullValueLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUnreservedWord(ctx *UnreservedWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSchemaName(ctx *SchemaNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitViewName(ctx *ViewNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMaterializedViewName(ctx *MaterializedViewNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitObjectName(ctx *ObjectNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitClusterName(ctx *ClusterNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndexName(ctx *IndexNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStatisticsTypeName(ctx *StatisticsTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPackageName(ctx *PackageNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndextypeName(ctx *IndextypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModelName(ctx *ModelNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOperatorName(ctx *OperatorNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConstraintName(ctx *ConstraintNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSavepointName(ctx *SavepointNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSynonymName(ctx *SynonymNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOwner(ctx *OwnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTablespaceName(ctx *TablespaceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTablespaceSetName(ctx *TablespaceSetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitServiceName(ctx *ServiceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIlmPolicyName(ctx *IlmPolicyNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPolicyName(ctx *PolicyNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDbLink(ctx *DbLinkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitParameterValue(ctx *ParameterValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDirectoryName(ctx *DirectoryNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDispatcherName(ctx *DispatcherNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitClientId(ctx *ClientIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOpaqueFormatSpec(ctx *OpaqueFormatSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAccessDriverType(ctx *AccessDriverTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitVarrayItem(ctx *VarrayItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNestedItem(ctx *NestedItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStorageTable(ctx *StorageTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLobSegname(ctx *LobSegnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLocationSpecifier(ctx *LocationSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitXmlSchemaURLName(ctx *XmlSchemaURLNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitElementName(ctx *ElementNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubpartitionName(ctx *SubpartitionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitParameterName(ctx *ParameterNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEditionName(ctx *EditionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitContainerName(ctx *ContainerNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartitionName(ctx *PartitionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartitionSetName(ctx *PartitionSetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartitionKeyValue(ctx *PartitionKeyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubpartitionKeyValue(ctx *SubpartitionKeyValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitZonemapName(ctx *ZonemapNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFlashbackArchiveName(ctx *FlashbackArchiveNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRoleName(ctx *RoleNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUserName(ctx *UserNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPassword(ctx *PasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLogGroupName(ctx *LogGroupNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnNames(ctx *ColumnNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTableNames(ctx *TableNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOracleId(ctx *OracleIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCollationName(ctx *CollationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnCollationName(ctx *ColumnCollationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDataTypeLength(ctx *DataTypeLengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPrimaryKey(ctx *PrimaryKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitExprs(ctx *ExprsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAndOperator(ctx *AndOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOrOperator(ctx *OrOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNotOperator(ctx *NotOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitBooleanPrimary(ctx *BooleanPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitBitExpr(ctx *BitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSimpleExpr(ctx *SimpleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAggregationFunction(ctx *AggregationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAggregationFunctionName(ctx *AggregationFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAnalyticClause(ctx *AnalyticClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitWindowingClause(ctx *WindowingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSpecialFunction(ctx *SpecialFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCastFunction(ctx *CastFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCharFunction(ctx *CharFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRegularFunction(ctx *RegularFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRegularFunctionName(ctx *RegularFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCaseExpression(ctx *CaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCaseWhen(ctx *CaseWhenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCaseElse(ctx *CaseElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOrderByItem(ctx *OrderByItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAttributeName(ctx *AttributeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSimpleExprs(ctx *SimpleExprsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLobItem(ctx *LobItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLobItems(ctx *LobItemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLobItemList(ctx *LobItemListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSpecialDatatype(ctx *SpecialDatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDataTypeName(ctx *DataTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDatetimeTypeSuffix(ctx *DatetimeTypeSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTreatFunction(ctx *TreatFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPrivateExprOfDb(ctx *PrivateExprOfDbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCaseExpr(ctx *CaseExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSimpleCaseExpr(ctx *SimpleCaseExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSearchedCaseExpr(ctx *SearchedCaseExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitElseClause(ctx *ElseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIntervalExpression(ctx *IntervalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitObjectAccessExpression(ctx *ObjectAccessExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConstructorExpr(ctx *ConstructorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIgnoredIdentifier(ctx *IgnoredIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIgnoredIdentifiers(ctx *IgnoredIdentifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMatchNone(ctx *MatchNoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHashSubpartitionQuantity(ctx *HashSubpartitionQuantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOdciParameters(ctx *OdciParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDatabaseName(ctx *DatabaseNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLocationName(ctx *LocationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFileName(ctx *FileNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAsmFileName(ctx *AsmFileNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFileNumber(ctx *FileNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInstanceName(ctx *InstanceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLogminerSessionName(ctx *LogminerSessionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTablespaceGroupName(ctx *TablespaceGroupNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCopyName(ctx *CopyNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMirrorName(ctx *MirrorNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUriString(ctx *UriStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQualifiedCredentialName(ctx *QualifiedCredentialNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPdbName(ctx *PdbNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDiskgroupName(ctx *DiskgroupNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTemplateName(ctx *TemplateNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAliasName(ctx *AliasNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDomain(ctx *DomainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDateValue(ctx *DateValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSessionId(ctx *SessionIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSerialNumber(ctx *SerialNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInstanceId(ctx *InstanceIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSqlId(ctx *SqlIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLogFileName(ctx *LogFileNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLogFileGroupsArchivedLocationName(ctx *LogFileGroupsArchivedLocationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAsmVersion(ctx *AsmVersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitWalletPassword(ctx *WalletPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHsmAuthString(ctx *HsmAuthStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTargetDbName(ctx *TargetDbNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCertificateId(ctx *CertificateIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCategoryName(ctx *CategoryNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOffset(ctx *OffsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowcount(ctx *RowcountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPercent(ctx *PercentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRollbackSegment(ctx *RollbackSegmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQueryName(ctx *QueryNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCycleValue(ctx *CycleValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNoCycleValue(ctx *NoCycleValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOrderingColumn(ctx *OrderingColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubavName(ctx *SubavNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitBaseAvName(ctx *BaseAvNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMeasName(ctx *MeasNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLevelRef(ctx *LevelRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOffsetExpr(ctx *OffsetExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMemberKeyExpr(ctx *MemberKeyExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDepthExpression(ctx *DepthExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUnitName(ctx *UnitNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitProcedureName(ctx *ProcedureNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCpuCost(ctx *CpuCostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIoCost(ctx *IoCostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNetworkCost(ctx *NetworkCostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDefaultSelectivity(ctx *DefaultSelectivityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDataItem(ctx *DataItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitVariableName(ctx *VariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitValidTimeColumn(ctx *ValidTimeColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAttrDim(ctx *AttrDimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHierarchyName(ctx *HierarchyNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAnalyticViewName(ctx *AnalyticViewNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSamplePercent(ctx *SamplePercentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSeedValue(ctx *SeedValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNamespace(ctx *NamespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRestorePoint(ctx *RestorePointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitScnValue(ctx *ScnValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTimestampValue(ctx *TimestampValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitScnTimestampExpr(ctx *ScnTimestampExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitReferenceModelName(ctx *ReferenceModelNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMainModelName(ctx *MainModelNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMeasureColumn(ctx *MeasureColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDimensionColumn(ctx *DimensionColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAnalyticFunctionName(ctx *AnalyticFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitComparisonCondition(ctx *ComparisonConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSimpleComparisonCondition(ctx *SimpleComparisonConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitGroupComparisonCondition(ctx *GroupComparisonConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFloatingPointCondition(ctx *FloatingPointConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLogicalCondition(ctx *LogicalConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModelCondition(ctx *ModelConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIsAnyCondition(ctx *IsAnyConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIsPresentCondition(ctx *IsPresentConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCellReference(ctx *CellReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMultisetCondition(ctx *MultisetConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIsASetCondition(ctx *IsASetConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIsEmptyCondition(ctx *IsEmptyConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMemberCondition(ctx *MemberConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubmultisetCondition(ctx *SubmultisetConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPatternMatchingCondition(ctx *PatternMatchingConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLikeCondition(ctx *LikeConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSearchValue(ctx *SearchValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEscapeChar(ctx *EscapeCharContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRegexpLikeCondition(ctx *RegexpLikeConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMatchParam(ctx *MatchParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRangeCondition(ctx *RangeConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNullCondition(ctx *NullConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitXmlCondition(ctx *XmlConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEqualsPathCondition(ctx *EqualsPathConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPathString(ctx *PathStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCorrelationInteger(ctx *CorrelationIntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUnderPathCondition(ctx *UnderPathConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLevels(ctx *LevelsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonCondition(ctx *JsonConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIsJsonCondition(ctx *IsJsonConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonEqualCondition(ctx *JsonEqualConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonExistsCondition(ctx *JsonExistsConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonPassingClause(ctx *JsonPassingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonExistsOnErrorClause(ctx *JsonExistsOnErrorClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonExistsOnEmptyClause(ctx *JsonExistsOnEmptyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonTextcontainsCondition(ctx *JsonTextcontainsConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonBasicPathExpr(ctx *JsonBasicPathExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonAbsolutePathExpr(ctx *JsonAbsolutePathExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonNonfunctionSteps(ctx *JsonNonfunctionStepsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonObjectStep(ctx *JsonObjectStepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonFieldName(ctx *JsonFieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLetter(ctx *LetterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDigit(ctx *DigitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonArrayStep(ctx *JsonArrayStepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonDescendentStep(ctx *JsonDescendentStepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonFunctionStep(ctx *JsonFunctionStepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonItemMethod(ctx *JsonItemMethodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonFilterExpr(ctx *JsonFilterExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonCond(ctx *JsonCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonDisjunction(ctx *JsonDisjunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonConjunction(ctx *JsonConjunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonNegation(ctx *JsonNegationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonExistsCond(ctx *JsonExistsCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonHasSubstringCond(ctx *JsonHasSubstringCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonStartsWithCond(ctx *JsonStartsWithCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonLikeCond(ctx *JsonLikeCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonLikeRegexCond(ctx *JsonLikeRegexCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonEqRegexCond(ctx *JsonEqRegexCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonInCond(ctx *JsonInCondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitValueList(ctx *ValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonComparison(ctx *JsonComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonRelativePathExpr(ctx *JsonRelativePathExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonComparePred(ctx *JsonComparePredContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonVar(ctx *JsonVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonScalar(ctx *JsonScalarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonNumber(ctx *JsonNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJsonString(ctx *JsonStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCompoundCondition(ctx *CompoundConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitExistsCondition(ctx *ExistsConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInCondition(ctx *InConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIsOfTypeCondition(ctx *IsOfTypeConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDatabaseCharset(ctx *DatabaseCharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNationalCharset(ctx *NationalCharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFilenamePattern(ctx *FilenamePatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConnectString(ctx *ConnectStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateTable(ctx *CreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateIndex(ctx *CreateIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterTable(ctx *AlterTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterIndex(ctx *AlterIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDropTable(ctx *DropTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDropIndex(ctx *DropIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTruncateTable(ctx *TruncateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateTableSpecification(ctx *CreateTableSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTablespaceClauseWithParen(ctx *TablespaceClauseWithParenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTablespaceClause(ctx *TablespaceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateSharingClause(ctx *CreateSharingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateDefinitionClause(ctx *CreateDefinitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateXMLTypeTableClause(ctx *CreateXMLTypeTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitXmlTypeStorageClause(ctx *XmlTypeStorageClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitXmlSchemaSpecClause(ctx *XmlSchemaSpecClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitXmlTypeVirtualColumnsClause(ctx *XmlTypeVirtualColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOidClause(ctx *OidClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOidIndexClause(ctx *OidIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateRelationalTableClause(ctx *CreateRelationalTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateMemOptimizeClause(ctx *CreateMemOptimizeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateParentClause(ctx *CreateParentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateObjectTableClause(ctx *CreateObjectTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRelationalProperties(ctx *RelationalPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRelationalProperty(ctx *RelationalPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitVisibleClause(ctx *VisibleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDefaultNullClause(ctx *DefaultNullClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIdentityClause(ctx *IdentityClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIdentifyOptions(ctx *IdentifyOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIdentityOption(ctx *IdentityOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEncryptionSpecification(ctx *EncryptionSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInlineConstraint(ctx *InlineConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitReferencesClause(ctx *ReferencesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConstraintState(ctx *ConstraintStateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNotDeferrable(ctx *NotDeferrableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInitiallyClause(ctx *InitiallyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitExceptionsClause(ctx *ExceptionsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUsingIndexClause(ctx *UsingIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateIndexClause(ctx *CreateIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInlineRefConstraint(ctx *InlineRefConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitVirtualColumnDefinition(ctx *VirtualColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOutOfLineConstraint(ctx *OutOfLineConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOutOfLineRefConstraint(ctx *OutOfLineRefConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateIndexSpecification(ctx *CreateIndexSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitClusterIndexClause(ctx *ClusterIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndexAttributes(ctx *IndexAttributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTableIndexClause(ctx *TableIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndexExpressions(ctx *IndexExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndexExpression(ctx *IndexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitBitmapJoinIndexClause(ctx *BitmapJoinIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnSortsClause_(ctx *ColumnSortsClause_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnSortClause_(ctx *ColumnSortClause_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateIndexDefinitionClause(ctx *CreateIndexDefinitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTableAlias(ctx *TableAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterDefinitionClause(ctx *AlterDefinitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterTableProperties(ctx *AlterTablePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRenameTableSpecification(ctx *RenameTableSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnClauses(ctx *ColumnClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOperateColumnClause(ctx *OperateColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAddColumnSpecification(ctx *AddColumnSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnOrVirtualDefinitions(ctx *ColumnOrVirtualDefinitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnOrVirtualDefinition(ctx *ColumnOrVirtualDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnProperties(ctx *ColumnPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnProperty(ctx *ColumnPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitObjectTypeColProperties(ctx *ObjectTypeColPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubstitutableColumnClause(ctx *SubstitutableColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModifyColumnSpecification(ctx *ModifyColumnSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModifyColProperties(ctx *ModifyColPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModifyColSubstitutable(ctx *ModifyColSubstitutableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDropColumnClause(ctx *DropColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDropColumnSpecification(ctx *DropColumnSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnOrColumnList(ctx *ColumnOrColumnListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCascadeOrInvalidate(ctx *CascadeOrInvalidateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCheckpointNumber(ctx *CheckpointNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRenameColumnClause(ctx *RenameColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConstraintClauses(ctx *ConstraintClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAddConstraintSpecification(ctx *AddConstraintSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModifyConstraintClause(ctx *ModifyConstraintClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConstraintWithName(ctx *ConstraintWithNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConstraintOption(ctx *ConstraintOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConstraintPrimaryOrUnique(ctx *ConstraintPrimaryOrUniqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRenameConstraintClause(ctx *RenameConstraintClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDropConstraintClause(ctx *DropConstraintClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterExternalTable(ctx *AlterExternalTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitObjectProperties(ctx *ObjectPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterIndexInformationClause(ctx *AlterIndexInformationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRenameIndexClause(ctx *RenameIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitObjectTableSubstitution(ctx *ObjectTableSubstitutionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMemOptimizeClause(ctx *MemOptimizeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMemOptimizeReadClause(ctx *MemOptimizeReadClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMemOptimizeWriteClause(ctx *MemOptimizeWriteClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEnableDisableClauses(ctx *EnableDisableClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEnableDisableClause(ctx *EnableDisableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEnableDisableOthers(ctx *EnableDisableOthersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRebuildClause(ctx *RebuildClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitParallelClause(ctx *ParallelClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUsableSpecification(ctx *UsableSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInvalidationSpecification(ctx *InvalidationSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMaterializedViewLogClause(ctx *MaterializedViewLogClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDropReuseClause(ctx *DropReuseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCollationClause(ctx *CollationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCommitClause(ctx *CommitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPhysicalProperties(ctx *PhysicalPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDeferredSegmentCreation(ctx *DeferredSegmentCreationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSegmentAttributesClause(ctx *SegmentAttributesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPhysicalAttributesClause(ctx *PhysicalAttributesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLoggingClause(ctx *LoggingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStorageClause(ctx *StorageClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSizeClause(ctx *SizeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMaxsizeClause(ctx *MaxsizeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTableCompression(ctx *TableCompressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInmemoryTableClause(ctx *InmemoryTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInmemoryAttributes(ctx *InmemoryAttributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInmemoryColumnClause(ctx *InmemoryColumnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInmemoryMemcompress(ctx *InmemoryMemcompressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInmemoryPriority(ctx *InmemoryPriorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInmemoryDistribute(ctx *InmemoryDistributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInmemoryDuplicate(ctx *InmemoryDuplicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIlmClause(ctx *IlmClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIlmPolicyClause(ctx *IlmPolicyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIlmCompressionPolicy(ctx *IlmCompressionPolicyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIlmTimePeriod(ctx *IlmTimePeriodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIlmTieringPolicy(ctx *IlmTieringPolicyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIlmInmemoryPolicy(ctx *IlmInmemoryPolicyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOrganizationClause(ctx *OrganizationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHeapOrgTableClause(ctx *HeapOrgTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndexOrgTableClause(ctx *IndexOrgTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitExternalTableClause(ctx *ExternalTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitExternalTableDataProps(ctx *ExternalTableDataPropsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMappingTableClause(ctx *MappingTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPrefixCompression(ctx *PrefixCompressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndexOrgOverflowClause(ctx *IndexOrgOverflowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitExternalPartitionClause(ctx *ExternalPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitClusterRelatedClause(ctx *ClusterRelatedClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTableProperties(ctx *TablePropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitReadOnlyClause(ctx *ReadOnlyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndexingClause(ctx *IndexingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTablePartitioningClauses(ctx *TablePartitioningClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRangePartitions(ctx *RangePartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRangeValuesClause(ctx *RangeValuesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTablePartitionDescription(ctx *TablePartitionDescriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInmemoryClause(ctx *InmemoryClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitVarrayColProperties(ctx *VarrayColPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNestedTableColProperties(ctx *NestedTableColPropertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLobStorageClause(ctx *LobStorageClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitVarrayStorageClause(ctx *VarrayStorageClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLobStorageParameters(ctx *LobStorageParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLobParameters(ctx *LobParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLobRetentionClause(ctx *LobRetentionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLobDeduplicateClause(ctx *LobDeduplicateClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLobCompressionClause(ctx *LobCompressionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitExternalPartSubpartDataProps(ctx *ExternalPartSubpartDataPropsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitListPartitions(ctx *ListPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitListValuesClause(ctx *ListValuesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitListValues(ctx *ListValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHashPartitions(ctx *HashPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHashPartitionsByQuantity(ctx *HashPartitionsByQuantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndexCompression(ctx *IndexCompressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAdvancedIndexCompression(ctx *AdvancedIndexCompressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndividualHashPartitions(ctx *IndividualHashPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartitioningStorageClause(ctx *PartitioningStorageClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLobPartitioningStorage(ctx *LobPartitioningStorageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCompositeRangePartitions(ctx *CompositeRangePartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubpartitionByRange(ctx *SubpartitionByRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubpartitionByList(ctx *SubpartitionByListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubpartitionByHash(ctx *SubpartitionByHashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSubpartitionTemplate(ctx *SubpartitionTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRangeSubpartitionDesc(ctx *RangeSubpartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitListSubpartitionDesc(ctx *ListSubpartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndividualHashSubparts(ctx *IndividualHashSubpartsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRangePartitionDesc(ctx *RangePartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCompositeListPartitions(ctx *CompositeListPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitListPartitionDesc(ctx *ListPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCompositeHashPartitions(ctx *CompositeHashPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitReferencePartitioning(ctx *ReferencePartitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitReferencePartitionDesc(ctx *ReferencePartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConstraint(ctx *ConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSystemPartitioning(ctx *SystemPartitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConsistentHashPartitions(ctx *ConsistentHashPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConsistentHashWithSubpartitions(ctx *ConsistentHashWithSubpartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartitionsetClauses(ctx *PartitionsetClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRangePartitionsetClause(ctx *RangePartitionsetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRangePartitionsetDesc(ctx *RangePartitionsetDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitListPartitionsetClause(ctx *ListPartitionsetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAttributeClusteringClause(ctx *AttributeClusteringClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitClusteringJoin(ctx *ClusteringJoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitClusterClause(ctx *ClusterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitClusteringColumns(ctx *ClusteringColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitClusteringColumnGroup(ctx *ClusteringColumnGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitClusteringWhen(ctx *ClusteringWhenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitZonemapClause(ctx *ZonemapClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowMovementClause(ctx *RowMovementClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFlashbackArchiveClause(ctx *FlashbackArchiveClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterSynonym(ctx *AlterSynonymContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterTablePartitioning(ctx *AlterTablePartitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModifyTablePartition(ctx *ModifyTablePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModifyRangePartition(ctx *ModifyRangePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModifyHashPartition(ctx *ModifyHashPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitModifyListPartition(ctx *ModifyListPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartitionExtendedName(ctx *PartitionExtendedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAddRangeSubpartition(ctx *AddRangeSubpartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDependentTablesClause(ctx *DependentTablesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAddHashSubpartition(ctx *AddHashSubpartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAddListSubpartition(ctx *AddListSubpartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCoalesceTableSubpartition(ctx *CoalesceTableSubpartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAllowDisallowClustering(ctx *AllowDisallowClusteringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterMappingTableClauses(ctx *AlterMappingTableClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDeallocateUnusedClause(ctx *DeallocateUnusedClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAllocateExtentClause(ctx *AllocateExtentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartitionSpec(ctx *PartitionSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartitionAttributes(ctx *PartitionAttributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitShrinkClause(ctx *ShrinkClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMoveTablePartition(ctx *MoveTablePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFilterCondition(ctx *FilterConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCoalesceTablePartition(ctx *CoalesceTablePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAddTablePartition(ctx *AddTablePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAddRangePartitionClause(ctx *AddRangePartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAddListPartitionClause(ctx *AddListPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitHashSubpartsByQuantity(ctx *HashSubpartsByQuantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAddSystemPartitionClause(ctx *AddSystemPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAddHashPartitionClause(ctx *AddHashPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDropTablePartition(ctx *DropTablePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartitionExtendedNames(ctx *PartitionExtendedNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartitionForClauses(ctx *PartitionForClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUpdateIndexClauses(ctx *UpdateIndexClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUpdateGlobalIndexClause(ctx *UpdateGlobalIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUpdateAllIndexesClause(ctx *UpdateAllIndexesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUpdateIndexPartition(ctx *UpdateIndexPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndexPartitionDesc(ctx *IndexPartitionDescContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndexSubpartitionClause(ctx *IndexSubpartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUpdateIndexSubpartition(ctx *UpdateIndexSubpartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSupplementalLoggingProps(ctx *SupplementalLoggingPropsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSupplementalLogGrpClause(ctx *SupplementalLogGrpClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSupplementalIdKeyClause(ctx *SupplementalIdKeyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterSession(ctx *AlterSessionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterSessionOption(ctx *AlterSessionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAdviseClause(ctx *AdviseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCloseDatabaseLinkClause(ctx *CloseDatabaseLinkClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCommitInProcedureClause(ctx *CommitInProcedureClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSecuriyClause(ctx *SecuriyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitParallelExecutionClause(ctx *ParallelExecutionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitResumableClause(ctx *ResumableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEnableResumableClause(ctx *EnableResumableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDisableResumableClause(ctx *DisableResumableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitShardDdlClause(ctx *ShardDdlClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSyncWithPrimaryClause(ctx *SyncWithPrimaryClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterSessionSetClause(ctx *AlterSessionSetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterSessionSetClauseOption(ctx *AlterSessionSetClauseOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitParameterClause(ctx *ParameterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEditionClause(ctx *EditionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitContainerClause(ctx *ContainerClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRowArchivalVisibilityClause(ctx *RowArchivalVisibilityClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterDatabase(ctx *AlterDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDatabaseClauses(ctx *DatabaseClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStartupClauses(ctx *StartupClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRecoveryClauses(ctx *RecoveryClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitGeneralRecovery(ctx *GeneralRecoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFullDatabaseRecovery(ctx *FullDatabaseRecoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartialDatabaseRecovery(ctx *PartialDatabaseRecoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitManagedStandbyRecovery(ctx *ManagedStandbyRecoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDatabaseFileClauses(ctx *DatabaseFileClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateDatafileClause(ctx *CreateDatafileClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFileSpecifications(ctx *FileSpecificationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFileSpecification(ctx *FileSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDatafileTempfileSpec(ctx *DatafileTempfileSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAutoextendClause(ctx *AutoextendClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRedoLogFileSpec(ctx *RedoLogFileSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterDatafileClause(ctx *AlterDatafileClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterTempfileClause(ctx *AlterTempfileClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLogfileClauses(ctx *LogfileClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLogfileDescriptor(ctx *LogfileDescriptorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAddLogfileClauses(ctx *AddLogfileClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitControlfileClauses(ctx *ControlfileClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTraceFileClause(ctx *TraceFileClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDropLogfileClauses(ctx *DropLogfileClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSwitchLogfileClause(ctx *SwitchLogfileClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSupplementalDbLogging(ctx *SupplementalDbLoggingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSupplementalPlsqlClause(ctx *SupplementalPlsqlClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSupplementalSubsetReplicationClause(ctx *SupplementalSubsetReplicationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStandbyDatabaseClauses(ctx *StandbyDatabaseClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitActivateStandbyDbClause(ctx *ActivateStandbyDbClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMaximizeStandbyDbClause(ctx *MaximizeStandbyDbClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRegisterLogfileClause(ctx *RegisterLogfileClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCommitSwitchoverClause(ctx *CommitSwitchoverClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStartStandbyClause(ctx *StartStandbyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStopStandbyClause(ctx *StopStandbyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSwitchoverClause(ctx *SwitchoverClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConvertDatabaseClause(ctx *ConvertDatabaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFailoverClause(ctx *FailoverClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDefaultSettingsClauses(ctx *DefaultSettingsClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSetTimeZoneClause(ctx *SetTimeZoneClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTimeZoneRegion(ctx *TimeZoneRegionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFlashbackModeClause(ctx *FlashbackModeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUndoModeClause(ctx *UndoModeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMoveDatafileClause(ctx *MoveDatafileClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInstanceClauses(ctx *InstanceClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSecurityClause(ctx *SecurityClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPrepareClause(ctx *PrepareClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDropMirrorCopy(ctx *DropMirrorCopyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLostWriteProtection(ctx *LostWriteProtectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCdbFleetClauses(ctx *CdbFleetClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLeadCdbClause(ctx *LeadCdbClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLeadCdbUriClause(ctx *LeadCdbUriClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPropertyClause(ctx *PropertyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterSystem(ctx *AlterSystemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterSystemOption(ctx *AlterSystemOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitArchiveLogClause(ctx *ArchiveLogClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCheckpointClause(ctx *CheckpointClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCheckDatafilesClause(ctx *CheckDatafilesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDistributedRecovClauses(ctx *DistributedRecovClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFlushClause(ctx *FlushClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEndSessionClauses(ctx *EndSessionClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterSystemSwitchLogfileClause(ctx *AlterSystemSwitchLogfileClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSuspendResumeClause(ctx *SuspendResumeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitQuiesceClauses(ctx *QuiesceClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRollingMigrationClauses(ctx *RollingMigrationClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRollingPatchClauses(ctx *RollingPatchClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterSystemSecurityClauses(ctx *AlterSystemSecurityClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAffinityClauses(ctx *AffinityClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitShutdownDispatcherClause(ctx *ShutdownDispatcherClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRegisterClause(ctx *RegisterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSetClause(ctx *SetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitResetClause(ctx *ResetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRelocateClientClause(ctx *RelocateClientClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCancelSqlClause(ctx *CancelSqlClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFlushPasswordfileMetadataCacheClause(ctx *FlushPasswordfileMetadataCacheClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitInstanceClause(ctx *InstanceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSequenceClause(ctx *SequenceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitChangeClause(ctx *ChangeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCurrentClause(ctx *CurrentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitGroupClause(ctx *GroupClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLogfileClause(ctx *LogfileClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNextClause(ctx *NextClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAllClause(ctx *AllClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitToLocationClause(ctx *ToLocationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFlushClauseOption(ctx *FlushClauseOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDisconnectSessionClause(ctx *DisconnectSessionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitKillSessionClause(ctx *KillSessionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStartRollingMigrationClause(ctx *StartRollingMigrationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStopRollingMigrationClause(ctx *StopRollingMigrationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStartRollingPatchClause(ctx *StartRollingPatchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStopRollingPatchClause(ctx *StopRollingPatchClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRestrictedSessionClause(ctx *RestrictedSessionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSetEncryptionWalletOpenClause(ctx *SetEncryptionWalletOpenClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSetEncryptionWalletCloseClause(ctx *SetEncryptionWalletCloseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSetEncryptionKeyClause(ctx *SetEncryptionKeyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEnableAffinityClause(ctx *EnableAffinityClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDisableAffinityClause(ctx *DisableAffinityClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterSystemSetClause(ctx *AlterSystemSetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterSystemResetClause(ctx *AlterSystemResetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSharedPoolClause(ctx *SharedPoolClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitGlobalContextClause(ctx *GlobalContextClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitBufferCacheClause(ctx *BufferCacheClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFlashCacheClause(ctx *FlashCacheClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRedoToClause(ctx *RedoToClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIdentifiedByWalletPassword(ctx *IdentifiedByWalletPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIdentifiedByHsmAuthString(ctx *IdentifiedByHsmAuthStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSetParameterClause(ctx *SetParameterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUseStoredOutlinesClause(ctx *UseStoredOutlinesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitGlobalTopicEnabledClause(ctx *GlobalTopicEnabledClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterSystemCommentClause(ctx *AlterSystemCommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitContainerCurrentAllClause(ctx *ContainerCurrentAllClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitScopeClause(ctx *ScopeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAnalyze(ctx *AnalyzeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPartitionExtensionClause(ctx *PartitionExtensionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitValidationClauses(ctx *ValidationClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAssociateStatistics(ctx *AssociateStatisticsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitColumnAssociation(ctx *ColumnAssociationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFunctionAssociation(ctx *FunctionAssociationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitStorageTableClause(ctx *StorageTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUsingStatisticsType(ctx *UsingStatisticsTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDefaultCostClause(ctx *DefaultCostClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDefaultSelectivityClause(ctx *DefaultSelectivityClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDisassociateStatistics(ctx *DisassociateStatisticsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAudit(ctx *AuditContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNoAudit(ctx *NoAuditContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAuditPolicyClause(ctx *AuditPolicyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitNoAuditPolicyClause(ctx *NoAuditPolicyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitByUsersWithRoles(ctx *ByUsersWithRolesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitContextClause(ctx *ContextClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitContextNamespaceAttributesClause(ctx *ContextNamespaceAttributesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFlashbackDatabase(ctx *FlashbackDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitScnTimestampClause(ctx *ScnTimestampClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRestorePointClause(ctx *RestorePointClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFlashbackTable(ctx *FlashbackTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRenameToTable(ctx *RenameToTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPurge(ctx *PurgeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRename(ctx *RenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateDatabase(ctx *CreateDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateDatabaseClauses(ctx *CreateDatabaseClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDatabaseLoggingClauses(ctx *DatabaseLoggingClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTablespaceClauses(ctx *TablespaceClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDefaultTablespace(ctx *DefaultTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDefaultTempTablespace(ctx *DefaultTempTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUndoTablespace(ctx *UndoTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitBigOrSmallFiles(ctx *BigOrSmallFilesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitExtentManagementClause(ctx *ExtentManagementClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEnablePluggableDatabase(ctx *EnablePluggableDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFileNameConvert(ctx *FileNameConvertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitReplaceFileNamePattern(ctx *ReplaceFileNamePatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTablespaceDatafileClauses(ctx *TablespaceDatafileClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateDatabaseLink(ctx *CreateDatabaseLinkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitConnectToClause(ctx *ConnectToClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDbLinkAuthentication(ctx *DbLinkAuthenticationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSetTransaction(ctx *SetTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCommit(ctx *CommitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCommentClause(ctx *CommentClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitWriteClause(ctx *WriteClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitForceClause(ctx *ForceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRollback(ctx *RollbackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSavepointClause(ctx *SavepointClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSavepoint(ctx *SavepointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSetConstraints(ctx *SetConstraintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitGrant(ctx *GrantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRevoke(ctx *RevokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitObjectPrivilegeClause(ctx *ObjectPrivilegeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSystemPrivilegeClause(ctx *SystemPrivilegeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRoleClause(ctx *RoleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitObjectPrivileges(ctx *ObjectPrivilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitObjectPrivilegeType(ctx *ObjectPrivilegeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOnObjectClause(ctx *OnObjectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSystemPrivilege(ctx *SystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSystemPrivilegeOperation(ctx *SystemPrivilegeOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAdvisorFrameworkSystemPrivilege(ctx *AdvisorFrameworkSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitClustersSystemPrivilege(ctx *ClustersSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitContextsSystemPrivilege(ctx *ContextsSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDataRedactionSystemPrivilege(ctx *DataRedactionSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDatabaseSystemPrivilege(ctx *DatabaseSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDatabaseLinksSystemPrivilege(ctx *DatabaseLinksSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDebuggingSystemPrivilege(ctx *DebuggingSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDictionariesSystemPrivilege(ctx *DictionariesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDimensionsSystemPrivilege(ctx *DimensionsSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDirectoriesSystemPrivilege(ctx *DirectoriesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitEditionsSystemPrivilege(ctx *EditionsSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitFlashbackDataArchivesPrivilege(ctx *FlashbackDataArchivesPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndexesSystemPrivilege(ctx *IndexesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitIndexTypesSystemPrivilege(ctx *IndexTypesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitJobSchedulerObjectsSystemPrivilege(ctx *JobSchedulerObjectsSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitKeyManagementFrameworkSystemPrivilege(ctx *KeyManagementFrameworkSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLibrariesFrameworkSystemPrivilege(ctx *LibrariesFrameworkSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitLogminerFrameworkSystemPrivilege(ctx *LogminerFrameworkSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMaterizlizedViewsSystemPrivilege(ctx *MaterizlizedViewsSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMiningModelsSystemPrivilege(ctx *MiningModelsSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOlapCubesSystemPrivilege(ctx *OlapCubesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOlapCubeMeasureFoldersSystemPrivilege(ctx *OlapCubeMeasureFoldersSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOlapCubeDiminsionsSystemPrivilege(ctx *OlapCubeDiminsionsSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOlapCubeBuildProcessesSystemPrivilege(ctx *OlapCubeBuildProcessesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOperatorsSystemPrivilege(ctx *OperatorsSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitOutlinesSystemPrivilege(ctx *OutlinesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPlanManagementSystemPrivilege(ctx *PlanManagementSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitPluggableDatabasesSystemPrivilege(ctx *PluggableDatabasesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitProceduresSystemPrivilege(ctx *ProceduresSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitProfilesSystemPrivilege(ctx *ProfilesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRolesSystemPrivilege(ctx *RolesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRollbackSegmentsSystemPrivilege(ctx *RollbackSegmentsSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSequencesSystemPrivilege(ctx *SequencesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSessionsSystemPrivilege(ctx *SessionsSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSqlTranslationProfilesSystemPrivilege(ctx *SqlTranslationProfilesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSynonymsSystemPrivilege(ctx *SynonymsSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTablesSystemPrivilege(ctx *TablesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTablespacesSystemPrivilege(ctx *TablespacesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTriggersSystemPrivilege(ctx *TriggersSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitTypesSystemPrivilege(ctx *TypesSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitUsersSystemPrivilege(ctx *UsersSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitViewsSystemPrivilege(ctx *ViewsSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitMiscellaneousSystemPrivilege(ctx *MiscellaneousSystemPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateUser(ctx *CreateUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDropUser(ctx *DropUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterUser(ctx *AlterUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCreateRole(ctx *CreateRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitDropRole(ctx *DropRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitAlterRole(ctx *AlterRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitSetRole(ctx *SetRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitRoleAssignment(ctx *RoleAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOracleStatementVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

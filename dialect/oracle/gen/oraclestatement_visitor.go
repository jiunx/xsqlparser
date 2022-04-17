// Code generated from E:/GoProject/src/github.com/2212442929/xsqlparser/dialect/oracle/grammer\OracleStatement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // OracleStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by OracleStatementParser.
type OracleStatementVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by OracleStatementParser#execute.
	VisitExecute(ctx *ExecuteContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#insert.
	VisitInsert(ctx *InsertContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#insertSingleTable.
	VisitInsertSingleTable(ctx *InsertSingleTableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#insertMultiTable.
	VisitInsertMultiTable(ctx *InsertMultiTableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#multiTableElement.
	VisitMultiTableElement(ctx *MultiTableElementContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#conditionalInsertClause.
	VisitConditionalInsertClause(ctx *ConditionalInsertClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#conditionalInsertWhenPart.
	VisitConditionalInsertWhenPart(ctx *ConditionalInsertWhenPartContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#conditionalInsertElsePart.
	VisitConditionalInsertElsePart(ctx *ConditionalInsertElsePartContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#insertIntoClause.
	VisitInsertIntoClause(ctx *InsertIntoClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#insertValuesClause.
	VisitInsertValuesClause(ctx *InsertValuesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#returningClause.
	VisitReturningClause(ctx *ReturningClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dmlTableExprClause.
	VisitDmlTableExprClause(ctx *DmlTableExprClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dmlTableClause.
	VisitDmlTableClause(ctx *DmlTableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partitionExtClause.
	VisitPartitionExtClause(ctx *PartitionExtClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dmlSubqueryClause.
	VisitDmlSubqueryClause(ctx *DmlSubqueryClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#subqueryRestrictionClause.
	VisitSubqueryRestrictionClause(ctx *SubqueryRestrictionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tableCollectionExpr.
	VisitTableCollectionExpr(ctx *TableCollectionExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#collectionExpr.
	VisitCollectionExpr(ctx *CollectionExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#update.
	VisitUpdate(ctx *UpdateContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#updateSpecification.
	VisitUpdateSpecification(ctx *UpdateSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#updateSetClause.
	VisitUpdateSetClause(ctx *UpdateSetClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#updateSetColumnList.
	VisitUpdateSetColumnList(ctx *UpdateSetColumnListContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#updateSetColumnClause.
	VisitUpdateSetColumnClause(ctx *UpdateSetColumnClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#updateSetValueClause.
	VisitUpdateSetValueClause(ctx *UpdateSetValueClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#assignmentValues.
	VisitAssignmentValues(ctx *AssignmentValuesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#assignmentValue.
	VisitAssignmentValue(ctx *AssignmentValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#delete.
	VisitDelete(ctx *DeleteContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#deleteSpecification.
	VisitDeleteSpecification(ctx *DeleteSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#select.
	VisitSelect(ctx *SelectContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#selectSubquery.
	VisitSelectSubquery(ctx *SelectSubqueryContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#selectUnionClause.
	VisitSelectUnionClause(ctx *SelectUnionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#parenthesisSelectSubquery.
	VisitParenthesisSelectSubquery(ctx *ParenthesisSelectSubqueryContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#queryBlock.
	VisitQueryBlock(ctx *QueryBlockContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#plsqlDeclarations.
	VisitPlsqlDeclarations(ctx *PlsqlDeclarationsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#functionHeading.
	VisitFunctionHeading(ctx *FunctionHeadingContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#parameterDeclaration.
	VisitParameterDeclaration(ctx *ParameterDeclarationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#procedureDeclaration.
	VisitProcedureDeclaration(ctx *ProcedureDeclarationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#procedureHeading.
	VisitProcedureHeading(ctx *ProcedureHeadingContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#procedureProperties.
	VisitProcedureProperties(ctx *ProcedurePropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#accessibleByClause.
	VisitAccessibleByClause(ctx *AccessibleByClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#accessor.
	VisitAccessor(ctx *AccessorContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#unitKind.
	VisitUnitKind(ctx *UnitKindContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#defaultCollationClause.
	VisitDefaultCollationClause(ctx *DefaultCollationClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#collationOption.
	VisitCollationOption(ctx *CollationOptionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#invokerRightsClause.
	VisitInvokerRightsClause(ctx *InvokerRightsClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#subqueryFactoringClause.
	VisitSubqueryFactoringClause(ctx *SubqueryFactoringClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#searchClause.
	VisitSearchClause(ctx *SearchClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#cycleClause.
	VisitCycleClause(ctx *CycleClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#subavFactoringClause.
	VisitSubavFactoringClause(ctx *SubavFactoringClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#subavClause.
	VisitSubavClause(ctx *SubavClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hierarchiesClause.
	VisitHierarchiesClause(ctx *HierarchiesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#filterClauses.
	VisitFilterClauses(ctx *FilterClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#filterClause.
	VisitFilterClause(ctx *FilterClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#addCalcsClause.
	VisitAddCalcsClause(ctx *AddCalcsClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#calcMeasClause.
	VisitCalcMeasClause(ctx *CalcMeasClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#calcMeasExpression.
	VisitCalcMeasExpression(ctx *CalcMeasExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#avExpression.
	VisitAvExpression(ctx *AvExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#avMeasExpression.
	VisitAvMeasExpression(ctx *AvMeasExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#leadLagExpression.
	VisitLeadLagExpression(ctx *LeadLagExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#leadLagFunctionName.
	VisitLeadLagFunctionName(ctx *LeadLagFunctionNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#leadLagClause.
	VisitLeadLagClause(ctx *LeadLagClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hierarchyRef.
	VisitHierarchyRef(ctx *HierarchyRefContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#windowExpression.
	VisitWindowExpression(ctx *WindowExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#windowClause.
	VisitWindowClause(ctx *WindowClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#precedingBoundary.
	VisitPrecedingBoundary(ctx *PrecedingBoundaryContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#followingBoundary.
	VisitFollowingBoundary(ctx *FollowingBoundaryContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rankExpression.
	VisitRankExpression(ctx *RankExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rankFunctionName.
	VisitRankFunctionName(ctx *RankFunctionNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rankClause.
	VisitRankClause(ctx *RankClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#calcMeasOrderByClause.
	VisitCalcMeasOrderByClause(ctx *CalcMeasOrderByClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#shareOfExpression.
	VisitShareOfExpression(ctx *ShareOfExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#shareClause.
	VisitShareClause(ctx *ShareClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#memberExpression.
	VisitMemberExpression(ctx *MemberExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#levelMemberLiteral.
	VisitLevelMemberLiteral(ctx *LevelMemberLiteralContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#posMemberKeys.
	VisitPosMemberKeys(ctx *PosMemberKeysContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#namedMemberKeys.
	VisitNamedMemberKeys(ctx *NamedMemberKeysContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hierNavigationExpression.
	VisitHierNavigationExpression(ctx *HierNavigationExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hierAncestorExpression.
	VisitHierAncestorExpression(ctx *HierAncestorExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hierParentExpression.
	VisitHierParentExpression(ctx *HierParentExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hierLeadLagExpression.
	VisitHierLeadLagExpression(ctx *HierLeadLagExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hierLeadLagClause.
	VisitHierLeadLagClause(ctx *HierLeadLagClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#qdrExpression.
	VisitQdrExpression(ctx *QdrExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#qualifier.
	VisitQualifier(ctx *QualifierContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#avHierExpression.
	VisitAvHierExpression(ctx *AvHierExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hierFunctionName.
	VisitHierFunctionName(ctx *HierFunctionNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#duplicateSpecification.
	VisitDuplicateSpecification(ctx *DuplicateSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#unqualifiedShorthand.
	VisitUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#selectList.
	VisitSelectList(ctx *SelectListContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#selectProjection.
	VisitSelectProjection(ctx *SelectProjectionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#selectProjectionExprClause.
	VisitSelectProjectionExprClause(ctx *SelectProjectionExprClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#selectFromClause.
	VisitSelectFromClause(ctx *SelectFromClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#fromClauseList.
	VisitFromClauseList(ctx *FromClauseListContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#fromClauseOption.
	VisitFromClauseOption(ctx *FromClauseOptionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#selectTableReference.
	VisitSelectTableReference(ctx *SelectTableReferenceContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#queryTableExprClause.
	VisitQueryTableExprClause(ctx *QueryTableExprClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#flashbackQueryClause.
	VisitFlashbackQueryClause(ctx *FlashbackQueryClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#queryTableExpr.
	VisitQueryTableExpr(ctx *QueryTableExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lateralClause.
	VisitLateralClause(ctx *LateralClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#queryTableExprSampleClause.
	VisitQueryTableExprSampleClause(ctx *QueryTableExprSampleClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#queryTableExprTableClause.
	VisitQueryTableExprTableClause(ctx *QueryTableExprTableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#queryTableExprViewClause.
	VisitQueryTableExprViewClause(ctx *QueryTableExprViewClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#queryTableExprAnalyticClause.
	VisitQueryTableExprAnalyticClause(ctx *QueryTableExprAnalyticClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inlineExternalTable.
	VisitInlineExternalTable(ctx *InlineExternalTableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inlineExternalTableProperties.
	VisitInlineExternalTableProperties(ctx *InlineExternalTablePropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#externalTableDataProperties.
	VisitExternalTableDataProperties(ctx *ExternalTableDataPropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#mofifiedExternalTable.
	VisitMofifiedExternalTable(ctx *MofifiedExternalTableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modifyExternalTableProperties.
	VisitModifyExternalTableProperties(ctx *ModifyExternalTablePropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#pivotClause.
	VisitPivotClause(ctx *PivotClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#pivotForClause.
	VisitPivotForClause(ctx *PivotForClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#pivotInClause.
	VisitPivotInClause(ctx *PivotInClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#unpivotClause.
	VisitUnpivotClause(ctx *UnpivotClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#unpivotInClause.
	VisitUnpivotInClause(ctx *UnpivotInClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#sampleClause.
	VisitSampleClause(ctx *SampleClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#containersClause.
	VisitContainersClause(ctx *ContainersClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#shardsClause.
	VisitShardsClause(ctx *ShardsClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#joinClause.
	VisitJoinClause(ctx *JoinClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#selectJoinOption.
	VisitSelectJoinOption(ctx *SelectJoinOptionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#innerCrossJoinClause.
	VisitInnerCrossJoinClause(ctx *InnerCrossJoinClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#selectJoinSpecification.
	VisitSelectJoinSpecification(ctx *SelectJoinSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#outerJoinClause.
	VisitOuterJoinClause(ctx *OuterJoinClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#queryPartitionClause.
	VisitQueryPartitionClause(ctx *QueryPartitionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#outerJoinType.
	VisitOuterJoinType(ctx *OuterJoinTypeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#crossOuterApplyClause.
	VisitCrossOuterApplyClause(ctx *CrossOuterApplyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inlineAnalyticView.
	VisitInlineAnalyticView(ctx *InlineAnalyticViewContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hierarchicalQueryClause.
	VisitHierarchicalQueryClause(ctx *HierarchicalQueryClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#groupByItem.
	VisitGroupByItem(ctx *GroupByItemContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rollupCubeClause.
	VisitRollupCubeClause(ctx *RollupCubeClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#groupingSetsClause.
	VisitGroupingSetsClause(ctx *GroupingSetsClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#groupingExprList.
	VisitGroupingExprList(ctx *GroupingExprListContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modelClause.
	VisitModelClause(ctx *ModelClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#cellReferenceOptions.
	VisitCellReferenceOptions(ctx *CellReferenceOptionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#returnRowsClause.
	VisitReturnRowsClause(ctx *ReturnRowsClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#referenceModel.
	VisitReferenceModel(ctx *ReferenceModelContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#mainModel.
	VisitMainModel(ctx *MainModelContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modelColumnClauses.
	VisitModelColumnClauses(ctx *ModelColumnClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modelRulesClause.
	VisitModelRulesClause(ctx *ModelRulesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modelIterateClause.
	VisitModelIterateClause(ctx *ModelIterateClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#cellAssignment.
	VisitCellAssignment(ctx *CellAssignmentContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#singleColumnForLoop.
	VisitSingleColumnForLoop(ctx *SingleColumnForLoopContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#multiColumnForLoop.
	VisitMultiColumnForLoop(ctx *MultiColumnForLoopContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modelExpr.
	VisitModelExpr(ctx *ModelExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#analyticFunction.
	VisitAnalyticFunction(ctx *AnalyticFunctionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#forUpdateClause.
	VisitForUpdateClause(ctx *ForUpdateClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#forUpdateClauseList.
	VisitForUpdateClauseList(ctx *ForUpdateClauseListContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#forUpdateClauseOption.
	VisitForUpdateClauseOption(ctx *ForUpdateClauseOptionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowLimitingClause.
	VisitRowLimitingClause(ctx *RowLimitingClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#merge.
	VisitMerge(ctx *MergeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hint.
	VisitHint(ctx *HintContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#intoClause.
	VisitIntoClause(ctx *IntoClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#usingClause.
	VisitUsingClause(ctx *UsingClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#mergeUpdateClause.
	VisitMergeUpdateClause(ctx *MergeUpdateClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#mergeSetAssignmentsClause.
	VisitMergeSetAssignmentsClause(ctx *MergeSetAssignmentsClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#mergeAssignment.
	VisitMergeAssignment(ctx *MergeAssignmentContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#mergeAssignmentValue.
	VisitMergeAssignmentValue(ctx *MergeAssignmentValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#deleteWhereClause.
	VisitDeleteWhereClause(ctx *DeleteWhereClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#mergeInsertClause.
	VisitMergeInsertClause(ctx *MergeInsertClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#mergeInsertColumn.
	VisitMergeInsertColumn(ctx *MergeInsertColumnContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#mergeColumnValue.
	VisitMergeColumnValue(ctx *MergeColumnValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#errorLoggingClause.
	VisitErrorLoggingClause(ctx *ErrorLoggingClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternClause.
	VisitRowPatternClause(ctx *RowPatternClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternPartitionBy.
	VisitRowPatternPartitionBy(ctx *RowPatternPartitionByContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternOrderBy.
	VisitRowPatternOrderBy(ctx *RowPatternOrderByContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternMeasures.
	VisitRowPatternMeasures(ctx *RowPatternMeasuresContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternMeasureColumn.
	VisitRowPatternMeasureColumn(ctx *RowPatternMeasureColumnContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternRowsPerMatch.
	VisitRowPatternRowsPerMatch(ctx *RowPatternRowsPerMatchContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternSkipTo.
	VisitRowPatternSkipTo(ctx *RowPatternSkipToContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPattern.
	VisitRowPattern(ctx *RowPatternContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternTerm.
	VisitRowPatternTerm(ctx *RowPatternTermContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternFactor.
	VisitRowPatternFactor(ctx *RowPatternFactorContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternPrimary.
	VisitRowPatternPrimary(ctx *RowPatternPrimaryContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternPermute.
	VisitRowPatternPermute(ctx *RowPatternPermuteContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternQuantifier.
	VisitRowPatternQuantifier(ctx *RowPatternQuantifierContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternSubsetClause.
	VisitRowPatternSubsetClause(ctx *RowPatternSubsetClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternSubsetItem.
	VisitRowPatternSubsetItem(ctx *RowPatternSubsetItemContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternDefinitionList.
	VisitRowPatternDefinitionList(ctx *RowPatternDefinitionListContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternDefinition.
	VisitRowPatternDefinition(ctx *RowPatternDefinitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternRecFunc.
	VisitRowPatternRecFunc(ctx *RowPatternRecFuncContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#patternMeasExpression.
	VisitPatternMeasExpression(ctx *PatternMeasExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternClassifierFunc.
	VisitRowPatternClassifierFunc(ctx *RowPatternClassifierFuncContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternMatchNumFunc.
	VisitRowPatternMatchNumFunc(ctx *RowPatternMatchNumFuncContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternNavigationFunc.
	VisitRowPatternNavigationFunc(ctx *RowPatternNavigationFuncContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternNavLogical.
	VisitRowPatternNavLogical(ctx *RowPatternNavLogicalContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternNavPhysical.
	VisitRowPatternNavPhysical(ctx *RowPatternNavPhysicalContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternNavCompound.
	VisitRowPatternNavCompound(ctx *RowPatternNavCompoundContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowPatternAggregateFunc.
	VisitRowPatternAggregateFunc(ctx *RowPatternAggregateFuncContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#parameterMarker.
	VisitParameterMarker(ctx *ParameterMarkerContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#literals.
	VisitLiterals(ctx *LiteralsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#stringLiterals.
	VisitStringLiterals(ctx *StringLiteralsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#numberLiterals.
	VisitNumberLiterals(ctx *NumberLiteralsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dateTimeLiterals.
	VisitDateTimeLiterals(ctx *DateTimeLiteralsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hexadecimalLiterals.
	VisitHexadecimalLiterals(ctx *HexadecimalLiteralsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#bitValueLiterals.
	VisitBitValueLiterals(ctx *BitValueLiteralsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#booleanLiterals.
	VisitBooleanLiterals(ctx *BooleanLiteralsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#nullValueLiterals.
	VisitNullValueLiterals(ctx *NullValueLiteralsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#unreservedWord.
	VisitUnreservedWord(ctx *UnreservedWordContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#schemaName.
	VisitSchemaName(ctx *SchemaNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#viewName.
	VisitViewName(ctx *ViewNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#materializedViewName.
	VisitMaterializedViewName(ctx *MaterializedViewNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#objectName.
	VisitObjectName(ctx *ObjectNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#clusterName.
	VisitClusterName(ctx *ClusterNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indexName.
	VisitIndexName(ctx *IndexNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#statisticsTypeName.
	VisitStatisticsTypeName(ctx *StatisticsTypeNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#packageName.
	VisitPackageName(ctx *PackageNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indextypeName.
	VisitIndextypeName(ctx *IndextypeNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modelName.
	VisitModelName(ctx *ModelNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#operatorName.
	VisitOperatorName(ctx *OperatorNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#constraintName.
	VisitConstraintName(ctx *ConstraintNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#savepointName.
	VisitSavepointName(ctx *SavepointNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#synonymName.
	VisitSynonymName(ctx *SynonymNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#owner.
	VisitOwner(ctx *OwnerContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tablespaceName.
	VisitTablespaceName(ctx *TablespaceNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tablespaceSetName.
	VisitTablespaceSetName(ctx *TablespaceSetNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#serviceName.
	VisitServiceName(ctx *ServiceNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#ilmPolicyName.
	VisitIlmPolicyName(ctx *IlmPolicyNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#policyName.
	VisitPolicyName(ctx *PolicyNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dbLink.
	VisitDbLink(ctx *DbLinkContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#parameterValue.
	VisitParameterValue(ctx *ParameterValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#directoryName.
	VisitDirectoryName(ctx *DirectoryNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dispatcherName.
	VisitDispatcherName(ctx *DispatcherNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#clientId.
	VisitClientId(ctx *ClientIdContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#opaqueFormatSpec.
	VisitOpaqueFormatSpec(ctx *OpaqueFormatSpecContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#accessDriverType.
	VisitAccessDriverType(ctx *AccessDriverTypeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#varrayItem.
	VisitVarrayItem(ctx *VarrayItemContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#nestedItem.
	VisitNestedItem(ctx *NestedItemContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#storageTable.
	VisitStorageTable(ctx *StorageTableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lobSegname.
	VisitLobSegname(ctx *LobSegnameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#locationSpecifier.
	VisitLocationSpecifier(ctx *LocationSpecifierContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#xmlSchemaURLName.
	VisitXmlSchemaURLName(ctx *XmlSchemaURLNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#elementName.
	VisitElementName(ctx *ElementNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#subpartitionName.
	VisitSubpartitionName(ctx *SubpartitionNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#parameterName.
	VisitParameterName(ctx *ParameterNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#editionName.
	VisitEditionName(ctx *EditionNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#containerName.
	VisitContainerName(ctx *ContainerNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partitionName.
	VisitPartitionName(ctx *PartitionNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partitionSetName.
	VisitPartitionSetName(ctx *PartitionSetNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partitionKeyValue.
	VisitPartitionKeyValue(ctx *PartitionKeyValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#subpartitionKeyValue.
	VisitSubpartitionKeyValue(ctx *SubpartitionKeyValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#zonemapName.
	VisitZonemapName(ctx *ZonemapNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#flashbackArchiveName.
	VisitFlashbackArchiveName(ctx *FlashbackArchiveNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#roleName.
	VisitRoleName(ctx *RoleNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#userName.
	VisitUserName(ctx *UserNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#password.
	VisitPassword(ctx *PasswordContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#logGroupName.
	VisitLogGroupName(ctx *LogGroupNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnNames.
	VisitColumnNames(ctx *ColumnNamesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tableNames.
	VisitTableNames(ctx *TableNamesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#oracleId.
	VisitOracleId(ctx *OracleIdContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#collationName.
	VisitCollationName(ctx *CollationNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnCollationName.
	VisitColumnCollationName(ctx *ColumnCollationNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dataTypeLength.
	VisitDataTypeLength(ctx *DataTypeLengthContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#primaryKey.
	VisitPrimaryKey(ctx *PrimaryKeyContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#exprs.
	VisitExprs(ctx *ExprsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#andOperator.
	VisitAndOperator(ctx *AndOperatorContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#orOperator.
	VisitOrOperator(ctx *OrOperatorContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#notOperator.
	VisitNotOperator(ctx *NotOperatorContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#booleanPrimary.
	VisitBooleanPrimary(ctx *BooleanPrimaryContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#bitExpr.
	VisitBitExpr(ctx *BitExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#simpleExpr.
	VisitSimpleExpr(ctx *SimpleExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#aggregationFunction.
	VisitAggregationFunction(ctx *AggregationFunctionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#aggregationFunctionName.
	VisitAggregationFunctionName(ctx *AggregationFunctionNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#analyticClause.
	VisitAnalyticClause(ctx *AnalyticClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#windowingClause.
	VisitWindowingClause(ctx *WindowingClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#specialFunction.
	VisitSpecialFunction(ctx *SpecialFunctionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#castFunction.
	VisitCastFunction(ctx *CastFunctionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#charFunction.
	VisitCharFunction(ctx *CharFunctionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#regularFunction.
	VisitRegularFunction(ctx *RegularFunctionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#regularFunctionName.
	VisitRegularFunctionName(ctx *RegularFunctionNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#caseExpression.
	VisitCaseExpression(ctx *CaseExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#caseWhen.
	VisitCaseWhen(ctx *CaseWhenContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#caseElse.
	VisitCaseElse(ctx *CaseElseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#orderByItem.
	VisitOrderByItem(ctx *OrderByItemContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#attributeName.
	VisitAttributeName(ctx *AttributeNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#simpleExprs.
	VisitSimpleExprs(ctx *SimpleExprsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lobItem.
	VisitLobItem(ctx *LobItemContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lobItems.
	VisitLobItems(ctx *LobItemsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lobItemList.
	VisitLobItemList(ctx *LobItemListContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#specialDatatype.
	VisitSpecialDatatype(ctx *SpecialDatatypeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dataTypeName.
	VisitDataTypeName(ctx *DataTypeNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#datetimeTypeSuffix.
	VisitDatetimeTypeSuffix(ctx *DatetimeTypeSuffixContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#treatFunction.
	VisitTreatFunction(ctx *TreatFunctionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#privateExprOfDb.
	VisitPrivateExprOfDb(ctx *PrivateExprOfDbContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#caseExpr.
	VisitCaseExpr(ctx *CaseExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#simpleCaseExpr.
	VisitSimpleCaseExpr(ctx *SimpleCaseExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#searchedCaseExpr.
	VisitSearchedCaseExpr(ctx *SearchedCaseExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#elseClause.
	VisitElseClause(ctx *ElseClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#intervalExpression.
	VisitIntervalExpression(ctx *IntervalExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#objectAccessExpression.
	VisitObjectAccessExpression(ctx *ObjectAccessExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#constructorExpr.
	VisitConstructorExpr(ctx *ConstructorExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#ignoredIdentifier.
	VisitIgnoredIdentifier(ctx *IgnoredIdentifierContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#ignoredIdentifiers.
	VisitIgnoredIdentifiers(ctx *IgnoredIdentifiersContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#matchNone.
	VisitMatchNone(ctx *MatchNoneContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hashSubpartitionQuantity.
	VisitHashSubpartitionQuantity(ctx *HashSubpartitionQuantityContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#odciParameters.
	VisitOdciParameters(ctx *OdciParametersContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#databaseName.
	VisitDatabaseName(ctx *DatabaseNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#locationName.
	VisitLocationName(ctx *LocationNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#fileName.
	VisitFileName(ctx *FileNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#asmFileName.
	VisitAsmFileName(ctx *AsmFileNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#fileNumber.
	VisitFileNumber(ctx *FileNumberContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#instanceName.
	VisitInstanceName(ctx *InstanceNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#logminerSessionName.
	VisitLogminerSessionName(ctx *LogminerSessionNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tablespaceGroupName.
	VisitTablespaceGroupName(ctx *TablespaceGroupNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#copyName.
	VisitCopyName(ctx *CopyNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#mirrorName.
	VisitMirrorName(ctx *MirrorNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#uriString.
	VisitUriString(ctx *UriStringContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#qualifiedCredentialName.
	VisitQualifiedCredentialName(ctx *QualifiedCredentialNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#pdbName.
	VisitPdbName(ctx *PdbNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#diskgroupName.
	VisitDiskgroupName(ctx *DiskgroupNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#templateName.
	VisitTemplateName(ctx *TemplateNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#aliasName.
	VisitAliasName(ctx *AliasNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#domain.
	VisitDomain(ctx *DomainContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dateValue.
	VisitDateValue(ctx *DateValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#sessionId.
	VisitSessionId(ctx *SessionIdContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#serialNumber.
	VisitSerialNumber(ctx *SerialNumberContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#instanceId.
	VisitInstanceId(ctx *InstanceIdContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#sqlId.
	VisitSqlId(ctx *SqlIdContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#logFileName.
	VisitLogFileName(ctx *LogFileNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#logFileGroupsArchivedLocationName.
	VisitLogFileGroupsArchivedLocationName(ctx *LogFileGroupsArchivedLocationNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#asmVersion.
	VisitAsmVersion(ctx *AsmVersionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#walletPassword.
	VisitWalletPassword(ctx *WalletPasswordContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hsmAuthString.
	VisitHsmAuthString(ctx *HsmAuthStringContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#targetDbName.
	VisitTargetDbName(ctx *TargetDbNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#certificateId.
	VisitCertificateId(ctx *CertificateIdContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#categoryName.
	VisitCategoryName(ctx *CategoryNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#offset.
	VisitOffset(ctx *OffsetContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowcount.
	VisitRowcount(ctx *RowcountContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#percent.
	VisitPercent(ctx *PercentContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rollbackSegment.
	VisitRollbackSegment(ctx *RollbackSegmentContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#queryName.
	VisitQueryName(ctx *QueryNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#cycleValue.
	VisitCycleValue(ctx *CycleValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#noCycleValue.
	VisitNoCycleValue(ctx *NoCycleValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#orderingColumn.
	VisitOrderingColumn(ctx *OrderingColumnContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#subavName.
	VisitSubavName(ctx *SubavNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#baseAvName.
	VisitBaseAvName(ctx *BaseAvNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#measName.
	VisitMeasName(ctx *MeasNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#levelRef.
	VisitLevelRef(ctx *LevelRefContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#offsetExpr.
	VisitOffsetExpr(ctx *OffsetExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#memberKeyExpr.
	VisitMemberKeyExpr(ctx *MemberKeyExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#depthExpression.
	VisitDepthExpression(ctx *DepthExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#unitName.
	VisitUnitName(ctx *UnitNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#procedureName.
	VisitProcedureName(ctx *ProcedureNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#cpuCost.
	VisitCpuCost(ctx *CpuCostContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#ioCost.
	VisitIoCost(ctx *IoCostContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#networkCost.
	VisitNetworkCost(ctx *NetworkCostContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#defaultSelectivity.
	VisitDefaultSelectivity(ctx *DefaultSelectivityContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dataItem.
	VisitDataItem(ctx *DataItemContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#variableName.
	VisitVariableName(ctx *VariableNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#validTimeColumn.
	VisitValidTimeColumn(ctx *ValidTimeColumnContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#attrDim.
	VisitAttrDim(ctx *AttrDimContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hierarchyName.
	VisitHierarchyName(ctx *HierarchyNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#analyticViewName.
	VisitAnalyticViewName(ctx *AnalyticViewNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#samplePercent.
	VisitSamplePercent(ctx *SamplePercentContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#seedValue.
	VisitSeedValue(ctx *SeedValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#namespace.
	VisitNamespace(ctx *NamespaceContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#restorePoint.
	VisitRestorePoint(ctx *RestorePointContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#scnValue.
	VisitScnValue(ctx *ScnValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#timestampValue.
	VisitTimestampValue(ctx *TimestampValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#scnTimestampExpr.
	VisitScnTimestampExpr(ctx *ScnTimestampExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#referenceModelName.
	VisitReferenceModelName(ctx *ReferenceModelNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#mainModelName.
	VisitMainModelName(ctx *MainModelNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#measureColumn.
	VisitMeasureColumn(ctx *MeasureColumnContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dimensionColumn.
	VisitDimensionColumn(ctx *DimensionColumnContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#analyticFunctionName.
	VisitAnalyticFunctionName(ctx *AnalyticFunctionNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#comparisonCondition.
	VisitComparisonCondition(ctx *ComparisonConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#simpleComparisonCondition.
	VisitSimpleComparisonCondition(ctx *SimpleComparisonConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#groupComparisonCondition.
	VisitGroupComparisonCondition(ctx *GroupComparisonConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#floatingPointCondition.
	VisitFloatingPointCondition(ctx *FloatingPointConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#logicalCondition.
	VisitLogicalCondition(ctx *LogicalConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modelCondition.
	VisitModelCondition(ctx *ModelConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#isAnyCondition.
	VisitIsAnyCondition(ctx *IsAnyConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#isPresentCondition.
	VisitIsPresentCondition(ctx *IsPresentConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#cellReference.
	VisitCellReference(ctx *CellReferenceContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#multisetCondition.
	VisitMultisetCondition(ctx *MultisetConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#isASetCondition.
	VisitIsASetCondition(ctx *IsASetConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#isEmptyCondition.
	VisitIsEmptyCondition(ctx *IsEmptyConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#memberCondition.
	VisitMemberCondition(ctx *MemberConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#submultisetCondition.
	VisitSubmultisetCondition(ctx *SubmultisetConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#patternMatchingCondition.
	VisitPatternMatchingCondition(ctx *PatternMatchingConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#likeCondition.
	VisitLikeCondition(ctx *LikeConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#searchValue.
	VisitSearchValue(ctx *SearchValueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#escapeChar.
	VisitEscapeChar(ctx *EscapeCharContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#regexpLikeCondition.
	VisitRegexpLikeCondition(ctx *RegexpLikeConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#matchParam.
	VisitMatchParam(ctx *MatchParamContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rangeCondition.
	VisitRangeCondition(ctx *RangeConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#nullCondition.
	VisitNullCondition(ctx *NullConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#xmlCondition.
	VisitXmlCondition(ctx *XmlConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#equalsPathCondition.
	VisitEqualsPathCondition(ctx *EqualsPathConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#pathString.
	VisitPathString(ctx *PathStringContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#correlationInteger.
	VisitCorrelationInteger(ctx *CorrelationIntegerContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#underPathCondition.
	VisitUnderPathCondition(ctx *UnderPathConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#levels.
	VisitLevels(ctx *LevelsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonCondition.
	VisitJsonCondition(ctx *JsonConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#isJsonCondition.
	VisitIsJsonCondition(ctx *IsJsonConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonEqualCondition.
	VisitJsonEqualCondition(ctx *JsonEqualConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonExistsCondition.
	VisitJsonExistsCondition(ctx *JsonExistsConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonPassingClause.
	VisitJsonPassingClause(ctx *JsonPassingClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonExistsOnErrorClause.
	VisitJsonExistsOnErrorClause(ctx *JsonExistsOnErrorClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonExistsOnEmptyClause.
	VisitJsonExistsOnEmptyClause(ctx *JsonExistsOnEmptyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonTextcontainsCondition.
	VisitJsonTextcontainsCondition(ctx *JsonTextcontainsConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonBasicPathExpr.
	VisitJsonBasicPathExpr(ctx *JsonBasicPathExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonAbsolutePathExpr.
	VisitJsonAbsolutePathExpr(ctx *JsonAbsolutePathExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonNonfunctionSteps.
	VisitJsonNonfunctionSteps(ctx *JsonNonfunctionStepsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonObjectStep.
	VisitJsonObjectStep(ctx *JsonObjectStepContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonFieldName.
	VisitJsonFieldName(ctx *JsonFieldNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#letter.
	VisitLetter(ctx *LetterContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#digit.
	VisitDigit(ctx *DigitContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonArrayStep.
	VisitJsonArrayStep(ctx *JsonArrayStepContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonDescendentStep.
	VisitJsonDescendentStep(ctx *JsonDescendentStepContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonFunctionStep.
	VisitJsonFunctionStep(ctx *JsonFunctionStepContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonItemMethod.
	VisitJsonItemMethod(ctx *JsonItemMethodContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonFilterExpr.
	VisitJsonFilterExpr(ctx *JsonFilterExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonCond.
	VisitJsonCond(ctx *JsonCondContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonDisjunction.
	VisitJsonDisjunction(ctx *JsonDisjunctionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonConjunction.
	VisitJsonConjunction(ctx *JsonConjunctionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonNegation.
	VisitJsonNegation(ctx *JsonNegationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonExistsCond.
	VisitJsonExistsCond(ctx *JsonExistsCondContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonHasSubstringCond.
	VisitJsonHasSubstringCond(ctx *JsonHasSubstringCondContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonStartsWithCond.
	VisitJsonStartsWithCond(ctx *JsonStartsWithCondContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonLikeCond.
	VisitJsonLikeCond(ctx *JsonLikeCondContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonLikeRegexCond.
	VisitJsonLikeRegexCond(ctx *JsonLikeRegexCondContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonEqRegexCond.
	VisitJsonEqRegexCond(ctx *JsonEqRegexCondContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonInCond.
	VisitJsonInCond(ctx *JsonInCondContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#valueList.
	VisitValueList(ctx *ValueListContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonComparison.
	VisitJsonComparison(ctx *JsonComparisonContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonRelativePathExpr.
	VisitJsonRelativePathExpr(ctx *JsonRelativePathExprContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonComparePred.
	VisitJsonComparePred(ctx *JsonComparePredContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonVar.
	VisitJsonVar(ctx *JsonVarContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonScalar.
	VisitJsonScalar(ctx *JsonScalarContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonNumber.
	VisitJsonNumber(ctx *JsonNumberContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jsonString.
	VisitJsonString(ctx *JsonStringContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#compoundCondition.
	VisitCompoundCondition(ctx *CompoundConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#existsCondition.
	VisitExistsCondition(ctx *ExistsConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inCondition.
	VisitInCondition(ctx *InConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#isOfTypeCondition.
	VisitIsOfTypeCondition(ctx *IsOfTypeConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#databaseCharset.
	VisitDatabaseCharset(ctx *DatabaseCharsetContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#nationalCharset.
	VisitNationalCharset(ctx *NationalCharsetContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#filenamePattern.
	VisitFilenamePattern(ctx *FilenamePatternContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#connectString.
	VisitConnectString(ctx *ConnectStringContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createTable.
	VisitCreateTable(ctx *CreateTableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createIndex.
	VisitCreateIndex(ctx *CreateIndexContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterTable.
	VisitAlterTable(ctx *AlterTableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterIndex.
	VisitAlterIndex(ctx *AlterIndexContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dropTable.
	VisitDropTable(ctx *DropTableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dropIndex.
	VisitDropIndex(ctx *DropIndexContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#truncateTable.
	VisitTruncateTable(ctx *TruncateTableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createTableSpecification.
	VisitCreateTableSpecification(ctx *CreateTableSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tablespaceClauseWithParen.
	VisitTablespaceClauseWithParen(ctx *TablespaceClauseWithParenContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tablespaceClause.
	VisitTablespaceClause(ctx *TablespaceClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createSharingClause.
	VisitCreateSharingClause(ctx *CreateSharingClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createDefinitionClause.
	VisitCreateDefinitionClause(ctx *CreateDefinitionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createXMLTypeTableClause.
	VisitCreateXMLTypeTableClause(ctx *CreateXMLTypeTableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#xmlTypeStorageClause.
	VisitXmlTypeStorageClause(ctx *XmlTypeStorageClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#xmlSchemaSpecClause.
	VisitXmlSchemaSpecClause(ctx *XmlSchemaSpecClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#xmlTypeVirtualColumnsClause.
	VisitXmlTypeVirtualColumnsClause(ctx *XmlTypeVirtualColumnsClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#oidClause.
	VisitOidClause(ctx *OidClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#oidIndexClause.
	VisitOidIndexClause(ctx *OidIndexClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createRelationalTableClause.
	VisitCreateRelationalTableClause(ctx *CreateRelationalTableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createMemOptimizeClause.
	VisitCreateMemOptimizeClause(ctx *CreateMemOptimizeClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createParentClause.
	VisitCreateParentClause(ctx *CreateParentClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createObjectTableClause.
	VisitCreateObjectTableClause(ctx *CreateObjectTableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#relationalProperties.
	VisitRelationalProperties(ctx *RelationalPropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#relationalProperty.
	VisitRelationalProperty(ctx *RelationalPropertyContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnDefinition.
	VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#visibleClause.
	VisitVisibleClause(ctx *VisibleClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#defaultNullClause.
	VisitDefaultNullClause(ctx *DefaultNullClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#identityClause.
	VisitIdentityClause(ctx *IdentityClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#identifyOptions.
	VisitIdentifyOptions(ctx *IdentifyOptionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#identityOption.
	VisitIdentityOption(ctx *IdentityOptionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#encryptionSpecification.
	VisitEncryptionSpecification(ctx *EncryptionSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inlineConstraint.
	VisitInlineConstraint(ctx *InlineConstraintContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#referencesClause.
	VisitReferencesClause(ctx *ReferencesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#constraintState.
	VisitConstraintState(ctx *ConstraintStateContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#notDeferrable.
	VisitNotDeferrable(ctx *NotDeferrableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#initiallyClause.
	VisitInitiallyClause(ctx *InitiallyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#exceptionsClause.
	VisitExceptionsClause(ctx *ExceptionsClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#usingIndexClause.
	VisitUsingIndexClause(ctx *UsingIndexClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createIndexClause.
	VisitCreateIndexClause(ctx *CreateIndexClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inlineRefConstraint.
	VisitInlineRefConstraint(ctx *InlineRefConstraintContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#virtualColumnDefinition.
	VisitVirtualColumnDefinition(ctx *VirtualColumnDefinitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#outOfLineConstraint.
	VisitOutOfLineConstraint(ctx *OutOfLineConstraintContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#outOfLineRefConstraint.
	VisitOutOfLineRefConstraint(ctx *OutOfLineRefConstraintContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createIndexSpecification.
	VisitCreateIndexSpecification(ctx *CreateIndexSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#clusterIndexClause.
	VisitClusterIndexClause(ctx *ClusterIndexClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indexAttributes.
	VisitIndexAttributes(ctx *IndexAttributesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tableIndexClause.
	VisitTableIndexClause(ctx *TableIndexClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indexExpressions.
	VisitIndexExpressions(ctx *IndexExpressionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indexExpression.
	VisitIndexExpression(ctx *IndexExpressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#bitmapJoinIndexClause.
	VisitBitmapJoinIndexClause(ctx *BitmapJoinIndexClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnSortsClause_.
	VisitColumnSortsClause_(ctx *ColumnSortsClause_Context) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnSortClause_.
	VisitColumnSortClause_(ctx *ColumnSortClause_Context) interface{}

	// Visit a parse tree produced by OracleStatementParser#createIndexDefinitionClause.
	VisitCreateIndexDefinitionClause(ctx *CreateIndexDefinitionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tableAlias.
	VisitTableAlias(ctx *TableAliasContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterDefinitionClause.
	VisitAlterDefinitionClause(ctx *AlterDefinitionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterTableProperties.
	VisitAlterTableProperties(ctx *AlterTablePropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#renameTableSpecification.
	VisitRenameTableSpecification(ctx *RenameTableSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnClauses.
	VisitColumnClauses(ctx *ColumnClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#operateColumnClause.
	VisitOperateColumnClause(ctx *OperateColumnClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#addColumnSpecification.
	VisitAddColumnSpecification(ctx *AddColumnSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnOrVirtualDefinitions.
	VisitColumnOrVirtualDefinitions(ctx *ColumnOrVirtualDefinitionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnOrVirtualDefinition.
	VisitColumnOrVirtualDefinition(ctx *ColumnOrVirtualDefinitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnProperties.
	VisitColumnProperties(ctx *ColumnPropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnProperty.
	VisitColumnProperty(ctx *ColumnPropertyContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#objectTypeColProperties.
	VisitObjectTypeColProperties(ctx *ObjectTypeColPropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#substitutableColumnClause.
	VisitSubstitutableColumnClause(ctx *SubstitutableColumnClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modifyColumnSpecification.
	VisitModifyColumnSpecification(ctx *ModifyColumnSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modifyColProperties.
	VisitModifyColProperties(ctx *ModifyColPropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modifyColSubstitutable.
	VisitModifyColSubstitutable(ctx *ModifyColSubstitutableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dropColumnClause.
	VisitDropColumnClause(ctx *DropColumnClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dropColumnSpecification.
	VisitDropColumnSpecification(ctx *DropColumnSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnOrColumnList.
	VisitColumnOrColumnList(ctx *ColumnOrColumnListContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#cascadeOrInvalidate.
	VisitCascadeOrInvalidate(ctx *CascadeOrInvalidateContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#checkpointNumber.
	VisitCheckpointNumber(ctx *CheckpointNumberContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#renameColumnClause.
	VisitRenameColumnClause(ctx *RenameColumnClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#constraintClauses.
	VisitConstraintClauses(ctx *ConstraintClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#addConstraintSpecification.
	VisitAddConstraintSpecification(ctx *AddConstraintSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modifyConstraintClause.
	VisitModifyConstraintClause(ctx *ModifyConstraintClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#constraintWithName.
	VisitConstraintWithName(ctx *ConstraintWithNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#constraintOption.
	VisitConstraintOption(ctx *ConstraintOptionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#constraintPrimaryOrUnique.
	VisitConstraintPrimaryOrUnique(ctx *ConstraintPrimaryOrUniqueContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#renameConstraintClause.
	VisitRenameConstraintClause(ctx *RenameConstraintClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dropConstraintClause.
	VisitDropConstraintClause(ctx *DropConstraintClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterExternalTable.
	VisitAlterExternalTable(ctx *AlterExternalTableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#objectProperties.
	VisitObjectProperties(ctx *ObjectPropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterIndexInformationClause.
	VisitAlterIndexInformationClause(ctx *AlterIndexInformationClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#renameIndexClause.
	VisitRenameIndexClause(ctx *RenameIndexClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#objectTableSubstitution.
	VisitObjectTableSubstitution(ctx *ObjectTableSubstitutionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#memOptimizeClause.
	VisitMemOptimizeClause(ctx *MemOptimizeClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#memOptimizeReadClause.
	VisitMemOptimizeReadClause(ctx *MemOptimizeReadClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#memOptimizeWriteClause.
	VisitMemOptimizeWriteClause(ctx *MemOptimizeWriteClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#enableDisableClauses.
	VisitEnableDisableClauses(ctx *EnableDisableClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#enableDisableClause.
	VisitEnableDisableClause(ctx *EnableDisableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#enableDisableOthers.
	VisitEnableDisableOthers(ctx *EnableDisableOthersContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rebuildClause.
	VisitRebuildClause(ctx *RebuildClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#parallelClause.
	VisitParallelClause(ctx *ParallelClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#usableSpecification.
	VisitUsableSpecification(ctx *UsableSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#invalidationSpecification.
	VisitInvalidationSpecification(ctx *InvalidationSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#materializedViewLogClause.
	VisitMaterializedViewLogClause(ctx *MaterializedViewLogClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dropReuseClause.
	VisitDropReuseClause(ctx *DropReuseClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#collationClause.
	VisitCollationClause(ctx *CollationClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#commitClause.
	VisitCommitClause(ctx *CommitClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#physicalProperties.
	VisitPhysicalProperties(ctx *PhysicalPropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#deferredSegmentCreation.
	VisitDeferredSegmentCreation(ctx *DeferredSegmentCreationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#segmentAttributesClause.
	VisitSegmentAttributesClause(ctx *SegmentAttributesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#physicalAttributesClause.
	VisitPhysicalAttributesClause(ctx *PhysicalAttributesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#loggingClause.
	VisitLoggingClause(ctx *LoggingClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#storageClause.
	VisitStorageClause(ctx *StorageClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#sizeClause.
	VisitSizeClause(ctx *SizeClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#maxsizeClause.
	VisitMaxsizeClause(ctx *MaxsizeClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tableCompression.
	VisitTableCompression(ctx *TableCompressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inmemoryTableClause.
	VisitInmemoryTableClause(ctx *InmemoryTableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inmemoryAttributes.
	VisitInmemoryAttributes(ctx *InmemoryAttributesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inmemoryColumnClause.
	VisitInmemoryColumnClause(ctx *InmemoryColumnClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inmemoryMemcompress.
	VisitInmemoryMemcompress(ctx *InmemoryMemcompressContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inmemoryPriority.
	VisitInmemoryPriority(ctx *InmemoryPriorityContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inmemoryDistribute.
	VisitInmemoryDistribute(ctx *InmemoryDistributeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inmemoryDuplicate.
	VisitInmemoryDuplicate(ctx *InmemoryDuplicateContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#ilmClause.
	VisitIlmClause(ctx *IlmClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#ilmPolicyClause.
	VisitIlmPolicyClause(ctx *IlmPolicyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#ilmCompressionPolicy.
	VisitIlmCompressionPolicy(ctx *IlmCompressionPolicyContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#ilmTimePeriod.
	VisitIlmTimePeriod(ctx *IlmTimePeriodContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#ilmTieringPolicy.
	VisitIlmTieringPolicy(ctx *IlmTieringPolicyContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#ilmInmemoryPolicy.
	VisitIlmInmemoryPolicy(ctx *IlmInmemoryPolicyContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#organizationClause.
	VisitOrganizationClause(ctx *OrganizationClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#heapOrgTableClause.
	VisitHeapOrgTableClause(ctx *HeapOrgTableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indexOrgTableClause.
	VisitIndexOrgTableClause(ctx *IndexOrgTableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#externalTableClause.
	VisitExternalTableClause(ctx *ExternalTableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#externalTableDataProps.
	VisitExternalTableDataProps(ctx *ExternalTableDataPropsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#mappingTableClause.
	VisitMappingTableClause(ctx *MappingTableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#prefixCompression.
	VisitPrefixCompression(ctx *PrefixCompressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indexOrgOverflowClause.
	VisitIndexOrgOverflowClause(ctx *IndexOrgOverflowClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#externalPartitionClause.
	VisitExternalPartitionClause(ctx *ExternalPartitionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#clusterRelatedClause.
	VisitClusterRelatedClause(ctx *ClusterRelatedClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tableProperties.
	VisitTableProperties(ctx *TablePropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#readOnlyClause.
	VisitReadOnlyClause(ctx *ReadOnlyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indexingClause.
	VisitIndexingClause(ctx *IndexingClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tablePartitioningClauses.
	VisitTablePartitioningClauses(ctx *TablePartitioningClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rangePartitions.
	VisitRangePartitions(ctx *RangePartitionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rangeValuesClause.
	VisitRangeValuesClause(ctx *RangeValuesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tablePartitionDescription.
	VisitTablePartitionDescription(ctx *TablePartitionDescriptionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#inmemoryClause.
	VisitInmemoryClause(ctx *InmemoryClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#varrayColProperties.
	VisitVarrayColProperties(ctx *VarrayColPropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#nestedTableColProperties.
	VisitNestedTableColProperties(ctx *NestedTableColPropertiesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lobStorageClause.
	VisitLobStorageClause(ctx *LobStorageClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#varrayStorageClause.
	VisitVarrayStorageClause(ctx *VarrayStorageClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lobStorageParameters.
	VisitLobStorageParameters(ctx *LobStorageParametersContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lobParameters.
	VisitLobParameters(ctx *LobParametersContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lobRetentionClause.
	VisitLobRetentionClause(ctx *LobRetentionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lobDeduplicateClause.
	VisitLobDeduplicateClause(ctx *LobDeduplicateClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lobCompressionClause.
	VisitLobCompressionClause(ctx *LobCompressionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#externalPartSubpartDataProps.
	VisitExternalPartSubpartDataProps(ctx *ExternalPartSubpartDataPropsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#listPartitions.
	VisitListPartitions(ctx *ListPartitionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#listValuesClause.
	VisitListValuesClause(ctx *ListValuesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#listValues.
	VisitListValues(ctx *ListValuesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hashPartitions.
	VisitHashPartitions(ctx *HashPartitionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hashPartitionsByQuantity.
	VisitHashPartitionsByQuantity(ctx *HashPartitionsByQuantityContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indexCompression.
	VisitIndexCompression(ctx *IndexCompressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#advancedIndexCompression.
	VisitAdvancedIndexCompression(ctx *AdvancedIndexCompressionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#individualHashPartitions.
	VisitIndividualHashPartitions(ctx *IndividualHashPartitionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partitioningStorageClause.
	VisitPartitioningStorageClause(ctx *PartitioningStorageClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lobPartitioningStorage.
	VisitLobPartitioningStorage(ctx *LobPartitioningStorageContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#compositeRangePartitions.
	VisitCompositeRangePartitions(ctx *CompositeRangePartitionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#subpartitionByRange.
	VisitSubpartitionByRange(ctx *SubpartitionByRangeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#subpartitionByList.
	VisitSubpartitionByList(ctx *SubpartitionByListContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#subpartitionByHash.
	VisitSubpartitionByHash(ctx *SubpartitionByHashContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#subpartitionTemplate.
	VisitSubpartitionTemplate(ctx *SubpartitionTemplateContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rangeSubpartitionDesc.
	VisitRangeSubpartitionDesc(ctx *RangeSubpartitionDescContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#listSubpartitionDesc.
	VisitListSubpartitionDesc(ctx *ListSubpartitionDescContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#individualHashSubparts.
	VisitIndividualHashSubparts(ctx *IndividualHashSubpartsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rangePartitionDesc.
	VisitRangePartitionDesc(ctx *RangePartitionDescContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#compositeListPartitions.
	VisitCompositeListPartitions(ctx *CompositeListPartitionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#listPartitionDesc.
	VisitListPartitionDesc(ctx *ListPartitionDescContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#compositeHashPartitions.
	VisitCompositeHashPartitions(ctx *CompositeHashPartitionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#referencePartitioning.
	VisitReferencePartitioning(ctx *ReferencePartitioningContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#referencePartitionDesc.
	VisitReferencePartitionDesc(ctx *ReferencePartitionDescContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#constraint.
	VisitConstraint(ctx *ConstraintContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#systemPartitioning.
	VisitSystemPartitioning(ctx *SystemPartitioningContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#consistentHashPartitions.
	VisitConsistentHashPartitions(ctx *ConsistentHashPartitionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#consistentHashWithSubpartitions.
	VisitConsistentHashWithSubpartitions(ctx *ConsistentHashWithSubpartitionsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partitionsetClauses.
	VisitPartitionsetClauses(ctx *PartitionsetClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rangePartitionsetClause.
	VisitRangePartitionsetClause(ctx *RangePartitionsetClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rangePartitionsetDesc.
	VisitRangePartitionsetDesc(ctx *RangePartitionsetDescContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#listPartitionsetClause.
	VisitListPartitionsetClause(ctx *ListPartitionsetClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#attributeClusteringClause.
	VisitAttributeClusteringClause(ctx *AttributeClusteringClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#clusteringJoin.
	VisitClusteringJoin(ctx *ClusteringJoinContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#clusterClause.
	VisitClusterClause(ctx *ClusterClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#clusteringColumns.
	VisitClusteringColumns(ctx *ClusteringColumnsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#clusteringColumnGroup.
	VisitClusteringColumnGroup(ctx *ClusteringColumnGroupContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#clusteringWhen.
	VisitClusteringWhen(ctx *ClusteringWhenContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#zonemapClause.
	VisitZonemapClause(ctx *ZonemapClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowMovementClause.
	VisitRowMovementClause(ctx *RowMovementClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#flashbackArchiveClause.
	VisitFlashbackArchiveClause(ctx *FlashbackArchiveClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterSynonym.
	VisitAlterSynonym(ctx *AlterSynonymContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterTablePartitioning.
	VisitAlterTablePartitioning(ctx *AlterTablePartitioningContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modifyTablePartition.
	VisitModifyTablePartition(ctx *ModifyTablePartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modifyRangePartition.
	VisitModifyRangePartition(ctx *ModifyRangePartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modifyHashPartition.
	VisitModifyHashPartition(ctx *ModifyHashPartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#modifyListPartition.
	VisitModifyListPartition(ctx *ModifyListPartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partitionExtendedName.
	VisitPartitionExtendedName(ctx *PartitionExtendedNameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#addRangeSubpartition.
	VisitAddRangeSubpartition(ctx *AddRangeSubpartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dependentTablesClause.
	VisitDependentTablesClause(ctx *DependentTablesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#addHashSubpartition.
	VisitAddHashSubpartition(ctx *AddHashSubpartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#addListSubpartition.
	VisitAddListSubpartition(ctx *AddListSubpartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#coalesceTableSubpartition.
	VisitCoalesceTableSubpartition(ctx *CoalesceTableSubpartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#allowDisallowClustering.
	VisitAllowDisallowClustering(ctx *AllowDisallowClusteringContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterMappingTableClauses.
	VisitAlterMappingTableClauses(ctx *AlterMappingTableClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#deallocateUnusedClause.
	VisitDeallocateUnusedClause(ctx *DeallocateUnusedClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#allocateExtentClause.
	VisitAllocateExtentClause(ctx *AllocateExtentClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partitionSpec.
	VisitPartitionSpec(ctx *PartitionSpecContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partitionAttributes.
	VisitPartitionAttributes(ctx *PartitionAttributesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#shrinkClause.
	VisitShrinkClause(ctx *ShrinkClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#moveTablePartition.
	VisitMoveTablePartition(ctx *MoveTablePartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#filterCondition.
	VisitFilterCondition(ctx *FilterConditionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#coalesceTablePartition.
	VisitCoalesceTablePartition(ctx *CoalesceTablePartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#addTablePartition.
	VisitAddTablePartition(ctx *AddTablePartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#addRangePartitionClause.
	VisitAddRangePartitionClause(ctx *AddRangePartitionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#addListPartitionClause.
	VisitAddListPartitionClause(ctx *AddListPartitionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#hashSubpartsByQuantity.
	VisitHashSubpartsByQuantity(ctx *HashSubpartsByQuantityContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#addSystemPartitionClause.
	VisitAddSystemPartitionClause(ctx *AddSystemPartitionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#addHashPartitionClause.
	VisitAddHashPartitionClause(ctx *AddHashPartitionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dropTablePartition.
	VisitDropTablePartition(ctx *DropTablePartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partitionExtendedNames.
	VisitPartitionExtendedNames(ctx *PartitionExtendedNamesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partitionForClauses.
	VisitPartitionForClauses(ctx *PartitionForClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#updateIndexClauses.
	VisitUpdateIndexClauses(ctx *UpdateIndexClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#updateGlobalIndexClause.
	VisitUpdateGlobalIndexClause(ctx *UpdateGlobalIndexClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#updateAllIndexesClause.
	VisitUpdateAllIndexesClause(ctx *UpdateAllIndexesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#updateIndexPartition.
	VisitUpdateIndexPartition(ctx *UpdateIndexPartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indexPartitionDesc.
	VisitIndexPartitionDesc(ctx *IndexPartitionDescContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indexSubpartitionClause.
	VisitIndexSubpartitionClause(ctx *IndexSubpartitionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#updateIndexSubpartition.
	VisitUpdateIndexSubpartition(ctx *UpdateIndexSubpartitionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#supplementalLoggingProps.
	VisitSupplementalLoggingProps(ctx *SupplementalLoggingPropsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#supplementalLogGrpClause.
	VisitSupplementalLogGrpClause(ctx *SupplementalLogGrpClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#supplementalIdKeyClause.
	VisitSupplementalIdKeyClause(ctx *SupplementalIdKeyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterSession.
	VisitAlterSession(ctx *AlterSessionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterSessionOption.
	VisitAlterSessionOption(ctx *AlterSessionOptionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#adviseClause.
	VisitAdviseClause(ctx *AdviseClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#closeDatabaseLinkClause.
	VisitCloseDatabaseLinkClause(ctx *CloseDatabaseLinkClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#commitInProcedureClause.
	VisitCommitInProcedureClause(ctx *CommitInProcedureClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#securiyClause.
	VisitSecuriyClause(ctx *SecuriyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#parallelExecutionClause.
	VisitParallelExecutionClause(ctx *ParallelExecutionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#resumableClause.
	VisitResumableClause(ctx *ResumableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#enableResumableClause.
	VisitEnableResumableClause(ctx *EnableResumableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#disableResumableClause.
	VisitDisableResumableClause(ctx *DisableResumableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#shardDdlClause.
	VisitShardDdlClause(ctx *ShardDdlClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#syncWithPrimaryClause.
	VisitSyncWithPrimaryClause(ctx *SyncWithPrimaryClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterSessionSetClause.
	VisitAlterSessionSetClause(ctx *AlterSessionSetClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterSessionSetClauseOption.
	VisitAlterSessionSetClauseOption(ctx *AlterSessionSetClauseOptionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#parameterClause.
	VisitParameterClause(ctx *ParameterClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#editionClause.
	VisitEditionClause(ctx *EditionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#containerClause.
	VisitContainerClause(ctx *ContainerClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rowArchivalVisibilityClause.
	VisitRowArchivalVisibilityClause(ctx *RowArchivalVisibilityClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterDatabase.
	VisitAlterDatabase(ctx *AlterDatabaseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#databaseClauses.
	VisitDatabaseClauses(ctx *DatabaseClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#startupClauses.
	VisitStartupClauses(ctx *StartupClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#recoveryClauses.
	VisitRecoveryClauses(ctx *RecoveryClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#generalRecovery.
	VisitGeneralRecovery(ctx *GeneralRecoveryContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#fullDatabaseRecovery.
	VisitFullDatabaseRecovery(ctx *FullDatabaseRecoveryContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partialDatabaseRecovery.
	VisitPartialDatabaseRecovery(ctx *PartialDatabaseRecoveryContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#managedStandbyRecovery.
	VisitManagedStandbyRecovery(ctx *ManagedStandbyRecoveryContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#databaseFileClauses.
	VisitDatabaseFileClauses(ctx *DatabaseFileClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createDatafileClause.
	VisitCreateDatafileClause(ctx *CreateDatafileClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#fileSpecifications.
	VisitFileSpecifications(ctx *FileSpecificationsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#fileSpecification.
	VisitFileSpecification(ctx *FileSpecificationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#datafileTempfileSpec.
	VisitDatafileTempfileSpec(ctx *DatafileTempfileSpecContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#autoextendClause.
	VisitAutoextendClause(ctx *AutoextendClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#redoLogFileSpec.
	VisitRedoLogFileSpec(ctx *RedoLogFileSpecContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterDatafileClause.
	VisitAlterDatafileClause(ctx *AlterDatafileClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterTempfileClause.
	VisitAlterTempfileClause(ctx *AlterTempfileClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#logfileClauses.
	VisitLogfileClauses(ctx *LogfileClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#logfileDescriptor.
	VisitLogfileDescriptor(ctx *LogfileDescriptorContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#addLogfileClauses.
	VisitAddLogfileClauses(ctx *AddLogfileClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#controlfileClauses.
	VisitControlfileClauses(ctx *ControlfileClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#traceFileClause.
	VisitTraceFileClause(ctx *TraceFileClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dropLogfileClauses.
	VisitDropLogfileClauses(ctx *DropLogfileClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#switchLogfileClause.
	VisitSwitchLogfileClause(ctx *SwitchLogfileClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#supplementalDbLogging.
	VisitSupplementalDbLogging(ctx *SupplementalDbLoggingContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#supplementalPlsqlClause.
	VisitSupplementalPlsqlClause(ctx *SupplementalPlsqlClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#supplementalSubsetReplicationClause.
	VisitSupplementalSubsetReplicationClause(ctx *SupplementalSubsetReplicationClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#standbyDatabaseClauses.
	VisitStandbyDatabaseClauses(ctx *StandbyDatabaseClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#activateStandbyDbClause.
	VisitActivateStandbyDbClause(ctx *ActivateStandbyDbClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#maximizeStandbyDbClause.
	VisitMaximizeStandbyDbClause(ctx *MaximizeStandbyDbClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#registerLogfileClause.
	VisitRegisterLogfileClause(ctx *RegisterLogfileClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#commitSwitchoverClause.
	VisitCommitSwitchoverClause(ctx *CommitSwitchoverClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#startStandbyClause.
	VisitStartStandbyClause(ctx *StartStandbyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#stopStandbyClause.
	VisitStopStandbyClause(ctx *StopStandbyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#switchoverClause.
	VisitSwitchoverClause(ctx *SwitchoverClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#convertDatabaseClause.
	VisitConvertDatabaseClause(ctx *ConvertDatabaseClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#failoverClause.
	VisitFailoverClause(ctx *FailoverClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#defaultSettingsClauses.
	VisitDefaultSettingsClauses(ctx *DefaultSettingsClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#setTimeZoneClause.
	VisitSetTimeZoneClause(ctx *SetTimeZoneClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#timeZoneRegion.
	VisitTimeZoneRegion(ctx *TimeZoneRegionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#flashbackModeClause.
	VisitFlashbackModeClause(ctx *FlashbackModeClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#undoModeClause.
	VisitUndoModeClause(ctx *UndoModeClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#moveDatafileClause.
	VisitMoveDatafileClause(ctx *MoveDatafileClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#instanceClauses.
	VisitInstanceClauses(ctx *InstanceClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#securityClause.
	VisitSecurityClause(ctx *SecurityClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#prepareClause.
	VisitPrepareClause(ctx *PrepareClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dropMirrorCopy.
	VisitDropMirrorCopy(ctx *DropMirrorCopyContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#lostWriteProtection.
	VisitLostWriteProtection(ctx *LostWriteProtectionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#cdbFleetClauses.
	VisitCdbFleetClauses(ctx *CdbFleetClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#leadCdbClause.
	VisitLeadCdbClause(ctx *LeadCdbClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#leadCdbUriClause.
	VisitLeadCdbUriClause(ctx *LeadCdbUriClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#propertyClause.
	VisitPropertyClause(ctx *PropertyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterSystem.
	VisitAlterSystem(ctx *AlterSystemContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterSystemOption.
	VisitAlterSystemOption(ctx *AlterSystemOptionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#archiveLogClause.
	VisitArchiveLogClause(ctx *ArchiveLogClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#checkpointClause.
	VisitCheckpointClause(ctx *CheckpointClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#checkDatafilesClause.
	VisitCheckDatafilesClause(ctx *CheckDatafilesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#distributedRecovClauses.
	VisitDistributedRecovClauses(ctx *DistributedRecovClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#flushClause.
	VisitFlushClause(ctx *FlushClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#endSessionClauses.
	VisitEndSessionClauses(ctx *EndSessionClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterSystemSwitchLogfileClause.
	VisitAlterSystemSwitchLogfileClause(ctx *AlterSystemSwitchLogfileClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#suspendResumeClause.
	VisitSuspendResumeClause(ctx *SuspendResumeClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#quiesceClauses.
	VisitQuiesceClauses(ctx *QuiesceClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rollingMigrationClauses.
	VisitRollingMigrationClauses(ctx *RollingMigrationClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rollingPatchClauses.
	VisitRollingPatchClauses(ctx *RollingPatchClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterSystemSecurityClauses.
	VisitAlterSystemSecurityClauses(ctx *AlterSystemSecurityClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#affinityClauses.
	VisitAffinityClauses(ctx *AffinityClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#shutdownDispatcherClause.
	VisitShutdownDispatcherClause(ctx *ShutdownDispatcherClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#registerClause.
	VisitRegisterClause(ctx *RegisterClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#setClause.
	VisitSetClause(ctx *SetClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#resetClause.
	VisitResetClause(ctx *ResetClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#relocateClientClause.
	VisitRelocateClientClause(ctx *RelocateClientClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#cancelSqlClause.
	VisitCancelSqlClause(ctx *CancelSqlClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#flushPasswordfileMetadataCacheClause.
	VisitFlushPasswordfileMetadataCacheClause(ctx *FlushPasswordfileMetadataCacheClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#instanceClause.
	VisitInstanceClause(ctx *InstanceClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#sequenceClause.
	VisitSequenceClause(ctx *SequenceClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#changeClause.
	VisitChangeClause(ctx *ChangeClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#currentClause.
	VisitCurrentClause(ctx *CurrentClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#groupClause.
	VisitGroupClause(ctx *GroupClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#logfileClause.
	VisitLogfileClause(ctx *LogfileClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#nextClause.
	VisitNextClause(ctx *NextClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#allClause.
	VisitAllClause(ctx *AllClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#toLocationClause.
	VisitToLocationClause(ctx *ToLocationClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#flushClauseOption.
	VisitFlushClauseOption(ctx *FlushClauseOptionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#disconnectSessionClause.
	VisitDisconnectSessionClause(ctx *DisconnectSessionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#killSessionClause.
	VisitKillSessionClause(ctx *KillSessionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#startRollingMigrationClause.
	VisitStartRollingMigrationClause(ctx *StartRollingMigrationClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#stopRollingMigrationClause.
	VisitStopRollingMigrationClause(ctx *StopRollingMigrationClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#startRollingPatchClause.
	VisitStartRollingPatchClause(ctx *StartRollingPatchClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#stopRollingPatchClause.
	VisitStopRollingPatchClause(ctx *StopRollingPatchClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#restrictedSessionClause.
	VisitRestrictedSessionClause(ctx *RestrictedSessionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#setEncryptionWalletOpenClause.
	VisitSetEncryptionWalletOpenClause(ctx *SetEncryptionWalletOpenClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#setEncryptionWalletCloseClause.
	VisitSetEncryptionWalletCloseClause(ctx *SetEncryptionWalletCloseClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#setEncryptionKeyClause.
	VisitSetEncryptionKeyClause(ctx *SetEncryptionKeyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#enableAffinityClause.
	VisitEnableAffinityClause(ctx *EnableAffinityClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#disableAffinityClause.
	VisitDisableAffinityClause(ctx *DisableAffinityClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterSystemSetClause.
	VisitAlterSystemSetClause(ctx *AlterSystemSetClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterSystemResetClause.
	VisitAlterSystemResetClause(ctx *AlterSystemResetClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#sharedPoolClause.
	VisitSharedPoolClause(ctx *SharedPoolClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#globalContextClause.
	VisitGlobalContextClause(ctx *GlobalContextClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#bufferCacheClause.
	VisitBufferCacheClause(ctx *BufferCacheClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#flashCacheClause.
	VisitFlashCacheClause(ctx *FlashCacheClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#redoToClause.
	VisitRedoToClause(ctx *RedoToClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#identifiedByWalletPassword.
	VisitIdentifiedByWalletPassword(ctx *IdentifiedByWalletPasswordContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#identifiedByHsmAuthString.
	VisitIdentifiedByHsmAuthString(ctx *IdentifiedByHsmAuthStringContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#setParameterClause.
	VisitSetParameterClause(ctx *SetParameterClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#useStoredOutlinesClause.
	VisitUseStoredOutlinesClause(ctx *UseStoredOutlinesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#globalTopicEnabledClause.
	VisitGlobalTopicEnabledClause(ctx *GlobalTopicEnabledClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterSystemCommentClause.
	VisitAlterSystemCommentClause(ctx *AlterSystemCommentClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#containerCurrentAllClause.
	VisitContainerCurrentAllClause(ctx *ContainerCurrentAllClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#scopeClause.
	VisitScopeClause(ctx *ScopeClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#analyze.
	VisitAnalyze(ctx *AnalyzeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#partitionExtensionClause.
	VisitPartitionExtensionClause(ctx *PartitionExtensionClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#validationClauses.
	VisitValidationClauses(ctx *ValidationClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#associateStatistics.
	VisitAssociateStatistics(ctx *AssociateStatisticsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#columnAssociation.
	VisitColumnAssociation(ctx *ColumnAssociationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#functionAssociation.
	VisitFunctionAssociation(ctx *FunctionAssociationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#storageTableClause.
	VisitStorageTableClause(ctx *StorageTableClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#usingStatisticsType.
	VisitUsingStatisticsType(ctx *UsingStatisticsTypeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#defaultCostClause.
	VisitDefaultCostClause(ctx *DefaultCostClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#defaultSelectivityClause.
	VisitDefaultSelectivityClause(ctx *DefaultSelectivityClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#disassociateStatistics.
	VisitDisassociateStatistics(ctx *DisassociateStatisticsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#audit.
	VisitAudit(ctx *AuditContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#noAudit.
	VisitNoAudit(ctx *NoAuditContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#auditPolicyClause.
	VisitAuditPolicyClause(ctx *AuditPolicyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#noAuditPolicyClause.
	VisitNoAuditPolicyClause(ctx *NoAuditPolicyClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#byUsersWithRoles.
	VisitByUsersWithRoles(ctx *ByUsersWithRolesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#contextClause.
	VisitContextClause(ctx *ContextClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#contextNamespaceAttributesClause.
	VisitContextNamespaceAttributesClause(ctx *ContextNamespaceAttributesClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#flashbackDatabase.
	VisitFlashbackDatabase(ctx *FlashbackDatabaseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#scnTimestampClause.
	VisitScnTimestampClause(ctx *ScnTimestampClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#restorePointClause.
	VisitRestorePointClause(ctx *RestorePointClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#flashbackTable.
	VisitFlashbackTable(ctx *FlashbackTableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#renameToTable.
	VisitRenameToTable(ctx *RenameToTableContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#purge.
	VisitPurge(ctx *PurgeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rename.
	VisitRename(ctx *RenameContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createDatabase.
	VisitCreateDatabase(ctx *CreateDatabaseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createDatabaseClauses.
	VisitCreateDatabaseClauses(ctx *CreateDatabaseClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#databaseLoggingClauses.
	VisitDatabaseLoggingClauses(ctx *DatabaseLoggingClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tablespaceClauses.
	VisitTablespaceClauses(ctx *TablespaceClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#defaultTablespace.
	VisitDefaultTablespace(ctx *DefaultTablespaceContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#defaultTempTablespace.
	VisitDefaultTempTablespace(ctx *DefaultTempTablespaceContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#undoTablespace.
	VisitUndoTablespace(ctx *UndoTablespaceContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#bigOrSmallFiles.
	VisitBigOrSmallFiles(ctx *BigOrSmallFilesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#extentManagementClause.
	VisitExtentManagementClause(ctx *ExtentManagementClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#enablePluggableDatabase.
	VisitEnablePluggableDatabase(ctx *EnablePluggableDatabaseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#fileNameConvert.
	VisitFileNameConvert(ctx *FileNameConvertContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#replaceFileNamePattern.
	VisitReplaceFileNamePattern(ctx *ReplaceFileNamePatternContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tablespaceDatafileClauses.
	VisitTablespaceDatafileClauses(ctx *TablespaceDatafileClausesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createDatabaseLink.
	VisitCreateDatabaseLink(ctx *CreateDatabaseLinkContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#connectToClause.
	VisitConnectToClause(ctx *ConnectToClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dbLinkAuthentication.
	VisitDbLinkAuthentication(ctx *DbLinkAuthenticationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#setTransaction.
	VisitSetTransaction(ctx *SetTransactionContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#commit.
	VisitCommit(ctx *CommitContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#commentClause.
	VisitCommentClause(ctx *CommentClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#writeClause.
	VisitWriteClause(ctx *WriteClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#forceClause.
	VisitForceClause(ctx *ForceClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rollback.
	VisitRollback(ctx *RollbackContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#savepointClause.
	VisitSavepointClause(ctx *SavepointClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#savepoint.
	VisitSavepoint(ctx *SavepointContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#setConstraints.
	VisitSetConstraints(ctx *SetConstraintsContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#grant.
	VisitGrant(ctx *GrantContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#revoke.
	VisitRevoke(ctx *RevokeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#objectPrivilegeClause.
	VisitObjectPrivilegeClause(ctx *ObjectPrivilegeClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#systemPrivilegeClause.
	VisitSystemPrivilegeClause(ctx *SystemPrivilegeClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#roleClause.
	VisitRoleClause(ctx *RoleClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#objectPrivileges.
	VisitObjectPrivileges(ctx *ObjectPrivilegesContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#objectPrivilegeType.
	VisitObjectPrivilegeType(ctx *ObjectPrivilegeTypeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#onObjectClause.
	VisitOnObjectClause(ctx *OnObjectClauseContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#systemPrivilege.
	VisitSystemPrivilege(ctx *SystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#systemPrivilegeOperation.
	VisitSystemPrivilegeOperation(ctx *SystemPrivilegeOperationContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#advisorFrameworkSystemPrivilege.
	VisitAdvisorFrameworkSystemPrivilege(ctx *AdvisorFrameworkSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#clustersSystemPrivilege.
	VisitClustersSystemPrivilege(ctx *ClustersSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#contextsSystemPrivilege.
	VisitContextsSystemPrivilege(ctx *ContextsSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dataRedactionSystemPrivilege.
	VisitDataRedactionSystemPrivilege(ctx *DataRedactionSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#databaseSystemPrivilege.
	VisitDatabaseSystemPrivilege(ctx *DatabaseSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#databaseLinksSystemPrivilege.
	VisitDatabaseLinksSystemPrivilege(ctx *DatabaseLinksSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#debuggingSystemPrivilege.
	VisitDebuggingSystemPrivilege(ctx *DebuggingSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dictionariesSystemPrivilege.
	VisitDictionariesSystemPrivilege(ctx *DictionariesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dimensionsSystemPrivilege.
	VisitDimensionsSystemPrivilege(ctx *DimensionsSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#directoriesSystemPrivilege.
	VisitDirectoriesSystemPrivilege(ctx *DirectoriesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#editionsSystemPrivilege.
	VisitEditionsSystemPrivilege(ctx *EditionsSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#flashbackDataArchivesPrivilege.
	VisitFlashbackDataArchivesPrivilege(ctx *FlashbackDataArchivesPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indexesSystemPrivilege.
	VisitIndexesSystemPrivilege(ctx *IndexesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#indexTypesSystemPrivilege.
	VisitIndexTypesSystemPrivilege(ctx *IndexTypesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#jobSchedulerObjectsSystemPrivilege.
	VisitJobSchedulerObjectsSystemPrivilege(ctx *JobSchedulerObjectsSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#keyManagementFrameworkSystemPrivilege.
	VisitKeyManagementFrameworkSystemPrivilege(ctx *KeyManagementFrameworkSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#librariesFrameworkSystemPrivilege.
	VisitLibrariesFrameworkSystemPrivilege(ctx *LibrariesFrameworkSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#logminerFrameworkSystemPrivilege.
	VisitLogminerFrameworkSystemPrivilege(ctx *LogminerFrameworkSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#materizlizedViewsSystemPrivilege.
	VisitMaterizlizedViewsSystemPrivilege(ctx *MaterizlizedViewsSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#miningModelsSystemPrivilege.
	VisitMiningModelsSystemPrivilege(ctx *MiningModelsSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#olapCubesSystemPrivilege.
	VisitOlapCubesSystemPrivilege(ctx *OlapCubesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#olapCubeMeasureFoldersSystemPrivilege.
	VisitOlapCubeMeasureFoldersSystemPrivilege(ctx *OlapCubeMeasureFoldersSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#olapCubeDiminsionsSystemPrivilege.
	VisitOlapCubeDiminsionsSystemPrivilege(ctx *OlapCubeDiminsionsSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#olapCubeBuildProcessesSystemPrivilege.
	VisitOlapCubeBuildProcessesSystemPrivilege(ctx *OlapCubeBuildProcessesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#operatorsSystemPrivilege.
	VisitOperatorsSystemPrivilege(ctx *OperatorsSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#outlinesSystemPrivilege.
	VisitOutlinesSystemPrivilege(ctx *OutlinesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#planManagementSystemPrivilege.
	VisitPlanManagementSystemPrivilege(ctx *PlanManagementSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#pluggableDatabasesSystemPrivilege.
	VisitPluggableDatabasesSystemPrivilege(ctx *PluggableDatabasesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#proceduresSystemPrivilege.
	VisitProceduresSystemPrivilege(ctx *ProceduresSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#profilesSystemPrivilege.
	VisitProfilesSystemPrivilege(ctx *ProfilesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rolesSystemPrivilege.
	VisitRolesSystemPrivilege(ctx *RolesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#rollbackSegmentsSystemPrivilege.
	VisitRollbackSegmentsSystemPrivilege(ctx *RollbackSegmentsSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#sequencesSystemPrivilege.
	VisitSequencesSystemPrivilege(ctx *SequencesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#sessionsSystemPrivilege.
	VisitSessionsSystemPrivilege(ctx *SessionsSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#sqlTranslationProfilesSystemPrivilege.
	VisitSqlTranslationProfilesSystemPrivilege(ctx *SqlTranslationProfilesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#synonymsSystemPrivilege.
	VisitSynonymsSystemPrivilege(ctx *SynonymsSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tablesSystemPrivilege.
	VisitTablesSystemPrivilege(ctx *TablesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#tablespacesSystemPrivilege.
	VisitTablespacesSystemPrivilege(ctx *TablespacesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#triggersSystemPrivilege.
	VisitTriggersSystemPrivilege(ctx *TriggersSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#typesSystemPrivilege.
	VisitTypesSystemPrivilege(ctx *TypesSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#usersSystemPrivilege.
	VisitUsersSystemPrivilege(ctx *UsersSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#viewsSystemPrivilege.
	VisitViewsSystemPrivilege(ctx *ViewsSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#miscellaneousSystemPrivilege.
	VisitMiscellaneousSystemPrivilege(ctx *MiscellaneousSystemPrivilegeContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createUser.
	VisitCreateUser(ctx *CreateUserContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dropUser.
	VisitDropUser(ctx *DropUserContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterUser.
	VisitAlterUser(ctx *AlterUserContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#createRole.
	VisitCreateRole(ctx *CreateRoleContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#dropRole.
	VisitDropRole(ctx *DropRoleContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#alterRole.
	VisitAlterRole(ctx *AlterRoleContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#setRole.
	VisitSetRole(ctx *SetRoleContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#roleAssignment.
	VisitRoleAssignment(ctx *RoleAssignmentContext) interface{}

	// Visit a parse tree produced by OracleStatementParser#call.
	VisitCall(ctx *CallContext) interface{}
}

// Code generated from E:/GoProject/src/github.com/2212442929/xsqlparser/dialect/oracle/grammer\OracleStatement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // OracleStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OracleStatementListener is a complete listener for a parse tree produced by OracleStatementParser.
type OracleStatementListener interface {
	antlr.ParseTreeListener

	// EnterExecute is called when entering the execute production.
	EnterExecute(c *ExecuteContext)

	// EnterInsert is called when entering the insert production.
	EnterInsert(c *InsertContext)

	// EnterInsertSingleTable is called when entering the insertSingleTable production.
	EnterInsertSingleTable(c *InsertSingleTableContext)

	// EnterInsertMultiTable is called when entering the insertMultiTable production.
	EnterInsertMultiTable(c *InsertMultiTableContext)

	// EnterMultiTableElement is called when entering the multiTableElement production.
	EnterMultiTableElement(c *MultiTableElementContext)

	// EnterConditionalInsertClause is called when entering the conditionalInsertClause production.
	EnterConditionalInsertClause(c *ConditionalInsertClauseContext)

	// EnterConditionalInsertWhenPart is called when entering the conditionalInsertWhenPart production.
	EnterConditionalInsertWhenPart(c *ConditionalInsertWhenPartContext)

	// EnterConditionalInsertElsePart is called when entering the conditionalInsertElsePart production.
	EnterConditionalInsertElsePart(c *ConditionalInsertElsePartContext)

	// EnterInsertIntoClause is called when entering the insertIntoClause production.
	EnterInsertIntoClause(c *InsertIntoClauseContext)

	// EnterInsertValuesClause is called when entering the insertValuesClause production.
	EnterInsertValuesClause(c *InsertValuesClauseContext)

	// EnterReturningClause is called when entering the returningClause production.
	EnterReturningClause(c *ReturningClauseContext)

	// EnterDmlTableExprClause is called when entering the dmlTableExprClause production.
	EnterDmlTableExprClause(c *DmlTableExprClauseContext)

	// EnterDmlTableClause is called when entering the dmlTableClause production.
	EnterDmlTableClause(c *DmlTableClauseContext)

	// EnterPartitionExtClause is called when entering the partitionExtClause production.
	EnterPartitionExtClause(c *PartitionExtClauseContext)

	// EnterDmlSubqueryClause is called when entering the dmlSubqueryClause production.
	EnterDmlSubqueryClause(c *DmlSubqueryClauseContext)

	// EnterSubqueryRestrictionClause is called when entering the subqueryRestrictionClause production.
	EnterSubqueryRestrictionClause(c *SubqueryRestrictionClauseContext)

	// EnterTableCollectionExpr is called when entering the tableCollectionExpr production.
	EnterTableCollectionExpr(c *TableCollectionExprContext)

	// EnterCollectionExpr is called when entering the collectionExpr production.
	EnterCollectionExpr(c *CollectionExprContext)

	// EnterUpdate is called when entering the update production.
	EnterUpdate(c *UpdateContext)

	// EnterUpdateSpecification is called when entering the updateSpecification production.
	EnterUpdateSpecification(c *UpdateSpecificationContext)

	// EnterUpdateSetClause is called when entering the updateSetClause production.
	EnterUpdateSetClause(c *UpdateSetClauseContext)

	// EnterUpdateSetColumnList is called when entering the updateSetColumnList production.
	EnterUpdateSetColumnList(c *UpdateSetColumnListContext)

	// EnterUpdateSetColumnClause is called when entering the updateSetColumnClause production.
	EnterUpdateSetColumnClause(c *UpdateSetColumnClauseContext)

	// EnterUpdateSetValueClause is called when entering the updateSetValueClause production.
	EnterUpdateSetValueClause(c *UpdateSetValueClauseContext)

	// EnterAssignmentValues is called when entering the assignmentValues production.
	EnterAssignmentValues(c *AssignmentValuesContext)

	// EnterAssignmentValue is called when entering the assignmentValue production.
	EnterAssignmentValue(c *AssignmentValueContext)

	// EnterDelete is called when entering the delete production.
	EnterDelete(c *DeleteContext)

	// EnterDeleteSpecification is called when entering the deleteSpecification production.
	EnterDeleteSpecification(c *DeleteSpecificationContext)

	// EnterSelect is called when entering the select production.
	EnterSelect(c *SelectContext)

	// EnterSelectSubquery is called when entering the selectSubquery production.
	EnterSelectSubquery(c *SelectSubqueryContext)

	// EnterSelectUnionClause is called when entering the selectUnionClause production.
	EnterSelectUnionClause(c *SelectUnionClauseContext)

	// EnterParenthesisSelectSubquery is called when entering the parenthesisSelectSubquery production.
	EnterParenthesisSelectSubquery(c *ParenthesisSelectSubqueryContext)

	// EnterQueryBlock is called when entering the queryBlock production.
	EnterQueryBlock(c *QueryBlockContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterPlsqlDeclarations is called when entering the plsqlDeclarations production.
	EnterPlsqlDeclarations(c *PlsqlDeclarationsContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionHeading is called when entering the functionHeading production.
	EnterFunctionHeading(c *FunctionHeadingContext)

	// EnterParameterDeclaration is called when entering the parameterDeclaration production.
	EnterParameterDeclaration(c *ParameterDeclarationContext)

	// EnterProcedureDeclaration is called when entering the procedureDeclaration production.
	EnterProcedureDeclaration(c *ProcedureDeclarationContext)

	// EnterProcedureHeading is called when entering the procedureHeading production.
	EnterProcedureHeading(c *ProcedureHeadingContext)

	// EnterProcedureProperties is called when entering the procedureProperties production.
	EnterProcedureProperties(c *ProcedurePropertiesContext)

	// EnterAccessibleByClause is called when entering the accessibleByClause production.
	EnterAccessibleByClause(c *AccessibleByClauseContext)

	// EnterAccessor is called when entering the accessor production.
	EnterAccessor(c *AccessorContext)

	// EnterUnitKind is called when entering the unitKind production.
	EnterUnitKind(c *UnitKindContext)

	// EnterDefaultCollationClause is called when entering the defaultCollationClause production.
	EnterDefaultCollationClause(c *DefaultCollationClauseContext)

	// EnterCollationOption is called when entering the collationOption production.
	EnterCollationOption(c *CollationOptionContext)

	// EnterInvokerRightsClause is called when entering the invokerRightsClause production.
	EnterInvokerRightsClause(c *InvokerRightsClauseContext)

	// EnterSubqueryFactoringClause is called when entering the subqueryFactoringClause production.
	EnterSubqueryFactoringClause(c *SubqueryFactoringClauseContext)

	// EnterSearchClause is called when entering the searchClause production.
	EnterSearchClause(c *SearchClauseContext)

	// EnterCycleClause is called when entering the cycleClause production.
	EnterCycleClause(c *CycleClauseContext)

	// EnterSubavFactoringClause is called when entering the subavFactoringClause production.
	EnterSubavFactoringClause(c *SubavFactoringClauseContext)

	// EnterSubavClause is called when entering the subavClause production.
	EnterSubavClause(c *SubavClauseContext)

	// EnterHierarchiesClause is called when entering the hierarchiesClause production.
	EnterHierarchiesClause(c *HierarchiesClauseContext)

	// EnterFilterClauses is called when entering the filterClauses production.
	EnterFilterClauses(c *FilterClausesContext)

	// EnterFilterClause is called when entering the filterClause production.
	EnterFilterClause(c *FilterClauseContext)

	// EnterAddCalcsClause is called when entering the addCalcsClause production.
	EnterAddCalcsClause(c *AddCalcsClauseContext)

	// EnterCalcMeasClause is called when entering the calcMeasClause production.
	EnterCalcMeasClause(c *CalcMeasClauseContext)

	// EnterCalcMeasExpression is called when entering the calcMeasExpression production.
	EnterCalcMeasExpression(c *CalcMeasExpressionContext)

	// EnterAvExpression is called when entering the avExpression production.
	EnterAvExpression(c *AvExpressionContext)

	// EnterAvMeasExpression is called when entering the avMeasExpression production.
	EnterAvMeasExpression(c *AvMeasExpressionContext)

	// EnterLeadLagExpression is called when entering the leadLagExpression production.
	EnterLeadLagExpression(c *LeadLagExpressionContext)

	// EnterLeadLagFunctionName is called when entering the leadLagFunctionName production.
	EnterLeadLagFunctionName(c *LeadLagFunctionNameContext)

	// EnterLeadLagClause is called when entering the leadLagClause production.
	EnterLeadLagClause(c *LeadLagClauseContext)

	// EnterHierarchyRef is called when entering the hierarchyRef production.
	EnterHierarchyRef(c *HierarchyRefContext)

	// EnterWindowExpression is called when entering the windowExpression production.
	EnterWindowExpression(c *WindowExpressionContext)

	// EnterWindowClause is called when entering the windowClause production.
	EnterWindowClause(c *WindowClauseContext)

	// EnterPrecedingBoundary is called when entering the precedingBoundary production.
	EnterPrecedingBoundary(c *PrecedingBoundaryContext)

	// EnterFollowingBoundary is called when entering the followingBoundary production.
	EnterFollowingBoundary(c *FollowingBoundaryContext)

	// EnterRankExpression is called when entering the rankExpression production.
	EnterRankExpression(c *RankExpressionContext)

	// EnterRankFunctionName is called when entering the rankFunctionName production.
	EnterRankFunctionName(c *RankFunctionNameContext)

	// EnterRankClause is called when entering the rankClause production.
	EnterRankClause(c *RankClauseContext)

	// EnterCalcMeasOrderByClause is called when entering the calcMeasOrderByClause production.
	EnterCalcMeasOrderByClause(c *CalcMeasOrderByClauseContext)

	// EnterShareOfExpression is called when entering the shareOfExpression production.
	EnterShareOfExpression(c *ShareOfExpressionContext)

	// EnterShareClause is called when entering the shareClause production.
	EnterShareClause(c *ShareClauseContext)

	// EnterMemberExpression is called when entering the memberExpression production.
	EnterMemberExpression(c *MemberExpressionContext)

	// EnterLevelMemberLiteral is called when entering the levelMemberLiteral production.
	EnterLevelMemberLiteral(c *LevelMemberLiteralContext)

	// EnterPosMemberKeys is called when entering the posMemberKeys production.
	EnterPosMemberKeys(c *PosMemberKeysContext)

	// EnterNamedMemberKeys is called when entering the namedMemberKeys production.
	EnterNamedMemberKeys(c *NamedMemberKeysContext)

	// EnterHierNavigationExpression is called when entering the hierNavigationExpression production.
	EnterHierNavigationExpression(c *HierNavigationExpressionContext)

	// EnterHierAncestorExpression is called when entering the hierAncestorExpression production.
	EnterHierAncestorExpression(c *HierAncestorExpressionContext)

	// EnterHierParentExpression is called when entering the hierParentExpression production.
	EnterHierParentExpression(c *HierParentExpressionContext)

	// EnterHierLeadLagExpression is called when entering the hierLeadLagExpression production.
	EnterHierLeadLagExpression(c *HierLeadLagExpressionContext)

	// EnterHierLeadLagClause is called when entering the hierLeadLagClause production.
	EnterHierLeadLagClause(c *HierLeadLagClauseContext)

	// EnterQdrExpression is called when entering the qdrExpression production.
	EnterQdrExpression(c *QdrExpressionContext)

	// EnterQualifier is called when entering the qualifier production.
	EnterQualifier(c *QualifierContext)

	// EnterAvHierExpression is called when entering the avHierExpression production.
	EnterAvHierExpression(c *AvHierExpressionContext)

	// EnterHierFunctionName is called when entering the hierFunctionName production.
	EnterHierFunctionName(c *HierFunctionNameContext)

	// EnterDuplicateSpecification is called when entering the duplicateSpecification production.
	EnterDuplicateSpecification(c *DuplicateSpecificationContext)

	// EnterUnqualifiedShorthand is called when entering the unqualifiedShorthand production.
	EnterUnqualifiedShorthand(c *UnqualifiedShorthandContext)

	// EnterSelectList is called when entering the selectList production.
	EnterSelectList(c *SelectListContext)

	// EnterSelectProjection is called when entering the selectProjection production.
	EnterSelectProjection(c *SelectProjectionContext)

	// EnterSelectProjectionExprClause is called when entering the selectProjectionExprClause production.
	EnterSelectProjectionExprClause(c *SelectProjectionExprClauseContext)

	// EnterSelectFromClause is called when entering the selectFromClause production.
	EnterSelectFromClause(c *SelectFromClauseContext)

	// EnterFromClauseList is called when entering the fromClauseList production.
	EnterFromClauseList(c *FromClauseListContext)

	// EnterFromClauseOption is called when entering the fromClauseOption production.
	EnterFromClauseOption(c *FromClauseOptionContext)

	// EnterSelectTableReference is called when entering the selectTableReference production.
	EnterSelectTableReference(c *SelectTableReferenceContext)

	// EnterQueryTableExprClause is called when entering the queryTableExprClause production.
	EnterQueryTableExprClause(c *QueryTableExprClauseContext)

	// EnterFlashbackQueryClause is called when entering the flashbackQueryClause production.
	EnterFlashbackQueryClause(c *FlashbackQueryClauseContext)

	// EnterQueryTableExpr is called when entering the queryTableExpr production.
	EnterQueryTableExpr(c *QueryTableExprContext)

	// EnterLateralClause is called when entering the lateralClause production.
	EnterLateralClause(c *LateralClauseContext)

	// EnterQueryTableExprSampleClause is called when entering the queryTableExprSampleClause production.
	EnterQueryTableExprSampleClause(c *QueryTableExprSampleClauseContext)

	// EnterQueryTableExprTableClause is called when entering the queryTableExprTableClause production.
	EnterQueryTableExprTableClause(c *QueryTableExprTableClauseContext)

	// EnterQueryTableExprViewClause is called when entering the queryTableExprViewClause production.
	EnterQueryTableExprViewClause(c *QueryTableExprViewClauseContext)

	// EnterQueryTableExprAnalyticClause is called when entering the queryTableExprAnalyticClause production.
	EnterQueryTableExprAnalyticClause(c *QueryTableExprAnalyticClauseContext)

	// EnterInlineExternalTable is called when entering the inlineExternalTable production.
	EnterInlineExternalTable(c *InlineExternalTableContext)

	// EnterInlineExternalTableProperties is called when entering the inlineExternalTableProperties production.
	EnterInlineExternalTableProperties(c *InlineExternalTablePropertiesContext)

	// EnterExternalTableDataProperties is called when entering the externalTableDataProperties production.
	EnterExternalTableDataProperties(c *ExternalTableDataPropertiesContext)

	// EnterMofifiedExternalTable is called when entering the mofifiedExternalTable production.
	EnterMofifiedExternalTable(c *MofifiedExternalTableContext)

	// EnterModifyExternalTableProperties is called when entering the modifyExternalTableProperties production.
	EnterModifyExternalTableProperties(c *ModifyExternalTablePropertiesContext)

	// EnterPivotClause is called when entering the pivotClause production.
	EnterPivotClause(c *PivotClauseContext)

	// EnterPivotForClause is called when entering the pivotForClause production.
	EnterPivotForClause(c *PivotForClauseContext)

	// EnterPivotInClause is called when entering the pivotInClause production.
	EnterPivotInClause(c *PivotInClauseContext)

	// EnterUnpivotClause is called when entering the unpivotClause production.
	EnterUnpivotClause(c *UnpivotClauseContext)

	// EnterUnpivotInClause is called when entering the unpivotInClause production.
	EnterUnpivotInClause(c *UnpivotInClauseContext)

	// EnterSampleClause is called when entering the sampleClause production.
	EnterSampleClause(c *SampleClauseContext)

	// EnterContainersClause is called when entering the containersClause production.
	EnterContainersClause(c *ContainersClauseContext)

	// EnterShardsClause is called when entering the shardsClause production.
	EnterShardsClause(c *ShardsClauseContext)

	// EnterJoinClause is called when entering the joinClause production.
	EnterJoinClause(c *JoinClauseContext)

	// EnterSelectJoinOption is called when entering the selectJoinOption production.
	EnterSelectJoinOption(c *SelectJoinOptionContext)

	// EnterInnerCrossJoinClause is called when entering the innerCrossJoinClause production.
	EnterInnerCrossJoinClause(c *InnerCrossJoinClauseContext)

	// EnterSelectJoinSpecification is called when entering the selectJoinSpecification production.
	EnterSelectJoinSpecification(c *SelectJoinSpecificationContext)

	// EnterOuterJoinClause is called when entering the outerJoinClause production.
	EnterOuterJoinClause(c *OuterJoinClauseContext)

	// EnterQueryPartitionClause is called when entering the queryPartitionClause production.
	EnterQueryPartitionClause(c *QueryPartitionClauseContext)

	// EnterOuterJoinType is called when entering the outerJoinType production.
	EnterOuterJoinType(c *OuterJoinTypeContext)

	// EnterCrossOuterApplyClause is called when entering the crossOuterApplyClause production.
	EnterCrossOuterApplyClause(c *CrossOuterApplyClauseContext)

	// EnterInlineAnalyticView is called when entering the inlineAnalyticView production.
	EnterInlineAnalyticView(c *InlineAnalyticViewContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterHierarchicalQueryClause is called when entering the hierarchicalQueryClause production.
	EnterHierarchicalQueryClause(c *HierarchicalQueryClauseContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterGroupByItem is called when entering the groupByItem production.
	EnterGroupByItem(c *GroupByItemContext)

	// EnterRollupCubeClause is called when entering the rollupCubeClause production.
	EnterRollupCubeClause(c *RollupCubeClauseContext)

	// EnterGroupingSetsClause is called when entering the groupingSetsClause production.
	EnterGroupingSetsClause(c *GroupingSetsClauseContext)

	// EnterGroupingExprList is called when entering the groupingExprList production.
	EnterGroupingExprList(c *GroupingExprListContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterModelClause is called when entering the modelClause production.
	EnterModelClause(c *ModelClauseContext)

	// EnterCellReferenceOptions is called when entering the cellReferenceOptions production.
	EnterCellReferenceOptions(c *CellReferenceOptionsContext)

	// EnterReturnRowsClause is called when entering the returnRowsClause production.
	EnterReturnRowsClause(c *ReturnRowsClauseContext)

	// EnterReferenceModel is called when entering the referenceModel production.
	EnterReferenceModel(c *ReferenceModelContext)

	// EnterMainModel is called when entering the mainModel production.
	EnterMainModel(c *MainModelContext)

	// EnterModelColumnClauses is called when entering the modelColumnClauses production.
	EnterModelColumnClauses(c *ModelColumnClausesContext)

	// EnterModelRulesClause is called when entering the modelRulesClause production.
	EnterModelRulesClause(c *ModelRulesClauseContext)

	// EnterModelIterateClause is called when entering the modelIterateClause production.
	EnterModelIterateClause(c *ModelIterateClauseContext)

	// EnterCellAssignment is called when entering the cellAssignment production.
	EnterCellAssignment(c *CellAssignmentContext)

	// EnterSingleColumnForLoop is called when entering the singleColumnForLoop production.
	EnterSingleColumnForLoop(c *SingleColumnForLoopContext)

	// EnterMultiColumnForLoop is called when entering the multiColumnForLoop production.
	EnterMultiColumnForLoop(c *MultiColumnForLoopContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterModelExpr is called when entering the modelExpr production.
	EnterModelExpr(c *ModelExprContext)

	// EnterAnalyticFunction is called when entering the analyticFunction production.
	EnterAnalyticFunction(c *AnalyticFunctionContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterForUpdateClause is called when entering the forUpdateClause production.
	EnterForUpdateClause(c *ForUpdateClauseContext)

	// EnterForUpdateClauseList is called when entering the forUpdateClauseList production.
	EnterForUpdateClauseList(c *ForUpdateClauseListContext)

	// EnterForUpdateClauseOption is called when entering the forUpdateClauseOption production.
	EnterForUpdateClauseOption(c *ForUpdateClauseOptionContext)

	// EnterRowLimitingClause is called when entering the rowLimitingClause production.
	EnterRowLimitingClause(c *RowLimitingClauseContext)

	// EnterMerge is called when entering the merge production.
	EnterMerge(c *MergeContext)

	// EnterHint is called when entering the hint production.
	EnterHint(c *HintContext)

	// EnterIntoClause is called when entering the intoClause production.
	EnterIntoClause(c *IntoClauseContext)

	// EnterUsingClause is called when entering the usingClause production.
	EnterUsingClause(c *UsingClauseContext)

	// EnterMergeUpdateClause is called when entering the mergeUpdateClause production.
	EnterMergeUpdateClause(c *MergeUpdateClauseContext)

	// EnterMergeSetAssignmentsClause is called when entering the mergeSetAssignmentsClause production.
	EnterMergeSetAssignmentsClause(c *MergeSetAssignmentsClauseContext)

	// EnterMergeAssignment is called when entering the mergeAssignment production.
	EnterMergeAssignment(c *MergeAssignmentContext)

	// EnterMergeAssignmentValue is called when entering the mergeAssignmentValue production.
	EnterMergeAssignmentValue(c *MergeAssignmentValueContext)

	// EnterDeleteWhereClause is called when entering the deleteWhereClause production.
	EnterDeleteWhereClause(c *DeleteWhereClauseContext)

	// EnterMergeInsertClause is called when entering the mergeInsertClause production.
	EnterMergeInsertClause(c *MergeInsertClauseContext)

	// EnterMergeInsertColumn is called when entering the mergeInsertColumn production.
	EnterMergeInsertColumn(c *MergeInsertColumnContext)

	// EnterMergeColumnValue is called when entering the mergeColumnValue production.
	EnterMergeColumnValue(c *MergeColumnValueContext)

	// EnterErrorLoggingClause is called when entering the errorLoggingClause production.
	EnterErrorLoggingClause(c *ErrorLoggingClauseContext)

	// EnterRowPatternClause is called when entering the rowPatternClause production.
	EnterRowPatternClause(c *RowPatternClauseContext)

	// EnterRowPatternPartitionBy is called when entering the rowPatternPartitionBy production.
	EnterRowPatternPartitionBy(c *RowPatternPartitionByContext)

	// EnterRowPatternOrderBy is called when entering the rowPatternOrderBy production.
	EnterRowPatternOrderBy(c *RowPatternOrderByContext)

	// EnterRowPatternMeasures is called when entering the rowPatternMeasures production.
	EnterRowPatternMeasures(c *RowPatternMeasuresContext)

	// EnterRowPatternMeasureColumn is called when entering the rowPatternMeasureColumn production.
	EnterRowPatternMeasureColumn(c *RowPatternMeasureColumnContext)

	// EnterRowPatternRowsPerMatch is called when entering the rowPatternRowsPerMatch production.
	EnterRowPatternRowsPerMatch(c *RowPatternRowsPerMatchContext)

	// EnterRowPatternSkipTo is called when entering the rowPatternSkipTo production.
	EnterRowPatternSkipTo(c *RowPatternSkipToContext)

	// EnterRowPattern is called when entering the rowPattern production.
	EnterRowPattern(c *RowPatternContext)

	// EnterRowPatternTerm is called when entering the rowPatternTerm production.
	EnterRowPatternTerm(c *RowPatternTermContext)

	// EnterRowPatternFactor is called when entering the rowPatternFactor production.
	EnterRowPatternFactor(c *RowPatternFactorContext)

	// EnterRowPatternPrimary is called when entering the rowPatternPrimary production.
	EnterRowPatternPrimary(c *RowPatternPrimaryContext)

	// EnterRowPatternPermute is called when entering the rowPatternPermute production.
	EnterRowPatternPermute(c *RowPatternPermuteContext)

	// EnterRowPatternQuantifier is called when entering the rowPatternQuantifier production.
	EnterRowPatternQuantifier(c *RowPatternQuantifierContext)

	// EnterRowPatternSubsetClause is called when entering the rowPatternSubsetClause production.
	EnterRowPatternSubsetClause(c *RowPatternSubsetClauseContext)

	// EnterRowPatternSubsetItem is called when entering the rowPatternSubsetItem production.
	EnterRowPatternSubsetItem(c *RowPatternSubsetItemContext)

	// EnterRowPatternDefinitionList is called when entering the rowPatternDefinitionList production.
	EnterRowPatternDefinitionList(c *RowPatternDefinitionListContext)

	// EnterRowPatternDefinition is called when entering the rowPatternDefinition production.
	EnterRowPatternDefinition(c *RowPatternDefinitionContext)

	// EnterRowPatternRecFunc is called when entering the rowPatternRecFunc production.
	EnterRowPatternRecFunc(c *RowPatternRecFuncContext)

	// EnterPatternMeasExpression is called when entering the patternMeasExpression production.
	EnterPatternMeasExpression(c *PatternMeasExpressionContext)

	// EnterRowPatternClassifierFunc is called when entering the rowPatternClassifierFunc production.
	EnterRowPatternClassifierFunc(c *RowPatternClassifierFuncContext)

	// EnterRowPatternMatchNumFunc is called when entering the rowPatternMatchNumFunc production.
	EnterRowPatternMatchNumFunc(c *RowPatternMatchNumFuncContext)

	// EnterRowPatternNavigationFunc is called when entering the rowPatternNavigationFunc production.
	EnterRowPatternNavigationFunc(c *RowPatternNavigationFuncContext)

	// EnterRowPatternNavLogical is called when entering the rowPatternNavLogical production.
	EnterRowPatternNavLogical(c *RowPatternNavLogicalContext)

	// EnterRowPatternNavPhysical is called when entering the rowPatternNavPhysical production.
	EnterRowPatternNavPhysical(c *RowPatternNavPhysicalContext)

	// EnterRowPatternNavCompound is called when entering the rowPatternNavCompound production.
	EnterRowPatternNavCompound(c *RowPatternNavCompoundContext)

	// EnterRowPatternAggregateFunc is called when entering the rowPatternAggregateFunc production.
	EnterRowPatternAggregateFunc(c *RowPatternAggregateFuncContext)

	// EnterParameterMarker is called when entering the parameterMarker production.
	EnterParameterMarker(c *ParameterMarkerContext)

	// EnterLiterals is called when entering the literals production.
	EnterLiterals(c *LiteralsContext)

	// EnterStringLiterals is called when entering the stringLiterals production.
	EnterStringLiterals(c *StringLiteralsContext)

	// EnterNumberLiterals is called when entering the numberLiterals production.
	EnterNumberLiterals(c *NumberLiteralsContext)

	// EnterDateTimeLiterals is called when entering the dateTimeLiterals production.
	EnterDateTimeLiterals(c *DateTimeLiteralsContext)

	// EnterHexadecimalLiterals is called when entering the hexadecimalLiterals production.
	EnterHexadecimalLiterals(c *HexadecimalLiteralsContext)

	// EnterBitValueLiterals is called when entering the bitValueLiterals production.
	EnterBitValueLiterals(c *BitValueLiteralsContext)

	// EnterBooleanLiterals is called when entering the booleanLiterals production.
	EnterBooleanLiterals(c *BooleanLiteralsContext)

	// EnterNullValueLiterals is called when entering the nullValueLiterals production.
	EnterNullValueLiterals(c *NullValueLiteralsContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterUnreservedWord is called when entering the unreservedWord production.
	EnterUnreservedWord(c *UnreservedWordContext)

	// EnterSchemaName is called when entering the schemaName production.
	EnterSchemaName(c *SchemaNameContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterViewName is called when entering the viewName production.
	EnterViewName(c *ViewNameContext)

	// EnterMaterializedViewName is called when entering the materializedViewName production.
	EnterMaterializedViewName(c *MaterializedViewNameContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterObjectName is called when entering the objectName production.
	EnterObjectName(c *ObjectNameContext)

	// EnterClusterName is called when entering the clusterName production.
	EnterClusterName(c *ClusterNameContext)

	// EnterIndexName is called when entering the indexName production.
	EnterIndexName(c *IndexNameContext)

	// EnterStatisticsTypeName is called when entering the statisticsTypeName production.
	EnterStatisticsTypeName(c *StatisticsTypeNameContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterPackageName is called when entering the packageName production.
	EnterPackageName(c *PackageNameContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterIndextypeName is called when entering the indextypeName production.
	EnterIndextypeName(c *IndextypeNameContext)

	// EnterModelName is called when entering the modelName production.
	EnterModelName(c *ModelNameContext)

	// EnterOperatorName is called when entering the operatorName production.
	EnterOperatorName(c *OperatorNameContext)

	// EnterConstraintName is called when entering the constraintName production.
	EnterConstraintName(c *ConstraintNameContext)

	// EnterSavepointName is called when entering the savepointName production.
	EnterSavepointName(c *SavepointNameContext)

	// EnterSynonymName is called when entering the synonymName production.
	EnterSynonymName(c *SynonymNameContext)

	// EnterOwner is called when entering the owner production.
	EnterOwner(c *OwnerContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterTablespaceName is called when entering the tablespaceName production.
	EnterTablespaceName(c *TablespaceNameContext)

	// EnterTablespaceSetName is called when entering the tablespaceSetName production.
	EnterTablespaceSetName(c *TablespaceSetNameContext)

	// EnterServiceName is called when entering the serviceName production.
	EnterServiceName(c *ServiceNameContext)

	// EnterIlmPolicyName is called when entering the ilmPolicyName production.
	EnterIlmPolicyName(c *IlmPolicyNameContext)

	// EnterPolicyName is called when entering the policyName production.
	EnterPolicyName(c *PolicyNameContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterDbLink is called when entering the dbLink production.
	EnterDbLink(c *DbLinkContext)

	// EnterParameterValue is called when entering the parameterValue production.
	EnterParameterValue(c *ParameterValueContext)

	// EnterDirectoryName is called when entering the directoryName production.
	EnterDirectoryName(c *DirectoryNameContext)

	// EnterDispatcherName is called when entering the dispatcherName production.
	EnterDispatcherName(c *DispatcherNameContext)

	// EnterClientId is called when entering the clientId production.
	EnterClientId(c *ClientIdContext)

	// EnterOpaqueFormatSpec is called when entering the opaqueFormatSpec production.
	EnterOpaqueFormatSpec(c *OpaqueFormatSpecContext)

	// EnterAccessDriverType is called when entering the accessDriverType production.
	EnterAccessDriverType(c *AccessDriverTypeContext)

	// EnterVarrayItem is called when entering the varrayItem production.
	EnterVarrayItem(c *VarrayItemContext)

	// EnterNestedItem is called when entering the nestedItem production.
	EnterNestedItem(c *NestedItemContext)

	// EnterStorageTable is called when entering the storageTable production.
	EnterStorageTable(c *StorageTableContext)

	// EnterLobSegname is called when entering the lobSegname production.
	EnterLobSegname(c *LobSegnameContext)

	// EnterLocationSpecifier is called when entering the locationSpecifier production.
	EnterLocationSpecifier(c *LocationSpecifierContext)

	// EnterXmlSchemaURLName is called when entering the xmlSchemaURLName production.
	EnterXmlSchemaURLName(c *XmlSchemaURLNameContext)

	// EnterElementName is called when entering the elementName production.
	EnterElementName(c *ElementNameContext)

	// EnterSubpartitionName is called when entering the subpartitionName production.
	EnterSubpartitionName(c *SubpartitionNameContext)

	// EnterParameterName is called when entering the parameterName production.
	EnterParameterName(c *ParameterNameContext)

	// EnterEditionName is called when entering the editionName production.
	EnterEditionName(c *EditionNameContext)

	// EnterContainerName is called when entering the containerName production.
	EnterContainerName(c *ContainerNameContext)

	// EnterPartitionName is called when entering the partitionName production.
	EnterPartitionName(c *PartitionNameContext)

	// EnterPartitionSetName is called when entering the partitionSetName production.
	EnterPartitionSetName(c *PartitionSetNameContext)

	// EnterPartitionKeyValue is called when entering the partitionKeyValue production.
	EnterPartitionKeyValue(c *PartitionKeyValueContext)

	// EnterSubpartitionKeyValue is called when entering the subpartitionKeyValue production.
	EnterSubpartitionKeyValue(c *SubpartitionKeyValueContext)

	// EnterZonemapName is called when entering the zonemapName production.
	EnterZonemapName(c *ZonemapNameContext)

	// EnterFlashbackArchiveName is called when entering the flashbackArchiveName production.
	EnterFlashbackArchiveName(c *FlashbackArchiveNameContext)

	// EnterRoleName is called when entering the roleName production.
	EnterRoleName(c *RoleNameContext)

	// EnterUserName is called when entering the userName production.
	EnterUserName(c *UserNameContext)

	// EnterPassword is called when entering the password production.
	EnterPassword(c *PasswordContext)

	// EnterLogGroupName is called when entering the logGroupName production.
	EnterLogGroupName(c *LogGroupNameContext)

	// EnterColumnNames is called when entering the columnNames production.
	EnterColumnNames(c *ColumnNamesContext)

	// EnterTableNames is called when entering the tableNames production.
	EnterTableNames(c *TableNamesContext)

	// EnterOracleId is called when entering the oracleId production.
	EnterOracleId(c *OracleIdContext)

	// EnterCollationName is called when entering the collationName production.
	EnterCollationName(c *CollationNameContext)

	// EnterColumnCollationName is called when entering the columnCollationName production.
	EnterColumnCollationName(c *ColumnCollationNameContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterDataTypeLength is called when entering the dataTypeLength production.
	EnterDataTypeLength(c *DataTypeLengthContext)

	// EnterPrimaryKey is called when entering the primaryKey production.
	EnterPrimaryKey(c *PrimaryKeyContext)

	// EnterExprs is called when entering the exprs production.
	EnterExprs(c *ExprsContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterAndOperator is called when entering the andOperator production.
	EnterAndOperator(c *AndOperatorContext)

	// EnterOrOperator is called when entering the orOperator production.
	EnterOrOperator(c *OrOperatorContext)

	// EnterNotOperator is called when entering the notOperator production.
	EnterNotOperator(c *NotOperatorContext)

	// EnterBooleanPrimary is called when entering the booleanPrimary production.
	EnterBooleanPrimary(c *BooleanPrimaryContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterBitExpr is called when entering the bitExpr production.
	EnterBitExpr(c *BitExprContext)

	// EnterSimpleExpr is called when entering the simpleExpr production.
	EnterSimpleExpr(c *SimpleExprContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterAggregationFunction is called when entering the aggregationFunction production.
	EnterAggregationFunction(c *AggregationFunctionContext)

	// EnterAggregationFunctionName is called when entering the aggregationFunctionName production.
	EnterAggregationFunctionName(c *AggregationFunctionNameContext)

	// EnterAnalyticClause is called when entering the analyticClause production.
	EnterAnalyticClause(c *AnalyticClauseContext)

	// EnterWindowingClause is called when entering the windowingClause production.
	EnterWindowingClause(c *WindowingClauseContext)

	// EnterSpecialFunction is called when entering the specialFunction production.
	EnterSpecialFunction(c *SpecialFunctionContext)

	// EnterCastFunction is called when entering the castFunction production.
	EnterCastFunction(c *CastFunctionContext)

	// EnterCharFunction is called when entering the charFunction production.
	EnterCharFunction(c *CharFunctionContext)

	// EnterRegularFunction is called when entering the regularFunction production.
	EnterRegularFunction(c *RegularFunctionContext)

	// EnterRegularFunctionName is called when entering the regularFunctionName production.
	EnterRegularFunctionName(c *RegularFunctionNameContext)

	// EnterCaseExpression is called when entering the caseExpression production.
	EnterCaseExpression(c *CaseExpressionContext)

	// EnterCaseWhen is called when entering the caseWhen production.
	EnterCaseWhen(c *CaseWhenContext)

	// EnterCaseElse is called when entering the caseElse production.
	EnterCaseElse(c *CaseElseContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderByItem is called when entering the orderByItem production.
	EnterOrderByItem(c *OrderByItemContext)

	// EnterAttributeName is called when entering the attributeName production.
	EnterAttributeName(c *AttributeNameContext)

	// EnterSimpleExprs is called when entering the simpleExprs production.
	EnterSimpleExprs(c *SimpleExprsContext)

	// EnterLobItem is called when entering the lobItem production.
	EnterLobItem(c *LobItemContext)

	// EnterLobItems is called when entering the lobItems production.
	EnterLobItems(c *LobItemsContext)

	// EnterLobItemList is called when entering the lobItemList production.
	EnterLobItemList(c *LobItemListContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterSpecialDatatype is called when entering the specialDatatype production.
	EnterSpecialDatatype(c *SpecialDatatypeContext)

	// EnterDataTypeName is called when entering the dataTypeName production.
	EnterDataTypeName(c *DataTypeNameContext)

	// EnterDatetimeTypeSuffix is called when entering the datetimeTypeSuffix production.
	EnterDatetimeTypeSuffix(c *DatetimeTypeSuffixContext)

	// EnterTreatFunction is called when entering the treatFunction production.
	EnterTreatFunction(c *TreatFunctionContext)

	// EnterPrivateExprOfDb is called when entering the privateExprOfDb production.
	EnterPrivateExprOfDb(c *PrivateExprOfDbContext)

	// EnterCaseExpr is called when entering the caseExpr production.
	EnterCaseExpr(c *CaseExprContext)

	// EnterSimpleCaseExpr is called when entering the simpleCaseExpr production.
	EnterSimpleCaseExpr(c *SimpleCaseExprContext)

	// EnterSearchedCaseExpr is called when entering the searchedCaseExpr production.
	EnterSearchedCaseExpr(c *SearchedCaseExprContext)

	// EnterElseClause is called when entering the elseClause production.
	EnterElseClause(c *ElseClauseContext)

	// EnterIntervalExpression is called when entering the intervalExpression production.
	EnterIntervalExpression(c *IntervalExpressionContext)

	// EnterObjectAccessExpression is called when entering the objectAccessExpression production.
	EnterObjectAccessExpression(c *ObjectAccessExpressionContext)

	// EnterConstructorExpr is called when entering the constructorExpr production.
	EnterConstructorExpr(c *ConstructorExprContext)

	// EnterIgnoredIdentifier is called when entering the ignoredIdentifier production.
	EnterIgnoredIdentifier(c *IgnoredIdentifierContext)

	// EnterIgnoredIdentifiers is called when entering the ignoredIdentifiers production.
	EnterIgnoredIdentifiers(c *IgnoredIdentifiersContext)

	// EnterMatchNone is called when entering the matchNone production.
	EnterMatchNone(c *MatchNoneContext)

	// EnterHashSubpartitionQuantity is called when entering the hashSubpartitionQuantity production.
	EnterHashSubpartitionQuantity(c *HashSubpartitionQuantityContext)

	// EnterOdciParameters is called when entering the odciParameters production.
	EnterOdciParameters(c *OdciParametersContext)

	// EnterDatabaseName is called when entering the databaseName production.
	EnterDatabaseName(c *DatabaseNameContext)

	// EnterLocationName is called when entering the locationName production.
	EnterLocationName(c *LocationNameContext)

	// EnterFileName is called when entering the fileName production.
	EnterFileName(c *FileNameContext)

	// EnterAsmFileName is called when entering the asmFileName production.
	EnterAsmFileName(c *AsmFileNameContext)

	// EnterFileNumber is called when entering the fileNumber production.
	EnterFileNumber(c *FileNumberContext)

	// EnterInstanceName is called when entering the instanceName production.
	EnterInstanceName(c *InstanceNameContext)

	// EnterLogminerSessionName is called when entering the logminerSessionName production.
	EnterLogminerSessionName(c *LogminerSessionNameContext)

	// EnterTablespaceGroupName is called when entering the tablespaceGroupName production.
	EnterTablespaceGroupName(c *TablespaceGroupNameContext)

	// EnterCopyName is called when entering the copyName production.
	EnterCopyName(c *CopyNameContext)

	// EnterMirrorName is called when entering the mirrorName production.
	EnterMirrorName(c *MirrorNameContext)

	// EnterUriString is called when entering the uriString production.
	EnterUriString(c *UriStringContext)

	// EnterQualifiedCredentialName is called when entering the qualifiedCredentialName production.
	EnterQualifiedCredentialName(c *QualifiedCredentialNameContext)

	// EnterPdbName is called when entering the pdbName production.
	EnterPdbName(c *PdbNameContext)

	// EnterDiskgroupName is called when entering the diskgroupName production.
	EnterDiskgroupName(c *DiskgroupNameContext)

	// EnterTemplateName is called when entering the templateName production.
	EnterTemplateName(c *TemplateNameContext)

	// EnterAliasName is called when entering the aliasName production.
	EnterAliasName(c *AliasNameContext)

	// EnterDomain is called when entering the domain production.
	EnterDomain(c *DomainContext)

	// EnterDateValue is called when entering the dateValue production.
	EnterDateValue(c *DateValueContext)

	// EnterSessionId is called when entering the sessionId production.
	EnterSessionId(c *SessionIdContext)

	// EnterSerialNumber is called when entering the serialNumber production.
	EnterSerialNumber(c *SerialNumberContext)

	// EnterInstanceId is called when entering the instanceId production.
	EnterInstanceId(c *InstanceIdContext)

	// EnterSqlId is called when entering the sqlId production.
	EnterSqlId(c *SqlIdContext)

	// EnterLogFileName is called when entering the logFileName production.
	EnterLogFileName(c *LogFileNameContext)

	// EnterLogFileGroupsArchivedLocationName is called when entering the logFileGroupsArchivedLocationName production.
	EnterLogFileGroupsArchivedLocationName(c *LogFileGroupsArchivedLocationNameContext)

	// EnterAsmVersion is called when entering the asmVersion production.
	EnterAsmVersion(c *AsmVersionContext)

	// EnterWalletPassword is called when entering the walletPassword production.
	EnterWalletPassword(c *WalletPasswordContext)

	// EnterHsmAuthString is called when entering the hsmAuthString production.
	EnterHsmAuthString(c *HsmAuthStringContext)

	// EnterTargetDbName is called when entering the targetDbName production.
	EnterTargetDbName(c *TargetDbNameContext)

	// EnterCertificateId is called when entering the certificateId production.
	EnterCertificateId(c *CertificateIdContext)

	// EnterCategoryName is called when entering the categoryName production.
	EnterCategoryName(c *CategoryNameContext)

	// EnterOffset is called when entering the offset production.
	EnterOffset(c *OffsetContext)

	// EnterRowcount is called when entering the rowcount production.
	EnterRowcount(c *RowcountContext)

	// EnterPercent is called when entering the percent production.
	EnterPercent(c *PercentContext)

	// EnterRollbackSegment is called when entering the rollbackSegment production.
	EnterRollbackSegment(c *RollbackSegmentContext)

	// EnterQueryName is called when entering the queryName production.
	EnterQueryName(c *QueryNameContext)

	// EnterCycleValue is called when entering the cycleValue production.
	EnterCycleValue(c *CycleValueContext)

	// EnterNoCycleValue is called when entering the noCycleValue production.
	EnterNoCycleValue(c *NoCycleValueContext)

	// EnterOrderingColumn is called when entering the orderingColumn production.
	EnterOrderingColumn(c *OrderingColumnContext)

	// EnterSubavName is called when entering the subavName production.
	EnterSubavName(c *SubavNameContext)

	// EnterBaseAvName is called when entering the baseAvName production.
	EnterBaseAvName(c *BaseAvNameContext)

	// EnterMeasName is called when entering the measName production.
	EnterMeasName(c *MeasNameContext)

	// EnterLevelRef is called when entering the levelRef production.
	EnterLevelRef(c *LevelRefContext)

	// EnterOffsetExpr is called when entering the offsetExpr production.
	EnterOffsetExpr(c *OffsetExprContext)

	// EnterMemberKeyExpr is called when entering the memberKeyExpr production.
	EnterMemberKeyExpr(c *MemberKeyExprContext)

	// EnterDepthExpression is called when entering the depthExpression production.
	EnterDepthExpression(c *DepthExpressionContext)

	// EnterUnitName is called when entering the unitName production.
	EnterUnitName(c *UnitNameContext)

	// EnterProcedureName is called when entering the procedureName production.
	EnterProcedureName(c *ProcedureNameContext)

	// EnterCpuCost is called when entering the cpuCost production.
	EnterCpuCost(c *CpuCostContext)

	// EnterIoCost is called when entering the ioCost production.
	EnterIoCost(c *IoCostContext)

	// EnterNetworkCost is called when entering the networkCost production.
	EnterNetworkCost(c *NetworkCostContext)

	// EnterDefaultSelectivity is called when entering the defaultSelectivity production.
	EnterDefaultSelectivity(c *DefaultSelectivityContext)

	// EnterDataItem is called when entering the dataItem production.
	EnterDataItem(c *DataItemContext)

	// EnterVariableName is called when entering the variableName production.
	EnterVariableName(c *VariableNameContext)

	// EnterValidTimeColumn is called when entering the validTimeColumn production.
	EnterValidTimeColumn(c *ValidTimeColumnContext)

	// EnterAttrDim is called when entering the attrDim production.
	EnterAttrDim(c *AttrDimContext)

	// EnterHierarchyName is called when entering the hierarchyName production.
	EnterHierarchyName(c *HierarchyNameContext)

	// EnterAnalyticViewName is called when entering the analyticViewName production.
	EnterAnalyticViewName(c *AnalyticViewNameContext)

	// EnterSamplePercent is called when entering the samplePercent production.
	EnterSamplePercent(c *SamplePercentContext)

	// EnterSeedValue is called when entering the seedValue production.
	EnterSeedValue(c *SeedValueContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterRestorePoint is called when entering the restorePoint production.
	EnterRestorePoint(c *RestorePointContext)

	// EnterScnValue is called when entering the scnValue production.
	EnterScnValue(c *ScnValueContext)

	// EnterTimestampValue is called when entering the timestampValue production.
	EnterTimestampValue(c *TimestampValueContext)

	// EnterScnTimestampExpr is called when entering the scnTimestampExpr production.
	EnterScnTimestampExpr(c *ScnTimestampExprContext)

	// EnterReferenceModelName is called when entering the referenceModelName production.
	EnterReferenceModelName(c *ReferenceModelNameContext)

	// EnterMainModelName is called when entering the mainModelName production.
	EnterMainModelName(c *MainModelNameContext)

	// EnterMeasureColumn is called when entering the measureColumn production.
	EnterMeasureColumn(c *MeasureColumnContext)

	// EnterDimensionColumn is called when entering the dimensionColumn production.
	EnterDimensionColumn(c *DimensionColumnContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterAnalyticFunctionName is called when entering the analyticFunctionName production.
	EnterAnalyticFunctionName(c *AnalyticFunctionNameContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterComparisonCondition is called when entering the comparisonCondition production.
	EnterComparisonCondition(c *ComparisonConditionContext)

	// EnterSimpleComparisonCondition is called when entering the simpleComparisonCondition production.
	EnterSimpleComparisonCondition(c *SimpleComparisonConditionContext)

	// EnterGroupComparisonCondition is called when entering the groupComparisonCondition production.
	EnterGroupComparisonCondition(c *GroupComparisonConditionContext)

	// EnterFloatingPointCondition is called when entering the floatingPointCondition production.
	EnterFloatingPointCondition(c *FloatingPointConditionContext)

	// EnterLogicalCondition is called when entering the logicalCondition production.
	EnterLogicalCondition(c *LogicalConditionContext)

	// EnterModelCondition is called when entering the modelCondition production.
	EnterModelCondition(c *ModelConditionContext)

	// EnterIsAnyCondition is called when entering the isAnyCondition production.
	EnterIsAnyCondition(c *IsAnyConditionContext)

	// EnterIsPresentCondition is called when entering the isPresentCondition production.
	EnterIsPresentCondition(c *IsPresentConditionContext)

	// EnterCellReference is called when entering the cellReference production.
	EnterCellReference(c *CellReferenceContext)

	// EnterMultisetCondition is called when entering the multisetCondition production.
	EnterMultisetCondition(c *MultisetConditionContext)

	// EnterIsASetCondition is called when entering the isASetCondition production.
	EnterIsASetCondition(c *IsASetConditionContext)

	// EnterIsEmptyCondition is called when entering the isEmptyCondition production.
	EnterIsEmptyCondition(c *IsEmptyConditionContext)

	// EnterMemberCondition is called when entering the memberCondition production.
	EnterMemberCondition(c *MemberConditionContext)

	// EnterSubmultisetCondition is called when entering the submultisetCondition production.
	EnterSubmultisetCondition(c *SubmultisetConditionContext)

	// EnterPatternMatchingCondition is called when entering the patternMatchingCondition production.
	EnterPatternMatchingCondition(c *PatternMatchingConditionContext)

	// EnterLikeCondition is called when entering the likeCondition production.
	EnterLikeCondition(c *LikeConditionContext)

	// EnterSearchValue is called when entering the searchValue production.
	EnterSearchValue(c *SearchValueContext)

	// EnterEscapeChar is called when entering the escapeChar production.
	EnterEscapeChar(c *EscapeCharContext)

	// EnterRegexpLikeCondition is called when entering the regexpLikeCondition production.
	EnterRegexpLikeCondition(c *RegexpLikeConditionContext)

	// EnterMatchParam is called when entering the matchParam production.
	EnterMatchParam(c *MatchParamContext)

	// EnterRangeCondition is called when entering the rangeCondition production.
	EnterRangeCondition(c *RangeConditionContext)

	// EnterNullCondition is called when entering the nullCondition production.
	EnterNullCondition(c *NullConditionContext)

	// EnterXmlCondition is called when entering the xmlCondition production.
	EnterXmlCondition(c *XmlConditionContext)

	// EnterEqualsPathCondition is called when entering the equalsPathCondition production.
	EnterEqualsPathCondition(c *EqualsPathConditionContext)

	// EnterPathString is called when entering the pathString production.
	EnterPathString(c *PathStringContext)

	// EnterCorrelationInteger is called when entering the correlationInteger production.
	EnterCorrelationInteger(c *CorrelationIntegerContext)

	// EnterUnderPathCondition is called when entering the underPathCondition production.
	EnterUnderPathCondition(c *UnderPathConditionContext)

	// EnterLevels is called when entering the levels production.
	EnterLevels(c *LevelsContext)

	// EnterJsonCondition is called when entering the jsonCondition production.
	EnterJsonCondition(c *JsonConditionContext)

	// EnterIsJsonCondition is called when entering the isJsonCondition production.
	EnterIsJsonCondition(c *IsJsonConditionContext)

	// EnterJsonEqualCondition is called when entering the jsonEqualCondition production.
	EnterJsonEqualCondition(c *JsonEqualConditionContext)

	// EnterJsonExistsCondition is called when entering the jsonExistsCondition production.
	EnterJsonExistsCondition(c *JsonExistsConditionContext)

	// EnterJsonPassingClause is called when entering the jsonPassingClause production.
	EnterJsonPassingClause(c *JsonPassingClauseContext)

	// EnterJsonExistsOnErrorClause is called when entering the jsonExistsOnErrorClause production.
	EnterJsonExistsOnErrorClause(c *JsonExistsOnErrorClauseContext)

	// EnterJsonExistsOnEmptyClause is called when entering the jsonExistsOnEmptyClause production.
	EnterJsonExistsOnEmptyClause(c *JsonExistsOnEmptyClauseContext)

	// EnterJsonTextcontainsCondition is called when entering the jsonTextcontainsCondition production.
	EnterJsonTextcontainsCondition(c *JsonTextcontainsConditionContext)

	// EnterJsonBasicPathExpr is called when entering the jsonBasicPathExpr production.
	EnterJsonBasicPathExpr(c *JsonBasicPathExprContext)

	// EnterJsonAbsolutePathExpr is called when entering the jsonAbsolutePathExpr production.
	EnterJsonAbsolutePathExpr(c *JsonAbsolutePathExprContext)

	// EnterJsonNonfunctionSteps is called when entering the jsonNonfunctionSteps production.
	EnterJsonNonfunctionSteps(c *JsonNonfunctionStepsContext)

	// EnterJsonObjectStep is called when entering the jsonObjectStep production.
	EnterJsonObjectStep(c *JsonObjectStepContext)

	// EnterJsonFieldName is called when entering the jsonFieldName production.
	EnterJsonFieldName(c *JsonFieldNameContext)

	// EnterLetter is called when entering the letter production.
	EnterLetter(c *LetterContext)

	// EnterDigit is called when entering the digit production.
	EnterDigit(c *DigitContext)

	// EnterJsonArrayStep is called when entering the jsonArrayStep production.
	EnterJsonArrayStep(c *JsonArrayStepContext)

	// EnterJsonDescendentStep is called when entering the jsonDescendentStep production.
	EnterJsonDescendentStep(c *JsonDescendentStepContext)

	// EnterJsonFunctionStep is called when entering the jsonFunctionStep production.
	EnterJsonFunctionStep(c *JsonFunctionStepContext)

	// EnterJsonItemMethod is called when entering the jsonItemMethod production.
	EnterJsonItemMethod(c *JsonItemMethodContext)

	// EnterJsonFilterExpr is called when entering the jsonFilterExpr production.
	EnterJsonFilterExpr(c *JsonFilterExprContext)

	// EnterJsonCond is called when entering the jsonCond production.
	EnterJsonCond(c *JsonCondContext)

	// EnterJsonDisjunction is called when entering the jsonDisjunction production.
	EnterJsonDisjunction(c *JsonDisjunctionContext)

	// EnterJsonConjunction is called when entering the jsonConjunction production.
	EnterJsonConjunction(c *JsonConjunctionContext)

	// EnterJsonNegation is called when entering the jsonNegation production.
	EnterJsonNegation(c *JsonNegationContext)

	// EnterJsonExistsCond is called when entering the jsonExistsCond production.
	EnterJsonExistsCond(c *JsonExistsCondContext)

	// EnterJsonHasSubstringCond is called when entering the jsonHasSubstringCond production.
	EnterJsonHasSubstringCond(c *JsonHasSubstringCondContext)

	// EnterJsonStartsWithCond is called when entering the jsonStartsWithCond production.
	EnterJsonStartsWithCond(c *JsonStartsWithCondContext)

	// EnterJsonLikeCond is called when entering the jsonLikeCond production.
	EnterJsonLikeCond(c *JsonLikeCondContext)

	// EnterJsonLikeRegexCond is called when entering the jsonLikeRegexCond production.
	EnterJsonLikeRegexCond(c *JsonLikeRegexCondContext)

	// EnterJsonEqRegexCond is called when entering the jsonEqRegexCond production.
	EnterJsonEqRegexCond(c *JsonEqRegexCondContext)

	// EnterJsonInCond is called when entering the jsonInCond production.
	EnterJsonInCond(c *JsonInCondContext)

	// EnterValueList is called when entering the valueList production.
	EnterValueList(c *ValueListContext)

	// EnterJsonComparison is called when entering the jsonComparison production.
	EnterJsonComparison(c *JsonComparisonContext)

	// EnterJsonRelativePathExpr is called when entering the jsonRelativePathExpr production.
	EnterJsonRelativePathExpr(c *JsonRelativePathExprContext)

	// EnterJsonComparePred is called when entering the jsonComparePred production.
	EnterJsonComparePred(c *JsonComparePredContext)

	// EnterJsonVar is called when entering the jsonVar production.
	EnterJsonVar(c *JsonVarContext)

	// EnterJsonScalar is called when entering the jsonScalar production.
	EnterJsonScalar(c *JsonScalarContext)

	// EnterJsonNumber is called when entering the jsonNumber production.
	EnterJsonNumber(c *JsonNumberContext)

	// EnterJsonString is called when entering the jsonString production.
	EnterJsonString(c *JsonStringContext)

	// EnterCompoundCondition is called when entering the compoundCondition production.
	EnterCompoundCondition(c *CompoundConditionContext)

	// EnterExistsCondition is called when entering the existsCondition production.
	EnterExistsCondition(c *ExistsConditionContext)

	// EnterInCondition is called when entering the inCondition production.
	EnterInCondition(c *InConditionContext)

	// EnterIsOfTypeCondition is called when entering the isOfTypeCondition production.
	EnterIsOfTypeCondition(c *IsOfTypeConditionContext)

	// EnterDatabaseCharset is called when entering the databaseCharset production.
	EnterDatabaseCharset(c *DatabaseCharsetContext)

	// EnterNationalCharset is called when entering the nationalCharset production.
	EnterNationalCharset(c *NationalCharsetContext)

	// EnterFilenamePattern is called when entering the filenamePattern production.
	EnterFilenamePattern(c *FilenamePatternContext)

	// EnterConnectString is called when entering the connectString production.
	EnterConnectString(c *ConnectStringContext)

	// EnterCreateTable is called when entering the createTable production.
	EnterCreateTable(c *CreateTableContext)

	// EnterCreateIndex is called when entering the createIndex production.
	EnterCreateIndex(c *CreateIndexContext)

	// EnterAlterTable is called when entering the alterTable production.
	EnterAlterTable(c *AlterTableContext)

	// EnterAlterIndex is called when entering the alterIndex production.
	EnterAlterIndex(c *AlterIndexContext)

	// EnterDropTable is called when entering the dropTable production.
	EnterDropTable(c *DropTableContext)

	// EnterDropIndex is called when entering the dropIndex production.
	EnterDropIndex(c *DropIndexContext)

	// EnterTruncateTable is called when entering the truncateTable production.
	EnterTruncateTable(c *TruncateTableContext)

	// EnterCreateTableSpecification is called when entering the createTableSpecification production.
	EnterCreateTableSpecification(c *CreateTableSpecificationContext)

	// EnterTablespaceClauseWithParen is called when entering the tablespaceClauseWithParen production.
	EnterTablespaceClauseWithParen(c *TablespaceClauseWithParenContext)

	// EnterTablespaceClause is called when entering the tablespaceClause production.
	EnterTablespaceClause(c *TablespaceClauseContext)

	// EnterCreateSharingClause is called when entering the createSharingClause production.
	EnterCreateSharingClause(c *CreateSharingClauseContext)

	// EnterCreateDefinitionClause is called when entering the createDefinitionClause production.
	EnterCreateDefinitionClause(c *CreateDefinitionClauseContext)

	// EnterCreateXMLTypeTableClause is called when entering the createXMLTypeTableClause production.
	EnterCreateXMLTypeTableClause(c *CreateXMLTypeTableClauseContext)

	// EnterXmlTypeStorageClause is called when entering the xmlTypeStorageClause production.
	EnterXmlTypeStorageClause(c *XmlTypeStorageClauseContext)

	// EnterXmlSchemaSpecClause is called when entering the xmlSchemaSpecClause production.
	EnterXmlSchemaSpecClause(c *XmlSchemaSpecClauseContext)

	// EnterXmlTypeVirtualColumnsClause is called when entering the xmlTypeVirtualColumnsClause production.
	EnterXmlTypeVirtualColumnsClause(c *XmlTypeVirtualColumnsClauseContext)

	// EnterOidClause is called when entering the oidClause production.
	EnterOidClause(c *OidClauseContext)

	// EnterOidIndexClause is called when entering the oidIndexClause production.
	EnterOidIndexClause(c *OidIndexClauseContext)

	// EnterCreateRelationalTableClause is called when entering the createRelationalTableClause production.
	EnterCreateRelationalTableClause(c *CreateRelationalTableClauseContext)

	// EnterCreateMemOptimizeClause is called when entering the createMemOptimizeClause production.
	EnterCreateMemOptimizeClause(c *CreateMemOptimizeClauseContext)

	// EnterCreateParentClause is called when entering the createParentClause production.
	EnterCreateParentClause(c *CreateParentClauseContext)

	// EnterCreateObjectTableClause is called when entering the createObjectTableClause production.
	EnterCreateObjectTableClause(c *CreateObjectTableClauseContext)

	// EnterRelationalProperties is called when entering the relationalProperties production.
	EnterRelationalProperties(c *RelationalPropertiesContext)

	// EnterRelationalProperty is called when entering the relationalProperty production.
	EnterRelationalProperty(c *RelationalPropertyContext)

	// EnterColumnDefinition is called when entering the columnDefinition production.
	EnterColumnDefinition(c *ColumnDefinitionContext)

	// EnterVisibleClause is called when entering the visibleClause production.
	EnterVisibleClause(c *VisibleClauseContext)

	// EnterDefaultNullClause is called when entering the defaultNullClause production.
	EnterDefaultNullClause(c *DefaultNullClauseContext)

	// EnterIdentityClause is called when entering the identityClause production.
	EnterIdentityClause(c *IdentityClauseContext)

	// EnterIdentifyOptions is called when entering the identifyOptions production.
	EnterIdentifyOptions(c *IdentifyOptionsContext)

	// EnterIdentityOption is called when entering the identityOption production.
	EnterIdentityOption(c *IdentityOptionContext)

	// EnterEncryptionSpecification is called when entering the encryptionSpecification production.
	EnterEncryptionSpecification(c *EncryptionSpecificationContext)

	// EnterInlineConstraint is called when entering the inlineConstraint production.
	EnterInlineConstraint(c *InlineConstraintContext)

	// EnterReferencesClause is called when entering the referencesClause production.
	EnterReferencesClause(c *ReferencesClauseContext)

	// EnterConstraintState is called when entering the constraintState production.
	EnterConstraintState(c *ConstraintStateContext)

	// EnterNotDeferrable is called when entering the notDeferrable production.
	EnterNotDeferrable(c *NotDeferrableContext)

	// EnterInitiallyClause is called when entering the initiallyClause production.
	EnterInitiallyClause(c *InitiallyClauseContext)

	// EnterExceptionsClause is called when entering the exceptionsClause production.
	EnterExceptionsClause(c *ExceptionsClauseContext)

	// EnterUsingIndexClause is called when entering the usingIndexClause production.
	EnterUsingIndexClause(c *UsingIndexClauseContext)

	// EnterCreateIndexClause is called when entering the createIndexClause production.
	EnterCreateIndexClause(c *CreateIndexClauseContext)

	// EnterInlineRefConstraint is called when entering the inlineRefConstraint production.
	EnterInlineRefConstraint(c *InlineRefConstraintContext)

	// EnterVirtualColumnDefinition is called when entering the virtualColumnDefinition production.
	EnterVirtualColumnDefinition(c *VirtualColumnDefinitionContext)

	// EnterOutOfLineConstraint is called when entering the outOfLineConstraint production.
	EnterOutOfLineConstraint(c *OutOfLineConstraintContext)

	// EnterOutOfLineRefConstraint is called when entering the outOfLineRefConstraint production.
	EnterOutOfLineRefConstraint(c *OutOfLineRefConstraintContext)

	// EnterCreateIndexSpecification is called when entering the createIndexSpecification production.
	EnterCreateIndexSpecification(c *CreateIndexSpecificationContext)

	// EnterClusterIndexClause is called when entering the clusterIndexClause production.
	EnterClusterIndexClause(c *ClusterIndexClauseContext)

	// EnterIndexAttributes is called when entering the indexAttributes production.
	EnterIndexAttributes(c *IndexAttributesContext)

	// EnterTableIndexClause is called when entering the tableIndexClause production.
	EnterTableIndexClause(c *TableIndexClauseContext)

	// EnterIndexExpressions is called when entering the indexExpressions production.
	EnterIndexExpressions(c *IndexExpressionsContext)

	// EnterIndexExpression is called when entering the indexExpression production.
	EnterIndexExpression(c *IndexExpressionContext)

	// EnterBitmapJoinIndexClause is called when entering the bitmapJoinIndexClause production.
	EnterBitmapJoinIndexClause(c *BitmapJoinIndexClauseContext)

	// EnterColumnSortsClause_ is called when entering the columnSortsClause_ production.
	EnterColumnSortsClause_(c *ColumnSortsClause_Context)

	// EnterColumnSortClause_ is called when entering the columnSortClause_ production.
	EnterColumnSortClause_(c *ColumnSortClause_Context)

	// EnterCreateIndexDefinitionClause is called when entering the createIndexDefinitionClause production.
	EnterCreateIndexDefinitionClause(c *CreateIndexDefinitionClauseContext)

	// EnterTableAlias is called when entering the tableAlias production.
	EnterTableAlias(c *TableAliasContext)

	// EnterAlterDefinitionClause is called when entering the alterDefinitionClause production.
	EnterAlterDefinitionClause(c *AlterDefinitionClauseContext)

	// EnterAlterTableProperties is called when entering the alterTableProperties production.
	EnterAlterTableProperties(c *AlterTablePropertiesContext)

	// EnterRenameTableSpecification is called when entering the renameTableSpecification production.
	EnterRenameTableSpecification(c *RenameTableSpecificationContext)

	// EnterColumnClauses is called when entering the columnClauses production.
	EnterColumnClauses(c *ColumnClausesContext)

	// EnterOperateColumnClause is called when entering the operateColumnClause production.
	EnterOperateColumnClause(c *OperateColumnClauseContext)

	// EnterAddColumnSpecification is called when entering the addColumnSpecification production.
	EnterAddColumnSpecification(c *AddColumnSpecificationContext)

	// EnterColumnOrVirtualDefinitions is called when entering the columnOrVirtualDefinitions production.
	EnterColumnOrVirtualDefinitions(c *ColumnOrVirtualDefinitionsContext)

	// EnterColumnOrVirtualDefinition is called when entering the columnOrVirtualDefinition production.
	EnterColumnOrVirtualDefinition(c *ColumnOrVirtualDefinitionContext)

	// EnterColumnProperties is called when entering the columnProperties production.
	EnterColumnProperties(c *ColumnPropertiesContext)

	// EnterColumnProperty is called when entering the columnProperty production.
	EnterColumnProperty(c *ColumnPropertyContext)

	// EnterObjectTypeColProperties is called when entering the objectTypeColProperties production.
	EnterObjectTypeColProperties(c *ObjectTypeColPropertiesContext)

	// EnterSubstitutableColumnClause is called when entering the substitutableColumnClause production.
	EnterSubstitutableColumnClause(c *SubstitutableColumnClauseContext)

	// EnterModifyColumnSpecification is called when entering the modifyColumnSpecification production.
	EnterModifyColumnSpecification(c *ModifyColumnSpecificationContext)

	// EnterModifyColProperties is called when entering the modifyColProperties production.
	EnterModifyColProperties(c *ModifyColPropertiesContext)

	// EnterModifyColSubstitutable is called when entering the modifyColSubstitutable production.
	EnterModifyColSubstitutable(c *ModifyColSubstitutableContext)

	// EnterDropColumnClause is called when entering the dropColumnClause production.
	EnterDropColumnClause(c *DropColumnClauseContext)

	// EnterDropColumnSpecification is called when entering the dropColumnSpecification production.
	EnterDropColumnSpecification(c *DropColumnSpecificationContext)

	// EnterColumnOrColumnList is called when entering the columnOrColumnList production.
	EnterColumnOrColumnList(c *ColumnOrColumnListContext)

	// EnterCascadeOrInvalidate is called when entering the cascadeOrInvalidate production.
	EnterCascadeOrInvalidate(c *CascadeOrInvalidateContext)

	// EnterCheckpointNumber is called when entering the checkpointNumber production.
	EnterCheckpointNumber(c *CheckpointNumberContext)

	// EnterRenameColumnClause is called when entering the renameColumnClause production.
	EnterRenameColumnClause(c *RenameColumnClauseContext)

	// EnterConstraintClauses is called when entering the constraintClauses production.
	EnterConstraintClauses(c *ConstraintClausesContext)

	// EnterAddConstraintSpecification is called when entering the addConstraintSpecification production.
	EnterAddConstraintSpecification(c *AddConstraintSpecificationContext)

	// EnterModifyConstraintClause is called when entering the modifyConstraintClause production.
	EnterModifyConstraintClause(c *ModifyConstraintClauseContext)

	// EnterConstraintWithName is called when entering the constraintWithName production.
	EnterConstraintWithName(c *ConstraintWithNameContext)

	// EnterConstraintOption is called when entering the constraintOption production.
	EnterConstraintOption(c *ConstraintOptionContext)

	// EnterConstraintPrimaryOrUnique is called when entering the constraintPrimaryOrUnique production.
	EnterConstraintPrimaryOrUnique(c *ConstraintPrimaryOrUniqueContext)

	// EnterRenameConstraintClause is called when entering the renameConstraintClause production.
	EnterRenameConstraintClause(c *RenameConstraintClauseContext)

	// EnterDropConstraintClause is called when entering the dropConstraintClause production.
	EnterDropConstraintClause(c *DropConstraintClauseContext)

	// EnterAlterExternalTable is called when entering the alterExternalTable production.
	EnterAlterExternalTable(c *AlterExternalTableContext)

	// EnterObjectProperties is called when entering the objectProperties production.
	EnterObjectProperties(c *ObjectPropertiesContext)

	// EnterAlterIndexInformationClause is called when entering the alterIndexInformationClause production.
	EnterAlterIndexInformationClause(c *AlterIndexInformationClauseContext)

	// EnterRenameIndexClause is called when entering the renameIndexClause production.
	EnterRenameIndexClause(c *RenameIndexClauseContext)

	// EnterObjectTableSubstitution is called when entering the objectTableSubstitution production.
	EnterObjectTableSubstitution(c *ObjectTableSubstitutionContext)

	// EnterMemOptimizeClause is called when entering the memOptimizeClause production.
	EnterMemOptimizeClause(c *MemOptimizeClauseContext)

	// EnterMemOptimizeReadClause is called when entering the memOptimizeReadClause production.
	EnterMemOptimizeReadClause(c *MemOptimizeReadClauseContext)

	// EnterMemOptimizeWriteClause is called when entering the memOptimizeWriteClause production.
	EnterMemOptimizeWriteClause(c *MemOptimizeWriteClauseContext)

	// EnterEnableDisableClauses is called when entering the enableDisableClauses production.
	EnterEnableDisableClauses(c *EnableDisableClausesContext)

	// EnterEnableDisableClause is called when entering the enableDisableClause production.
	EnterEnableDisableClause(c *EnableDisableClauseContext)

	// EnterEnableDisableOthers is called when entering the enableDisableOthers production.
	EnterEnableDisableOthers(c *EnableDisableOthersContext)

	// EnterRebuildClause is called when entering the rebuildClause production.
	EnterRebuildClause(c *RebuildClauseContext)

	// EnterParallelClause is called when entering the parallelClause production.
	EnterParallelClause(c *ParallelClauseContext)

	// EnterUsableSpecification is called when entering the usableSpecification production.
	EnterUsableSpecification(c *UsableSpecificationContext)

	// EnterInvalidationSpecification is called when entering the invalidationSpecification production.
	EnterInvalidationSpecification(c *InvalidationSpecificationContext)

	// EnterMaterializedViewLogClause is called when entering the materializedViewLogClause production.
	EnterMaterializedViewLogClause(c *MaterializedViewLogClauseContext)

	// EnterDropReuseClause is called when entering the dropReuseClause production.
	EnterDropReuseClause(c *DropReuseClauseContext)

	// EnterCollationClause is called when entering the collationClause production.
	EnterCollationClause(c *CollationClauseContext)

	// EnterCommitClause is called when entering the commitClause production.
	EnterCommitClause(c *CommitClauseContext)

	// EnterPhysicalProperties is called when entering the physicalProperties production.
	EnterPhysicalProperties(c *PhysicalPropertiesContext)

	// EnterDeferredSegmentCreation is called when entering the deferredSegmentCreation production.
	EnterDeferredSegmentCreation(c *DeferredSegmentCreationContext)

	// EnterSegmentAttributesClause is called when entering the segmentAttributesClause production.
	EnterSegmentAttributesClause(c *SegmentAttributesClauseContext)

	// EnterPhysicalAttributesClause is called when entering the physicalAttributesClause production.
	EnterPhysicalAttributesClause(c *PhysicalAttributesClauseContext)

	// EnterLoggingClause is called when entering the loggingClause production.
	EnterLoggingClause(c *LoggingClauseContext)

	// EnterStorageClause is called when entering the storageClause production.
	EnterStorageClause(c *StorageClauseContext)

	// EnterSizeClause is called when entering the sizeClause production.
	EnterSizeClause(c *SizeClauseContext)

	// EnterMaxsizeClause is called when entering the maxsizeClause production.
	EnterMaxsizeClause(c *MaxsizeClauseContext)

	// EnterTableCompression is called when entering the tableCompression production.
	EnterTableCompression(c *TableCompressionContext)

	// EnterInmemoryTableClause is called when entering the inmemoryTableClause production.
	EnterInmemoryTableClause(c *InmemoryTableClauseContext)

	// EnterInmemoryAttributes is called when entering the inmemoryAttributes production.
	EnterInmemoryAttributes(c *InmemoryAttributesContext)

	// EnterInmemoryColumnClause is called when entering the inmemoryColumnClause production.
	EnterInmemoryColumnClause(c *InmemoryColumnClauseContext)

	// EnterInmemoryMemcompress is called when entering the inmemoryMemcompress production.
	EnterInmemoryMemcompress(c *InmemoryMemcompressContext)

	// EnterInmemoryPriority is called when entering the inmemoryPriority production.
	EnterInmemoryPriority(c *InmemoryPriorityContext)

	// EnterInmemoryDistribute is called when entering the inmemoryDistribute production.
	EnterInmemoryDistribute(c *InmemoryDistributeContext)

	// EnterInmemoryDuplicate is called when entering the inmemoryDuplicate production.
	EnterInmemoryDuplicate(c *InmemoryDuplicateContext)

	// EnterIlmClause is called when entering the ilmClause production.
	EnterIlmClause(c *IlmClauseContext)

	// EnterIlmPolicyClause is called when entering the ilmPolicyClause production.
	EnterIlmPolicyClause(c *IlmPolicyClauseContext)

	// EnterIlmCompressionPolicy is called when entering the ilmCompressionPolicy production.
	EnterIlmCompressionPolicy(c *IlmCompressionPolicyContext)

	// EnterIlmTimePeriod is called when entering the ilmTimePeriod production.
	EnterIlmTimePeriod(c *IlmTimePeriodContext)

	// EnterIlmTieringPolicy is called when entering the ilmTieringPolicy production.
	EnterIlmTieringPolicy(c *IlmTieringPolicyContext)

	// EnterIlmInmemoryPolicy is called when entering the ilmInmemoryPolicy production.
	EnterIlmInmemoryPolicy(c *IlmInmemoryPolicyContext)

	// EnterOrganizationClause is called when entering the organizationClause production.
	EnterOrganizationClause(c *OrganizationClauseContext)

	// EnterHeapOrgTableClause is called when entering the heapOrgTableClause production.
	EnterHeapOrgTableClause(c *HeapOrgTableClauseContext)

	// EnterIndexOrgTableClause is called when entering the indexOrgTableClause production.
	EnterIndexOrgTableClause(c *IndexOrgTableClauseContext)

	// EnterExternalTableClause is called when entering the externalTableClause production.
	EnterExternalTableClause(c *ExternalTableClauseContext)

	// EnterExternalTableDataProps is called when entering the externalTableDataProps production.
	EnterExternalTableDataProps(c *ExternalTableDataPropsContext)

	// EnterMappingTableClause is called when entering the mappingTableClause production.
	EnterMappingTableClause(c *MappingTableClauseContext)

	// EnterPrefixCompression is called when entering the prefixCompression production.
	EnterPrefixCompression(c *PrefixCompressionContext)

	// EnterIndexOrgOverflowClause is called when entering the indexOrgOverflowClause production.
	EnterIndexOrgOverflowClause(c *IndexOrgOverflowClauseContext)

	// EnterExternalPartitionClause is called when entering the externalPartitionClause production.
	EnterExternalPartitionClause(c *ExternalPartitionClauseContext)

	// EnterClusterRelatedClause is called when entering the clusterRelatedClause production.
	EnterClusterRelatedClause(c *ClusterRelatedClauseContext)

	// EnterTableProperties is called when entering the tableProperties production.
	EnterTableProperties(c *TablePropertiesContext)

	// EnterReadOnlyClause is called when entering the readOnlyClause production.
	EnterReadOnlyClause(c *ReadOnlyClauseContext)

	// EnterIndexingClause is called when entering the indexingClause production.
	EnterIndexingClause(c *IndexingClauseContext)

	// EnterTablePartitioningClauses is called when entering the tablePartitioningClauses production.
	EnterTablePartitioningClauses(c *TablePartitioningClausesContext)

	// EnterRangePartitions is called when entering the rangePartitions production.
	EnterRangePartitions(c *RangePartitionsContext)

	// EnterRangeValuesClause is called when entering the rangeValuesClause production.
	EnterRangeValuesClause(c *RangeValuesClauseContext)

	// EnterTablePartitionDescription is called when entering the tablePartitionDescription production.
	EnterTablePartitionDescription(c *TablePartitionDescriptionContext)

	// EnterInmemoryClause is called when entering the inmemoryClause production.
	EnterInmemoryClause(c *InmemoryClauseContext)

	// EnterVarrayColProperties is called when entering the varrayColProperties production.
	EnterVarrayColProperties(c *VarrayColPropertiesContext)

	// EnterNestedTableColProperties is called when entering the nestedTableColProperties production.
	EnterNestedTableColProperties(c *NestedTableColPropertiesContext)

	// EnterLobStorageClause is called when entering the lobStorageClause production.
	EnterLobStorageClause(c *LobStorageClauseContext)

	// EnterVarrayStorageClause is called when entering the varrayStorageClause production.
	EnterVarrayStorageClause(c *VarrayStorageClauseContext)

	// EnterLobStorageParameters is called when entering the lobStorageParameters production.
	EnterLobStorageParameters(c *LobStorageParametersContext)

	// EnterLobParameters is called when entering the lobParameters production.
	EnterLobParameters(c *LobParametersContext)

	// EnterLobRetentionClause is called when entering the lobRetentionClause production.
	EnterLobRetentionClause(c *LobRetentionClauseContext)

	// EnterLobDeduplicateClause is called when entering the lobDeduplicateClause production.
	EnterLobDeduplicateClause(c *LobDeduplicateClauseContext)

	// EnterLobCompressionClause is called when entering the lobCompressionClause production.
	EnterLobCompressionClause(c *LobCompressionClauseContext)

	// EnterExternalPartSubpartDataProps is called when entering the externalPartSubpartDataProps production.
	EnterExternalPartSubpartDataProps(c *ExternalPartSubpartDataPropsContext)

	// EnterListPartitions is called when entering the listPartitions production.
	EnterListPartitions(c *ListPartitionsContext)

	// EnterListValuesClause is called when entering the listValuesClause production.
	EnterListValuesClause(c *ListValuesClauseContext)

	// EnterListValues is called when entering the listValues production.
	EnterListValues(c *ListValuesContext)

	// EnterHashPartitions is called when entering the hashPartitions production.
	EnterHashPartitions(c *HashPartitionsContext)

	// EnterHashPartitionsByQuantity is called when entering the hashPartitionsByQuantity production.
	EnterHashPartitionsByQuantity(c *HashPartitionsByQuantityContext)

	// EnterIndexCompression is called when entering the indexCompression production.
	EnterIndexCompression(c *IndexCompressionContext)

	// EnterAdvancedIndexCompression is called when entering the advancedIndexCompression production.
	EnterAdvancedIndexCompression(c *AdvancedIndexCompressionContext)

	// EnterIndividualHashPartitions is called when entering the individualHashPartitions production.
	EnterIndividualHashPartitions(c *IndividualHashPartitionsContext)

	// EnterPartitioningStorageClause is called when entering the partitioningStorageClause production.
	EnterPartitioningStorageClause(c *PartitioningStorageClauseContext)

	// EnterLobPartitioningStorage is called when entering the lobPartitioningStorage production.
	EnterLobPartitioningStorage(c *LobPartitioningStorageContext)

	// EnterCompositeRangePartitions is called when entering the compositeRangePartitions production.
	EnterCompositeRangePartitions(c *CompositeRangePartitionsContext)

	// EnterSubpartitionByRange is called when entering the subpartitionByRange production.
	EnterSubpartitionByRange(c *SubpartitionByRangeContext)

	// EnterSubpartitionByList is called when entering the subpartitionByList production.
	EnterSubpartitionByList(c *SubpartitionByListContext)

	// EnterSubpartitionByHash is called when entering the subpartitionByHash production.
	EnterSubpartitionByHash(c *SubpartitionByHashContext)

	// EnterSubpartitionTemplate is called when entering the subpartitionTemplate production.
	EnterSubpartitionTemplate(c *SubpartitionTemplateContext)

	// EnterRangeSubpartitionDesc is called when entering the rangeSubpartitionDesc production.
	EnterRangeSubpartitionDesc(c *RangeSubpartitionDescContext)

	// EnterListSubpartitionDesc is called when entering the listSubpartitionDesc production.
	EnterListSubpartitionDesc(c *ListSubpartitionDescContext)

	// EnterIndividualHashSubparts is called when entering the individualHashSubparts production.
	EnterIndividualHashSubparts(c *IndividualHashSubpartsContext)

	// EnterRangePartitionDesc is called when entering the rangePartitionDesc production.
	EnterRangePartitionDesc(c *RangePartitionDescContext)

	// EnterCompositeListPartitions is called when entering the compositeListPartitions production.
	EnterCompositeListPartitions(c *CompositeListPartitionsContext)

	// EnterListPartitionDesc is called when entering the listPartitionDesc production.
	EnterListPartitionDesc(c *ListPartitionDescContext)

	// EnterCompositeHashPartitions is called when entering the compositeHashPartitions production.
	EnterCompositeHashPartitions(c *CompositeHashPartitionsContext)

	// EnterReferencePartitioning is called when entering the referencePartitioning production.
	EnterReferencePartitioning(c *ReferencePartitioningContext)

	// EnterReferencePartitionDesc is called when entering the referencePartitionDesc production.
	EnterReferencePartitionDesc(c *ReferencePartitionDescContext)

	// EnterConstraint is called when entering the constraint production.
	EnterConstraint(c *ConstraintContext)

	// EnterSystemPartitioning is called when entering the systemPartitioning production.
	EnterSystemPartitioning(c *SystemPartitioningContext)

	// EnterConsistentHashPartitions is called when entering the consistentHashPartitions production.
	EnterConsistentHashPartitions(c *ConsistentHashPartitionsContext)

	// EnterConsistentHashWithSubpartitions is called when entering the consistentHashWithSubpartitions production.
	EnterConsistentHashWithSubpartitions(c *ConsistentHashWithSubpartitionsContext)

	// EnterPartitionsetClauses is called when entering the partitionsetClauses production.
	EnterPartitionsetClauses(c *PartitionsetClausesContext)

	// EnterRangePartitionsetClause is called when entering the rangePartitionsetClause production.
	EnterRangePartitionsetClause(c *RangePartitionsetClauseContext)

	// EnterRangePartitionsetDesc is called when entering the rangePartitionsetDesc production.
	EnterRangePartitionsetDesc(c *RangePartitionsetDescContext)

	// EnterListPartitionsetClause is called when entering the listPartitionsetClause production.
	EnterListPartitionsetClause(c *ListPartitionsetClauseContext)

	// EnterAttributeClusteringClause is called when entering the attributeClusteringClause production.
	EnterAttributeClusteringClause(c *AttributeClusteringClauseContext)

	// EnterClusteringJoin is called when entering the clusteringJoin production.
	EnterClusteringJoin(c *ClusteringJoinContext)

	// EnterClusterClause is called when entering the clusterClause production.
	EnterClusterClause(c *ClusterClauseContext)

	// EnterClusteringColumns is called when entering the clusteringColumns production.
	EnterClusteringColumns(c *ClusteringColumnsContext)

	// EnterClusteringColumnGroup is called when entering the clusteringColumnGroup production.
	EnterClusteringColumnGroup(c *ClusteringColumnGroupContext)

	// EnterClusteringWhen is called when entering the clusteringWhen production.
	EnterClusteringWhen(c *ClusteringWhenContext)

	// EnterZonemapClause is called when entering the zonemapClause production.
	EnterZonemapClause(c *ZonemapClauseContext)

	// EnterRowMovementClause is called when entering the rowMovementClause production.
	EnterRowMovementClause(c *RowMovementClauseContext)

	// EnterFlashbackArchiveClause is called when entering the flashbackArchiveClause production.
	EnterFlashbackArchiveClause(c *FlashbackArchiveClauseContext)

	// EnterAlterSynonym is called when entering the alterSynonym production.
	EnterAlterSynonym(c *AlterSynonymContext)

	// EnterAlterTablePartitioning is called when entering the alterTablePartitioning production.
	EnterAlterTablePartitioning(c *AlterTablePartitioningContext)

	// EnterModifyTablePartition is called when entering the modifyTablePartition production.
	EnterModifyTablePartition(c *ModifyTablePartitionContext)

	// EnterModifyRangePartition is called when entering the modifyRangePartition production.
	EnterModifyRangePartition(c *ModifyRangePartitionContext)

	// EnterModifyHashPartition is called when entering the modifyHashPartition production.
	EnterModifyHashPartition(c *ModifyHashPartitionContext)

	// EnterModifyListPartition is called when entering the modifyListPartition production.
	EnterModifyListPartition(c *ModifyListPartitionContext)

	// EnterPartitionExtendedName is called when entering the partitionExtendedName production.
	EnterPartitionExtendedName(c *PartitionExtendedNameContext)

	// EnterAddRangeSubpartition is called when entering the addRangeSubpartition production.
	EnterAddRangeSubpartition(c *AddRangeSubpartitionContext)

	// EnterDependentTablesClause is called when entering the dependentTablesClause production.
	EnterDependentTablesClause(c *DependentTablesClauseContext)

	// EnterAddHashSubpartition is called when entering the addHashSubpartition production.
	EnterAddHashSubpartition(c *AddHashSubpartitionContext)

	// EnterAddListSubpartition is called when entering the addListSubpartition production.
	EnterAddListSubpartition(c *AddListSubpartitionContext)

	// EnterCoalesceTableSubpartition is called when entering the coalesceTableSubpartition production.
	EnterCoalesceTableSubpartition(c *CoalesceTableSubpartitionContext)

	// EnterAllowDisallowClustering is called when entering the allowDisallowClustering production.
	EnterAllowDisallowClustering(c *AllowDisallowClusteringContext)

	// EnterAlterMappingTableClauses is called when entering the alterMappingTableClauses production.
	EnterAlterMappingTableClauses(c *AlterMappingTableClausesContext)

	// EnterDeallocateUnusedClause is called when entering the deallocateUnusedClause production.
	EnterDeallocateUnusedClause(c *DeallocateUnusedClauseContext)

	// EnterAllocateExtentClause is called when entering the allocateExtentClause production.
	EnterAllocateExtentClause(c *AllocateExtentClauseContext)

	// EnterPartitionSpec is called when entering the partitionSpec production.
	EnterPartitionSpec(c *PartitionSpecContext)

	// EnterPartitionAttributes is called when entering the partitionAttributes production.
	EnterPartitionAttributes(c *PartitionAttributesContext)

	// EnterShrinkClause is called when entering the shrinkClause production.
	EnterShrinkClause(c *ShrinkClauseContext)

	// EnterMoveTablePartition is called when entering the moveTablePartition production.
	EnterMoveTablePartition(c *MoveTablePartitionContext)

	// EnterFilterCondition is called when entering the filterCondition production.
	EnterFilterCondition(c *FilterConditionContext)

	// EnterCoalesceTablePartition is called when entering the coalesceTablePartition production.
	EnterCoalesceTablePartition(c *CoalesceTablePartitionContext)

	// EnterAddTablePartition is called when entering the addTablePartition production.
	EnterAddTablePartition(c *AddTablePartitionContext)

	// EnterAddRangePartitionClause is called when entering the addRangePartitionClause production.
	EnterAddRangePartitionClause(c *AddRangePartitionClauseContext)

	// EnterAddListPartitionClause is called when entering the addListPartitionClause production.
	EnterAddListPartitionClause(c *AddListPartitionClauseContext)

	// EnterHashSubpartsByQuantity is called when entering the hashSubpartsByQuantity production.
	EnterHashSubpartsByQuantity(c *HashSubpartsByQuantityContext)

	// EnterAddSystemPartitionClause is called when entering the addSystemPartitionClause production.
	EnterAddSystemPartitionClause(c *AddSystemPartitionClauseContext)

	// EnterAddHashPartitionClause is called when entering the addHashPartitionClause production.
	EnterAddHashPartitionClause(c *AddHashPartitionClauseContext)

	// EnterDropTablePartition is called when entering the dropTablePartition production.
	EnterDropTablePartition(c *DropTablePartitionContext)

	// EnterPartitionExtendedNames is called when entering the partitionExtendedNames production.
	EnterPartitionExtendedNames(c *PartitionExtendedNamesContext)

	// EnterPartitionForClauses is called when entering the partitionForClauses production.
	EnterPartitionForClauses(c *PartitionForClausesContext)

	// EnterUpdateIndexClauses is called when entering the updateIndexClauses production.
	EnterUpdateIndexClauses(c *UpdateIndexClausesContext)

	// EnterUpdateGlobalIndexClause is called when entering the updateGlobalIndexClause production.
	EnterUpdateGlobalIndexClause(c *UpdateGlobalIndexClauseContext)

	// EnterUpdateAllIndexesClause is called when entering the updateAllIndexesClause production.
	EnterUpdateAllIndexesClause(c *UpdateAllIndexesClauseContext)

	// EnterUpdateIndexPartition is called when entering the updateIndexPartition production.
	EnterUpdateIndexPartition(c *UpdateIndexPartitionContext)

	// EnterIndexPartitionDesc is called when entering the indexPartitionDesc production.
	EnterIndexPartitionDesc(c *IndexPartitionDescContext)

	// EnterIndexSubpartitionClause is called when entering the indexSubpartitionClause production.
	EnterIndexSubpartitionClause(c *IndexSubpartitionClauseContext)

	// EnterUpdateIndexSubpartition is called when entering the updateIndexSubpartition production.
	EnterUpdateIndexSubpartition(c *UpdateIndexSubpartitionContext)

	// EnterSupplementalLoggingProps is called when entering the supplementalLoggingProps production.
	EnterSupplementalLoggingProps(c *SupplementalLoggingPropsContext)

	// EnterSupplementalLogGrpClause is called when entering the supplementalLogGrpClause production.
	EnterSupplementalLogGrpClause(c *SupplementalLogGrpClauseContext)

	// EnterSupplementalIdKeyClause is called when entering the supplementalIdKeyClause production.
	EnterSupplementalIdKeyClause(c *SupplementalIdKeyClauseContext)

	// EnterAlterSession is called when entering the alterSession production.
	EnterAlterSession(c *AlterSessionContext)

	// EnterAlterSessionOption is called when entering the alterSessionOption production.
	EnterAlterSessionOption(c *AlterSessionOptionContext)

	// EnterAdviseClause is called when entering the adviseClause production.
	EnterAdviseClause(c *AdviseClauseContext)

	// EnterCloseDatabaseLinkClause is called when entering the closeDatabaseLinkClause production.
	EnterCloseDatabaseLinkClause(c *CloseDatabaseLinkClauseContext)

	// EnterCommitInProcedureClause is called when entering the commitInProcedureClause production.
	EnterCommitInProcedureClause(c *CommitInProcedureClauseContext)

	// EnterSecuriyClause is called when entering the securiyClause production.
	EnterSecuriyClause(c *SecuriyClauseContext)

	// EnterParallelExecutionClause is called when entering the parallelExecutionClause production.
	EnterParallelExecutionClause(c *ParallelExecutionClauseContext)

	// EnterResumableClause is called when entering the resumableClause production.
	EnterResumableClause(c *ResumableClauseContext)

	// EnterEnableResumableClause is called when entering the enableResumableClause production.
	EnterEnableResumableClause(c *EnableResumableClauseContext)

	// EnterDisableResumableClause is called when entering the disableResumableClause production.
	EnterDisableResumableClause(c *DisableResumableClauseContext)

	// EnterShardDdlClause is called when entering the shardDdlClause production.
	EnterShardDdlClause(c *ShardDdlClauseContext)

	// EnterSyncWithPrimaryClause is called when entering the syncWithPrimaryClause production.
	EnterSyncWithPrimaryClause(c *SyncWithPrimaryClauseContext)

	// EnterAlterSessionSetClause is called when entering the alterSessionSetClause production.
	EnterAlterSessionSetClause(c *AlterSessionSetClauseContext)

	// EnterAlterSessionSetClauseOption is called when entering the alterSessionSetClauseOption production.
	EnterAlterSessionSetClauseOption(c *AlterSessionSetClauseOptionContext)

	// EnterParameterClause is called when entering the parameterClause production.
	EnterParameterClause(c *ParameterClauseContext)

	// EnterEditionClause is called when entering the editionClause production.
	EnterEditionClause(c *EditionClauseContext)

	// EnterContainerClause is called when entering the containerClause production.
	EnterContainerClause(c *ContainerClauseContext)

	// EnterRowArchivalVisibilityClause is called when entering the rowArchivalVisibilityClause production.
	EnterRowArchivalVisibilityClause(c *RowArchivalVisibilityClauseContext)

	// EnterAlterDatabase is called when entering the alterDatabase production.
	EnterAlterDatabase(c *AlterDatabaseContext)

	// EnterDatabaseClauses is called when entering the databaseClauses production.
	EnterDatabaseClauses(c *DatabaseClausesContext)

	// EnterStartupClauses is called when entering the startupClauses production.
	EnterStartupClauses(c *StartupClausesContext)

	// EnterRecoveryClauses is called when entering the recoveryClauses production.
	EnterRecoveryClauses(c *RecoveryClausesContext)

	// EnterGeneralRecovery is called when entering the generalRecovery production.
	EnterGeneralRecovery(c *GeneralRecoveryContext)

	// EnterFullDatabaseRecovery is called when entering the fullDatabaseRecovery production.
	EnterFullDatabaseRecovery(c *FullDatabaseRecoveryContext)

	// EnterPartialDatabaseRecovery is called when entering the partialDatabaseRecovery production.
	EnterPartialDatabaseRecovery(c *PartialDatabaseRecoveryContext)

	// EnterManagedStandbyRecovery is called when entering the managedStandbyRecovery production.
	EnterManagedStandbyRecovery(c *ManagedStandbyRecoveryContext)

	// EnterDatabaseFileClauses is called when entering the databaseFileClauses production.
	EnterDatabaseFileClauses(c *DatabaseFileClausesContext)

	// EnterCreateDatafileClause is called when entering the createDatafileClause production.
	EnterCreateDatafileClause(c *CreateDatafileClauseContext)

	// EnterFileSpecifications is called when entering the fileSpecifications production.
	EnterFileSpecifications(c *FileSpecificationsContext)

	// EnterFileSpecification is called when entering the fileSpecification production.
	EnterFileSpecification(c *FileSpecificationContext)

	// EnterDatafileTempfileSpec is called when entering the datafileTempfileSpec production.
	EnterDatafileTempfileSpec(c *DatafileTempfileSpecContext)

	// EnterAutoextendClause is called when entering the autoextendClause production.
	EnterAutoextendClause(c *AutoextendClauseContext)

	// EnterRedoLogFileSpec is called when entering the redoLogFileSpec production.
	EnterRedoLogFileSpec(c *RedoLogFileSpecContext)

	// EnterAlterDatafileClause is called when entering the alterDatafileClause production.
	EnterAlterDatafileClause(c *AlterDatafileClauseContext)

	// EnterAlterTempfileClause is called when entering the alterTempfileClause production.
	EnterAlterTempfileClause(c *AlterTempfileClauseContext)

	// EnterLogfileClauses is called when entering the logfileClauses production.
	EnterLogfileClauses(c *LogfileClausesContext)

	// EnterLogfileDescriptor is called when entering the logfileDescriptor production.
	EnterLogfileDescriptor(c *LogfileDescriptorContext)

	// EnterAddLogfileClauses is called when entering the addLogfileClauses production.
	EnterAddLogfileClauses(c *AddLogfileClausesContext)

	// EnterControlfileClauses is called when entering the controlfileClauses production.
	EnterControlfileClauses(c *ControlfileClausesContext)

	// EnterTraceFileClause is called when entering the traceFileClause production.
	EnterTraceFileClause(c *TraceFileClauseContext)

	// EnterDropLogfileClauses is called when entering the dropLogfileClauses production.
	EnterDropLogfileClauses(c *DropLogfileClausesContext)

	// EnterSwitchLogfileClause is called when entering the switchLogfileClause production.
	EnterSwitchLogfileClause(c *SwitchLogfileClauseContext)

	// EnterSupplementalDbLogging is called when entering the supplementalDbLogging production.
	EnterSupplementalDbLogging(c *SupplementalDbLoggingContext)

	// EnterSupplementalPlsqlClause is called when entering the supplementalPlsqlClause production.
	EnterSupplementalPlsqlClause(c *SupplementalPlsqlClauseContext)

	// EnterSupplementalSubsetReplicationClause is called when entering the supplementalSubsetReplicationClause production.
	EnterSupplementalSubsetReplicationClause(c *SupplementalSubsetReplicationClauseContext)

	// EnterStandbyDatabaseClauses is called when entering the standbyDatabaseClauses production.
	EnterStandbyDatabaseClauses(c *StandbyDatabaseClausesContext)

	// EnterActivateStandbyDbClause is called when entering the activateStandbyDbClause production.
	EnterActivateStandbyDbClause(c *ActivateStandbyDbClauseContext)

	// EnterMaximizeStandbyDbClause is called when entering the maximizeStandbyDbClause production.
	EnterMaximizeStandbyDbClause(c *MaximizeStandbyDbClauseContext)

	// EnterRegisterLogfileClause is called when entering the registerLogfileClause production.
	EnterRegisterLogfileClause(c *RegisterLogfileClauseContext)

	// EnterCommitSwitchoverClause is called when entering the commitSwitchoverClause production.
	EnterCommitSwitchoverClause(c *CommitSwitchoverClauseContext)

	// EnterStartStandbyClause is called when entering the startStandbyClause production.
	EnterStartStandbyClause(c *StartStandbyClauseContext)

	// EnterStopStandbyClause is called when entering the stopStandbyClause production.
	EnterStopStandbyClause(c *StopStandbyClauseContext)

	// EnterSwitchoverClause is called when entering the switchoverClause production.
	EnterSwitchoverClause(c *SwitchoverClauseContext)

	// EnterConvertDatabaseClause is called when entering the convertDatabaseClause production.
	EnterConvertDatabaseClause(c *ConvertDatabaseClauseContext)

	// EnterFailoverClause is called when entering the failoverClause production.
	EnterFailoverClause(c *FailoverClauseContext)

	// EnterDefaultSettingsClauses is called when entering the defaultSettingsClauses production.
	EnterDefaultSettingsClauses(c *DefaultSettingsClausesContext)

	// EnterSetTimeZoneClause is called when entering the setTimeZoneClause production.
	EnterSetTimeZoneClause(c *SetTimeZoneClauseContext)

	// EnterTimeZoneRegion is called when entering the timeZoneRegion production.
	EnterTimeZoneRegion(c *TimeZoneRegionContext)

	// EnterFlashbackModeClause is called when entering the flashbackModeClause production.
	EnterFlashbackModeClause(c *FlashbackModeClauseContext)

	// EnterUndoModeClause is called when entering the undoModeClause production.
	EnterUndoModeClause(c *UndoModeClauseContext)

	// EnterMoveDatafileClause is called when entering the moveDatafileClause production.
	EnterMoveDatafileClause(c *MoveDatafileClauseContext)

	// EnterInstanceClauses is called when entering the instanceClauses production.
	EnterInstanceClauses(c *InstanceClausesContext)

	// EnterSecurityClause is called when entering the securityClause production.
	EnterSecurityClause(c *SecurityClauseContext)

	// EnterPrepareClause is called when entering the prepareClause production.
	EnterPrepareClause(c *PrepareClauseContext)

	// EnterDropMirrorCopy is called when entering the dropMirrorCopy production.
	EnterDropMirrorCopy(c *DropMirrorCopyContext)

	// EnterLostWriteProtection is called when entering the lostWriteProtection production.
	EnterLostWriteProtection(c *LostWriteProtectionContext)

	// EnterCdbFleetClauses is called when entering the cdbFleetClauses production.
	EnterCdbFleetClauses(c *CdbFleetClausesContext)

	// EnterLeadCdbClause is called when entering the leadCdbClause production.
	EnterLeadCdbClause(c *LeadCdbClauseContext)

	// EnterLeadCdbUriClause is called when entering the leadCdbUriClause production.
	EnterLeadCdbUriClause(c *LeadCdbUriClauseContext)

	// EnterPropertyClause is called when entering the propertyClause production.
	EnterPropertyClause(c *PropertyClauseContext)

	// EnterAlterSystem is called when entering the alterSystem production.
	EnterAlterSystem(c *AlterSystemContext)

	// EnterAlterSystemOption is called when entering the alterSystemOption production.
	EnterAlterSystemOption(c *AlterSystemOptionContext)

	// EnterArchiveLogClause is called when entering the archiveLogClause production.
	EnterArchiveLogClause(c *ArchiveLogClauseContext)

	// EnterCheckpointClause is called when entering the checkpointClause production.
	EnterCheckpointClause(c *CheckpointClauseContext)

	// EnterCheckDatafilesClause is called when entering the checkDatafilesClause production.
	EnterCheckDatafilesClause(c *CheckDatafilesClauseContext)

	// EnterDistributedRecovClauses is called when entering the distributedRecovClauses production.
	EnterDistributedRecovClauses(c *DistributedRecovClausesContext)

	// EnterFlushClause is called when entering the flushClause production.
	EnterFlushClause(c *FlushClauseContext)

	// EnterEndSessionClauses is called when entering the endSessionClauses production.
	EnterEndSessionClauses(c *EndSessionClausesContext)

	// EnterAlterSystemSwitchLogfileClause is called when entering the alterSystemSwitchLogfileClause production.
	EnterAlterSystemSwitchLogfileClause(c *AlterSystemSwitchLogfileClauseContext)

	// EnterSuspendResumeClause is called when entering the suspendResumeClause production.
	EnterSuspendResumeClause(c *SuspendResumeClauseContext)

	// EnterQuiesceClauses is called when entering the quiesceClauses production.
	EnterQuiesceClauses(c *QuiesceClausesContext)

	// EnterRollingMigrationClauses is called when entering the rollingMigrationClauses production.
	EnterRollingMigrationClauses(c *RollingMigrationClausesContext)

	// EnterRollingPatchClauses is called when entering the rollingPatchClauses production.
	EnterRollingPatchClauses(c *RollingPatchClausesContext)

	// EnterAlterSystemSecurityClauses is called when entering the alterSystemSecurityClauses production.
	EnterAlterSystemSecurityClauses(c *AlterSystemSecurityClausesContext)

	// EnterAffinityClauses is called when entering the affinityClauses production.
	EnterAffinityClauses(c *AffinityClausesContext)

	// EnterShutdownDispatcherClause is called when entering the shutdownDispatcherClause production.
	EnterShutdownDispatcherClause(c *ShutdownDispatcherClauseContext)

	// EnterRegisterClause is called when entering the registerClause production.
	EnterRegisterClause(c *RegisterClauseContext)

	// EnterSetClause is called when entering the setClause production.
	EnterSetClause(c *SetClauseContext)

	// EnterResetClause is called when entering the resetClause production.
	EnterResetClause(c *ResetClauseContext)

	// EnterRelocateClientClause is called when entering the relocateClientClause production.
	EnterRelocateClientClause(c *RelocateClientClauseContext)

	// EnterCancelSqlClause is called when entering the cancelSqlClause production.
	EnterCancelSqlClause(c *CancelSqlClauseContext)

	// EnterFlushPasswordfileMetadataCacheClause is called when entering the flushPasswordfileMetadataCacheClause production.
	EnterFlushPasswordfileMetadataCacheClause(c *FlushPasswordfileMetadataCacheClauseContext)

	// EnterInstanceClause is called when entering the instanceClause production.
	EnterInstanceClause(c *InstanceClauseContext)

	// EnterSequenceClause is called when entering the sequenceClause production.
	EnterSequenceClause(c *SequenceClauseContext)

	// EnterChangeClause is called when entering the changeClause production.
	EnterChangeClause(c *ChangeClauseContext)

	// EnterCurrentClause is called when entering the currentClause production.
	EnterCurrentClause(c *CurrentClauseContext)

	// EnterGroupClause is called when entering the groupClause production.
	EnterGroupClause(c *GroupClauseContext)

	// EnterLogfileClause is called when entering the logfileClause production.
	EnterLogfileClause(c *LogfileClauseContext)

	// EnterNextClause is called when entering the nextClause production.
	EnterNextClause(c *NextClauseContext)

	// EnterAllClause is called when entering the allClause production.
	EnterAllClause(c *AllClauseContext)

	// EnterToLocationClause is called when entering the toLocationClause production.
	EnterToLocationClause(c *ToLocationClauseContext)

	// EnterFlushClauseOption is called when entering the flushClauseOption production.
	EnterFlushClauseOption(c *FlushClauseOptionContext)

	// EnterDisconnectSessionClause is called when entering the disconnectSessionClause production.
	EnterDisconnectSessionClause(c *DisconnectSessionClauseContext)

	// EnterKillSessionClause is called when entering the killSessionClause production.
	EnterKillSessionClause(c *KillSessionClauseContext)

	// EnterStartRollingMigrationClause is called when entering the startRollingMigrationClause production.
	EnterStartRollingMigrationClause(c *StartRollingMigrationClauseContext)

	// EnterStopRollingMigrationClause is called when entering the stopRollingMigrationClause production.
	EnterStopRollingMigrationClause(c *StopRollingMigrationClauseContext)

	// EnterStartRollingPatchClause is called when entering the startRollingPatchClause production.
	EnterStartRollingPatchClause(c *StartRollingPatchClauseContext)

	// EnterStopRollingPatchClause is called when entering the stopRollingPatchClause production.
	EnterStopRollingPatchClause(c *StopRollingPatchClauseContext)

	// EnterRestrictedSessionClause is called when entering the restrictedSessionClause production.
	EnterRestrictedSessionClause(c *RestrictedSessionClauseContext)

	// EnterSetEncryptionWalletOpenClause is called when entering the setEncryptionWalletOpenClause production.
	EnterSetEncryptionWalletOpenClause(c *SetEncryptionWalletOpenClauseContext)

	// EnterSetEncryptionWalletCloseClause is called when entering the setEncryptionWalletCloseClause production.
	EnterSetEncryptionWalletCloseClause(c *SetEncryptionWalletCloseClauseContext)

	// EnterSetEncryptionKeyClause is called when entering the setEncryptionKeyClause production.
	EnterSetEncryptionKeyClause(c *SetEncryptionKeyClauseContext)

	// EnterEnableAffinityClause is called when entering the enableAffinityClause production.
	EnterEnableAffinityClause(c *EnableAffinityClauseContext)

	// EnterDisableAffinityClause is called when entering the disableAffinityClause production.
	EnterDisableAffinityClause(c *DisableAffinityClauseContext)

	// EnterAlterSystemSetClause is called when entering the alterSystemSetClause production.
	EnterAlterSystemSetClause(c *AlterSystemSetClauseContext)

	// EnterAlterSystemResetClause is called when entering the alterSystemResetClause production.
	EnterAlterSystemResetClause(c *AlterSystemResetClauseContext)

	// EnterSharedPoolClause is called when entering the sharedPoolClause production.
	EnterSharedPoolClause(c *SharedPoolClauseContext)

	// EnterGlobalContextClause is called when entering the globalContextClause production.
	EnterGlobalContextClause(c *GlobalContextClauseContext)

	// EnterBufferCacheClause is called when entering the bufferCacheClause production.
	EnterBufferCacheClause(c *BufferCacheClauseContext)

	// EnterFlashCacheClause is called when entering the flashCacheClause production.
	EnterFlashCacheClause(c *FlashCacheClauseContext)

	// EnterRedoToClause is called when entering the redoToClause production.
	EnterRedoToClause(c *RedoToClauseContext)

	// EnterIdentifiedByWalletPassword is called when entering the identifiedByWalletPassword production.
	EnterIdentifiedByWalletPassword(c *IdentifiedByWalletPasswordContext)

	// EnterIdentifiedByHsmAuthString is called when entering the identifiedByHsmAuthString production.
	EnterIdentifiedByHsmAuthString(c *IdentifiedByHsmAuthStringContext)

	// EnterSetParameterClause is called when entering the setParameterClause production.
	EnterSetParameterClause(c *SetParameterClauseContext)

	// EnterUseStoredOutlinesClause is called when entering the useStoredOutlinesClause production.
	EnterUseStoredOutlinesClause(c *UseStoredOutlinesClauseContext)

	// EnterGlobalTopicEnabledClause is called when entering the globalTopicEnabledClause production.
	EnterGlobalTopicEnabledClause(c *GlobalTopicEnabledClauseContext)

	// EnterAlterSystemCommentClause is called when entering the alterSystemCommentClause production.
	EnterAlterSystemCommentClause(c *AlterSystemCommentClauseContext)

	// EnterContainerCurrentAllClause is called when entering the containerCurrentAllClause production.
	EnterContainerCurrentAllClause(c *ContainerCurrentAllClauseContext)

	// EnterScopeClause is called when entering the scopeClause production.
	EnterScopeClause(c *ScopeClauseContext)

	// EnterAnalyze is called when entering the analyze production.
	EnterAnalyze(c *AnalyzeContext)

	// EnterPartitionExtensionClause is called when entering the partitionExtensionClause production.
	EnterPartitionExtensionClause(c *PartitionExtensionClauseContext)

	// EnterValidationClauses is called when entering the validationClauses production.
	EnterValidationClauses(c *ValidationClausesContext)

	// EnterAssociateStatistics is called when entering the associateStatistics production.
	EnterAssociateStatistics(c *AssociateStatisticsContext)

	// EnterColumnAssociation is called when entering the columnAssociation production.
	EnterColumnAssociation(c *ColumnAssociationContext)

	// EnterFunctionAssociation is called when entering the functionAssociation production.
	EnterFunctionAssociation(c *FunctionAssociationContext)

	// EnterStorageTableClause is called when entering the storageTableClause production.
	EnterStorageTableClause(c *StorageTableClauseContext)

	// EnterUsingStatisticsType is called when entering the usingStatisticsType production.
	EnterUsingStatisticsType(c *UsingStatisticsTypeContext)

	// EnterDefaultCostClause is called when entering the defaultCostClause production.
	EnterDefaultCostClause(c *DefaultCostClauseContext)

	// EnterDefaultSelectivityClause is called when entering the defaultSelectivityClause production.
	EnterDefaultSelectivityClause(c *DefaultSelectivityClauseContext)

	// EnterDisassociateStatistics is called when entering the disassociateStatistics production.
	EnterDisassociateStatistics(c *DisassociateStatisticsContext)

	// EnterAudit is called when entering the audit production.
	EnterAudit(c *AuditContext)

	// EnterNoAudit is called when entering the noAudit production.
	EnterNoAudit(c *NoAuditContext)

	// EnterAuditPolicyClause is called when entering the auditPolicyClause production.
	EnterAuditPolicyClause(c *AuditPolicyClauseContext)

	// EnterNoAuditPolicyClause is called when entering the noAuditPolicyClause production.
	EnterNoAuditPolicyClause(c *NoAuditPolicyClauseContext)

	// EnterByUsersWithRoles is called when entering the byUsersWithRoles production.
	EnterByUsersWithRoles(c *ByUsersWithRolesContext)

	// EnterContextClause is called when entering the contextClause production.
	EnterContextClause(c *ContextClauseContext)

	// EnterContextNamespaceAttributesClause is called when entering the contextNamespaceAttributesClause production.
	EnterContextNamespaceAttributesClause(c *ContextNamespaceAttributesClauseContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterFlashbackDatabase is called when entering the flashbackDatabase production.
	EnterFlashbackDatabase(c *FlashbackDatabaseContext)

	// EnterScnTimestampClause is called when entering the scnTimestampClause production.
	EnterScnTimestampClause(c *ScnTimestampClauseContext)

	// EnterRestorePointClause is called when entering the restorePointClause production.
	EnterRestorePointClause(c *RestorePointClauseContext)

	// EnterFlashbackTable is called when entering the flashbackTable production.
	EnterFlashbackTable(c *FlashbackTableContext)

	// EnterRenameToTable is called when entering the renameToTable production.
	EnterRenameToTable(c *RenameToTableContext)

	// EnterPurge is called when entering the purge production.
	EnterPurge(c *PurgeContext)

	// EnterRename is called when entering the rename production.
	EnterRename(c *RenameContext)

	// EnterCreateDatabase is called when entering the createDatabase production.
	EnterCreateDatabase(c *CreateDatabaseContext)

	// EnterCreateDatabaseClauses is called when entering the createDatabaseClauses production.
	EnterCreateDatabaseClauses(c *CreateDatabaseClausesContext)

	// EnterDatabaseLoggingClauses is called when entering the databaseLoggingClauses production.
	EnterDatabaseLoggingClauses(c *DatabaseLoggingClausesContext)

	// EnterTablespaceClauses is called when entering the tablespaceClauses production.
	EnterTablespaceClauses(c *TablespaceClausesContext)

	// EnterDefaultTablespace is called when entering the defaultTablespace production.
	EnterDefaultTablespace(c *DefaultTablespaceContext)

	// EnterDefaultTempTablespace is called when entering the defaultTempTablespace production.
	EnterDefaultTempTablespace(c *DefaultTempTablespaceContext)

	// EnterUndoTablespace is called when entering the undoTablespace production.
	EnterUndoTablespace(c *UndoTablespaceContext)

	// EnterBigOrSmallFiles is called when entering the bigOrSmallFiles production.
	EnterBigOrSmallFiles(c *BigOrSmallFilesContext)

	// EnterExtentManagementClause is called when entering the extentManagementClause production.
	EnterExtentManagementClause(c *ExtentManagementClauseContext)

	// EnterEnablePluggableDatabase is called when entering the enablePluggableDatabase production.
	EnterEnablePluggableDatabase(c *EnablePluggableDatabaseContext)

	// EnterFileNameConvert is called when entering the fileNameConvert production.
	EnterFileNameConvert(c *FileNameConvertContext)

	// EnterReplaceFileNamePattern is called when entering the replaceFileNamePattern production.
	EnterReplaceFileNamePattern(c *ReplaceFileNamePatternContext)

	// EnterTablespaceDatafileClauses is called when entering the tablespaceDatafileClauses production.
	EnterTablespaceDatafileClauses(c *TablespaceDatafileClausesContext)

	// EnterCreateDatabaseLink is called when entering the createDatabaseLink production.
	EnterCreateDatabaseLink(c *CreateDatabaseLinkContext)

	// EnterConnectToClause is called when entering the connectToClause production.
	EnterConnectToClause(c *ConnectToClauseContext)

	// EnterDbLinkAuthentication is called when entering the dbLinkAuthentication production.
	EnterDbLinkAuthentication(c *DbLinkAuthenticationContext)

	// EnterSetTransaction is called when entering the setTransaction production.
	EnterSetTransaction(c *SetTransactionContext)

	// EnterCommit is called when entering the commit production.
	EnterCommit(c *CommitContext)

	// EnterCommentClause is called when entering the commentClause production.
	EnterCommentClause(c *CommentClauseContext)

	// EnterWriteClause is called when entering the writeClause production.
	EnterWriteClause(c *WriteClauseContext)

	// EnterForceClause is called when entering the forceClause production.
	EnterForceClause(c *ForceClauseContext)

	// EnterRollback is called when entering the rollback production.
	EnterRollback(c *RollbackContext)

	// EnterSavepointClause is called when entering the savepointClause production.
	EnterSavepointClause(c *SavepointClauseContext)

	// EnterSavepoint is called when entering the savepoint production.
	EnterSavepoint(c *SavepointContext)

	// EnterSetConstraints is called when entering the setConstraints production.
	EnterSetConstraints(c *SetConstraintsContext)

	// EnterGrant is called when entering the grant production.
	EnterGrant(c *GrantContext)

	// EnterRevoke is called when entering the revoke production.
	EnterRevoke(c *RevokeContext)

	// EnterObjectPrivilegeClause is called when entering the objectPrivilegeClause production.
	EnterObjectPrivilegeClause(c *ObjectPrivilegeClauseContext)

	// EnterSystemPrivilegeClause is called when entering the systemPrivilegeClause production.
	EnterSystemPrivilegeClause(c *SystemPrivilegeClauseContext)

	// EnterRoleClause is called when entering the roleClause production.
	EnterRoleClause(c *RoleClauseContext)

	// EnterObjectPrivileges is called when entering the objectPrivileges production.
	EnterObjectPrivileges(c *ObjectPrivilegesContext)

	// EnterObjectPrivilegeType is called when entering the objectPrivilegeType production.
	EnterObjectPrivilegeType(c *ObjectPrivilegeTypeContext)

	// EnterOnObjectClause is called when entering the onObjectClause production.
	EnterOnObjectClause(c *OnObjectClauseContext)

	// EnterSystemPrivilege is called when entering the systemPrivilege production.
	EnterSystemPrivilege(c *SystemPrivilegeContext)

	// EnterSystemPrivilegeOperation is called when entering the systemPrivilegeOperation production.
	EnterSystemPrivilegeOperation(c *SystemPrivilegeOperationContext)

	// EnterAdvisorFrameworkSystemPrivilege is called when entering the advisorFrameworkSystemPrivilege production.
	EnterAdvisorFrameworkSystemPrivilege(c *AdvisorFrameworkSystemPrivilegeContext)

	// EnterClustersSystemPrivilege is called when entering the clustersSystemPrivilege production.
	EnterClustersSystemPrivilege(c *ClustersSystemPrivilegeContext)

	// EnterContextsSystemPrivilege is called when entering the contextsSystemPrivilege production.
	EnterContextsSystemPrivilege(c *ContextsSystemPrivilegeContext)

	// EnterDataRedactionSystemPrivilege is called when entering the dataRedactionSystemPrivilege production.
	EnterDataRedactionSystemPrivilege(c *DataRedactionSystemPrivilegeContext)

	// EnterDatabaseSystemPrivilege is called when entering the databaseSystemPrivilege production.
	EnterDatabaseSystemPrivilege(c *DatabaseSystemPrivilegeContext)

	// EnterDatabaseLinksSystemPrivilege is called when entering the databaseLinksSystemPrivilege production.
	EnterDatabaseLinksSystemPrivilege(c *DatabaseLinksSystemPrivilegeContext)

	// EnterDebuggingSystemPrivilege is called when entering the debuggingSystemPrivilege production.
	EnterDebuggingSystemPrivilege(c *DebuggingSystemPrivilegeContext)

	// EnterDictionariesSystemPrivilege is called when entering the dictionariesSystemPrivilege production.
	EnterDictionariesSystemPrivilege(c *DictionariesSystemPrivilegeContext)

	// EnterDimensionsSystemPrivilege is called when entering the dimensionsSystemPrivilege production.
	EnterDimensionsSystemPrivilege(c *DimensionsSystemPrivilegeContext)

	// EnterDirectoriesSystemPrivilege is called when entering the directoriesSystemPrivilege production.
	EnterDirectoriesSystemPrivilege(c *DirectoriesSystemPrivilegeContext)

	// EnterEditionsSystemPrivilege is called when entering the editionsSystemPrivilege production.
	EnterEditionsSystemPrivilege(c *EditionsSystemPrivilegeContext)

	// EnterFlashbackDataArchivesPrivilege is called when entering the flashbackDataArchivesPrivilege production.
	EnterFlashbackDataArchivesPrivilege(c *FlashbackDataArchivesPrivilegeContext)

	// EnterIndexesSystemPrivilege is called when entering the indexesSystemPrivilege production.
	EnterIndexesSystemPrivilege(c *IndexesSystemPrivilegeContext)

	// EnterIndexTypesSystemPrivilege is called when entering the indexTypesSystemPrivilege production.
	EnterIndexTypesSystemPrivilege(c *IndexTypesSystemPrivilegeContext)

	// EnterJobSchedulerObjectsSystemPrivilege is called when entering the jobSchedulerObjectsSystemPrivilege production.
	EnterJobSchedulerObjectsSystemPrivilege(c *JobSchedulerObjectsSystemPrivilegeContext)

	// EnterKeyManagementFrameworkSystemPrivilege is called when entering the keyManagementFrameworkSystemPrivilege production.
	EnterKeyManagementFrameworkSystemPrivilege(c *KeyManagementFrameworkSystemPrivilegeContext)

	// EnterLibrariesFrameworkSystemPrivilege is called when entering the librariesFrameworkSystemPrivilege production.
	EnterLibrariesFrameworkSystemPrivilege(c *LibrariesFrameworkSystemPrivilegeContext)

	// EnterLogminerFrameworkSystemPrivilege is called when entering the logminerFrameworkSystemPrivilege production.
	EnterLogminerFrameworkSystemPrivilege(c *LogminerFrameworkSystemPrivilegeContext)

	// EnterMaterizlizedViewsSystemPrivilege is called when entering the materizlizedViewsSystemPrivilege production.
	EnterMaterizlizedViewsSystemPrivilege(c *MaterizlizedViewsSystemPrivilegeContext)

	// EnterMiningModelsSystemPrivilege is called when entering the miningModelsSystemPrivilege production.
	EnterMiningModelsSystemPrivilege(c *MiningModelsSystemPrivilegeContext)

	// EnterOlapCubesSystemPrivilege is called when entering the olapCubesSystemPrivilege production.
	EnterOlapCubesSystemPrivilege(c *OlapCubesSystemPrivilegeContext)

	// EnterOlapCubeMeasureFoldersSystemPrivilege is called when entering the olapCubeMeasureFoldersSystemPrivilege production.
	EnterOlapCubeMeasureFoldersSystemPrivilege(c *OlapCubeMeasureFoldersSystemPrivilegeContext)

	// EnterOlapCubeDiminsionsSystemPrivilege is called when entering the olapCubeDiminsionsSystemPrivilege production.
	EnterOlapCubeDiminsionsSystemPrivilege(c *OlapCubeDiminsionsSystemPrivilegeContext)

	// EnterOlapCubeBuildProcessesSystemPrivilege is called when entering the olapCubeBuildProcessesSystemPrivilege production.
	EnterOlapCubeBuildProcessesSystemPrivilege(c *OlapCubeBuildProcessesSystemPrivilegeContext)

	// EnterOperatorsSystemPrivilege is called when entering the operatorsSystemPrivilege production.
	EnterOperatorsSystemPrivilege(c *OperatorsSystemPrivilegeContext)

	// EnterOutlinesSystemPrivilege is called when entering the outlinesSystemPrivilege production.
	EnterOutlinesSystemPrivilege(c *OutlinesSystemPrivilegeContext)

	// EnterPlanManagementSystemPrivilege is called when entering the planManagementSystemPrivilege production.
	EnterPlanManagementSystemPrivilege(c *PlanManagementSystemPrivilegeContext)

	// EnterPluggableDatabasesSystemPrivilege is called when entering the pluggableDatabasesSystemPrivilege production.
	EnterPluggableDatabasesSystemPrivilege(c *PluggableDatabasesSystemPrivilegeContext)

	// EnterProceduresSystemPrivilege is called when entering the proceduresSystemPrivilege production.
	EnterProceduresSystemPrivilege(c *ProceduresSystemPrivilegeContext)

	// EnterProfilesSystemPrivilege is called when entering the profilesSystemPrivilege production.
	EnterProfilesSystemPrivilege(c *ProfilesSystemPrivilegeContext)

	// EnterRolesSystemPrivilege is called when entering the rolesSystemPrivilege production.
	EnterRolesSystemPrivilege(c *RolesSystemPrivilegeContext)

	// EnterRollbackSegmentsSystemPrivilege is called when entering the rollbackSegmentsSystemPrivilege production.
	EnterRollbackSegmentsSystemPrivilege(c *RollbackSegmentsSystemPrivilegeContext)

	// EnterSequencesSystemPrivilege is called when entering the sequencesSystemPrivilege production.
	EnterSequencesSystemPrivilege(c *SequencesSystemPrivilegeContext)

	// EnterSessionsSystemPrivilege is called when entering the sessionsSystemPrivilege production.
	EnterSessionsSystemPrivilege(c *SessionsSystemPrivilegeContext)

	// EnterSqlTranslationProfilesSystemPrivilege is called when entering the sqlTranslationProfilesSystemPrivilege production.
	EnterSqlTranslationProfilesSystemPrivilege(c *SqlTranslationProfilesSystemPrivilegeContext)

	// EnterSynonymsSystemPrivilege is called when entering the synonymsSystemPrivilege production.
	EnterSynonymsSystemPrivilege(c *SynonymsSystemPrivilegeContext)

	// EnterTablesSystemPrivilege is called when entering the tablesSystemPrivilege production.
	EnterTablesSystemPrivilege(c *TablesSystemPrivilegeContext)

	// EnterTablespacesSystemPrivilege is called when entering the tablespacesSystemPrivilege production.
	EnterTablespacesSystemPrivilege(c *TablespacesSystemPrivilegeContext)

	// EnterTriggersSystemPrivilege is called when entering the triggersSystemPrivilege production.
	EnterTriggersSystemPrivilege(c *TriggersSystemPrivilegeContext)

	// EnterTypesSystemPrivilege is called when entering the typesSystemPrivilege production.
	EnterTypesSystemPrivilege(c *TypesSystemPrivilegeContext)

	// EnterUsersSystemPrivilege is called when entering the usersSystemPrivilege production.
	EnterUsersSystemPrivilege(c *UsersSystemPrivilegeContext)

	// EnterViewsSystemPrivilege is called when entering the viewsSystemPrivilege production.
	EnterViewsSystemPrivilege(c *ViewsSystemPrivilegeContext)

	// EnterMiscellaneousSystemPrivilege is called when entering the miscellaneousSystemPrivilege production.
	EnterMiscellaneousSystemPrivilege(c *MiscellaneousSystemPrivilegeContext)

	// EnterCreateUser is called when entering the createUser production.
	EnterCreateUser(c *CreateUserContext)

	// EnterDropUser is called when entering the dropUser production.
	EnterDropUser(c *DropUserContext)

	// EnterAlterUser is called when entering the alterUser production.
	EnterAlterUser(c *AlterUserContext)

	// EnterCreateRole is called when entering the createRole production.
	EnterCreateRole(c *CreateRoleContext)

	// EnterDropRole is called when entering the dropRole production.
	EnterDropRole(c *DropRoleContext)

	// EnterAlterRole is called when entering the alterRole production.
	EnterAlterRole(c *AlterRoleContext)

	// EnterSetRole is called when entering the setRole production.
	EnterSetRole(c *SetRoleContext)

	// EnterRoleAssignment is called when entering the roleAssignment production.
	EnterRoleAssignment(c *RoleAssignmentContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)

	// ExitInsert is called when exiting the insert production.
	ExitInsert(c *InsertContext)

	// ExitInsertSingleTable is called when exiting the insertSingleTable production.
	ExitInsertSingleTable(c *InsertSingleTableContext)

	// ExitInsertMultiTable is called when exiting the insertMultiTable production.
	ExitInsertMultiTable(c *InsertMultiTableContext)

	// ExitMultiTableElement is called when exiting the multiTableElement production.
	ExitMultiTableElement(c *MultiTableElementContext)

	// ExitConditionalInsertClause is called when exiting the conditionalInsertClause production.
	ExitConditionalInsertClause(c *ConditionalInsertClauseContext)

	// ExitConditionalInsertWhenPart is called when exiting the conditionalInsertWhenPart production.
	ExitConditionalInsertWhenPart(c *ConditionalInsertWhenPartContext)

	// ExitConditionalInsertElsePart is called when exiting the conditionalInsertElsePart production.
	ExitConditionalInsertElsePart(c *ConditionalInsertElsePartContext)

	// ExitInsertIntoClause is called when exiting the insertIntoClause production.
	ExitInsertIntoClause(c *InsertIntoClauseContext)

	// ExitInsertValuesClause is called when exiting the insertValuesClause production.
	ExitInsertValuesClause(c *InsertValuesClauseContext)

	// ExitReturningClause is called when exiting the returningClause production.
	ExitReturningClause(c *ReturningClauseContext)

	// ExitDmlTableExprClause is called when exiting the dmlTableExprClause production.
	ExitDmlTableExprClause(c *DmlTableExprClauseContext)

	// ExitDmlTableClause is called when exiting the dmlTableClause production.
	ExitDmlTableClause(c *DmlTableClauseContext)

	// ExitPartitionExtClause is called when exiting the partitionExtClause production.
	ExitPartitionExtClause(c *PartitionExtClauseContext)

	// ExitDmlSubqueryClause is called when exiting the dmlSubqueryClause production.
	ExitDmlSubqueryClause(c *DmlSubqueryClauseContext)

	// ExitSubqueryRestrictionClause is called when exiting the subqueryRestrictionClause production.
	ExitSubqueryRestrictionClause(c *SubqueryRestrictionClauseContext)

	// ExitTableCollectionExpr is called when exiting the tableCollectionExpr production.
	ExitTableCollectionExpr(c *TableCollectionExprContext)

	// ExitCollectionExpr is called when exiting the collectionExpr production.
	ExitCollectionExpr(c *CollectionExprContext)

	// ExitUpdate is called when exiting the update production.
	ExitUpdate(c *UpdateContext)

	// ExitUpdateSpecification is called when exiting the updateSpecification production.
	ExitUpdateSpecification(c *UpdateSpecificationContext)

	// ExitUpdateSetClause is called when exiting the updateSetClause production.
	ExitUpdateSetClause(c *UpdateSetClauseContext)

	// ExitUpdateSetColumnList is called when exiting the updateSetColumnList production.
	ExitUpdateSetColumnList(c *UpdateSetColumnListContext)

	// ExitUpdateSetColumnClause is called when exiting the updateSetColumnClause production.
	ExitUpdateSetColumnClause(c *UpdateSetColumnClauseContext)

	// ExitUpdateSetValueClause is called when exiting the updateSetValueClause production.
	ExitUpdateSetValueClause(c *UpdateSetValueClauseContext)

	// ExitAssignmentValues is called when exiting the assignmentValues production.
	ExitAssignmentValues(c *AssignmentValuesContext)

	// ExitAssignmentValue is called when exiting the assignmentValue production.
	ExitAssignmentValue(c *AssignmentValueContext)

	// ExitDelete is called when exiting the delete production.
	ExitDelete(c *DeleteContext)

	// ExitDeleteSpecification is called when exiting the deleteSpecification production.
	ExitDeleteSpecification(c *DeleteSpecificationContext)

	// ExitSelect is called when exiting the select production.
	ExitSelect(c *SelectContext)

	// ExitSelectSubquery is called when exiting the selectSubquery production.
	ExitSelectSubquery(c *SelectSubqueryContext)

	// ExitSelectUnionClause is called when exiting the selectUnionClause production.
	ExitSelectUnionClause(c *SelectUnionClauseContext)

	// ExitParenthesisSelectSubquery is called when exiting the parenthesisSelectSubquery production.
	ExitParenthesisSelectSubquery(c *ParenthesisSelectSubqueryContext)

	// ExitQueryBlock is called when exiting the queryBlock production.
	ExitQueryBlock(c *QueryBlockContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitPlsqlDeclarations is called when exiting the plsqlDeclarations production.
	ExitPlsqlDeclarations(c *PlsqlDeclarationsContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionHeading is called when exiting the functionHeading production.
	ExitFunctionHeading(c *FunctionHeadingContext)

	// ExitParameterDeclaration is called when exiting the parameterDeclaration production.
	ExitParameterDeclaration(c *ParameterDeclarationContext)

	// ExitProcedureDeclaration is called when exiting the procedureDeclaration production.
	ExitProcedureDeclaration(c *ProcedureDeclarationContext)

	// ExitProcedureHeading is called when exiting the procedureHeading production.
	ExitProcedureHeading(c *ProcedureHeadingContext)

	// ExitProcedureProperties is called when exiting the procedureProperties production.
	ExitProcedureProperties(c *ProcedurePropertiesContext)

	// ExitAccessibleByClause is called when exiting the accessibleByClause production.
	ExitAccessibleByClause(c *AccessibleByClauseContext)

	// ExitAccessor is called when exiting the accessor production.
	ExitAccessor(c *AccessorContext)

	// ExitUnitKind is called when exiting the unitKind production.
	ExitUnitKind(c *UnitKindContext)

	// ExitDefaultCollationClause is called when exiting the defaultCollationClause production.
	ExitDefaultCollationClause(c *DefaultCollationClauseContext)

	// ExitCollationOption is called when exiting the collationOption production.
	ExitCollationOption(c *CollationOptionContext)

	// ExitInvokerRightsClause is called when exiting the invokerRightsClause production.
	ExitInvokerRightsClause(c *InvokerRightsClauseContext)

	// ExitSubqueryFactoringClause is called when exiting the subqueryFactoringClause production.
	ExitSubqueryFactoringClause(c *SubqueryFactoringClauseContext)

	// ExitSearchClause is called when exiting the searchClause production.
	ExitSearchClause(c *SearchClauseContext)

	// ExitCycleClause is called when exiting the cycleClause production.
	ExitCycleClause(c *CycleClauseContext)

	// ExitSubavFactoringClause is called when exiting the subavFactoringClause production.
	ExitSubavFactoringClause(c *SubavFactoringClauseContext)

	// ExitSubavClause is called when exiting the subavClause production.
	ExitSubavClause(c *SubavClauseContext)

	// ExitHierarchiesClause is called when exiting the hierarchiesClause production.
	ExitHierarchiesClause(c *HierarchiesClauseContext)

	// ExitFilterClauses is called when exiting the filterClauses production.
	ExitFilterClauses(c *FilterClausesContext)

	// ExitFilterClause is called when exiting the filterClause production.
	ExitFilterClause(c *FilterClauseContext)

	// ExitAddCalcsClause is called when exiting the addCalcsClause production.
	ExitAddCalcsClause(c *AddCalcsClauseContext)

	// ExitCalcMeasClause is called when exiting the calcMeasClause production.
	ExitCalcMeasClause(c *CalcMeasClauseContext)

	// ExitCalcMeasExpression is called when exiting the calcMeasExpression production.
	ExitCalcMeasExpression(c *CalcMeasExpressionContext)

	// ExitAvExpression is called when exiting the avExpression production.
	ExitAvExpression(c *AvExpressionContext)

	// ExitAvMeasExpression is called when exiting the avMeasExpression production.
	ExitAvMeasExpression(c *AvMeasExpressionContext)

	// ExitLeadLagExpression is called when exiting the leadLagExpression production.
	ExitLeadLagExpression(c *LeadLagExpressionContext)

	// ExitLeadLagFunctionName is called when exiting the leadLagFunctionName production.
	ExitLeadLagFunctionName(c *LeadLagFunctionNameContext)

	// ExitLeadLagClause is called when exiting the leadLagClause production.
	ExitLeadLagClause(c *LeadLagClauseContext)

	// ExitHierarchyRef is called when exiting the hierarchyRef production.
	ExitHierarchyRef(c *HierarchyRefContext)

	// ExitWindowExpression is called when exiting the windowExpression production.
	ExitWindowExpression(c *WindowExpressionContext)

	// ExitWindowClause is called when exiting the windowClause production.
	ExitWindowClause(c *WindowClauseContext)

	// ExitPrecedingBoundary is called when exiting the precedingBoundary production.
	ExitPrecedingBoundary(c *PrecedingBoundaryContext)

	// ExitFollowingBoundary is called when exiting the followingBoundary production.
	ExitFollowingBoundary(c *FollowingBoundaryContext)

	// ExitRankExpression is called when exiting the rankExpression production.
	ExitRankExpression(c *RankExpressionContext)

	// ExitRankFunctionName is called when exiting the rankFunctionName production.
	ExitRankFunctionName(c *RankFunctionNameContext)

	// ExitRankClause is called when exiting the rankClause production.
	ExitRankClause(c *RankClauseContext)

	// ExitCalcMeasOrderByClause is called when exiting the calcMeasOrderByClause production.
	ExitCalcMeasOrderByClause(c *CalcMeasOrderByClauseContext)

	// ExitShareOfExpression is called when exiting the shareOfExpression production.
	ExitShareOfExpression(c *ShareOfExpressionContext)

	// ExitShareClause is called when exiting the shareClause production.
	ExitShareClause(c *ShareClauseContext)

	// ExitMemberExpression is called when exiting the memberExpression production.
	ExitMemberExpression(c *MemberExpressionContext)

	// ExitLevelMemberLiteral is called when exiting the levelMemberLiteral production.
	ExitLevelMemberLiteral(c *LevelMemberLiteralContext)

	// ExitPosMemberKeys is called when exiting the posMemberKeys production.
	ExitPosMemberKeys(c *PosMemberKeysContext)

	// ExitNamedMemberKeys is called when exiting the namedMemberKeys production.
	ExitNamedMemberKeys(c *NamedMemberKeysContext)

	// ExitHierNavigationExpression is called when exiting the hierNavigationExpression production.
	ExitHierNavigationExpression(c *HierNavigationExpressionContext)

	// ExitHierAncestorExpression is called when exiting the hierAncestorExpression production.
	ExitHierAncestorExpression(c *HierAncestorExpressionContext)

	// ExitHierParentExpression is called when exiting the hierParentExpression production.
	ExitHierParentExpression(c *HierParentExpressionContext)

	// ExitHierLeadLagExpression is called when exiting the hierLeadLagExpression production.
	ExitHierLeadLagExpression(c *HierLeadLagExpressionContext)

	// ExitHierLeadLagClause is called when exiting the hierLeadLagClause production.
	ExitHierLeadLagClause(c *HierLeadLagClauseContext)

	// ExitQdrExpression is called when exiting the qdrExpression production.
	ExitQdrExpression(c *QdrExpressionContext)

	// ExitQualifier is called when exiting the qualifier production.
	ExitQualifier(c *QualifierContext)

	// ExitAvHierExpression is called when exiting the avHierExpression production.
	ExitAvHierExpression(c *AvHierExpressionContext)

	// ExitHierFunctionName is called when exiting the hierFunctionName production.
	ExitHierFunctionName(c *HierFunctionNameContext)

	// ExitDuplicateSpecification is called when exiting the duplicateSpecification production.
	ExitDuplicateSpecification(c *DuplicateSpecificationContext)

	// ExitUnqualifiedShorthand is called when exiting the unqualifiedShorthand production.
	ExitUnqualifiedShorthand(c *UnqualifiedShorthandContext)

	// ExitSelectList is called when exiting the selectList production.
	ExitSelectList(c *SelectListContext)

	// ExitSelectProjection is called when exiting the selectProjection production.
	ExitSelectProjection(c *SelectProjectionContext)

	// ExitSelectProjectionExprClause is called when exiting the selectProjectionExprClause production.
	ExitSelectProjectionExprClause(c *SelectProjectionExprClauseContext)

	// ExitSelectFromClause is called when exiting the selectFromClause production.
	ExitSelectFromClause(c *SelectFromClauseContext)

	// ExitFromClauseList is called when exiting the fromClauseList production.
	ExitFromClauseList(c *FromClauseListContext)

	// ExitFromClauseOption is called when exiting the fromClauseOption production.
	ExitFromClauseOption(c *FromClauseOptionContext)

	// ExitSelectTableReference is called when exiting the selectTableReference production.
	ExitSelectTableReference(c *SelectTableReferenceContext)

	// ExitQueryTableExprClause is called when exiting the queryTableExprClause production.
	ExitQueryTableExprClause(c *QueryTableExprClauseContext)

	// ExitFlashbackQueryClause is called when exiting the flashbackQueryClause production.
	ExitFlashbackQueryClause(c *FlashbackQueryClauseContext)

	// ExitQueryTableExpr is called when exiting the queryTableExpr production.
	ExitQueryTableExpr(c *QueryTableExprContext)

	// ExitLateralClause is called when exiting the lateralClause production.
	ExitLateralClause(c *LateralClauseContext)

	// ExitQueryTableExprSampleClause is called when exiting the queryTableExprSampleClause production.
	ExitQueryTableExprSampleClause(c *QueryTableExprSampleClauseContext)

	// ExitQueryTableExprTableClause is called when exiting the queryTableExprTableClause production.
	ExitQueryTableExprTableClause(c *QueryTableExprTableClauseContext)

	// ExitQueryTableExprViewClause is called when exiting the queryTableExprViewClause production.
	ExitQueryTableExprViewClause(c *QueryTableExprViewClauseContext)

	// ExitQueryTableExprAnalyticClause is called when exiting the queryTableExprAnalyticClause production.
	ExitQueryTableExprAnalyticClause(c *QueryTableExprAnalyticClauseContext)

	// ExitInlineExternalTable is called when exiting the inlineExternalTable production.
	ExitInlineExternalTable(c *InlineExternalTableContext)

	// ExitInlineExternalTableProperties is called when exiting the inlineExternalTableProperties production.
	ExitInlineExternalTableProperties(c *InlineExternalTablePropertiesContext)

	// ExitExternalTableDataProperties is called when exiting the externalTableDataProperties production.
	ExitExternalTableDataProperties(c *ExternalTableDataPropertiesContext)

	// ExitMofifiedExternalTable is called when exiting the mofifiedExternalTable production.
	ExitMofifiedExternalTable(c *MofifiedExternalTableContext)

	// ExitModifyExternalTableProperties is called when exiting the modifyExternalTableProperties production.
	ExitModifyExternalTableProperties(c *ModifyExternalTablePropertiesContext)

	// ExitPivotClause is called when exiting the pivotClause production.
	ExitPivotClause(c *PivotClauseContext)

	// ExitPivotForClause is called when exiting the pivotForClause production.
	ExitPivotForClause(c *PivotForClauseContext)

	// ExitPivotInClause is called when exiting the pivotInClause production.
	ExitPivotInClause(c *PivotInClauseContext)

	// ExitUnpivotClause is called when exiting the unpivotClause production.
	ExitUnpivotClause(c *UnpivotClauseContext)

	// ExitUnpivotInClause is called when exiting the unpivotInClause production.
	ExitUnpivotInClause(c *UnpivotInClauseContext)

	// ExitSampleClause is called when exiting the sampleClause production.
	ExitSampleClause(c *SampleClauseContext)

	// ExitContainersClause is called when exiting the containersClause production.
	ExitContainersClause(c *ContainersClauseContext)

	// ExitShardsClause is called when exiting the shardsClause production.
	ExitShardsClause(c *ShardsClauseContext)

	// ExitJoinClause is called when exiting the joinClause production.
	ExitJoinClause(c *JoinClauseContext)

	// ExitSelectJoinOption is called when exiting the selectJoinOption production.
	ExitSelectJoinOption(c *SelectJoinOptionContext)

	// ExitInnerCrossJoinClause is called when exiting the innerCrossJoinClause production.
	ExitInnerCrossJoinClause(c *InnerCrossJoinClauseContext)

	// ExitSelectJoinSpecification is called when exiting the selectJoinSpecification production.
	ExitSelectJoinSpecification(c *SelectJoinSpecificationContext)

	// ExitOuterJoinClause is called when exiting the outerJoinClause production.
	ExitOuterJoinClause(c *OuterJoinClauseContext)

	// ExitQueryPartitionClause is called when exiting the queryPartitionClause production.
	ExitQueryPartitionClause(c *QueryPartitionClauseContext)

	// ExitOuterJoinType is called when exiting the outerJoinType production.
	ExitOuterJoinType(c *OuterJoinTypeContext)

	// ExitCrossOuterApplyClause is called when exiting the crossOuterApplyClause production.
	ExitCrossOuterApplyClause(c *CrossOuterApplyClauseContext)

	// ExitInlineAnalyticView is called when exiting the inlineAnalyticView production.
	ExitInlineAnalyticView(c *InlineAnalyticViewContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitHierarchicalQueryClause is called when exiting the hierarchicalQueryClause production.
	ExitHierarchicalQueryClause(c *HierarchicalQueryClauseContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitGroupByItem is called when exiting the groupByItem production.
	ExitGroupByItem(c *GroupByItemContext)

	// ExitRollupCubeClause is called when exiting the rollupCubeClause production.
	ExitRollupCubeClause(c *RollupCubeClauseContext)

	// ExitGroupingSetsClause is called when exiting the groupingSetsClause production.
	ExitGroupingSetsClause(c *GroupingSetsClauseContext)

	// ExitGroupingExprList is called when exiting the groupingExprList production.
	ExitGroupingExprList(c *GroupingExprListContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitModelClause is called when exiting the modelClause production.
	ExitModelClause(c *ModelClauseContext)

	// ExitCellReferenceOptions is called when exiting the cellReferenceOptions production.
	ExitCellReferenceOptions(c *CellReferenceOptionsContext)

	// ExitReturnRowsClause is called when exiting the returnRowsClause production.
	ExitReturnRowsClause(c *ReturnRowsClauseContext)

	// ExitReferenceModel is called when exiting the referenceModel production.
	ExitReferenceModel(c *ReferenceModelContext)

	// ExitMainModel is called when exiting the mainModel production.
	ExitMainModel(c *MainModelContext)

	// ExitModelColumnClauses is called when exiting the modelColumnClauses production.
	ExitModelColumnClauses(c *ModelColumnClausesContext)

	// ExitModelRulesClause is called when exiting the modelRulesClause production.
	ExitModelRulesClause(c *ModelRulesClauseContext)

	// ExitModelIterateClause is called when exiting the modelIterateClause production.
	ExitModelIterateClause(c *ModelIterateClauseContext)

	// ExitCellAssignment is called when exiting the cellAssignment production.
	ExitCellAssignment(c *CellAssignmentContext)

	// ExitSingleColumnForLoop is called when exiting the singleColumnForLoop production.
	ExitSingleColumnForLoop(c *SingleColumnForLoopContext)

	// ExitMultiColumnForLoop is called when exiting the multiColumnForLoop production.
	ExitMultiColumnForLoop(c *MultiColumnForLoopContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitModelExpr is called when exiting the modelExpr production.
	ExitModelExpr(c *ModelExprContext)

	// ExitAnalyticFunction is called when exiting the analyticFunction production.
	ExitAnalyticFunction(c *AnalyticFunctionContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitForUpdateClause is called when exiting the forUpdateClause production.
	ExitForUpdateClause(c *ForUpdateClauseContext)

	// ExitForUpdateClauseList is called when exiting the forUpdateClauseList production.
	ExitForUpdateClauseList(c *ForUpdateClauseListContext)

	// ExitForUpdateClauseOption is called when exiting the forUpdateClauseOption production.
	ExitForUpdateClauseOption(c *ForUpdateClauseOptionContext)

	// ExitRowLimitingClause is called when exiting the rowLimitingClause production.
	ExitRowLimitingClause(c *RowLimitingClauseContext)

	// ExitMerge is called when exiting the merge production.
	ExitMerge(c *MergeContext)

	// ExitHint is called when exiting the hint production.
	ExitHint(c *HintContext)

	// ExitIntoClause is called when exiting the intoClause production.
	ExitIntoClause(c *IntoClauseContext)

	// ExitUsingClause is called when exiting the usingClause production.
	ExitUsingClause(c *UsingClauseContext)

	// ExitMergeUpdateClause is called when exiting the mergeUpdateClause production.
	ExitMergeUpdateClause(c *MergeUpdateClauseContext)

	// ExitMergeSetAssignmentsClause is called when exiting the mergeSetAssignmentsClause production.
	ExitMergeSetAssignmentsClause(c *MergeSetAssignmentsClauseContext)

	// ExitMergeAssignment is called when exiting the mergeAssignment production.
	ExitMergeAssignment(c *MergeAssignmentContext)

	// ExitMergeAssignmentValue is called when exiting the mergeAssignmentValue production.
	ExitMergeAssignmentValue(c *MergeAssignmentValueContext)

	// ExitDeleteWhereClause is called when exiting the deleteWhereClause production.
	ExitDeleteWhereClause(c *DeleteWhereClauseContext)

	// ExitMergeInsertClause is called when exiting the mergeInsertClause production.
	ExitMergeInsertClause(c *MergeInsertClauseContext)

	// ExitMergeInsertColumn is called when exiting the mergeInsertColumn production.
	ExitMergeInsertColumn(c *MergeInsertColumnContext)

	// ExitMergeColumnValue is called when exiting the mergeColumnValue production.
	ExitMergeColumnValue(c *MergeColumnValueContext)

	// ExitErrorLoggingClause is called when exiting the errorLoggingClause production.
	ExitErrorLoggingClause(c *ErrorLoggingClauseContext)

	// ExitRowPatternClause is called when exiting the rowPatternClause production.
	ExitRowPatternClause(c *RowPatternClauseContext)

	// ExitRowPatternPartitionBy is called when exiting the rowPatternPartitionBy production.
	ExitRowPatternPartitionBy(c *RowPatternPartitionByContext)

	// ExitRowPatternOrderBy is called when exiting the rowPatternOrderBy production.
	ExitRowPatternOrderBy(c *RowPatternOrderByContext)

	// ExitRowPatternMeasures is called when exiting the rowPatternMeasures production.
	ExitRowPatternMeasures(c *RowPatternMeasuresContext)

	// ExitRowPatternMeasureColumn is called when exiting the rowPatternMeasureColumn production.
	ExitRowPatternMeasureColumn(c *RowPatternMeasureColumnContext)

	// ExitRowPatternRowsPerMatch is called when exiting the rowPatternRowsPerMatch production.
	ExitRowPatternRowsPerMatch(c *RowPatternRowsPerMatchContext)

	// ExitRowPatternSkipTo is called when exiting the rowPatternSkipTo production.
	ExitRowPatternSkipTo(c *RowPatternSkipToContext)

	// ExitRowPattern is called when exiting the rowPattern production.
	ExitRowPattern(c *RowPatternContext)

	// ExitRowPatternTerm is called when exiting the rowPatternTerm production.
	ExitRowPatternTerm(c *RowPatternTermContext)

	// ExitRowPatternFactor is called when exiting the rowPatternFactor production.
	ExitRowPatternFactor(c *RowPatternFactorContext)

	// ExitRowPatternPrimary is called when exiting the rowPatternPrimary production.
	ExitRowPatternPrimary(c *RowPatternPrimaryContext)

	// ExitRowPatternPermute is called when exiting the rowPatternPermute production.
	ExitRowPatternPermute(c *RowPatternPermuteContext)

	// ExitRowPatternQuantifier is called when exiting the rowPatternQuantifier production.
	ExitRowPatternQuantifier(c *RowPatternQuantifierContext)

	// ExitRowPatternSubsetClause is called when exiting the rowPatternSubsetClause production.
	ExitRowPatternSubsetClause(c *RowPatternSubsetClauseContext)

	// ExitRowPatternSubsetItem is called when exiting the rowPatternSubsetItem production.
	ExitRowPatternSubsetItem(c *RowPatternSubsetItemContext)

	// ExitRowPatternDefinitionList is called when exiting the rowPatternDefinitionList production.
	ExitRowPatternDefinitionList(c *RowPatternDefinitionListContext)

	// ExitRowPatternDefinition is called when exiting the rowPatternDefinition production.
	ExitRowPatternDefinition(c *RowPatternDefinitionContext)

	// ExitRowPatternRecFunc is called when exiting the rowPatternRecFunc production.
	ExitRowPatternRecFunc(c *RowPatternRecFuncContext)

	// ExitPatternMeasExpression is called when exiting the patternMeasExpression production.
	ExitPatternMeasExpression(c *PatternMeasExpressionContext)

	// ExitRowPatternClassifierFunc is called when exiting the rowPatternClassifierFunc production.
	ExitRowPatternClassifierFunc(c *RowPatternClassifierFuncContext)

	// ExitRowPatternMatchNumFunc is called when exiting the rowPatternMatchNumFunc production.
	ExitRowPatternMatchNumFunc(c *RowPatternMatchNumFuncContext)

	// ExitRowPatternNavigationFunc is called when exiting the rowPatternNavigationFunc production.
	ExitRowPatternNavigationFunc(c *RowPatternNavigationFuncContext)

	// ExitRowPatternNavLogical is called when exiting the rowPatternNavLogical production.
	ExitRowPatternNavLogical(c *RowPatternNavLogicalContext)

	// ExitRowPatternNavPhysical is called when exiting the rowPatternNavPhysical production.
	ExitRowPatternNavPhysical(c *RowPatternNavPhysicalContext)

	// ExitRowPatternNavCompound is called when exiting the rowPatternNavCompound production.
	ExitRowPatternNavCompound(c *RowPatternNavCompoundContext)

	// ExitRowPatternAggregateFunc is called when exiting the rowPatternAggregateFunc production.
	ExitRowPatternAggregateFunc(c *RowPatternAggregateFuncContext)

	// ExitParameterMarker is called when exiting the parameterMarker production.
	ExitParameterMarker(c *ParameterMarkerContext)

	// ExitLiterals is called when exiting the literals production.
	ExitLiterals(c *LiteralsContext)

	// ExitStringLiterals is called when exiting the stringLiterals production.
	ExitStringLiterals(c *StringLiteralsContext)

	// ExitNumberLiterals is called when exiting the numberLiterals production.
	ExitNumberLiterals(c *NumberLiteralsContext)

	// ExitDateTimeLiterals is called when exiting the dateTimeLiterals production.
	ExitDateTimeLiterals(c *DateTimeLiteralsContext)

	// ExitHexadecimalLiterals is called when exiting the hexadecimalLiterals production.
	ExitHexadecimalLiterals(c *HexadecimalLiteralsContext)

	// ExitBitValueLiterals is called when exiting the bitValueLiterals production.
	ExitBitValueLiterals(c *BitValueLiteralsContext)

	// ExitBooleanLiterals is called when exiting the booleanLiterals production.
	ExitBooleanLiterals(c *BooleanLiteralsContext)

	// ExitNullValueLiterals is called when exiting the nullValueLiterals production.
	ExitNullValueLiterals(c *NullValueLiteralsContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitUnreservedWord is called when exiting the unreservedWord production.
	ExitUnreservedWord(c *UnreservedWordContext)

	// ExitSchemaName is called when exiting the schemaName production.
	ExitSchemaName(c *SchemaNameContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitViewName is called when exiting the viewName production.
	ExitViewName(c *ViewNameContext)

	// ExitMaterializedViewName is called when exiting the materializedViewName production.
	ExitMaterializedViewName(c *MaterializedViewNameContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitObjectName is called when exiting the objectName production.
	ExitObjectName(c *ObjectNameContext)

	// ExitClusterName is called when exiting the clusterName production.
	ExitClusterName(c *ClusterNameContext)

	// ExitIndexName is called when exiting the indexName production.
	ExitIndexName(c *IndexNameContext)

	// ExitStatisticsTypeName is called when exiting the statisticsTypeName production.
	ExitStatisticsTypeName(c *StatisticsTypeNameContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitPackageName is called when exiting the packageName production.
	ExitPackageName(c *PackageNameContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitIndextypeName is called when exiting the indextypeName production.
	ExitIndextypeName(c *IndextypeNameContext)

	// ExitModelName is called when exiting the modelName production.
	ExitModelName(c *ModelNameContext)

	// ExitOperatorName is called when exiting the operatorName production.
	ExitOperatorName(c *OperatorNameContext)

	// ExitConstraintName is called when exiting the constraintName production.
	ExitConstraintName(c *ConstraintNameContext)

	// ExitSavepointName is called when exiting the savepointName production.
	ExitSavepointName(c *SavepointNameContext)

	// ExitSynonymName is called when exiting the synonymName production.
	ExitSynonymName(c *SynonymNameContext)

	// ExitOwner is called when exiting the owner production.
	ExitOwner(c *OwnerContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitTablespaceName is called when exiting the tablespaceName production.
	ExitTablespaceName(c *TablespaceNameContext)

	// ExitTablespaceSetName is called when exiting the tablespaceSetName production.
	ExitTablespaceSetName(c *TablespaceSetNameContext)

	// ExitServiceName is called when exiting the serviceName production.
	ExitServiceName(c *ServiceNameContext)

	// ExitIlmPolicyName is called when exiting the ilmPolicyName production.
	ExitIlmPolicyName(c *IlmPolicyNameContext)

	// ExitPolicyName is called when exiting the policyName production.
	ExitPolicyName(c *PolicyNameContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitDbLink is called when exiting the dbLink production.
	ExitDbLink(c *DbLinkContext)

	// ExitParameterValue is called when exiting the parameterValue production.
	ExitParameterValue(c *ParameterValueContext)

	// ExitDirectoryName is called when exiting the directoryName production.
	ExitDirectoryName(c *DirectoryNameContext)

	// ExitDispatcherName is called when exiting the dispatcherName production.
	ExitDispatcherName(c *DispatcherNameContext)

	// ExitClientId is called when exiting the clientId production.
	ExitClientId(c *ClientIdContext)

	// ExitOpaqueFormatSpec is called when exiting the opaqueFormatSpec production.
	ExitOpaqueFormatSpec(c *OpaqueFormatSpecContext)

	// ExitAccessDriverType is called when exiting the accessDriverType production.
	ExitAccessDriverType(c *AccessDriverTypeContext)

	// ExitVarrayItem is called when exiting the varrayItem production.
	ExitVarrayItem(c *VarrayItemContext)

	// ExitNestedItem is called when exiting the nestedItem production.
	ExitNestedItem(c *NestedItemContext)

	// ExitStorageTable is called when exiting the storageTable production.
	ExitStorageTable(c *StorageTableContext)

	// ExitLobSegname is called when exiting the lobSegname production.
	ExitLobSegname(c *LobSegnameContext)

	// ExitLocationSpecifier is called when exiting the locationSpecifier production.
	ExitLocationSpecifier(c *LocationSpecifierContext)

	// ExitXmlSchemaURLName is called when exiting the xmlSchemaURLName production.
	ExitXmlSchemaURLName(c *XmlSchemaURLNameContext)

	// ExitElementName is called when exiting the elementName production.
	ExitElementName(c *ElementNameContext)

	// ExitSubpartitionName is called when exiting the subpartitionName production.
	ExitSubpartitionName(c *SubpartitionNameContext)

	// ExitParameterName is called when exiting the parameterName production.
	ExitParameterName(c *ParameterNameContext)

	// ExitEditionName is called when exiting the editionName production.
	ExitEditionName(c *EditionNameContext)

	// ExitContainerName is called when exiting the containerName production.
	ExitContainerName(c *ContainerNameContext)

	// ExitPartitionName is called when exiting the partitionName production.
	ExitPartitionName(c *PartitionNameContext)

	// ExitPartitionSetName is called when exiting the partitionSetName production.
	ExitPartitionSetName(c *PartitionSetNameContext)

	// ExitPartitionKeyValue is called when exiting the partitionKeyValue production.
	ExitPartitionKeyValue(c *PartitionKeyValueContext)

	// ExitSubpartitionKeyValue is called when exiting the subpartitionKeyValue production.
	ExitSubpartitionKeyValue(c *SubpartitionKeyValueContext)

	// ExitZonemapName is called when exiting the zonemapName production.
	ExitZonemapName(c *ZonemapNameContext)

	// ExitFlashbackArchiveName is called when exiting the flashbackArchiveName production.
	ExitFlashbackArchiveName(c *FlashbackArchiveNameContext)

	// ExitRoleName is called when exiting the roleName production.
	ExitRoleName(c *RoleNameContext)

	// ExitUserName is called when exiting the userName production.
	ExitUserName(c *UserNameContext)

	// ExitPassword is called when exiting the password production.
	ExitPassword(c *PasswordContext)

	// ExitLogGroupName is called when exiting the logGroupName production.
	ExitLogGroupName(c *LogGroupNameContext)

	// ExitColumnNames is called when exiting the columnNames production.
	ExitColumnNames(c *ColumnNamesContext)

	// ExitTableNames is called when exiting the tableNames production.
	ExitTableNames(c *TableNamesContext)

	// ExitOracleId is called when exiting the oracleId production.
	ExitOracleId(c *OracleIdContext)

	// ExitCollationName is called when exiting the collationName production.
	ExitCollationName(c *CollationNameContext)

	// ExitColumnCollationName is called when exiting the columnCollationName production.
	ExitColumnCollationName(c *ColumnCollationNameContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitDataTypeLength is called when exiting the dataTypeLength production.
	ExitDataTypeLength(c *DataTypeLengthContext)

	// ExitPrimaryKey is called when exiting the primaryKey production.
	ExitPrimaryKey(c *PrimaryKeyContext)

	// ExitExprs is called when exiting the exprs production.
	ExitExprs(c *ExprsContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitAndOperator is called when exiting the andOperator production.
	ExitAndOperator(c *AndOperatorContext)

	// ExitOrOperator is called when exiting the orOperator production.
	ExitOrOperator(c *OrOperatorContext)

	// ExitNotOperator is called when exiting the notOperator production.
	ExitNotOperator(c *NotOperatorContext)

	// ExitBooleanPrimary is called when exiting the booleanPrimary production.
	ExitBooleanPrimary(c *BooleanPrimaryContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitBitExpr is called when exiting the bitExpr production.
	ExitBitExpr(c *BitExprContext)

	// ExitSimpleExpr is called when exiting the simpleExpr production.
	ExitSimpleExpr(c *SimpleExprContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitAggregationFunction is called when exiting the aggregationFunction production.
	ExitAggregationFunction(c *AggregationFunctionContext)

	// ExitAggregationFunctionName is called when exiting the aggregationFunctionName production.
	ExitAggregationFunctionName(c *AggregationFunctionNameContext)

	// ExitAnalyticClause is called when exiting the analyticClause production.
	ExitAnalyticClause(c *AnalyticClauseContext)

	// ExitWindowingClause is called when exiting the windowingClause production.
	ExitWindowingClause(c *WindowingClauseContext)

	// ExitSpecialFunction is called when exiting the specialFunction production.
	ExitSpecialFunction(c *SpecialFunctionContext)

	// ExitCastFunction is called when exiting the castFunction production.
	ExitCastFunction(c *CastFunctionContext)

	// ExitCharFunction is called when exiting the charFunction production.
	ExitCharFunction(c *CharFunctionContext)

	// ExitRegularFunction is called when exiting the regularFunction production.
	ExitRegularFunction(c *RegularFunctionContext)

	// ExitRegularFunctionName is called when exiting the regularFunctionName production.
	ExitRegularFunctionName(c *RegularFunctionNameContext)

	// ExitCaseExpression is called when exiting the caseExpression production.
	ExitCaseExpression(c *CaseExpressionContext)

	// ExitCaseWhen is called when exiting the caseWhen production.
	ExitCaseWhen(c *CaseWhenContext)

	// ExitCaseElse is called when exiting the caseElse production.
	ExitCaseElse(c *CaseElseContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderByItem is called when exiting the orderByItem production.
	ExitOrderByItem(c *OrderByItemContext)

	// ExitAttributeName is called when exiting the attributeName production.
	ExitAttributeName(c *AttributeNameContext)

	// ExitSimpleExprs is called when exiting the simpleExprs production.
	ExitSimpleExprs(c *SimpleExprsContext)

	// ExitLobItem is called when exiting the lobItem production.
	ExitLobItem(c *LobItemContext)

	// ExitLobItems is called when exiting the lobItems production.
	ExitLobItems(c *LobItemsContext)

	// ExitLobItemList is called when exiting the lobItemList production.
	ExitLobItemList(c *LobItemListContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitSpecialDatatype is called when exiting the specialDatatype production.
	ExitSpecialDatatype(c *SpecialDatatypeContext)

	// ExitDataTypeName is called when exiting the dataTypeName production.
	ExitDataTypeName(c *DataTypeNameContext)

	// ExitDatetimeTypeSuffix is called when exiting the datetimeTypeSuffix production.
	ExitDatetimeTypeSuffix(c *DatetimeTypeSuffixContext)

	// ExitTreatFunction is called when exiting the treatFunction production.
	ExitTreatFunction(c *TreatFunctionContext)

	// ExitPrivateExprOfDb is called when exiting the privateExprOfDb production.
	ExitPrivateExprOfDb(c *PrivateExprOfDbContext)

	// ExitCaseExpr is called when exiting the caseExpr production.
	ExitCaseExpr(c *CaseExprContext)

	// ExitSimpleCaseExpr is called when exiting the simpleCaseExpr production.
	ExitSimpleCaseExpr(c *SimpleCaseExprContext)

	// ExitSearchedCaseExpr is called when exiting the searchedCaseExpr production.
	ExitSearchedCaseExpr(c *SearchedCaseExprContext)

	// ExitElseClause is called when exiting the elseClause production.
	ExitElseClause(c *ElseClauseContext)

	// ExitIntervalExpression is called when exiting the intervalExpression production.
	ExitIntervalExpression(c *IntervalExpressionContext)

	// ExitObjectAccessExpression is called when exiting the objectAccessExpression production.
	ExitObjectAccessExpression(c *ObjectAccessExpressionContext)

	// ExitConstructorExpr is called when exiting the constructorExpr production.
	ExitConstructorExpr(c *ConstructorExprContext)

	// ExitIgnoredIdentifier is called when exiting the ignoredIdentifier production.
	ExitIgnoredIdentifier(c *IgnoredIdentifierContext)

	// ExitIgnoredIdentifiers is called when exiting the ignoredIdentifiers production.
	ExitIgnoredIdentifiers(c *IgnoredIdentifiersContext)

	// ExitMatchNone is called when exiting the matchNone production.
	ExitMatchNone(c *MatchNoneContext)

	// ExitHashSubpartitionQuantity is called when exiting the hashSubpartitionQuantity production.
	ExitHashSubpartitionQuantity(c *HashSubpartitionQuantityContext)

	// ExitOdciParameters is called when exiting the odciParameters production.
	ExitOdciParameters(c *OdciParametersContext)

	// ExitDatabaseName is called when exiting the databaseName production.
	ExitDatabaseName(c *DatabaseNameContext)

	// ExitLocationName is called when exiting the locationName production.
	ExitLocationName(c *LocationNameContext)

	// ExitFileName is called when exiting the fileName production.
	ExitFileName(c *FileNameContext)

	// ExitAsmFileName is called when exiting the asmFileName production.
	ExitAsmFileName(c *AsmFileNameContext)

	// ExitFileNumber is called when exiting the fileNumber production.
	ExitFileNumber(c *FileNumberContext)

	// ExitInstanceName is called when exiting the instanceName production.
	ExitInstanceName(c *InstanceNameContext)

	// ExitLogminerSessionName is called when exiting the logminerSessionName production.
	ExitLogminerSessionName(c *LogminerSessionNameContext)

	// ExitTablespaceGroupName is called when exiting the tablespaceGroupName production.
	ExitTablespaceGroupName(c *TablespaceGroupNameContext)

	// ExitCopyName is called when exiting the copyName production.
	ExitCopyName(c *CopyNameContext)

	// ExitMirrorName is called when exiting the mirrorName production.
	ExitMirrorName(c *MirrorNameContext)

	// ExitUriString is called when exiting the uriString production.
	ExitUriString(c *UriStringContext)

	// ExitQualifiedCredentialName is called when exiting the qualifiedCredentialName production.
	ExitQualifiedCredentialName(c *QualifiedCredentialNameContext)

	// ExitPdbName is called when exiting the pdbName production.
	ExitPdbName(c *PdbNameContext)

	// ExitDiskgroupName is called when exiting the diskgroupName production.
	ExitDiskgroupName(c *DiskgroupNameContext)

	// ExitTemplateName is called when exiting the templateName production.
	ExitTemplateName(c *TemplateNameContext)

	// ExitAliasName is called when exiting the aliasName production.
	ExitAliasName(c *AliasNameContext)

	// ExitDomain is called when exiting the domain production.
	ExitDomain(c *DomainContext)

	// ExitDateValue is called when exiting the dateValue production.
	ExitDateValue(c *DateValueContext)

	// ExitSessionId is called when exiting the sessionId production.
	ExitSessionId(c *SessionIdContext)

	// ExitSerialNumber is called when exiting the serialNumber production.
	ExitSerialNumber(c *SerialNumberContext)

	// ExitInstanceId is called when exiting the instanceId production.
	ExitInstanceId(c *InstanceIdContext)

	// ExitSqlId is called when exiting the sqlId production.
	ExitSqlId(c *SqlIdContext)

	// ExitLogFileName is called when exiting the logFileName production.
	ExitLogFileName(c *LogFileNameContext)

	// ExitLogFileGroupsArchivedLocationName is called when exiting the logFileGroupsArchivedLocationName production.
	ExitLogFileGroupsArchivedLocationName(c *LogFileGroupsArchivedLocationNameContext)

	// ExitAsmVersion is called when exiting the asmVersion production.
	ExitAsmVersion(c *AsmVersionContext)

	// ExitWalletPassword is called when exiting the walletPassword production.
	ExitWalletPassword(c *WalletPasswordContext)

	// ExitHsmAuthString is called when exiting the hsmAuthString production.
	ExitHsmAuthString(c *HsmAuthStringContext)

	// ExitTargetDbName is called when exiting the targetDbName production.
	ExitTargetDbName(c *TargetDbNameContext)

	// ExitCertificateId is called when exiting the certificateId production.
	ExitCertificateId(c *CertificateIdContext)

	// ExitCategoryName is called when exiting the categoryName production.
	ExitCategoryName(c *CategoryNameContext)

	// ExitOffset is called when exiting the offset production.
	ExitOffset(c *OffsetContext)

	// ExitRowcount is called when exiting the rowcount production.
	ExitRowcount(c *RowcountContext)

	// ExitPercent is called when exiting the percent production.
	ExitPercent(c *PercentContext)

	// ExitRollbackSegment is called when exiting the rollbackSegment production.
	ExitRollbackSegment(c *RollbackSegmentContext)

	// ExitQueryName is called when exiting the queryName production.
	ExitQueryName(c *QueryNameContext)

	// ExitCycleValue is called when exiting the cycleValue production.
	ExitCycleValue(c *CycleValueContext)

	// ExitNoCycleValue is called when exiting the noCycleValue production.
	ExitNoCycleValue(c *NoCycleValueContext)

	// ExitOrderingColumn is called when exiting the orderingColumn production.
	ExitOrderingColumn(c *OrderingColumnContext)

	// ExitSubavName is called when exiting the subavName production.
	ExitSubavName(c *SubavNameContext)

	// ExitBaseAvName is called when exiting the baseAvName production.
	ExitBaseAvName(c *BaseAvNameContext)

	// ExitMeasName is called when exiting the measName production.
	ExitMeasName(c *MeasNameContext)

	// ExitLevelRef is called when exiting the levelRef production.
	ExitLevelRef(c *LevelRefContext)

	// ExitOffsetExpr is called when exiting the offsetExpr production.
	ExitOffsetExpr(c *OffsetExprContext)

	// ExitMemberKeyExpr is called when exiting the memberKeyExpr production.
	ExitMemberKeyExpr(c *MemberKeyExprContext)

	// ExitDepthExpression is called when exiting the depthExpression production.
	ExitDepthExpression(c *DepthExpressionContext)

	// ExitUnitName is called when exiting the unitName production.
	ExitUnitName(c *UnitNameContext)

	// ExitProcedureName is called when exiting the procedureName production.
	ExitProcedureName(c *ProcedureNameContext)

	// ExitCpuCost is called when exiting the cpuCost production.
	ExitCpuCost(c *CpuCostContext)

	// ExitIoCost is called when exiting the ioCost production.
	ExitIoCost(c *IoCostContext)

	// ExitNetworkCost is called when exiting the networkCost production.
	ExitNetworkCost(c *NetworkCostContext)

	// ExitDefaultSelectivity is called when exiting the defaultSelectivity production.
	ExitDefaultSelectivity(c *DefaultSelectivityContext)

	// ExitDataItem is called when exiting the dataItem production.
	ExitDataItem(c *DataItemContext)

	// ExitVariableName is called when exiting the variableName production.
	ExitVariableName(c *VariableNameContext)

	// ExitValidTimeColumn is called when exiting the validTimeColumn production.
	ExitValidTimeColumn(c *ValidTimeColumnContext)

	// ExitAttrDim is called when exiting the attrDim production.
	ExitAttrDim(c *AttrDimContext)

	// ExitHierarchyName is called when exiting the hierarchyName production.
	ExitHierarchyName(c *HierarchyNameContext)

	// ExitAnalyticViewName is called when exiting the analyticViewName production.
	ExitAnalyticViewName(c *AnalyticViewNameContext)

	// ExitSamplePercent is called when exiting the samplePercent production.
	ExitSamplePercent(c *SamplePercentContext)

	// ExitSeedValue is called when exiting the seedValue production.
	ExitSeedValue(c *SeedValueContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitRestorePoint is called when exiting the restorePoint production.
	ExitRestorePoint(c *RestorePointContext)

	// ExitScnValue is called when exiting the scnValue production.
	ExitScnValue(c *ScnValueContext)

	// ExitTimestampValue is called when exiting the timestampValue production.
	ExitTimestampValue(c *TimestampValueContext)

	// ExitScnTimestampExpr is called when exiting the scnTimestampExpr production.
	ExitScnTimestampExpr(c *ScnTimestampExprContext)

	// ExitReferenceModelName is called when exiting the referenceModelName production.
	ExitReferenceModelName(c *ReferenceModelNameContext)

	// ExitMainModelName is called when exiting the mainModelName production.
	ExitMainModelName(c *MainModelNameContext)

	// ExitMeasureColumn is called when exiting the measureColumn production.
	ExitMeasureColumn(c *MeasureColumnContext)

	// ExitDimensionColumn is called when exiting the dimensionColumn production.
	ExitDimensionColumn(c *DimensionColumnContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitAnalyticFunctionName is called when exiting the analyticFunctionName production.
	ExitAnalyticFunctionName(c *AnalyticFunctionNameContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitComparisonCondition is called when exiting the comparisonCondition production.
	ExitComparisonCondition(c *ComparisonConditionContext)

	// ExitSimpleComparisonCondition is called when exiting the simpleComparisonCondition production.
	ExitSimpleComparisonCondition(c *SimpleComparisonConditionContext)

	// ExitGroupComparisonCondition is called when exiting the groupComparisonCondition production.
	ExitGroupComparisonCondition(c *GroupComparisonConditionContext)

	// ExitFloatingPointCondition is called when exiting the floatingPointCondition production.
	ExitFloatingPointCondition(c *FloatingPointConditionContext)

	// ExitLogicalCondition is called when exiting the logicalCondition production.
	ExitLogicalCondition(c *LogicalConditionContext)

	// ExitModelCondition is called when exiting the modelCondition production.
	ExitModelCondition(c *ModelConditionContext)

	// ExitIsAnyCondition is called when exiting the isAnyCondition production.
	ExitIsAnyCondition(c *IsAnyConditionContext)

	// ExitIsPresentCondition is called when exiting the isPresentCondition production.
	ExitIsPresentCondition(c *IsPresentConditionContext)

	// ExitCellReference is called when exiting the cellReference production.
	ExitCellReference(c *CellReferenceContext)

	// ExitMultisetCondition is called when exiting the multisetCondition production.
	ExitMultisetCondition(c *MultisetConditionContext)

	// ExitIsASetCondition is called when exiting the isASetCondition production.
	ExitIsASetCondition(c *IsASetConditionContext)

	// ExitIsEmptyCondition is called when exiting the isEmptyCondition production.
	ExitIsEmptyCondition(c *IsEmptyConditionContext)

	// ExitMemberCondition is called when exiting the memberCondition production.
	ExitMemberCondition(c *MemberConditionContext)

	// ExitSubmultisetCondition is called when exiting the submultisetCondition production.
	ExitSubmultisetCondition(c *SubmultisetConditionContext)

	// ExitPatternMatchingCondition is called when exiting the patternMatchingCondition production.
	ExitPatternMatchingCondition(c *PatternMatchingConditionContext)

	// ExitLikeCondition is called when exiting the likeCondition production.
	ExitLikeCondition(c *LikeConditionContext)

	// ExitSearchValue is called when exiting the searchValue production.
	ExitSearchValue(c *SearchValueContext)

	// ExitEscapeChar is called when exiting the escapeChar production.
	ExitEscapeChar(c *EscapeCharContext)

	// ExitRegexpLikeCondition is called when exiting the regexpLikeCondition production.
	ExitRegexpLikeCondition(c *RegexpLikeConditionContext)

	// ExitMatchParam is called when exiting the matchParam production.
	ExitMatchParam(c *MatchParamContext)

	// ExitRangeCondition is called when exiting the rangeCondition production.
	ExitRangeCondition(c *RangeConditionContext)

	// ExitNullCondition is called when exiting the nullCondition production.
	ExitNullCondition(c *NullConditionContext)

	// ExitXmlCondition is called when exiting the xmlCondition production.
	ExitXmlCondition(c *XmlConditionContext)

	// ExitEqualsPathCondition is called when exiting the equalsPathCondition production.
	ExitEqualsPathCondition(c *EqualsPathConditionContext)

	// ExitPathString is called when exiting the pathString production.
	ExitPathString(c *PathStringContext)

	// ExitCorrelationInteger is called when exiting the correlationInteger production.
	ExitCorrelationInteger(c *CorrelationIntegerContext)

	// ExitUnderPathCondition is called when exiting the underPathCondition production.
	ExitUnderPathCondition(c *UnderPathConditionContext)

	// ExitLevels is called when exiting the levels production.
	ExitLevels(c *LevelsContext)

	// ExitJsonCondition is called when exiting the jsonCondition production.
	ExitJsonCondition(c *JsonConditionContext)

	// ExitIsJsonCondition is called when exiting the isJsonCondition production.
	ExitIsJsonCondition(c *IsJsonConditionContext)

	// ExitJsonEqualCondition is called when exiting the jsonEqualCondition production.
	ExitJsonEqualCondition(c *JsonEqualConditionContext)

	// ExitJsonExistsCondition is called when exiting the jsonExistsCondition production.
	ExitJsonExistsCondition(c *JsonExistsConditionContext)

	// ExitJsonPassingClause is called when exiting the jsonPassingClause production.
	ExitJsonPassingClause(c *JsonPassingClauseContext)

	// ExitJsonExistsOnErrorClause is called when exiting the jsonExistsOnErrorClause production.
	ExitJsonExistsOnErrorClause(c *JsonExistsOnErrorClauseContext)

	// ExitJsonExistsOnEmptyClause is called when exiting the jsonExistsOnEmptyClause production.
	ExitJsonExistsOnEmptyClause(c *JsonExistsOnEmptyClauseContext)

	// ExitJsonTextcontainsCondition is called when exiting the jsonTextcontainsCondition production.
	ExitJsonTextcontainsCondition(c *JsonTextcontainsConditionContext)

	// ExitJsonBasicPathExpr is called when exiting the jsonBasicPathExpr production.
	ExitJsonBasicPathExpr(c *JsonBasicPathExprContext)

	// ExitJsonAbsolutePathExpr is called when exiting the jsonAbsolutePathExpr production.
	ExitJsonAbsolutePathExpr(c *JsonAbsolutePathExprContext)

	// ExitJsonNonfunctionSteps is called when exiting the jsonNonfunctionSteps production.
	ExitJsonNonfunctionSteps(c *JsonNonfunctionStepsContext)

	// ExitJsonObjectStep is called when exiting the jsonObjectStep production.
	ExitJsonObjectStep(c *JsonObjectStepContext)

	// ExitJsonFieldName is called when exiting the jsonFieldName production.
	ExitJsonFieldName(c *JsonFieldNameContext)

	// ExitLetter is called when exiting the letter production.
	ExitLetter(c *LetterContext)

	// ExitDigit is called when exiting the digit production.
	ExitDigit(c *DigitContext)

	// ExitJsonArrayStep is called when exiting the jsonArrayStep production.
	ExitJsonArrayStep(c *JsonArrayStepContext)

	// ExitJsonDescendentStep is called when exiting the jsonDescendentStep production.
	ExitJsonDescendentStep(c *JsonDescendentStepContext)

	// ExitJsonFunctionStep is called when exiting the jsonFunctionStep production.
	ExitJsonFunctionStep(c *JsonFunctionStepContext)

	// ExitJsonItemMethod is called when exiting the jsonItemMethod production.
	ExitJsonItemMethod(c *JsonItemMethodContext)

	// ExitJsonFilterExpr is called when exiting the jsonFilterExpr production.
	ExitJsonFilterExpr(c *JsonFilterExprContext)

	// ExitJsonCond is called when exiting the jsonCond production.
	ExitJsonCond(c *JsonCondContext)

	// ExitJsonDisjunction is called when exiting the jsonDisjunction production.
	ExitJsonDisjunction(c *JsonDisjunctionContext)

	// ExitJsonConjunction is called when exiting the jsonConjunction production.
	ExitJsonConjunction(c *JsonConjunctionContext)

	// ExitJsonNegation is called when exiting the jsonNegation production.
	ExitJsonNegation(c *JsonNegationContext)

	// ExitJsonExistsCond is called when exiting the jsonExistsCond production.
	ExitJsonExistsCond(c *JsonExistsCondContext)

	// ExitJsonHasSubstringCond is called when exiting the jsonHasSubstringCond production.
	ExitJsonHasSubstringCond(c *JsonHasSubstringCondContext)

	// ExitJsonStartsWithCond is called when exiting the jsonStartsWithCond production.
	ExitJsonStartsWithCond(c *JsonStartsWithCondContext)

	// ExitJsonLikeCond is called when exiting the jsonLikeCond production.
	ExitJsonLikeCond(c *JsonLikeCondContext)

	// ExitJsonLikeRegexCond is called when exiting the jsonLikeRegexCond production.
	ExitJsonLikeRegexCond(c *JsonLikeRegexCondContext)

	// ExitJsonEqRegexCond is called when exiting the jsonEqRegexCond production.
	ExitJsonEqRegexCond(c *JsonEqRegexCondContext)

	// ExitJsonInCond is called when exiting the jsonInCond production.
	ExitJsonInCond(c *JsonInCondContext)

	// ExitValueList is called when exiting the valueList production.
	ExitValueList(c *ValueListContext)

	// ExitJsonComparison is called when exiting the jsonComparison production.
	ExitJsonComparison(c *JsonComparisonContext)

	// ExitJsonRelativePathExpr is called when exiting the jsonRelativePathExpr production.
	ExitJsonRelativePathExpr(c *JsonRelativePathExprContext)

	// ExitJsonComparePred is called when exiting the jsonComparePred production.
	ExitJsonComparePred(c *JsonComparePredContext)

	// ExitJsonVar is called when exiting the jsonVar production.
	ExitJsonVar(c *JsonVarContext)

	// ExitJsonScalar is called when exiting the jsonScalar production.
	ExitJsonScalar(c *JsonScalarContext)

	// ExitJsonNumber is called when exiting the jsonNumber production.
	ExitJsonNumber(c *JsonNumberContext)

	// ExitJsonString is called when exiting the jsonString production.
	ExitJsonString(c *JsonStringContext)

	// ExitCompoundCondition is called when exiting the compoundCondition production.
	ExitCompoundCondition(c *CompoundConditionContext)

	// ExitExistsCondition is called when exiting the existsCondition production.
	ExitExistsCondition(c *ExistsConditionContext)

	// ExitInCondition is called when exiting the inCondition production.
	ExitInCondition(c *InConditionContext)

	// ExitIsOfTypeCondition is called when exiting the isOfTypeCondition production.
	ExitIsOfTypeCondition(c *IsOfTypeConditionContext)

	// ExitDatabaseCharset is called when exiting the databaseCharset production.
	ExitDatabaseCharset(c *DatabaseCharsetContext)

	// ExitNationalCharset is called when exiting the nationalCharset production.
	ExitNationalCharset(c *NationalCharsetContext)

	// ExitFilenamePattern is called when exiting the filenamePattern production.
	ExitFilenamePattern(c *FilenamePatternContext)

	// ExitConnectString is called when exiting the connectString production.
	ExitConnectString(c *ConnectStringContext)

	// ExitCreateTable is called when exiting the createTable production.
	ExitCreateTable(c *CreateTableContext)

	// ExitCreateIndex is called when exiting the createIndex production.
	ExitCreateIndex(c *CreateIndexContext)

	// ExitAlterTable is called when exiting the alterTable production.
	ExitAlterTable(c *AlterTableContext)

	// ExitAlterIndex is called when exiting the alterIndex production.
	ExitAlterIndex(c *AlterIndexContext)

	// ExitDropTable is called when exiting the dropTable production.
	ExitDropTable(c *DropTableContext)

	// ExitDropIndex is called when exiting the dropIndex production.
	ExitDropIndex(c *DropIndexContext)

	// ExitTruncateTable is called when exiting the truncateTable production.
	ExitTruncateTable(c *TruncateTableContext)

	// ExitCreateTableSpecification is called when exiting the createTableSpecification production.
	ExitCreateTableSpecification(c *CreateTableSpecificationContext)

	// ExitTablespaceClauseWithParen is called when exiting the tablespaceClauseWithParen production.
	ExitTablespaceClauseWithParen(c *TablespaceClauseWithParenContext)

	// ExitTablespaceClause is called when exiting the tablespaceClause production.
	ExitTablespaceClause(c *TablespaceClauseContext)

	// ExitCreateSharingClause is called when exiting the createSharingClause production.
	ExitCreateSharingClause(c *CreateSharingClauseContext)

	// ExitCreateDefinitionClause is called when exiting the createDefinitionClause production.
	ExitCreateDefinitionClause(c *CreateDefinitionClauseContext)

	// ExitCreateXMLTypeTableClause is called when exiting the createXMLTypeTableClause production.
	ExitCreateXMLTypeTableClause(c *CreateXMLTypeTableClauseContext)

	// ExitXmlTypeStorageClause is called when exiting the xmlTypeStorageClause production.
	ExitXmlTypeStorageClause(c *XmlTypeStorageClauseContext)

	// ExitXmlSchemaSpecClause is called when exiting the xmlSchemaSpecClause production.
	ExitXmlSchemaSpecClause(c *XmlSchemaSpecClauseContext)

	// ExitXmlTypeVirtualColumnsClause is called when exiting the xmlTypeVirtualColumnsClause production.
	ExitXmlTypeVirtualColumnsClause(c *XmlTypeVirtualColumnsClauseContext)

	// ExitOidClause is called when exiting the oidClause production.
	ExitOidClause(c *OidClauseContext)

	// ExitOidIndexClause is called when exiting the oidIndexClause production.
	ExitOidIndexClause(c *OidIndexClauseContext)

	// ExitCreateRelationalTableClause is called when exiting the createRelationalTableClause production.
	ExitCreateRelationalTableClause(c *CreateRelationalTableClauseContext)

	// ExitCreateMemOptimizeClause is called when exiting the createMemOptimizeClause production.
	ExitCreateMemOptimizeClause(c *CreateMemOptimizeClauseContext)

	// ExitCreateParentClause is called when exiting the createParentClause production.
	ExitCreateParentClause(c *CreateParentClauseContext)

	// ExitCreateObjectTableClause is called when exiting the createObjectTableClause production.
	ExitCreateObjectTableClause(c *CreateObjectTableClauseContext)

	// ExitRelationalProperties is called when exiting the relationalProperties production.
	ExitRelationalProperties(c *RelationalPropertiesContext)

	// ExitRelationalProperty is called when exiting the relationalProperty production.
	ExitRelationalProperty(c *RelationalPropertyContext)

	// ExitColumnDefinition is called when exiting the columnDefinition production.
	ExitColumnDefinition(c *ColumnDefinitionContext)

	// ExitVisibleClause is called when exiting the visibleClause production.
	ExitVisibleClause(c *VisibleClauseContext)

	// ExitDefaultNullClause is called when exiting the defaultNullClause production.
	ExitDefaultNullClause(c *DefaultNullClauseContext)

	// ExitIdentityClause is called when exiting the identityClause production.
	ExitIdentityClause(c *IdentityClauseContext)

	// ExitIdentifyOptions is called when exiting the identifyOptions production.
	ExitIdentifyOptions(c *IdentifyOptionsContext)

	// ExitIdentityOption is called when exiting the identityOption production.
	ExitIdentityOption(c *IdentityOptionContext)

	// ExitEncryptionSpecification is called when exiting the encryptionSpecification production.
	ExitEncryptionSpecification(c *EncryptionSpecificationContext)

	// ExitInlineConstraint is called when exiting the inlineConstraint production.
	ExitInlineConstraint(c *InlineConstraintContext)

	// ExitReferencesClause is called when exiting the referencesClause production.
	ExitReferencesClause(c *ReferencesClauseContext)

	// ExitConstraintState is called when exiting the constraintState production.
	ExitConstraintState(c *ConstraintStateContext)

	// ExitNotDeferrable is called when exiting the notDeferrable production.
	ExitNotDeferrable(c *NotDeferrableContext)

	// ExitInitiallyClause is called when exiting the initiallyClause production.
	ExitInitiallyClause(c *InitiallyClauseContext)

	// ExitExceptionsClause is called when exiting the exceptionsClause production.
	ExitExceptionsClause(c *ExceptionsClauseContext)

	// ExitUsingIndexClause is called when exiting the usingIndexClause production.
	ExitUsingIndexClause(c *UsingIndexClauseContext)

	// ExitCreateIndexClause is called when exiting the createIndexClause production.
	ExitCreateIndexClause(c *CreateIndexClauseContext)

	// ExitInlineRefConstraint is called when exiting the inlineRefConstraint production.
	ExitInlineRefConstraint(c *InlineRefConstraintContext)

	// ExitVirtualColumnDefinition is called when exiting the virtualColumnDefinition production.
	ExitVirtualColumnDefinition(c *VirtualColumnDefinitionContext)

	// ExitOutOfLineConstraint is called when exiting the outOfLineConstraint production.
	ExitOutOfLineConstraint(c *OutOfLineConstraintContext)

	// ExitOutOfLineRefConstraint is called when exiting the outOfLineRefConstraint production.
	ExitOutOfLineRefConstraint(c *OutOfLineRefConstraintContext)

	// ExitCreateIndexSpecification is called when exiting the createIndexSpecification production.
	ExitCreateIndexSpecification(c *CreateIndexSpecificationContext)

	// ExitClusterIndexClause is called when exiting the clusterIndexClause production.
	ExitClusterIndexClause(c *ClusterIndexClauseContext)

	// ExitIndexAttributes is called when exiting the indexAttributes production.
	ExitIndexAttributes(c *IndexAttributesContext)

	// ExitTableIndexClause is called when exiting the tableIndexClause production.
	ExitTableIndexClause(c *TableIndexClauseContext)

	// ExitIndexExpressions is called when exiting the indexExpressions production.
	ExitIndexExpressions(c *IndexExpressionsContext)

	// ExitIndexExpression is called when exiting the indexExpression production.
	ExitIndexExpression(c *IndexExpressionContext)

	// ExitBitmapJoinIndexClause is called when exiting the bitmapJoinIndexClause production.
	ExitBitmapJoinIndexClause(c *BitmapJoinIndexClauseContext)

	// ExitColumnSortsClause_ is called when exiting the columnSortsClause_ production.
	ExitColumnSortsClause_(c *ColumnSortsClause_Context)

	// ExitColumnSortClause_ is called when exiting the columnSortClause_ production.
	ExitColumnSortClause_(c *ColumnSortClause_Context)

	// ExitCreateIndexDefinitionClause is called when exiting the createIndexDefinitionClause production.
	ExitCreateIndexDefinitionClause(c *CreateIndexDefinitionClauseContext)

	// ExitTableAlias is called when exiting the tableAlias production.
	ExitTableAlias(c *TableAliasContext)

	// ExitAlterDefinitionClause is called when exiting the alterDefinitionClause production.
	ExitAlterDefinitionClause(c *AlterDefinitionClauseContext)

	// ExitAlterTableProperties is called when exiting the alterTableProperties production.
	ExitAlterTableProperties(c *AlterTablePropertiesContext)

	// ExitRenameTableSpecification is called when exiting the renameTableSpecification production.
	ExitRenameTableSpecification(c *RenameTableSpecificationContext)

	// ExitColumnClauses is called when exiting the columnClauses production.
	ExitColumnClauses(c *ColumnClausesContext)

	// ExitOperateColumnClause is called when exiting the operateColumnClause production.
	ExitOperateColumnClause(c *OperateColumnClauseContext)

	// ExitAddColumnSpecification is called when exiting the addColumnSpecification production.
	ExitAddColumnSpecification(c *AddColumnSpecificationContext)

	// ExitColumnOrVirtualDefinitions is called when exiting the columnOrVirtualDefinitions production.
	ExitColumnOrVirtualDefinitions(c *ColumnOrVirtualDefinitionsContext)

	// ExitColumnOrVirtualDefinition is called when exiting the columnOrVirtualDefinition production.
	ExitColumnOrVirtualDefinition(c *ColumnOrVirtualDefinitionContext)

	// ExitColumnProperties is called when exiting the columnProperties production.
	ExitColumnProperties(c *ColumnPropertiesContext)

	// ExitColumnProperty is called when exiting the columnProperty production.
	ExitColumnProperty(c *ColumnPropertyContext)

	// ExitObjectTypeColProperties is called when exiting the objectTypeColProperties production.
	ExitObjectTypeColProperties(c *ObjectTypeColPropertiesContext)

	// ExitSubstitutableColumnClause is called when exiting the substitutableColumnClause production.
	ExitSubstitutableColumnClause(c *SubstitutableColumnClauseContext)

	// ExitModifyColumnSpecification is called when exiting the modifyColumnSpecification production.
	ExitModifyColumnSpecification(c *ModifyColumnSpecificationContext)

	// ExitModifyColProperties is called when exiting the modifyColProperties production.
	ExitModifyColProperties(c *ModifyColPropertiesContext)

	// ExitModifyColSubstitutable is called when exiting the modifyColSubstitutable production.
	ExitModifyColSubstitutable(c *ModifyColSubstitutableContext)

	// ExitDropColumnClause is called when exiting the dropColumnClause production.
	ExitDropColumnClause(c *DropColumnClauseContext)

	// ExitDropColumnSpecification is called when exiting the dropColumnSpecification production.
	ExitDropColumnSpecification(c *DropColumnSpecificationContext)

	// ExitColumnOrColumnList is called when exiting the columnOrColumnList production.
	ExitColumnOrColumnList(c *ColumnOrColumnListContext)

	// ExitCascadeOrInvalidate is called when exiting the cascadeOrInvalidate production.
	ExitCascadeOrInvalidate(c *CascadeOrInvalidateContext)

	// ExitCheckpointNumber is called when exiting the checkpointNumber production.
	ExitCheckpointNumber(c *CheckpointNumberContext)

	// ExitRenameColumnClause is called when exiting the renameColumnClause production.
	ExitRenameColumnClause(c *RenameColumnClauseContext)

	// ExitConstraintClauses is called when exiting the constraintClauses production.
	ExitConstraintClauses(c *ConstraintClausesContext)

	// ExitAddConstraintSpecification is called when exiting the addConstraintSpecification production.
	ExitAddConstraintSpecification(c *AddConstraintSpecificationContext)

	// ExitModifyConstraintClause is called when exiting the modifyConstraintClause production.
	ExitModifyConstraintClause(c *ModifyConstraintClauseContext)

	// ExitConstraintWithName is called when exiting the constraintWithName production.
	ExitConstraintWithName(c *ConstraintWithNameContext)

	// ExitConstraintOption is called when exiting the constraintOption production.
	ExitConstraintOption(c *ConstraintOptionContext)

	// ExitConstraintPrimaryOrUnique is called when exiting the constraintPrimaryOrUnique production.
	ExitConstraintPrimaryOrUnique(c *ConstraintPrimaryOrUniqueContext)

	// ExitRenameConstraintClause is called when exiting the renameConstraintClause production.
	ExitRenameConstraintClause(c *RenameConstraintClauseContext)

	// ExitDropConstraintClause is called when exiting the dropConstraintClause production.
	ExitDropConstraintClause(c *DropConstraintClauseContext)

	// ExitAlterExternalTable is called when exiting the alterExternalTable production.
	ExitAlterExternalTable(c *AlterExternalTableContext)

	// ExitObjectProperties is called when exiting the objectProperties production.
	ExitObjectProperties(c *ObjectPropertiesContext)

	// ExitAlterIndexInformationClause is called when exiting the alterIndexInformationClause production.
	ExitAlterIndexInformationClause(c *AlterIndexInformationClauseContext)

	// ExitRenameIndexClause is called when exiting the renameIndexClause production.
	ExitRenameIndexClause(c *RenameIndexClauseContext)

	// ExitObjectTableSubstitution is called when exiting the objectTableSubstitution production.
	ExitObjectTableSubstitution(c *ObjectTableSubstitutionContext)

	// ExitMemOptimizeClause is called when exiting the memOptimizeClause production.
	ExitMemOptimizeClause(c *MemOptimizeClauseContext)

	// ExitMemOptimizeReadClause is called when exiting the memOptimizeReadClause production.
	ExitMemOptimizeReadClause(c *MemOptimizeReadClauseContext)

	// ExitMemOptimizeWriteClause is called when exiting the memOptimizeWriteClause production.
	ExitMemOptimizeWriteClause(c *MemOptimizeWriteClauseContext)

	// ExitEnableDisableClauses is called when exiting the enableDisableClauses production.
	ExitEnableDisableClauses(c *EnableDisableClausesContext)

	// ExitEnableDisableClause is called when exiting the enableDisableClause production.
	ExitEnableDisableClause(c *EnableDisableClauseContext)

	// ExitEnableDisableOthers is called when exiting the enableDisableOthers production.
	ExitEnableDisableOthers(c *EnableDisableOthersContext)

	// ExitRebuildClause is called when exiting the rebuildClause production.
	ExitRebuildClause(c *RebuildClauseContext)

	// ExitParallelClause is called when exiting the parallelClause production.
	ExitParallelClause(c *ParallelClauseContext)

	// ExitUsableSpecification is called when exiting the usableSpecification production.
	ExitUsableSpecification(c *UsableSpecificationContext)

	// ExitInvalidationSpecification is called when exiting the invalidationSpecification production.
	ExitInvalidationSpecification(c *InvalidationSpecificationContext)

	// ExitMaterializedViewLogClause is called when exiting the materializedViewLogClause production.
	ExitMaterializedViewLogClause(c *MaterializedViewLogClauseContext)

	// ExitDropReuseClause is called when exiting the dropReuseClause production.
	ExitDropReuseClause(c *DropReuseClauseContext)

	// ExitCollationClause is called when exiting the collationClause production.
	ExitCollationClause(c *CollationClauseContext)

	// ExitCommitClause is called when exiting the commitClause production.
	ExitCommitClause(c *CommitClauseContext)

	// ExitPhysicalProperties is called when exiting the physicalProperties production.
	ExitPhysicalProperties(c *PhysicalPropertiesContext)

	// ExitDeferredSegmentCreation is called when exiting the deferredSegmentCreation production.
	ExitDeferredSegmentCreation(c *DeferredSegmentCreationContext)

	// ExitSegmentAttributesClause is called when exiting the segmentAttributesClause production.
	ExitSegmentAttributesClause(c *SegmentAttributesClauseContext)

	// ExitPhysicalAttributesClause is called when exiting the physicalAttributesClause production.
	ExitPhysicalAttributesClause(c *PhysicalAttributesClauseContext)

	// ExitLoggingClause is called when exiting the loggingClause production.
	ExitLoggingClause(c *LoggingClauseContext)

	// ExitStorageClause is called when exiting the storageClause production.
	ExitStorageClause(c *StorageClauseContext)

	// ExitSizeClause is called when exiting the sizeClause production.
	ExitSizeClause(c *SizeClauseContext)

	// ExitMaxsizeClause is called when exiting the maxsizeClause production.
	ExitMaxsizeClause(c *MaxsizeClauseContext)

	// ExitTableCompression is called when exiting the tableCompression production.
	ExitTableCompression(c *TableCompressionContext)

	// ExitInmemoryTableClause is called when exiting the inmemoryTableClause production.
	ExitInmemoryTableClause(c *InmemoryTableClauseContext)

	// ExitInmemoryAttributes is called when exiting the inmemoryAttributes production.
	ExitInmemoryAttributes(c *InmemoryAttributesContext)

	// ExitInmemoryColumnClause is called when exiting the inmemoryColumnClause production.
	ExitInmemoryColumnClause(c *InmemoryColumnClauseContext)

	// ExitInmemoryMemcompress is called when exiting the inmemoryMemcompress production.
	ExitInmemoryMemcompress(c *InmemoryMemcompressContext)

	// ExitInmemoryPriority is called when exiting the inmemoryPriority production.
	ExitInmemoryPriority(c *InmemoryPriorityContext)

	// ExitInmemoryDistribute is called when exiting the inmemoryDistribute production.
	ExitInmemoryDistribute(c *InmemoryDistributeContext)

	// ExitInmemoryDuplicate is called when exiting the inmemoryDuplicate production.
	ExitInmemoryDuplicate(c *InmemoryDuplicateContext)

	// ExitIlmClause is called when exiting the ilmClause production.
	ExitIlmClause(c *IlmClauseContext)

	// ExitIlmPolicyClause is called when exiting the ilmPolicyClause production.
	ExitIlmPolicyClause(c *IlmPolicyClauseContext)

	// ExitIlmCompressionPolicy is called when exiting the ilmCompressionPolicy production.
	ExitIlmCompressionPolicy(c *IlmCompressionPolicyContext)

	// ExitIlmTimePeriod is called when exiting the ilmTimePeriod production.
	ExitIlmTimePeriod(c *IlmTimePeriodContext)

	// ExitIlmTieringPolicy is called when exiting the ilmTieringPolicy production.
	ExitIlmTieringPolicy(c *IlmTieringPolicyContext)

	// ExitIlmInmemoryPolicy is called when exiting the ilmInmemoryPolicy production.
	ExitIlmInmemoryPolicy(c *IlmInmemoryPolicyContext)

	// ExitOrganizationClause is called when exiting the organizationClause production.
	ExitOrganizationClause(c *OrganizationClauseContext)

	// ExitHeapOrgTableClause is called when exiting the heapOrgTableClause production.
	ExitHeapOrgTableClause(c *HeapOrgTableClauseContext)

	// ExitIndexOrgTableClause is called when exiting the indexOrgTableClause production.
	ExitIndexOrgTableClause(c *IndexOrgTableClauseContext)

	// ExitExternalTableClause is called when exiting the externalTableClause production.
	ExitExternalTableClause(c *ExternalTableClauseContext)

	// ExitExternalTableDataProps is called when exiting the externalTableDataProps production.
	ExitExternalTableDataProps(c *ExternalTableDataPropsContext)

	// ExitMappingTableClause is called when exiting the mappingTableClause production.
	ExitMappingTableClause(c *MappingTableClauseContext)

	// ExitPrefixCompression is called when exiting the prefixCompression production.
	ExitPrefixCompression(c *PrefixCompressionContext)

	// ExitIndexOrgOverflowClause is called when exiting the indexOrgOverflowClause production.
	ExitIndexOrgOverflowClause(c *IndexOrgOverflowClauseContext)

	// ExitExternalPartitionClause is called when exiting the externalPartitionClause production.
	ExitExternalPartitionClause(c *ExternalPartitionClauseContext)

	// ExitClusterRelatedClause is called when exiting the clusterRelatedClause production.
	ExitClusterRelatedClause(c *ClusterRelatedClauseContext)

	// ExitTableProperties is called when exiting the tableProperties production.
	ExitTableProperties(c *TablePropertiesContext)

	// ExitReadOnlyClause is called when exiting the readOnlyClause production.
	ExitReadOnlyClause(c *ReadOnlyClauseContext)

	// ExitIndexingClause is called when exiting the indexingClause production.
	ExitIndexingClause(c *IndexingClauseContext)

	// ExitTablePartitioningClauses is called when exiting the tablePartitioningClauses production.
	ExitTablePartitioningClauses(c *TablePartitioningClausesContext)

	// ExitRangePartitions is called when exiting the rangePartitions production.
	ExitRangePartitions(c *RangePartitionsContext)

	// ExitRangeValuesClause is called when exiting the rangeValuesClause production.
	ExitRangeValuesClause(c *RangeValuesClauseContext)

	// ExitTablePartitionDescription is called when exiting the tablePartitionDescription production.
	ExitTablePartitionDescription(c *TablePartitionDescriptionContext)

	// ExitInmemoryClause is called when exiting the inmemoryClause production.
	ExitInmemoryClause(c *InmemoryClauseContext)

	// ExitVarrayColProperties is called when exiting the varrayColProperties production.
	ExitVarrayColProperties(c *VarrayColPropertiesContext)

	// ExitNestedTableColProperties is called when exiting the nestedTableColProperties production.
	ExitNestedTableColProperties(c *NestedTableColPropertiesContext)

	// ExitLobStorageClause is called when exiting the lobStorageClause production.
	ExitLobStorageClause(c *LobStorageClauseContext)

	// ExitVarrayStorageClause is called when exiting the varrayStorageClause production.
	ExitVarrayStorageClause(c *VarrayStorageClauseContext)

	// ExitLobStorageParameters is called when exiting the lobStorageParameters production.
	ExitLobStorageParameters(c *LobStorageParametersContext)

	// ExitLobParameters is called when exiting the lobParameters production.
	ExitLobParameters(c *LobParametersContext)

	// ExitLobRetentionClause is called when exiting the lobRetentionClause production.
	ExitLobRetentionClause(c *LobRetentionClauseContext)

	// ExitLobDeduplicateClause is called when exiting the lobDeduplicateClause production.
	ExitLobDeduplicateClause(c *LobDeduplicateClauseContext)

	// ExitLobCompressionClause is called when exiting the lobCompressionClause production.
	ExitLobCompressionClause(c *LobCompressionClauseContext)

	// ExitExternalPartSubpartDataProps is called when exiting the externalPartSubpartDataProps production.
	ExitExternalPartSubpartDataProps(c *ExternalPartSubpartDataPropsContext)

	// ExitListPartitions is called when exiting the listPartitions production.
	ExitListPartitions(c *ListPartitionsContext)

	// ExitListValuesClause is called when exiting the listValuesClause production.
	ExitListValuesClause(c *ListValuesClauseContext)

	// ExitListValues is called when exiting the listValues production.
	ExitListValues(c *ListValuesContext)

	// ExitHashPartitions is called when exiting the hashPartitions production.
	ExitHashPartitions(c *HashPartitionsContext)

	// ExitHashPartitionsByQuantity is called when exiting the hashPartitionsByQuantity production.
	ExitHashPartitionsByQuantity(c *HashPartitionsByQuantityContext)

	// ExitIndexCompression is called when exiting the indexCompression production.
	ExitIndexCompression(c *IndexCompressionContext)

	// ExitAdvancedIndexCompression is called when exiting the advancedIndexCompression production.
	ExitAdvancedIndexCompression(c *AdvancedIndexCompressionContext)

	// ExitIndividualHashPartitions is called when exiting the individualHashPartitions production.
	ExitIndividualHashPartitions(c *IndividualHashPartitionsContext)

	// ExitPartitioningStorageClause is called when exiting the partitioningStorageClause production.
	ExitPartitioningStorageClause(c *PartitioningStorageClauseContext)

	// ExitLobPartitioningStorage is called when exiting the lobPartitioningStorage production.
	ExitLobPartitioningStorage(c *LobPartitioningStorageContext)

	// ExitCompositeRangePartitions is called when exiting the compositeRangePartitions production.
	ExitCompositeRangePartitions(c *CompositeRangePartitionsContext)

	// ExitSubpartitionByRange is called when exiting the subpartitionByRange production.
	ExitSubpartitionByRange(c *SubpartitionByRangeContext)

	// ExitSubpartitionByList is called when exiting the subpartitionByList production.
	ExitSubpartitionByList(c *SubpartitionByListContext)

	// ExitSubpartitionByHash is called when exiting the subpartitionByHash production.
	ExitSubpartitionByHash(c *SubpartitionByHashContext)

	// ExitSubpartitionTemplate is called when exiting the subpartitionTemplate production.
	ExitSubpartitionTemplate(c *SubpartitionTemplateContext)

	// ExitRangeSubpartitionDesc is called when exiting the rangeSubpartitionDesc production.
	ExitRangeSubpartitionDesc(c *RangeSubpartitionDescContext)

	// ExitListSubpartitionDesc is called when exiting the listSubpartitionDesc production.
	ExitListSubpartitionDesc(c *ListSubpartitionDescContext)

	// ExitIndividualHashSubparts is called when exiting the individualHashSubparts production.
	ExitIndividualHashSubparts(c *IndividualHashSubpartsContext)

	// ExitRangePartitionDesc is called when exiting the rangePartitionDesc production.
	ExitRangePartitionDesc(c *RangePartitionDescContext)

	// ExitCompositeListPartitions is called when exiting the compositeListPartitions production.
	ExitCompositeListPartitions(c *CompositeListPartitionsContext)

	// ExitListPartitionDesc is called when exiting the listPartitionDesc production.
	ExitListPartitionDesc(c *ListPartitionDescContext)

	// ExitCompositeHashPartitions is called when exiting the compositeHashPartitions production.
	ExitCompositeHashPartitions(c *CompositeHashPartitionsContext)

	// ExitReferencePartitioning is called when exiting the referencePartitioning production.
	ExitReferencePartitioning(c *ReferencePartitioningContext)

	// ExitReferencePartitionDesc is called when exiting the referencePartitionDesc production.
	ExitReferencePartitionDesc(c *ReferencePartitionDescContext)

	// ExitConstraint is called when exiting the constraint production.
	ExitConstraint(c *ConstraintContext)

	// ExitSystemPartitioning is called when exiting the systemPartitioning production.
	ExitSystemPartitioning(c *SystemPartitioningContext)

	// ExitConsistentHashPartitions is called when exiting the consistentHashPartitions production.
	ExitConsistentHashPartitions(c *ConsistentHashPartitionsContext)

	// ExitConsistentHashWithSubpartitions is called when exiting the consistentHashWithSubpartitions production.
	ExitConsistentHashWithSubpartitions(c *ConsistentHashWithSubpartitionsContext)

	// ExitPartitionsetClauses is called when exiting the partitionsetClauses production.
	ExitPartitionsetClauses(c *PartitionsetClausesContext)

	// ExitRangePartitionsetClause is called when exiting the rangePartitionsetClause production.
	ExitRangePartitionsetClause(c *RangePartitionsetClauseContext)

	// ExitRangePartitionsetDesc is called when exiting the rangePartitionsetDesc production.
	ExitRangePartitionsetDesc(c *RangePartitionsetDescContext)

	// ExitListPartitionsetClause is called when exiting the listPartitionsetClause production.
	ExitListPartitionsetClause(c *ListPartitionsetClauseContext)

	// ExitAttributeClusteringClause is called when exiting the attributeClusteringClause production.
	ExitAttributeClusteringClause(c *AttributeClusteringClauseContext)

	// ExitClusteringJoin is called when exiting the clusteringJoin production.
	ExitClusteringJoin(c *ClusteringJoinContext)

	// ExitClusterClause is called when exiting the clusterClause production.
	ExitClusterClause(c *ClusterClauseContext)

	// ExitClusteringColumns is called when exiting the clusteringColumns production.
	ExitClusteringColumns(c *ClusteringColumnsContext)

	// ExitClusteringColumnGroup is called when exiting the clusteringColumnGroup production.
	ExitClusteringColumnGroup(c *ClusteringColumnGroupContext)

	// ExitClusteringWhen is called when exiting the clusteringWhen production.
	ExitClusteringWhen(c *ClusteringWhenContext)

	// ExitZonemapClause is called when exiting the zonemapClause production.
	ExitZonemapClause(c *ZonemapClauseContext)

	// ExitRowMovementClause is called when exiting the rowMovementClause production.
	ExitRowMovementClause(c *RowMovementClauseContext)

	// ExitFlashbackArchiveClause is called when exiting the flashbackArchiveClause production.
	ExitFlashbackArchiveClause(c *FlashbackArchiveClauseContext)

	// ExitAlterSynonym is called when exiting the alterSynonym production.
	ExitAlterSynonym(c *AlterSynonymContext)

	// ExitAlterTablePartitioning is called when exiting the alterTablePartitioning production.
	ExitAlterTablePartitioning(c *AlterTablePartitioningContext)

	// ExitModifyTablePartition is called when exiting the modifyTablePartition production.
	ExitModifyTablePartition(c *ModifyTablePartitionContext)

	// ExitModifyRangePartition is called when exiting the modifyRangePartition production.
	ExitModifyRangePartition(c *ModifyRangePartitionContext)

	// ExitModifyHashPartition is called when exiting the modifyHashPartition production.
	ExitModifyHashPartition(c *ModifyHashPartitionContext)

	// ExitModifyListPartition is called when exiting the modifyListPartition production.
	ExitModifyListPartition(c *ModifyListPartitionContext)

	// ExitPartitionExtendedName is called when exiting the partitionExtendedName production.
	ExitPartitionExtendedName(c *PartitionExtendedNameContext)

	// ExitAddRangeSubpartition is called when exiting the addRangeSubpartition production.
	ExitAddRangeSubpartition(c *AddRangeSubpartitionContext)

	// ExitDependentTablesClause is called when exiting the dependentTablesClause production.
	ExitDependentTablesClause(c *DependentTablesClauseContext)

	// ExitAddHashSubpartition is called when exiting the addHashSubpartition production.
	ExitAddHashSubpartition(c *AddHashSubpartitionContext)

	// ExitAddListSubpartition is called when exiting the addListSubpartition production.
	ExitAddListSubpartition(c *AddListSubpartitionContext)

	// ExitCoalesceTableSubpartition is called when exiting the coalesceTableSubpartition production.
	ExitCoalesceTableSubpartition(c *CoalesceTableSubpartitionContext)

	// ExitAllowDisallowClustering is called when exiting the allowDisallowClustering production.
	ExitAllowDisallowClustering(c *AllowDisallowClusteringContext)

	// ExitAlterMappingTableClauses is called when exiting the alterMappingTableClauses production.
	ExitAlterMappingTableClauses(c *AlterMappingTableClausesContext)

	// ExitDeallocateUnusedClause is called when exiting the deallocateUnusedClause production.
	ExitDeallocateUnusedClause(c *DeallocateUnusedClauseContext)

	// ExitAllocateExtentClause is called when exiting the allocateExtentClause production.
	ExitAllocateExtentClause(c *AllocateExtentClauseContext)

	// ExitPartitionSpec is called when exiting the partitionSpec production.
	ExitPartitionSpec(c *PartitionSpecContext)

	// ExitPartitionAttributes is called when exiting the partitionAttributes production.
	ExitPartitionAttributes(c *PartitionAttributesContext)

	// ExitShrinkClause is called when exiting the shrinkClause production.
	ExitShrinkClause(c *ShrinkClauseContext)

	// ExitMoveTablePartition is called when exiting the moveTablePartition production.
	ExitMoveTablePartition(c *MoveTablePartitionContext)

	// ExitFilterCondition is called when exiting the filterCondition production.
	ExitFilterCondition(c *FilterConditionContext)

	// ExitCoalesceTablePartition is called when exiting the coalesceTablePartition production.
	ExitCoalesceTablePartition(c *CoalesceTablePartitionContext)

	// ExitAddTablePartition is called when exiting the addTablePartition production.
	ExitAddTablePartition(c *AddTablePartitionContext)

	// ExitAddRangePartitionClause is called when exiting the addRangePartitionClause production.
	ExitAddRangePartitionClause(c *AddRangePartitionClauseContext)

	// ExitAddListPartitionClause is called when exiting the addListPartitionClause production.
	ExitAddListPartitionClause(c *AddListPartitionClauseContext)

	// ExitHashSubpartsByQuantity is called when exiting the hashSubpartsByQuantity production.
	ExitHashSubpartsByQuantity(c *HashSubpartsByQuantityContext)

	// ExitAddSystemPartitionClause is called when exiting the addSystemPartitionClause production.
	ExitAddSystemPartitionClause(c *AddSystemPartitionClauseContext)

	// ExitAddHashPartitionClause is called when exiting the addHashPartitionClause production.
	ExitAddHashPartitionClause(c *AddHashPartitionClauseContext)

	// ExitDropTablePartition is called when exiting the dropTablePartition production.
	ExitDropTablePartition(c *DropTablePartitionContext)

	// ExitPartitionExtendedNames is called when exiting the partitionExtendedNames production.
	ExitPartitionExtendedNames(c *PartitionExtendedNamesContext)

	// ExitPartitionForClauses is called when exiting the partitionForClauses production.
	ExitPartitionForClauses(c *PartitionForClausesContext)

	// ExitUpdateIndexClauses is called when exiting the updateIndexClauses production.
	ExitUpdateIndexClauses(c *UpdateIndexClausesContext)

	// ExitUpdateGlobalIndexClause is called when exiting the updateGlobalIndexClause production.
	ExitUpdateGlobalIndexClause(c *UpdateGlobalIndexClauseContext)

	// ExitUpdateAllIndexesClause is called when exiting the updateAllIndexesClause production.
	ExitUpdateAllIndexesClause(c *UpdateAllIndexesClauseContext)

	// ExitUpdateIndexPartition is called when exiting the updateIndexPartition production.
	ExitUpdateIndexPartition(c *UpdateIndexPartitionContext)

	// ExitIndexPartitionDesc is called when exiting the indexPartitionDesc production.
	ExitIndexPartitionDesc(c *IndexPartitionDescContext)

	// ExitIndexSubpartitionClause is called when exiting the indexSubpartitionClause production.
	ExitIndexSubpartitionClause(c *IndexSubpartitionClauseContext)

	// ExitUpdateIndexSubpartition is called when exiting the updateIndexSubpartition production.
	ExitUpdateIndexSubpartition(c *UpdateIndexSubpartitionContext)

	// ExitSupplementalLoggingProps is called when exiting the supplementalLoggingProps production.
	ExitSupplementalLoggingProps(c *SupplementalLoggingPropsContext)

	// ExitSupplementalLogGrpClause is called when exiting the supplementalLogGrpClause production.
	ExitSupplementalLogGrpClause(c *SupplementalLogGrpClauseContext)

	// ExitSupplementalIdKeyClause is called when exiting the supplementalIdKeyClause production.
	ExitSupplementalIdKeyClause(c *SupplementalIdKeyClauseContext)

	// ExitAlterSession is called when exiting the alterSession production.
	ExitAlterSession(c *AlterSessionContext)

	// ExitAlterSessionOption is called when exiting the alterSessionOption production.
	ExitAlterSessionOption(c *AlterSessionOptionContext)

	// ExitAdviseClause is called when exiting the adviseClause production.
	ExitAdviseClause(c *AdviseClauseContext)

	// ExitCloseDatabaseLinkClause is called when exiting the closeDatabaseLinkClause production.
	ExitCloseDatabaseLinkClause(c *CloseDatabaseLinkClauseContext)

	// ExitCommitInProcedureClause is called when exiting the commitInProcedureClause production.
	ExitCommitInProcedureClause(c *CommitInProcedureClauseContext)

	// ExitSecuriyClause is called when exiting the securiyClause production.
	ExitSecuriyClause(c *SecuriyClauseContext)

	// ExitParallelExecutionClause is called when exiting the parallelExecutionClause production.
	ExitParallelExecutionClause(c *ParallelExecutionClauseContext)

	// ExitResumableClause is called when exiting the resumableClause production.
	ExitResumableClause(c *ResumableClauseContext)

	// ExitEnableResumableClause is called when exiting the enableResumableClause production.
	ExitEnableResumableClause(c *EnableResumableClauseContext)

	// ExitDisableResumableClause is called when exiting the disableResumableClause production.
	ExitDisableResumableClause(c *DisableResumableClauseContext)

	// ExitShardDdlClause is called when exiting the shardDdlClause production.
	ExitShardDdlClause(c *ShardDdlClauseContext)

	// ExitSyncWithPrimaryClause is called when exiting the syncWithPrimaryClause production.
	ExitSyncWithPrimaryClause(c *SyncWithPrimaryClauseContext)

	// ExitAlterSessionSetClause is called when exiting the alterSessionSetClause production.
	ExitAlterSessionSetClause(c *AlterSessionSetClauseContext)

	// ExitAlterSessionSetClauseOption is called when exiting the alterSessionSetClauseOption production.
	ExitAlterSessionSetClauseOption(c *AlterSessionSetClauseOptionContext)

	// ExitParameterClause is called when exiting the parameterClause production.
	ExitParameterClause(c *ParameterClauseContext)

	// ExitEditionClause is called when exiting the editionClause production.
	ExitEditionClause(c *EditionClauseContext)

	// ExitContainerClause is called when exiting the containerClause production.
	ExitContainerClause(c *ContainerClauseContext)

	// ExitRowArchivalVisibilityClause is called when exiting the rowArchivalVisibilityClause production.
	ExitRowArchivalVisibilityClause(c *RowArchivalVisibilityClauseContext)

	// ExitAlterDatabase is called when exiting the alterDatabase production.
	ExitAlterDatabase(c *AlterDatabaseContext)

	// ExitDatabaseClauses is called when exiting the databaseClauses production.
	ExitDatabaseClauses(c *DatabaseClausesContext)

	// ExitStartupClauses is called when exiting the startupClauses production.
	ExitStartupClauses(c *StartupClausesContext)

	// ExitRecoveryClauses is called when exiting the recoveryClauses production.
	ExitRecoveryClauses(c *RecoveryClausesContext)

	// ExitGeneralRecovery is called when exiting the generalRecovery production.
	ExitGeneralRecovery(c *GeneralRecoveryContext)

	// ExitFullDatabaseRecovery is called when exiting the fullDatabaseRecovery production.
	ExitFullDatabaseRecovery(c *FullDatabaseRecoveryContext)

	// ExitPartialDatabaseRecovery is called when exiting the partialDatabaseRecovery production.
	ExitPartialDatabaseRecovery(c *PartialDatabaseRecoveryContext)

	// ExitManagedStandbyRecovery is called when exiting the managedStandbyRecovery production.
	ExitManagedStandbyRecovery(c *ManagedStandbyRecoveryContext)

	// ExitDatabaseFileClauses is called when exiting the databaseFileClauses production.
	ExitDatabaseFileClauses(c *DatabaseFileClausesContext)

	// ExitCreateDatafileClause is called when exiting the createDatafileClause production.
	ExitCreateDatafileClause(c *CreateDatafileClauseContext)

	// ExitFileSpecifications is called when exiting the fileSpecifications production.
	ExitFileSpecifications(c *FileSpecificationsContext)

	// ExitFileSpecification is called when exiting the fileSpecification production.
	ExitFileSpecification(c *FileSpecificationContext)

	// ExitDatafileTempfileSpec is called when exiting the datafileTempfileSpec production.
	ExitDatafileTempfileSpec(c *DatafileTempfileSpecContext)

	// ExitAutoextendClause is called when exiting the autoextendClause production.
	ExitAutoextendClause(c *AutoextendClauseContext)

	// ExitRedoLogFileSpec is called when exiting the redoLogFileSpec production.
	ExitRedoLogFileSpec(c *RedoLogFileSpecContext)

	// ExitAlterDatafileClause is called when exiting the alterDatafileClause production.
	ExitAlterDatafileClause(c *AlterDatafileClauseContext)

	// ExitAlterTempfileClause is called when exiting the alterTempfileClause production.
	ExitAlterTempfileClause(c *AlterTempfileClauseContext)

	// ExitLogfileClauses is called when exiting the logfileClauses production.
	ExitLogfileClauses(c *LogfileClausesContext)

	// ExitLogfileDescriptor is called when exiting the logfileDescriptor production.
	ExitLogfileDescriptor(c *LogfileDescriptorContext)

	// ExitAddLogfileClauses is called when exiting the addLogfileClauses production.
	ExitAddLogfileClauses(c *AddLogfileClausesContext)

	// ExitControlfileClauses is called when exiting the controlfileClauses production.
	ExitControlfileClauses(c *ControlfileClausesContext)

	// ExitTraceFileClause is called when exiting the traceFileClause production.
	ExitTraceFileClause(c *TraceFileClauseContext)

	// ExitDropLogfileClauses is called when exiting the dropLogfileClauses production.
	ExitDropLogfileClauses(c *DropLogfileClausesContext)

	// ExitSwitchLogfileClause is called when exiting the switchLogfileClause production.
	ExitSwitchLogfileClause(c *SwitchLogfileClauseContext)

	// ExitSupplementalDbLogging is called when exiting the supplementalDbLogging production.
	ExitSupplementalDbLogging(c *SupplementalDbLoggingContext)

	// ExitSupplementalPlsqlClause is called when exiting the supplementalPlsqlClause production.
	ExitSupplementalPlsqlClause(c *SupplementalPlsqlClauseContext)

	// ExitSupplementalSubsetReplicationClause is called when exiting the supplementalSubsetReplicationClause production.
	ExitSupplementalSubsetReplicationClause(c *SupplementalSubsetReplicationClauseContext)

	// ExitStandbyDatabaseClauses is called when exiting the standbyDatabaseClauses production.
	ExitStandbyDatabaseClauses(c *StandbyDatabaseClausesContext)

	// ExitActivateStandbyDbClause is called when exiting the activateStandbyDbClause production.
	ExitActivateStandbyDbClause(c *ActivateStandbyDbClauseContext)

	// ExitMaximizeStandbyDbClause is called when exiting the maximizeStandbyDbClause production.
	ExitMaximizeStandbyDbClause(c *MaximizeStandbyDbClauseContext)

	// ExitRegisterLogfileClause is called when exiting the registerLogfileClause production.
	ExitRegisterLogfileClause(c *RegisterLogfileClauseContext)

	// ExitCommitSwitchoverClause is called when exiting the commitSwitchoverClause production.
	ExitCommitSwitchoverClause(c *CommitSwitchoverClauseContext)

	// ExitStartStandbyClause is called when exiting the startStandbyClause production.
	ExitStartStandbyClause(c *StartStandbyClauseContext)

	// ExitStopStandbyClause is called when exiting the stopStandbyClause production.
	ExitStopStandbyClause(c *StopStandbyClauseContext)

	// ExitSwitchoverClause is called when exiting the switchoverClause production.
	ExitSwitchoverClause(c *SwitchoverClauseContext)

	// ExitConvertDatabaseClause is called when exiting the convertDatabaseClause production.
	ExitConvertDatabaseClause(c *ConvertDatabaseClauseContext)

	// ExitFailoverClause is called when exiting the failoverClause production.
	ExitFailoverClause(c *FailoverClauseContext)

	// ExitDefaultSettingsClauses is called when exiting the defaultSettingsClauses production.
	ExitDefaultSettingsClauses(c *DefaultSettingsClausesContext)

	// ExitSetTimeZoneClause is called when exiting the setTimeZoneClause production.
	ExitSetTimeZoneClause(c *SetTimeZoneClauseContext)

	// ExitTimeZoneRegion is called when exiting the timeZoneRegion production.
	ExitTimeZoneRegion(c *TimeZoneRegionContext)

	// ExitFlashbackModeClause is called when exiting the flashbackModeClause production.
	ExitFlashbackModeClause(c *FlashbackModeClauseContext)

	// ExitUndoModeClause is called when exiting the undoModeClause production.
	ExitUndoModeClause(c *UndoModeClauseContext)

	// ExitMoveDatafileClause is called when exiting the moveDatafileClause production.
	ExitMoveDatafileClause(c *MoveDatafileClauseContext)

	// ExitInstanceClauses is called when exiting the instanceClauses production.
	ExitInstanceClauses(c *InstanceClausesContext)

	// ExitSecurityClause is called when exiting the securityClause production.
	ExitSecurityClause(c *SecurityClauseContext)

	// ExitPrepareClause is called when exiting the prepareClause production.
	ExitPrepareClause(c *PrepareClauseContext)

	// ExitDropMirrorCopy is called when exiting the dropMirrorCopy production.
	ExitDropMirrorCopy(c *DropMirrorCopyContext)

	// ExitLostWriteProtection is called when exiting the lostWriteProtection production.
	ExitLostWriteProtection(c *LostWriteProtectionContext)

	// ExitCdbFleetClauses is called when exiting the cdbFleetClauses production.
	ExitCdbFleetClauses(c *CdbFleetClausesContext)

	// ExitLeadCdbClause is called when exiting the leadCdbClause production.
	ExitLeadCdbClause(c *LeadCdbClauseContext)

	// ExitLeadCdbUriClause is called when exiting the leadCdbUriClause production.
	ExitLeadCdbUriClause(c *LeadCdbUriClauseContext)

	// ExitPropertyClause is called when exiting the propertyClause production.
	ExitPropertyClause(c *PropertyClauseContext)

	// ExitAlterSystem is called when exiting the alterSystem production.
	ExitAlterSystem(c *AlterSystemContext)

	// ExitAlterSystemOption is called when exiting the alterSystemOption production.
	ExitAlterSystemOption(c *AlterSystemOptionContext)

	// ExitArchiveLogClause is called when exiting the archiveLogClause production.
	ExitArchiveLogClause(c *ArchiveLogClauseContext)

	// ExitCheckpointClause is called when exiting the checkpointClause production.
	ExitCheckpointClause(c *CheckpointClauseContext)

	// ExitCheckDatafilesClause is called when exiting the checkDatafilesClause production.
	ExitCheckDatafilesClause(c *CheckDatafilesClauseContext)

	// ExitDistributedRecovClauses is called when exiting the distributedRecovClauses production.
	ExitDistributedRecovClauses(c *DistributedRecovClausesContext)

	// ExitFlushClause is called when exiting the flushClause production.
	ExitFlushClause(c *FlushClauseContext)

	// ExitEndSessionClauses is called when exiting the endSessionClauses production.
	ExitEndSessionClauses(c *EndSessionClausesContext)

	// ExitAlterSystemSwitchLogfileClause is called when exiting the alterSystemSwitchLogfileClause production.
	ExitAlterSystemSwitchLogfileClause(c *AlterSystemSwitchLogfileClauseContext)

	// ExitSuspendResumeClause is called when exiting the suspendResumeClause production.
	ExitSuspendResumeClause(c *SuspendResumeClauseContext)

	// ExitQuiesceClauses is called when exiting the quiesceClauses production.
	ExitQuiesceClauses(c *QuiesceClausesContext)

	// ExitRollingMigrationClauses is called when exiting the rollingMigrationClauses production.
	ExitRollingMigrationClauses(c *RollingMigrationClausesContext)

	// ExitRollingPatchClauses is called when exiting the rollingPatchClauses production.
	ExitRollingPatchClauses(c *RollingPatchClausesContext)

	// ExitAlterSystemSecurityClauses is called when exiting the alterSystemSecurityClauses production.
	ExitAlterSystemSecurityClauses(c *AlterSystemSecurityClausesContext)

	// ExitAffinityClauses is called when exiting the affinityClauses production.
	ExitAffinityClauses(c *AffinityClausesContext)

	// ExitShutdownDispatcherClause is called when exiting the shutdownDispatcherClause production.
	ExitShutdownDispatcherClause(c *ShutdownDispatcherClauseContext)

	// ExitRegisterClause is called when exiting the registerClause production.
	ExitRegisterClause(c *RegisterClauseContext)

	// ExitSetClause is called when exiting the setClause production.
	ExitSetClause(c *SetClauseContext)

	// ExitResetClause is called when exiting the resetClause production.
	ExitResetClause(c *ResetClauseContext)

	// ExitRelocateClientClause is called when exiting the relocateClientClause production.
	ExitRelocateClientClause(c *RelocateClientClauseContext)

	// ExitCancelSqlClause is called when exiting the cancelSqlClause production.
	ExitCancelSqlClause(c *CancelSqlClauseContext)

	// ExitFlushPasswordfileMetadataCacheClause is called when exiting the flushPasswordfileMetadataCacheClause production.
	ExitFlushPasswordfileMetadataCacheClause(c *FlushPasswordfileMetadataCacheClauseContext)

	// ExitInstanceClause is called when exiting the instanceClause production.
	ExitInstanceClause(c *InstanceClauseContext)

	// ExitSequenceClause is called when exiting the sequenceClause production.
	ExitSequenceClause(c *SequenceClauseContext)

	// ExitChangeClause is called when exiting the changeClause production.
	ExitChangeClause(c *ChangeClauseContext)

	// ExitCurrentClause is called when exiting the currentClause production.
	ExitCurrentClause(c *CurrentClauseContext)

	// ExitGroupClause is called when exiting the groupClause production.
	ExitGroupClause(c *GroupClauseContext)

	// ExitLogfileClause is called when exiting the logfileClause production.
	ExitLogfileClause(c *LogfileClauseContext)

	// ExitNextClause is called when exiting the nextClause production.
	ExitNextClause(c *NextClauseContext)

	// ExitAllClause is called when exiting the allClause production.
	ExitAllClause(c *AllClauseContext)

	// ExitToLocationClause is called when exiting the toLocationClause production.
	ExitToLocationClause(c *ToLocationClauseContext)

	// ExitFlushClauseOption is called when exiting the flushClauseOption production.
	ExitFlushClauseOption(c *FlushClauseOptionContext)

	// ExitDisconnectSessionClause is called when exiting the disconnectSessionClause production.
	ExitDisconnectSessionClause(c *DisconnectSessionClauseContext)

	// ExitKillSessionClause is called when exiting the killSessionClause production.
	ExitKillSessionClause(c *KillSessionClauseContext)

	// ExitStartRollingMigrationClause is called when exiting the startRollingMigrationClause production.
	ExitStartRollingMigrationClause(c *StartRollingMigrationClauseContext)

	// ExitStopRollingMigrationClause is called when exiting the stopRollingMigrationClause production.
	ExitStopRollingMigrationClause(c *StopRollingMigrationClauseContext)

	// ExitStartRollingPatchClause is called when exiting the startRollingPatchClause production.
	ExitStartRollingPatchClause(c *StartRollingPatchClauseContext)

	// ExitStopRollingPatchClause is called when exiting the stopRollingPatchClause production.
	ExitStopRollingPatchClause(c *StopRollingPatchClauseContext)

	// ExitRestrictedSessionClause is called when exiting the restrictedSessionClause production.
	ExitRestrictedSessionClause(c *RestrictedSessionClauseContext)

	// ExitSetEncryptionWalletOpenClause is called when exiting the setEncryptionWalletOpenClause production.
	ExitSetEncryptionWalletOpenClause(c *SetEncryptionWalletOpenClauseContext)

	// ExitSetEncryptionWalletCloseClause is called when exiting the setEncryptionWalletCloseClause production.
	ExitSetEncryptionWalletCloseClause(c *SetEncryptionWalletCloseClauseContext)

	// ExitSetEncryptionKeyClause is called when exiting the setEncryptionKeyClause production.
	ExitSetEncryptionKeyClause(c *SetEncryptionKeyClauseContext)

	// ExitEnableAffinityClause is called when exiting the enableAffinityClause production.
	ExitEnableAffinityClause(c *EnableAffinityClauseContext)

	// ExitDisableAffinityClause is called when exiting the disableAffinityClause production.
	ExitDisableAffinityClause(c *DisableAffinityClauseContext)

	// ExitAlterSystemSetClause is called when exiting the alterSystemSetClause production.
	ExitAlterSystemSetClause(c *AlterSystemSetClauseContext)

	// ExitAlterSystemResetClause is called when exiting the alterSystemResetClause production.
	ExitAlterSystemResetClause(c *AlterSystemResetClauseContext)

	// ExitSharedPoolClause is called when exiting the sharedPoolClause production.
	ExitSharedPoolClause(c *SharedPoolClauseContext)

	// ExitGlobalContextClause is called when exiting the globalContextClause production.
	ExitGlobalContextClause(c *GlobalContextClauseContext)

	// ExitBufferCacheClause is called when exiting the bufferCacheClause production.
	ExitBufferCacheClause(c *BufferCacheClauseContext)

	// ExitFlashCacheClause is called when exiting the flashCacheClause production.
	ExitFlashCacheClause(c *FlashCacheClauseContext)

	// ExitRedoToClause is called when exiting the redoToClause production.
	ExitRedoToClause(c *RedoToClauseContext)

	// ExitIdentifiedByWalletPassword is called when exiting the identifiedByWalletPassword production.
	ExitIdentifiedByWalletPassword(c *IdentifiedByWalletPasswordContext)

	// ExitIdentifiedByHsmAuthString is called when exiting the identifiedByHsmAuthString production.
	ExitIdentifiedByHsmAuthString(c *IdentifiedByHsmAuthStringContext)

	// ExitSetParameterClause is called when exiting the setParameterClause production.
	ExitSetParameterClause(c *SetParameterClauseContext)

	// ExitUseStoredOutlinesClause is called when exiting the useStoredOutlinesClause production.
	ExitUseStoredOutlinesClause(c *UseStoredOutlinesClauseContext)

	// ExitGlobalTopicEnabledClause is called when exiting the globalTopicEnabledClause production.
	ExitGlobalTopicEnabledClause(c *GlobalTopicEnabledClauseContext)

	// ExitAlterSystemCommentClause is called when exiting the alterSystemCommentClause production.
	ExitAlterSystemCommentClause(c *AlterSystemCommentClauseContext)

	// ExitContainerCurrentAllClause is called when exiting the containerCurrentAllClause production.
	ExitContainerCurrentAllClause(c *ContainerCurrentAllClauseContext)

	// ExitScopeClause is called when exiting the scopeClause production.
	ExitScopeClause(c *ScopeClauseContext)

	// ExitAnalyze is called when exiting the analyze production.
	ExitAnalyze(c *AnalyzeContext)

	// ExitPartitionExtensionClause is called when exiting the partitionExtensionClause production.
	ExitPartitionExtensionClause(c *PartitionExtensionClauseContext)

	// ExitValidationClauses is called when exiting the validationClauses production.
	ExitValidationClauses(c *ValidationClausesContext)

	// ExitAssociateStatistics is called when exiting the associateStatistics production.
	ExitAssociateStatistics(c *AssociateStatisticsContext)

	// ExitColumnAssociation is called when exiting the columnAssociation production.
	ExitColumnAssociation(c *ColumnAssociationContext)

	// ExitFunctionAssociation is called when exiting the functionAssociation production.
	ExitFunctionAssociation(c *FunctionAssociationContext)

	// ExitStorageTableClause is called when exiting the storageTableClause production.
	ExitStorageTableClause(c *StorageTableClauseContext)

	// ExitUsingStatisticsType is called when exiting the usingStatisticsType production.
	ExitUsingStatisticsType(c *UsingStatisticsTypeContext)

	// ExitDefaultCostClause is called when exiting the defaultCostClause production.
	ExitDefaultCostClause(c *DefaultCostClauseContext)

	// ExitDefaultSelectivityClause is called when exiting the defaultSelectivityClause production.
	ExitDefaultSelectivityClause(c *DefaultSelectivityClauseContext)

	// ExitDisassociateStatistics is called when exiting the disassociateStatistics production.
	ExitDisassociateStatistics(c *DisassociateStatisticsContext)

	// ExitAudit is called when exiting the audit production.
	ExitAudit(c *AuditContext)

	// ExitNoAudit is called when exiting the noAudit production.
	ExitNoAudit(c *NoAuditContext)

	// ExitAuditPolicyClause is called when exiting the auditPolicyClause production.
	ExitAuditPolicyClause(c *AuditPolicyClauseContext)

	// ExitNoAuditPolicyClause is called when exiting the noAuditPolicyClause production.
	ExitNoAuditPolicyClause(c *NoAuditPolicyClauseContext)

	// ExitByUsersWithRoles is called when exiting the byUsersWithRoles production.
	ExitByUsersWithRoles(c *ByUsersWithRolesContext)

	// ExitContextClause is called when exiting the contextClause production.
	ExitContextClause(c *ContextClauseContext)

	// ExitContextNamespaceAttributesClause is called when exiting the contextNamespaceAttributesClause production.
	ExitContextNamespaceAttributesClause(c *ContextNamespaceAttributesClauseContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitFlashbackDatabase is called when exiting the flashbackDatabase production.
	ExitFlashbackDatabase(c *FlashbackDatabaseContext)

	// ExitScnTimestampClause is called when exiting the scnTimestampClause production.
	ExitScnTimestampClause(c *ScnTimestampClauseContext)

	// ExitRestorePointClause is called when exiting the restorePointClause production.
	ExitRestorePointClause(c *RestorePointClauseContext)

	// ExitFlashbackTable is called when exiting the flashbackTable production.
	ExitFlashbackTable(c *FlashbackTableContext)

	// ExitRenameToTable is called when exiting the renameToTable production.
	ExitRenameToTable(c *RenameToTableContext)

	// ExitPurge is called when exiting the purge production.
	ExitPurge(c *PurgeContext)

	// ExitRename is called when exiting the rename production.
	ExitRename(c *RenameContext)

	// ExitCreateDatabase is called when exiting the createDatabase production.
	ExitCreateDatabase(c *CreateDatabaseContext)

	// ExitCreateDatabaseClauses is called when exiting the createDatabaseClauses production.
	ExitCreateDatabaseClauses(c *CreateDatabaseClausesContext)

	// ExitDatabaseLoggingClauses is called when exiting the databaseLoggingClauses production.
	ExitDatabaseLoggingClauses(c *DatabaseLoggingClausesContext)

	// ExitTablespaceClauses is called when exiting the tablespaceClauses production.
	ExitTablespaceClauses(c *TablespaceClausesContext)

	// ExitDefaultTablespace is called when exiting the defaultTablespace production.
	ExitDefaultTablespace(c *DefaultTablespaceContext)

	// ExitDefaultTempTablespace is called when exiting the defaultTempTablespace production.
	ExitDefaultTempTablespace(c *DefaultTempTablespaceContext)

	// ExitUndoTablespace is called when exiting the undoTablespace production.
	ExitUndoTablespace(c *UndoTablespaceContext)

	// ExitBigOrSmallFiles is called when exiting the bigOrSmallFiles production.
	ExitBigOrSmallFiles(c *BigOrSmallFilesContext)

	// ExitExtentManagementClause is called when exiting the extentManagementClause production.
	ExitExtentManagementClause(c *ExtentManagementClauseContext)

	// ExitEnablePluggableDatabase is called when exiting the enablePluggableDatabase production.
	ExitEnablePluggableDatabase(c *EnablePluggableDatabaseContext)

	// ExitFileNameConvert is called when exiting the fileNameConvert production.
	ExitFileNameConvert(c *FileNameConvertContext)

	// ExitReplaceFileNamePattern is called when exiting the replaceFileNamePattern production.
	ExitReplaceFileNamePattern(c *ReplaceFileNamePatternContext)

	// ExitTablespaceDatafileClauses is called when exiting the tablespaceDatafileClauses production.
	ExitTablespaceDatafileClauses(c *TablespaceDatafileClausesContext)

	// ExitCreateDatabaseLink is called when exiting the createDatabaseLink production.
	ExitCreateDatabaseLink(c *CreateDatabaseLinkContext)

	// ExitConnectToClause is called when exiting the connectToClause production.
	ExitConnectToClause(c *ConnectToClauseContext)

	// ExitDbLinkAuthentication is called when exiting the dbLinkAuthentication production.
	ExitDbLinkAuthentication(c *DbLinkAuthenticationContext)

	// ExitSetTransaction is called when exiting the setTransaction production.
	ExitSetTransaction(c *SetTransactionContext)

	// ExitCommit is called when exiting the commit production.
	ExitCommit(c *CommitContext)

	// ExitCommentClause is called when exiting the commentClause production.
	ExitCommentClause(c *CommentClauseContext)

	// ExitWriteClause is called when exiting the writeClause production.
	ExitWriteClause(c *WriteClauseContext)

	// ExitForceClause is called when exiting the forceClause production.
	ExitForceClause(c *ForceClauseContext)

	// ExitRollback is called when exiting the rollback production.
	ExitRollback(c *RollbackContext)

	// ExitSavepointClause is called when exiting the savepointClause production.
	ExitSavepointClause(c *SavepointClauseContext)

	// ExitSavepoint is called when exiting the savepoint production.
	ExitSavepoint(c *SavepointContext)

	// ExitSetConstraints is called when exiting the setConstraints production.
	ExitSetConstraints(c *SetConstraintsContext)

	// ExitGrant is called when exiting the grant production.
	ExitGrant(c *GrantContext)

	// ExitRevoke is called when exiting the revoke production.
	ExitRevoke(c *RevokeContext)

	// ExitObjectPrivilegeClause is called when exiting the objectPrivilegeClause production.
	ExitObjectPrivilegeClause(c *ObjectPrivilegeClauseContext)

	// ExitSystemPrivilegeClause is called when exiting the systemPrivilegeClause production.
	ExitSystemPrivilegeClause(c *SystemPrivilegeClauseContext)

	// ExitRoleClause is called when exiting the roleClause production.
	ExitRoleClause(c *RoleClauseContext)

	// ExitObjectPrivileges is called when exiting the objectPrivileges production.
	ExitObjectPrivileges(c *ObjectPrivilegesContext)

	// ExitObjectPrivilegeType is called when exiting the objectPrivilegeType production.
	ExitObjectPrivilegeType(c *ObjectPrivilegeTypeContext)

	// ExitOnObjectClause is called when exiting the onObjectClause production.
	ExitOnObjectClause(c *OnObjectClauseContext)

	// ExitSystemPrivilege is called when exiting the systemPrivilege production.
	ExitSystemPrivilege(c *SystemPrivilegeContext)

	// ExitSystemPrivilegeOperation is called when exiting the systemPrivilegeOperation production.
	ExitSystemPrivilegeOperation(c *SystemPrivilegeOperationContext)

	// ExitAdvisorFrameworkSystemPrivilege is called when exiting the advisorFrameworkSystemPrivilege production.
	ExitAdvisorFrameworkSystemPrivilege(c *AdvisorFrameworkSystemPrivilegeContext)

	// ExitClustersSystemPrivilege is called when exiting the clustersSystemPrivilege production.
	ExitClustersSystemPrivilege(c *ClustersSystemPrivilegeContext)

	// ExitContextsSystemPrivilege is called when exiting the contextsSystemPrivilege production.
	ExitContextsSystemPrivilege(c *ContextsSystemPrivilegeContext)

	// ExitDataRedactionSystemPrivilege is called when exiting the dataRedactionSystemPrivilege production.
	ExitDataRedactionSystemPrivilege(c *DataRedactionSystemPrivilegeContext)

	// ExitDatabaseSystemPrivilege is called when exiting the databaseSystemPrivilege production.
	ExitDatabaseSystemPrivilege(c *DatabaseSystemPrivilegeContext)

	// ExitDatabaseLinksSystemPrivilege is called when exiting the databaseLinksSystemPrivilege production.
	ExitDatabaseLinksSystemPrivilege(c *DatabaseLinksSystemPrivilegeContext)

	// ExitDebuggingSystemPrivilege is called when exiting the debuggingSystemPrivilege production.
	ExitDebuggingSystemPrivilege(c *DebuggingSystemPrivilegeContext)

	// ExitDictionariesSystemPrivilege is called when exiting the dictionariesSystemPrivilege production.
	ExitDictionariesSystemPrivilege(c *DictionariesSystemPrivilegeContext)

	// ExitDimensionsSystemPrivilege is called when exiting the dimensionsSystemPrivilege production.
	ExitDimensionsSystemPrivilege(c *DimensionsSystemPrivilegeContext)

	// ExitDirectoriesSystemPrivilege is called when exiting the directoriesSystemPrivilege production.
	ExitDirectoriesSystemPrivilege(c *DirectoriesSystemPrivilegeContext)

	// ExitEditionsSystemPrivilege is called when exiting the editionsSystemPrivilege production.
	ExitEditionsSystemPrivilege(c *EditionsSystemPrivilegeContext)

	// ExitFlashbackDataArchivesPrivilege is called when exiting the flashbackDataArchivesPrivilege production.
	ExitFlashbackDataArchivesPrivilege(c *FlashbackDataArchivesPrivilegeContext)

	// ExitIndexesSystemPrivilege is called when exiting the indexesSystemPrivilege production.
	ExitIndexesSystemPrivilege(c *IndexesSystemPrivilegeContext)

	// ExitIndexTypesSystemPrivilege is called when exiting the indexTypesSystemPrivilege production.
	ExitIndexTypesSystemPrivilege(c *IndexTypesSystemPrivilegeContext)

	// ExitJobSchedulerObjectsSystemPrivilege is called when exiting the jobSchedulerObjectsSystemPrivilege production.
	ExitJobSchedulerObjectsSystemPrivilege(c *JobSchedulerObjectsSystemPrivilegeContext)

	// ExitKeyManagementFrameworkSystemPrivilege is called when exiting the keyManagementFrameworkSystemPrivilege production.
	ExitKeyManagementFrameworkSystemPrivilege(c *KeyManagementFrameworkSystemPrivilegeContext)

	// ExitLibrariesFrameworkSystemPrivilege is called when exiting the librariesFrameworkSystemPrivilege production.
	ExitLibrariesFrameworkSystemPrivilege(c *LibrariesFrameworkSystemPrivilegeContext)

	// ExitLogminerFrameworkSystemPrivilege is called when exiting the logminerFrameworkSystemPrivilege production.
	ExitLogminerFrameworkSystemPrivilege(c *LogminerFrameworkSystemPrivilegeContext)

	// ExitMaterizlizedViewsSystemPrivilege is called when exiting the materizlizedViewsSystemPrivilege production.
	ExitMaterizlizedViewsSystemPrivilege(c *MaterizlizedViewsSystemPrivilegeContext)

	// ExitMiningModelsSystemPrivilege is called when exiting the miningModelsSystemPrivilege production.
	ExitMiningModelsSystemPrivilege(c *MiningModelsSystemPrivilegeContext)

	// ExitOlapCubesSystemPrivilege is called when exiting the olapCubesSystemPrivilege production.
	ExitOlapCubesSystemPrivilege(c *OlapCubesSystemPrivilegeContext)

	// ExitOlapCubeMeasureFoldersSystemPrivilege is called when exiting the olapCubeMeasureFoldersSystemPrivilege production.
	ExitOlapCubeMeasureFoldersSystemPrivilege(c *OlapCubeMeasureFoldersSystemPrivilegeContext)

	// ExitOlapCubeDiminsionsSystemPrivilege is called when exiting the olapCubeDiminsionsSystemPrivilege production.
	ExitOlapCubeDiminsionsSystemPrivilege(c *OlapCubeDiminsionsSystemPrivilegeContext)

	// ExitOlapCubeBuildProcessesSystemPrivilege is called when exiting the olapCubeBuildProcessesSystemPrivilege production.
	ExitOlapCubeBuildProcessesSystemPrivilege(c *OlapCubeBuildProcessesSystemPrivilegeContext)

	// ExitOperatorsSystemPrivilege is called when exiting the operatorsSystemPrivilege production.
	ExitOperatorsSystemPrivilege(c *OperatorsSystemPrivilegeContext)

	// ExitOutlinesSystemPrivilege is called when exiting the outlinesSystemPrivilege production.
	ExitOutlinesSystemPrivilege(c *OutlinesSystemPrivilegeContext)

	// ExitPlanManagementSystemPrivilege is called when exiting the planManagementSystemPrivilege production.
	ExitPlanManagementSystemPrivilege(c *PlanManagementSystemPrivilegeContext)

	// ExitPluggableDatabasesSystemPrivilege is called when exiting the pluggableDatabasesSystemPrivilege production.
	ExitPluggableDatabasesSystemPrivilege(c *PluggableDatabasesSystemPrivilegeContext)

	// ExitProceduresSystemPrivilege is called when exiting the proceduresSystemPrivilege production.
	ExitProceduresSystemPrivilege(c *ProceduresSystemPrivilegeContext)

	// ExitProfilesSystemPrivilege is called when exiting the profilesSystemPrivilege production.
	ExitProfilesSystemPrivilege(c *ProfilesSystemPrivilegeContext)

	// ExitRolesSystemPrivilege is called when exiting the rolesSystemPrivilege production.
	ExitRolesSystemPrivilege(c *RolesSystemPrivilegeContext)

	// ExitRollbackSegmentsSystemPrivilege is called when exiting the rollbackSegmentsSystemPrivilege production.
	ExitRollbackSegmentsSystemPrivilege(c *RollbackSegmentsSystemPrivilegeContext)

	// ExitSequencesSystemPrivilege is called when exiting the sequencesSystemPrivilege production.
	ExitSequencesSystemPrivilege(c *SequencesSystemPrivilegeContext)

	// ExitSessionsSystemPrivilege is called when exiting the sessionsSystemPrivilege production.
	ExitSessionsSystemPrivilege(c *SessionsSystemPrivilegeContext)

	// ExitSqlTranslationProfilesSystemPrivilege is called when exiting the sqlTranslationProfilesSystemPrivilege production.
	ExitSqlTranslationProfilesSystemPrivilege(c *SqlTranslationProfilesSystemPrivilegeContext)

	// ExitSynonymsSystemPrivilege is called when exiting the synonymsSystemPrivilege production.
	ExitSynonymsSystemPrivilege(c *SynonymsSystemPrivilegeContext)

	// ExitTablesSystemPrivilege is called when exiting the tablesSystemPrivilege production.
	ExitTablesSystemPrivilege(c *TablesSystemPrivilegeContext)

	// ExitTablespacesSystemPrivilege is called when exiting the tablespacesSystemPrivilege production.
	ExitTablespacesSystemPrivilege(c *TablespacesSystemPrivilegeContext)

	// ExitTriggersSystemPrivilege is called when exiting the triggersSystemPrivilege production.
	ExitTriggersSystemPrivilege(c *TriggersSystemPrivilegeContext)

	// ExitTypesSystemPrivilege is called when exiting the typesSystemPrivilege production.
	ExitTypesSystemPrivilege(c *TypesSystemPrivilegeContext)

	// ExitUsersSystemPrivilege is called when exiting the usersSystemPrivilege production.
	ExitUsersSystemPrivilege(c *UsersSystemPrivilegeContext)

	// ExitViewsSystemPrivilege is called when exiting the viewsSystemPrivilege production.
	ExitViewsSystemPrivilege(c *ViewsSystemPrivilegeContext)

	// ExitMiscellaneousSystemPrivilege is called when exiting the miscellaneousSystemPrivilege production.
	ExitMiscellaneousSystemPrivilege(c *MiscellaneousSystemPrivilegeContext)

	// ExitCreateUser is called when exiting the createUser production.
	ExitCreateUser(c *CreateUserContext)

	// ExitDropUser is called when exiting the dropUser production.
	ExitDropUser(c *DropUserContext)

	// ExitAlterUser is called when exiting the alterUser production.
	ExitAlterUser(c *AlterUserContext)

	// ExitCreateRole is called when exiting the createRole production.
	ExitCreateRole(c *CreateRoleContext)

	// ExitDropRole is called when exiting the dropRole production.
	ExitDropRole(c *DropRoleContext)

	// ExitAlterRole is called when exiting the alterRole production.
	ExitAlterRole(c *AlterRoleContext)

	// ExitSetRole is called when exiting the setRole production.
	ExitSetRole(c *SetRoleContext)

	// ExitRoleAssignment is called when exiting the roleAssignment production.
	ExitRoleAssignment(c *RoleAssignmentContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)
}

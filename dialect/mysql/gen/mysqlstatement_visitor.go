// Code generated from E:/GoProject/src/github.com/2212442929/xsqlparser/dialect/mysql/grammer\MySQLStatement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // MySQLStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by MySQLStatementParser.
type MySQLStatementVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MySQLStatementParser#execute.
	VisitExecute(ctx *ExecuteContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#insert.
	VisitInsert(ctx *InsertContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#insertSpecification.
	VisitInsertSpecification(ctx *InsertSpecificationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#insertValuesClause.
	VisitInsertValuesClause(ctx *InsertValuesClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#fields.
	VisitFields(ctx *FieldsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#insertIdentifier.
	VisitInsertIdentifier(ctx *InsertIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableWild.
	VisitTableWild(ctx *TableWildContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#insertSelectClause.
	VisitInsertSelectClause(ctx *InsertSelectClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#onDuplicateKeyClause.
	VisitOnDuplicateKeyClause(ctx *OnDuplicateKeyClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#valueReference.
	VisitValueReference(ctx *ValueReferenceContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#derivedColumns.
	VisitDerivedColumns(ctx *DerivedColumnsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#replace.
	VisitReplace(ctx *ReplaceContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#replaceSpecification.
	VisitReplaceSpecification(ctx *ReplaceSpecificationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#replaceValuesClause.
	VisitReplaceValuesClause(ctx *ReplaceValuesClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#replaceSelectClause.
	VisitReplaceSelectClause(ctx *ReplaceSelectClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#update.
	VisitUpdate(ctx *UpdateContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#updateSpecification_.
	VisitUpdateSpecification_(ctx *UpdateSpecification_Context) interface{}

	// Visit a parse tree produced by MySQLStatementParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#setAssignmentsClause.
	VisitSetAssignmentsClause(ctx *SetAssignmentsClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#assignmentValues.
	VisitAssignmentValues(ctx *AssignmentValuesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#assignmentValue.
	VisitAssignmentValue(ctx *AssignmentValueContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#blobValue.
	VisitBlobValue(ctx *BlobValueContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#delete.
	VisitDelete(ctx *DeleteContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#deleteSpecification.
	VisitDeleteSpecification(ctx *DeleteSpecificationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#singleTableClause.
	VisitSingleTableClause(ctx *SingleTableClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#multipleTablesClause.
	VisitMultipleTablesClause(ctx *MultipleTablesClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#select.
	VisitSelect(ctx *SelectContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#selectWithInto.
	VisitSelectWithInto(ctx *SelectWithIntoContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#queryExpression.
	VisitQueryExpression(ctx *QueryExpressionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#queryExpressionBody.
	VisitQueryExpressionBody(ctx *QueryExpressionBodyContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#unionClause.
	VisitUnionClause(ctx *UnionClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#queryExpressionParens.
	VisitQueryExpressionParens(ctx *QueryExpressionParensContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#queryPrimary.
	VisitQueryPrimary(ctx *QueryPrimaryContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#querySpecification.
	VisitQuerySpecification(ctx *QuerySpecificationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#doStatement.
	VisitDoStatement(ctx *DoStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#handlerStatement.
	VisitHandlerStatement(ctx *HandlerStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#handlerOpenStatement.
	VisitHandlerOpenStatement(ctx *HandlerOpenStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#handlerReadIndexStatement.
	VisitHandlerReadIndexStatement(ctx *HandlerReadIndexStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#handlerReadStatement.
	VisitHandlerReadStatement(ctx *HandlerReadStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#handlerCloseStatement.
	VisitHandlerCloseStatement(ctx *HandlerCloseStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#loadStatement.
	VisitLoadStatement(ctx *LoadStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#loadDataStatement.
	VisitLoadDataStatement(ctx *LoadDataStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#loadXmlStatement.
	VisitLoadXmlStatement(ctx *LoadXmlStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#explicitTable.
	VisitExplicitTable(ctx *ExplicitTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableValueConstructor.
	VisitTableValueConstructor(ctx *TableValueConstructorContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#rowConstructorList.
	VisitRowConstructorList(ctx *RowConstructorListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#cteClause.
	VisitCteClause(ctx *CteClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#selectSpecification.
	VisitSelectSpecification(ctx *SelectSpecificationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#duplicateSpecification.
	VisitDuplicateSpecification(ctx *DuplicateSpecificationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#projections.
	VisitProjections(ctx *ProjectionsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#projection.
	VisitProjection(ctx *ProjectionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#unqualifiedShorthand.
	VisitUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#qualifiedShorthand.
	VisitQualifiedShorthand(ctx *QualifiedShorthandContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableReferences.
	VisitTableReferences(ctx *TableReferencesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#escapedTableReference.
	VisitEscapedTableReference(ctx *EscapedTableReferenceContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableReference.
	VisitTableReference(ctx *TableReferenceContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableFactor.
	VisitTableFactor(ctx *TableFactorContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#partitionNames.
	VisitPartitionNames(ctx *PartitionNamesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#indexHintList.
	VisitIndexHintList(ctx *IndexHintListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#indexHint.
	VisitIndexHint(ctx *IndexHintContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#joinedTable.
	VisitJoinedTable(ctx *JoinedTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#innerJoinType.
	VisitInnerJoinType(ctx *InnerJoinTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#outerJoinType.
	VisitOuterJoinType(ctx *OuterJoinTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#naturalJoinType.
	VisitNaturalJoinType(ctx *NaturalJoinTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#joinSpecification.
	VisitJoinSpecification(ctx *JoinSpecificationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#limitRowCount.
	VisitLimitRowCount(ctx *LimitRowCountContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#limitOffset.
	VisitLimitOffset(ctx *LimitOffsetContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#windowClause.
	VisitWindowClause(ctx *WindowClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#windowItem.
	VisitWindowItem(ctx *WindowItemContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#selectLinesInto.
	VisitSelectLinesInto(ctx *SelectLinesIntoContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#selectFieldsInto.
	VisitSelectFieldsInto(ctx *SelectFieldsIntoContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#selectIntoExpression.
	VisitSelectIntoExpression(ctx *SelectIntoExpressionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#lockClause.
	VisitLockClause(ctx *LockClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#lockClauseList.
	VisitLockClauseList(ctx *LockClauseListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#lockStrength.
	VisitLockStrength(ctx *LockStrengthContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#lockedRowAction.
	VisitLockedRowAction(ctx *LockedRowActionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableLockingList.
	VisitTableLockingList(ctx *TableLockingListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableIdentOptWild.
	VisitTableIdentOptWild(ctx *TableIdentOptWildContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableAliasRefList.
	VisitTableAliasRefList(ctx *TableAliasRefListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#parameterMarker.
	VisitParameterMarker(ctx *ParameterMarkerContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#customKeyword.
	VisitCustomKeyword(ctx *CustomKeywordContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#literals.
	VisitLiterals(ctx *LiteralsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#string_.
	VisitString_(ctx *String_Context) interface{}

	// Visit a parse tree produced by MySQLStatementParser#stringLiterals.
	VisitStringLiterals(ctx *StringLiteralsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#numberLiterals.
	VisitNumberLiterals(ctx *NumberLiteralsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#temporalLiterals.
	VisitTemporalLiterals(ctx *TemporalLiteralsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#hexadecimalLiterals.
	VisitHexadecimalLiterals(ctx *HexadecimalLiteralsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#bitValueLiterals.
	VisitBitValueLiterals(ctx *BitValueLiteralsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#booleanLiterals.
	VisitBooleanLiterals(ctx *BooleanLiteralsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#nullValueLiterals.
	VisitNullValueLiterals(ctx *NullValueLiteralsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#collationName.
	VisitCollationName(ctx *CollationNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#identifierKeywordsUnambiguous.
	VisitIdentifierKeywordsUnambiguous(ctx *IdentifierKeywordsUnambiguousContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#identifierKeywordsAmbiguous1RolesAndLabels.
	VisitIdentifierKeywordsAmbiguous1RolesAndLabels(ctx *IdentifierKeywordsAmbiguous1RolesAndLabelsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#identifierKeywordsAmbiguous2Labels.
	VisitIdentifierKeywordsAmbiguous2Labels(ctx *IdentifierKeywordsAmbiguous2LabelsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#identifierKeywordsAmbiguous3Roles.
	VisitIdentifierKeywordsAmbiguous3Roles(ctx *IdentifierKeywordsAmbiguous3RolesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#identifierKeywordsAmbiguous4SystemVariables.
	VisitIdentifierKeywordsAmbiguous4SystemVariables(ctx *IdentifierKeywordsAmbiguous4SystemVariablesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#textOrIdentifier.
	VisitTextOrIdentifier(ctx *TextOrIdentifierContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#userVariable.
	VisitUserVariable(ctx *UserVariableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#systemVariable.
	VisitSystemVariable(ctx *SystemVariableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#setSystemVariable.
	VisitSetSystemVariable(ctx *SetSystemVariableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#optionType.
	VisitOptionType(ctx *OptionTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#internalVariableName.
	VisitInternalVariableName(ctx *InternalVariableNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#setExprOrDefault.
	VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#transactionCharacteristics.
	VisitTransactionCharacteristics(ctx *TransactionCharacteristicsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#isolationLevel.
	VisitIsolationLevel(ctx *IsolationLevelContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#isolationTypes.
	VisitIsolationTypes(ctx *IsolationTypesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#transactionAccessMode.
	VisitTransactionAccessMode(ctx *TransactionAccessModeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#schemaName.
	VisitSchemaName(ctx *SchemaNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#schemaNames.
	VisitSchemaNames(ctx *SchemaNamesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#charsetName.
	VisitCharsetName(ctx *CharsetNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#schemaPairs.
	VisitSchemaPairs(ctx *SchemaPairsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#schemaPair.
	VisitSchemaPair(ctx *SchemaPairContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#indexName.
	VisitIndexName(ctx *IndexNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#constraintName.
	VisitConstraintName(ctx *ConstraintNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#userIdentifierOrText.
	VisitUserIdentifierOrText(ctx *UserIdentifierOrTextContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#userName.
	VisitUserName(ctx *UserNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#eventName.
	VisitEventName(ctx *EventNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#serverName.
	VisitServerName(ctx *ServerNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#wrapperName.
	VisitWrapperName(ctx *WrapperNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#viewName.
	VisitViewName(ctx *ViewNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#owner.
	VisitOwner(ctx *OwnerContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableList.
	VisitTableList(ctx *TableListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#viewNames.
	VisitViewNames(ctx *ViewNamesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#columnNames.
	VisitColumnNames(ctx *ColumnNamesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#groupName.
	VisitGroupName(ctx *GroupNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#routineName.
	VisitRoutineName(ctx *RoutineNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#shardLibraryName.
	VisitShardLibraryName(ctx *ShardLibraryNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#componentName.
	VisitComponentName(ctx *ComponentNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#pluginName.
	VisitPluginName(ctx *PluginNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#hostName.
	VisitHostName(ctx *HostNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#port.
	VisitPort(ctx *PortContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#cloneInstance.
	VisitCloneInstance(ctx *CloneInstanceContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#cloneDir.
	VisitCloneDir(ctx *CloneDirContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#channelName.
	VisitChannelName(ctx *ChannelNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#logName.
	VisitLogName(ctx *LogNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#roleName.
	VisitRoleName(ctx *RoleNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#roleIdentifierOrText.
	VisitRoleIdentifierOrText(ctx *RoleIdentifierOrTextContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#engineRef.
	VisitEngineRef(ctx *EngineRefContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#triggerName.
	VisitTriggerName(ctx *TriggerNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#triggerTime.
	VisitTriggerTime(ctx *TriggerTimeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableOrTables.
	VisitTableOrTables(ctx *TableOrTablesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#userOrRole.
	VisitUserOrRole(ctx *UserOrRoleContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#partitionName.
	VisitPartitionName(ctx *PartitionNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#allOrPartitionNameList.
	VisitAllOrPartitionNameList(ctx *AllOrPartitionNameListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#triggerEvent.
	VisitTriggerEvent(ctx *TriggerEventContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#triggerOrder.
	VisitTriggerOrder(ctx *TriggerOrderContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#andOperator.
	VisitAndOperator(ctx *AndOperatorContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#orOperator.
	VisitOrOperator(ctx *OrOperatorContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#notOperator.
	VisitNotOperator(ctx *NotOperatorContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#booleanPrimary.
	VisitBooleanPrimary(ctx *BooleanPrimaryContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#bitExpr.
	VisitBitExpr(ctx *BitExprContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#simpleExpr.
	VisitSimpleExpr(ctx *SimpleExprContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#columnRef.
	VisitColumnRef(ctx *ColumnRefContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#columnRefList.
	VisitColumnRefList(ctx *ColumnRefListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#aggregationFunction.
	VisitAggregationFunction(ctx *AggregationFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#aggregationFunctionName.
	VisitAggregationFunctionName(ctx *AggregationFunctionNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#distinct.
	VisitDistinct(ctx *DistinctContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#overClause.
	VisitOverClause(ctx *OverClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#windowSpecification.
	VisitWindowSpecification(ctx *WindowSpecificationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#frameClause.
	VisitFrameClause(ctx *FrameClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#frameStart.
	VisitFrameStart(ctx *FrameStartContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#frameEnd.
	VisitFrameEnd(ctx *FrameEndContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#frameBetween.
	VisitFrameBetween(ctx *FrameBetweenContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#specialFunction.
	VisitSpecialFunction(ctx *SpecialFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#currentUserFunction.
	VisitCurrentUserFunction(ctx *CurrentUserFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#groupConcatFunction.
	VisitGroupConcatFunction(ctx *GroupConcatFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#windowFunction.
	VisitWindowFunction(ctx *WindowFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#windowingClause.
	VisitWindowingClause(ctx *WindowingClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#leadLagInfo.
	VisitLeadLagInfo(ctx *LeadLagInfoContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#nullTreatment.
	VisitNullTreatment(ctx *NullTreatmentContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#checkType.
	VisitCheckType(ctx *CheckTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#repairType.
	VisitRepairType(ctx *RepairTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#castFunction.
	VisitCastFunction(ctx *CastFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#convertFunction.
	VisitConvertFunction(ctx *ConvertFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#castType.
	VisitCastType(ctx *CastTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#nchar.
	VisitNchar(ctx *NcharContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#positionFunction.
	VisitPositionFunction(ctx *PositionFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#substringFunction.
	VisitSubstringFunction(ctx *SubstringFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#extractFunction.
	VisitExtractFunction(ctx *ExtractFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#charFunction.
	VisitCharFunction(ctx *CharFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#trimFunction.
	VisitTrimFunction(ctx *TrimFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#valuesFunction.
	VisitValuesFunction(ctx *ValuesFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#weightStringFunction.
	VisitWeightStringFunction(ctx *WeightStringFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#levelClause.
	VisitLevelClause(ctx *LevelClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#levelInWeightListElement.
	VisitLevelInWeightListElement(ctx *LevelInWeightListElementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#regularFunction.
	VisitRegularFunction(ctx *RegularFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#shorthandRegularFunction.
	VisitShorthandRegularFunction(ctx *ShorthandRegularFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#completeRegularFunction.
	VisitCompleteRegularFunction(ctx *CompleteRegularFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#regularFunctionName.
	VisitRegularFunctionName(ctx *RegularFunctionNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#matchExpression.
	VisitMatchExpression(ctx *MatchExpressionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#matchSearchModifier.
	VisitMatchSearchModifier(ctx *MatchSearchModifierContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#caseExpression.
	VisitCaseExpression(ctx *CaseExpressionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#datetimeExpr.
	VisitDatetimeExpr(ctx *DatetimeExprContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#binaryLogFileIndexNumber.
	VisitBinaryLogFileIndexNumber(ctx *BinaryLogFileIndexNumberContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#caseWhen.
	VisitCaseWhen(ctx *CaseWhenContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#caseElse.
	VisitCaseElse(ctx *CaseElseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#intervalExpression.
	VisitIntervalExpression(ctx *IntervalExpressionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#intervalValue.
	VisitIntervalValue(ctx *IntervalValueContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#intervalUnit.
	VisitIntervalUnit(ctx *IntervalUnitContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#orderByItem.
	VisitOrderByItem(ctx *OrderByItemContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#stringList.
	VisitStringList(ctx *StringListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#textString.
	VisitTextString(ctx *TextStringContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#textStringHash.
	VisitTextStringHash(ctx *TextStringHashContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#fieldOptions.
	VisitFieldOptions(ctx *FieldOptionsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#precision.
	VisitPrecision(ctx *PrecisionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#typeDatetimePrecision.
	VisitTypeDatetimePrecision(ctx *TypeDatetimePrecisionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#charsetWithOptBinary.
	VisitCharsetWithOptBinary(ctx *CharsetWithOptBinaryContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#ascii.
	VisitAscii(ctx *AsciiContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#unicode.
	VisitUnicode(ctx *UnicodeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#charset.
	VisitCharset(ctx *CharsetContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#defaultCollation.
	VisitDefaultCollation(ctx *DefaultCollationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#defaultEncryption.
	VisitDefaultEncryption(ctx *DefaultEncryptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#defaultCharset.
	VisitDefaultCharset(ctx *DefaultCharsetContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#signedLiteral.
	VisitSignedLiteral(ctx *SignedLiteralContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#now.
	VisitNow(ctx *NowContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#columnFormat.
	VisitColumnFormat(ctx *ColumnFormatContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#storageMedia.
	VisitStorageMedia(ctx *StorageMediaContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#direction.
	VisitDirection(ctx *DirectionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#keyOrIndex.
	VisitKeyOrIndex(ctx *KeyOrIndexContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#fieldLength.
	VisitFieldLength(ctx *FieldLengthContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#characterSet.
	VisitCharacterSet(ctx *CharacterSetContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#collateClause.
	VisitCollateClause(ctx *CollateClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#fieldOrVarSpec.
	VisitFieldOrVarSpec(ctx *FieldOrVarSpecContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#notExistClause.
	VisitNotExistClause(ctx *NotExistClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#existClause.
	VisitExistClause(ctx *ExistClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#connectionId.
	VisitConnectionId(ctx *ConnectionIdContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#labelName.
	VisitLabelName(ctx *LabelNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#cursorName.
	VisitCursorName(ctx *CursorNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#conditionName.
	VisitConditionName(ctx *ConditionNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#unionOption.
	VisitUnionOption(ctx *UnionOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#noWriteToBinLog.
	VisitNoWriteToBinLog(ctx *NoWriteToBinLogContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#channelOption.
	VisitChannelOption(ctx *ChannelOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#preparedStatement.
	VisitPreparedStatement(ctx *PreparedStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#executeStatement.
	VisitExecuteStatement(ctx *ExecuteStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#executeVarList.
	VisitExecuteVarList(ctx *ExecuteVarListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterStatement.
	VisitAlterStatement(ctx *AlterStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createTable.
	VisitCreateTable(ctx *CreateTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#partitionClause.
	VisitPartitionClause(ctx *PartitionClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#partitionTypeDef.
	VisitPartitionTypeDef(ctx *PartitionTypeDefContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#subPartitions.
	VisitSubPartitions(ctx *SubPartitionsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#partitionKeyAlgorithm.
	VisitPartitionKeyAlgorithm(ctx *PartitionKeyAlgorithmContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#duplicateAsQueryExpression.
	VisitDuplicateAsQueryExpression(ctx *DuplicateAsQueryExpressionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterTable.
	VisitAlterTable(ctx *AlterTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#standaloneAlterTableAction.
	VisitStandaloneAlterTableAction(ctx *StandaloneAlterTableActionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterTableActions.
	VisitAlterTableActions(ctx *AlterTableActionsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterTablePartitionOptions.
	VisitAlterTablePartitionOptions(ctx *AlterTablePartitionOptionsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterCommandList.
	VisitAlterCommandList(ctx *AlterCommandListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterList.
	VisitAlterList(ctx *AlterListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createTableOptionsSpaceSeparated.
	VisitCreateTableOptionsSpaceSeparated(ctx *CreateTableOptionsSpaceSeparatedContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#addColumn.
	VisitAddColumn(ctx *AddColumnContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#addTableConstraint.
	VisitAddTableConstraint(ctx *AddTableConstraintContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#changeColumn.
	VisitChangeColumn(ctx *ChangeColumnContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#modifyColumn.
	VisitModifyColumn(ctx *ModifyColumnContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterTableDrop.
	VisitAlterTableDrop(ctx *AlterTableDropContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#disableKeys.
	VisitDisableKeys(ctx *DisableKeysContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#enableKeys.
	VisitEnableKeys(ctx *EnableKeysContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterColumn.
	VisitAlterColumn(ctx *AlterColumnContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterIndex.
	VisitAlterIndex(ctx *AlterIndexContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterCheck.
	VisitAlterCheck(ctx *AlterCheckContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterConstraint.
	VisitAlterConstraint(ctx *AlterConstraintContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#renameColumn.
	VisitRenameColumn(ctx *RenameColumnContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterRenameTable.
	VisitAlterRenameTable(ctx *AlterRenameTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#renameIndex.
	VisitRenameIndex(ctx *RenameIndexContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterConvert.
	VisitAlterConvert(ctx *AlterConvertContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterTableForce.
	VisitAlterTableForce(ctx *AlterTableForceContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterTableOrder.
	VisitAlterTableOrder(ctx *AlterTableOrderContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterOrderList.
	VisitAlterOrderList(ctx *AlterOrderListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableConstraintDef.
	VisitTableConstraintDef(ctx *TableConstraintDefContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterCommandsModifierList.
	VisitAlterCommandsModifierList(ctx *AlterCommandsModifierListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterCommandsModifier.
	VisitAlterCommandsModifier(ctx *AlterCommandsModifierContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#withValidation.
	VisitWithValidation(ctx *WithValidationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#standaloneAlterCommands.
	VisitStandaloneAlterCommands(ctx *StandaloneAlterCommandsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterPartition.
	VisitAlterPartition(ctx *AlterPartitionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#constraintClause.
	VisitConstraintClause(ctx *ConstraintClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableElementList.
	VisitTableElementList(ctx *TableElementListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableElement.
	VisitTableElement(ctx *TableElementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#restrict.
	VisitRestrict(ctx *RestrictContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#fulltextIndexOption.
	VisitFulltextIndexOption(ctx *FulltextIndexOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropTable.
	VisitDropTable(ctx *DropTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropIndex.
	VisitDropIndex(ctx *DropIndexContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterAlgorithmOption.
	VisitAlterAlgorithmOption(ctx *AlterAlgorithmOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterLockOption.
	VisitAlterLockOption(ctx *AlterLockOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#truncateTable.
	VisitTruncateTable(ctx *TruncateTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createIndex.
	VisitCreateIndex(ctx *CreateIndexContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createDatabase.
	VisitCreateDatabase(ctx *CreateDatabaseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterDatabase.
	VisitAlterDatabase(ctx *AlterDatabaseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createDatabaseSpecification_.
	VisitCreateDatabaseSpecification_(ctx *CreateDatabaseSpecification_Context) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterDatabaseSpecification_.
	VisitAlterDatabaseSpecification_(ctx *AlterDatabaseSpecification_Context) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropDatabase.
	VisitDropDatabase(ctx *DropDatabaseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterInstance.
	VisitAlterInstance(ctx *AlterInstanceContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#instanceAction.
	VisitInstanceAction(ctx *InstanceActionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#channel.
	VisitChannel(ctx *ChannelContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createEvent.
	VisitCreateEvent(ctx *CreateEventContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterEvent.
	VisitAlterEvent(ctx *AlterEventContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropEvent.
	VisitDropEvent(ctx *DropEventContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createFunction.
	VisitCreateFunction(ctx *CreateFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterFunction.
	VisitAlterFunction(ctx *AlterFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropFunction.
	VisitDropFunction(ctx *DropFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createProcedure.
	VisitCreateProcedure(ctx *CreateProcedureContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterProcedure.
	VisitAlterProcedure(ctx *AlterProcedureContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropProcedure.
	VisitDropProcedure(ctx *DropProcedureContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createServer.
	VisitCreateServer(ctx *CreateServerContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterServer.
	VisitAlterServer(ctx *AlterServerContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropServer.
	VisitDropServer(ctx *DropServerContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createView.
	VisitCreateView(ctx *CreateViewContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterView.
	VisitAlterView(ctx *AlterViewContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropView.
	VisitDropView(ctx *DropViewContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createTablespace.
	VisitCreateTablespace(ctx *CreateTablespaceContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createTablespaceInnodb.
	VisitCreateTablespaceInnodb(ctx *CreateTablespaceInnodbContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createTablespaceNdb.
	VisitCreateTablespaceNdb(ctx *CreateTablespaceNdbContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterTablespace.
	VisitAlterTablespace(ctx *AlterTablespaceContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterTablespaceNdb.
	VisitAlterTablespaceNdb(ctx *AlterTablespaceNdbContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterTablespaceInnodb.
	VisitAlterTablespaceInnodb(ctx *AlterTablespaceInnodbContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropTablespace.
	VisitDropTablespace(ctx *DropTablespaceContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createLogfileGroup.
	VisitCreateLogfileGroup(ctx *CreateLogfileGroupContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterLogfileGroup.
	VisitAlterLogfileGroup(ctx *AlterLogfileGroupContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropLogfileGroup.
	VisitDropLogfileGroup(ctx *DropLogfileGroupContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createTrigger.
	VisitCreateTrigger(ctx *CreateTriggerContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropTrigger.
	VisitDropTrigger(ctx *DropTriggerContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#renameTable.
	VisitRenameTable(ctx *RenameTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createDefinitionClause.
	VisitCreateDefinitionClause(ctx *CreateDefinitionClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#columnDefinition.
	VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#fieldDefinition.
	VisitFieldDefinition(ctx *FieldDefinitionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#columnAttribute.
	VisitColumnAttribute(ctx *ColumnAttributeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#checkConstraint.
	VisitCheckConstraint(ctx *CheckConstraintContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#constraintEnforcement.
	VisitConstraintEnforcement(ctx *ConstraintEnforcementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#generatedOption.
	VisitGeneratedOption(ctx *GeneratedOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#referenceDefinition.
	VisitReferenceDefinition(ctx *ReferenceDefinitionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#onUpdateDelete.
	VisitOnUpdateDelete(ctx *OnUpdateDeleteContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#referenceOption.
	VisitReferenceOption(ctx *ReferenceOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#indexNameAndType.
	VisitIndexNameAndType(ctx *IndexNameAndTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#indexType.
	VisitIndexType(ctx *IndexTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#indexTypeClause.
	VisitIndexTypeClause(ctx *IndexTypeClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#keyParts.
	VisitKeyParts(ctx *KeyPartsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#keyPart.
	VisitKeyPart(ctx *KeyPartContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#keyPartWithExpression.
	VisitKeyPartWithExpression(ctx *KeyPartWithExpressionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#keyListWithExpression.
	VisitKeyListWithExpression(ctx *KeyListWithExpressionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#indexOption.
	VisitIndexOption(ctx *IndexOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#commonIndexOption.
	VisitCommonIndexOption(ctx *CommonIndexOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#visibility.
	VisitVisibility(ctx *VisibilityContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createLikeClause.
	VisitCreateLikeClause(ctx *CreateLikeClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createIndexSpecification.
	VisitCreateIndexSpecification(ctx *CreateIndexSpecificationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createTableOptions.
	VisitCreateTableOptions(ctx *CreateTableOptionsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createTableOption.
	VisitCreateTableOption(ctx *CreateTableOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createSRSStatement.
	VisitCreateSRSStatement(ctx *CreateSRSStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropSRSStatement.
	VisitDropSRSStatement(ctx *DropSRSStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#srsAttribute.
	VisitSrsAttribute(ctx *SrsAttributeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#place.
	VisitPlace(ctx *PlaceContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#partitionDefinitions.
	VisitPartitionDefinitions(ctx *PartitionDefinitionsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#partitionDefinition.
	VisitPartitionDefinition(ctx *PartitionDefinitionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#partitionLessThanValue.
	VisitPartitionLessThanValue(ctx *PartitionLessThanValueContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#partitionValueList.
	VisitPartitionValueList(ctx *PartitionValueListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#partitionDefinitionOption.
	VisitPartitionDefinitionOption(ctx *PartitionDefinitionOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#subpartitionDefinition.
	VisitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#ownerStatement.
	VisitOwnerStatement(ctx *OwnerStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#scheduleExpression.
	VisitScheduleExpression(ctx *ScheduleExpressionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#timestampValue.
	VisitTimestampValue(ctx *TimestampValueContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#routineBody.
	VisitRoutineBody(ctx *RoutineBodyContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#serverOption.
	VisitServerOption(ctx *ServerOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#routineOption.
	VisitRoutineOption(ctx *RoutineOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#procedureParameter.
	VisitProcedureParameter(ctx *ProcedureParameterContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#fileSizeLiteral.
	VisitFileSizeLiteral(ctx *FileSizeLiteralContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#simpleStatement.
	VisitSimpleStatement(ctx *SimpleStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#validStatement.
	VisitValidStatement(ctx *ValidStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#beginStatement.
	VisitBeginStatement(ctx *BeginStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#declareStatement.
	VisitDeclareStatement(ctx *DeclareStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#flowControlStatement.
	VisitFlowControlStatement(ctx *FlowControlStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#caseStatement.
	VisitCaseStatement(ctx *CaseStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#iterateStatement.
	VisitIterateStatement(ctx *IterateStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#leaveStatement.
	VisitLeaveStatement(ctx *LeaveStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#repeatStatement.
	VisitRepeatStatement(ctx *RepeatStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#cursorStatement.
	VisitCursorStatement(ctx *CursorStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#cursorCloseStatement.
	VisitCursorCloseStatement(ctx *CursorCloseStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#cursorDeclareStatement.
	VisitCursorDeclareStatement(ctx *CursorDeclareStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#cursorFetchStatement.
	VisitCursorFetchStatement(ctx *CursorFetchStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#cursorOpenStatement.
	VisitCursorOpenStatement(ctx *CursorOpenStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#conditionHandlingStatement.
	VisitConditionHandlingStatement(ctx *ConditionHandlingStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#declareConditionStatement.
	VisitDeclareConditionStatement(ctx *DeclareConditionStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#declareHandlerStatement.
	VisitDeclareHandlerStatement(ctx *DeclareHandlerStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#getDiagnosticsStatement.
	VisitGetDiagnosticsStatement(ctx *GetDiagnosticsStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#statementInformationItem.
	VisitStatementInformationItem(ctx *StatementInformationItemContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#conditionInformationItem.
	VisitConditionInformationItem(ctx *ConditionInformationItemContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#conditionNumber.
	VisitConditionNumber(ctx *ConditionNumberContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#statementInformationItemName.
	VisitStatementInformationItemName(ctx *StatementInformationItemNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#conditionInformationItemName.
	VisitConditionInformationItemName(ctx *ConditionInformationItemNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#handlerAction.
	VisitHandlerAction(ctx *HandlerActionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#conditionValue.
	VisitConditionValue(ctx *ConditionValueContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#resignalStatement.
	VisitResignalStatement(ctx *ResignalStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#signalStatement.
	VisitSignalStatement(ctx *SignalStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#signalInformationItem.
	VisitSignalInformationItem(ctx *SignalInformationItemContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#use.
	VisitUse(ctx *UseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#help.
	VisitHelp(ctx *HelpContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#explain.
	VisitExplain(ctx *ExplainContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showDatabases.
	VisitShowDatabases(ctx *ShowDatabasesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showTables.
	VisitShowTables(ctx *ShowTablesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showTableStatus.
	VisitShowTableStatus(ctx *ShowTableStatusContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showColumns.
	VisitShowColumns(ctx *ShowColumnsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showIndex.
	VisitShowIndex(ctx *ShowIndexContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showCreateTable.
	VisitShowCreateTable(ctx *ShowCreateTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#fromSchema.
	VisitFromSchema(ctx *FromSchemaContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#fromTable.
	VisitFromTable(ctx *FromTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showLike.
	VisitShowLike(ctx *ShowLikeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showWhereClause.
	VisitShowWhereClause(ctx *ShowWhereClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showFilter.
	VisitShowFilter(ctx *ShowFilterContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showProfileType.
	VisitShowProfileType(ctx *ShowProfileTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#setVariable.
	VisitSetVariable(ctx *SetVariableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#optionValueList.
	VisitOptionValueList(ctx *OptionValueListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#optionValueNoOptionType.
	VisitOptionValueNoOptionType(ctx *OptionValueNoOptionTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#optionValue.
	VisitOptionValue(ctx *OptionValueContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showBinaryLogs.
	VisitShowBinaryLogs(ctx *ShowBinaryLogsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showBinlogEvents.
	VisitShowBinlogEvents(ctx *ShowBinlogEventsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showCharacterSet.
	VisitShowCharacterSet(ctx *ShowCharacterSetContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showCollation.
	VisitShowCollation(ctx *ShowCollationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showCreateDatabase.
	VisitShowCreateDatabase(ctx *ShowCreateDatabaseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showCreateEvent.
	VisitShowCreateEvent(ctx *ShowCreateEventContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showCreateFunction.
	VisitShowCreateFunction(ctx *ShowCreateFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showCreateProcedure.
	VisitShowCreateProcedure(ctx *ShowCreateProcedureContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showCreateTrigger.
	VisitShowCreateTrigger(ctx *ShowCreateTriggerContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showCreateUser.
	VisitShowCreateUser(ctx *ShowCreateUserContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showCreateView.
	VisitShowCreateView(ctx *ShowCreateViewContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showEngine.
	VisitShowEngine(ctx *ShowEngineContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showEngines.
	VisitShowEngines(ctx *ShowEnginesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showCharset.
	VisitShowCharset(ctx *ShowCharsetContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showErrors.
	VisitShowErrors(ctx *ShowErrorsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showEvents.
	VisitShowEvents(ctx *ShowEventsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showFunctionCode.
	VisitShowFunctionCode(ctx *ShowFunctionCodeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showFunctionStatus.
	VisitShowFunctionStatus(ctx *ShowFunctionStatusContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showGrant.
	VisitShowGrant(ctx *ShowGrantContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showMasterStatus.
	VisitShowMasterStatus(ctx *ShowMasterStatusContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showOpenTables.
	VisitShowOpenTables(ctx *ShowOpenTablesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showPlugins.
	VisitShowPlugins(ctx *ShowPluginsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showPrivileges.
	VisitShowPrivileges(ctx *ShowPrivilegesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showProcedureCode.
	VisitShowProcedureCode(ctx *ShowProcedureCodeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showProcedureStatus.
	VisitShowProcedureStatus(ctx *ShowProcedureStatusContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showProcesslist.
	VisitShowProcesslist(ctx *ShowProcesslistContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showProfile.
	VisitShowProfile(ctx *ShowProfileContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showProfiles.
	VisitShowProfiles(ctx *ShowProfilesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showRelaylogEvent.
	VisitShowRelaylogEvent(ctx *ShowRelaylogEventContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showSlavehost.
	VisitShowSlavehost(ctx *ShowSlavehostContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showSlaveStatus.
	VisitShowSlaveStatus(ctx *ShowSlaveStatusContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showStatus.
	VisitShowStatus(ctx *ShowStatusContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showTriggers.
	VisitShowTriggers(ctx *ShowTriggersContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showVariables.
	VisitShowVariables(ctx *ShowVariablesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showWarnings.
	VisitShowWarnings(ctx *ShowWarningsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showReplicas.
	VisitShowReplicas(ctx *ShowReplicasContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#showReplicaStatus.
	VisitShowReplicaStatus(ctx *ShowReplicaStatusContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#setCharacter.
	VisitSetCharacter(ctx *SetCharacterContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#clone.
	VisitClone(ctx *CloneContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#cloneAction.
	VisitCloneAction(ctx *CloneActionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createLoadableFunction.
	VisitCreateLoadableFunction(ctx *CreateLoadableFunctionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#install.
	VisitInstall(ctx *InstallContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#uninstall.
	VisitUninstall(ctx *UninstallContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#installComponent.
	VisitInstallComponent(ctx *InstallComponentContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#installPlugin.
	VisitInstallPlugin(ctx *InstallPluginContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#uninstallComponent.
	VisitUninstallComponent(ctx *UninstallComponentContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#uninstallPlugin.
	VisitUninstallPlugin(ctx *UninstallPluginContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#analyzeTable.
	VisitAnalyzeTable(ctx *AnalyzeTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#histogram.
	VisitHistogram(ctx *HistogramContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#checkTable.
	VisitCheckTable(ctx *CheckTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#checkTableOption.
	VisitCheckTableOption(ctx *CheckTableOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#checksumTable.
	VisitChecksumTable(ctx *ChecksumTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#optimizeTable.
	VisitOptimizeTable(ctx *OptimizeTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#repairTable.
	VisitRepairTable(ctx *RepairTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterResourceGroup.
	VisitAlterResourceGroup(ctx *AlterResourceGroupContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#vcpuSpec.
	VisitVcpuSpec(ctx *VcpuSpecContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createResourceGroup.
	VisitCreateResourceGroup(ctx *CreateResourceGroupContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropResourceGroup.
	VisitDropResourceGroup(ctx *DropResourceGroupContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#setResourceGroup.
	VisitSetResourceGroup(ctx *SetResourceGroupContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#binlog.
	VisitBinlog(ctx *BinlogContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#cacheIndex.
	VisitCacheIndex(ctx *CacheIndexContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#cacheTableIndexList.
	VisitCacheTableIndexList(ctx *CacheTableIndexListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#partitionList.
	VisitPartitionList(ctx *PartitionListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#flush.
	VisitFlush(ctx *FlushContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#flushOption.
	VisitFlushOption(ctx *FlushOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tablesOption.
	VisitTablesOption(ctx *TablesOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#kill.
	VisitKill(ctx *KillContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#loadIndexInfo.
	VisitLoadIndexInfo(ctx *LoadIndexInfoContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#loadTableIndexList.
	VisitLoadTableIndexList(ctx *LoadTableIndexListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#resetStatement.
	VisitResetStatement(ctx *ResetStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#resetOption.
	VisitResetOption(ctx *ResetOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#resetPersist.
	VisitResetPersist(ctx *ResetPersistContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#restart.
	VisitRestart(ctx *RestartContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#shutdown.
	VisitShutdown(ctx *ShutdownContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#explainType.
	VisitExplainType(ctx *ExplainTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#explainableStatement.
	VisitExplainableStatement(ctx *ExplainableStatementContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#formatName.
	VisitFormatName(ctx *FormatNameContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#show.
	VisitShow(ctx *ShowContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#setTransaction.
	VisitSetTransaction(ctx *SetTransactionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#setAutoCommit.
	VisitSetAutoCommit(ctx *SetAutoCommitContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#beginTransaction.
	VisitBeginTransaction(ctx *BeginTransactionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#transactionCharacteristic.
	VisitTransactionCharacteristic(ctx *TransactionCharacteristicContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#commit.
	VisitCommit(ctx *CommitContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#rollback.
	VisitRollback(ctx *RollbackContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#savepoint.
	VisitSavepoint(ctx *SavepointContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#begin.
	VisitBegin(ctx *BeginContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#lock.
	VisitLock(ctx *LockContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#unlock.
	VisitUnlock(ctx *UnlockContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#releaseSavepoint.
	VisitReleaseSavepoint(ctx *ReleaseSavepointContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#xa.
	VisitXa(ctx *XaContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#optionChain.
	VisitOptionChain(ctx *OptionChainContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#optionRelease.
	VisitOptionRelease(ctx *OptionReleaseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tableLock.
	VisitTableLock(ctx *TableLockContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#lockOption.
	VisitLockOption(ctx *LockOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#xid.
	VisitXid(ctx *XidContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#grantRoleOrPrivilegeTo.
	VisitGrantRoleOrPrivilegeTo(ctx *GrantRoleOrPrivilegeToContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#grantRoleOrPrivilegeOnTo.
	VisitGrantRoleOrPrivilegeOnTo(ctx *GrantRoleOrPrivilegeOnToContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#grantProxy.
	VisitGrantProxy(ctx *GrantProxyContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#revokeFrom.
	VisitRevokeFrom(ctx *RevokeFromContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#revokeOnFrom.
	VisitRevokeOnFrom(ctx *RevokeOnFromContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#userList.
	VisitUserList(ctx *UserListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#roleOrPrivileges.
	VisitRoleOrPrivileges(ctx *RoleOrPrivilegesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#roleOrDynamicPrivilege.
	VisitRoleOrDynamicPrivilege(ctx *RoleOrDynamicPrivilegeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#roleAtHost.
	VisitRoleAtHost(ctx *RoleAtHostContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeSelect.
	VisitStaticPrivilegeSelect(ctx *StaticPrivilegeSelectContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeInsert.
	VisitStaticPrivilegeInsert(ctx *StaticPrivilegeInsertContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeUpdate.
	VisitStaticPrivilegeUpdate(ctx *StaticPrivilegeUpdateContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeReferences.
	VisitStaticPrivilegeReferences(ctx *StaticPrivilegeReferencesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeDelete.
	VisitStaticPrivilegeDelete(ctx *StaticPrivilegeDeleteContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeUsage.
	VisitStaticPrivilegeUsage(ctx *StaticPrivilegeUsageContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeIndex.
	VisitStaticPrivilegeIndex(ctx *StaticPrivilegeIndexContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeAlter.
	VisitStaticPrivilegeAlter(ctx *StaticPrivilegeAlterContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeCreate.
	VisitStaticPrivilegeCreate(ctx *StaticPrivilegeCreateContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeDrop.
	VisitStaticPrivilegeDrop(ctx *StaticPrivilegeDropContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeExecute.
	VisitStaticPrivilegeExecute(ctx *StaticPrivilegeExecuteContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeReload.
	VisitStaticPrivilegeReload(ctx *StaticPrivilegeReloadContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeShutdown.
	VisitStaticPrivilegeShutdown(ctx *StaticPrivilegeShutdownContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeProcess.
	VisitStaticPrivilegeProcess(ctx *StaticPrivilegeProcessContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeFile.
	VisitStaticPrivilegeFile(ctx *StaticPrivilegeFileContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeGrant.
	VisitStaticPrivilegeGrant(ctx *StaticPrivilegeGrantContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeShowDatabases.
	VisitStaticPrivilegeShowDatabases(ctx *StaticPrivilegeShowDatabasesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeSuper.
	VisitStaticPrivilegeSuper(ctx *StaticPrivilegeSuperContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeCreateTemporaryTables.
	VisitStaticPrivilegeCreateTemporaryTables(ctx *StaticPrivilegeCreateTemporaryTablesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeLockTables.
	VisitStaticPrivilegeLockTables(ctx *StaticPrivilegeLockTablesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeReplicationSlave.
	VisitStaticPrivilegeReplicationSlave(ctx *StaticPrivilegeReplicationSlaveContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeReplicationClient.
	VisitStaticPrivilegeReplicationClient(ctx *StaticPrivilegeReplicationClientContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeCreateView.
	VisitStaticPrivilegeCreateView(ctx *StaticPrivilegeCreateViewContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeShowView.
	VisitStaticPrivilegeShowView(ctx *StaticPrivilegeShowViewContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeCreateRoutine.
	VisitStaticPrivilegeCreateRoutine(ctx *StaticPrivilegeCreateRoutineContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeAlterRoutine.
	VisitStaticPrivilegeAlterRoutine(ctx *StaticPrivilegeAlterRoutineContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeCreateUser.
	VisitStaticPrivilegeCreateUser(ctx *StaticPrivilegeCreateUserContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeEvent.
	VisitStaticPrivilegeEvent(ctx *StaticPrivilegeEventContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeTrigger.
	VisitStaticPrivilegeTrigger(ctx *StaticPrivilegeTriggerContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeCreateTablespace.
	VisitStaticPrivilegeCreateTablespace(ctx *StaticPrivilegeCreateTablespaceContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeCreateRole.
	VisitStaticPrivilegeCreateRole(ctx *StaticPrivilegeCreateRoleContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#staticPrivilegeDropRole.
	VisitStaticPrivilegeDropRole(ctx *StaticPrivilegeDropRoleContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#aclType.
	VisitAclType(ctx *AclTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#grantLevelGlobal.
	VisitGrantLevelGlobal(ctx *GrantLevelGlobalContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#grantLevelSchemaGlobal.
	VisitGrantLevelSchemaGlobal(ctx *GrantLevelSchemaGlobalContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#grantLevelTable.
	VisitGrantLevelTable(ctx *GrantLevelTableContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createUser.
	VisitCreateUser(ctx *CreateUserContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createUserEntryNoOption.
	VisitCreateUserEntryNoOption(ctx *CreateUserEntryNoOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createUserEntryIdentifiedBy.
	VisitCreateUserEntryIdentifiedBy(ctx *CreateUserEntryIdentifiedByContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createUserEntryIdentifiedWith.
	VisitCreateUserEntryIdentifiedWith(ctx *CreateUserEntryIdentifiedWithContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createUserList.
	VisitCreateUserList(ctx *CreateUserListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#defaultRoleClause.
	VisitDefaultRoleClause(ctx *DefaultRoleClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#requireClause.
	VisitRequireClause(ctx *RequireClauseContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#connectOptions.
	VisitConnectOptions(ctx *ConnectOptionsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#accountLockPasswordExpireOptions.
	VisitAccountLockPasswordExpireOptions(ctx *AccountLockPasswordExpireOptionsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#accountLockPasswordExpireOption.
	VisitAccountLockPasswordExpireOption(ctx *AccountLockPasswordExpireOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterUser.
	VisitAlterUser(ctx *AlterUserContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterUserEntry.
	VisitAlterUserEntry(ctx *AlterUserEntryContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#alterUserList.
	VisitAlterUserList(ctx *AlterUserListContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropUser.
	VisitDropUser(ctx *DropUserContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#createRole.
	VisitCreateRole(ctx *CreateRoleContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#dropRole.
	VisitDropRole(ctx *DropRoleContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#renameUser.
	VisitRenameUser(ctx *RenameUserContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#setDefaultRole.
	VisitSetDefaultRole(ctx *SetDefaultRoleContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#setRole.
	VisitSetRole(ctx *SetRoleContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#setPassword.
	VisitSetPassword(ctx *SetPasswordContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#authOption.
	VisitAuthOption(ctx *AuthOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#withGrantOption.
	VisitWithGrantOption(ctx *WithGrantOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#userOrRoles.
	VisitUserOrRoles(ctx *UserOrRolesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#roles.
	VisitRoles(ctx *RolesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#grantAs.
	VisitGrantAs(ctx *GrantAsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#withRoles.
	VisitWithRoles(ctx *WithRolesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#userAuthOption.
	VisitUserAuthOption(ctx *UserAuthOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#identifiedBy.
	VisitIdentifiedBy(ctx *IdentifiedByContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#identifiedWith.
	VisitIdentifiedWith(ctx *IdentifiedWithContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#connectOption.
	VisitConnectOption(ctx *ConnectOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#tlsOption.
	VisitTlsOption(ctx *TlsOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#userFuncAuthOption.
	VisitUserFuncAuthOption(ctx *UserFuncAuthOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#change.
	VisitChange(ctx *ChangeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#changeMasterTo.
	VisitChangeMasterTo(ctx *ChangeMasterToContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#changeReplicationFilter.
	VisitChangeReplicationFilter(ctx *ChangeReplicationFilterContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#startSlave.
	VisitStartSlave(ctx *StartSlaveContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#stopSlave.
	VisitStopSlave(ctx *StopSlaveContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#groupReplication.
	VisitGroupReplication(ctx *GroupReplicationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#startGroupReplication.
	VisitStartGroupReplication(ctx *StartGroupReplicationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#stopGroupReplication.
	VisitStopGroupReplication(ctx *StopGroupReplicationContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#purgeBinaryLog.
	VisitPurgeBinaryLog(ctx *PurgeBinaryLogContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#threadTypes.
	VisitThreadTypes(ctx *ThreadTypesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#threadType.
	VisitThreadType(ctx *ThreadTypeContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#utilOption.
	VisitUtilOption(ctx *UtilOptionContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#connectionOptions.
	VisitConnectionOptions(ctx *ConnectionOptionsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#masterDefs.
	VisitMasterDefs(ctx *MasterDefsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#masterDef.
	VisitMasterDef(ctx *MasterDefContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#ignoreServerIds.
	VisitIgnoreServerIds(ctx *IgnoreServerIdsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#ignoreServerId.
	VisitIgnoreServerId(ctx *IgnoreServerIdContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#filterDefs.
	VisitFilterDefs(ctx *FilterDefsContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#filterDef.
	VisitFilterDef(ctx *FilterDefContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#wildTables.
	VisitWildTables(ctx *WildTablesContext) interface{}

	// Visit a parse tree produced by MySQLStatementParser#wildTable.
	VisitWildTable(ctx *WildTableContext) interface{}
}

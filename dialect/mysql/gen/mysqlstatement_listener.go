// Code generated from E:/GoProject/src/github.com/2212442929/xsqlparser/dialect/mysql/grammer\MySQLStatement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // MySQLStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MySQLStatementListener is a complete listener for a parse tree produced by MySQLStatementParser.
type MySQLStatementListener interface {
	antlr.ParseTreeListener

	// EnterExecute is called when entering the execute production.
	EnterExecute(c *ExecuteContext)

	// EnterInsert is called when entering the insert production.
	EnterInsert(c *InsertContext)

	// EnterInsertSpecification is called when entering the insertSpecification production.
	EnterInsertSpecification(c *InsertSpecificationContext)

	// EnterInsertValuesClause is called when entering the insertValuesClause production.
	EnterInsertValuesClause(c *InsertValuesClauseContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterInsertIdentifier is called when entering the insertIdentifier production.
	EnterInsertIdentifier(c *InsertIdentifierContext)

	// EnterTableWild is called when entering the tableWild production.
	EnterTableWild(c *TableWildContext)

	// EnterInsertSelectClause is called when entering the insertSelectClause production.
	EnterInsertSelectClause(c *InsertSelectClauseContext)

	// EnterOnDuplicateKeyClause is called when entering the onDuplicateKeyClause production.
	EnterOnDuplicateKeyClause(c *OnDuplicateKeyClauseContext)

	// EnterValueReference is called when entering the valueReference production.
	EnterValueReference(c *ValueReferenceContext)

	// EnterDerivedColumns is called when entering the derivedColumns production.
	EnterDerivedColumns(c *DerivedColumnsContext)

	// EnterReplace is called when entering the replace production.
	EnterReplace(c *ReplaceContext)

	// EnterReplaceSpecification is called when entering the replaceSpecification production.
	EnterReplaceSpecification(c *ReplaceSpecificationContext)

	// EnterReplaceValuesClause is called when entering the replaceValuesClause production.
	EnterReplaceValuesClause(c *ReplaceValuesClauseContext)

	// EnterReplaceSelectClause is called when entering the replaceSelectClause production.
	EnterReplaceSelectClause(c *ReplaceSelectClauseContext)

	// EnterUpdate is called when entering the update production.
	EnterUpdate(c *UpdateContext)

	// EnterUpdateSpecification_ is called when entering the updateSpecification_ production.
	EnterUpdateSpecification_(c *UpdateSpecification_Context)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterSetAssignmentsClause is called when entering the setAssignmentsClause production.
	EnterSetAssignmentsClause(c *SetAssignmentsClauseContext)

	// EnterAssignmentValues is called when entering the assignmentValues production.
	EnterAssignmentValues(c *AssignmentValuesContext)

	// EnterAssignmentValue is called when entering the assignmentValue production.
	EnterAssignmentValue(c *AssignmentValueContext)

	// EnterBlobValue is called when entering the blobValue production.
	EnterBlobValue(c *BlobValueContext)

	// EnterDelete is called when entering the delete production.
	EnterDelete(c *DeleteContext)

	// EnterDeleteSpecification is called when entering the deleteSpecification production.
	EnterDeleteSpecification(c *DeleteSpecificationContext)

	// EnterSingleTableClause is called when entering the singleTableClause production.
	EnterSingleTableClause(c *SingleTableClauseContext)

	// EnterMultipleTablesClause is called when entering the multipleTablesClause production.
	EnterMultipleTablesClause(c *MultipleTablesClauseContext)

	// EnterSelect is called when entering the select production.
	EnterSelect(c *SelectContext)

	// EnterSelectWithInto is called when entering the selectWithInto production.
	EnterSelectWithInto(c *SelectWithIntoContext)

	// EnterQueryExpression is called when entering the queryExpression production.
	EnterQueryExpression(c *QueryExpressionContext)

	// EnterQueryExpressionBody is called when entering the queryExpressionBody production.
	EnterQueryExpressionBody(c *QueryExpressionBodyContext)

	// EnterUnionClause is called when entering the unionClause production.
	EnterUnionClause(c *UnionClauseContext)

	// EnterQueryExpressionParens is called when entering the queryExpressionParens production.
	EnterQueryExpressionParens(c *QueryExpressionParensContext)

	// EnterQueryPrimary is called when entering the queryPrimary production.
	EnterQueryPrimary(c *QueryPrimaryContext)

	// EnterQuerySpecification is called when entering the querySpecification production.
	EnterQuerySpecification(c *QuerySpecificationContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterDoStatement is called when entering the doStatement production.
	EnterDoStatement(c *DoStatementContext)

	// EnterHandlerStatement is called when entering the handlerStatement production.
	EnterHandlerStatement(c *HandlerStatementContext)

	// EnterHandlerOpenStatement is called when entering the handlerOpenStatement production.
	EnterHandlerOpenStatement(c *HandlerOpenStatementContext)

	// EnterHandlerReadIndexStatement is called when entering the handlerReadIndexStatement production.
	EnterHandlerReadIndexStatement(c *HandlerReadIndexStatementContext)

	// EnterHandlerReadStatement is called when entering the handlerReadStatement production.
	EnterHandlerReadStatement(c *HandlerReadStatementContext)

	// EnterHandlerCloseStatement is called when entering the handlerCloseStatement production.
	EnterHandlerCloseStatement(c *HandlerCloseStatementContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterLoadStatement is called when entering the loadStatement production.
	EnterLoadStatement(c *LoadStatementContext)

	// EnterLoadDataStatement is called when entering the loadDataStatement production.
	EnterLoadDataStatement(c *LoadDataStatementContext)

	// EnterLoadXmlStatement is called when entering the loadXmlStatement production.
	EnterLoadXmlStatement(c *LoadXmlStatementContext)

	// EnterExplicitTable is called when entering the explicitTable production.
	EnterExplicitTable(c *ExplicitTableContext)

	// EnterTableValueConstructor is called when entering the tableValueConstructor production.
	EnterTableValueConstructor(c *TableValueConstructorContext)

	// EnterRowConstructorList is called when entering the rowConstructorList production.
	EnterRowConstructorList(c *RowConstructorListContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterCteClause is called when entering the cteClause production.
	EnterCteClause(c *CteClauseContext)

	// EnterSelectSpecification is called when entering the selectSpecification production.
	EnterSelectSpecification(c *SelectSpecificationContext)

	// EnterDuplicateSpecification is called when entering the duplicateSpecification production.
	EnterDuplicateSpecification(c *DuplicateSpecificationContext)

	// EnterProjections is called when entering the projections production.
	EnterProjections(c *ProjectionsContext)

	// EnterProjection is called when entering the projection production.
	EnterProjection(c *ProjectionContext)

	// EnterUnqualifiedShorthand is called when entering the unqualifiedShorthand production.
	EnterUnqualifiedShorthand(c *UnqualifiedShorthandContext)

	// EnterQualifiedShorthand is called when entering the qualifiedShorthand production.
	EnterQualifiedShorthand(c *QualifiedShorthandContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterTableReferences is called when entering the tableReferences production.
	EnterTableReferences(c *TableReferencesContext)

	// EnterEscapedTableReference is called when entering the escapedTableReference production.
	EnterEscapedTableReference(c *EscapedTableReferenceContext)

	// EnterTableReference is called when entering the tableReference production.
	EnterTableReference(c *TableReferenceContext)

	// EnterTableFactor is called when entering the tableFactor production.
	EnterTableFactor(c *TableFactorContext)

	// EnterPartitionNames is called when entering the partitionNames production.
	EnterPartitionNames(c *PartitionNamesContext)

	// EnterIndexHintList is called when entering the indexHintList production.
	EnterIndexHintList(c *IndexHintListContext)

	// EnterIndexHint is called when entering the indexHint production.
	EnterIndexHint(c *IndexHintContext)

	// EnterJoinedTable is called when entering the joinedTable production.
	EnterJoinedTable(c *JoinedTableContext)

	// EnterInnerJoinType is called when entering the innerJoinType production.
	EnterInnerJoinType(c *InnerJoinTypeContext)

	// EnterOuterJoinType is called when entering the outerJoinType production.
	EnterOuterJoinType(c *OuterJoinTypeContext)

	// EnterNaturalJoinType is called when entering the naturalJoinType production.
	EnterNaturalJoinType(c *NaturalJoinTypeContext)

	// EnterJoinSpecification is called when entering the joinSpecification production.
	EnterJoinSpecification(c *JoinSpecificationContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterLimitRowCount is called when entering the limitRowCount production.
	EnterLimitRowCount(c *LimitRowCountContext)

	// EnterLimitOffset is called when entering the limitOffset production.
	EnterLimitOffset(c *LimitOffsetContext)

	// EnterWindowClause is called when entering the windowClause production.
	EnterWindowClause(c *WindowClauseContext)

	// EnterWindowItem is called when entering the windowItem production.
	EnterWindowItem(c *WindowItemContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterSelectLinesInto is called when entering the selectLinesInto production.
	EnterSelectLinesInto(c *SelectLinesIntoContext)

	// EnterSelectFieldsInto is called when entering the selectFieldsInto production.
	EnterSelectFieldsInto(c *SelectFieldsIntoContext)

	// EnterSelectIntoExpression is called when entering the selectIntoExpression production.
	EnterSelectIntoExpression(c *SelectIntoExpressionContext)

	// EnterLockClause is called when entering the lockClause production.
	EnterLockClause(c *LockClauseContext)

	// EnterLockClauseList is called when entering the lockClauseList production.
	EnterLockClauseList(c *LockClauseListContext)

	// EnterLockStrength is called when entering the lockStrength production.
	EnterLockStrength(c *LockStrengthContext)

	// EnterLockedRowAction is called when entering the lockedRowAction production.
	EnterLockedRowAction(c *LockedRowActionContext)

	// EnterTableLockingList is called when entering the tableLockingList production.
	EnterTableLockingList(c *TableLockingListContext)

	// EnterTableIdentOptWild is called when entering the tableIdentOptWild production.
	EnterTableIdentOptWild(c *TableIdentOptWildContext)

	// EnterTableAliasRefList is called when entering the tableAliasRefList production.
	EnterTableAliasRefList(c *TableAliasRefListContext)

	// EnterParameterMarker is called when entering the parameterMarker production.
	EnterParameterMarker(c *ParameterMarkerContext)

	// EnterCustomKeyword is called when entering the customKeyword production.
	EnterCustomKeyword(c *CustomKeywordContext)

	// EnterLiterals is called when entering the literals production.
	EnterLiterals(c *LiteralsContext)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterStringLiterals is called when entering the stringLiterals production.
	EnterStringLiterals(c *StringLiteralsContext)

	// EnterNumberLiterals is called when entering the numberLiterals production.
	EnterNumberLiterals(c *NumberLiteralsContext)

	// EnterTemporalLiterals is called when entering the temporalLiterals production.
	EnterTemporalLiterals(c *TemporalLiteralsContext)

	// EnterHexadecimalLiterals is called when entering the hexadecimalLiterals production.
	EnterHexadecimalLiterals(c *HexadecimalLiteralsContext)

	// EnterBitValueLiterals is called when entering the bitValueLiterals production.
	EnterBitValueLiterals(c *BitValueLiteralsContext)

	// EnterBooleanLiterals is called when entering the booleanLiterals production.
	EnterBooleanLiterals(c *BooleanLiteralsContext)

	// EnterNullValueLiterals is called when entering the nullValueLiterals production.
	EnterNullValueLiterals(c *NullValueLiteralsContext)

	// EnterCollationName is called when entering the collationName production.
	EnterCollationName(c *CollationNameContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIdentifierKeywordsUnambiguous is called when entering the identifierKeywordsUnambiguous production.
	EnterIdentifierKeywordsUnambiguous(c *IdentifierKeywordsUnambiguousContext)

	// EnterIdentifierKeywordsAmbiguous1RolesAndLabels is called when entering the identifierKeywordsAmbiguous1RolesAndLabels production.
	EnterIdentifierKeywordsAmbiguous1RolesAndLabels(c *IdentifierKeywordsAmbiguous1RolesAndLabelsContext)

	// EnterIdentifierKeywordsAmbiguous2Labels is called when entering the identifierKeywordsAmbiguous2Labels production.
	EnterIdentifierKeywordsAmbiguous2Labels(c *IdentifierKeywordsAmbiguous2LabelsContext)

	// EnterIdentifierKeywordsAmbiguous3Roles is called when entering the identifierKeywordsAmbiguous3Roles production.
	EnterIdentifierKeywordsAmbiguous3Roles(c *IdentifierKeywordsAmbiguous3RolesContext)

	// EnterIdentifierKeywordsAmbiguous4SystemVariables is called when entering the identifierKeywordsAmbiguous4SystemVariables production.
	EnterIdentifierKeywordsAmbiguous4SystemVariables(c *IdentifierKeywordsAmbiguous4SystemVariablesContext)

	// EnterTextOrIdentifier is called when entering the textOrIdentifier production.
	EnterTextOrIdentifier(c *TextOrIdentifierContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterUserVariable is called when entering the userVariable production.
	EnterUserVariable(c *UserVariableContext)

	// EnterSystemVariable is called when entering the systemVariable production.
	EnterSystemVariable(c *SystemVariableContext)

	// EnterSetSystemVariable is called when entering the setSystemVariable production.
	EnterSetSystemVariable(c *SetSystemVariableContext)

	// EnterOptionType is called when entering the optionType production.
	EnterOptionType(c *OptionTypeContext)

	// EnterInternalVariableName is called when entering the internalVariableName production.
	EnterInternalVariableName(c *InternalVariableNameContext)

	// EnterSetExprOrDefault is called when entering the setExprOrDefault production.
	EnterSetExprOrDefault(c *SetExprOrDefaultContext)

	// EnterTransactionCharacteristics is called when entering the transactionCharacteristics production.
	EnterTransactionCharacteristics(c *TransactionCharacteristicsContext)

	// EnterIsolationLevel is called when entering the isolationLevel production.
	EnterIsolationLevel(c *IsolationLevelContext)

	// EnterIsolationTypes is called when entering the isolationTypes production.
	EnterIsolationTypes(c *IsolationTypesContext)

	// EnterTransactionAccessMode is called when entering the transactionAccessMode production.
	EnterTransactionAccessMode(c *TransactionAccessModeContext)

	// EnterSchemaName is called when entering the schemaName production.
	EnterSchemaName(c *SchemaNameContext)

	// EnterSchemaNames is called when entering the schemaNames production.
	EnterSchemaNames(c *SchemaNamesContext)

	// EnterCharsetName is called when entering the charsetName production.
	EnterCharsetName(c *CharsetNameContext)

	// EnterSchemaPairs is called when entering the schemaPairs production.
	EnterSchemaPairs(c *SchemaPairsContext)

	// EnterSchemaPair is called when entering the schemaPair production.
	EnterSchemaPair(c *SchemaPairContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterIndexName is called when entering the indexName production.
	EnterIndexName(c *IndexNameContext)

	// EnterConstraintName is called when entering the constraintName production.
	EnterConstraintName(c *ConstraintNameContext)

	// EnterUserIdentifierOrText is called when entering the userIdentifierOrText production.
	EnterUserIdentifierOrText(c *UserIdentifierOrTextContext)

	// EnterUserName is called when entering the userName production.
	EnterUserName(c *UserNameContext)

	// EnterEventName is called when entering the eventName production.
	EnterEventName(c *EventNameContext)

	// EnterServerName is called when entering the serverName production.
	EnterServerName(c *ServerNameContext)

	// EnterWrapperName is called when entering the wrapperName production.
	EnterWrapperName(c *WrapperNameContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterViewName is called when entering the viewName production.
	EnterViewName(c *ViewNameContext)

	// EnterOwner is called when entering the owner production.
	EnterOwner(c *OwnerContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterTableList is called when entering the tableList production.
	EnterTableList(c *TableListContext)

	// EnterViewNames is called when entering the viewNames production.
	EnterViewNames(c *ViewNamesContext)

	// EnterColumnNames is called when entering the columnNames production.
	EnterColumnNames(c *ColumnNamesContext)

	// EnterGroupName is called when entering the groupName production.
	EnterGroupName(c *GroupNameContext)

	// EnterRoutineName is called when entering the routineName production.
	EnterRoutineName(c *RoutineNameContext)

	// EnterShardLibraryName is called when entering the shardLibraryName production.
	EnterShardLibraryName(c *ShardLibraryNameContext)

	// EnterComponentName is called when entering the componentName production.
	EnterComponentName(c *ComponentNameContext)

	// EnterPluginName is called when entering the pluginName production.
	EnterPluginName(c *PluginNameContext)

	// EnterHostName is called when entering the hostName production.
	EnterHostName(c *HostNameContext)

	// EnterPort is called when entering the port production.
	EnterPort(c *PortContext)

	// EnterCloneInstance is called when entering the cloneInstance production.
	EnterCloneInstance(c *CloneInstanceContext)

	// EnterCloneDir is called when entering the cloneDir production.
	EnterCloneDir(c *CloneDirContext)

	// EnterChannelName is called when entering the channelName production.
	EnterChannelName(c *ChannelNameContext)

	// EnterLogName is called when entering the logName production.
	EnterLogName(c *LogNameContext)

	// EnterRoleName is called when entering the roleName production.
	EnterRoleName(c *RoleNameContext)

	// EnterRoleIdentifierOrText is called when entering the roleIdentifierOrText production.
	EnterRoleIdentifierOrText(c *RoleIdentifierOrTextContext)

	// EnterEngineRef is called when entering the engineRef production.
	EnterEngineRef(c *EngineRefContext)

	// EnterTriggerName is called when entering the triggerName production.
	EnterTriggerName(c *TriggerNameContext)

	// EnterTriggerTime is called when entering the triggerTime production.
	EnterTriggerTime(c *TriggerTimeContext)

	// EnterTableOrTables is called when entering the tableOrTables production.
	EnterTableOrTables(c *TableOrTablesContext)

	// EnterUserOrRole is called when entering the userOrRole production.
	EnterUserOrRole(c *UserOrRoleContext)

	// EnterPartitionName is called when entering the partitionName production.
	EnterPartitionName(c *PartitionNameContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterAllOrPartitionNameList is called when entering the allOrPartitionNameList production.
	EnterAllOrPartitionNameList(c *AllOrPartitionNameListContext)

	// EnterTriggerEvent is called when entering the triggerEvent production.
	EnterTriggerEvent(c *TriggerEventContext)

	// EnterTriggerOrder is called when entering the triggerOrder production.
	EnterTriggerOrder(c *TriggerOrderContext)

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

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterBitExpr is called when entering the bitExpr production.
	EnterBitExpr(c *BitExprContext)

	// EnterSimpleExpr is called when entering the simpleExpr production.
	EnterSimpleExpr(c *SimpleExprContext)

	// EnterColumnRef is called when entering the columnRef production.
	EnterColumnRef(c *ColumnRefContext)

	// EnterColumnRefList is called when entering the columnRefList production.
	EnterColumnRefList(c *ColumnRefListContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterAggregationFunction is called when entering the aggregationFunction production.
	EnterAggregationFunction(c *AggregationFunctionContext)

	// EnterAggregationFunctionName is called when entering the aggregationFunctionName production.
	EnterAggregationFunctionName(c *AggregationFunctionNameContext)

	// EnterDistinct is called when entering the distinct production.
	EnterDistinct(c *DistinctContext)

	// EnterOverClause is called when entering the overClause production.
	EnterOverClause(c *OverClauseContext)

	// EnterWindowSpecification is called when entering the windowSpecification production.
	EnterWindowSpecification(c *WindowSpecificationContext)

	// EnterFrameClause is called when entering the frameClause production.
	EnterFrameClause(c *FrameClauseContext)

	// EnterFrameStart is called when entering the frameStart production.
	EnterFrameStart(c *FrameStartContext)

	// EnterFrameEnd is called when entering the frameEnd production.
	EnterFrameEnd(c *FrameEndContext)

	// EnterFrameBetween is called when entering the frameBetween production.
	EnterFrameBetween(c *FrameBetweenContext)

	// EnterSpecialFunction is called when entering the specialFunction production.
	EnterSpecialFunction(c *SpecialFunctionContext)

	// EnterCurrentUserFunction is called when entering the currentUserFunction production.
	EnterCurrentUserFunction(c *CurrentUserFunctionContext)

	// EnterGroupConcatFunction is called when entering the groupConcatFunction production.
	EnterGroupConcatFunction(c *GroupConcatFunctionContext)

	// EnterWindowFunction is called when entering the windowFunction production.
	EnterWindowFunction(c *WindowFunctionContext)

	// EnterWindowingClause is called when entering the windowingClause production.
	EnterWindowingClause(c *WindowingClauseContext)

	// EnterLeadLagInfo is called when entering the leadLagInfo production.
	EnterLeadLagInfo(c *LeadLagInfoContext)

	// EnterNullTreatment is called when entering the nullTreatment production.
	EnterNullTreatment(c *NullTreatmentContext)

	// EnterCheckType is called when entering the checkType production.
	EnterCheckType(c *CheckTypeContext)

	// EnterRepairType is called when entering the repairType production.
	EnterRepairType(c *RepairTypeContext)

	// EnterCastFunction is called when entering the castFunction production.
	EnterCastFunction(c *CastFunctionContext)

	// EnterConvertFunction is called when entering the convertFunction production.
	EnterConvertFunction(c *ConvertFunctionContext)

	// EnterCastType is called when entering the castType production.
	EnterCastType(c *CastTypeContext)

	// EnterNchar is called when entering the nchar production.
	EnterNchar(c *NcharContext)

	// EnterPositionFunction is called when entering the positionFunction production.
	EnterPositionFunction(c *PositionFunctionContext)

	// EnterSubstringFunction is called when entering the substringFunction production.
	EnterSubstringFunction(c *SubstringFunctionContext)

	// EnterExtractFunction is called when entering the extractFunction production.
	EnterExtractFunction(c *ExtractFunctionContext)

	// EnterCharFunction is called when entering the charFunction production.
	EnterCharFunction(c *CharFunctionContext)

	// EnterTrimFunction is called when entering the trimFunction production.
	EnterTrimFunction(c *TrimFunctionContext)

	// EnterValuesFunction is called when entering the valuesFunction production.
	EnterValuesFunction(c *ValuesFunctionContext)

	// EnterWeightStringFunction is called when entering the weightStringFunction production.
	EnterWeightStringFunction(c *WeightStringFunctionContext)

	// EnterLevelClause is called when entering the levelClause production.
	EnterLevelClause(c *LevelClauseContext)

	// EnterLevelInWeightListElement is called when entering the levelInWeightListElement production.
	EnterLevelInWeightListElement(c *LevelInWeightListElementContext)

	// EnterRegularFunction is called when entering the regularFunction production.
	EnterRegularFunction(c *RegularFunctionContext)

	// EnterShorthandRegularFunction is called when entering the shorthandRegularFunction production.
	EnterShorthandRegularFunction(c *ShorthandRegularFunctionContext)

	// EnterCompleteRegularFunction is called when entering the completeRegularFunction production.
	EnterCompleteRegularFunction(c *CompleteRegularFunctionContext)

	// EnterRegularFunctionName is called when entering the regularFunctionName production.
	EnterRegularFunctionName(c *RegularFunctionNameContext)

	// EnterMatchExpression is called when entering the matchExpression production.
	EnterMatchExpression(c *MatchExpressionContext)

	// EnterMatchSearchModifier is called when entering the matchSearchModifier production.
	EnterMatchSearchModifier(c *MatchSearchModifierContext)

	// EnterCaseExpression is called when entering the caseExpression production.
	EnterCaseExpression(c *CaseExpressionContext)

	// EnterDatetimeExpr is called when entering the datetimeExpr production.
	EnterDatetimeExpr(c *DatetimeExprContext)

	// EnterBinaryLogFileIndexNumber is called when entering the binaryLogFileIndexNumber production.
	EnterBinaryLogFileIndexNumber(c *BinaryLogFileIndexNumberContext)

	// EnterCaseWhen is called when entering the caseWhen production.
	EnterCaseWhen(c *CaseWhenContext)

	// EnterCaseElse is called when entering the caseElse production.
	EnterCaseElse(c *CaseElseContext)

	// EnterIntervalExpression is called when entering the intervalExpression production.
	EnterIntervalExpression(c *IntervalExpressionContext)

	// EnterIntervalValue is called when entering the intervalValue production.
	EnterIntervalValue(c *IntervalValueContext)

	// EnterIntervalUnit is called when entering the intervalUnit production.
	EnterIntervalUnit(c *IntervalUnitContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderByItem is called when entering the orderByItem production.
	EnterOrderByItem(c *OrderByItemContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterStringList is called when entering the stringList production.
	EnterStringList(c *StringListContext)

	// EnterTextString is called when entering the textString production.
	EnterTextString(c *TextStringContext)

	// EnterTextStringHash is called when entering the textStringHash production.
	EnterTextStringHash(c *TextStringHashContext)

	// EnterFieldOptions is called when entering the fieldOptions production.
	EnterFieldOptions(c *FieldOptionsContext)

	// EnterPrecision is called when entering the precision production.
	EnterPrecision(c *PrecisionContext)

	// EnterTypeDatetimePrecision is called when entering the typeDatetimePrecision production.
	EnterTypeDatetimePrecision(c *TypeDatetimePrecisionContext)

	// EnterCharsetWithOptBinary is called when entering the charsetWithOptBinary production.
	EnterCharsetWithOptBinary(c *CharsetWithOptBinaryContext)

	// EnterAscii is called when entering the ascii production.
	EnterAscii(c *AsciiContext)

	// EnterUnicode is called when entering the unicode production.
	EnterUnicode(c *UnicodeContext)

	// EnterCharset is called when entering the charset production.
	EnterCharset(c *CharsetContext)

	// EnterDefaultCollation is called when entering the defaultCollation production.
	EnterDefaultCollation(c *DefaultCollationContext)

	// EnterDefaultEncryption is called when entering the defaultEncryption production.
	EnterDefaultEncryption(c *DefaultEncryptionContext)

	// EnterDefaultCharset is called when entering the defaultCharset production.
	EnterDefaultCharset(c *DefaultCharsetContext)

	// EnterSignedLiteral is called when entering the signedLiteral production.
	EnterSignedLiteral(c *SignedLiteralContext)

	// EnterNow is called when entering the now production.
	EnterNow(c *NowContext)

	// EnterColumnFormat is called when entering the columnFormat production.
	EnterColumnFormat(c *ColumnFormatContext)

	// EnterStorageMedia is called when entering the storageMedia production.
	EnterStorageMedia(c *StorageMediaContext)

	// EnterDirection is called when entering the direction production.
	EnterDirection(c *DirectionContext)

	// EnterKeyOrIndex is called when entering the keyOrIndex production.
	EnterKeyOrIndex(c *KeyOrIndexContext)

	// EnterFieldLength is called when entering the fieldLength production.
	EnterFieldLength(c *FieldLengthContext)

	// EnterCharacterSet is called when entering the characterSet production.
	EnterCharacterSet(c *CharacterSetContext)

	// EnterCollateClause is called when entering the collateClause production.
	EnterCollateClause(c *CollateClauseContext)

	// EnterFieldOrVarSpec is called when entering the fieldOrVarSpec production.
	EnterFieldOrVarSpec(c *FieldOrVarSpecContext)

	// EnterNotExistClause is called when entering the notExistClause production.
	EnterNotExistClause(c *NotExistClauseContext)

	// EnterExistClause is called when entering the existClause production.
	EnterExistClause(c *ExistClauseContext)

	// EnterConnectionId is called when entering the connectionId production.
	EnterConnectionId(c *ConnectionIdContext)

	// EnterLabelName is called when entering the labelName production.
	EnterLabelName(c *LabelNameContext)

	// EnterCursorName is called when entering the cursorName production.
	EnterCursorName(c *CursorNameContext)

	// EnterConditionName is called when entering the conditionName production.
	EnterConditionName(c *ConditionNameContext)

	// EnterUnionOption is called when entering the unionOption production.
	EnterUnionOption(c *UnionOptionContext)

	// EnterNoWriteToBinLog is called when entering the noWriteToBinLog production.
	EnterNoWriteToBinLog(c *NoWriteToBinLogContext)

	// EnterChannelOption is called when entering the channelOption production.
	EnterChannelOption(c *ChannelOptionContext)

	// EnterPreparedStatement is called when entering the preparedStatement production.
	EnterPreparedStatement(c *PreparedStatementContext)

	// EnterExecuteStatement is called when entering the executeStatement production.
	EnterExecuteStatement(c *ExecuteStatementContext)

	// EnterExecuteVarList is called when entering the executeVarList production.
	EnterExecuteVarList(c *ExecuteVarListContext)

	// EnterAlterStatement is called when entering the alterStatement production.
	EnterAlterStatement(c *AlterStatementContext)

	// EnterCreateTable is called when entering the createTable production.
	EnterCreateTable(c *CreateTableContext)

	// EnterPartitionClause is called when entering the partitionClause production.
	EnterPartitionClause(c *PartitionClauseContext)

	// EnterPartitionTypeDef is called when entering the partitionTypeDef production.
	EnterPartitionTypeDef(c *PartitionTypeDefContext)

	// EnterSubPartitions is called when entering the subPartitions production.
	EnterSubPartitions(c *SubPartitionsContext)

	// EnterPartitionKeyAlgorithm is called when entering the partitionKeyAlgorithm production.
	EnterPartitionKeyAlgorithm(c *PartitionKeyAlgorithmContext)

	// EnterDuplicateAsQueryExpression is called when entering the duplicateAsQueryExpression production.
	EnterDuplicateAsQueryExpression(c *DuplicateAsQueryExpressionContext)

	// EnterAlterTable is called when entering the alterTable production.
	EnterAlterTable(c *AlterTableContext)

	// EnterStandaloneAlterTableAction is called when entering the standaloneAlterTableAction production.
	EnterStandaloneAlterTableAction(c *StandaloneAlterTableActionContext)

	// EnterAlterTableActions is called when entering the alterTableActions production.
	EnterAlterTableActions(c *AlterTableActionsContext)

	// EnterAlterTablePartitionOptions is called when entering the alterTablePartitionOptions production.
	EnterAlterTablePartitionOptions(c *AlterTablePartitionOptionsContext)

	// EnterAlterCommandList is called when entering the alterCommandList production.
	EnterAlterCommandList(c *AlterCommandListContext)

	// EnterAlterList is called when entering the alterList production.
	EnterAlterList(c *AlterListContext)

	// EnterCreateTableOptionsSpaceSeparated is called when entering the createTableOptionsSpaceSeparated production.
	EnterCreateTableOptionsSpaceSeparated(c *CreateTableOptionsSpaceSeparatedContext)

	// EnterAddColumn is called when entering the addColumn production.
	EnterAddColumn(c *AddColumnContext)

	// EnterAddTableConstraint is called when entering the addTableConstraint production.
	EnterAddTableConstraint(c *AddTableConstraintContext)

	// EnterChangeColumn is called when entering the changeColumn production.
	EnterChangeColumn(c *ChangeColumnContext)

	// EnterModifyColumn is called when entering the modifyColumn production.
	EnterModifyColumn(c *ModifyColumnContext)

	// EnterAlterTableDrop is called when entering the alterTableDrop production.
	EnterAlterTableDrop(c *AlterTableDropContext)

	// EnterDisableKeys is called when entering the disableKeys production.
	EnterDisableKeys(c *DisableKeysContext)

	// EnterEnableKeys is called when entering the enableKeys production.
	EnterEnableKeys(c *EnableKeysContext)

	// EnterAlterColumn is called when entering the alterColumn production.
	EnterAlterColumn(c *AlterColumnContext)

	// EnterAlterIndex is called when entering the alterIndex production.
	EnterAlterIndex(c *AlterIndexContext)

	// EnterAlterCheck is called when entering the alterCheck production.
	EnterAlterCheck(c *AlterCheckContext)

	// EnterAlterConstraint is called when entering the alterConstraint production.
	EnterAlterConstraint(c *AlterConstraintContext)

	// EnterRenameColumn is called when entering the renameColumn production.
	EnterRenameColumn(c *RenameColumnContext)

	// EnterAlterRenameTable is called when entering the alterRenameTable production.
	EnterAlterRenameTable(c *AlterRenameTableContext)

	// EnterRenameIndex is called when entering the renameIndex production.
	EnterRenameIndex(c *RenameIndexContext)

	// EnterAlterConvert is called when entering the alterConvert production.
	EnterAlterConvert(c *AlterConvertContext)

	// EnterAlterTableForce is called when entering the alterTableForce production.
	EnterAlterTableForce(c *AlterTableForceContext)

	// EnterAlterTableOrder is called when entering the alterTableOrder production.
	EnterAlterTableOrder(c *AlterTableOrderContext)

	// EnterAlterOrderList is called when entering the alterOrderList production.
	EnterAlterOrderList(c *AlterOrderListContext)

	// EnterTableConstraintDef is called when entering the tableConstraintDef production.
	EnterTableConstraintDef(c *TableConstraintDefContext)

	// EnterAlterCommandsModifierList is called when entering the alterCommandsModifierList production.
	EnterAlterCommandsModifierList(c *AlterCommandsModifierListContext)

	// EnterAlterCommandsModifier is called when entering the alterCommandsModifier production.
	EnterAlterCommandsModifier(c *AlterCommandsModifierContext)

	// EnterWithValidation is called when entering the withValidation production.
	EnterWithValidation(c *WithValidationContext)

	// EnterStandaloneAlterCommands is called when entering the standaloneAlterCommands production.
	EnterStandaloneAlterCommands(c *StandaloneAlterCommandsContext)

	// EnterAlterPartition is called when entering the alterPartition production.
	EnterAlterPartition(c *AlterPartitionContext)

	// EnterConstraintClause is called when entering the constraintClause production.
	EnterConstraintClause(c *ConstraintClauseContext)

	// EnterTableElementList is called when entering the tableElementList production.
	EnterTableElementList(c *TableElementListContext)

	// EnterTableElement is called when entering the tableElement production.
	EnterTableElement(c *TableElementContext)

	// EnterRestrict is called when entering the restrict production.
	EnterRestrict(c *RestrictContext)

	// EnterFulltextIndexOption is called when entering the fulltextIndexOption production.
	EnterFulltextIndexOption(c *FulltextIndexOptionContext)

	// EnterDropTable is called when entering the dropTable production.
	EnterDropTable(c *DropTableContext)

	// EnterDropIndex is called when entering the dropIndex production.
	EnterDropIndex(c *DropIndexContext)

	// EnterAlterAlgorithmOption is called when entering the alterAlgorithmOption production.
	EnterAlterAlgorithmOption(c *AlterAlgorithmOptionContext)

	// EnterAlterLockOption is called when entering the alterLockOption production.
	EnterAlterLockOption(c *AlterLockOptionContext)

	// EnterTruncateTable is called when entering the truncateTable production.
	EnterTruncateTable(c *TruncateTableContext)

	// EnterCreateIndex is called when entering the createIndex production.
	EnterCreateIndex(c *CreateIndexContext)

	// EnterCreateDatabase is called when entering the createDatabase production.
	EnterCreateDatabase(c *CreateDatabaseContext)

	// EnterAlterDatabase is called when entering the alterDatabase production.
	EnterAlterDatabase(c *AlterDatabaseContext)

	// EnterCreateDatabaseSpecification_ is called when entering the createDatabaseSpecification_ production.
	EnterCreateDatabaseSpecification_(c *CreateDatabaseSpecification_Context)

	// EnterAlterDatabaseSpecification_ is called when entering the alterDatabaseSpecification_ production.
	EnterAlterDatabaseSpecification_(c *AlterDatabaseSpecification_Context)

	// EnterDropDatabase is called when entering the dropDatabase production.
	EnterDropDatabase(c *DropDatabaseContext)

	// EnterAlterInstance is called when entering the alterInstance production.
	EnterAlterInstance(c *AlterInstanceContext)

	// EnterInstanceAction is called when entering the instanceAction production.
	EnterInstanceAction(c *InstanceActionContext)

	// EnterChannel is called when entering the channel production.
	EnterChannel(c *ChannelContext)

	// EnterCreateEvent is called when entering the createEvent production.
	EnterCreateEvent(c *CreateEventContext)

	// EnterAlterEvent is called when entering the alterEvent production.
	EnterAlterEvent(c *AlterEventContext)

	// EnterDropEvent is called when entering the dropEvent production.
	EnterDropEvent(c *DropEventContext)

	// EnterCreateFunction is called when entering the createFunction production.
	EnterCreateFunction(c *CreateFunctionContext)

	// EnterAlterFunction is called when entering the alterFunction production.
	EnterAlterFunction(c *AlterFunctionContext)

	// EnterDropFunction is called when entering the dropFunction production.
	EnterDropFunction(c *DropFunctionContext)

	// EnterCreateProcedure is called when entering the createProcedure production.
	EnterCreateProcedure(c *CreateProcedureContext)

	// EnterAlterProcedure is called when entering the alterProcedure production.
	EnterAlterProcedure(c *AlterProcedureContext)

	// EnterDropProcedure is called when entering the dropProcedure production.
	EnterDropProcedure(c *DropProcedureContext)

	// EnterCreateServer is called when entering the createServer production.
	EnterCreateServer(c *CreateServerContext)

	// EnterAlterServer is called when entering the alterServer production.
	EnterAlterServer(c *AlterServerContext)

	// EnterDropServer is called when entering the dropServer production.
	EnterDropServer(c *DropServerContext)

	// EnterCreateView is called when entering the createView production.
	EnterCreateView(c *CreateViewContext)

	// EnterAlterView is called when entering the alterView production.
	EnterAlterView(c *AlterViewContext)

	// EnterDropView is called when entering the dropView production.
	EnterDropView(c *DropViewContext)

	// EnterCreateTablespace is called when entering the createTablespace production.
	EnterCreateTablespace(c *CreateTablespaceContext)

	// EnterCreateTablespaceInnodb is called when entering the createTablespaceInnodb production.
	EnterCreateTablespaceInnodb(c *CreateTablespaceInnodbContext)

	// EnterCreateTablespaceNdb is called when entering the createTablespaceNdb production.
	EnterCreateTablespaceNdb(c *CreateTablespaceNdbContext)

	// EnterAlterTablespace is called when entering the alterTablespace production.
	EnterAlterTablespace(c *AlterTablespaceContext)

	// EnterAlterTablespaceNdb is called when entering the alterTablespaceNdb production.
	EnterAlterTablespaceNdb(c *AlterTablespaceNdbContext)

	// EnterAlterTablespaceInnodb is called when entering the alterTablespaceInnodb production.
	EnterAlterTablespaceInnodb(c *AlterTablespaceInnodbContext)

	// EnterDropTablespace is called when entering the dropTablespace production.
	EnterDropTablespace(c *DropTablespaceContext)

	// EnterCreateLogfileGroup is called when entering the createLogfileGroup production.
	EnterCreateLogfileGroup(c *CreateLogfileGroupContext)

	// EnterAlterLogfileGroup is called when entering the alterLogfileGroup production.
	EnterAlterLogfileGroup(c *AlterLogfileGroupContext)

	// EnterDropLogfileGroup is called when entering the dropLogfileGroup production.
	EnterDropLogfileGroup(c *DropLogfileGroupContext)

	// EnterCreateTrigger is called when entering the createTrigger production.
	EnterCreateTrigger(c *CreateTriggerContext)

	// EnterDropTrigger is called when entering the dropTrigger production.
	EnterDropTrigger(c *DropTriggerContext)

	// EnterRenameTable is called when entering the renameTable production.
	EnterRenameTable(c *RenameTableContext)

	// EnterCreateDefinitionClause is called when entering the createDefinitionClause production.
	EnterCreateDefinitionClause(c *CreateDefinitionClauseContext)

	// EnterColumnDefinition is called when entering the columnDefinition production.
	EnterColumnDefinition(c *ColumnDefinitionContext)

	// EnterFieldDefinition is called when entering the fieldDefinition production.
	EnterFieldDefinition(c *FieldDefinitionContext)

	// EnterColumnAttribute is called when entering the columnAttribute production.
	EnterColumnAttribute(c *ColumnAttributeContext)

	// EnterCheckConstraint is called when entering the checkConstraint production.
	EnterCheckConstraint(c *CheckConstraintContext)

	// EnterConstraintEnforcement is called when entering the constraintEnforcement production.
	EnterConstraintEnforcement(c *ConstraintEnforcementContext)

	// EnterGeneratedOption is called when entering the generatedOption production.
	EnterGeneratedOption(c *GeneratedOptionContext)

	// EnterReferenceDefinition is called when entering the referenceDefinition production.
	EnterReferenceDefinition(c *ReferenceDefinitionContext)

	// EnterOnUpdateDelete is called when entering the onUpdateDelete production.
	EnterOnUpdateDelete(c *OnUpdateDeleteContext)

	// EnterReferenceOption is called when entering the referenceOption production.
	EnterReferenceOption(c *ReferenceOptionContext)

	// EnterIndexNameAndType is called when entering the indexNameAndType production.
	EnterIndexNameAndType(c *IndexNameAndTypeContext)

	// EnterIndexType is called when entering the indexType production.
	EnterIndexType(c *IndexTypeContext)

	// EnterIndexTypeClause is called when entering the indexTypeClause production.
	EnterIndexTypeClause(c *IndexTypeClauseContext)

	// EnterKeyParts is called when entering the keyParts production.
	EnterKeyParts(c *KeyPartsContext)

	// EnterKeyPart is called when entering the keyPart production.
	EnterKeyPart(c *KeyPartContext)

	// EnterKeyPartWithExpression is called when entering the keyPartWithExpression production.
	EnterKeyPartWithExpression(c *KeyPartWithExpressionContext)

	// EnterKeyListWithExpression is called when entering the keyListWithExpression production.
	EnterKeyListWithExpression(c *KeyListWithExpressionContext)

	// EnterIndexOption is called when entering the indexOption production.
	EnterIndexOption(c *IndexOptionContext)

	// EnterCommonIndexOption is called when entering the commonIndexOption production.
	EnterCommonIndexOption(c *CommonIndexOptionContext)

	// EnterVisibility is called when entering the visibility production.
	EnterVisibility(c *VisibilityContext)

	// EnterCreateLikeClause is called when entering the createLikeClause production.
	EnterCreateLikeClause(c *CreateLikeClauseContext)

	// EnterCreateIndexSpecification is called when entering the createIndexSpecification production.
	EnterCreateIndexSpecification(c *CreateIndexSpecificationContext)

	// EnterCreateTableOptions is called when entering the createTableOptions production.
	EnterCreateTableOptions(c *CreateTableOptionsContext)

	// EnterCreateTableOption is called when entering the createTableOption production.
	EnterCreateTableOption(c *CreateTableOptionContext)

	// EnterCreateSRSStatement is called when entering the createSRSStatement production.
	EnterCreateSRSStatement(c *CreateSRSStatementContext)

	// EnterDropSRSStatement is called when entering the dropSRSStatement production.
	EnterDropSRSStatement(c *DropSRSStatementContext)

	// EnterSrsAttribute is called when entering the srsAttribute production.
	EnterSrsAttribute(c *SrsAttributeContext)

	// EnterPlace is called when entering the place production.
	EnterPlace(c *PlaceContext)

	// EnterPartitionDefinitions is called when entering the partitionDefinitions production.
	EnterPartitionDefinitions(c *PartitionDefinitionsContext)

	// EnterPartitionDefinition is called when entering the partitionDefinition production.
	EnterPartitionDefinition(c *PartitionDefinitionContext)

	// EnterPartitionLessThanValue is called when entering the partitionLessThanValue production.
	EnterPartitionLessThanValue(c *PartitionLessThanValueContext)

	// EnterPartitionValueList is called when entering the partitionValueList production.
	EnterPartitionValueList(c *PartitionValueListContext)

	// EnterPartitionDefinitionOption is called when entering the partitionDefinitionOption production.
	EnterPartitionDefinitionOption(c *PartitionDefinitionOptionContext)

	// EnterSubpartitionDefinition is called when entering the subpartitionDefinition production.
	EnterSubpartitionDefinition(c *SubpartitionDefinitionContext)

	// EnterOwnerStatement is called when entering the ownerStatement production.
	EnterOwnerStatement(c *OwnerStatementContext)

	// EnterScheduleExpression is called when entering the scheduleExpression production.
	EnterScheduleExpression(c *ScheduleExpressionContext)

	// EnterTimestampValue is called when entering the timestampValue production.
	EnterTimestampValue(c *TimestampValueContext)

	// EnterRoutineBody is called when entering the routineBody production.
	EnterRoutineBody(c *RoutineBodyContext)

	// EnterServerOption is called when entering the serverOption production.
	EnterServerOption(c *ServerOptionContext)

	// EnterRoutineOption is called when entering the routineOption production.
	EnterRoutineOption(c *RoutineOptionContext)

	// EnterProcedureParameter is called when entering the procedureParameter production.
	EnterProcedureParameter(c *ProcedureParameterContext)

	// EnterFileSizeLiteral is called when entering the fileSizeLiteral production.
	EnterFileSizeLiteral(c *FileSizeLiteralContext)

	// EnterSimpleStatement is called when entering the simpleStatement production.
	EnterSimpleStatement(c *SimpleStatementContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterValidStatement is called when entering the validStatement production.
	EnterValidStatement(c *ValidStatementContext)

	// EnterBeginStatement is called when entering the beginStatement production.
	EnterBeginStatement(c *BeginStatementContext)

	// EnterDeclareStatement is called when entering the declareStatement production.
	EnterDeclareStatement(c *DeclareStatementContext)

	// EnterFlowControlStatement is called when entering the flowControlStatement production.
	EnterFlowControlStatement(c *FlowControlStatementContext)

	// EnterCaseStatement is called when entering the caseStatement production.
	EnterCaseStatement(c *CaseStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterIterateStatement is called when entering the iterateStatement production.
	EnterIterateStatement(c *IterateStatementContext)

	// EnterLeaveStatement is called when entering the leaveStatement production.
	EnterLeaveStatement(c *LeaveStatementContext)

	// EnterLoopStatement is called when entering the loopStatement production.
	EnterLoopStatement(c *LoopStatementContext)

	// EnterRepeatStatement is called when entering the repeatStatement production.
	EnterRepeatStatement(c *RepeatStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterCursorStatement is called when entering the cursorStatement production.
	EnterCursorStatement(c *CursorStatementContext)

	// EnterCursorCloseStatement is called when entering the cursorCloseStatement production.
	EnterCursorCloseStatement(c *CursorCloseStatementContext)

	// EnterCursorDeclareStatement is called when entering the cursorDeclareStatement production.
	EnterCursorDeclareStatement(c *CursorDeclareStatementContext)

	// EnterCursorFetchStatement is called when entering the cursorFetchStatement production.
	EnterCursorFetchStatement(c *CursorFetchStatementContext)

	// EnterCursorOpenStatement is called when entering the cursorOpenStatement production.
	EnterCursorOpenStatement(c *CursorOpenStatementContext)

	// EnterConditionHandlingStatement is called when entering the conditionHandlingStatement production.
	EnterConditionHandlingStatement(c *ConditionHandlingStatementContext)

	// EnterDeclareConditionStatement is called when entering the declareConditionStatement production.
	EnterDeclareConditionStatement(c *DeclareConditionStatementContext)

	// EnterDeclareHandlerStatement is called when entering the declareHandlerStatement production.
	EnterDeclareHandlerStatement(c *DeclareHandlerStatementContext)

	// EnterGetDiagnosticsStatement is called when entering the getDiagnosticsStatement production.
	EnterGetDiagnosticsStatement(c *GetDiagnosticsStatementContext)

	// EnterStatementInformationItem is called when entering the statementInformationItem production.
	EnterStatementInformationItem(c *StatementInformationItemContext)

	// EnterConditionInformationItem is called when entering the conditionInformationItem production.
	EnterConditionInformationItem(c *ConditionInformationItemContext)

	// EnterConditionNumber is called when entering the conditionNumber production.
	EnterConditionNumber(c *ConditionNumberContext)

	// EnterStatementInformationItemName is called when entering the statementInformationItemName production.
	EnterStatementInformationItemName(c *StatementInformationItemNameContext)

	// EnterConditionInformationItemName is called when entering the conditionInformationItemName production.
	EnterConditionInformationItemName(c *ConditionInformationItemNameContext)

	// EnterHandlerAction is called when entering the handlerAction production.
	EnterHandlerAction(c *HandlerActionContext)

	// EnterConditionValue is called when entering the conditionValue production.
	EnterConditionValue(c *ConditionValueContext)

	// EnterResignalStatement is called when entering the resignalStatement production.
	EnterResignalStatement(c *ResignalStatementContext)

	// EnterSignalStatement is called when entering the signalStatement production.
	EnterSignalStatement(c *SignalStatementContext)

	// EnterSignalInformationItem is called when entering the signalInformationItem production.
	EnterSignalInformationItem(c *SignalInformationItemContext)

	// EnterUse is called when entering the use production.
	EnterUse(c *UseContext)

	// EnterHelp is called when entering the help production.
	EnterHelp(c *HelpContext)

	// EnterExplain is called when entering the explain production.
	EnterExplain(c *ExplainContext)

	// EnterShowDatabases is called when entering the showDatabases production.
	EnterShowDatabases(c *ShowDatabasesContext)

	// EnterShowTables is called when entering the showTables production.
	EnterShowTables(c *ShowTablesContext)

	// EnterShowTableStatus is called when entering the showTableStatus production.
	EnterShowTableStatus(c *ShowTableStatusContext)

	// EnterShowColumns is called when entering the showColumns production.
	EnterShowColumns(c *ShowColumnsContext)

	// EnterShowIndex is called when entering the showIndex production.
	EnterShowIndex(c *ShowIndexContext)

	// EnterShowCreateTable is called when entering the showCreateTable production.
	EnterShowCreateTable(c *ShowCreateTableContext)

	// EnterFromSchema is called when entering the fromSchema production.
	EnterFromSchema(c *FromSchemaContext)

	// EnterFromTable is called when entering the fromTable production.
	EnterFromTable(c *FromTableContext)

	// EnterShowLike is called when entering the showLike production.
	EnterShowLike(c *ShowLikeContext)

	// EnterShowWhereClause is called when entering the showWhereClause production.
	EnterShowWhereClause(c *ShowWhereClauseContext)

	// EnterShowFilter is called when entering the showFilter production.
	EnterShowFilter(c *ShowFilterContext)

	// EnterShowProfileType is called when entering the showProfileType production.
	EnterShowProfileType(c *ShowProfileTypeContext)

	// EnterSetVariable is called when entering the setVariable production.
	EnterSetVariable(c *SetVariableContext)

	// EnterOptionValueList is called when entering the optionValueList production.
	EnterOptionValueList(c *OptionValueListContext)

	// EnterOptionValueNoOptionType is called when entering the optionValueNoOptionType production.
	EnterOptionValueNoOptionType(c *OptionValueNoOptionTypeContext)

	// EnterOptionValue is called when entering the optionValue production.
	EnterOptionValue(c *OptionValueContext)

	// EnterShowBinaryLogs is called when entering the showBinaryLogs production.
	EnterShowBinaryLogs(c *ShowBinaryLogsContext)

	// EnterShowBinlogEvents is called when entering the showBinlogEvents production.
	EnterShowBinlogEvents(c *ShowBinlogEventsContext)

	// EnterShowCharacterSet is called when entering the showCharacterSet production.
	EnterShowCharacterSet(c *ShowCharacterSetContext)

	// EnterShowCollation is called when entering the showCollation production.
	EnterShowCollation(c *ShowCollationContext)

	// EnterShowCreateDatabase is called when entering the showCreateDatabase production.
	EnterShowCreateDatabase(c *ShowCreateDatabaseContext)

	// EnterShowCreateEvent is called when entering the showCreateEvent production.
	EnterShowCreateEvent(c *ShowCreateEventContext)

	// EnterShowCreateFunction is called when entering the showCreateFunction production.
	EnterShowCreateFunction(c *ShowCreateFunctionContext)

	// EnterShowCreateProcedure is called when entering the showCreateProcedure production.
	EnterShowCreateProcedure(c *ShowCreateProcedureContext)

	// EnterShowCreateTrigger is called when entering the showCreateTrigger production.
	EnterShowCreateTrigger(c *ShowCreateTriggerContext)

	// EnterShowCreateUser is called when entering the showCreateUser production.
	EnterShowCreateUser(c *ShowCreateUserContext)

	// EnterShowCreateView is called when entering the showCreateView production.
	EnterShowCreateView(c *ShowCreateViewContext)

	// EnterShowEngine is called when entering the showEngine production.
	EnterShowEngine(c *ShowEngineContext)

	// EnterShowEngines is called when entering the showEngines production.
	EnterShowEngines(c *ShowEnginesContext)

	// EnterShowCharset is called when entering the showCharset production.
	EnterShowCharset(c *ShowCharsetContext)

	// EnterShowErrors is called when entering the showErrors production.
	EnterShowErrors(c *ShowErrorsContext)

	// EnterShowEvents is called when entering the showEvents production.
	EnterShowEvents(c *ShowEventsContext)

	// EnterShowFunctionCode is called when entering the showFunctionCode production.
	EnterShowFunctionCode(c *ShowFunctionCodeContext)

	// EnterShowFunctionStatus is called when entering the showFunctionStatus production.
	EnterShowFunctionStatus(c *ShowFunctionStatusContext)

	// EnterShowGrant is called when entering the showGrant production.
	EnterShowGrant(c *ShowGrantContext)

	// EnterShowMasterStatus is called when entering the showMasterStatus production.
	EnterShowMasterStatus(c *ShowMasterStatusContext)

	// EnterShowOpenTables is called when entering the showOpenTables production.
	EnterShowOpenTables(c *ShowOpenTablesContext)

	// EnterShowPlugins is called when entering the showPlugins production.
	EnterShowPlugins(c *ShowPluginsContext)

	// EnterShowPrivileges is called when entering the showPrivileges production.
	EnterShowPrivileges(c *ShowPrivilegesContext)

	// EnterShowProcedureCode is called when entering the showProcedureCode production.
	EnterShowProcedureCode(c *ShowProcedureCodeContext)

	// EnterShowProcedureStatus is called when entering the showProcedureStatus production.
	EnterShowProcedureStatus(c *ShowProcedureStatusContext)

	// EnterShowProcesslist is called when entering the showProcesslist production.
	EnterShowProcesslist(c *ShowProcesslistContext)

	// EnterShowProfile is called when entering the showProfile production.
	EnterShowProfile(c *ShowProfileContext)

	// EnterShowProfiles is called when entering the showProfiles production.
	EnterShowProfiles(c *ShowProfilesContext)

	// EnterShowRelaylogEvent is called when entering the showRelaylogEvent production.
	EnterShowRelaylogEvent(c *ShowRelaylogEventContext)

	// EnterShowSlavehost is called when entering the showSlavehost production.
	EnterShowSlavehost(c *ShowSlavehostContext)

	// EnterShowSlaveStatus is called when entering the showSlaveStatus production.
	EnterShowSlaveStatus(c *ShowSlaveStatusContext)

	// EnterShowStatus is called when entering the showStatus production.
	EnterShowStatus(c *ShowStatusContext)

	// EnterShowTriggers is called when entering the showTriggers production.
	EnterShowTriggers(c *ShowTriggersContext)

	// EnterShowVariables is called when entering the showVariables production.
	EnterShowVariables(c *ShowVariablesContext)

	// EnterShowWarnings is called when entering the showWarnings production.
	EnterShowWarnings(c *ShowWarningsContext)

	// EnterShowReplicas is called when entering the showReplicas production.
	EnterShowReplicas(c *ShowReplicasContext)

	// EnterShowReplicaStatus is called when entering the showReplicaStatus production.
	EnterShowReplicaStatus(c *ShowReplicaStatusContext)

	// EnterSetCharacter is called when entering the setCharacter production.
	EnterSetCharacter(c *SetCharacterContext)

	// EnterClone is called when entering the clone production.
	EnterClone(c *CloneContext)

	// EnterCloneAction is called when entering the cloneAction production.
	EnterCloneAction(c *CloneActionContext)

	// EnterCreateLoadableFunction is called when entering the createLoadableFunction production.
	EnterCreateLoadableFunction(c *CreateLoadableFunctionContext)

	// EnterInstall is called when entering the install production.
	EnterInstall(c *InstallContext)

	// EnterUninstall is called when entering the uninstall production.
	EnterUninstall(c *UninstallContext)

	// EnterInstallComponent is called when entering the installComponent production.
	EnterInstallComponent(c *InstallComponentContext)

	// EnterInstallPlugin is called when entering the installPlugin production.
	EnterInstallPlugin(c *InstallPluginContext)

	// EnterUninstallComponent is called when entering the uninstallComponent production.
	EnterUninstallComponent(c *UninstallComponentContext)

	// EnterUninstallPlugin is called when entering the uninstallPlugin production.
	EnterUninstallPlugin(c *UninstallPluginContext)

	// EnterAnalyzeTable is called when entering the analyzeTable production.
	EnterAnalyzeTable(c *AnalyzeTableContext)

	// EnterHistogram is called when entering the histogram production.
	EnterHistogram(c *HistogramContext)

	// EnterCheckTable is called when entering the checkTable production.
	EnterCheckTable(c *CheckTableContext)

	// EnterCheckTableOption is called when entering the checkTableOption production.
	EnterCheckTableOption(c *CheckTableOptionContext)

	// EnterChecksumTable is called when entering the checksumTable production.
	EnterChecksumTable(c *ChecksumTableContext)

	// EnterOptimizeTable is called when entering the optimizeTable production.
	EnterOptimizeTable(c *OptimizeTableContext)

	// EnterRepairTable is called when entering the repairTable production.
	EnterRepairTable(c *RepairTableContext)

	// EnterAlterResourceGroup is called when entering the alterResourceGroup production.
	EnterAlterResourceGroup(c *AlterResourceGroupContext)

	// EnterVcpuSpec is called when entering the vcpuSpec production.
	EnterVcpuSpec(c *VcpuSpecContext)

	// EnterCreateResourceGroup is called when entering the createResourceGroup production.
	EnterCreateResourceGroup(c *CreateResourceGroupContext)

	// EnterDropResourceGroup is called when entering the dropResourceGroup production.
	EnterDropResourceGroup(c *DropResourceGroupContext)

	// EnterSetResourceGroup is called when entering the setResourceGroup production.
	EnterSetResourceGroup(c *SetResourceGroupContext)

	// EnterBinlog is called when entering the binlog production.
	EnterBinlog(c *BinlogContext)

	// EnterCacheIndex is called when entering the cacheIndex production.
	EnterCacheIndex(c *CacheIndexContext)

	// EnterCacheTableIndexList is called when entering the cacheTableIndexList production.
	EnterCacheTableIndexList(c *CacheTableIndexListContext)

	// EnterPartitionList is called when entering the partitionList production.
	EnterPartitionList(c *PartitionListContext)

	// EnterFlush is called when entering the flush production.
	EnterFlush(c *FlushContext)

	// EnterFlushOption is called when entering the flushOption production.
	EnterFlushOption(c *FlushOptionContext)

	// EnterTablesOption is called when entering the tablesOption production.
	EnterTablesOption(c *TablesOptionContext)

	// EnterKill is called when entering the kill production.
	EnterKill(c *KillContext)

	// EnterLoadIndexInfo is called when entering the loadIndexInfo production.
	EnterLoadIndexInfo(c *LoadIndexInfoContext)

	// EnterLoadTableIndexList is called when entering the loadTableIndexList production.
	EnterLoadTableIndexList(c *LoadTableIndexListContext)

	// EnterResetStatement is called when entering the resetStatement production.
	EnterResetStatement(c *ResetStatementContext)

	// EnterResetOption is called when entering the resetOption production.
	EnterResetOption(c *ResetOptionContext)

	// EnterResetPersist is called when entering the resetPersist production.
	EnterResetPersist(c *ResetPersistContext)

	// EnterRestart is called when entering the restart production.
	EnterRestart(c *RestartContext)

	// EnterShutdown is called when entering the shutdown production.
	EnterShutdown(c *ShutdownContext)

	// EnterExplainType is called when entering the explainType production.
	EnterExplainType(c *ExplainTypeContext)

	// EnterExplainableStatement is called when entering the explainableStatement production.
	EnterExplainableStatement(c *ExplainableStatementContext)

	// EnterFormatName is called when entering the formatName production.
	EnterFormatName(c *FormatNameContext)

	// EnterShow is called when entering the show production.
	EnterShow(c *ShowContext)

	// EnterSetTransaction is called when entering the setTransaction production.
	EnterSetTransaction(c *SetTransactionContext)

	// EnterSetAutoCommit is called when entering the setAutoCommit production.
	EnterSetAutoCommit(c *SetAutoCommitContext)

	// EnterBeginTransaction is called when entering the beginTransaction production.
	EnterBeginTransaction(c *BeginTransactionContext)

	// EnterTransactionCharacteristic is called when entering the transactionCharacteristic production.
	EnterTransactionCharacteristic(c *TransactionCharacteristicContext)

	// EnterCommit is called when entering the commit production.
	EnterCommit(c *CommitContext)

	// EnterRollback is called when entering the rollback production.
	EnterRollback(c *RollbackContext)

	// EnterSavepoint is called when entering the savepoint production.
	EnterSavepoint(c *SavepointContext)

	// EnterBegin is called when entering the begin production.
	EnterBegin(c *BeginContext)

	// EnterLock is called when entering the lock production.
	EnterLock(c *LockContext)

	// EnterUnlock is called when entering the unlock production.
	EnterUnlock(c *UnlockContext)

	// EnterReleaseSavepoint is called when entering the releaseSavepoint production.
	EnterReleaseSavepoint(c *ReleaseSavepointContext)

	// EnterXa is called when entering the xa production.
	EnterXa(c *XaContext)

	// EnterOptionChain is called when entering the optionChain production.
	EnterOptionChain(c *OptionChainContext)

	// EnterOptionRelease is called when entering the optionRelease production.
	EnterOptionRelease(c *OptionReleaseContext)

	// EnterTableLock is called when entering the tableLock production.
	EnterTableLock(c *TableLockContext)

	// EnterLockOption is called when entering the lockOption production.
	EnterLockOption(c *LockOptionContext)

	// EnterXid is called when entering the xid production.
	EnterXid(c *XidContext)

	// EnterGrantRoleOrPrivilegeTo is called when entering the grantRoleOrPrivilegeTo production.
	EnterGrantRoleOrPrivilegeTo(c *GrantRoleOrPrivilegeToContext)

	// EnterGrantRoleOrPrivilegeOnTo is called when entering the grantRoleOrPrivilegeOnTo production.
	EnterGrantRoleOrPrivilegeOnTo(c *GrantRoleOrPrivilegeOnToContext)

	// EnterGrantProxy is called when entering the grantProxy production.
	EnterGrantProxy(c *GrantProxyContext)

	// EnterRevokeFrom is called when entering the revokeFrom production.
	EnterRevokeFrom(c *RevokeFromContext)

	// EnterRevokeOnFrom is called when entering the revokeOnFrom production.
	EnterRevokeOnFrom(c *RevokeOnFromContext)

	// EnterUserList is called when entering the userList production.
	EnterUserList(c *UserListContext)

	// EnterRoleOrPrivileges is called when entering the roleOrPrivileges production.
	EnterRoleOrPrivileges(c *RoleOrPrivilegesContext)

	// EnterRoleOrDynamicPrivilege is called when entering the roleOrDynamicPrivilege production.
	EnterRoleOrDynamicPrivilege(c *RoleOrDynamicPrivilegeContext)

	// EnterRoleAtHost is called when entering the roleAtHost production.
	EnterRoleAtHost(c *RoleAtHostContext)

	// EnterStaticPrivilegeSelect is called when entering the staticPrivilegeSelect production.
	EnterStaticPrivilegeSelect(c *StaticPrivilegeSelectContext)

	// EnterStaticPrivilegeInsert is called when entering the staticPrivilegeInsert production.
	EnterStaticPrivilegeInsert(c *StaticPrivilegeInsertContext)

	// EnterStaticPrivilegeUpdate is called when entering the staticPrivilegeUpdate production.
	EnterStaticPrivilegeUpdate(c *StaticPrivilegeUpdateContext)

	// EnterStaticPrivilegeReferences is called when entering the staticPrivilegeReferences production.
	EnterStaticPrivilegeReferences(c *StaticPrivilegeReferencesContext)

	// EnterStaticPrivilegeDelete is called when entering the staticPrivilegeDelete production.
	EnterStaticPrivilegeDelete(c *StaticPrivilegeDeleteContext)

	// EnterStaticPrivilegeUsage is called when entering the staticPrivilegeUsage production.
	EnterStaticPrivilegeUsage(c *StaticPrivilegeUsageContext)

	// EnterStaticPrivilegeIndex is called when entering the staticPrivilegeIndex production.
	EnterStaticPrivilegeIndex(c *StaticPrivilegeIndexContext)

	// EnterStaticPrivilegeAlter is called when entering the staticPrivilegeAlter production.
	EnterStaticPrivilegeAlter(c *StaticPrivilegeAlterContext)

	// EnterStaticPrivilegeCreate is called when entering the staticPrivilegeCreate production.
	EnterStaticPrivilegeCreate(c *StaticPrivilegeCreateContext)

	// EnterStaticPrivilegeDrop is called when entering the staticPrivilegeDrop production.
	EnterStaticPrivilegeDrop(c *StaticPrivilegeDropContext)

	// EnterStaticPrivilegeExecute is called when entering the staticPrivilegeExecute production.
	EnterStaticPrivilegeExecute(c *StaticPrivilegeExecuteContext)

	// EnterStaticPrivilegeReload is called when entering the staticPrivilegeReload production.
	EnterStaticPrivilegeReload(c *StaticPrivilegeReloadContext)

	// EnterStaticPrivilegeShutdown is called when entering the staticPrivilegeShutdown production.
	EnterStaticPrivilegeShutdown(c *StaticPrivilegeShutdownContext)

	// EnterStaticPrivilegeProcess is called when entering the staticPrivilegeProcess production.
	EnterStaticPrivilegeProcess(c *StaticPrivilegeProcessContext)

	// EnterStaticPrivilegeFile is called when entering the staticPrivilegeFile production.
	EnterStaticPrivilegeFile(c *StaticPrivilegeFileContext)

	// EnterStaticPrivilegeGrant is called when entering the staticPrivilegeGrant production.
	EnterStaticPrivilegeGrant(c *StaticPrivilegeGrantContext)

	// EnterStaticPrivilegeShowDatabases is called when entering the staticPrivilegeShowDatabases production.
	EnterStaticPrivilegeShowDatabases(c *StaticPrivilegeShowDatabasesContext)

	// EnterStaticPrivilegeSuper is called when entering the staticPrivilegeSuper production.
	EnterStaticPrivilegeSuper(c *StaticPrivilegeSuperContext)

	// EnterStaticPrivilegeCreateTemporaryTables is called when entering the staticPrivilegeCreateTemporaryTables production.
	EnterStaticPrivilegeCreateTemporaryTables(c *StaticPrivilegeCreateTemporaryTablesContext)

	// EnterStaticPrivilegeLockTables is called when entering the staticPrivilegeLockTables production.
	EnterStaticPrivilegeLockTables(c *StaticPrivilegeLockTablesContext)

	// EnterStaticPrivilegeReplicationSlave is called when entering the staticPrivilegeReplicationSlave production.
	EnterStaticPrivilegeReplicationSlave(c *StaticPrivilegeReplicationSlaveContext)

	// EnterStaticPrivilegeReplicationClient is called when entering the staticPrivilegeReplicationClient production.
	EnterStaticPrivilegeReplicationClient(c *StaticPrivilegeReplicationClientContext)

	// EnterStaticPrivilegeCreateView is called when entering the staticPrivilegeCreateView production.
	EnterStaticPrivilegeCreateView(c *StaticPrivilegeCreateViewContext)

	// EnterStaticPrivilegeShowView is called when entering the staticPrivilegeShowView production.
	EnterStaticPrivilegeShowView(c *StaticPrivilegeShowViewContext)

	// EnterStaticPrivilegeCreateRoutine is called when entering the staticPrivilegeCreateRoutine production.
	EnterStaticPrivilegeCreateRoutine(c *StaticPrivilegeCreateRoutineContext)

	// EnterStaticPrivilegeAlterRoutine is called when entering the staticPrivilegeAlterRoutine production.
	EnterStaticPrivilegeAlterRoutine(c *StaticPrivilegeAlterRoutineContext)

	// EnterStaticPrivilegeCreateUser is called when entering the staticPrivilegeCreateUser production.
	EnterStaticPrivilegeCreateUser(c *StaticPrivilegeCreateUserContext)

	// EnterStaticPrivilegeEvent is called when entering the staticPrivilegeEvent production.
	EnterStaticPrivilegeEvent(c *StaticPrivilegeEventContext)

	// EnterStaticPrivilegeTrigger is called when entering the staticPrivilegeTrigger production.
	EnterStaticPrivilegeTrigger(c *StaticPrivilegeTriggerContext)

	// EnterStaticPrivilegeCreateTablespace is called when entering the staticPrivilegeCreateTablespace production.
	EnterStaticPrivilegeCreateTablespace(c *StaticPrivilegeCreateTablespaceContext)

	// EnterStaticPrivilegeCreateRole is called when entering the staticPrivilegeCreateRole production.
	EnterStaticPrivilegeCreateRole(c *StaticPrivilegeCreateRoleContext)

	// EnterStaticPrivilegeDropRole is called when entering the staticPrivilegeDropRole production.
	EnterStaticPrivilegeDropRole(c *StaticPrivilegeDropRoleContext)

	// EnterAclType is called when entering the aclType production.
	EnterAclType(c *AclTypeContext)

	// EnterGrantLevelGlobal is called when entering the grantLevelGlobal production.
	EnterGrantLevelGlobal(c *GrantLevelGlobalContext)

	// EnterGrantLevelSchemaGlobal is called when entering the grantLevelSchemaGlobal production.
	EnterGrantLevelSchemaGlobal(c *GrantLevelSchemaGlobalContext)

	// EnterGrantLevelTable is called when entering the grantLevelTable production.
	EnterGrantLevelTable(c *GrantLevelTableContext)

	// EnterCreateUser is called when entering the createUser production.
	EnterCreateUser(c *CreateUserContext)

	// EnterCreateUserEntryNoOption is called when entering the createUserEntryNoOption production.
	EnterCreateUserEntryNoOption(c *CreateUserEntryNoOptionContext)

	// EnterCreateUserEntryIdentifiedBy is called when entering the createUserEntryIdentifiedBy production.
	EnterCreateUserEntryIdentifiedBy(c *CreateUserEntryIdentifiedByContext)

	// EnterCreateUserEntryIdentifiedWith is called when entering the createUserEntryIdentifiedWith production.
	EnterCreateUserEntryIdentifiedWith(c *CreateUserEntryIdentifiedWithContext)

	// EnterCreateUserList is called when entering the createUserList production.
	EnterCreateUserList(c *CreateUserListContext)

	// EnterDefaultRoleClause is called when entering the defaultRoleClause production.
	EnterDefaultRoleClause(c *DefaultRoleClauseContext)

	// EnterRequireClause is called when entering the requireClause production.
	EnterRequireClause(c *RequireClauseContext)

	// EnterConnectOptions is called when entering the connectOptions production.
	EnterConnectOptions(c *ConnectOptionsContext)

	// EnterAccountLockPasswordExpireOptions is called when entering the accountLockPasswordExpireOptions production.
	EnterAccountLockPasswordExpireOptions(c *AccountLockPasswordExpireOptionsContext)

	// EnterAccountLockPasswordExpireOption is called when entering the accountLockPasswordExpireOption production.
	EnterAccountLockPasswordExpireOption(c *AccountLockPasswordExpireOptionContext)

	// EnterAlterUser is called when entering the alterUser production.
	EnterAlterUser(c *AlterUserContext)

	// EnterAlterUserEntry is called when entering the alterUserEntry production.
	EnterAlterUserEntry(c *AlterUserEntryContext)

	// EnterAlterUserList is called when entering the alterUserList production.
	EnterAlterUserList(c *AlterUserListContext)

	// EnterDropUser is called when entering the dropUser production.
	EnterDropUser(c *DropUserContext)

	// EnterCreateRole is called when entering the createRole production.
	EnterCreateRole(c *CreateRoleContext)

	// EnterDropRole is called when entering the dropRole production.
	EnterDropRole(c *DropRoleContext)

	// EnterRenameUser is called when entering the renameUser production.
	EnterRenameUser(c *RenameUserContext)

	// EnterSetDefaultRole is called when entering the setDefaultRole production.
	EnterSetDefaultRole(c *SetDefaultRoleContext)

	// EnterSetRole is called when entering the setRole production.
	EnterSetRole(c *SetRoleContext)

	// EnterSetPassword is called when entering the setPassword production.
	EnterSetPassword(c *SetPasswordContext)

	// EnterAuthOption is called when entering the authOption production.
	EnterAuthOption(c *AuthOptionContext)

	// EnterWithGrantOption is called when entering the withGrantOption production.
	EnterWithGrantOption(c *WithGrantOptionContext)

	// EnterUserOrRoles is called when entering the userOrRoles production.
	EnterUserOrRoles(c *UserOrRolesContext)

	// EnterRoles is called when entering the roles production.
	EnterRoles(c *RolesContext)

	// EnterGrantAs is called when entering the grantAs production.
	EnterGrantAs(c *GrantAsContext)

	// EnterWithRoles is called when entering the withRoles production.
	EnterWithRoles(c *WithRolesContext)

	// EnterUserAuthOption is called when entering the userAuthOption production.
	EnterUserAuthOption(c *UserAuthOptionContext)

	// EnterIdentifiedBy is called when entering the identifiedBy production.
	EnterIdentifiedBy(c *IdentifiedByContext)

	// EnterIdentifiedWith is called when entering the identifiedWith production.
	EnterIdentifiedWith(c *IdentifiedWithContext)

	// EnterConnectOption is called when entering the connectOption production.
	EnterConnectOption(c *ConnectOptionContext)

	// EnterTlsOption is called when entering the tlsOption production.
	EnterTlsOption(c *TlsOptionContext)

	// EnterUserFuncAuthOption is called when entering the userFuncAuthOption production.
	EnterUserFuncAuthOption(c *UserFuncAuthOptionContext)

	// EnterChange is called when entering the change production.
	EnterChange(c *ChangeContext)

	// EnterChangeMasterTo is called when entering the changeMasterTo production.
	EnterChangeMasterTo(c *ChangeMasterToContext)

	// EnterChangeReplicationFilter is called when entering the changeReplicationFilter production.
	EnterChangeReplicationFilter(c *ChangeReplicationFilterContext)

	// EnterStartSlave is called when entering the startSlave production.
	EnterStartSlave(c *StartSlaveContext)

	// EnterStopSlave is called when entering the stopSlave production.
	EnterStopSlave(c *StopSlaveContext)

	// EnterGroupReplication is called when entering the groupReplication production.
	EnterGroupReplication(c *GroupReplicationContext)

	// EnterStartGroupReplication is called when entering the startGroupReplication production.
	EnterStartGroupReplication(c *StartGroupReplicationContext)

	// EnterStopGroupReplication is called when entering the stopGroupReplication production.
	EnterStopGroupReplication(c *StopGroupReplicationContext)

	// EnterPurgeBinaryLog is called when entering the purgeBinaryLog production.
	EnterPurgeBinaryLog(c *PurgeBinaryLogContext)

	// EnterThreadTypes is called when entering the threadTypes production.
	EnterThreadTypes(c *ThreadTypesContext)

	// EnterThreadType is called when entering the threadType production.
	EnterThreadType(c *ThreadTypeContext)

	// EnterUtilOption is called when entering the utilOption production.
	EnterUtilOption(c *UtilOptionContext)

	// EnterConnectionOptions is called when entering the connectionOptions production.
	EnterConnectionOptions(c *ConnectionOptionsContext)

	// EnterMasterDefs is called when entering the masterDefs production.
	EnterMasterDefs(c *MasterDefsContext)

	// EnterMasterDef is called when entering the masterDef production.
	EnterMasterDef(c *MasterDefContext)

	// EnterIgnoreServerIds is called when entering the ignoreServerIds production.
	EnterIgnoreServerIds(c *IgnoreServerIdsContext)

	// EnterIgnoreServerId is called when entering the ignoreServerId production.
	EnterIgnoreServerId(c *IgnoreServerIdContext)

	// EnterFilterDefs is called when entering the filterDefs production.
	EnterFilterDefs(c *FilterDefsContext)

	// EnterFilterDef is called when entering the filterDef production.
	EnterFilterDef(c *FilterDefContext)

	// EnterWildTables is called when entering the wildTables production.
	EnterWildTables(c *WildTablesContext)

	// EnterWildTable is called when entering the wildTable production.
	EnterWildTable(c *WildTableContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)

	// ExitInsert is called when exiting the insert production.
	ExitInsert(c *InsertContext)

	// ExitInsertSpecification is called when exiting the insertSpecification production.
	ExitInsertSpecification(c *InsertSpecificationContext)

	// ExitInsertValuesClause is called when exiting the insertValuesClause production.
	ExitInsertValuesClause(c *InsertValuesClauseContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitInsertIdentifier is called when exiting the insertIdentifier production.
	ExitInsertIdentifier(c *InsertIdentifierContext)

	// ExitTableWild is called when exiting the tableWild production.
	ExitTableWild(c *TableWildContext)

	// ExitInsertSelectClause is called when exiting the insertSelectClause production.
	ExitInsertSelectClause(c *InsertSelectClauseContext)

	// ExitOnDuplicateKeyClause is called when exiting the onDuplicateKeyClause production.
	ExitOnDuplicateKeyClause(c *OnDuplicateKeyClauseContext)

	// ExitValueReference is called when exiting the valueReference production.
	ExitValueReference(c *ValueReferenceContext)

	// ExitDerivedColumns is called when exiting the derivedColumns production.
	ExitDerivedColumns(c *DerivedColumnsContext)

	// ExitReplace is called when exiting the replace production.
	ExitReplace(c *ReplaceContext)

	// ExitReplaceSpecification is called when exiting the replaceSpecification production.
	ExitReplaceSpecification(c *ReplaceSpecificationContext)

	// ExitReplaceValuesClause is called when exiting the replaceValuesClause production.
	ExitReplaceValuesClause(c *ReplaceValuesClauseContext)

	// ExitReplaceSelectClause is called when exiting the replaceSelectClause production.
	ExitReplaceSelectClause(c *ReplaceSelectClauseContext)

	// ExitUpdate is called when exiting the update production.
	ExitUpdate(c *UpdateContext)

	// ExitUpdateSpecification_ is called when exiting the updateSpecification_ production.
	ExitUpdateSpecification_(c *UpdateSpecification_Context)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitSetAssignmentsClause is called when exiting the setAssignmentsClause production.
	ExitSetAssignmentsClause(c *SetAssignmentsClauseContext)

	// ExitAssignmentValues is called when exiting the assignmentValues production.
	ExitAssignmentValues(c *AssignmentValuesContext)

	// ExitAssignmentValue is called when exiting the assignmentValue production.
	ExitAssignmentValue(c *AssignmentValueContext)

	// ExitBlobValue is called when exiting the blobValue production.
	ExitBlobValue(c *BlobValueContext)

	// ExitDelete is called when exiting the delete production.
	ExitDelete(c *DeleteContext)

	// ExitDeleteSpecification is called when exiting the deleteSpecification production.
	ExitDeleteSpecification(c *DeleteSpecificationContext)

	// ExitSingleTableClause is called when exiting the singleTableClause production.
	ExitSingleTableClause(c *SingleTableClauseContext)

	// ExitMultipleTablesClause is called when exiting the multipleTablesClause production.
	ExitMultipleTablesClause(c *MultipleTablesClauseContext)

	// ExitSelect is called when exiting the select production.
	ExitSelect(c *SelectContext)

	// ExitSelectWithInto is called when exiting the selectWithInto production.
	ExitSelectWithInto(c *SelectWithIntoContext)

	// ExitQueryExpression is called when exiting the queryExpression production.
	ExitQueryExpression(c *QueryExpressionContext)

	// ExitQueryExpressionBody is called when exiting the queryExpressionBody production.
	ExitQueryExpressionBody(c *QueryExpressionBodyContext)

	// ExitUnionClause is called when exiting the unionClause production.
	ExitUnionClause(c *UnionClauseContext)

	// ExitQueryExpressionParens is called when exiting the queryExpressionParens production.
	ExitQueryExpressionParens(c *QueryExpressionParensContext)

	// ExitQueryPrimary is called when exiting the queryPrimary production.
	ExitQueryPrimary(c *QueryPrimaryContext)

	// ExitQuerySpecification is called when exiting the querySpecification production.
	ExitQuerySpecification(c *QuerySpecificationContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitDoStatement is called when exiting the doStatement production.
	ExitDoStatement(c *DoStatementContext)

	// ExitHandlerStatement is called when exiting the handlerStatement production.
	ExitHandlerStatement(c *HandlerStatementContext)

	// ExitHandlerOpenStatement is called when exiting the handlerOpenStatement production.
	ExitHandlerOpenStatement(c *HandlerOpenStatementContext)

	// ExitHandlerReadIndexStatement is called when exiting the handlerReadIndexStatement production.
	ExitHandlerReadIndexStatement(c *HandlerReadIndexStatementContext)

	// ExitHandlerReadStatement is called when exiting the handlerReadStatement production.
	ExitHandlerReadStatement(c *HandlerReadStatementContext)

	// ExitHandlerCloseStatement is called when exiting the handlerCloseStatement production.
	ExitHandlerCloseStatement(c *HandlerCloseStatementContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitLoadStatement is called when exiting the loadStatement production.
	ExitLoadStatement(c *LoadStatementContext)

	// ExitLoadDataStatement is called when exiting the loadDataStatement production.
	ExitLoadDataStatement(c *LoadDataStatementContext)

	// ExitLoadXmlStatement is called when exiting the loadXmlStatement production.
	ExitLoadXmlStatement(c *LoadXmlStatementContext)

	// ExitExplicitTable is called when exiting the explicitTable production.
	ExitExplicitTable(c *ExplicitTableContext)

	// ExitTableValueConstructor is called when exiting the tableValueConstructor production.
	ExitTableValueConstructor(c *TableValueConstructorContext)

	// ExitRowConstructorList is called when exiting the rowConstructorList production.
	ExitRowConstructorList(c *RowConstructorListContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitCteClause is called when exiting the cteClause production.
	ExitCteClause(c *CteClauseContext)

	// ExitSelectSpecification is called when exiting the selectSpecification production.
	ExitSelectSpecification(c *SelectSpecificationContext)

	// ExitDuplicateSpecification is called when exiting the duplicateSpecification production.
	ExitDuplicateSpecification(c *DuplicateSpecificationContext)

	// ExitProjections is called when exiting the projections production.
	ExitProjections(c *ProjectionsContext)

	// ExitProjection is called when exiting the projection production.
	ExitProjection(c *ProjectionContext)

	// ExitUnqualifiedShorthand is called when exiting the unqualifiedShorthand production.
	ExitUnqualifiedShorthand(c *UnqualifiedShorthandContext)

	// ExitQualifiedShorthand is called when exiting the qualifiedShorthand production.
	ExitQualifiedShorthand(c *QualifiedShorthandContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitTableReferences is called when exiting the tableReferences production.
	ExitTableReferences(c *TableReferencesContext)

	// ExitEscapedTableReference is called when exiting the escapedTableReference production.
	ExitEscapedTableReference(c *EscapedTableReferenceContext)

	// ExitTableReference is called when exiting the tableReference production.
	ExitTableReference(c *TableReferenceContext)

	// ExitTableFactor is called when exiting the tableFactor production.
	ExitTableFactor(c *TableFactorContext)

	// ExitPartitionNames is called when exiting the partitionNames production.
	ExitPartitionNames(c *PartitionNamesContext)

	// ExitIndexHintList is called when exiting the indexHintList production.
	ExitIndexHintList(c *IndexHintListContext)

	// ExitIndexHint is called when exiting the indexHint production.
	ExitIndexHint(c *IndexHintContext)

	// ExitJoinedTable is called when exiting the joinedTable production.
	ExitJoinedTable(c *JoinedTableContext)

	// ExitInnerJoinType is called when exiting the innerJoinType production.
	ExitInnerJoinType(c *InnerJoinTypeContext)

	// ExitOuterJoinType is called when exiting the outerJoinType production.
	ExitOuterJoinType(c *OuterJoinTypeContext)

	// ExitNaturalJoinType is called when exiting the naturalJoinType production.
	ExitNaturalJoinType(c *NaturalJoinTypeContext)

	// ExitJoinSpecification is called when exiting the joinSpecification production.
	ExitJoinSpecification(c *JoinSpecificationContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitLimitRowCount is called when exiting the limitRowCount production.
	ExitLimitRowCount(c *LimitRowCountContext)

	// ExitLimitOffset is called when exiting the limitOffset production.
	ExitLimitOffset(c *LimitOffsetContext)

	// ExitWindowClause is called when exiting the windowClause production.
	ExitWindowClause(c *WindowClauseContext)

	// ExitWindowItem is called when exiting the windowItem production.
	ExitWindowItem(c *WindowItemContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitSelectLinesInto is called when exiting the selectLinesInto production.
	ExitSelectLinesInto(c *SelectLinesIntoContext)

	// ExitSelectFieldsInto is called when exiting the selectFieldsInto production.
	ExitSelectFieldsInto(c *SelectFieldsIntoContext)

	// ExitSelectIntoExpression is called when exiting the selectIntoExpression production.
	ExitSelectIntoExpression(c *SelectIntoExpressionContext)

	// ExitLockClause is called when exiting the lockClause production.
	ExitLockClause(c *LockClauseContext)

	// ExitLockClauseList is called when exiting the lockClauseList production.
	ExitLockClauseList(c *LockClauseListContext)

	// ExitLockStrength is called when exiting the lockStrength production.
	ExitLockStrength(c *LockStrengthContext)

	// ExitLockedRowAction is called when exiting the lockedRowAction production.
	ExitLockedRowAction(c *LockedRowActionContext)

	// ExitTableLockingList is called when exiting the tableLockingList production.
	ExitTableLockingList(c *TableLockingListContext)

	// ExitTableIdentOptWild is called when exiting the tableIdentOptWild production.
	ExitTableIdentOptWild(c *TableIdentOptWildContext)

	// ExitTableAliasRefList is called when exiting the tableAliasRefList production.
	ExitTableAliasRefList(c *TableAliasRefListContext)

	// ExitParameterMarker is called when exiting the parameterMarker production.
	ExitParameterMarker(c *ParameterMarkerContext)

	// ExitCustomKeyword is called when exiting the customKeyword production.
	ExitCustomKeyword(c *CustomKeywordContext)

	// ExitLiterals is called when exiting the literals production.
	ExitLiterals(c *LiteralsContext)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitStringLiterals is called when exiting the stringLiterals production.
	ExitStringLiterals(c *StringLiteralsContext)

	// ExitNumberLiterals is called when exiting the numberLiterals production.
	ExitNumberLiterals(c *NumberLiteralsContext)

	// ExitTemporalLiterals is called when exiting the temporalLiterals production.
	ExitTemporalLiterals(c *TemporalLiteralsContext)

	// ExitHexadecimalLiterals is called when exiting the hexadecimalLiterals production.
	ExitHexadecimalLiterals(c *HexadecimalLiteralsContext)

	// ExitBitValueLiterals is called when exiting the bitValueLiterals production.
	ExitBitValueLiterals(c *BitValueLiteralsContext)

	// ExitBooleanLiterals is called when exiting the booleanLiterals production.
	ExitBooleanLiterals(c *BooleanLiteralsContext)

	// ExitNullValueLiterals is called when exiting the nullValueLiterals production.
	ExitNullValueLiterals(c *NullValueLiteralsContext)

	// ExitCollationName is called when exiting the collationName production.
	ExitCollationName(c *CollationNameContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIdentifierKeywordsUnambiguous is called when exiting the identifierKeywordsUnambiguous production.
	ExitIdentifierKeywordsUnambiguous(c *IdentifierKeywordsUnambiguousContext)

	// ExitIdentifierKeywordsAmbiguous1RolesAndLabels is called when exiting the identifierKeywordsAmbiguous1RolesAndLabels production.
	ExitIdentifierKeywordsAmbiguous1RolesAndLabels(c *IdentifierKeywordsAmbiguous1RolesAndLabelsContext)

	// ExitIdentifierKeywordsAmbiguous2Labels is called when exiting the identifierKeywordsAmbiguous2Labels production.
	ExitIdentifierKeywordsAmbiguous2Labels(c *IdentifierKeywordsAmbiguous2LabelsContext)

	// ExitIdentifierKeywordsAmbiguous3Roles is called when exiting the identifierKeywordsAmbiguous3Roles production.
	ExitIdentifierKeywordsAmbiguous3Roles(c *IdentifierKeywordsAmbiguous3RolesContext)

	// ExitIdentifierKeywordsAmbiguous4SystemVariables is called when exiting the identifierKeywordsAmbiguous4SystemVariables production.
	ExitIdentifierKeywordsAmbiguous4SystemVariables(c *IdentifierKeywordsAmbiguous4SystemVariablesContext)

	// ExitTextOrIdentifier is called when exiting the textOrIdentifier production.
	ExitTextOrIdentifier(c *TextOrIdentifierContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitUserVariable is called when exiting the userVariable production.
	ExitUserVariable(c *UserVariableContext)

	// ExitSystemVariable is called when exiting the systemVariable production.
	ExitSystemVariable(c *SystemVariableContext)

	// ExitSetSystemVariable is called when exiting the setSystemVariable production.
	ExitSetSystemVariable(c *SetSystemVariableContext)

	// ExitOptionType is called when exiting the optionType production.
	ExitOptionType(c *OptionTypeContext)

	// ExitInternalVariableName is called when exiting the internalVariableName production.
	ExitInternalVariableName(c *InternalVariableNameContext)

	// ExitSetExprOrDefault is called when exiting the setExprOrDefault production.
	ExitSetExprOrDefault(c *SetExprOrDefaultContext)

	// ExitTransactionCharacteristics is called when exiting the transactionCharacteristics production.
	ExitTransactionCharacteristics(c *TransactionCharacteristicsContext)

	// ExitIsolationLevel is called when exiting the isolationLevel production.
	ExitIsolationLevel(c *IsolationLevelContext)

	// ExitIsolationTypes is called when exiting the isolationTypes production.
	ExitIsolationTypes(c *IsolationTypesContext)

	// ExitTransactionAccessMode is called when exiting the transactionAccessMode production.
	ExitTransactionAccessMode(c *TransactionAccessModeContext)

	// ExitSchemaName is called when exiting the schemaName production.
	ExitSchemaName(c *SchemaNameContext)

	// ExitSchemaNames is called when exiting the schemaNames production.
	ExitSchemaNames(c *SchemaNamesContext)

	// ExitCharsetName is called when exiting the charsetName production.
	ExitCharsetName(c *CharsetNameContext)

	// ExitSchemaPairs is called when exiting the schemaPairs production.
	ExitSchemaPairs(c *SchemaPairsContext)

	// ExitSchemaPair is called when exiting the schemaPair production.
	ExitSchemaPair(c *SchemaPairContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitIndexName is called when exiting the indexName production.
	ExitIndexName(c *IndexNameContext)

	// ExitConstraintName is called when exiting the constraintName production.
	ExitConstraintName(c *ConstraintNameContext)

	// ExitUserIdentifierOrText is called when exiting the userIdentifierOrText production.
	ExitUserIdentifierOrText(c *UserIdentifierOrTextContext)

	// ExitUserName is called when exiting the userName production.
	ExitUserName(c *UserNameContext)

	// ExitEventName is called when exiting the eventName production.
	ExitEventName(c *EventNameContext)

	// ExitServerName is called when exiting the serverName production.
	ExitServerName(c *ServerNameContext)

	// ExitWrapperName is called when exiting the wrapperName production.
	ExitWrapperName(c *WrapperNameContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitViewName is called when exiting the viewName production.
	ExitViewName(c *ViewNameContext)

	// ExitOwner is called when exiting the owner production.
	ExitOwner(c *OwnerContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitTableList is called when exiting the tableList production.
	ExitTableList(c *TableListContext)

	// ExitViewNames is called when exiting the viewNames production.
	ExitViewNames(c *ViewNamesContext)

	// ExitColumnNames is called when exiting the columnNames production.
	ExitColumnNames(c *ColumnNamesContext)

	// ExitGroupName is called when exiting the groupName production.
	ExitGroupName(c *GroupNameContext)

	// ExitRoutineName is called when exiting the routineName production.
	ExitRoutineName(c *RoutineNameContext)

	// ExitShardLibraryName is called when exiting the shardLibraryName production.
	ExitShardLibraryName(c *ShardLibraryNameContext)

	// ExitComponentName is called when exiting the componentName production.
	ExitComponentName(c *ComponentNameContext)

	// ExitPluginName is called when exiting the pluginName production.
	ExitPluginName(c *PluginNameContext)

	// ExitHostName is called when exiting the hostName production.
	ExitHostName(c *HostNameContext)

	// ExitPort is called when exiting the port production.
	ExitPort(c *PortContext)

	// ExitCloneInstance is called when exiting the cloneInstance production.
	ExitCloneInstance(c *CloneInstanceContext)

	// ExitCloneDir is called when exiting the cloneDir production.
	ExitCloneDir(c *CloneDirContext)

	// ExitChannelName is called when exiting the channelName production.
	ExitChannelName(c *ChannelNameContext)

	// ExitLogName is called when exiting the logName production.
	ExitLogName(c *LogNameContext)

	// ExitRoleName is called when exiting the roleName production.
	ExitRoleName(c *RoleNameContext)

	// ExitRoleIdentifierOrText is called when exiting the roleIdentifierOrText production.
	ExitRoleIdentifierOrText(c *RoleIdentifierOrTextContext)

	// ExitEngineRef is called when exiting the engineRef production.
	ExitEngineRef(c *EngineRefContext)

	// ExitTriggerName is called when exiting the triggerName production.
	ExitTriggerName(c *TriggerNameContext)

	// ExitTriggerTime is called when exiting the triggerTime production.
	ExitTriggerTime(c *TriggerTimeContext)

	// ExitTableOrTables is called when exiting the tableOrTables production.
	ExitTableOrTables(c *TableOrTablesContext)

	// ExitUserOrRole is called when exiting the userOrRole production.
	ExitUserOrRole(c *UserOrRoleContext)

	// ExitPartitionName is called when exiting the partitionName production.
	ExitPartitionName(c *PartitionNameContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitAllOrPartitionNameList is called when exiting the allOrPartitionNameList production.
	ExitAllOrPartitionNameList(c *AllOrPartitionNameListContext)

	// ExitTriggerEvent is called when exiting the triggerEvent production.
	ExitTriggerEvent(c *TriggerEventContext)

	// ExitTriggerOrder is called when exiting the triggerOrder production.
	ExitTriggerOrder(c *TriggerOrderContext)

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

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitBitExpr is called when exiting the bitExpr production.
	ExitBitExpr(c *BitExprContext)

	// ExitSimpleExpr is called when exiting the simpleExpr production.
	ExitSimpleExpr(c *SimpleExprContext)

	// ExitColumnRef is called when exiting the columnRef production.
	ExitColumnRef(c *ColumnRefContext)

	// ExitColumnRefList is called when exiting the columnRefList production.
	ExitColumnRefList(c *ColumnRefListContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitAggregationFunction is called when exiting the aggregationFunction production.
	ExitAggregationFunction(c *AggregationFunctionContext)

	// ExitAggregationFunctionName is called when exiting the aggregationFunctionName production.
	ExitAggregationFunctionName(c *AggregationFunctionNameContext)

	// ExitDistinct is called when exiting the distinct production.
	ExitDistinct(c *DistinctContext)

	// ExitOverClause is called when exiting the overClause production.
	ExitOverClause(c *OverClauseContext)

	// ExitWindowSpecification is called when exiting the windowSpecification production.
	ExitWindowSpecification(c *WindowSpecificationContext)

	// ExitFrameClause is called when exiting the frameClause production.
	ExitFrameClause(c *FrameClauseContext)

	// ExitFrameStart is called when exiting the frameStart production.
	ExitFrameStart(c *FrameStartContext)

	// ExitFrameEnd is called when exiting the frameEnd production.
	ExitFrameEnd(c *FrameEndContext)

	// ExitFrameBetween is called when exiting the frameBetween production.
	ExitFrameBetween(c *FrameBetweenContext)

	// ExitSpecialFunction is called when exiting the specialFunction production.
	ExitSpecialFunction(c *SpecialFunctionContext)

	// ExitCurrentUserFunction is called when exiting the currentUserFunction production.
	ExitCurrentUserFunction(c *CurrentUserFunctionContext)

	// ExitGroupConcatFunction is called when exiting the groupConcatFunction production.
	ExitGroupConcatFunction(c *GroupConcatFunctionContext)

	// ExitWindowFunction is called when exiting the windowFunction production.
	ExitWindowFunction(c *WindowFunctionContext)

	// ExitWindowingClause is called when exiting the windowingClause production.
	ExitWindowingClause(c *WindowingClauseContext)

	// ExitLeadLagInfo is called when exiting the leadLagInfo production.
	ExitLeadLagInfo(c *LeadLagInfoContext)

	// ExitNullTreatment is called when exiting the nullTreatment production.
	ExitNullTreatment(c *NullTreatmentContext)

	// ExitCheckType is called when exiting the checkType production.
	ExitCheckType(c *CheckTypeContext)

	// ExitRepairType is called when exiting the repairType production.
	ExitRepairType(c *RepairTypeContext)

	// ExitCastFunction is called when exiting the castFunction production.
	ExitCastFunction(c *CastFunctionContext)

	// ExitConvertFunction is called when exiting the convertFunction production.
	ExitConvertFunction(c *ConvertFunctionContext)

	// ExitCastType is called when exiting the castType production.
	ExitCastType(c *CastTypeContext)

	// ExitNchar is called when exiting the nchar production.
	ExitNchar(c *NcharContext)

	// ExitPositionFunction is called when exiting the positionFunction production.
	ExitPositionFunction(c *PositionFunctionContext)

	// ExitSubstringFunction is called when exiting the substringFunction production.
	ExitSubstringFunction(c *SubstringFunctionContext)

	// ExitExtractFunction is called when exiting the extractFunction production.
	ExitExtractFunction(c *ExtractFunctionContext)

	// ExitCharFunction is called when exiting the charFunction production.
	ExitCharFunction(c *CharFunctionContext)

	// ExitTrimFunction is called when exiting the trimFunction production.
	ExitTrimFunction(c *TrimFunctionContext)

	// ExitValuesFunction is called when exiting the valuesFunction production.
	ExitValuesFunction(c *ValuesFunctionContext)

	// ExitWeightStringFunction is called when exiting the weightStringFunction production.
	ExitWeightStringFunction(c *WeightStringFunctionContext)

	// ExitLevelClause is called when exiting the levelClause production.
	ExitLevelClause(c *LevelClauseContext)

	// ExitLevelInWeightListElement is called when exiting the levelInWeightListElement production.
	ExitLevelInWeightListElement(c *LevelInWeightListElementContext)

	// ExitRegularFunction is called when exiting the regularFunction production.
	ExitRegularFunction(c *RegularFunctionContext)

	// ExitShorthandRegularFunction is called when exiting the shorthandRegularFunction production.
	ExitShorthandRegularFunction(c *ShorthandRegularFunctionContext)

	// ExitCompleteRegularFunction is called when exiting the completeRegularFunction production.
	ExitCompleteRegularFunction(c *CompleteRegularFunctionContext)

	// ExitRegularFunctionName is called when exiting the regularFunctionName production.
	ExitRegularFunctionName(c *RegularFunctionNameContext)

	// ExitMatchExpression is called when exiting the matchExpression production.
	ExitMatchExpression(c *MatchExpressionContext)

	// ExitMatchSearchModifier is called when exiting the matchSearchModifier production.
	ExitMatchSearchModifier(c *MatchSearchModifierContext)

	// ExitCaseExpression is called when exiting the caseExpression production.
	ExitCaseExpression(c *CaseExpressionContext)

	// ExitDatetimeExpr is called when exiting the datetimeExpr production.
	ExitDatetimeExpr(c *DatetimeExprContext)

	// ExitBinaryLogFileIndexNumber is called when exiting the binaryLogFileIndexNumber production.
	ExitBinaryLogFileIndexNumber(c *BinaryLogFileIndexNumberContext)

	// ExitCaseWhen is called when exiting the caseWhen production.
	ExitCaseWhen(c *CaseWhenContext)

	// ExitCaseElse is called when exiting the caseElse production.
	ExitCaseElse(c *CaseElseContext)

	// ExitIntervalExpression is called when exiting the intervalExpression production.
	ExitIntervalExpression(c *IntervalExpressionContext)

	// ExitIntervalValue is called when exiting the intervalValue production.
	ExitIntervalValue(c *IntervalValueContext)

	// ExitIntervalUnit is called when exiting the intervalUnit production.
	ExitIntervalUnit(c *IntervalUnitContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderByItem is called when exiting the orderByItem production.
	ExitOrderByItem(c *OrderByItemContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitStringList is called when exiting the stringList production.
	ExitStringList(c *StringListContext)

	// ExitTextString is called when exiting the textString production.
	ExitTextString(c *TextStringContext)

	// ExitTextStringHash is called when exiting the textStringHash production.
	ExitTextStringHash(c *TextStringHashContext)

	// ExitFieldOptions is called when exiting the fieldOptions production.
	ExitFieldOptions(c *FieldOptionsContext)

	// ExitPrecision is called when exiting the precision production.
	ExitPrecision(c *PrecisionContext)

	// ExitTypeDatetimePrecision is called when exiting the typeDatetimePrecision production.
	ExitTypeDatetimePrecision(c *TypeDatetimePrecisionContext)

	// ExitCharsetWithOptBinary is called when exiting the charsetWithOptBinary production.
	ExitCharsetWithOptBinary(c *CharsetWithOptBinaryContext)

	// ExitAscii is called when exiting the ascii production.
	ExitAscii(c *AsciiContext)

	// ExitUnicode is called when exiting the unicode production.
	ExitUnicode(c *UnicodeContext)

	// ExitCharset is called when exiting the charset production.
	ExitCharset(c *CharsetContext)

	// ExitDefaultCollation is called when exiting the defaultCollation production.
	ExitDefaultCollation(c *DefaultCollationContext)

	// ExitDefaultEncryption is called when exiting the defaultEncryption production.
	ExitDefaultEncryption(c *DefaultEncryptionContext)

	// ExitDefaultCharset is called when exiting the defaultCharset production.
	ExitDefaultCharset(c *DefaultCharsetContext)

	// ExitSignedLiteral is called when exiting the signedLiteral production.
	ExitSignedLiteral(c *SignedLiteralContext)

	// ExitNow is called when exiting the now production.
	ExitNow(c *NowContext)

	// ExitColumnFormat is called when exiting the columnFormat production.
	ExitColumnFormat(c *ColumnFormatContext)

	// ExitStorageMedia is called when exiting the storageMedia production.
	ExitStorageMedia(c *StorageMediaContext)

	// ExitDirection is called when exiting the direction production.
	ExitDirection(c *DirectionContext)

	// ExitKeyOrIndex is called when exiting the keyOrIndex production.
	ExitKeyOrIndex(c *KeyOrIndexContext)

	// ExitFieldLength is called when exiting the fieldLength production.
	ExitFieldLength(c *FieldLengthContext)

	// ExitCharacterSet is called when exiting the characterSet production.
	ExitCharacterSet(c *CharacterSetContext)

	// ExitCollateClause is called when exiting the collateClause production.
	ExitCollateClause(c *CollateClauseContext)

	// ExitFieldOrVarSpec is called when exiting the fieldOrVarSpec production.
	ExitFieldOrVarSpec(c *FieldOrVarSpecContext)

	// ExitNotExistClause is called when exiting the notExistClause production.
	ExitNotExistClause(c *NotExistClauseContext)

	// ExitExistClause is called when exiting the existClause production.
	ExitExistClause(c *ExistClauseContext)

	// ExitConnectionId is called when exiting the connectionId production.
	ExitConnectionId(c *ConnectionIdContext)

	// ExitLabelName is called when exiting the labelName production.
	ExitLabelName(c *LabelNameContext)

	// ExitCursorName is called when exiting the cursorName production.
	ExitCursorName(c *CursorNameContext)

	// ExitConditionName is called when exiting the conditionName production.
	ExitConditionName(c *ConditionNameContext)

	// ExitUnionOption is called when exiting the unionOption production.
	ExitUnionOption(c *UnionOptionContext)

	// ExitNoWriteToBinLog is called when exiting the noWriteToBinLog production.
	ExitNoWriteToBinLog(c *NoWriteToBinLogContext)

	// ExitChannelOption is called when exiting the channelOption production.
	ExitChannelOption(c *ChannelOptionContext)

	// ExitPreparedStatement is called when exiting the preparedStatement production.
	ExitPreparedStatement(c *PreparedStatementContext)

	// ExitExecuteStatement is called when exiting the executeStatement production.
	ExitExecuteStatement(c *ExecuteStatementContext)

	// ExitExecuteVarList is called when exiting the executeVarList production.
	ExitExecuteVarList(c *ExecuteVarListContext)

	// ExitAlterStatement is called when exiting the alterStatement production.
	ExitAlterStatement(c *AlterStatementContext)

	// ExitCreateTable is called when exiting the createTable production.
	ExitCreateTable(c *CreateTableContext)

	// ExitPartitionClause is called when exiting the partitionClause production.
	ExitPartitionClause(c *PartitionClauseContext)

	// ExitPartitionTypeDef is called when exiting the partitionTypeDef production.
	ExitPartitionTypeDef(c *PartitionTypeDefContext)

	// ExitSubPartitions is called when exiting the subPartitions production.
	ExitSubPartitions(c *SubPartitionsContext)

	// ExitPartitionKeyAlgorithm is called when exiting the partitionKeyAlgorithm production.
	ExitPartitionKeyAlgorithm(c *PartitionKeyAlgorithmContext)

	// ExitDuplicateAsQueryExpression is called when exiting the duplicateAsQueryExpression production.
	ExitDuplicateAsQueryExpression(c *DuplicateAsQueryExpressionContext)

	// ExitAlterTable is called when exiting the alterTable production.
	ExitAlterTable(c *AlterTableContext)

	// ExitStandaloneAlterTableAction is called when exiting the standaloneAlterTableAction production.
	ExitStandaloneAlterTableAction(c *StandaloneAlterTableActionContext)

	// ExitAlterTableActions is called when exiting the alterTableActions production.
	ExitAlterTableActions(c *AlterTableActionsContext)

	// ExitAlterTablePartitionOptions is called when exiting the alterTablePartitionOptions production.
	ExitAlterTablePartitionOptions(c *AlterTablePartitionOptionsContext)

	// ExitAlterCommandList is called when exiting the alterCommandList production.
	ExitAlterCommandList(c *AlterCommandListContext)

	// ExitAlterList is called when exiting the alterList production.
	ExitAlterList(c *AlterListContext)

	// ExitCreateTableOptionsSpaceSeparated is called when exiting the createTableOptionsSpaceSeparated production.
	ExitCreateTableOptionsSpaceSeparated(c *CreateTableOptionsSpaceSeparatedContext)

	// ExitAddColumn is called when exiting the addColumn production.
	ExitAddColumn(c *AddColumnContext)

	// ExitAddTableConstraint is called when exiting the addTableConstraint production.
	ExitAddTableConstraint(c *AddTableConstraintContext)

	// ExitChangeColumn is called when exiting the changeColumn production.
	ExitChangeColumn(c *ChangeColumnContext)

	// ExitModifyColumn is called when exiting the modifyColumn production.
	ExitModifyColumn(c *ModifyColumnContext)

	// ExitAlterTableDrop is called when exiting the alterTableDrop production.
	ExitAlterTableDrop(c *AlterTableDropContext)

	// ExitDisableKeys is called when exiting the disableKeys production.
	ExitDisableKeys(c *DisableKeysContext)

	// ExitEnableKeys is called when exiting the enableKeys production.
	ExitEnableKeys(c *EnableKeysContext)

	// ExitAlterColumn is called when exiting the alterColumn production.
	ExitAlterColumn(c *AlterColumnContext)

	// ExitAlterIndex is called when exiting the alterIndex production.
	ExitAlterIndex(c *AlterIndexContext)

	// ExitAlterCheck is called when exiting the alterCheck production.
	ExitAlterCheck(c *AlterCheckContext)

	// ExitAlterConstraint is called when exiting the alterConstraint production.
	ExitAlterConstraint(c *AlterConstraintContext)

	// ExitRenameColumn is called when exiting the renameColumn production.
	ExitRenameColumn(c *RenameColumnContext)

	// ExitAlterRenameTable is called when exiting the alterRenameTable production.
	ExitAlterRenameTable(c *AlterRenameTableContext)

	// ExitRenameIndex is called when exiting the renameIndex production.
	ExitRenameIndex(c *RenameIndexContext)

	// ExitAlterConvert is called when exiting the alterConvert production.
	ExitAlterConvert(c *AlterConvertContext)

	// ExitAlterTableForce is called when exiting the alterTableForce production.
	ExitAlterTableForce(c *AlterTableForceContext)

	// ExitAlterTableOrder is called when exiting the alterTableOrder production.
	ExitAlterTableOrder(c *AlterTableOrderContext)

	// ExitAlterOrderList is called when exiting the alterOrderList production.
	ExitAlterOrderList(c *AlterOrderListContext)

	// ExitTableConstraintDef is called when exiting the tableConstraintDef production.
	ExitTableConstraintDef(c *TableConstraintDefContext)

	// ExitAlterCommandsModifierList is called when exiting the alterCommandsModifierList production.
	ExitAlterCommandsModifierList(c *AlterCommandsModifierListContext)

	// ExitAlterCommandsModifier is called when exiting the alterCommandsModifier production.
	ExitAlterCommandsModifier(c *AlterCommandsModifierContext)

	// ExitWithValidation is called when exiting the withValidation production.
	ExitWithValidation(c *WithValidationContext)

	// ExitStandaloneAlterCommands is called when exiting the standaloneAlterCommands production.
	ExitStandaloneAlterCommands(c *StandaloneAlterCommandsContext)

	// ExitAlterPartition is called when exiting the alterPartition production.
	ExitAlterPartition(c *AlterPartitionContext)

	// ExitConstraintClause is called when exiting the constraintClause production.
	ExitConstraintClause(c *ConstraintClauseContext)

	// ExitTableElementList is called when exiting the tableElementList production.
	ExitTableElementList(c *TableElementListContext)

	// ExitTableElement is called when exiting the tableElement production.
	ExitTableElement(c *TableElementContext)

	// ExitRestrict is called when exiting the restrict production.
	ExitRestrict(c *RestrictContext)

	// ExitFulltextIndexOption is called when exiting the fulltextIndexOption production.
	ExitFulltextIndexOption(c *FulltextIndexOptionContext)

	// ExitDropTable is called when exiting the dropTable production.
	ExitDropTable(c *DropTableContext)

	// ExitDropIndex is called when exiting the dropIndex production.
	ExitDropIndex(c *DropIndexContext)

	// ExitAlterAlgorithmOption is called when exiting the alterAlgorithmOption production.
	ExitAlterAlgorithmOption(c *AlterAlgorithmOptionContext)

	// ExitAlterLockOption is called when exiting the alterLockOption production.
	ExitAlterLockOption(c *AlterLockOptionContext)

	// ExitTruncateTable is called when exiting the truncateTable production.
	ExitTruncateTable(c *TruncateTableContext)

	// ExitCreateIndex is called when exiting the createIndex production.
	ExitCreateIndex(c *CreateIndexContext)

	// ExitCreateDatabase is called when exiting the createDatabase production.
	ExitCreateDatabase(c *CreateDatabaseContext)

	// ExitAlterDatabase is called when exiting the alterDatabase production.
	ExitAlterDatabase(c *AlterDatabaseContext)

	// ExitCreateDatabaseSpecification_ is called when exiting the createDatabaseSpecification_ production.
	ExitCreateDatabaseSpecification_(c *CreateDatabaseSpecification_Context)

	// ExitAlterDatabaseSpecification_ is called when exiting the alterDatabaseSpecification_ production.
	ExitAlterDatabaseSpecification_(c *AlterDatabaseSpecification_Context)

	// ExitDropDatabase is called when exiting the dropDatabase production.
	ExitDropDatabase(c *DropDatabaseContext)

	// ExitAlterInstance is called when exiting the alterInstance production.
	ExitAlterInstance(c *AlterInstanceContext)

	// ExitInstanceAction is called when exiting the instanceAction production.
	ExitInstanceAction(c *InstanceActionContext)

	// ExitChannel is called when exiting the channel production.
	ExitChannel(c *ChannelContext)

	// ExitCreateEvent is called when exiting the createEvent production.
	ExitCreateEvent(c *CreateEventContext)

	// ExitAlterEvent is called when exiting the alterEvent production.
	ExitAlterEvent(c *AlterEventContext)

	// ExitDropEvent is called when exiting the dropEvent production.
	ExitDropEvent(c *DropEventContext)

	// ExitCreateFunction is called when exiting the createFunction production.
	ExitCreateFunction(c *CreateFunctionContext)

	// ExitAlterFunction is called when exiting the alterFunction production.
	ExitAlterFunction(c *AlterFunctionContext)

	// ExitDropFunction is called when exiting the dropFunction production.
	ExitDropFunction(c *DropFunctionContext)

	// ExitCreateProcedure is called when exiting the createProcedure production.
	ExitCreateProcedure(c *CreateProcedureContext)

	// ExitAlterProcedure is called when exiting the alterProcedure production.
	ExitAlterProcedure(c *AlterProcedureContext)

	// ExitDropProcedure is called when exiting the dropProcedure production.
	ExitDropProcedure(c *DropProcedureContext)

	// ExitCreateServer is called when exiting the createServer production.
	ExitCreateServer(c *CreateServerContext)

	// ExitAlterServer is called when exiting the alterServer production.
	ExitAlterServer(c *AlterServerContext)

	// ExitDropServer is called when exiting the dropServer production.
	ExitDropServer(c *DropServerContext)

	// ExitCreateView is called when exiting the createView production.
	ExitCreateView(c *CreateViewContext)

	// ExitAlterView is called when exiting the alterView production.
	ExitAlterView(c *AlterViewContext)

	// ExitDropView is called when exiting the dropView production.
	ExitDropView(c *DropViewContext)

	// ExitCreateTablespace is called when exiting the createTablespace production.
	ExitCreateTablespace(c *CreateTablespaceContext)

	// ExitCreateTablespaceInnodb is called when exiting the createTablespaceInnodb production.
	ExitCreateTablespaceInnodb(c *CreateTablespaceInnodbContext)

	// ExitCreateTablespaceNdb is called when exiting the createTablespaceNdb production.
	ExitCreateTablespaceNdb(c *CreateTablespaceNdbContext)

	// ExitAlterTablespace is called when exiting the alterTablespace production.
	ExitAlterTablespace(c *AlterTablespaceContext)

	// ExitAlterTablespaceNdb is called when exiting the alterTablespaceNdb production.
	ExitAlterTablespaceNdb(c *AlterTablespaceNdbContext)

	// ExitAlterTablespaceInnodb is called when exiting the alterTablespaceInnodb production.
	ExitAlterTablespaceInnodb(c *AlterTablespaceInnodbContext)

	// ExitDropTablespace is called when exiting the dropTablespace production.
	ExitDropTablespace(c *DropTablespaceContext)

	// ExitCreateLogfileGroup is called when exiting the createLogfileGroup production.
	ExitCreateLogfileGroup(c *CreateLogfileGroupContext)

	// ExitAlterLogfileGroup is called when exiting the alterLogfileGroup production.
	ExitAlterLogfileGroup(c *AlterLogfileGroupContext)

	// ExitDropLogfileGroup is called when exiting the dropLogfileGroup production.
	ExitDropLogfileGroup(c *DropLogfileGroupContext)

	// ExitCreateTrigger is called when exiting the createTrigger production.
	ExitCreateTrigger(c *CreateTriggerContext)

	// ExitDropTrigger is called when exiting the dropTrigger production.
	ExitDropTrigger(c *DropTriggerContext)

	// ExitRenameTable is called when exiting the renameTable production.
	ExitRenameTable(c *RenameTableContext)

	// ExitCreateDefinitionClause is called when exiting the createDefinitionClause production.
	ExitCreateDefinitionClause(c *CreateDefinitionClauseContext)

	// ExitColumnDefinition is called when exiting the columnDefinition production.
	ExitColumnDefinition(c *ColumnDefinitionContext)

	// ExitFieldDefinition is called when exiting the fieldDefinition production.
	ExitFieldDefinition(c *FieldDefinitionContext)

	// ExitColumnAttribute is called when exiting the columnAttribute production.
	ExitColumnAttribute(c *ColumnAttributeContext)

	// ExitCheckConstraint is called when exiting the checkConstraint production.
	ExitCheckConstraint(c *CheckConstraintContext)

	// ExitConstraintEnforcement is called when exiting the constraintEnforcement production.
	ExitConstraintEnforcement(c *ConstraintEnforcementContext)

	// ExitGeneratedOption is called when exiting the generatedOption production.
	ExitGeneratedOption(c *GeneratedOptionContext)

	// ExitReferenceDefinition is called when exiting the referenceDefinition production.
	ExitReferenceDefinition(c *ReferenceDefinitionContext)

	// ExitOnUpdateDelete is called when exiting the onUpdateDelete production.
	ExitOnUpdateDelete(c *OnUpdateDeleteContext)

	// ExitReferenceOption is called when exiting the referenceOption production.
	ExitReferenceOption(c *ReferenceOptionContext)

	// ExitIndexNameAndType is called when exiting the indexNameAndType production.
	ExitIndexNameAndType(c *IndexNameAndTypeContext)

	// ExitIndexType is called when exiting the indexType production.
	ExitIndexType(c *IndexTypeContext)

	// ExitIndexTypeClause is called when exiting the indexTypeClause production.
	ExitIndexTypeClause(c *IndexTypeClauseContext)

	// ExitKeyParts is called when exiting the keyParts production.
	ExitKeyParts(c *KeyPartsContext)

	// ExitKeyPart is called when exiting the keyPart production.
	ExitKeyPart(c *KeyPartContext)

	// ExitKeyPartWithExpression is called when exiting the keyPartWithExpression production.
	ExitKeyPartWithExpression(c *KeyPartWithExpressionContext)

	// ExitKeyListWithExpression is called when exiting the keyListWithExpression production.
	ExitKeyListWithExpression(c *KeyListWithExpressionContext)

	// ExitIndexOption is called when exiting the indexOption production.
	ExitIndexOption(c *IndexOptionContext)

	// ExitCommonIndexOption is called when exiting the commonIndexOption production.
	ExitCommonIndexOption(c *CommonIndexOptionContext)

	// ExitVisibility is called when exiting the visibility production.
	ExitVisibility(c *VisibilityContext)

	// ExitCreateLikeClause is called when exiting the createLikeClause production.
	ExitCreateLikeClause(c *CreateLikeClauseContext)

	// ExitCreateIndexSpecification is called when exiting the createIndexSpecification production.
	ExitCreateIndexSpecification(c *CreateIndexSpecificationContext)

	// ExitCreateTableOptions is called when exiting the createTableOptions production.
	ExitCreateTableOptions(c *CreateTableOptionsContext)

	// ExitCreateTableOption is called when exiting the createTableOption production.
	ExitCreateTableOption(c *CreateTableOptionContext)

	// ExitCreateSRSStatement is called when exiting the createSRSStatement production.
	ExitCreateSRSStatement(c *CreateSRSStatementContext)

	// ExitDropSRSStatement is called when exiting the dropSRSStatement production.
	ExitDropSRSStatement(c *DropSRSStatementContext)

	// ExitSrsAttribute is called when exiting the srsAttribute production.
	ExitSrsAttribute(c *SrsAttributeContext)

	// ExitPlace is called when exiting the place production.
	ExitPlace(c *PlaceContext)

	// ExitPartitionDefinitions is called when exiting the partitionDefinitions production.
	ExitPartitionDefinitions(c *PartitionDefinitionsContext)

	// ExitPartitionDefinition is called when exiting the partitionDefinition production.
	ExitPartitionDefinition(c *PartitionDefinitionContext)

	// ExitPartitionLessThanValue is called when exiting the partitionLessThanValue production.
	ExitPartitionLessThanValue(c *PartitionLessThanValueContext)

	// ExitPartitionValueList is called when exiting the partitionValueList production.
	ExitPartitionValueList(c *PartitionValueListContext)

	// ExitPartitionDefinitionOption is called when exiting the partitionDefinitionOption production.
	ExitPartitionDefinitionOption(c *PartitionDefinitionOptionContext)

	// ExitSubpartitionDefinition is called when exiting the subpartitionDefinition production.
	ExitSubpartitionDefinition(c *SubpartitionDefinitionContext)

	// ExitOwnerStatement is called when exiting the ownerStatement production.
	ExitOwnerStatement(c *OwnerStatementContext)

	// ExitScheduleExpression is called when exiting the scheduleExpression production.
	ExitScheduleExpression(c *ScheduleExpressionContext)

	// ExitTimestampValue is called when exiting the timestampValue production.
	ExitTimestampValue(c *TimestampValueContext)

	// ExitRoutineBody is called when exiting the routineBody production.
	ExitRoutineBody(c *RoutineBodyContext)

	// ExitServerOption is called when exiting the serverOption production.
	ExitServerOption(c *ServerOptionContext)

	// ExitRoutineOption is called when exiting the routineOption production.
	ExitRoutineOption(c *RoutineOptionContext)

	// ExitProcedureParameter is called when exiting the procedureParameter production.
	ExitProcedureParameter(c *ProcedureParameterContext)

	// ExitFileSizeLiteral is called when exiting the fileSizeLiteral production.
	ExitFileSizeLiteral(c *FileSizeLiteralContext)

	// ExitSimpleStatement is called when exiting the simpleStatement production.
	ExitSimpleStatement(c *SimpleStatementContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitValidStatement is called when exiting the validStatement production.
	ExitValidStatement(c *ValidStatementContext)

	// ExitBeginStatement is called when exiting the beginStatement production.
	ExitBeginStatement(c *BeginStatementContext)

	// ExitDeclareStatement is called when exiting the declareStatement production.
	ExitDeclareStatement(c *DeclareStatementContext)

	// ExitFlowControlStatement is called when exiting the flowControlStatement production.
	ExitFlowControlStatement(c *FlowControlStatementContext)

	// ExitCaseStatement is called when exiting the caseStatement production.
	ExitCaseStatement(c *CaseStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitIterateStatement is called when exiting the iterateStatement production.
	ExitIterateStatement(c *IterateStatementContext)

	// ExitLeaveStatement is called when exiting the leaveStatement production.
	ExitLeaveStatement(c *LeaveStatementContext)

	// ExitLoopStatement is called when exiting the loopStatement production.
	ExitLoopStatement(c *LoopStatementContext)

	// ExitRepeatStatement is called when exiting the repeatStatement production.
	ExitRepeatStatement(c *RepeatStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitCursorStatement is called when exiting the cursorStatement production.
	ExitCursorStatement(c *CursorStatementContext)

	// ExitCursorCloseStatement is called when exiting the cursorCloseStatement production.
	ExitCursorCloseStatement(c *CursorCloseStatementContext)

	// ExitCursorDeclareStatement is called when exiting the cursorDeclareStatement production.
	ExitCursorDeclareStatement(c *CursorDeclareStatementContext)

	// ExitCursorFetchStatement is called when exiting the cursorFetchStatement production.
	ExitCursorFetchStatement(c *CursorFetchStatementContext)

	// ExitCursorOpenStatement is called when exiting the cursorOpenStatement production.
	ExitCursorOpenStatement(c *CursorOpenStatementContext)

	// ExitConditionHandlingStatement is called when exiting the conditionHandlingStatement production.
	ExitConditionHandlingStatement(c *ConditionHandlingStatementContext)

	// ExitDeclareConditionStatement is called when exiting the declareConditionStatement production.
	ExitDeclareConditionStatement(c *DeclareConditionStatementContext)

	// ExitDeclareHandlerStatement is called when exiting the declareHandlerStatement production.
	ExitDeclareHandlerStatement(c *DeclareHandlerStatementContext)

	// ExitGetDiagnosticsStatement is called when exiting the getDiagnosticsStatement production.
	ExitGetDiagnosticsStatement(c *GetDiagnosticsStatementContext)

	// ExitStatementInformationItem is called when exiting the statementInformationItem production.
	ExitStatementInformationItem(c *StatementInformationItemContext)

	// ExitConditionInformationItem is called when exiting the conditionInformationItem production.
	ExitConditionInformationItem(c *ConditionInformationItemContext)

	// ExitConditionNumber is called when exiting the conditionNumber production.
	ExitConditionNumber(c *ConditionNumberContext)

	// ExitStatementInformationItemName is called when exiting the statementInformationItemName production.
	ExitStatementInformationItemName(c *StatementInformationItemNameContext)

	// ExitConditionInformationItemName is called when exiting the conditionInformationItemName production.
	ExitConditionInformationItemName(c *ConditionInformationItemNameContext)

	// ExitHandlerAction is called when exiting the handlerAction production.
	ExitHandlerAction(c *HandlerActionContext)

	// ExitConditionValue is called when exiting the conditionValue production.
	ExitConditionValue(c *ConditionValueContext)

	// ExitResignalStatement is called when exiting the resignalStatement production.
	ExitResignalStatement(c *ResignalStatementContext)

	// ExitSignalStatement is called when exiting the signalStatement production.
	ExitSignalStatement(c *SignalStatementContext)

	// ExitSignalInformationItem is called when exiting the signalInformationItem production.
	ExitSignalInformationItem(c *SignalInformationItemContext)

	// ExitUse is called when exiting the use production.
	ExitUse(c *UseContext)

	// ExitHelp is called when exiting the help production.
	ExitHelp(c *HelpContext)

	// ExitExplain is called when exiting the explain production.
	ExitExplain(c *ExplainContext)

	// ExitShowDatabases is called when exiting the showDatabases production.
	ExitShowDatabases(c *ShowDatabasesContext)

	// ExitShowTables is called when exiting the showTables production.
	ExitShowTables(c *ShowTablesContext)

	// ExitShowTableStatus is called when exiting the showTableStatus production.
	ExitShowTableStatus(c *ShowTableStatusContext)

	// ExitShowColumns is called when exiting the showColumns production.
	ExitShowColumns(c *ShowColumnsContext)

	// ExitShowIndex is called when exiting the showIndex production.
	ExitShowIndex(c *ShowIndexContext)

	// ExitShowCreateTable is called when exiting the showCreateTable production.
	ExitShowCreateTable(c *ShowCreateTableContext)

	// ExitFromSchema is called when exiting the fromSchema production.
	ExitFromSchema(c *FromSchemaContext)

	// ExitFromTable is called when exiting the fromTable production.
	ExitFromTable(c *FromTableContext)

	// ExitShowLike is called when exiting the showLike production.
	ExitShowLike(c *ShowLikeContext)

	// ExitShowWhereClause is called when exiting the showWhereClause production.
	ExitShowWhereClause(c *ShowWhereClauseContext)

	// ExitShowFilter is called when exiting the showFilter production.
	ExitShowFilter(c *ShowFilterContext)

	// ExitShowProfileType is called when exiting the showProfileType production.
	ExitShowProfileType(c *ShowProfileTypeContext)

	// ExitSetVariable is called when exiting the setVariable production.
	ExitSetVariable(c *SetVariableContext)

	// ExitOptionValueList is called when exiting the optionValueList production.
	ExitOptionValueList(c *OptionValueListContext)

	// ExitOptionValueNoOptionType is called when exiting the optionValueNoOptionType production.
	ExitOptionValueNoOptionType(c *OptionValueNoOptionTypeContext)

	// ExitOptionValue is called when exiting the optionValue production.
	ExitOptionValue(c *OptionValueContext)

	// ExitShowBinaryLogs is called when exiting the showBinaryLogs production.
	ExitShowBinaryLogs(c *ShowBinaryLogsContext)

	// ExitShowBinlogEvents is called when exiting the showBinlogEvents production.
	ExitShowBinlogEvents(c *ShowBinlogEventsContext)

	// ExitShowCharacterSet is called when exiting the showCharacterSet production.
	ExitShowCharacterSet(c *ShowCharacterSetContext)

	// ExitShowCollation is called when exiting the showCollation production.
	ExitShowCollation(c *ShowCollationContext)

	// ExitShowCreateDatabase is called when exiting the showCreateDatabase production.
	ExitShowCreateDatabase(c *ShowCreateDatabaseContext)

	// ExitShowCreateEvent is called when exiting the showCreateEvent production.
	ExitShowCreateEvent(c *ShowCreateEventContext)

	// ExitShowCreateFunction is called when exiting the showCreateFunction production.
	ExitShowCreateFunction(c *ShowCreateFunctionContext)

	// ExitShowCreateProcedure is called when exiting the showCreateProcedure production.
	ExitShowCreateProcedure(c *ShowCreateProcedureContext)

	// ExitShowCreateTrigger is called when exiting the showCreateTrigger production.
	ExitShowCreateTrigger(c *ShowCreateTriggerContext)

	// ExitShowCreateUser is called when exiting the showCreateUser production.
	ExitShowCreateUser(c *ShowCreateUserContext)

	// ExitShowCreateView is called when exiting the showCreateView production.
	ExitShowCreateView(c *ShowCreateViewContext)

	// ExitShowEngine is called when exiting the showEngine production.
	ExitShowEngine(c *ShowEngineContext)

	// ExitShowEngines is called when exiting the showEngines production.
	ExitShowEngines(c *ShowEnginesContext)

	// ExitShowCharset is called when exiting the showCharset production.
	ExitShowCharset(c *ShowCharsetContext)

	// ExitShowErrors is called when exiting the showErrors production.
	ExitShowErrors(c *ShowErrorsContext)

	// ExitShowEvents is called when exiting the showEvents production.
	ExitShowEvents(c *ShowEventsContext)

	// ExitShowFunctionCode is called when exiting the showFunctionCode production.
	ExitShowFunctionCode(c *ShowFunctionCodeContext)

	// ExitShowFunctionStatus is called when exiting the showFunctionStatus production.
	ExitShowFunctionStatus(c *ShowFunctionStatusContext)

	// ExitShowGrant is called when exiting the showGrant production.
	ExitShowGrant(c *ShowGrantContext)

	// ExitShowMasterStatus is called when exiting the showMasterStatus production.
	ExitShowMasterStatus(c *ShowMasterStatusContext)

	// ExitShowOpenTables is called when exiting the showOpenTables production.
	ExitShowOpenTables(c *ShowOpenTablesContext)

	// ExitShowPlugins is called when exiting the showPlugins production.
	ExitShowPlugins(c *ShowPluginsContext)

	// ExitShowPrivileges is called when exiting the showPrivileges production.
	ExitShowPrivileges(c *ShowPrivilegesContext)

	// ExitShowProcedureCode is called when exiting the showProcedureCode production.
	ExitShowProcedureCode(c *ShowProcedureCodeContext)

	// ExitShowProcedureStatus is called when exiting the showProcedureStatus production.
	ExitShowProcedureStatus(c *ShowProcedureStatusContext)

	// ExitShowProcesslist is called when exiting the showProcesslist production.
	ExitShowProcesslist(c *ShowProcesslistContext)

	// ExitShowProfile is called when exiting the showProfile production.
	ExitShowProfile(c *ShowProfileContext)

	// ExitShowProfiles is called when exiting the showProfiles production.
	ExitShowProfiles(c *ShowProfilesContext)

	// ExitShowRelaylogEvent is called when exiting the showRelaylogEvent production.
	ExitShowRelaylogEvent(c *ShowRelaylogEventContext)

	// ExitShowSlavehost is called when exiting the showSlavehost production.
	ExitShowSlavehost(c *ShowSlavehostContext)

	// ExitShowSlaveStatus is called when exiting the showSlaveStatus production.
	ExitShowSlaveStatus(c *ShowSlaveStatusContext)

	// ExitShowStatus is called when exiting the showStatus production.
	ExitShowStatus(c *ShowStatusContext)

	// ExitShowTriggers is called when exiting the showTriggers production.
	ExitShowTriggers(c *ShowTriggersContext)

	// ExitShowVariables is called when exiting the showVariables production.
	ExitShowVariables(c *ShowVariablesContext)

	// ExitShowWarnings is called when exiting the showWarnings production.
	ExitShowWarnings(c *ShowWarningsContext)

	// ExitShowReplicas is called when exiting the showReplicas production.
	ExitShowReplicas(c *ShowReplicasContext)

	// ExitShowReplicaStatus is called when exiting the showReplicaStatus production.
	ExitShowReplicaStatus(c *ShowReplicaStatusContext)

	// ExitSetCharacter is called when exiting the setCharacter production.
	ExitSetCharacter(c *SetCharacterContext)

	// ExitClone is called when exiting the clone production.
	ExitClone(c *CloneContext)

	// ExitCloneAction is called when exiting the cloneAction production.
	ExitCloneAction(c *CloneActionContext)

	// ExitCreateLoadableFunction is called when exiting the createLoadableFunction production.
	ExitCreateLoadableFunction(c *CreateLoadableFunctionContext)

	// ExitInstall is called when exiting the install production.
	ExitInstall(c *InstallContext)

	// ExitUninstall is called when exiting the uninstall production.
	ExitUninstall(c *UninstallContext)

	// ExitInstallComponent is called when exiting the installComponent production.
	ExitInstallComponent(c *InstallComponentContext)

	// ExitInstallPlugin is called when exiting the installPlugin production.
	ExitInstallPlugin(c *InstallPluginContext)

	// ExitUninstallComponent is called when exiting the uninstallComponent production.
	ExitUninstallComponent(c *UninstallComponentContext)

	// ExitUninstallPlugin is called when exiting the uninstallPlugin production.
	ExitUninstallPlugin(c *UninstallPluginContext)

	// ExitAnalyzeTable is called when exiting the analyzeTable production.
	ExitAnalyzeTable(c *AnalyzeTableContext)

	// ExitHistogram is called when exiting the histogram production.
	ExitHistogram(c *HistogramContext)

	// ExitCheckTable is called when exiting the checkTable production.
	ExitCheckTable(c *CheckTableContext)

	// ExitCheckTableOption is called when exiting the checkTableOption production.
	ExitCheckTableOption(c *CheckTableOptionContext)

	// ExitChecksumTable is called when exiting the checksumTable production.
	ExitChecksumTable(c *ChecksumTableContext)

	// ExitOptimizeTable is called when exiting the optimizeTable production.
	ExitOptimizeTable(c *OptimizeTableContext)

	// ExitRepairTable is called when exiting the repairTable production.
	ExitRepairTable(c *RepairTableContext)

	// ExitAlterResourceGroup is called when exiting the alterResourceGroup production.
	ExitAlterResourceGroup(c *AlterResourceGroupContext)

	// ExitVcpuSpec is called when exiting the vcpuSpec production.
	ExitVcpuSpec(c *VcpuSpecContext)

	// ExitCreateResourceGroup is called when exiting the createResourceGroup production.
	ExitCreateResourceGroup(c *CreateResourceGroupContext)

	// ExitDropResourceGroup is called when exiting the dropResourceGroup production.
	ExitDropResourceGroup(c *DropResourceGroupContext)

	// ExitSetResourceGroup is called when exiting the setResourceGroup production.
	ExitSetResourceGroup(c *SetResourceGroupContext)

	// ExitBinlog is called when exiting the binlog production.
	ExitBinlog(c *BinlogContext)

	// ExitCacheIndex is called when exiting the cacheIndex production.
	ExitCacheIndex(c *CacheIndexContext)

	// ExitCacheTableIndexList is called when exiting the cacheTableIndexList production.
	ExitCacheTableIndexList(c *CacheTableIndexListContext)

	// ExitPartitionList is called when exiting the partitionList production.
	ExitPartitionList(c *PartitionListContext)

	// ExitFlush is called when exiting the flush production.
	ExitFlush(c *FlushContext)

	// ExitFlushOption is called when exiting the flushOption production.
	ExitFlushOption(c *FlushOptionContext)

	// ExitTablesOption is called when exiting the tablesOption production.
	ExitTablesOption(c *TablesOptionContext)

	// ExitKill is called when exiting the kill production.
	ExitKill(c *KillContext)

	// ExitLoadIndexInfo is called when exiting the loadIndexInfo production.
	ExitLoadIndexInfo(c *LoadIndexInfoContext)

	// ExitLoadTableIndexList is called when exiting the loadTableIndexList production.
	ExitLoadTableIndexList(c *LoadTableIndexListContext)

	// ExitResetStatement is called when exiting the resetStatement production.
	ExitResetStatement(c *ResetStatementContext)

	// ExitResetOption is called when exiting the resetOption production.
	ExitResetOption(c *ResetOptionContext)

	// ExitResetPersist is called when exiting the resetPersist production.
	ExitResetPersist(c *ResetPersistContext)

	// ExitRestart is called when exiting the restart production.
	ExitRestart(c *RestartContext)

	// ExitShutdown is called when exiting the shutdown production.
	ExitShutdown(c *ShutdownContext)

	// ExitExplainType is called when exiting the explainType production.
	ExitExplainType(c *ExplainTypeContext)

	// ExitExplainableStatement is called when exiting the explainableStatement production.
	ExitExplainableStatement(c *ExplainableStatementContext)

	// ExitFormatName is called when exiting the formatName production.
	ExitFormatName(c *FormatNameContext)

	// ExitShow is called when exiting the show production.
	ExitShow(c *ShowContext)

	// ExitSetTransaction is called when exiting the setTransaction production.
	ExitSetTransaction(c *SetTransactionContext)

	// ExitSetAutoCommit is called when exiting the setAutoCommit production.
	ExitSetAutoCommit(c *SetAutoCommitContext)

	// ExitBeginTransaction is called when exiting the beginTransaction production.
	ExitBeginTransaction(c *BeginTransactionContext)

	// ExitTransactionCharacteristic is called when exiting the transactionCharacteristic production.
	ExitTransactionCharacteristic(c *TransactionCharacteristicContext)

	// ExitCommit is called when exiting the commit production.
	ExitCommit(c *CommitContext)

	// ExitRollback is called when exiting the rollback production.
	ExitRollback(c *RollbackContext)

	// ExitSavepoint is called when exiting the savepoint production.
	ExitSavepoint(c *SavepointContext)

	// ExitBegin is called when exiting the begin production.
	ExitBegin(c *BeginContext)

	// ExitLock is called when exiting the lock production.
	ExitLock(c *LockContext)

	// ExitUnlock is called when exiting the unlock production.
	ExitUnlock(c *UnlockContext)

	// ExitReleaseSavepoint is called when exiting the releaseSavepoint production.
	ExitReleaseSavepoint(c *ReleaseSavepointContext)

	// ExitXa is called when exiting the xa production.
	ExitXa(c *XaContext)

	// ExitOptionChain is called when exiting the optionChain production.
	ExitOptionChain(c *OptionChainContext)

	// ExitOptionRelease is called when exiting the optionRelease production.
	ExitOptionRelease(c *OptionReleaseContext)

	// ExitTableLock is called when exiting the tableLock production.
	ExitTableLock(c *TableLockContext)

	// ExitLockOption is called when exiting the lockOption production.
	ExitLockOption(c *LockOptionContext)

	// ExitXid is called when exiting the xid production.
	ExitXid(c *XidContext)

	// ExitGrantRoleOrPrivilegeTo is called when exiting the grantRoleOrPrivilegeTo production.
	ExitGrantRoleOrPrivilegeTo(c *GrantRoleOrPrivilegeToContext)

	// ExitGrantRoleOrPrivilegeOnTo is called when exiting the grantRoleOrPrivilegeOnTo production.
	ExitGrantRoleOrPrivilegeOnTo(c *GrantRoleOrPrivilegeOnToContext)

	// ExitGrantProxy is called when exiting the grantProxy production.
	ExitGrantProxy(c *GrantProxyContext)

	// ExitRevokeFrom is called when exiting the revokeFrom production.
	ExitRevokeFrom(c *RevokeFromContext)

	// ExitRevokeOnFrom is called when exiting the revokeOnFrom production.
	ExitRevokeOnFrom(c *RevokeOnFromContext)

	// ExitUserList is called when exiting the userList production.
	ExitUserList(c *UserListContext)

	// ExitRoleOrPrivileges is called when exiting the roleOrPrivileges production.
	ExitRoleOrPrivileges(c *RoleOrPrivilegesContext)

	// ExitRoleOrDynamicPrivilege is called when exiting the roleOrDynamicPrivilege production.
	ExitRoleOrDynamicPrivilege(c *RoleOrDynamicPrivilegeContext)

	// ExitRoleAtHost is called when exiting the roleAtHost production.
	ExitRoleAtHost(c *RoleAtHostContext)

	// ExitStaticPrivilegeSelect is called when exiting the staticPrivilegeSelect production.
	ExitStaticPrivilegeSelect(c *StaticPrivilegeSelectContext)

	// ExitStaticPrivilegeInsert is called when exiting the staticPrivilegeInsert production.
	ExitStaticPrivilegeInsert(c *StaticPrivilegeInsertContext)

	// ExitStaticPrivilegeUpdate is called when exiting the staticPrivilegeUpdate production.
	ExitStaticPrivilegeUpdate(c *StaticPrivilegeUpdateContext)

	// ExitStaticPrivilegeReferences is called when exiting the staticPrivilegeReferences production.
	ExitStaticPrivilegeReferences(c *StaticPrivilegeReferencesContext)

	// ExitStaticPrivilegeDelete is called when exiting the staticPrivilegeDelete production.
	ExitStaticPrivilegeDelete(c *StaticPrivilegeDeleteContext)

	// ExitStaticPrivilegeUsage is called when exiting the staticPrivilegeUsage production.
	ExitStaticPrivilegeUsage(c *StaticPrivilegeUsageContext)

	// ExitStaticPrivilegeIndex is called when exiting the staticPrivilegeIndex production.
	ExitStaticPrivilegeIndex(c *StaticPrivilegeIndexContext)

	// ExitStaticPrivilegeAlter is called when exiting the staticPrivilegeAlter production.
	ExitStaticPrivilegeAlter(c *StaticPrivilegeAlterContext)

	// ExitStaticPrivilegeCreate is called when exiting the staticPrivilegeCreate production.
	ExitStaticPrivilegeCreate(c *StaticPrivilegeCreateContext)

	// ExitStaticPrivilegeDrop is called when exiting the staticPrivilegeDrop production.
	ExitStaticPrivilegeDrop(c *StaticPrivilegeDropContext)

	// ExitStaticPrivilegeExecute is called when exiting the staticPrivilegeExecute production.
	ExitStaticPrivilegeExecute(c *StaticPrivilegeExecuteContext)

	// ExitStaticPrivilegeReload is called when exiting the staticPrivilegeReload production.
	ExitStaticPrivilegeReload(c *StaticPrivilegeReloadContext)

	// ExitStaticPrivilegeShutdown is called when exiting the staticPrivilegeShutdown production.
	ExitStaticPrivilegeShutdown(c *StaticPrivilegeShutdownContext)

	// ExitStaticPrivilegeProcess is called when exiting the staticPrivilegeProcess production.
	ExitStaticPrivilegeProcess(c *StaticPrivilegeProcessContext)

	// ExitStaticPrivilegeFile is called when exiting the staticPrivilegeFile production.
	ExitStaticPrivilegeFile(c *StaticPrivilegeFileContext)

	// ExitStaticPrivilegeGrant is called when exiting the staticPrivilegeGrant production.
	ExitStaticPrivilegeGrant(c *StaticPrivilegeGrantContext)

	// ExitStaticPrivilegeShowDatabases is called when exiting the staticPrivilegeShowDatabases production.
	ExitStaticPrivilegeShowDatabases(c *StaticPrivilegeShowDatabasesContext)

	// ExitStaticPrivilegeSuper is called when exiting the staticPrivilegeSuper production.
	ExitStaticPrivilegeSuper(c *StaticPrivilegeSuperContext)

	// ExitStaticPrivilegeCreateTemporaryTables is called when exiting the staticPrivilegeCreateTemporaryTables production.
	ExitStaticPrivilegeCreateTemporaryTables(c *StaticPrivilegeCreateTemporaryTablesContext)

	// ExitStaticPrivilegeLockTables is called when exiting the staticPrivilegeLockTables production.
	ExitStaticPrivilegeLockTables(c *StaticPrivilegeLockTablesContext)

	// ExitStaticPrivilegeReplicationSlave is called when exiting the staticPrivilegeReplicationSlave production.
	ExitStaticPrivilegeReplicationSlave(c *StaticPrivilegeReplicationSlaveContext)

	// ExitStaticPrivilegeReplicationClient is called when exiting the staticPrivilegeReplicationClient production.
	ExitStaticPrivilegeReplicationClient(c *StaticPrivilegeReplicationClientContext)

	// ExitStaticPrivilegeCreateView is called when exiting the staticPrivilegeCreateView production.
	ExitStaticPrivilegeCreateView(c *StaticPrivilegeCreateViewContext)

	// ExitStaticPrivilegeShowView is called when exiting the staticPrivilegeShowView production.
	ExitStaticPrivilegeShowView(c *StaticPrivilegeShowViewContext)

	// ExitStaticPrivilegeCreateRoutine is called when exiting the staticPrivilegeCreateRoutine production.
	ExitStaticPrivilegeCreateRoutine(c *StaticPrivilegeCreateRoutineContext)

	// ExitStaticPrivilegeAlterRoutine is called when exiting the staticPrivilegeAlterRoutine production.
	ExitStaticPrivilegeAlterRoutine(c *StaticPrivilegeAlterRoutineContext)

	// ExitStaticPrivilegeCreateUser is called when exiting the staticPrivilegeCreateUser production.
	ExitStaticPrivilegeCreateUser(c *StaticPrivilegeCreateUserContext)

	// ExitStaticPrivilegeEvent is called when exiting the staticPrivilegeEvent production.
	ExitStaticPrivilegeEvent(c *StaticPrivilegeEventContext)

	// ExitStaticPrivilegeTrigger is called when exiting the staticPrivilegeTrigger production.
	ExitStaticPrivilegeTrigger(c *StaticPrivilegeTriggerContext)

	// ExitStaticPrivilegeCreateTablespace is called when exiting the staticPrivilegeCreateTablespace production.
	ExitStaticPrivilegeCreateTablespace(c *StaticPrivilegeCreateTablespaceContext)

	// ExitStaticPrivilegeCreateRole is called when exiting the staticPrivilegeCreateRole production.
	ExitStaticPrivilegeCreateRole(c *StaticPrivilegeCreateRoleContext)

	// ExitStaticPrivilegeDropRole is called when exiting the staticPrivilegeDropRole production.
	ExitStaticPrivilegeDropRole(c *StaticPrivilegeDropRoleContext)

	// ExitAclType is called when exiting the aclType production.
	ExitAclType(c *AclTypeContext)

	// ExitGrantLevelGlobal is called when exiting the grantLevelGlobal production.
	ExitGrantLevelGlobal(c *GrantLevelGlobalContext)

	// ExitGrantLevelSchemaGlobal is called when exiting the grantLevelSchemaGlobal production.
	ExitGrantLevelSchemaGlobal(c *GrantLevelSchemaGlobalContext)

	// ExitGrantLevelTable is called when exiting the grantLevelTable production.
	ExitGrantLevelTable(c *GrantLevelTableContext)

	// ExitCreateUser is called when exiting the createUser production.
	ExitCreateUser(c *CreateUserContext)

	// ExitCreateUserEntryNoOption is called when exiting the createUserEntryNoOption production.
	ExitCreateUserEntryNoOption(c *CreateUserEntryNoOptionContext)

	// ExitCreateUserEntryIdentifiedBy is called when exiting the createUserEntryIdentifiedBy production.
	ExitCreateUserEntryIdentifiedBy(c *CreateUserEntryIdentifiedByContext)

	// ExitCreateUserEntryIdentifiedWith is called when exiting the createUserEntryIdentifiedWith production.
	ExitCreateUserEntryIdentifiedWith(c *CreateUserEntryIdentifiedWithContext)

	// ExitCreateUserList is called when exiting the createUserList production.
	ExitCreateUserList(c *CreateUserListContext)

	// ExitDefaultRoleClause is called when exiting the defaultRoleClause production.
	ExitDefaultRoleClause(c *DefaultRoleClauseContext)

	// ExitRequireClause is called when exiting the requireClause production.
	ExitRequireClause(c *RequireClauseContext)

	// ExitConnectOptions is called when exiting the connectOptions production.
	ExitConnectOptions(c *ConnectOptionsContext)

	// ExitAccountLockPasswordExpireOptions is called when exiting the accountLockPasswordExpireOptions production.
	ExitAccountLockPasswordExpireOptions(c *AccountLockPasswordExpireOptionsContext)

	// ExitAccountLockPasswordExpireOption is called when exiting the accountLockPasswordExpireOption production.
	ExitAccountLockPasswordExpireOption(c *AccountLockPasswordExpireOptionContext)

	// ExitAlterUser is called when exiting the alterUser production.
	ExitAlterUser(c *AlterUserContext)

	// ExitAlterUserEntry is called when exiting the alterUserEntry production.
	ExitAlterUserEntry(c *AlterUserEntryContext)

	// ExitAlterUserList is called when exiting the alterUserList production.
	ExitAlterUserList(c *AlterUserListContext)

	// ExitDropUser is called when exiting the dropUser production.
	ExitDropUser(c *DropUserContext)

	// ExitCreateRole is called when exiting the createRole production.
	ExitCreateRole(c *CreateRoleContext)

	// ExitDropRole is called when exiting the dropRole production.
	ExitDropRole(c *DropRoleContext)

	// ExitRenameUser is called when exiting the renameUser production.
	ExitRenameUser(c *RenameUserContext)

	// ExitSetDefaultRole is called when exiting the setDefaultRole production.
	ExitSetDefaultRole(c *SetDefaultRoleContext)

	// ExitSetRole is called when exiting the setRole production.
	ExitSetRole(c *SetRoleContext)

	// ExitSetPassword is called when exiting the setPassword production.
	ExitSetPassword(c *SetPasswordContext)

	// ExitAuthOption is called when exiting the authOption production.
	ExitAuthOption(c *AuthOptionContext)

	// ExitWithGrantOption is called when exiting the withGrantOption production.
	ExitWithGrantOption(c *WithGrantOptionContext)

	// ExitUserOrRoles is called when exiting the userOrRoles production.
	ExitUserOrRoles(c *UserOrRolesContext)

	// ExitRoles is called when exiting the roles production.
	ExitRoles(c *RolesContext)

	// ExitGrantAs is called when exiting the grantAs production.
	ExitGrantAs(c *GrantAsContext)

	// ExitWithRoles is called when exiting the withRoles production.
	ExitWithRoles(c *WithRolesContext)

	// ExitUserAuthOption is called when exiting the userAuthOption production.
	ExitUserAuthOption(c *UserAuthOptionContext)

	// ExitIdentifiedBy is called when exiting the identifiedBy production.
	ExitIdentifiedBy(c *IdentifiedByContext)

	// ExitIdentifiedWith is called when exiting the identifiedWith production.
	ExitIdentifiedWith(c *IdentifiedWithContext)

	// ExitConnectOption is called when exiting the connectOption production.
	ExitConnectOption(c *ConnectOptionContext)

	// ExitTlsOption is called when exiting the tlsOption production.
	ExitTlsOption(c *TlsOptionContext)

	// ExitUserFuncAuthOption is called when exiting the userFuncAuthOption production.
	ExitUserFuncAuthOption(c *UserFuncAuthOptionContext)

	// ExitChange is called when exiting the change production.
	ExitChange(c *ChangeContext)

	// ExitChangeMasterTo is called when exiting the changeMasterTo production.
	ExitChangeMasterTo(c *ChangeMasterToContext)

	// ExitChangeReplicationFilter is called when exiting the changeReplicationFilter production.
	ExitChangeReplicationFilter(c *ChangeReplicationFilterContext)

	// ExitStartSlave is called when exiting the startSlave production.
	ExitStartSlave(c *StartSlaveContext)

	// ExitStopSlave is called when exiting the stopSlave production.
	ExitStopSlave(c *StopSlaveContext)

	// ExitGroupReplication is called when exiting the groupReplication production.
	ExitGroupReplication(c *GroupReplicationContext)

	// ExitStartGroupReplication is called when exiting the startGroupReplication production.
	ExitStartGroupReplication(c *StartGroupReplicationContext)

	// ExitStopGroupReplication is called when exiting the stopGroupReplication production.
	ExitStopGroupReplication(c *StopGroupReplicationContext)

	// ExitPurgeBinaryLog is called when exiting the purgeBinaryLog production.
	ExitPurgeBinaryLog(c *PurgeBinaryLogContext)

	// ExitThreadTypes is called when exiting the threadTypes production.
	ExitThreadTypes(c *ThreadTypesContext)

	// ExitThreadType is called when exiting the threadType production.
	ExitThreadType(c *ThreadTypeContext)

	// ExitUtilOption is called when exiting the utilOption production.
	ExitUtilOption(c *UtilOptionContext)

	// ExitConnectionOptions is called when exiting the connectionOptions production.
	ExitConnectionOptions(c *ConnectionOptionsContext)

	// ExitMasterDefs is called when exiting the masterDefs production.
	ExitMasterDefs(c *MasterDefsContext)

	// ExitMasterDef is called when exiting the masterDef production.
	ExitMasterDef(c *MasterDefContext)

	// ExitIgnoreServerIds is called when exiting the ignoreServerIds production.
	ExitIgnoreServerIds(c *IgnoreServerIdsContext)

	// ExitIgnoreServerId is called when exiting the ignoreServerId production.
	ExitIgnoreServerId(c *IgnoreServerIdContext)

	// ExitFilterDefs is called when exiting the filterDefs production.
	ExitFilterDefs(c *FilterDefsContext)

	// ExitFilterDef is called when exiting the filterDef production.
	ExitFilterDef(c *FilterDefContext)

	// ExitWildTables is called when exiting the wildTables production.
	ExitWildTables(c *WildTablesContext)

	// ExitWildTable is called when exiting the wildTable production.
	ExitWildTable(c *WildTableContext)
}

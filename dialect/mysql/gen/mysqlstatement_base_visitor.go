// Code generated from E:/GoProject/src/github.com/2212442929/xsqlparser/dialect/mysql/grammer\MySQLStatement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // MySQLStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseMySQLStatementVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMySQLStatementVisitor) VisitExecute(ctx *ExecuteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitInsert(ctx *InsertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitInsertSpecification(ctx *InsertSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitInsertValuesClause(ctx *InsertValuesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFields(ctx *FieldsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitInsertIdentifier(ctx *InsertIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableWild(ctx *TableWildContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitInsertSelectClause(ctx *InsertSelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOnDuplicateKeyClause(ctx *OnDuplicateKeyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitValueReference(ctx *ValueReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDerivedColumns(ctx *DerivedColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitReplace(ctx *ReplaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitReplaceSpecification(ctx *ReplaceSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitReplaceValuesClause(ctx *ReplaceValuesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitReplaceSelectClause(ctx *ReplaceSelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUpdate(ctx *UpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUpdateSpecification_(ctx *UpdateSpecification_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSetAssignmentsClause(ctx *SetAssignmentsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAssignmentValues(ctx *AssignmentValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAssignmentValue(ctx *AssignmentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitBlobValue(ctx *BlobValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDelete(ctx *DeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDeleteSpecification(ctx *DeleteSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSingleTableClause(ctx *SingleTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitMultipleTablesClause(ctx *MultipleTablesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSelect(ctx *SelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSelectWithInto(ctx *SelectWithIntoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitQueryExpression(ctx *QueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitQueryExpressionBody(ctx *QueryExpressionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUnionClause(ctx *UnionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitQueryExpressionParens(ctx *QueryExpressionParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitQueryPrimary(ctx *QueryPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDoStatement(ctx *DoStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitHandlerStatement(ctx *HandlerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitHandlerOpenStatement(ctx *HandlerOpenStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitHandlerReadIndexStatement(ctx *HandlerReadIndexStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitHandlerReadStatement(ctx *HandlerReadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitHandlerCloseStatement(ctx *HandlerCloseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLoadStatement(ctx *LoadStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLoadDataStatement(ctx *LoadDataStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLoadXmlStatement(ctx *LoadXmlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitExplicitTable(ctx *ExplicitTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableValueConstructor(ctx *TableValueConstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRowConstructorList(ctx *RowConstructorListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCteClause(ctx *CteClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSelectSpecification(ctx *SelectSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDuplicateSpecification(ctx *DuplicateSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitProjections(ctx *ProjectionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitProjection(ctx *ProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitQualifiedShorthand(ctx *QualifiedShorthandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableReferences(ctx *TableReferencesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitEscapedTableReference(ctx *EscapedTableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableReference(ctx *TableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableFactor(ctx *TableFactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPartitionNames(ctx *PartitionNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIndexHintList(ctx *IndexHintListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIndexHint(ctx *IndexHintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitJoinedTable(ctx *JoinedTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitInnerJoinType(ctx *InnerJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOuterJoinType(ctx *OuterJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitNaturalJoinType(ctx *NaturalJoinTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitJoinSpecification(ctx *JoinSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLimitRowCount(ctx *LimitRowCountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLimitOffset(ctx *LimitOffsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWindowClause(ctx *WindowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWindowItem(ctx *WindowItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSelectLinesInto(ctx *SelectLinesIntoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSelectFieldsInto(ctx *SelectFieldsIntoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSelectIntoExpression(ctx *SelectIntoExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLockClause(ctx *LockClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLockClauseList(ctx *LockClauseListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLockStrength(ctx *LockStrengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLockedRowAction(ctx *LockedRowActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableLockingList(ctx *TableLockingListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableIdentOptWild(ctx *TableIdentOptWildContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableAliasRefList(ctx *TableAliasRefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitParameterMarker(ctx *ParameterMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCustomKeyword(ctx *CustomKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLiterals(ctx *LiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitString_(ctx *String_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStringLiterals(ctx *StringLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitNumberLiterals(ctx *NumberLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTemporalLiterals(ctx *TemporalLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitHexadecimalLiterals(ctx *HexadecimalLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitBitValueLiterals(ctx *BitValueLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitBooleanLiterals(ctx *BooleanLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitNullValueLiterals(ctx *NullValueLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCollationName(ctx *CollationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIdentifierKeywordsUnambiguous(ctx *IdentifierKeywordsUnambiguousContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIdentifierKeywordsAmbiguous1RolesAndLabels(ctx *IdentifierKeywordsAmbiguous1RolesAndLabelsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIdentifierKeywordsAmbiguous2Labels(ctx *IdentifierKeywordsAmbiguous2LabelsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIdentifierKeywordsAmbiguous3Roles(ctx *IdentifierKeywordsAmbiguous3RolesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIdentifierKeywordsAmbiguous4SystemVariables(ctx *IdentifierKeywordsAmbiguous4SystemVariablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTextOrIdentifier(ctx *TextOrIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUserVariable(ctx *UserVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSystemVariable(ctx *SystemVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSetSystemVariable(ctx *SetSystemVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOptionType(ctx *OptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitInternalVariableName(ctx *InternalVariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSetExprOrDefault(ctx *SetExprOrDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTransactionCharacteristics(ctx *TransactionCharacteristicsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIsolationLevel(ctx *IsolationLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIsolationTypes(ctx *IsolationTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTransactionAccessMode(ctx *TransactionAccessModeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSchemaName(ctx *SchemaNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSchemaNames(ctx *SchemaNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCharsetName(ctx *CharsetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSchemaPairs(ctx *SchemaPairsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSchemaPair(ctx *SchemaPairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIndexName(ctx *IndexNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConstraintName(ctx *ConstraintNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUserIdentifierOrText(ctx *UserIdentifierOrTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUserName(ctx *UserNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitEventName(ctx *EventNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitServerName(ctx *ServerNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWrapperName(ctx *WrapperNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitViewName(ctx *ViewNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOwner(ctx *OwnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableList(ctx *TableListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitViewNames(ctx *ViewNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitColumnNames(ctx *ColumnNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGroupName(ctx *GroupNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRoutineName(ctx *RoutineNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShardLibraryName(ctx *ShardLibraryNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitComponentName(ctx *ComponentNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPluginName(ctx *PluginNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitHostName(ctx *HostNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPort(ctx *PortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCloneInstance(ctx *CloneInstanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCloneDir(ctx *CloneDirContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitChannelName(ctx *ChannelNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLogName(ctx *LogNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRoleName(ctx *RoleNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRoleIdentifierOrText(ctx *RoleIdentifierOrTextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitEngineRef(ctx *EngineRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTriggerName(ctx *TriggerNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTriggerTime(ctx *TriggerTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableOrTables(ctx *TableOrTablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUserOrRole(ctx *UserOrRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPartitionName(ctx *PartitionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIdentifierList(ctx *IdentifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAllOrPartitionNameList(ctx *AllOrPartitionNameListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTriggerEvent(ctx *TriggerEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTriggerOrder(ctx *TriggerOrderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAndOperator(ctx *AndOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOrOperator(ctx *OrOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitNotOperator(ctx *NotOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitBooleanPrimary(ctx *BooleanPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAssignmentOperator(ctx *AssignmentOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitBitExpr(ctx *BitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSimpleExpr(ctx *SimpleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitColumnRef(ctx *ColumnRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitColumnRefList(ctx *ColumnRefListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAggregationFunction(ctx *AggregationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAggregationFunctionName(ctx *AggregationFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDistinct(ctx *DistinctContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOverClause(ctx *OverClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWindowSpecification(ctx *WindowSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFrameClause(ctx *FrameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFrameStart(ctx *FrameStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFrameEnd(ctx *FrameEndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFrameBetween(ctx *FrameBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSpecialFunction(ctx *SpecialFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCurrentUserFunction(ctx *CurrentUserFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGroupConcatFunction(ctx *GroupConcatFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWindowFunction(ctx *WindowFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWindowingClause(ctx *WindowingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLeadLagInfo(ctx *LeadLagInfoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitNullTreatment(ctx *NullTreatmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCheckType(ctx *CheckTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRepairType(ctx *RepairTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCastFunction(ctx *CastFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConvertFunction(ctx *ConvertFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCastType(ctx *CastTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitNchar(ctx *NcharContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPositionFunction(ctx *PositionFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSubstringFunction(ctx *SubstringFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitExtractFunction(ctx *ExtractFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCharFunction(ctx *CharFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTrimFunction(ctx *TrimFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitValuesFunction(ctx *ValuesFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWeightStringFunction(ctx *WeightStringFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLevelClause(ctx *LevelClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLevelInWeightListElement(ctx *LevelInWeightListElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRegularFunction(ctx *RegularFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShorthandRegularFunction(ctx *ShorthandRegularFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCompleteRegularFunction(ctx *CompleteRegularFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRegularFunctionName(ctx *RegularFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitMatchExpression(ctx *MatchExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitMatchSearchModifier(ctx *MatchSearchModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCaseExpression(ctx *CaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDatetimeExpr(ctx *DatetimeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitBinaryLogFileIndexNumber(ctx *BinaryLogFileIndexNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCaseWhen(ctx *CaseWhenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCaseElse(ctx *CaseElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIntervalExpression(ctx *IntervalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIntervalValue(ctx *IntervalValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIntervalUnit(ctx *IntervalUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOrderByItem(ctx *OrderByItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStringList(ctx *StringListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTextString(ctx *TextStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTextStringHash(ctx *TextStringHashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFieldOptions(ctx *FieldOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPrecision(ctx *PrecisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTypeDatetimePrecision(ctx *TypeDatetimePrecisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCharsetWithOptBinary(ctx *CharsetWithOptBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAscii(ctx *AsciiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUnicode(ctx *UnicodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCharset(ctx *CharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDefaultCollation(ctx *DefaultCollationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDefaultEncryption(ctx *DefaultEncryptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDefaultCharset(ctx *DefaultCharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSignedLiteral(ctx *SignedLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitNow(ctx *NowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitColumnFormat(ctx *ColumnFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStorageMedia(ctx *StorageMediaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDirection(ctx *DirectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitKeyOrIndex(ctx *KeyOrIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFieldLength(ctx *FieldLengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCharacterSet(ctx *CharacterSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCollateClause(ctx *CollateClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFieldOrVarSpec(ctx *FieldOrVarSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitNotExistClause(ctx *NotExistClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitExistClause(ctx *ExistClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConnectionId(ctx *ConnectionIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLabelName(ctx *LabelNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCursorName(ctx *CursorNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConditionName(ctx *ConditionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUnionOption(ctx *UnionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitNoWriteToBinLog(ctx *NoWriteToBinLogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitChannelOption(ctx *ChannelOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPreparedStatement(ctx *PreparedStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitExecuteStatement(ctx *ExecuteStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitExecuteVarList(ctx *ExecuteVarListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterStatement(ctx *AlterStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateTable(ctx *CreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPartitionClause(ctx *PartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPartitionTypeDef(ctx *PartitionTypeDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSubPartitions(ctx *SubPartitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPartitionKeyAlgorithm(ctx *PartitionKeyAlgorithmContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDuplicateAsQueryExpression(ctx *DuplicateAsQueryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterTable(ctx *AlterTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStandaloneAlterTableAction(ctx *StandaloneAlterTableActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterTableActions(ctx *AlterTableActionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterTablePartitionOptions(ctx *AlterTablePartitionOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterCommandList(ctx *AlterCommandListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterList(ctx *AlterListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateTableOptionsSpaceSeparated(ctx *CreateTableOptionsSpaceSeparatedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAddColumn(ctx *AddColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAddTableConstraint(ctx *AddTableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitChangeColumn(ctx *ChangeColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitModifyColumn(ctx *ModifyColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterTableDrop(ctx *AlterTableDropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDisableKeys(ctx *DisableKeysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitEnableKeys(ctx *EnableKeysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterColumn(ctx *AlterColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterIndex(ctx *AlterIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterCheck(ctx *AlterCheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterConstraint(ctx *AlterConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRenameColumn(ctx *RenameColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterRenameTable(ctx *AlterRenameTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRenameIndex(ctx *RenameIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterConvert(ctx *AlterConvertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterTableForce(ctx *AlterTableForceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterTableOrder(ctx *AlterTableOrderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterOrderList(ctx *AlterOrderListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableConstraintDef(ctx *TableConstraintDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterCommandsModifierList(ctx *AlterCommandsModifierListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterCommandsModifier(ctx *AlterCommandsModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWithValidation(ctx *WithValidationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStandaloneAlterCommands(ctx *StandaloneAlterCommandsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterPartition(ctx *AlterPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConstraintClause(ctx *ConstraintClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableElementList(ctx *TableElementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableElement(ctx *TableElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRestrict(ctx *RestrictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFulltextIndexOption(ctx *FulltextIndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropTable(ctx *DropTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropIndex(ctx *DropIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterAlgorithmOption(ctx *AlterAlgorithmOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterLockOption(ctx *AlterLockOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTruncateTable(ctx *TruncateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateIndex(ctx *CreateIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateDatabase(ctx *CreateDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterDatabase(ctx *AlterDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateDatabaseSpecification_(ctx *CreateDatabaseSpecification_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterDatabaseSpecification_(ctx *AlterDatabaseSpecification_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropDatabase(ctx *DropDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterInstance(ctx *AlterInstanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitInstanceAction(ctx *InstanceActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitChannel(ctx *ChannelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateEvent(ctx *CreateEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterEvent(ctx *AlterEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropEvent(ctx *DropEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateFunction(ctx *CreateFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterFunction(ctx *AlterFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropFunction(ctx *DropFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateProcedure(ctx *CreateProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterProcedure(ctx *AlterProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropProcedure(ctx *DropProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateServer(ctx *CreateServerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterServer(ctx *AlterServerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropServer(ctx *DropServerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateView(ctx *CreateViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterView(ctx *AlterViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropView(ctx *DropViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateTablespace(ctx *CreateTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateTablespaceInnodb(ctx *CreateTablespaceInnodbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateTablespaceNdb(ctx *CreateTablespaceNdbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterTablespace(ctx *AlterTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterTablespaceNdb(ctx *AlterTablespaceNdbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterTablespaceInnodb(ctx *AlterTablespaceInnodbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropTablespace(ctx *DropTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateLogfileGroup(ctx *CreateLogfileGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterLogfileGroup(ctx *AlterLogfileGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropLogfileGroup(ctx *DropLogfileGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateTrigger(ctx *CreateTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropTrigger(ctx *DropTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRenameTable(ctx *RenameTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateDefinitionClause(ctx *CreateDefinitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFieldDefinition(ctx *FieldDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitColumnAttribute(ctx *ColumnAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCheckConstraint(ctx *CheckConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConstraintEnforcement(ctx *ConstraintEnforcementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGeneratedOption(ctx *GeneratedOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitReferenceDefinition(ctx *ReferenceDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOnUpdateDelete(ctx *OnUpdateDeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitReferenceOption(ctx *ReferenceOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIndexNameAndType(ctx *IndexNameAndTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIndexType(ctx *IndexTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIndexTypeClause(ctx *IndexTypeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitKeyParts(ctx *KeyPartsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitKeyPart(ctx *KeyPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitKeyPartWithExpression(ctx *KeyPartWithExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitKeyListWithExpression(ctx *KeyListWithExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIndexOption(ctx *IndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCommonIndexOption(ctx *CommonIndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitVisibility(ctx *VisibilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateLikeClause(ctx *CreateLikeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateIndexSpecification(ctx *CreateIndexSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateTableOptions(ctx *CreateTableOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateTableOption(ctx *CreateTableOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateSRSStatement(ctx *CreateSRSStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropSRSStatement(ctx *DropSRSStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSrsAttribute(ctx *SrsAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPlace(ctx *PlaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPartitionDefinitions(ctx *PartitionDefinitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPartitionDefinition(ctx *PartitionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPartitionLessThanValue(ctx *PartitionLessThanValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPartitionValueList(ctx *PartitionValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPartitionDefinitionOption(ctx *PartitionDefinitionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOwnerStatement(ctx *OwnerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitScheduleExpression(ctx *ScheduleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTimestampValue(ctx *TimestampValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRoutineBody(ctx *RoutineBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitServerOption(ctx *ServerOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRoutineOption(ctx *RoutineOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitProcedureParameter(ctx *ProcedureParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFileSizeLiteral(ctx *FileSizeLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSimpleStatement(ctx *SimpleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCompoundStatement(ctx *CompoundStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitValidStatement(ctx *ValidStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitBeginStatement(ctx *BeginStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDeclareStatement(ctx *DeclareStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFlowControlStatement(ctx *FlowControlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCaseStatement(ctx *CaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIterateStatement(ctx *IterateStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLeaveStatement(ctx *LeaveStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRepeatStatement(ctx *RepeatStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCursorStatement(ctx *CursorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCursorCloseStatement(ctx *CursorCloseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCursorDeclareStatement(ctx *CursorDeclareStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCursorFetchStatement(ctx *CursorFetchStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCursorOpenStatement(ctx *CursorOpenStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConditionHandlingStatement(ctx *ConditionHandlingStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDeclareConditionStatement(ctx *DeclareConditionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDeclareHandlerStatement(ctx *DeclareHandlerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGetDiagnosticsStatement(ctx *GetDiagnosticsStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStatementInformationItem(ctx *StatementInformationItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConditionInformationItem(ctx *ConditionInformationItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConditionNumber(ctx *ConditionNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStatementInformationItemName(ctx *StatementInformationItemNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConditionInformationItemName(ctx *ConditionInformationItemNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitHandlerAction(ctx *HandlerActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConditionValue(ctx *ConditionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitResignalStatement(ctx *ResignalStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSignalStatement(ctx *SignalStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSignalInformationItem(ctx *SignalInformationItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUse(ctx *UseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitHelp(ctx *HelpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitExplain(ctx *ExplainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowDatabases(ctx *ShowDatabasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowTables(ctx *ShowTablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowTableStatus(ctx *ShowTableStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowColumns(ctx *ShowColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowIndex(ctx *ShowIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowCreateTable(ctx *ShowCreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFromSchema(ctx *FromSchemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFromTable(ctx *FromTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowLike(ctx *ShowLikeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowWhereClause(ctx *ShowWhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowFilter(ctx *ShowFilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowProfileType(ctx *ShowProfileTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSetVariable(ctx *SetVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOptionValueList(ctx *OptionValueListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOptionValueNoOptionType(ctx *OptionValueNoOptionTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOptionValue(ctx *OptionValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowBinaryLogs(ctx *ShowBinaryLogsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowBinlogEvents(ctx *ShowBinlogEventsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowCharacterSet(ctx *ShowCharacterSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowCollation(ctx *ShowCollationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowCreateDatabase(ctx *ShowCreateDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowCreateEvent(ctx *ShowCreateEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowCreateFunction(ctx *ShowCreateFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowCreateProcedure(ctx *ShowCreateProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowCreateTrigger(ctx *ShowCreateTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowCreateUser(ctx *ShowCreateUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowCreateView(ctx *ShowCreateViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowEngine(ctx *ShowEngineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowEngines(ctx *ShowEnginesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowCharset(ctx *ShowCharsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowErrors(ctx *ShowErrorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowEvents(ctx *ShowEventsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowFunctionCode(ctx *ShowFunctionCodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowFunctionStatus(ctx *ShowFunctionStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowGrant(ctx *ShowGrantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowMasterStatus(ctx *ShowMasterStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowOpenTables(ctx *ShowOpenTablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowPlugins(ctx *ShowPluginsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowPrivileges(ctx *ShowPrivilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowProcedureCode(ctx *ShowProcedureCodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowProcedureStatus(ctx *ShowProcedureStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowProcesslist(ctx *ShowProcesslistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowProfile(ctx *ShowProfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowProfiles(ctx *ShowProfilesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowRelaylogEvent(ctx *ShowRelaylogEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowSlavehost(ctx *ShowSlavehostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowSlaveStatus(ctx *ShowSlaveStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowStatus(ctx *ShowStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowTriggers(ctx *ShowTriggersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowVariables(ctx *ShowVariablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowWarnings(ctx *ShowWarningsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowReplicas(ctx *ShowReplicasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShowReplicaStatus(ctx *ShowReplicaStatusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSetCharacter(ctx *SetCharacterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitClone(ctx *CloneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCloneAction(ctx *CloneActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateLoadableFunction(ctx *CreateLoadableFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitInstall(ctx *InstallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUninstall(ctx *UninstallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitInstallComponent(ctx *InstallComponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitInstallPlugin(ctx *InstallPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUninstallComponent(ctx *UninstallComponentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUninstallPlugin(ctx *UninstallPluginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAnalyzeTable(ctx *AnalyzeTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitHistogram(ctx *HistogramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCheckTable(ctx *CheckTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCheckTableOption(ctx *CheckTableOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitChecksumTable(ctx *ChecksumTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOptimizeTable(ctx *OptimizeTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRepairTable(ctx *RepairTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterResourceGroup(ctx *AlterResourceGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitVcpuSpec(ctx *VcpuSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateResourceGroup(ctx *CreateResourceGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropResourceGroup(ctx *DropResourceGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSetResourceGroup(ctx *SetResourceGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitBinlog(ctx *BinlogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCacheIndex(ctx *CacheIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCacheTableIndexList(ctx *CacheTableIndexListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPartitionList(ctx *PartitionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFlush(ctx *FlushContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFlushOption(ctx *FlushOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTablesOption(ctx *TablesOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitKill(ctx *KillContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLoadIndexInfo(ctx *LoadIndexInfoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLoadTableIndexList(ctx *LoadTableIndexListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitResetStatement(ctx *ResetStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitResetOption(ctx *ResetOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitResetPersist(ctx *ResetPersistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRestart(ctx *RestartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShutdown(ctx *ShutdownContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitExplainType(ctx *ExplainTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitExplainableStatement(ctx *ExplainableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFormatName(ctx *FormatNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitShow(ctx *ShowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSetTransaction(ctx *SetTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSetAutoCommit(ctx *SetAutoCommitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitBeginTransaction(ctx *BeginTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTransactionCharacteristic(ctx *TransactionCharacteristicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCommit(ctx *CommitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRollback(ctx *RollbackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSavepoint(ctx *SavepointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitBegin(ctx *BeginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLock(ctx *LockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUnlock(ctx *UnlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitReleaseSavepoint(ctx *ReleaseSavepointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitXa(ctx *XaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOptionChain(ctx *OptionChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitOptionRelease(ctx *OptionReleaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTableLock(ctx *TableLockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitLockOption(ctx *LockOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitXid(ctx *XidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGrantRoleOrPrivilegeTo(ctx *GrantRoleOrPrivilegeToContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGrantRoleOrPrivilegeOnTo(ctx *GrantRoleOrPrivilegeOnToContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGrantProxy(ctx *GrantProxyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRevokeFrom(ctx *RevokeFromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRevokeOnFrom(ctx *RevokeOnFromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUserList(ctx *UserListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRoleOrPrivileges(ctx *RoleOrPrivilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRoleOrDynamicPrivilege(ctx *RoleOrDynamicPrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRoleAtHost(ctx *RoleAtHostContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeSelect(ctx *StaticPrivilegeSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeInsert(ctx *StaticPrivilegeInsertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeUpdate(ctx *StaticPrivilegeUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeReferences(ctx *StaticPrivilegeReferencesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeDelete(ctx *StaticPrivilegeDeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeUsage(ctx *StaticPrivilegeUsageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeIndex(ctx *StaticPrivilegeIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeAlter(ctx *StaticPrivilegeAlterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeCreate(ctx *StaticPrivilegeCreateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeDrop(ctx *StaticPrivilegeDropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeExecute(ctx *StaticPrivilegeExecuteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeReload(ctx *StaticPrivilegeReloadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeShutdown(ctx *StaticPrivilegeShutdownContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeProcess(ctx *StaticPrivilegeProcessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeFile(ctx *StaticPrivilegeFileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeGrant(ctx *StaticPrivilegeGrantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeShowDatabases(ctx *StaticPrivilegeShowDatabasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeSuper(ctx *StaticPrivilegeSuperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeCreateTemporaryTables(ctx *StaticPrivilegeCreateTemporaryTablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeLockTables(ctx *StaticPrivilegeLockTablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeReplicationSlave(ctx *StaticPrivilegeReplicationSlaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeReplicationClient(ctx *StaticPrivilegeReplicationClientContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeCreateView(ctx *StaticPrivilegeCreateViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeShowView(ctx *StaticPrivilegeShowViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeCreateRoutine(ctx *StaticPrivilegeCreateRoutineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeAlterRoutine(ctx *StaticPrivilegeAlterRoutineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeCreateUser(ctx *StaticPrivilegeCreateUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeEvent(ctx *StaticPrivilegeEventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeTrigger(ctx *StaticPrivilegeTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeCreateTablespace(ctx *StaticPrivilegeCreateTablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeCreateRole(ctx *StaticPrivilegeCreateRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStaticPrivilegeDropRole(ctx *StaticPrivilegeDropRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAclType(ctx *AclTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGrantLevelGlobal(ctx *GrantLevelGlobalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGrantLevelSchemaGlobal(ctx *GrantLevelSchemaGlobalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGrantLevelTable(ctx *GrantLevelTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateUser(ctx *CreateUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateUserEntryNoOption(ctx *CreateUserEntryNoOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateUserEntryIdentifiedBy(ctx *CreateUserEntryIdentifiedByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateUserEntryIdentifiedWith(ctx *CreateUserEntryIdentifiedWithContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateUserList(ctx *CreateUserListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDefaultRoleClause(ctx *DefaultRoleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRequireClause(ctx *RequireClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConnectOptions(ctx *ConnectOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAccountLockPasswordExpireOptions(ctx *AccountLockPasswordExpireOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAccountLockPasswordExpireOption(ctx *AccountLockPasswordExpireOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterUser(ctx *AlterUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterUserEntry(ctx *AlterUserEntryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAlterUserList(ctx *AlterUserListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropUser(ctx *DropUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitCreateRole(ctx *CreateRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitDropRole(ctx *DropRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRenameUser(ctx *RenameUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSetDefaultRole(ctx *SetDefaultRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSetRole(ctx *SetRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitSetPassword(ctx *SetPasswordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitAuthOption(ctx *AuthOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWithGrantOption(ctx *WithGrantOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUserOrRoles(ctx *UserOrRolesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitRoles(ctx *RolesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGrantAs(ctx *GrantAsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWithRoles(ctx *WithRolesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUserAuthOption(ctx *UserAuthOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIdentifiedBy(ctx *IdentifiedByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIdentifiedWith(ctx *IdentifiedWithContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConnectOption(ctx *ConnectOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitTlsOption(ctx *TlsOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUserFuncAuthOption(ctx *UserFuncAuthOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitChange(ctx *ChangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitChangeMasterTo(ctx *ChangeMasterToContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitChangeReplicationFilter(ctx *ChangeReplicationFilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStartSlave(ctx *StartSlaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStopSlave(ctx *StopSlaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitGroupReplication(ctx *GroupReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStartGroupReplication(ctx *StartGroupReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitStopGroupReplication(ctx *StopGroupReplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitPurgeBinaryLog(ctx *PurgeBinaryLogContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitThreadTypes(ctx *ThreadTypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitThreadType(ctx *ThreadTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitUtilOption(ctx *UtilOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitConnectionOptions(ctx *ConnectionOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitMasterDefs(ctx *MasterDefsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitMasterDef(ctx *MasterDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIgnoreServerIds(ctx *IgnoreServerIdsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitIgnoreServerId(ctx *IgnoreServerIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFilterDefs(ctx *FilterDefsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitFilterDef(ctx *FilterDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWildTables(ctx *WildTablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMySQLStatementVisitor) VisitWildTable(ctx *WildTableContext) interface{} {
	return v.VisitChildren(ctx)
}

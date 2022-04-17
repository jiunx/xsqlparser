// Code generated from E:/GoProject/src/github.com/2212442929/xsqlparser/dialect/sqlserver/grammer\SQLServerStatement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQLServerStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSQLServerStatementVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSQLServerStatementVisitor) VisitExecute(ctx *ExecuteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitInsert(ctx *InsertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitInsertDefaultValue(ctx *InsertDefaultValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitInsertValuesClause(ctx *InsertValuesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitInsertSelectClause(ctx *InsertSelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitUpdate(ctx *UpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSetAssignmentsClause(ctx *SetAssignmentsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAssignmentValues(ctx *AssignmentValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAssignmentValue(ctx *AssignmentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDelete(ctx *DeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSingleTableClause(ctx *SingleTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitMultipleTablesClause(ctx *MultipleTablesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitMultipleTableNames(ctx *MultipleTableNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSelect(ctx *SelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAggregationClause(ctx *AggregationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSelectClause(ctx *SelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDuplicateSpecification(ctx *DuplicateSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitProjections(ctx *ProjectionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitProjection(ctx *ProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTop(ctx *TopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTopNum(ctx *TopNumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitQualifiedShorthand(ctx *QualifiedShorthandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableReferences(ctx *TableReferencesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableReference(ctx *TableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableFactor(ctx *TableFactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitJoinedTable(ctx *JoinedTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitJoinSpecification(ctx *JoinSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCteClause(ctx *CteClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOutputClause(ctx *OutputClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOutputWithColumns(ctx *OutputWithColumnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOutputWithColumn(ctx *OutputWithColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOutputWithAaterisk(ctx *OutputWithAateriskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOutputTableName(ctx *OutputTableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitQueryHint(ctx *QueryHintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitUseHitName(ctx *UseHitNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitParameterMarker(ctx *ParameterMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitLiterals(ctx *LiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitStringLiterals(ctx *StringLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitNumberLiterals(ctx *NumberLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDateTimeLiterals(ctx *DateTimeLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitHexadecimalLiterals(ctx *HexadecimalLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitBitValueLiterals(ctx *BitValueLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitBooleanLiterals(ctx *BooleanLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitNullValueLiterals(ctx *NullValueLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitUnreservedWord(ctx *UnreservedWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDatabaseName(ctx *DatabaseNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSchemaName(ctx *SchemaNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFunctionName(ctx *FunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitProcedureName(ctx *ProcedureNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitViewName(ctx *ViewNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTriggerName(ctx *TriggerNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSequenceName(ctx *SequenceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitQueueName(ctx *QueueNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitContractName(ctx *ContractNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitServiceName(ctx *ServiceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOwner(ctx *OwnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitColumnNames(ctx *ColumnNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitColumnNamesWithSort(ctx *ColumnNamesWithSortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableNames(ctx *TableNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitIndexName(ctx *IndexNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitConstraintName(ctx *ConstraintNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCollationName(ctx *CollationNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDataTypeLength(ctx *DataTypeLengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPrimaryKey(ctx *PrimaryKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAndOperator(ctx *AndOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOrOperator(ctx *OrOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitNotOperator(ctx *NotOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitBooleanPrimary(ctx *BooleanPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitBitExpr(ctx *BitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSimpleExpr(ctx *SimpleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAggregationFunction(ctx *AggregationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAggregationFunctionName(ctx *AggregationFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDistinct(ctx *DistinctContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSpecialFunction(ctx *SpecialFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCastFunction(ctx *CastFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCharFunction(ctx *CharFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitRegularFunction(ctx *RegularFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitRegularFunctionName(ctx *RegularFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCaseExpression(ctx *CaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCaseWhen(ctx *CaseWhenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCaseElse(ctx *CaseElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPrivateExprOfDb(ctx *PrivateExprOfDbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOrderByItem(ctx *OrderByItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDataTypeName(ctx *DataTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAtTimeZoneExpr(ctx *AtTimeZoneExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCastExpr(ctx *CastExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitConvertExpr(ctx *ConvertExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWindowedFunction(ctx *WindowedFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOverClause(ctx *OverClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPartitionByClause(ctx *PartitionByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitRowRangeClause(ctx *RowRangeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWindowFrameExtent(ctx *WindowFrameExtentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWindowFrameBetween(ctx *WindowFrameBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWindowFrameBound(ctx *WindowFrameBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWindowFramePreceding(ctx *WindowFramePrecedingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWindowFrameFollowing(ctx *WindowFrameFollowingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitColumnNameWithSort(ctx *ColumnNameWithSortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitIndexOption(ctx *IndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCompressionOption(ctx *CompressionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitEqTime(ctx *EqTimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitEqOnOffOption(ctx *EqOnOffOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitEqKey(ctx *EqKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitEqOnOff(ctx *EqOnOffContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOnPartitionClause(ctx *OnPartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPartitionExpressions(ctx *PartitionExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPartitionExpression(ctx *PartitionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitNumberRange(ctx *NumberRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitLowPriorityLockWait(ctx *LowPriorityLockWaitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOnLowPriorLockWait(ctx *OnLowPriorLockWaitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitIgnoredIdentifier(ctx *IgnoredIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitIgnoredIdentifiers(ctx *IgnoredIdentifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitMatchNone(ctx *MatchNoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitVariableName(ctx *VariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitExecuteAsClause(ctx *ExecuteAsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTransactionName(ctx *TransactionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTransactionVariableName(ctx *TransactionVariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSavepointName(ctx *SavepointNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSavepointVariableName(ctx *SavepointVariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitEntityType(ctx *EntityTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateTable(ctx *CreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateTableClause(ctx *CreateTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateIndex(ctx *CreateIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateDatabase(ctx *CreateDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateFunction(ctx *CreateFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateProcedure(ctx *CreateProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateView(ctx *CreateViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateTrigger(ctx *CreateTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateSequence(ctx *CreateSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateService(ctx *CreateServiceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateSchema(ctx *CreateSchemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterTable(ctx *AlterTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterIndex(ctx *AlterIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterDatabase(ctx *AlterDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterProcedure(ctx *AlterProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterFunction(ctx *AlterFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterView(ctx *AlterViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterTrigger(ctx *AlterTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterSequence(ctx *AlterSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterService(ctx *AlterServiceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterSchema(ctx *AlterSchemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropTable(ctx *DropTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropIndex(ctx *DropIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropDatabase(ctx *DropDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropFunction(ctx *DropFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropProcedure(ctx *DropProcedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropView(ctx *DropViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropTrigger(ctx *DropTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropSequence(ctx *DropSequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropService(ctx *DropServiceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropSchema(ctx *DropSchemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTruncateTable(ctx *TruncateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFileTableClause(ctx *FileTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateDefinitionClause(ctx *CreateDefinitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateTableDefinitions(ctx *CreateTableDefinitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateTableDefinition(ctx *CreateTableDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitColumnDefinitionOption(ctx *ColumnDefinitionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitEncryptedOptions(ctx *EncryptedOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitColumnConstraint(ctx *ColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitComputedColumnConstraint(ctx *ComputedColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitComputedColumnForeignKeyConstraint(ctx *ComputedColumnForeignKeyConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitComputedColumnForeignKeyOnAction(ctx *ComputedColumnForeignKeyOnActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPrimaryKeyConstraint(ctx *PrimaryKeyConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDiskTablePrimaryKeyConstraintOption(ctx *DiskTablePrimaryKeyConstraintOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitClusterOption(ctx *ClusterOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPrimaryKeyWithClause(ctx *PrimaryKeyWithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPrimaryKeyOnClause(ctx *PrimaryKeyOnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOnSchemaColumn(ctx *OnSchemaColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOnFileGroup(ctx *OnFileGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOnString(ctx *OnStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitMemoryTablePrimaryKeyConstraintOption(ctx *MemoryTablePrimaryKeyConstraintOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWithBucket(ctx *WithBucketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitColumnForeignKeyConstraint(ctx *ColumnForeignKeyConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitForeignKeyOnAction(ctx *ForeignKeyOnActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitForeignKeyOn(ctx *ForeignKeyOnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCheckConstraint(ctx *CheckConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitColumnIndex(ctx *ColumnIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWithIndexOption(ctx *WithIndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitIndexOnClause(ctx *IndexOnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOnDefault(ctx *OnDefaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFileStreamOn(ctx *FileStreamOnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitColumnConstraints(ctx *ColumnConstraintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitComputedColumnDefinition(ctx *ComputedColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitColumnSetDefinition(ctx *ColumnSetDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableConstraint(ctx *TableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTablePrimaryConstraint(ctx *TablePrimaryConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPrimaryKeyUnique(ctx *PrimaryKeyUniqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDiskTablePrimaryConstraintOption(ctx *DiskTablePrimaryConstraintOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitMemoryTablePrimaryConstraintOption(ctx *MemoryTablePrimaryConstraintOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitHashWithBucket(ctx *HashWithBucketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableForeignKeyConstraint(ctx *TableForeignKeyConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableIndex(ctx *TableIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitIndexNameOption(ctx *IndexNameOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitIndexOptions(ctx *IndexOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPeriodClause(ctx *PeriodClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPartitionScheme(ctx *PartitionSchemeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFileGroup(ctx *FileGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableOptions(ctx *TableOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableOption(ctx *TableOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDataDelectionOption(ctx *DataDelectionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableStretchOptions(ctx *TableStretchOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableStretchOption(ctx *TableStretchOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitMigrationState_(ctx *MigrationState_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableOperationOption(ctx *TableOperationOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDistributionOption(ctx *DistributionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDataWareHouseTableOption(ctx *DataWareHouseTableOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDataWareHousePartitionOption(ctx *DataWareHousePartitionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateIndexSpecification(ctx *CreateIndexSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterDefinitionClause(ctx *AlterDefinitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAddColumnSpecification(ctx *AddColumnSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitModifyColumnSpecification(ctx *ModifyColumnSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterColumnOperation(ctx *AlterColumnOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterColumnAddOptions(ctx *AlterColumnAddOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterColumnAddOption(ctx *AlterColumnAddOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitConstraintForColumn(ctx *ConstraintForColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitGeneratedColumnNamesClause(ctx *GeneratedColumnNamesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitGeneratedColumnNameClause(ctx *GeneratedColumnNameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitGeneratedColumnName(ctx *GeneratedColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterDrop(ctx *AlterDropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterTableDropConstraint(ctx *AlterTableDropConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropConstraintName(ctx *DropConstraintNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropConstraintWithClause(ctx *DropConstraintWithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropConstraintOption(ctx *DropConstraintOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOnOffOption(ctx *OnOffOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropColumnSpecification(ctx *DropColumnSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropIndexSpecification(ctx *DropIndexSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterCheckConstraint(ctx *AlterCheckConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterTableTrigger(ctx *AlterTableTriggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterSwitch(ctx *AlterSwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterSet(ctx *AlterSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSetFileStreamClause(ctx *SetFileStreamClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSetSystemVersionClause(ctx *SetSystemVersionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterSetOnClause(ctx *AlterSetOnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDataConsistencyCheckClause(ctx *DataConsistencyCheckClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitHistoryRetentionPeriodClause(ctx *HistoryRetentionPeriodClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitHistoryRetentionPeriod(ctx *HistoryRetentionPeriodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterTableTableIndex(ctx *AlterTableTableIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitIndexWithName(ctx *IndexWithNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitIndexNonClusterClause(ctx *IndexNonClusterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterTableIndexOnClause(ctx *AlterTableIndexOnClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitIndexClusterClause(ctx *IndexClusterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterTableOption(ctx *AlterTableOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOnHistoryTableClause(ctx *OnHistoryTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitIfExist(ctx *IfExistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateDatabaseClause(ctx *CreateDatabaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFileDefinitionClause(ctx *FileDefinitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDatabaseOption(ctx *DatabaseOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFileStreamOption(ctx *FileStreamOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFileSpec(ctx *FileSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDatabaseFileSpecOption(ctx *DatabaseFileSpecOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDatabaseFileGroup(ctx *DatabaseFileGroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDatabaseFileGroupContains(ctx *DatabaseFileGroupContainsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDatabaseLogOns(ctx *DatabaseLogOnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDeclareVariable(ctx *DeclareVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableVariable(ctx *TableVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitVariTableTypeDefinition(ctx *VariTableTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableVariableClause(ctx *TableVariableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitVariableTableColumnDefinition(ctx *VariableTableColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitVariableTableColumnConstraint(ctx *VariableTableColumnConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitVariableTableConstraint(ctx *VariableTableConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSetVariable(ctx *SetVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSetVariableClause(ctx *SetVariableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCursorVariable(ctx *CursorVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCursorClause(ctx *CursorClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCompoundOperation(ctx *CompoundOperationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFuncParameters(ctx *FuncParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFuncReturns(ctx *FuncReturnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFuncMutiReturn(ctx *FuncMutiReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFuncInlineReturn(ctx *FuncInlineReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFuncScalarReturn(ctx *FuncScalarReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTableTypeDefinition(ctx *TableTypeDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCompoundStatement(ctx *CompoundStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFunctionOption(ctx *FunctionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitValidStatement(ctx *ValidStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitProcParameters(ctx *ProcParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitProcParameter(ctx *ProcParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateOrAlterProcClause(ctx *CreateOrAlterProcClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWithCreateProcOption(ctx *WithCreateProcOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitProcOption(ctx *ProcOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitProcAsClause(ctx *ProcAsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitProcSetOption(ctx *ProcSetOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateOrAlterViewClause(ctx *CreateOrAlterViewClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitViewAttribute(ctx *ViewAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWithCommonTableExpr(ctx *WithCommonTableExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCommonTableExpr(ctx *CommonTableExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateTriggerClause(ctx *CreateTriggerClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDmlTriggerOption(ctx *DmlTriggerOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitMethodSpecifier(ctx *MethodSpecifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTriggerTarget(ctx *TriggerTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateOrAlterSequenceClause(ctx *CreateOrAlterSequenceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateIndexClause(ctx *CreateIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFilterPredicate(ctx *FilterPredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitConjunct(ctx *ConjunctContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterIndexClause(ctx *AlterIndexClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitRelationalIndexOption(ctx *RelationalIndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPartitionNumberRange(ctx *PartitionNumberRangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitReorganizeOption(ctx *ReorganizeOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSetIndexOption(ctx *SetIndexOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitResumableIndexOptions(ctx *ResumableIndexOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterDatabaseClause(ctx *AlterDatabaseClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAddSecondaryOption(ctx *AddSecondaryOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitEditionOptions(ctx *EditionOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitServiceObjective(ctx *ServiceObjectiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterDatabaseOptionSpec(ctx *AlterDatabaseOptionSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFileAndFilegroupOptions(ctx *FileAndFilegroupOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAddOrModifyFilegroups(ctx *AddOrModifyFilegroupsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFilegroupUpdatabilityOption(ctx *FilegroupUpdatabilityOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAddOrModifyFiles(ctx *AddOrModifyFilesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAcceleratedDatabaseRecovery(ctx *AcceleratedDatabaseRecoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAutoOption(ctx *AutoOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAutomaticTuningOption(ctx *AutomaticTuningOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitChangeTrackingOption(ctx *ChangeTrackingOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitChangeTrackingOptionList(ctx *ChangeTrackingOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCursorOption(ctx *CursorOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitExternalAccessOption(ctx *ExternalAccessOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitQueryStoreOptions(ctx *QueryStoreOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitQueryStoreOptionList(ctx *QueryStoreOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitQueryCapturePolicyOptionList(ctx *QueryCapturePolicyOptionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitRecoveryOption(ctx *RecoveryOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSqlOption(ctx *SqlOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSnapshotOption(ctx *SnapshotOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitServiceBrokerOption(ctx *ServiceBrokerOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTargetRecoveryTimeOption(ctx *TargetRecoveryTimeOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitTermination(ctx *TerminationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateServiceClause(ctx *CreateServiceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterServiceClause(ctx *AlterServiceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterServiceOptArg(ctx *AlterServiceOptArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSchemaNameClause(ctx *SchemaNameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSchemaElement(ctx *SchemaElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateTableAsSelectClause(ctx *CreateTableAsSelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateTableAsSelect(ctx *CreateTableAsSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateRemoteTableAsSelect(ctx *CreateRemoteTableAsSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitWithDistributionOption(ctx *WithDistributionOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOptionQueryHintClause(ctx *OptionQueryHintClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitGrant(ctx *GrantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitRevoke(ctx *RevokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDeny(ctx *DenyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitClassPrivilegesClause(ctx *ClassPrivilegesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitClassTypePrivilegesClause(ctx *ClassTypePrivilegesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOptionForClause(ctx *OptionForClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitClassPrivileges(ctx *ClassPrivilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOnClassClause(ctx *OnClassClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitClassTypePrivileges(ctx *ClassTypePrivilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitOnClassTypeClause(ctx *OnClassTypeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitBasicPermission(ctx *BasicPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitObjectPermission(ctx *ObjectPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitServerPermission(ctx *ServerPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitServerPrincipalPermission(ctx *ServerPrincipalPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDatabasePermission(ctx *DatabasePermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDatabasePrincipalPermission(ctx *DatabasePrincipalPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSchemaPermission(ctx *SchemaPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitServiceBrokerPermission(ctx *ServiceBrokerPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitEndpointPermission(ctx *EndpointPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCertificatePermission(ctx *CertificatePermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSymmetricKeyPermission(ctx *SymmetricKeyPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAsymmetricKeyPermission(ctx *AsymmetricKeyPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAssemblyPermission(ctx *AssemblyPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAvailabilityGroupPermission(ctx *AvailabilityGroupPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitFullTextPermission(ctx *FullTextPermissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitClass_(ctx *Class_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitClassType(ctx *ClassTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitRoleClause(ctx *RoleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateUser(ctx *CreateUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropUser(ctx *DropUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterUser(ctx *AlterUserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateRole(ctx *CreateRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropRole(ctx *DropRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterRole(ctx *AlterRoleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCreateLogin(ctx *CreateLoginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitDropLogin(ctx *DropLoginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitAlterLogin(ctx *AlterLoginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSetTransaction(ctx *SetTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSetImplicitTransactions(ctx *SetImplicitTransactionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitImplicitTransactionsValue(ctx *ImplicitTransactionsValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitBeginTransaction(ctx *BeginTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitBeginDistributedTransaction(ctx *BeginDistributedTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCommit(ctx *CommitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCommitWork(ctx *CommitWorkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitRollback(ctx *RollbackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitRollbackWork(ctx *RollbackWorkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitSavepoint(ctx *SavepointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitExplain(ctx *ExplainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQLServerStatementVisitor) VisitExplainableStatement(ctx *ExplainableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

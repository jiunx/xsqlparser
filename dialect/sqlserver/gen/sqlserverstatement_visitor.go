// Code generated from E:/GoProject/src/github.com/2212442929/xsqlparser/dialect/sqlserver/grammer\SQLServerStatement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQLServerStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SQLServerStatementParser.
type SQLServerStatementVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SQLServerStatementParser#execute.
	VisitExecute(ctx *ExecuteContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#insert.
	VisitInsert(ctx *InsertContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#insertDefaultValue.
	VisitInsertDefaultValue(ctx *InsertDefaultValueContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#insertValuesClause.
	VisitInsertValuesClause(ctx *InsertValuesClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#insertSelectClause.
	VisitInsertSelectClause(ctx *InsertSelectClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#update.
	VisitUpdate(ctx *UpdateContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#setAssignmentsClause.
	VisitSetAssignmentsClause(ctx *SetAssignmentsClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#assignmentValues.
	VisitAssignmentValues(ctx *AssignmentValuesContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#assignmentValue.
	VisitAssignmentValue(ctx *AssignmentValueContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#delete.
	VisitDelete(ctx *DeleteContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#singleTableClause.
	VisitSingleTableClause(ctx *SingleTableClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#multipleTablesClause.
	VisitMultipleTablesClause(ctx *MultipleTablesClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#multipleTableNames.
	VisitMultipleTableNames(ctx *MultipleTableNamesContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#select.
	VisitSelect(ctx *SelectContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#aggregationClause.
	VisitAggregationClause(ctx *AggregationClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#selectClause.
	VisitSelectClause(ctx *SelectClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#duplicateSpecification.
	VisitDuplicateSpecification(ctx *DuplicateSpecificationContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#projections.
	VisitProjections(ctx *ProjectionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#projection.
	VisitProjection(ctx *ProjectionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#top.
	VisitTop(ctx *TopContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#topNum.
	VisitTopNum(ctx *TopNumContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#unqualifiedShorthand.
	VisitUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#qualifiedShorthand.
	VisitQualifiedShorthand(ctx *QualifiedShorthandContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableReferences.
	VisitTableReferences(ctx *TableReferencesContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableReference.
	VisitTableReference(ctx *TableReferenceContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableFactor.
	VisitTableFactor(ctx *TableFactorContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#joinedTable.
	VisitJoinedTable(ctx *JoinedTableContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#joinSpecification.
	VisitJoinSpecification(ctx *JoinSpecificationContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#cteClause.
	VisitCteClause(ctx *CteClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#outputClause.
	VisitOutputClause(ctx *OutputClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#outputWithColumns.
	VisitOutputWithColumns(ctx *OutputWithColumnsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#outputWithColumn.
	VisitOutputWithColumn(ctx *OutputWithColumnContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#outputWithAaterisk.
	VisitOutputWithAaterisk(ctx *OutputWithAateriskContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#outputTableName.
	VisitOutputTableName(ctx *OutputTableNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#queryHint.
	VisitQueryHint(ctx *QueryHintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#useHitName.
	VisitUseHitName(ctx *UseHitNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#parameterMarker.
	VisitParameterMarker(ctx *ParameterMarkerContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#literals.
	VisitLiterals(ctx *LiteralsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#stringLiterals.
	VisitStringLiterals(ctx *StringLiteralsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#numberLiterals.
	VisitNumberLiterals(ctx *NumberLiteralsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dateTimeLiterals.
	VisitDateTimeLiterals(ctx *DateTimeLiteralsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#hexadecimalLiterals.
	VisitHexadecimalLiterals(ctx *HexadecimalLiteralsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#bitValueLiterals.
	VisitBitValueLiterals(ctx *BitValueLiteralsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#booleanLiterals.
	VisitBooleanLiterals(ctx *BooleanLiteralsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#nullValueLiterals.
	VisitNullValueLiterals(ctx *NullValueLiteralsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#unreservedWord.
	VisitUnreservedWord(ctx *UnreservedWordContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#databaseName.
	VisitDatabaseName(ctx *DatabaseNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#schemaName.
	VisitSchemaName(ctx *SchemaNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#functionName.
	VisitFunctionName(ctx *FunctionNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#procedureName.
	VisitProcedureName(ctx *ProcedureNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#viewName.
	VisitViewName(ctx *ViewNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#triggerName.
	VisitTriggerName(ctx *TriggerNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#sequenceName.
	VisitSequenceName(ctx *SequenceNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#queueName.
	VisitQueueName(ctx *QueueNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#contractName.
	VisitContractName(ctx *ContractNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#serviceName.
	VisitServiceName(ctx *ServiceNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#owner.
	VisitOwner(ctx *OwnerContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#columnNames.
	VisitColumnNames(ctx *ColumnNamesContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#columnNamesWithSort.
	VisitColumnNamesWithSort(ctx *ColumnNamesWithSortContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableNames.
	VisitTableNames(ctx *TableNamesContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#indexName.
	VisitIndexName(ctx *IndexNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#constraintName.
	VisitConstraintName(ctx *ConstraintNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#collationName.
	VisitCollationName(ctx *CollationNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dataTypeLength.
	VisitDataTypeLength(ctx *DataTypeLengthContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#primaryKey.
	VisitPrimaryKey(ctx *PrimaryKeyContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#andOperator.
	VisitAndOperator(ctx *AndOperatorContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#orOperator.
	VisitOrOperator(ctx *OrOperatorContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#notOperator.
	VisitNotOperator(ctx *NotOperatorContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#booleanPrimary.
	VisitBooleanPrimary(ctx *BooleanPrimaryContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#bitExpr.
	VisitBitExpr(ctx *BitExprContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#simpleExpr.
	VisitSimpleExpr(ctx *SimpleExprContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#aggregationFunction.
	VisitAggregationFunction(ctx *AggregationFunctionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#aggregationFunctionName.
	VisitAggregationFunctionName(ctx *AggregationFunctionNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#distinct.
	VisitDistinct(ctx *DistinctContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#specialFunction.
	VisitSpecialFunction(ctx *SpecialFunctionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#castFunction.
	VisitCastFunction(ctx *CastFunctionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#charFunction.
	VisitCharFunction(ctx *CharFunctionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#regularFunction.
	VisitRegularFunction(ctx *RegularFunctionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#regularFunctionName.
	VisitRegularFunctionName(ctx *RegularFunctionNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#caseExpression.
	VisitCaseExpression(ctx *CaseExpressionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#caseWhen.
	VisitCaseWhen(ctx *CaseWhenContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#caseElse.
	VisitCaseElse(ctx *CaseElseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#privateExprOfDb.
	VisitPrivateExprOfDb(ctx *PrivateExprOfDbContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#orderByItem.
	VisitOrderByItem(ctx *OrderByItemContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dataTypeName.
	VisitDataTypeName(ctx *DataTypeNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#atTimeZoneExpr.
	VisitAtTimeZoneExpr(ctx *AtTimeZoneExprContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#castExpr.
	VisitCastExpr(ctx *CastExprContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#convertExpr.
	VisitConvertExpr(ctx *ConvertExprContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#windowedFunction.
	VisitWindowedFunction(ctx *WindowedFunctionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#overClause.
	VisitOverClause(ctx *OverClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#partitionByClause.
	VisitPartitionByClause(ctx *PartitionByClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#rowRangeClause.
	VisitRowRangeClause(ctx *RowRangeClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#windowFrameExtent.
	VisitWindowFrameExtent(ctx *WindowFrameExtentContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#windowFrameBetween.
	VisitWindowFrameBetween(ctx *WindowFrameBetweenContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#windowFrameBound.
	VisitWindowFrameBound(ctx *WindowFrameBoundContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#windowFramePreceding.
	VisitWindowFramePreceding(ctx *WindowFramePrecedingContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#windowFrameFollowing.
	VisitWindowFrameFollowing(ctx *WindowFrameFollowingContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#columnNameWithSort.
	VisitColumnNameWithSort(ctx *ColumnNameWithSortContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#indexOption.
	VisitIndexOption(ctx *IndexOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#compressionOption.
	VisitCompressionOption(ctx *CompressionOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#eqTime.
	VisitEqTime(ctx *EqTimeContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#eqOnOffOption.
	VisitEqOnOffOption(ctx *EqOnOffOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#eqKey.
	VisitEqKey(ctx *EqKeyContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#eqOnOff.
	VisitEqOnOff(ctx *EqOnOffContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#onPartitionClause.
	VisitOnPartitionClause(ctx *OnPartitionClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#partitionExpressions.
	VisitPartitionExpressions(ctx *PartitionExpressionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#partitionExpression.
	VisitPartitionExpression(ctx *PartitionExpressionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#numberRange.
	VisitNumberRange(ctx *NumberRangeContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#lowPriorityLockWait.
	VisitLowPriorityLockWait(ctx *LowPriorityLockWaitContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#onLowPriorLockWait.
	VisitOnLowPriorLockWait(ctx *OnLowPriorLockWaitContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#ignoredIdentifier.
	VisitIgnoredIdentifier(ctx *IgnoredIdentifierContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#ignoredIdentifiers.
	VisitIgnoredIdentifiers(ctx *IgnoredIdentifiersContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#matchNone.
	VisitMatchNone(ctx *MatchNoneContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#variableName.
	VisitVariableName(ctx *VariableNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#executeAsClause.
	VisitExecuteAsClause(ctx *ExecuteAsClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#transactionName.
	VisitTransactionName(ctx *TransactionNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#transactionVariableName.
	VisitTransactionVariableName(ctx *TransactionVariableNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#savepointName.
	VisitSavepointName(ctx *SavepointNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#savepointVariableName.
	VisitSavepointVariableName(ctx *SavepointVariableNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#entityType.
	VisitEntityType(ctx *EntityTypeContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createTable.
	VisitCreateTable(ctx *CreateTableContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createTableClause.
	VisitCreateTableClause(ctx *CreateTableClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createIndex.
	VisitCreateIndex(ctx *CreateIndexContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createDatabase.
	VisitCreateDatabase(ctx *CreateDatabaseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createFunction.
	VisitCreateFunction(ctx *CreateFunctionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createProcedure.
	VisitCreateProcedure(ctx *CreateProcedureContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createView.
	VisitCreateView(ctx *CreateViewContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createTrigger.
	VisitCreateTrigger(ctx *CreateTriggerContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createSequence.
	VisitCreateSequence(ctx *CreateSequenceContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createService.
	VisitCreateService(ctx *CreateServiceContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createSchema.
	VisitCreateSchema(ctx *CreateSchemaContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterTable.
	VisitAlterTable(ctx *AlterTableContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterIndex.
	VisitAlterIndex(ctx *AlterIndexContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterDatabase.
	VisitAlterDatabase(ctx *AlterDatabaseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterProcedure.
	VisitAlterProcedure(ctx *AlterProcedureContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterFunction.
	VisitAlterFunction(ctx *AlterFunctionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterView.
	VisitAlterView(ctx *AlterViewContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterTrigger.
	VisitAlterTrigger(ctx *AlterTriggerContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterSequence.
	VisitAlterSequence(ctx *AlterSequenceContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterService.
	VisitAlterService(ctx *AlterServiceContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterSchema.
	VisitAlterSchema(ctx *AlterSchemaContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropTable.
	VisitDropTable(ctx *DropTableContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropIndex.
	VisitDropIndex(ctx *DropIndexContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropDatabase.
	VisitDropDatabase(ctx *DropDatabaseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropFunction.
	VisitDropFunction(ctx *DropFunctionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropProcedure.
	VisitDropProcedure(ctx *DropProcedureContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropView.
	VisitDropView(ctx *DropViewContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropTrigger.
	VisitDropTrigger(ctx *DropTriggerContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropSequence.
	VisitDropSequence(ctx *DropSequenceContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropService.
	VisitDropService(ctx *DropServiceContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropSchema.
	VisitDropSchema(ctx *DropSchemaContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#truncateTable.
	VisitTruncateTable(ctx *TruncateTableContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#fileTableClause.
	VisitFileTableClause(ctx *FileTableClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createDefinitionClause.
	VisitCreateDefinitionClause(ctx *CreateDefinitionClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createTableDefinitions.
	VisitCreateTableDefinitions(ctx *CreateTableDefinitionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createTableDefinition.
	VisitCreateTableDefinition(ctx *CreateTableDefinitionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#columnDefinition.
	VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#columnDefinitionOption.
	VisitColumnDefinitionOption(ctx *ColumnDefinitionOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#encryptedOptions.
	VisitEncryptedOptions(ctx *EncryptedOptionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#columnConstraint.
	VisitColumnConstraint(ctx *ColumnConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#computedColumnConstraint.
	VisitComputedColumnConstraint(ctx *ComputedColumnConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#computedColumnForeignKeyConstraint.
	VisitComputedColumnForeignKeyConstraint(ctx *ComputedColumnForeignKeyConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#computedColumnForeignKeyOnAction.
	VisitComputedColumnForeignKeyOnAction(ctx *ComputedColumnForeignKeyOnActionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#primaryKeyConstraint.
	VisitPrimaryKeyConstraint(ctx *PrimaryKeyConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#diskTablePrimaryKeyConstraintOption.
	VisitDiskTablePrimaryKeyConstraintOption(ctx *DiskTablePrimaryKeyConstraintOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#clusterOption.
	VisitClusterOption(ctx *ClusterOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#primaryKeyWithClause.
	VisitPrimaryKeyWithClause(ctx *PrimaryKeyWithClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#primaryKeyOnClause.
	VisitPrimaryKeyOnClause(ctx *PrimaryKeyOnClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#onSchemaColumn.
	VisitOnSchemaColumn(ctx *OnSchemaColumnContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#onFileGroup.
	VisitOnFileGroup(ctx *OnFileGroupContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#onString.
	VisitOnString(ctx *OnStringContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#memoryTablePrimaryKeyConstraintOption.
	VisitMemoryTablePrimaryKeyConstraintOption(ctx *MemoryTablePrimaryKeyConstraintOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#withBucket.
	VisitWithBucket(ctx *WithBucketContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#columnForeignKeyConstraint.
	VisitColumnForeignKeyConstraint(ctx *ColumnForeignKeyConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#foreignKeyOnAction.
	VisitForeignKeyOnAction(ctx *ForeignKeyOnActionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#foreignKeyOn.
	VisitForeignKeyOn(ctx *ForeignKeyOnContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#checkConstraint.
	VisitCheckConstraint(ctx *CheckConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#columnIndex.
	VisitColumnIndex(ctx *ColumnIndexContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#withIndexOption.
	VisitWithIndexOption(ctx *WithIndexOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#indexOnClause.
	VisitIndexOnClause(ctx *IndexOnClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#onDefault.
	VisitOnDefault(ctx *OnDefaultContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#fileStreamOn.
	VisitFileStreamOn(ctx *FileStreamOnContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#columnConstraints.
	VisitColumnConstraints(ctx *ColumnConstraintsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#computedColumnDefinition.
	VisitComputedColumnDefinition(ctx *ComputedColumnDefinitionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#columnSetDefinition.
	VisitColumnSetDefinition(ctx *ColumnSetDefinitionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableConstraint.
	VisitTableConstraint(ctx *TableConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tablePrimaryConstraint.
	VisitTablePrimaryConstraint(ctx *TablePrimaryConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#primaryKeyUnique.
	VisitPrimaryKeyUnique(ctx *PrimaryKeyUniqueContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#diskTablePrimaryConstraintOption.
	VisitDiskTablePrimaryConstraintOption(ctx *DiskTablePrimaryConstraintOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#memoryTablePrimaryConstraintOption.
	VisitMemoryTablePrimaryConstraintOption(ctx *MemoryTablePrimaryConstraintOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#hashWithBucket.
	VisitHashWithBucket(ctx *HashWithBucketContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableForeignKeyConstraint.
	VisitTableForeignKeyConstraint(ctx *TableForeignKeyConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableIndex.
	VisitTableIndex(ctx *TableIndexContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#indexNameOption.
	VisitIndexNameOption(ctx *IndexNameOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#indexOptions.
	VisitIndexOptions(ctx *IndexOptionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#periodClause.
	VisitPeriodClause(ctx *PeriodClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#partitionScheme.
	VisitPartitionScheme(ctx *PartitionSchemeContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#fileGroup.
	VisitFileGroup(ctx *FileGroupContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableOptions.
	VisitTableOptions(ctx *TableOptionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableOption.
	VisitTableOption(ctx *TableOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dataDelectionOption.
	VisitDataDelectionOption(ctx *DataDelectionOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableStretchOptions.
	VisitTableStretchOptions(ctx *TableStretchOptionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableStretchOption.
	VisitTableStretchOption(ctx *TableStretchOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#migrationState_.
	VisitMigrationState_(ctx *MigrationState_Context) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableOperationOption.
	VisitTableOperationOption(ctx *TableOperationOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#distributionOption.
	VisitDistributionOption(ctx *DistributionOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dataWareHouseTableOption.
	VisitDataWareHouseTableOption(ctx *DataWareHouseTableOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dataWareHousePartitionOption.
	VisitDataWareHousePartitionOption(ctx *DataWareHousePartitionOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createIndexSpecification.
	VisitCreateIndexSpecification(ctx *CreateIndexSpecificationContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterDefinitionClause.
	VisitAlterDefinitionClause(ctx *AlterDefinitionClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#addColumnSpecification.
	VisitAddColumnSpecification(ctx *AddColumnSpecificationContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#modifyColumnSpecification.
	VisitModifyColumnSpecification(ctx *ModifyColumnSpecificationContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterColumnOperation.
	VisitAlterColumnOperation(ctx *AlterColumnOperationContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterColumnAddOptions.
	VisitAlterColumnAddOptions(ctx *AlterColumnAddOptionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterColumnAddOption.
	VisitAlterColumnAddOption(ctx *AlterColumnAddOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#constraintForColumn.
	VisitConstraintForColumn(ctx *ConstraintForColumnContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#generatedColumnNamesClause.
	VisitGeneratedColumnNamesClause(ctx *GeneratedColumnNamesClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#generatedColumnNameClause.
	VisitGeneratedColumnNameClause(ctx *GeneratedColumnNameClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#generatedColumnName.
	VisitGeneratedColumnName(ctx *GeneratedColumnNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterDrop.
	VisitAlterDrop(ctx *AlterDropContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterTableDropConstraint.
	VisitAlterTableDropConstraint(ctx *AlterTableDropConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropConstraintName.
	VisitDropConstraintName(ctx *DropConstraintNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropConstraintWithClause.
	VisitDropConstraintWithClause(ctx *DropConstraintWithClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropConstraintOption.
	VisitDropConstraintOption(ctx *DropConstraintOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#onOffOption.
	VisitOnOffOption(ctx *OnOffOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropColumnSpecification.
	VisitDropColumnSpecification(ctx *DropColumnSpecificationContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropIndexSpecification.
	VisitDropIndexSpecification(ctx *DropIndexSpecificationContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterCheckConstraint.
	VisitAlterCheckConstraint(ctx *AlterCheckConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterTableTrigger.
	VisitAlterTableTrigger(ctx *AlterTableTriggerContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterSwitch.
	VisitAlterSwitch(ctx *AlterSwitchContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterSet.
	VisitAlterSet(ctx *AlterSetContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#setFileStreamClause.
	VisitSetFileStreamClause(ctx *SetFileStreamClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#setSystemVersionClause.
	VisitSetSystemVersionClause(ctx *SetSystemVersionClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterSetOnClause.
	VisitAlterSetOnClause(ctx *AlterSetOnClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dataConsistencyCheckClause.
	VisitDataConsistencyCheckClause(ctx *DataConsistencyCheckClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#historyRetentionPeriodClause.
	VisitHistoryRetentionPeriodClause(ctx *HistoryRetentionPeriodClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#historyRetentionPeriod.
	VisitHistoryRetentionPeriod(ctx *HistoryRetentionPeriodContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterTableTableIndex.
	VisitAlterTableTableIndex(ctx *AlterTableTableIndexContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#indexWithName.
	VisitIndexWithName(ctx *IndexWithNameContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#indexNonClusterClause.
	VisitIndexNonClusterClause(ctx *IndexNonClusterClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterTableIndexOnClause.
	VisitAlterTableIndexOnClause(ctx *AlterTableIndexOnClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#indexClusterClause.
	VisitIndexClusterClause(ctx *IndexClusterClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterTableOption.
	VisitAlterTableOption(ctx *AlterTableOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#onHistoryTableClause.
	VisitOnHistoryTableClause(ctx *OnHistoryTableClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#ifExist.
	VisitIfExist(ctx *IfExistContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createDatabaseClause.
	VisitCreateDatabaseClause(ctx *CreateDatabaseClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#fileDefinitionClause.
	VisitFileDefinitionClause(ctx *FileDefinitionClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#databaseOption.
	VisitDatabaseOption(ctx *DatabaseOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#fileStreamOption.
	VisitFileStreamOption(ctx *FileStreamOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#fileSpec.
	VisitFileSpec(ctx *FileSpecContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#databaseFileSpecOption.
	VisitDatabaseFileSpecOption(ctx *DatabaseFileSpecOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#databaseFileGroup.
	VisitDatabaseFileGroup(ctx *DatabaseFileGroupContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#databaseFileGroupContains.
	VisitDatabaseFileGroupContains(ctx *DatabaseFileGroupContainsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#databaseLogOns.
	VisitDatabaseLogOns(ctx *DatabaseLogOnsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#declareVariable.
	VisitDeclareVariable(ctx *DeclareVariableContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableVariable.
	VisitTableVariable(ctx *TableVariableContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#variTableTypeDefinition.
	VisitVariTableTypeDefinition(ctx *VariTableTypeDefinitionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableVariableClause.
	VisitTableVariableClause(ctx *TableVariableClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#variableTableColumnDefinition.
	VisitVariableTableColumnDefinition(ctx *VariableTableColumnDefinitionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#variableTableColumnConstraint.
	VisitVariableTableColumnConstraint(ctx *VariableTableColumnConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#variableTableConstraint.
	VisitVariableTableConstraint(ctx *VariableTableConstraintContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#setVariable.
	VisitSetVariable(ctx *SetVariableContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#setVariableClause.
	VisitSetVariableClause(ctx *SetVariableClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#cursorVariable.
	VisitCursorVariable(ctx *CursorVariableContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#cursorClause.
	VisitCursorClause(ctx *CursorClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#compoundOperation.
	VisitCompoundOperation(ctx *CompoundOperationContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#funcParameters.
	VisitFuncParameters(ctx *FuncParametersContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#funcReturns.
	VisitFuncReturns(ctx *FuncReturnsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#funcMutiReturn.
	VisitFuncMutiReturn(ctx *FuncMutiReturnContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#funcInlineReturn.
	VisitFuncInlineReturn(ctx *FuncInlineReturnContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#funcScalarReturn.
	VisitFuncScalarReturn(ctx *FuncScalarReturnContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#tableTypeDefinition.
	VisitTableTypeDefinition(ctx *TableTypeDefinitionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#functionOption.
	VisitFunctionOption(ctx *FunctionOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#validStatement.
	VisitValidStatement(ctx *ValidStatementContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#procParameters.
	VisitProcParameters(ctx *ProcParametersContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#procParameter.
	VisitProcParameter(ctx *ProcParameterContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createOrAlterProcClause.
	VisitCreateOrAlterProcClause(ctx *CreateOrAlterProcClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#withCreateProcOption.
	VisitWithCreateProcOption(ctx *WithCreateProcOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#procOption.
	VisitProcOption(ctx *ProcOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#procAsClause.
	VisitProcAsClause(ctx *ProcAsClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#procSetOption.
	VisitProcSetOption(ctx *ProcSetOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createOrAlterViewClause.
	VisitCreateOrAlterViewClause(ctx *CreateOrAlterViewClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#viewAttribute.
	VisitViewAttribute(ctx *ViewAttributeContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#withCommonTableExpr.
	VisitWithCommonTableExpr(ctx *WithCommonTableExprContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#commonTableExpr.
	VisitCommonTableExpr(ctx *CommonTableExprContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createTriggerClause.
	VisitCreateTriggerClause(ctx *CreateTriggerClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dmlTriggerOption.
	VisitDmlTriggerOption(ctx *DmlTriggerOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#methodSpecifier.
	VisitMethodSpecifier(ctx *MethodSpecifierContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#triggerTarget.
	VisitTriggerTarget(ctx *TriggerTargetContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createOrAlterSequenceClause.
	VisitCreateOrAlterSequenceClause(ctx *CreateOrAlterSequenceClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createIndexClause.
	VisitCreateIndexClause(ctx *CreateIndexClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#filterPredicate.
	VisitFilterPredicate(ctx *FilterPredicateContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#conjunct.
	VisitConjunct(ctx *ConjunctContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterIndexClause.
	VisitAlterIndexClause(ctx *AlterIndexClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#relationalIndexOption.
	VisitRelationalIndexOption(ctx *RelationalIndexOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#partitionNumberRange.
	VisitPartitionNumberRange(ctx *PartitionNumberRangeContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#reorganizeOption.
	VisitReorganizeOption(ctx *ReorganizeOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#setIndexOption.
	VisitSetIndexOption(ctx *SetIndexOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#resumableIndexOptions.
	VisitResumableIndexOptions(ctx *ResumableIndexOptionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterDatabaseClause.
	VisitAlterDatabaseClause(ctx *AlterDatabaseClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#addSecondaryOption.
	VisitAddSecondaryOption(ctx *AddSecondaryOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#editionOptions.
	VisitEditionOptions(ctx *EditionOptionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#serviceObjective.
	VisitServiceObjective(ctx *ServiceObjectiveContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterDatabaseOptionSpec.
	VisitAlterDatabaseOptionSpec(ctx *AlterDatabaseOptionSpecContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#fileAndFilegroupOptions.
	VisitFileAndFilegroupOptions(ctx *FileAndFilegroupOptionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#addOrModifyFilegroups.
	VisitAddOrModifyFilegroups(ctx *AddOrModifyFilegroupsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#filegroupUpdatabilityOption.
	VisitFilegroupUpdatabilityOption(ctx *FilegroupUpdatabilityOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#addOrModifyFiles.
	VisitAddOrModifyFiles(ctx *AddOrModifyFilesContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#acceleratedDatabaseRecovery.
	VisitAcceleratedDatabaseRecovery(ctx *AcceleratedDatabaseRecoveryContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#autoOption.
	VisitAutoOption(ctx *AutoOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#automaticTuningOption.
	VisitAutomaticTuningOption(ctx *AutomaticTuningOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#changeTrackingOption.
	VisitChangeTrackingOption(ctx *ChangeTrackingOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#changeTrackingOptionList.
	VisitChangeTrackingOptionList(ctx *ChangeTrackingOptionListContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#cursorOption.
	VisitCursorOption(ctx *CursorOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#externalAccessOption.
	VisitExternalAccessOption(ctx *ExternalAccessOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#queryStoreOptions.
	VisitQueryStoreOptions(ctx *QueryStoreOptionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#queryStoreOptionList.
	VisitQueryStoreOptionList(ctx *QueryStoreOptionListContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#queryCapturePolicyOptionList.
	VisitQueryCapturePolicyOptionList(ctx *QueryCapturePolicyOptionListContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#recoveryOption.
	VisitRecoveryOption(ctx *RecoveryOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#sqlOption.
	VisitSqlOption(ctx *SqlOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#snapshotOption.
	VisitSnapshotOption(ctx *SnapshotOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#serviceBrokerOption.
	VisitServiceBrokerOption(ctx *ServiceBrokerOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#targetRecoveryTimeOption.
	VisitTargetRecoveryTimeOption(ctx *TargetRecoveryTimeOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#termination.
	VisitTermination(ctx *TerminationContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createServiceClause.
	VisitCreateServiceClause(ctx *CreateServiceClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterServiceClause.
	VisitAlterServiceClause(ctx *AlterServiceClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterServiceOptArg.
	VisitAlterServiceOptArg(ctx *AlterServiceOptArgContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#schemaNameClause.
	VisitSchemaNameClause(ctx *SchemaNameClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#schemaElement.
	VisitSchemaElement(ctx *SchemaElementContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createTableAsSelectClause.
	VisitCreateTableAsSelectClause(ctx *CreateTableAsSelectClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createTableAsSelect.
	VisitCreateTableAsSelect(ctx *CreateTableAsSelectContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createRemoteTableAsSelect.
	VisitCreateRemoteTableAsSelect(ctx *CreateRemoteTableAsSelectContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#withDistributionOption.
	VisitWithDistributionOption(ctx *WithDistributionOptionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#optionQueryHintClause.
	VisitOptionQueryHintClause(ctx *OptionQueryHintClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#grant.
	VisitGrant(ctx *GrantContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#revoke.
	VisitRevoke(ctx *RevokeContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#deny.
	VisitDeny(ctx *DenyContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#classPrivilegesClause.
	VisitClassPrivilegesClause(ctx *ClassPrivilegesClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#classTypePrivilegesClause.
	VisitClassTypePrivilegesClause(ctx *ClassTypePrivilegesClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#optionForClause.
	VisitOptionForClause(ctx *OptionForClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#classPrivileges.
	VisitClassPrivileges(ctx *ClassPrivilegesContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#onClassClause.
	VisitOnClassClause(ctx *OnClassClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#classTypePrivileges.
	VisitClassTypePrivileges(ctx *ClassTypePrivilegesContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#onClassTypeClause.
	VisitOnClassTypeClause(ctx *OnClassTypeClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#privilegeType.
	VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#basicPermission.
	VisitBasicPermission(ctx *BasicPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#objectPermission.
	VisitObjectPermission(ctx *ObjectPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#serverPermission.
	VisitServerPermission(ctx *ServerPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#serverPrincipalPermission.
	VisitServerPrincipalPermission(ctx *ServerPrincipalPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#databasePermission.
	VisitDatabasePermission(ctx *DatabasePermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#databasePrincipalPermission.
	VisitDatabasePrincipalPermission(ctx *DatabasePrincipalPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#schemaPermission.
	VisitSchemaPermission(ctx *SchemaPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#serviceBrokerPermission.
	VisitServiceBrokerPermission(ctx *ServiceBrokerPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#endpointPermission.
	VisitEndpointPermission(ctx *EndpointPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#certificatePermission.
	VisitCertificatePermission(ctx *CertificatePermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#symmetricKeyPermission.
	VisitSymmetricKeyPermission(ctx *SymmetricKeyPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#asymmetricKeyPermission.
	VisitAsymmetricKeyPermission(ctx *AsymmetricKeyPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#assemblyPermission.
	VisitAssemblyPermission(ctx *AssemblyPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#availabilityGroupPermission.
	VisitAvailabilityGroupPermission(ctx *AvailabilityGroupPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#fullTextPermission.
	VisitFullTextPermission(ctx *FullTextPermissionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#class_.
	VisitClass_(ctx *Class_Context) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#classType.
	VisitClassType(ctx *ClassTypeContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#roleClause.
	VisitRoleClause(ctx *RoleClauseContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createUser.
	VisitCreateUser(ctx *CreateUserContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropUser.
	VisitDropUser(ctx *DropUserContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterUser.
	VisitAlterUser(ctx *AlterUserContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createRole.
	VisitCreateRole(ctx *CreateRoleContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropRole.
	VisitDropRole(ctx *DropRoleContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterRole.
	VisitAlterRole(ctx *AlterRoleContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#createLogin.
	VisitCreateLogin(ctx *CreateLoginContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#dropLogin.
	VisitDropLogin(ctx *DropLoginContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#alterLogin.
	VisitAlterLogin(ctx *AlterLoginContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#setTransaction.
	VisitSetTransaction(ctx *SetTransactionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#setImplicitTransactions.
	VisitSetImplicitTransactions(ctx *SetImplicitTransactionsContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#implicitTransactionsValue.
	VisitImplicitTransactionsValue(ctx *ImplicitTransactionsValueContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#beginTransaction.
	VisitBeginTransaction(ctx *BeginTransactionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#beginDistributedTransaction.
	VisitBeginDistributedTransaction(ctx *BeginDistributedTransactionContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#commit.
	VisitCommit(ctx *CommitContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#commitWork.
	VisitCommitWork(ctx *CommitWorkContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#rollback.
	VisitRollback(ctx *RollbackContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#rollbackWork.
	VisitRollbackWork(ctx *RollbackWorkContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#savepoint.
	VisitSavepoint(ctx *SavepointContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#call.
	VisitCall(ctx *CallContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#explain.
	VisitExplain(ctx *ExplainContext) interface{}

	// Visit a parse tree produced by SQLServerStatementParser#explainableStatement.
	VisitExplainableStatement(ctx *ExplainableStatementContext) interface{}
}

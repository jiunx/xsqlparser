// Code generated from E:/GoProject/src/github.com/2212442929/xsqlparser/dialect/sqlserver/grammer\SQLServerStatement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQLServerStatement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SQLServerStatementListener is a complete listener for a parse tree produced by SQLServerStatementParser.
type SQLServerStatementListener interface {
	antlr.ParseTreeListener

	// EnterExecute is called when entering the execute production.
	EnterExecute(c *ExecuteContext)

	// EnterInsert is called when entering the insert production.
	EnterInsert(c *InsertContext)

	// EnterInsertDefaultValue is called when entering the insertDefaultValue production.
	EnterInsertDefaultValue(c *InsertDefaultValueContext)

	// EnterInsertValuesClause is called when entering the insertValuesClause production.
	EnterInsertValuesClause(c *InsertValuesClauseContext)

	// EnterInsertSelectClause is called when entering the insertSelectClause production.
	EnterInsertSelectClause(c *InsertSelectClauseContext)

	// EnterUpdate is called when entering the update production.
	EnterUpdate(c *UpdateContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterSetAssignmentsClause is called when entering the setAssignmentsClause production.
	EnterSetAssignmentsClause(c *SetAssignmentsClauseContext)

	// EnterAssignmentValues is called when entering the assignmentValues production.
	EnterAssignmentValues(c *AssignmentValuesContext)

	// EnterAssignmentValue is called when entering the assignmentValue production.
	EnterAssignmentValue(c *AssignmentValueContext)

	// EnterDelete is called when entering the delete production.
	EnterDelete(c *DeleteContext)

	// EnterSingleTableClause is called when entering the singleTableClause production.
	EnterSingleTableClause(c *SingleTableClauseContext)

	// EnterMultipleTablesClause is called when entering the multipleTablesClause production.
	EnterMultipleTablesClause(c *MultipleTablesClauseContext)

	// EnterMultipleTableNames is called when entering the multipleTableNames production.
	EnterMultipleTableNames(c *MultipleTableNamesContext)

	// EnterSelect is called when entering the select production.
	EnterSelect(c *SelectContext)

	// EnterAggregationClause is called when entering the aggregationClause production.
	EnterAggregationClause(c *AggregationClauseContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterDuplicateSpecification is called when entering the duplicateSpecification production.
	EnterDuplicateSpecification(c *DuplicateSpecificationContext)

	// EnterProjections is called when entering the projections production.
	EnterProjections(c *ProjectionsContext)

	// EnterProjection is called when entering the projection production.
	EnterProjection(c *ProjectionContext)

	// EnterTop is called when entering the top production.
	EnterTop(c *TopContext)

	// EnterTopNum is called when entering the topNum production.
	EnterTopNum(c *TopNumContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterUnqualifiedShorthand is called when entering the unqualifiedShorthand production.
	EnterUnqualifiedShorthand(c *UnqualifiedShorthandContext)

	// EnterQualifiedShorthand is called when entering the qualifiedShorthand production.
	EnterQualifiedShorthand(c *QualifiedShorthandContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterTableReferences is called when entering the tableReferences production.
	EnterTableReferences(c *TableReferencesContext)

	// EnterTableReference is called when entering the tableReference production.
	EnterTableReference(c *TableReferenceContext)

	// EnterTableFactor is called when entering the tableFactor production.
	EnterTableFactor(c *TableFactorContext)

	// EnterJoinedTable is called when entering the joinedTable production.
	EnterJoinedTable(c *JoinedTableContext)

	// EnterJoinSpecification is called when entering the joinSpecification production.
	EnterJoinSpecification(c *JoinSpecificationContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterCteClause is called when entering the cteClause production.
	EnterCteClause(c *CteClauseContext)

	// EnterOutputClause is called when entering the outputClause production.
	EnterOutputClause(c *OutputClauseContext)

	// EnterOutputWithColumns is called when entering the outputWithColumns production.
	EnterOutputWithColumns(c *OutputWithColumnsContext)

	// EnterOutputWithColumn is called when entering the outputWithColumn production.
	EnterOutputWithColumn(c *OutputWithColumnContext)

	// EnterOutputWithAaterisk is called when entering the outputWithAaterisk production.
	EnterOutputWithAaterisk(c *OutputWithAateriskContext)

	// EnterOutputTableName is called when entering the outputTableName production.
	EnterOutputTableName(c *OutputTableNameContext)

	// EnterQueryHint is called when entering the queryHint production.
	EnterQueryHint(c *QueryHintContext)

	// EnterUseHitName is called when entering the useHitName production.
	EnterUseHitName(c *UseHitNameContext)

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

	// EnterDatabaseName is called when entering the databaseName production.
	EnterDatabaseName(c *DatabaseNameContext)

	// EnterSchemaName is called when entering the schemaName production.
	EnterSchemaName(c *SchemaNameContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterProcedureName is called when entering the procedureName production.
	EnterProcedureName(c *ProcedureNameContext)

	// EnterViewName is called when entering the viewName production.
	EnterViewName(c *ViewNameContext)

	// EnterTriggerName is called when entering the triggerName production.
	EnterTriggerName(c *TriggerNameContext)

	// EnterSequenceName is called when entering the sequenceName production.
	EnterSequenceName(c *SequenceNameContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterQueueName is called when entering the queueName production.
	EnterQueueName(c *QueueNameContext)

	// EnterContractName is called when entering the contractName production.
	EnterContractName(c *ContractNameContext)

	// EnterServiceName is called when entering the serviceName production.
	EnterServiceName(c *ServiceNameContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterOwner is called when entering the owner production.
	EnterOwner(c *OwnerContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterColumnNames is called when entering the columnNames production.
	EnterColumnNames(c *ColumnNamesContext)

	// EnterColumnNamesWithSort is called when entering the columnNamesWithSort production.
	EnterColumnNamesWithSort(c *ColumnNamesWithSortContext)

	// EnterTableNames is called when entering the tableNames production.
	EnterTableNames(c *TableNamesContext)

	// EnterIndexName is called when entering the indexName production.
	EnterIndexName(c *IndexNameContext)

	// EnterConstraintName is called when entering the constraintName production.
	EnterConstraintName(c *ConstraintNameContext)

	// EnterCollationName is called when entering the collationName production.
	EnterCollationName(c *CollationNameContext)

	// EnterDataTypeLength is called when entering the dataTypeLength production.
	EnterDataTypeLength(c *DataTypeLengthContext)

	// EnterPrimaryKey is called when entering the primaryKey production.
	EnterPrimaryKey(c *PrimaryKeyContext)

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

	// EnterDistinct is called when entering the distinct production.
	EnterDistinct(c *DistinctContext)

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

	// EnterPrivateExprOfDb is called when entering the privateExprOfDb production.
	EnterPrivateExprOfDb(c *PrivateExprOfDbContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderByItem is called when entering the orderByItem production.
	EnterOrderByItem(c *OrderByItemContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterDataTypeName is called when entering the dataTypeName production.
	EnterDataTypeName(c *DataTypeNameContext)

	// EnterAtTimeZoneExpr is called when entering the atTimeZoneExpr production.
	EnterAtTimeZoneExpr(c *AtTimeZoneExprContext)

	// EnterCastExpr is called when entering the castExpr production.
	EnterCastExpr(c *CastExprContext)

	// EnterConvertExpr is called when entering the convertExpr production.
	EnterConvertExpr(c *ConvertExprContext)

	// EnterWindowedFunction is called when entering the windowedFunction production.
	EnterWindowedFunction(c *WindowedFunctionContext)

	// EnterOverClause is called when entering the overClause production.
	EnterOverClause(c *OverClauseContext)

	// EnterPartitionByClause is called when entering the partitionByClause production.
	EnterPartitionByClause(c *PartitionByClauseContext)

	// EnterRowRangeClause is called when entering the rowRangeClause production.
	EnterRowRangeClause(c *RowRangeClauseContext)

	// EnterWindowFrameExtent is called when entering the windowFrameExtent production.
	EnterWindowFrameExtent(c *WindowFrameExtentContext)

	// EnterWindowFrameBetween is called when entering the windowFrameBetween production.
	EnterWindowFrameBetween(c *WindowFrameBetweenContext)

	// EnterWindowFrameBound is called when entering the windowFrameBound production.
	EnterWindowFrameBound(c *WindowFrameBoundContext)

	// EnterWindowFramePreceding is called when entering the windowFramePreceding production.
	EnterWindowFramePreceding(c *WindowFramePrecedingContext)

	// EnterWindowFrameFollowing is called when entering the windowFrameFollowing production.
	EnterWindowFrameFollowing(c *WindowFrameFollowingContext)

	// EnterColumnNameWithSort is called when entering the columnNameWithSort production.
	EnterColumnNameWithSort(c *ColumnNameWithSortContext)

	// EnterIndexOption is called when entering the indexOption production.
	EnterIndexOption(c *IndexOptionContext)

	// EnterCompressionOption is called when entering the compressionOption production.
	EnterCompressionOption(c *CompressionOptionContext)

	// EnterEqTime is called when entering the eqTime production.
	EnterEqTime(c *EqTimeContext)

	// EnterEqOnOffOption is called when entering the eqOnOffOption production.
	EnterEqOnOffOption(c *EqOnOffOptionContext)

	// EnterEqKey is called when entering the eqKey production.
	EnterEqKey(c *EqKeyContext)

	// EnterEqOnOff is called when entering the eqOnOff production.
	EnterEqOnOff(c *EqOnOffContext)

	// EnterOnPartitionClause is called when entering the onPartitionClause production.
	EnterOnPartitionClause(c *OnPartitionClauseContext)

	// EnterPartitionExpressions is called when entering the partitionExpressions production.
	EnterPartitionExpressions(c *PartitionExpressionsContext)

	// EnterPartitionExpression is called when entering the partitionExpression production.
	EnterPartitionExpression(c *PartitionExpressionContext)

	// EnterNumberRange is called when entering the numberRange production.
	EnterNumberRange(c *NumberRangeContext)

	// EnterLowPriorityLockWait is called when entering the lowPriorityLockWait production.
	EnterLowPriorityLockWait(c *LowPriorityLockWaitContext)

	// EnterOnLowPriorLockWait is called when entering the onLowPriorLockWait production.
	EnterOnLowPriorLockWait(c *OnLowPriorLockWaitContext)

	// EnterIgnoredIdentifier is called when entering the ignoredIdentifier production.
	EnterIgnoredIdentifier(c *IgnoredIdentifierContext)

	// EnterIgnoredIdentifiers is called when entering the ignoredIdentifiers production.
	EnterIgnoredIdentifiers(c *IgnoredIdentifiersContext)

	// EnterMatchNone is called when entering the matchNone production.
	EnterMatchNone(c *MatchNoneContext)

	// EnterVariableName is called when entering the variableName production.
	EnterVariableName(c *VariableNameContext)

	// EnterExecuteAsClause is called when entering the executeAsClause production.
	EnterExecuteAsClause(c *ExecuteAsClauseContext)

	// EnterTransactionName is called when entering the transactionName production.
	EnterTransactionName(c *TransactionNameContext)

	// EnterTransactionVariableName is called when entering the transactionVariableName production.
	EnterTransactionVariableName(c *TransactionVariableNameContext)

	// EnterSavepointName is called when entering the savepointName production.
	EnterSavepointName(c *SavepointNameContext)

	// EnterSavepointVariableName is called when entering the savepointVariableName production.
	EnterSavepointVariableName(c *SavepointVariableNameContext)

	// EnterEntityType is called when entering the entityType production.
	EnterEntityType(c *EntityTypeContext)

	// EnterCreateTable is called when entering the createTable production.
	EnterCreateTable(c *CreateTableContext)

	// EnterCreateTableClause is called when entering the createTableClause production.
	EnterCreateTableClause(c *CreateTableClauseContext)

	// EnterCreateIndex is called when entering the createIndex production.
	EnterCreateIndex(c *CreateIndexContext)

	// EnterCreateDatabase is called when entering the createDatabase production.
	EnterCreateDatabase(c *CreateDatabaseContext)

	// EnterCreateFunction is called when entering the createFunction production.
	EnterCreateFunction(c *CreateFunctionContext)

	// EnterCreateProcedure is called when entering the createProcedure production.
	EnterCreateProcedure(c *CreateProcedureContext)

	// EnterCreateView is called when entering the createView production.
	EnterCreateView(c *CreateViewContext)

	// EnterCreateTrigger is called when entering the createTrigger production.
	EnterCreateTrigger(c *CreateTriggerContext)

	// EnterCreateSequence is called when entering the createSequence production.
	EnterCreateSequence(c *CreateSequenceContext)

	// EnterCreateService is called when entering the createService production.
	EnterCreateService(c *CreateServiceContext)

	// EnterCreateSchema is called when entering the createSchema production.
	EnterCreateSchema(c *CreateSchemaContext)

	// EnterAlterTable is called when entering the alterTable production.
	EnterAlterTable(c *AlterTableContext)

	// EnterAlterIndex is called when entering the alterIndex production.
	EnterAlterIndex(c *AlterIndexContext)

	// EnterAlterDatabase is called when entering the alterDatabase production.
	EnterAlterDatabase(c *AlterDatabaseContext)

	// EnterAlterProcedure is called when entering the alterProcedure production.
	EnterAlterProcedure(c *AlterProcedureContext)

	// EnterAlterFunction is called when entering the alterFunction production.
	EnterAlterFunction(c *AlterFunctionContext)

	// EnterAlterView is called when entering the alterView production.
	EnterAlterView(c *AlterViewContext)

	// EnterAlterTrigger is called when entering the alterTrigger production.
	EnterAlterTrigger(c *AlterTriggerContext)

	// EnterAlterSequence is called when entering the alterSequence production.
	EnterAlterSequence(c *AlterSequenceContext)

	// EnterAlterService is called when entering the alterService production.
	EnterAlterService(c *AlterServiceContext)

	// EnterAlterSchema is called when entering the alterSchema production.
	EnterAlterSchema(c *AlterSchemaContext)

	// EnterDropTable is called when entering the dropTable production.
	EnterDropTable(c *DropTableContext)

	// EnterDropIndex is called when entering the dropIndex production.
	EnterDropIndex(c *DropIndexContext)

	// EnterDropDatabase is called when entering the dropDatabase production.
	EnterDropDatabase(c *DropDatabaseContext)

	// EnterDropFunction is called when entering the dropFunction production.
	EnterDropFunction(c *DropFunctionContext)

	// EnterDropProcedure is called when entering the dropProcedure production.
	EnterDropProcedure(c *DropProcedureContext)

	// EnterDropView is called when entering the dropView production.
	EnterDropView(c *DropViewContext)

	// EnterDropTrigger is called when entering the dropTrigger production.
	EnterDropTrigger(c *DropTriggerContext)

	// EnterDropSequence is called when entering the dropSequence production.
	EnterDropSequence(c *DropSequenceContext)

	// EnterDropService is called when entering the dropService production.
	EnterDropService(c *DropServiceContext)

	// EnterDropSchema is called when entering the dropSchema production.
	EnterDropSchema(c *DropSchemaContext)

	// EnterTruncateTable is called when entering the truncateTable production.
	EnterTruncateTable(c *TruncateTableContext)

	// EnterFileTableClause is called when entering the fileTableClause production.
	EnterFileTableClause(c *FileTableClauseContext)

	// EnterCreateDefinitionClause is called when entering the createDefinitionClause production.
	EnterCreateDefinitionClause(c *CreateDefinitionClauseContext)

	// EnterCreateTableDefinitions is called when entering the createTableDefinitions production.
	EnterCreateTableDefinitions(c *CreateTableDefinitionsContext)

	// EnterCreateTableDefinition is called when entering the createTableDefinition production.
	EnterCreateTableDefinition(c *CreateTableDefinitionContext)

	// EnterColumnDefinition is called when entering the columnDefinition production.
	EnterColumnDefinition(c *ColumnDefinitionContext)

	// EnterColumnDefinitionOption is called when entering the columnDefinitionOption production.
	EnterColumnDefinitionOption(c *ColumnDefinitionOptionContext)

	// EnterEncryptedOptions is called when entering the encryptedOptions production.
	EnterEncryptedOptions(c *EncryptedOptionsContext)

	// EnterColumnConstraint is called when entering the columnConstraint production.
	EnterColumnConstraint(c *ColumnConstraintContext)

	// EnterComputedColumnConstraint is called when entering the computedColumnConstraint production.
	EnterComputedColumnConstraint(c *ComputedColumnConstraintContext)

	// EnterComputedColumnForeignKeyConstraint is called when entering the computedColumnForeignKeyConstraint production.
	EnterComputedColumnForeignKeyConstraint(c *ComputedColumnForeignKeyConstraintContext)

	// EnterComputedColumnForeignKeyOnAction is called when entering the computedColumnForeignKeyOnAction production.
	EnterComputedColumnForeignKeyOnAction(c *ComputedColumnForeignKeyOnActionContext)

	// EnterPrimaryKeyConstraint is called when entering the primaryKeyConstraint production.
	EnterPrimaryKeyConstraint(c *PrimaryKeyConstraintContext)

	// EnterDiskTablePrimaryKeyConstraintOption is called when entering the diskTablePrimaryKeyConstraintOption production.
	EnterDiskTablePrimaryKeyConstraintOption(c *DiskTablePrimaryKeyConstraintOptionContext)

	// EnterClusterOption is called when entering the clusterOption production.
	EnterClusterOption(c *ClusterOptionContext)

	// EnterPrimaryKeyWithClause is called when entering the primaryKeyWithClause production.
	EnterPrimaryKeyWithClause(c *PrimaryKeyWithClauseContext)

	// EnterPrimaryKeyOnClause is called when entering the primaryKeyOnClause production.
	EnterPrimaryKeyOnClause(c *PrimaryKeyOnClauseContext)

	// EnterOnSchemaColumn is called when entering the onSchemaColumn production.
	EnterOnSchemaColumn(c *OnSchemaColumnContext)

	// EnterOnFileGroup is called when entering the onFileGroup production.
	EnterOnFileGroup(c *OnFileGroupContext)

	// EnterOnString is called when entering the onString production.
	EnterOnString(c *OnStringContext)

	// EnterMemoryTablePrimaryKeyConstraintOption is called when entering the memoryTablePrimaryKeyConstraintOption production.
	EnterMemoryTablePrimaryKeyConstraintOption(c *MemoryTablePrimaryKeyConstraintOptionContext)

	// EnterWithBucket is called when entering the withBucket production.
	EnterWithBucket(c *WithBucketContext)

	// EnterColumnForeignKeyConstraint is called when entering the columnForeignKeyConstraint production.
	EnterColumnForeignKeyConstraint(c *ColumnForeignKeyConstraintContext)

	// EnterForeignKeyOnAction is called when entering the foreignKeyOnAction production.
	EnterForeignKeyOnAction(c *ForeignKeyOnActionContext)

	// EnterForeignKeyOn is called when entering the foreignKeyOn production.
	EnterForeignKeyOn(c *ForeignKeyOnContext)

	// EnterCheckConstraint is called when entering the checkConstraint production.
	EnterCheckConstraint(c *CheckConstraintContext)

	// EnterColumnIndex is called when entering the columnIndex production.
	EnterColumnIndex(c *ColumnIndexContext)

	// EnterWithIndexOption is called when entering the withIndexOption production.
	EnterWithIndexOption(c *WithIndexOptionContext)

	// EnterIndexOnClause is called when entering the indexOnClause production.
	EnterIndexOnClause(c *IndexOnClauseContext)

	// EnterOnDefault is called when entering the onDefault production.
	EnterOnDefault(c *OnDefaultContext)

	// EnterFileStreamOn is called when entering the fileStreamOn production.
	EnterFileStreamOn(c *FileStreamOnContext)

	// EnterColumnConstraints is called when entering the columnConstraints production.
	EnterColumnConstraints(c *ColumnConstraintsContext)

	// EnterComputedColumnDefinition is called when entering the computedColumnDefinition production.
	EnterComputedColumnDefinition(c *ComputedColumnDefinitionContext)

	// EnterColumnSetDefinition is called when entering the columnSetDefinition production.
	EnterColumnSetDefinition(c *ColumnSetDefinitionContext)

	// EnterTableConstraint is called when entering the tableConstraint production.
	EnterTableConstraint(c *TableConstraintContext)

	// EnterTablePrimaryConstraint is called when entering the tablePrimaryConstraint production.
	EnterTablePrimaryConstraint(c *TablePrimaryConstraintContext)

	// EnterPrimaryKeyUnique is called when entering the primaryKeyUnique production.
	EnterPrimaryKeyUnique(c *PrimaryKeyUniqueContext)

	// EnterDiskTablePrimaryConstraintOption is called when entering the diskTablePrimaryConstraintOption production.
	EnterDiskTablePrimaryConstraintOption(c *DiskTablePrimaryConstraintOptionContext)

	// EnterMemoryTablePrimaryConstraintOption is called when entering the memoryTablePrimaryConstraintOption production.
	EnterMemoryTablePrimaryConstraintOption(c *MemoryTablePrimaryConstraintOptionContext)

	// EnterHashWithBucket is called when entering the hashWithBucket production.
	EnterHashWithBucket(c *HashWithBucketContext)

	// EnterTableForeignKeyConstraint is called when entering the tableForeignKeyConstraint production.
	EnterTableForeignKeyConstraint(c *TableForeignKeyConstraintContext)

	// EnterTableIndex is called when entering the tableIndex production.
	EnterTableIndex(c *TableIndexContext)

	// EnterIndexNameOption is called when entering the indexNameOption production.
	EnterIndexNameOption(c *IndexNameOptionContext)

	// EnterIndexOptions is called when entering the indexOptions production.
	EnterIndexOptions(c *IndexOptionsContext)

	// EnterPeriodClause is called when entering the periodClause production.
	EnterPeriodClause(c *PeriodClauseContext)

	// EnterPartitionScheme is called when entering the partitionScheme production.
	EnterPartitionScheme(c *PartitionSchemeContext)

	// EnterFileGroup is called when entering the fileGroup production.
	EnterFileGroup(c *FileGroupContext)

	// EnterTableOptions is called when entering the tableOptions production.
	EnterTableOptions(c *TableOptionsContext)

	// EnterTableOption is called when entering the tableOption production.
	EnterTableOption(c *TableOptionContext)

	// EnterDataDelectionOption is called when entering the dataDelectionOption production.
	EnterDataDelectionOption(c *DataDelectionOptionContext)

	// EnterTableStretchOptions is called when entering the tableStretchOptions production.
	EnterTableStretchOptions(c *TableStretchOptionsContext)

	// EnterTableStretchOption is called when entering the tableStretchOption production.
	EnterTableStretchOption(c *TableStretchOptionContext)

	// EnterMigrationState_ is called when entering the migrationState_ production.
	EnterMigrationState_(c *MigrationState_Context)

	// EnterTableOperationOption is called when entering the tableOperationOption production.
	EnterTableOperationOption(c *TableOperationOptionContext)

	// EnterDistributionOption is called when entering the distributionOption production.
	EnterDistributionOption(c *DistributionOptionContext)

	// EnterDataWareHouseTableOption is called when entering the dataWareHouseTableOption production.
	EnterDataWareHouseTableOption(c *DataWareHouseTableOptionContext)

	// EnterDataWareHousePartitionOption is called when entering the dataWareHousePartitionOption production.
	EnterDataWareHousePartitionOption(c *DataWareHousePartitionOptionContext)

	// EnterCreateIndexSpecification is called when entering the createIndexSpecification production.
	EnterCreateIndexSpecification(c *CreateIndexSpecificationContext)

	// EnterAlterDefinitionClause is called when entering the alterDefinitionClause production.
	EnterAlterDefinitionClause(c *AlterDefinitionClauseContext)

	// EnterAddColumnSpecification is called when entering the addColumnSpecification production.
	EnterAddColumnSpecification(c *AddColumnSpecificationContext)

	// EnterModifyColumnSpecification is called when entering the modifyColumnSpecification production.
	EnterModifyColumnSpecification(c *ModifyColumnSpecificationContext)

	// EnterAlterColumnOperation is called when entering the alterColumnOperation production.
	EnterAlterColumnOperation(c *AlterColumnOperationContext)

	// EnterAlterColumnAddOptions is called when entering the alterColumnAddOptions production.
	EnterAlterColumnAddOptions(c *AlterColumnAddOptionsContext)

	// EnterAlterColumnAddOption is called when entering the alterColumnAddOption production.
	EnterAlterColumnAddOption(c *AlterColumnAddOptionContext)

	// EnterConstraintForColumn is called when entering the constraintForColumn production.
	EnterConstraintForColumn(c *ConstraintForColumnContext)

	// EnterGeneratedColumnNamesClause is called when entering the generatedColumnNamesClause production.
	EnterGeneratedColumnNamesClause(c *GeneratedColumnNamesClauseContext)

	// EnterGeneratedColumnNameClause is called when entering the generatedColumnNameClause production.
	EnterGeneratedColumnNameClause(c *GeneratedColumnNameClauseContext)

	// EnterGeneratedColumnName is called when entering the generatedColumnName production.
	EnterGeneratedColumnName(c *GeneratedColumnNameContext)

	// EnterAlterDrop is called when entering the alterDrop production.
	EnterAlterDrop(c *AlterDropContext)

	// EnterAlterTableDropConstraint is called when entering the alterTableDropConstraint production.
	EnterAlterTableDropConstraint(c *AlterTableDropConstraintContext)

	// EnterDropConstraintName is called when entering the dropConstraintName production.
	EnterDropConstraintName(c *DropConstraintNameContext)

	// EnterDropConstraintWithClause is called when entering the dropConstraintWithClause production.
	EnterDropConstraintWithClause(c *DropConstraintWithClauseContext)

	// EnterDropConstraintOption is called when entering the dropConstraintOption production.
	EnterDropConstraintOption(c *DropConstraintOptionContext)

	// EnterOnOffOption is called when entering the onOffOption production.
	EnterOnOffOption(c *OnOffOptionContext)

	// EnterDropColumnSpecification is called when entering the dropColumnSpecification production.
	EnterDropColumnSpecification(c *DropColumnSpecificationContext)

	// EnterDropIndexSpecification is called when entering the dropIndexSpecification production.
	EnterDropIndexSpecification(c *DropIndexSpecificationContext)

	// EnterAlterCheckConstraint is called when entering the alterCheckConstraint production.
	EnterAlterCheckConstraint(c *AlterCheckConstraintContext)

	// EnterAlterTableTrigger is called when entering the alterTableTrigger production.
	EnterAlterTableTrigger(c *AlterTableTriggerContext)

	// EnterAlterSwitch is called when entering the alterSwitch production.
	EnterAlterSwitch(c *AlterSwitchContext)

	// EnterAlterSet is called when entering the alterSet production.
	EnterAlterSet(c *AlterSetContext)

	// EnterSetFileStreamClause is called when entering the setFileStreamClause production.
	EnterSetFileStreamClause(c *SetFileStreamClauseContext)

	// EnterSetSystemVersionClause is called when entering the setSystemVersionClause production.
	EnterSetSystemVersionClause(c *SetSystemVersionClauseContext)

	// EnterAlterSetOnClause is called when entering the alterSetOnClause production.
	EnterAlterSetOnClause(c *AlterSetOnClauseContext)

	// EnterDataConsistencyCheckClause is called when entering the dataConsistencyCheckClause production.
	EnterDataConsistencyCheckClause(c *DataConsistencyCheckClauseContext)

	// EnterHistoryRetentionPeriodClause is called when entering the historyRetentionPeriodClause production.
	EnterHistoryRetentionPeriodClause(c *HistoryRetentionPeriodClauseContext)

	// EnterHistoryRetentionPeriod is called when entering the historyRetentionPeriod production.
	EnterHistoryRetentionPeriod(c *HistoryRetentionPeriodContext)

	// EnterAlterTableTableIndex is called when entering the alterTableTableIndex production.
	EnterAlterTableTableIndex(c *AlterTableTableIndexContext)

	// EnterIndexWithName is called when entering the indexWithName production.
	EnterIndexWithName(c *IndexWithNameContext)

	// EnterIndexNonClusterClause is called when entering the indexNonClusterClause production.
	EnterIndexNonClusterClause(c *IndexNonClusterClauseContext)

	// EnterAlterTableIndexOnClause is called when entering the alterTableIndexOnClause production.
	EnterAlterTableIndexOnClause(c *AlterTableIndexOnClauseContext)

	// EnterIndexClusterClause is called when entering the indexClusterClause production.
	EnterIndexClusterClause(c *IndexClusterClauseContext)

	// EnterAlterTableOption is called when entering the alterTableOption production.
	EnterAlterTableOption(c *AlterTableOptionContext)

	// EnterOnHistoryTableClause is called when entering the onHistoryTableClause production.
	EnterOnHistoryTableClause(c *OnHistoryTableClauseContext)

	// EnterIfExist is called when entering the ifExist production.
	EnterIfExist(c *IfExistContext)

	// EnterCreateDatabaseClause is called when entering the createDatabaseClause production.
	EnterCreateDatabaseClause(c *CreateDatabaseClauseContext)

	// EnterFileDefinitionClause is called when entering the fileDefinitionClause production.
	EnterFileDefinitionClause(c *FileDefinitionClauseContext)

	// EnterDatabaseOption is called when entering the databaseOption production.
	EnterDatabaseOption(c *DatabaseOptionContext)

	// EnterFileStreamOption is called when entering the fileStreamOption production.
	EnterFileStreamOption(c *FileStreamOptionContext)

	// EnterFileSpec is called when entering the fileSpec production.
	EnterFileSpec(c *FileSpecContext)

	// EnterDatabaseFileSpecOption is called when entering the databaseFileSpecOption production.
	EnterDatabaseFileSpecOption(c *DatabaseFileSpecOptionContext)

	// EnterDatabaseFileGroup is called when entering the databaseFileGroup production.
	EnterDatabaseFileGroup(c *DatabaseFileGroupContext)

	// EnterDatabaseFileGroupContains is called when entering the databaseFileGroupContains production.
	EnterDatabaseFileGroupContains(c *DatabaseFileGroupContainsContext)

	// EnterDatabaseLogOns is called when entering the databaseLogOns production.
	EnterDatabaseLogOns(c *DatabaseLogOnsContext)

	// EnterDeclareVariable is called when entering the declareVariable production.
	EnterDeclareVariable(c *DeclareVariableContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterTableVariable is called when entering the tableVariable production.
	EnterTableVariable(c *TableVariableContext)

	// EnterVariTableTypeDefinition is called when entering the variTableTypeDefinition production.
	EnterVariTableTypeDefinition(c *VariTableTypeDefinitionContext)

	// EnterTableVariableClause is called when entering the tableVariableClause production.
	EnterTableVariableClause(c *TableVariableClauseContext)

	// EnterVariableTableColumnDefinition is called when entering the variableTableColumnDefinition production.
	EnterVariableTableColumnDefinition(c *VariableTableColumnDefinitionContext)

	// EnterVariableTableColumnConstraint is called when entering the variableTableColumnConstraint production.
	EnterVariableTableColumnConstraint(c *VariableTableColumnConstraintContext)

	// EnterVariableTableConstraint is called when entering the variableTableConstraint production.
	EnterVariableTableConstraint(c *VariableTableConstraintContext)

	// EnterSetVariable is called when entering the setVariable production.
	EnterSetVariable(c *SetVariableContext)

	// EnterSetVariableClause is called when entering the setVariableClause production.
	EnterSetVariableClause(c *SetVariableClauseContext)

	// EnterCursorVariable is called when entering the cursorVariable production.
	EnterCursorVariable(c *CursorVariableContext)

	// EnterCursorClause is called when entering the cursorClause production.
	EnterCursorClause(c *CursorClauseContext)

	// EnterCompoundOperation is called when entering the compoundOperation production.
	EnterCompoundOperation(c *CompoundOperationContext)

	// EnterFuncParameters is called when entering the funcParameters production.
	EnterFuncParameters(c *FuncParametersContext)

	// EnterFuncReturns is called when entering the funcReturns production.
	EnterFuncReturns(c *FuncReturnsContext)

	// EnterFuncMutiReturn is called when entering the funcMutiReturn production.
	EnterFuncMutiReturn(c *FuncMutiReturnContext)

	// EnterFuncInlineReturn is called when entering the funcInlineReturn production.
	EnterFuncInlineReturn(c *FuncInlineReturnContext)

	// EnterFuncScalarReturn is called when entering the funcScalarReturn production.
	EnterFuncScalarReturn(c *FuncScalarReturnContext)

	// EnterTableTypeDefinition is called when entering the tableTypeDefinition production.
	EnterTableTypeDefinition(c *TableTypeDefinitionContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterFunctionOption is called when entering the functionOption production.
	EnterFunctionOption(c *FunctionOptionContext)

	// EnterValidStatement is called when entering the validStatement production.
	EnterValidStatement(c *ValidStatementContext)

	// EnterProcParameters is called when entering the procParameters production.
	EnterProcParameters(c *ProcParametersContext)

	// EnterProcParameter is called when entering the procParameter production.
	EnterProcParameter(c *ProcParameterContext)

	// EnterCreateOrAlterProcClause is called when entering the createOrAlterProcClause production.
	EnterCreateOrAlterProcClause(c *CreateOrAlterProcClauseContext)

	// EnterWithCreateProcOption is called when entering the withCreateProcOption production.
	EnterWithCreateProcOption(c *WithCreateProcOptionContext)

	// EnterProcOption is called when entering the procOption production.
	EnterProcOption(c *ProcOptionContext)

	// EnterProcAsClause is called when entering the procAsClause production.
	EnterProcAsClause(c *ProcAsClauseContext)

	// EnterProcSetOption is called when entering the procSetOption production.
	EnterProcSetOption(c *ProcSetOptionContext)

	// EnterCreateOrAlterViewClause is called when entering the createOrAlterViewClause production.
	EnterCreateOrAlterViewClause(c *CreateOrAlterViewClauseContext)

	// EnterViewAttribute is called when entering the viewAttribute production.
	EnterViewAttribute(c *ViewAttributeContext)

	// EnterWithCommonTableExpr is called when entering the withCommonTableExpr production.
	EnterWithCommonTableExpr(c *WithCommonTableExprContext)

	// EnterCommonTableExpr is called when entering the commonTableExpr production.
	EnterCommonTableExpr(c *CommonTableExprContext)

	// EnterCreateTriggerClause is called when entering the createTriggerClause production.
	EnterCreateTriggerClause(c *CreateTriggerClauseContext)

	// EnterDmlTriggerOption is called when entering the dmlTriggerOption production.
	EnterDmlTriggerOption(c *DmlTriggerOptionContext)

	// EnterMethodSpecifier is called when entering the methodSpecifier production.
	EnterMethodSpecifier(c *MethodSpecifierContext)

	// EnterTriggerTarget is called when entering the triggerTarget production.
	EnterTriggerTarget(c *TriggerTargetContext)

	// EnterCreateOrAlterSequenceClause is called when entering the createOrAlterSequenceClause production.
	EnterCreateOrAlterSequenceClause(c *CreateOrAlterSequenceClauseContext)

	// EnterCreateIndexClause is called when entering the createIndexClause production.
	EnterCreateIndexClause(c *CreateIndexClauseContext)

	// EnterFilterPredicate is called when entering the filterPredicate production.
	EnterFilterPredicate(c *FilterPredicateContext)

	// EnterConjunct is called when entering the conjunct production.
	EnterConjunct(c *ConjunctContext)

	// EnterAlterIndexClause is called when entering the alterIndexClause production.
	EnterAlterIndexClause(c *AlterIndexClauseContext)

	// EnterRelationalIndexOption is called when entering the relationalIndexOption production.
	EnterRelationalIndexOption(c *RelationalIndexOptionContext)

	// EnterPartitionNumberRange is called when entering the partitionNumberRange production.
	EnterPartitionNumberRange(c *PartitionNumberRangeContext)

	// EnterReorganizeOption is called when entering the reorganizeOption production.
	EnterReorganizeOption(c *ReorganizeOptionContext)

	// EnterSetIndexOption is called when entering the setIndexOption production.
	EnterSetIndexOption(c *SetIndexOptionContext)

	// EnterResumableIndexOptions is called when entering the resumableIndexOptions production.
	EnterResumableIndexOptions(c *ResumableIndexOptionsContext)

	// EnterAlterDatabaseClause is called when entering the alterDatabaseClause production.
	EnterAlterDatabaseClause(c *AlterDatabaseClauseContext)

	// EnterAddSecondaryOption is called when entering the addSecondaryOption production.
	EnterAddSecondaryOption(c *AddSecondaryOptionContext)

	// EnterEditionOptions is called when entering the editionOptions production.
	EnterEditionOptions(c *EditionOptionsContext)

	// EnterServiceObjective is called when entering the serviceObjective production.
	EnterServiceObjective(c *ServiceObjectiveContext)

	// EnterAlterDatabaseOptionSpec is called when entering the alterDatabaseOptionSpec production.
	EnterAlterDatabaseOptionSpec(c *AlterDatabaseOptionSpecContext)

	// EnterFileAndFilegroupOptions is called when entering the fileAndFilegroupOptions production.
	EnterFileAndFilegroupOptions(c *FileAndFilegroupOptionsContext)

	// EnterAddOrModifyFilegroups is called when entering the addOrModifyFilegroups production.
	EnterAddOrModifyFilegroups(c *AddOrModifyFilegroupsContext)

	// EnterFilegroupUpdatabilityOption is called when entering the filegroupUpdatabilityOption production.
	EnterFilegroupUpdatabilityOption(c *FilegroupUpdatabilityOptionContext)

	// EnterAddOrModifyFiles is called when entering the addOrModifyFiles production.
	EnterAddOrModifyFiles(c *AddOrModifyFilesContext)

	// EnterAcceleratedDatabaseRecovery is called when entering the acceleratedDatabaseRecovery production.
	EnterAcceleratedDatabaseRecovery(c *AcceleratedDatabaseRecoveryContext)

	// EnterAutoOption is called when entering the autoOption production.
	EnterAutoOption(c *AutoOptionContext)

	// EnterAutomaticTuningOption is called when entering the automaticTuningOption production.
	EnterAutomaticTuningOption(c *AutomaticTuningOptionContext)

	// EnterChangeTrackingOption is called when entering the changeTrackingOption production.
	EnterChangeTrackingOption(c *ChangeTrackingOptionContext)

	// EnterChangeTrackingOptionList is called when entering the changeTrackingOptionList production.
	EnterChangeTrackingOptionList(c *ChangeTrackingOptionListContext)

	// EnterCursorOption is called when entering the cursorOption production.
	EnterCursorOption(c *CursorOptionContext)

	// EnterExternalAccessOption is called when entering the externalAccessOption production.
	EnterExternalAccessOption(c *ExternalAccessOptionContext)

	// EnterQueryStoreOptions is called when entering the queryStoreOptions production.
	EnterQueryStoreOptions(c *QueryStoreOptionsContext)

	// EnterQueryStoreOptionList is called when entering the queryStoreOptionList production.
	EnterQueryStoreOptionList(c *QueryStoreOptionListContext)

	// EnterQueryCapturePolicyOptionList is called when entering the queryCapturePolicyOptionList production.
	EnterQueryCapturePolicyOptionList(c *QueryCapturePolicyOptionListContext)

	// EnterRecoveryOption is called when entering the recoveryOption production.
	EnterRecoveryOption(c *RecoveryOptionContext)

	// EnterSqlOption is called when entering the sqlOption production.
	EnterSqlOption(c *SqlOptionContext)

	// EnterSnapshotOption is called when entering the snapshotOption production.
	EnterSnapshotOption(c *SnapshotOptionContext)

	// EnterServiceBrokerOption is called when entering the serviceBrokerOption production.
	EnterServiceBrokerOption(c *ServiceBrokerOptionContext)

	// EnterTargetRecoveryTimeOption is called when entering the targetRecoveryTimeOption production.
	EnterTargetRecoveryTimeOption(c *TargetRecoveryTimeOptionContext)

	// EnterTermination is called when entering the termination production.
	EnterTermination(c *TerminationContext)

	// EnterCreateServiceClause is called when entering the createServiceClause production.
	EnterCreateServiceClause(c *CreateServiceClauseContext)

	// EnterAlterServiceClause is called when entering the alterServiceClause production.
	EnterAlterServiceClause(c *AlterServiceClauseContext)

	// EnterAlterServiceOptArg is called when entering the alterServiceOptArg production.
	EnterAlterServiceOptArg(c *AlterServiceOptArgContext)

	// EnterSchemaNameClause is called when entering the schemaNameClause production.
	EnterSchemaNameClause(c *SchemaNameClauseContext)

	// EnterSchemaElement is called when entering the schemaElement production.
	EnterSchemaElement(c *SchemaElementContext)

	// EnterCreateTableAsSelectClause is called when entering the createTableAsSelectClause production.
	EnterCreateTableAsSelectClause(c *CreateTableAsSelectClauseContext)

	// EnterCreateTableAsSelect is called when entering the createTableAsSelect production.
	EnterCreateTableAsSelect(c *CreateTableAsSelectContext)

	// EnterCreateRemoteTableAsSelect is called when entering the createRemoteTableAsSelect production.
	EnterCreateRemoteTableAsSelect(c *CreateRemoteTableAsSelectContext)

	// EnterWithDistributionOption is called when entering the withDistributionOption production.
	EnterWithDistributionOption(c *WithDistributionOptionContext)

	// EnterOptionQueryHintClause is called when entering the optionQueryHintClause production.
	EnterOptionQueryHintClause(c *OptionQueryHintClauseContext)

	// EnterGrant is called when entering the grant production.
	EnterGrant(c *GrantContext)

	// EnterRevoke is called when entering the revoke production.
	EnterRevoke(c *RevokeContext)

	// EnterDeny is called when entering the deny production.
	EnterDeny(c *DenyContext)

	// EnterClassPrivilegesClause is called when entering the classPrivilegesClause production.
	EnterClassPrivilegesClause(c *ClassPrivilegesClauseContext)

	// EnterClassTypePrivilegesClause is called when entering the classTypePrivilegesClause production.
	EnterClassTypePrivilegesClause(c *ClassTypePrivilegesClauseContext)

	// EnterOptionForClause is called when entering the optionForClause production.
	EnterOptionForClause(c *OptionForClauseContext)

	// EnterClassPrivileges is called when entering the classPrivileges production.
	EnterClassPrivileges(c *ClassPrivilegesContext)

	// EnterOnClassClause is called when entering the onClassClause production.
	EnterOnClassClause(c *OnClassClauseContext)

	// EnterClassTypePrivileges is called when entering the classTypePrivileges production.
	EnterClassTypePrivileges(c *ClassTypePrivilegesContext)

	// EnterOnClassTypeClause is called when entering the onClassTypeClause production.
	EnterOnClassTypeClause(c *OnClassTypeClauseContext)

	// EnterPrivilegeType is called when entering the privilegeType production.
	EnterPrivilegeType(c *PrivilegeTypeContext)

	// EnterBasicPermission is called when entering the basicPermission production.
	EnterBasicPermission(c *BasicPermissionContext)

	// EnterObjectPermission is called when entering the objectPermission production.
	EnterObjectPermission(c *ObjectPermissionContext)

	// EnterServerPermission is called when entering the serverPermission production.
	EnterServerPermission(c *ServerPermissionContext)

	// EnterServerPrincipalPermission is called when entering the serverPrincipalPermission production.
	EnterServerPrincipalPermission(c *ServerPrincipalPermissionContext)

	// EnterDatabasePermission is called when entering the databasePermission production.
	EnterDatabasePermission(c *DatabasePermissionContext)

	// EnterDatabasePrincipalPermission is called when entering the databasePrincipalPermission production.
	EnterDatabasePrincipalPermission(c *DatabasePrincipalPermissionContext)

	// EnterSchemaPermission is called when entering the schemaPermission production.
	EnterSchemaPermission(c *SchemaPermissionContext)

	// EnterServiceBrokerPermission is called when entering the serviceBrokerPermission production.
	EnterServiceBrokerPermission(c *ServiceBrokerPermissionContext)

	// EnterEndpointPermission is called when entering the endpointPermission production.
	EnterEndpointPermission(c *EndpointPermissionContext)

	// EnterCertificatePermission is called when entering the certificatePermission production.
	EnterCertificatePermission(c *CertificatePermissionContext)

	// EnterSymmetricKeyPermission is called when entering the symmetricKeyPermission production.
	EnterSymmetricKeyPermission(c *SymmetricKeyPermissionContext)

	// EnterAsymmetricKeyPermission is called when entering the asymmetricKeyPermission production.
	EnterAsymmetricKeyPermission(c *AsymmetricKeyPermissionContext)

	// EnterAssemblyPermission is called when entering the assemblyPermission production.
	EnterAssemblyPermission(c *AssemblyPermissionContext)

	// EnterAvailabilityGroupPermission is called when entering the availabilityGroupPermission production.
	EnterAvailabilityGroupPermission(c *AvailabilityGroupPermissionContext)

	// EnterFullTextPermission is called when entering the fullTextPermission production.
	EnterFullTextPermission(c *FullTextPermissionContext)

	// EnterClass_ is called when entering the class_ production.
	EnterClass_(c *Class_Context)

	// EnterClassType is called when entering the classType production.
	EnterClassType(c *ClassTypeContext)

	// EnterRoleClause is called when entering the roleClause production.
	EnterRoleClause(c *RoleClauseContext)

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

	// EnterCreateLogin is called when entering the createLogin production.
	EnterCreateLogin(c *CreateLoginContext)

	// EnterDropLogin is called when entering the dropLogin production.
	EnterDropLogin(c *DropLoginContext)

	// EnterAlterLogin is called when entering the alterLogin production.
	EnterAlterLogin(c *AlterLoginContext)

	// EnterSetTransaction is called when entering the setTransaction production.
	EnterSetTransaction(c *SetTransactionContext)

	// EnterSetImplicitTransactions is called when entering the setImplicitTransactions production.
	EnterSetImplicitTransactions(c *SetImplicitTransactionsContext)

	// EnterImplicitTransactionsValue is called when entering the implicitTransactionsValue production.
	EnterImplicitTransactionsValue(c *ImplicitTransactionsValueContext)

	// EnterBeginTransaction is called when entering the beginTransaction production.
	EnterBeginTransaction(c *BeginTransactionContext)

	// EnterBeginDistributedTransaction is called when entering the beginDistributedTransaction production.
	EnterBeginDistributedTransaction(c *BeginDistributedTransactionContext)

	// EnterCommit is called when entering the commit production.
	EnterCommit(c *CommitContext)

	// EnterCommitWork is called when entering the commitWork production.
	EnterCommitWork(c *CommitWorkContext)

	// EnterRollback is called when entering the rollback production.
	EnterRollback(c *RollbackContext)

	// EnterRollbackWork is called when entering the rollbackWork production.
	EnterRollbackWork(c *RollbackWorkContext)

	// EnterSavepoint is called when entering the savepoint production.
	EnterSavepoint(c *SavepointContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterExplain is called when entering the explain production.
	EnterExplain(c *ExplainContext)

	// EnterExplainableStatement is called when entering the explainableStatement production.
	EnterExplainableStatement(c *ExplainableStatementContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)

	// ExitInsert is called when exiting the insert production.
	ExitInsert(c *InsertContext)

	// ExitInsertDefaultValue is called when exiting the insertDefaultValue production.
	ExitInsertDefaultValue(c *InsertDefaultValueContext)

	// ExitInsertValuesClause is called when exiting the insertValuesClause production.
	ExitInsertValuesClause(c *InsertValuesClauseContext)

	// ExitInsertSelectClause is called when exiting the insertSelectClause production.
	ExitInsertSelectClause(c *InsertSelectClauseContext)

	// ExitUpdate is called when exiting the update production.
	ExitUpdate(c *UpdateContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitSetAssignmentsClause is called when exiting the setAssignmentsClause production.
	ExitSetAssignmentsClause(c *SetAssignmentsClauseContext)

	// ExitAssignmentValues is called when exiting the assignmentValues production.
	ExitAssignmentValues(c *AssignmentValuesContext)

	// ExitAssignmentValue is called when exiting the assignmentValue production.
	ExitAssignmentValue(c *AssignmentValueContext)

	// ExitDelete is called when exiting the delete production.
	ExitDelete(c *DeleteContext)

	// ExitSingleTableClause is called when exiting the singleTableClause production.
	ExitSingleTableClause(c *SingleTableClauseContext)

	// ExitMultipleTablesClause is called when exiting the multipleTablesClause production.
	ExitMultipleTablesClause(c *MultipleTablesClauseContext)

	// ExitMultipleTableNames is called when exiting the multipleTableNames production.
	ExitMultipleTableNames(c *MultipleTableNamesContext)

	// ExitSelect is called when exiting the select production.
	ExitSelect(c *SelectContext)

	// ExitAggregationClause is called when exiting the aggregationClause production.
	ExitAggregationClause(c *AggregationClauseContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitDuplicateSpecification is called when exiting the duplicateSpecification production.
	ExitDuplicateSpecification(c *DuplicateSpecificationContext)

	// ExitProjections is called when exiting the projections production.
	ExitProjections(c *ProjectionsContext)

	// ExitProjection is called when exiting the projection production.
	ExitProjection(c *ProjectionContext)

	// ExitTop is called when exiting the top production.
	ExitTop(c *TopContext)

	// ExitTopNum is called when exiting the topNum production.
	ExitTopNum(c *TopNumContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitUnqualifiedShorthand is called when exiting the unqualifiedShorthand production.
	ExitUnqualifiedShorthand(c *UnqualifiedShorthandContext)

	// ExitQualifiedShorthand is called when exiting the qualifiedShorthand production.
	ExitQualifiedShorthand(c *QualifiedShorthandContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitTableReferences is called when exiting the tableReferences production.
	ExitTableReferences(c *TableReferencesContext)

	// ExitTableReference is called when exiting the tableReference production.
	ExitTableReference(c *TableReferenceContext)

	// ExitTableFactor is called when exiting the tableFactor production.
	ExitTableFactor(c *TableFactorContext)

	// ExitJoinedTable is called when exiting the joinedTable production.
	ExitJoinedTable(c *JoinedTableContext)

	// ExitJoinSpecification is called when exiting the joinSpecification production.
	ExitJoinSpecification(c *JoinSpecificationContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitCteClause is called when exiting the cteClause production.
	ExitCteClause(c *CteClauseContext)

	// ExitOutputClause is called when exiting the outputClause production.
	ExitOutputClause(c *OutputClauseContext)

	// ExitOutputWithColumns is called when exiting the outputWithColumns production.
	ExitOutputWithColumns(c *OutputWithColumnsContext)

	// ExitOutputWithColumn is called when exiting the outputWithColumn production.
	ExitOutputWithColumn(c *OutputWithColumnContext)

	// ExitOutputWithAaterisk is called when exiting the outputWithAaterisk production.
	ExitOutputWithAaterisk(c *OutputWithAateriskContext)

	// ExitOutputTableName is called when exiting the outputTableName production.
	ExitOutputTableName(c *OutputTableNameContext)

	// ExitQueryHint is called when exiting the queryHint production.
	ExitQueryHint(c *QueryHintContext)

	// ExitUseHitName is called when exiting the useHitName production.
	ExitUseHitName(c *UseHitNameContext)

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

	// ExitDatabaseName is called when exiting the databaseName production.
	ExitDatabaseName(c *DatabaseNameContext)

	// ExitSchemaName is called when exiting the schemaName production.
	ExitSchemaName(c *SchemaNameContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitProcedureName is called when exiting the procedureName production.
	ExitProcedureName(c *ProcedureNameContext)

	// ExitViewName is called when exiting the viewName production.
	ExitViewName(c *ViewNameContext)

	// ExitTriggerName is called when exiting the triggerName production.
	ExitTriggerName(c *TriggerNameContext)

	// ExitSequenceName is called when exiting the sequenceName production.
	ExitSequenceName(c *SequenceNameContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitQueueName is called when exiting the queueName production.
	ExitQueueName(c *QueueNameContext)

	// ExitContractName is called when exiting the contractName production.
	ExitContractName(c *ContractNameContext)

	// ExitServiceName is called when exiting the serviceName production.
	ExitServiceName(c *ServiceNameContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitOwner is called when exiting the owner production.
	ExitOwner(c *OwnerContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitColumnNames is called when exiting the columnNames production.
	ExitColumnNames(c *ColumnNamesContext)

	// ExitColumnNamesWithSort is called when exiting the columnNamesWithSort production.
	ExitColumnNamesWithSort(c *ColumnNamesWithSortContext)

	// ExitTableNames is called when exiting the tableNames production.
	ExitTableNames(c *TableNamesContext)

	// ExitIndexName is called when exiting the indexName production.
	ExitIndexName(c *IndexNameContext)

	// ExitConstraintName is called when exiting the constraintName production.
	ExitConstraintName(c *ConstraintNameContext)

	// ExitCollationName is called when exiting the collationName production.
	ExitCollationName(c *CollationNameContext)

	// ExitDataTypeLength is called when exiting the dataTypeLength production.
	ExitDataTypeLength(c *DataTypeLengthContext)

	// ExitPrimaryKey is called when exiting the primaryKey production.
	ExitPrimaryKey(c *PrimaryKeyContext)

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

	// ExitDistinct is called when exiting the distinct production.
	ExitDistinct(c *DistinctContext)

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

	// ExitPrivateExprOfDb is called when exiting the privateExprOfDb production.
	ExitPrivateExprOfDb(c *PrivateExprOfDbContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderByItem is called when exiting the orderByItem production.
	ExitOrderByItem(c *OrderByItemContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitDataTypeName is called when exiting the dataTypeName production.
	ExitDataTypeName(c *DataTypeNameContext)

	// ExitAtTimeZoneExpr is called when exiting the atTimeZoneExpr production.
	ExitAtTimeZoneExpr(c *AtTimeZoneExprContext)

	// ExitCastExpr is called when exiting the castExpr production.
	ExitCastExpr(c *CastExprContext)

	// ExitConvertExpr is called when exiting the convertExpr production.
	ExitConvertExpr(c *ConvertExprContext)

	// ExitWindowedFunction is called when exiting the windowedFunction production.
	ExitWindowedFunction(c *WindowedFunctionContext)

	// ExitOverClause is called when exiting the overClause production.
	ExitOverClause(c *OverClauseContext)

	// ExitPartitionByClause is called when exiting the partitionByClause production.
	ExitPartitionByClause(c *PartitionByClauseContext)

	// ExitRowRangeClause is called when exiting the rowRangeClause production.
	ExitRowRangeClause(c *RowRangeClauseContext)

	// ExitWindowFrameExtent is called when exiting the windowFrameExtent production.
	ExitWindowFrameExtent(c *WindowFrameExtentContext)

	// ExitWindowFrameBetween is called when exiting the windowFrameBetween production.
	ExitWindowFrameBetween(c *WindowFrameBetweenContext)

	// ExitWindowFrameBound is called when exiting the windowFrameBound production.
	ExitWindowFrameBound(c *WindowFrameBoundContext)

	// ExitWindowFramePreceding is called when exiting the windowFramePreceding production.
	ExitWindowFramePreceding(c *WindowFramePrecedingContext)

	// ExitWindowFrameFollowing is called when exiting the windowFrameFollowing production.
	ExitWindowFrameFollowing(c *WindowFrameFollowingContext)

	// ExitColumnNameWithSort is called when exiting the columnNameWithSort production.
	ExitColumnNameWithSort(c *ColumnNameWithSortContext)

	// ExitIndexOption is called when exiting the indexOption production.
	ExitIndexOption(c *IndexOptionContext)

	// ExitCompressionOption is called when exiting the compressionOption production.
	ExitCompressionOption(c *CompressionOptionContext)

	// ExitEqTime is called when exiting the eqTime production.
	ExitEqTime(c *EqTimeContext)

	// ExitEqOnOffOption is called when exiting the eqOnOffOption production.
	ExitEqOnOffOption(c *EqOnOffOptionContext)

	// ExitEqKey is called when exiting the eqKey production.
	ExitEqKey(c *EqKeyContext)

	// ExitEqOnOff is called when exiting the eqOnOff production.
	ExitEqOnOff(c *EqOnOffContext)

	// ExitOnPartitionClause is called when exiting the onPartitionClause production.
	ExitOnPartitionClause(c *OnPartitionClauseContext)

	// ExitPartitionExpressions is called when exiting the partitionExpressions production.
	ExitPartitionExpressions(c *PartitionExpressionsContext)

	// ExitPartitionExpression is called when exiting the partitionExpression production.
	ExitPartitionExpression(c *PartitionExpressionContext)

	// ExitNumberRange is called when exiting the numberRange production.
	ExitNumberRange(c *NumberRangeContext)

	// ExitLowPriorityLockWait is called when exiting the lowPriorityLockWait production.
	ExitLowPriorityLockWait(c *LowPriorityLockWaitContext)

	// ExitOnLowPriorLockWait is called when exiting the onLowPriorLockWait production.
	ExitOnLowPriorLockWait(c *OnLowPriorLockWaitContext)

	// ExitIgnoredIdentifier is called when exiting the ignoredIdentifier production.
	ExitIgnoredIdentifier(c *IgnoredIdentifierContext)

	// ExitIgnoredIdentifiers is called when exiting the ignoredIdentifiers production.
	ExitIgnoredIdentifiers(c *IgnoredIdentifiersContext)

	// ExitMatchNone is called when exiting the matchNone production.
	ExitMatchNone(c *MatchNoneContext)

	// ExitVariableName is called when exiting the variableName production.
	ExitVariableName(c *VariableNameContext)

	// ExitExecuteAsClause is called when exiting the executeAsClause production.
	ExitExecuteAsClause(c *ExecuteAsClauseContext)

	// ExitTransactionName is called when exiting the transactionName production.
	ExitTransactionName(c *TransactionNameContext)

	// ExitTransactionVariableName is called when exiting the transactionVariableName production.
	ExitTransactionVariableName(c *TransactionVariableNameContext)

	// ExitSavepointName is called when exiting the savepointName production.
	ExitSavepointName(c *SavepointNameContext)

	// ExitSavepointVariableName is called when exiting the savepointVariableName production.
	ExitSavepointVariableName(c *SavepointVariableNameContext)

	// ExitEntityType is called when exiting the entityType production.
	ExitEntityType(c *EntityTypeContext)

	// ExitCreateTable is called when exiting the createTable production.
	ExitCreateTable(c *CreateTableContext)

	// ExitCreateTableClause is called when exiting the createTableClause production.
	ExitCreateTableClause(c *CreateTableClauseContext)

	// ExitCreateIndex is called when exiting the createIndex production.
	ExitCreateIndex(c *CreateIndexContext)

	// ExitCreateDatabase is called when exiting the createDatabase production.
	ExitCreateDatabase(c *CreateDatabaseContext)

	// ExitCreateFunction is called when exiting the createFunction production.
	ExitCreateFunction(c *CreateFunctionContext)

	// ExitCreateProcedure is called when exiting the createProcedure production.
	ExitCreateProcedure(c *CreateProcedureContext)

	// ExitCreateView is called when exiting the createView production.
	ExitCreateView(c *CreateViewContext)

	// ExitCreateTrigger is called when exiting the createTrigger production.
	ExitCreateTrigger(c *CreateTriggerContext)

	// ExitCreateSequence is called when exiting the createSequence production.
	ExitCreateSequence(c *CreateSequenceContext)

	// ExitCreateService is called when exiting the createService production.
	ExitCreateService(c *CreateServiceContext)

	// ExitCreateSchema is called when exiting the createSchema production.
	ExitCreateSchema(c *CreateSchemaContext)

	// ExitAlterTable is called when exiting the alterTable production.
	ExitAlterTable(c *AlterTableContext)

	// ExitAlterIndex is called when exiting the alterIndex production.
	ExitAlterIndex(c *AlterIndexContext)

	// ExitAlterDatabase is called when exiting the alterDatabase production.
	ExitAlterDatabase(c *AlterDatabaseContext)

	// ExitAlterProcedure is called when exiting the alterProcedure production.
	ExitAlterProcedure(c *AlterProcedureContext)

	// ExitAlterFunction is called when exiting the alterFunction production.
	ExitAlterFunction(c *AlterFunctionContext)

	// ExitAlterView is called when exiting the alterView production.
	ExitAlterView(c *AlterViewContext)

	// ExitAlterTrigger is called when exiting the alterTrigger production.
	ExitAlterTrigger(c *AlterTriggerContext)

	// ExitAlterSequence is called when exiting the alterSequence production.
	ExitAlterSequence(c *AlterSequenceContext)

	// ExitAlterService is called when exiting the alterService production.
	ExitAlterService(c *AlterServiceContext)

	// ExitAlterSchema is called when exiting the alterSchema production.
	ExitAlterSchema(c *AlterSchemaContext)

	// ExitDropTable is called when exiting the dropTable production.
	ExitDropTable(c *DropTableContext)

	// ExitDropIndex is called when exiting the dropIndex production.
	ExitDropIndex(c *DropIndexContext)

	// ExitDropDatabase is called when exiting the dropDatabase production.
	ExitDropDatabase(c *DropDatabaseContext)

	// ExitDropFunction is called when exiting the dropFunction production.
	ExitDropFunction(c *DropFunctionContext)

	// ExitDropProcedure is called when exiting the dropProcedure production.
	ExitDropProcedure(c *DropProcedureContext)

	// ExitDropView is called when exiting the dropView production.
	ExitDropView(c *DropViewContext)

	// ExitDropTrigger is called when exiting the dropTrigger production.
	ExitDropTrigger(c *DropTriggerContext)

	// ExitDropSequence is called when exiting the dropSequence production.
	ExitDropSequence(c *DropSequenceContext)

	// ExitDropService is called when exiting the dropService production.
	ExitDropService(c *DropServiceContext)

	// ExitDropSchema is called when exiting the dropSchema production.
	ExitDropSchema(c *DropSchemaContext)

	// ExitTruncateTable is called when exiting the truncateTable production.
	ExitTruncateTable(c *TruncateTableContext)

	// ExitFileTableClause is called when exiting the fileTableClause production.
	ExitFileTableClause(c *FileTableClauseContext)

	// ExitCreateDefinitionClause is called when exiting the createDefinitionClause production.
	ExitCreateDefinitionClause(c *CreateDefinitionClauseContext)

	// ExitCreateTableDefinitions is called when exiting the createTableDefinitions production.
	ExitCreateTableDefinitions(c *CreateTableDefinitionsContext)

	// ExitCreateTableDefinition is called when exiting the createTableDefinition production.
	ExitCreateTableDefinition(c *CreateTableDefinitionContext)

	// ExitColumnDefinition is called when exiting the columnDefinition production.
	ExitColumnDefinition(c *ColumnDefinitionContext)

	// ExitColumnDefinitionOption is called when exiting the columnDefinitionOption production.
	ExitColumnDefinitionOption(c *ColumnDefinitionOptionContext)

	// ExitEncryptedOptions is called when exiting the encryptedOptions production.
	ExitEncryptedOptions(c *EncryptedOptionsContext)

	// ExitColumnConstraint is called when exiting the columnConstraint production.
	ExitColumnConstraint(c *ColumnConstraintContext)

	// ExitComputedColumnConstraint is called when exiting the computedColumnConstraint production.
	ExitComputedColumnConstraint(c *ComputedColumnConstraintContext)

	// ExitComputedColumnForeignKeyConstraint is called when exiting the computedColumnForeignKeyConstraint production.
	ExitComputedColumnForeignKeyConstraint(c *ComputedColumnForeignKeyConstraintContext)

	// ExitComputedColumnForeignKeyOnAction is called when exiting the computedColumnForeignKeyOnAction production.
	ExitComputedColumnForeignKeyOnAction(c *ComputedColumnForeignKeyOnActionContext)

	// ExitPrimaryKeyConstraint is called when exiting the primaryKeyConstraint production.
	ExitPrimaryKeyConstraint(c *PrimaryKeyConstraintContext)

	// ExitDiskTablePrimaryKeyConstraintOption is called when exiting the diskTablePrimaryKeyConstraintOption production.
	ExitDiskTablePrimaryKeyConstraintOption(c *DiskTablePrimaryKeyConstraintOptionContext)

	// ExitClusterOption is called when exiting the clusterOption production.
	ExitClusterOption(c *ClusterOptionContext)

	// ExitPrimaryKeyWithClause is called when exiting the primaryKeyWithClause production.
	ExitPrimaryKeyWithClause(c *PrimaryKeyWithClauseContext)

	// ExitPrimaryKeyOnClause is called when exiting the primaryKeyOnClause production.
	ExitPrimaryKeyOnClause(c *PrimaryKeyOnClauseContext)

	// ExitOnSchemaColumn is called when exiting the onSchemaColumn production.
	ExitOnSchemaColumn(c *OnSchemaColumnContext)

	// ExitOnFileGroup is called when exiting the onFileGroup production.
	ExitOnFileGroup(c *OnFileGroupContext)

	// ExitOnString is called when exiting the onString production.
	ExitOnString(c *OnStringContext)

	// ExitMemoryTablePrimaryKeyConstraintOption is called when exiting the memoryTablePrimaryKeyConstraintOption production.
	ExitMemoryTablePrimaryKeyConstraintOption(c *MemoryTablePrimaryKeyConstraintOptionContext)

	// ExitWithBucket is called when exiting the withBucket production.
	ExitWithBucket(c *WithBucketContext)

	// ExitColumnForeignKeyConstraint is called when exiting the columnForeignKeyConstraint production.
	ExitColumnForeignKeyConstraint(c *ColumnForeignKeyConstraintContext)

	// ExitForeignKeyOnAction is called when exiting the foreignKeyOnAction production.
	ExitForeignKeyOnAction(c *ForeignKeyOnActionContext)

	// ExitForeignKeyOn is called when exiting the foreignKeyOn production.
	ExitForeignKeyOn(c *ForeignKeyOnContext)

	// ExitCheckConstraint is called when exiting the checkConstraint production.
	ExitCheckConstraint(c *CheckConstraintContext)

	// ExitColumnIndex is called when exiting the columnIndex production.
	ExitColumnIndex(c *ColumnIndexContext)

	// ExitWithIndexOption is called when exiting the withIndexOption production.
	ExitWithIndexOption(c *WithIndexOptionContext)

	// ExitIndexOnClause is called when exiting the indexOnClause production.
	ExitIndexOnClause(c *IndexOnClauseContext)

	// ExitOnDefault is called when exiting the onDefault production.
	ExitOnDefault(c *OnDefaultContext)

	// ExitFileStreamOn is called when exiting the fileStreamOn production.
	ExitFileStreamOn(c *FileStreamOnContext)

	// ExitColumnConstraints is called when exiting the columnConstraints production.
	ExitColumnConstraints(c *ColumnConstraintsContext)

	// ExitComputedColumnDefinition is called when exiting the computedColumnDefinition production.
	ExitComputedColumnDefinition(c *ComputedColumnDefinitionContext)

	// ExitColumnSetDefinition is called when exiting the columnSetDefinition production.
	ExitColumnSetDefinition(c *ColumnSetDefinitionContext)

	// ExitTableConstraint is called when exiting the tableConstraint production.
	ExitTableConstraint(c *TableConstraintContext)

	// ExitTablePrimaryConstraint is called when exiting the tablePrimaryConstraint production.
	ExitTablePrimaryConstraint(c *TablePrimaryConstraintContext)

	// ExitPrimaryKeyUnique is called when exiting the primaryKeyUnique production.
	ExitPrimaryKeyUnique(c *PrimaryKeyUniqueContext)

	// ExitDiskTablePrimaryConstraintOption is called when exiting the diskTablePrimaryConstraintOption production.
	ExitDiskTablePrimaryConstraintOption(c *DiskTablePrimaryConstraintOptionContext)

	// ExitMemoryTablePrimaryConstraintOption is called when exiting the memoryTablePrimaryConstraintOption production.
	ExitMemoryTablePrimaryConstraintOption(c *MemoryTablePrimaryConstraintOptionContext)

	// ExitHashWithBucket is called when exiting the hashWithBucket production.
	ExitHashWithBucket(c *HashWithBucketContext)

	// ExitTableForeignKeyConstraint is called when exiting the tableForeignKeyConstraint production.
	ExitTableForeignKeyConstraint(c *TableForeignKeyConstraintContext)

	// ExitTableIndex is called when exiting the tableIndex production.
	ExitTableIndex(c *TableIndexContext)

	// ExitIndexNameOption is called when exiting the indexNameOption production.
	ExitIndexNameOption(c *IndexNameOptionContext)

	// ExitIndexOptions is called when exiting the indexOptions production.
	ExitIndexOptions(c *IndexOptionsContext)

	// ExitPeriodClause is called when exiting the periodClause production.
	ExitPeriodClause(c *PeriodClauseContext)

	// ExitPartitionScheme is called when exiting the partitionScheme production.
	ExitPartitionScheme(c *PartitionSchemeContext)

	// ExitFileGroup is called when exiting the fileGroup production.
	ExitFileGroup(c *FileGroupContext)

	// ExitTableOptions is called when exiting the tableOptions production.
	ExitTableOptions(c *TableOptionsContext)

	// ExitTableOption is called when exiting the tableOption production.
	ExitTableOption(c *TableOptionContext)

	// ExitDataDelectionOption is called when exiting the dataDelectionOption production.
	ExitDataDelectionOption(c *DataDelectionOptionContext)

	// ExitTableStretchOptions is called when exiting the tableStretchOptions production.
	ExitTableStretchOptions(c *TableStretchOptionsContext)

	// ExitTableStretchOption is called when exiting the tableStretchOption production.
	ExitTableStretchOption(c *TableStretchOptionContext)

	// ExitMigrationState_ is called when exiting the migrationState_ production.
	ExitMigrationState_(c *MigrationState_Context)

	// ExitTableOperationOption is called when exiting the tableOperationOption production.
	ExitTableOperationOption(c *TableOperationOptionContext)

	// ExitDistributionOption is called when exiting the distributionOption production.
	ExitDistributionOption(c *DistributionOptionContext)

	// ExitDataWareHouseTableOption is called when exiting the dataWareHouseTableOption production.
	ExitDataWareHouseTableOption(c *DataWareHouseTableOptionContext)

	// ExitDataWareHousePartitionOption is called when exiting the dataWareHousePartitionOption production.
	ExitDataWareHousePartitionOption(c *DataWareHousePartitionOptionContext)

	// ExitCreateIndexSpecification is called when exiting the createIndexSpecification production.
	ExitCreateIndexSpecification(c *CreateIndexSpecificationContext)

	// ExitAlterDefinitionClause is called when exiting the alterDefinitionClause production.
	ExitAlterDefinitionClause(c *AlterDefinitionClauseContext)

	// ExitAddColumnSpecification is called when exiting the addColumnSpecification production.
	ExitAddColumnSpecification(c *AddColumnSpecificationContext)

	// ExitModifyColumnSpecification is called when exiting the modifyColumnSpecification production.
	ExitModifyColumnSpecification(c *ModifyColumnSpecificationContext)

	// ExitAlterColumnOperation is called when exiting the alterColumnOperation production.
	ExitAlterColumnOperation(c *AlterColumnOperationContext)

	// ExitAlterColumnAddOptions is called when exiting the alterColumnAddOptions production.
	ExitAlterColumnAddOptions(c *AlterColumnAddOptionsContext)

	// ExitAlterColumnAddOption is called when exiting the alterColumnAddOption production.
	ExitAlterColumnAddOption(c *AlterColumnAddOptionContext)

	// ExitConstraintForColumn is called when exiting the constraintForColumn production.
	ExitConstraintForColumn(c *ConstraintForColumnContext)

	// ExitGeneratedColumnNamesClause is called when exiting the generatedColumnNamesClause production.
	ExitGeneratedColumnNamesClause(c *GeneratedColumnNamesClauseContext)

	// ExitGeneratedColumnNameClause is called when exiting the generatedColumnNameClause production.
	ExitGeneratedColumnNameClause(c *GeneratedColumnNameClauseContext)

	// ExitGeneratedColumnName is called when exiting the generatedColumnName production.
	ExitGeneratedColumnName(c *GeneratedColumnNameContext)

	// ExitAlterDrop is called when exiting the alterDrop production.
	ExitAlterDrop(c *AlterDropContext)

	// ExitAlterTableDropConstraint is called when exiting the alterTableDropConstraint production.
	ExitAlterTableDropConstraint(c *AlterTableDropConstraintContext)

	// ExitDropConstraintName is called when exiting the dropConstraintName production.
	ExitDropConstraintName(c *DropConstraintNameContext)

	// ExitDropConstraintWithClause is called when exiting the dropConstraintWithClause production.
	ExitDropConstraintWithClause(c *DropConstraintWithClauseContext)

	// ExitDropConstraintOption is called when exiting the dropConstraintOption production.
	ExitDropConstraintOption(c *DropConstraintOptionContext)

	// ExitOnOffOption is called when exiting the onOffOption production.
	ExitOnOffOption(c *OnOffOptionContext)

	// ExitDropColumnSpecification is called when exiting the dropColumnSpecification production.
	ExitDropColumnSpecification(c *DropColumnSpecificationContext)

	// ExitDropIndexSpecification is called when exiting the dropIndexSpecification production.
	ExitDropIndexSpecification(c *DropIndexSpecificationContext)

	// ExitAlterCheckConstraint is called when exiting the alterCheckConstraint production.
	ExitAlterCheckConstraint(c *AlterCheckConstraintContext)

	// ExitAlterTableTrigger is called when exiting the alterTableTrigger production.
	ExitAlterTableTrigger(c *AlterTableTriggerContext)

	// ExitAlterSwitch is called when exiting the alterSwitch production.
	ExitAlterSwitch(c *AlterSwitchContext)

	// ExitAlterSet is called when exiting the alterSet production.
	ExitAlterSet(c *AlterSetContext)

	// ExitSetFileStreamClause is called when exiting the setFileStreamClause production.
	ExitSetFileStreamClause(c *SetFileStreamClauseContext)

	// ExitSetSystemVersionClause is called when exiting the setSystemVersionClause production.
	ExitSetSystemVersionClause(c *SetSystemVersionClauseContext)

	// ExitAlterSetOnClause is called when exiting the alterSetOnClause production.
	ExitAlterSetOnClause(c *AlterSetOnClauseContext)

	// ExitDataConsistencyCheckClause is called when exiting the dataConsistencyCheckClause production.
	ExitDataConsistencyCheckClause(c *DataConsistencyCheckClauseContext)

	// ExitHistoryRetentionPeriodClause is called when exiting the historyRetentionPeriodClause production.
	ExitHistoryRetentionPeriodClause(c *HistoryRetentionPeriodClauseContext)

	// ExitHistoryRetentionPeriod is called when exiting the historyRetentionPeriod production.
	ExitHistoryRetentionPeriod(c *HistoryRetentionPeriodContext)

	// ExitAlterTableTableIndex is called when exiting the alterTableTableIndex production.
	ExitAlterTableTableIndex(c *AlterTableTableIndexContext)

	// ExitIndexWithName is called when exiting the indexWithName production.
	ExitIndexWithName(c *IndexWithNameContext)

	// ExitIndexNonClusterClause is called when exiting the indexNonClusterClause production.
	ExitIndexNonClusterClause(c *IndexNonClusterClauseContext)

	// ExitAlterTableIndexOnClause is called when exiting the alterTableIndexOnClause production.
	ExitAlterTableIndexOnClause(c *AlterTableIndexOnClauseContext)

	// ExitIndexClusterClause is called when exiting the indexClusterClause production.
	ExitIndexClusterClause(c *IndexClusterClauseContext)

	// ExitAlterTableOption is called when exiting the alterTableOption production.
	ExitAlterTableOption(c *AlterTableOptionContext)

	// ExitOnHistoryTableClause is called when exiting the onHistoryTableClause production.
	ExitOnHistoryTableClause(c *OnHistoryTableClauseContext)

	// ExitIfExist is called when exiting the ifExist production.
	ExitIfExist(c *IfExistContext)

	// ExitCreateDatabaseClause is called when exiting the createDatabaseClause production.
	ExitCreateDatabaseClause(c *CreateDatabaseClauseContext)

	// ExitFileDefinitionClause is called when exiting the fileDefinitionClause production.
	ExitFileDefinitionClause(c *FileDefinitionClauseContext)

	// ExitDatabaseOption is called when exiting the databaseOption production.
	ExitDatabaseOption(c *DatabaseOptionContext)

	// ExitFileStreamOption is called when exiting the fileStreamOption production.
	ExitFileStreamOption(c *FileStreamOptionContext)

	// ExitFileSpec is called when exiting the fileSpec production.
	ExitFileSpec(c *FileSpecContext)

	// ExitDatabaseFileSpecOption is called when exiting the databaseFileSpecOption production.
	ExitDatabaseFileSpecOption(c *DatabaseFileSpecOptionContext)

	// ExitDatabaseFileGroup is called when exiting the databaseFileGroup production.
	ExitDatabaseFileGroup(c *DatabaseFileGroupContext)

	// ExitDatabaseFileGroupContains is called when exiting the databaseFileGroupContains production.
	ExitDatabaseFileGroupContains(c *DatabaseFileGroupContainsContext)

	// ExitDatabaseLogOns is called when exiting the databaseLogOns production.
	ExitDatabaseLogOns(c *DatabaseLogOnsContext)

	// ExitDeclareVariable is called when exiting the declareVariable production.
	ExitDeclareVariable(c *DeclareVariableContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitTableVariable is called when exiting the tableVariable production.
	ExitTableVariable(c *TableVariableContext)

	// ExitVariTableTypeDefinition is called when exiting the variTableTypeDefinition production.
	ExitVariTableTypeDefinition(c *VariTableTypeDefinitionContext)

	// ExitTableVariableClause is called when exiting the tableVariableClause production.
	ExitTableVariableClause(c *TableVariableClauseContext)

	// ExitVariableTableColumnDefinition is called when exiting the variableTableColumnDefinition production.
	ExitVariableTableColumnDefinition(c *VariableTableColumnDefinitionContext)

	// ExitVariableTableColumnConstraint is called when exiting the variableTableColumnConstraint production.
	ExitVariableTableColumnConstraint(c *VariableTableColumnConstraintContext)

	// ExitVariableTableConstraint is called when exiting the variableTableConstraint production.
	ExitVariableTableConstraint(c *VariableTableConstraintContext)

	// ExitSetVariable is called when exiting the setVariable production.
	ExitSetVariable(c *SetVariableContext)

	// ExitSetVariableClause is called when exiting the setVariableClause production.
	ExitSetVariableClause(c *SetVariableClauseContext)

	// ExitCursorVariable is called when exiting the cursorVariable production.
	ExitCursorVariable(c *CursorVariableContext)

	// ExitCursorClause is called when exiting the cursorClause production.
	ExitCursorClause(c *CursorClauseContext)

	// ExitCompoundOperation is called when exiting the compoundOperation production.
	ExitCompoundOperation(c *CompoundOperationContext)

	// ExitFuncParameters is called when exiting the funcParameters production.
	ExitFuncParameters(c *FuncParametersContext)

	// ExitFuncReturns is called when exiting the funcReturns production.
	ExitFuncReturns(c *FuncReturnsContext)

	// ExitFuncMutiReturn is called when exiting the funcMutiReturn production.
	ExitFuncMutiReturn(c *FuncMutiReturnContext)

	// ExitFuncInlineReturn is called when exiting the funcInlineReturn production.
	ExitFuncInlineReturn(c *FuncInlineReturnContext)

	// ExitFuncScalarReturn is called when exiting the funcScalarReturn production.
	ExitFuncScalarReturn(c *FuncScalarReturnContext)

	// ExitTableTypeDefinition is called when exiting the tableTypeDefinition production.
	ExitTableTypeDefinition(c *TableTypeDefinitionContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitFunctionOption is called when exiting the functionOption production.
	ExitFunctionOption(c *FunctionOptionContext)

	// ExitValidStatement is called when exiting the validStatement production.
	ExitValidStatement(c *ValidStatementContext)

	// ExitProcParameters is called when exiting the procParameters production.
	ExitProcParameters(c *ProcParametersContext)

	// ExitProcParameter is called when exiting the procParameter production.
	ExitProcParameter(c *ProcParameterContext)

	// ExitCreateOrAlterProcClause is called when exiting the createOrAlterProcClause production.
	ExitCreateOrAlterProcClause(c *CreateOrAlterProcClauseContext)

	// ExitWithCreateProcOption is called when exiting the withCreateProcOption production.
	ExitWithCreateProcOption(c *WithCreateProcOptionContext)

	// ExitProcOption is called when exiting the procOption production.
	ExitProcOption(c *ProcOptionContext)

	// ExitProcAsClause is called when exiting the procAsClause production.
	ExitProcAsClause(c *ProcAsClauseContext)

	// ExitProcSetOption is called when exiting the procSetOption production.
	ExitProcSetOption(c *ProcSetOptionContext)

	// ExitCreateOrAlterViewClause is called when exiting the createOrAlterViewClause production.
	ExitCreateOrAlterViewClause(c *CreateOrAlterViewClauseContext)

	// ExitViewAttribute is called when exiting the viewAttribute production.
	ExitViewAttribute(c *ViewAttributeContext)

	// ExitWithCommonTableExpr is called when exiting the withCommonTableExpr production.
	ExitWithCommonTableExpr(c *WithCommonTableExprContext)

	// ExitCommonTableExpr is called when exiting the commonTableExpr production.
	ExitCommonTableExpr(c *CommonTableExprContext)

	// ExitCreateTriggerClause is called when exiting the createTriggerClause production.
	ExitCreateTriggerClause(c *CreateTriggerClauseContext)

	// ExitDmlTriggerOption is called when exiting the dmlTriggerOption production.
	ExitDmlTriggerOption(c *DmlTriggerOptionContext)

	// ExitMethodSpecifier is called when exiting the methodSpecifier production.
	ExitMethodSpecifier(c *MethodSpecifierContext)

	// ExitTriggerTarget is called when exiting the triggerTarget production.
	ExitTriggerTarget(c *TriggerTargetContext)

	// ExitCreateOrAlterSequenceClause is called when exiting the createOrAlterSequenceClause production.
	ExitCreateOrAlterSequenceClause(c *CreateOrAlterSequenceClauseContext)

	// ExitCreateIndexClause is called when exiting the createIndexClause production.
	ExitCreateIndexClause(c *CreateIndexClauseContext)

	// ExitFilterPredicate is called when exiting the filterPredicate production.
	ExitFilterPredicate(c *FilterPredicateContext)

	// ExitConjunct is called when exiting the conjunct production.
	ExitConjunct(c *ConjunctContext)

	// ExitAlterIndexClause is called when exiting the alterIndexClause production.
	ExitAlterIndexClause(c *AlterIndexClauseContext)

	// ExitRelationalIndexOption is called when exiting the relationalIndexOption production.
	ExitRelationalIndexOption(c *RelationalIndexOptionContext)

	// ExitPartitionNumberRange is called when exiting the partitionNumberRange production.
	ExitPartitionNumberRange(c *PartitionNumberRangeContext)

	// ExitReorganizeOption is called when exiting the reorganizeOption production.
	ExitReorganizeOption(c *ReorganizeOptionContext)

	// ExitSetIndexOption is called when exiting the setIndexOption production.
	ExitSetIndexOption(c *SetIndexOptionContext)

	// ExitResumableIndexOptions is called when exiting the resumableIndexOptions production.
	ExitResumableIndexOptions(c *ResumableIndexOptionsContext)

	// ExitAlterDatabaseClause is called when exiting the alterDatabaseClause production.
	ExitAlterDatabaseClause(c *AlterDatabaseClauseContext)

	// ExitAddSecondaryOption is called when exiting the addSecondaryOption production.
	ExitAddSecondaryOption(c *AddSecondaryOptionContext)

	// ExitEditionOptions is called when exiting the editionOptions production.
	ExitEditionOptions(c *EditionOptionsContext)

	// ExitServiceObjective is called when exiting the serviceObjective production.
	ExitServiceObjective(c *ServiceObjectiveContext)

	// ExitAlterDatabaseOptionSpec is called when exiting the alterDatabaseOptionSpec production.
	ExitAlterDatabaseOptionSpec(c *AlterDatabaseOptionSpecContext)

	// ExitFileAndFilegroupOptions is called when exiting the fileAndFilegroupOptions production.
	ExitFileAndFilegroupOptions(c *FileAndFilegroupOptionsContext)

	// ExitAddOrModifyFilegroups is called when exiting the addOrModifyFilegroups production.
	ExitAddOrModifyFilegroups(c *AddOrModifyFilegroupsContext)

	// ExitFilegroupUpdatabilityOption is called when exiting the filegroupUpdatabilityOption production.
	ExitFilegroupUpdatabilityOption(c *FilegroupUpdatabilityOptionContext)

	// ExitAddOrModifyFiles is called when exiting the addOrModifyFiles production.
	ExitAddOrModifyFiles(c *AddOrModifyFilesContext)

	// ExitAcceleratedDatabaseRecovery is called when exiting the acceleratedDatabaseRecovery production.
	ExitAcceleratedDatabaseRecovery(c *AcceleratedDatabaseRecoveryContext)

	// ExitAutoOption is called when exiting the autoOption production.
	ExitAutoOption(c *AutoOptionContext)

	// ExitAutomaticTuningOption is called when exiting the automaticTuningOption production.
	ExitAutomaticTuningOption(c *AutomaticTuningOptionContext)

	// ExitChangeTrackingOption is called when exiting the changeTrackingOption production.
	ExitChangeTrackingOption(c *ChangeTrackingOptionContext)

	// ExitChangeTrackingOptionList is called when exiting the changeTrackingOptionList production.
	ExitChangeTrackingOptionList(c *ChangeTrackingOptionListContext)

	// ExitCursorOption is called when exiting the cursorOption production.
	ExitCursorOption(c *CursorOptionContext)

	// ExitExternalAccessOption is called when exiting the externalAccessOption production.
	ExitExternalAccessOption(c *ExternalAccessOptionContext)

	// ExitQueryStoreOptions is called when exiting the queryStoreOptions production.
	ExitQueryStoreOptions(c *QueryStoreOptionsContext)

	// ExitQueryStoreOptionList is called when exiting the queryStoreOptionList production.
	ExitQueryStoreOptionList(c *QueryStoreOptionListContext)

	// ExitQueryCapturePolicyOptionList is called when exiting the queryCapturePolicyOptionList production.
	ExitQueryCapturePolicyOptionList(c *QueryCapturePolicyOptionListContext)

	// ExitRecoveryOption is called when exiting the recoveryOption production.
	ExitRecoveryOption(c *RecoveryOptionContext)

	// ExitSqlOption is called when exiting the sqlOption production.
	ExitSqlOption(c *SqlOptionContext)

	// ExitSnapshotOption is called when exiting the snapshotOption production.
	ExitSnapshotOption(c *SnapshotOptionContext)

	// ExitServiceBrokerOption is called when exiting the serviceBrokerOption production.
	ExitServiceBrokerOption(c *ServiceBrokerOptionContext)

	// ExitTargetRecoveryTimeOption is called when exiting the targetRecoveryTimeOption production.
	ExitTargetRecoveryTimeOption(c *TargetRecoveryTimeOptionContext)

	// ExitTermination is called when exiting the termination production.
	ExitTermination(c *TerminationContext)

	// ExitCreateServiceClause is called when exiting the createServiceClause production.
	ExitCreateServiceClause(c *CreateServiceClauseContext)

	// ExitAlterServiceClause is called when exiting the alterServiceClause production.
	ExitAlterServiceClause(c *AlterServiceClauseContext)

	// ExitAlterServiceOptArg is called when exiting the alterServiceOptArg production.
	ExitAlterServiceOptArg(c *AlterServiceOptArgContext)

	// ExitSchemaNameClause is called when exiting the schemaNameClause production.
	ExitSchemaNameClause(c *SchemaNameClauseContext)

	// ExitSchemaElement is called when exiting the schemaElement production.
	ExitSchemaElement(c *SchemaElementContext)

	// ExitCreateTableAsSelectClause is called when exiting the createTableAsSelectClause production.
	ExitCreateTableAsSelectClause(c *CreateTableAsSelectClauseContext)

	// ExitCreateTableAsSelect is called when exiting the createTableAsSelect production.
	ExitCreateTableAsSelect(c *CreateTableAsSelectContext)

	// ExitCreateRemoteTableAsSelect is called when exiting the createRemoteTableAsSelect production.
	ExitCreateRemoteTableAsSelect(c *CreateRemoteTableAsSelectContext)

	// ExitWithDistributionOption is called when exiting the withDistributionOption production.
	ExitWithDistributionOption(c *WithDistributionOptionContext)

	// ExitOptionQueryHintClause is called when exiting the optionQueryHintClause production.
	ExitOptionQueryHintClause(c *OptionQueryHintClauseContext)

	// ExitGrant is called when exiting the grant production.
	ExitGrant(c *GrantContext)

	// ExitRevoke is called when exiting the revoke production.
	ExitRevoke(c *RevokeContext)

	// ExitDeny is called when exiting the deny production.
	ExitDeny(c *DenyContext)

	// ExitClassPrivilegesClause is called when exiting the classPrivilegesClause production.
	ExitClassPrivilegesClause(c *ClassPrivilegesClauseContext)

	// ExitClassTypePrivilegesClause is called when exiting the classTypePrivilegesClause production.
	ExitClassTypePrivilegesClause(c *ClassTypePrivilegesClauseContext)

	// ExitOptionForClause is called when exiting the optionForClause production.
	ExitOptionForClause(c *OptionForClauseContext)

	// ExitClassPrivileges is called when exiting the classPrivileges production.
	ExitClassPrivileges(c *ClassPrivilegesContext)

	// ExitOnClassClause is called when exiting the onClassClause production.
	ExitOnClassClause(c *OnClassClauseContext)

	// ExitClassTypePrivileges is called when exiting the classTypePrivileges production.
	ExitClassTypePrivileges(c *ClassTypePrivilegesContext)

	// ExitOnClassTypeClause is called when exiting the onClassTypeClause production.
	ExitOnClassTypeClause(c *OnClassTypeClauseContext)

	// ExitPrivilegeType is called when exiting the privilegeType production.
	ExitPrivilegeType(c *PrivilegeTypeContext)

	// ExitBasicPermission is called when exiting the basicPermission production.
	ExitBasicPermission(c *BasicPermissionContext)

	// ExitObjectPermission is called when exiting the objectPermission production.
	ExitObjectPermission(c *ObjectPermissionContext)

	// ExitServerPermission is called when exiting the serverPermission production.
	ExitServerPermission(c *ServerPermissionContext)

	// ExitServerPrincipalPermission is called when exiting the serverPrincipalPermission production.
	ExitServerPrincipalPermission(c *ServerPrincipalPermissionContext)

	// ExitDatabasePermission is called when exiting the databasePermission production.
	ExitDatabasePermission(c *DatabasePermissionContext)

	// ExitDatabasePrincipalPermission is called when exiting the databasePrincipalPermission production.
	ExitDatabasePrincipalPermission(c *DatabasePrincipalPermissionContext)

	// ExitSchemaPermission is called when exiting the schemaPermission production.
	ExitSchemaPermission(c *SchemaPermissionContext)

	// ExitServiceBrokerPermission is called when exiting the serviceBrokerPermission production.
	ExitServiceBrokerPermission(c *ServiceBrokerPermissionContext)

	// ExitEndpointPermission is called when exiting the endpointPermission production.
	ExitEndpointPermission(c *EndpointPermissionContext)

	// ExitCertificatePermission is called when exiting the certificatePermission production.
	ExitCertificatePermission(c *CertificatePermissionContext)

	// ExitSymmetricKeyPermission is called when exiting the symmetricKeyPermission production.
	ExitSymmetricKeyPermission(c *SymmetricKeyPermissionContext)

	// ExitAsymmetricKeyPermission is called when exiting the asymmetricKeyPermission production.
	ExitAsymmetricKeyPermission(c *AsymmetricKeyPermissionContext)

	// ExitAssemblyPermission is called when exiting the assemblyPermission production.
	ExitAssemblyPermission(c *AssemblyPermissionContext)

	// ExitAvailabilityGroupPermission is called when exiting the availabilityGroupPermission production.
	ExitAvailabilityGroupPermission(c *AvailabilityGroupPermissionContext)

	// ExitFullTextPermission is called when exiting the fullTextPermission production.
	ExitFullTextPermission(c *FullTextPermissionContext)

	// ExitClass_ is called when exiting the class_ production.
	ExitClass_(c *Class_Context)

	// ExitClassType is called when exiting the classType production.
	ExitClassType(c *ClassTypeContext)

	// ExitRoleClause is called when exiting the roleClause production.
	ExitRoleClause(c *RoleClauseContext)

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

	// ExitCreateLogin is called when exiting the createLogin production.
	ExitCreateLogin(c *CreateLoginContext)

	// ExitDropLogin is called when exiting the dropLogin production.
	ExitDropLogin(c *DropLoginContext)

	// ExitAlterLogin is called when exiting the alterLogin production.
	ExitAlterLogin(c *AlterLoginContext)

	// ExitSetTransaction is called when exiting the setTransaction production.
	ExitSetTransaction(c *SetTransactionContext)

	// ExitSetImplicitTransactions is called when exiting the setImplicitTransactions production.
	ExitSetImplicitTransactions(c *SetImplicitTransactionsContext)

	// ExitImplicitTransactionsValue is called when exiting the implicitTransactionsValue production.
	ExitImplicitTransactionsValue(c *ImplicitTransactionsValueContext)

	// ExitBeginTransaction is called when exiting the beginTransaction production.
	ExitBeginTransaction(c *BeginTransactionContext)

	// ExitBeginDistributedTransaction is called when exiting the beginDistributedTransaction production.
	ExitBeginDistributedTransaction(c *BeginDistributedTransactionContext)

	// ExitCommit is called when exiting the commit production.
	ExitCommit(c *CommitContext)

	// ExitCommitWork is called when exiting the commitWork production.
	ExitCommitWork(c *CommitWorkContext)

	// ExitRollback is called when exiting the rollback production.
	ExitRollback(c *RollbackContext)

	// ExitRollbackWork is called when exiting the rollbackWork production.
	ExitRollbackWork(c *RollbackWorkContext)

	// ExitSavepoint is called when exiting the savepoint production.
	ExitSavepoint(c *SavepointContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitExplain is called when exiting the explain production.
	ExitExplain(c *ExplainContext)

	// ExitExplainableStatement is called when exiting the explainableStatement production.
	ExitExplainableStatement(c *ExplainableStatementContext)
}

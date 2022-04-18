// Code generated from E:/GoProject/src/github.com/jiunx/xsqlparser/dialect/sql92/grammer\SQL92Statement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQL92Statement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SQL92StatementListener is a complete listener for a parse tree produced by SQL92StatementParser.
type SQL92StatementListener interface {
	antlr.ParseTreeListener

	// EnterExecute is called when entering the execute production.
	EnterExecute(c *ExecuteContext)

	// EnterInsert is called when entering the insert production.
	EnterInsert(c *InsertContext)

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

	// EnterBlobValue is called when entering the blobValue production.
	EnterBlobValue(c *BlobValueContext)

	// EnterDelete is called when entering the delete production.
	EnterDelete(c *DeleteContext)

	// EnterSingleTableClause is called when entering the singleTableClause production.
	EnterSingleTableClause(c *SingleTableClauseContext)

	// EnterSelect is called when entering the select production.
	EnterSelect(c *SelectContext)

	// EnterUnionClause is called when entering the unionClause production.
	EnterUnionClause(c *UnionClauseContext)

	// EnterSelectClause is called when entering the selectClause production.
	EnterSelectClause(c *SelectClauseContext)

	// EnterSelectSpecification is called when entering the selectSpecification production.
	EnterSelectSpecification(c *SelectSpecificationContext)

	// EnterDuplicateSpecification is called when entering the duplicateSpecification production.
	EnterDuplicateSpecification(c *DuplicateSpecificationContext)

	// EnterProjections is called when entering the projections production.
	EnterProjections(c *ProjectionsContext)

	// EnterProjection is called when entering the projection production.
	EnterProjection(c *ProjectionContext)

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

	// EnterEscapedTableReference is called when entering the escapedTableReference production.
	EnterEscapedTableReference(c *EscapedTableReferenceContext)

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

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterLimitRowCount is called when entering the limitRowCount production.
	EnterLimitRowCount(c *LimitRowCountContext)

	// EnterLimitOffset is called when entering the limitOffset production.
	EnterLimitOffset(c *LimitOffsetContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

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

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterSchemaName is called when entering the schemaName production.
	EnterSchemaName(c *SchemaNameContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterViewName is called when entering the viewName production.
	EnterViewName(c *ViewNameContext)

	// EnterOwner is called when entering the owner production.
	EnterOwner(c *OwnerContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterConstraintName is called when entering the constraintName production.
	EnterConstraintName(c *ConstraintNameContext)

	// EnterColumnNames is called when entering the columnNames production.
	EnterColumnNames(c *ColumnNamesContext)

	// EnterTableNames is called when entering the tableNames production.
	EnterTableNames(c *TableNamesContext)

	// EnterCharacterSetName is called when entering the characterSetName production.
	EnterCharacterSetName(c *CharacterSetNameContext)

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

	// EnterConvertFunction is called when entering the convertFunction production.
	EnterConvertFunction(c *ConvertFunctionContext)

	// EnterPositionFunction is called when entering the positionFunction production.
	EnterPositionFunction(c *PositionFunctionContext)

	// EnterSubstringFunction is called when entering the substringFunction production.
	EnterSubstringFunction(c *SubstringFunctionContext)

	// EnterExtractFunction is called when entering the extractFunction production.
	EnterExtractFunction(c *ExtractFunctionContext)

	// EnterTrimFunction is called when entering the trimFunction production.
	EnterTrimFunction(c *TrimFunctionContext)

	// EnterRegularFunction is called when entering the regularFunction production.
	EnterRegularFunction(c *RegularFunctionContext)

	// EnterRegularFunctionName is called when entering the regularFunctionName production.
	EnterRegularFunctionName(c *RegularFunctionNameContext)

	// EnterMatchExpression is called when entering the matchExpression production.
	EnterMatchExpression(c *MatchExpressionContext)

	// EnterCaseExpression is called when entering the caseExpression production.
	EnterCaseExpression(c *CaseExpressionContext)

	// EnterCaseWhen is called when entering the caseWhen production.
	EnterCaseWhen(c *CaseWhenContext)

	// EnterCaseElse is called when entering the caseElse production.
	EnterCaseElse(c *CaseElseContext)

	// EnterIntervalExpression is called when entering the intervalExpression production.
	EnterIntervalExpression(c *IntervalExpressionContext)

	// EnterIntervalUnit is called when entering the intervalUnit production.
	EnterIntervalUnit(c *IntervalUnitContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderByItem is called when entering the orderByItem production.
	EnterOrderByItem(c *OrderByItemContext)

	// EnterDataType is called when entering the dataType production.
	EnterDataType(c *DataTypeContext)

	// EnterDataTypeName is called when entering the dataTypeName production.
	EnterDataTypeName(c *DataTypeNameContext)

	// EnterDataTypeLength is called when entering the dataTypeLength production.
	EnterDataTypeLength(c *DataTypeLengthContext)

	// EnterCharacterSet is called when entering the characterSet production.
	EnterCharacterSet(c *CharacterSetContext)

	// EnterCollateClause is called when entering the collateClause production.
	EnterCollateClause(c *CollateClauseContext)

	// EnterIgnoredIdentifier is called when entering the ignoredIdentifier production.
	EnterIgnoredIdentifier(c *IgnoredIdentifierContext)

	// EnterDropBehaviour is called when entering the dropBehaviour production.
	EnterDropBehaviour(c *DropBehaviourContext)

	// EnterCreateTable is called when entering the createTable production.
	EnterCreateTable(c *CreateTableContext)

	// EnterAlterTable is called when entering the alterTable production.
	EnterAlterTable(c *AlterTableContext)

	// EnterDropTable is called when entering the dropTable production.
	EnterDropTable(c *DropTableContext)

	// EnterCreateDatabase is called when entering the createDatabase production.
	EnterCreateDatabase(c *CreateDatabaseContext)

	// EnterDropDatabase is called when entering the dropDatabase production.
	EnterDropDatabase(c *DropDatabaseContext)

	// EnterCreateView is called when entering the createView production.
	EnterCreateView(c *CreateViewContext)

	// EnterDropView is called when entering the dropView production.
	EnterDropView(c *DropViewContext)

	// EnterCreateTableSpecification is called when entering the createTableSpecification production.
	EnterCreateTableSpecification(c *CreateTableSpecificationContext)

	// EnterCreateDefinitionClause is called when entering the createDefinitionClause production.
	EnterCreateDefinitionClause(c *CreateDefinitionClauseContext)

	// EnterCreateDatabaseSpecification_ is called when entering the createDatabaseSpecification_ production.
	EnterCreateDatabaseSpecification_(c *CreateDatabaseSpecification_Context)

	// EnterCreateDefinition is called when entering the createDefinition production.
	EnterCreateDefinition(c *CreateDefinitionContext)

	// EnterColumnDefinition is called when entering the columnDefinition production.
	EnterColumnDefinition(c *ColumnDefinitionContext)

	// EnterDataTypeOption is called when entering the dataTypeOption production.
	EnterDataTypeOption(c *DataTypeOptionContext)

	// EnterCheckConstraintDefinition is called when entering the checkConstraintDefinition production.
	EnterCheckConstraintDefinition(c *CheckConstraintDefinitionContext)

	// EnterReferenceDefinition is called when entering the referenceDefinition production.
	EnterReferenceDefinition(c *ReferenceDefinitionContext)

	// EnterReferenceOption is called when entering the referenceOption production.
	EnterReferenceOption(c *ReferenceOptionContext)

	// EnterKeyParts is called when entering the keyParts production.
	EnterKeyParts(c *KeyPartsContext)

	// EnterKeyPart is called when entering the keyPart production.
	EnterKeyPart(c *KeyPartContext)

	// EnterConstraintDefinition is called when entering the constraintDefinition production.
	EnterConstraintDefinition(c *ConstraintDefinitionContext)

	// EnterPrimaryKeyOption is called when entering the primaryKeyOption production.
	EnterPrimaryKeyOption(c *PrimaryKeyOptionContext)

	// EnterPrimaryKey is called when entering the primaryKey production.
	EnterPrimaryKey(c *PrimaryKeyContext)

	// EnterUniqueOption is called when entering the uniqueOption production.
	EnterUniqueOption(c *UniqueOptionContext)

	// EnterForeignKeyOption is called when entering the foreignKeyOption production.
	EnterForeignKeyOption(c *ForeignKeyOptionContext)

	// EnterCreateLikeClause is called when entering the createLikeClause production.
	EnterCreateLikeClause(c *CreateLikeClauseContext)

	// EnterAlterDefinitionClause is called when entering the alterDefinitionClause production.
	EnterAlterDefinitionClause(c *AlterDefinitionClauseContext)

	// EnterAddColumnSpecification is called when entering the addColumnSpecification production.
	EnterAddColumnSpecification(c *AddColumnSpecificationContext)

	// EnterModifyColumnSpecification is called when entering the modifyColumnSpecification production.
	EnterModifyColumnSpecification(c *ModifyColumnSpecificationContext)

	// EnterDropColumnSpecification is called when entering the dropColumnSpecification production.
	EnterDropColumnSpecification(c *DropColumnSpecificationContext)

	// EnterAddConstraintSpecification is called when entering the addConstraintSpecification production.
	EnterAddConstraintSpecification(c *AddConstraintSpecificationContext)

	// EnterDropConstraintSpecification is called when entering the dropConstraintSpecification production.
	EnterDropConstraintSpecification(c *DropConstraintSpecificationContext)

	// EnterSetTransaction is called when entering the setTransaction production.
	EnterSetTransaction(c *SetTransactionContext)

	// EnterCommit is called when entering the commit production.
	EnterCommit(c *CommitContext)

	// EnterRollback is called when entering the rollback production.
	EnterRollback(c *RollbackContext)

	// EnterLevelOfIsolation is called when entering the levelOfIsolation production.
	EnterLevelOfIsolation(c *LevelOfIsolationContext)

	// EnterGrant is called when entering the grant production.
	EnterGrant(c *GrantContext)

	// EnterRevoke is called when entering the revoke production.
	EnterRevoke(c *RevokeContext)

	// EnterPrivilegeClause is called when entering the privilegeClause production.
	EnterPrivilegeClause(c *PrivilegeClauseContext)

	// EnterPrivileges is called when entering the privileges production.
	EnterPrivileges(c *PrivilegesContext)

	// EnterPrivilegeType is called when entering the privilegeType production.
	EnterPrivilegeType(c *PrivilegeTypeContext)

	// EnterGrantee is called when entering the grantee production.
	EnterGrantee(c *GranteeContext)

	// EnterOnObjectClause is called when entering the onObjectClause production.
	EnterOnObjectClause(c *OnObjectClauseContext)

	// EnterObjectType is called when entering the objectType production.
	EnterObjectType(c *ObjectTypeContext)

	// EnterPrivilegeLevel is called when entering the privilegeLevel production.
	EnterPrivilegeLevel(c *PrivilegeLevelContext)

	// ExitExecute is called when exiting the execute production.
	ExitExecute(c *ExecuteContext)

	// ExitInsert is called when exiting the insert production.
	ExitInsert(c *InsertContext)

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

	// ExitBlobValue is called when exiting the blobValue production.
	ExitBlobValue(c *BlobValueContext)

	// ExitDelete is called when exiting the delete production.
	ExitDelete(c *DeleteContext)

	// ExitSingleTableClause is called when exiting the singleTableClause production.
	ExitSingleTableClause(c *SingleTableClauseContext)

	// ExitSelect is called when exiting the select production.
	ExitSelect(c *SelectContext)

	// ExitUnionClause is called when exiting the unionClause production.
	ExitUnionClause(c *UnionClauseContext)

	// ExitSelectClause is called when exiting the selectClause production.
	ExitSelectClause(c *SelectClauseContext)

	// ExitSelectSpecification is called when exiting the selectSpecification production.
	ExitSelectSpecification(c *SelectSpecificationContext)

	// ExitDuplicateSpecification is called when exiting the duplicateSpecification production.
	ExitDuplicateSpecification(c *DuplicateSpecificationContext)

	// ExitProjections is called when exiting the projections production.
	ExitProjections(c *ProjectionsContext)

	// ExitProjection is called when exiting the projection production.
	ExitProjection(c *ProjectionContext)

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

	// ExitEscapedTableReference is called when exiting the escapedTableReference production.
	ExitEscapedTableReference(c *EscapedTableReferenceContext)

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

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitLimitRowCount is called when exiting the limitRowCount production.
	ExitLimitRowCount(c *LimitRowCountContext)

	// ExitLimitOffset is called when exiting the limitOffset production.
	ExitLimitOffset(c *LimitOffsetContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

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

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitSchemaName is called when exiting the schemaName production.
	ExitSchemaName(c *SchemaNameContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitViewName is called when exiting the viewName production.
	ExitViewName(c *ViewNameContext)

	// ExitOwner is called when exiting the owner production.
	ExitOwner(c *OwnerContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitConstraintName is called when exiting the constraintName production.
	ExitConstraintName(c *ConstraintNameContext)

	// ExitColumnNames is called when exiting the columnNames production.
	ExitColumnNames(c *ColumnNamesContext)

	// ExitTableNames is called when exiting the tableNames production.
	ExitTableNames(c *TableNamesContext)

	// ExitCharacterSetName is called when exiting the characterSetName production.
	ExitCharacterSetName(c *CharacterSetNameContext)

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

	// ExitConvertFunction is called when exiting the convertFunction production.
	ExitConvertFunction(c *ConvertFunctionContext)

	// ExitPositionFunction is called when exiting the positionFunction production.
	ExitPositionFunction(c *PositionFunctionContext)

	// ExitSubstringFunction is called when exiting the substringFunction production.
	ExitSubstringFunction(c *SubstringFunctionContext)

	// ExitExtractFunction is called when exiting the extractFunction production.
	ExitExtractFunction(c *ExtractFunctionContext)

	// ExitTrimFunction is called when exiting the trimFunction production.
	ExitTrimFunction(c *TrimFunctionContext)

	// ExitRegularFunction is called when exiting the regularFunction production.
	ExitRegularFunction(c *RegularFunctionContext)

	// ExitRegularFunctionName is called when exiting the regularFunctionName production.
	ExitRegularFunctionName(c *RegularFunctionNameContext)

	// ExitMatchExpression is called when exiting the matchExpression production.
	ExitMatchExpression(c *MatchExpressionContext)

	// ExitCaseExpression is called when exiting the caseExpression production.
	ExitCaseExpression(c *CaseExpressionContext)

	// ExitCaseWhen is called when exiting the caseWhen production.
	ExitCaseWhen(c *CaseWhenContext)

	// ExitCaseElse is called when exiting the caseElse production.
	ExitCaseElse(c *CaseElseContext)

	// ExitIntervalExpression is called when exiting the intervalExpression production.
	ExitIntervalExpression(c *IntervalExpressionContext)

	// ExitIntervalUnit is called when exiting the intervalUnit production.
	ExitIntervalUnit(c *IntervalUnitContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderByItem is called when exiting the orderByItem production.
	ExitOrderByItem(c *OrderByItemContext)

	// ExitDataType is called when exiting the dataType production.
	ExitDataType(c *DataTypeContext)

	// ExitDataTypeName is called when exiting the dataTypeName production.
	ExitDataTypeName(c *DataTypeNameContext)

	// ExitDataTypeLength is called when exiting the dataTypeLength production.
	ExitDataTypeLength(c *DataTypeLengthContext)

	// ExitCharacterSet is called when exiting the characterSet production.
	ExitCharacterSet(c *CharacterSetContext)

	// ExitCollateClause is called when exiting the collateClause production.
	ExitCollateClause(c *CollateClauseContext)

	// ExitIgnoredIdentifier is called when exiting the ignoredIdentifier production.
	ExitIgnoredIdentifier(c *IgnoredIdentifierContext)

	// ExitDropBehaviour is called when exiting the dropBehaviour production.
	ExitDropBehaviour(c *DropBehaviourContext)

	// ExitCreateTable is called when exiting the createTable production.
	ExitCreateTable(c *CreateTableContext)

	// ExitAlterTable is called when exiting the alterTable production.
	ExitAlterTable(c *AlterTableContext)

	// ExitDropTable is called when exiting the dropTable production.
	ExitDropTable(c *DropTableContext)

	// ExitCreateDatabase is called when exiting the createDatabase production.
	ExitCreateDatabase(c *CreateDatabaseContext)

	// ExitDropDatabase is called when exiting the dropDatabase production.
	ExitDropDatabase(c *DropDatabaseContext)

	// ExitCreateView is called when exiting the createView production.
	ExitCreateView(c *CreateViewContext)

	// ExitDropView is called when exiting the dropView production.
	ExitDropView(c *DropViewContext)

	// ExitCreateTableSpecification is called when exiting the createTableSpecification production.
	ExitCreateTableSpecification(c *CreateTableSpecificationContext)

	// ExitCreateDefinitionClause is called when exiting the createDefinitionClause production.
	ExitCreateDefinitionClause(c *CreateDefinitionClauseContext)

	// ExitCreateDatabaseSpecification_ is called when exiting the createDatabaseSpecification_ production.
	ExitCreateDatabaseSpecification_(c *CreateDatabaseSpecification_Context)

	// ExitCreateDefinition is called when exiting the createDefinition production.
	ExitCreateDefinition(c *CreateDefinitionContext)

	// ExitColumnDefinition is called when exiting the columnDefinition production.
	ExitColumnDefinition(c *ColumnDefinitionContext)

	// ExitDataTypeOption is called when exiting the dataTypeOption production.
	ExitDataTypeOption(c *DataTypeOptionContext)

	// ExitCheckConstraintDefinition is called when exiting the checkConstraintDefinition production.
	ExitCheckConstraintDefinition(c *CheckConstraintDefinitionContext)

	// ExitReferenceDefinition is called when exiting the referenceDefinition production.
	ExitReferenceDefinition(c *ReferenceDefinitionContext)

	// ExitReferenceOption is called when exiting the referenceOption production.
	ExitReferenceOption(c *ReferenceOptionContext)

	// ExitKeyParts is called when exiting the keyParts production.
	ExitKeyParts(c *KeyPartsContext)

	// ExitKeyPart is called when exiting the keyPart production.
	ExitKeyPart(c *KeyPartContext)

	// ExitConstraintDefinition is called when exiting the constraintDefinition production.
	ExitConstraintDefinition(c *ConstraintDefinitionContext)

	// ExitPrimaryKeyOption is called when exiting the primaryKeyOption production.
	ExitPrimaryKeyOption(c *PrimaryKeyOptionContext)

	// ExitPrimaryKey is called when exiting the primaryKey production.
	ExitPrimaryKey(c *PrimaryKeyContext)

	// ExitUniqueOption is called when exiting the uniqueOption production.
	ExitUniqueOption(c *UniqueOptionContext)

	// ExitForeignKeyOption is called when exiting the foreignKeyOption production.
	ExitForeignKeyOption(c *ForeignKeyOptionContext)

	// ExitCreateLikeClause is called when exiting the createLikeClause production.
	ExitCreateLikeClause(c *CreateLikeClauseContext)

	// ExitAlterDefinitionClause is called when exiting the alterDefinitionClause production.
	ExitAlterDefinitionClause(c *AlterDefinitionClauseContext)

	// ExitAddColumnSpecification is called when exiting the addColumnSpecification production.
	ExitAddColumnSpecification(c *AddColumnSpecificationContext)

	// ExitModifyColumnSpecification is called when exiting the modifyColumnSpecification production.
	ExitModifyColumnSpecification(c *ModifyColumnSpecificationContext)

	// ExitDropColumnSpecification is called when exiting the dropColumnSpecification production.
	ExitDropColumnSpecification(c *DropColumnSpecificationContext)

	// ExitAddConstraintSpecification is called when exiting the addConstraintSpecification production.
	ExitAddConstraintSpecification(c *AddConstraintSpecificationContext)

	// ExitDropConstraintSpecification is called when exiting the dropConstraintSpecification production.
	ExitDropConstraintSpecification(c *DropConstraintSpecificationContext)

	// ExitSetTransaction is called when exiting the setTransaction production.
	ExitSetTransaction(c *SetTransactionContext)

	// ExitCommit is called when exiting the commit production.
	ExitCommit(c *CommitContext)

	// ExitRollback is called when exiting the rollback production.
	ExitRollback(c *RollbackContext)

	// ExitLevelOfIsolation is called when exiting the levelOfIsolation production.
	ExitLevelOfIsolation(c *LevelOfIsolationContext)

	// ExitGrant is called when exiting the grant production.
	ExitGrant(c *GrantContext)

	// ExitRevoke is called when exiting the revoke production.
	ExitRevoke(c *RevokeContext)

	// ExitPrivilegeClause is called when exiting the privilegeClause production.
	ExitPrivilegeClause(c *PrivilegeClauseContext)

	// ExitPrivileges is called when exiting the privileges production.
	ExitPrivileges(c *PrivilegesContext)

	// ExitPrivilegeType is called when exiting the privilegeType production.
	ExitPrivilegeType(c *PrivilegeTypeContext)

	// ExitGrantee is called when exiting the grantee production.
	ExitGrantee(c *GranteeContext)

	// ExitOnObjectClause is called when exiting the onObjectClause production.
	ExitOnObjectClause(c *OnObjectClauseContext)

	// ExitObjectType is called when exiting the objectType production.
	ExitObjectType(c *ObjectTypeContext)

	// ExitPrivilegeLevel is called when exiting the privilegeLevel production.
	ExitPrivilegeLevel(c *PrivilegeLevelContext)
}

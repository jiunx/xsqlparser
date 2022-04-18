// Code generated from E:/GoProject/src/github.com/jiunx/xsqlparser/dialect/sql92/grammer\SQL92Statement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQL92Statement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SQL92StatementParser.
type SQL92StatementVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SQL92StatementParser#execute.
	VisitExecute(ctx *ExecuteContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#insert.
	VisitInsert(ctx *InsertContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#insertValuesClause.
	VisitInsertValuesClause(ctx *InsertValuesClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#insertSelectClause.
	VisitInsertSelectClause(ctx *InsertSelectClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#update.
	VisitUpdate(ctx *UpdateContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#setAssignmentsClause.
	VisitSetAssignmentsClause(ctx *SetAssignmentsClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#assignmentValues.
	VisitAssignmentValues(ctx *AssignmentValuesContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#assignmentValue.
	VisitAssignmentValue(ctx *AssignmentValueContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#blobValue.
	VisitBlobValue(ctx *BlobValueContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#delete.
	VisitDelete(ctx *DeleteContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#singleTableClause.
	VisitSingleTableClause(ctx *SingleTableClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#select.
	VisitSelect(ctx *SelectContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#unionClause.
	VisitUnionClause(ctx *UnionClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#selectClause.
	VisitSelectClause(ctx *SelectClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#selectSpecification.
	VisitSelectSpecification(ctx *SelectSpecificationContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#duplicateSpecification.
	VisitDuplicateSpecification(ctx *DuplicateSpecificationContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#projections.
	VisitProjections(ctx *ProjectionsContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#projection.
	VisitProjection(ctx *ProjectionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#unqualifiedShorthand.
	VisitUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#qualifiedShorthand.
	VisitQualifiedShorthand(ctx *QualifiedShorthandContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#tableReferences.
	VisitTableReferences(ctx *TableReferencesContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#escapedTableReference.
	VisitEscapedTableReference(ctx *EscapedTableReferenceContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#tableReference.
	VisitTableReference(ctx *TableReferenceContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#tableFactor.
	VisitTableFactor(ctx *TableFactorContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#joinedTable.
	VisitJoinedTable(ctx *JoinedTableContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#joinSpecification.
	VisitJoinSpecification(ctx *JoinSpecificationContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#limitRowCount.
	VisitLimitRowCount(ctx *LimitRowCountContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#limitOffset.
	VisitLimitOffset(ctx *LimitOffsetContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#subquery.
	VisitSubquery(ctx *SubqueryContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#parameterMarker.
	VisitParameterMarker(ctx *ParameterMarkerContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#literals.
	VisitLiterals(ctx *LiteralsContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#stringLiterals.
	VisitStringLiterals(ctx *StringLiteralsContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#numberLiterals.
	VisitNumberLiterals(ctx *NumberLiteralsContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#dateTimeLiterals.
	VisitDateTimeLiterals(ctx *DateTimeLiteralsContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#hexadecimalLiterals.
	VisitHexadecimalLiterals(ctx *HexadecimalLiteralsContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#bitValueLiterals.
	VisitBitValueLiterals(ctx *BitValueLiteralsContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#booleanLiterals.
	VisitBooleanLiterals(ctx *BooleanLiteralsContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#nullValueLiterals.
	VisitNullValueLiterals(ctx *NullValueLiteralsContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#unreservedWord.
	VisitUnreservedWord(ctx *UnreservedWordContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#schemaName.
	VisitSchemaName(ctx *SchemaNameContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#tableName.
	VisitTableName(ctx *TableNameContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#columnName.
	VisitColumnName(ctx *ColumnNameContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#viewName.
	VisitViewName(ctx *ViewNameContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#owner.
	VisitOwner(ctx *OwnerContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#constraintName.
	VisitConstraintName(ctx *ConstraintNameContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#columnNames.
	VisitColumnNames(ctx *ColumnNamesContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#tableNames.
	VisitTableNames(ctx *TableNamesContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#characterSetName.
	VisitCharacterSetName(ctx *CharacterSetNameContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#andOperator.
	VisitAndOperator(ctx *AndOperatorContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#orOperator.
	VisitOrOperator(ctx *OrOperatorContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#notOperator.
	VisitNotOperator(ctx *NotOperatorContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#booleanPrimary.
	VisitBooleanPrimary(ctx *BooleanPrimaryContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#predicate.
	VisitPredicate(ctx *PredicateContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#bitExpr.
	VisitBitExpr(ctx *BitExprContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#simpleExpr.
	VisitSimpleExpr(ctx *SimpleExprContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#aggregationFunction.
	VisitAggregationFunction(ctx *AggregationFunctionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#aggregationFunctionName.
	VisitAggregationFunctionName(ctx *AggregationFunctionNameContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#distinct.
	VisitDistinct(ctx *DistinctContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#specialFunction.
	VisitSpecialFunction(ctx *SpecialFunctionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#castFunction.
	VisitCastFunction(ctx *CastFunctionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#convertFunction.
	VisitConvertFunction(ctx *ConvertFunctionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#positionFunction.
	VisitPositionFunction(ctx *PositionFunctionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#substringFunction.
	VisitSubstringFunction(ctx *SubstringFunctionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#extractFunction.
	VisitExtractFunction(ctx *ExtractFunctionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#trimFunction.
	VisitTrimFunction(ctx *TrimFunctionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#regularFunction.
	VisitRegularFunction(ctx *RegularFunctionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#regularFunctionName.
	VisitRegularFunctionName(ctx *RegularFunctionNameContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#matchExpression.
	VisitMatchExpression(ctx *MatchExpressionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#caseExpression.
	VisitCaseExpression(ctx *CaseExpressionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#caseWhen.
	VisitCaseWhen(ctx *CaseWhenContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#caseElse.
	VisitCaseElse(ctx *CaseElseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#intervalExpression.
	VisitIntervalExpression(ctx *IntervalExpressionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#intervalUnit.
	VisitIntervalUnit(ctx *IntervalUnitContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#orderByItem.
	VisitOrderByItem(ctx *OrderByItemContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#dataType.
	VisitDataType(ctx *DataTypeContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#dataTypeName.
	VisitDataTypeName(ctx *DataTypeNameContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#dataTypeLength.
	VisitDataTypeLength(ctx *DataTypeLengthContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#characterSet.
	VisitCharacterSet(ctx *CharacterSetContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#collateClause.
	VisitCollateClause(ctx *CollateClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#ignoredIdentifier.
	VisitIgnoredIdentifier(ctx *IgnoredIdentifierContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#dropBehaviour.
	VisitDropBehaviour(ctx *DropBehaviourContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#createTable.
	VisitCreateTable(ctx *CreateTableContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#alterTable.
	VisitAlterTable(ctx *AlterTableContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#dropTable.
	VisitDropTable(ctx *DropTableContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#createDatabase.
	VisitCreateDatabase(ctx *CreateDatabaseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#dropDatabase.
	VisitDropDatabase(ctx *DropDatabaseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#createView.
	VisitCreateView(ctx *CreateViewContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#dropView.
	VisitDropView(ctx *DropViewContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#createTableSpecification.
	VisitCreateTableSpecification(ctx *CreateTableSpecificationContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#createDefinitionClause.
	VisitCreateDefinitionClause(ctx *CreateDefinitionClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#createDatabaseSpecification_.
	VisitCreateDatabaseSpecification_(ctx *CreateDatabaseSpecification_Context) interface{}

	// Visit a parse tree produced by SQL92StatementParser#createDefinition.
	VisitCreateDefinition(ctx *CreateDefinitionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#columnDefinition.
	VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#dataTypeOption.
	VisitDataTypeOption(ctx *DataTypeOptionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#checkConstraintDefinition.
	VisitCheckConstraintDefinition(ctx *CheckConstraintDefinitionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#referenceDefinition.
	VisitReferenceDefinition(ctx *ReferenceDefinitionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#referenceOption.
	VisitReferenceOption(ctx *ReferenceOptionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#keyParts.
	VisitKeyParts(ctx *KeyPartsContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#keyPart.
	VisitKeyPart(ctx *KeyPartContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#constraintDefinition.
	VisitConstraintDefinition(ctx *ConstraintDefinitionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#primaryKeyOption.
	VisitPrimaryKeyOption(ctx *PrimaryKeyOptionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#primaryKey.
	VisitPrimaryKey(ctx *PrimaryKeyContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#uniqueOption.
	VisitUniqueOption(ctx *UniqueOptionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#foreignKeyOption.
	VisitForeignKeyOption(ctx *ForeignKeyOptionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#createLikeClause.
	VisitCreateLikeClause(ctx *CreateLikeClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#alterDefinitionClause.
	VisitAlterDefinitionClause(ctx *AlterDefinitionClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#addColumnSpecification.
	VisitAddColumnSpecification(ctx *AddColumnSpecificationContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#modifyColumnSpecification.
	VisitModifyColumnSpecification(ctx *ModifyColumnSpecificationContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#dropColumnSpecification.
	VisitDropColumnSpecification(ctx *DropColumnSpecificationContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#addConstraintSpecification.
	VisitAddConstraintSpecification(ctx *AddConstraintSpecificationContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#dropConstraintSpecification.
	VisitDropConstraintSpecification(ctx *DropConstraintSpecificationContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#setTransaction.
	VisitSetTransaction(ctx *SetTransactionContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#commit.
	VisitCommit(ctx *CommitContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#rollback.
	VisitRollback(ctx *RollbackContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#levelOfIsolation.
	VisitLevelOfIsolation(ctx *LevelOfIsolationContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#grant.
	VisitGrant(ctx *GrantContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#revoke.
	VisitRevoke(ctx *RevokeContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#privilegeClause.
	VisitPrivilegeClause(ctx *PrivilegeClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#privileges.
	VisitPrivileges(ctx *PrivilegesContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#privilegeType.
	VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#grantee.
	VisitGrantee(ctx *GranteeContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#onObjectClause.
	VisitOnObjectClause(ctx *OnObjectClauseContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#objectType.
	VisitObjectType(ctx *ObjectTypeContext) interface{}

	// Visit a parse tree produced by SQL92StatementParser#privilegeLevel.
	VisitPrivilegeLevel(ctx *PrivilegeLevelContext) interface{}
}

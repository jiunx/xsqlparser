// Code generated from E:/GoProject/src/github.com/jiunx/xsqlparser/dialect/sql92/grammer\SQL92Statement.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // SQL92Statement

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSQL92StatementVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSQL92StatementVisitor) VisitExecute(ctx *ExecuteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitInsert(ctx *InsertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitInsertValuesClause(ctx *InsertValuesClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitInsertSelectClause(ctx *InsertSelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitUpdate(ctx *UpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitSetAssignmentsClause(ctx *SetAssignmentsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitAssignmentValues(ctx *AssignmentValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitAssignmentValue(ctx *AssignmentValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitBlobValue(ctx *BlobValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDelete(ctx *DeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitSingleTableClause(ctx *SingleTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitSelect(ctx *SelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitUnionClause(ctx *UnionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitSelectClause(ctx *SelectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitSelectSpecification(ctx *SelectSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDuplicateSpecification(ctx *DuplicateSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitProjections(ctx *ProjectionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitProjection(ctx *ProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitUnqualifiedShorthand(ctx *UnqualifiedShorthandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitQualifiedShorthand(ctx *QualifiedShorthandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitTableReferences(ctx *TableReferencesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitEscapedTableReference(ctx *EscapedTableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitTableReference(ctx *TableReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitTableFactor(ctx *TableFactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitJoinedTable(ctx *JoinedTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitJoinSpecification(ctx *JoinSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitLimitRowCount(ctx *LimitRowCountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitLimitOffset(ctx *LimitOffsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitParameterMarker(ctx *ParameterMarkerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitLiterals(ctx *LiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitStringLiterals(ctx *StringLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitNumberLiterals(ctx *NumberLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDateTimeLiterals(ctx *DateTimeLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitHexadecimalLiterals(ctx *HexadecimalLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitBitValueLiterals(ctx *BitValueLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitBooleanLiterals(ctx *BooleanLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitNullValueLiterals(ctx *NullValueLiteralsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitUnreservedWord(ctx *UnreservedWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitSchemaName(ctx *SchemaNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitTableName(ctx *TableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitColumnName(ctx *ColumnNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitViewName(ctx *ViewNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitOwner(ctx *OwnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitConstraintName(ctx *ConstraintNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitColumnNames(ctx *ColumnNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitTableNames(ctx *TableNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCharacterSetName(ctx *CharacterSetNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitAndOperator(ctx *AndOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitOrOperator(ctx *OrOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitNotOperator(ctx *NotOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitBooleanPrimary(ctx *BooleanPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitPredicate(ctx *PredicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitBitExpr(ctx *BitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitSimpleExpr(ctx *SimpleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitAggregationFunction(ctx *AggregationFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitAggregationFunctionName(ctx *AggregationFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDistinct(ctx *DistinctContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitSpecialFunction(ctx *SpecialFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCastFunction(ctx *CastFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitConvertFunction(ctx *ConvertFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitPositionFunction(ctx *PositionFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitSubstringFunction(ctx *SubstringFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitExtractFunction(ctx *ExtractFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitTrimFunction(ctx *TrimFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitRegularFunction(ctx *RegularFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitRegularFunctionName(ctx *RegularFunctionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitMatchExpression(ctx *MatchExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCaseExpression(ctx *CaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCaseWhen(ctx *CaseWhenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCaseElse(ctx *CaseElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitIntervalExpression(ctx *IntervalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitIntervalUnit(ctx *IntervalUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitOrderByItem(ctx *OrderByItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDataType(ctx *DataTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDataTypeName(ctx *DataTypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDataTypeLength(ctx *DataTypeLengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCharacterSet(ctx *CharacterSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCollateClause(ctx *CollateClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitIgnoredIdentifier(ctx *IgnoredIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDropBehaviour(ctx *DropBehaviourContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCreateTable(ctx *CreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitAlterTable(ctx *AlterTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDropTable(ctx *DropTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCreateDatabase(ctx *CreateDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDropDatabase(ctx *DropDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCreateView(ctx *CreateViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDropView(ctx *DropViewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCreateTableSpecification(ctx *CreateTableSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCreateDefinitionClause(ctx *CreateDefinitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCreateDatabaseSpecification_(ctx *CreateDatabaseSpecification_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCreateDefinition(ctx *CreateDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitColumnDefinition(ctx *ColumnDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDataTypeOption(ctx *DataTypeOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCheckConstraintDefinition(ctx *CheckConstraintDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitReferenceDefinition(ctx *ReferenceDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitReferenceOption(ctx *ReferenceOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitKeyParts(ctx *KeyPartsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitKeyPart(ctx *KeyPartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitConstraintDefinition(ctx *ConstraintDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitPrimaryKeyOption(ctx *PrimaryKeyOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitPrimaryKey(ctx *PrimaryKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitUniqueOption(ctx *UniqueOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitForeignKeyOption(ctx *ForeignKeyOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCreateLikeClause(ctx *CreateLikeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitAlterDefinitionClause(ctx *AlterDefinitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitAddColumnSpecification(ctx *AddColumnSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitModifyColumnSpecification(ctx *ModifyColumnSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDropColumnSpecification(ctx *DropColumnSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitAddConstraintSpecification(ctx *AddConstraintSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitDropConstraintSpecification(ctx *DropConstraintSpecificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitSetTransaction(ctx *SetTransactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitCommit(ctx *CommitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitRollback(ctx *RollbackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitLevelOfIsolation(ctx *LevelOfIsolationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitGrant(ctx *GrantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitRevoke(ctx *RevokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitPrivilegeClause(ctx *PrivilegeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitPrivileges(ctx *PrivilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitPrivilegeType(ctx *PrivilegeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitGrantee(ctx *GranteeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitOnObjectClause(ctx *OnObjectClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitObjectType(ctx *ObjectTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSQL92StatementVisitor) VisitPrivilegeLevel(ctx *PrivilegeLevelContext) interface{} {
	return v.VisitChildren(ctx)
}

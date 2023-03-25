// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GoliteParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// GoliteParserListener is a complete listener for a parse tree produced by GoliteParser.
type GoliteParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterTypedeclaration is called when entering the typedeclaration production.
	EnterTypedeclaration(c *TypedeclarationContext)

	// EnterFields is called when entering the fields production.
	EnterFields(c *FieldsContext)

	// EnterFieldsprime is called when entering the fieldsprime production.
	EnterFieldsprime(c *FieldsprimeContext)

	// EnterDecl is called when entering the decl production.
	EnterDecl(c *DeclContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterDeclarationprime is called when entering the declarationprime production.
	EnterDeclarationprime(c *DeclarationprimeContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFuncprime is called when entering the funcprime production.
	EnterFuncprime(c *FuncprimeContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterParametersprime is called when entering the parametersprime production.
	EnterParametersprime(c *ParametersprimeContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterDelete is called when entering the delete production.
	EnterDelete(c *DeleteContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterAssignmentprime is called when entering the assignmentprime production.
	EnterAssignmentprime(c *AssignmentprimeContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterConditional is called when entering the conditional production.
	EnterConditional(c *ConditionalContext)

	// EnterConditionalprime is called when entering the conditionalprime production.
	EnterConditionalprime(c *ConditionalprimeContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterRtrn is called when entering the rtrn production.
	EnterRtrn(c *RtrnContext)

	// EnterRtrnprime is called when entering the rtrnprime production.
	EnterRtrnprime(c *RtrnprimeContext)

	// EnterRead is called when entering the read production.
	EnterRead(c *ReadContext)

	// EnterInvocation is called when entering the invocation production.
	EnterInvocation(c *InvocationContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgumentsprime is called when entering the argumentsprime production.
	EnterArgumentsprime(c *ArgumentsprimeContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpressionprime is called when entering the expressionprime production.
	EnterExpressionprime(c *ExpressionprimeContext)

	// EnterBoolterm is called when entering the boolterm production.
	EnterBoolterm(c *BooltermContext)

	// EnterBoolprime is called when entering the boolprime production.
	EnterBoolprime(c *BoolprimeContext)

	// EnterEqualterm is called when entering the equalterm production.
	EnterEqualterm(c *EqualtermContext)

	// EnterEqualprime is called when entering the equalprime production.
	EnterEqualprime(c *EqualprimeContext)

	// EnterRelationterm is called when entering the relationterm production.
	EnterRelationterm(c *RelationtermContext)

	// EnterRelationprime is called when entering the relationprime production.
	EnterRelationprime(c *RelationprimeContext)

	// EnterSimpleterm is called when entering the simpleterm production.
	EnterSimpleterm(c *SimpletermContext)

	// EnterSimpleprime is called when entering the simpleprime production.
	EnterSimpleprime(c *SimpleprimeContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterTermprime is called when entering the termprime production.
	EnterTermprime(c *TermprimeContext)

	// EnterUnaryterm is called when entering the unaryterm production.
	EnterUnaryterm(c *UnarytermContext)

	// EnterSelectorterm is called when entering the selectorterm production.
	EnterSelectorterm(c *SelectortermContext)

	// EnterSelectorprime is called when entering the selectorprime production.
	EnterSelectorprime(c *SelectorprimeContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitTypedeclaration is called when exiting the typedeclaration production.
	ExitTypedeclaration(c *TypedeclarationContext)

	// ExitFields is called when exiting the fields production.
	ExitFields(c *FieldsContext)

	// ExitFieldsprime is called when exiting the fieldsprime production.
	ExitFieldsprime(c *FieldsprimeContext)

	// ExitDecl is called when exiting the decl production.
	ExitDecl(c *DeclContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitDeclarationprime is called when exiting the declarationprime production.
	ExitDeclarationprime(c *DeclarationprimeContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFuncprime is called when exiting the funcprime production.
	ExitFuncprime(c *FuncprimeContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitParametersprime is called when exiting the parametersprime production.
	ExitParametersprime(c *ParametersprimeContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitDelete is called when exiting the delete production.
	ExitDelete(c *DeleteContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitAssignmentprime is called when exiting the assignmentprime production.
	ExitAssignmentprime(c *AssignmentprimeContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitConditional is called when exiting the conditional production.
	ExitConditional(c *ConditionalContext)

	// ExitConditionalprime is called when exiting the conditionalprime production.
	ExitConditionalprime(c *ConditionalprimeContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitRtrn is called when exiting the rtrn production.
	ExitRtrn(c *RtrnContext)

	// ExitRtrnprime is called when exiting the rtrnprime production.
	ExitRtrnprime(c *RtrnprimeContext)

	// ExitRead is called when exiting the read production.
	ExitRead(c *ReadContext)

	// ExitInvocation is called when exiting the invocation production.
	ExitInvocation(c *InvocationContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgumentsprime is called when exiting the argumentsprime production.
	ExitArgumentsprime(c *ArgumentsprimeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpressionprime is called when exiting the expressionprime production.
	ExitExpressionprime(c *ExpressionprimeContext)

	// ExitBoolterm is called when exiting the boolterm production.
	ExitBoolterm(c *BooltermContext)

	// ExitBoolprime is called when exiting the boolprime production.
	ExitBoolprime(c *BoolprimeContext)

	// ExitEqualterm is called when exiting the equalterm production.
	ExitEqualterm(c *EqualtermContext)

	// ExitEqualprime is called when exiting the equalprime production.
	ExitEqualprime(c *EqualprimeContext)

	// ExitRelationterm is called when exiting the relationterm production.
	ExitRelationterm(c *RelationtermContext)

	// ExitRelationprime is called when exiting the relationprime production.
	ExitRelationprime(c *RelationprimeContext)

	// ExitSimpleterm is called when exiting the simpleterm production.
	ExitSimpleterm(c *SimpletermContext)

	// ExitSimpleprime is called when exiting the simpleprime production.
	ExitSimpleprime(c *SimpleprimeContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitTermprime is called when exiting the termprime production.
	ExitTermprime(c *TermprimeContext)

	// ExitUnaryterm is called when exiting the unaryterm production.
	ExitUnaryterm(c *UnarytermContext)

	// ExitSelectorterm is called when exiting the selectorterm production.
	ExitSelectorterm(c *SelectortermContext)

	// ExitSelectorprime is called when exiting the selectorprime production.
	ExitSelectorprime(c *SelectorprimeContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)
}

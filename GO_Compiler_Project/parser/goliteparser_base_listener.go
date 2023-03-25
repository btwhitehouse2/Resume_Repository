// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GoliteParser
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseGoliteParserListener is a complete listener for a parse tree produced by GoliteParser.
type BaseGoliteParserListener struct{}

var _ GoliteParserListener = &BaseGoliteParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoliteParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoliteParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoliteParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoliteParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGoliteParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGoliteParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterTypedeclaration is called when production typedeclaration is entered.
func (s *BaseGoliteParserListener) EnterTypedeclaration(ctx *TypedeclarationContext) {}

// ExitTypedeclaration is called when production typedeclaration is exited.
func (s *BaseGoliteParserListener) ExitTypedeclaration(ctx *TypedeclarationContext) {}

// EnterFields is called when production fields is entered.
func (s *BaseGoliteParserListener) EnterFields(ctx *FieldsContext) {}

// ExitFields is called when production fields is exited.
func (s *BaseGoliteParserListener) ExitFields(ctx *FieldsContext) {}

// EnterFieldsprime is called when production fieldsprime is entered.
func (s *BaseGoliteParserListener) EnterFieldsprime(ctx *FieldsprimeContext) {}

// ExitFieldsprime is called when production fieldsprime is exited.
func (s *BaseGoliteParserListener) ExitFieldsprime(ctx *FieldsprimeContext) {}

// EnterDecl is called when production decl is entered.
func (s *BaseGoliteParserListener) EnterDecl(ctx *DeclContext) {}

// ExitDecl is called when production decl is exited.
func (s *BaseGoliteParserListener) ExitDecl(ctx *DeclContext) {}

// EnterType is called when production type is entered.
func (s *BaseGoliteParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseGoliteParserListener) ExitType(ctx *TypeContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseGoliteParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseGoliteParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterDeclarationprime is called when production declarationprime is entered.
func (s *BaseGoliteParserListener) EnterDeclarationprime(ctx *DeclarationprimeContext) {}

// ExitDeclarationprime is called when production declarationprime is exited.
func (s *BaseGoliteParserListener) ExitDeclarationprime(ctx *DeclarationprimeContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseGoliteParserListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseGoliteParserListener) ExitFunction(ctx *FunctionContext) {}

// EnterFuncprime is called when production funcprime is entered.
func (s *BaseGoliteParserListener) EnterFuncprime(ctx *FuncprimeContext) {}

// ExitFuncprime is called when production funcprime is exited.
func (s *BaseGoliteParserListener) ExitFuncprime(ctx *FuncprimeContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseGoliteParserListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseGoliteParserListener) ExitParameters(ctx *ParametersContext) {}

// EnterParametersprime is called when production parametersprime is entered.
func (s *BaseGoliteParserListener) EnterParametersprime(ctx *ParametersprimeContext) {}

// ExitParametersprime is called when production parametersprime is exited.
func (s *BaseGoliteParserListener) ExitParametersprime(ctx *ParametersprimeContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseGoliteParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseGoliteParserListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGoliteParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGoliteParserListener) ExitBlock(ctx *BlockContext) {}

// EnterDelete is called when production delete is entered.
func (s *BaseGoliteParserListener) EnterDelete(ctx *DeleteContext) {}

// ExitDelete is called when production delete is exited.
func (s *BaseGoliteParserListener) ExitDelete(ctx *DeleteContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseGoliteParserListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseGoliteParserListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterAssignmentprime is called when production assignmentprime is entered.
func (s *BaseGoliteParserListener) EnterAssignmentprime(ctx *AssignmentprimeContext) {}

// ExitAssignmentprime is called when production assignmentprime is exited.
func (s *BaseGoliteParserListener) ExitAssignmentprime(ctx *AssignmentprimeContext) {}

// EnterPrint is called when production print is entered.
func (s *BaseGoliteParserListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BaseGoliteParserListener) ExitPrint(ctx *PrintContext) {}

// EnterConditional is called when production conditional is entered.
func (s *BaseGoliteParserListener) EnterConditional(ctx *ConditionalContext) {}

// ExitConditional is called when production conditional is exited.
func (s *BaseGoliteParserListener) ExitConditional(ctx *ConditionalContext) {}

// EnterConditionalprime is called when production conditionalprime is entered.
func (s *BaseGoliteParserListener) EnterConditionalprime(ctx *ConditionalprimeContext) {}

// ExitConditionalprime is called when production conditionalprime is exited.
func (s *BaseGoliteParserListener) ExitConditionalprime(ctx *ConditionalprimeContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseGoliteParserListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseGoliteParserListener) ExitLoop(ctx *LoopContext) {}

// EnterRtrn is called when production rtrn is entered.
func (s *BaseGoliteParserListener) EnterRtrn(ctx *RtrnContext) {}

// ExitRtrn is called when production rtrn is exited.
func (s *BaseGoliteParserListener) ExitRtrn(ctx *RtrnContext) {}

// EnterRtrnprime is called when production rtrnprime is entered.
func (s *BaseGoliteParserListener) EnterRtrnprime(ctx *RtrnprimeContext) {}

// ExitRtrnprime is called when production rtrnprime is exited.
func (s *BaseGoliteParserListener) ExitRtrnprime(ctx *RtrnprimeContext) {}

// EnterRead is called when production read is entered.
func (s *BaseGoliteParserListener) EnterRead(ctx *ReadContext) {}

// ExitRead is called when production read is exited.
func (s *BaseGoliteParserListener) ExitRead(ctx *ReadContext) {}

// EnterInvocation is called when production invocation is entered.
func (s *BaseGoliteParserListener) EnterInvocation(ctx *InvocationContext) {}

// ExitInvocation is called when production invocation is exited.
func (s *BaseGoliteParserListener) ExitInvocation(ctx *InvocationContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGoliteParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGoliteParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgumentsprime is called when production argumentsprime is entered.
func (s *BaseGoliteParserListener) EnterArgumentsprime(ctx *ArgumentsprimeContext) {}

// ExitArgumentsprime is called when production argumentsprime is exited.
func (s *BaseGoliteParserListener) ExitArgumentsprime(ctx *ArgumentsprimeContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGoliteParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGoliteParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpressionprime is called when production expressionprime is entered.
func (s *BaseGoliteParserListener) EnterExpressionprime(ctx *ExpressionprimeContext) {}

// ExitExpressionprime is called when production expressionprime is exited.
func (s *BaseGoliteParserListener) ExitExpressionprime(ctx *ExpressionprimeContext) {}

// EnterBoolterm is called when production boolterm is entered.
func (s *BaseGoliteParserListener) EnterBoolterm(ctx *BooltermContext) {}

// ExitBoolterm is called when production boolterm is exited.
func (s *BaseGoliteParserListener) ExitBoolterm(ctx *BooltermContext) {}

// EnterBoolprime is called when production boolprime is entered.
func (s *BaseGoliteParserListener) EnterBoolprime(ctx *BoolprimeContext) {}

// ExitBoolprime is called when production boolprime is exited.
func (s *BaseGoliteParserListener) ExitBoolprime(ctx *BoolprimeContext) {}

// EnterEqualterm is called when production equalterm is entered.
func (s *BaseGoliteParserListener) EnterEqualterm(ctx *EqualtermContext) {}

// ExitEqualterm is called when production equalterm is exited.
func (s *BaseGoliteParserListener) ExitEqualterm(ctx *EqualtermContext) {}

// EnterEqualprime is called when production equalprime is entered.
func (s *BaseGoliteParserListener) EnterEqualprime(ctx *EqualprimeContext) {}

// ExitEqualprime is called when production equalprime is exited.
func (s *BaseGoliteParserListener) ExitEqualprime(ctx *EqualprimeContext) {}

// EnterRelationterm is called when production relationterm is entered.
func (s *BaseGoliteParserListener) EnterRelationterm(ctx *RelationtermContext) {}

// ExitRelationterm is called when production relationterm is exited.
func (s *BaseGoliteParserListener) ExitRelationterm(ctx *RelationtermContext) {}

// EnterRelationprime is called when production relationprime is entered.
func (s *BaseGoliteParserListener) EnterRelationprime(ctx *RelationprimeContext) {}

// ExitRelationprime is called when production relationprime is exited.
func (s *BaseGoliteParserListener) ExitRelationprime(ctx *RelationprimeContext) {}

// EnterSimpleterm is called when production simpleterm is entered.
func (s *BaseGoliteParserListener) EnterSimpleterm(ctx *SimpletermContext) {}

// ExitSimpleterm is called when production simpleterm is exited.
func (s *BaseGoliteParserListener) ExitSimpleterm(ctx *SimpletermContext) {}

// EnterSimpleprime is called when production simpleprime is entered.
func (s *BaseGoliteParserListener) EnterSimpleprime(ctx *SimpleprimeContext) {}

// ExitSimpleprime is called when production simpleprime is exited.
func (s *BaseGoliteParserListener) ExitSimpleprime(ctx *SimpleprimeContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseGoliteParserListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseGoliteParserListener) ExitTerm(ctx *TermContext) {}

// EnterTermprime is called when production termprime is entered.
func (s *BaseGoliteParserListener) EnterTermprime(ctx *TermprimeContext) {}

// ExitTermprime is called when production termprime is exited.
func (s *BaseGoliteParserListener) ExitTermprime(ctx *TermprimeContext) {}

// EnterUnaryterm is called when production unaryterm is entered.
func (s *BaseGoliteParserListener) EnterUnaryterm(ctx *UnarytermContext) {}

// ExitUnaryterm is called when production unaryterm is exited.
func (s *BaseGoliteParserListener) ExitUnaryterm(ctx *UnarytermContext) {}

// EnterSelectorterm is called when production selectorterm is entered.
func (s *BaseGoliteParserListener) EnterSelectorterm(ctx *SelectortermContext) {}

// ExitSelectorterm is called when production selectorterm is exited.
func (s *BaseGoliteParserListener) ExitSelectorterm(ctx *SelectortermContext) {}

// EnterSelectorprime is called when production selectorprime is entered.
func (s *BaseGoliteParserListener) EnterSelectorprime(ctx *SelectorprimeContext) {}

// ExitSelectorprime is called when production selectorprime is exited.
func (s *BaseGoliteParserListener) ExitSelectorprime(ctx *SelectorprimeContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseGoliteParserListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseGoliteParserListener) ExitFactor(ctx *FactorContext) {}

package parser

import (
	"fmt"
	"golite/ast"
	"golite/context"
	"golite/lexer"
	"golite/token"
	"golite/types"
	ty "golite/types"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Parser interface {
	Parse() *ast.Program
	GetErrors() []*context.CompilerError
}

type goliteParser struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	*BaseGoliteParserListener
	parser *GoliteParser
	lexer  lexer.Lexer
	errors []*context.CompilerError
	nodes  map[string]interface{}
}

func (gParser *goliteParser) GetErrors() []*context.CompilerError {
	return gParser.errors
}

func (gParser *goliteParser) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	gParser.errors = append(gParser.errors, &context.CompilerError{
		Line:   line,
		Column: column,
		Msg:    msg,
		Phase:  context.PARSER,
	})
}

func NewParser(lexer lexer.Lexer) Parser {
	wrappedParser := &goliteParser{antlr.NewDefaultErrorListener(), &BaseGoliteParserListener{}, nil, lexer, nil, make(map[string]interface{})}
	p := NewGoliteParser(lexer.GetTokenStream())
	p.RemoveErrorListeners()
	p.AddErrorListener(wrappedParser)
	wrappedParser.parser = p
	return wrappedParser
}
func (gParser *goliteParser) Parse() *ast.Program {
	ctx := gParser.parser.Program()
	if context.HasErrors(gParser.lexer.GetErrors()) ||
		context.HasErrors(gParser.GetErrors()) {
		fmt.Println("has errors")
		return nil
	}
	antlr.ParseTreeWalkerDefault.Walk(gParser, ctx)
	_, _, key := GetTokenInfo(ctx)
	return gParser.nodes[key].(*ast.Program)
}

/******************* Implementation of the Listeners **************************/

// GetTokenInfo gerates a unique key for each node in the ParserTree
func GetTokenInfo(ctx antlr.ParserRuleContext) (int, int, string) {
	key := fmt.Sprintf("%d,%d", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	return ctx.GetStart().GetLine(), ctx.GetStart().GetColumn(), key
}


func (gParser *goliteParser) ExitProgram(c *ProgramContext) {
	//fmt.Printf("-----Got to Program:%s-----\n",c.GetText())
	line, col, key := GetTokenInfo(c)
	var tdecls []*ast.TypeDecl
	var decls []*ast.Declaration
	var funcs []*ast.Function

	for _, tdeclCtx := range c.AllTypedeclaration() {
		_, _, tdkey := GetTokenInfo(tdeclCtx)
		tdeclStmt := gParser.nodes[tdkey].(*ast.TypeDecl)
		tdecls = append(tdecls, tdeclStmt)
	} 
	for _, declCtx := range c.AllDeclaration() {
		_, _, dckey := GetTokenInfo(declCtx)
		declStmt := gParser.nodes[dckey].(*ast.Declaration)
		decls = append(decls, declStmt)
	}
	for _, fCtx := range c.AllFunction() {
		_, _, fnkey := GetTokenInfo(fCtx)
		fStmt := gParser.nodes[fnkey].(*ast.Function)
		funcs = append(funcs, fStmt)
	}
	gParser.nodes[key] = ast.NewProgram(tdecls, decls, funcs, token.NewToken(line, col))
}


func (gParser *goliteParser) ExitTypedeclaration(c *TypedeclarationContext) {
	//fmt.Printf("-----Got to TypeDecl:%s-----\n",c.GetText())
	line, col, key := GetTokenInfo(c)
	fldsFactor := c.GetFlds()
	_, _, fldKey := GetTokenInfo(fldsFactor)
	fld := gParser.nodes[fldKey].(*ast.Fields)
	tdecl := ast.NewTypeDecl(c.ID().GetText(),fld, token.NewToken(line, col))
	gParser.nodes[key] = tdecl
}

func (gParser *goliteParser) ExitFields(c *FieldsContext){
	line, col, key := GetTokenInfo(c)
	termContexts := c.AllFieldsprime()
	declFactor := c.GetDec()
	var declList []*ast.Decl 
	_, _, declKey := GetTokenInfo(declFactor)
	decl := gParser.nodes[declKey].(*ast.Decl)
	declList = append(declList, decl)
	gParser.nodes[key] = decl
	for _, dcl := range(termContexts) {
		_, _, declKey := GetTokenInfo(dcl)
		decl = gParser.nodes[declKey].(*ast.Decl)
		declList = append(declList, decl)
	}
	gParser.nodes[key] = ast.NewFields(declList, token.NewToken(line,col))
}


//'id' Type 
//'id' ('int' | 'bool' | '*''id')
func (gParser *goliteParser) ExitDecl(c *DeclContext) {
	line, col, key := GetTokenInfo(c)
	//id
	idFactor := c.GetId()
	_, _, typeKey := GetTokenInfo(c.GetTy())
		actualTy := gParser.nodes[typeKey].(ty.Type)
		gParser.nodes[key] = ast.NewDeclStm(idFactor.GetText(), actualTy, token.NewToken(line, col))
}


func (gParser *goliteParser) ExitType(c *TypeContext) {
	_, _, key := GetTokenInfo(c)
	// int | bool | '*' 'id'
	if intFactor := c.INT() ; intFactor != nil {
		intTy := ty.StringToType("int")
		gParser.nodes[key] = intTy
	}else if boolFactor := c.BOOL() ; boolFactor != nil {
		boolTy := ty.StringToType("bool")
		gParser.nodes[key] = boolTy
	}else if astFactor := c.ASTERISK() ; astFactor != nil {
		idFactor := c.ID() 
		gParser.nodes[key] = &ty.StructTy{idFactor.GetText()}
		}	
}

//var id, id, id type
//need to fix
func (gParser *goliteParser) ExitDeclaration(c *DeclarationContext) {
	//fmt.Printf("-----Got to Declaration:%s-----\n",c.GetText())
	line, col, key := GetTokenInfo(c)
	//id 
	termContexts := c.AllDeclarationprime()
	var idList []*ast.Variable
	id := ast.NewVariable(c.ID().GetText(), token.NewToken(line, col))
	idList = append(idList, id)
	_, _, typeKey := GetTokenInfo(c.Type_())
		actualTy := gParser.nodes[typeKey].(ty.Type)
	for _, termContext := range termContexts {
		id := ast.NewVariable(termContext.GetId().GetText(), token.NewToken(line, col))
		idList = append(idList, id)
	}
		gParser.nodes[key] = ast.NewDeclaration(idList, actualTy, token.NewToken(line, col))	
}

func (gParser *goliteParser) ExitFunction(c *FunctionContext) {
	//fmt.Printf("-----Got to ExitFunction:%s-----\n",c.GetText())
	line, col, key := GetTokenInfo(c)
	var returnType ty.Type
	var listDec []*ast.Declaration
	var stmts []ast.Statement  
	//func -don't care about func term
	//id
	idFactor := c.ID()
	gParser.nodes[key] = ast.NewVariable(idFactor.GetText(), token.NewToken(line, col))
	//parameters
	paramFactor := c.Parameters()
	_, _, parKey := GetTokenInfo(paramFactor)
	par := gParser.nodes[parKey].(*ast.Parameters)
	//[type]
	if typeFac := c.Funcprime(); typeFac != nil{	
		_, _, tKey := GetTokenInfo(typeFac)
		returnType = gParser.nodes[tKey].(types.Type)
	}else{
		returnType = types.StringToType("void")
	}
	//declarations
	declFactor := c.AllDeclaration()
	for _, decl := range declFactor {
		_, _, declKey := GetTokenInfo(decl)
		dec := gParser.nodes[declKey].(*ast.Declaration)
		listDec = append(listDec, dec)
	}
	//statements
	stmntFactor := c.AllStatement()
	for _, stmt := range stmntFactor {
	_, _, stKey := GetTokenInfo(stmt)
	st := gParser.nodes[stKey].(ast.Statement)
	stmts = append(stmts, st)
	}
	gParser.nodes[key] = ast.NewFunction(idFactor.GetText(),par,returnType,listDec,stmts,token.NewToken(line,col))

}

func (gParser *goliteParser) ExitParameters(c *ParametersContext) {
// '(' [DECL (',' DECL)* ]   ')'
//prime: (, DECL)* 
//double prime: [DECL (',' DECL)* ] 
	line, col, key := GetTokenInfo(c)
	termContexts := c.AllParametersprime()
	var declList []*ast.Decl 
	if declFactor := c.Decl(); declFactor != nil {
	_, _, declKey := GetTokenInfo(declFactor)
	decl := gParser.nodes[declKey].(*ast.Decl)
	declList = append(declList, decl)
	gParser.nodes[key] = decl
		for _, dcl := range(termContexts) {
			_, _, declKey := GetTokenInfo(dcl)
			decl = gParser.nodes[declKey].(*ast.Decl)
			declList = append(declList, decl)
		}
	}
	gParser.nodes[key] = ast.NewParameters(declList, token.NewToken(line,col))
}


func (gParser *goliteParser) ExitStatement(c *StatementContext) {
	//needs to include block or { }
	_, _, key := GetTokenInfo(c)
	//block
	//fmt.Printf("-----Got to Statement:%s-----\n",c.GetText())
	if blockFactor := c.Block(); blockFactor != nil { 
		_, _, blockKey := GetTokenInfo(blockFactor)
		//expr = gParser.nodes[blockKey].(ast.Expression)
		blk := gParser.nodes[blockKey].(*ast.Block)
		gParser.nodes[key] = blk
	//assignment: first part is ID
	//need to confirm that it is not invocation
	} else if assignmentFactor := c.Assignment(); assignmentFactor != nil {
		_, _, assKey := GetTokenInfo(assignmentFactor)
		expr := gParser.nodes[assKey].(*ast.Assignment)
		gParser.nodes[key] = expr
	//print: first part is printf
	}else if printFactor := c.Print_(); printFactor != nil {
		_, _, printKey := GetTokenInfo(printFactor)
		//expr = gParser.nodes[printKey].(ast.Expression)
		pr := gParser.nodes[printKey].(*ast.Print)
		gParser.nodes[key] = pr
	//delete: first part is delete
	}else if delFactor := c.Delete_(); delFactor != nil {
		_, _, delKey := GetTokenInfo(delFactor)
		//expr = gParser.nodes[delKey].(ast.Expression)
		del := gParser.nodes[delKey].(*ast.DelLiteral)
		gParser.nodes[key] = del
	//conditional: first part is if
	} else if condFactor := c.Conditional(); condFactor != nil {
		_, _, condKey := GetTokenInfo(condFactor)
		//expr = gParser.nodes[condKey].(ast.Expression)
		cond := gParser.nodes[condKey].(*ast.CondLit)
		gParser.nodes[key] = cond
	//loop : first part is for
	} else if loopFactor := c.Loop(); loopFactor != nil {
		_, _, loopKey := GetTokenInfo(loopFactor)
		//expr = gParser.nodes[condKey].(ast.Expression)
		f := gParser.nodes[loopKey].(*ast.ForLiteral)
		gParser.nodes[key] = f
	//return: first part is return
	} else if returnFactor := c.Rtrn(); returnFactor != nil {
		_, _, returnKey := GetTokenInfo(returnFactor)
		//expr = gParser.nodes[condKey].(ast.Expression)
		ret := gParser.nodes[returnKey].(*ast.ReturnLit)
		fmt.Println(returnKey,"return", ret)
		gParser.nodes[key] = ret
	//read: first part is scan
	} else if readFactor := c.Read(); readFactor != nil {
		_, _, readKey := GetTokenInfo(readFactor)
		//expr = gParser.nodes[readKey].(ast.Expression)
		read := gParser.nodes[readKey].(*ast.ReadLiteral)
		gParser.nodes[key] = read
	//invocation
	} else if invocationFactor := c.Invocation(); invocationFactor != nil{
		_, _, invKey := GetTokenInfo(invocationFactor)
		//expr = gParser.nodes[invKey].(ast.Expression)
		invoc := gParser.nodes[invKey].(*ast.InvocacationExpression)
		gParser.nodes[key] = invoc	
}
}

func (gParser *goliteParser) ExitBlock(c *BlockContext){
	line, col, key := GetTokenInfo(c)
	termContexts := c.AllStatement()
	var listStat	[]ast.Statement
	for _,termContext := range termContexts {
		_, _, blockKey := GetTokenInfo(termContext)
		stat := gParser.nodes[blockKey].(ast.Statement)
		listStat = append(listStat, stat)
	}
	gParser.nodes[key] = ast.NewBlockStm(listStat, token.NewToken(line,col))
}

func (gParser *goliteParser) ExitDelete(c *DeleteContext) {
	line, col, key := GetTokenInfo(c)
	_, _, exprKey := GetTokenInfo(c.Expression())
	//expression
	expr := gParser.nodes[exprKey].(ast.Expression)
	gParser.nodes[key] = ast.NewDelLit(expr, token.NewToken(line,col))

}

// EnterAssignment is called when entering the assignment production.
// Will include LValue
func (gParser *goliteParser) ExitAssignment(c *AssignmentContext) {
	line, col, key := GetTokenInfo(c)
	termContexts := c.AllAssignmentprime()
	_, _, exprKey := GetTokenInfo(c.Expression())
	var assignVar string
	//id
	expr := gParser.nodes[exprKey].(ast.Expression)
	if len(termContexts) == 0 {
		assignVar = c.ID().GetText()
	//(.id)*
	} else {
		assignVar = c.ID().GetText()
		for _, termContext := range(termContexts){
			assign := termContext.GetText()
			assignVar = assignVar + "." + assign
		}
	}
	//expression
	gParser.nodes[key] = ast.NewAssignStm(assignVar, expr, token.NewToken(line, col))
	}


// update: picks at least a string and maybe and expression
// EnterPrint is called when entering the print production.
func (gParser *goliteParser) ExitPrint(c *PrintContext) {
	line, col, key := GetTokenInfo(c)
	termContexts := c.AllExpression()
	stringFactor := c.STRING()
	//printf
	var listExpr []ast.Expression
	//expression
	for _,termContext := range(termContexts) {
		_, _, exprKey := GetTokenInfo(termContext)
		expr := gParser.nodes[exprKey].(ast.Expression)
		gParser.nodes[key] = expr
		listExpr = append(listExpr, expr)
	}
	gParser.nodes[key] = ast.NewPrintStm(listExpr, stringFactor.GetText(), token.NewToken(line,col))
}



func (gParser *goliteParser) ExitConditional(c *ConditionalContext) {
	line, col, key := GetTokenInfo(c)
	_, _, exprKey := GetTokenInfo(c.Expression())
	_, _, ifblockKey := GetTokenInfo(c.Block())
	expr := gParser.nodes[exprKey].(ast.Expression)
	//block
	ifblock := gParser.nodes[ifblockKey].(*ast.Block)
	//[else block]
	if termContexts := c.Conditionalprime(); termContexts != nil {
		_, _, elseblockKey := GetTokenInfo(c.Block())
		elseblock := gParser.nodes[elseblockKey].(*ast.Block)
		gParser.nodes[key] = ast.NewCondLit(expr,ifblock,elseblock,token.NewToken(line,col))
	}else {//no else statemetn so will insert a nil value
		var elseblock	*ast.Block //declared an empty block to pass		
		gParser.nodes[key] = ast.NewCondLit(expr,ifblock,elseblock,token.NewToken(line,col))
	}
}

//for loop
func (gParser *goliteParser) ExitLoop(c *LoopContext) {
	line, col, key := GetTokenInfo(c)
	_, _, exprKey := GetTokenInfo(c.Expression())
	_, _, blockKey := GetTokenInfo(c.Block())
	//for
	
	//expression
	expr := gParser.nodes[exprKey].(ast.Expression)
	//block = []Statments
	block := gParser.nodes[blockKey].(*ast.Block)
	gParser.nodes[key] = ast.NewForLit(expr,block, token.NewToken(line,col))
}

func (gParser *goliteParser) ExitRtrn(c *RtrnContext) {
	line, col, key := GetTokenInfo(c)
	//if a return value exists
	if returnFac := c.Rtrnprime(); returnFac != nil {
		_,_, exprKey := GetTokenInfo(returnFac)
		expr := gParser.nodes[exprKey].(ast.Expression)
		gParser.nodes[key] = ast.NewReturnLit(expr, token.NewToken(line,col))
	}else { // if a return value does not exist
		var expr ast.Expression
		gParser.nodes[key] = ast.NewReturnLit(expr, token.NewToken(line,col))
	}
}

func (gParser *goliteParser) ExitRead(c *ReadContext) {
	line, col, key := GetTokenInfo(c)
	_, _, exprKey := GetTokenInfo(c.Expression())
	expr := gParser.nodes[exprKey].(ast.Expression)
	//read
	gParser.nodes[key] = ast.NewReadLit(expr, token.NewToken(line, col))
}

func (gParser *goliteParser) ExitInvocation(c *InvocationContext) {
	line, col, key := GetTokenInfo(c)
	_,_, argKey := GetTokenInfo(c.Arguments())
	args := gParser.nodes[argKey].(*ast.Arguments)
	//ID + Expression
	gParser.nodes[key] = ast.NewInvocExpr(key, args, token.NewToken(line, col))
	}

//do we need commas
func (gParser *goliteParser) ExitArguments(c *ArgumentsContext) {
	line, col, key := GetTokenInfo(c)
	termContexts := c.AllArgumentsprime()
	_, _, eprKey := GetTokenInfo(c.Expression())
	//only 1 expression then
	if len(termContexts) == 0{
		//check to make sure not empty
		if exprCtx := c.Expression(); exprCtx != nil {
			expr := gParser.nodes[eprKey].(ast.Expression)
			var expressions []ast.Expression
			expressions = append(expressions, expr)
			gParser.nodes[key] = ast.NewArgStm(expressions,token.NewToken(line,col))
		}

	}else{
			//call initial expression
			expr := gParser.nodes[eprKey].(ast.Expression)
			var expressions []ast.Expression
			expressions = append(expressions, expr)
			for _,termContext := range(termContexts) {
				_, _, expressionKey := GetTokenInfo(termContext)
				//expression call
				expression := gParser.nodes[expressionKey].(ast.Expression)
				expressions = append(expressions,expression)
			} 
			//new arg statement 
			gParser.nodes[key] = ast.NewArgStm(expressions,token.NewToken(line,col))
	}

}

func (gParser *goliteParser) ExitExpression(c *ExpressionContext) {
	//fmt.Printf("-----Got to Expression:%s-----\n",c.GetText())
	line, col, key := GetTokenInfo(c)
	termContexts := c.AllExpressionprime()
	_, _, eprKey := GetTokenInfo(c.Boolterm())
	//if no || and bool term then just log bool 
	if len(termContexts) == 0 {
		expr := gParser.nodes[eprKey].(ast.Expression)
		gParser.nodes[key] = expr
	} else {
		//if bool + || + bool ...
		//bool
		expr := gParser.nodes[eprKey].(ast.Expression)
		gParser.nodes[key] = expr
		for _,termContext := range(termContexts) {
			_, _, exprKey := GetTokenInfo(termContext.GetTrm())
			orFac := termContext.GetOr()
			boolExpr := gParser.nodes[exprKey].(ast.Expression)
			binOpEx := ast.NewBinOp(ast.StrToOp(orFac.GetText()),boolExpr,expr, token.NewToken(line,col))
			gParser.nodes[key] = binOpEx
			expr = binOpEx
		}
	}
}

func (gParser *goliteParser) ExitBoolTerm(c *BooltermContext) {
	line, col, key := GetTokenInfo(c)
	termContexts := c.AllBoolprime()
	_, _, eprKey := GetTokenInfo(c.Equalterm())
	
	if len(termContexts) == 0 {
		expr := gParser.nodes[eprKey].(ast.Expression)
		gParser.nodes[key] = expr
	} else {
		//call next token must be '&&'
		//call boolterm by using ast.callexpr.go
		expr := gParser.nodes[eprKey].(ast.Expression)
		//equal + && + equal ...
		for _,termContext := range(termContexts) {
			// && term
			_, _, boolKey := GetTokenInfo(termContext.GetTrm())
			andFac := termContext.GetAnd()
			equalExpr := gParser.nodes[boolKey].(ast.Expression)
			boolExpr := ast.NewBinOp(ast.StrToOp(andFac.GetText()),equalExpr,expr, token.NewToken(line,col))
			gParser.nodes[key] = boolExpr
			expr = boolExpr
		}
		}
	}

func (gParser *goliteParser) ExitEqualTerm(c *EqualtermContext) {
	line, col, key := GetTokenInfo(c)
	termContexts := c.AllEqualprime()
	_, _, eprKey := GetTokenInfo(c.Relationterm())
	
	if len(termContexts) == 0 {
		expr := gParser.nodes[eprKey].(ast.Expression)
		gParser.nodes[key] = expr
	} else {
		//call next token may be  '==' | '!='
		expr := gParser.nodes[eprKey].(ast.Expression)
		//relation + "==" + relation ...
		for _,termContext := range(termContexts) {
			// relation term
			_, _, equalKey := GetTokenInfo(termContext.GetTrm())
			eqFac := termContext.GetEq()
			equalExpr := gParser.nodes[equalKey].(ast.Expression)
			binOpExpr := ast.NewBinOp(ast.StrToOp(eqFac.GetText()), equalExpr, expr, token.NewToken(line, col))
			gParser.nodes[key] = binOpExpr
			expr = binOpExpr
		}
	}
}

func (gParser *goliteParser) ExitRelationTerm(c *RelationtermContext) {
	line, col, key := GetTokenInfo(c)
	termContexts := c.AllRelationprime()
	_, _, eprKey := GetTokenInfo(c.Simpleterm())
	
	if len(termContexts) == 0 {
		expr := gParser.nodes[eprKey].(ast.Expression)
		gParser.nodes[key] = expr
	} else {
		//call next token maybe '>' | '<' | '>=' | '<='
		//next is simple term
		expr := gParser.nodes[eprKey].(ast.Expression)
		// simple + '>' + simple ...
		for _,termContext := range(termContexts) {
			//comparison term
			_, _, simpleKey := GetTokenInfo(termContext.GetTrm())
			compFac := termContext.GetComp()
			simpleExpr := gParser.nodes[simpleKey].(ast.Expression)
			binOpExpr := ast.NewBinOp(ast.StrToOp(compFac.GetText()), simpleExpr, expr, token.NewToken(line, col))
			gParser.nodes[key] = binOpExpr
			expr = binOpExpr
		}
	}
}

func (gParser *goliteParser) ExitSimpleterm(c *SimpletermContext) {
	//fmt.Printf("-----Got to Simple Term:%s-----\n",c.GetText())
	
	line, col, key := GetTokenInfo(c)
	termContexts := c.AllSimpleprime()
	_, _, eprKey := GetTokenInfo(c.Term())

	if len(termContexts) == 0 {
		expr := gParser.nodes[eprKey].(ast.Expression)
		gParser.nodes[key] = expr
	} else {
		var termExpr ast.Expression 
		expr := gParser.nodes[eprKey].(ast.Expression)
		// term + '-' term 
		for _,termContext := range(termContexts){
			_, _, rightKey := GetTokenInfo(termContext.GetTrm())
			plusminFactor := termContext.GetPm()
			termExpr = gParser.nodes[rightKey].(ast.Expression)
			binoExpr := ast.NewBinOp(ast.StrToOp(plusminFactor.GetText()), termExpr, expr, token.NewToken(line, col))
			gParser.nodes[key] = binoExpr
			expr = 	binoExpr

		}
	}
}

func (gParser *goliteParser) ExitTerm(c *TermContext) {
	//fmt.Printf("-----Got to Term:%s-----\n",c.GetText())
	line, col, key := GetTokenInfo(c)
	termContexts := c.AllTermprime()
	_, _, eprKey := GetTokenInfo(c.Unaryterm())
	if len(termContexts) == 0 {
		expr := gParser.nodes[eprKey].(ast.Expression)
		gParser.nodes[key] = expr
	} else {
		//call af (asterisk forwardslash)
		//after call unary term
		expr := gParser.nodes[eprKey].(ast.Expression)
		//unary term '*'|'/'  unary term ...
		for _,termContext := range(termContexts){
			_, _, utKey := GetTokenInfo(termContext.GetTrm())
			afFac := termContext.GetAf()
			utExpr := gParser.nodes[utKey].(ast.Expression)
			binOpExpr := ast.NewBinOp(ast.StrToOp(afFac.GetText()),utExpr ,expr ,token.NewToken(line, col))
			gParser.nodes[key] = binOpExpr
			expr = binOpExpr
			}
	}
}

func (gParser *goliteParser) ExitUnaryTerm(c *UnarytermContext) {
	//fmt.Printf("-----Got to Exit Unary Term:%s-----\n",c.GetText())
	line, col, key := GetTokenInfo(c)
	// ! selectorTerm
	 if notTerm := c.NOT(); notTerm != nil {
		//!
		_, _, selKey := GetTokenInfo(c.Selectorterm())
		//gParser.nodes[key] = ast.Expression(c.NOT().GetText(), token.NewToken(line, col))
		//selector term
		stExpr := gParser.nodes[selKey].(ast.Expression)
		gParser.nodes[key] = stExpr
		gParser.nodes[key] = ast.NewUnaryOp(ast.StrToOp(notTerm.GetText()), stExpr, token.NewToken(line, col))

	// - selectorTerm
	}else if negTerm := c.SUB(); negTerm != nil {
		//-
		_, _, selKey := GetTokenInfo(c.Selectorterm())
		//selector term
		stExpr := gParser.nodes[selKey].(ast.Expression)
		gParser.nodes[key] = stExpr
		gParser.nodes[key] = ast.NewUnaryOp(ast.StrToOp(notTerm.GetText()), stExpr, token.NewToken(line, col))
		//this exists for the SelectorTerm Option
		}else {
		_, _, selKey := GetTokenInfo(c.Selectorterm())
		expr := gParser.nodes[selKey].(*ast.Arguments)
		// gParser.nodes[key] = expr
		// listExpr = append(listExpr, expr)
		gParser.nodes[key] = ast.NewCallExpr(selKey,expr,token.NewToken(line,col))
		//still need to make a call with an ast node pro
}
}

func (gParser *goliteParser) ExitSelectorTerm(c *SelectortermContext) {
	line,col, key := GetTokenInfo(c)
	termContexts := c.AllSelectorprime()
	// _ , _, idKey := GetTokenInfo(c.Factor())
	var idList []string
	//this exists for the SelectorTerm Option
	_, _, facKey := GetTokenInfo(c.Factor())
	if len(termContexts) == 0{
		expr := gParser.nodes[facKey].(ast.Expression)
		gParser.nodes[key] = expr
	} else {
		// this exist for factor .'id' ....
		// still need to call for factor
		//..... need to confirm call CallExpr
		expr := gParser.nodes[facKey].(ast.Expression)
		gParser.nodes[key] = expr
		
		//now for the .'id' ...
		for _, termContext := range(termContexts){
			idFactor := termContext.GetId().GetText()
			idList = append(idList,idFactor)
		}
		gParser.nodes[key] = ast.NewAttributeCall(expr,idList, token.NewToken(line, col))
	}
}


func (gParser *goliteParser) ExitFactor(c *FactorContext) {
	//fmt.Printf("-----Got to Exit Factor:%s-----\n",c.GetText())
	line, col, key := GetTokenInfo(c)
	//expression
	if exprCtx := c.GetExpr(); exprCtx != nil { 
		var expr ast.Expression 
		_, _, exprKey := GetTokenInfo(exprCtx)
		expr = gParser.nodes[exprKey].(ast.Expression)
		gParser.nodes[key] = expr
	//id [arguments]
	//
	} else if idFactor := c.ID(); idFactor != nil {
		//check to see if it is a function call
		if parenFac := c.LPAREN() ; parenFac != nil {
			argumentsFactor := c.GetArg()
			_, _, argsKey := GetTokenInfo(argumentsFactor)
			arg := gParser.nodes[argsKey].(*ast.Arguments)
			gParser.nodes[key] = ast.NewCallExpr(idFactor.GetText(),arg, token.NewToken(line, col))
		}else { // we know it is just an id, already been declared.
			gParser.nodes[key] = ast.NewVariable(idFactor.GetText(), token.NewToken(line, col))
		}
		if  argumentsFactor := c.GetArg(); argumentsFactor != nil  {
			_, _, argsKey := GetTokenInfo(argumentsFactor)
			//check to see if it is a function call
			
			a := gParser.nodes[argsKey].(*ast.Arguments)
			gParser.nodes[key] = ast.NewVariable(idFactor.GetText(), token.NewToken(line, col))
			ast.NewCallExpr(idFactor.GetText(),a,token.NewToken(line,col))
		}
	//'new' 'id'
	}else if newFactor := c.NEW(); newFactor != nil {
		if idFactor := c.ID() ;idFactor != nil{
			//newLit will have both new and id only one call
			gParser.nodes[key] = ast.NewNewLit(idFactor.GetText(), token.NewToken(line, col))
		}
	//int
	}else if intFactor := c.NUMBER(); intFactor != nil {
		intValue, _ := strconv.Atoi(intFactor.GetText())
		gParser.nodes[key] = ast.NewIntLit(int64(intValue), token.NewToken(line, col))
		//true
	} else if trueFactor := c.TRUE(); trueFactor != nil {
		gParser.nodes[key] = ast.NewBoolLit(true, token.NewToken(line, col))
		//false
	} else if falseFactor := c.FALSE(); falseFactor != nil {
		gParser.nodes[key] = ast.NewBoolLit(false, token.NewToken(line, col))		
	
	//nil
	//} else if NIL := c.NIL(); NIL == nil {
	} else if NIL := c.NIL(); NIL != nil {
		gParser.nodes[key] = ast.NewVariable(NIL.GetText(), token.NewToken(line, col))
	}
}
 
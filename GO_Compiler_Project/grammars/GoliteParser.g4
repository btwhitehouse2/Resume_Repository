parser grammar GoliteParser; 

options {
    tokenVocab = GoliteLexer;
}

program: tyDec=typedeclaration* dec=declaration* func=function* EOF ;

typedeclaration: TYPE ID STRUCT LCURLYBRACE flds=fields  RCURLYBRACE SEMICOLON ;
fields: dec=decl SEMICOLON fieldsprime* ;
fieldsprime: decl SEMICOLON  ;

decl: id=ID ty=type ;

type: INT | BOOL | (ASTERISK ID) ;

declaration: varlit=VAR id=ID declarationprime* type SEMICOLON ;
declarationprime:  COMMA id=ID;

function: FUNC ID parameters funcprime? LCURLYBRACE declaration* statement* RCURLYBRACE ;
funcprime: type ;
parameters: LPAREN (dec=decl parametersprime*)?  RPAREN ;
parametersprime: COMMA decl ;
statement: block | ass=assignment | print | del=delete | conditional | loop | rtrn | read | inv=invocation ;

block: LCURLYBRACE statement* RCURLYBRACE ;

delete: DELETE expr=expression SEMICOLON ;

assignment: ID assignmentprime* EQUALS expression SEMICOLON ;
assignmentprime: (PERIOD ID) ; 
print: PRINTF LPAREN STRING (COMMA exp=expression)* RPAREN SEMICOLON ;

conditional: cond=IF LPAREN expression RPAREN block conditionalprime? ;
conditionalprime: (cond=ELSE block) ; 
loop: lp=FOR LPAREN expression RPAREN block ;

rtrn: RETURN rtrnprime? SEMICOLON ;
rtrnprime: expression ;
read: SCAN expression SEMICOLON ;

invocation: ID arguments SEMICOLON ;

arguments: (LPAREN (expression argumentsprime*)? RPAREN) ;
argumentsprime: (COMMA expression) ;
expression: bt=boolterm expressionprime* ; // create primes
expressionprime: (or=OR trm=boolterm) ;
boolterm: eq=equalterm boolprime* ;
boolprime: (and=ANDCOMP trm=equalterm) ;
equalterm: rt=relationterm equalprime* ;
equalprime: (eq=(ISEQUAL|NOTEQUAL) trm=relationterm) ;
relationterm: st=simpleterm relationprime* ;
relationprime: (comp=(GREATERTHAN|LESSTHAN|LESSEQUAL|GREATEREQUAL) trm=simpleterm) ;
simpleterm: t=term simpleprime* ;
simpleprime: (pm=(PLUS|SUB) trm=term) ;
term: ut=unaryterm termprime* ;
termprime: (af=(ASTERISK|FSLASH) trm=unaryterm) ;
unaryterm: (not=NOT st=selectorterm) | (neg=SUB selectorterm) | selectorterm ;

selectorterm: f=factor selectorprime* ;
selectorprime: (PERIOD id=ID) ;
factor: (LPAREN expr=expression RPAREN) | ID arg=arguments? | NUMBER | NEW ID | TRUE | FALSE | NIL ;
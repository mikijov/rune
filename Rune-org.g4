grammar Rune;

// put imports here
@header {
}

// generally not needed at all as everything in current and imported packages is accessible
@members {
}

/* tokens { EOF } */

module: stmt* EOF
    ;

stmt
    : simpleStmt
    | compoundStmt
    ;
simpleStmt
    : ( declarationStmt
      | returnStmt
      | breakStmt
      | continueStmt
      | expression
      ) ';'
    ;
declarationStmt
    : (NAME ':' typeName)
    | declarationExpr
    ;
returnStmt
    : 'return'
    | 'return' expression
    ;
breakStmt
    : 'break' NAME?
    ;
continueStmt
    : 'continue' NAME?
    ;
compoundStmt
    : ifStmt
    | loopWhile
    | loopUntil
    | loopForEach
    ;
ifStmt
    : 'if' expression scope ('else' 'if' expression scope)* ('else' scope)?
    ;
loopWhile
    : 'loop' NAME? 'while' expression scope
    ;
loopUntil
    : 'loop' NAME? 'until' expression scope
    ;
loopForEach
    : 'loop' NAME? 'foreach' NAME 'in' expression scope
    ;
scope: '{' stmt* '}'
    ;

typeName: 'int' | 'string' | 'bool' | 'list' | 'map'
    ;

expression: declarationExpr
    ;
declarationExpr returns [expr :Expression]
    : assignment
    | assignment ':=' assignment
    | assignment ':' typeName '=' assignment
    ;
assignment returns [expr :Expression]
    : logicalOr
    | logicalOr assignmentOp logicalOr
    ;
assignmentOp
    : '+=' | '-=' | '*=' | '/=' | '%=' | '&=' |
      '|=' | '^=' | '<<=' | '>>=' | '**=' | '//='
    ;
logicalOr returns [expr :Expression]
    : logicalAnd
    | left=logicalAnd op='or' right=logicalAnd
    { $expr = NewBinaryExpression($left.expr, $op.text, $right.expr) }
    ;
logicalAnd returns [expr :Expression]
    : logicalNot
    | left=logicalNot op='and' right=logicalNot
    { $expr = NewBinaryExpression($left.expr, $op.text, $right.expr) }
    ;
logicalNot returns [expr :Expression]
    : op='not' logicalNot
    { fmt.Printf("op: %s\n", $op.text) }
    | comparison
    ;
comparison returns [expr :Expression]
    : orExpr
    | left=orExpr op=compOp right=orExpr
    { $expr = NewBinaryExpression($left.expr, $op.text, $right.expr) }
    ;
compOp
    : '<'|'>'|'=='|'>='|'<='|'!='|'in'|'not' 'in'|'is'|'is' 'not'
    ;
orExpr returns [expr :Expression]
    : xorExpr
    | left=xorExpr op='|' right=xorExpr
    ;
xorExpr returns [expr :Expression]
    : andExpr
    | left=andExpr op='^' right=andExpr
    ;
andExpr returns [expr :Expression]
    : shiftExpr
    | left=shiftExpr op='&' right=shiftExpr
    ;
shiftExpr returns [expr :Expression]
    : arithExpr
    | left=arithExpr op=('<<'|'>>') right=arithExpr
    ;
arithExpr returns [expr :Expression]
    : term
    | left=term op=('+'|'-') right=term
    { $expr = NewBinaryExpression($left.expr, $op.text, $right.expr) }
    ;
term returns [expr :Expression]
    : factor
    | left=factor op=('*'|'/'|'%'|'//') right=factor
    { $expr = NewBinaryExpression($left.expr, $op.text, $right.expr) }
    ;
factor returns [expr :Expression]
    : ('+'|'-'|'~') factor | power
    ;
power
    : atom trailer* ('**' factor)?
    ;
trailer
    : tuple
    | '[' expression ']'
    | '.' NAME
    ;
atom returns [expr :Expression]
    : lambda
    | array
    | map_
    | val=NAME
    | val=NUMBER { $expr = NewIntegerLiteral($val.text) }
    | STRING
    ;

lambda
    : 'func' paramDecl (':' typeName)? scope
    ;
paramDecl
    : '(' (combinedParam (',' combinedParam)*)? ')'
    ;
combinedParam
    : NAME (',' NAME)* ':' typeName
    ;

tuple
    : '(' expressionList ')'
    ;
array
    : '[' expressionList ']'
    ;
expressionList
    : expression (',' expression)? (',')?
    ;
map_
    : '{' ( mapitem (',' mapitem) (',')? )? '}'
    ;
mapitem
    : NAME ':' expression
    ;

NAME: [a-zA-Z_] [a-zA-Z0-9_]*
    ;

NUMBER
    :   '0' ([xX] [0-9a-fA-F]+  ([eE] [+-]? [0-9]+)?
    |        [oO] [0-7]+
    |        [bB] [01]+
            )
    | ([0-9]+ '.' [0-9]* | '.' [0-9]+) ([eE] [+-]? [0-9]+)?
    |  [0-9]+                          ([eE] [+-]? [0-9]+)?
    ;

STRING
    :
    ( '"'    ('\\' (([ \t]+ ('\r'? '\n')?)|.) | ~[\\\r\n"])*  '"'
    | '"""'  ('\\' .                          | ~'\\'     )*? '"""'
    )
    ;

LINENDING: '\r'? '\n' -> skip
    ;

WHITESPACE: ('\t' | ' ')+ -> skip
    ;

COMMENT: '#' ~[\r\n]* -> skip;

grammar Rune;

// put imports here
@header {
import (
    "mikijov/rune/vm"
)

var _ vm.Type // inhibit unused import error
}

// generally not needed at all as everything in current and imported packages is accessible
@members {
}

/* tokens { EOF } */

program
    : statements+=statement* EOF
    ;

statement
    : declaration ';'
    | returnStatement ';'
    | expression ';'
    | function
    | ifStatement
    | loop
    ;

declaration
    : identifier=IDENTIFIER ':=' value=expression
    | identifier=IDENTIFIER ':' type_=typeName ('=' value=expression)?
    ;
typeName
    : 'int' | 'real' | 'string' | 'bool'
    // | 'list' | 'map'
    ;

function
    : 'func' identifier=IDENTIFIER params=paramDecl (':' returnType=typeName)? body=scope
    ;
paramDecl
    : '(' (paramGroup+=combinedParam (',' paramGroup+=combinedParam)*)? ')'
    ;
combinedParam
    : names+=IDENTIFIER (',' names+=IDENTIFIER)* ':' paramType=typeName
    ;
scope
    : '{' statements+=statement* '}'
    ;

returnStatement
    : 'return' (retVal=expression)?
    ;

ifStatement
    : 'if' conditions+=expression effects+=scope
            ('else' 'if' conditions+=expression effects+=scope)*
            ('else' alternative=scope)?
    ;

loop: 'loop' (kind=('while'|'until') condition=expression)? body=scope
    ;

expression: expression2
    ;

expression2
    : '(' value=expression2 ')' # ExpressionPassthrough
    | op='-' value=expression2 # UnaryExpression
    | left=IDENTIFIER op=assignmentOp right=expression2 # Assignment
    | left=expression2 op='or' right=expression2 # BinaryExpression
    | left=expression2 op='and' right=expression2 # BinaryExpression
    | left=expression2 op=('<'|'>'|'=='|'>='|'<='|'!=') right=expression2 # BinaryExpression
    | left=expression2 op='|' right=expression2 # BinaryExpression
    | left=expression2 op='^' right=expression2 # BinaryExpression
    | left=expression2 op='&' right=expression2 # BinaryExpression
    | left=expression2 op=('*'|'/'|'%') right=expression2 # BinaryExpression
    | left=expression2 op=('+'|'-') right=expression2 # BinaryExpression
    | value=literal # LiteralPassthrough
    | name=IDENTIFIER # VariableExpression
    | name=IDENTIFIER '(' (params+=expression (',' params+=expression)*)? ')' # FunctionCall
    ;

assignmentOp
    /* : '+=' | '-=' | '*=' | '/=' | '%=' | '&=' | '|=' | '^=' | '<<=' | '>>=' | '**=' | '//=' */
    : '=' | '+=' | '-=' | '*=' | '/=' | '%=' | '&=' | '|=' | '^='
    ;
literal
    : value=REAL_LITERAL # RealLiteral
    | value=INTEGER_LITERAL # IntegerLiteral
    | value=BOOLEAN_LITERAL # BooleanLiteral
    ;

INTEGER_LITERAL: [0-9]+ ;
REAL_LITERAL: [0-9]* '.' [0-9]+ ;
BOOLEAN_LITERAL: 'true' | 'false' ;

IDENTIFIER: [a-zA-Z_] [a-zA-Z0-9_]* ;

LINENDING: '\r'? '\n' -> skip;

WHITESPACE: ('\t' | ' ')+ -> skip;

COMMENT: '#' ~[\r\n]* -> skip;

grammar Rune;

// put imports here
@header {
import (
    "mikijov/rune-antlr/vm"
)

var _ vm.Type // inhibit unused import error
}

// generally not needed at all as everything in current and imported packages is accessible
@members {
}

/* tokens { EOF } */

program
    : statement* EOF
    ;

statement
    : declaration ';'
    | expression ';'
    ;

declaration
    : identifier=IDENTIFIER ':=' value=expression
    | identifier=IDENTIFIER ':' type_=typeName ('=' value=expression)?
    ;
typeName
    : 'int' | 'real' | 'string' | 'bool'
    // | 'list' | 'map'
    ;

expression: expression2 ;

expression2
    : '(' value=expression2 ')' # ExpressionPassthrough
    | op='-' value=expression2 # UnaryExpression
    | left=expression2 op=('*'|'/'|'%') right=expression2 # BinaryExpression
    | left=expression2 op=('+'|'-') right=expression2 # BinaryExpression
    | value=literal # LiteralPassthrough
    ;

literal
    : value=REAL_LITERAL # RealLiteral
    | value=INTEGER_LITERAL # IntegerLiteral
    ;

INTEGER_LITERAL: [0-9]+ ;
REAL_LITERAL: [0-9]* '.' [0-9]+ ;
BOOLEAN_LITERAL: 'true' | 'false' ;

IDENTIFIER: [a-zA-Z_] [a-zA-Z0-9_]* ;

LINENDING: '\r'? '\n' -> skip;

WHITESPACE: ('\t' | ' ')+ -> skip;

COMMENT: '#' ~[\r\n]* -> skip;

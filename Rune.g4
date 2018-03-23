grammar Rune;

// put imports here
@header {
import (
    "github.com/mikijov/rune/vm"
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

//programStatement
//: declaration
//| typeDeclaration
//| function
//;
statement
    : declaration
    | typeDeclaration
    | function
    | returnStatement
    | expression ';'
    | ifStatement
    | loop
    ;

declaration
    : identifier=IDENTIFIER ':=' value=expression ';'
    | identifier=IDENTIFIER ':' type_=typeName ('=' value=expression ';')?
    ;
typeDeclaration
    : 'type' identifier=IDENTIFIER ':' type_=typeName
    ;
typeName: typeName2 (isarray='[' ']')?;
typeName2
    : 'int' # SimpleType
    | 'real' # SimpleType
    | 'string' # SimpleType
    | 'bool' # SimpleType
    | 'func' '(' (paramTypes+=typeName2 (',' paramTypes+=typeName2)*)? ')' (':' returnType=typeName2)? # FunctionType
    | 'struct' '{' combinedField* '}' # StructType
    | name=IDENTIFIER # CustomType
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
combinedField
    : names+=IDENTIFIER (',' names+=IDENTIFIER)* ':' paramType=typeName
    ;
scope
    : '{' statements+=statement* '}'
    ;

returnStatement
    : 'return' (retVal=expression ';')?
    ;

ifStatement
    : 'if' conditions+=expression effects+=scope
            ('else' 'if' conditions+=expression effects+=scope)*
            ('else' alternative=scope)?
    ;

// loop: 'loop' (kind=('while'|'until') condition=expression)? body=scope
loop: 'loop' (kind='while' condition=expression)? body=scope
    ;

expression: expression2;
expression2
    : '(' value=expression2 ')' # ExpressionPassthrough
    | name=IDENTIFIER '(' (params+=expression2 (',' params+=expression2)*)? ')' # FunctionCall
    | base=expression2 '.' field=IDENTIFIER # FieldSelector
    | array=expression2 '[' index=expression2 ']' # ArraySelector
    | op=unaryOp value=expression2 # UnaryExpression
    | left=expression2 op=('*'|'/'|'%'|'&') right=expression2 # BinaryExpression
    | left=expression2 op=('+'|'-'|'|'|'^') right=expression2 # BinaryExpression
    // | left=expression2 op=('<<'|'>>') right=expression2 # BinaryExpression
    | left=expression2 op=('<'|'>'|'=='|'>='|'<='|'!=') right=expression2 # BinaryExpression
    | left=expression2 op='and' right=expression2 # BinaryExpression
    | left=expression2 op='or' right=expression2 # BinaryExpression
    | left=expression2 op=assignmentOp right=expression2 # Assignment
    | value=literal # LiteralPassthrough
    | 'func' params=paramDecl (':' returnType=typeName)? body=scope # Lambda
    | name=IDENTIFIER # VariableExpression
    ;

unaryOp
    /* 'not' | '!' | '+' | '*' | '&' | '~' etc... */
    : '-'
    ;
assignmentOp
    /* : '<<=' | '>>=' | '**=' | '//=' */
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

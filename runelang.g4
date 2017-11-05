grammar runelang;

tokens { INDENT, DEDENT, NEWLINE, ENDMARKER }

module: (NEWLINE | expressionList)* ENDMARKER
    ;
expressionList: expression (',' expression)? (',')?
    ;
expression
    : function
    | declaration
    | test
    ;

function: 'func' NAME parameterDeclaration type_? expressionList
    ;
parameterDeclaration: '(' (combinedParameter (',' combinedParameter)*)? ')'
    ;
combinedParameter: NAME (',' NAME)* type_
    ;

declaration
    : NAME ( (type_ ('=' test)?)
           | (':=' test)
           )
    ;

type_: ':' typeName
    ;
typeName: 'int' | 'string' | 'bool' | 'list' | 'map_'
    ;


test: or_test ('if' or_test 'else' test)?
    ;
or_test: and_test ('or' and_test)*
    ;
and_test: not_test ('and' not_test)*
    ;
not_test: 'not' not_test | comparison
    ;
comparison: expr (comp_op expr)*
    ;
comp_op: '<'|'>'|'=='|'>='|'<='|'!='|'in'|'not' 'in'|'is'|'is' 'not'
    ;
expr: xor_expr ('|' xor_expr)*
    ;
xor_expr: and_expr ('^' and_expr)*
    ;
and_expr: shift_expr ('&' shift_expr)*
    ;
shift_expr: arith_expr (('<<'|'>>') arith_expr)*
    ;
arith_expr: term (('+'|'-') term)*
    ;
term: factor (('*'|'/'|'%'|'//') factor)*
    ;
factor: ('+'|'-'|'~') factor | power
    ;
power: atom trailer* ('**' factor)?
    ;
trailer
    : '(' expressionList? ')'
    | '[' expression ']'
    | '.' NAME
    ;
atom: lambda | list | map_ | NAME | NUMBER | STRING
    ;
lambda: 'func' parameterDeclaration type_? expressionList
    ;
list: '(' expressionList ')'
    ;
map_: '{' ( kvpair (',' kvpair) (',')? )? '}'
    ;
kvpair: NAME ':' expression
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

LINENDING:             (('\r'? '\n')+
    |      '\\'  [ \t]* ('\r'? '\n')  )
{
}   -> channel(HIDDEN)
    ;

WHITESPACE: ('\t' | ' ')+
{
}   -> channel(HIDDEN)
    ;

COMMENT: '#' ~[\r\n]* -> skip;

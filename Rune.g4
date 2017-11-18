grammar Rune;

// put imports here
@header {
}

// generally not needed at all as everything in current and imported packages is accessible
@members {
}

/* tokens { EOF } */

module
    returns [m :Module]
    @init {
        $m = NewModule()
    }
    @after {
        for _, stmt := range $ctx.AllStmt() {
            $m.AddStatement(stmt.GetS())
        }
    }
    : stmt* EOF
    ;

stmt returns [s :Statement]
    : simpleStmt
    { $s = $simpleStmt.s }
    ;
simpleStmt returns [s :Statement]
    :
    ( expression
    { $s = &expressionStatement{$expression.expr} }
    ) ';'
    ;

expression returns [expr :Expression]
    : e=arithExpr
    { $expr = $e.expr }
    ;
arithExpr returns [expr :Expression]
    : e=term
    { $expr = $e.expr }
    | left=arithExpr op=('+'|'-') right=term
    { $expr = NewBinaryExpression($left.expr, $op.text, $right.expr) }
    ;
term returns [expr :Expression]
    : e=atom
    { $expr = $e.expr }
    | left=term op=('*'|'/'|'%') right=atom
    { $expr = NewBinaryExpression($left.expr, $op.text, $right.expr) }
    ;
atom returns [expr :Expression]
    : val=INTEGER_LITERAL
    { $expr = NewIntegerLiteral($val.text) }
    | val=REAL_LITERAL
    { $expr = NewRealLiteral($val.text) }
    ;

INTEGER_LITERAL: [0-9]+ ;
REAL_LITERAL: [0-9]* '.' [0-9]+ ;

LINENDING: '\r'? '\n' -> skip;

WHITESPACE: ('\t' | ' ')+ -> skip;

COMMENT: '#' ~[\r\n]* -> skip;

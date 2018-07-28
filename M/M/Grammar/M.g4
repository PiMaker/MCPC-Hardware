grammar M;

calcExpression
    :   Identifier
    |   StringLiteral
	|   Constant
	|   '{' initializerList '}'
	|   calcExpression calcOperator calcExpression
	|   unaryCalcOperator calcExpression
    |   calcExpression unaryCalcOperator
	|   '(' typeSpecifier ')' calcExpression
	|   '(' calcExpression ')'
	|   typeSpecifier '(' calcExpression ')'
    ;

calcOperator
    :   (Less | LessEqual | Greater | GreaterEqual | Plus | Minus | Star | Div | Mod | And | Or | AndAnd | OrOr | Xor)
	;

unaryCalcOperator
    :   (PlusPlus | MinusMinus | Tilde | Not | And)
	;

initializerList
    :   Constant
    |   initializerList ',' Constant
    ;

assignmentExpression
    :   typeSpecifier? Identifier ('=' calcExpression)?
    ;

topLevelAssignmentExpression
    :   typeSpecifier Identifier '=' calcExpression
    ;

typeSpecifier
    :   ('char'
    |   'int'
	|   'void')
    ;

nestedParenthesesBlock
    :   (   ~('(' | ')')
        |   '(' nestedParenthesesBlock ')'
        )*
    ;

parameterDeclarationList
    :   parameterDeclaration
    |   parameterDeclarationList ',' parameterDeclaration
    ;

parameterDeclaration
    :   typeSpecifier Identifier
    ;

paramterPassList
    :   calcExpression
    |   paramterPassList ',' calcExpression
    ;

parameterCallList
    :   calcExpression
    |   parameterCallList ',' calcExpression
    ;

statement
    :   labeledStatement
    |   compoundStatement
    |   assignmentStatement
    |   selectionStatement
    |   iterationStatement
	|   callStatement
    |   jumpStatement
	|   asmStatement
	|   shiftStatement
    ;
	
shiftStatement
    :   Identifier (LeftShift | RightShift) Constant ';'
	;

assignmentStatement
    :   assignmentExpression ';'
	;

callStatement
    :   Identifier '(' parameterCallList ')' ';'
	;

labeledStatement
    :   Identifier ':' statement
    ;

compoundStatement
    :   '{' statementList? '}'
    ;

statementList
    :   statement
    |   statementList statement
    ;

selectionStatement
    :   'if' '(' calcExpression ')' statement ('else' statement)?
    ;

iterationStatement
    :   While '(' calcExpression ')' statement
    |   Do statement While '(' calcExpression ')' ';'
    |   For '(' forCondition ')' statement
    ;

forCondition
	:   assignmentExpression ';' calcExpression? ';' assignmentExpression?
	;

jumpStatement
    :   'goto' Identifier ';'
    |   'continue' ';'
    |   'break' ';'
    |   'return' calcExpression? ';'
    ;

compilationUnit
    :   translationUnit EOF
    ;

translationUnit
    :   topLevelDeclaration
    |   translationUnit topLevelDeclaration
    ;

topLevelDeclaration
    :   functionDefinition
	|   topLevelAssignmentExpression
	|   preprocessorDirective
    ;

functionDefinition
    :   'inline'? typeSpecifier Identifier '(' parameterDeclarationList? ')' statement
    ;

asmStatement
    :   'asm' '{'* '{' ~'}'* '}'
    ;
	
preprocessorDirective
	:   '#' 'include' Filename
	;

Break : 'break';
Char : 'char';
Continue : 'continue';
Do : 'do';
Else : 'else';
For : 'for';
Goto : 'goto';
If : 'if';
Inline : 'inline';
Int : 'int';
Return : 'return';
Void : 'void';
While : 'while';

LeftParen : '(';
RightParen : ')';
LeftBracket : '[';
RightBracket : ']';
LeftBrace : '{';
RightBrace : '}';

Less : '<';
LessEqual : '<=';
Greater : '>';
GreaterEqual : '>=';
LeftShift : '<<';
RightShift : '>>';

Plus : '+';
PlusPlus : '++';
Minus : '-';
MinusMinus : '--';
Star : '*';
Div : '/';
Mod : '%';

And : '&';
Or : '|';
AndAnd : '&&';
OrOr : '||';
Xor : '^';
Not : '!';
Tilde : '~';

Question : '?';
Colon : ':';
Semi : ';';
Comma : ',';

Assign : '=';

Equal : '==';
NotEqual : '!=';
Dot : '.';

Identifier
    :   IdentifierNondigit
        (   IdentifierNondigit
        |   Digit
        )*
    ;

Filename
	:   Identifier '.m'
	;

fragment
IdentifierNondigit
    :   Nondigit
    |   UniversalCharacterName
    ;

fragment
Nondigit
    :   [a-zA-Z_]
    ;

fragment
Digit
    :   [0-9]
    ;

fragment
UniversalCharacterName
    :   '\\u' HexQuad
    |   '\\U' HexQuad HexQuad
    ;

fragment
HexQuad
    :   HexadecimalDigit HexadecimalDigit HexadecimalDigit HexadecimalDigit
    ;

Constant
    :   IntegerConstant
    |   CharacterConstant
    ;

fragment
IntegerConstant
    :   DecimalConstant
    |   OctalConstant
    |   HexadecimalConstant
    |	BinaryConstant
    ;

fragment
BinaryConstant
	:	'0' [bB] [0-1]+
	;

fragment
DecimalConstant
    :   NonzeroDigit Digit*
    ;

fragment
OctalConstant
    :   '0' OctalDigit*
    ;

fragment
HexadecimalConstant
    :   HexadecimalPrefix HexadecimalDigit+
    ;

fragment
HexadecimalPrefix
    :   '0' [xX]
    ;

fragment
NonzeroDigit
    :   [1-9]
    ;

fragment
OctalDigit
    :   [0-7]
    ;

fragment
HexadecimalDigit
    :   [0-9a-fA-F]
    ;

fragment
ExponentPart
    :   'e' Sign? DigitSequence
    |   'E' Sign? DigitSequence
    ;

fragment
Sign
    :   '+' | '-'
    ;

DigitSequence
    :   Digit+
    ;

fragment
BinaryExponentPart
    :   'p' Sign? DigitSequence
    |   'P' Sign? DigitSequence
    ;

fragment
HexadecimalDigitSequence
    :   HexadecimalDigit+
    ;

fragment
CharacterConstant
    :   '\'' CCharSequence '\''
    |   'L\'' CCharSequence '\''
    |   'u\'' CCharSequence '\''
    |   'U\'' CCharSequence '\''
    ;

fragment
CCharSequence
    :   CChar+
    ;

fragment
CChar
    :   ~['\\\r\n]
    |   EscapeSequence
    ;

fragment
EscapeSequence
    :   SimpleEscapeSequence
    |   OctalEscapeSequence
    |   HexadecimalEscapeSequence
    |   UniversalCharacterName
    ;

fragment
SimpleEscapeSequence
    :   '\\' ['"?abfnrtv\\]
    ;

fragment
OctalEscapeSequence
    :   '\\' OctalDigit
    |   '\\' OctalDigit OctalDigit
    |   '\\' OctalDigit OctalDigit OctalDigit
    ;

fragment
HexadecimalEscapeSequence
    :   '\\x' HexadecimalDigit+
    ;

StringLiteral
    :   '"' SCharSequence? '"'
    ;

fragment
SCharSequence
    :   SChar+
    ;

fragment
SChar
    :   ~["\\\r\n]
    |   EscapeSequence
    |   '\\\n'   // Added line
    |   '\\\r\n' // Added line
    ;

Whitespace
    :   [ \t\r\n]+
        -> skip
    ;

Newline
    :   (   '\r' '\n'?
        |   '\n'
        )
        -> skip
    ;

BlockComment
    :   '/*' .*? '*/'
        -> skip
    ;

LineComment
    :   '//' ~[\r\n]*
        -> skip
    ;

swiftFile: Statement*;

Statement: (Declaration  | Expression | DoBlock | ReturnStmt) '\n';
Expression: Try? (DeclareVarExpr | DeclareLetExpr | AssignVarExpr | FunctionCall | Identifier | ArrayLiteral | IfElseExpr | PredicateExpr);
FunctionCall: FunctionName '(' IdentifierList? ')' ClosureExpr? NewLine?;
ArrayLiteral: '[' ExpressionList ']';
ExpressionList: Expression ExpressionC*;
ExpressionC: ',' Expression;
IfElseExpr: 'if' Expression StmtBlock;
PredicateExpr: Lhs CompareOp Rhs;
Lhs: Expression;
Rhs: Expression;
DoBlock: 'do' StmtBlock CatchBlock?;
CatchBlock: 'catch' WhereClause? StmtBlock;
WhereClause: 'where' Expression;
ReturnStmt: 'return' Expression?;
Declaration: StructDecl | ImportDecl | ConstantDecl | VarDecl | EnumDecl | ClassDecl | FuncDecl | InitFuncDecl | TypeAlias;

ImportDecl: Attribute* 'import' ImportKind? ImportPath;
ImportKind: 'typealias' | 'struct' | 'class' | 'enum' | 'protocol' | 'let' | 'var' | 'func';

ConstantDecl: AccessModifier? 'let' Identifier ':' TypeName;
VarDecl: AccessModifier? 'var' Identifier ':' TypeName;
DeclareVarExpr: 'var' VarName '=' Expression;
DeclareLetExpr: 'let' VarName '=' Expression;
AssignVarExpr: VarName '=' Expression;

TypeAlias: 'typealias' TypeName '=' ExistingType;

FuncDecl: AccessModifier? 'func' FuncName '(' Parameters? ')' StmtBlock?;
InitFuncDecl: AccessModifier? FuncName '(' Parameters? ')' StmtBlock?;
Parameters: Parameter ParameterC*;
ParameterC: ',' Parameter;
Parameter: Attribute? Label? ParamName ':' TypeName;
StmtBlock: '{' '\n' Statement* '\n}' '\n';
ClosureExpr: '{' ClosureSignature?  Statement* '\n}' '\n';
ClosureSignature: IdentifierList 'in' '\n';
StructDecl: 'struct' StructName Body;
Body: '{' '\n' Member+ '}' '\n';
Member: '\t' Declaration '\n';

ClassDecl: 'class' ClassName Body;

EnumDecl: 'enum' EnumName Inherit? EnumBlock;
Inherit: ':' IdentifierList;
IdentifierList: (identifier | Expression) IdentifierListC*;
IdentifierListC: ',' (identifier | Expression);

EnumBlock: '{' '\n' EnumMember+ '}' '\n';
EnumMember: Declaration | EnumCase;
EnumCase: 'case' EnumCaseName EnumValueAssignment?;
EnumValueAssignment: '=' EnumRawValue '\n';


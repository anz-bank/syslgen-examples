swiftFile: Statement*;

Statement: Declaration | Expression ;

Expression: DeclareVarExpr | DeclareLetExpr | AssignVarExpr | FunctionCall | Identifier | ArrayLiteral | IfElseExpr | PredicateExpr;
FunctionCall: FunctionName '(' IdentifierList? ')' ClosureExpr? NewLine?;
ArrayLiteral: '[' ExpressionList ']';
ExpressionList: Expression ExpressionC*;
ExpressionC: ',' Expression;
IfElseExpr: 'if' Expression StmtBlock;
PredicateExpr: Lhs CompareOp Rhs;
Lhs: Expression;
Rhs: Expression;
Declaration: StructDecl | ImportDecl | ConstantDecl | VarDecl | EnumDecl | ClassDecl | FuncDecl | InitFuncDecl | TypeAlias;

ImportDecl: Attribute* 'import' ImportKind? ImportPath '\n';
ImportKind: 'typealias' | 'struct' | 'class' | 'enum' | 'protocol' | 'let' | 'var' | 'func';

ConstantDecl: AccessModifier? 'let' Identifier ':' TypeName;
VarDecl: AccessModifier? 'var' Identifier ':' TypeName '\n';
DeclareVarExpr: 'var' VarName '=' Expression '\n';
DeclareLetExpr: 'let' VarName '=' Expression '\n';
AssignVarExpr: VarName '=' Expression '\n';

TypeAlias: 'typealias' TypeName '=' ExistingType '\n';

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
Member: '\t' Declaration;

ClassDecl: 'class' ClassName Body;

EnumDecl: 'enum' EnumName Inherit? EnumBlock;
Inherit: ':' IdentifierList;
IdentifierList: (identifier | Expression) IdentifierListC*;
IdentifierListC: ',' (identifier | Expression);

EnumBlock: '{' '\n' EnumMember+ '}' '\n';
EnumMember: Declaration | EnumCase;
EnumCase: 'case' EnumCaseName EnumValueAssignment?;
EnumValueAssignment: '=' EnumRawValue '\n';


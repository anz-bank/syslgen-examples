goFile: PackageClause '\n' ImportDecl? '\n' TopLevelDecl+ '\n';
PackageClause: 'package' PackageName ';\n';

ImportDecl: 'import' '(\n' ImportSpec* '\n)\n';
ImportSpec: Import '\n';
TopLevelDecl: Comment '\n' (Declaration | FunctionDecl | MethodDecl);
Declaration: VarDecl | ConstDecl | StructType | InterfaceType | AliasDecl;
StructType : 'type' StructName 'struct' '{\n' FieldDecl* '}\n\n';
FieldDecl: '\t' identifier Type? Tag? '\n';
IdentifierList: identifier IdentifierListC*;
IdentifierListC: ',' identifier;

VarDecl: 'var' IdentifierList TypeName;
ConstDecl: 'const' '(\n'  ConstSpec '\n)\n';
ConstSpec: VarName TypeName '=' ConstValue '\n';

FunctionDecl   : 'func' FunctionName Signature? Block? ;
Signature: Parameters Result?;
Parameters: '(' ParameterList? ')';
Result         : ReturnTypes | TypeName;
ReturnTypes: '(' ResultTypeList* ')';
ResultTypeList: TypeName ',';
TypeList:  TypeName;
ParameterList     : ParameterDecl ParameterDeclC*;
ParameterDecl  : Identifier TypeName;
ParameterDeclC: ',' ParameterDecl;

InterfaceType      : 'type' InterfaceName 'interface'  '{\n'  MethodSpec* '}\n\n' MethodDecl*;
MethodSpec         : '\t' MethodName Signature '\n' | InterfaceTypeName ;
MethodDecl: 'func' Receiver FunctionName Signature? Block? '\n';
Receiver: '(' ReceiverType ')';
AliasDecl: 'type' identifier Type? ';\n\n';

Block: '{\n' StatementList* '}\n\n';
StatementList: Statement ';\n';
Statement: ReturnStmt |  DeclareAndAssignStmt | AssignStmt | IfElseStmt | IncrementVarByStmt | FunctionCall;

AssignStmt: Variables '=' Expression;
IfElseStmt: 'if' Expression Block;
IncrementVarByStmt: Variables '+=' Expression;
ReturnStmt: 'return' (PayLoad | Expression);
DeclareAndAssignStmt: Variables ':=' Expression;

Expression: FunctionCall | NewStruct | GetArg |  ValueExpr | NewSlice;

GetArg: LHS '.' RHS;
NewSlice: '[]' TypeName '{' SliceValues? '}';
FunctionCall: FunctionName '(' FunctionArgs? ')';
FunctionArgs: Expression FuncArgsRest*;
FuncArgsRest: ',' Expression;
NewStruct: StructName '{}';

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
ReturnTypes: '(' ResultTypeList 'error' ')';
ResultTypeList: TypeName ',';
TypeList:  TypeName;
ParameterList     : ParameterDecl ParameterDeclC*;
ParameterDecl  : Identifier TypeName;
ParameterDeclC: ',' ParameterDecl;

InterfaceType      : 'type' InterfaceName 'interface'  '{\n'  MethodSpec* '}\n\n' MethodDecl*;
MethodSpec         : '\t' MethodName Signature '\n' | InterfaceTypeName ;
MethodDecl: 'func' Receiver FunctionName Signature? Block? '\n';
Receiver: '(' ReceiverType ')';

Block: '{\n' StatementList* '}\n\n';
StatementList: Statement ';\n';
Statement: ReturnStmt | FunctionCall | DeclareAndAssignStmt | NewStruct;
ReturnStmt: 'return' (PayLoad | FunctionCall);
FunctionCall: FunctionName '(' FunctionArgs? ')';
DeclareAndAssignStmt: Variables ':=' Statement;
NewStruct: StructName '{}';
AliasDecl: 'type' identifier Type? ';\n\n';


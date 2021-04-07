namespace Graphite.Core

type PrimitiveType = String
type Type = PrimitiveType

type MethodParameterDeclaration = { name: string; ``type``: Type }

//type InternalMethodDeclaration = { name: string; parameters: list<MethodParameterDeclaration>; statements: list<Statement> }
//type InternalModuleDeclaration = { name: string; methods: list<InternalMethodDeclaration> }
//type InternalLibraryDeclaration = { name: string; modules: list<InternalModuleDeclaration> }

type ExternalMethodDeclaration = { name: string; parameters: list<MethodParameterDeclaration> }
type ExternalModuleDeclaration = { name: string; methods: list<ExternalMethodDeclaration> }
type ExternalLibraryDeclaration = { name: string; modules: list<ExternalModuleDeclaration> }

type MethodDeclaration = 
    //| Internal of InternalMethodDeclaration 
    | External of ExternalMethodDeclaration  
    
type ModuleDeclaration = 
    //| Internal of InternalModuleDeclaration 
    | External of ExternalModuleDeclaration      
    
type LibraryDeclaration = 
    //| Internal of InternalLibraryDeclaration 
    | External of ExternalLibraryDeclaration

type StringLiteral = string
type Value = StringLiteral

type Argument = { parameter: MethodParameterDeclaration; value: Value }
type MethodInvocation = { method: MethodDeclaration; arguments: list<Argument> }

type Statement =
    | Invocation of MethodInvocation

type Program = { libraries:list<ExternalLibraryDeclaration>; entrypoint: list<Statement> }
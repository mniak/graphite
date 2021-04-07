namespace Graphite.Core

type PrimitiveType = String
type Type = PrimitiveType

type MethodParameterDeclaration = { name: string; ``type``: Type }

type InternalMethodDeclaration = { name: string; parameters: array<MethodParameterDeclaration>; (*statements: array<IStatement>*) }
type InternalModuleDeclaration = { name: string; methods: array<InternalMethodDeclaration> }
type InternalLibraryDeclaration = { name: string; modules: array<InternalModuleDeclaration> }

type ExternalMethodDeclaration = { name: string; parameters: array<MethodParameterDeclaration> }
type ExternalModuleDeclaration = { name: string; methods: array<ExternalMethodDeclaration> }
type ExternalLibraryDeclaration = { name: string; modules: array<ExternalModuleDeclaration> }

type MethodDeclaration = 
    | Internal of InternalMethodDeclaration 
    | External of ExternalMethodDeclaration  
    
type ModuleDeclaration = 
    | Internal of InternalModuleDeclaration 
    | External of ExternalModuleDeclaration      
    
type LibraryDeclaration = 
    | Internal of InternalLibraryDeclaration 
    | External of ExternalLibraryDeclaration

type StringLiteral = string
type Value = StringLiteral

type Argument = { parameter: MethodParameterDeclaration; value: Value }
type MethodInvocation = { method: MethodDeclaration; arguments: array<Argument> }

type Statement =
    | Invocation of MethodInvocation

type Program = { libraries:array<ExternalLibraryDeclaration>; entrypoint: array<Statement> }
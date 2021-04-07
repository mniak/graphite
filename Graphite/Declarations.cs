namespace Graphite
{
    public record LibraryDeclaration(string name, ModuleDeclaration[] modules) { }
    public record ModuleDeclaration(string name, MethodDeclaration[] methods) { }
    public record MethodDeclaration(string name, MethodParameterDeclaration[] parameters, IStatement[] statements) : IMethodDeclaration { }
    public record MethodParameterDeclaration(string name, Type type) { }

    public record ExternalLibraryDeclaration(string name, ExternalModuleDeclaration[] modules) { }
    public record ExternalModuleDeclaration(string name, ExternalMethodDeclaration[] methods) { }
    public record ExternalMethodDeclaration(string name, MethodParameterDeclaration[] parameters) : IMethodDeclaration { }

    public interface IMethodDeclaration
    {
        string name { get; }
    }

}
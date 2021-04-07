namespace Graphite
{
    public record MethodInvocation(IMethodDeclaration method, Argument[] arguments) : IStatement { }
}
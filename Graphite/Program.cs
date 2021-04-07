namespace Graphite
{
    public record Program(ExternalLibraryDeclaration[] libraries, IStatement[] entrypoint) { }
}
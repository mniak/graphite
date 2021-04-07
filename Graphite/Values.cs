namespace Graphite
{
    public interface IValue { }
    public record StringLiteral(string value) : IValue { }
}
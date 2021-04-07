using System;
using System.Linq;
using System.Text;

namespace Graphite.Serialization.HumanReadable
{
    public class HumanReadableSerializer
    {
        public string SerializeSyntax(Program syntaxGraph)
        {
            return string.Join(Environment.NewLine, syntaxGraph.entrypoint.Select(SerializeStatement));
        }

        public string SerializeStatement(IStatement statement)
        {
            return statement switch
            {
                MethodInvocation mi => SerializeMethodInvocation(mi),
                var s => throw new SourceCodeSerializeException($"Statement could not be serialized: {s}"),
            };
        }
        public string SerializeMethodInvocation(MethodInvocation methodInvocation)
        {
            var sb = new StringBuilder();
            sb.Append(methodInvocation.method.name);
            sb.Append('(');
            foreach (var arg in methodInvocation.arguments)
            {
                sb.Append(arg.parameter.name);
                sb.Append('=');
                sb.Append(SerializeValue(arg.value));
            }
            sb.Append(')');
            return sb.ToString();
        }

        public string SerializeValue(IValue value)
        {
            return value switch
            {
                StringLiteral sl => "\"" + sl.value + "\"",
                var s => throw new SourceCodeSerializeException($"Value could not be serialized: {s}"),
            };
        }
    }
}

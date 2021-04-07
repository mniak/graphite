using System;

namespace Graphite.Serialization.HumanReadable
{
    [Serializable]
    public class SourceCodeSerializeException : Exception
    {
        public SourceCodeSerializeException() { }
        public SourceCodeSerializeException(string message) : base(message) { }
        public SourceCodeSerializeException(string message, Exception inner) : base(message, inner) { }
        protected SourceCodeSerializeException(
          System.Runtime.Serialization.SerializationInfo info,
          System.Runtime.Serialization.StreamingContext context) : base(info, context) { }
    }
}

namespace Graphite
{
    public record Type() { }
    public record PrimitiveType(string name) : Type
    {
        public static PrimitiveType String => new PrimitiveType("string");
    }

    public enum PrimitiveTypes
    {
        String,

        //// Integers
        //UInt8,
        //SInt8,
        //UInt16,
        //SInt16,
        //UInt32,
        //SInt32,
    }
}
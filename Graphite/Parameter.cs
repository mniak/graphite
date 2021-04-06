namespace Graphite
{
    public class Parameter
    {
        private PrimitiveTypes type;
        private string name;

        public Parameter(PrimitiveTypes type, string name)
        {
            this.type = type;
            this.name = name;
        }
    }
}
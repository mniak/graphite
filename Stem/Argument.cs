namespace Stem
{
    public class Argument
    {
        private Parameter parameter;
        private StringLiteral value;

        public Argument(Parameter parameter, StringLiteral value)
        {
            this.parameter = parameter;
            this.value = value;
        }
    }
}
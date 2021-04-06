namespace Stem
{
    public class Method
    {
        private string v;
        private Parameter[] parameters;

        public Method(string v, Parameter[] parameters)
        {
            this.v = v;
            this.parameters = parameters;
        }
    }
}
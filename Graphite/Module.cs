namespace Graphite
{
    public class Module
    {
        private string v;
        private Method[] methods;

        public Module(string v, Method[] methods)
        {
            this.v = v;
            this.methods = methods;
        }
    }
}
namespace Graphite
{
    public class Library
    {
        private string name;
        private Module[] modules;

        public Library(string name, Module[] modules)
        {
            this.name = name;
            this.modules = modules;
        }
    }
}
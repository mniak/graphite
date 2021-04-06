namespace Graphite
{
    public class SyntaxGraph
    {
        private Library[] libraries;
        private Block entrypoint;

        public SyntaxGraph(Library[] libraries, Block entrypoint)
        {
            this.libraries = libraries;
            this.entrypoint = entrypoint;
        }
    }
}
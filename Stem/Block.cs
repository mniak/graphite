namespace Stem
{
    public class Block
    {
        private MethodInvocation[] methodInvocations;

        public Block(MethodInvocation[] methodInvocations)
        {
            this.methodInvocations = methodInvocations;
        }
    }
}
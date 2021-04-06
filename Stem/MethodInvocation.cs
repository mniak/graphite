namespace Stem
{
    public class MethodInvocation
    {
        private Method method;
        private Argument[] argumentList;

        public MethodInvocation(Method method, Argument[] argumentList)
        {
            this.method = method;
            this.argumentList = argumentList;
        }
    }
}
using Graphite;

namespace Demo
{
    class Program
    {
        static void Main(string[] args)
        {
            HelloWorld();
        }

        static void HelloWorld()
        {
            /*
                #include <stdio.h>
                void main() {
                    printf("Hello, World!");
                }
            */
            var printfParam0 = new Parameter(
                type: PrimitiveTypes.String,
                name: "text"
            );
            var printf = new Method("printf", new[] { printfParam0 });
            var standardLibrary = new Library(
                name: "stdlib",
                modules: new[] { new Module("io", new[] { printf }) }
            );
            var syntax = new SyntaxGraph(
                libraries: new[] { standardLibrary },
                entrypoint: new Block(new[] {
                    new MethodInvocation(
                        method: printf,
                        argumentList: new[] {
                            new Argument(
                                parameter: printfParam0,
                                value: new StringLiteral("Hello, World!")
                            )
                        }
                    )
                })
            );
        }
    }
}

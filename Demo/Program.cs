using Graphite.Core;
using System;

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
            var printfParam0 = new MethodParameterDeclaration(
                type: PrimitiveType.String,
                name: "text"
            );
            var printf = new ExternalMethodDeclaration("printf", new[] { printfParam0 });
            var standardLibrary = new ExternalLibraryDeclaration(
                name: "stdlib",
                modules: new[] { new ExternalModuleDeclaration("io", new[] { printf }) }
            );
            var syntax = new Graphite.Core.Program(
                libraries: new[] { standardLibrary },
                entrypoint: new Statement[] {
                    Statement.NewMethod(new MethodInvocation(
                        method: MethodDeclaration.NewExternal(printf),
                        arguments: new[] {
                            new Argument(
                                parameter: printfParam0,
                                value: "Hello, World!"
                            )
                        }
                    )),
                }
            );


            //var renderer = new HumanReadableSerializer();
            //var csharpCode = renderer.SerializeSyntax(syntax);

            //Console.WriteLine(csharpCode);
            Console.ReadLine();
        }
    }
}

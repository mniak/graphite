#include "graphite.h"
#include "program.h"
#include <iostream>

using namespace std;

int main()
{
    /*
        #include <stdio.h>
        void main() {
            printf("Hello, World!");
        }
    */
    MethodParameterDeclaration printfParam0;
    printfParam0.type = PrimitiveType::String();
    printfParam0.name = "text";

    ExternalMethodDeclaration printf;
    printf.name = "printf";
    printf.parameters.push_back(printfParam0);

    ExternalModuleDeclaration externalStdlib;
    externalStdlib.name = "io";
    externalStdlib.methods.push_back(printf);

    ExternalLibraryDeclaration standardLibrary;
    standardLibrary.name = "stdlib";
    standardLibrary.modules.push_back(externalStdlib);

    MethodInvocation printfInvocation;
    printfInvocation.method = printf;
    printfInvocation.arguments.push_back(
        Argument(
            printfParam0,
            StringLiteral("Hello, World!")));

    Program syntax;
    syntax.libraries.push_back(standardLibrary);
    syntax.entrypoint.push_back(printfInvocation);

    

    // auto renderer = new HumanReadableSerializer();
    // auto csharpCode = renderer.SerializeSyntax(syntax);

    // cout << csharpCode << end;
    return 0;
}
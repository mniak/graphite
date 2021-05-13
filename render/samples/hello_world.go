package samples

import "github.com/mniak/graphite"

func HelloWorld() graphite.Program {
	/*
	   #include <stdio.h>
	   void main{} {
	       printf{"Hello, World!"};
	   }
	*/
	//printfParam0 := graphite.MethodParameterDeclaration{
	//	Name:    "text",
	//	TheType: graphite.PrimitiveString(),
	//}
	//printf := graphite.ExternalMethodDeclaration{
	//	Name: "printf",
	//	Parameters: []graphite.MethodParameterDeclaration{
	//		printfParam0,
	//	},
	//}
	//standardLibrary := graphite.ExternalLibraryDeclaration{
	//	Name: "native]",
	//	Modules: []graphite.ExternalModuleDeclaration{
	//		{
	//			Name: "io",
	//			Methods: []graphite.ExternalMethodDeclaration{
	//				printf,
	//			},
	//		},
	//	},
	//}
	//
	//program := graphite.Program{
	//	Libraries: []graphite.ExternalLibraryDeclaration{standardLibrary},
	//	Entrypoint: []graphite.IStatement{
	//		graphite.MethodInvocation{
	//			Method: printf,
	//			Arguments: []graphite.Argument{
	//				{
	//					Parameter: printfParam0,
	//					Value:     graphite.StringLiteral{Value: "Hello, World!"},
	//				},
	//			},
	//		},
	//	},
	//}
	//return program
	return nil
}

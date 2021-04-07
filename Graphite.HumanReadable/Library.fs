module Graphite.HumanReadable

open Graphite.Core
open System

let SerializeValue value = "\"" + value + "\""
let SerializeStatement statement =
    match statement with
    | Invocation inv ->
        match inv.method with
        //| MethodDeclaration.Internal i -> i.name
        | MethodDeclaration.External e -> e.name
        + "("
        + (inv.arguments 
            |> Seq.map (fun arg -> 
                arg.parameter.name
                + "="
                + SerializeValue(arg.value)) 
            |> String.concat ", "
            )
        + ")"

let SerializeSyntax syntax = 
    syntax.entrypoint
    |> Seq.map SerializeStatement
    |> String.concat Environment.NewLine
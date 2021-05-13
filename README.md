![License](https://img.shields.io/github/license/mniak/graphite.svg?style=flat-square)
![Tag](https://img.shields.io/github/tag/mniak/graphite.svg?style=flat-square)

# Graphite: Languageless compiler

This is an experiment trying to create a compiler that skips all the syntactic parsing.

The proposal is that the code is already born in the editor as an [abstract semantic graph](https://en.wikipedia.org/wiki/Abstract_semantic_graph).

I intend to use the [LLVM infrastructure](https://llvm.org/) to create the compiler.

## Roadmap

I'm starting with the part of starting to draw the semantic graph via code and then continue with the following features:

- [ ] Serializer in _human readable_ format with configurable syntax.
- [ ] Serializer and deserializer in _machine readable_ format so that the "code" can be stored and read.
  It is interesting that this is in some format that is still slightly _human readable_ to facilitate the differential comparison and therefore better integration with git. Maybe a YAML.
- [ ] [Structured Editor] (https://en.wikipedia.org/wiki/Structure_editor) to directly edit the semantic graph
- [ ] "Precompiler" to [LLVM IR](https://en.wikipedia.org/wiki/LLVM#Intermediate_representation)
  - [x] Entry point
  - [ ] Basic arithmetic operations
    - [x] Addition `+`
    - [x] Multiplication `*`
  - [x] Method declaration
  - [x] Method invocation
  - [ ] Declare external method (in other library)
- [ ] Integration with LLVM so that I can use a single command to perform the complete compilation from the "source" (graph) to the executable binary. 
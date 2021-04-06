# Graphite: Code without syntax

Este é um experimento tentando criar um compilador que pula toda a parte de parseamento sintático.

A proposta é que o código já nasca no editor como um [grafo semântico abstrato](https://en.wikipedia.org/wiki/Abstract_semantic_graph).

Pretendo usar a [infraestrutura LLVM](https://llvm.org/) para criar o compilador.

Estou iniciando pela parte de começar a desenhar o grafo semântico via código para depois seguir com as seguintes funcionalidades:
- Serializador em formato _human readable_ de sintaxe configurável.
- Serializador e desserializador em formato _machine readable_ para que o "código" possa ser armazenado e lido. É interessante que este seja em algum formato que ainda seja ligeiramente _human readable_ para facilitar a comparação diferencial e por conseguinte, melhor integração com o git. Talvez um YAML.
- [Editor Estruturado](https://en.wikipedia.org/wiki/Structure_editor) para editar diretamente o grafo semântico
- "Precompilador" para representação [LLVM IR](https://en.wikipedia.org/wiki/LLVM#Intermediate_representation)
- Integração com o LLVM de maneira que eu possa através de um só comando realizar a compilação completa desde o "fonte" (grafo) até o binário executável.
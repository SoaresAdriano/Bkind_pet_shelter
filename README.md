# Bkind Pet Shelter

O projeto Bkind chegou como um desafio do meu Tech lead Luís, da Dito. Em uma conversa de 1:1 surgiu a ideia onde a partir de um CRUD inicial eu fosse capaz de dar meus primeiros passos na criação de uma aplicação que realize os conceitos básicos de um CRUD. Para esse propósito então decidi iniciar a criação de um sistema capaz de realizar operações de um abrigo para animais. A documentação do projeto inteiro ainda será abordada.

## Fluxo de dados na aplicação  

`petCadaster_service`

- Pessoas com acesso a API podem realizar operações básicas de CRUD.

## Como executar
Instale e configure [Go](https://golang.org/doc/install).

Copie o conteúdo do arquivo `.env.example` para um novo arquivo chamado `.env`. O arquivo `.env.example` contém todas as variáveis necessárias para executar o projeto.

Execute o comando `make rest-api`.


## Como testar


## Comandos Makefile

`make setup`
 
Baixa os binários de Go necessários para rodar o conjunto de testes e pacotes necessários para rodar o serviço.
 
`make rest-api`

Inicializa o serviço para comunicação via REST API.
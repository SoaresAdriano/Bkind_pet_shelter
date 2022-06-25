# Bkind Pet Shelter

O projeto Bkind chegou como um desafio do meu Tech lead Luís, da Dito. Em uma conversa de 1:1 surgiu a ideia onde a partir de um CRUD inicial eu fosse capaz de dar meus primeiros passos na criação de uma aplicação que realize os conceitos básicos de um CRUD. Para esse propósito então decidi iniciar a criação de um sistema capaz de realizar operações de um abrigo para animais. A documentação do projeto inteiro ainda será abordada.


## Ubiquitous Language

`Adoption`

A process in which a rescued pet is placed into a permanent home.

`Adoption Guarantee`

Operational philosophy of sheltering 
organizations. This language is often used interchangeably with “no kill” with all of the same variations 
in meaning. For example, organizations may variably use this term to indicate that all animals entering a 
sheltering program will remain there until adopted, only medically or behaviorally healthy or treatable 
animals will be cared for until adoption, or that only animals for which an adoption can be guaranteed 
are admitted into the organization’s care. Actual admission, adoption, or euthanasia practices of an 
organization cannot be inferred by the use of this terminology.

`Castration`
The removal of the testicles of a male animal.

`Capacity for Care`

Is an organization’s ability to appropriately care for the animals it serves. This is based
on a range of parameters including, but not limited to, the number of appropriate housing units; staffing 
for programs or services; staff training; average length of stay; and the total number of reclaims, 
adoptions, transfers, releases, or other outcomes.  

`Caregiver`

A person overseeing and providing food and veterinary care for a feral cat colony.

`Closed Adoption`

Adoption policies with strict rules, guidelines, or background checks.

`Cruelty`

Physical harm or injury inflicted upon an animal by an individual or group of individuals.

`Euthanasia`

Ending the life of a pet.

`Feral`

An animal who is not social with humans and shows little to no signs of being interested in becoming social with humans.

`Neuter`

Castrate or spay (a domestic animal). 

`Limited Admission`

Shelters are typically privately‐funded shelters without municipal roles or animal 
control contracts.  As a private organization, the shelter may accept animals based on self‐defined 
criteria and mission.  Because they are not publicly funded, limited admission shelters are under no 
obligation to take stray animals, and often focus on services for owner‐surrendered animals.

`Live Release Rate (LRR)`

Is an indication of the number of animals leaving a facility by means other than 
euthanasia or in‐shelter death. Live outcomes are usually achieved through adoption, reclaim by owner, 
transfer to another agency or other life‐saving actions. LRR is usually expressed by a percentage that can 
be calculated in three ways, depending on organizational preferences.

`Pet Shelter`

A place to treat abandoned, injured or lost animals and provide the care they need until they get a new place to be.

`Spay`

The surgical removal of a female pet's uterus and ovaries to prevent reproduction.

More terms can be seeing here:

[sheltervet.org](https://www.sheltervet.org/assets/PDFs/shelter%20terminology.pdf)

[coastalpetrescue.org](https://coastalpetrescue.org/humane-education/glossary/)

### Fluxo de dados na aplicação  

`petCadaster_service`

- Pessoas com acesso a API podem realizar operações básicas de CRUD.

### Como executar
Instale e configure [Go](https://golang.org/doc/install).

Copie o conteúdo do arquivo `.env.example` para um novo arquivo chamado `.env`. O arquivo `.env.example` contém todas as variáveis necessárias para executar o projeto.

Execute o comando `make rest-api`.


### Como testar


### Comandos Makefile

`make setup`
 
Baixa os binários de Go necessários para rodar o conjunto de testes e pacotes necessários para rodar o serviço.
 
`make rest-api`

Inicializa o serviço para comunicação via REST API.
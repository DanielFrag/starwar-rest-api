# starwars-rest-api (B2W)
API escrita em Golang para fornecer recursos de acordo com os requisitos abaixo.

# Requisitos

- A API deve ser REST;
- Para cada planeta, os seguintes dados devem ser obtidos do banco de dados da aplicação, sendo inserido manualmente:
```
Nome
Clima
Terreno
```
- Para cada planeta também devemos ter a quantidade de aparições em filmes, que podem ser obtidas pela API pública do Star Wars:  https://swapi.co/

### Funcionalidades desejadas: 

- Adicionar um planeta (com nome, clima e terreno);
- Listar planetas;
- Buscar por nome;
- Buscar por ID;
- Remover planeta.

# Descrição da Solução

A aplicação foi organizada nos seguintes pacotes:
- dto - contém a definição da estrutura de dados usada para tráfego e comunicação com a aplicação de terceiros.
- handler - contém a lógica de manipulação de dados e disponibilização dos recursos.
- infra - contém o código de acesso e manipulação da base de dados (MongoDB).
- mock - contém dados "mocados" usados nos testes.
- model - contém a definição das estruturas de dados a serem usadas na aplicação (objetos do domínio).
- repository - contém a definição das interfaces usadas pela camada de negócio para acesso a base de dados (interna e externa).
- router - contém a configuração das formas de disponibilização dos recursos.
- service - contém as entidades que realizam a interação com a aplicação de terceiros.
- utils - contém as ferramentas utilizadas em diversas partes da aplicação que não fazem parte da lógica do domínio.

## Testes

Para rodar os testes dos pacotes, basta navegar até o diretório que contém o pacote e digitar o comando:
```
go test
```
Vale ressaltar que alguns testes fazem acesso tanto aos dados do pacote mock quanto acessos a dados reais externos à aplicação.

Para rodar os testes do pacote de 'infra', são necessárias algumas observações.
Seus testes fazem acesso real ao Mongo por isso, o mesmo precisa estar rodando.
Os testes desse pacote incluem e manipulam dados em uma base de teste que deve ser excluída quando o teste for concluído.

Para rodar um teste específico, basta navegar até o diretório que contém o teste e digitar

```
go test -run <nome_do_teste>
//ex: go test -run TestCorsSetup (na pasta handler)
```

## Observações quanto aos Middlewares

Para isolar responsabilidades e organizar a hierarquia de execução das funções que atendem aos recursos solicitados, foi utilizada uma abordagem das funções semelhante a organização de middlewares adotada pelo framework Express.js do Node.js.
Esta organização impõe que todo middleware receba como parâmetro um adapter 'http.HandlerFunc' e que retorne um adapter do mesmo tipo.
Assim, as funções são organizadas e injetadas pelas instâncias da struct do tipo 'HandlerFuncInjector'.

## Variáveis de ambiente

As seguintes variáveis de ambiente possuem valores padrões que podem ser redefinidos para esta aplicação:
- MONGODB_URI contém a url de conexão com o banco (mongodb://localhost:27017).
- DB_NAME contém o nome da base de dados (widgets-spa-rv).
- PORT contém a porta a ser escutada pelo servidor (8080).
- TIMEOUT_SECONDS contém o tempo de expiração de uma requisição em segundos (30).

## Bibliotecas de terceiros

As seguintes bibliotecas foram usadas neste projeto:
- github.com/gorilla/context
- github.com/gorilla/mux
- gopkg.in/mgo.v2
- gopkg.in/mgo.v2/bson

## Rodar a API

Após instalar as dependências listadas acima, basta rodar no terminal o comando:

```
go build
```

Um novo executável será criado na raiz do projeto. Então, basta digitar no terminal o comando:

```
./<nome_do_executável>
```

## Teste das rotas

O arquivo 'starwarsapi.postman_collection.json' pode ser importado pelo software 'Postman'.
Ao importar o arquivo, será criada uma nova coleção no Postman. Esta coleção contém os dados para requisitar os recursos da API, além de alguns resultados já gravados como exemplo.

## Rotas

### GET /api/planets

Retorna os dados dos planetas presentes na base de dados. Pode receber como parâmetro na url as propriedades 'id' ou 'name' do planeta, retornando 0 ou 1 planeta. Caso um identificador não seja passado na url, retorna todos os planetas da base de dados.
#### Exemplos
```
url: http://<domínio>/api/planets
url: http://<domínio>/api/planets?id=5aa8134cabb4442ef287e0eb
url: http://<domínio>/api/planets?name=Tatooine
```
### POST /api/planets

Insere um planeta na base de dados. Recebe como parâmetro, no corpo da requisição, um json contendo as propriedades 'name', 'climate', 'terrain' do planeta. Retorna um status code 201 e no cabeçalho da resposta o endereço do recurso criado (Location) ou os seguintes erros:
- caso um recurso com um mesmo 'name' já exista (500).
- caso o content-type da requisição não seja 'application/json' (415).
- caso os dados do planeta estejam incompletos (faltando ou 'name', ou 'climate', ou 'terrain') (400).
#### Exemplos
```
url: http://<domínio>/api/planets

body:
{
	"name": "Tatooine",
	"climate": "hot",
	"terrain": "sand"
}
```
### DELETE /api/planets/:id

Remove um planeta da base de dados caso ele exista. Retorna o código 204 caso o planeta tenha sido removido ou 404 caso não o encontre.
#### Exemplos
```
url: http://<domínio>/api/planets/5aa8134cabb4442ef287e0eb
```

## Integração externa

A quantidade de filmes em que um planeta aparece não é armazenada na base de dados da api, ela é requisitada à uma aplicação de terceiros (SWAPI), hospedada no endereço
```
https://www.swapi.co/
```
Todas as vezes que um planeta é solicitado à aplicação, ela requisita os dados deste planeta à SWAPI e contabiliza a quantidade de filmes em que este planeta aparece. Caso os dados não sejam encontrados, o valor 0 é atribuído à quantidade de aparições do planeta em filmes.   
Assim, caso os dados de um planeta sejam atualizados, a aplicação se manterá coerente (single source of truth).   
Ao solictar todos os planetas da aplicação, dois algorítmos podem ser usados para popular os dados de quantidade de aparições dos planetas em filmes.   
A escolha entre estes dois algorítmos depende da quantidade de registros a serem populados. A escolha entre estes dois algoritmos visa diminuir o número de requisições a dados externos.   
Em ambos os algoritmos a complexidade é próxima à quantidade de planetas (O(n)). Porém, no primeiro, a quantidade de requisições externas cresce de acordo com a quantidade de registros, e no segundo a quantidade de requisições externas possui um limite.
### 1º algoritmo
Caso existam poucos registros (menos de 9), os dados dos planetas são buscados individualmente, pelo nome ou pelo id do planeta, na SWAPI. Essas requisições são gerenciadas por um 'worker' capaz de processar até 3 solicitações em paralelo.
### 2º algoritmo
Caso a quantidade de registros seja numerosa, cria-se um dicinário (hashmap), cuja chave é a propriedade 'name' do registro e o valor é o próprio registro. Requisita-se então todos os planetas à SWAPI e varre-se a coleção retornada buscando o nome do planeta no dicinário. Caso o registro seja encotrado, a quantidade de filmes é populada.

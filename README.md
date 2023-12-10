# Platform-API Gateway

A Platform-API Gateway é uma API intermediária que atua como um ponto de entrada para as APIs internas, fornecendo roteamento, autenticação e outros serviços comuns de middleware.

## Funcionalidades

- **Roteamento de solicitações**: Direciona as chamadas de entrada para os serviços internos corretos.
- **Autenticação**: Verifica as credenciais das solicitações e rejeita aquelas que não atendem aos requisitos de autenticação.
- **Logging**: Mantém um registro de todas as solicitações e respostas para fins de monitoramento e depuração.

## Configuração

Antes de iniciar a API Gateway, você deve configurar as seguintes variáveis de ambiente:

- `DATABASE_URL`: URL de conexão para o banco de dados PostgreSQL.
- `API_KEY`: Chave de API para autenticação nas rotas protegidas.

Instale as dependências do projeto com o seguinte comando:

```bash
go mod tidy
```

## Iniciando a API Gateway
Para executar a API Gateway, use o seguinte comando:
```bash
go run ./cmd/server/main.go
```

## Endpoints
Endpoint cadastrado na [API](https://github.com/jcr04/AUAUPETS.go) como forma de exemplo, em teste real adicione sua propria api com as configurações necessarias

### Rotas para Animais
* GET `/animals`: Lista todos os registros de animais.


## Autenticação

Todas as rotas protegidas exigem um cabeçalho de autenticação na solicitação. A autenticação é realizada via token JWT. Inclua o seguinte cabeçalho em suas solicitações para rotas protegidas:
* Authorization: `Bearer YOUR_JWT_TOKEN`


## Erros

A API usa o seguinte formato de erro:

```json
{
  "success": false,
  "error": {
    "code": "código_de_erro",
    "message": "Mensagem de erro descritiva."
  }
}
```
## Rate Limiting
Para garantir a qualidade do serviço, a API impõe limites de taxa de solicitação. Se você exceder esses limites, receberá uma resposta 429 Too Many Requests.

## Exemplos de Solicitações
* - ![Screenshot_2](https://github.com/jcr04/Plataform-API/assets/70778525/f3240772-0697-45fe-94d1-d032101058c6)


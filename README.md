# Enviar Mensagem usando Golang, RabbitMQ e MySQL

Aplicação Web que permite cadastrar mensagem, esta aplicação utiliza os serviços  [Message API](https://github.com/marceloagmelo/go-message-api). Esta aplicação possue as seguintes funcionalidades.

- [Listar Mensagens](#listar-mensagens)
- [Cadastrar Mensagem](#enviar-mensagem)
- [Excluir Mensagem](#atualizar-mensagem)

----

# Instalação

```
go get -v github.com/marceloagmelo/go-message-send
```
```
cd go-message-send
```

## Build da Aplicação

```
./image-build.sh
```

## Iniciar as Aplicações de Dependências
```
./dependecy-start.sh
```

## Preparar o MySQL

```
docker  exec -it mysqldb bash -c "mysql -u root -p"
```
- Criar a tabela
	> use gomessagedb;
	
	> CREATE TABLE mensagem (
id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
titulo VARCHAR(100), texto VARCHAR(255),
status INTEGER,
PRIMARY KEY (id)
);

## Iniciar a Aplicação
```
./start.sh
```
```
http://localhost:7070
```

## Finalizar a Aplicação
```
./stop.sh
```

## Finalizar a Todas as Aplicações
```
./stop-all.sh
```

# Fucionalidades
Lista das funcionalidas:

### Listar Mensagens
[Clique aqui para ver a imagem](https://github.com/marceloagmelo/go-message-send/blob/master/imagens/tela-listar-mensagens.png?raw=true)

### Cadastrar Mensagem
[Clique aqui para ver a imagem](https://github.com/marceloagmelo/go-message-send/blob/master/imagens/tela-cadastrar-mensagem.png?raw=true)


### Apagar Mensagem
[Clique aqui para ver a imagem](https://github.com/marceloagmelo/go-message-send/blob/master/imagens/tela-listar-mensagens.png?raw=true)
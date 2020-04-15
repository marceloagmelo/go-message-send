package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/marceloagmelo/go-message-send/logger"
	"github.com/marceloagmelo/go-message-send/models"
	"github.com/marceloagmelo/go-message-send/utils"
	"github.com/marceloagmelo/go-message-send/variaveis"
)

var api = "go-message/api/v1"

//Health testar conexão com a API
func Health() (mensagemHealth models.MensagemHealth, erro error) {
	endpoint := variaveis.ApiURL + "/" + api + "/health"

	resposta, err := utils.GetRequest(endpoint)
	defer resposta.Body.Close()
	if err != nil {
		return mensagemHealth, err
	}
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o conteudo recebido", err.Error())
			logger.Erro.Println(mensagem)
			return mensagemHealth, err
		}
		mensagemHealth = models.MensagemHealth{}
		err = json.Unmarshal(corpo, &mensagemHealth)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter o retorno JSON", err.Error())
			logger.Erro.Println(mensagem)
			return mensagemHealth, err
		}
	}
	return mensagemHealth, nil
}

//ListaMensagens listar mensagens
func ListaMensagens() (mensagens models.Mensagens, erro error) {
	endpoint := variaveis.ApiURL + "/" + api + "/mensagens"

	resposta, err := utils.GetRequest(endpoint)
	if err != nil {
		return nil, err
	}
	defer resposta.Body.Close()
	if resposta.StatusCode == 200 {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o conteudo recebido", err.Error())
			logger.Erro.Println(mensagem)
			return nil, err
		}
		mensagens = models.Mensagens{}
		err = json.Unmarshal(corpo, &mensagens)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter o retorno JSON", err.Error())
			logger.Erro.Println(mensagem)
			return nil, err
		}
	}
	return mensagens, nil
}

//EnviarMensagem enviar a mensagem
func EnviarMensagem(novaMensagem models.Mensagem) (mensagemRetorno models.Mensagem, erro error) {
	endpoint := variaveis.ApiURL + "/" + api + "/mensagem/criar"

	resposta, err := utils.PostRequest(endpoint, novaMensagem)
	defer resposta.Body.Close()
	if err != nil {
		return mensagemRetorno, err
	}
	if resposta.StatusCode == http.StatusCreated {
		corpo, err := ioutil.ReadAll(resposta.Body)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao ler o conteudo recebido", err.Error())
			logger.Erro.Println(mensagem)
			return mensagemRetorno, err
		}
		mensagemRetorno = models.Mensagem{}
		err = json.Unmarshal(corpo, &mensagemRetorno)
		if err != nil {
			mensagem := fmt.Sprintf("%s: %s", "Erro ao converter o retorno JSON", err.Error())
			logger.Erro.Println(mensagem)
			return mensagemRetorno, err
		}
	}
	return mensagemRetorno, nil
}

//ApagarMensagem apagar mensagem
func ApagarMensagem(id string) error {
	endpoint := variaveis.ApiURL + "/" + api + "/mensagem/apagar/" + id

	err := utils.DeleteRequest(endpoint)
	if err != nil {
		return err
	}
	return nil
}

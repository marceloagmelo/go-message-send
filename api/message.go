package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/marceloagmelo/go-message-send/logger"
	"github.com/marceloagmelo/go-message-send/models"
	"github.com/marceloagmelo/go-message-send/variaveis"
)

var api = "go-message/api/v1"

// getRequest recuperar a requisição
func getRequest(endpoint string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	defer tr.CloseIdleConnections()

	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 180,
	}

	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao criar um request", err.Error())
		logger.Erro.Println(mensagem)
		return nil, err
	}

	resposta, err := cliente.Do(request)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao abrir o request", err.Error())
		logger.Erro.Println(mensagem)
		return nil, err
	}
	return resposta, nil
}

// postRequest envio de uma requisição
func postRequest(endpoint string, mensagem models.Mensagem) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	defer tr.CloseIdleConnections()

	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 30,
	}

	conteudoEnviar, err := json.Marshal(&mensagem)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao gerar o objeto com o JSON lido", err.Error())
		logger.Erro.Println(mensagem)
		return nil, err
	}

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(conteudoEnviar))
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao criar o request com a mensagem", err.Error())
		logger.Erro.Println(mensagem)
		return nil, err
	}

	resposta, err := cliente.Do(request)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao executar o post da mensagem", err.Error())
		logger.Erro.Println(mensagem)
		return nil, err
	}
	return resposta, nil
}

//Health testar conexão com a API
func Health() (mensagemHealth models.MensagemHealth, erro error) {
	endpoint := variaveis.ApiURL + "/" + api + "/health"

	resposta, err := getRequest(endpoint)
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

	resposta, err := getRequest(endpoint)
	defer resposta.Body.Close()
	if err != nil {
		return nil, err
	}
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

	resposta, err := postRequest(endpoint, novaMensagem)
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

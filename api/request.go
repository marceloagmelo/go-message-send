package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/marceloagmelo/go-message-send/logger"
	"github.com/marceloagmelo/go-message-send/models"
)

// GetRequest recuperar a requisição
func GetRequest(endpoint string) (*http.Response, error) {
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

// PostRequest envio de uma requisição
func PostRequest(endpoint string, mensagem models.Mensagem) (*http.Response, error) {
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

// DeleteRequest requisição para deletar
func DeleteRequest(endpoint string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	defer tr.CloseIdleConnections()

	cliente := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 180,
	}

	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao criar um request", err.Error())
		logger.Erro.Println(mensagem)
		return err
	}

	resposta, err := cliente.Do(request)
	if err != nil {
		mensagem := fmt.Sprintf("%s: %s", "Erro ao abrir o request", err.Error())
		logger.Erro.Println(mensagem)
		return err
	}
	defer resposta.Body.Close()

	return nil
}

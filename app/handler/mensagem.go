package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marceloagmelo/go-message-send/api"
	"github.com/marceloagmelo/go-message-send/logger"
	"github.com/marceloagmelo/go-message-send/models"
)

var view = template.Must(template.ParseGlob("views/*.html"))

//Health testa conexão com o mysql e rabbitmq
func Health(w http.ResponseWriter, r *http.Request) {
	mensagem, err := api.Health()
	if err != nil {
		mensagemErro := fmt.Sprintf("%s: %s", "Erro verificar o heal check", err)
		data := map[string]interface{}{
			"titulo":       "Lista de Mensagens",
			"mensagemErro": mensagemErro,
		}

		err := view.ExecuteTemplate(w, "Erro", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	data := map[string]interface{}{
		"titulo":   "Lista de Mensagens",
		"mensagem": mensagem.Mensagem,
	}

	err = view.ExecuteTemplate(w, "Health", data)
	if err != nil {
		mensagemErro := fmt.Sprintf("%s: %s", "Erro ao chamar a página de health check", err)
		data := map[string]interface{}{
			"titulo":       "Lista de Mensagens",
			"mensagemErro": mensagemErro,
		}

		err := view.ExecuteTemplate(w, "Erro", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
}

//Home primeira página
func Home(w http.ResponseWriter, r *http.Request) {
	mensagens, _ := api.ListaMensagens()

	data := map[string]interface{}{
		"titulo":    "Lista de Mensagens",
		"mensagens": mensagens,
	}

	err := view.ExecuteTemplate(w, "Index", data)
	if err != nil {
		mensagemErro := fmt.Sprintf("%s: %s", "Erro ao chamar a página home", err)
		data := map[string]interface{}{
			"titulo":       "Lista de Mensagens",
			"mensagemErro": mensagemErro,
		}

		err := view.ExecuteTemplate(w, "Erro", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
}

//New página de edição de uma nova mensagem
func New(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"titulo":   "Nova Mensagem",
		"mensagem": "",
	}

	view.ExecuteTemplate(w, "New", data)
}

//Insert mensagem
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			mensagemErro := fmt.Sprintf("%s: %s", "Erro no parse do formulário", err)
			data := map[string]interface{}{
				"titulo":       "Lista de Mensagens",
				"mensagemErro": mensagemErro,
			}

			err := view.ExecuteTemplate(w, "Erro", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		titulo := r.FormValue("titulo")
		texto := r.FormValue("texto")

		if titulo != "" && texto != "" {
			var mensagemForm models.Mensagem
			mensagemForm.ID = 0
			mensagemForm.Titulo = titulo
			mensagemForm.Texto = texto
			mensagemForm.Status = 1

			mensagemRetorno, err := api.EnviarMensagem(mensagemForm)
			if err != nil {
				mensagemErro := fmt.Sprintf("%s: %s", "Erro ao listar todas as mensagens", err)
				data := map[string]interface{}{
					"titulo":       "Lista de Mensagens",
					"mensagemErro": mensagemErro,
				}

				err = view.ExecuteTemplate(w, "Erro", data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}

			mensagem := fmt.Sprintf("Mensagem %v enviada com sucesso!", mensagemRetorno.ID)
			logger.Info.Println(mensagem)

			data := map[string]interface{}{
				"titulo":   "Lista de Mensagens",
				"mensagem": mensagem,
			}

			err = view.ExecuteTemplate(w, "Sucesso", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

//Apagar mensagem
func Apagar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := api.ApagarMensagem(id)
	if err != nil {
		mensagemErro := fmt.Sprintf("%s: %s", "Erro no parse do formulário", err)
		data := map[string]interface{}{
			"titulo":       "Lista de Mensagens",
			"mensagemErro": mensagemErro,
		}

		err := view.ExecuteTemplate(w, "Erro", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	mensagem := fmt.Sprintf("Mensagem %v apagada com sucesso!", id)
	logger.Info.Println(mensagem)

	http.Redirect(w, r, "/", 301)
}

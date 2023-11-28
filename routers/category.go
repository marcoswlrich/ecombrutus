package routers

import (
	"encoding/json"
	"strconv"

	"github.com/marcoswlrich/ecombrutus/bd"
	"github.com/marcoswlrich/ecombrutus/models"
)

func InsertCategory(body string, User string) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Erro nos dados recebidos " + err.Error()
	}

	if len(t.CategName) == 0 {
		return 400, "Deve-se especificar o Nome (Title) da categoria"
	}

	if len(t.CategPath) == 0 {
		return 400, "Deve-se especificar o caminho (Rota) da categoria"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, err2 := bd.InsertCategory(t)
	if err2 != nil {
		return 400, "Ocorreu um erro ao tentar realizar o registro da categoria" + t.CategName + " > " + err2.Error()
	}

	return 200, "{ CategID: " + strconv.Itoa(int(result)) + "}"
}

func UpdateCategory(body string, User string, id int) (int, string) {
	var t models.Category

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Erro nos dados recebidos " + err.Error()
	}

	if len(t.CategName) == 0 && len(t.CategPath) == 0 {
		return 400, "Deve-se especificar CategName e CategPath para atualizar"
	}

	isAdmin, msg := bd.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	t.CategID = id
	err2 := bd.UpdateCategory(t)

	if err2 != nil {
		return 400, "Ocorreu um erro ao tentar realizar o UPDATE da categoría " + strconv.Itoa(
			id,
		) + " > " + err.Error()
	}

	return 200, "Update OK"
}

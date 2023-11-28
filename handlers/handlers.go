package handlers

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"

	"github.com/marcoswlrich/ecombrutus/auth"
	"github.com/marcoswlrich/ecombrutus/routers"
)

func Handlers(
	path string,
	method string,
	body string,
	headers map[string]string,
	request events.APIGatewayV2HTTPRequest,
) (int, string) {
	fmt.Println("Processando" + path + " > " + method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := validateAuth(path, method, headers)
	if !isOk {
		return statusCode, user
	}

	switch path[0:4] {
	case "user":
		return ProcessUsers(body, path, method, user, id, request)
	case "prod":
		return ProcessProducts(body, path, method, user, idn, request)
	case "stoc":
		return ProcessStock(body, path, method, user, idn, request)
	case "addr":
		return ProcessAddress(body, path, method, user, idn, request)
	case "cate":
		return ProcessCategory(body, path, method, user, idn, request)
	case "orde":
		return ProcessOrder(body, path, method, user, idn, request)
	}

	return 400, "Method Invalid"
}

func validateAuth(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") || (path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido"
	}

	todoOk, err, msg := auth.ValidateToken(token)
	if !todoOk {
		if err != nil {
			fmt.Println("Error Token" + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error Token" + msg)
			return false, 401, msg
		}
	}

	fmt.Println("Token Ok")
	return true, 200, msg
}

func ProcessUsers(
	body string,
	path string,
	method string,
	user string,
	id string,
	request events.APIGatewayV2HTTPRequest,
) (int, string) {
	return 400, "Metodo invalido"
}

func ProcessProducts(
	body string,
	path string,
	method string,
	user string,
	id int,
	request events.APIGatewayV2HTTPRequest,
) (int, string) {
	return 400, "Metodo invalido"
}

func ProcessCategory(
	body string,
	path string,
	method string,
	user string,
	id int,
	request events.APIGatewayV2HTTPRequest,
) (int, string) {
	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	case "PUT":
		return routers.UpdateCategory(body, user, id)
	}
	return 400, "Metodo invalido"
}

func ProcessStock(
	body string,
	path string,
	method string,
	user string,
	id int,
	request events.APIGatewayV2HTTPRequest,
) (int, string) {
	return 400, "Metodo invalido"
}

func ProcessAddress(
	body string,
	path string,
	method string,
	user string,
	id int,
	request events.APIGatewayV2HTTPRequest,
) (int, string) {
	return 400, "Metodo invalido"
}

func ProcessOrder(
	body string,
	path string,
	method string,
	user string,
	id int,
	request events.APIGatewayV2HTTPRequest,
) (int, string) {
	return 400, "Metodo invalido"
}

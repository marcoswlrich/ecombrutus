package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidateToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")

	if len(parts) < 2 {
		fmt.Println("Token invalido ")
		return false, nil, "Token invalido"
	}

	userInfo, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Nao foi possivel decodificar o Token: ", err.Error())
		return false, err, err.Error()
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("Nao foi possivel decodificar a estrutura JSON", err.Error())
		return false, err, err.Error()
	}

	thour := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(thour) {
		fmt.Println("Tempo de expiracao token = " + tm.String())
		fmt.Println("Token expirado!")
		return false, err, "Token expirado !!"
	}

	return true, nil, tkj.Username
}

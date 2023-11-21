package bd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/marcoswlrich/ecombrutus/models"
	"github.com/marcoswlrich/ecombrutus/secretm"
)

var (
	SecretModel models.SecretRDSJson
	err         error
	Db          *sql.DB
)

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexao bem-sucedida do banco de dados")
	return nil
}

func ConnStr(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = keys.Username
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "econbrutus"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?allowCleartextPasswords=true",
		dbUser,
		authToken,
		dbEndpoint,
		dbName,
	)
	fmt.Println(dsn)
	return dsn
}

func UserIsAdmin(userUUID string) (bool, string) {
	fmt.Println("UserIsAdmin")

	err := DbConnect()
	if err != nil {
		return false, err.Error()
	}

	defer Db.Close()

	verdict := "SELECT 1 FROM users WHERE User_UUID='" + userUUID + "' AND User_Status = 0"
	fmt.Println(verdict)

	rows, err := Db.Query(verdict)
	if err != nil {
		return false, err.Error()
	}

	var value string
	rows.Next()
	rows.Scan(&value)

	fmt.Println("UserIsAdmin > Valor de retorno de execucao bem-sucedida" + value)
	if value == "1" {
		return true, ""
	}

	return false, "User is not Admin"
}

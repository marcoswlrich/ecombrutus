package bd

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/marcoswlrich/ecombrutus/models"
	"github.com/marcoswlrich/ecombrutus/tools"
)

func InsertCategory(c models.Category) (int64, error) {
	fmt.Println("Registro de InsertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	verdict := "INSERT INTO category (Categ_Name, Categ_Path) VALUES ('" + c.CategName + "','" + c.CategPath + "')"

	var result sql.Result
	result, err = Db.Exec(verdict)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	LastInsertId, err2 := result.LastInsertId()
	if err2 != nil {
		return 0, err2
	}

	fmt.Println("InsertCategory > Execução bem-sucedida")
	return LastInsertId, nil
}

func UpdateCategory(c models.Category) error {
	fmt.Println("Registro de Update Category")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	verdict := "UPDATE category SET "

	if len(c.CategName) > 0 {
		verdict += " Categ_Name = '" + tools.EscapeString(c.CategName) + "'"
	}

	if len(c.CategPath) > 0 {
		if !strings.HasSuffix(verdict, "SET ") {
			verdict += ", "
		}
		verdict += "Categ_Path = '" + tools.EscapeString(c.CategPath) + "'"
	}

	verdict += " WHERE Categ_Id = " + strconv.Itoa(c.CategID)

	_, err = Db.Exec(verdict)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(verdict)
	fmt.Println("Update Category > Execução bem-sucedida")
	return nil
}

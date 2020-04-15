package users_db

import (
	"database/sql"
	"fmt"
	"github.com/AlexisDevGrp/bookstore_users_api/domain/users"
	"github.com/AlexisDevGrp/bookstore_users_api/utils/date_utils"
	"github.com/AlexisDevGrp/bookstore_users_api/utils/mess"
	_ "github.com/godror/godror"
	"log"
	"os"
)
const (
	ora_username = "ora_username"
	ora_pwd = "ora_pwd"
	ora_host = "ora_host"
	ora_schema = "ora_schema"
	ora_port = "ora_port"
)

var(

	oraUsr = os.Getenv(ora_username)
	oraPwd = os.Getenv(ora_pwd)
	oraH = os.Getenv(ora_host)
	oraSchema = os.Getenv(ora_schema)
	oraP = os.Getenv(ora_port)
	Client *sql.DB
)
func Init() {

	configOra := fmt.Sprintf("%s/%s@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=tcp)" +
		"(HOST=%s)(PORT=%s)))(CONNECT_DATA=(SERVICE_NAME=%s)))", oraUsr, oraPwd, oraH,oraP, oraSchema)
	fmt.Println("configOra: ", configOra)
	fmt.Println("we are opening the database")
	Client, err := sql.Open("godror", configOra)
	if err != nil {
		fmt.Println("Pb opening the database")
		panic(err)
	}
	defer Client.Close()
	fmt.Println("Ping")
	if err = Client.Ping(); err != nil {
		fmt.Println("Pb Ping")
		panic(err)
	}
	log.Println("Database successfully configured")

}
func SaveUser(q string, u *users.User )  *mess.RestMsg {
	Init()
	stmt, err := Client.Prepare(q)
	if err != nil {
		return mess.ItemNotCreated(err.Error())
	}

	defer stmt.Close()

	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated)
	if err != nil {
		return mess.ItemNotCreated(err.Error())
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mess.ItemNotCreated(err.Error())
	}
	u.Id = userId
	u.DateCreated = date_utils.GetNowString()

}

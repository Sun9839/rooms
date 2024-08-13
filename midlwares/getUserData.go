package midlwares

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

type User struct {
	id string `json:"id"`
    login     string `json:"login"`
    password string `json:"password"`
	name string `json:"name"`
	surname string `json:"surname"`
	wasborn string `"json:"wasborn"`
	wascreate string `json:"wascreate"`
	
}

func GetUserData(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	var req string = "select * from test_database3.users WHERE login = '" + vars["user"] + "'"
	db, err := sql.Open("mysql", "root:YT93epyfNM!@/test_database3")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	user, err := db.Query(req)
	if err != nil {
        log.Println(err)
    }
    defer user.Close()

	users := []User{}

	for user.Next(){
        u := User{}
        err := user.Scan(&u.id, &u.login, &u.password, &u.name, &u.surname, &u.wasborn, &u.wascreate)
        if err != nil{
            fmt.Println(err)
            continue
        }
        users = append(users, u)
    }

	json.NewEncoder(w).Encode(users)
}
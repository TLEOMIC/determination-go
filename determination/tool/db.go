package tool

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"fmt"
)

var dbs map[string]*sql.DB

func init(){
	dbs = make(map[string]*sql.DB)
}
func DbInit(){
	DataBaseConfig := DataBaseConfigAll()
	for i := range DataBaseConfig{
		DBC := DataBaseConfig[i].(map[string]string)
		db, err := sql.Open("mysql", DBC["USER"]+":"+DBC["PASSWORD"]+"@"+DBC["NETWORK_PROTOCOL"]+"("+DBC["IP"]+":"+DBC["PORT"]+")/"+DBC["DATABASE"])
		if err != nil {
			panic(err)
		}
		var num int
		num , _ =strconv.Atoi(DBC["CONN_MAX_LIFE_TIME"])
		db.SetConnMaxLifetime(time.Minute * time.Duration(num))
		num , _ =strconv.Atoi(DBC["MAX_OPEN_CONNS"])
		db.SetMaxOpenConns(num)
		num , _ =strconv.Atoi(DBC["MAX_IDLE_CONNS"])
		db.SetMaxIdleConns(num)
		db.Ping()
		dbs[i] = db
	}

}
func Db(database string) *sql.DB{
	return dbs[database]
}
func Select(database string,rule map[string]interface{}) (Rta []map[string]interface{}){
	db := Db(database)
	field := ""
	where := ""
	table := ""
	var whereP []interface{}
	order := ""
	limit := ""
	var fields []interface{}
	var fieldIndexes [][]string
	for index, value := range rule {
		switch index {
			case "field":
				var fieldArr []string
				for fieldindex, fieldvalue := range value.(map[string]string) {

					switch fieldvalue {
						case "string":
							fieldIndexes = append(fieldIndexes,[]string{"string",fieldindex})
							fields = append(fields,new(string))
						case "int":
							fieldIndexes = append(fieldIndexes,[]string{"int",fieldindex})
							fields = append(fields,new(int))
					}
					fieldArr = append(fieldArr,fieldindex)
				}
				field = Implode(fieldArr)
			case "where":
				where = " where "+value.(string)
			case "table":
				table = value.(string)
			case "whereP":
				whereP = value.([]interface{})
			case "order":
				order = " order by "+ value.(string)
			case "limit":
				limit = " limit "+ value.(string)
		}
	}
	rows, err := db.Query("select "+field+" from "+table+" "+where+" "+order+" "+limit,whereP...)
	defer rows.Close()
	if err != nil {  
		fmt.Println(err)
	}
	for rows.Next(){
		err := rows.Scan(fields...)
		if err != nil {  
			fmt.Println(err)
		}
		Data := make(map[string]interface{})
		for fieldindex, fieldvalue := range fieldIndexes {
			switch fieldvalue[0] {
				case "string":
					Data[fieldvalue[1]] = *fields[fieldindex].(*string)
				case "int":
					Data[fieldvalue[1]] = *fields[fieldindex].(*int)
			}
		}
		Rta = append(Rta,Data)
	}
	return Rta
}


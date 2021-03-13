package tool

import (
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
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
		dbs[i] = db
	}

}
func Db(database string) *sql.DB{
	return dbs[database]
}


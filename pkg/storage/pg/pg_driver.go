package pg

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mitchellh/mapstructure"

	"app/config"
	"app/pkg"
)

/* -------------------------------------------------------------------------- */
/*                                Query To Map                                */
/* -------------------------------------------------------------------------- */
// m := pg.QueryToMap("SELECT * FROM models WHERE ID = $1", 1)
// res, _ := pkg.Marshal(m)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func QueryToMap(query string, args ...interface{}) map[string]interface{} {
	db := CreatePgConnection()
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}

	var m map[string]interface{}
	cols, _ := rows.Columns()
	for rows.Next() {
		m = parseToMap(rows, cols)
	}

	return m
}

/* -------------------------------------------------------------------------- */
/*                               Query To Struct                              */
/* -------------------------------------------------------------------------- */
// model := pg.QueryToStruct[model.Model]("SELECT * FROM models WHERE ID = $1", 1)
// fmt.Println(model.Model)
/* -------------------------------------------------------------------------- */
func QueryToStruct[K comparable](query string, args ...interface{}) K {
	var result K
	m := QueryToMap(query, args...)
	err := mapstructure.Decode(m, &result)
	if err != nil {
		pkg.Logger.Critical(err.Error())
		panic(err)
	}

	return result
}

/* -------------------------------------------------------------------------- */
/*                              Private Functions                             */
/* -------------------------------------------------------------------------- */
func CreatePgConnection() *sql.DB {
	appConfig := config.AppConfig

	dsn := fmt.Sprintf("host=%s users=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		appConfig.DbHost,
		appConfig.DbUser,
		appConfig.DbPassword,
		appConfig.DbName,
		appConfig.DbPort,
		appConfig.DbSslMode,
		appConfig.DbTimezone,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		pkg.Logger.Critical(err.Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		pkg.Logger.Critical(err.Error())
		panic(err)
	}

	return db
}

func parseToMap(rows *sql.Rows, cols []string) map[string]interface{} {
	values := make([]interface{}, len(cols))
	pointers := make([]interface{}, len(cols))
	for i := range values {
		pointers[i] = &values[i]
	}

	err := rows.Scan(pointers...)
	if err != nil {
		panic(err)
	}

	m := make(map[string]interface{})
	for i, colName := range cols {
		if values[i] == nil {
			m[colName] = nil
		} else {
			m[colName] = values[i]
		}
	}
	return m
}

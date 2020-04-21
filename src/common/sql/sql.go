package Sql

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	// "runtime"
	"strconv"
	"strings"

	//第三方---------------------------
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

//
var _driver, _source string
var _db *sql.DB

//连接/重连
func Connect(driver, source string) *sql.DB {
	_driver, _source = driver, source
	var err error
	_db, err = sql.Open(driver, source)
	if err != nil {
		panic(err)
	}
	return _db
}

func GetSqlPath() string {
	wd, _ := os.Getwd()
	return wd + "/db/"
}

//TODO 初始化及增量更新
func UpdateDB(serverId int64, dbname string) {
	sqlPath := GetSqlPath()

	Exec("set @@global.auto_increment_increment=1000;")
	Exec("set @@global.auto_increment_offset=?;", serverId) //服ID
	//Exec("show variables like 'auto_increment%';") //ID分服,可能无效

	tx, err := _db.Begin()
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Exec("set @@auto_increment_increment=1000;")
	tx.Exec("set @@auto_increment_offset=?;", serverId) //服ID
	rows, err := tx.Query("select count(*) from INFORMATION_SCHEMA.TABLES where TABLE_SCHEMA=? and TABLE_NAME='version';", dbname)
	//rows, err := tx.Query("select count(*) from pg_class where relname='version'")
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	rows.Next()
	var count int32 = 0
	err = rows.Scan(&count)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	rows.Close()
	fmt.Println("count:", count)
	var oldVer, newVer int
	oldVer, newVer = 0, 0

	files, _ := ioutil.ReadDir(sqlPath)
	for _, f := range files {
		if f.Name() != "install.sql" {
			ver, err := strconv.Atoi(strings.Split(f.Name(), ".")[0])
			if err == nil {
				newVer = int(math.Max(float64(newVer), float64(ver)))
			}
		}
	}

	if count == 0 { //初始化
		fmt.Println(">>DataBase Initing!______________")
		//mysql

		//run install.sql
		fstr, err := ioutil.ReadFile(sqlPath + "install.sql")
		if err != nil {
			panic(err)
		}
		strs := strings.Split(string(fstr), ";")
		for _, str := range strs[:len(strs)-1] {
			fmt.Println(">>Sql.Exec:", str)
			_, err = tx.Exec(str)
			if err != nil {
				tx.Rollback()
				panic(err)
			}
		}
		_, err = tx.Exec("insert into version(sid, ver) values (?, ?);", serverId, newVer) //服ID
		if err != nil {
			tx.Rollback()
			panic(err)
		}
	} else { //增量更新
		fmt.Println(">>DBUpdate versionUp!______________")
		rows, err = tx.Query("select ver from version;")
		rows.Next()
		err = rows.Scan(&oldVer)
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		rows.Close()
		fmt.Println(">>DB Old Version:", oldVer, "______________")
		fmt.Println(">>DB New Version:", newVer, "______________")

		for i := oldVer + 1; i <= newVer; i++ {
			//run i.sql
			fstr, err := ioutil.ReadFile(sqlPath + strconv.Itoa(i) + ".sql")
			if err != nil {
				panic(err)
			}
			strs := strings.Split(string(fstr), ";")
			for _, str := range strs[:len(strs)-1] {
				fmt.Println(">>Sql.Exec:", str)
				_, err = tx.Exec(str)
				if err != nil {
					tx.Rollback()
					panic(err)
				}
			}
			_, err = tx.Exec("update version set ver = ?;", newVer) //服ID
			if err != nil {
				tx.Rollback()
				panic(err)
			}
		}
	}
	tx.Commit()
}

func Begin() (*sql.Tx, error) {
	tx, err := _db.Begin()

	return tx, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	return _db.Query(query, args...)
}

func Query2Map(query string, args ...interface{}) []map[string]interface{} {
	rows, err := _db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	columns, _ := rows.Columns()
	length := len(columns)

	var ret []map[string]interface{}
	for rows.Next() {
		value := make([]interface{}, length)
		columnPointers := make([]interface{}, length)
		for i := 0; i < length; i++ {
			columnPointers[i] = &value[i]
		}
		rows.Scan(columnPointers...)
		data := make(map[string]interface{})
		for i := 0; i < length; i++ {
			key := columns[i]
			val := columnPointers[i].(*interface{})
			data[key] = *val
		}
		ret = append(ret, data)
		//fmt.Println(data)
	}
	return ret
}

func Query2Map1(str string, args ...interface{}) map[string]interface{} {
	r := Query2Map(str, args...)
	return r[0]
}

func Exec(str string, args ...interface{}) (int64, int64, error) {
	res, err := _db.Exec(str, args...)
	if err != nil {
		return 0, 0, err
	}
	a, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	l, err := res.LastInsertId()
	if err != nil {
		return 0, 0, err
	}
	return a, l, err
}

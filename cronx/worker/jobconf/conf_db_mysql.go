package jobconf

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlDb struct {
	dsn     string // data source name
	dsnDesc string // 数据源描述，不包含敏感信息，日志中使用
}

func (sf *mysqlDb) Init() error {
	// 数据库配置 TODO
	host := "127.0.0.1"
	port := 3306
	user := "yxy"
	passwd := "123456aabc"
	dbName := "cronx"
	sf.dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, passwd, host, port, dbName)
	sf.dsnDesc = fmt.Sprintf("%s:**@tcp(%s:%d)/%s?charset=utf8", user, host, port, dbName)
	return nil
}

func (sf *mysqlDb) Load() (bool, error) {
	// 连接数据库
	db, err := sql.Open("mysql", sf.dsn)
	if err != nil {
		return false, fmt.Errorf("open_mysql_failed:%v. db=%s", err, sf.dsnDesc)
	}
	defer db.Close()

	// 加载对象表
	if err = loadObjs(db); err != nil {
		return false, fmt.Errorf("load_objs_failed:%v. db=%s", err, sf.dsnDesc)
	}

	// 加载规则表

	return true, nil
}

func loadObjs(db *sql.DB) error {
	sqlStr := "" +
		"SELECT type, obj_group, obj_key, version, flags " +
		"FROM job_conf_obj " +
		"WHERE obj_group = ''"
	rows, err := db.Query(sqlStr)
	if err != nil {
		return fmt.Errorf("db_query:%v", err)
	}
	defer rows.Close()

	var (
		tp       string
		objGroup string
		objKey   string
		version  string
		flags    string
	)
	for rows.Next() {
		err = rows.Scan(&tp, &objGroup, &objKey, &version, &flags)
		fmt.Println(tp, objGroup, objKey)
	}
	return nil
}

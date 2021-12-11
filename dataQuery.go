package week2

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func QueryUserName() error {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 主要代码
	var name string
	id := 1
	err = db.QueryRow("select name from users where id = ?", id).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			// there were no rows, but otherwise no error occurred
			// sql.ErrNoRows 由 QueryRow方法返回
			// 大多数情况下需要处理
            // 对sql.ErrNoRows包装后返回
			return errors.Wrapf(err, "[用户查询]失败  不存在id=%d的用户", id)
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(name)
	return nil
}

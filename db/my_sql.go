package db

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

/**
*@author: 廖理
*@date:2022/12/22
**/

func dbMysql() {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/adwords"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("链接成功", db)
	defer db.Close() // 注意这行代码要写在上面err判断的下面
}

//为什么上面代码中的defer db_sql.Close()语句不应该写在if err != nil的前面呢？

//Open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库的连接。如果要检查数据源的名称是否真实有效，应该调用Ping方法。
//
//返回的DB对象可以安全地被多个goroutine并发使用，并且维护其自己的空闲连接池。因此，Open函数应该仅被调用一次，很少需要关闭这个DB对象。
//
//接下来，我们定义一个全局变量db，用来保存数据库连接对象。将上面的示例代码拆分出一个独立的initDB函数，
//只需要在程序启动时调用一次该函数完成全局变量db的初始化，其他函数中就可以直接使用全局变量db了。（注意下方的注意）
// 定义一个全局对象db
var db_sql *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:123456@tcp(127.0.0.1:3306)/adwords?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db_sql, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db_sql.Ping()
	if err != nil {
		return err
	}
	fmt.Println("链接成功", db_sql)
	return nil

}

//其中sql.DB是表示连接的数据库对象（结构体实例），它保存了连接数据库相关的所有信息。
//它内部维护着一个具有零到多个底层连接的连接池，它可以安全地被多个goroutine同时使用。
//
//SetMaxOpenConns
//func (db_sql *DB) SetMaxOpenConns(n int)
//SetMaxOpenConns设置与数据库建立连接的最大数目。 如果n大于0且小于最大闲置连接数，
//会将最大闲置连接数减小到匹配最大开启连接数的限制。 如果n<=0，不会限制最大开启连接数，默认为0（无限制）。
//
//SetMaxIdleConns
//func (db_sql *DB) SetMaxIdleConns(n int)
//SetMaxIdleConns设置连接池中的最大闲置连接数。 如果n大于最大开启连接数，
//则新的最大闲置连接数会减小到匹配最大开启连接数的限制。 如果n<=0，不会保留闲置连接。

type user struct {
	ID   int64
	Name string
	Age  int64
}

// 查询单条数据示例
func queryRowDemo(id int) {
	sqlStr := `select id, name, age from user where id=?`
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db_sql.QueryRow(sqlStr, id).Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
}

// 查询多条数据示例
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := db_sql.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
	}
}

// 查询多条数据示例
func queryMultiRowDemo2(ids []int64) {
	sqlStr := "select id, name, age from user where id in ?"
	rows, err := db_sql.Query(sqlStr, ids)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.ID, u.Name, u.Age)
	}
}

// 插入数据
func insertRowDemo(name string, age int64) {
	sqlStr := "insert into user(name, age) values (?,?)"
	ret, err := db_sql.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo(id, age int64) {
	sqlStr := "update user set age=? where id = ?"
	ret, err := db_sql.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo(id int64) {
	sqlStr := "delete from user where id = ?"
	ret, err := db_sql.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func DbDemo() {
	//dbMysql()
	initDB()

	//insertRowDemo("玩哈哈",38)
	//queryRowDemo(1)
	//queryMultiRowDemo()
	//queryMultiRowDemo2([]int64{1,2})
	//updateRowDemo(1,100)

	//deleteRowDemo(2)
	//prepareQueryDemo()
	//prepareInsertDemo()

	transactionDemo()
}

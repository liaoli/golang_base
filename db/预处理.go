package db

import "fmt"

/**
*@author: 廖理
*@date:2022/12/22
**/

//MySQL预处理
//什么是预处理？
//普通SQL语句执行过程：
//
//客户端对SQL语句进行占位符替换得到完整的SQL语句。
//客户端发送完整SQL语句到MySQL服务端
//MySQL服务端执行完整的SQL语句并将结果返回给客户端。
//预处理执行过程：
//
//把SQL语句分成两部分，命令部分与数据部分。
//先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
//然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
//MySQL服务端执行完整的SQL语句并将结果返回给客户端。
//为什么要预处理？
//优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
//避免SQL注入问题。

//Go实现MySQL预处理
//database/sql中使用下面的Prepare方法来实现预处理操作。
//
//func (db_sql *DB) Prepare(query string) (*Stmt, error)
//Prepare方法会先将sql语句发送给MySQL服务端，返回一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。
//
//查询操作的预处理示例代码如下：

// 预处理查询示例
func prepareQueryDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := db_sql.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
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

// 预处理插入示例
func prepareInsertDemo() {
	sqlStr := "insert into user(name, age) values (?,?)"
	stmt, err := db_sql.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("小王子", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("沙河娜扎", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")
}

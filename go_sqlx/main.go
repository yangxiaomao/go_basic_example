package main

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

// type user struct {
// 	Id   int
// 	Age  int
// 	Name string
// }

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/task?charset=utf8mb4&parseTime=True"
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

// // 查询单条数据实例
// func queryRowDemo() {
// 	sqlStr := "select id, name, age from user where id=?"
// 	var u user
// 	err := db.Get(&u, sqlStr, 2)
// 	if err != nil {
// 		fmt.Printf("get failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)

// }

// // 查询多条数据示例
// func queryMultiRowDemo() {
// 	sqlStr := "select id, name, age from user where id > ?"
// 	var users []user
// 	err := db.Select(&users, sqlStr, 0)
// 	if err != nil {
// 		fmt.Printf("query failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Printf("users:%#v\n", users)
// }

// // 插入数据
// func insertRowDemo() {
// 	sqlStr := "insert into user(name, age) values (?,?)"
// 	ret, err := db.Exec(sqlStr, "通州毛总", 19)
// 	if err != nil {
// 		fmt.Printf("insert failed, err:%v\n", err)
// 		return
// 	}
// 	theID, err := ret.LastInsertId() // 新插入数据的ID
// 	if err != nil {
// 		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Printf("insert success, the id is %d.\n", theID)
// }

// // 更新数据
// func updateRowDemo() {
// 	sqlStr := "update user set age=? where id = ?"
// 	ret, err := db.Exec(sqlStr, 39, 6)
// 	if err != nil {
// 		fmt.Printf("update failed, err:%v\n", err)
// 		return
// 	}
// 	n, err := ret.RowsAffected() // 操作影响的行数
// 	if err != nil {
// 		fmt.Printf("get RowsAffected failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Printf("update success, affected rows:%d\n", n)
// }

// // 删除数据
// func deleteRowDemo(id int64) {
// 	sqlStr := "delete from user where id = ?"
// 	ret, err := db.Exec(sqlStr, id)
// 	if err != nil {
// 		fmt.Printf("delete failed, err:%v\n", err)
// 		return
// 	}
// 	n, err := ret.RowsAffected() // 操作影响的行数
// 	if err != nil {
// 		fmt.Printf("get RowsAffected failed, err:%v\n", err)
// 		return
// 	}
// 	fmt.Printf("delete success, affected rows:%d\n", n)

// }

// // DB.NamedExec 方法用来绑定SQL语句与结构体或map中的同名字段
// func insertUserDemo() (err error) {
// 	sqlStr := "INSERT INTO user (name, age) VALUES (:name, :age)"
// 	_, err = db.NamedExec(sqlStr,
// 		map[string]interface{}{
// 			"name": "七米",
// 			"age":  28,
// 		})
// 	return
// }

// // DB.NamedQuery 与 DB.NamedExec 同理，这里是支持查询。
// func namedQuery() {
// 	sqlStr := "SELECT * FROM user WHERE name=:name"
// 	// 使用map做命名查询
// 	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "七米"})
// 	if err != nil {
// 		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
// 		return
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var u user
// 		err := rows.StructScan(&u)
// 		if err != nil {
// 			fmt.Printf("scan failed, err:%v\n", err)
// 			continue
// 		}
// 		fmt.Printf("user:%#v\n", u)
// 	}

// 	u := user{
// 		Name: "七米",
// 	}
// 	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
// 	rows, err = db.NamedQuery(sqlStr, u)
// 	if err != nil {
// 		fmt.Printf("db.NameQuery failed, err:%v\n", err)
// 		return
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var u user
// 		err := rows.StructScan(&u)
// 		if err != nil {
// 			fmt.Printf("scan failed, err:%v\n", err)
// 			continue
// 		}
// 		fmt.Printf("user:%v\n", u)
// 	}
// }

// // 事务操作 ， sqlx中提供的db.Beginx() 和 tx.Exec()
// func transactionDemo2() (err error) {
// 	tx, err := db.Beginx() // 开启事务
// 	if err != nil {
// 		fmt.Printf("begin trans failed, err:%v\n", err)
// 		return err
// 	}
// 	defer func() {
// 		if p := recover(); p != nil {
// 			tx.Rollback()
// 			panic(p) // re-throw panic after Rollback
// 		} else if err != nil {
// 			fmt.Println("rollback, err=", err)
// 			tx.Rollback() // err is non-nil; don't change it
// 		} else {
// 			err = tx.Commit() // err is nil; if Commit returns error update err
// 			fmt.Println("commit")
// 		}
// 	}()

// 	sqlStr1 := "Update user set age=20 where id=?"

// 	rs, err := tx.Exec(sqlStr1, 4)
// 	if err != nil {
// 		return err
// 	}
// 	n, err := rs.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if n != 1 {
// 		return errors.New("exec sqlStr1 failed")
// 	}
// 	sqlStr2 := "Update user set age=50 where id=?"
// 	rs, err = tx.Exec(sqlStr2, 5)
// 	if err != nil {
// 		return err
// 	}
// 	n, err = rs.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if n != 1 {
// 		return errors.New("exec sqlStr1 failed")
// 	}
// 	return err
// }

// // BatchInsertUsers 自行构造批量插入的语句, 比较笨，但很好理解
// func BatchInsertUsers(users []*User) error {
// 	// 存放（?,?）的slice
// 	valueStrings := make([]string, 0, len(users))
// 	// 存放values的slice
// 	valueArgs := make([]interface{}, 0, len(users)*2)
// 	// 遍历users准备相关数据
// 	for _, u := range users {
// 		// 此处占位符要与插入值的个数对应
// 		valueStrings = append(valueStrings, "(?, ?)")
// 		valueArgs = append(valueArgs, u.Name)
// 		valueArgs = append(valueArgs, u.Age)
// 	}
// 	// 自行拼接要执行的具体语句
// 	stmt := fmt.Sprintf("INSERT INTO user (name, age) VALUES %s",
// 		strings.Join(valueStrings, ","))
// 	_, err := db.Exec(stmt, valueArgs...)
// 	return err
// }

// // 使用 sqlx.In 实现批量插入代码如下：
// // BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数， 注意传入的参数是[]interface{}
// func BatchInsertUsers2(users []interface{}) error {
// 	query, args, _ := sqlx.In(
// 		"INSERT INTO user (name, age) VALUES (?),(?),(?)",
// 		users..., //如果arg实现了driver.Valuer, sqlx.In 会通过调用 Value()来展开它
// 	)
// 	fmt.Println(query) // 查看生成的queryString
// 	fmt.Println(args)  // 查看生成的args
// 	_, err := db.Exec(query, args...)
// 	return err
// }

// // 使用NamedExec 实现批量插入的代码如下：
// // BatchInsertUsers3 使用NamedExec实现批量插入
// func BatchInsertUsers3(users []*User) error {
// 	_, err := db.NamedExec("INSERT INTO user (name, age) VALUES (:name, :age)", users)
// 	return err
// }

// // in查询
// // QueryByIDs 根据给定ID查询
// func QueryByIDs(ids []int) (users []User, err error) {
// 	// 动态填充id
// 	query, args, err := sqlx.In("select name, age from user where id in (?)", ids)
// 	if err != nil {
// 		return
// 	}
// 	// sqlx.In 返回带`?`bindvar的查询语句， 我们使用Rebind()重新绑定它
// 	query = db.Rebind(query)

// 	err = db.Select(&users, query, args...)
// 	return
// }

// in 查询和FIND_IN_SET函数
// 查询id在给定id集合的数据并维持给定id集合的顺序
func QueryAndOrderByIDs(ids []int) (users []User, err error) {
	// 动态填充id
	strIDs := make([]string, 0, len(ids))
	for _, id := range ids {
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	query, args, err := sqlx.In("select name, age from user where id in (?) order by FIND_IN_SET(id, ?)", ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}

	// sqlx.In 返回带`?`bindvar的查询语句， 我们使用Rebind()重新绑定它
	query = db.Rebind(query)

	err = db.Select(&users, query, args...)
	return
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	ids := []int{1, 7, 3, 10}
	users, err := QueryAndOrderByIDs(ids)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)

	// ids := []int{1, 7, 3}
	// users, err := QueryByIDs(ids)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(users)
	// defer db.Close()
	// u1 := User{Name: "八米11", Age: 18}
	// u2 := User{Name: "毛总211", Age: 20}
	// u3 := User{Name: "小王子11", Age: 28}

	// // 方法1
	// users := []*User{&u1, &u2, &u3}
	// err = BatchInsertUsers(users)
	// if err != nil {
	// 	fmt.Printf("BatchInsertUsers failed, err:%v\n", err)
	// }

	// // 方法2
	// users2 := []interface{}{u1, u2, u3}
	// err = BatchInsertUsers2(users2)
	// if err != nil {
	// 	fmt.Printf("BatchInsertUsers2 failed, err:%v\n", err)
	// }

	// // 方法3
	// users3 := []*User{&u1, &u2, &u3}
	// err = BatchInsertUsers3(users3)
	// if err != nil {
	// 	fmt.Printf("BatchInsertUsers3 failed, err:%v\n", err)
	// }

	// fmt.Println("sqlx 成功")
	// queryRowDemo()

	// queryMultiRowDemo()

	// insertRowDemo()

	// updateRowDemo()

	// deleteRowDemo(2)

	// insertUserDemo()

	// namedQuery()

	// transactionDemo2()
}

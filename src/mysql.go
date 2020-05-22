package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {
	//数据库配置
	db,err2 := sql.Open("mysql","root:@(127.0.0.1:3306)/studygo")
	if err2 != nil{
		fmt.Println(err2)
	}
	defer db.Close()
	//链接数据库
	err := db.Ping()
	if err != nil{
		fmt.Println(err)
		return
	}
	//操作数据库
	//sql :="insert into go_user values(5,'wsq',1)"
	//result,_:= db.Exec(sql)
	//n,err :=result.LastInsertId()
	//fmt.Println(n)

	//执行预处理
	//stu :=[][]string{{"szl","1"},{"szl","1"}}
	//stmt,err:=db.Prepare("insert into go_user(name,sex) values(?,?)")
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//for _,user := range stu{
	//	stmt.Exec(user[0],user[1])
	//}


	//查询单行数据
	//var id,name,sex string
	//row := db.QueryRow("select * from go_user where id =4")
	//row.Scan(&id,&name,&sex)
	//fmt.Println(id)
	//fmt.Println(name)
	//fmt.Println(sex)

	//获取多行数据
	var id,name,sex string
	rows,_ := db.Query("select * from go_user")
	for rows.Next(){
		rows.Scan(&id,&name,&sex)
		fmt.Println(id,name,sex)
	}

}

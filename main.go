package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	DbConnect 		*sql.DB
)

const (
	host     = "zhangzhen.c.cloudtogo.cn"
	port     = 36413
	user     = "rp123456"
	password = "rp123456"
	dbname   = "postgres"
)

func Init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	Db, ersmg := sql.Open("postgres", psqlInfo)
	if ersmg != nil {
		panic(ersmg)
	}
	ersmg = Db.Ping()
	if ersmg != nil {
		panic(ersmg)
	}
	fmt.Println("Successfully connected!")
	DbConnect=Db
}


func main() {
	Init()
	err:=DbConnect.Ping()
	if err!=nil {
		fmt.Println(err.Error())
	}
	//Insert() //新增
	//Del() //删除
	//Update()//修改
	query() //查询
}

func Insert()  {
	sqlStatement := ` INSERT INTO user_info (uid, user_name,dept_name) VALUES ($1, $2,$3) RETURNING uid`
	id := 5
	err := DbConnect.QueryRow(sqlStatement, id, "小汪223","市场部223").Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
}

func Del() {
	sqlStatement := ` delete from user_info where uid = $1`
	id := 5
	result, err := DbConnect.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	result_int, _ := result.RowsAffected()
	if result_int > 0 {
		fmt.Printf("操作成功 ID is:%d", result_int)
	} else {
		fmt.Printf("操作失败 id is:%d", result_int)
	}
}

func Update()  {
	sqlStatement := `update user_info set user_name =$1 where  uid= $2`
	id := 4
	result, err := DbConnect.Exec(sqlStatement, "小李",id)
	if err != nil {
		panic(err)
	}
	result_int, _ := result.RowsAffected()
	if result_int > 0 {
		fmt.Printf("操作成功 ID is:%d", result_int)
	} else {
		fmt.Printf("操作失败 id is:%d", result_int)
	}
}

func query()  {
	sqlStatement := ` select * from user_info where uid = $1`
	id := 4
	result,err := DbConnect.Query(sqlStatement, id)
	if err!=nil {
		fmt.Println(err.Error())
	}
	var m UserInfo
	for result.Next(){
		result.Scan(&m.Uid,&m.UserName,&m.DeptName)
	}
	fmt.Println(m)
}

type UserInfo struct {
	Uid int
	UserName string
	DeptName string
}
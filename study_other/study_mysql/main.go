package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func queryOne() {
	var user User
	err := DB.QueryRow("select name,age from users;").Scan(&user.Name, &user.Age)
	if err != nil {
		fmt.Println("scan filed err:", err.Error())
	}
	fmt.Println(user)
}

func queryMany() {
	var user []*User
	sql := "select name,age from users;"
	rows, err := DB.Query(sql)
	if err != nil {
		fmt.Println("query err:", err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u User
		err := rows.Scan(&u.Name, &u.Age)
		if err != nil {
			fmt.Println("scan filed err:", err.Error())
			return
		}
		user = append(user, &u)
	}
	fmt.Println(len(user))
	fmt.Println(user)
}

func main() {
	InitMySQL()
	queryOne()
	queryMany()
}

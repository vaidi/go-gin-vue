package dao

import (
	"database/sql"
	"fmt"
	"go-gin-vue/core/core_inside"
	"log"
	"time"
)

type UserInfo struct {
	Id         int64     `db:"id"`
	Password   string    `db:"password"`
	Mobile     string    `db:"mobile"`
	RoleId     int64     `db:"role_id"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}

func (u UserInfo) TableName() string {
	return "user_info"
}
func (u UserInfo) Query(useId int64) UserInfo {
	rows, err := core_inside.Db.Query("select id,password,mobile,role_id,create_time,update_time from user_info where id=?", useId)
	if err != nil {
		fmt.Println("失败了")
		panic(err)
	}
	defer rows.Close()
	var userInfo UserInfo
	for rows.Next() {

		if err := rows.Scan(&userInfo.Id, &userInfo.Password, &userInfo.Mobile, &userInfo.RoleId, &userInfo.CreateTime, &userInfo.UpdateTime); err != nil {
			log.Fatal(err)
		}
		//local := time.Local
		//location, timeErr := time.ParseInLocation("2006-01-02 15:04:05", createTime, local)
		//if timeErr != nil {
		//	panic(timeErr)
		//}
		//fmt.Print("location:%v", location)
		//userInfo.CreateTime = location
		//
		//updateLocation, _ := time.ParseInLocation("2006-01-02 15:04:05", updateTime, local)
		//userInfo.UpdateTime = updateLocation
		// 处理user结构体
		fmt.Printf("%#v\n", userInfo)
	}
	return userInfo
}
func QueryOne() {

}

func (u UserInfo) Add(uu UserInfo) bool {
	rows, _ := core_inside.Db.Query("select mobile from user_info where mobile=?", uu.Mobile)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)
	if rows.Next() {
		var mobile string
		rows.Scan(&mobile)
		fmt.Println("已存在手机号", mobile)
		return false
	}

	exec, err := core_inside.Db.Exec("insert into user_info(mobile,password,role_id) values (?,?,?)", uu.Mobile, uu.Password, uu.RoleId)
	if err != nil {
		fmt.Println("insert error:", err.Error())
	}
	res, _ := exec.RowsAffected()
	return res > 0
}

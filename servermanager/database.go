package servermanager

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	. "github.com/KouKouChan/CSO2-Server/blademaster/typestruct"
	. "github.com/KouKouChan/CSO2-Server/verbose"
)

var (
	DB     *sql.DB
	Dblock sync.Mutex
	DBPath string
)

//从数据库中读取用户数据
//如果是新用户则保存到数据库中
func GetUserFromDatabase(loginname string, passwd []byte) (*User, int) {
	u := GetNewUser()
	var bytes []byte
	if DB == nil {
		filepath := DBPath + loginname
		rst, _ := PathExists(filepath)
		if rst {

			Dblock.Lock()
			bytes, _ = ioutil.ReadFile(filepath)

			Dblock.Unlock()

			u.SetID(GetNewUserID())

		} else {
			return nil, USER_NOT_FOUND
		}
	} else {
		sql := "SELECT * FROM UserData Where username = '" + loginname + "'"
		rows, err := DB.Query(sql)
		defer rows.Close()
		if err != nil {
			DebugInfo(1, "Searching User "+loginname+"'s data failed !")
			return nil, USER_UNKOWN_ERROR
		}

		if rows.Next() {
			rows.Scan(&u.Userid, &u.UserName, &u.IngameName, &u.Password, &u.UserMail, &bytes)
		} else {
			return nil, USER_NOT_FOUND
		}

	}
	err := json.Unmarshal(bytes, &u)
	if err != nil {
		DebugInfo(1, "Suffered a error while getting User", string(loginname)+"'s data !", err)
		return nil, USER_UNKOWN_ERROR
	}

	//检查密码
	str := fmt.Sprintf("%x", md5.Sum([]byte(string(loginname)+string(passwd))))
	for i := 0; i < 32; i++ {
		if str[i] != u.Password[i] {
			return nil, USER_PASSWD_ERROR
		}
	}

	DebugInfo(1, "User", u.UserName, "data found !")

	return &u, USER_LOGIN_SUCCESS
}

func AddUserToDB(u *User) error {
	if u == nil {
		return nil
	}
	data, _ := json.MarshalIndent(u, "", "     ")
	if DB == nil { //json
		filepath := DBPath + u.UserName
		Dblock.Lock()
		err := ioutil.WriteFile(filepath, data, 0644)

		Dblock.Unlock()
		if err != nil {
			return err
		}

		filepath = DBPath + u.IngameName + ".check"
		Dblock.Lock()
		err = ioutil.WriteFile(filepath, []byte(u.UserName), 0644)
		Dblock.Unlock()
		if err != nil {
			return err
		}
		return nil
	}
	//mysql
	stmt, _ := DB.Prepare(`INSERT INTO UserData (username,gamename,password,mail,data) VALUES (?, ?, ?, ?, ?)`)
	defer stmt.Close()

	_, err := stmt.Exec(u.UserName, u.IngameName, u.Password, u.UserMail, data)
	if err != nil {
		DebugInfo(1, "Insert User", u.UserName, "data failed !")
		return err
	}

	return nil
}

func UpdateUserToDB(u *User) error {
	if u == nil {
		return nil
	}
	data, _ := json.MarshalIndent(u, "", "     ")
	if DB == nil {
		filepath := DBPath + u.UserName
		Dblock.Lock()
		err := ioutil.WriteFile(filepath, data, 0644)

		Dblock.Unlock()
		if err != nil {
			return err
		}
		return nil
	}
	stmt, _ := DB.Prepare("UPDATE UserData set username=?,gamename=?,password=?,mail=?,data=? where uid=?")
	defer stmt.Close()

	_, err := stmt.Exec(u.UserName, u.IngameName, u.Password, u.UserMail, data, u.Userid)
	if err != nil {
		DebugInfo(1, "Update User", u.UserName, "data failed !")
		return err
	}
	return nil
}

// func IsExistsMail(mail []byte) bool { //该函数已作废
// 	if DB != nil {
// 		query, err := DB.Prepare("SELECT * FROM userinfo WHERE UserMail = ?")
// 		if err == nil {
// 			defer query.Close()
// 			Dblock.Lock()
// 			rows, err := query.Query(mail)
// 			Dblock.Unlock()
// 			if err != nil {
// 				return false
// 			}
// 			defer rows.Close()
// 			if rows.Next() {
// 				return true
// 			}
// 		}
// 		//存在风险，如果出错时候其实该用户存在，那么会出现冗余
// 		return false
// 	}
// 	return false
// }

func IsExistsUser(username []byte) bool {
	if DB == nil {
		filepath := DBPath + string(username)
		rst, _ := PathExists(filepath)
		if rst {
			return true
		}
		return false
	}
	sql := "SELECT * FROM UserData Where username = '" + string(username) + "'"
	rows, err := DB.Query(sql)
	defer rows.Close()
	if err != nil {
		DebugInfo(1, "Searching User "+string(username)+"'s data failed !")
		return false
	}

	if rows.Next() {
		return true
	}
	return false

}

func IsExistsIngameName(name []byte) bool {
	if DB == nil {
		filepath := DBPath + string(name) + ".check"
		rst, _ := PathExists(filepath)
		if rst {
			return true
		}
		return false
	}
	sql := "SELECT * FROM UserData Where gamename = '" + string(name) + "'"
	rows, err := DB.Query(sql)
	defer rows.Close()
	if err != nil {
		DebugInfo(1, "Searching User "+string(name)+"'s data failed !")
		return false
	}

	if rows.Next() {
		return true
	}
	return false
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func SaveAllUsers() bool {
	UsersManager.Lock.Lock()
	defer UsersManager.Lock.Unlock()
	for _, v := range UsersManager.Users {
		if v == nil {
			continue
		}

		if UpdateUserToDB(v) != nil {
			return false
		}
	}
	return true
}

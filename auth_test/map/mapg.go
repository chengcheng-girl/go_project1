package main

import "encoding/json"
type response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data  string     `json:"data"`
	Status bool    `json:"status"`
}
var Map map[string]func(user,pwd string)bool
func f1(user,pwd string)  bool {
	if user =="admin" && pwd=="admin"{
		return true
	}

	return false
}
func f2(user,pwd string)  bool {
	if user =="ldap" && pwd=="ldap"{
		return true
	}

	return false
}
func f3(user,pwd string)  bool {
	if user =="docker" && pwd=="docker"{
		return true
	}

	return false
}

func init() {
	Map = make(map[string]func(user,pwd string)bool)
	Map["admin"] = f1
	Map["ldap"] = f2
	Map["docker"] = f3
}


func main() {
	user:="docker"
	pwd:="docker"
	succ := &response{
		Code:   200,
		Msg:    "认证成功",
		Data:   "ok",
		Status: true,
	}
	fail := &response{
		Code:   500,
		Msg:    "认证失败，账号为user密码为pwd",
		Data:   "ok",
		Status: true,
	}
	s, _ := json.Marshal(&succ)
	f, _ := json.Marshal(&fail)
	for _,v:=range Map{
		a:=v(user,pwd)
		if a{
			print(string(s))
			return
		}
	}
	print(string(f))
}

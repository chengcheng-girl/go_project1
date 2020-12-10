package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data  string     `json:"data"`
	Status bool    `json:"status"`
}
type Auth interface{
	auth(user, pwd  string)bool
}

type DockerVerfiy struct{
}
func(v *DockerVerfiy)auth(user, pwd  string)bool {
	if user == "docker" && pwd == "docker" {
		return true
	}
	return false
}
type LDAPVerfiy struct{
}
func(v *LDAPVerfiy)auth(user, pwd  string)bool {
	if user == "ldap" && pwd == "ldap" {
		return true
	}
	return false
}

type DBVerfiy struct{
}
func(v *DBVerfiy)auth(user, pwd  string)bool {
	if user == "admin" && pwd == "admin" {
		return true
	}
	return false
}
var sMap map[string]Auth
func init(){
	sMap=make(map[string]Auth)
	sMap["admin"]=new(DockerVerfiy)
	sMap["ldap"]=new(DBVerfiy)
	sMap["docker"]=new(LDAPVerfiy)
}
func main() {
	user := "docker"
	pwd := "docker"
	succ := &Response{
		Code:   200,
		Msg:    "认证成功",
		Data:   "ok",
		Status: true,
	}
	fail := &Response{
		Code:   500,
		Msg:    "认证失败，账号为user密码为pwd",
		Data:   "ok",
		Status: true,
	}
	s, _ := json.Marshal(&succ)
	f, _ := json.Marshal(&fail)
	for _, v := range sMap {
		a := v.auth(user, pwd)
		if a {
			print(a)
			fmt.Println(string(s))
			return
		}
	}
	fmt.Println(string(f))


}

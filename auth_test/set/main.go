package main

import (
	"sync"
	"fmt"
)
type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

func New() *Set  {

	return &Set{
		m:map[interface{}]bool{},
	}
}

func (s *Set) Add(item interface{}){
	s.Lock()
	defer s.Unlock()
	s.m[item]=true
}

func (s *Set) Remove(item interface{})  {
	s.Lock()
	defer s.Unlock()
	delete(s.m,item)
}

func (s *Set) Has(item interface{}) bool  {
	s.RLock()
	defer s.RUnlock()
	_,ok:=s.m[item]
	return ok
}

func (s *Set) Clear()  {
	s.Lock()
	defer s.Unlock()
	s.m=map[interface{}]bool{}
}

func (s * Set) List() []interface{} {

	s.RLock()
	defer s.RUnlock()
	list:=[]interface{}{}
	for item:=range s.m{
		list=append(list,item)
	}

	return list
}

func (s *Set) Len() int  {
	return len(s.List())
}

func (s *Set) IsEmpty() bool  {

	if s.Len()==0{
		return true
	}else {
		return false
	}
}

func main()  {

	//初始化
	s:=New()

	//添加
	s.Add("宅男帮")
	s.Add(1234567890)
	s.Add("http://www.zhainanbang.net")
	s.Add(true)
	s.Add(3.14)
	//打印集中的元素和长度
	fmt.Println(s.List(),s.Len())
	s.Add(3.14)
	//打印集中的元素和长度
	fmt.Println(s.List(),s.Len())
	//移除
	s.Remove(3.14)
	//打印集中的元素和长度
	fmt.Println(s.List(),s.Len())
	//判断某个元素是否在集合中
	fmt.Println("宅男帮是否在集合中:",s.Has("宅男帮"))
	//移除
	s.Remove("http://www.zhainanbang.net")
	s.Remove(1234567890)
	//判断是否为空
	fmt.Println("集合是否为空:",s.IsEmpty())
	//打印集中的元素和长度
	fmt.Println(s.List(),s.Len())
	//清空
	s.Clear()
	//判断是否为空
	fmt.Println("集合是否为空:",s.IsEmpty())
	//打印集中的元素和长度
	fmt.Println(s.List(),s.Len())



}
package singletonpattern

import "sync"

//Message Message
type Message struct {
	a int
}

var msg = &Message{a: 1}

//Instance 饿汉式
func Instance() *Message {
	return msg
}

var once = &sync.Once{}
var msgLazy *Message

//InstanceLazy 懒汉式
func InstanceLazy() *Message {

	once.Do(func() {
		msgLazy = &Message{
			a: 1,
		}
	})
	return msgLazy

}

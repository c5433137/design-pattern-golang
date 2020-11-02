package builderpattern

import "sync"

/*
	建造者模式：
		适用于复杂结构体初始化，使用函数的方式屏蔽具体初始化的细节。很多开源项目都是这样用的，增加可读性。
*/

//Header Header
type Header struct {
	SrcAddr  string
	SrcPort  uint64
	DestAddr string
	DestPort uint64
	Items    map[string]string
}

//Body Body
type Body struct {
	Items []string
}

//Message Message
type Message struct {
	Header *Header
	Body   *Body
}

//BuildStu Message对象的Builder对象
type BuildStu struct {
	once *sync.Once
	msg  *Message
}

//Builder Builder
func Builder() *BuildStu {
	return &BuildStu{
		once: &sync.Once{},
		msg:  &Message{Header: &Header{}, Body: &Body{}},
	}
}

//WithSrcAddr WithSrcAddr
func (b *BuildStu) WithSrcAddr(srcAddr string) *BuildStu {
	b.msg.Header.SrcAddr = srcAddr
	return b
}

//WithSrcPort WithSrcPort
func (b *BuildStu) WithSrcPort(srcPort uint64) *BuildStu {
	b.msg.Header.SrcPort = srcPort
	return b
}

//WithDestAddr WithDestAddr
func (b *BuildStu) WithDestAddr(destAddr string) *BuildStu {
	b.msg.Header.DestAddr = destAddr
	return b
}

//WithDestPort WithDestPort
func (b *BuildStu) WithDestPort(destPort uint64) *BuildStu {
	b.msg.Header.DestPort = destPort
	return b
}

//WithHeaderItem WithHeaderItem
func (b *BuildStu) WithHeaderItem(key, value string) *BuildStu {
	b.once.Do(func() {
		b.msg.Header.Items = make(map[string]string)
	})
	b.msg.Header.Items[key] = value
	return b
}

//WithBodyItem WithBodyItem
func (b *BuildStu) WithBodyItem(record string) *BuildStu {
	b.msg.Body.Items = append(b.msg.Body.Items, record)
	return b
}

//Build 创建Message对象，在最后一步调用
func (b *BuildStu) Build() *Message {
	return b.msg
}

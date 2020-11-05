package adapterpattern

/*
	适配器模式：类似于电源转接头。
	适用场景：1、全部替换接口时 有一部分接口 历史原因 还不能立即替换为新接口  就新增一个适配器，无需修改老接口
*/

//Records Records
type Records struct {
	Items []string
}
 
//Consumer Consumer
type Consumer interface {
	Poll() Records
}

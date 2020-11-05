package compositepattern

/*
	组合模式（has-a）：组合优于继承，两种方式 直接组合、嵌入组合（匿名成员特性）
*/

/*
	设计思想：
		struct不依赖interface
		1. 包含角色：
			1). 共同的接口MenuComponent，为root和leaf结构体共有的方法
			2). root结构体(包含leaf列表)和leaf结构体
			3). 将结构体中共同部分的数据抽离构成一个interface，使用匿名组合的方式实现2中的两类结构体
*/

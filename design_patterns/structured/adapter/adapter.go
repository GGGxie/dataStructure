package adapter

import "fmt"

//为两个不兼容的接口之间的建立桥梁，通过适配器进行接口转换

// ICreateServer 创建云主机，目标接口
type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

// AWSClient aws sdk
type AWSClient struct{}

// RunInstance 启动实例
func (c *AWSClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws client run success, cpu： %f, mem: %f", cpu, mem)
	return nil
}

// AwsClientAdapter 适配器
type AwsClientAdapter struct {
	Client AWSClient
}

// CreateServer 启动实例
// 适配器提供了CreateServer方法，而原类提供了RunInstance方法
func (a *AwsClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.RunInstance(cpu, mem)
	return nil
}

// AliyunClient aliyun sdk
type AliyunClient struct{}

// CreateServer 启动实例
func (c *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aws client run success, cpu： %d, mem: %d", cpu, mem)
	return nil
}

// AliyunClientAdapter 适配器
type AliyunClientAdapter struct {
	Client AliyunClient
}

// CreateServer 启动实例
// 适配器提供了float64类型的入参，而原类提供的是int类型的入参
func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.CreateServer(int(cpu), int(mem))
	return nil
}

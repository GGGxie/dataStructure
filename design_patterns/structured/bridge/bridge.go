package bridge

import "fmt"

//将抽象类和实现类解耦，使得两者能自主变化
//EmailMsgSender和PhoneMsgSender是实现类，都有send方法，INotification是抽象类，调用了它们的send方法

// IMsgSender IMsgSender
type IMsgSender interface {
	Send(msg string) error
}

// EmailMsgSender 发送邮件
// 可能还有 电话、短信等各种实现
type EmailMsgSender struct {
	emails []string
}

// NewEmailMsgSender NewEmailMsgSender
func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

// Send Send
func (s *EmailMsgSender) Send(msg string) error {
	// 这里去发送消息
	fmt.Printf("send email %s\n", msg)
	return nil
}

// PhoneMsgSender 发送短信
type PhoneMsgSender struct {
	msg []string
}

// NewEmailMsgSender NewEmailMsgSender
func NewPhoneMsgSender(msg []string) *PhoneMsgSender {
	return &PhoneMsgSender{msg: msg}
}

// Send Send
func (p *PhoneMsgSender) Send(msg string) error {
	// 这里去发送消息
	fmt.Printf("send msg %s\n", msg)
	return nil
}

// INotification 通知接口
type INotification interface {
	Notify(msg string) error
}

// ErrorNotification 错误通知
// 后面可能还有 warning 各种级别
type ErrorNotification struct {
	sender IMsgSender
}

// NewErrorNotification NewErrorNotification
func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

// Notify 发送通知
func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}

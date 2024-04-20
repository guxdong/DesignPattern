package observer

import (
	"fmt"
)

// MsgType 信用卡消息类型
type MsgType int

const (
	// 消费消息
	ConsumeType MsgType = iota
	// 账单消息
	BillType
	// 逾期消息
	ExpireType
)

type Observer interface {
	Name() string
	Update(message string)
}

type ShortMessage struct{}

func (s *ShortMessage) Name() string {
	return "短信"
}

func (s *ShortMessage) Update(message string) {
	fmt.Printf("通过 %s 发送消息：%s。\n", s.Name(), message)
}

type Email struct{}

func (e *Email) Name() string {
	return "邮件"
}

func (e *Email) Update(message string) {
	fmt.Printf("通过 %s 发送消息：%s。\n", e.Name(), message)
}

type Telephone struct{}

func (t *Telephone) Name() string {
	return "电话"
}

func (t *Telephone) Update(message string) {
	fmt.Printf("通过 %s 发送消息：%s。\n", t.Name(), message)
}

type CreditCard struct {
	Holder        string                 // 持卡人姓名
	ConsumeSum    float32                // 消费金额
	ObserverGroup map[MsgType][]Observer // 观察者分组
}

func NewCreditCard(name string) *CreditCard {
	return &CreditCard{
		Holder:        name,
		ObserverGroup: make(map[MsgType][]Observer),
	}
}

// Register 注册观察者
func (c *CreditCard) Register(observer Observer, msgTypes ...MsgType) {
	for _, msgType := range msgTypes {
		c.ObserverGroup[msgType] = append(c.ObserverGroup[msgType], observer)
	}
}

// Unregister 取消注册观察者
func (c *CreditCard) Unregister(observer Observer, msgTypes ...MsgType) {
	for _, msgType := range msgTypes {
		observers, isExist := c.ObserverGroup[msgType]
		if !isExist {
			continue
		}
		c.ObserverGroup[msgType] = c.removeObserver(observer, observers)
	}
}

func (c *CreditCard) removeObserver(toBeRemoved Observer, observers []Observer) []Observer {
	length := len(observers)
	for i, observer := range observers {
		if toBeRemoved.Name() != observer.Name() {
			continue
		}
		observers[length-1], observers[i] = observers[i], observers[length-1]
		return observers[:length-1]
	}
	return observers
}

func (c *CreditCard) Notify(msgType MsgType, message string) {
	for _, observer := range c.ObserverGroup[msgType] {
		observer.Update(message)
	}
}

func (c *CreditCard) Consume(money float32) {
	c.ConsumeSum += money
	message := fmt.Sprintf("尊敬的客户 %s, 您的信用卡消费 %.2f 元", c.Holder, money)
	c.Notify(ConsumeType, message)
}

func (c *CreditCard) SendBill() {
	message := fmt.Sprintf("尊敬的客户 %s, 您本月的账单已出, 共消费 %.2f 元", c.Holder, c.ConsumeSum)
	c.Notify(BillType, message)
}

func (c *CreditCard) Expire() {
	message := fmt.Sprintf("尊敬的客户 %s, 您本月的账单已逾期, 请及时还款, 总额 %.2f 元", c.Holder, c.ConsumeSum)
	c.Notify(ExpireType, message)
}

func Logic() {
	creditCard := NewCreditCard("张三")
	// 日常消费以及账单逾期通过短信通知
	creditCard.Register(new(ShortMessage), ConsumeType, ExpireType)
	// 消费账单以及账单逾期通过邮件通知
	creditCard.Register(new(Email), BillType, ExpireType)
	// 账单逾期通过电话通知
	creditCard.Register(new(Telephone), ExpireType)

	creditCard.Consume(78.23)
	creditCard.Consume(100)
	creditCard.SendBill()
	creditCard.Expire()

	// 账单逾期消息取消邮件通知以及短信通知
	creditCard.Unregister(new(Email), ExpireType)
	creditCard.Unregister(new(ShortMessage), ExpireType)
	creditCard.Consume(64.19)
	creditCard.Expire()
}

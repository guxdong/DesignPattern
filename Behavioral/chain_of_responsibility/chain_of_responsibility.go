package chainofresponsibility

import "fmt"

// Passenger 旅客
type Passenger struct {
	Name                  string // 姓名
	HasBoardingPass       bool   // 是否对应登机牌
	HasLuggage            bool   // 是否有行李
	IsPassIdentityCheck   bool   // 是否通过身份检查
	IsPassSecurityCheck   bool   // 是否通过安检
	IsCompleteForBoarding bool   // 是否完成登机
}

// BoardingProcessor 登机过程各节点处理接口
type BoardingProcessor interface {
	SetSuccessor(successor BoardingProcessor)
	Process(passenager *Passenger)
}

// BaseBoardingProcessor 登机过程处理基类
type BaseBoardingProcessor struct {
	successor BoardingProcessor
}

func (b *BaseBoardingProcessor) SetSuccessor(processor BoardingProcessor) {
	b.successor = processor
}

func (b *BaseBoardingProcessor) Process(passenager *Passenger) {
	if b.successor != nil {
		b.successor.Process(passenager)
	}
}

// BoardingPassProcessor 登机牌
type BoardingPassProcessor struct {
	BaseBoardingProcessor
}

func (b *BoardingPassProcessor) Process(passenger *Passenger) {
	if !passenger.HasBoardingPass {
		fmt.Printf("为旅客%s办理登机牌;\n", passenger.Name)
		passenger.HasBoardingPass = true
	}
	// 成功打印登机牌后进入下一环节进行处理
	b.successor.Process(passenger)
}

type LuggageProcessor struct {
	BaseBoardingProcessor
}

func (l *LuggageProcessor) Process(passenger *Passenger) {
	if !passenger.HasBoardingPass {
		fmt.Printf("旅客%s未打印登机牌, 不能办理行李托运;\n", passenger.Name)
		return
	}
	if !passenger.HasLuggage {
		fmt.Printf("为旅客%s办理行李托运;\n", passenger.Name)
	}
	l.successor.Process(passenger)
}

type IdentityCheckProcessor struct {
	BaseBoardingProcessor
}

func (i *IdentityCheckProcessor) Process(passenger *Passenger) {
	if !passenger.HasBoardingPass {
		fmt.Printf("旅客%s未打印登机牌, 不能办理行李托运;\n", passenger.Name)
		return
	}
	if !passenger.IsPassIdentityCheck {
		fmt.Printf("为旅客%s核实身份信息;\n", passenger.Name)
		passenger.IsPassIdentityCheck = true
	}
	i.successor.Process(passenger)
}

type SecurityCheckProcessor struct {
	BaseBoardingProcessor
}

func (s *SecurityCheckProcessor) Process(passenger *Passenger) {
	if !passenger.HasBoardingPass {
		fmt.Printf("旅客%s未打印登机牌, 不能进行安检;\n", passenger.Name)
		return
	}

	if !passenger.IsPassIdentityCheck {
		fmt.Printf("旅客%s未通过身份检查, 不能进行安检;\n", passenger.Name)
		return
	}
	if !passenger.IsPassSecurityCheck {
		fmt.Printf("为旅客%s进行安检;\n", passenger.Name)
		passenger.IsPassSecurityCheck = true
	}
	s.successor.Process(passenger)
}

type CompleteBoardingProcessor struct {
	BaseBoardingProcessor
}

func (c *CompleteBoardingProcessor) Process(passenger *Passenger) {
	if !passenger.HasBoardingPass ||
		!passenger.IsPassIdentityCheck ||
		!passenger.IsPassSecurityCheck {
		fmt.Printf("旅客%s登机检查过程未完成, 不能登机;\n", passenger.Name)
		return
	}
	passenger.IsCompleteForBoarding = true
	fmt.Printf("旅客%s成功登机;\n", passenger.Name)
}

// 业务逻辑
func Logic() {
	passenger := Passenger{
		Name: "张三",
	}
	completeBoarding := new(CompleteBoardingProcessor)

	securityCheck := new(SecurityCheckProcessor)
	securityCheck.SetSuccessor(completeBoarding)

	identityCheck := new(IdentityCheckProcessor)
	identityCheck.SetSuccessor(securityCheck)

	luggage := new(LuggageProcessor)
	luggage.SetSuccessor(identityCheck)

	boardingPass := new(BoardingPassProcessor)
	boardingPass.SetSuccessor(luggage)

	boardingPass.Process(&passenger)
}

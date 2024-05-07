package visitor

import "fmt"

/* 访问者接口 */
type Visitor interface {
	VisitForPDF(pdf *PDF)
	VisitForPPT(ppt *PPT)
}

/* 具体访问者 */
type ExtractVisitor struct{}

func (e *ExtractVisitor) VisitForPDF(pdf *PDF) {
	fmt.Printf("该 PDF 共 %d 页, ", pdf.page)
	fmt.Println("从 PDF 中提取文本")
}

func (e *ExtractVisitor) VisitForPPT(ppt *PPT) {
	fmt.Printf("该 PPT 使用了 %s 模板, ", ppt.template)
	fmt.Println("从 PPT 中提取文本")
}

type CompressVisitor struct{}

func (c *CompressVisitor) VisitForPDF(pdf *PDF) {
	fmt.Printf("该 PDF 共 %d 页, ", pdf.page)
	fmt.Println("对 PDF 进行压缩")
}

func (c *CompressVisitor) VisitForPPT(ppt *PPT) {
	fmt.Printf("该 PPT 使用了 %s 模板, ", ppt.template)
	fmt.Println("对 PPT 进行压缩")
}

/* 元素接口 */
type Element interface {
	Accept(visitor Visitor)
}

/* 具体元素 */
type PDF struct {
	page int // PDF的页数
}

func NewPDF(page int) *PDF {
	return &PDF{
		page: page,
	}
}

func (p *PDF) Accept(visitor Visitor) {
	visitor.VisitForPDF(p)
}

type PPT struct {
	template string // PPT所使用的模板
}

func NewPPT(template string) *PPT {
	return &PPT{
		template: template,
	}
}

func (p *PPT) Accept(visitor Visitor) {
	visitor.VisitForPPT(p)
}

// Logic 业务逻辑
func Logic() {
	var visitor Visitor
	visitor = new(ExtractVisitor)

	pdf := NewPDF(123)
	ppt := NewPPT("空白")

	pdf.Accept(visitor)
	ppt.Accept(visitor)

	visitor = new(CompressVisitor)

	pdf.Accept(visitor)
	ppt.Accept(visitor)
}

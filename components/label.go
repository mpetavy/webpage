package components

import "github.com/mpetavy/common"

const (
	FOR     = "for"
	CAPTION = "caption"
)

type Label struct {
	For     StringProperty
	Caption StringProperty
}

func (label Label) updateElement(element *common.Element) {
	element.AddAttr(FOR, label.For.HTML())
	element.Text = label.Caption.HTML()
}

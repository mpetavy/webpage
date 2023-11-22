package components

import (
	"fmt"
	"github.com/mpetavy/common"
	"strings"
)

type Page struct {
	Components []*Component
}

func NewPage() (*Page, error) {
	return &Page{}, nil
}

func (page Page) HTML() (string, error) {
	sb := strings.Builder{}
	sb.WriteString("<!DOCTYPE html>\n")

	for _, component := range page.Components {
		html, err := component.HTML()
		if common.Error(err) {
			return "", err
		}
		sb.WriteString(fmt.Sprintf("%s\n", html))
	}

	return sb.String(), nil
}

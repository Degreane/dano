package batch

import "github.com/shopspring/decimal"

type node struct {
	Name string          `yaml:"name"`
	AR   decimal.Decimal `yaml:"ar"`
	IR   decimal.Decimal `yaml:"ir"`
	N    int64           `yaml:"n"`
}

func newNode() *node {
	return new(node)
}

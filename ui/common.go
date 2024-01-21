package ui

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type Item struct {
	name  string
	width int32
}

func NewItem(n string, w int32) Item {
	return Item{
		name:  n,
		width: w,
	}
}

func addColV2(l *vcl.TListView, i []Item) {
	index := l.Columns().Add()
	index.SetCaption("序")
	index.SetWidth(0)
	index.SetAlignment(types.TaCenter)
	index.SetAutoSize(false)

	for _, v := range i {
		col := l.Columns().Add()
		col.SetCaption(v.name)
		col.SetWidth(v.width)
		col.SetAlignment(types.TaCenter)
	}
}

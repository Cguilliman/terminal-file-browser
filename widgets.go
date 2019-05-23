package main

// import (
// 	"fmt"
// 	ui "github.com/gizak/termui"
// )

// type Header struct {
// 	Time   *ui.Par
// 	Count  *ui.Par
// 	Filter *ui.Par
// 	bg     *ui.Par
// }

// func NewHeader() *Header {
// 	return &Header{
// 		Time:  headerPar(2, ""),
// 		Count: headerPar(24, "-"),
// 		Count: headerPar(40, ""),
// 		Count: headerBg(),
// 	}
// }

// func headerBg() *ui.Par {
// 	bg := ui.NewPar("")
// 	bg.X = 1
// 	bg.Height = 1
// 	bg.Border = false
// 	bg.Bg = ui.ThemeAttr("header.bg")
// 	return bg
// }

// func headerPar(x int, s string) *ui.Par {
// 	p := ui.NewPar(fmt.Sprintf(" %s", s))
// 	p.X = x
// 	p.Border = false
// 	p.Height = 1
// 	p.Width = 20
// 	p.Bg = ui.ThemeAttr("header.bg")
// 	p.TextFgColor = ui.ThemeAttr("header.fg")
// 	p.TextBgColor = ui.ThemeAttr("header.bg")
// 	return p
// }

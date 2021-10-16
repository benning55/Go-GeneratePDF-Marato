package main

import (
	"fmt"
	"os"

	"github.com/benning55/fruitfull-pdf/data"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	buildFruitTable(m)

	err := m.OutputFileAndClose("pdfs/div_rhino_fruit.pdf")
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
}

func buildHeading(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/logo.jpeg", props.Rect{
					Center:  true,
					Percent: 75,
				})

				if err != nil {
					fmt.Println("Image file was not loaded 😱 - ", err)
				}

			})
		})
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Test Document", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurpleColor(),
			})
		})
	})
}

func buildFruitTable(m pdf.Maroto) {
	headings := getHeadings()
	contents := data.FruitList(20)
	// contents := [][]string{{"Apple", "Red and juicy", "2.00"}, {"Orange", "Orange and juicy", "3.00"}}

	m.SetBackgroundColor(getTealColor())

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Products", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.Row(20, func() {
		m.Col(12, func() {
			m.TableList(headings, contents, props.TableList{
				HeaderProp: props.TableListContent{
					Size:      9,
					GridSizes: []uint{3, 7, 2},
				},
				ContentProp: props.TableListContent{
					Size:      7,
					GridSizes: []uint{3, 7, 2},
				},
				Align:              consts.Left,
				HeaderContentSpace: 3,
				Line:               false,
			})
		})
	})

	// m.Row(20, func() {
	// 	m.ColSpace(7)
	// 	m.Col(2, func() {
	// 		m.Text("Total:", props.Text{
	// 			Top:   5,
	// 			Style: consts.Bold,
	// 			Size:  8,
	// 			Align: consts.Right,
	// 		})
	// 	})
	// 	m.Col(3, func() {
	// 		m.Text("$ XXXX.00", props.Text{
	// 			Top:   5,
	// 			Style: consts.Bold,
	// 			Size:  8,
	// 			Align: consts.Center,
	// 		})
	// 	})
	// })
}

func getHeadings() []string {
	return []string{"Fruit", "Description", "Price"}
}

func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}

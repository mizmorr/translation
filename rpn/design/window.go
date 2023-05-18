package design

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func get_orig(num int) string {
	file, err := os.ReadFile("data/sample" + fmt.Sprint(num) + ".txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}
func anl(num int) string {
	file, err := os.ReadFile("data/data" + fmt.Sprint(num) + "_analz.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}
func rpn(num int) string {
	file, err := os.ReadFile("data/sample" + fmt.Sprint(num) + "_rpn.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}
func Run() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Syntax analyzer")
	label_rpn := widget.NewLabel("")
	myWindow.Resize(fyne.NewSize(350, 500))
	myWindow.SetFixedSize(true)
	tabs := container.NewDocTabs(
		container.NewTabItem("1", widget.NewLabel(get_orig(1))),
		container.NewTabItem("2", widget.NewLabel(get_orig(2))),
		container.NewTabItem("3", widget.NewLabel(get_orig(3))),
		container.NewTabItem("4", widget.NewLabel(get_orig(4))),
	)
	label_analyzed := widget.NewLabel("")
	button_analyzed := widget.NewButton("Analysis", func() {
		label_analyzed.SetText(anl(tabs.SelectedIndex() + 1))
	})
	button_rpn := widget.NewButton("RPN", func() {
		label_rpn.SetText(rpn(tabs.SelectedIndex() + 1))

	})
	container_analyzed := container.NewGridWithRows(2, container.NewPadded(container.NewScroll(label_analyzed)), container.NewCenter(button_analyzed))
	container_ := container.NewGridWithRows(2, container.NewVSplit(tabs, layout.NewSpacer()))
	container_rpn := container.NewGridWithRows(2, container.NewPadded(container.NewScroll(label_rpn)), container.NewCenter(button_rpn))
	group := container.NewGridWithColumns(3, container_, container_analyzed, container_rpn)
	myWindow.SetContent(group)
	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow.ShowAndRun()

}

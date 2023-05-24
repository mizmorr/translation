package design

import (
	"fmt"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func get_orig(num int) string {
	file, err := os.ReadFile("../rpn/data/sample" + fmt.Sprint(num) + ".txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}
func anl(num int) string {
	file, err := os.ReadFile("../rpn/data/data" + fmt.Sprint(num) + "_analz.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}
func rpn(num int) string {
	file, err := os.ReadFile("../rpn/data/sample" + fmt.Sprint(num) + "_rpn.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}
func pyth(num int) string {
	file, err := os.ReadFile("data/task" + fmt.Sprint(num) + "/result.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}
func Show() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Translator")
	label_rpn := widget.NewLabel("")
	myWindow.Resize(fyne.NewSize(1200, 800))
	myWindow.SetFixedSize(true)
	tabs := container.NewDocTabs(
		container.NewTabItem("1", widget.NewLabel(get_orig(1))),
		container.NewTabItem("2", widget.NewLabel(get_orig(2))),
		container.NewTabItem("3", widget.NewLabel(get_orig(3))),
		container.NewTabItem("4", widget.NewLabel(get_orig(4))),
	)
	green := color.NRGBA{R: 133, G: 133, B: 133, A: 255}
	label := canvas.NewText("Транслятор из Java в Python", green)
	label.Alignment = fyne.TextAlignTrailing
	label.TextSize = 26
	label.TextStyle.Bold = true
	label_analyzed := widget.NewLabel("")
	button_to_rpn := widget.NewButtonWithIcon("Провести лексический анализ и сформировать ОПЗ", theme.ConfirmIcon(), func() {
		label_rpn.SetText(rpn(tabs.SelectedIndex() + 1))
		label_analyzed.SetText(anl(tabs.SelectedIndex() + 1))

	})
	label_python := widget.NewLabel("")
	button_to_python := widget.NewButtonWithIcon("Перевести ОПЗ в код ЯП Python", theme.ConfirmIcon(), func() {
		label_python.SetText(pyth(tabs.SelectedIndex() + 1))

	})
	container_python := container.NewGridWithRows(3, container.NewScroll(label_python), layout.NewSpacer(), container.NewPadded(container.NewGridWithRows(3, button_to_python)))
	container_to_rpn := container.NewGridWithRows(3, container.NewScroll(label_analyzed), container.NewScroll(label_rpn), container.NewPadded(container.NewGridWithRows(3, button_to_rpn)))
	container_ := container.NewGridWithRows(2, container.NewVSplit(tabs, label))
	group := container.NewGridWithColumns(3, container_, container_to_rpn, container_python)
	myWindow.SetContent(group)
	myWindow.CenterOnScreen()
	myWindow.SetIcon(theme.ComputerIcon())
	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow.ShowAndRun()

}

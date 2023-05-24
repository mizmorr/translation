package design

import (
	"fmt"
	"image/color"
	"os"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Show() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Analyzer")
	myWindow.Resize(fyne.NewSize(600, 600))
	myWindow.SetFixedSize(true)
	green := color.NRGBA{R: 133, G: 133, B: 133, A: 255}
	label := canvas.NewText("Синтаксический анализатор", green)
	label.Alignment = fyne.TextAlignCenter
	label.TextSize = 26
	label.TextStyle.Bold = true
	entry := widget.NewMultiLineEntry()
	progress := widget.NewProgressBar()

	button := widget.NewButtonWithIcon("Провести анализ", theme.DocumentIcon(), func() {
		go func() {
			for i := 0.0; i <= 1.0; i += 0.1 {
				time.Sleep(time.Millisecond * 10)
				progress.SetValue(i)
			}
		}()
		java := entry.Text
		f, _ := os.Create("script/java.txt")
		f.Write([]byte(java))
		f.Close()
		cmd := exec.Command("python", "script/script.py")
		cmd.Run()
		res, _ := os.ReadFile("design/result.txt")
		if string(res) == "-1" {
			dialog.ShowInformation("", "Код корректен\nОшибок не найдено.", myWindow)
		} else {
			dialog.ShowInformation("", fmt.Sprint("Найдена синтаксическая ошибка в ", string(res), " строке.\n Проверьте код повторно!"), myWindow)
		}
	})
	content := container.NewGridWithColumns(2, container.NewVSplit(container.NewGridWithRows(4, label, layout.NewSpacer(), container.NewGridWithRows(2, progress)), container.NewPadded(button)), container.NewPadded(entry))
	myWindow.SetContent(content)
	myWindow.CenterOnScreen()
	myWindow.SetIcon(theme.ComputerIcon())
	myApp.Settings().SetTheme(theme.LightTheme())
	myWindow.ShowAndRun()
}

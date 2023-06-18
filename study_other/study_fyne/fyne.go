package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

/*
fyne包结构划分
fyne将功能划分到多个子包中：
fyne.io/fyne：提供所有fyne应用程序代码共用的基础定义，包括数据类型和接口；
fyne.io/fyne/app：提供创建应用程序的 API；
fyne.io/fyne/canvas：提供Fyne使用的绘制 API；
fyne.io/fyne/dialog：提供对话框组件；
fyne.io/fyne/layout：提供多种界面布局；
fyne.io/fyne/widget：提供多种组件，fyne所有的窗体控件和交互元素都在这个子包中。
*/

func main() {
	a := app.New()
	w := a.NewWindow("测试一下")
	w.Resize(fyne.NewSize(200, 200))

	hello := widget.NewLabel("say hi")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}

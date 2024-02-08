package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Run() {
	myApp := app.New()
	//title := fmt.Sprintf("Zurok Wallet (%s)", getHostname())
	myWindow := myApp.NewWindow("Zurok Wallet")

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Wallet", theme.FolderOpenIcon(), createWalletContent()),
		container.NewTabItem("Send", createSendContent(myWindow)),
		container.NewTabItem("Receive", createReceiveContent()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.ShowAndRun()
}

func createReceiveContent() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Receive Content"),
		// Add receive-related widgets here
	)
}

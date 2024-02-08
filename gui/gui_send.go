package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func createSendContent(parent fyne.Window) fyne.CanvasObject {
	fromWallet := widget.NewEntry()
	fromWallet.SetPlaceHolder("Wallet to send from...")

	toWallet := widget.NewEntry()
	toWallet.SetPlaceHolder("Wallet to send to...")

	content := container.NewVBox(fromWallet, toWallet, widget.NewButton("Send", func() {
		//log.Println("Content was:", fromWallet.Text)
		confirmSend(parent)
	}))

	return content
}

func confirmSend(window fyne.Window) {
	d := dialog.NewConfirm("Confirmation", "Are you sure you want to send this?", func(response bool) {
		if response {
			fmt.Println("Sending...")
		} else {
			fmt.Println("Cancelled Send")
		}
	}, window)

	d.SetDismissText("Cancel")
	d.SetConfirmText("Continue")
	d.Show()
}

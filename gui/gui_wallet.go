package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func createWalletContent() fyne.CanvasObject {
	node := "3000"

	// Get wallets with balances
	addresses := listAddresses(node)
	var balances []string

	for _, addr := range addresses {
		newString := fmt.Sprintf("%s: %d", addr, getBalance(addr, node))
		balances = append(balances, newString)
	}

	data := binding.BindStringList(
		&balances,
	)

	// Create list to display balances
	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	// Create a new wallet button (needs a confirmation button)
	newWalletButton := widget.NewButton("Create New Wallet", func() {
		newWallet := createWallet(node)
		listVal := fmt.Sprintf("%s: 0", newWallet)
		data.Append(listVal)
	})

	content := container.NewBorder(nil, newWalletButton, nil, nil, list)
	return content
}

package wallet

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type PageWrapper struct {
	Pages *tview.Pages
}

var currentWallet *Wallet

func Run() {

	app := tview.NewApplication()

	pages := tview.NewPages()

	pwrap := PageWrapper{
		Pages: pages,
	}

	menuFlex := buildMenuPage(app, pwrap)
	createFlex := buildWalletCreationPage(app, pwrap)
	noWalletFlex := buildNoWalletPage(app, pwrap)

	pages.AddPage("menu", menuFlex, true, true)
	pages.AddPage("create", createFlex, true, false)
	pages.AddPage("nowallet", noWalletFlex, true, false)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}

func (pw PageWrapper) Switch(s string) {
	pw.Pages.SwitchToPage(s)
}

func buildMenuPage(app *tview.Application, pwrap PageWrapper) *tview.Flex {
	menuList := tview.NewList().
		AddItem("Create New Alpenglow Wallet", "Generate Keys and Save Wallet to Disk", 'c', func() { pwrap.Switch("create") }).
		AddItem("Open Current Wallet", "Open Wallet that is Currently In Use", 'o', func() {
			if currentWallet != nil {
				pwrap.Switch("account")
			} else {
				pwrap.Switch("nowallet")
			}
		}).
		AddItem("Load Wallet", "Load Wallet from File", 'l', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	menuList.SetBorder(true).SetTitle(" Alpenglow Wallet ").SetBackgroundColor(tcell.ColorDefault)

	menuFlex := tview.NewFlex().
		AddItem(menuList, 0, 1, true)

	return menuFlex
}

func buildWalletCreationPage(app *tview.Application, pwrap PageWrapper) *tview.Flex {
	createList := tview.NewList().
		AddItem("Confirm Wallet Creation", "This will generate a new ED25519 Key Pair", 'c', func() { keyPairInitialization(app, pwrap) }).
		AddItem("Back", "Return to Menu", 'b', func() { pwrap.Switch("menu") })

	createList.SetBorder(true).SetTitle(" Create Wallet ").SetBackgroundColor(tcell.ColorDefault)

	createFlex := tview.NewFlex().
		AddItem(createList, 0, 1, true)

	return createFlex
}

func buildAccountPage(app *tview.Application, pwrap PageWrapper) *tview.Flex {
	text := tview.NewTextView()
	text.SetBorder(true).SetTitle(" Account View ").SetBackgroundColor(tcell.ColorDefault)
	text.SetTextAlign(tview.AlignCenter)
	text.SetText(fmt.Sprintf("Account Name: %s\nPublic Key: 0x%x", currentWallet.Name, *currentWallet.PublicKey))

	accountList := tview.NewList().
		AddItem("New Transaction", "Create a New Transaction on Alpine", 'n', nil).
		AddItem("Save Wallet", "Save Alpenglow Wallet", 's', nil).
		AddItem("Back", "Return to Menu", 'b', func() { pwrap.Switch("menu") })

	accountList.SetBorder(true).SetTitle(" Account Menu ").SetBackgroundColor(tcell.ColorDefault)

	accountFlex := tview.NewFlex().
		AddItem(text, 0, 3, false).
		AddItem(accountList, 0, 1, true)

	accountFlex.InputHandler()

	return accountFlex
}

func buildNoWalletPage(app *tview.Application, pwrap PageWrapper) *tview.Pages {
	noWalletModal := tview.NewModal().
		SetText("No Wallet Currently in Use").
		AddButtons([]string{"Return to Menu", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Quit" {
				app.Stop()
			} else if buttonLabel == "Return to Menu" {
				pwrap.Switch("menu")
			}
		})

	noWalletModal.SetBorder(true).SetTitle(" No Wallet ").SetBackgroundColor(tcell.ColorDefault)

	background := tview.NewBox().
		SetBackgroundColor(tcell.ColorDefault)

	pages := tview.NewPages().
		AddPage("background", background, true, true).
		AddPage("modal", noWalletModal, true, true)

	return pages

}

func keyPairInitialization(app *tview.Application, pwrap PageWrapper) {
	currentWallet = CreateNewWallet("test")
	pwrap.Pages.AddAndSwitchToPage("account", buildAccountPage(app, pwrap), true)
}

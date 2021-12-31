package wallet

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type PageWrapper struct {
	Pages *tview.Pages
}

func Run() {
	pubkey, privkey := GenerateEd25519KeyPair()
	fmt.Printf("%x\n", *privkey)
	fmt.Printf("%x\n", *pubkey)

	app := tview.NewApplication()

	pages := tview.NewPages()

	pwrap := PageWrapper{
		Pages: pages,
	}

	menuList := tview.NewList().
		AddItem("Create New Alpenglow Wallet", "Generate Keys and Save Wallet to Disk", 'c', func() { pwrap.Switch("create") }).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})

	menuList.SetBorder(true).SetTitle(" Alpenglow Wallet ").SetBackgroundColor(tcell.ColorDefault)

	menuFlex := tview.NewFlex().
		AddItem(menuList, 0, 1, true)

	createList := tview.NewList().
		AddItem("Back", "Return to Menu", 'b', func() { pwrap.Switch("menu") })

	createList.SetBorder(true).SetTitle(" Create Wallet ").SetBackgroundColor(tcell.ColorDefault)

	createFlex := tview.NewFlex().
		AddItem(createList, 0, 1, true)

	pages.AddPage("menu", menuFlex, true, true)
	pages.AddPage("create", createFlex, true, false)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		panic(err)
	}
}

func (pw PageWrapper) Switch(s string) {
	pw.Pages.SwitchToPage(s)
}

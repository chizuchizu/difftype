package main

import (
	"github.com/rivo/tview"
  "github.com/gdamore/tcell/v2"
  "strings"
  // "fmt"
)

func main() {
  app := tview.NewApplication()
  text := `
  ╔═══╗─╔═╗╔═╦════╗
  ╚╗╔╗║─║╔╝║╔╣╔╗╔╗║
  ─║║║╠╦╝╚╦╝╚╬╝║║╠╣─╔╦══╦══╗
  ─║║║╠╬╗╔╩╗╔╝─║║║║─║║╔╗║║═╣
  ╔╝╚╝║║║║─║║──║║║╚═╝║╚╝║║═╣
  ╚═══╩╝╚╝─╚╝──╚╝╚═╗╔╣╔═╩══╝
  ───────────────╔═╝║║║
  ───────────────╚══╝╚╝

  ██████╗░██╗███████╗███████╗████████╗██╗░░░██╗██████╗░███████╗
  ██╔══██╗██║██╔════╝██╔════╝╚══██╔══╝╚██╗░██╔╝██╔══██╗██╔════╝
  ██║░░██║██║█████╗░░█████╗░░░░░██║░░░░╚████╔╝░██████╔╝█████╗░░
  ██║░░██║██║██╔══╝░░██╔══╝░░░░░██║░░░░░╚██╔╝░░██╔═══╝░██╔══╝░░
  ██████╔╝██║██║░░░░░██║░░░░░░░░██║░░░░░░██║░░░██║░░░░░███████╗
  ╚═════╝░╚═╝╚═╝░░░░░╚═╝░░░░░░░░╚═╝░░░░░░╚═╝░░░╚═╝░░░░░╚══════╝
  `
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
      AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
      AddItem(nil, 0, 1, false).
      AddItem(p, height, 1, true).
      AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

  background := tview.NewTextView().
  SetTextColor(tcell.ColorOrange).
  SetText(strings.Repeat(text, 100))

  box := tview.NewBox().
  SetBorder(true).
  SetTitle("Centered Box")

  pages := tview.NewPages().
  AddPage("background", background, true, true).
  AddPage("box", modal(box, 40, 10), true, true)

  if err := app.SetRoot(pages, true).Run(); err != nil {
    panic(err)
  }
  
}

package main

import (
	"github.com/andlabs/ui"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	respChan := make(chan QR)

	err := ui.Main(func() {

		var prox []string
		var uniqueProxies []string

		//Creating elements

		//Creating input field
		input := ui.NewEntry()

		//Creating buttons
		button := ui.NewButton("Open File")
		button1 := ui.NewButton("Start Checking")
		button3 := ui.NewButton("Exit")

		bt := ui.NewCombobox()

		bt.Append("HTTP")
		bt.Append("SOCKS")

		//Creating labels
		greeting := ui.NewLabel("")
		res := ui.NewLabel("")

		//Creating progress bar
		pb := ui.NewProgressBar()

		//Creating box
		box := ui.NewVerticalBox()

		//Appending elements to box
		box.Append(ui.NewLabel("Select proxy type"), false)
		box.Append(bt, false)
		box.Append(ui.NewLabel("\n"), false)
		box.Append(ui.NewLabel("Path to file with proxies"), false)
		box.Append(input, false)
		box.Append(ui.NewLabel("\n"), false)
		box.Append(button, false)
		box.Append(button1, false)
		box.Append(greeting, false)
		box.Append(ui.NewLabel("Progress"), false)
		box.Append(pb, false)
		box.Append(ui.NewLabel("\n"), false)
		box.Append(res, false)
		box.Append(ui.NewLabel("\n"), false)
		box.Append(button3, false)

		//Creating window
		window := ui.NewWindow("ProxyChecker by trigun117", 500, 200, false)
		window.SetMargined(true)
		window.SetChild(box)

		button1.Disable()
		bt.SetSelected(1)
		button3.Hide()

		//Button click event
		button.OnClicked(func(*ui.Button) {

			//Open file
			input.SetText(ui.OpenFile(window))

			switch bt.Selected() {
			case 0:
				prox = readFromFile(input.Text(), 0)
				uniqueProxies = unique(prox)
			case 1:
				prox = readFromFile(input.Text(), 1)
				uniqueProxies = unique(prox)
			}

			button1.Enable()
		})

		//Button click event
		button1.OnClicked(func(*ui.Button) {

			//Updating progress bar value
			pb.SetValue(20)

			switch bt.Selected() {
			case 0:
				realIP := getRealIP()
				for _, proxy := range uniqueProxies {
					go checkProxyHTTP(proxy, respChan, realIP)
				}
			case 1:
				for _, proxy := range uniqueProxies {
					go checkProxySOCKS(proxy, respChan)
				}
			}

			//Updating progress bar value
			pb.SetValue(50)

			for range uniqueProxies {
				r := <-respChan
				if r.Res {
					writeToFile(r.Addr)
				}
			}

			//Updating progress bar value
			pb.SetValue(100)

			res.SetText("Finish, check your proxies in live-proxies.txt")

			button3.Show()

			//Button click event
			button3.OnClicked(func(*ui.Button) {

				//Close window if button clicked
				ui.Quit()
			})
		})

		//Event when window closing
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
	})

	if err != nil {
		panic(err)
	}
}

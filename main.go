package main

import (
	"github.com/andlabs/ui"
	"os"
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

		//Creating labels
		greeting := ui.NewLabel("")
		res := ui.NewLabel("")

		//Creating progress bar
		pb := ui.NewProgressBar()

		//Creating box
		box := ui.NewVerticalBox()

		//Appending elements to box
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
		window := ui.NewWindow("ProxyChecker", 150, 150, false)
		window.SetMargined(true)
		window.SetChild(box)

		button3.Hide()

		//Button click event
		button.OnClicked(func(*ui.Button) {

			//Open file
			input.SetText(ui.OpenFile(window))

			prox = readFromFile(input.Text())
			uniqueProxies = unique(prox)
		})

		//Button click event
		button1.OnClicked(func(*ui.Button) {

			//Updating progress bar value
			pb.SetValue(30)

			realIP := getRealIP()
			for _, proxy := range uniqueProxies {
				go checkProxy(proxy, respChan, realIP)
			}

			//Updating progress bar value
			pb.SetValue(80)

			os.Create(`live-proxies.txt`)
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

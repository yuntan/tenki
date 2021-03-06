package main

import (
	ui "github.com/gizak/termui"
)

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	getWeatherInfo()
	processData()
	setupView()
	waitLoop()
}

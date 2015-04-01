package main

import (
	ui "github.com/gizak/termui"
	tm "github.com/nsf/termbox-go"

	"fmt"
	"time"
)

const (
	BAR_WIDTH  = 5
	BAR_GAP    = 3
	SPLINE_NUM = BAR_WIDTH + BAR_GAP
)

var (
	view []ui.Bufferer

	cityName      string
	temp          float64
	weather       string
	chartLabels   []string
	lineChartData []float64
	barChartData  []int
	dataTime      time.Time
)

func processData() {
	tempData := []float64{14.8, 16.1, 17.3, 18.6, 16.1, 9.0, 7.1, 5.1}

	cityName = "Osaka-shi, JP"
	temp = 14.9
	weather = "Rain"
	chartLabels = []string{"01  6", "01  9", "01 12", "01 15", "01 18", "01 21", "02  0", "02  3"}
	lineChartData = make([]float64, len(chartLabels)*SPLINE_NUM)
	for i := range tempData {
		for j := 0; j < SPLINE_NUM; j++ {
			lineChartData[i*SPLINE_NUM+j] = tempData[i]
		}
	}
	barChartData = []int{5, 2, 1, 0, 1, 0, 0, 0}
	dataTime, _ = time.Parse("2006.01.02 15:04", "2015.04.01 06:41")
}

func setupView() {
	ui.SetTheme(ui.ColorScheme{
		BlockBg:           ui.ColorWhite,
		HasBorder:         false,
		BorderFg:          ui.ColorBlue,
		BorderBg:          ui.ColorWhite,
		BorderLabelTextBg: ui.ColorWhite,
		BorderLabelTextFg: ui.ColorGreen,
		ParTextBg:         ui.ColorWhite,
		ParTextFg:         ui.ColorBlack,
		LineChartLine:     ui.ColorRed | ui.AttrBold,
		LineChartAxes:     ui.ColorBlack,
		LineChartText:     ui.ColorRed,
		BarChartBar:       ui.ColorGreen,
		BarChartNum:       ui.ColorWhite,
		BarChartText:      ui.ColorGreen,
	})

	block := ui.NewBlock()
	block.Width = 100
	block.Height = 20

	cityLabel := ui.NewPar(cityName)
	cityLabel.X = 4
	cityLabel.Y = 2
	cityLabel.Width = 15
	cityLabel.TextFgColor = ui.ColorMagenta | ui.AttrBold | ui.AttrReverse

	weatherLabel := ui.NewPar(weather)
	weatherLabel.X = 4
	weatherLabel.Y = 4
	weatherLabel.Width = 10
	weatherLabel.TextFgColor = ui.ColorBlue | ui.AttrBold

	tempLabel := ui.NewPar(fmt.Sprintf("%4.1fC", temp))
	tempLabel.X = 15
	tempLabel.Y = 4
	tempLabel.Width = 10
	tempLabel.TextFgColor = ui.ColorMagenta | ui.AttrBold | ui.AttrUnderline

	infoPar := ui.NewPar(fmt.Sprintf("Max: %13.1fC\nMin: %13.1fC\nWind: %10.2fm/s\nClouds: %10d%%\nPressure: %6.2fhpa",
		18.8, 7.2, 1.23, 92, 997.31))
	infoPar.X = 4
	infoPar.Y = 6
	infoPar.Width = 30
	infoPar.Height = 5

	timeLabel := ui.NewPar(dataTime.Format("get at 01/02 15:04"))
	timeLabel.X = 4
	timeLabel.Y = 19
	timeLabel.Width = 20
	timeLabel.Height = 1

	titleLabel := ui.NewPar("Temperature and Precipitation")
	titleLabel.X = 50
	titleLabel.Y = 3
	titleLabel.Width = 30

	tempChart := ui.NewLineChart()
	tempChart.X = 26
	tempChart.Y = 4
	tempChart.Width = 8 * len(chartLabels)
	tempChart.Height = 15
	tempChart.Mode = "dot"
	//	tempChart.DotStyle = '-'
	tempChart.Data = lineChartData
	tempChart.DataLabels = []string{""}

	rainChart := ui.NewBarChart()
	rainChart.X = 33
	rainChart.Y = 14
	rainChart.Width = 8 * len(chartLabels)
	rainChart.Height = 5
	rainChart.Data = barChartData
	rainChart.DataLabels = chartLabels
	rainChart.BarWidth = BAR_WIDTH
	rainChart.BarGap = BAR_GAP

	view = []ui.Bufferer{
		block,
		cityLabel,
		weatherLabel,
		tempLabel,
		infoPar,
		timeLabel,
		titleLabel,
		tempChart,
		rainChart,
	}
}

func waitLoop() {
	event := make(chan tm.Event)
	go func() {
		for {
			event <- tm.PollEvent()
		}
	}()

	for {
		select {
		case e := <-event:
			if e.Type == tm.EventKey && e.Ch == 'q' {
				return
			}
		default:
			ui.Render(view...)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

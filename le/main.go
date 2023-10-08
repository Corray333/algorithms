// package main

// import (
// 	"net/http"
// 	"os/exec"
// )

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello, world!"))
// 	})
// 	exec.Command("cmd", "/c start http://localhost:8080").Start()
// 	http.ListenAndServe("127.0.0.1:8080", nil)
// }

package main

import (
	"net/http"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

func f(x float64) float64 {
	return x * x
}
func gen(a, b, step float64) []float64 {
	res := []float64{}
	for i := a; i <= b; i += step {
		res = append(res, i)
	}
	return res
}

// generate random data for line chart
func generateLineItems() []opts.LineData {
	items := []opts.LineData{}
	for i := 1.0; i < 16; i += 1 {
		items = append(items, opts.LineData{Value: f(i)})
	}
	items = append(items, opts.LineData{Value: f(16)})
	return items
}

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 1.0; i < 16; i += 1 {
		items = append(items, opts.BarData{Value: f(i)})
	}
	items = append(items, opts.BarData{Value: f(16)})
	return items
}

func httpserver(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))
	bar.SetGlobalOptions(
		charts.WithColorsOpts(
			opts.Colors{"white"},
		),
	)

	// Put data into instance
	line.SetXAxis(gen(1, 16, 1)).
		AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: false}))
	bar.SetXAxis(gen(1, 16, 1)).
		AddSeries("Category A", generateBarItems())
	line.Overlap(bar)
	line.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	exec.Command("cmd", "/c start http://localhost:8081").Start()
	http.ListenAndServe("127.0.0.1:8081", nil)
}

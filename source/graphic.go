package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func MakeGraphic(p *plot.Plot, min *float64, max *float64, step *float64, mpoints []float64, name string) {
	var points plotter.XYs

	for i := 0; i < len(mpoints); i += 2 {
		points = append(points, plotter.XY{
			X: mpoints[i],
			Y: mpoints[i+1],
		})
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		log.Fatal(err)
	}

	p.Add(line)
	p.Add(plotter.NewGrid())

	err = p.Save(8*vg.Inch, 6*vg.Inch, name)
	if err != nil {
		log.Fatal(err)
	}
}

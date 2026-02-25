package main

import (
	"flag"

	"gonum.org/v1/plot"
)

func main() {

	formula := flag.String("f", "x^2", "formula")
	min := flag.Float64("min", -10, "min x")
	max := flag.Float64("max", 10, "max x")
	step := flag.Float64("step", 0.1, "step")
	output := flag.String("out", "graph.png", "output file")

	flag.Parse()

	tokens := MakeTokens(*formula)
	parser := &MathParser{}

	p := plot.New()
	p.Title.Text = *formula
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	points := []float64{}

	for x := *min; x <= *max; x += *step {
		y := parser.Eval(tokens, x)
		points = append(points, x)
		points = append(points, y)
	}

	MakeGraphic(p, min, max, step, points, *output)
}

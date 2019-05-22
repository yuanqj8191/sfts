package cum

import "math"

type CumAPI interface {
	App(val float64) float64
}

type CumSum struct {
	sum float64
}

type CumPrd struct {
	prd float64
}

type CumMax struct {
	max float64
}

type CumMin struct {
	min float64
}

type CumCnt struct {
	cnt float64
}

type CumAvg struct {
	sum, cnt float64
}

func NewCumSum() *CumSum {
	return &CumSum{}
}

func (cum *CumSum) App(val float64) (sum float64) {
	cum.sum += val
	return cum.sum
}

func NewCumPrd() *CumPrd {
	return &CumPrd{prd: 1}
}

func (cum *CumPrd) App(val float64) (prd float64) {
	cum.prd *= val
	return cum.prd
}

func NewCumMax() *CumMax {
	return &CumMax{max: math.Inf(-1)}
}

func (cum *CumMax) App(val float64) (max float64) {
	if cum.max < val {
		cum.max = val
	}
	return cum.max
}

func NewCumMin() *CumMin {
	return &CumMin{min: math.Inf(1)}
}

func (cum *CumMin) App(val float64) (min float64) {
	if cum.min > val {
		cum.min = val
	}
	return cum.min
}

func NewCumCnt() *CumCnt {
	return &CumCnt{}
}

func (cum *CumCnt) App(val float64) (cnt float64) {
	cum.cnt++
	return cum.cnt
}

func NewCumAvg() *CumAvg {
	return &CumAvg{}
}

func (cum *CumAvg) App(val float64) (avg float64) {
	cum.cnt++
	cum.sum += val
	return cum.sum / cum.cnt
}
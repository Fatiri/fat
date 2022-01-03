package indikators

type Rsi struct {
	n     int64
	up    *Ema
	dn    *Ema
	prevC float64
	sz    int64
}

func NewRsi(n int64) *Rsi {
	k := 1.0 / float64(n)
	return &Rsi{
		n:     n,
		up:    NewEma(n, k),
		dn:    NewEma(n, k),
		prevC: 0,
		sz:    0,
	}
}

func (r *Rsi) Update(v float64) float64 {
	r.sz++

	chg := v - r.prevC
	r.prevC = v

	if r.sz == 1 {
		return 0
	}

	var up, dn float64
	if chg > 0 {
		up = r.up.Update(chg)
		dn = r.dn.Update(0)
	} else {
		up = r.up.Update(0)
		dn = r.dn.Update(-chg)
	}

	if r.sz <= r.n {
		return 0
	}

	sum := up + dn
	if almostZero(sum) {
		return 0
	}

	return up / sum * 100.0
}

func (r *Rsi) InitPeriod() int64 {
	return r.n
}

func (r *Rsi) Valid() bool {
	return r.sz > r.InitPeriod()
}
func RsiArr(in []float64, n int64) []float64 {
	out := make([]float64, len(in))

	r := NewRsi(n)
	for i, v := range in {
		out[i] = r.Update(v)
	}

	return out
}

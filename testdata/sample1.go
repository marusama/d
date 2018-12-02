package main

import (
	"time"

	"github.com/marusama/d"
)

func main() {
	a := 123
	b := "hello world"
	c := 3.1415926
	d_ := func(n int) bool { return n > 0 }(1)
	e := []int{1, 2, 3}
	f := []byte("goodbye world")
	g := e[1:]
	h := time.Now()
	i := &h

	d.D(a, b, c, d_, e, f, g, h, i)

	s := struct {
		id    int
		i     int
		f     float64
		t     *time.Time
		t2    time.Time
		tNil  *time.Time
		tZero time.Time
		tUTC  time.Time
	}{
		id:   1,
		i:    2,
		f:    3.4,
		t:    i,
		t2:   h,
		tUTC: h.UTC(),
	}
	d.D(s)

	ts := []time.Time{h, time.Now()}
	d.D(ts)

	ts2 := []*time.Time{i}
	d.D(ts2)

	tmap := map[int]time.Time{
		1: h,
		2: time.Now(),
	}
	d.D(tmap)

	tmap2 := map[int]*time.Time{
		1: i,
	}
	d.D(tmap2)
}

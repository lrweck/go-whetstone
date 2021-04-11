package main

import (
	"fmt"
	"math"
	"time"
)

var (
	t1, t2, t3 float64
	j, k, l    int
	e1         []float64
)

func main() {

	var isave int
	var n2, n3, n4, n6, n7, n8, n9, n11 int
	var inner, outer, kount, npass, max_pass int

	var x, y, z float64
	var whet_save, kilowhet float32
	var error, whet_err, percent_err float32
	var begin_time, end_time, dif_time float32
	var dif_save float32

	kilowhet = 0
	whet_save = 0

	e1 = make([]float64, 5)

	npass = 0
	max_pass = 2

	inner = 3000
	outer = 1000

	for npass < max_pass {
		fmt.Printf("Pass number: %d\n", npass)

		kount = 0
		begin_time = seconds()

		for kount < outer {
			t1 = 0.49995
			t2 = 0.50025
			t3 = 2.0

			isave = inner

			n2 = 12 * inner
			n3 = 14 * inner
			n4 = 345 * inner
			n6 = 210 * inner
			n7 = 32 * inner
			n8 = 899 * inner
			n9 = 616 * inner
			n11 = 93 * inner

			e1[1] = 1.0
			e1[2] = -1.0
			e1[3] = -1.0
			e1[4] = -1.0

			// Loop 2
			for i := 1; i <= n2; i++ {
				e1[1] = (e1[1] + e1[2] + e1[3] - e1[4]) * t1
				e1[2] = (e1[1] + e1[2] - e1[3] + e1[4]) * t1
				e1[3] = (e1[1] - e1[2] + e1[3] + e1[4]) * t1
				e1[4] = (-e1[1] + e1[2] + e1[3] + e1[4]) * t1
			}

			// Loop 3
			for i := 1; i <= n3; i++ {
				sub1(e1)
			}

			// Loop 4
			j = 1
			for i := 1; i <= n4; i++ {
				if j-1 != 0 {
					j = 2
				} else {
					j = 3
				}
				if j-2 != 0 {
					j = 1
				} else {
					j = 0
				}
				if j-1 != 0 {
					j = 1
				} else {
					j = 0
				}

			}

			// Loop 6
			j = 1
			k = 2
			l = 3
			for i := 1; i <= n6; i++ {
				j = j * (k - j) * (l - k)
				k = l*k - (l-j)*k
				l = (l - k) * (k + j)
				e1[l-1] = float64(j + k + l)
				e1[k-1] = float64(j * k * l)
			}

			// Loop 7
			x = 0.5
			y = 0.5
			for i := 1; i <= n7; i++ {
				x = t1 * math.Atan(t3*math.Sin(x)*math.Cos(x)/(math.Cos(x+y)+math.Cos(x-y)-1.0))
				y = t1 * math.Atan(t3*math.Sin(y)*math.Cos(y)/(math.Cos(x+y)+math.Cos(x-y)-1.0))
			}

			// Loop 8
			x = 1.0
			y = 1.0
			z = 1.0
			for i := 1; i <= n8; i++ {
				sub2(x, y, &z)
			}

			// Loop 9
			j = 1
			k = 2
			l = 3
			e1[1] = 1.0
			e1[2] = 2.0
			e1[3] = 3.0
			for i := 1; i <= n9; i++ {
				sub3()
			}

			// Loop 11
			x = 0.75
			for i := 1; i <= n11; i++ {
				x = math.Sqrt(math.Exp(math.Log(x) / t1))
			}

			inner = isave
			kount += 1

		}

		end_time = seconds()

		dif_time = end_time - begin_time
		kilowhet = float32(100*inner*outer) / dif_time
		fmt.Printf("Elapsed time: %s - MIPS: %f\n", time.Duration(dif_time*float32(time.Second)), kilowhet/1000)

		npass += 1
		if npass < max_pass {
			dif_save = dif_time
			whet_save = kilowhet
			inner *= max_pass
		}
	}

	error = dif_time - (dif_save * float32(max_pass))
	whet_err = whet_save - kilowhet
	percent_err = whet_err * 100 / kilowhet
	fmt.Println("------------------------------------------")
	fmt.Printf("error: %f - whet_err: %f - percent_err: %f", error, whet_err, percent_err)

}

func sub1(e []float64) {
	for i := 1; i <= 6; i++ {
		e[1] = (e[1] + e[2] + e[3] - e[4]) * t1
		e[2] = (e[1] + e[2] - e[3] + e[4]) * t1
		e[3] = (e[1] - e[2] + e[3] + e[4]) * t1
		e[4] = (-e[1] + e[2] + e[3] + e[4]) / t3
	}
}

func sub2(x float64, y float64, z *float64) {
	var x1, y1 float64
	x1 = x
	y1 = y
	x1 = (x1 + y1) * t1
	y1 = (x1 + y1) * t1
	*z = (x1 + y1) / t3
}

func sub3() {
	e1[j] = e1[k]
	e1[k] = e1[l]
	e1[l] = e1[j]
}

func seconds() float32 {
	now := time.Now()

	mins := now.Minute()
	secs := now.Second()
	mills := now.Nanosecond() / 1000000
	hour := now.Hour()
	s1 := float32((hour*3600 + mins*60 + secs)) + (float32(mills) * 0.001)

	// fmt.Printf("hour: %d - mins: %d - secs: %d - mils: %d\n", hour, mins, secs, mills)

	return s1
}

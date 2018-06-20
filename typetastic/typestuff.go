package main

import "C"
import (
	"fmt"
)

func main() {

	s := slice{0, 1, 2, 3, 5, 7, 11, 13, 17, 19} //Go will autmatically type the numbers
	x := []num{1, 2, 3, 4, 5}                    // this will be stored as a string of nums
	s = append(s, x...)                          //have to append like this (duh)
	w := s[3:7]                                  //grabbing a slice of s
	s = append(s, w...)                          // appending w to s (interesting , but expected)
	fmt.Println(s)
	//	ints := []int(s)         ///  ----  This won't work
	//  nums := []num(s) //  -- This works
	//	ints := []int(nums)  ///  ----but this doesn't

	// ughhh, we have to do this
	ints := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		ints[i] = int(s[i])
	}
	fmt.Println("This is ints:", ints)

	//Now a common c type thang
	var number num
	var pnt point
	pnt = &number
	fmt.Println(pnt, &number)
	fmt.Println(*pnt, number)
	number = 1
	fmt.Println(*pnt, number)
	*pnt = 300
	fmt.Println(*pnt, number)
	pnts := make(points, len(ints))
	//others := make([]other, len(pnts))
	holders := make([]holder, len(pnts))
	for i := 0; i < len(pnts); i++ {
		pnts[i] = &s[i]
		holders[i].o = other(pnts[i])

	}
	others := holderstoothers(holders)
	fmt.Println(pnts)
	fmt.Println(holders)
	*pnts[4] = *pnt //Whaaaa

	fmt.Println(s, *others[4], *holders[4].o)

}

type num int

type slice []num

type point *num

type points []point

type other point

type holder struct {
	o other
}

func otherstoholders(o []other) []holder {
	h := make([]holder, len(o))
	for i := 0; i < len(h); i++ {
		h[i].o = o[i]
	}
	return h
}
func holderstoothers(h []holder) []other {
	o := make([]other, len(h))
	for i := 0; i < len(o); i++ {
		o[i] = h[i].o
	}
	return o
}

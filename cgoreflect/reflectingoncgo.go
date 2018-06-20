package main

/*
#include <stdio.h>

typedef struct{
	int first;
	int second;
	int third;
} Results;

Results MakeResults(int first,int second,int third){
	Results Place;
	Place.first=first;
	Place.second=second;
	Place.third=third;
	return Place;
}

*/
import "C"
import (
	"fmt"
	"reflect"
)

func main() {

	x := C.MakeResults(1, 2, 3)
	fmt.Println("C.MakeResults:", x)
	v := reflect.ValueOf(&x)
	fmt.Println("\nv:=reflect.ValueOf(x), \nv=", v)
	fmt.Println("v.IsValid() returns: ", v.IsValid())
	vnopoint := reflect.ValueOf(x)
	fmt.Println("NumberofFields", vnopoint.NumField())
	setable := v.Elem()
	fmt.Println("v.Kind()", v.Kind())
	fmt.Println("setable.Kind()", setable.Kind())
	fmt.Println("vnopoint.Kind()", vnopoint.Kind())

	fmt.Println("\nsetable:=v.Elm() \nsettable=", setable)
	fmt.Println("setable.CanSet() returns: ", setable.CanSet())
	fmt.Println("NumberofFields", setable.NumField())
	t := reflect.TypeOf(x)
	fmt.Println("\nt := reflect.TypeOf(x), \nt=", t)

}

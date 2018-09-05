package main

import (
	"fmt"
)

func main() {
	// define variable with literal
	array := [...]int{
		10, 20, 30,
		5: 10,
	}

	// constant variable
	const x = "hello"
	// x = "hello!"

	// define variable without literal
	var y = "world"
	y = "world!"
	fmt.Println(array, x, y)

	// allocate a type
	arrayOfPointer := [...]*int{new(int), new(int)}
	*arrayOfPointer[1] = 10
	fmt.Println(arrayOfPointer, *arrayOfPointer[0], *arrayOfPointer[1])

	// without literal
	// var arrayByNewKeyword [5]int
	// arrayByNewKeyword = new(Array, ..???...
	// seems not supported, with new() func.
	// see go/types package to find out more
	// .. the answer was make([]int, 5) which uses heap.

	// clone an array
	arrayOfPointer2 := arrayOfPointer
	fmt.Println(arrayOfPointer2)

	// clone all values? or copy the reference?
	arrayOfInts := [...]int{1, 2, 3, 4}
	arrayOfInts2 := arrayOfInts
	arrayOfInts[0] = 5
	arrayOfInts3 := arrayOfInts2

	// cloned..
	// **equality check is performed strictly (in binary level)**
	fmt.Println(arrayOfInts, arrayOfInts2, arrayOfInts == arrayOfInts2, arrayOfInts2 == arrayOfInts3)

	// how to copy the reference?
	// before figuring out it, see multi-dim array
	// [...][...] not works
	arrayOfArrayOfInt := [...][2]int{{1, 2}, {3, 4}, {5, 6}}
	arrayOfArrayOfInt2 := arrayOfArrayOfInt

	arrayOfArrayOfInt2[0][1] *= 10
	fmt.Println(arrayOfArrayOfInt, arrayOfArrayOfInt2, arrayOfArrayOfInt == arrayOfArrayOfInt2)
	// := deep copy the whole array data

	// here let's copy the reference.
	pointerOfArrayOfArrayOfInt := &arrayOfArrayOfInt
	(*pointerOfArrayOfArrayOfInt)[0][1] *= 10
	fmt.Println(*pointerOfArrayOfArrayOfInt, arrayOfArrayOfInt, pointerOfArrayOfArrayOfInt == &arrayOfArrayOfInt, *pointerOfArrayOfArrayOfInt == arrayOfArrayOfInt)

	// just a moment, see number notations, ..., hexa, octal, binary....?
	fmt.Println(123, 123.123, -10e-3*1.5343E10, 0x1234, 01234)
	// binary literal is not supported yet: https://github.com/golang/go/issues/19308
	// a := 0b1234

	// **As see from above, array in golang is treated as a primitive value**
	// and here see the **slice**: an abstracted dynamic array type.
	sliceOfInt := make([]int, 5) // can pass capacity as like , 10)
	fmt.Println(sliceOfInt)

	// and the literal for making slice was, just to omit ... from [...] part
	sliceOfInt2 := []int{1, 2, 3}
	var nilSlice []int     // can omit allocation part for empty slice
	nilSlice2 := []int{}   // equal to above line, and the length of each expressions are same,.. var <-> :{}, pick anything
	var emptyInt int       // allocation! as 0
	var emptyString string // allocation! as "" empty string
	var emptyFloat float32 // as 0
	var emptyArray [3]int
	fmt.Println(sliceOfInt2, nilSlice, &sliceOfInt2[0], nilSlice2, emptyInt, emptyString, emptyFloat, emptyArray)

	// copy the reference?
	sliceOfInt3 := sliceOfInt2
	fmt.Println(sliceOfInt2, sliceOfInt3) // , sliceOfInt2 == sliceOfInt3) is invalid comparison

	// slice the slice.
	sliceOfInt2 = append(sliceOfInt2, 1, 2, 3, 4, 5)
	partOfSliceOfInt2 := sliceOfInt2[0 : len(sliceOfInt2)-2] // cap(sliceOfInt2) = capacity of slice
	partOfSliceOfInt2[0] *= 10
	fmt.Println(partOfSliceOfInt2, sliceOfInt2)

	// slice the array.
	partOfArrayOfInts := arrayOfInts[0 : len(arrayOfInts)-2]
	partOfArrayOfInts[0] *= 10 // ** the new array is a slice which refers the old array **
	partOfArrayOfInts[1] *= 10
	newCopyOfArrayOfInts := arrayOfInts
	newCopyOfArrayOfInts[1] *= 777
	arrayOfInts[1] *= 7
	fmt.Println(partOfArrayOfInts, arrayOfInts, newCopyOfArrayOfInts)

	// and when append to the slice which refer the array?
	partOfArrayOfInts = append(partOfArrayOfInts, 1, 2, 3, 4, 5, 6)
	fmt.Println(partOfArrayOfInts, arrayOfInts, newCopyOfArrayOfInts)

	// append element to the array.
	// intArray := [3]int{1, 2, 3}
	// intArray = append(intArray, 4) error!

	// [1, 2, 3, 4, 5, 6]
	slice1 := []int{1, 2, 3, 4, 5, 6}
	// [1, 2, 3, 4, ...]
	slice2 := slice1[0:4]
	// overwrite the slice1 then...
	slice2 = append(slice2, 55)
	// slice1[4] overwitten!
	fmt.Println(slice1, slice2)

	// to prevent this behavior
	// slice strictly by give slicing capacity arg also
	// to branch the new slice after exceeding the given capacity
	slice3 := slice1[0:4:4]
	slice3 = append(slice3, 99)
	fmt.Println(slice1, slice3)

	// unpack slice or array as arguments
	fmt.Println(append(slice3, partOfArrayOfInts...))

	// recSlice... is refer? or copy?
	recSlice := []int{1, 2, 3}
	recSlice = append(recSlice, recSlice...)
	recSlice[0] *= 100
	// it was copy
	fmt.Println(recSlice)

	// create array from slice
	// var arrayFromSlice [4]int
	// arrayFromSlice = recSlice
	// fmt.Println(arrayFromSlice, recSlice)
	// impossible, array is not created dynamically

	/** Caution: array and slice is primitive value, not the abstract type! **/

	// iteration
	// here k, v is copied value
	for k, v := range slice1 {
		fmt.Printf("%d-%d (%x-%x)\t", k, v, &k, &v)
	}

	for k := 1; k < len(slice1); k++ {
		// here k, v is copied value
		fmt.Printf("%d-%d (%x-%x)\t", k, slice1[k], &k, &slice1[k])
	}

	// multi-dim slice
	multiDimSlice := [][]int32{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}}
	multiDimSlice[0] = append(multiDimSlice[0], 1, 2, 3)
	fmt.Println(multiDimSlice)

	// map (keys are unordered)
	var dict = make(map[int]int)
	dict[0] = 123
	fmt.Println(dict)

	// literal
	dict2 := map[string]int8{"a": 1, "b": 2, "c": 012}
	fmt.Println(dict2, dict2["c"])

	var nilArray [2]int // allocation occured
	nilArray[0] = 10
	var nilMap map[int]int
	// nilMap[0] = 1 error! because allocation not ouccured (only pointer)
	nilMap = dict // use like this
	var nilSlice3 []int
	nilSlice3 = append(nilSlice3, 123)
	fmt.Println(nilMap, nilArray, nilSlice3) //, nilMap)

	// existence of key in map
	key := "c"
	dict2v, dict2exists := dict2[key]
	if dict2exists {
		fmt.Printf("dict2 %s=%d", key, dict2v)
	} else {
		fmt.Printf("dict2 not exists.. %s=%d", key, dict2v)
	}

	// delete map element
	delete(dict2, "c")

	// iterate map
	for k, v := range dict2 {
		fmt.Println(k, v)
	}

	/* map, slice ... are called by reference!!! */
}

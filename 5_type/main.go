package main

import (
	"fmt"
)

// firstType type is blabla...
type firstType struct {
	id   int64
	name string
}

type secondType struct {
	first firstType
	level string
}

// type alias (base type)
type scondTypeAlias secondType

// field and type of field are same
type whatType struct {
	firstType
}

// method
func (s secondType) updateLevel() secondType { // call by value
	s.level = "updated level"
	return s
}

func (s *secondType) updateLevelByRef() { // call by ref
	s.level = "updated level"
}

func main() {
	first := *new(firstType) // it returns pointer of firstType
	first.id = 1
	first.name = "hey"
	fmt.Println(first)

	first2 := firstType{ // use literal
		id:   1,
		name: "hey",
	}
	fmt.Println(first2)

	var first3 firstType // allocation occurs as struct from C
	first3.id = 1
	first3.name = "abc"
	fmt.Println(first3)

	// embedding (composition)
	second := secondType{
		first: firstType{
			id:   2,
			name: "world",
		},
		level: "super",
	}

	second2 := secondType{
		first: second.first,
		level: "normal",
	}

	second2.first.name = "updated"

	fmt.Println(second, second2)

	secondAlias := scondTypeAlias{
		first: second.first,
		level: second.level,
	}

	var secondAlias2 scondTypeAlias
	secondAlias2 = scondTypeAlias(second) // can be assigned type conversion when aliased type

	fmt.Println(secondAlias, secondAlias2)

	// method
	second3 := secondType{
		first: second.first,
		level: "hello level",
	}

	second3_2 := second3.updateLevel()

	second4 := second3         // copy
	second4.updateLevelByRef() // update it self

	second4_2 := (&second4).updateLevel() // receiver's pointer type is automatically adjusted for method proto
	fmt.Println(second3, second3_2, second4, second4_2)

	/* golang's reference types: slice, map, channel, interface, func */

	// wait...
	id := int64(1234)
	name := "name...."
	whatType := whatType{
		firstType{ // when equal to variable name, field name can be omiited
			id,
			name,
		},
	}
	fmt.Println(whatType)

	// update slice by function
	slice5 := []int{1, 2, 3}
	updateSlice(slice5) // call by reference
	fmt.Println(slice5)

	slice6 := sliceOfIntType{1, 2, 3}
	slice6.updateSliceOfIntType() // call by reference
	fmt.Println(slice6)

	// assignment
	slice5 = slice6 // share the reference of data
	slice5[0] = 3000
	slice5Casted := sliceOfIntType(slice5)
	fmt.Println(slice5, slice6)
	fmt.Printf("%x == %x ?\n", &slice5Casted[0], &slice6[0])
}

/* basically all the types (structs) copy all values, but when the some field have pointer, it works like call by reference */

type sliceOfIntType []int

func updateSlice(slice []int) {
	slice[0] = 10000
}

func (s sliceOfIntType) updateSliceOfIntType() {
	s[0] = 2000
}

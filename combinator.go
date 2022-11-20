package combinator

import (
	"reflect"
	"sort"
)

func Combinations(container interface{}, length uint) interface{} {
	containerType := reflect.TypeOf(container)
	if containerType.Kind() != reflect.Slice {
		panic("argument must be a slice")
	}

	type Position = int
	nPerM := func(n Position, m Position) Position {
		res := 1
		for i := 0; i < m; i++ {
			res *= n - i
			res /= i + 1
		}
		return res
	}

	elementType := containerType.Elem()
	containerValue := reflect.ValueOf(container)
	nrElements := containerValue.Len()
	combinations := reflect.MakeSlice(
		reflect.SliceOf(reflect.SliceOf(elementType)),
		0,
		nPerM(nrElements, int(length)),
	)

	positions := make([]Position, length)
	for i := 0; i < Position(length); i++ {
		positions[i] = i
	}
	maxPosition := func(position Position) Position {
		return nrElements + position - Position(length)
	}
	var increment func(Position) bool
	increment = func(position Position) bool {
		positions[position]++
		if positions[position] > maxPosition(position) {
			if position == 0 {
				return false
			} else {
				res := increment(position - 1)
				positions[position] = positions[position-1] + 1
				return res
			}
		}
		return true
	}
	addCombination := func() {
		combination := reflect.MakeSlice(reflect.SliceOf(elementType), 0, int(length))
		for _, position := range positions {
			combination = reflect.Append(combination, containerValue.Index(position))
		}
		combinations = reflect.Append(combinations, combination)
	}
	addCombination()
	for increment(Position(length) - 1) {
		addCombination()
	}

	return combinations.Interface()
}

func Permutations(container interface{}, length uint) interface{} {
	containerType := reflect.TypeOf(container)
	if containerType.Kind() != reflect.Slice {
		panic("argument must be a slice")
	}

	type Position = int
	var nPerM func(Position, Position) Position
	nPerM = func(n Position, m Position) Position {
		if m > 1 {
			return n * nPerM(n-1, m-1)
		}
		return n
	}

	elementType := containerType.Elem()
	containerValue := reflect.ValueOf(container)
	nrElements := containerValue.Len()
	size := nPerM(nrElements, int(length))
	permutations := reflect.MakeSlice(reflect.SliceOf(reflect.SliceOf(elementType)), 0, size)

	positions := make([]Position, length)
	for i := 0; i < Position(length); i++ {
		positions[i] = i
	}
	insertUnique := func(position Position, value Position) {
		prevValue := -1
		for {
			add := 0
			for i := 0; i < position; i++ {
				if (prevValue+1 < positions[i]+1) && (positions[i] <= value) {
					add++
				}
			}
			prevValue = value
			value += add
			if prevValue == value {
				break
			}
		}
		positions[position] = value
	}
	goTo := func(index Position) {
		nrElementsCopy := nrElements
		for i := 0; i < int(length); i++ {
			insertUnique(i, index%nrElementsCopy)
			index /= nrElementsCopy
			nrElementsCopy--
		}
	}
	addPermutation := func() {
		permutation := reflect.MakeSlice(reflect.SliceOf(elementType), 0, int(length))
		for _, position := range positions {
			permutation = reflect.Append(permutation, containerValue.Index(position))
		}
		permutations = reflect.Append(permutations, permutation)
	}
	for i := 0; i < size; i++ {
		goTo(i)
		addPermutation()
	}

	return permutations.Interface()
}

func CartesianProducts(containers interface{}) interface{} {
	containersType := reflect.TypeOf(containers)
	if containersType.Kind() != reflect.Slice {
		panic("argument must be a slice")
	}
	containerType := containersType.Elem()
	if containerType.Kind() != reflect.Slice {
		panic("argument must be a slice")
	}

	type Position = int
	elementType := containerType.Elem()
	containersValue := reflect.ValueOf(containers)
	nrContainers := containersValue.Len()
	size := 1
	for i := 0; i < nrContainers; i++ {
		size *= containersValue.Index(i).Len()
	}
	cartesianProducts := reflect.MakeSlice(reflect.SliceOf(reflect.SliceOf(elementType)), 0, size)

	positions := make([]Position, nrContainers)
	var increment func(Position)
	increment = func(position Position) {
		positions[position]++
		if positions[position] == containersValue.Index(position).Len() {
			positions[position] = 0
			increment(position + 1)
		}
	}
	addCartesian := func() {
		cartesian := reflect.MakeSlice(reflect.SliceOf(elementType), 0, nrContainers)
		for i, position := range positions {
			cartesian = reflect.Append(cartesian, containersValue.Index(i).Index(position))
		}
		cartesianProducts = reflect.Append(cartesianProducts, cartesian)
	}
	addCartesian()
	for i := 0; i < size-1; i++ {
		increment(0)
		addCartesian()
	}
	return cartesianProducts.Interface()
}

func Twines(containers interface{}) interface{} {
	containersType := reflect.TypeOf(containers)
	if containersType.Kind() != reflect.Slice {
		panic("argument must be a slice")
	}
	containerType := containersType.Elem()
	if containerType.Kind() != reflect.Slice {
		panic("argument must be a slice")
	}

	type Position = int
	containersValue := reflect.ValueOf(containers)
	nrContainers := containersValue.Len()
	elementType := containerType.Elem()

	size := 0
	for i := 0; i < nrContainers; i++ {
		size += containersValue.Index(i).Len()
	}

	containersElements := make([][][]Position, 0, nrContainers) // TODO: rename
	nrTwines := 1
	for i := 0; i < nrContainers; i++ {
		containerElements := make([]Position, 0, size)
		for j := 0; j < size; j++ {
			containerElements = append(containerElements, j)
		}
		containerSize := containersValue.Index(i).Len()
		combinations := Combinations(containerElements, uint(containerSize)).([][]Position)
		containersElements = append(containersElements, combinations)
		nrTwines *= len(combinations)
		size -= containerSize
	}

	twineSize := 0
	for i := 0; i < nrContainers; i++ {
		twineSize += containersValue.Index(i).Len()
	}

	twines := reflect.MakeSlice(reflect.SliceOf(reflect.SliceOf(elementType)), 0, nrTwines)
	for _, positionsIndices := range CartesianProducts(containersElements).([][][]Position) {
		twine := reflect.MakeSlice(reflect.SliceOf(elementType), twineSize, twineSize)
		filledPositions := make([]Position, 0, twineSize)
		for containerIdx, positionIndices := range positionsIndices {
			filledByContainer := make([]Position, 0, containersValue.Index(containerIdx).Len())
			for containerPosition, twinePosition := range positionIndices {
				for _, filledPosition := range filledPositions {
					if twinePosition >= filledPosition {
						twinePosition++
					}
				}
				twine.Index(twinePosition).Set(containersValue.Index(containerIdx).Index(containerPosition))
				filledByContainer = append(filledByContainer, twinePosition)
			}
			filledPositions = append(filledPositions, filledByContainer...)
			sort.Ints(filledPositions)
		}
		twines = reflect.Append(twines, twine)
	}
	return twines.Interface()
}

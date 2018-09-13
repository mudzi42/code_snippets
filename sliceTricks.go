package codesnippets

//Slice Tricks

func sliceTricks() {
	// AppendVector
	a = append(a, b...)

	// Copy
	b = make([]T, len(a))
	copy(b, a)
	// or
	b = append([]T(nil), a...)

	// Cut
	a = append(a[:i], a[j:]...)

	//Delete
	a = append(a[:i], a[i+1:]...)
	// or
	a = a[:i+copy(a[i:], a[i+1:])]

	//Delete without preserving order
	a[i] = a[len(a)-1]
	a = a[:len(a)-1]
}

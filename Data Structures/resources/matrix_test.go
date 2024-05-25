// go test -run none -bench . -benchtime 3s

// Tests to show how Data Oriented Design matters.

package main

import "testing"

var fa int

// capture the time it takes to perform a linkedList traversal
func BenchmarkLinkedListTraverse(b *testing.B) {
	var a int

	// all the code we want to benchmark should be inside this loop
	for i := 0; i < b.N; i++ {
		a = LinkedListTraverse()
	}
	fa = a
}

// capture the time it takes to perform a column traversal
func BenchmarkColumnTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = ColumnTraverse()
	}
	fa = a
}

// capture the time it takes to perform a row traversal
func BenchmarkRowTraverse(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = RowTraverse()
	}
	fa = a
}

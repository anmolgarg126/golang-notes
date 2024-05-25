// Package caching provides code to show why Data Oriented Design matters. How
// data layouts matter more to performance than algorithm efficiency.
package main

import "fmt"

// create a square matrix of 2Mb * 2Mb
const (
	rows = 2 * 1024
	cols = 2 * 1024
)

// represents a matrix with a large number of columns per row
var matrix [rows][cols]byte

// data struct represents a data node of our linked list
type data struct {
	v byte
	p *data
}

// head points to the head node
var head *data

func init() {
	// Create a link list with the same number of elements.
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {

			// Create a new node and link it backwards.
			head = &data{p: head}

			// Add a value to all even elements.
			if row%2 == 0 {
				matrix[row][col] = 0xFF
				head.v = 0xFF
			}
		}
	}

	// Count the number of elements in the link list.
	var ctr int
	d := head
	for d != nil {
		ctr++
		d = d.p
	}

	fmt.Println("Elements in the link list", ctr)
	fmt.Println("Elements in the matrix", rows*cols)
}

// LinkedListTraverse traverses linearly
func LinkedListTraverse() int {
	var ctr int
	d := head

	for d != nil {
		if d.v == 0xFF {
			ctr++
		}
		d = d.p
	}
	return ctr
}

// ColumnTraverse traverses linearly down each column
func ColumnTraverse() int {
	var ctr int
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}
	return ctr
}

// RowTraverse traverses linearly down each row
func RowTraverse() int {
	var ctr int
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				ctr++
			}
		}
	}
	return ctr
}

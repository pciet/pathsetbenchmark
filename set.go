// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import ()

type Item interface {
	Equal(Item) bool
}

type Set interface {
	Add(Item) Set
	Remove(Item) Set
	Combine(...Set) Set
	Reduce() Set
	Has(Item) bool
	Equal(Set) bool
	Diff(Set) Set
}

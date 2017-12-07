// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import ()

type SlicePathSet []Path

func SliceCopy(the SlicePathSet) SlicePathSet {
	next := make(SlicePathSet, len(the))
	for i, item := range the {
		next[i] = item
	}
	return next
}

func SliceAdd(the SlicePathSet, add Path) SlicePathSet {
	the = append(the, add)
	return the
}

func (the SlicePathSet) Add(an Item) Set {
	return SliceAdd(the, an.(Path))
}

func SliceDelete(the SlicePathSet, remove Path) SlicePathSet {
	out := make(SlicePathSet, 0, len(the))
	for _, item := range the {
		if item.PathEqual(remove) {
			continue
		}
		out = append(out, item)
	}
	return out
}

func (the SlicePathSet) Remove(an Item) Set {
	return SliceDelete(the, an.(Path))
}

func SliceCombine(sets ...SlicePathSet) SlicePathSet {
	out := make(SlicePathSet, 0, len(sets[0])*len(sets))
	for _, set := range sets {
		for _, item := range set {
			out = append(out, item)
		}
	}
	return out
}

func (the SlicePathSet) Combine(sets ...Set) Set {
	combine := make([]SlicePathSet, len(sets)+1)
	combine[0] = the
	for i, set := range sets {
		combine[i+1] = set.(SlicePathSet)
	}
	return SliceCombine(combine...)
}

func SliceReduce(the SlicePathSet) SlicePathSet {
	out := make(SlicePathSet, 0, len(the))
	for _, item := range the {
		if SliceHas(out, item) {
			continue
		}
		out = append(out, item)
	}
	return out
}

func (the SlicePathSet) Reduce() Set {
	return SliceReduce(the)
}

func SliceHas(the SlicePathSet, item Path) bool {
	for _, path := range the {
		if path.PathEqual(item) {
			return true
		}
	}
	return false
}

func (the SlicePathSet) Has(an Item) bool {
	return SliceHas(the, an.(Path))
}

func SliceEqual(sets ...SlicePathSet) bool {
	first := sets[0]
	length := len(first)
	for _, set := range sets {
		if length != len(set) {
			return false
		}
		// since slices cannot be compared there will be an unnecessary check on first
		for _, item := range first {
			if SliceHas(set, item) == false {
				return false
			}
		}
	}
	return true
}

func (the SlicePathSet) Equal(a Set) bool {
	return SliceEqual(the, a.(SlicePathSet))
}

func SliceDiff(a SlicePathSet, b SlicePathSet) SlicePathSet {
	out := make(SlicePathSet, 0, len(a))
	for _, item := range a {
		if SliceHas(b, item) == false {
			out = append(out, item)
		}
	}
	for _, item := range b {
		if SliceHas(a, item) == false {
			out = append(out, item)
		}
	}
	return out
}

func (the SlicePathSet) Diff(with Set) Set {
	return SliceDiff(the, with.(SlicePathSet))
}

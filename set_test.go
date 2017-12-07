// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"testing"
)

type AddCase struct {
	Set
	Item
	Out Set
}

var AddCases = []AddCase{
	{
		Set:  SlicePathSet{},
		Item: Path{{0, 0}},
		Out: SlicePathSet{
			{{0, 0}},
		},
	},
	{
		Set: SlicePathSet{
			{{1, 1}},
		},
		Item: Path{{2, 2}},
		Out: SlicePathSet{
			{{2, 2}},
			{{1, 1}},
		},
	},
	{
		Set: SlicePathSet{
			{{0, 0}, {0, 1}},
			{{0, 2}, {1, 1}},
		},
		Item: Path{{0, 0}, {0, 3}, {2, 2}},
		Out: SlicePathSet{
			{{0, 0}, {0, 3}, {2, 2}},
			{{0, 0}, {0, 1}},
			{{0, 2}, {1, 1}},
		},
	},
	{
		Set:  MapPathSet{},
		Item: Path{{0, 0}},
		Out: MapPathSet{
			&Path{{0, 0}}: {},
		},
	},
	{
		Set: MapPathSet{
			&Path{{1, 1}}: {},
		},
		Item: Path{{2, 2}},
		Out: MapPathSet{
			&Path{{2, 2}}: {},
			&Path{{1, 1}}: {},
		},
	},
	{
		Set: MapPathSet{
			&Path{{0, 0}, {0, 1}}: {},
			&Path{{0, 2}, {1, 1}}: {},
		},
		Item: Path{{0, 0}, {0, 3}, {2, 2}},
		Out: MapPathSet{
			&Path{{0, 0}, {0, 3}, {2, 2}}: {},
			&Path{{0, 0}, {0, 1}}:         {},
			&Path{{0, 2}, {1, 1}}:         {},
		},
	},
}

func TestAdd(t *testing.T) {
	for i, c := range AddCases {
		if c.Out.Equal(c.Set.Add(c.Item)) == false {
			t.Fatalf("%v: out not equal", i)
		}
	}
}

type RemoveCase struct {
	Set
	Item
	Out Set
}

var RemoveCases = []RemoveCase{
	{
		Set: MapPathSet{
			&Path{{0, 0}}:         {},
			&Path{{0, 1}, {1, 1}}: {},
		},
		Item: Path{{0, 0}},
		Out: MapPathSet{
			&Path{{0, 1}, {1, 1}}: {},
		},
	},
	{
		Set: MapPathSet{
			&Path{{5, 5}}: {},
		},
		Item: Path{{5, 5}},
		Out:  MapPathSet{},
	},
	{
		Set: MapPathSet{
			&Path{{5, 5}, {4, 5}, {3, 3}}: {},
			&Path{{3, 3}, {2, 2}}:         {},
			&Path{{4, 4}}:                 {},
		},
		Item: Path{{5, 5}, {4, 5}, {3, 3}},
		Out: MapPathSet{
			&Path{{3, 3}, {2, 2}}: {},
			&Path{{4, 4}}:         {},
		},
	},
	{
		Set: SlicePathSet{
			{{0, 0}},
			{{0, 1}, {1, 1}},
		},
		Item: Path{{0, 0}},
		Out: SlicePathSet{
			{{0, 1}, {1, 1}},
		},
	},
	{
		Set: SlicePathSet{
			{{5, 5}},
		},
		Item: Path{{5, 5}},
		Out:  SlicePathSet{},
	},
	{
		Set: SlicePathSet{
			{{5, 5}, {4, 5}, {3, 3}},
			{{3, 3}, {2, 2}},
			{{4, 4}},
		},
		Item: Path{{5, 5}, {4, 5}, {3, 3}},
		Out: SlicePathSet{
			{{3, 3}, {2, 2}},
			{{4, 4}},
		},
	},
}

func TestRemove(t *testing.T) {
	for i, c := range RemoveCases {
		if c.Out.Equal(c.Set.Remove(c.Item)) == false {
			t.Fatalf("%v failed", i)
		}
	}
}

type CombineCase struct {
	A   Set
	B   Set
	C   Set
	D   Set
	Out Set
}

var CombineCases = []CombineCase{
	{
		A: SlicePathSet{
			{{1, 1}, {2, 2}},
		},
		B: SlicePathSet{
			{{1, 1}},
		},
		Out: SlicePathSet{
			{{1, 1}, {2, 2}},
			{{1, 1}},
		},
	},
	{
		A: SlicePathSet{
			{{1, 1}, {2, 2}},
			{{1, 1}},
		},
		B: SlicePathSet{
			{{1, 1}},
		},
		Out: SlicePathSet{
			{{1, 1}, {2, 2}},
			{{1, 1}},
			{{1, 1}},
		},
	},
	{
		A: SlicePathSet{
			{{1, 1}, {2, 2}},
		},
		B: SlicePathSet{
			{{1, 1}},
		},
		C: SlicePathSet{
			{{3, 3}, {4, 4}},
		},
		Out: SlicePathSet{
			{{3, 3}, {4, 4}},
			{{1, 1}, {2, 2}},
			{{1, 1}},
		},
	},
	{
		A: SlicePathSet{
			{{0, 2}, {5, 3}},
		},
		B: SlicePathSet{
			{{1, 7}, {2, 6}},
		},
		C: SlicePathSet{
			{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}},
		},
		D: SlicePathSet{
			{{3, 3}, {4, 4}},
		},
		Out: SlicePathSet{
			{{3, 3}, {4, 4}},
			{{1, 7}, {2, 6}},
			{{0, 2}, {5, 3}},
			{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}},
		},
	},
	{
		A: MapPathSet{
			&Path{{1, 1}, {2, 2}}: {},
		},
		B: MapPathSet{
			&Path{{1, 1}}: {},
		},
		Out: MapPathSet{
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{1, 1}}:         {},
		},
	},
	{
		A: MapPathSet{
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{1, 1}}:         {},
		},
		B: MapPathSet{
			&Path{{1, 1}}: {},
		},
		Out: MapPathSet{
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{1, 1}}:         {},
			&Path{{1, 1}}:         {},
		},
	},
	{
		A: MapPathSet{
			&Path{{1, 1}, {2, 2}}: {},
		},
		B: MapPathSet{
			&Path{{1, 1}}: {},
		},
		C: MapPathSet{
			&Path{{3, 3}, {4, 4}}: {},
		},
		Out: MapPathSet{
			&Path{{3, 3}, {4, 4}}: {},
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{1, 1}}:         {},
		},
	},
	{
		A: MapPathSet{
			&Path{{0, 2}, {5, 3}}: {},
		},
		B: MapPathSet{
			&Path{{1, 7}, {2, 6}}: {},
		},
		C: MapPathSet{
			&Path{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}}: {},
		},
		D: MapPathSet{
			&Path{{3, 3}, {4, 4}}: {},
		},
		Out: MapPathSet{
			&Path{{3, 3}, {4, 4}}:                         {},
			&Path{{1, 7}, {2, 6}}:                         {},
			&Path{{0, 2}, {5, 3}}:                         {},
			&Path{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}}: {},
		},
	},
}

func TestCombine(t *testing.T) {
	for i, c := range CombineCases {
		args := make([]Set, 0, 4)
		if c.A != nil {
			args = append(args, c.A)
		}
		if c.B != nil {
			args = append(args, c.B)
		}
		if c.C != nil {
			args = append(args, c.C)
		}
		if c.D != nil {
			args = append(args, c.D)
		}
		if len(args) < 2 {
			t.Fatalf("%v: not enough sets", i)
		}
		if c.Out.Equal(args[0].Combine(args[1:]...)) == false {
			t.Fatalf("%v failed", i)
		}
	}
}

type ReduceCase struct {
	Set
	Out Set
}

var ReduceCases = []ReduceCase{
	{
		Set: MapPathSet{
			&Path{{0, 0}}: {},
			&Path{{0, 0}}: {},
		},
		Out: MapPathSet{
			&Path{{0, 0}}: {},
		},
	},
	{
		Set: MapPathSet{
			&Path{{0, 0}}: {},
		},
		Out: MapPathSet{
			&Path{{0, 0}}: {},
		},
	},
	{
		Set: MapPathSet{
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{3, 3}, {1, 1}}: {},
			&Path{{1, 1}, {1, 1}}: {},
			&Path{{3, 3}, {1, 1}}: {},
			&Path{{1, 1}, {2, 2}}: {},
		},
		Out: MapPathSet{
			&Path{{1, 1}, {1, 1}}: {},
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{3, 3}, {1, 1}}: {},
		},
	},
	{
		Set: SlicePathSet{
			{{0, 0}},
			{{0, 0}},
		},
		Out: SlicePathSet{
			{{0, 0}},
		},
	},
	{
		Set: SlicePathSet{
			{{0, 0}},
		},
		Out: SlicePathSet{
			{{0, 0}},
		},
	},
	{
		Set: SlicePathSet{
			{{1, 1}, {2, 2}},
			{{3, 3}, {1, 1}},
			{{1, 1}, {1, 1}},
			{{3, 3}, {1, 1}},
			{{1, 1}, {2, 2}},
		},
		Out: SlicePathSet{
			{{1, 1}, {1, 1}},
			{{1, 1}, {2, 2}},
			{{3, 3}, {1, 1}},
		},
	},
}

func TestReduce(t *testing.T) {
	for i, c := range ReduceCases {
		if c.Out.Equal(c.Set.Reduce()) == false {
			t.Fatalf("%v failed", i)
		}
	}
}

type HasCase struct {
	Has bool
	Set
	Item
}

var HasCases = []HasCase{
	{
		Has: true,
		Set: SlicePathSet{
			{{0, 0}, {1, 1}, {2, 2}},
			{{0, 0}},
		},
		Item: Path{{0, 0}, {1, 1}, {2, 2}},
	},
	{
		Has: false,
		Set: SlicePathSet{
			{{0, 0}, {1, 1}, {2, 2}},
			{{0, 0}},
		},
		Item: Path{{0, 0}, {2, 2}, {1, 1}},
	},
	{
		Has: true,
		Set: MapPathSet{
			&Path{{0, 0}, {1, 1}, {2, 2}}: {},
			&Path{{0, 0}}:                 {},
		},
		Item: Path{{0, 0}, {1, 1}, {2, 2}},
	},
	{
		Has: false,
		Set: MapPathSet{
			&Path{{0, 0}, {1, 1}, {2, 2}}: {},
			&Path{{0, 0}}:                 {},
		},
		Item: Path{{0, 0}, {2, 2}, {1, 1}},
	},
}

func TestHas(t *testing.T) {
	for i, c := range HasCases {
		if c.Set.Has(c.Item) != c.Has {
			t.Fatalf("%v failed", i)
		}
	}
}

type EqualCase struct {
	Equal bool
	A     Set
	B     Set
	C     Set
	D     Set
}

var EqualCases = []EqualCase{
	{
		Equal: false,
		A:     SlicePathSet{},
		B: SlicePathSet{
			{{0, 0}},
		},
		C: SlicePathSet{},
		D: SlicePathSet{
			{{2, 2}},
		},
	},
	{
		Equal: true,
		A:     SlicePathSet{},
		B:     SlicePathSet{},
	},
	{
		Equal: true,
		A: SlicePathSet{
			{{0, 0}, {1, 1}, {2, 2}},
			{{1, 2}, {3, 2}},
		},
		B: SlicePathSet{
			{{1, 2}, {3, 2}},
			{{0, 0}, {1, 1}, {2, 2}},
		},
	},
	{
		Equal: true,
		A: SlicePathSet{
			{{1, 5}},
			{{2, 6}},
			{{1, 1}, {2, 2}},
		},
		B: SlicePathSet{
			{{1, 5}},
			{{1, 1}, {2, 2}},
			{{2, 6}},
		},
		C: SlicePathSet{
			{{2, 6}},
			{{1, 5}},
			{{1, 1}, {2, 2}},
		},
	},
	{
		Equal: false,
		A: SlicePathSet{
			{{1, 5}, {2, 6}},
			{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}},
			{{1, 1}},
		},
		B: SlicePathSet{
			{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}},
			{{1, 5}, {2, 6}, {3, 7}},
			{{1, 1}},
		},
		C: SlicePathSet{
			{{1, 5}, {2, 6}},
			{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}},
			{{1, 1}},
		},
	},
	{
		Equal: false,
		A:     MapPathSet{},
		B: MapPathSet{
			&Path{{0, 0}}: {},
		},
		C: MapPathSet{},
		D: MapPathSet{
			&Path{{2, 2}}: {},
		},
	},
	{
		Equal: true,
		A:     MapPathSet{},
		B:     MapPathSet{},
	},
	{
		Equal: true,
		A: MapPathSet{
			&Path{{0, 0}, {1, 1}, {2, 2}}: {},
			&Path{{1, 2}, {3, 2}}:         {},
		},
		B: MapPathSet{
			&Path{{1, 2}, {3, 2}}:         {},
			&Path{{0, 0}, {1, 1}, {2, 2}}: {},
		},
	},
	{
		Equal: true,
		A: MapPathSet{
			&Path{{1, 5}}:         {},
			&Path{{2, 6}}:         {},
			&Path{{1, 1}, {2, 2}}: {},
		},
		B: MapPathSet{
			&Path{{1, 5}}:         {},
			&Path{{1, 1}, {2, 2}}: {},
			&Path{{2, 6}}:         {},
		},
		C: MapPathSet{
			&Path{{2, 6}}:         {},
			&Path{{1, 5}}:         {},
			&Path{{1, 1}, {2, 2}}: {},
		},
	},
	{
		Equal: false,
		A: MapPathSet{
			&Path{{1, 5}, {2, 6}}:                                 {},
			&Path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}: {},
			&Path{{1, 1}}: {},
		},
		B: MapPathSet{
			&Path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}: {},
			&Path{{1, 5}, {2, 6}, {3, 7}}:                         {},
			&Path{{1, 1}}:                                         {},
		},
		C: MapPathSet{
			&Path{{1, 5}, {2, 6}}:                                 {},
			&Path{{1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}}: {},
			&Path{{1, 1}}: {},
		},
	},
}

func TestEqual(t *testing.T) {
	for i, c := range EqualCases {
		args := make([]Set, 0, 4)
		if c.A != nil {
			args = append(args, c.A)
		}
		if c.B != nil {
			args = append(args, c.B)
		}
		if c.C != nil {
			args = append(args, c.C)
		}
		if c.D != nil {
			args = append(args, c.D)
		}
		oneFalse := false
		for j := 0; j < 4; j++ {
			if len(args) < (j + 1) {
				break
			}
			for k := 0; k < 4; k++ {
				if len(args) < (k + 1) {
					break
				}
				if j <= k {
					continue
				}
				if c.Equal == false {
					if args[j].Equal(args[k]) == false {
						oneFalse = true
					}
				} else {
					if args[j].Equal(args[k]) == false {
						t.Fatalf("%v failed", i)
					}
				}
			}
		}
		// Set is one argument but the concrete types take variadic arguments, so to translate "not all are equal" we need one false comparison to pass
		if (c.Equal == false) && (oneFalse == false) {
			t.Fatalf("%v failed", i)
		}
	}
}

type DiffCase struct {
	A   Set
	B   Set
	Out Set
}

var DiffCases = []DiffCase{
	{
		A: MapPathSet{
			&Path{{0, 0}, {1, 1}, {2, 2}, {3, 3}}: {},
			&Path{{0, 0}, {1, 5}}:                 {},
			&Path{{0, 0}, {1, 1}}:                 {},
		},
		B: MapPathSet{
			&Path{{0, 0}, {1, 1}}: {},
		},
		Out: MapPathSet{
			&Path{{0, 0}, {1, 1}, {2, 2}, {3, 3}}: {},
			&Path{{0, 0}, {1, 5}}:                 {},
		},
	},
	{
		A: MapPathSet{
			&Path{{0, 0}}: {},
		},
		B: MapPathSet{
			&Path{{0, 0}}: {},
		},
		Out: MapPathSet{},
	},
	{
		A: MapPathSet{
			&Path{{5, 5}, {4, 4}, {3, 3}}: {},
			&Path{{3, 3}, {4, 4}, {5, 5}}: {},
		},
		B: MapPathSet{
			&Path{{3, 3}, {5, 5}, {4, 4}}: {},
		},
		Out: MapPathSet{
			&Path{{5, 5}, {4, 4}, {3, 3}}: {},
			&Path{{3, 3}, {4, 4}, {5, 5}}: {},
			&Path{{3, 3}, {5, 5}, {4, 4}}: {},
		},
	},
	{
		A: SlicePathSet{
			{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
			{{0, 0}, {1, 5}},
			{{0, 0}, {1, 1}},
		},
		B: SlicePathSet{
			{{0, 0}, {1, 1}},
		},
		Out: SlicePathSet{
			{{0, 0}, {1, 1}, {2, 2}, {3, 3}},
			{{0, 0}, {1, 5}},
		},
	},
	{
		A: SlicePathSet{
			{{0, 0}},
		},
		B: SlicePathSet{
			{{0, 0}},
		},
		Out: SlicePathSet{},
	},
	{
		A: SlicePathSet{
			{{5, 5}, {4, 4}, {3, 3}},
			{{3, 3}, {4, 4}, {5, 5}},
		},
		B: SlicePathSet{
			{{3, 3}, {5, 5}, {4, 4}},
		},
		Out: SlicePathSet{
			{{5, 5}, {4, 4}, {3, 3}},
			{{3, 3}, {4, 4}, {5, 5}},
			{{3, 3}, {5, 5}, {4, 4}},
		},
	},
}

func TestDiff(t *testing.T) {
	for i, c := range DiffCases {
		if c.Out.Equal(c.A.Diff(c.B)) == false {
			t.Fatalf("%v failed", i)
		}
	}
}

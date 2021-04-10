package set

import (
	"testing"
)

func TestUnion(t *testing.T) {
	set1 := populateSet(3, 0)
	set2 := populateSet(2, 3)

	set3 := set1.Union(set2)

	if len(set3.Items()) != 5 {
		t.Errorf("wrong count, expected 5 and got %d", set3.Size())
	}
	// don't edit original sets
	if len(set1.Items()) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", set1.Size())
	}
	if len(set2.Items()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", set2.Size())
	}
}

func TestIntersection(t *testing.T) {
	set1 := populateSet(3, 0)
	set2 := populateSet(2, 0)

	set3 := set1.Intersection(set2)

	if len(set3.Items()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", set3.Size())
	}
	// don't edit original sets
	if len(set1.Items()) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", set1.Size())
	}
	if len(set2.Items()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", set2.Size())
	}
}

func TestDifference(t *testing.T) {
	set1 := populateSet(3, 0)
	set2 := populateSet(2, 0)

	set3 := set1.Difference(set2)

	if len(set3.Items()) != 1 {
		t.Errorf("wrong count, expected 2 and got %d", set3.Size())
	}
	// don't edit original sets
	if len(set1.Items()) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", set1.Size())
	}
	if len(set2.Items()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", set2.Size())
	}
}

func TestSubset(t *testing.T) {
	set1 := populateSet(3, 0)
	set2 := populateSet(2, 0)

	if set1.Subset(set2) {
		t.Errorf("expected false and got true")
	}

	// don't edit original sets
	if len(set1.Items()) != 3 {
		t.Errorf("wrong count, expected 3 and got %d", set1.Size())
	}
	if len(set2.Items()) != 2 {
		t.Errorf("wrong count, expected 2 and got %d", set2.Size())
	}

	// try real subsets
	set1 = populateSet(2, 0)
	if !set1.Subset(set2) {
		t.Errorf("expected true and got false")
	}

	set1 = populateSet(1, 0)
	if !set1.Subset(set2) {
		t.Errorf("expected true and got false")
	}
}

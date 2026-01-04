package bindings

import (
	"fmt"
	"testing"
)

// TestBindingsLink verifies that the FAISS C libraries are correctly linked
// by creating a simple index and checking its properties.
func TestBindingsLink(t *testing.T) {
	// Create a flat L2 index with dimension 64
	ptr := CreateTestIndex(64)
	defer FreeIndex(ptr)

	// Verify dimension
	d := GetIndexDimension(ptr)
	if d != 64 {
		t.Errorf("expected dimension 64, got %d", d)
	}

	// Verify ntotal is 0 (empty index)
	ntotal := GetIndexNtotal(ptr)
	if ntotal != 0 {
		t.Errorf("expected ntotal 0, got %d", ntotal)
	}

	t.Log("FAISS bindings linked and working correctly")
}

// TestMultipleDimensions tests creating indexes with different dimensions.
func TestMultipleDimensions(t *testing.T) {
	dimensions := []int{32, 64, 128, 256, 512}

	for _, dim := range dimensions {
		dim := dim // capture for closure
		t.Run(fmt.Sprintf("dim_%d", dim), func(t *testing.T) {
			ptr := CreateTestIndex(dim)
			defer FreeIndex(ptr)

			got := GetIndexDimension(ptr)
			if got != dim {
				t.Errorf("expected dimension %d, got %d", dim, got)
			}
		})
	}
}

// TestZeroDimension verifies that CreateTestIndex handles edge case dimensions.
func TestZeroDimension(t *testing.T) {
	// FAISS accepts dimension 0 (creates an empty index)
	ptr := CreateTestIndex(0)
	defer FreeIndex(ptr)

	d := GetIndexDimension(ptr)
	if d != 0 {
		t.Errorf("expected dimension 0, got %d", d)
	}
}

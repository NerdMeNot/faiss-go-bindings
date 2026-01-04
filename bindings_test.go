package bindings

import (
	"fmt"
	"testing"
)

// TestBindingsLink verifies that the FAISS C libraries are correctly linked
// by creating a simple index and checking its properties.
func TestBindingsLink(t *testing.T) {
	// Create a flat L2 index with dimension 64
	ptr, errCode := CreateTestIndex(64)
	if errCode != 0 {
		t.Fatalf("CreateTestIndex failed with code %d", errCode)
	}
	if ptr == 0 {
		t.Fatal("CreateTestIndex returned nil index")
	}
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
			ptr, errCode := CreateTestIndex(dim)
			if errCode != 0 {
				t.Fatalf("CreateTestIndex(%d) failed with code %d", dim, errCode)
			}
			if ptr == 0 {
				t.Fatalf("CreateTestIndex(%d) returned nil", dim)
			}
			defer FreeIndex(ptr)

			got := GetIndexDimension(ptr)
			if got != dim {
				t.Errorf("expected dimension %d, got %d", dim, got)
			}
		})
	}
}

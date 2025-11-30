package physics

import "testing"

func TestMemoryQuadtreeQuery(t *testing.T) {
	qt := NewMemoryQuadtree(Rect{X: 0, Y: 0, Width: 10, Height: 10})
	blobInside := NewBasicBlob("inside", Vector{X: 5, Y: 5})
	blobOutside := NewBasicBlob("outside", Vector{X: 20, Y: 20})

	qt.Insert(blobInside)
	qt.Insert(blobOutside)

	results := qt.Query(Rect{X: 0, Y: 0, Width: 10, Height: 10})
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].ID() != blobInside.ID() {
		t.Fatalf("unexpected blob returned: %s", results[0].ID())
	}
}

package physics

// Rect defines a simple rectangular region.
type Rect struct {
	X, Y          float64
	Width, Height float64
}

// Quadtree provides basic collision queries for blobs.
type Quadtree interface {
	Insert(blob Blob)
	Query(rangeRect Rect) []Blob
}

// MemoryQuadtree is a trivial in-memory implementation for scaffolding purposes.
type MemoryQuadtree struct {
	bounds Rect
	blobs  []Blob
}

// NewMemoryQuadtree constructs a MemoryQuadtree with the provided bounds.
func NewMemoryQuadtree(bounds Rect) *MemoryQuadtree {
	return &MemoryQuadtree{bounds: bounds, blobs: make([]Blob, 0)}
}

// Insert adds a blob to the quadtree without subdivision.
func (q *MemoryQuadtree) Insert(blob Blob) {
	q.blobs = append(q.blobs, blob)
}

// Query returns any blobs that fall within the provided range. This
// implementation uses a simple bounding check sufficient for the initial
// scaffolding.
func (q *MemoryQuadtree) Query(rangeRect Rect) []Blob {
	matches := make([]Blob, 0)
	for _, blob := range q.blobs {
		position := blob.Position()
		if position.X >= rangeRect.X && position.X <= rangeRect.X+rangeRect.Width &&
			position.Y >= rangeRect.Y && position.Y <= rangeRect.Y+rangeRect.Height {
			matches = append(matches, blob)
		}
	}
	return matches
}

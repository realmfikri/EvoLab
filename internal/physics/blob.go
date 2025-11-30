package physics

import "time"

// Vector is a simple 2D vector representation used by blob entities.
type Vector struct {
	X float64
	Y float64
}

// Blob represents an entity within the simulation that can be updated and
// queried for a position.
type Blob interface {
	ID() string
	Position() Vector
	Update(delta time.Duration)
}

// BasicBlob is a minimal blob implementation suitable for wiring tests.
type BasicBlob struct {
	id       string
	position Vector
	velocity Vector
}

// NewBasicBlob constructs a BasicBlob with the provided identifier and starting
// position.
func NewBasicBlob(id string, position Vector) *BasicBlob {
	return &BasicBlob{
		id:       id,
		position: position,
		velocity: Vector{X: 1, Y: 1},
	}
}

// ID returns the blob identifier.
func (b *BasicBlob) ID() string { return b.id }

// Position returns the current position.
func (b *BasicBlob) Position() Vector { return b.position }

// Update advances the blob position using its velocity and the supplied delta
// time.
func (b *BasicBlob) Update(delta time.Duration) {
	seconds := delta.Seconds()
	b.position.X += b.velocity.X * seconds
	b.position.Y += b.velocity.Y * seconds
}

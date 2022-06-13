package component

import (
	ray "github.com/gen2brain/raylib-go/raylib"
	lua "github.com/yuin/gopher-lua"
)

type Texture struct {
	Loaded ray.Texture2D
}

func (Texture) Name() string {
	return "texture"
}

type Model struct {
	Loaded ray.Model
	Scale  float32
}

func (Model) Name() string {
	return "model"
}

type Position struct {
	ray.Vector3
	Scale float32
}

func (Position) Name() string {
	return "position"
}

type Scriptable struct {
	Script *lua.LState
}

func (Scriptable) Name() string {
	return "scriptable"
}

// Acceleration describes an entity's rate at which it increases it's velocity.
type Acceleration struct {
	ray.Vector3
}

func (Acceleration) Name() string {
	return "acceleration"
}

// Colour represents a set of bytes to show colour on a display in an RGB format.
type Colour struct {
	Red, Green, Blue byte
}

func (Colour) Name() string {
	return "colour"
}

// Dimensions is a representation of the 2D geomtry that makes up an object in
// the game world.
type Dimensions struct {
	Width, Height, Radius float32
}

func (Dimensions) Name() string {
	return "dimensions"
}

// Gravity represents the amount of gravitational force the entity is under.
type Gravity struct {
	Force ray.Vector3
}

func (Gravity) Name() string {
	return "gravity"
}

// RigidBody is an entity with a solid body in which deformation is zero or so
// small it can be neglected. The distance between any two given points on a
// rigid body remains constant in time regardless of external forces exerted on
// it.
type RigidBody struct {
	Acceleration
	Velocity ray.Vector3
}

func (RigidBody) Name() string {
	return "rigid_body"
}

// Rotation describes an entity's angle transformation.
type Rotation struct {
	X, Y, Z int32
}

func (Rotation) Name() string {
	return "rotation"
}

// Transform describes an entity which has a position, rotation, and scale.
type Transform struct {
	Position
	Rotation
	Dimensions
	Scale float32
}

func (Transform) Name() string {
	return "transform"
}

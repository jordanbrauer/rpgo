package game

import (
	"github.com/bits-and-blooms/bitset"
	"github.com/jordanbrauer/rpgo/game/component"
	"github.com/jordanbrauer/rpgo/game/entity"
	"github.com/jordanbrauer/rpgo/game/system"
)

// CreateWorld returns a pointer to an empty world in memory.
func CreateWorld() World {
	var world = new(world)
	world.components = component.CreateComponentManager()
	world.entities = entity.CreateEntityManager()
	world.systems = system.CreateSystemManager()

	return world
}

// World instances act as coordinators between the entity, component, and system
// managers.
type World interface {
	Update(name string, dt float32)

	System(name string) system.System

	// Component will return a component interface for the given entity and
	// component name. The caller will have to use type assertions to extract
	// the real value from the result.
	Component(entity entity.Entity, name string) component.Component

	// Entity will return the signature for the given entity ID.
	Entity(entity entity.Entity) *bitset.BitSet

	// Entities is a total count of entities living in the current world space.
	Entities() int

	// Destroy will remove the entity and all of it's dedicated component/system
	// resources from the world, freeing up room for another (new) entity.
	Destroy(entity entity.Entity)

	// RegisterComponent will reserve a new component ID with the given name.
	RegisterComponent(name string)

	// RegisterComponent will reserve a new component ID with the given name.
	RegisterSystem(system system.System, components ...string)

	// CreateEntity will add a new living entity to the world and return it.
	CreateEntity() entity.Entity

	// SignEntity will set the given entity's signature for the given component names.
	SignEntity(entity entity.Entity, components ...string)

	// AttachComponent will assign the given component name and data to the entity.
	AttachComponent(entity entity.Entity, component component.Component)
}

type world struct {
	components component.ComponentManager
	entities   entity.EntityManager
	systems    system.SystemManager
}

func (world *world) Component(entity entity.Entity, name string) component.Component {
	return world.components.Read(entity, name)
}

func (world *world) Destroy(entity entity.Entity) {
	world.entities.Destroy(entity)
	world.systems.Destroy(entity)
	world.components.Destroy(entity)
}

func (world *world) Entity(entity entity.Entity) *bitset.BitSet {
	return world.entities.Read(entity)
}

func (world *world) Entities() int {
	return world.entities.Living()
}

func (world *world) CreateEntity() entity.Entity {
	return world.entities.Create()
}

func (world *world) RegisterComponent(name string) {
	world.components.Register(name)
}

func (world *world) RegisterSystem(system system.System, components ...string) {
	var name = system.Name()

	world.systems.Register(name, system)
	system.Updates(world)

	if len(components) > 0 {
		world.systems.Use(name, world.components.Sign(components...))
	}
}

func (world *world) SignEntity(entity entity.Entity, components ...string) {
	world.entities.Sign(entity, world.components.Sign(components...))
}

func (world *world) AttachComponent(entity entity.Entity, component component.Component) {
	var name = component.Name()

	world.components.Attach(entity, name, component)

	var signature = world.Entity(entity)

	if nil == signature {
		signature = new(bitset.BitSet)
	}

	signature.Set(uint(world.components.Signature(name)))
	world.entities.Sign(entity, signature)
	world.systems.Change(entity, signature)
}

func (world *world) System(name string) system.System {
	return world.systems.Read(name)
}

func (world *world) Update(name string, dt float32) {
	world.System(name).Update(dt)
}

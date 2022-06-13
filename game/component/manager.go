package component

import (
	"github.com/bits-and-blooms/bitset"
	"github.com/jordanbrauer/rpgo/game/entity"
)

// MaxComponents is the total amount of components that each entity is allowed
// to have.
const MaxComponents = 10

// CreateComponentManager will new up an empty manager with no components or
// signatures registered.
func CreateComponentManager() ComponentManager {
	var manager = new(componentManager)
	manager.components = make(map[string]*componentEntityMap)
	manager.signatures = make(map[string]int)

	return manager
}

// Component is a generic struct with data in them for systems to utilize.
type Component interface {
	Name() string
}

// ComponentManager takes care of creating, deleting, reading, and signing
// components to entities.
type ComponentManager interface {
	// Register will reserve a new ID by the given name for a component.
	Register(name string)

	// Read will return the component data by name for the given entity.
	Read(entity entity.Entity, name string) Component

	// Remove will delete the given component by name on the given entity.
	Remove(entity entity.Entity, name string)

	// Attach will assign a component data by name to the given entity.
	Attach(entity entity.Entity, name string, component Component)

	// Sign will create a signature for the given components by name that can be
	// assigned to an entity.
	Sign(names ...string) *bitset.BitSet

	// Signature will return the ID of a component by name.
	Signature(name string) int

	Destroy(entity entity.Entity)
}

type componentManager struct {
	next       int
	components map[string]*componentEntityMap
	signatures map[string]int
}

func (manager *componentManager) Destroy(entity entity.Entity) {
	for _, component := range manager.components {
		component.remove(entity)
	}
}

func (manager *componentManager) Attach(entity entity.Entity, name string, component Component) {
	manager.components[name].insert(entity, component)
}

func (manager *componentManager) Read(entity entity.Entity, name string) Component {
	return manager.components[name].read(entity)
}

func (manager *componentManager) Register(name string) {
	manager.signatures[name] = manager.next
	var components = new(componentEntityMap)
	components.entityComponents = make(map[int]Component)
	components.entityIndexMap = make(map[entity.Entity]int)
	components.indexEntityMap = make(map[int]entity.Entity)
	manager.components[name] = components

	manager.next++
}

func (manager *componentManager) Remove(entity entity.Entity, name string) {
}

func (manager *componentManager) Sign(names ...string) *bitset.BitSet {
	var signature = new(bitset.BitSet)

	for _, name := range names {
		signature.Set(uint(manager.Signature(name)))
	}

	return signature
}

func (manager *componentManager) Signature(name string) int {
	return manager.signatures[name]
}

type componentEntityMap struct {
	entityComponents map[int]Component
	entityIndexMap   map[entity.Entity]int
	indexEntityMap   map[int]entity.Entity
	size             int
}

func (pack *componentEntityMap) insert(entity entity.Entity, component Component) {
	var newIndex = pack.size
	pack.entityIndexMap[entity] = newIndex
	pack.indexEntityMap[newIndex] = entity
	pack.entityComponents[newIndex] = component

	pack.size++
}

func (pack *componentEntityMap) remove(entity entity.Entity) {
	delete(pack.entityComponents, pack.entityIndexMap[entity])
	// pack.entityComponents[pack.entityIndexMap[entity]]

	pack.size--
}

func (pack *componentEntityMap) read(entity entity.Entity) Component {
	return pack.entityComponents[pack.entityIndexMap[entity]]
}

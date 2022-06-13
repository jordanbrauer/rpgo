package system

import (
	"github.com/bits-and-blooms/bitset"
	"github.com/jordanbrauer/rpgo/game/component"
	"github.com/jordanbrauer/rpgo/game/entity"
)

type World interface {
	Component(entity entity.Entity, name string) component.Component
	Entities() int
}

// System describes the required functions to be implemented by a system.
type System interface {
	Update(dt float32)
	Updates(world World)
	Unsubscribe(entity entity.Entity)
	Subscribe(entity entity.Entity)
	Subscribed(entity entity.Entity) bool
	Name() string
}

type SystemAccess struct {
	entities []entity.Entity
	world    World
}

func (system *SystemAccess) Subscribed(entity entity.Entity) bool {
	for _, subscription := range system.entities {
		if entity == subscription {
			return true
		}
	}

	return false
}

func (system *SystemAccess) Subscribe(entity entity.Entity) {
	system.entities = append(system.entities, entity)
}

func (system *SystemAccess) Unsubscribe(entity entity.Entity) {
	// TODO
}

func (system *SystemAccess) Updates(world World) {
	system.world = world
}

func (system *SystemAccess) Component(entity entity.Entity, name string) component.Component {
	return system.world.Component(entity, name)
}

func (system *SystemAccess) Entities() []entity.Entity {
	return system.entities
}

func (system *SystemAccess) World() World {
	return system.world
}

func CreateSystemManager() SystemManager {
	var manager = new(systemManager)
	manager.signatures = make(map[string]*bitset.BitSet)
	manager.systems = make(map[string]System)

	return manager
}

type SystemManager interface {
	Register(name string, system System)
	Read(name string) System
	Destroy(entity entity.Entity)
	Change(entity entity.Entity, signature *bitset.BitSet)
	Use(name string, signature *bitset.BitSet)
}

type systemManager struct {
	signatures map[string]*bitset.BitSet
	systems    map[string]System
}

func (manager *systemManager) Register(name string, system System) {
	manager.systems[name] = system
}

func (manager *systemManager) Read(name string) System {
	return manager.systems[name]
}

func (manager *systemManager) Destroy(entity entity.Entity) {
	// TODO
}

func (manager *systemManager) Use(name string, signature *bitset.BitSet) {
	manager.signatures[name] = signature
}

func (manager *systemManager) Change(entity entity.Entity, signature *bitset.BitSet) {
	for name, system := range manager.systems {
		var systemSignature = manager.signatures[name]

		if systemSignature != nil && !system.Subscribed(entity) && signature.IsSuperSet(systemSignature) {
			system.Subscribe(entity)

			continue
		}

		system.Unsubscribe(entity)
	}
}

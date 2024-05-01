
package ecs

import "github.com/laykku/genecs/gen"

type System func(world *World, deltaTime float32)
type InitSystem func(world *World)

type World struct {
	store   *ComponentStore
	systems []System
	initSystems []InitSystem
	entities map[int]int
	container *Container
}

type Query struct {
	world *World
	entities []int
}

func (world *World) Tick(deltaTime float32) {
	initSystems := world.initSystems
	world.initSystems = make([]InitSystem, 0)
	for _, system := range initSystems {
		system(world)
	}

	for _, system := range world.systems {
		system(world, deltaTime)
	}
}

func CreateWorld() *World {
	return &World{
		store:   &ComponentStore{
		},
		systems: make([]System, 0),
		initSystems: make([]InitSystem, 0),
		entities: make(map[int]int),
		container: &Container{
			items: make([]interface{}, 0),
		},
	}
}

func (world *World) CreateEntity() int {
	for entity, compCount := range world.entities {
		if compCount == 0 {
			return entity
		}
	}
	return len(world.entities)
}

func (world *World) Register(system System) *World {
	world.systems = append(world.systems, system)
	return world
}

func (world *World) RegisterInit(system InitSystem) *World {
	world.initSystems = append(world.initSystems, system)
	return world
}

func (world *World) Stat() (int, int) {
	compCount := 0
	
	compCount += len(world.store.MouseMove)
	
	compCount += len(world.store.MoveInput)
	
	return len(world.entities), compCount
}

type ComponentStore struct {

MouseMove []MouseMove

MoveInput []MoveInput

}


func (component MouseMove) GetEntity() int {
return component.Entity
}
func (component MouseMove) IsOneFrame() bool {
	return component.OneFrame
	}

func (component MoveInput) GetEntity() int {
return component.Entity
}
func (component MoveInput) IsOneFrame() bool {
	return component.OneFrame
	}


func (world *World) Query() *Query {
	return &Query {
		world: world,
		entities: make([]int, 0),
	}
}


func (query *Query) MouseMove() *Query {
	query.entities = append(query.entities, 0)
	return query
}

func (query *Query) MoveInput() *Query {
	query.entities = append(query.entities, 1)
	return query
}


func (query *Query) Fetch() []int {
	var entities []int = nil
	for _, componentType := range query.entities {
		var newEntities []int = make([]int, 0)
		switch componentType {
			
		case 0:
			for i := 0; i < len(query.world.store.MouseMove); i++ {
				comp := &query.world.store.MouseMove[i]
				newEntities = append(newEntities, comp.GetEntity())
			}
			
		case 1:
			for i := 0; i < len(query.world.store.MoveInput); i++ {
				comp := &query.world.store.MoveInput[i]
				newEntities = append(newEntities, comp.GetEntity())
			}
			
		}
		entities = FilterQuery(entities, newEntities)
	}
	return entities
}

func FilterQuery(entities, newEntities []int) []int {
	filtered := make([]int, 0)
	if entities != nil {
		for _, v := range entities {
			if gen.Present(newEntities, v) {
				filtered = append(filtered, v)
			}
		}
	} else {
		filtered = newEntities
	}
	return filtered
}

	
	func (component MouseMove) Add(world *World) {
		world.store.MouseMove = append(world.store.MouseMove, component)
		world.entities[component.Entity]++
	}
	
	func (component MoveInput) Add(world *World) {
		world.store.MoveInput = append(world.store.MoveInput, component)
		world.entities[component.Entity]++
	}
	


func (world *World) GetMouseMove(entity int) *MouseMove {
	for i := 0; i < len(world.store.MouseMove); i++ {
		comp := &world.store.MouseMove[i]
		if comp.GetEntity() == entity {
			return comp
		}
	}
	return nil
}

func (world *World) GetMoveInput(entity int) *MoveInput {
	for i := 0; i < len(world.store.MoveInput); i++ {
		comp := &world.store.MoveInput[i]
		if comp.GetEntity() == entity {
			return comp
		}
	}
	return nil
}


    
	func (component *MouseMove) Remove(world *World) {
		for i := 0; i < len(world.store.MouseMove); i++ {
			comp := &world.store.MouseMove[i]
			if comp.GetEntity() == component.GetEntity() {
				world.store.MouseMove = append(world.store.MouseMove[:i], world.store.MouseMove[i+1:]...)
				world.entities[component.GetEntity()]--
			}
		}
	}
    
	func (component *MoveInput) Remove(world *World) {
		for i := 0; i < len(world.store.MoveInput); i++ {
			comp := &world.store.MoveInput[i]
			if comp.GetEntity() == component.GetEntity() {
				world.store.MoveInput = append(world.store.MoveInput[:i], world.store.MoveInput[i+1:]...)
				world.entities[component.GetEntity()]--
			}
		}
	}
    

type Container struct {
	items []interface{}
}

func (world *World) Inject(item interface{}) {
	world.container.items = append(world.container.items, item)
}

func Resolve[T interface{}](world *World) *T {
	for _, item := range world.container.items {
		if v, ok := item.(*T); ok {
			return v
		}
	}
	return nil
}
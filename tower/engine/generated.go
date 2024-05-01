
package engine

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
	
	compCount += len(world.store.TransformComponent)
	
	compCount += len(world.store.CameraComponent)
	
	compCount += len(world.store.MeshComponent)
	
	compCount += len(world.store.BatchList)
	
	return len(world.entities), compCount
}

type ComponentStore struct {

TransformComponent []TransformComponent

CameraComponent []CameraComponent

MeshComponent []MeshComponent

BatchList []BatchList

}


func (component TransformComponent) GetEntity() int {
return component.Entity
}
func (component TransformComponent) IsOneFrame() bool {
	return component.OneFrame
	}

func (component CameraComponent) GetEntity() int {
return component.Entity
}
func (component CameraComponent) IsOneFrame() bool {
	return component.OneFrame
	}

func (component MeshComponent) GetEntity() int {
return component.Entity
}
func (component MeshComponent) IsOneFrame() bool {
	return component.OneFrame
	}

func (component BatchList) GetEntity() int {
return component.Entity
}
func (component BatchList) IsOneFrame() bool {
	return component.OneFrame
	}


func (world *World) Query() *Query {
	return &Query {
		world: world,
		entities: make([]int, 0),
	}
}


func (query *Query) TransformComponent() *Query {
	query.entities = append(query.entities, 0)
	return query
}

func (query *Query) CameraComponent() *Query {
	query.entities = append(query.entities, 1)
	return query
}

func (query *Query) MeshComponent() *Query {
	query.entities = append(query.entities, 2)
	return query
}

func (query *Query) BatchList() *Query {
	query.entities = append(query.entities, 3)
	return query
}


func (query *Query) Fetch() []int {
	var entities []int = nil
	for _, componentType := range query.entities {
		var newEntities []int = make([]int, 0)
		switch componentType {
			
		case 0:
			for i := 0; i < len(query.world.store.TransformComponent); i++ {
				comp := &query.world.store.TransformComponent[i]
				newEntities = append(newEntities, comp.GetEntity())
			}
			
		case 1:
			for i := 0; i < len(query.world.store.CameraComponent); i++ {
				comp := &query.world.store.CameraComponent[i]
				newEntities = append(newEntities, comp.GetEntity())
			}
			
		case 2:
			for i := 0; i < len(query.world.store.MeshComponent); i++ {
				comp := &query.world.store.MeshComponent[i]
				newEntities = append(newEntities, comp.GetEntity())
			}
			
		case 3:
			for i := 0; i < len(query.world.store.BatchList); i++ {
				comp := &query.world.store.BatchList[i]
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

	
	func (component TransformComponent) Add(world *World) {
		world.store.TransformComponent = append(world.store.TransformComponent, component)
		world.entities[component.Entity]++
	}
	
	func (component CameraComponent) Add(world *World) {
		world.store.CameraComponent = append(world.store.CameraComponent, component)
		world.entities[component.Entity]++
	}
	
	func (component MeshComponent) Add(world *World) {
		world.store.MeshComponent = append(world.store.MeshComponent, component)
		world.entities[component.Entity]++
	}
	
	func (component BatchList) Add(world *World) {
		world.store.BatchList = append(world.store.BatchList, component)
		world.entities[component.Entity]++
	}
	


func (world *World) GetTransformComponent(entity int) *TransformComponent {
	for i := 0; i < len(world.store.TransformComponent); i++ {
		comp := &world.store.TransformComponent[i]
		if comp.GetEntity() == entity {
			return comp
		}
	}
	return nil
}

func (world *World) GetCameraComponent(entity int) *CameraComponent {
	for i := 0; i < len(world.store.CameraComponent); i++ {
		comp := &world.store.CameraComponent[i]
		if comp.GetEntity() == entity {
			return comp
		}
	}
	return nil
}

func (world *World) GetMeshComponent(entity int) *MeshComponent {
	for i := 0; i < len(world.store.MeshComponent); i++ {
		comp := &world.store.MeshComponent[i]
		if comp.GetEntity() == entity {
			return comp
		}
	}
	return nil
}

func (world *World) GetBatchList(entity int) *BatchList {
	for i := 0; i < len(world.store.BatchList); i++ {
		comp := &world.store.BatchList[i]
		if comp.GetEntity() == entity {
			return comp
		}
	}
	return nil
}


    
	func (component *TransformComponent) Remove(world *World) {
		for i := 0; i < len(world.store.TransformComponent); i++ {
			comp := &world.store.TransformComponent[i]
			if comp.GetEntity() == component.GetEntity() {
				world.store.TransformComponent = append(world.store.TransformComponent[:i], world.store.TransformComponent[i+1:]...)
				world.entities[component.GetEntity()]--
			}
		}
	}
    
	func (component *CameraComponent) Remove(world *World) {
		for i := 0; i < len(world.store.CameraComponent); i++ {
			comp := &world.store.CameraComponent[i]
			if comp.GetEntity() == component.GetEntity() {
				world.store.CameraComponent = append(world.store.CameraComponent[:i], world.store.CameraComponent[i+1:]...)
				world.entities[component.GetEntity()]--
			}
		}
	}
    
	func (component *MeshComponent) Remove(world *World) {
		for i := 0; i < len(world.store.MeshComponent); i++ {
			comp := &world.store.MeshComponent[i]
			if comp.GetEntity() == component.GetEntity() {
				world.store.MeshComponent = append(world.store.MeshComponent[:i], world.store.MeshComponent[i+1:]...)
				world.entities[component.GetEntity()]--
			}
		}
	}
    
	func (component *BatchList) Remove(world *World) {
		for i := 0; i < len(world.store.BatchList); i++ {
			comp := &world.store.BatchList[i]
			if comp.GetEntity() == component.GetEntity() {
				world.store.BatchList = append(world.store.BatchList[:i], world.store.BatchList[i+1:]...)
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
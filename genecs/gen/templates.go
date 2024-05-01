package gen

var codeTemplates = map[string]string{
	"world": `
package {{.Package}}

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
	{{range .Components}}
	compCount += len(world.store.{{.}})
	{{end}}
	return len(world.entities), compCount
}
`,

	"store": `
type ComponentStore struct {
{{range .Components}}
{{.}} []{{.}}
{{end}}
}
`,

	"component": `
{{range .Components}}
func (component {{.}}) GetEntity() int {
return component.Entity
}
func (component {{.}}) IsOneFrame() bool {
	return component.OneFrame
	}
{{end}}
`,

	"query": `
func (world *World) Query() *Query {
	return &Query {
		world: world,
		entities: make([]int, 0),
	}
}

{{range $i, $c := .Components}}
func (query *Query) {{$c}}() *Query {
	query.entities = append(query.entities, {{$i}})
	return query
}
{{end}}

func (query *Query) Fetch() []int {
	var entities []int = nil
	for _, componentType := range query.entities {
		var newEntities []int = make([]int, 0)
		switch componentType {
			{{range $i, $c := .Components}}
		case {{$i}}:
			for i := 0; i < len(query.world.store.{{$c}}); i++ {
				comp := &query.world.store.{{$c}}[i]
				newEntities = append(newEntities, comp.GetEntity())
			}
			{{end}}
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
`,

	"add": `
	{{range .Components}}
	func (component {{.}}) Add(world *World) {
		world.store.{{.}} = append(world.store.{{.}}, component)
		world.entities[component.Entity]++
	}
	{{end}}
`,

	"get": `
{{range .Components}}
func (world *World) Get{{.}}(entity int) *{{.}} {
	for i := 0; i < len(world.store.{{.}}); i++ {
		comp := &world.store.{{.}}[i]
		if comp.GetEntity() == entity {
			return comp
		}
	}
	return nil
}
{{end}}
`,

	"remove": `
    {{range .Components}}
	func (component *{{.}}) Remove(world *World) {
		for i := 0; i < len(world.store.{{.}}); i++ {
			comp := &world.store.{{.}}[i]
			if comp.GetEntity() == component.GetEntity() {
				world.store.{{.}} = append(world.store.{{.}}[:i], world.store.{{.}}[i+1:]...)
				world.entities[component.GetEntity()]--
			}
		}
	}
    {{end}}
`,
	"container": `
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
}`,
}

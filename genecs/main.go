package main

import (
	"fmt"

	"github.com/laykku/genecs/ecs"
)

func main() {
	run()
}

func run() {
	world := ecs.CreateWorld()

	e := world.CreateEntity()

	ecs.Person{
		Entity: e,
		Name:   "Bob",
	}.Add(world)

	ecs.Greeter{
		Entity:   e,
		Greeting: "Hello, ",
	}.Add(world)

	world.Register(GreetingSystem)

	for {
		world.Tick(0)
	}
}

func GreetingSystem(world *ecs.World, deltaTime float32) {
	q := world.Query().Greeter().Person().Fetch()
	for _, e := range q {
		p := world.GetPerson(e)
		g := world.GetGreeter(e)
		fmt.Printf("%s%s!\n", g.Greeting, p.Name)
		greeter := world.GetGreeter(e)
		greeter.Remove(world)
	}
}

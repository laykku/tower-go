package main

import (
	"testing"

	"github.com/laykku/genecs/ecs"
)

func TestAddComponent(t *testing.T) {
	world := ecs.CreateWorld()
	entity := world.CreateEntity()
	ecs.Person{Entity: entity, Name: "Bob"}.Add(world)
	person := world.GetPerson(entity)
	if person == nil {
		t.Fatal("person is nil")
	}
}

func TestRemoveComponent(t *testing.T) {
	world := ecs.CreateWorld()
	entity := world.CreateEntity()
	ecs.Person{Entity: entity, Name: "Bob"}.Add(world)

	q := world.Query().Person().Fetch()

	if len(q) != 1 {
		t.Fatal("len(q) != 1")
	}

	world.GetPerson(q[0]).Remove(world)
	q = world.Query().Person().Fetch()

	if len(q) != 0 {
		t.Fatal("len(q) != 0")
	}
}

func TestRemoveComponentDoesntAffectOtherComponents(t *testing.T) {
	world := ecs.CreateWorld()
	entity := world.CreateEntity()
	ecs.Person{Entity: entity, Name: "Bob"}.Add(world)
	ecs.Greeter{OneFrame: true, Entity: entity, Greeting: "Hello"}.Add(world)

	q := world.Query().Person().Fetch()

	if len(q) != 1 {
		t.Fatal("len(q) != 1")
	}

	world.GetPerson(q[0]).Remove(world)

	q = world.Query().Greeter().Fetch()

	if len(q) != 1 {
		t.Fatal("len(q) != 1")
	}
}

type SomeService struct {
	name string
}

func TestContainer(t *testing.T) {
	world := ecs.CreateWorld()
	ecs.Inject(world, &SomeService{
		name: "Some service",
	})

	s := ecs.Resolve[SomeService](world)

	if s == nil {
		t.Fatal("Service not resolved")
	}

	if s.name != "Some service" {
		t.Fatal("Wrong service resolved")
	}

	s.name = "Test"

	if s.name != "Test" {
		t.Fatal("Error changing service state")
	}
}

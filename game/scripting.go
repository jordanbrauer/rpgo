package game

import (
	lua "github.com/yuin/gopher-lua"
)

type Library struct {
	World
}

func (lib Library) Spawn(state *lua.LState) int {
	entity := Spawn(
		lib.World,
		state.ToString(1),
		state.ToString(2),
		float32(state.ToNumber(3)),
		float32(state.ToNumber(4)),
		float32(state.ToNumber(5)),
		float32(state.ToNumber(6)),
	)

	state.Push(lua.LNumber(entity))

	return 1
}

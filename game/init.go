package game

import (
	"github.com/hajimehoshi/ebiten/mobile"
	"github.com/reconquest/mrdig/game/engine"
)

func Init() {}

func init() {
	mobile.SetGame(engine.New())
}

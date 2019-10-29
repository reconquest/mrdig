package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Engine struct{}

func New() *Engine {
	return &Engine{}
}

func (engine *Engine) Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\n", ebiten.CurrentTPS()),
	)

	return nil
}

func (engine *Engine) Layout(width int, height int) (int, int) {
	return width, height
}

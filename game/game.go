package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/mobile"
)

type game struct{}

func (game *game) Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\n", ebiten.CurrentTPS()),
	)

	return nil
}

func (game *game) Layout(width int, height int) (int, int) {
	return width, height
}

func init() {
	mobile.SetGame(&game{})
}

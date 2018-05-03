package world
/**
usage: World is everything withing the screen
and in times it will update the position of every
object registered
 */
import (
	"github.com/hajimehoshi/ebiten"
	"go/types"
)

type Worlder interface{
	Update()
}

type Hitter interface{
	Update()
}

type World struct {
	screen *ebiten.Image
	objects types.Slice
}

func (world *World) Update()  {

	fallDown()
}

func (world *World) Hit()  {

	fallDown()
}


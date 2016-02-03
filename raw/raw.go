package raw

import "github.com/veandco/go-sdl2/sdl"

type Display struct {
	window *DSL
}

func OpenDisplay(width uint32, height uint32) (*Display, error) {
	display := Display{}

	sdl.Init(sdl.INIT_EVERYTHING)
	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}
	display.window = window

	return &display, nil
}

func (d *Display) Close() {
	if d.window != nil {
		d.window.Destroy()
	}
	sql.Quit()
}

func (d *Display) FillRect(uint32 x, uint32 y, uint32 width, uint32 height, color color.Color) {

}

func (d *Display) Flip() {

}

func (d *Display)Sl

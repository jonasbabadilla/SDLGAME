package levels

import (
	"github.com/veandco/go-sdl2/sdl"
)

// cutscene
func LevelNine(renderer *sdl.Renderer) (levelData []Object, LevelBG Object, pStart StartData, err error) {

	Surf, _ := sdl.LoadBMP("LEVELS/LevelNineSprites/levelLayout.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)

	levelData = CreateLevel(Surf, Tex)

	defer Surf.Free()

	Surf, _ = sdl.LoadBMP("LEVELS/LevelNineSprites/BG.bmp")
	BG, _ := renderer.CreateTextureFromSurface(Surf)

	backgroundData = Object{
		Tex:          BG,
		X:            0,
		Y:            0,
		ObjectWidth:  1280,
		ObjectHeight: 720,
	}

	defer Surf.Free()

	return levelData, backgroundData, PlayerStart, nil

}

package levels

import (
	"github.com/veandco/go-sdl2/sdl"
)

func LevelFive(renderer *sdl.Renderer) (levelData []Object, LevelBG Object, PlayerStart StartData, err error) {

	Surf, _ := sdl.LoadBMP("LEVELS/LevelFiveSprites/levelLayout.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)

	levelData = CreateLevel(Surf, Tex)

	defer Surf.Free()

	Surf, _ = sdl.LoadBMP("LEVELS/LevelFiveSprites/BG.bmp")
	BG, _ := renderer.CreateTextureFromSurface(Surf)

	backgroundData = Object{
		Tex:          BG,
		X:            0,
		Y:            0,
		ObjectWidth:  1280,
		ObjectHeight: 720,
	}

	defer Surf.Free()

	PlayerStart = StartData{X: 624, Y: 624, EndData: struct {
		X int
		Y int
	}{X: 593, Y: 185}}

	return levelData, backgroundData, PlayerStart, nil

}
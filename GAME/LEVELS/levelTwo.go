package levels

// cutscene
import (
	"github.com/veandco/go-sdl2/sdl"
)

func LevelTwo(renderer *sdl.Renderer) (levelData []Object, LevelBG Object, PlayerStart StartData, err error) {

	Surf, _ := sdl.LoadBMP("LEVELS/LevelTwoSprites/levelLayout.bmp")
	Tex, _ := renderer.CreateTextureFromSurface(Surf)

	levelData = CreateLevel(Surf, Tex)

	defer Surf.Free()

	Surf, _ = sdl.LoadBMP("LEVELS/LevelTwoSprites/BG.bmp")
	BG, _ := renderer.CreateTextureFromSurface(Surf)

	backgroundData = Object{
		Tex:          BG,
		X:            0,
		Y:            0,
		ObjectWidth:  1280,
		ObjectHeight: 720,
	}

	defer Surf.Free()

	PlayerStart = StartData{X: 10, Y: 588, EndData: struct {
		X int
		Y int
	}{X: 1184, Y: 652}}

	return levelData, backgroundData, PlayerStart, nil

}

package main

import (
	levels "chaseGame/GAME/LEVELS"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var Renderer *sdl.Renderer

const (
	screenWidth  = 1280
	screenHeight = 720
)

var ObjectData []levels.Object
var backgroundData levels.Object
var levelType levels.LevelUpdate

var pChar player
var delay int
var visible bool
var switched bool

func checkCollision() {
	//Check if player is on same Y level
	for i := range ObjectData {
		if pChar.y+float64(pChar.playerHeight*4) >= float64(ObjectData[i].Y) && pChar.y+float64(pChar.playerHeight*4) < float64(ObjectData[i].Y+(ObjectData[i].ObjectHeight/2)) {

			if pChar.x >= float64(ObjectData[i].X-16) && pChar.x <= float64(ObjectData[i].X+ObjectData[i].ObjectWidth-16) {
				pChar.y -= float64(gravity)
				if Keys[sdl.SCANCODE_UP] == 0 {
					JumpState = false
					JumpTimer = 0
					PlayerSpeedY = 7.00
					gravity = 5.00
				}
			}
		}
	}

	if pChar.y+float64(pChar.playerHeight*4) >= screenHeight {
		pChar.x = float64(pStart.X)
		pChar.y = float64(pStart.Y)
	}

	if pChar.x >= float64(pStart.EndData.X-16) {
		if Keys[sdl.SCANCODE_SPACE] == 1 {
			if currentLvl+1 <= totalLvl {
				currentLvl += 1
			}
			ObjectData, backgroundData, pStart, textData = decideLevel()
			pChar.x = float64(pStart.X)
			pChar.y = float64(pStart.Y)
		}

	}

}

func main() {
	if err := sdl.Init(uint32(sdl.INIT_EVERYTHING)); err != nil {
		fmt.Println("initializing sdl:", err)
		return
	}
	window, err := sdl.CreateWindow(
		"The Horizon",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()
	Renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing Renderer:", err)
		return
	}
	defer Renderer.Destroy()

	ObjectData, backgroundData, pStart, textData = decideLevel()

	pChar, err = NewPlayer(Renderer)
	if err != nil {
		fmt.Println("creating player:", err)
		return
	}

	BotDirection = sdl.FLIP_HORIZONTAL
	visible = true

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		Renderer.SetDrawColor(144, 144, 144, 0)
		Renderer.Clear()

		//levels.ObjectData[0].Draw(Renderer, levels.ObjectData)
		Draw(Renderer, ObjectData, textData)

		pChar.Draw(Renderer)
		switch currentLvl {
		case 2:
			Update()
		case 5:
			pChar.DrawTwo(Renderer)
		case 6:
			Update()
			pChar.DrawTwo(Renderer)
		case 7:
			pChar.DrawTwo(Renderer)
		case 8:
			pChar.DrawTwo(Renderer)
		case 12:
			pChar.DrawTwo(Renderer)
		case 13:
			pChar.DrawTwo(Renderer)
		case 15:
			pChar.FinalDraw(Renderer)
		}
		pChar.Update()

		Renderer.Present()
		checkCollision()

	}
}

func Update() {

	if visible == true {
		delay++
		if delay >= 100 {
			BotX += 2
			BotDirection = sdl.FLIP_NONE
		}
		pChar.BotOne(Renderer)
	}

	if visible == false {
		if pChar.x > 640 {
			visible = true
		}
	}

}

func Draw(Renderer *sdl.Renderer, ObjectData []levels.Object, textData levels.Object) {

	backgroundData.Tex.SetBlendMode(sdl.BLENDMODE_BLEND)
	Renderer.Copy(backgroundData.Tex, nil, nil)

	for _, k := range ObjectData {
		Renderer.Copy(k.Tex,
			&sdl.Rect{X: int32(k.X), Y: int32(k.Y), W: int32(k.ObjectWidth), H: int32(k.ObjectHeight)},
			&sdl.Rect{X: int32(k.X), Y: int32(k.Y), W: int32(k.ObjectWidth), H: int32(k.ObjectHeight)},
		)
	}

	Renderer.Copy(
		textData.Tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(textData.ObjectWidth), H: int32(textData.ObjectHeight)},
		&sdl.Rect{X: 100, Y: 100, W: int32(textData.ObjectWidth), H: int32(textData.ObjectHeight)},
	)

}

///DHFSGFYSDUHATFYSAGFDYSAUGFFSJKAGFSAFDGSAF

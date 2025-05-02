package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

func drawString(screen tcell.Screen, x, y int, msg string) {
	for idx, char := range msg {
		screen.SetContent(x+idx, y, char, nil, tcell.StyleDefault)
	}

}

func main() {

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()

	err = screen.Init()
	if err != nil {
		log.Fatal(err)
	}

	// game init
	player := NewSprite('#', 10, 10)

	// game loop
	running := true
	for running {
		// draw
		screen.Clear()
		player.Draw(screen)


		// ui
		drawString(
			screen,
			5,
			2,
			fmt.Sprintf("Hello Kuba %d.", 5),
		)
		screen.Show()

		// update
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Rune() {
			case 'q':
				running = false
			case 'w':
				player.Y -= 1
			case 's':
				player.Y += 1
			case 'a':
				player.X -= 1
			case 'd':
				player.X += 1
			}
		}


	}

}

package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
)


func drawString(screen tcell.Screen, x, y int, msg string, style tcell.Style) {
	for idx, char := range msg {
		screen.SetContent(x+idx, y, char, nil, style)
	}
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()

	if err := screen.Init(); err != nil {
		log.Fatal(err)
	}

	game := NewGameState(screen)
	quit := make(chan struct{})
	events := make(chan tcell.Event)

	// Event polling goroutine
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				ev := game.screen.PollEvent()
				select {
				case events <- ev:
				case <-quit:
					return
				}
			}
		}
	}()

	// Game timer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				game.Update()
				select {
				case game.redraw <- struct{}{}: // Request redraw
				default: // Skip if redraw already pending
				}
			case <-quit:
				return
			}
		}
	}()

	// Main game loop
	running := true
	for running {
		// Non-blocking check for redraw requests
		select {
		case ev := <-events: // Receive from channel
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Rune() == 'q' {
					running = false
				}
				if ev.Rune() == '1' || ev.Rune() == '+' {
					game.Buy(1)
				} else if ev.Rune() == '2' || ev.Rune() == 'ě' {
					game.Buy(2)
				} else if ev.Rune() == '3' || ev.Rune() == 'š' {
					game.Buy(3)
				} else if ev.Rune() == '4' || ev.Rune() == ';' {
					game.Buy(4)
				}
			}
		case <-game.redraw: // Redraw channel
			game.Draw()
		case <-time.After(33 * time.Millisecond): // ~30 FPS
			// Optional actoins
		}
	}

	close(quit)
}

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
		ticker := time.NewTicker(1000 * time.Millisecond)
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
				if ev.Rune() == 'i' {
					game.Buy(game.selectedIdx)
				} else if ev.Rune() == 'j' {
					game.selectedIdx++
					if game.selectedIdx >= game.visibleItems || game.selectedIdx >= len(game.menuItems) {
						game.selectedIdx = 0
					}
				} else if ev.Rune() == 'k' {
					game.selectedIdx--
					if game.selectedIdx < 0 {
						game.selectedIdx = min(game.visibleItems, len(game.menuItems))-1
					}
				}
				
				// Request redraw
				select {
				case game.redraw <- struct{}{}: // Request redraw
				default: // Skip if redraw already pending
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

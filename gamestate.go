package main

import (
	"fmt"
	"sync/atomic"

	"github.com/gdamore/tcell/v2"
)

// GameState holds all game state with atomic counters for thread safety
type GameState struct {
	total_points int64 // Use atomic package for concurrent access
	base_points  int64 // Use atomic package for concurrent access
	screen       tcell.Screen
	redraw       chan struct{} // Channel to force screen redraws

	// add menu with upgrades
	menuItems    []*MenuItem
	visibleItems int
	page         int
	selectedIdx  int
}

func NewGameState(screen tcell.Screen) *GameState {
	return &GameState{
		base_points:  5,
		total_points: 0,
		redraw:       make(chan struct{}, 1), // Buffered channel to prevent blocking
		screen:       screen,

		menuItems:    InitalizeItems(),
		visibleItems: 2,
		page:         1,
		selectedIdx:  0,
	}
}

func (gs *GameState) DrawMenu(y int) {
	first_idx := (gs.selectedIdx / 10) * 10
	last_idx := min(first_idx+10, len(gs.menuItems))
	for i, item := range gs.menuItems[first_idx:last_idx] {
		if i >= gs.visibleItems {
			break
		}

		style := tcell.StyleDefault
		description := ""
		if !item.IsUnlocked {
			style = style.Foreground(tcell.ColorGray)
			description = fmt.Sprintf("(Earn %d to unlock)", item.UnlockScore)
		} else {
			description = fmt.Sprintf("(%d)", item.Count)
		}

		// Highlight seleted item
		if gs.selectedIdx%10 == i {
			style = style.Background(tcell.ColorBlue)
			if !item.IsUnlocked {
				style = style.Background(tcell.ColorLightPink)
			}

			width, _ := gs.screen.Size()
			for x := 2; x < width-30; x++ {
				gs.screen.SetContent(x, y+i, ' ', nil, style)
			}
		}

		// name
		text := fmt.Sprintf(
			"[%d] %s",
			first_idx+i+1,
			item.Name,
		)
		drawString(
			gs.screen,
			2,
			y+i,
			text,
			style,
		)
		// Count or locked
		drawString(
			gs.screen,
			30,
			y+i,
			description,
			style,
		)
		if item.IsUnlocked {
			// Cost
			drawString(
				gs.screen,
				60,
				y+i,
				fmt.Sprintf("%d", item.Cost),
				style,
			)
			// Amount produced per second
			drawString(
				gs.screen,
				70,
				y+i,
				fmt.Sprintf("%d/s", item.ScorePerSecond),
				style,
			)
		}

	}

}

func (gs *GameState) Draw() {
	gs.screen.Clear()

	// UI Elements
	gs.DrawMenu(4)
	drawString(
		gs.screen,
		10,
		2,
		fmt.Sprintf("Loudness Points: %d", atomic.LoadInt64(&gs.base_points)),
		tcell.StyleDefault.Foreground(tcell.ColorYellow),
	)
	// drawString(
	// 	gs.screen,
	// 	10,
	// 	1,
	// 	fmt.Sprintf("Total acumulated: %d", atomic.LoadInt64(&gs.total_points)),
	// 	tcell.StyleDefault.Foreground(tcell.ColorLightGray),
	// )
	gs.screen.Show()
}

func (gs *GameState) Buy(idx int) {
	// idx = idx - 1
	// Out of range or not unlocked yet
	if idx >= len(gs.menuItems) || idx >= gs.visibleItems-1 {
		return
	}

	item := gs.menuItems[idx]
	if item.Cost > gs.base_points || !item.IsUnlocked {
		return
	}

	atomic.AddInt64(&gs.base_points, -item.Cost)
	atomic.AddInt64(&item.Count, 1)

	increment := float64(item.Cost) * 0.1
	atomic.AddInt64(&item.Cost, max(1, int64(increment)))

	gs.menuItems[idx] = item
}

func (gs *GameState) Update() {
	updateAmount := int64(0)
	for i, item := range gs.menuItems {
		if i >= gs.visibleItems {
			break
		}

		updateAmount += item.Count * item.ScorePerSecond
	}

	atomic.AddInt64(&gs.base_points, updateAmount)
	atomic.AddInt64(&gs.total_points, updateAmount)

	if gs.visibleItems <= len(gs.menuItems) {
		if gs.base_points >= gs.menuItems[gs.visibleItems-1].UnlockScore {
			gs.menuItems[gs.visibleItems-1].IsUnlocked = true
			gs.visibleItems += 1
		}
	}

	select {
	case gs.redraw <- struct{}{}: // Request redraw
	default: // Skip if redraw already pending
	}
}

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
}

func NewGameState(screen tcell.Screen) *GameState {
	return &GameState{
		base_points: 1000,
		total_points: 0,
		redraw:      make(chan struct{}, 1), // Buffered channel to prevent blocking
		screen:      screen,

		menuItems:    InitalizeItems(),
		visibleItems: 2,
	}
}

func (gs *GameState) DrawMenu(y int) {
	for i, item := range gs.menuItems {
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

		// name
		text := fmt.Sprintf(
			"[%d] %s",
			i+1,
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
		fmt.Sprintf("Calculation Points: %d", atomic.LoadInt64(&gs.base_points)),
		tcell.StyleDefault.Foreground(tcell.ColorYellow),
	)
	drawString(
		gs.screen,
		10,
		1,
		fmt.Sprintf("Total acumulated: %d", atomic.LoadInt64(&gs.total_points)),
		tcell.StyleDefault.Foreground(tcell.ColorLightGray),
	)
	gs.screen.Show()
}

func (gs *GameState) Buy(idx int) {
	idx = idx - 1
	// Out of range or not unlocked yet
	if idx >= len(gs.menuItems) || idx >= gs.visibleItems-1 {
		return
	}

	item := gs.menuItems[idx]
	if item.Cost > gs.base_points {
		return
	}

	atomic.AddInt64(&gs.base_points, -item.Cost)
	atomic.AddInt64(&item.Count, 1)

	gs.menuItems[idx] = item

	select {
	case gs.redraw <- struct{}{}: // Request redraw
	default: // Skip if redraw already pending
	}
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
		if gs.total_points >= gs.menuItems[gs.visibleItems-1].UnlockScore {
			gs.menuItems[gs.visibleItems-1].IsUnlocked = true
			gs.visibleItems += 1
		}
	}

	select {
	case gs.redraw <- struct{}{}: // Request redraw
	default: // Skip if redraw already pending
	}
}

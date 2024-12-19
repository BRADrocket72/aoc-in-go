package main

import (
	"image"
	"testing"
)

func Test_moveZero(t *testing.T) {
	bot := bot{position: image.Point{0, 0}, velocity: image.Point{1, 1}}
	if bot.position.X != 0 || bot.position.Y != 0 {
		t.Error("test error")
	}

}
func Test_moveOnce(t *testing.T) {
	bot := bot{position: image.Point{0, 0}, velocity: image.Point{1, 1}}
	bot = moveBot(bot, 11, 7, 1, true)
	if bot.position.X != 1 || bot.position.Y != 1 {
		t.Error("test error")
	}

}

func Test_moveOnce_edge(t *testing.T) {
	bot := bot{position: image.Point{11, 7}, velocity: image.Point{1, 1}}
	bot = moveBot(bot, 11, 7, 1, true)
	if bot.position.X != 0 || bot.position.Y != 0 {
		t.Error("failed edge move", bot.position)
	}

}

func Test_moveOnce_edge2(t *testing.T) {
	bot := bot{position: image.Point{0, 0}, velocity: image.Point{-1, -1}}
	bot = moveBot(bot, 11, 7, 1, true)
	if bot.position.X != 11 || bot.position.Y != 7 {
		t.Error("failed edge move", bot.position)
	}

}

//227268405 too high

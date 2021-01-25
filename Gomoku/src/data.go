package main

import (
	"time"

	"github.com/faiface/pixel/text"

	_ "image/jpeg"

	"github.com/faiface/pixel"
)

var IATimer time.Duration
var timer time.Time
var mycfg config
var state gameState
var redraw bool
var transpaDrawn bool
var lastpy int8
var lastpx int8
var X float64
var Y float64
var lastX float64
var lastY float64
var lastTurnState []gameState
var second <-chan time.Time
var deepth int
var width int
var routine bool
var hinter hintInfo
var saveNotif notif
var showCapturable bool
var winPos [2]int8
var loosePos [2]int8
var iavsia bool

var NodeAddresses map[int][]*Node

type config struct {
	Fullscreen  bool
	RefreshRate float64
	Smooth      bool
	Vsync       bool
	X           float64
	Y           float64
	Sound       bool
}

type notif struct {
	Draw  bool
	Text  string
	Timer time.Time
}

type saveGame struct {
	State         gameState
	Timer         time.Time
	LastTurnState []gameState
}

type hintInfo struct {
	Draw      bool
	X         int8
	Y         int8
	Timestamp time.Time
}

type gameState struct {
	Menu             bool
	IA               bool
	TwoP             bool
	Capture          bool
	DoubleThree      bool
	Standard         bool
	Pro              bool
	Long             bool
	Board            [19][19]int8
	CapturableBoard  [19][19]int8
	ThreeBoard       [19][19]int8
	Turn             int
	CurrentPlayer    int8
	LastPlayed       [2]int8
	CapturePlayerOne int8
	CapturePlayerTwo int8
	WinOne           bool
	WinTwo           bool
	WinConditionOne  bool
	WinConditionTwo  bool
	Winnable         []winInProgress
	StonesPlayed     [][2]int8
}

type alignementHeuristic struct {
	Length int8
	Player int8
	Start  [2]int8
	End    [2]int8
}

type winInProgress struct {
	Player int8
	Start  [2]int8
	End    [2]int8
}

type chanMove struct {
	One   chan bool
	Two   chan bool
	Three chan bool
	Four  chan bool
	Five  chan bool
	Six   chan bool
	Seven chan bool
	Eight chan bool
}

type chanSprite struct {
	Menu             chan *pixel.Sprite
	Board            chan *pixel.Sprite
	Off              chan *pixel.Sprite
	On               chan *pixel.Sprite
	Whitestone       chan *pixel.Sprite
	Blackstone       chan *pixel.Sprite
	Goldstone        chan *pixel.Sprite
	RedStone         chan *pixel.Sprite
	BlueStone        chan *pixel.Sprite
	Transpblackstone chan *pixel.Sprite
	Transpwhitestone chan *pixel.Sprite
	Soundoff         chan *pixel.Sprite
	Soundon          chan *pixel.Sprite
	Winplayerone     chan *pixel.Sprite
	Winplayertwo     chan *pixel.Sprite
	Winia            chan *pixel.Sprite
	Wood             chan *pixel.Sprite
	LogoSoundOn      chan *pixel.Sprite
	LogoSoundOff     chan *pixel.Sprite
	Atlas            chan *text.Atlas
}

type Sprite struct {
	Menu                   *pixel.Sprite
	Board                  *pixel.Sprite
	Off                    *pixel.Sprite
	On                     *pixel.Sprite
	SoundOn                *pixel.Sprite
	SoundOff               *pixel.Sprite
	BlackStone             *pixel.Sprite
	WhiteStone             *pixel.Sprite
	Goldstone              *pixel.Sprite
	RedStone               *pixel.Sprite
	BlueStone              *pixel.Sprite
	TransparencyWhiteStone *pixel.Sprite
	TransparencyBlackStone *pixel.Sprite
	WinPlayerOne           *pixel.Sprite
	WinPlayerTwo           *pixel.Sprite
	WinIA                  *pixel.Sprite
	Wood                   *pixel.Sprite
	LogoSoundOff           *pixel.Sprite
	LogoSoundOn            *pixel.Sprite
	Atlas                  *text.Atlas
}

type Node struct {
	Pos              [2]int8
	CurrentPlayer    int8
	Alpha            int
	Beta             int
	Board            [19][19]int8
	ThreeBoard       [19][19]int8
	CapturableBoard  [19][19]int8
	StonesPlayed     [][2]int8
	CapturePlayerOne int8
	CapturePlayerTwo int8
	Children         []*Node
	Cut              bool
}

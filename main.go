package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000;
	screenHeight = 450;
)

var (
	running = true

  bgColor = rl.NewColor(147, 211, 196, 255);

	grassSprite  rl.Texture2D;
	playerSprite rl.Texture2D;

	playerSrc  rl.Rectangle;
	playerDest rl.Rectangle;

  playerSpeed float32 = 3.0;

  isMusicPlaying = true; 
  bgMusic rl.Music 
  bgMusicVolume float32 = 0.65; 
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White);
  rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White);
}

func input() {
  if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
    playerDest.Y-=playerSpeed;
  }

  if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyUp) {
    playerDest.Y+=playerSpeed;
  }

  if rl. IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
    playerDest.X-=playerSpeed;
  }

  if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
    playerDest.X+=playerSpeed;
  }

  if rl.IsKeyPressed(rl.KeyQ) {
    isMusicPlaying = !isMusicPlaying; 
  }
  
  if rl.IsKeyPressed(rl.KeyMinus) {
    bgMusicVolume = float32(math.Max(0, float64(bgMusicVolume-0.1))); 
  } else if rl.IsKeyPressed(rl.KeyEqual) {
    bgMusicVolume = float32(math.Min(1, float64(bgMusicVolume+0.1)));
  }
}

func update() {
  running = !rl.WindowShouldClose();

  rl.UpdateMusicStream(bgMusic); 
  if isMusicPlaying {
    rl.ResumeMusicStream(bgMusic);
  } else {
    rl.PauseMusicStream(bgMusic);
  }

  rl.SetMusicVolume(bgMusic, bgMusicVolume);
}

func render() {
	rl.BeginDrawing();
	rl.ClearBackground(bgColor);
	drawScene();
	rl.EndDrawing();
}

func quit() {
	rl.UnloadTexture(playerSprite);
	rl.UnloadTexture(grassSprite);
  rl.UnloadMusicStream(bgMusic); 
  rl.CloseAudioDevice();
	rl.CloseWindow();
}

func init() {
	rl.InitWindow(screenWidth, screenHeight, "Celeste");
	rl.SetExitKey(0);
	rl.SetTargetFPS(60);

	grassSprite = rl.LoadTexture("assets/tiles/Grass.png");
	playerSprite = rl.LoadTexture("assets/characters/Basic Charakter Spritesheet.png");

	playerSrc = rl.NewRectangle(0, 0, 48, 48);
	playerDest = rl.NewRectangle(200, 200, 100, 100);

  rl.InitAudioDevice();
  bgMusic = rl.LoadMusicStream("assets/averys-farm.mp3");
  rl.PlayMusicStream(bgMusic);
}

func main() {
  defer quit();

	for running {
		input();
		update();
		render();
	}
}

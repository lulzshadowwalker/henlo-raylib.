package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 1000;
	screenHeight = 450;

  movingUp = 0x0001;
  movingDown = 0x0010;
  movingRight = 0x0100;
  movingLeft = 0x1000;
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
  bgMusic rl.Music;
  bgMusicVolume float32 = 0.65; 

  camera rl.Camera2D;

  frameCount = 0;
  frameSpeed = 8; 
  playerFrame = 0; 

  movementDir = 0x0000;

  playerDir = 1; // 1 Down, 0 Up, 2 Left, 3 Right based on the spritesheet
)

func drawScene() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White);
  rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White);
}

func input() {
  if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
    movementDir |= movingUp;
    playerDir = 1; 
  }

  if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
    movementDir |= movingDown;
    playerDir = 0; 
  }

  if rl. IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
    movementDir |= movingLeft;
    playerDir = 2;
  }

  if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
    movementDir |= movingRight;
    playerDir = 3; 
  }

  if rl.IsKeyPressed(rl.KeyQ) {
    isMusicPlaying = !isMusicPlaying; 
  }
  
  if rl.IsKeyPressed(rl.KeyMinus) {
    bgMusicVolume = float32(math.Max(0, float64(bgMusicVolume-0.1))); 
    rl.SetMusicVolume(bgMusic, bgMusicVolume);
 } else if rl.IsKeyPressed(rl.KeyEqual) {
    bgMusicVolume = float32(math.Min(1, float64(bgMusicVolume+0.1)));
    rl.SetMusicVolume(bgMusic, bgMusicVolume);
  }
}

func update() {
  running = !rl.WindowShouldClose();

  frameCount++; 

  if movementDir != 0 {
    if movementDir&movingUp != 0 { 
      playerDest.Y -= playerSpeed;
    }
    if movementDir&movingDown != 0 {
      playerDest.Y += playerSpeed;
    }

    if movementDir&movingRight != 0 { 
      playerDest.X += playerSpeed;
    } 
    if movementDir&movingLeft != 0 {
      playerDest.X -= playerSpeed;
    }

    if frameCount % 8 == 1 { playerFrame++ };
    if playerFrame > 3 { playerFrame = 0; }
  } else {
    if frameCount % 45 == 1 { playerFrame++; }
    if playerFrame > 1 { playerFrame = 0; }
  }


  playerSrc.X = playerSrc.Width * float32(playerFrame); 
  playerSrc.Y = playerSrc.Height * float32(playerDir); 

  rl.UpdateMusicStream(bgMusic); 
  if isMusicPlaying {
    rl.ResumeMusicStream(bgMusic);
  } else {
    rl.PauseMusicStream(bgMusic);
  }

  camera.Target = rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2)));

  movementDir = 0;
}

func render() {
	rl.BeginDrawing();
	rl.ClearBackground(bgColor);

  rl.BeginMode2D(camera); 

	drawScene();

  rl.EndMode2D();
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

  camera = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/8), float32 (screenHeight/8)), rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))),
0.0, 1.0)
}

func main() {
  defer quit();

	for running {
		input();
		update();
		render();
	}
}

package main

import (
  "log"
  "math"
  "os"
  "strconv"
  "strings"

  rl "github.com/gen2brain/raylib-go/raylib"
)

const (
  screenWidth  = 1000;
  screenHeight = 450;

  movingUp = 0b0001;
  movingDown = 0b0010;
  movingRight = 0b0100;
  movingLeft = 0b1000;
)

var (
  running = true

  bgColor = rl.NewColor(147, 211, 196, 255);

  tex rl.Texture2D; 
  grassSprite  rl.Texture2D;
  hillSprite rl.Texture2D;
  fenceSprite rl.Texture2D;
  houseSprite rl.Texture2D;
  waterSprite rl.Texture2D;
  tilledSprite rl.Texture2D;
  waterprite rl.Texture2D;
  playerSprite rl.Texture2D;

  playerSrc  rl.Rectangle;
  playerDest rl.Rectangle;

  playerSpeed float32 = 1.4;

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

var (
  // the sprite files we wanna pull from 
  srcMap []string; 

  // i.e. the map that we wanna draw onto the screen that has nothing to do with any sprite files 
  // we only draw the sprites we load onto this map
  mapHeight int;
  mapWidth int;
  // i.e. which blocks of the map we do or do not want to draw onto 
  // TODO, might wanna make this a boolean
  tileMap[]int; 

  // where we wanna pull the sprite from (from the file)
  tileSrc rl.Rectangle;
  // where we wanna draw the tile onto the screen
  tileDest rl.Rectangle;
)

func loadMap(filePath string) {
  file, err := os.ReadFile(filePath); 
  if err != nil { 
    log.Fatalf("cannot open map file at %q: error: %q", filePath, err); 
  }

  data := strings.Split(strings.ReplaceAll(string(file), "\n", " "), " ");

  var mapSize int; 
  for i := range data { 
    val64, _ := strconv.ParseInt(data[i], 10, 64); 

    val := int(val64); 


    if i == 0 {
      mapWidth = val; 
    } else if i == 1 {
      mapHeight = val; 
    } else if mapSize = mapHeight * mapWidth; i < mapSize+2 {
      tileMap = append(tileMap, val); 
    } else {
      srcMap = append(srcMap, data[i]);
    }
  }

  if len(tileMap) > mapSize { tileMap = tileMap[:len(tileMap)-1]; }
}

func drawScene() {
  for i := range tileMap {
    if tileMap[i] != 0 {
      switch srcMap[i] {
      case "h": tex = houseSprite; 
      case "l": tex = hillSprite; 
      case "w": tex = waterSprite;
      case "t": tex = tilledSprite;
      case "f": tex = fenceSprite;
      case "g": tex = grassSprite; 
      default: continue; 
    }

      tileDest.X = tileDest.Width * float32(i % mapWidth); 
      tileDest.Y = tileDest.Height * float32(i / mapWidth); 

      tileSrc.X = tileSrc.Width * float32((tileMap[i]-1) % int(tex.Width/int32(tileSrc.Width)));
      tileSrc.Y = tileSrc.Width * float32((tileMap[i]-1) / int(tex.Width/int32(tileSrc.Width)));

      rl.DrawTexturePro(tex, tileSrc, tileDest, rl.NewVector2(tileDest.Width, tileDest.Height), 0, rl.White);
    }
  }


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
    } else if movementDir&movingDown != 0 {
      playerDest.Y += playerSpeed;
    }

    if movementDir&movingRight != 0 { 
      playerDest.X += playerSpeed;
    } else if movementDir&movingLeft != 0 {
      playerDest.X -= playerSpeed;
    }

    if frameCount % 6 == 1 { playerFrame++ };
    if playerFrame > 3 { playerFrame = 0; }
  } else {
    if frameCount % 30 == 1 { playerFrame++; }
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
  camera.Zoom = 1.8;

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

func initializeGame() {
  rl.InitWindow(screenWidth, screenHeight, "Celeste");
  rl.SetExitKey(0);
  rl.SetTargetFPS(60);

  grassSprite = rl.LoadTexture("assets/tiles/Grass.png");
  houseSprite = rl.LoadTexture("assets/tiles/Wooden House.png");
  waterSprite = rl.LoadTexture("assets/tiles/Water.png");
  tilledSprite = rl.LoadTexture("assets/tiles/Tilled Dirt.png");
  fenceSprite = rl.LoadTexture("assets/tiles/Fences.png");
  hillSprite = rl.LoadTexture("assets/tiles/Hills.png");

  tileSrc = rl.NewRectangle(0, 0, 16, 16); 
  tileDest = rl.NewRectangle(0, 0, 16, 16); 

  playerSprite = rl.LoadTexture("assets/characters/Basic Charakter Spritesheet.png");

  playerSrc = rl.NewRectangle(0, 0, 48, 48);
  playerDest = rl.NewRectangle(200, 200, 60, 60);

  rl.InitAudioDevice();
  bgMusic = rl.LoadMusicStream("assets/averys-farm.mp3");
  rl.PlayMusicStream(bgMusic);

  camera = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32 (screenHeight/2)), rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))),
    0.0, 1.0);

  loadMap("./world.map");
}

func main() {
  initializeGame();
  defer quit();

  for running {
    input();
    update();
    render();
  }
}

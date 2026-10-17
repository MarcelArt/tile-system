package engine

import rl "github.com/gen2brain/raylib-go/raylib"

var texture2Ds = make(map[string]rl.Texture2D)

func SetTexture2D(id string, path string) rl.Texture2D {
	texture := rl.LoadTexture(path)
	texture2Ds[id] = texture
	return texture2Ds[id]
}

func GetTexture2D(id string) rl.Texture2D {
	return texture2Ds[id]
}

func UnloadTexture2D(id string) {
	rl.UnloadTexture(texture2Ds[id])
	delete(texture2Ds, id)
}

func UnloadAllTexture2D() {
	for k := range texture2Ds {
		UnloadTexture2D(k)
	}
}

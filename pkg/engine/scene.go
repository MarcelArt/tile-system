package engine

type SceneResult struct {
	NextScene IScene
}

type IScene interface {
	Update() SceneResult
	Draw()
	GetID() string
}

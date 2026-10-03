package no

type Copy struct{}

func (*Copy) Lock()   {}
func (*Copy) Unlock() {}

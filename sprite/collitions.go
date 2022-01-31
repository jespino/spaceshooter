package sprite

type Collition struct {
	Member1 Sprite
	Member2 Sprite
}

func SpritesCollides(s1 Sprite, s2 Sprite) *Collition {
	if s1.IsAlive() && s2.IsAlive() {
		rect1 := s1.Rect()
		rect2 := s2.Rect()
		if rect1.Overlaps(rect2.Rectangle) {
			return &Collition{s1, s2}
		}
	}
	return nil
}

func SpritesGroupsCollides(g1 SpritesGroup, g2 SpritesGroup) []*Collition {
	collitions := []*Collition{}
	for i := 0; i < g1.num; i++ {
		for j := 0; j < g2.num; j++ {
			collition := SpritesCollides(g1.sprites[i], g2.sprites[j])
			if collition != nil {
				collitions = append(collitions, collition)
			}
		}
	}
	return collitions
}

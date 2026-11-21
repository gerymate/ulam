package ulam

type SpiralMove struct {
	// Handle movement in a spiral iterator style.
	x, y, steps, steps_left, dir int
}

func NewSpiralMove(x, y int) *SpiralMove {
	// Create new movement iterator. The origin of the spiral is (x, y).
	return &SpiralMove{x: x, y: y, steps: 1, steps_left: 1, dir: 0}
}

func (s *SpiralMove) NextMove() (int, int) {
	// Make one step in the spiral. First right, then up, then twice left, then twice down, then three times right, etc.
	s.steps_left -= 1
	switch s.dir {
	case 0:
		s.x += 1
		if s.steps_left == 0 {
			s.steps_left = s.steps
			s.dir = 1
		}
	case 1:
		s.y -= 1
		if s.steps_left == 0 {
			s.steps += 1
			s.steps_left = s.steps
			s.dir = 2
		}
	case 2:
		s.x -= 1
		if s.steps_left == 0 {
			s.steps_left = s.steps
			s.dir = 3
		}
	case 3:
		s.y += 1
		if s.steps_left == 0 {
			s.steps += 1
			s.steps_left = s.steps
			s.dir = 0
		}
	}
	return s.x, s.y
}

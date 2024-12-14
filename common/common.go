package common

type Pos struct {
	X int
	Y int
}

func (p Pos) Add(q Pos) Pos {
	return Pos{p.X + q.X, p.Y + q.Y}
}

func (p Pos) Mul(n int) Pos {
	return Pos{p.X * n, p.Y * n}
}
func (p Pos) Eq(q Pos) bool {
	return p.X == q.X && p.Y == q.Y
}

func (p Pos) Dist(q Pos) int {
	res := 0
	if p.X < q.X {
		res += q.X - p.X
	} else {
		res += p.X - q.X
	}
	if p.Y < q.Y {
		res += q.Y - p.Y
	} else {
		res += p.Y - q.Y
	}
	return res
}

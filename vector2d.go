package main

import "math"

type Vector2d struct {
	x float64
	y float64
}

func (v1 Vector2d) Add(v2 Vector2d) Vector2d {
	return Vector2d{v1.x + v2.x, v1.y + v2.y}
}

func (v1 Vector2d) Subtract(v2 Vector2d) Vector2d {
	return Vector2d{v1.x - v2.x, v1.y - v2.y}
}

func (v1 Vector2d) Miltiply(v2 Vector2d) Vector2d {
	return Vector2d{v1.x * v2.x, v1.y * v2.y}
}

func (v1 Vector2d) AddV(d float64) Vector2d {
	return Vector2d{v1.x + d, v1.y + d}
}

func (v1 Vector2d) MultiplyV(d float64) Vector2d {
	return Vector2d{v1.x * d, v1.y * d}
}

func (v1 Vector2d) SubtractV(d float64) Vector2d {
	return Vector2d{v1.x - d, v1.y - d}
}

func (v1 Vector2d) DivisionV(d float64) Vector2d {
	return Vector2d{v1.x / d, v1.y / d}
}

func (v1 Vector2d) limit(lower, upper float64) Vector2d {
	return Vector2d{math.Min(math.Max(v1.x, lower), upper),
		math.Min(math.Max(v1.y, lower), upper)}
}

func (v1 Vector2d) Distance(v2 Vector2d) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}

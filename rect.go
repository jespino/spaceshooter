package main

import "image"

type Rect struct {
	image.Rectangle
}

func NewRect(x0, y0, x1, y1 int) Rect {
	return Rect{
		Rectangle: image.Rect(x0, y0, x1, y1),
	}
}

func (r *Rect) CenterY() int {
	return r.Min.Y + (r.Height() / 2)
}

func (r *Rect) CenterX() int {
	return r.Min.X + (r.Width() / 2)
}

func (r *Rect) Top() int {
	return r.Min.Y
}

func (r *Rect) Bottom() int {
	return r.Max.Y
}

func (r *Rect) Left() int {
	return r.Min.X
}

func (r *Rect) Right() int {
	return r.Max.X
}

func (r *Rect) Width() int {
	return r.Max.X - r.Min.X
}

func (r *Rect) Height() int {
	return r.Max.Y - r.Min.Y
}

func (r *Rect) SetCenterX(x int) {
	diff := x - r.CenterX()
	r.Max.X += diff
	r.Min.X += diff
}

func (r *Rect) SetCenterY(y int) {
	diff := y - r.CenterY()
	r.Max.Y += diff
	r.Min.Y += diff
}

func (r *Rect) SetRight(x int) {
	diff := x - r.Right()
	r.Max.X += diff
	r.Min.X += diff
}

func (r *Rect) SetLeft(x int) {
	diff := x - r.Left()
	r.Max.X += diff
	r.Min.X += diff
}

func (r *Rect) SetTop(y int) {
	diff := y - r.Top()
	r.Max.Y += diff
	r.Min.Y += diff
}

func (r *Rect) SetBottom(y int) {
	diff := y - r.Bottom()
	r.Max.Y += diff
	r.Min.Y += diff
}

func (r *Rect) MoveX(x int) {
	r.SetLeft(r.Left() + x)
}

func (r *Rect) MoveY(y int) {
	r.SetTop(r.Top() + y)
}

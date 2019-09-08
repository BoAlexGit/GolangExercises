// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// geometry.go
// See page 156.

// Package geometry defines simple types for plane geometry.
//!+point
package geometry

import (
	"math"
	"time"
)

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//!-point

//!+path

// A Path is a journey connecting the points with straight lines.
type Path []Point

// Distance returns the distance traveled along the path.
func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

//!-path
*************************************************************************
const day = 24 * time.Hour
fmt.Println(day.SecondsQ) // "86400"
// Мы также определяли собственный метод в разделе 2.5 — метод String типа
// Celsius :
func (с Celsius) String() string { return fmt.Sprintf("%g°C",c) }

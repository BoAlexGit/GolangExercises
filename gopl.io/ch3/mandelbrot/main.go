// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// Mandelbrot.go
// See page 61.
//!+
// Приведенная далее программа использует арифметику complex128 для генера­
//ции множества Мандельброта.
// Mandelbrot создает PNG-изображение фрактала Мандельброта

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}


/***************************************************************
Два вложенных цикла проходят по всем точкам растрового изображения размером
1024x1024 в оттенках серого цвета, представляющего часть комплексной плоскости
от - 2 до +2. Программа проверяет, позволяет ли многократное возведение в квадрат
и добавление числа, представляющего точку, “сбежать” из круга радиусом 2. Если по­
зволяет, то данная точка закрашивается оттенком, соответствующим количеству итера­
ций, потребовавшихся для “побега”. Если не позволяет, данное значение принадлежит
множеству Мандельброта, и точка остается черной. Наконец программа записывает в
стандартный поток вывода изображение в PNG-кодировке, показанное на рис. 3.3.




	var х complex128 = complex(1, 2) // 1+2i
	var у complex128 = complex(3, 4) // 3+4i
	fmt.Println(x*y) // "(5+10i)"
	fmt.Println(real(x*y)) // "5"
	fmt.Println(imag(x*y)) // "10"
Если непосредственно за литералом с плавающей точкой или за десятичным цело­
численным литералом следует i , например 3 .1 4 1 5 9 2 i или 2 i , такой литерал ста­
новится мнимым литералом, обозначающим комплексное число с нулевым действи­
тельным компонентом:
	fmt.Println(1i * li) // "(-1+0i)", i2 = -1
Согласно правилам константной арифметики комплексные константы могут быть
прибавлены к другим константам (целочисленным или с плавающей точкой, действи­
тельным или мнимым), позволяя нам записывать комплексные числа естественным
образом, как, например, 1 + 2 i или, что то же самое, 2 i + 1. Показанные выше объ­
явления х и у могут быть упрощены:
х := 1 + 2i
у := 3 + 4i
Комплексные числа можно проверять на равенство с помощью операторов ==
и !=. Два комплексных числа равны тогда и только тогда, когда равны их действи­
тельные части и их мнимые части.
Пакет math/cmplx предоставляет библиотечные функции для работы с комплекс­
ными числами, такие как комплексный квадратный корень или возведение в степень:
	fmt.Println(cmplx.Sqrt(-1)) // "(0+1i)

*/
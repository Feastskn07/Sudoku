package main

import (
	"fmt"
	"os"
	"strconv"
)

// tahta değişkenini global tanımladık
// bu sayede her fonksiyonda tahtayı ayrı şekilde çağırmak zorunda değiliz
var tahta [][]int

// ana fonksiyon
func main() {
	if sudoku() {
		if cozum(0, 0) {
			ciz(tahta)
		}
	} else {
		fmt.Println("Error")
	}
}

// sudoku tahtasını alan fonksiyon
func sudoku() bool {
	tahta = make([][]int, 9)
	for i, satir := range os.Args[1:] {
		if len(satir) > 9 {
			return false
		}
		tahta[i] = make([]int, 9)
		// satırdaki her elemanı tam sayıya çevirdik
		for j, karakter := range satir {
			if karakter != '.' {
				sayi, err := strconv.Atoi(string(karakter))
				tahta[i][j] = sayi
				if err != nil {
					fmt.Println("Donusum hatasi")
				}
			}
		}
	}
	return kontrolArgs()
}

// arguman kontrolü
func kontrolArgs() bool {
	if len(os.Args) != 10 {
		return false
	}
	for i := 1; i < 10; i++ {
		if len(os.Args[i]) != 9 {
			return false
		}
	}
	for _, satir := range os.Args[1:] {
		for _, karakter := range satir {
			if (karakter < '0' || karakter > '9') && karakter != '.' {
				return false
			}
			for i := 0; i < 9; i++ {
				for j := 1 + i; j < 9; j++ {
					if satir[i] != '.' && (satir[i] == satir[j]) {
						return false
					}
				}
			}
			for i := 1; i < 10; i++ {
				iSatir := os.Args[i]
				for x := 0; x < 9; x++ {
					for j := i + 1; j < 10; j++ {
						jSatir := os.Args[j]
						if iSatir[x] != '.' && (iSatir[x] == jSatir[x]) {
							return false
						}
					}
				}
			}
			for k := 0; k < 9; k += 3 {
				for l := 0; l < 9; l += 3 {
					if !kontrolArgKare(l, k) {
						return false
					}
				}
			}
		}
	}

	return true
}

// argumanlarda 3*3 alanları kontrol edebilmek için destek fonksiyon
func kontrolArgKare(basSatir, basSutun int) bool {
	set := make(map[int]bool)
	for i := basSatir; i < basSatir+3; i++ {
		for j := basSutun; j < basSutun+3; j++ {
			num := tahta[i][j]
			if num != 0 {
				if set[num] {
					return false
				}
				set[num] = true
			}
		}
	}
	return true
}

// çözümü yapan fonksiyon
func cozum(x int, y int) bool {
	if y == 9 {
		return true
	}
	if tahta[y][x] != 0 {
		return cozum(siradakiHucre(x, y))
	} else {
		for i := range [9]int{} {
			v := i + 1
			if kontrol(x, y, v) {
				tahta[y][x] = v
				if cozum(siradakiHucre(x, y)) {
					return true
				}
				tahta[y][x] = 0
			}
		}
		return false
	}
}

// kontrol edilen hücreyi sonrakine geçiren fonksiyon
func siradakiHucre(x int, y int) (int, int) {
	srdX, srdY := (x+1)%9, y
	if srdX == 0 {
		srdY = y + 1
	}
	return srdX, srdY
}

// sırasıyla dikeyde, yatayda ve 3*3 alanda v değeri uygun mu kontrol eden fonksiyonlar
func kontrolDikey(x int, y int, deger int) bool {
	for i := range [9]int{} {
		if tahta[i][x] == deger {
			return true
		}
	}
	return false
}

func kontrolYatay(x int, y int, deger int) bool {
	for i := range [9]int{} {
		if tahta[y][i] == deger {
			return true
		}
	}
	return false
}

func kontrolKare(x int, y int, deger int) bool {
	sx, sy := int(x/3)*3, int(y/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if tahta[sy+dy][sx+dx] == deger {
				return true
			}
		}
	}
	return false
}

// kontrolleri kontrol eden fonksiyon
func kontrol(x int, y int, deger int) bool {
	return !kontrolDikey(x, y, deger) && !kontrolYatay(x, y, deger) && !kontrolKare(x, y, deger)
}

// tahtayı çizen fonksiyon
func ciz(tahta [][]int) {
	for _, satir := range tahta {
		// bu kısmı çıktıda köşeli parantez almamak için kullandık
		for _, eleman := range satir {
			fmt.Print(eleman, " ")
		}
		fmt.Println()
	}
}

// Öğrenci Yurdu Kapasitesi Hesaplama

package main

import "fmt"

// Kat 1 Tanımlama
func firstFloor(single int, totalroom int) int {

	return single * totalroom
}

// Kat 2 Tanımlama
func secondFloor(double int, totalroom int) int {

	return double * totalroom
}

// Kat 1 Kapasitesini Tanımlama
func firstfloorCap() int {
	firstfloorcap := firstFloor(1, 18)
	return firstfloorcap
}

// Kat 2 Kapasitesini Tanımlama
func secondfloorCap() int {
	secondfloorcap := secondFloor(2, 12)
	return secondfloorcap
}

// Toplam Kapasiteyi Tanımlama
func total() int {
	total := firstfloorCap() + secondfloorCap()
	return total
}

// Blok Sayısını Tanımlama
func block() int {

	block := 3
	return block
}

//Hesaplama İşlemleri
func calc() {

	if firstfloorCap() > secondfloorCap() {
		fmt.Println("Tek Kişilik Odalarda Kalan Öğrenci Sayısı Çift Kişilik Odalarda Kalan Öğrenci Sayısından", firstfloorCap()-secondfloorCap(), "Fazladır")

	} else {
		fmt.Println("Çift Kişilik Odalarda Kalan Öğrenci Sayısı Tek Kişilik Odalarda Kalan Öğrenci Sayısından", secondfloorCap()-firstfloorCap(), "Fazladır")

	}
	fmt.Println("Bir Blokta Bulunan Toplam Öğrenci Sayısı", total())
	if block() == 1 {
		fmt.Println("Yurtta Bulunan Toplam Öğrenci Sayısı", total())
	} else {
		fmt.Println("Yurtta Bulunan Toplam Öğrenci Sayısı", block()*total())
	}
}

//Hesaplama İşlemlerini Gerçekleştirilen Fonksiyonu Çağırma
func main() {
	calc()
}

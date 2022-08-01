// Öğrenci Yurdu Kapasitesi Hesaplama
package main

import "fmt"

// Kat 1 Tanımlama
func firstFloor(single int, totalRoom int) int {

	return single * totalRoom
}

// Kat 2 Tanımlama
func secondFloor(double int, totalRoom int) int {

	return double * totalRoom
}

// Kat 1 Kapasitesi Tanımlama
func firstFloorCap() int {
	firstFloorcap := firstFloor(1, 18)
	return firstFloorcap
}

// Kat 2 Kapasitesi Tanımlama
func secondFloorCap() int {
	secondFloorcap := secondFloor(2, 9)
	return secondFloorcap
}

// Toplam Kapasiteyi Tanımlama
func total() int {
	total := firstFloorCap() + secondFloorCap()
	return total
}

//Yurt Blok Sayısını Tanımlama
func block() int {

	block := 3
	return block
}

//Hesaplama İşlemleri
func calc() {
	diffCap := firstFloorCap() > secondFloorCap()
	diffVar := firstFloorCap() - secondFloorCap()
	diffNull := firstFloorCap() == secondFloorCap()
	if diffCap {
		fmt.Println("Tek Kişilik Odalarda Kalan Öğrenci Sayısı Çift Kişilik Odalarda Kalan Öğrenci Sayısından", diffVar, "Fazladır.")

	} else {
		if diffNull {			 // Öğrenci Sayılarının Eşit Olma Koşulu Eklenmiştir.

			fmt.Println("Tek Kişilik Odalarda Kalan Öğrenci Sayısı Çift Kişilik Odalarda Kalan Öğrenci Sayısına Eşittir.")
		} else {
			fmt.Println("Çift Kişilik Odalarda Kalan Öğrenci Sayısı Tek Kişilik Odalarda Kalan Öğrenci Sayısından", -1*diffVar, "Fazladır.")
		}
	}

	fmt.Println("Bir Blokta Bulunan Toplam Öğrenci Sayısı", total())

	//Fonksiyon dışarıdan değer almadığından blok sayısı elle girilip kaç blok var ise tüm yurt kapasitesi hesaplanmaktadır.
	
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

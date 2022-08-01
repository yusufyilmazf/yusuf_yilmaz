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

// Kat 1 Kapasitesini Tanımlama
func firstFloorCap() int {
	firstFloorcap := firstFloor(1, 18)
	return firstFloorcap
}

// Kat 2 Kapasitesini Tanımlama
func secondFloorCap() int {
	secondFloorCap := secondFloor(2, 12)
	return secondFloorCap
}

// Toplam Kapasiteyi Tanımlama
func total() int {
	total := firstFloorCap() + secondFloorCap()
	return total
}

// Blok Sayısını Tanımlama
func block() int {

	block := 3
	return block
}

//Hesaplama İşlemleri
func calc() {
	
	diffCap := firstFloorCap() > secondFloorCap()

	if diffCap > 0 {
		fmt.Println("Tek Kişilik Odalarda Kalan Öğrenci Sayısı Çift Kişilik Odalarda Kalan Öğrenci Sayısından", diffCap, "Fazladır")

	} else {
		fmt.Println("Çift Kişilik Odalarda Kalan Öğrenci Sayısı Tek Kişilik Odalarda Kalan Öğrenci Sayısından", -1 * diffCap, "Fazladır")

	}
	//TODO 
	// Öğrenci sayılar eşit ise sonuç ne olacaktır ?
	
	
	fmt.Println("Bir Blokta Bulunan Toplam Öğrenci Sayısı", total())
	
	//TODO
	// Aşağıda tanımlanan if Bloğunun her hangi bir geçerliliği var mı ? 
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

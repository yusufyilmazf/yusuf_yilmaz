// Öğrenci Yurdu Kapasitesi Hesaplama

package main

import "fmt"

func firstFloor(single int, totalroom int) int {

	return single * totalroom
}

func secondFloor(double int, totalroom int) int {

	return double * totalroom
}

func firstfloorCap() int {
	firstfloorcap := firstFloor(1, 18)
	return firstfloorcap
}

func seconfloorCap() int {
	seconfloorcap := secondFloor(2, 12)
	return seconfloorcap
}

func total() int {
	total := firstfloorCap() + seconfloorCap()
	return total
}

func block() int {

	block := 3
	return block
}

func main() {

	if firstfloorCap() > seconfloorCap() {
		fmt.Println("Tek Kişilik Odalarda Kalan Öğrenci Sayısı Çift Kişilik Odalarda Kalan Öğrenci Sayısından", firstfloorCap()-seconfloorCap(), "Fazladır")

	} else {
		fmt.Println("Çift Kişilik Odalarda Kalan Öğrenci Sayısı Tek Kişilik Odalarda Kalan Öğrenci Sayısından", seconfloorCap()-firstfloorCap(), "Fazladır")

	}
	fmt.Println("Bir Blokta Bulunan Toplam Öğrenci Sayısı", total())
	if block() == 1 {

	} else {
		fmt.Println("Yurtta Bulunan Toplam Öğrenci Sayısı", block()*total())
	}

}

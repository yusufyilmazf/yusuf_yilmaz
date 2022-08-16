//AVM İçerisinde Öğrenci Sinema Bileti Alana Rastgele Yemek Kuponu Veren Kampanya

package main

import (
	"fmt"
	"math/rand"
)

// İndirim Kuponlarının Sayılarını Tutan Fonksiyon
func main() {

	var couponCapacity int = 25
	numberOfStudents := studentDiscount(40, 35)
	var numberOfCouponsRemaining int = couponCapacity - numberOfStudents
	for i := couponCapacity; i > 0; i-- {
		fmt.Println(i, ". Yemek Kuponunu Kullandınız")
		if i == 0 {
			fmt.Println("Yemek Kuponumuz Kalmamıştır.")
		} else if i == numberOfCouponsRemaining+1 {
			i--
			fmt.Println("Kalan Yemek Kuponu :", i)
			break
		}

	}

}

//Rastgele Yemek Kuponu Veren Fonksiyon
func randomFoodCoupon() ([6]string, [7]string) {
	var incomingCustomer [6]string = [6]string{"Yusuf", "Mücahit", "Abdülkadir", "Furkan", "Babür", "Talha"}
	var restaurants [7]string = [7]string{"Mc Donalds", "Pidem", "Burger King", "Hd İskender", "Sayrem", "KFC", "Bay İskender"}

	for i := 0; i < len(incomingCustomer); i++ {

		randomFoodCoupon := restaurants[rand.Intn(len(incomingCustomer))]

		fmt.Println("Kuponu Kullanabileceğiniz Restaurantlar :", randomFoodCoupon)

	}
	return incomingCustomer, restaurants
}

// Müşterilerin Yaşlarına Göre Öğrenci Koşulunu Sağlayıp Sağlamadığını Kontrol Eden Fonksiyon
func studentDiscount(movieTicketPrice int, discountedMovieTicketPrice int) int {

	var customerAges [6]int = [6]int{21, 24, 16, 29, 31, 46}
	var customerQueues [6]int = [6]int{1, 2, 3, 4, 5, 6}
	var studentDiscount int
	for i := 0; i < len(customerAges); i++ {

		if customerAges[i] > 30 && customerAges[i] < 18 {
			fmt.Println(customerQueues[i], ". Müşteri İçin Ödenecek Olan Tam Bilet Tutarı :", movieTicketPrice)

		} else {
			fmt.Println(customerQueues[i], ". Müşteri İçin Ödenecek Olan Öğrenci Bileti Tutarı :", discountedMovieTicketPrice)
			fmt.Println("Öğrenci İndirimimizden Yararlandığınızdan Dolayı Yemek Kuponu Kazandınız!")
			randomFoodCoupon()
			studentDiscount++
		}

	}
	return studentDiscount
}

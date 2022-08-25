//HASTANE RANDEVU SİSTEMİ

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Öğle Arası Saati ve Öğle Arası Bitimine Kalan Zamanın Tanımlandığı Fonksiyon
func main() {

	var lunchBreak int = 12
	date := time.Now().Hour()
	dateMinute := 60 + lunchBreak - time.Now().Minute()

	checkLunchBreakTime(&lunchBreak, &date, &dateMinute)

}

//Anlık Zamanın Öğle Arasına Denk Gelip Gelmediğini Kontrol Eden Fonksiyon
func checkLunchBreakTime(lunchBreak *int, date *int, dateMinute *int) {

	if *date >= *lunchBreak && *date < *lunchBreak+1 {
		fmt.Println("Öğle Arası", *dateMinute-*lunchBreak, "Dakika Sonra Tekrar Geliniz")
	} else {
		checkAppointment()
	}
}

//Muayene Olunmak İstenen Bölümün Seçildiği ve Seçtiğimiz Bölüme Göre Rastgele Doktor Ataması Yapan Fonksiyon
func hospitalAppointmentSystem() {

	var pickHospitalBranch string
	var hospitalBranch [5]string = [5]string{"Dahiliye", "Kardiyoloji", "Radyoloji", "Cildiye", "Nöroloji"}
	fmt.Printf("Muayene Olmak İstediğiniz Hastane Bölümünü Seçiniz : %s\n", hospitalBranch)
	fmt.Scan(&pickHospitalBranch)

	var internalMedicineDoctors [5]string = [5]string{"Dr. Tuğçe", "Dr. Kazım", "Dr. Hasan", "Dr. Babür", "Dr. Selami"}
	var cardiologyDoctors [5]string = [5]string{"Dr. Ahmet", "Dr. Mehmet", "Dr. Canan", "Dr. Cemre", "Dr. Kemal"}
	var radiologyDoctors [5]string = [5]string{"Dr. Sude", "Dr. Furkan", "Dr. Talha", "Dr. Seyit", "Dr. Mahmut"}
	var dermatologyDoctors [5]string = [5]string{"Dr. Ali", "Dr. Fatih", "Dr. Şevval", "Dr. Hüseyin", "Dr. Ecrin"}
	var neurologyDoctors [5]string = [5]string{"Dr. Furkan", "Dr. Abdülkadir", "Dr. Eda", "Dr. Sultan", "Dr. Ömer"}

	randomDoctorAppointment := internalMedicineDoctors[rand.Intn(len(internalMedicineDoctors))]
	randomDoctorAppointment1 := cardiologyDoctors[rand.Intn(len(cardiologyDoctors))]
	randomDoctorAppointment2 := radiologyDoctors[rand.Intn(len(radiologyDoctors))]
	randomDoctorAppointment3 := dermatologyDoctors[rand.Intn(len(dermatologyDoctors))]
	randomDoctorAppointment4 := neurologyDoctors[rand.Intn(len(neurologyDoctors))]

	switch pickHospitalBranch {
	case hospitalBranch[0]:
		fmt.Println("Dahiliye Bölümü İçin Muayene Olacağınız Doktor :", randomDoctorAppointment)
	case hospitalBranch[1]:
		fmt.Println("Kardiyoloji Bölümü İçin Muayene Olacağınız Doktor :", randomDoctorAppointment1)
	case hospitalBranch[2]:
		fmt.Println("Radyoloji Bölümü İçin Muayene Olacağınız Doktor :", randomDoctorAppointment2)
	case hospitalBranch[3]:
		fmt.Println("Cildiye Bölümü İçin Muayene Olacağınız Doktor :", randomDoctorAppointment3)
	case hospitalBranch[4]:
		fmt.Println("Nöroloji Bölümü İçin Muayene Olacağınız Doktor :", randomDoctorAppointment4)
	default:
		fmt.Println("Hatalı Bölüm Adı Girdiniz. Lütfen Bölüm Adını Doğru Girdiğinizden Emin Olun")
	}

}

//Randevumuzun Olup Olmadığını Soran Fonksiyon
func checkAppointment() {

	var haveAppointment string
	fmt.Println("Randevunuz Var mı?")
	fmt.Scan(&haveAppointment)
	switch haveAppointment {
	case "Evet":
		fmt.Println("Randevu Saatinizde Doktorunuzun Yanına Gidebilirsiniz.")
	case "Hayır":
		hospitalAppointmentSystem()
	}
}

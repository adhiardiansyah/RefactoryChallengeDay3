package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	fibonacci()
	fmt.Println("\n--------------------------------")
	palindrome()
	fmt.Println("\n--------------------------------")
	capitalize()
	fmt.Println("--------------------------------")
	fmt.Println(randomString(10))
	fmt.Println("\n--------------------------------")
	var f = do
	f()
	f = func() { fmt.Println("Berhenti lakukan!") }
	f()
	fmt.Println("\n--------------------------------")
}

func fibonacci() {
	var bilangan, sesudah int
	var sebelum, saatini int = 0, 1
	fmt.Println("Fibonacci")
	fmt.Print("Masukkan jumlah bilangan: ")
	fmt.Scanln(&bilangan)
	if bilangan < 2 {
		fmt.Print("\nJumlah minimal bilangan pada Fibonacci adalah dua")
	} else {
		fmt.Print("\nFibonacci:\n")
		for i := 0; i < bilangan; i++ {
			fmt.Print(sebelum, " ")
			sesudah = sebelum + saatini
			sebelum = saatini
			saatini = sesudah
		}
	}
	fmt.Println()
}

func palindrome() {
	var angka, sisa, temp int
	var reverse int = 0

	fmt.Println()
	fmt.Println("Palindrom")
	fmt.Print("Masukkan beberapa angka positif : ")
	fmt.Scanln(&angka)

	temp = angka

	for {
		sisa = angka % 10
		reverse = reverse*10 + sisa
		angka /= 10

		if angka == 0 {
			break // Break digunakan untuk keluar dari loop
		}
	}

	if temp == reverse {
		fmt.Print(temp, " adalah Palindrom")
	} else {
		fmt.Print(temp, " bukan Palindrom")
	}
	fmt.Println()
}

func capitalize() {
	fmt.Println()
	fmt.Println("Kalimat Kapital")
	fmt.Print("Masukkan kalimat : ")
	inputReader := bufio.NewReader(os.Stdin)
	kalimat, _ := inputReader.ReadString('\n')

	kalimatKapital := strings.ToUpper(kalimat)

	fmt.Println("Sebelum :", kalimat)
	fmt.Println("Sesudah :", kalimatKapital)
}

func randomString(n int) string {
	fmt.Println()
	fmt.Println("Kalimat random sebanyak", n, ":")
	var kata = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = kata[rand.Intn(len(kata))]
	}
	return string(s)
}

func do() {
	fmt.Println()
	fmt.Println("Function Expression")
	fmt.Println()
	fmt.Println("Lakukan!")
}

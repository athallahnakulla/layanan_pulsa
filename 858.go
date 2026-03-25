package main

import (
	"database/sql"
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var step = 0

func connectDB() {
	db, _ = sql.Open("sqlite3", "./pulsa.db")
}

func getPulsa() float64 {
	var saldo float64
	db.QueryRow("SELECT saldo FROM users WHERE id = 1").Scan(&saldo)
	return saldo
}

func getKuota() (float64, string) {
	var kuota float64
	var masa string
	db.QueryRow("SELECT kuota, masa_aktif FROM users WHERE id = 1").Scan(&kuota, &masa)
	return kuota, masa
}

func beliKuota(harga float64, paket float64, tambahHari int) string {
	saldo := getPulsa()

	if saldo < harga {
		return "Saldo tidak cukup\n\n0. Kembali"
	}

	db.Exec("UPDATE users SET saldo = saldo - ? WHERE id = 1", harga)
	db.Exec("UPDATE users SET kuota = kuota + ? WHERE id = 1", paket)

	_, masaLama := getKuota()
	var waktu time.Time

	if masaLama == "" {
		waktu = time.Now()
	} else {
		waktu, _ = time.Parse("2006-01-02", masaLama)
	}

	waktuBaru := waktu.AddDate(0, 0, tambahHari)
	masaBaru := waktuBaru.Format("2006-01-02")

	db.Exec("UPDATE users SET masa_aktif = ? WHERE id = 1", masaBaru)

	return fmt.Sprintf(
		"Pembelian berhasil\nPulsa: Rp %.0f\nKuota: +%.0f GB\nMasa Aktif: %s\n\n0. Kembali",
		saldo-harga, paket, masaBaru,
	)
}

func main() {
	connectDB()

	myApp := app.New()
	w := myApp.NewWindow("USSD *858#")

	label := widget.NewLabel("Ketik *858# untuk mulai")
	input := widget.NewEntry()

	button := widget.NewButton("Kirim", func() {
		text := input.Text

		if text == "0" {
			label.SetText("Menu:\n1. Cek Pulsa\n2. Cek Kuota\n3. Beli Kuota\n\n0. Menu Utama\n9. Keluar")
			step = 1
			input.SetText("")
			return
		}

		if text == "9" && step == 1 {
			label.SetText("Terima kasih telah menggunakan layanan *858#")
			step = 0
			input.SetText("")
			return
		}

		if step == 0 {
			if text == "*858#" {
				label.SetText("Menu:\n1. Cek Pulsa\n2. Cek Kuota\n3. Beli Kuota\n\n0. Menu Utama\n9. Keluar")
				step = 1
			} else {
				label.SetText("Kode USSD salah!\nKetik *858#")
			}

		} else if step == 1 {

			if text == "1" {
				saldo := getPulsa()
				label.SetText(fmt.Sprintf("Pulsa: Rp %.0f\n\n0. Kembali", saldo))
				step = 3

			} else if text == "2" {
				kuota, masa := getKuota()
				label.SetText(fmt.Sprintf("Kuota: %.1f GB\nMasa Aktif: %s\n\n0. Kembali", kuota, masa))
				step = 4

			} else if text == "3" {
				label.SetText("Pilih paket:\n1. 1GB (10k / 7 hari)\n2. 2GB (15k / 15 hari)\n\n0. Kembali")
				step = 2

			} else {
				label.SetText("Pilihan salah\n\n0. Kembali\n9. Keluar")
			}

		} else if step == 2 {
			if text == "1" {
				label.SetText(beliKuota(10000, 1, 7))
				step = 5

			} else if text == "2" {
				label.SetText(beliKuota(15000, 2, 15))
				step = 5

			} else {
				label.SetText("Pilihan salah\n\n0. Kembali")
			}

		} else if step == 3 || step == 4 || step == 5 {
			label.SetText("Tekan:\n0. Kembali")
		}

		input.SetText("")
	})

	w.SetContent(container.NewVBox(label, input, button))
	w.ShowAndRun()
}

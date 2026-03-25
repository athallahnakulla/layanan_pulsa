# 📱 USSD Pulsa App (*858#)

Simulasi sistem USSD seperti pada handphone (contoh: *805#) yang dibuat menggunakan **Golang**, **Fyne (GUI)**, dan **SQLite**.

Aplikasi ini meniru cara kerja layanan operator seluler untuk:
- Cek pulsa
- Cek kuota internet
- Pembelian paket data
- Isi ulang pulsa (top-up)
- Manajemen masa aktif



## 🚀 Fitur Utama

- 📞 Akses menggunakan kode USSD `*858#`
- 💰 Cek saldo pulsa
- 📶 Cek sisa kuota & masa aktif
- 🛒 Pembelian paket kuota (otomatis menambah masa aktif)
- 🔄 Top-up pulsa
- 🚪 Exit system (seperti USSD asli)
- 🧭 Navigasi berbasis menu (state machine)



## 🛠️ Tech Stack

- **Golang** → Logic & backend
- **Fyne** → GUI Desktop App
- **SQLite** → Database lokal



## 🧱 Struktur Database

Tabel: `users`

| Field        | Tipe   | Keterangan            |
|-------------|--------|----------------------|
| id          | int    | Primary key          |
| saldo       | float  | Saldo pulsa          |
| kuota       | float  | Kuota internet (GB)  |
| masa_aktif  | text   | Masa aktif (YYYY-MM-DD) |



## ▶️ Cara Menjalankan

1. Clone repository:
git clone https://github.com/athallahnakulla/layanan_pulsa.git

2. Masuk ke folder:
cd layanan_pulsa

3. Install dependency:
go mod tidy

4. Jalankan aplikasi:
go run main.go

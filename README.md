Berikut adalah contoh **README** untuk proyek di GitHub yang memuat penjelasan tentang tujuan proyek, bagaimana cara menjalankan kode, dan lainnya.

---

# **Optimasi Penjadwalan Shift Pegawai dengan Pemrograman Linear Integer**

Proyek ini bertujuan untuk mengembangkan model optimasi penjadwalan shift pegawai di minimarket menggunakan pendekatan **Pemrograman Linear Integer (ILP)**. Model ini mengalokasikan pegawai ke shift (Pagi, Siang, Malam) dengan **minimisasi biaya total**.

## **Deskripsi Proyek**

Penjadwalan shift pegawai adalah tugas yang penting dalam operasi sehari-hari minimarket dan bisnis lain dengan shift kerja. Dengan menggunakan **model optimasi**, proyek ini bertujuan untuk meminimalkan total biaya yang dikeluarkan untuk membayar upah pegawai, sambil memastikan bahwa kebutuhan jumlah pegawai per shift (demand) terpenuhi.

### **Tujuan Proyek:**

- Menyusun jadwal kerja yang memenuhi kebutuhan pegawai untuk setiap shift.
- Meminimalkan total biaya yang dikeluarkan untuk membayar upah pegawai.
- Menerapkan **Pemrograman Linear Integer (ILP)** menggunakan **Go** dan **GLPK** untuk menyelesaikan masalah ini.

## **Fitur:**

- Model optimasi untuk penjadwalan shift pegawai.
- **GLPK** solver untuk menyelesaikan masalah optimasi.
- Aplikasi menggunakan **Go (Golang)** dengan **GLPK binding** untuk melakukan komputasi.
- Output berupa **jadwal pegawai** per shift yang optimal, dengan **biaya total** yang terminimalisasi.

## **Persyaratan:**

- **Go** (Golang) versi 1.18 atau lebih tinggi.
- **GLPK** library (GNU Linear Programming Kit) yang sudah terpasang di sistem Anda.

## **Instalasi dan Pengaturan**

### 1. **Clone Repository**

```bash
git clone https://github.com/mhmdiamd/go-model-employee-shifting.git
cd go-model-employee-shifting
```

### 2. **Instalasi Go dan GLPK**

#### Menginstal **Go**:

Ikuti instruksi penginstalan Go di [Go Lang Official](https://golang.org/dl/).

#### Menginstal **GLPK** di macOS (untuk Linux atau Windows, ikuti instruksi masing-masing):

```bash
brew install glpk
```

#### Menginstal GLPK binding untuk Go (go-glpk):

```bash
go get github.com/lukpank/go-glpk
```

### 3. **Menjalankan Program**

Setelah semua persyaratan diinstal, jalankan program dengan perintah berikut:

```bash
go run main.go
```

Ini akan menjalankan model optimasi penjadwalan dan menghasilkan output berupa penugasan pegawai ke shift yang optimal, beserta total biaya yang dihitung.

## **Struktur Direktori**

```
shift-scheduling-optimization/
│
├── main.go              # File utama untuk menjalankan optimasi
├── model/
│   └── scenario.go      # Model dan solver penjadwalan shift
├── README.md            # Dokumen ini
├── go.mod               # Modul dependensi Go
└── go.sum               # Hash dependensi Go
```

## **Contoh Output**

Output dari program ini akan menampilkan penugasan pegawai ke shift serta biaya total:

```plaintext
MIP Assignments:
  G → Siang
  I → Siang
  A → Siang
  K → Malam
  C → Pagi
  D → Pagi
  F → Malam
  H → Siang
  J → Malam
  B → Siang

Jumlah pegawai per shift:
- Pagi: 3 orang
- Siang: 5 orang
- Malam: 3 orang

Total Biaya: Rp 556000
```

## **Kontribusi**

Jika Anda ingin berkontribusi pada proyek ini, silakan fork repositori ini, buat cabang baru, dan ajukan pull request dengan perubahan atau fitur yang diinginkan.

### **Langkah-langkah kontribusi:**

1. Fork repositori ini
2. Clone repositori ke lokal Anda
3. Buat cabang baru untuk fitur yang akan ditambahkan
4. Lakukan perubahan dan commit
5. Push perubahan ke repositori forked Anda
6. Ajukan pull request ke repositori utama

## **Lisensi**

Proyek ini dilisensikan di bawah **MIT License** – lihat [LICENSE.md](LICENSE.md) untuk detail lebih lanjut.

---

Dengan **README** ini, pengguna dapat dengan mudah memahami tujuan proyek, cara instalasi, dan cara menjalankan aplikasi, serta mendapatkan contoh output dan instruksi untuk berkontribusi. Apakah ada hal yang ingin ditambahkan atau diperbarui?

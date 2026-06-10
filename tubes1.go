package main
import "fmt"

const NMAX int = 100

type CatatanMood struct {
	id        int
	skorEmosi int
	deskripsi string
	tanggal   string
	bulan     int
	mingguKe  int
	tahun     int
}

type Tugas struct {
	id        int
	namaTugas string
	durasi    int // dalam menit
	selesai   bool
	prioritas int
}

type arrMood [NMAX]CatatanMood
type arrTugas [NMAX]Tugas

func inputString() string {
	var result string
	fmt.Scanln(&result)
	return result
}

func inputInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func hitungMingguKe(tanggal string) (int, int, int) {
	var d, m, y int = 0, 0, 0
	var i int = 0
	var bagian int = 0
	var num int = 0

	for i < len(tanggal) {
		if tanggal[i] >= '0' && tanggal[i] <= '9' {
			num = num*10 + int(tanggal[i]-'0')
		} else if tanggal[i] == '/' {
			if bagian == 0 {
				d = num
			} else if bagian == 1 {
				m = num
			}
			num = 0
			bagian = bagian + 1
		}
		i = i + 1
	}
	y = num

	if d < 1 {
		d = 1
	}
	if m < 1 {
		m = 1
	}
	if d <= 7 {
		return 1, m, y
	} else if d <= 14 {
		return 2, m, y
	} else if d <= 21 {
		return 3, m, y
	}
	return 4, m, y
}

func skorEmosiString(skor int) string {
	if skor == 1 {
		return "1 (Terrible)"
	} else if skor == 2 {
		return "2 (Bad)"
	} else if skor == 3 {
		return "3 (Fine)"
	} else if skor == 4 {
		return "4 (Good)"
	} else if skor == 5 {
		return "5 (Awesome!)"
	} else {
		return " "
	}
}

func tambahMood(T *arrMood, n *int, idCounter *int) {
	if *n >= NMAX {
		fmt.Println("Data penuh, tidak dapat menambah catatan mood.")
		return
	}
	*idCounter = *idCounter + 1
	var skor int
	var desc string
	var tanggal string
	fmt.Println("Masukkan skor emosi (1-5):")
	fmt.Println("1. Terrible")
	fmt.Println("2. Bad")
	fmt.Println("3. Fine")
	fmt.Println("4. Good")
	fmt.Println("5. Awesome!")
	fmt.Print("Pilih skor: ")
	skor = inputInt()
	for skor < 1 || skor > 5 {
		fmt.Println("Skor tidak valid! Harap masukkan angka 1-5.")
		fmt.Print("Pilih skor: ")
		skor = inputInt()
	}
	fmt.Print("Masukkan deskripsi perasaan: ")
	desc = inputString()
	fmt.Print("Masukkan tanggal (DD/MM/YYYY): ")
	tanggal = inputString()

	T[*n].id = *idCounter
	T[*n].skorEmosi = skor
	T[*n].deskripsi = desc
	T[*n].tanggal = tanggal
	T[*n].mingguKe, T[*n].bulan, T[*n].tahun = hitungMingguKe(tanggal)
	*n = *n + 1
	fmt.Println("Catatan mood berhasil ditambahkan.")
}

func cariMoodByID(T arrMood, n int, id int) int {
	var found int = -1
	var i int = 0
	for i < n && found == -1 {
		if T[i].id == id {
			found = i
		}
		i = i + 1
	}
	return found
}

func ubahMood(T *arrMood, n int) {
	var id int
	fmt.Print("Masukkan ID catatan mood yang ingin diubah: ")
	id = inputInt()
	var idx int = cariMoodByID(*T, n, id)
	if idx == -1 {
		fmt.Println("Data dengan ID tersebut tidak ditemukan.")
		return
	}
	var skor int
	var desc string
	var tanggal string
	fmt.Println("Masukkan skor emosi baru (1-5):")
	fmt.Println("1. Terrible")
	fmt.Println("2. Bad")
	fmt.Println("3. Fine")
	fmt.Println("4. Good")
	fmt.Println("5. Awesome")
	fmt.Print("Pilih skor: ")
	skor = inputInt()
	for skor < 1 || skor > 5 {
		fmt.Println("Skor tidak valid! Harap masukkan angka 1-5.")
		fmt.Print("Pilih skor: ")
		skor = inputInt()
	}
	fmt.Print("Masukkan deskripsi perasaan baru: ")
	desc = inputString()
	fmt.Print("Masukkan tanggal baru (DD/MM/YYYY): ")
	tanggal = inputString()
	T[idx].skorEmosi = skor
	T[idx].deskripsi = desc
	T[idx].tanggal = tanggal
	T[idx].mingguKe, T[idx].bulan, T[idx].tahun = hitungMingguKe(tanggal)
	fmt.Println("Catatan mood berhasil diubah.")
}

func hapusMood(T *arrMood, n *int) {
	var id int
	fmt.Print("Masukkan ID catatan mood yang ingin dihapus: ")
	id = inputInt()
	var idx int = cariMoodByID(*T, *n, id)
	if idx == -1 {
		fmt.Println("Data dengan ID tersebut tidak ditemukan.")
		return
	}
	var i int = idx
	for i < *n-1 {
		T[i] = T[i+1]
		i = i + 1
	}
	*n = *n - 1
	fmt.Println("Catatan mood berhasil dihapus.")
}

func tampilMood(T arrMood, n int) {
	if n == 0 {
		fmt.Println("Belum ada catatan mood.")
		return
	}
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Printf("%-5s %-15s %-20s %-15s %-10s\n", "ID", "Skor Emosi", "Deskripsi", "Tanggal", "Minggu Ke")
	fmt.Println("--------------------------------------------------------------------------")
	var i int = 0
	for i < n {
		fmt.Printf("%-5d %-15s %-20s %-15s %-10d\n", T[i].id, skorEmosiString(T[i].skorEmosi), T[i].deskripsi, T[i].tanggal, T[i].mingguKe)
		i = i + 1
	}
	fmt.Println("--------------------------------------------------------------------------")
}

func tambahTugas(T *arrTugas, n *int, idCounter *int) {
	if *n >= NMAX {
		fmt.Println("Data penuh, tidak dapat menambah tugas.")
		return
	}
	*idCounter = *idCounter + 1
	var nama string
	var durasi int
	var prioritas int
	fmt.Print("Masukkan nama tugas: ")
	nama = inputString()
	fmt.Print("Masukkan durasi pengerjaan (menit): ")
	durasi = inputInt()
	fmt.Print("Masukkan prioritas (1=Tinggi, 2=Sedang, 3=Rendah): ")
	prioritas = inputInt()

	T[*n].id = *idCounter
	T[*n].namaTugas = nama
	T[*n].durasi = durasi
	T[*n].prioritas = prioritas
	T[*n].selesai = false
	*n = *n + 1
	fmt.Println("Tugas berhasil ditambahkan.")
}

func cariTugasByID(T arrTugas, n int, id int) int {
	var found int = -1
	var i int = 0
	for i < n && found == -1 {
		if T[i].id == id {
			found = i
		}
		i = i + 1
	}
	return found
}

func ubahTugas(T *arrTugas, n int) {
	var id int
	fmt.Print("Masukkan ID tugas yang ingin diubah: ")
	id = inputInt()
	var idx int = cariTugasByID(*T, n, id)
	if idx == -1 {
		fmt.Println("Data dengan ID tersebut tidak ditemukan.")
		return
	}
	var nama string
	var durasi int
	var prioritas int
	fmt.Print("Masukkan nama tugas baru: ")
	nama = inputString()
	fmt.Print("Masukkan durasi pengerjaan baru (menit): ")
	durasi = inputInt()
	fmt.Print("Masukkan prioritas baru (1=Tinggi, 2=Sedang, 3=Rendah): ")
	prioritas = inputInt()
	T[idx].namaTugas = nama
	T[idx].durasi = durasi
	T[idx].prioritas = prioritas
	fmt.Println("Tugas berhasil diubah.")
}

func hapusTugas(T *arrTugas, n *int) {
	var id int
	fmt.Print("Masukkan ID tugas yang ingin dihapus: ")
	id = inputInt()
	var idx int = cariTugasByID(*T, *n, id)
	if idx == -1 {
		fmt.Println("Data dengan ID tersebut tidak ditemukan.")
		return
	}
	var i int = idx
	for i < *n-1 {
		T[i] = T[i+1]
		i = i + 1
	}
	*n = *n - 1
	fmt.Println("Tugas berhasil dihapus.")
}

func tandaiSelesai(T *arrTugas, n int) {
	var id int
	fmt.Print("Masukkan ID tugas yang ingin ditandai selesai: ")
	id = inputInt()
	var idx int = cariTugasByID(*T, n, id)
	if idx == -1 {
		fmt.Println("Data dengan ID tersebut tidak ditemukan.")
		return
	}
	T[idx].selesai = true
	fmt.Println("Tugas berhasil ditandai selesai.")
}

func tampilTugas(T arrTugas, n int) {
	if n == 0 {
		fmt.Println("Belum ada daftar tugas.")
		return
	}
	fmt.Println("------------------------------------------------------------------")
	fmt.Printf("%-5s %-25s %-15s %-10s %-10s\n", "ID", "Nama Tugas", "Durasi (menit)", "Prioritas", "Status")
	fmt.Println("------------------------------------------------------------------")
	var i int = 0
	for i < n {
		var labelPrioritas string
		if T[i].prioritas == 1 {
			labelPrioritas = "Tinggi"7
		} else if T[i].prioritas == 2 {
			labelPrioritas = "Sedang"
		} else {
			labelPrioritas = "Rendah"
		}
		var labelSelesai string
		if T[i].selesai {
			labelSelesai = "Selesai"
		} else {
			labelSelesai = "Belum"
		}
		fmt.Printf("%-5d %-25s %-15d %-10s %-10s\n", T[i].id, T[i].namaTugas, T[i].durasi, labelPrioritas, labelSelesai)
		i = i + 1
	}
	fmt.Println("------------------------------------------------------------------")
}

func seqSearchMood(T arrMood, n int, katakunci string) int {
	var found int = -1
	var i int = 0
	for i < n && found == -1 {
		if T[i].tanggal == katakunci {
			found = i
		}
		i = i + 1
	}
	return found
}

func seqSearchTugas(T arrTugas, n int, katakunci string) int {
	var found int = -1
	var i int = 0
	for i < n && found == -1 {
		if T[i].namaTugas == katakunci {
			found = i
		}
		i = i + 1
	}
	return found
}

func binarySearchMood(T arrMood, n int, skor int) int {
	var found int = -1
	var kr int = 0
	var kn int = n - 1
	var med int
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if skor < T[med].skorEmosi {
			kn = med - 1
		} else if skor > T[med].skorEmosi {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}

func binarySearchTugas(T arrTugas, n int, durasi int) int {
	var found int = -1
	var kr int = 0
	var kn int = n - 1
	var med int
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if durasi < T[med].durasi {
			kn = med - 1
		} else if durasi > T[med].durasi {
			kr = med + 1
		} else {
			found = med
		}
	}
	return found
}

func menuCariMood(T *arrMood, n int) {
	var pilihanCari int
	fmt.Println("\n--- Pencarian Catatan Mood ---")
	fmt.Println("1. Sequential Search (berdasarkan tanggal)")
	fmt.Println("2. Binary Search (berdasarkan emosi)")
	fmt.Print("Pilih metode pencarian: ")
	pilihanCari = inputInt()

	if pilihanCari == 1 {
		var kunci string
		fmt.Print("Masukkan tanggal yang dicari (DD/MM/YYYY): ")
		kunci = inputString()
		var idx int = seqSearchMood(*T, n, kunci)
		if idx == -1 {
			fmt.Println("Data tidak ditemukan.")
		} else {
			fmt.Println("Data ditemukan:")
			var i int = 0
			for i < n {
				if T[i].tanggal == kunci {
					fmt.Printf("ID: %d | Skor: %s | Deskripsi: %s | Tanggal: %s | Minggu Ke: %d\n",
						T[i].id, skorEmosiString(T[i].skorEmosi), T[i].deskripsi, T[i].tanggal, T[i].mingguKe)
				}
				i = i + 1
			}
		}
	} else if pilihanCari == 2 {
		var skor int
		fmt.Println("1. Terrible")
		fmt.Println("2. Bad")
		fmt.Println("3. Fine")
		fmt.Println("4. Good")
		fmt.Println("5. Awesome")
		fmt.Print("Masukkan emosi yang dicari (1-5): ")
		skor = inputInt()

		selectionSortMood(T, n)

		var idx int = binarySearchMood(*T, n, skor)
		if idx == -1 {
			fmt.Println("Data tidak ditemukan.")
		} else {
			fmt.Println("Data ditemukan:")
			var left int = idx
			for left > 0 && T[left-1].skorEmosi == skor {
				left = left - 1
			}
			var right int = idx
			for right < n-1 && T[right+1].skorEmosi == skor {
				right = right + 1
			}
			var i int = left
			for i <= right {
				fmt.Printf("ID: %d | Skor: %s | Deskripsi: %s | Tanggal: %s | Minggu Ke: %d\n",
					T[i].id, skorEmosiString(T[i].skorEmosi), T[i].deskripsi, T[i].tanggal, T[i].mingguKe)
				i = i + 1
			}
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func menuCariTugas(T *arrTugas, n int) {
	var pilihanCari int
	fmt.Println("\n--- Pencarian Tugas ---")
	fmt.Println("1. Sequential Search (berdasarkan nama tugas)")
	fmt.Println("2. Binary Search (berdasarkan durasi)")
	fmt.Print("Pilih metode pencarian: ")
	pilihanCari = inputInt()

	if pilihanCari == 1 {
		var kunci string
		fmt.Print("Masukkan nama tugas yang dicari: ")
		kunci = inputString()
		var idx int = seqSearchTugas(*T, n, kunci)
		if idx == -1 {
			fmt.Println("Data tidak ditemukan.")
		} else {
			fmt.Println("Data ditemukan:")
			var i int = 0
			for i < n {
				if T[i].namaTugas == kunci {
					var labelPrioritas string
					if T[i].prioritas == 1 {
						labelPrioritas = "Tinggi"
					} else if T[i].prioritas == 2 {
						labelPrioritas = "Sedang"
					} else {
						labelPrioritas = "Rendah"
					}
					fmt.Printf("ID: %d | Tugas: %s | Durasi: %d menit | Prioritas: %s\n",
						T[i].id, T[i].namaTugas, T[i].durasi, labelPrioritas)
				}
				i = i + 1
			}
		}
	} else if pilihanCari == 2 {
		var durasi int
		fmt.Print("Masukkan durasi yang dicari (menit): ")
		durasi = inputInt()

		insertionSortTugasDurasi(T, n)

		var idx int = binarySearchTugas(*T, n, durasi)
		if idx == -1 {
			fmt.Println("Data tidak ditemukan.")
		} else {
			fmt.Println("Data ditemukan:")
			var left int = idx
			for left > 0 && T[left-1].durasi == durasi {
				left = left - 1
			}
			var right int = idx
			for right < n-1 && T[right+1].durasi == durasi {
				right = right + 1
			}
			var i int = left
			for i <= right {
				fmt.Printf("ID: %d | Tugas: %s | Durasi: %d menit\n",
					T[i].id, T[i].namaTugas, T[i].durasi)
				i = i + 1
			}
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func selectionSortMood(T *arrMood, n int) {
	var i, j, idxMin int
	var temp CatatanMood
	i = 1
	for i <= n-1 {
		idxMin = i - 1
		j = i
		for j < n {
			if T[idxMin].skorEmosi > T[j].skorEmosi {
				idxMin = j
			}
			j = j + 1
		}
		temp = T[idxMin]
		T[idxMin] = T[i-1]
		T[i-1] = temp
		i = i + 1
	}
	fmt.Println("Catatan mood berhasil diurutkan berdasarkan skor emosi (ascending) - Selection Sort.")
}

func insertionSortTugasPrioritas(T *arrTugas, n int) {
	var i, j int
	var temp Tugas
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp.prioritas < T[j-1].prioritas {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
	fmt.Println("Tugas berhasil diurutkan berdasarkan prioritas (Tinggi -> Rendah) - Insertion Sort.")
}

func insertionSortTugasDurasi(T *arrTugas, n int) {
	var i, j int
	var temp Tugas
	i = 1
	for i <= n-1 {
		j = i
		temp = T[j]
		for j > 0 && temp.durasi < T[j-1].durasi {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
		i = i + 1
	}
	fmt.Println("Tugas berhasil diurutkan berdasarkan durasi (terpendek -> terlama) - Insertion Sort.")
}

func menuUrutMood(T *arrMood, n int) {
	var pilihanUrut int
	fmt.Println("\n--- Pengurutan Catatan Mood ---")
	fmt.Println("1. Selection Sort berdasarkan skor emosi (ascending)")
	fmt.Print("Pilih metode pengurutan: ")
	pilihanUrut = inputInt()
	if pilihanUrut == 1 {
		selectionSortMood(T, n)
		tampilMood(*T, n)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func menuUrutTugas(T *arrTugas, n int) {
	var pilihanUrut int
	fmt.Println("\n--- Pengurutan Tugas ---")
	fmt.Println("1. Insertion Sort berdasarkan prioritas (Tinggi -> Rendah)")
	fmt.Println("2. Insertion Sort berdasarkan durasi (terpendek -> terlama)")
	fmt.Print("Pilih metode pengurutan: ")
	pilihanUrut = inputInt()
	if pilihanUrut == 1 {
		insertionSortTugasPrioritas(T, n)
		tampilTugas(*T, n)
	} else if pilihanUrut == 2 {
		insertionSortTugasDurasi(T, n)
		tampilTugas(*T, n)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func statistikMood(T arrMood, n int) {
	if n == 0 {
		fmt.Println("Belum ada catatan mood untuk ditampilkan statistiknya.")
		return
	}

	fmt.Println("\n=== Statistik Tren Suasana Hati Mingguan (per Bulan) ===")
	fmt.Println("--------------------------------------------------------")

	var mingguSudahDiproses [NMAX]int
	var nUnik int = 0

	var i int = 0
	for i < n {
		var mg int = T[i].tahun*1000 + T[i].bulan*10 + T[i].mingguKe
		var sudah bool = false
		var j int = 0
		for j < nUnik {
			if mingguSudahDiproses[j]8 == mg {
				sudah = true
			}
			j = j + 1
		}

		if !sudah {
			mingguSudahDiproses[nUnik] = mg
			nUnik = nUnik + 1
		}
		i = i + 1
	}

	var x int = 1
	for x < nUnik {
		var y int = x
		var temp int = mingguSudahDiproses[y]
		for y > 0 && temp < mingguSudahDiproses[y-1] {
			mingguSudahDiproses[y] = mingguSudahDiproses[y-1]
			y = y - 1
		}
		mingguSudahDiproses[y] = temp
		x = x + 1
	}

	var u int = 0
	for u < nUnik {
		var mg int = mingguSudahDiproses[u]
		var total int = 0
		var jumlah int = 0
		var k int = 0
		for k < n {
			if T[k].tahun*1000+T[k].bulan*10+T[k].mingguKe == mg {
				total = total + T[k].skorEmosi
				jumlah = jumlah + 1
			}
			k = k + 1
		}

		var rata float64 = float64(total) / float64(jumlah)
		var tahun int = mg / 1000
		var bulan int = (mg / 10) % 100
		var minggu int = mg % 10
		fmt.Printf("Tahun %d Bulan %d Minggu ke-%d : %.2f (dari %d catatan)\n", tahun, bulan, minggu, rata, jumlah)
		u = u + 1
	}
	fmt.Println("--------------------------------------------------------")
}

func statistikTugas(T arrTugas, n int) {
	if n == 0 {
		fmt.Println("Belum ada tugas untuk ditampilkan statistiknya.")
		return
	}
	var selesai int = 0
	var i int = 0
	for i < n {
		if T[i].selesai {
			selesai = selesai + 1
		}
		i = i + 1
	}
	var belum int = n - selesai
	var persentase float64 = float64(selesai) / float64(n) * 100.0

	fmt.Println("\n=== Statistik Tingkat Penyelesaian Tugas ===")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Total Tugas    : %d\n", n)
	fmt.Printf("Selesai        : %d\n", selesai)
	fmt.Printf("Belum Selesai  : %d\n", belum)
	fmt.Printf("Persentase     : %.2f%%\n", persentase)
	fmt.Println("---------------------------------------------")
}

func main() {
	var dataMood arrMood
	var dataTugas arrTugas
	var nMood int = 0
	var nTugas int = 0
	var idMood int = 0
	var idTugas int = 0

	var pilihan int
	var subPilihan int

	fmt.Println("=================================================")
	fmt.Println("           Selamat Datang di MindFlow            ")
	fmt.Println("Asisten Virtual Kesehatan Mental & Produktivitas")
	fmt.Println("=================================================")

	pilihan = 0
	for pilihan != 4 {
		fmt.Println("\n=== MENU UTAMA ===")
		fmt.Println("1. Catatan Suasana Hati (Mood)")
		fmt.Println("2. Daftar Tugas Harian")
		fmt.Println("3. Statistik")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")
		pilihan = inputInt()

		if pilihan == 1 {
			subPilihan = 0
			for subPilihan != 7 {
				fmt.Println("\n--- Menu Catatan Mood ---")
				fmt.Println("1. Tambah Catatan Mood")
				fmt.Println("2. Tampilkan Semua Catatan Mood")
				fmt.Println("3. Ubah Catatan Mood")
				fmt.Println("4. Hapus Catatan Mood")
				fmt.Println("5. Cari Catatan Mood")
				fmt.Println("6. Urutkan Catatan Mood")
				fmt.Println("7. Kembali ke Menu Utama")
				fmt.Print("Pilih menu: ")
				subPilihan = inputInt()

				if subPilihan == 1 {
					tambahMood(&dataMood, &nMood, &idMood)
				} else if subPilihan == 2 {
					tampilMood(dataMood, nMood)
				} else if subPilihan == 3 {
					ubahMood(&dataMood, nMood)
				} else if subPilihan == 4 {
					hapusMood(&dataMood, &nMood)
				} else if subPilihan == 5 {
					menuCariMood(&dataMood, nMood)
				} else if subPilihan == 6 {
					menuUrutMood(&dataMood, nMood)
				} else if subPilihan == 7 {
					fmt.Println("Kembali ke menu utama...")
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		} else if pilihan == 2 {
			subPilihan = 0
			for subPilihan != 8 {
				fmt.Println("\n--- Menu Daftar Tugas ---")
				fmt.Println("1. Tambah Tugas")
				fmt.Println("2. Tampilkan Semua Tugas")
				fmt.Println("3. Ubah Tugas")
				fmt.Println("4. Hapus Tugas")
				fmt.Println("5. Cari Tugas")
				fmt.Println("6. Urutkan Tugas")
				fmt.Println("7. Tandai Tugas Selesai")
				fmt.Println("8. Kembali ke Menu Utama")
				fmt.Print("Pilih menu: ")
				subPilihan = inputInt()

				if subPilihan == 1 {
					tambahTugas(&dataTugas, &nTugas, &idTugas)
				} else if subPilihan == 2 {
					tampilTugas(dataTugas, nTugas)
				} else if subPilihan == 3 {
					ubahTugas(&dataTugas, nTugas)
				} else if subPilihan == 4 {
					hapusTugas(&dataTugas, &nTugas)
				} else if subPilihan == 5 {
					menuCariTugas(&dataTugas, nTugas)
				} else if subPilihan == 6 {
					menuUrutTugas(&dataTugas, nTugas)
				} else if subPilihan == 7 {
					tandaiSelesai(&dataTugas, nTugas)
				} else if subPilihan == 8 {
					fmt.Println("Kembali ke menu utama...")
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		} else if pilihan == 3 {
			subPilihan = 0
			for subPilihan != 3 {
				fmt.Println("\n--- Menu Statistik ---")
				fmt.Println("1. Tren Suasana Hati Mingguan")
				fmt.Println("2. Tingkat Penyelesaian Tugas")
				fmt.Println("3. Kembali ke Menu Utama")
				fmt.Print("Pilih menu: ")
				subPilihan = inputInt()

				if subPilihan == 1 {
					statistikMood(dataMood, nMood)
				} else if subPilihan == 2 {
					statistikTugas(dataTugas, nTugas)
				} else if subPilihan == 3 {
					fmt.Println("Kembali ke menu utama...")
				} else {
					fmt.Println("Pilihan tidak valid.")
				}
			}
		} else if pilihan == 4 {
			fmt.Println("Terima kasih telah menggunakan MindFlow. Sampai jumpa!")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

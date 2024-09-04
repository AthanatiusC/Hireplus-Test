package main

import (
	"fmt"
	"time"
)

// CutiPribadi struct untuk menyimpan informasi cuti
type CutiPribadi struct {
	JumlahCutiBersama int
	TanggalJoin       time.Time
	TanggalCuti       time.Time
	DurasiCuti        int
}

// BolehCutiPribadi function untuk menentukan apakah karyawan boleh cuti
func BolehCutiPribadi(cuti CutiPribadi) (bool, string) {
	sisaHariTahunPertama := 365 - int(cuti.TanggalJoin.YearDay()) + 180

	kuotaCutiPribadi := int(float64(sisaHariTahunPertama) / 365 * float64(14-cuti.JumlahCutiBersama))

	if cuti.TanggalCuti.Before(cuti.TanggalJoin.AddDate(0, 0, 180)) {
		return false, "Belum melewati masa tunggu 180 hari"
	}

	if cuti.DurasiCuti > kuotaCutiPribadi {
		return false, "Kuota cuti pribadi tidak mencukupi"
	}

	if cuti.DurasiCuti > 3 {
		return false, "Durasi cuti melebihi batas maksimal 3 hari"
	}

	return true, "Boleh mengambil cuti"
}

func main() {
	var jumlahCutiBersama, durasiCuti int
	var tanggalJoin, tanggalCuti time.Time
	var tglJStr, tglCStr string
	fmt.Print("Jumlah Cuti Bersama : ")
	fmt.Scanln(&jumlahCutiBersama)
	fmt.Print("Tanggal Join : ")
	fmt.Scanln(&tglJStr)
	tanggalJoin, err := time.Parse("2006-01-02", tglJStr)
	if err != nil {
		panic("tanggal join tidak valid")
	}
	fmt.Print("Tanggal Cuti :")
	fmt.Scanln(&tglCStr)
	tanggalCuti, err = time.Parse("2006-01-02", tglCStr)
	if err != nil {
		panic("tanggal cuti tidak valid")
	}
	fmt.Print("Durasi Cuti :")
	fmt.Scanln(&durasiCuti)

	cuti := CutiPribadi{
		JumlahCutiBersama: jumlahCutiBersama,
		TanggalJoin:       tanggalJoin,
		TanggalCuti:       tanggalCuti,
		DurasiCuti:        durasiCuti,
	}

	boleh, alasan := BolehCutiPribadi(cuti)
	fmt.Printf("Boleh cuti: %t, Alasan: %s\n", boleh, alasan)
}

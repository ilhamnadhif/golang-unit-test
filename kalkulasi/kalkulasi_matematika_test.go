package kalkulasi

import "testing"

func TestTambah(t *testing.T) {
	result := Tambah(3, 5)
	if result != 8 {
		panic("error")
	}
}

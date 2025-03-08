# Result

Hasil dari codereview dari kode yang diberikan terdapat beberapa masalah seperti


### Race Condition
Race condition pada kode tersebut disebabkan variable sharedCounter diakses secara concurrent oleh 3 goroutine
Solusi: harus ada mekanisme lock, bisa menggnakan Mutex

mu.Lock()
sharedCounter++
mu.Unlock()


### Deadlock
Deadlock pada kode tersebut disebabkan resultChan sudah ditutup pada processData, ini terjadi karena menggunakan goroutine sehingga kita tidak bisa memastikan fungsi mana yang akan dijalankan terlebih dahulu
Solusi: tutup semua channel setelah semua worker selesai

wg.Wait() // Pastikan semua worker selesai
close(resultChan)
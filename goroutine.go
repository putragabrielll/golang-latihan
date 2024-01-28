package main

import (
	"fmt"
	"sync"
	"time"
)


func cetak(num int){
	fmt.Println("Urutan", num)
	time.Sleep(time.Second)
	fmt.Println("Lanjutan", num)
}


func Routine(angka int){
	var wg sync.WaitGroup // menggunakan WaitGroup
	for i := 0; i < angka; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done() // menunggu higga selesai terlebih dahulu baru menjalankan yg lain.
			cetak(i)
		}()
	}
	wg.Wait()
}
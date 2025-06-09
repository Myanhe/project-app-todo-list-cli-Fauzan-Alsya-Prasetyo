package main

import (
    "flag"
    "fmt"
    "todoapp/models"
    "todoapp/utils"
)

func main() {
    // Definisi flag CLI
    addFlag := flag.String("add", "", "Menambahkan tugas baru dengan judul")
    descFlag := flag.String("desc", "", "Deskripsi tugas (opsional)")
    listFlag := flag.Bool("list", false, "Menampilkan semua tugas")
    doneFlag := flag.Int("done", 0, "Menandai tugas selesai berdasarkan ID")
    deleteFlag := flag.Int("delete", 0, "Menghapus tugas berdasarkan ID")
    searchFlag := flag.String("search", "", "Mencari tugas berdasarkan kata kunci")
    flag.Parse()

    // Nama file data tugas
    const filename = "tasks.json"
    // Muat data tugas dari file (jika belum ada, akan kosong)
    tasks, err := utils.LoadTasks(filename)
    if err != nil {
        fmt.Println("Gagal membaca file tugas:", err)
        return
    }

    // Jika flag --list diaktifkan, tampilkan semua tugas
    if *listFlag {
        utils.ListTasks(tasks)
        return
    }
    // Jika flag --search diaktifkan, cari dan tampilkan tugas yang cocok
    if *searchFlag != "" {
        utils.SearchTasks(tasks, *searchFlag)
        return
    }
    // Jika flag --add diaktifkan, tambahkan tugas baru
    if *addFlag != "" {
        // Panggil fungsi AddTask untuk validasi dan penambahan
        if err := utils.AddTask(&tasks, *addFlag, *descFlag); err != nil {
            fmt.Println("Error:", err)
            return
        }
        // Simpan perubahan ke file JSON
        if err := utils.SaveTasks(filename, tasks); err != nil {
            fmt.Println("Gagal menyimpan tugas:", err)
        } else {
            fmt.Println("Tugas berhasil ditambahkan.")
        }
        return
    }
    // Jika flag --done diaktifkan, tandai tugas selesai
    if *doneFlag != 0 {
        if err := utils.MarkTaskDone(&tasks, *doneFlag); err != nil {
            fmt.Println("Error:", err)
            return
        }
        if err := utils.SaveTasks(filename, tasks); err != nil {
            fmt.Println("Gagal menyimpan perubahan:", err)
        } else {
            fmt.Println("Tugas berhasil ditandai selesai.")
        }
        return
    }
    // Jika flag --delete diaktifkan, hapus tugas
    if *deleteFlag != 0 {
        if err := utils.DeleteTask(&tasks, *deleteFlag); err != nil {
            fmt.Println("Error:", err)
            return
        }
        if err := utils.SaveTasks(filename, tasks); err != nil {
            fmt.Println("Gagal menyimpan perubahan:", err)
        } else {
            fmt.Println("Tugas berhasil dihapus.")
        }
        return
    }
}

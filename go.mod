module todo-list

go 1.22.2

// project-app-todo-list-cli/
//
// ├── cmd/              # Berisi command CLI (cobra)
// │   ├── add.go
// │   ├── list.go
// │   ├── done.go
// │   └── delete.go
// ├── data/             # Folder untuk menyimpan database JSON
// │   └── tasks.json    # File database berisi daftar to-do
// ├── models/           # Berisi struct dan fungsi terkait Task
// │   └── task.go
// ├── utils/            # Berisi fungsi bantu (load/save JSON, validasi)
// │   └── file.go
// ├── main.go           # Entry point
// ├── go.mod
// └── go.sum
package main

import (
    "fmt"
    "os"

    "coobak/internal/chunker"
    "coobak/internal/store"
    "coobak/internal/versioning"
    "coobak/pkg/types"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: coobak backup|restore <filepath>")
        return
    }

    command := os.Args[1]
    filePath := os.Args[2]
    cas := store.CAS{BasePath: ".coobak"}

    switch command {
    case "backup":
        chunks, err := chunker.ChunkFile(filePath)
        if err != nil {
            fmt.Println("Chunking error:", err)
            return
        }

        var hashes [][32]byte
        for _, chunk := range chunks {
            cas.Store(chunk.Hash, chunk.Data)
            hashes = append(hashes, chunk.Hash)
        }

        manifest := types.Manifest{
            FilePath: filePath,
            Chunks:   hashes,
            Time:     "manual-backup",
        }

        err = versioning.SaveManifest(".coobak", filePath, manifest)
        if err != nil {
            fmt.Println("Manifest save error:", err)
        } else {
            fmt.Println("Backup complete.")
        }
    
    default:
        fmt.Println("Unsupported command.")
    }



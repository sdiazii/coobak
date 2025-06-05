package main

import (
    "encoding/json"
    "fmt"
    "os"
    "path/filepath"
    "time"

    "coobak/internal/chunker"
    "coobak/internal/store"
    "coobak/internal/versioning"
    "coobak/pkg/types"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage:")
        fmt.Println("  coobak backup <file>")
        fmt.Println("  coobak restore <file> <version>")
        fmt.Println("  coobak list <file>")
        return
    }

    command := os.Args[1]
    filePath := os.Args[2]

    switch command {
    case "backup":
        err := backupFile(filePath)
        if err != nil {
            fmt.Println("Backup error:", err)
        } else {
            fmt.Println("Backup complete.")
        }

    case "restore":
        if len(os.Args) < 4 {
            fmt.Println("Usage: coobak restore <file> <version>")
            return
        }
        version := os.Args[3]
        err := restoreFile(filePath, version)
        if err != nil {
            fmt.Println("Restore error:", err)
        }

    case "list":
        err := listVersions(filePath)
        if err != nil {
            fmt.Println("List error:", err)
        }

    default:
        fmt.Println("Unsupported command:", command)
    }
}

func backupFile(filePath string) error {
    chunks, err := chunker.ChunkFile(filePath)
    if err != nil {
        return fmt.Errorf("chunking error: %w", err)
    }

    cas := store.CAS{BasePath: ".coobak"}
    var hashes [][32]byte
    for _, chunk := range chunks {
        cas.Store(chunk.Hash, chunk.Data)
        hashes = append(hashes, chunk.Hash)
    }

    manifest := types.Manifest{
        FilePath: filePath,
        Chunks:   hashes,
        Time:     time.Now().Format("20060102_150405"),
    }

    return versioning.SaveManifest(".coobak", filePath, manifest)
}

func restoreFile(filePath string, version string) error {
    manifestFile := version + ".json"
    manifestPath := filepath.Join(".coobak", "versions", filePath, manifestFile)

    manifestData, err := os.ReadFile(manifestPath)
    if err != nil {
        return fmt.Errorf("cannot read manifest: %w", err)
    }

    var manifest types.Manifest
    if err := json.Unmarshal(manifestData, &manifest); err != nil {
        return fmt.Errorf("invalid manifest: %w", err)
    }

    cas := store.CAS{BasePath: ".coobak"}
    outFile, err := os.Create(filePath + ".restored")
    if err != nil {
        return fmt.Errorf("cannot create restored file: %w", err)
    }
    defer outFile.Close()

    for _, hash := range manifest.Chunks {
        block, err := cas.Retrieve(hash)
        if err != nil {
            return fmt.Errorf("missing block: %x", hash)
        }
        _, err = outFile.Write(block)
        if err != nil {
            return fmt.Errorf("write error: %w", err)
        }
    }

    fmt.Println("Restored to:", filePath + ".restored")
    return nil
}

func listVersions(filePath string) error {
    dir := filepath.Join(".coobak", "versions", filePath)
    entries, err := os.ReadDir(dir)
    if err != nil {
        return fmt.Errorf("cannot list versions: %w", err)
    }

    fmt.Println("Available versions for", filePath)
    for _, entry := range entries {
        if !entry.IsDir() && entry.Name() != "latest.json" {
            fmt.Println("  ", entry.Name())
        }
    }
    return nil
}
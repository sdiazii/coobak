package store

import (
    "encoding/hex"
    "os"
    "path/filepath"
)

type CAS struct {
    BasePath string
}

func (c *CAS) Store(hash [32]byte, data []byte) error {
    hexHash := hex.EncodeToString(hash[:])
    path := filepath.Join(c.BasePath, "blocks", hexHash)

    if _, err := os.Stat(path); err == nil {
        return nil
    }

    os.MkdirAll(filepath.Dir(path), 0755)
    return os.WriteFile(path, data, 0644)
}

func (c *CAS) Retrieve(hash [32]byte) ([]byte, error) {
    hexHash := hex.EncodeToString(hash[:])
    path := filepath.Join(c.BasePath, "blocks", hexHash)
    return os.ReadFile(path)
}

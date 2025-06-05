package chunker

import (
    "crypto/sha256"
    "io"
    "os"
)

const ChunkSize = 4096

type Chunk struct {
    Hash [32]byte
    Data []byte
}

func ChunkFile(path string) ([]Chunk, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var chunks []Chunk
    buf := make([]byte, ChunkSize)

    for {
        n, err := file.Read(buf)
        if err != nil && err != io.EOF {
            return nil, err
        }
        if n == 0 {
            break
        }
        data := make([]byte, n)
        copy(data, buf[:n])
        chunks = append(chunks, Chunk{
            Hash: sha256.Sum256(data),
            Data: data,
        })
    }

    return chunks, nil
}

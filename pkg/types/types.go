package types

type Manifest struct {
    FilePath string      `json:"file_path"`
    Chunks   [][32]byte  `json:"chunks"`
    Time     string      `json:"timestamp"`
}

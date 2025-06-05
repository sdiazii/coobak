package versioning

import (
    "encoding/json"
    "os"
    "path/filepath"
    "fmt"
    "time"
    "coobak/pkg/types"
)

func SaveManifest(basePath, filePath string, manifest types.Manifest) error {
    dir := filepath.Join(basePath, "versions", filePath)
    os.MkdirAll(dir, 0755)

    out, _ := json.MarshalIndent(manifest, "", "  ")
    timestamp := time.Now().Format("20060102_150405")
    filename := fmt.Sprintf("%s.json", timestamp)
    return os.WriteFile(filepath.Join(dir, filename), out, 0644)
}

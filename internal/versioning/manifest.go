package versioning

import (
    "encoding/json"
    "os"
    "path/filepath"

    "coobak/pkg/types"
)

func SaveManifest(basePath, filePath string, manifest types.Manifest) error {
    dir := filepath.Join(basePath, "versions", filePath)
    os.MkdirAll(dir, 0755)

    out, _ := json.MarshalIndent(manifest, "", "  ")
    return os.WriteFile(filepath.Join(dir, "version.json"), out, 0644)
}

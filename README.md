# 🧱 coobak — Cooperative Backup System

**coobak** is a lightweight, versioned backup tool inspired by Venti and Elephant File System. It splits files into deduplicated, content-addressable chunks and stores them locally along with a manifest of historical versions — making it easy to restore any previous state of your files.

---

## 📦 Features

- ✅ Chunk-based deduplication (content-addressable storage)
- 📁 Automatic version tracking (timestamped manifests)
- 🔁 File restoration from any saved version
- 📂 Modular and extendable Go codebase
- 🪄 Cross-platform support (Windows, Linux, macOS)

---

## 🚀 Quick Start

### 🔧 Build

Make sure Go is installed.

```bash
go build -o coobak.exe ./cmd/coobak
```

### 📥 Backup a File

```bash
./coobak.exe backup example.txt
```

This will:
- Chunk the file into 4KB blocks
- Save each block into `.coobak/blocks/`
- Save a manifest into `.coobak/versions/example.txt/YYYYMMDD_HHMMSS.json`

---

### 📋 List All Versions

```bash
./coobak.exe list example.txt
```

Example output:

```
Available versions for example.txt
   20250604_230312.json
   20250605_104120.json
```

---

### 🔄 Restore a Version

```bash
./coobak.exe restore example.txt 20250604_230312
```

Result:

```
Restored to: example.txt.restored
```

Creates a new file with the restored content based on the specified version.

---

## 🛠 Project Structure

```
coobak/
├── cmd/coobak/         # CLI entry point
├── internal/
│   ├── chunker/        # File chunking and hashing
│   ├── store/          # Content-addressable storage (CAS)
│   ├── versioning/     # Version manifest creation
├── pkg/types/          # Shared data types
```

---

## 📚 How It Works

- Files are split into 4KB chunks
- Each chunk is stored as a file named by its SHA-256 hash
- A JSON manifest tracks the order of chunks and the version timestamp
- The restore command reassembles chunks into the original file

---

## 📁 Output Directory

All backups are stored in a hidden `.coobak/` directory:

```
.coobak/
├── blocks/             # Stored file chunks by hash
├── versions/           # One folder per file
│   └── example.txt/
│       ├── 20250604_230312.json
│       ├── 20250605_104120.json
```

---

## 🧩 Future Improvements

- 🔐 Optional encryption
- ☁️ Peer-to-peer replication
- 🧹 Garbage collection
- 🔄 Auto-backup via filesystem watcher

---

## 📜 License

MIT License. Feel free to adapt and extend this system for your needs.

---

## 👤 Author

Designed and prototyped in Go as a personal exploration of content-addressable, versioned, and decentralized storage systems.
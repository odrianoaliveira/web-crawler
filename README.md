# 🕷️ Web Crawler in Go

A web crawler written in Go. It starts with a list of seed URLs, downloads and parses HTML content, extracts useful information, and stores it in MongoDB.

---

## 🧩 Features

- Load initial seed URLs
- Maintain a URL frontier to avoid duplicates
- Download pages via HTTP
- Parse HTML to extract content and links
- Store structured data in MongoDB
- Built using Go



## 🚀 Getting Started

### Prerequisites

- [Go 1.20+](https://golang.org/doc/install)
- [MongoDB](https://www.mongodb.com/try/download/community)

### Clone and Build

```bash
git clone https://github.com/yourusername/web-crawler.git
cd web-crawler
go build -o crawler ./cmd/crawler

```

### Run the Crawler

```bash
./crawler -seed-url urls.txt
```

### Configuration
Edit configs/config.yaml to set the MongoDB connection and other runtime parameters:

```yaml
mongo_uri: "mongodb://localhost:27017"
database: "crawler"
collection: "pages"
```

### Running Tests

```
go test ./...
```
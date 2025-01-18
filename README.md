# 🌐 Web Crawler (Go)

This repository contains the **Web Crawler** project, developed as part of the [Boot.dev course](https://www.boot.dev/courses/build-web-crawler-golang). The goal of this project is to create a CLI-based web crawler in Go, reinforcing key backend development concepts.

## 🚀 Features

- **Recursive Crawling**: Traverse and fetch links recursively from web pages.
- **Concurrency**: Leverage Go routines for efficient parallel crawling.
- **Custom Depth Control**: Set limits on how deeply the crawler traverses links.
- **Error Handling**: Gracefully manage timeouts and invalid URLs.
- **Output Summary**: Present crawled URLs in a clear, readable format.

## 🛠️ Technologies Used

- **Go**: Core language for development.
- **Concurrency**: Using Go routines and channels.
- **CLI Design**: Build and manage command-line interactions.
- **Testing**: Robust unit tests with Go's `testing` package.

## 📚 What I Learned

- Implementing concurrency with Go routines and channels.
- Parsing and managing HTML content in Go.
- Error handling and timeouts in HTTP requests.
- Designing effective CLI tools in Go.
- Writing clean, maintainable, and testable Go code.

## 🧪 Testing

Unit tests were written to verify the core functionality of the crawler, including:

- Proper traversal of links.
- Handling invalid or unreachable URLs.
- Adhering to depth limits during recursion.

Run tests with:
```bash
go test ./...
```

## 🌟 Why This Project?

The Web Crawler project was built to deepen my understanding of backend development, specifically:

- Gaining practical experience with Go's concurrency model.
- Exploring the challenges of web scraping and crawling.
- Building a scalable and efficient tool for recursive link traversal.

## 📂 Project Structure

```
├── crawler/         # Core crawler logic
├── cmd/             # CLI implementation
├── tests/           # Unit tests
└── README.md        # Project documentation
```

## 🔗 Related Resources

- [Boot.dev Course](https://www.boot.dev/courses/build-web-crawler-golang)
- [Official Go Documentation](https://golang.org/doc/)

---

Feel free to explore, test, and contribute to this project! 🚀

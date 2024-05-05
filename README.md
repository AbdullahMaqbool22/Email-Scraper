## Email Scraper

This Go program is designed to scrape email addresses associated with a given domain from various APIs and save them to a file. It utilizes JSON API responses for retrieving email data and writing them to a text file.

### Prerequisites

- Go programming language installed
- Access to APIs with domain email data

### Usage

1. Clone this repository to your local machine.
2. Replace placeholders in the `main` function of `main.go` with your domain, API URLs, and API keys.
3. Run the program by executing `go run main.go`.
4. The program will scrape email addresses associated with the provided domain from the specified APIs and save them to a file named `domain_emails.txt`.

### Configuration

In the `main` function of `main.go`, you can configure the following:

- `domain.Domain`: Replace `"Your domain here"` with the target domain.
- `domain.APIs`: Add API URLs and their respective API keys within the slice. You can add multiple APIs if necessary.

### Contributing

Contributions are welcome! If you find any bugs or have suggestions for improvements, feel free to open an issue or submit a pull request.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### Acknowledgments

Special thanks to contributors and developers of the libraries used in this project.

### Disclaimer

This program is intended for educational and informational purposes only. Use it responsibly and ensure compliance with applicable laws and regulations when scraping email data.

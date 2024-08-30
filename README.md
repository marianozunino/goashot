# Go Ashot

Go Ashot is a Proof of Concept (POC) project that demonstrates the power of Go libraries and frameworks in creating a web application for automating food orders from Rappi.

## 🚀 Project Overview

This project showcases the integration of modern web technologies and Go frameworks to create a seamless food ordering experience. It was built to explore and learn:

- [Go Rod](https://go-rod.github.io/#/): A high-level Chrome DevTools Protocol driver for web automation
- [HTMX](https://htmx.org/): A library for creating dynamic web applications without JavaScript
- [templ](https://github.com/a-h/templ): A type-safe HTML templating language for Go
- [Echo](https://echo.labstack.com/): A high-performance, extensible web framework for Go
- [Cobra](https://github.com/spf13/cobra): A powerful library for creating modern CLI applications

> 💡 Inspiration: We're a group of Shawarma enthusiasts 🥙 at the office who decided to streamline our lunch ordering process!

## ✨ Features

- Intuitive web interface for CRUD operations on food orders
- Automated crawler that places orders on Rappi based on web interface input
- Seamless integration with Rappi's two-factor authentication (2FA) system
- Robust data persistence layer for order management
- Dynamic web interfaces powered by HTMX, eliminating the need for client-side JavaScript
- Efficient and type-safe HTML templating with templ
- Flexible command-line interface for configuration using Cobra

## 🛠 Installation & Usage

### Building the Project

```bash
make build
```

This command generates a binary in the `bin` directory.

### Running the Application

Start the web server:

```bash
./bin/app backoffice [-p 5000] [-a 0.0.0.0]
```

Launch the crawler:

```bash
./bin/app crawler [-p 5001] [-a 0.0.0.0]
```

Default ports are 5000 for the web server and 5001 for the crawler, both binding to `localhost`.

### Crawler Usage

1. The crawler navigates to the Rappi login page.
2. Complete the login process manually, including 2FA if required.
3. Once authenticated, the crawler automatically processes orders created through the web interface.

## 📁 Project Structure

```
go-ashot/
├── internal/
│   ├── handlers/    # HTTP request handlers
│   ├── models/      # Data structures and business logic
│   ├── services/    # Business logic and external service integrations
│   ├── templates/   # HTMX and templ UI components
│   └── repository/  # Data persistence layer
├── cmd/             # Command-line interface definitions
```

## 🛡️ License

This project is open-source and available under the [MIT License](LICENSE).

## 📞 Support

None

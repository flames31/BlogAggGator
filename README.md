# BlogAggGator 📰🚀

BlogAggGator is an RSS feed aggregator that continuously fetches blog posts from multiple sources and keeps track of the latest updates. It prioritizes fetching the oldest unseen feeds first, ensuring an up-to-date and efficient feed scraping process.

## Features 🌟
- Add new RSS feeds easily  
- Automatically fetch new posts from subscribed feeds  
- Prioritizes fetching the least recently updated feeds  
- Stores fetched posts in a database for easy retrieval  
- Command-line interface for managing feeds  

## Installation ⚙️
### Prerequisites
- Go 1.20+ installed  
- PostgreSQL database set up  

### Clone the Repository
```sh
git clone https://github.com/flames31/BlogAggGator.git
cd BlogAggGator
```

### Install Dependencies
```sh
go mod tidy
```

### Build the Executable
```sh
go build -o blogagg
```

Or install it globally:
```sh
go install .
```

## Usage 🛠️
### Adding a Feed
```sh
./blogagg addfeed "TechCrunch" "https://techcrunch.com/feed/"
```

### Running the Fetcher
```sh
./blogagg fetch
```

### Database Migrations (using Goose)
```sh
goose up
```

## Configuration ⚙️
Edit the `config.yaml` file to set up database credentials and other settings.

## Contributing 🤝
Pull requests are welcome! If you find a bug or have an idea for an improvement, feel free to open an issue or submit a PR.

## License 📝
This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

---
Made with ❤️ by [Flames31](https://github.com/flames31)


# Simple Web Scraper

A basic web scraper written in Go using the Colly framework to collect data from the Best Buy website on macbooks.

## Features

- Scrapes macbook information from the Best Buy website.
- Collects macbook's names, prices, and image URLs.
- Exports the scraped data to a JSON file for analysis.

## Prerequisites

- You should have Go installed on your local machine.

## Getting Started

1. Clone this repository to your local machine:

   ```shell
   git clone https://github.com/AdamElmaghraby/Simple-Web-Scraper---First-Project.git

2. Install Go dependencies:

    go get github.com/gocolly/colly

3. Run the web scraper:

    go run main.go


4. The scraped data will be exported to a JSON file named "bestbuy_macbook_data.json."

## Customization

- Modify the web scraper to target other websites or collect additional data by updating the scraping logic in `main.go`.

## Note

Due to Best Buy's page structure, it's currently limited to only collecting css elements when using colly. Therefore, only collects first 3 listing's prices.


## Contributing

If you have ideas for improvements or find any issues, please feel free to contribute. You can fork the repository and submit a pull request.



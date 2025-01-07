Golang Web Crawler

This Golang-based web crawler fetches product URLs from multiple e-commerce websites concurrently using goroutines and channels for efficient parallel processing.

Features

Crawls multiple e-commerce websites concurrently.
Skips unsupported or invalid URLs.
Prevents duplicate crawling using a hasVisited map.
Concurrency controlled with worker pool pattern.
Secure HTTP client with TLS configuration.


Install Golang: Ensure you have Go installed on your machine.
git clone <repo-url>
cd <repo-folder>

go run main.go


How it Works

Initialization:

  A list of e-commerce URLs is provided in the main() function.
  A channel (queue) is used to manage the URLs to be processed.
  A worker pool is created with a fixed number of goroutines (numWorkers).

Worker Goroutines:

  Each worker fetches a URL from the queue channel.
  If the URL hasn't been visited, it is processed by calling ProductUrls.

Processing URLs:

  The ProductUrls function fetches the page content.
  Extracts all links using the extractlinks library.
  Filters and validates the extracted links before adding them back to the queue for further crawling.

Concurrency Management:

  Synchronization is handled using a sync.WaitGroup.
  The map hasVisited prevents duplicate crawling.

Error Handling:
  If a URL is invalid or results in an error, it is skipped, preventing the crawler from crashing.


main(): Initializes the worker pool and seeds the initial URLs.
worker(): Worker function that processes URLs from the queue.
ProductUrls(): Fetches and processes product URLs from a given webpage.
MakeCompleteProductUrl(): Converts relative URLs to absolute URLs.
isValidURL(): Validates if a URL is properly formatted and supported.
checkErr(): Handles errors without terminating the application.

Dependencies
github.com/steelx/extractlinks: Used for extracting links from webpages.


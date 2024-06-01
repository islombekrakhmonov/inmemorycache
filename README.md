# InMemoryCache

## Overview

InMemoryCache is a simple and efficient in-memory key-value store written in Go. It provides basic operations such as setting, getting, and deleting items from the cache. This implementation can be useful for applications that require fast access to frequently used data.

## Features

- **Set**: Store a key-value pair in the cache.
- **Get**: Retrieve the value associated with a key.
- **Delete**: Remove a key-value pair from the cache.
- **Concurrency Safe**: Supports concurrent access with proper synchronization.

## Installation

To use InMemoryCache in your project, you need to have Go installed. You can get the package using:

```sh
go get github.com/islombekrakhmonov/inmemorycache

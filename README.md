# Prometheus Buienradar Exporter
[![GoDoc][badge]][godoc]

Simple exporter for scraping data from Buienradar.

## Installation

Download & install using:

```
go get git.slxh.eu/prometheus/buienradar_exporter
```

## Usage
Run buienradar with the regions you want to scrape, eg:

```
./buienradar_exporter -regions 'Utrecht'
```

Check the [Buienradar website][regions] for a list of all regions.

[badge]: https://godoc.org/git.slxh.eu/prometheus/buienradar_exporter?status.svg
[godoc]: https://godoc.org/git.slxh.eu/prometheus/buienradar_exporter
[regions]: https://json.buienradar.nl

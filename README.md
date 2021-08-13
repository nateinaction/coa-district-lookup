# COA District Lookup

This command line application accepts a CSV file in the format:
```csv
streetnumber,streetname,zipcode
```

It will lookup the district for each address and return the following:
```csv
streetnumber,streetname,zipcode,district
```

## Usage
```bash
coa-district-lookup path/to/csv/file.csv
```

## Requirements
- [Golang](https://golang.org/doc/install)

## How to download
```bash
git clone http://github.com/nateinaction/coa-district-lookup.git
```

## How to build
```bash
go build
```

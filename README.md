# TopN

A simple application to find the top *n* numbers in a given file that has one integer on each line.

## Usage

Generate a file of random numbers:

```shell
shuf -i 1-1000000000 -n 1000000000 -o <filename> # 9.2 GiB
```

Build the application:

```shell
go build ./...
```

Run the application for a given file:

```shell
./go-topn -f <filename>
```

Set the number of top entries to return (default=10):

```shell
./go-topn -f <filename> -n <number>
```
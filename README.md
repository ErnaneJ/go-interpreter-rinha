# Go interpreter for rinha

## Runnind with go CLI

```bash
go run main.go
```

## Running with docker image

Build a docker image from project
```bash
docker build -t go-interpreter-rinha:1.0 .
```
And execute it 

```bash
docker run -dp 3000:3000 go-interpreter-rinha:1.0

# Or

docker run -it --rm go-interpreter-rinha:1.0 # with this instructions you can see outputs from execution
```

## Tokenize source.rinha to source.rinha.json

```bash
# install rust
# cargo install rinha
# run rinha ./var/rinha/files/print.rinha > ./var/rinha/files/print.rinha.json for exemple

# use the json file
```

## Execute

```bash
go run main.go

# or

bin/run file_name # ex: bin/run print for /var/rinha/files/print.rinha

# bin/run file need permissions to run. For this, execute
chmod +x bin/run
```
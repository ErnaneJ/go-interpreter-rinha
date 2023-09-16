# Go Interpreter for Rinha

This project is a Go interpreter for the Rinha programming language.

## Running with the Go CLI

You can run the interpreter using the Go command-line interface (CLI):

```bash
go run main.go
```

## Running with a Docker Image

You can also run the interpreter using a Docker image. First, build a Docker image from the project:

```bash
docker build -t go-interpreter-rinha:1.0 .
```

Then, execute the Docker container:

```bash
docker run -dp 3000:3000 go-interpreter-rinha:1.0
```

Alternatively, if you want to see the outputs from the execution:

```bash
docker run -it --rm go-interpreter-rinha:1.0
```

## Tokenizing Source Files

You can tokenize a Rinha source file into a JSON file using the `rinha` tool. First, make sure you have Rust installed. Then, install the `rinha` tool:

```bash
cargo install rinha
```

After installing `rinha`, you can tokenize a Rinha source file to JSON like this:

```bash
rinha ./var/rinha/files/print.rinha > ./var/rinha/files/print.rinha.json
```

You can then use the generated JSON file as needed.

## Execution

To execute a Rinha program, use the following commands:

```bash
go run main.go file_name
```

or

```bash
bin/run file_name
```

For example:

```bash
bin/run print
```

The `bin/run` script needs execution permissions. You can grant the necessary permissions by running:

```bash
chmod +x bin/run
```

Note: `file_name` represents a file present in the `/var/rinha/files` folder and corresponds to a file without an extension.
# Go Interpreter for Rinha 🚀


Welcome to Go Interpreter for the Rinha programming language!

See more at [Rinha de compiler](https://github.com/aripiprazole/rinha-de-compiler).

## Getting Started

### Running with the Go CLI 🏃‍♂️

To execute the interpreter using the Go command-line interface (CLI), simply enter the following command in your terminal:

```bash
go run main.go
```

### Running with Docker 🐳

Start by building a Docker image from the project:

```bash
docker build -t go-interpreter-rinha:1.0 .
```

Once the image is built, you can run the interpreter within a Docker container.

```bash
docker run -dp 3000:3000 go-interpreter-rinha:1.0
```

If you prefer to see the interpreter's output interactively:

```bash
docker run -it --rm go-interpreter-rinha:1.0
```

**Note:** All executions (CLI or docker) use the content present in `var/rinha/source.rinha.json` to interpret.

## Tokenizing Source Files 📝

To tokenize a Rinha source file into a JSON format, you can utilize the `rinha` tool. First, ensure you have Rust installed. Then, proceed with the installation of the `rinha` tool:

```bash
cargo install rinha
```

Once `rinha` is installed, you can tokenize a Rinha source file to JSON with this command:

```bash
rinha ./var/rinha/<FILE_NAME>.rinha > ./var/rinha/<FILE_NAME>.rinha.json
```

You can subsequently use the generated JSON file as required.

## Execution 🚀

To run a `.rinha` program on your machine, make sure the source code is present in the `/var/rinha/` folder. After that, you can run the command below to generate the JSON and then interpret it by showing the output in your terminal.

```bash
bin/run file_name

#For example:
bin/run hello
```

Ensure that the `bin/run` script has execution permissions. You can grant these permissions by running:

```bash
chmod +x bin/run
```

**Note:** `file_name` should correspond to a file in the `/var/rinha/files` directory and should exclude any file extensions.

## TODO List 📋

Here's a list of tasks that have been completed and those that are yet to be tackled:

- [x] Int
- [x] Str
- [x] Call
- [x] Binary
- [x] Function
- [x] Let
- [x] If
- [x] Print
- [x] First
- [x] Second
- [x] Bool
- [x] Tuple
- [x] Var
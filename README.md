# md2pdf-go-cli

A simple command-line tool to convert Markdown files to PDFs.

## Installation

To install the CLI, you need to have Go installed on your machine. You can then clone the repository and build the executable:

```bash
git clone https://github.com/manishramanandi/md2pdf-go-cli.git
cd md2pdf-go-cli
go build
```

This will create an executable file named `md2pdf-go-cli` in the project directory.

## Usage

To convert a Markdown file to a PDF, use the `convert` command with the `--mdfile` and `--pdffile` flags:

```bash
./md2pdf-go-cli convert --mdfile <path/to/your/markdown/file.md> --pdffile <path/to/your/output/file.pdf>
```

### Options

You can customize the PDF output with the following optional flags:

*   `--fontsize` or `-f`: Set the font size (default: "12").
*   `--fontstyle` or `-s`: Set the font style (default: "B" for bold).
*   `--textcolorR` or `-r`: Set the red component of the text color (0-255) (default: "200").
*   `--textcolorG` or `-g`: Set the green component of the text color (0-255) (default: "200").
*   `--textcolorB` or `-b`: Set the blue component of the text color (0-255) (default: "200").

Example with options:

```bash
./md2pdf-go-cli convert -m test.md -p test.pdf -f 14 -s I
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue on the GitHub repository.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

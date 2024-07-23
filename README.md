# LinguistMap

LinguistMap is a powerful CLI tool that generates JSON mappings of programming languages and their file extensions based on GitHub's Linguist YAML file.

## Features

- Fetches the latest language data from GitHub's Linguist repository
- Falls back to an embedded YAML file if network fetch fails
- Generates simple extension-to-language and language-to-extension mappings by default
- Optionally creates detailed language maps with additional metadata in both directions
- Embedded YAML file ensures the tool always works, even offline

## Installation

1. Ensure you have Go installed on your system.
2. Clone this repository:
   ```
   git clone https://github.com/gusanmaz/LinguistMap.git
   cd LinguistMap
   ```
3. Build the program:
   ```
   go build -o linguistmap main.go
   ```

## Usage

### Basic Usage

To generate simple mapping files:

```
./linguistmap
```

This will create two JSON files:
- `extension_to_language.json`: Maps file extensions to languages
- `language_to_extension.json`: Maps languages to file extensions

### Detailed Mapping

To generate detailed language maps with additional metadata:

```
./linguistmap -d
```
or
```
./linguistmap --detailed
```

This will create two JSON files:
- `detailed_extension_to_language.json`: Maps file extensions to detailed language information
- `detailed_language_to_extension.json`: Maps languages to their detailed information, including extensions and filenames

### Command-line Options

```
Usage of ./linguistmap:
  -d, --detailed    Generate detailed language maps
```

## Output Files

1. `extension_to_language.json`: Maps extensions to languages (simple mode)
2. `language_to_extension.json`: Maps languages to extensions (simple mode)
3. `detailed_extension_to_language.json`: Maps extensions to detailed language information (detailed mode)
4. `detailed_language_to_extension.json`: Maps languages to their detailed information (detailed mode)

## Fallback Mechanism

If the program fails to fetch the YAML file from GitHub, it will automatically use the embedded YAML file included in the binary. This ensures the tool always works, even without an internet connection.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Author

Created and maintained by [gusanmaz](https://github.com/gusanmaz).

## Links

- GitHub Repository: [https://github.com/gusanmaz/LinguistMap](https://github.com/gusanmaz/LinguistMap)
- Issues: [https://github.com/gusanmaz/LinguistMap/issues](https://github.com/gusanmaz/LinguistMap/issues)

## Acknowledgments

This tool is based on the language data provided by [GitHub Linguist](https://github.com/github/linguist).git 
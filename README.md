# go-text-Modifier
## Overview

Text Modifier is a command-line text processing tool built in Go that performs automated text corrections and transformations. The program reads an input file, applies a series of predefined modifications based on special markers, and outputs the corrected text to a specified file.

This project demonstrates practical application of string manipulation, file I/O operations, and building a functional CLI tool that solves real-world text editing problems.

## Features

The tool performs the following text transformations:

- **Number Base Conversion**
  - `(hex)` - Converts hexadecimal numbers to decimal
  - `(bin)` - Converts binary numbers to decimal

- **Case Modifications**
  - `(up)` - Converts words to UPPERCASE
  - `(low)` - Converts words to lowercase
  - `(cap)` - Capitalizes words
  - Supports numeric parameters: `(up, 2)` applies transformation to the previous 2 words

- **Punctuation Formatting**
  - Automatically spaces punctuation marks (`.`, `,`, `!`, `?`, `:`, `;`) correctly
  - Handles punctuation groups (`...`, `!?`) without extra spacing
  - Formats single quotes `'` to wrap words without spaces

- **Article Correction**
  - Automatically converts `a` to `an` before words starting with vowels or 'h'

## Usage

```bash
go run . <input_file> <output_file>
```

### Example

```bash
# Input file (sample.txt)
it (cap) was the best of times, it was the worst of times (up)

# Run the program
go run . sample.txt result.txt

# Output file (result.txt)
It was the best of times, it was the worst of TIMES
```

## What I Learned

Through this project, I developed proficiency in:

- **File System Operations**: Working with Go's `fs` API to read from and write to files
- **String Manipulation**: Parsing and transforming text using Go's string manipulation functions
- **Number System Conversions**: Converting between binary, hexadecimal, and decimal number systems
- **Regular Expressions & Pattern Matching**: Identifying and processing special markers within text
- **Code Organization**: Structuring a Go project with reusable functions and clean separation of concerns

## Technologies

- **Language**: Go
- **Concepts**: File I/O, String Processing, CLI Applications

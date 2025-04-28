
- `main.go`: The main Go program that generates a ZIP file containing a CSV file.
- `example.csv`: A sample CSV file with data (used as a reference).
- `example.zip`: The output ZIP file created by the program.
- `go.mod`: The Go module file.

## Prerequisites

- Go 1.24.1 or later installed on your system.

## How It Works

1. The program creates a ZIP file named `example.zip`.
2. Inside the ZIP file, it creates a CSV file named `example.csv`.
3. The CSV file contains the following data:

   | Name    | Age | Email              |
   |---------|-----|--------------------|
   | Alice   | 30  | alice@example.com  |
   | Bob     | 25  | bob@example.com    |
   | Charlie | 35  | charlie@example.com|

4. The program writes the header and data rows to the CSV file and flushes the data to the ZIP file.

## How to Run

1. Clone this repository or copy the files into a directory.
2. Run the following command to execute the program:

   ```bash
   go run main.go

3. After running the program, you will find a file named example.zip in the same directory. Extract it to see the example.csv file.

Error Handling
The program includes error handling for the following scenarios:

Failure to create the ZIP file.
Failure to create the CSV file inside the ZIP.
Failure to write the header or data rows to the CSV file.
Failure to flush the CSV writer.
Output
Upon successful execution, the program logs the following message: สร้าง ZIP ที่มี CSV เรียบร้อยแล้ว!

This means "ZIP containing CSV created successfully!" in Thai.
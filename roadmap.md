# Step by step of every new features

## Parsing the Excel file (.xlsx)

The purpose is to use a library (written in Go), like ["Excelize"](https://github.com/360EntSecGroup-Skylar/excelize), then Read the Excel file and create a slice out of it:

    data [][]string

### installing the library

go get github.com/360EntSecGroup-Skylar/excelize

### Reading the Excel file

    package main
   
    import (
      "fmt"
      "github.com/360EntSecGroup-Skylar/excelize"
    )
   
    func main() {
      xlsx, err := excelize.OpenFile("./filenames.xlsx")
      if err != nil {
        fmt.Println(err)
        return
      }
      // Get all the rows in the Sheet1.
      rows, err := xlsx.GetRows("Sheet1")
      if err != nil {
        fmt.Println(err)
        return
      }
      // skip the first line (header)
      for _, row := range rows[1:] {
        for _, colCell := range row {
          fmt.Print(colCell, "\t")
        }
        fmt.Println()
      }
    }
   
### Store the Excel file data into a slice

    var data [][2]string
    // skip the first line (header)
    for _, row := range rows[1:] {
      line := [2]string{"",""}
      for i, colCell := range row {
        line[i] = colCell
      }
      // skip the last empty line
      if line[0] != "" && line[1] != "" {
        data = append(data, line)
      }
    }
    fmt.Println(data)

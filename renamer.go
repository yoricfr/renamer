// This  program runs from the command line and will rename all the files in the search_folder, following the renaming rules in the Excel file sitting in the root directory (filenames.xlsx)
// Author: Yoric Mangeart
// Date: 2020/03/31
// Revision: 1.0 original release
//           1.1 Print out all file names with no renaming rules
//           1.2 Print out all Excel names not found in search_folder

package main

import (
	"fmt"
  "os"
  "strings"
  "path/filepath"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {

  // Open the Excel file containing the renaming rules
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

  // where we store the data from the Excel file
  var data [][3]string

  // skip the first line (header)
  for _, row := range rows[1:] {
    line := [3]string{"","",""}
		for i, colCell := range row {
      line[i] = colCell
		}
    // skip the last empry line
    if line[0] != "" && line[1] != "" {
      data = append(data, line)
    }
	}

  filesRenamed := 0

  // browse every file from "search_folder"
  err = filepath.Walk("search_folder", func(path string, info os.FileInfo, err error) error {
    // skip folders and file name starting with dot
    if info.Mode().IsDir() || strings.HasPrefix(info.Name(), ".") {
      return nil
    }
    // check if there's a renaming rule for this specific file
    for j,l := range data {
      if l[0] == info.Name() {
        err = os.Rename(path, filepath.Join(filepath.Dir(path),l[1]))
          if err != nil {
            fmt.Printf("error renaming file: ", err)
            return nil
        }
        filesRenamed +=1
        // we mark this line as "used" for post-processing the listing
        data[j][2]="u"
        break
      }
    }
    return nil
  })
  if err != nil {
		fmt.Printf("error walking the path %q: %v\n", "search_folder", err)
		return
	}
  for _,l := range data {
    if l[2]=="" {
      fmt.Println("Excel name not found:", l[0])
    }
  }
  fmt.Println("Done:", filesRenamed, "file(s) renamed.")
}

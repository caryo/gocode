package main

import (
	"fmt"
	"os"
)

func main() {
   envStr := os.Environ()

   fmt.Printf("len(envStr) = %d\n", len(envStr))
   for i := 0; i < len(envStr); i++ {
      fmt.Printf("%s\n", envStr[i])
   }
   pathStr := os.ExpandEnv("${PATH}")
   fmt.Printf("len(pathStr) = %d\n", len(pathStr))
   fmt.Printf("%s\n", pathStr)
   userStr := os.Getenv("USERNAME")
   fmt.Printf("len(userStr) = %d\n", len(userStr))
   fmt.Printf("%s\n", userStr)
   hostStr,_ := os.Hostname()
   fmt.Printf("len(hostStr) = %d\n", len(hostStr))
   fmt.Printf("%s\n", hostStr)
   tempStr := os.TempDir()
   fmt.Printf("len(tempStr) = %d\n", len(tempStr))
   fmt.Printf("%s\n", tempStr)

   tempDir,tempDirErr := os.Open(tempStr)
   if (tempDirErr != nil) {
      fmt.Printf("tempDirErr = %s\n", tempDirErr)
   }
   if (tempDir != nil) {
      tempFiles,tempFilesErr := tempDir.Readdirnames(0)
      if (tempFilesErr != nil) {
         fmt.Printf("tempDirErr = %s\n", tempDirErr)
      }
      fmt.Printf("len(tempFiles) = %d\n", len(tempFiles))
      for i := 0; i < len(tempFiles); i++ {
         fmt.Printf("%s\n", tempFiles[i])
      }
   }

   fileStr := tempStr + "\\file1.txt"
   fmt.Printf("fileStr = %s\n", fileStr)
   file1,fileErr := os.Create(fileStr)
   fmt.Printf("file1 = %v\n", file1)
   if (fileErr != nil) {
      fmt.Printf("fileErr = %s\n", fileErr)
   }
   if (file1 != nil) {
      fileInfo,fileInfoErr := file1.Stat()
      if (fileInfoErr != nil) {
         fmt.Printf("fileInfoErr = %s\n", fileInfoErr)
      }
      if (fileInfo != nil) {
         fmt.Printf("fileInfo.Name = %s\n", fileInfo.Name())
         fmt.Printf("fileInfo.Size = %d\n", fileInfo.Size())
         fmt.Printf("fileInfo.Mode = %s\n", fileInfo.Mode())
         fmt.Printf("fileInfo.IsDir = %b\n", fileInfo.IsDir())
      }
   }
}

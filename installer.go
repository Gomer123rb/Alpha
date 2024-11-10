package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
    alphaURL   = "https://example.com/path/to/Alpha_installer.exe"  // Replace with the actual URL
    outputDir  = "C:\\AlphaDownloader"
    outputFile = "Alpha_installer.exe"
)

func main() {
    err := os.MkdirAll(outputDir, os.ModePerm)
    if err != nil {
        fmt.Println("Error creating output directory:", err)
        return
    }

    outputPath := filepath.Join(outputDir, outputFile)
    err = downloadFile(outputPath, alphaURL)
    if err != nil {
        fmt.Println("Error downloading file:", err)
        return
    }

    fmt.Println("Download completed:", outputPath)

    err = runInstaller(outputPath)
    if err != nil {
        fmt.Println("Error running installer:", err)
        return
    }

    fmt.Println("Alpha installed successfully!")
}

func downloadFile(filepath string, url string) error {
    response, err := http.Get(url)
    if err != nil {
        return err
    }
    defer response.Body.Close()

    out, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, response.Body)
    return err
}

func runInstaller(filepath string) error {
    cmd := exec.Command(filepath)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}

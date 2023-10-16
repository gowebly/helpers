package gowebly

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsExistInFolder(t *testing.T) {
	// Test case 1: File exists in current folder.
	fileName := "testfile.txt"
	filePath := filepath.Join(".", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	file.Close()
	defer os.Remove(filePath)

	exists := isExistInFolder(fileName, false)
	if !exists {
		t.Errorf("Expected file to exist, but it does not.")
	}

	// Test case 2: Folder exists in current folder.
	folderName := "testfolder"
	folderPath := filepath.Join(".", folderName)
	err = os.Mkdir(folderPath, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create test folder: %v", err)
	}
	defer os.Remove(folderPath)

	exists = isExistInFolder(folderName, true)
	if !exists {
		t.Errorf("Expected folder to exist, but it does not.")
	}

	// Test case 3: File does not exist in current folder.
	notExistFileName := "nonexistent.txt"
	exists = isExistInFolder(notExistFileName, false)
	if exists {
		t.Errorf("Expected file not to exist, but it does.")
	}

	// Test case 4: Folder does not exist in current folder.
	notExistFolderName := "nonexistent"
	exists = isExistInFolder(notExistFolderName, true)
	if exists {
		t.Errorf("Expected folder not to exist, but it does.")
	}
}

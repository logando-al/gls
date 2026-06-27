package scanner

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

// FileInfo contains information about a file
type FileInfo struct {
	Name     string
	Path     string
	Size     int64
	Mode     fs.FileMode
	ModTime  time.Time
	IsDir    bool
	IsLink   bool
	LinkPath string
}

// Options configures the scanner behavior
type Options struct {
	ShowHidden bool
	Recursive  bool
	All        bool // Show current and parent directories with -a
}

// Scan scans a directory and returns file information
func Scan(dir string, options Options) ([]FileInfo, error) {
	var entries []FileInfo
	
	baseInfo, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	
	if !baseInfo.IsDir() {
		entry, err := createFileInfo(dir, baseInfo)
		if err != nil {
			return nil, err
		}
		return []FileInfo{entry}, nil
	}
	
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	
	for _, file := range files {
		// Skip hidden files unless ShowHidden is true
		if !options.ShowHidden && isHidden(file.Name()) {
			continue
		}
		
		// Skip . and .. unless All is true
		if !options.All && (file.Name() == "." || file.Name() == "..") {
			continue
		}
		
		absPath := filepath.Join(dir, file.Name())
		info, err := file.Info()
		if err != nil {
			continue
		}
		
		entry, err := createFileInfo(absPath, info)
		if err != nil {
			continue
		}
		
		entries = append(entries, entry)
		
		// Recursive scanning if needed
		if options.Recursive && file.IsDir() && file.Name() != "." && file.Name() != ".." {
			subEntries, err := Scan(absPath, options)
			if err == nil {
				entries = append(entries, subEntries...)
			}
		}
	}
	
	return entries, nil
}

// createFileInfo creates a FileInfo from an fs.FileInfo
func createFileInfo(path string, info fs.FileInfo) (FileInfo, error) {
	entry := FileInfo{
		Name:    info.Name(),
		Path:    path,
		Size:    info.Size(),
		Mode:    info.Mode(),
		ModTime: info.ModTime(),
		IsDir:   info.IsDir(),
	}
	
	// Check if it's a symbolic link
	if entry.Mode&os.ModeSymlink != 0 {
		entry.IsLink = true
		if linkPath, err := os.Readlink(path); err == nil {
			entry.LinkPath = linkPath
		}
	}
	
	return entry, nil
}

// isHidden checks if a file is hidden (starts with a dot on Unix-like systems)
func isHidden(filename string) bool {
	return len(filename) > 0 && filename[0] == '.'
}
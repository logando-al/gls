package output

import (
	"fmt"
	"github.com/logando-al/gls/pkg/colors"
	"github.com/logando-al/gls/pkg/scanner"
	"path/filepath"
)

// Format defines the type of output format
type Format int

const (
	FormatGrid Format = iota
	FormatList
	FormatTree
)

// Options for output formatting
type Options struct {
	Format      Format
	ShowDetails bool // Show file details like permissions, size, etc.
	ShowIcons   bool
	ShowGit     bool
	WithColors  bool
	HumanSize   bool
	ShowHidden  bool
	ShowAll     bool   // Show . and .. directory entries when true
	SortBy      string // Name, Size, Time, etc.
	Reverse     bool
	GroupDirs   bool // Group directories first or last
	TreeDepth   int  // For tree format, -1 for unlimited
}

// PrintFiles formats and prints files according to the specified options
func PrintFiles(files []scanner.FileInfo, options Options) error {
	switch options.Format {
	case FormatGrid:
		return printGrid(files, options)
	case FormatList:
		return printList(files, options)
	case FormatTree:
		return printTree(files, options)
	default:
		return printGrid(files, options)
	}
}

// printGrid prints files in a grid format (default ls behavior)
func printGrid(files []scanner.FileInfo, options Options) error {
	if !options.ShowDetails {
		// Simple grid view, just print names
		for _, file := range files {
			fmt.Printf("%s", formatFileName(file, options))

			// Add slash for directories
			if file.IsDir {
				fmt.Printf("/")
			}

			// Add @ for symlinks
			if file.IsLink {
				fmt.Printf("@")
			}

			fmt.Println()
		}
		return nil
	}

	// Grid with details (one per line)
	printDetails(files, options)
	return nil
}

// printList prints files in a list format
func printList(files []scanner.FileInfo, options Options) error {
	if options.ShowDetails {
		printDetails(files, options)
	} else {
		// Simple list view
		for _, file := range files {
			fmt.Print(formatFileName(file, options))

			// Add slash for directories
			if file.IsDir {
				fmt.Print("/")
			}

			fmt.Println()
		}
	}
	return nil
}

// printTree prints files in a tree format
func printTree(files []scanner.FileInfo, options Options) error {
	rootDir := "."
	// If there are files passed in, try to determine if they belong to the same directory
	if len(files) > 0 {
		rootDir = filepath.Dir(files[0].Path)

		// Check if all files are in the same directory
		sameDir := true
		for _, file := range files {
			if filepath.Dir(file.Path) != rootDir {
				sameDir = false
				break
			}
		}

		if sameDir {
			// All files are in the same directory, print them as a tree
			printDirTree(rootDir, 0, "", options)
			return nil
		}
	}

	// If we get here, the files passed in are multiple directories or different paths
	for _, file := range files {
		if file.IsDir {
			fmt.Print(formatFileName(file, options) + "/")
			fmt.Println()
			printDirTree(file.Path, 1, "", options)
		} else {
			fmt.Println(formatFileName(file, options))
		}
	}

	return nil
}

// printDirTree recursively prints a directory tree
func printDirTree(dir string, depth int, prefix string, options Options) error {
	// Stop if reached max depth
	if options.TreeDepth >= 0 && depth > options.TreeDepth {
		return nil
	}

	// Scan the current directory
	scanOptions := scanner.Options{
		ShowHidden: options.ShowHidden,
		Recursive:  false,
	}

	// Use ShowAll if available, otherwise default to false
	if options.ShowAll {
		scanOptions.All = true
	}

	files, err := scanner.Scan(dir, scanOptions)
	if err != nil {
		return err
	}

	// Sort files
	files = sortFiles(files, options.SortBy, options.Reverse, options.GroupDirs)

	for i, file := range files {
		isLast := i == len(files)-1

		if depth == 0 {
			// At root level, just print the name
			fmt.Println(formatFileName(file, options))
			if file.IsDir {
				// For directories at root level, add children
				if isLast {
					printDirTree(filepath.Join(dir, file.Name), depth+1, "    ", options)
				} else {
					printDirTree(filepath.Join(dir, file.Name), depth+1, "│   ", options)
				}
			}
		} else {
			// At subdirectory level
			var currentPrefix, newPrefix string
			if isLast {
				currentPrefix = prefix + "└── "
				newPrefix = prefix + "    "
			} else {
				currentPrefix = prefix + "├── "
				newPrefix = prefix + "│   "
			}

			fmt.Printf("%s%s", currentPrefix, formatFileName(file, options))

			// Add directory indicator
			if file.IsDir {
				fmt.Print("/")
			}

			// Add link indicator
			if file.IsLink {
				fmt.Printf(" -> %s", file.LinkPath)
			}

			fmt.Println()

			// If it's a directory, print its contents
			if file.IsDir && !file.IsLink {
				printDirTree(filepath.Join(dir, file.Name), depth+1, newPrefix, options)
			}
		}
	}

	return nil
}

// printDetails prints file details (permissions, size, timestamp)
func printDetails(files []scanner.FileInfo, options Options) error {
	// Print table header if details are requested
	if options.ShowDetails {
		fmt.Print("Permissions")
		fmt.Print("\t")
		fmt.Print("Size")
		fmt.Print("\t")
		fmt.Print("Modified")
		fmt.Print("\t")
		fmt.Print("Name")
		fmt.Println()
	}

	for _, file := range files {
		// Print permissions
		fmt.Print(file.Mode.String())
		fmt.Print("\t")

		// Print size
		if options.HumanSize {
			fmt.Print(humanSize(file.Size))
		} else {
			fmt.Print(file.Size)
		}
		fmt.Print("\t")

		// Print modification time
		fmt.Print(file.ModTime.Format("Jan 02 15:04"))
		fmt.Print("\t")

		// Print name with formatting
		fmt.Print(formatFileName(file, options))

		// Add slash for directories
		if file.IsDir {
			fmt.Print("/")
		}

		// Add @ for symlinks with link destination
		if file.IsLink {
			if file.LinkPath != "" {
				fmt.Printf(" -> %s", file.LinkPath)
			} else {
				fmt.Printf("@")
			}
		}

		fmt.Println()
	}
	return nil
}

// formatFileName formats a filename according to the options
var colorManager *colors.ColorManager

func formatFileName(file scanner.FileInfo, options Options) string {
	if options.WithColors {
		if colorManager == nil {
			colorManager = colors.NewColorManager()
			colorManager.SetupExtensionColors()
		}
		return colorManager.FormatFileName(file)
	}
	return file.Name
}

// humanSize formats a size in bytes as a human-readable size
func humanSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d", size)
	}

	units := []string{"", "K", "M", "G", "T"}
	for i, unit := range units {
		if size < 1024 || (size == 1024 && len(units) == i+1) {
			return fmt.Sprintf("%.1f%s", float64(size)/1024, unit)
		}
		size = size / 1024
	}

	return fmt.Sprintf("%dT", size)
}

// sortFiles sorts files according to the specified criteria
func sortFiles(files []scanner.FileInfo, sortBy string, reverse bool, groupDirs bool) []scanner.FileInfo {
	result := make([]scanner.FileInfo, len(files))
	copy(result, files)

	// Case 1: Group directories first or last
	if groupDirs {
		var dirs, nonDirs []scanner.FileInfo
		for _, file := range files {
			if file.IsDir {
				dirs = append(dirs, file)
			} else {
				nonDirs = append(nonDirs, file)
			}
		}

		// If groupDirs is true, directories come first
		result = append(dirs, nonDirs...)
	}

	// Now sort according to the specified criteria
	switch sortBy {
	case "name":
		sortSlice(result, func(a, b scanner.FileInfo) bool {
			return a.Name < b.Name
		}, reverse)
	case "size":
		sortSlice(result, func(a, b scanner.FileInfo) bool {
			return a.Size < b.Size
		}, reverse)
	case "time":
		sortSlice(result, func(a, b scanner.FileInfo) bool {
			return a.ModTime.Before(b.ModTime)
		}, reverse)
	case "extension":
		sortSlice(result, func(a, b scanner.FileInfo) bool {
			extA := filepath.Ext(a.Name)
			extB := filepath.Ext(b.Name)
			if extA != extB {
				return extA < extB
			}
			return a.Name < b.Name // If extensions are same, sort by name
		}, reverse)
	default:
		// Default to sorting by name
		sortSlice(result, func(a, b scanner.FileInfo) bool {
			return a.Name < b.Name
		}, reverse)
	}

	return result
}

// sortSlice sorts a slice of FileInfo using the provided comparison function
func sortSlice(files []scanner.FileInfo, compare func(a, b scanner.FileInfo) bool, reverse bool) {
	n := len(files)
	if n <= 1 {
		return
	}

	// Simple bubble sort
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			shouldSwap := compare(files[j], files[j+1])
			if reverse {
				shouldSwap = !shouldSwap
			}
			if shouldSwap {
				files[j], files[j+1] = files[j+1], files[j]
			}
		}
	}
}

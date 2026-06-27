package output

import (
	"fmt"
	"github.com/logando-al/go-ls/pkg/scanner"
	"github.com/logando-al/go-ls/pkg/colors"
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
	Format       Format
	ShowDetails  bool // Show file details like permissions, size, etc.
	ShowIcons    bool
	ShowGit      bool
	WithColors   bool
	HumanSize    bool
	ShowHidden   bool
	SortBy       string // Name, Size, Time, etc.
	Reverse      bool
	GroupDirs    bool // Group directories first or last
	TreeDepth    int  // For tree format, -1 for unlimited
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
			fmt.Printf(formatFileName(file, options))
			
			// Add slash for directories
			if file.IsDir {
				fmt.Printf("/")
			}
			
			fmt.Println()
		}
	}
	return nil
}

// printTree prints files in a tree format
func printTree(files []scanner.FileInfo, options Options) error {
	fmt.Println("Tree format not implemented yet")
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
		fmt.Printf(formatFileName(file, options))
		
		// Add slash for directories
		if file.IsDir {
			fmt.Printf("/")
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
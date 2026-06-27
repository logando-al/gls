package main

import (
	"fmt"
	"os"
	"errors"

	"github.com/spf13/cobra"
	"github.com/logando-al/go-ls/pkg/scanner"
	"github.com/logando-al/go-ls/pkg/output"
)

var (
	showHidden      bool
	showDetails     bool
	showAll         bool
	recursive       bool
	treeView        bool
	showIcons       bool
	showGit         bool
	withColors      bool
	humanSize       bool
	sortBy          string
	reverse         bool
	groupDirs       bool
	treeDepth       int
	color           string
)

var rootCmd = &cobra.Command{
	Use:   "go-ls [directory]",
	Short: "A modern Go-based alternative to ls",
	Long: `go-ls is a modern alternative to the traditional 'ls' command, 
written in Go. It provides enhanced features like color-coding, file 
metadata display, tree view, and Git integration.

If no directory is specified, the current directory is used.`,
	RunE: runLs,
}

func init() {
	// Basic options
	rootCmd.Flags().BoolVarP(&showDetails, "long", "l", false, "Show file details (permissions, size, timestamp)")
	rootCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show hidden files (files starting with .)")
	rootCmd.Flags().BoolVar(&showIcons, "icons", false, "Show icons for file types")
	rootCmd.Flags().BoolVar(&withColors, "colors", false, "Enable color output")
	
	// Sorting options
	rootCmd.Flags().StringVar(&sortBy, "sort", "name", "Sort files by: name, size, time")
	rootCmd.Flags().BoolVar(&reverse, "reverse", false, "Reverse sort order")
	rootCmd.Flags().BoolVar(&groupDirs, "group-dirs", false, "Group directories first")
	
	// Display options
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "R", false, "Recursively list subdirectories")
	rootCmd.Flags().BoolVarP(&treeView, "tree", "T", false, "Show files in a tree view")
	rootCmd.Flags().BoolVar(&humanSize, "human-readable", false, "Show human-readable file sizes")
	
	// Filtering options
	rootCmd.Flags().BoolVar(&showHidden, "hidden", false, "Show only hidden files")
	rootCmd.Flags().BoolVar(&showGit, "git", false, "Show git status")
	
	// Color options
	rootCmd.Flags().StringVar(&color, "color", "",
		"When to use color. 'always', 'never', or 'auto'. Default is 'auto'")
}

func runLs(cmd *cobra.Command, args []string) error {
	// Determine the directory to list
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}
	
	// Determine if we should show hidden files
	showHidden = showAll || showHidden
	
	// Determine color usage
	if color != "" {
		switch color {
		case "always":
			withColors = true
		case "never":
			withColors = false
		case "auto":
			withColors = true // Let colors package decide
		default:
			return errors.New("invalid color option, must be 'always', 'never', or 'auto'")
		}
	} else {
		// Default to auto if not specified
		withColors = true
	}
	
	// Set up scanner options
	scanOptions := scanner.Options{
		ShowHidden: showHidden,
		Recursive:  recursive,
		All:        showAll,
	}
	
	// Scan the directory
	files, err := scanner.Scan(dir, scanOptions)
	if err != nil {
		return fmt.Errorf("error scanning directory: %w", err)
	}
	
	// Set up output options
	outputOpt := output.Options{}
	
	// Set display format
	if treeView {
		outputOpt.Format = output.FormatTree
	} else if showDetails {
		outputOpt.Format = output.FormatList
	} else {
		outputOpt.Format = output.FormatGrid
	}
	
	outputOpt.ShowDetails = showDetails
	outputOpt.ShowIcons = showIcons
	outputOpt.ShowGit = showGit
	outputOpt.WithColors = withColors
	outputOpt.HumanSize = humanSize
	outputOpt.SortBy = sortBy
	outputOpt.Reverse = reverse
	outputOpt.GroupDirs = groupDirs
	outputOpt.TreeDepth = treeDepth
	
	// Print the files
	return output.PrintFiles(files, outputOpt)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
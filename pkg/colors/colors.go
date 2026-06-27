package colors

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	
	"github.com/fatih/color"
	"github.com/logando-al/gls/pkg/scanner"
)

// ColorManager manages color-coding for files
type ColorManager struct {
	directoryColor    *color.Color
	symlinkColor      *color.Color
	executableColor   *color.Color
	regularFileColor  *color.Color
	hiddenFileColor   *color.Color
	extensionColors   map[string]*color.Color
	colorEnabled      bool
}

// String constants for known file extensions and their associated colors
const (
	// Code files extension
	packageGo     = ".go"
	packageJs     = ".js"
	packageJsx    = ".jsx"
	packageTs     = ".ts"
	packageTsx    = ".tsx"
	packageRs     = ".rs"
	packageRb     = ".rb"
	packagePy     = ".py"
	packageLua    = ".lua"
	packageJava   = ".java"
	packageScala  = ".scala"
	packageClj    = ".clj"
	packageCljc   = ".cljc"
	packageClojure = ".clojure"
	packageC      = ".c"
	packageCpp    = ".cpp"
	packageCc     = ".cc"
	packageCxx    = ".cxx"
	packageH      = ".h"
	packageHpp    = ".hpp"
	packageHxx    = ".hxx"
	packageCs     = ".cs"
	packagePhp    = ".php"
	packagePerl   = ".pl"
	packageRpm    = ".rpm"
	packageDocker = ".Dockerfile"
	packageRsSpec = "Podspec"
	packageRsBlock = "rb"
	
	// Config files extension
	packageYaml  = ".yaml"
	packageYml   = ".yml"
	packageToml  = ".toml"
	packageJson  = ".json"
	packageXml   = ".xml"
	packageIni   = ".ini"
	packageTml   = ".tml"
	packageConf  = ".conf"
	packageGit   = ".gitignore"
	
	// Shell/Batch files extension
	packageSh    = ".sh"
	packageBash  = ".bash"
	packageBatch = ".bat"
	packageCmd   = ".cmd"
	packageZsh   = ".zsh"
	packageFish  = ".fish"
	
	// Archive files extension
	packageZip   = ".zip"
	packageTar   = ".tar"
	packageGz    = ".gz"
	packageBz2   = ".bz2"
	packageXz    = ".xz"
	package7z    = ".7z"
	packageRar   = ".rar"
	packageApk   = ".apk"
	packageDeb   = ".deb"
	packageRpm2  = ".rpm"
	
	// Media files extension
	packageMov   = ".mov"
	packageMp4   = ".mp4"
	packageAvi   = ".avi"
	packageMkv   = ".mkv"
	packageFlv   = ".flv"
	packageWmv   = ".wmv"
	package3gp   = ".3gp"
	packageMpeg  = ".mpeg"
	packageMp3   = ".mp3"
	packageFlac  = ".flac"
	packageOgg   = ".ogg"
	packageJpg   = ".jpg"
	packageJpeg  = ".jpeg"
	packageSsh   = ".ssh"
	packagePng   = ".png"
	packageGitInt= ".git"
	packageGif   = ".gif"
	packageWebp  = ".webp"
	packageBmp   = ".bmp"
	packageTiff  = ".tiff"
	packageTif   = ".tif"
	packagePng2  = ".png"
	packageSvg   = ".svg"
	
	// Office files extension
	packagePdf   = ".pdf"
	packageDoc   = ".doc"
	packageDocx  = ".docx"
	packageOdt   = ".odt"
	packageXls   = ".xls"
	packageXlsx  = ".xlsx"
	packageOds   = ".ods"
	packagePpt   = ".ppt"
	packagePptx  = ".pptx"
	packageOdp   = ".odp"
)

// NewColorManager creates a new ColorManager with default colors
func NewColorManager() *ColorManager {
	return &ColorManager{
		directoryColor:    color.New(color.BgBlue, color.FgWhite).Add(color.Bold),   // Blue bg, white fg, bold for directories
		symlinkColor:      color.New(color.FgCyan, color.Bold),                       // Cyan bold for symlinks
		executableColor:   color.New(color.FgGreen, color.Bold),                      // Green bold for executables
		regularFileColor:  color.New(color.FgWhite),                                  // White for regular files
		hiddenFileColor:   color.New(color.FgHiBlack, color.Italic),                  // Dim gray for hidden files
		extensionColors:   make(map[string]*color.Color),
		colorEnabled:      canUseColor(),
	}
}

// SetupExtensionColors sets up colors for different file extensions
func (cm *ColorManager) SetupExtensionColors() {
	// Code files
	cm.extensionColors[packageGo] = color.New(color.FgCyan)
	cm.extensionColors[packageJs] = color.New(color.FgYellow)
	cm.extensionColors[packageJsx] = color.New(color.FgHiYellow)
	cm.extensionColors[packageTs] = color.New(color.FgHiCyan)
	cm.extensionColors[packageTsx] = color.New(color.FgHiGreen)
	cm.extensionColors[packageRs] = color.New(color.FgRed)
	cm.extensionColors[packageRb] = color.New(color.FgHiRed)
	cm.extensionColors[packagePy] = color.New(color.FgHiBlue)
	cm.extensionColors[packageLua] = color.New(color.FgHiMagenta)
	cm.extensionColors[packageJava] = color.New(color.FgHiCyan)
	cm.extensionColors[packageScala] = color.New(color.FgRed)
	cm.extensionColors[packageC] = color.New(color.FgHiBlue)
	cm.extensionColors[packageCpp] = color.New(color.FgHiYellow)
	cm.extensionColors[packageCs] = color.New(color.FgCyan)
	cm.extensionColors[packagePhp] = color.New(color.FgMagenta)
	cm.extensionColors[packagePerl] = color.New(color.FgHiGreen)
	
	// Config files
	cm.extensionColors[packageYaml] = color.New(color.FgCyan)
	cm.extensionColors[packageYml] = color.New(color.FgCyan)
	cm.extensionColors[packageToml] = color.New(color.FgHiCyan)
	
	// Archive files
	cm.extensionColors[packageZip] = color.New(color.FgRed)
	cm.extensionColors[packageTar] = color.New(color.FgHiRed)
	cm.extensionColors[packageGz] = color.New(color.FgRed)
	
	// Media files
	cm.extensionColors[packageMov] = color.New(color.FgMagenta)
	cm.extensionColors[packageMp4] = color.New(color.FgHiMagenta)
	cm.extensionColors[packageJpg] = color.New(color.FgYellow)
	cm.extensionColors[packageJpeg] = color.New(color.FgHiYellow)
	cm.extensionColors[packagePng] = color.New(color.FgYellow)
	cm.extensionColors[packageGif] = color.New(color.FgHiYellow)
	cm.extensionColors[packageMp3] = color.New(color.FgCyan)
	
	// Office files
	cm.extensionColors[packagePdf] = color.New(color.FgRed)
	cm.extensionColors[packageDoc] = color.New(color.FgHiBlue)
	cm.extensionColors[packageDocx] = color.New(color.FgHiBlue)
	cm.extensionColors[packageXls] = color.New(color.FgHiGreen)
	cm.extensionColors[packageXlsx] = color.New(color.FgHiGreen)
}

// FormatFileName formats a file name with the appropriate color
func (cm *ColorManager) FormatFileName(file scanner.FileInfo) string {
	if !cm.colorEnabled {
		return file.Name
	}
	
	// Check if it's a directory
	if file.IsDir {
		return cm.directoryColor.Sprintf("%s", file.Name)
	}
	
	// Check if it's a symbolic link
	if file.IsLink {
		fname := cm.symlinkColor.Sprintf("%s", file.Name)
		if file.LinkPath != "" {
			// Add arrow and colored link destination
			return fmt.Sprintf("%s%s%s", fname, 
				color.WhiteString(" -> "),
				cm.FormatFileNameByName(file.LinkPath))
		}
		return fname
	}
	
	// Check if it's hidden
	if isHiddenFile(file.Name) {
		return cm.hiddenFileColor.Sprintf("%s", file.Name)
	}
	
	// Check if it's executable
	if file.Mode&0111 != 0 {
		return cm.executableColor.Sprintf("%s", file.Name)
	}
	
	// Check for extension-specific coloring
	ext := filepath.Ext(file.Name)
	if color, ok := cm.extensionColors[ext]; ok {
		return color.Sprintf("%s", file.Name)
	}
	
	// Default color for regular files
	return cm.regularFileColor.Sprintf("%s", file.Name)
}

// FormatFileNameByName formats a file name with the appropriate color by name only
func (cm *ColorManager) FormatFileNameByName(filename string) string {
	if !cm.colorEnabled {
		return filename
	}
	
	// Check for extension-specific coloring
	ext := filepath.Ext(filename)
	if color, ok := cm.extensionColors[ext]; ok {
		return color.Sprintf("%s", filename)
	}
	
	// Default color
	return cm.regularFileColor.Sprintf("%s", filename)
}

// isHiddenFile checks if a file is hidden
func isHiddenFile(filename string) bool {
	return len(filename) > 0 && filename[0] == '.'
}

// canUseColor determines if we can use color output
func canUseColor() bool {
	// If NO_COLOR environment variable is set, disable colors
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		return false
	}
	
	// Check if the TERM is set and contains "color"
	if term, ok := os.LookupEnv("TERM"); ok && strings.Contains(term, "color") {
		return true
	}
	
	// Default: enable on non-Windows systems
	return runtime.GOOS != "windows"
}
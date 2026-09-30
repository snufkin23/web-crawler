package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// copy-paste from AI model
// also i love using logger during development,helps me
// did some testing,good for me

type color string
type icon string
type lvl string

const (
	iconInfo    icon = "ℹ️ "
	iconSuccess icon = "✅ "
	iconWarn    icon = "⚠️ "
	iconError   icon = "❌"
	iconDebug   icon = "🔍"
)

const (
	colorReset   color = "\033[0m"
	colorRed     color = "\033[31m"
	colorGreen   color = "\033[32m"
	colorYellow  color = "\033[33m"
	colorBlue    color = "\033[34m"
	colorCyan    color = "\033[36m"
	colorDim     color = "\033[90m"
	colorBoldRed color = "\033[1;31m"
)

func colorizeLevel(level lvl) string {
	switch level {
	case "INFO":
		return string(colorBlue) + string(level) + string(colorReset)

	case "SUCCESS":
		return string(colorGreen) + string(level) + string(colorReset)

	case "WARN":
		return string(colorYellow) + string(level) + string(colorReset)

	case "ERROR":
		return string(colorRed) + string(level) + string(colorReset)

	case "FATAL":
		return string(colorBoldRed) + string(level) + string(colorReset)

	case "DEBUG":
		return string(colorCyan) + string(level) + string(colorReset)

	default:
		return string(level)
	}
}

func logMessage(level lvl, ic icon, formatStr string, a ...any) {
	timestamp := time.Now().Format("15:04:05")

	_, file, line, ok := runtime.Caller(2)

	caller := "unknown:0"

	if ok {
		caller = fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}

	rawMessage := fmt.Sprintf(formatStr, a...)
	lines := strings.Split(rawMessage, "\n")

	// 1. Level: tight bracket padding (fixed 9 visual columns)
	rawLevelBracket := fmt.Sprintf("[%s]", level)
	levelPadLen := max(9-len(rawLevelBracket), 0)

	coloredLevel := "[" +
		colorizeLevel(level) +
		"]" +
		strings.Repeat(" ", levelPadLen)

	// 2. Caller: tight bracket padding (fixed 13 visual columns)
	rawCallerBracket := fmt.Sprintf("[%s]", caller)
	callerPadLen := max(13-len(rawCallerBracket), 0)

	coloredCaller :=
		string(colorDim) + "[" + string(colorReset) +
			string(colorCyan) + caller + string(colorReset) +
			string(colorDim) + "]" + string(colorReset) +
			strings.Repeat(" ", callerPadLen)

	coloredTime :=
		string(colorDim) +
			timestamp +
			string(colorReset)

	// Prefix width:
	// 8 (time) + 1 + 9 (level) + 1 + 13 (caller) + 1 = 33 spaces
	fmt.Printf(
		"%s %s %s %s %s\n",
		coloredTime,
		coloredLevel,
		coloredCaller,
		ic,
		lines[0],
	)

	// Branch connector sits exactly at column 33 under the icon
	prefixIndent := strings.Repeat(" ", 33)

	for i := 1; i < len(lines); i++ {
		connector := "├─ "

		if i == len(lines)-1 {
			connector = "└─ "
		}

		branchMsg := lines[i]

		if level == "ERROR" || level == "FATAL" {
			branchMsg =
				string(colorRed) +
					branchMsg +
					string(colorReset)
		}

		fmt.Printf(
			"%s%s%s\n",
			prefixIndent,
			connector,
			branchMsg,
		)
	}
}

func NewLine() {
	fmt.Println()
}

func Divider() {
	fmt.Println(
		string(colorDim) +
			"----------------------------------------------------------------------------------" +
			string(colorReset),
	)
}

func Info(formatStr string, a ...any) {
	logMessage("INFO", iconInfo, formatStr, a...)
}

func Success(formatStr string, a ...any) {
	logMessage("SUCCESS", iconSuccess, formatStr, a...)
}

func Warn(formatStr string, a ...any) {
	logMessage("WARN", iconWarn, formatStr, a...)
}

func Error(formatStr string, a ...any) {
	logMessage("ERROR", iconError, formatStr, a...)
}

func Debug(formatStr string, a ...any) {
	logMessage("DEBUG", iconDebug, formatStr, a...)
}

func Fatal(formatStr string, a ...any) {
	logMessage("FATAL", iconError, formatStr, a...)
	os.Exit(1)
}

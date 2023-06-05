package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func LogResponse(method, path string, statusCode int) {
	statusColor := getColorForStatusCode(statusCode)
	methodColor := color.New(color.FgCyan).SprintFunc()

	fmt.Printf("Response: %s %s - Status: %s\n", methodColor(method), path, statusColor(statusCode))
}

func getColorForStatusCode(statusCode int) func(a ...interface{}) string {
	switch {
	case statusCode >= 200 && statusCode < 300:
		return color.New(color.FgGreen).SprintFunc()
	case statusCode >= 300 && statusCode < 400:
		return color.New(color.FgYellow).SprintFunc()
	case statusCode >= 400 && statusCode < 500:
		return color.New(color.FgMagenta).SprintFunc()
	default:
		return color.New(color.FgRed).SprintFunc()
	}
}

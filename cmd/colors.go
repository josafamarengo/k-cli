package cmd

import "github.com/fatih/color"

var (
    Cyan    = color.New(color.FgCyan).SprintFunc()
    Yellow  = color.New(color.FgYellow).SprintFunc()
    Red     = color.New(color.FgRed).SprintFunc()
    Green   = color.New(color.FgGreen).SprintFunc()
    Blue    = color.New(color.FgBlue).SprintFunc()
    Magenta = color.New(color.FgMagenta).SprintFunc()

    InvalidArg = color.New(color.FgRed).Add(color.Underline).SprintFunc()
)


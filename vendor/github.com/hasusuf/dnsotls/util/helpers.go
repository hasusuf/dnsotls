package util

import (
	"bufio"
	"fmt"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"io"
	"os"
	"reflect"
)

func GetFlagString(cmd *cobra.Command, flag string) string {
	s, err := cmd.Flags().GetString(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return s
}

// GetFlagStringSlice can be used to accept multiple argument with flag repetition (e.g. -f arg1,arg2 -f arg3 ...)
func GetFlagStringSlice(cmd *cobra.Command, flag string) []string {
	s, err := cmd.Flags().GetStringSlice(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return s
}

// GetFlagStringArray can be used to accept multiple argument with flag repetition (e.g. -f arg1 -f arg2 ...)
func GetFlagStringArray(cmd *cobra.Command, flag string) []string {
	s, err := cmd.Flags().GetStringArray(flag)
	if err != nil {
		glog.Fatalf("err accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return s
}

func GetFlagBool(cmd *cobra.Command, flag string) bool {
	b, err := cmd.Flags().GetBool(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return b
}

// Assumes the flag has a default value.
func GetFlagInt(cmd *cobra.Command, flag string) int {
	i, err := cmd.Flags().GetInt(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return i
}

// Assumes the flag has a default value.
func GetFlagInt32(cmd *cobra.Command, flag string) int32 {
	i, err := cmd.Flags().GetInt32(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return i
}

// Assumes the flag has a default value.
func GetFlagInt64(cmd *cobra.Command, flag string) int64 {
	i, err := cmd.Flags().GetInt64(flag)
	if err != nil {
		glog.Fatalf("error accessing flag %s for command %s: %v", flag, cmd.Name(), err)
	}
	return i
}

func IsFlagPresent(cmd *cobra.Command, name string) bool {
	b := cmd.Flags().Lookup(name)

	return b.Changed
}

func GetType(args interface{}) string {
	return fmt.Sprint(reflect.TypeOf(args))
}

func IsEmpty(needle interface{}) bool {

	if GetType(needle) == "string" {
		return len(needle.(string)) == 0
	}

	return false
}

func GetFirstLineFromFile(file string) (line string, err error) {
	lastLine := 0

	fileHandler, _ := os.Open(file)
	defer fileHandler.Close()

	scanHandler := bufio.NewScanner(fileHandler)
	for scanHandler.Scan() {
		lastLine++
		if lastLine == 1 {
			line = scanHandler.Text()

			return scanHandler.Text(), scanHandler.Err()
		}
	}

	return line, io.EOF
}

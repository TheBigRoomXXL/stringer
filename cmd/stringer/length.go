package stringer

import (
    "fmt"

    "github.com/TheBigRoomXXL/stringer/pkg/stringer"
    "github.com/spf13/cobra"
)

var onlyDigits bool
var onlyRunes bool

var lengthCmd = &cobra.Command{
    Use:   "length",
    Aliases: []string{"len"},
    Short:  "Give the len of a string",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {

        i := args[0]
        res, kind := stringer.Length(i, onlyDigits, onlyRunes)

        pluralS := "s"
        if res == 1 {
            pluralS = ""
        }
        fmt.Printf("'%s' has a %d %s%s.\n", i, res, kind, pluralS)
    },
}

func init() {
    lengthCmd.Flags().BoolVarP(&onlyDigits, "digit", "d", false, "Count only the digits")
    lengthCmd.Flags().BoolVarP(&onlyRunes, "rune", "r", false, "Count only the runes")
    rootCmd.AddCommand(lengthCmd)
}

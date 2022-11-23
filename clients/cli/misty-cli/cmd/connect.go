/*
Copyright © 2022 Rohit Singh

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect is used to connect to a misty broker instance",
	Long: `
connect is used to connect to a misty broker instance
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the hostname from the provided flags
		hostname, err := cmd.Flags().GetString("host")
		if err != nil {
			return fmt.Errorf("error while getting hostname flag input: %q", err)
		}
		// Get the port from the provided flags
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			return fmt.Errorf("error while getting port flag input: %q", err)
		}
		// Get the name from the provided flags
		brokerName, err := cmd.Flags().GetString("name")
		if err != nil {
			return fmt.Errorf("error while getting name flag input: %q", err)
		}

		fmt.Printf("attempting to connect to %s=%s:%d...\n", brokerName, hostname, port)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringP("host", "H", "localhost", "misty broker hostname")
	connectCmd.Flags().IntP("port", "p", 1315, "misty broker port number")
	connectCmd.Flags().StringP("name", "n", "default", "misty broker local reference name")
}

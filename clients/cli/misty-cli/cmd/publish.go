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
	"rohitsingh/misty-cli/actions"

	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish -H {BROKER_HOSTNAME} -p {BROKER_PORT} -t {TOPIC} -m {MESSAGE}",
	Short: "publish sends a payload to a message topic",
	Long: `
publish sends a payload to a message topic
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the hostname from the provided flags
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			return fmt.Errorf("error while getting hostname flag input: %q", err)
		}
		// Get the port from the provided flags
		port, err := cmd.Flags().GetInt("port")
		if err != nil {
			return fmt.Errorf("error while getting port flag input: %q", err)
		}
		// Get the topic name from the provided flags
		topic, err := cmd.Flags().GetString("topic")
		if err != nil {
			return fmt.Errorf("error while getting topic flag input: %q", err)
		}
		if topic == "" {
			return fmt.Errorf("missing flag for topic (-t): field is required")
		}
		// Get the message from the provided flags
		message, err := cmd.Flags().GetString("message")
		if err != nil {
			return fmt.Errorf("error while getting message flag input: %q", err)
		}
		if message == "" {
			return fmt.Errorf("missing flag for message (-m): field is required")
		}
		return actions.Publish(host, port, topic, message)
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
	publishCmd.Flags().StringP("host", "H", "localhost", "misty broker hostname")
	publishCmd.Flags().IntP("port", "p", 1315, "misty broker port number")
	publishCmd.Flags().StringP("topic", "t", "", "topic on which to publish message (required)")
	publishCmd.Flags().StringP("message", "m", "", "message to be published (required)")
}

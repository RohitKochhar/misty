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

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "subscribe listens for message on a specified topic",
	Long: `
subscribe listens for message on a specified topic
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the hostname from the provided flags
		brokerHost, err := cmd.Flags().GetString("host")
		if err != nil {
			return fmt.Errorf("error while getting hostname flag input: %q", err)
		}
		// Get the port from the provided flags
		brokerPort, err := cmd.Flags().GetInt("port")
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
		return actions.SubscribeAction(brokerHost, brokerPort, topic)
	},
}

func init() {
	rootCmd.AddCommand(subscribeCmd)
	subscribeCmd.Flags().StringP("topic", "t", "", "topic on which to publish message (required)")
	subscribeCmd.Flags().StringP("host", "H", "localhost", "misty broker hostname")
	subscribeCmd.Flags().IntP("port", "p", 1315, "misty broker port number")
}

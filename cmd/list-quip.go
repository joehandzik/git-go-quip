// Copyright © 2016 Joe Handzik <joseph.t.handzik@hpe.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
        "github.com/mduvall/go-quip"
)

// list-quipCmd represents the list-quip command
var list_quipCmd = &cobra.Command{
	Use:   "list-quip",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
                quip_token, _ := cmd.Flags().GetString("quip-token")
                quip_client := quip.NewClient(quip_token)
                getThreadStruct := quip.GetRecentThreadsParams{
                }
	        threads := quip_client.GetRecentThreads(&getThreadStruct)
                fmt.Println("Here are your Quip Threads:")
                for _, thread := range threads {
                    thread_id := thread.Thread["id"]
                    expanded_thread := quip_client.GetThread(thread_id)
                    fmt.Println("Thread:", thread_id)
                    for key, thread_map_element := range expanded_thread.Thread {
                        fmt.Println("Key:", key, "Value:", thread_map_element)
                    }
                }
        },
}

func init() {
	RootCmd.AddCommand(list_quipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// list-quipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// list-quipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

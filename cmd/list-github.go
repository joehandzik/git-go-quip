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
        "github.com/google/go-github/github"
        "golang.org/x/oauth2"
)

// list-githubCmd represents the list-github command
var list_githubCmd = &cobra.Command{
	Use:   "list-github",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
                github_token, _ := cmd.Flags().GetString("github-token")

                token_source := oauth2.StaticTokenSource(
                                  &oauth2.Token{AccessToken: github_token})
                token_client := oauth2.NewClient(oauth2.NoContext, token_source)

                github_client := github.NewClient(token_client)

                repos, _, _ := github_client.Repositories.List("", nil)

                for index, repo := range repos {
                    fmt.Printf("Repo %d: %s\n", index, *repo.FullName)
		}
	},
}

func init() {
	RootCmd.AddCommand(list_githubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// list-githubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// list-githubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

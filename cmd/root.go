// Copyright © 2018 Jacob Adlers <jacob.adlers@gmail.com>
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
	"os"
	"strings"

	"github.com/jadlers/zzng/apiseeds"
	"github.com/jadlers/zzng/genius"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zzng",
	Short: "Get lyrics for songs.",
	Long: `CLI to get information on those songs and artists that might be playing
	around you at this exact moment. Or maybe you just have to know more about
	that song you've got stuck in your head.

Use genius to get links to listen on YouTube or Spotify and links to lyrics on
their website.

Use apiseeds to get lyrics directly back as text, not nearly as many songs are supported.`,
}

var cmdGenius = &cobra.Command{
	Use:   "genius [search]",
	Short: "Get links for lyrics, Spotify, and Spotify from Genius.",

	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		search := strings.Join(args, " ")
		res := genius.GetLinksFor(search)
		fmt.Println(res)
	},
}

var cmdApiseeds = &cobra.Command{
	Use:   "apiseeds [artist] [song]",
	Short: "Use Apiseeds to get lyrics.",

	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		lyrics, err := apiseeds.GetLyrics(args[0], args[1])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Println(lyrics)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(cmdGenius)
	rootCmd.AddCommand(cmdApiseeds)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

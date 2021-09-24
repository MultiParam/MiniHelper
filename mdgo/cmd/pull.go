package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/multiparam/minihelper/mdgo/pkg/mdfile"

	"github.com/spf13/cobra"
)

type PullOptions struct {
	StorePath string
	Update    bool

	MdFiles []string
}

type Pull struct {
	Options PullOptions
}

var pullOpts = &PullOptions{}

var pullCmd = &cobra.Command{
	Use:  "pull",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pull := &Pull{
			Options: *pullOpts,
		}

		if err := pull.VerifyAndComplete(args); err != nil {
			return err
		}

		if err := pull.Run(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().StringVarP(&pullOpts.StorePath, "path", "p", "", "The path to store pictures.")
	pullCmd.Flags().BoolVarP(&pullOpts.Update, "update", "u", false, "To update the markdown file.")
}

func (pull *Pull) VerifyAndComplete(args []string) error {
	pull.Options.MdFiles = args

	if len(pull.Options.StorePath) == 0 {
		return fmt.Errorf("you must specify the store path of markdown pictures, `--path ./images/`")
	}

	return nil
}

func (pull *Pull) Run() error {
	// pull all pictures file by file
	for _, file := range pull.Options.MdFiles {
		if err := pull.Pull(file); err != nil {
			return err
		}

		if pull.Options.Update {
			pull.UpdateLinks(file)
		}
	}

	return nil
}

func (pull *Pull) Pull(file string) error {
	links, err := mdfile.GetAllPictureLinks(file)
	if err != nil {
		return err
	}

	if !mdfile.PathExists(pull.Options.StorePath) {
		if err := os.MkdirAll(pull.Options.StorePath, os.ModePerm); err != nil {
			return err
		}
	}

	dir := filepath.Dir(file)

	for _, link := range links {
		fmt.Println(link)

		// 处理网络链接的情况
		if mdfile.IsHttpLink(link) {
			p := filepath.Join(pull.Options.StorePath, filepath.Base(link))
			err := mdfile.DownloadPic(link, p)
			if err != nil {
				return err
			}
			fmt.Println(p)

			continue
		}

		// copy files
		dstPath := filepath.Join(pull.Options.StorePath, filepath.Base(link))
		if !filepath.IsAbs(link) {
			link = filepath.Join(dir, link)
			if link == dstPath {
				fmt.Println("hello")
			}
		}

		if err := mdfile.CopyRegularFile(link, dstPath); err != nil {
			return err
		}
	}

	return nil
}

func (pull *Pull) UpdateLinks(file string) error {
	dir := filepath.Dir(file)
	prefix, err := filepath.Rel(dir, pull.Options.StorePath)
	if err != nil {
		return err
	}

	if err := mdfile.ModifyAllPictureLines(file, prefix); err != nil {
		return err
	}

	return nil
}

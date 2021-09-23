package cmd

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/multiparam/minihelper/mdgo/pkg/mdfile"
	"github.com/multiparam/minihelper/mdgo/pkg/oss"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type PushOptions struct {
	EndPoint   string
	BucketName string
	AccessID   string
	AccessKey  string
	Path       string

	Domain  string
	Config  string
	MdFiles []string
}

type Push struct {
	Options PushOptions
}

const TempPicDir = "./temp-mdgo-pics"

var pushOpts = &PushOptions{}

var pushCmd = &cobra.Command{
	Use:  "push",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		push := &Push{
			Options: *pushOpts,
		}

		if err := push.VerifyAndComplete(args); err != nil {
			return err
		}

		if err := push.Run(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	pushCmd.Flags().StringVar(&pushOpts.EndPoint, "endpoint", "", "The endpoint of OSS picture bed.")
	pushCmd.Flags().StringVar(&pushOpts.BucketName, "bucket", "", "The bucket of OSS picture bed.")
	pushCmd.Flags().StringVar(&pushOpts.AccessID, "accessid", "", "The access id of OSS picture bed.")
	pushCmd.Flags().StringVar(&pushOpts.AccessKey, "accesskey", "", "The access secret of OSS picture bed.")
	pushCmd.Flags().StringVar(&pushOpts.Path, "path", "", "The store path of OSS picture bed.")
	pushCmd.Flags().StringVarP(&pushOpts.Config, "config", "c", "", "The config of the OSS picture bed.")
}

func (push *Push) VerifyAndComplete(args []string) error {
	push.Options.MdFiles = args

	if len(pushOpts.Config) != 0 {
		viper.SetConfigFile(pushOpts.Config)
		viper.SetConfigType("yaml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	if len(push.Options.EndPoint) == 0 {
		push.Options.EndPoint = viper.GetString("pushOptions.endPoint")
	}

	if len(push.Options.BucketName) == 0 {
		push.Options.BucketName = viper.GetString("pushOptions.bucketName")
	}

	if len(push.Options.AccessID) == 0 {
		push.Options.AccessID = viper.GetString("pushOptions.accessID")
	}

	if len(push.Options.AccessKey) == 0 {
		push.Options.AccessKey = viper.GetString("pushOptions.accessKey")
	}

	if len(push.Options.Path) == 0 {
		push.Options.Path = viper.GetString("pushOptions.path")
	}

	if len(push.Options.Domain) == 0 {
		push.Options.Domain = viper.GetString("pushOptions.domain")
	}

	if len(push.Options.EndPoint) == 0 || len(push.Options.BucketName) == 0 || len(push.Options.AccessID) == 0 || len(push.Options.AccessKey) == 0 || len(push.Options.Path) == 0 {
		return fmt.Errorf("you must specify all the arguments from command or config file.")
	}

	return nil
}

func (push *Push) Run() error {
	// get oss picture bed connection
	picBed := &oss.OSSPictureBed{
		EndPoint:   push.Options.EndPoint,
		BucketName: push.Options.BucketName,
		AccessID:   push.Options.AccessID,
		AccessKey:  push.Options.AccessKey,
	}
	if err := picBed.Connect(); err != nil {
		return err
	}

	// get prefix of new pic link
	if len(push.Options.Domain) == 0 {
		push.Options.Domain = "http://" + picBed.BucketName + "." + picBed.EndPoint
	}
	u, err := url.Parse(push.Options.Domain)
	if err != nil {
		return err
	}
	u.Path = filepath.Join(u.Path, push.Options.Path)

	// push all pictures and modify the markdown files.
	for _, file := range push.Options.MdFiles {
		if err := push.Push(file, picBed); err != nil {
			return err
		}
		if err := mdfile.ModifyAllPictureLines(file, u.String()); err != nil {
			return err
		}
	}

	return nil
}

func (push *Push) Push(file string, picBed *oss.OSSPictureBed) error {
	dir := filepath.Dir(file)

	links, err := mdfile.GetAllPictureLinks(file)
	if err != nil {
		return err
	}

	if !mdfile.PathExists(TempPicDir) {
		if err := os.MkdirAll(TempPicDir, os.ModePerm); err != nil {
			return err
		}
		defer os.Remove(TempPicDir)
	}

	for _, link := range links {
		fmt.Println(link)

		// 处理网络链接的情况
		if mdfile.IsHttpLink(link) {
			p := filepath.Join(TempPicDir, filepath.Base(link))
			err := mdfile.DownloadPic(link, p)
			if err != nil {
				return err
			}
			// fmt.Println(p)

			if err := picBed.UploadPic(filepath.Join(push.Options.Path, filepath.Base(p)), p); err != nil {
				return err
			}

			if err := os.Remove(p); err != nil {
				return err
			}

			continue
		}

		// 处理绝对路径的情况
		if filepath.IsAbs(link) {
			// fmt.Println(filepath.Join(push.options.Path, filepath.Base(link)))

			if err := picBed.UploadPic(filepath.Join(push.Options.Path, filepath.Base(link)), link); err != nil {
				return err
			}

			continue
		}

		// 处理相对路径的情况
		if !filepath.IsAbs(link) {
			p := filepath.Join(dir, link)
			// fmt.Println(filepath.Join(push.options.Path, filepath.Base(p)))

			if err := picBed.UploadPic(filepath.Join(push.Options.Path, filepath.Base(p)), p); err != nil {
				return err
			}

			continue
		}
	}

	return nil
}

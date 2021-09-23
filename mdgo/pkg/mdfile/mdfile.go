package mdfile

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const PictureFormat = ".*!\\[.*\\](.*).*" // markdown image link format
const TempFileName = ".temp.md"

// GetAllPictureLinks returns all the picture links from the file named mdfileName.
func GetAllPictureLinks(mdFilename string) ([]string, error) {
	file, err := os.Open(mdFilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	res := []string{}

	reg, err := regexp.Compile(PictureFormat)
	if err != nil {
		return []string{}, err
	}

	for true {
		line, _, err := reader.ReadLine()

		if err != nil && err == io.EOF {
			break
		}

		if err != nil {
			return []string{}, err
		}

		match := reg.Match(line)
		if match {
			link := GetPictureLink(string(line))
			res = append(res, link)
		}
	}

	return res, nil
}

// ModifyAllPictureLines modifies all the picture links using
func ModifyAllPictureLines(mdFilename, prefix string) error {
	file, err := os.Open(mdFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	fileTemp, err := os.OpenFile(TempFileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer fileTemp.Close()
	writer := bufio.NewWriter(fileTemp)

	reg, err := regexp.Compile(PictureFormat)
	if err != nil {
		return err
	}

	for true {
		line, _, err := reader.ReadLine()

		if err != nil && err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		match := reg.Match(line)
		if match {
			newLine, err := GetNewPictureLine(string(line), prefix)
			if err != nil {
				return err
			}
			writer.WriteString(newLine + "\n")
		} else {
			writer.WriteString(string(line) + "\n")
		}
	}
	writer.Flush()

	err = os.Rename(TempFileName, mdFilename)
	if err != nil {
		return err
	}

	return nil
}

// GetNewPictureLine returns the new picture line.
func GetNewPictureLine(line, prefix string) (string, error) {
	oldLink := GetPictureLink(line)
	if IsHttpLink(prefix) {
		u, err := url.Parse(prefix)
		if err != nil {
			return "", err
		}
		u.Path = filepath.Join(u.Path, filepath.Base(oldLink))
		newLine := strings.Replace(line, oldLink, u.String(), -1)

		return newLine, nil
	}

	return "", nil
}

// GetPictureLink returns the picture link from line.
func GetPictureLink(line string) string {
	return strings.Split(strings.Split(line, "](")[1], ")")[0]
}

func IsHttpLink(link string) bool {
	return strings.HasPrefix(link, "http://") || strings.HasPrefix(link, "https://")
}

func DownloadPic(picURL, path string) error {
	resp, err := http.Get(picURL)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(resp.Status, "200") {
		return fmt.Errorf("download pic %s error, the status of respond is %s", picURL, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}

	if _, err := io.Copy(out, bytes.NewReader(body)); err != nil {
		return err
	}

	return nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}

	return true
}
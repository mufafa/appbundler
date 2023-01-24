package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mufafa/AppBundler/builder"
	"github.com/mufafa/AppBundler/plist"
)

const (
	basedir = "dist"
)

type Icon interface{}

type AppleBundle struct {
	BundleIdentifier string
	ExecutableName   string
	ProductName      string
	IconFile         Icon
}

func GeneratePlistFile(app AppleBundle) string {
	vars := map[string]string{
		"EXECUTABLE_NAME":   app.ExecutableName,
		"PRODUCT_NAME":      app.ProductName,
		"BUNDLE_IDENTIFIER": app.BundleIdentifier,
		"ICON_FILE":         app.IconFile.(string),
	}

	var tmpFile string = plist.PListFile

	for key, value := range vars {
		old := "${" + key + "}"
		tmpFile = strings.ReplaceAll(tmpFile, old, value)
	}
	return tmpFile
}

func CopyFile(src string, dest string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = ioutil.WriteFile(dest, input, fs.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func createFolder(path string) error {
	err := os.MkdirAll(path, 0750)
	if err != nil && !os.IsExist(err) {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func createInfoPlist(path string, app AppleBundle, str string) error {
	plist_file := path + "/info.plist"
	err := os.WriteFile(plist_file, []byte(str), 0660)
	if err != nil {
		return err
	}
	return nil
}

func isDirExist(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func main() {
	nameFlag := flag.String("name", "My Application", "Application Product name")
	iconFlag := flag.String("icon", "", "Application icon file")
	identiferFlag := flag.String("id", "com.myapplication.org", "Application Bundle identtifier")
	buildFlag := flag.Bool("build", false, "Build go project")
	flag.Parse()

	appBundle := AppleBundle{
		BundleIdentifier: *identiferFlag,
		ProductName:      *nameFlag,
		ExecutableName:   *nameFlag,
		IconFile:         *iconFlag,
	}

	if *buildFlag {
		builder.Build("", appBundle.ProductName)
	}

	dest := basedir + "/" + appBundle.ProductName + ".app"
	checkDir := isDirExist(dest)
	if !checkDir {
		err := createFolder(dest)
		if err != nil {
			log.Fatal(err)
		}
	}
	content := dest + "/Contents"
	resources := content + "/Resources"
	macos := content + "/MacOS"

	dirs := []string{
		content,
		resources,
		macos,
	}

	for _, v := range dirs {
		createFolder(v)
	}

	//create plist file
	plist := GeneratePlistFile(appBundle)
	err := createInfoPlist(content, appBundle, plist)
	if err != nil {
		log.Fatal(err)
	}

	//move executable file
	executable_path := macos + "/" + appBundle.ProductName
	executable_file := "./" + appBundle.ProductName
	err = os.Rename(executable_file, executable_path)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//copy icon file
	icon_file_name := appBundle.IconFile.(string)
	fileExtension := filepath.Ext(icon_file_name)

	if len(icon_file_name) != 0 {
		if fileExtension == "icns" {
			icon_name := appBundle.IconFile.(string)
			icon_file := icon_name
			icon_dest := resources + "/" + icon_name
			err = CopyFile(icon_file, icon_dest)
			if err != nil {
				log.Fatalln(err.Error())
			}
		} else {
			fmt.Println("icon file should be 'icns' image")
		}
	}

	fmt.Printf("%v bundle created on /dist folder", appBundle.ExecutableName)
}

# Go Apple Bundle Builder

**AppBundler** is an tool for macOS. It lets you create Apple Application Bundle.

## Build AppBundler

These steps will put the AppBundler repository under /Users/yourlogin/Documents.
```
$ cd ~/Documents   #or where you want to store the repository#
$ git clone https://github.com/mufafa/appbundler
$ cd appbundler

$ ./go build -o AppBundler
```

## Install AppBundler
These steps will put the AppBundler executable under /usr/local/bin. 
```
$ cd ~/Documents   #or where you have built executable file#
$ mv AppleBunler /usr/local/bin
$ echo 'export PATH=/usr/local/bin:$PATH' >> ~/Desktop/zshrc.txt  #or wherever you want to install#
```
## Using AppBundler
You can use AppBundler both go project or any other macOS executable binary file.

```
$ cd ~/Desktop/<executable_file_path>
$ AppBundler \
        -name=<application_name> \ #default: My Application
        -id=<bundle_identification> \ #default: com.myapplication.org
        -icon=<icns_file_path>  \ #eg: ~/Desktop/<your_path>/icon.icns
        -build=<true|false> #default:false <set true if you need to build go file>
            
$ cd ./dist 
```
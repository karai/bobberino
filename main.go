package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// Attribution constants
const appName = "Bobberino"
const appDev = "The TurtleCoin Developers"
const appDescription = appName + " - Cross compile all the things"
const appLicense = "https://choosealicense.com/licenses/mit/"
const appRepository = "https://github.com/karai/bobberino"
const appURL = "https://karai.io https://turtlecoin.lol"

var licensevar bool
var filename string

var t time.Time = time.Now()
var timestamp string = t.Format("20060102150405")

// Version string
func semverInfo() string {
	var majorSemver, minorSemver, patchSemver, wholeString string
	majorSemver = "0"
	minorSemver = "1"
	patchSemver = "0"
	wholeString = majorSemver + "." + minorSemver + "." + patchSemver
	return wholeString
}

func parseFlags() {
	flag.StringVar(&filename, "file", "main.go", "File to compile, main.go by default")
	flag.BoolVar(&licensevar, "license", false, "Show the license before we get started.")
	flag.Parse()
}

func main() {
	parseFlags()
	announce()
	if licensevar {
		printLicense()
	}

	crossCompile("aix", "ppc64", "", filename)
	crossCompile("android", "amd64", "", filename)
	crossCompile("android", "arm64", "", filename)
	crossCompile("darwin", "amd64", "", filename)
	crossCompile("darwin", "arm64", "", filename)
	crossCompile("dragonfly", "amd64", "", filename)
	crossCompile("freebsd", "amd64", "", filename)
	crossCompile("freebsd", "arm64", "", filename)
	crossCompile("js", "wasm", ".wasm", filename)
	crossCompile("linux", "amd64", "", filename)
	crossCompile("linux", "arm64", "", filename)
	crossCompile("netbsd", "amd64", "", filename)
	crossCompile("netbsd", "arm64", "", filename)
	crossCompile("openbsd", "amd64", "", filename)
	crossCompile("openbsd", "arm64", "", filename)
	crossCompile("plan9", "amd64", "", filename)
	crossCompile("solaris", "amd64", "", filename)
	crossCompile("windows", "amd64", ".exe", filename)
	crossCompile("windows", "arm", ".exe", filename)
}

func announce() {
	color.Set(color.FgHiWhite)
	logrus.Info(appName + " v" + semverInfo() + " running!")
	logrus.Info("Building: ", filename)
	logrus.Info("Starting in 2 seconds, press CTRL C to interrupt.")
	time.Sleep(2 * time.Second)
}

func zipFile(filename string) {

}

func crossCompile(osName, archName, extension, filename string) {
	gobin := "go"
	gobuild := "build"
	goflags := "-o"
	descriptiveFilename := "builds/" + osName + archName + "/" + strings.TrimRight(filename, ".go") + extension
	fullZipDir := "builds/" + osName + archName + "/"
	gofilename := filename
	osENV := "GOOS=" + osName
	archENV := "GOARCH=" + archName
	cmd := exec.Command(gobin, gobuild, goflags, descriptiveFilename, gofilename)
	newEnv := append(os.Environ(), osENV, archENV)
	cmd.Env = newEnv

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		color.Set(color.FgHiRed)
		fmt.Println(osName + "/" + archName + "\t - ⛔ PROBLEM")
		logrus.Debug(err)
	} else {
		color.Set(color.FgHiGreen)
		fmt.Println(osName + "/" + archName + "\t - ✔️ DONE")

		app0 := "zip"
		app1 := "-r"
		app2 := "./builds/" + osName + "-" + archName + "-" + timestamp + ".zip"
		app3 := fullZipDir
		cmd := exec.Command(app0, app1, app2, app3)
		stdout, _ := cmd.Output()
		logrus.Debug(stdout)
	}
	logrus.Debug(string(stdout))
}

// Print the license for the user
func printLicense() {
	fmt.Printf(color.GreenString("\n"+appName+" v"+semverInfo()) + color.WhiteString(" by "+appDev))
	color.Set(color.FgGreen)
	fmt.Println("\n" + appRepository + "\n" + appURL + "\n")
	color.Set(color.FgHiWhite)
	fmt.Println("\nMIT License\nCopyright (c) 2020-2021 RockSteady, TurtleCoin Developers")
	color.Set(color.FgHiBlack)
	fmt.Println("\nPermission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the 'Software'), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:\n\nThe above copyright notice and this permission notice shall be included in allcopies or substantial portions of the Software.\n\nTHE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.")
	fmt.Println()
}

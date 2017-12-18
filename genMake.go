package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Make file generator for Golang (xgo)")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nWorking golang directory: ")
	wdir, _ := reader.ReadString('\n')
	wdir = normalize(wdir)
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Build output directory: ")
	bdir, _ := reader.ReadString('\n')
	bdir = normalize(bdir)
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Go version: ")
	goVersion, _ := reader.ReadString('\n')
	goVersion = normalize(goVersion)
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Test build targets: ")
	Tt, _ := reader.ReadString('\n')
	Tt = normalize(Tt)
	Ttargets := strings.Fields(Tt)
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Full build targets: ")
	Ft, _ := reader.ReadString('\n')
	Ft = normalize(Ft)
	Ftargets := strings.Fields(Ft)
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Test build prefix: ")
	TBpr, _ := reader.ReadString('\n')
	TBpr = normalize(TBpr)
	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Full build prefix: ")
	FBpr, _ := reader.ReadString('\n')
	FBpr = normalize(FBpr)

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("\nMakefile export directory: ")
	expDir, _ := reader.ReadString('\n')
	expDir = normalize(expDir)

	makefile := ""

	// Build
	makefile += "default:\n	cd " + bdir + " && xgo --targets="
	for i, a := range Ftargets {
		if i+1 != len(Ftargets) {
			makefile += a + ","
		} else {
			makefile += a
		}
	}
	makefile += " -go " + goVersion
	makefile += " -out " + FBpr
	makefile += " " + wdir

	// Build and Run
	makefile += "\n\nrbuild:\n	cd " + bdir + " && xgo --targets="
	for i, a := range Ttargets {
		if i+1 != len(Ttargets) {
			makefile += a + ","
		} else {
			makefile += a
		}
	}
	makefile += " -go " + goVersion
	makefile += " -out " + TBpr
	makefile += " " + wdir
	makefile += " && ./" + TBpr + "-linux-amd64"

	// Test build
	makefile += "\n\ntbuild:\n	cd " + bdir + " && xgo --targets="
	for i, a := range Ttargets {
		if i+1 != len(Ttargets) {
			makefile += a + ","
		} else {
			makefile += a
		}
	}
	makefile += " -go " + goVersion
	makefile += " -out " + TBpr
	makefile += " " + wdir

	file, err := os.Create(expDir + "/Makefile")
	if err != nil {
		fmt.Println(err)
	}

	file.Close()

	file, err = os.OpenFile(expDir+"/Makefile", os.O_RDWR, 0644)

	_, err = file.WriteString(makefile)
	if err != nil {
		fmt.Println(err)
	}

	file.Close()

}

func normalize(input string) string {
	return input[:len(input)-1]
}

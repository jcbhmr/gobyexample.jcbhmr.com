// 2>/dev/null; exec go run "$0" "$@"
//go:build ignore
package main

import (
	"log"
	"os/exec"
	"path/filepath"
	"flag"
	"os"
	"fmt"
	"regexp"
	"strings"
	"errors"
)

type args struct {
	dir string
	destDir *string
}

func argsParse() *args {
	a := &args{}
	f := flag.NewFlagSet("mdbook-build-i18n", flag.ExitOnError)
	f.Usage = func() {
		fmt.Fprintf(f.Output(), `Builds a book from its markdown files for a specific language

Usage: mdbook-build-i18n [OPTIONS] [dir]

Arguments:
  [dir]  Root directory for the book
         (Defaults to the current directory when omitted)

Options:
  -d, --dest-dir <dest-dir>  Output directory for the book
                             Relative paths are interpreted relative to the book's root directory.
                             If omitted, mdBook uses build.build-dir from book.toml or defaults to 'book'.
`)
	}
	destDir := f.String("dest-dir", "", "")
	f.StringVar(destDir, "d", "", "")
	f.Parse(os.Args[1:])
	dir := f.Arg(0)
	if dir == "" {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Fprint(f.Output(), err)
			os.Exit(1)
		}
		a.dir = wd
	} else {
		dir2, err := filepath.Abs(dir)
		if err != nil {
			fmt.Fprint(f.Output(), err)
			os.Exit(1)
		}
		a.dir = dir2
	}
	if *destDir != "" {
		destDir2 := filepath.Clean(filepath.Join(a.dir, *destDir))
		a.destDir = &destDir2
	}
	return a
}

type book struct {
	build bookBuild
	output bookOutput
	preprocessor bookPreprocessor
}
type bookBuild struct {
	buildDir string
}
type bookOutput struct {
	html bookOutputHTML
}
type bookOutputHTML struct {
	siteURL string
}
type bookPreprocessor struct {
	gettext bookPreprocessorGettext
}
type bookPreprocessorGettext struct {
	poDir string
}

func bookLoadDir(dir string) (*book, error) {
	bookTOMLBytes, err := os.ReadFile(filepath.Join(dir, "book.toml"))
	if err != nil {
		return nil, err
	}
	bookTOMLString := string(bookTOMLBytes)
	
	b := &book{}

	var buildDir string
	buildRe := regexp.MustCompile(`\[build\]\n[\s\S]*?build-dir\s*=\s*(".*?"|'.*?')`)
	buildMatch := buildRe.FindStringSubmatch(bookTOMLString)
	if buildMatch != nil {
		buildString := buildMatch[1]
		if strings.Contains(buildString, "\\") {
			return nil, errors.New("build-dir contains backslashes")
		}
		buildDir = buildString[1:len(buildString)-1]
	} else {
		buildDir = "book"
	}
	buildDir2 := filepath.Clean(filepath.Join(dir, buildDir))
	b.build.buildDir = buildDir2

	var siteURL string
	siteURLRe := regexp.MustCompile(`\[output\.html\]\n[\s\S]*?site-url\s*=\s*(".*?"|'.*?')`)
	siteURLMatch := siteURLRe.FindStringSubmatch(bookTOMLString)
	if siteURLMatch != nil {
		siteURLString := siteURLMatch[1]
		if strings.Contains(siteURLString, "\\") {
			return nil, errors.New("site-url contains backslashes")
		}
		siteURL = siteURLString[1:len(siteURLString)-1]
	} else {
		siteURL = "/"
	}
	b.output.html.siteURL = siteURL

	var poDir string
	poDirRe := regexp.MustCompile(`\[preprocessor\.gettext\]\n[\s\S]*?po-dir\s*=\s*(".*?"|'.*?')`)
	poDirMatch := poDirRe.FindStringSubmatch(bookTOMLString)
	if poDirMatch != nil {
		poDirString := poDirMatch[1]
		if strings.Contains(poDirString, "\\") {
			return nil, errors.New("po-dir contains backslashes")
		}
		poDir = poDirString[1:len(poDirString)-1]
	} else {
		poDir = "po"
	}
	poDir2 := filepath.Clean(filepath.Join(dir, poDir))
	b.preprocessor.gettext.poDir = poDir2

	return b, nil
}

func main() {
	log.SetFlags(0)

	args2 := argsParse()

	book, err := bookLoadDir(args2.dir)
	if err != nil {
		log.Fatalf("could not load book: %v", err)
	}
	var destDir string
	if args2.destDir != nil {
		destDir = *args2.destDir
	} else {
		destDir = book.build.buildDir
	}

	files, err := os.ReadDir(book.preprocessor.gettext.poDir)
	if err != nil {
		log.Fatalf("could not read directory %s: %v", book.preprocessor.gettext.poDir, err)
	}
	langs := make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if strings.HasSuffix(name, ".po") {
			langs = append(langs, name[:len(name)-3])
		}
	}

	cmd := exec.Command("mdbook", "build", "--dest-dir", destDir, args2.dir)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("$ %s", cmd)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("failed to build book: %v", err)
	}

	for _, lang := range langs {
		cmd := exec.Command("mdbook", "build", "--dest-dir", filepath.Join(destDir, lang), args2.dir)
		cmd.Env = append(os.Environ(), "MDBOOK_BOOK__LANGUAGE="+lang, "MDBOOK_OUTPUT__HTML__SITE_URL="+book.output.html.siteURL+lang+"/", "MDBOOK_OUTPUT__HTML__REDIRECT={}")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Printf("$ %s", cmd)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("failed to build book %s: %v", lang, err)
		}
	}

	if _, err := exec.LookPath("i18n-report"); err == nil {
		args := []string{"report", filepath.Join(destDir, "report.html")}
		for _, lang := range langs {
			args = append(args, filepath.Join(book.preprocessor.gettext.poDir, lang+".po"))
		}
		cmd := exec.Command("i18n-report", args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		log.Printf("$ %s", cmd)
		err := cmd.Run()
		if err != nil {
			log.Fatalf("failed to generate i18n report: %v", err)
		}
	} else {
		log.Println("skipping report.html: i18n-report not found")
	}
}

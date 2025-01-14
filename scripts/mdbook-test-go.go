//go:build ignore

// 2>/dev/null; exec go run "$0" "$@"
//
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strings"
)

type args struct {
	dir     string
	destDir *string
	chapter *string
}

func argsParse() *args {
	a := &args{}
	f := flag.NewFlagSet("mdbook-test-go", flag.ExitOnError)
	f.Usage = func() {
		fmt.Fprintf(f.Output(), `Tests that a book's Go code samples compile

Usage: mdbook-test-go [OPTIONS] [dir]

Arguments:
  [dir]  Root directory for the book
         (Defaults to the current directory when omitted)

Options:
  -d, --dest-dir <dest-dir>  Output directory for the book
                             Relative paths are interpreted relative to the book's root directory.
                             If omitted, mdBook uses build.build-dir from book.toml or defaults to 'book'.
  -c, --chapter <chapter>
  -h, --help                 Print help
  -v, --version              Print version
`)
	}
	destDir := f.String("dest-dir", "", "")
	f.StringVar(destDir, "d", "", "")
	chapter := f.String("chapter", "", "")
	f.StringVar(chapter, "c", "", "")
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
	if *chapter != "" {
		chapter2 := *chapter
		a.chapter = &chapter2
	}
	return a
}

type book struct {
	book  bookBook
	build bookBuild
}
type bookBook struct {
	src string
}
type bookBuild struct {
	buildDir string
}

func bookLoadDir(dir string) (*book, error) {
	bookTOMLBytes, err := os.ReadFile(filepath.Join(dir, "book.toml"))
	if err != nil {
		return nil, err
	}
	bookTOMLString := string(bookTOMLBytes)

	b := &book{}

	var src string
	srcRe := regexp.MustCompile(`\[book\]\n[\s\S]*?src\s*=\s*(".*?"|'.*?')`)
	srcMatch := srcRe.FindStringSubmatch(bookTOMLString)
	if srcMatch != nil {
		srcString := srcMatch[1]
		if strings.Contains(srcString, "\\") {
			return nil, errors.New("src contains backslashes")
		}
		src = srcString[1 : len(srcString)-1]
	} else {
		src = "src"
	}
	src2, err := filepath.Abs(filepath.Join(dir, src))
	if err != nil {
		return nil, err
	}
	b.book.src = src2

	var buildDir string
	buildRe := regexp.MustCompile(`\[build\]\n[\s\S]*?build-dir\s*=\s*(".*?"|'.*?')`)
	buildMatch := buildRe.FindStringSubmatch(bookTOMLString)
	if buildMatch != nil {
		buildString := buildMatch[1]
		if strings.Contains(buildString, "\\") {
			return nil, errors.New("build-dir contains backslashes")
		}
		buildDir = buildString[1 : len(buildString)-1]
	} else {
		buildDir = "book"
	}
	buildDir2 := filepath.Clean(filepath.Join(dir, buildDir))
	b.build.buildDir = buildDir2

	return b, nil
}

type summary struct {
	chapters []summaryChapter
}
type summaryChapter struct {
	title string
	path  string
}

func summaryLoadDir(dir string) (*summary, error) {
	summaryMDBytes, err := os.ReadFile(filepath.Join(dir, "SUMMARY.md"))
	if err != nil {
		return nil, err
	}
	summaryMDString := string(summaryMDBytes)

	s := &summary{}

	chapterRe := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	chapterMatches := chapterRe.FindAllStringSubmatch(summaryMDString, -1)
	s.chapters = []summaryChapter{}
	for _, chapterMatch := range chapterMatches {
		title := chapterMatch[1]
		relativePath := chapterMatch[2]
		if relativePath == "" {
			continue
		}
		path := filepath.Clean(filepath.Join(dir, relativePath))
		s.chapters = append(s.chapters, summaryChapter{title, path})
	}

	return s, nil
}

type goCodeBlock struct {
	code        string
	ignore      bool
	compileFail bool
	noRun       bool
	shouldPanic bool
}

func goCodeBlockFindAllStringSubmatch(s string) ([]goCodeBlock, error) {
	goCodeBlockRe := regexp.MustCompile("```" + `(go.*?)\n([\s\S]*?)\n` + "```")
	goCodeBlockMatches := goCodeBlockRe.FindAllStringSubmatch(s, -1)
	goCodeBlocks := []goCodeBlock{}
	for _, goCodeBlockMatch := range goCodeBlockMatches {
		code := goCodeBlockMatch[2]
		attributes := strings.Split(goCodeBlockMatch[1], ",")[1:]
		ignore := slices.Contains(attributes, "ignore")
		compileFail := slices.Contains(attributes, "compile_fail")
		noRun := slices.Contains(attributes, "no_run")
		shouldPanic := slices.Contains(attributes, "should_panic")
		goCodeBlocks = append(goCodeBlocks, goCodeBlock{code, ignore, compileFail, noRun, shouldPanic})
	}
	return goCodeBlocks, nil
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
	_ = destDir

	summary, err := summaryLoadDir(book.book.src)
	if err != nil {
		log.Fatalf("could not load summary: %v", err)
	}

	chapters := []summaryChapter{}
	if args2.chapter != nil {
		for _, chapter := range summary.chapters {
			if chapter.title == *args2.chapter {
				chapters = append(chapters, chapter)
				break
			}
		}
	} else {
		chapters = summary.chapters
	}

	for _, chapter := range chapters {
		relativeChapterPath, err := filepath.Rel(book.book.src, chapter.path)
		if err != nil {
			log.Fatalf("could not get relative chapter path: %v", err)
		}
		log.Printf(`Testing chapter '%s': "%s"`, chapter.title, relativeChapterPath)

		mdBytes, err := os.ReadFile(chapter.path)
		if err != nil {
			log.Fatalf("could not read md file: %v", err)
		}
		mdString := string(mdBytes)

		goCodeBlocks, err := goCodeBlockFindAllStringSubmatch(mdString)
		if err != nil {
			log.Fatalf("could not find go code blocks: %v", err)
		}

		for _, goCodeBlock := range goCodeBlocks {
			if goCodeBlock.ignore {
				continue
			}

			tmpDir, err := os.MkdirTemp("", "mdbook-test-go")
			if err != nil {
				log.Fatalf("could not create temp dir: %v", err)
			}
			defer os.RemoveAll(tmpDir)

			err = os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte(goCodeBlock.code), 0644)
			if err != nil {
				log.Fatalf("could not write main.go: %v", err)
			}

			cmd := exec.Command("go", "build", "./main.go")
			cmd.Dir = tmpDir
			_, err = cmd.CombinedOutput()
			if goCodeBlock.compileFail {
				if err == nil {
					log.Fatalf("expected go code to not compile")
				}
				continue
			}
			if err != nil {
				log.Fatalf("could not compile go code: %v", err)
			}

			if goCodeBlock.noRun {
				continue
			}

			var exeExt string
			if runtime.GOOS == "windows" {
				exeExt = ".exe"
			}
			cmd = exec.Command(filepath.Join(tmpDir, "main"+exeExt))
			_, err = cmd.CombinedOutput()
			if goCodeBlock.shouldPanic {
				if err == nil {
					log.Fatalf("expected go code to panic")
				}
				continue
			}
			if err != nil {
				log.Fatalf("running go code failed: %v", err)
			}
		}
	}
}

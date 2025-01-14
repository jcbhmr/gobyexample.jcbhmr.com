// 2>/dev/null; exec go run "$0" "$@"
//go:build ignore

package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strings"

	"github.com/BurntSushi/toml"
	flag "github.com/spf13/pflag"
)

var dir string
var destDir = flag.StringP("dest-dir", "d", "", "Output directory for the book\nRelative paths are interpreted relative to the book's root directory.\nIf omitted, mdBook uses build.build-dir from book.toml or defaults to 'book'")
var chapter = flag.StringP("chapter", "c", "", "")

func main() {
	log.SetFlags(0)

	flag.Parse()
	dir, _ = filepath.Abs(flag.Arg(0))
	*destDir = filepath.Clean(filepath.Join(dir, *destDir))

	book, err := bookLoadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	if *destDir == "" {
		*destDir = filepath.Clean(filepath.Join(dir, book.build.buildDir))
	}

	summary, err := summaryLoadDir(book.book.src)
	if err != nil {
		log.Fatal(err)
	}

	if *chapter != "" {
		slices.DeleteFunc(summary.chapters, func(c summaryChapter) bool {
			return c.title != *chapter
		})
	}

	for _, chapter := range summary.chapters {
		relativeChapterPath, _ := filepath.Rel(book.book.src, chapter.path)
		log.Printf(`Testing chapter '%s': "%s"`, chapter.title, relativeChapterPath)

		mdBytes, err := os.ReadFile(chapter.path)
		if err != nil {
			log.Fatal(err)
		}
		goCodeBlocks, err := goCodeBlockFindAllSubmatch(mdBytes)
		if err != nil {
			log.Fatal(err)
		}
		for _, goCodeBlock := range goCodeBlocks {
			if goCodeBlock.ignore {
				continue
			}

			tmpDir, err := os.MkdirTemp("", "mdbook-test-go")
			if err != nil {
				log.Fatal(err)
			}
			defer os.RemoveAll(tmpDir)

			err = os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte(goCodeBlock.code), 0644)
			if err != nil {
				log.Fatal(err)
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
				log.Fatalf("compiling go code failed: %v", err)
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

type book struct {
	book struct {
		src string
	}
	build struct {
		buildDir string `toml:"build-dir"`
	}
}

func bookLoadDir(dir string) (*book, error) {
	bookTOMLBytes, err := os.ReadFile(filepath.Join(dir, "book.toml"))
	if err != nil {
		return nil, err
	}
	b := &book{}
	err = toml.Unmarshal(bookTOMLBytes, b)
	if err != nil {
		return nil, err
	}
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
	s := &summary{}
	// TODO: Use a Markdown parser instead.
	chapterRe := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	chapterMatches := chapterRe.FindAllSubmatch(summaryMDBytes, -1)
	s.chapters = []summaryChapter{}
	for _, chapterMatch := range chapterMatches {
		title := string(chapterMatch[1])
		relativePath := string(chapterMatch[2])
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

func goCodeBlockFindAllSubmatch(b []byte) ([]goCodeBlock, error) {
	// TODO: Use a Markdown parser instead.
	goCodeBlockRe := regexp.MustCompile("```" + `(go.*?)\n([\s\S]*?)\n` + "```")
	goCodeBlockMatches := goCodeBlockRe.FindAllSubmatch(b, -1)
	goCodeBlocks := []goCodeBlock{}
	for _, goCodeBlockMatch := range goCodeBlockMatches {
		code := string(goCodeBlockMatch[2])
		attributes := strings.Split(string(goCodeBlockMatch[1]), ",")[1:]
		ignore := slices.Contains(attributes, "ignore")
		compileFail := slices.Contains(attributes, "compile_fail")
		noRun := slices.Contains(attributes, "no_run")
		shouldPanic := slices.Contains(attributes, "should_panic")
		goCodeBlocks = append(goCodeBlocks, goCodeBlock{code, ignore, compileFail, noRun, shouldPanic})
	}
	return goCodeBlocks, nil
}

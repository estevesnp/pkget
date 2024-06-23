package fetch

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/estevesnp/pkgo/internal/text"
)

func FetchPackages(pkg string, limit int) ([]string, error) {
	url := fmt.Sprintf("https://pkg.go.dev/search?limit=%d&m=package&q=%s", limit, pkg)
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching search results: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error fetching search results: %s", res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error parsing document: %w", err)
	}

	pkgs := make([]string, 0, limit)
	doc.Find(".SearchSnippet-header-path").Each(func(i int, s *goquery.Selection) {
		if i < limit {
			pkgs = append(pkgs, strings.Trim(s.Text(), "()"))
		}
	})

	return pkgs, nil
}

func SpinWhileFetching(pkg string, limit int) ([]string, error) {
	done := make(chan bool)

	txt := fmt.Sprintf("Fetching packages for %q ", pkg)
	go text.Spinner(txt, text.Basic, 100*time.Millisecond, done)

	t := time.Now()
	pkgs, err := FetchPackages(pkg, limit)
	done <- true
	fmt.Printf("\r%s \n", txt)
	fmt.Printf("Done in %vms\n\n", time.Since(t).Milliseconds())

	return pkgs, err
}

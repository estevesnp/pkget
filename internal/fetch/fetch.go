package fetch

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FetchPackages(pkg string, limit int) ([]string, bool, error) {
	url := fmt.Sprintf("https://pkg.go.dev/search?limit=%d&m=package&q=%s", limit, pkg)
	res, err := http.Get(url)
	if err != nil {
		return nil, false, fmt.Errorf("error fetching search results: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, false, fmt.Errorf("error fetching search results: %s", res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, false, fmt.Errorf("error parsing document: %w", err)
	}

	pkgs := make([]string, 0, limit)
	doc.Find(".SearchSnippet-header-path").Each(func(i int, s *goquery.Selection) {
		if i < limit {
			pkgs = append(pkgs, strings.Trim(s.Text(), "()"))
		}
	})

	if len(pkgs) == 0 {
		return nil, false, nil
	}

	return pkgs, true, nil
}

package jenny

import (
	//"github.com/bndr/gojenkins"
	"fmt"
	"github.com/pkg/browser"
)

var client string

func OpenLink(j Jenkins) {
	l := fmt.Sprintf("%s/%s", j.Uri, j.Project)
	browser.OpenURL(l)
}

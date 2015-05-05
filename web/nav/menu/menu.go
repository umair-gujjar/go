package menu

import (
    "bytes"
    "html/template"
    "sort"
    "strings"
    "web/nav"
)

var funcMap = template.FuncMap{
    "sort": func(p *nav.Pages) *nav.Pages {
        sort.Sort(p)
        return p
    },
}

var Html string = `
<li><a href="{{ .Url }}">{{ .Label }}</a>
    {{ if .Pages  }}
    <ul>
        {{ range sort .Pages }}
        {{ with . }}{{ template "menu" . }}{{ end }}
        {{ end }}
    </ul>
</li>
{{ else }}
</li>
{{ end }}`

type Menu struct {
    *nav.Nav
}

func New(n *nav.Nav) *Menu {
    return &Menu{n}
}

func PageRender(p *nav.Page) string {
    var output bytes.Buffer
    t := template.Must(template.New("menu").Funcs(funcMap).Parse(Html))
    t.Execute(&output, p)
    return output.String()
}

func HtmlMinify(s string) string {
    s = strings.Replace(s, "\n", "", -1)
    s = strings.Replace(s, "\t", "", -1)
    return s
}

func (m Menu) String() string {
    out := `<ul>`
    for _, p := range m.Pages {
        out += PageRender(p)
    }
    out += `</ul>`
    return HtmlMinify(out)
}

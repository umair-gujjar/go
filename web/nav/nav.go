package nav

type Page struct {
	Label   string
	Url     string
	Active  bool
	Visible bool
	Order   int
	Pages   Pages
}

type Pages []*Page

func (p Pages) Len() int           { return len(p) }
func (p Pages) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Pages) Less(i, j int) bool { return p[i].Order < p[j].Order }

func NewPage(label, url string) *Page {
	return &Page{Label: label, Url: url, Order: 0, Pages: make(Pages, 0)}
}

func (p *Page) Add(page *Page) {
	p.Pages = append(p.Pages, page)
}

type Nav struct {
	Name  string
	Pages map[string]*Page
}

func NewNav(name string) *Nav {
	return &Nav{name, make(map[string]*Page)}
}

func (n *Nav) Add(name string, page *Page) {
	n.Pages[name] = page
}

func (n Nav) GetName() string {
	return n.Name
}

func (n Nav) GetPage(name string) *Page {
	if p, ok := n.Pages[name]; ok {
		return p
	}
	return nil
}

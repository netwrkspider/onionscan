package webui

import (
	"github.com/s-rah/onionscan/config"
	"github.com/s-rah/onionscan/crawldb"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"log"
)

type WebUI struct {
	osc  *config.OnionScanConfig
	Done chan bool
}

type Section struct {
        Title template.HTML
        Relationships []crawldb.Relationship
}

type Content struct {
	SearchTerm  string
	NumResults  int
	Results     []crawldb.Relationship
	Error       string
	OnionRecord crawldb.CrawlRecord
	Sections    []Section
}

func (wui *WebUI) Index(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")
	var content Content
	if search != "" {
		content.SearchTerm = search
                results, err := wui.osc.Database.GetRelationshipsWithIdentifier(search)

		if err == nil {
			content.NumResults = len(results)
		        sections := make(map[string][]crawldb.Relationship)
		        for _,relationship := range results {
		                sections[relationship.From] = append(sections[relationship.From], relationship)
		        }
		        
		        log.Printf("Got %d Relationship types", len(sections))
		        
		        for k,v := range sections {
		                content.Sections = append(content.Sections, Section{template.HTML(k),v})
		        }
		} else {
		        content.Error = err.Error()
		}
	}

	var templates = template.Must(template.ParseFiles("templates/index.html"))
	templates.ExecuteTemplate(w, "index.html", content)
}

func (wui *WebUI) Onion(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("onion")
	var content Content
	if search != "" {
		content.SearchTerm = search
		if strings.HasSuffix(search, ".onion") {
			results, err := wui.osc.Database.GetIdentifierWithOnion(search)

			if err == nil {
				content.NumResults = len(results)
			        sections := make(map[string][]crawldb.Relationship)
			        for _,relationship := range results {
			                sections[relationship.From] = append(sections[relationship.From], relationship)
			        }
			        
			        log.Printf("Got %d Relationship types", len(sections))
			        
			        for k,v := range sections {
			                content.Sections = append(content.Sections, Section{template.HTML(k),v})
			        }
			} else {
			        content.Error = err.Error()
			}
		}
	}

	var templates = template.Must(template.ParseFiles("templates/index.html"))
	templates.ExecuteTemplate(w, "index.html", content)
}

func (wui *WebUI) Listen(osc *config.OnionScanConfig, port int) {
	wui.osc = osc
	http.HandleFunc("/", wui.Index)
	http.HandleFunc("/onion", wui.Onion)

	fs := http.FileServer(http.Dir("./templates/style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))
	portstr := strconv.Itoa(port)
	http.ListenAndServe("127.0.0.1:"+portstr, nil)
}

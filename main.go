package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var PAGEL = [11]string{
	"summary",      // show best of each page here
	"bloglist",     // social media posts and computer journaling
	"researchlist", // program research posted here
	"projectlist",  // paid & personal projects
	"resume",       // a fancy UI decorated resume, make docx version downloadable
	"githubresume", // a github generated resume
	"wisdom",       // wisdom
}

type blog struct {
	title       string `json:"title"`
	date        string `json:"date"`
	description string `json:"description"`
	article     string `json:"article"`
}

type Blog struct {
	blogL []blog `json:"blogL"`
}

type research struct {
	title       string `json:"title"`
	date        string `json:"date"`
	description string `json:"description"`
	article     string `json:"article"`
}

type Research struct {
	researchL []research `json:"researchL"`
}

type project struct {
	title       string `json:"title"`
	date        string `json:"date"`
	description string `json:"description"`
	article     string `json:"article"`
}

type Project struct {
	projectL []project `json:"projectL"`
}

var _blogL Blog
var blogL []blog

var _researchL Research
var researchL []research

var _projectL Project
var projectL []project

func main() {
	fmt.Println("app is running")

	var blogsJSONFile, err = os.Open("./blogL.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened blogL.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer blogsJSONFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(blogsJSONFile)
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &_blogL)
	//
	//
	var researchLJSONFile, err1 = os.Open("./researchL.json")
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened researchL.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer researchLJSONFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ = ioutil.ReadAll(researchLJSONFile)
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &_researchL)
	//
	//
	var projectLJSONFile, err2 = os.Open("./projectL.json")
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened projectL.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer projectLJSONFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ = ioutil.ReadAll(projectLJSONFile)
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &_projectL)

	researchL = _researchL.researchL
	blogL = _blogL.blogL
	projectL = _projectL.projectL

	// The first thing to do is to associate the hello component with a path.
	//
	// This is done by calling the Route() function,  which tells go-app what
	// component to display for a given path, on both client and server-side.
	// app.Route("/", &RootApp{})
	app.Route("/summary", &Summary{})
	app.Route("/bloglist", &BlogList{})
	app.Route("/researchlist", &ResearchList{})
	app.Route("/projectlist", &ProjectList{})
	app.Route("/resume", &Resume{})
	app.Route("/githubresume", &GithubResume{})
	app.Route("/wisdom", &Wisdom{})
	// app.RouteWithRegexp("^/bar.*", &bar) // bar component is associated with all paths that start with /bar.

	// Once the routes set up, the next thing to do is to either launch the app
	// or the server that serves the app.
	//
	// When executed on the client-side, the RunWhenOnBrowser() function
	// launches the app,  starting a loop that listens for app events and
	// executes client instructions. Since it is a blocking call, the code below
	// it will never be executed.
	//
	// When executed on the server-side, RunWhenOnBrowser() does nothing, which
	// lets room for server implementation without the need for precompiling
	// instructions.
	app.RunWhenOnBrowser()

	// Finally, launching the server that serves the app is done by using the Go
	// standard HTTP package.
	//
	// The Handler is an HTTP handler that serves the client and all its
	// required resources to make it work into a web browser. Here it is
	// configured to handle requests with a path that starts with "/".
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		// PreRenderCache: app.NewPreRenderLRUCache(100*100000, time.Hour), // 10MB/1hour
		// Icon: app.Icon{
		// 	Default:    "/web/logo.png",       // Specify default favicon.
		// 	AppleTouch: "/web/logo-apple.png", // Specify icon on IOS devices.
		// },
		// Styles: []string{
		// 	"/web/hello.css", // Loads hello.css file.
		// },
		// Scripts: []string{
		// 	"/web/hello.js", // Loads hello.js file.
		// },
		// app.GitHubPages("REPOSITORY_NAME")
		// Resources: app.RemoteBucket("https://storage.googleapis.com/myapp.appspot.com"),
		// Resources:   app.LocalDir("/tmp/web"),
		// RawHeaders: []string{
		// 	`<!-- Global site tag (gtag.js) - Google Analytics -->
		// 	<script async src="https://www.googletagmanager.com/gtag/js?id=UA-xxxxxxx-x"></script>
		// 	<script>
		// 	  window.dataLayer = window.dataLayer || [];
		// 	  function gtag(){dataLayer.push(arguments);}
		// 	  gtag('js', new Date());
		// 	  gtag('config', 'UA-xxxxxx-x');
		// 	</script>
		// 	`,
		// },
	})
	for i := range blogL {
		route := fmt.Sprintf("/blogL/%d.txt", i)
		http.Handle(route, &app.Handler{})
	}
	for i := range researchL {
		route := fmt.Sprintf("/researchL/%d.txt", i)
		http.Handle(route, &app.Handler{})
	}
	for i := range projectL {
		route := fmt.Sprintf("/projectL/%d.txt", i)
		http.Handle(route, &app.Handler{})
	}
	// err := app.GenerateStaticWebsite(".", &app.Handler{
	// 	Name:        "Hello",
	// 	Description: "An Hello World! example",
	// })

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type RootApp struct {
	app.Compo
	// apptitle       string
	// appdescription string
	Name        string
	Description string
}

// func (r *RootApp) Render() app.UI {
// 	return app.Div().Body(
// 		&Navbar{},
// 	// currentPage
// 	// &Footer{},
// 	)
// }

type Navbar struct {
	app.Compo
}

// func (n *Navbar) Render( /*items []string*/ ) app.UI {
// fmt.Println(items)
// return app.Ul().Body(
// 	app.Range(items).Slice(func(i int) app.UI {
// 		return app.Li().Text(items[i])
// 	}),
// )
// }

func (n *Navbar) Render( /*items []string*/ ) app.UI {
	return app.Ul().Body(
		app.Range(PAGEL).Slice(func(i int) app.UI {
			// return app.Li()
			// .Text(PAGES[i])
			str := `<a href="/` + PAGEL[i] + `">` + PAGEL[i] + `</a>`
			return app.Raw(`<li>` + str + `</li>`)
		}),
	)
}

// import 'bootstrap/dist/css/bootstrap.css';
// import Nav from 'react-bootstrap/Nav';

// type Footer struct {
// 	app.Compo
// }

// func (f *Footer) Render(items []string) app.UI {
// 	return app.Ul().Body(
// 		app.Range(items).Slice(func(i int) app.UI {
// 			return app.Li().Text(items[i])
// 		}),
// 	)
// }

// type Page struct {
// 	app.Compo
// }

type Summary struct {
	app.Compo
}

func (s *Summary) Render() app.UI {
	return app.Div().Body().Text("Summary")
}

/*

These combined together can be reused for
all kinds of database models, but I am not sure
how to make the styling work?

type ListPage struct {

}

type itemPage struct {

}
*/

type BlogList struct {
	app.Compo
}

func (b *BlogList) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Blogs"),
		app.Range(blogL).Slice(func(i int) app.UI {
			return &BlogItem{
				title:       blogL[i].title,
				description: blogL[i].description,
				date:        blogL[i].date,
				article:     blogL[i].article,
			}
		}),
	)
}

type BlogItem struct {
	app.Compo
	title       string
	description string
	date        string
	article     string
}

func (b *BlogItem) Render() app.UI {
	return app.Div().Body(
		app.H3().Text(b.title),
		app.P().Text(b.description),
		app.P().Text(b.date),
		app.P().Text(b.article),
	)
}

/*

 */

type ResearchList struct {
	app.Compo
}

func (r *ResearchList) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Research"),
		app.Range(researchL).Slice(func(i int) app.UI {
			return &ResearchItem{
				title:       researchL[i].title,
				description: researchL[i].description,
				date:        researchL[i].date,
				article:     researchL[i].article,
			}
		}),
	)
}

type ResearchItem struct {
	app.Compo
	title       string
	description string
	date        string
	article     string
}

func (r *ResearchItem) Render() app.UI {
	return app.Div().Body(
		app.H3().Text(r.title),
		app.P().Text(r.description),
		app.P().Text(r.date),
		app.P().Text(r.article),
	)
}

type ProjectList struct {
	app.Compo
}

func (p *ProjectList) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Research"),
		app.Range(projectL).Slice(func(i int) app.UI {
			return &ProjectItem{
				title:       projectL[i].title,
				description: projectL[i].description,
				date:        projectL[i].date,
				article:     projectL[i].article,
			}
		}),
	)
}

type ProjectItem struct {
	app.Compo
	title       string
	description string
	date        string
	article     string
}

func (p *ProjectItem) Render() app.UI {
	return app.Div().Body(
		app.H3().Text(p.title),
		app.P().Text(p.description),
		app.P().Text(p.date),
		app.P().Text(p.article),
	)
}

type Resume struct {
	app.Compo
}

func (r *Resume) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Resume Link"),
		app.IFrame().Src("https://docs.google.com/document/d/1jyROKcHxufZY7Np-Ls7RtOninCGHjABfuY8phqWSijk/edit?usp=sharing"),
	)
}

type GithubResume struct {
	app.Compo
}

func (r *GithubResume) Render() app.UI {
	return app.Div().Body(
		app.H1().Text("Github Resume Link"),
		app.IFrame().Src("https://resume.github.io/?godjtinkers365"),
	)
}

type Wisdom struct {
	app.Compo
}

func (o *Wisdom) Render() app.UI {
	return app.Div().Body().Text("Wisdom")
}

package home

import "html/template"

var VarHomeCards []HomeCard = []HomeCard{
	{
		Title:       "What is GoX?",
		Description: template.HTML("Understand what GoX solves"),
	},
	{
		Title:       "Using GoX",
		Description: template.HTML("Learn how to use GoX"),
	},
	{
		Title:       "Additions",
		Description: template.HTML("Additional technologies to use with GoX"),
	},
	{
		Title:       "Contribute",
		Description: template.HTML("Find out how you can contribute to GoX"),
	},
}

var VarHomeSections []HomeSection = []HomeSection{
	{
		Title:       "What's GoX?",
		Description: "While Go and HTMX are a joy to work with, there is no established way of using these technologies together in a way that scales; resulting in complexity in your codecase. GoX structures your code and hides this complexity behind certain primitives.",
		Section:     1,
		Cards:       cards1,
	},
	{
		Title:       "Using GoX",
		Description: "Reading documentation can be difficult, so we have created some examples to illutrate common usecases.",
		Section:     2,
		Cards:       cards2,
	},
	{
		Title:       "Additions",
		Description: "Go and HTMX alone are not enough to create a rich user experience. Below are some recommended technologies to add to your toolbox.",
		Section:     3,
		Cards:       cards3,
	},
	{
		Title:       "Contribute",
		Description: "GoX is an open-source project. PRs, Issues, Proposals will all be looked at",
		Section:     3,
		Cards:       cards4,
	},
}

var cards1 []HomeCard = []HomeCard{
	{
		Title:       "Pages",
		Description: template.HTML("Pages are a simple way to handle routing to either dynamic or static pages"),
		Link:        "/docs#pages",
	},
	{
		Title:       "Data",
		Description: template.HTML("Separate data from your HTML powering re-usuable components with dynamic data fetching"),
		Link:        "/docs#data-handling",
	},
	{
		Title:       "Render",
		Description: template.HTML("Render your page and html partials at build time"),
		Link:        "/docs#custom-rendering",
	},
	{
		Title:       "Handle",
		Description: template.HTML("Handlers used for html partials or full page reloads"),
		Link:        "/docs#custom-handling",
	},
}
var cards3 []HomeCard = []HomeCard{
	// {
	// 	Title:       "Go http/template",
	// 	Description: template.HTML("html/template implements data-driven templates for generating HTML output safe against code injection."),
	// 	Link:        "/",
	// },
	// {
	// 	Title:       "HTMX",
	// 	Description: template.HTML("HTMX is a library that allows you to access modern browser features directly from HTML, rather than using javascript."),
	// 	Link:        "/",
	// },
	{
		Title:       "Alpine",
		Description: template.HTML("Alpine is a rugged, minimal tool for composing behavior directly in your markup."),
		Link:        "/",
	},
	{
		Title:       "Material 3",
		Description: template.HTML("Material 3 is the latest version of Google’s open-source design system."),
		Link:        "/",
	},
	{
		Title:       "Tailwind",
		Description: template.HTML("A utility-first CSS framework directly in your markup."),
		Link:        "/",
	},
}

var cards2 []HomeCard = []HomeCard{
	{
		Title:       "SPA vs MPA",
		Description: template.HTML("Navigating to a new page does not always require a full refresh/render. Using HTMX and separate client/server state, we can get the best of both worlds."),
		Link:        "/",
	},
	{
		Title:       "SSR",
		Description: template.HTML("Even if the page is static, GoX will dynamically render components on the server."),
		Link:        "/",
	},
}

var cards4 []HomeCard = []HomeCard{
	{
		Title:       "GoX-Starter",
		Description: template.HTML("The GoX starter-kit is open-source."),
		Link:        "/",
	},
	{
		Title:       "GoX-Website",
		Description: template.HTML("The code for this current website is public."),
		Link:        "/",
	},
	{
		Title:       "TODOs",
		Description: template.HTML("GoX is a new project and far from production ready. Feel free to help with the TODO's!"),
		Link:        "/",
	},
}

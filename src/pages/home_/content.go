package home

import "html/template"

var VarHomeCards []HomeCard = []HomeCard{
	{
		Title:       "What is GoX?",
		Description: template.HTML("Understand what GoX solves"),
		Link:        "#what-is-gox",
	},
	{
		Title:       "Using GoX",
		Description: template.HTML("Learn how to use GoX"),
		Link:        "#using-gox",
	},
	{
		Title:       "Additions",
		Description: template.HTML("Additional technologies to use with GoX"),
		Link:        "#additions",
	},
	{
		Title:       "Contribute",
		Description: template.HTML("Find out how you can contribute to GoX"),
		Link:        "#contribute",
	},
}

var VarHomeSections []HomeSection = []HomeSection{
	{
		Title:       "What's GoX?",
		Description: "While Go and HTMX are a joy to work with, there is no established way of using these technologies together in a way that scales; resulting in complexity in your codecase. GoX structures your code and hides this complexity behind certain primitives.",
		Id:          "what-is-gox",
		Section:     1,
		Cards:       cards1,
		ImageVer:    "/static/assets/gox-mascot-vert.png",
		ImageHor:    "/static/assets/gox-mascot-hor.png",
	},
	{
		Title:       "Using GoX",
		Description: "Reading documentation can be difficult, so we have created some examples to illutrate common usecases.",
		Id:          "using-gox",
		Section:     2,
		Cards:       cards2,
		ImageVer:    "https://lh3.googleusercontent.com/RbMYWtKTLGv7tEAwsya6Z7NHcUYpn4gwkrp3zy9dVhN0jFppAE7VR12r1Hpgh1fZI3MhK3jUsG2xnCSrzpaLiJTxIHWO_CIARwfhe6naTzWo5VIdg8Y=w2400",
		ImageHor:    "https://lh3.googleusercontent.com/7wuUlEFgaFEA075w4_OLDilE1vwTPGX7_G5_tiF9iARu8xXu1b-K27vD4cA3KLdhXdwABG1_I6YxvPjeUfzaYe1oVkuIJ0wTvh4ng6k7pEKAQVJzzw=w2400",
	},
	{
		Title:       "Additions",
		Description: "Go and HTMX alone are not enough to create a rich user experience. Below are some recommended technologies to add to your toolbox.",
		Id:          "additions",
		Section:     3,
		Cards:       cards3,
		ImageVer:    "https://lh3.googleusercontent.com/k5ZNz1pSrSN1NMrxHqxkO8Je_n5S9Q3TVYpZriGx9J2w_mCbya0INifAWw4KGRoaZTavD8e_rbaW3UqRUjk5UC5dKYXuh5M3qyLfrhwE6jIfVrCw9aeL=w2400",
		ImageHor:    "https://lh3.googleusercontent.com/S7a0PvKXaPN1BjW3kffTstRaen4CVUD2nO1nfEk5z_GmjP5AD2fjfPdgfCig4cJMDpu11iGxdceqMJ9PrfU_PnP3bBtzwIN-IrXj4rTJGS6I7zx0=w2400",
	},
	{
		Title:       "Contribute",
		Description: "GoX is an open-source project. PRs, Issues, Proposals will all be looked at",
		Id:          "contribute",
		Section:     3,
		Cards:       cards4,
		ImageVer:    "https://lh3.googleusercontent.com/k5ZNz1pSrSN1NMrxHqxkO8Je_n5S9Q3TVYpZriGx9J2w_mCbya0INifAWw4KGRoaZTavD8e_rbaW3UqRUjk5UC5dKYXuh5M3qyLfrhwE6jIfVrCw9aeL=w2400",
		ImageHor:    "https://lh3.googleusercontent.com/S7a0PvKXaPN1BjW3kffTstRaen4CVUD2nO1nfEk5z_GmjP5AD2fjfPdgfCig4cJMDpu11iGxdceqMJ9PrfU_PnP3bBtzwIN-IrXj4rTJGS6I7zx0=w2400",
	},
}

var cards1 []HomeCard = []HomeCard{
	{
		Title:       "Pages",
		Description: template.HTML("Pages are a simple way to define and handle route specific UI"),
		Link:        "/docs/pages",
	},
	{
		Title:       "Data",
		Description: template.HTML("Separate data from your HTML powering reusuable components with dynamic data fetching"),
		Link:        "/docs/data",
	},
	{
		Title:       "Render",
		Description: template.HTML("Bespoke rendering processes for routes at build time"),
		Link:        "/docs/render",
	},
	{
		Title:       "Handle",
		Description: template.HTML("Custom handlers for html partials or full pages"),
		Link:        "/docs/handle",
	},
}
var cards3 []HomeCard = []HomeCard{
	{
		Title:       "Alpine",
		Description: template.HTML("Alpine is a rugged, minimal tool for composing behavior directly in your markup."),
		Link:        "https://alpinejs.dev/",
	},
	{
		Title:       "Material 3",
		Description: template.HTML("Material 3 is the latest version of Google’s open-source design system."),
		Link:        "https://m3.material.io/",
	},
	{
		Title:       "Tailwind",
		Description: template.HTML("A utility-first CSS framework directly in your markup."),
		Link:        "https://tailwindcss.com/",
	},
}

var cards2 []HomeCard = []HomeCard{
	{
		Title:       "Todo",
		Description: template.HTML("Shows you how to return static and dynamic HTML"),
		Link:        "/examples/todo",
	},
	{
		Title:       "Data Fetching",
		Description: template.HTML("Per-request data fetching for your pages"),
		Link:        "/examples/data",
	},
	{
		Title:       "Custom Rendering",
		Description: template.HTML("Have a build step for your HTML"),
		Link:        "/examples/render",
	},
	{
		Title:       "Pages",
		Description: template.HTML("Custom Rendering"),
		Link:        "/examples/pages",
	},
}

var cards4 []HomeCard = []HomeCard{
	{
		Title:       "gox-starter",
		Description: template.HTML("The GoX starter-kit is open-source."),
		Link:        "/",
	},
	{
		Title:       "gox-website",
		Description: template.HTML("The code for this current website is public."),
		Link:        "/",
	},
	{
		Title:       "TODOs",
		Description: template.HTML("GoX is a new project and far from production ready. Feel free to help with the TODO's!"),
		Link:        "/",
	},
}

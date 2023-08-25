package home

import "html/template"

var VarHomeCards []HomeCard = []HomeCard{
	{
		Title:       "What is GoX?",
		Description: template.HTML("Learn how GoX enables <i class='font-medium'>blazingly fast</i> speeds in production"),
	},
	{
		Title:       "Using GoX",
		Description: template.HTML("Get to know Gox best practices"),
	},
	{
		Title:       "Design",
		Description: template.HTML("Understand how to use Material 3 and Tailwind to create stunning design"),
	},
	{
		Title:       "Contribute",
		Description: template.HTML("Find out how you can contribute to GoX"),
	},
}

var VarHomeSections []HomeSection = []HomeSection{
	{
		Title:       "What's GoX?",
		Description: "GoX is an opinionated, modern, lightweight stack that focuses on Server-Side Rendering. Go html/templates and HTMX, with opt-in Alpine.js, allows you to access modern browser features within your HTML. For design, GoX outsources all standard component and design requirements to Material 3, supercharged by Tailwind integration.",
		Section:     1,
		Cards:       cards1,
	},
	{
		Title:       "Using GoX",
		Description: "GoX is not just a stack, but rather a set of opinionated rules within a set of technologies. These rules will maintain the usability, readability and performance of your code.",
		Section:     2,
		Cards:       cards2,
	},
	{
		Title:       "Design",
		Description: "Design as an engineering problem has been solved. GoX outsources all standard design to Matieral 3, which provides us with an opinionated theming system and an extensive component library. Additionally, GoX combines Material 3 with HTMX, Alpine & Tailwind, allowing us to keep our code within the HTML reducing additional complexity.",
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
		Title:       "Go http/template",
		Description: template.HTML("html/template implements data-driven templates for generating HTML output safe against code injection."),
		Link:        "/",
	},
	{
		Title:       "HTMX",
		Description: template.HTML("HTMX is a library that allows you to access modern browser features directly from HTML, rather than using javascript."),
		Link:        "/",
	},
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

var cards3 []HomeCard = []HomeCard{
	{
		Title:       "Tailwind vs CSS properties",
		Description: template.HTML("While Tailwind has access to all the theming from Material 3, Tailwind is not used to style Material 3 components."),
		Link:        "/",
	},
	{
		Title:       "Material 3 (pre-release)",
		Description: template.HTML("Material 3 for web is still in pre-release, resulting in some unstable components and missing documentaion."),
		Link:        "/",
	},
}

var cards4 []HomeCard = []HomeCard{
	{
		Title:       "Website Repository",
		Description: template.HTML("The code for the current website is public."),
		Link:        "/",
	},
	{
		Title:       "TODOs",
		Description: template.HTML("GoX is a new project and far from a complete stack. Feel free to help with the TODO's!"),
		Link:        "/",
	},
}

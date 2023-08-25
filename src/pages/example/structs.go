package example

type Task struct {
	Text string
}
type ExampleContent struct {
	HomeActive    bool
	ExampleActive bool
	DocsActive    bool
	ActiveTabId   string
	Tasks         []Task
}

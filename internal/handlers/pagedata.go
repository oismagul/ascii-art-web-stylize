package handlers

type PageData struct {
	Title       string
	Banner      string
	InputText   string
	AsciiResult string
	ErrorMsg    string
}

type ErrInfo struct {
	Title     string
	ErrorType int
	ErrorMsg  string
}

package utils

type PageData struct {
	Title string
	datas map[string]interface{}
}

func NewPageData(title string) *PageData {
	if title == "" {
		title = "LCV"
	}
	return &PageData{
		Title: title,
		datas: make(map[string]interface{}),
	}
}

package libraries

type Page struct {
	pageUrl string
}

var initPage *Page

/**
 * init
 */
func init()  {
	initPage = new(Page);
}

/**
 * set url page url
 */
func SetUrl(pageUrl string) string {
	initPage.pageUrl = pageUrl
	return initPage.pageUrl
}

func GetUrl() string{
	return initPage.pageUrl;
}

package interfaces

type Controller interface {
	RunHandler(url_pattern string)
}

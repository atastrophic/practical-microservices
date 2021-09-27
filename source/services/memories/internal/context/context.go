package context

type Context interface {
}

type _Context struct {
}

func GetApplicationContext() Context {
	return _Context{}
}

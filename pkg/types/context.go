package types

type ContextKey string

func (c ContextKey) String() string {
	return string(c)
}

var (
	ContextKeyBuildInformation = ContextKey("buildInformation")
	ContextKeyTemplatesFiles   = ContextKey("templatesFiles")
	ContextKeyStaticFiles      = ContextKey("staticFiles")
	ContextKeyIsDebug          = ContextKey("isDebug")
	ContextKeyRouteData        = ContextKey("routeData")
)

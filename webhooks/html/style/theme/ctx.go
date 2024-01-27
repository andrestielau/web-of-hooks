package theme

import "context"

type themeContextKey struct{}

func Set(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, themeContextKey{}, v)
}

func Get(ctx context.Context) string {
	if theme, ok := ctx.Value(themeContextKey{}).(string); ok {
		return theme
	}
	return ""
}

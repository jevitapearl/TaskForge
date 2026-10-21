package contextx

import "context"

type contextKey struct{}

type User struct {
  ID   string
  Role string
}

func SetUser(ctx context.Context, user User) context.Context {
  return context.WithValue(ctx, contextKey{}, user)
}

func FromContext(ctx context.Context) (User, bool) {
	user, ok := ctx.Value(contextKey{}).(User)
	return user, ok
}

func UserID(ctx context.Context) string {
	user, ok := FromContext(ctx)
	if !ok {
		return ""
	}
	return user.ID
}

func UserRole(ctx context.Context) string {
	user, ok := FromContext(ctx)
	if !ok {
			return ""
	}
	return user.Role
}
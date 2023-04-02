package selects

type Scope struct {
	Name  string
	Apply ScopeFunc
}
type ScopeFunc func(b *SubBuilder) *SubBuilder
type scopes []*Scope

func (s scopes) Clone() scopes {
	return s
}
func (s scopes) WithScope(scope *Scope) scopes {
	return append(s, scope)
}

func (s scopes) WithoutScope(scope *Scope) scopes {
	newScopes := make(scopes, 0, len(s))
	for _, sc := range s {
		if sc.Name != scope.Name {
			newScopes = append(newScopes, sc)
		}
	}
	return newScopes
}

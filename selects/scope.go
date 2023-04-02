package selects

import "github.com/abibby/bob/set"

type Scoper interface {
	Scopes() []*Scope
}

type Scope struct {
	Name  string
	Apply ScopeFunc
}
type ScopeFunc func(b *SubBuilder) *SubBuilder
type scopes struct {
	parent              any
	scopes              []*Scope
	withoutGlobalScopes set.Set[string]
}

func newScopes() *scopes {
	return &scopes{
		scopes:              []*Scope{},
		withoutGlobalScopes: set.New[string](),
	}
}

func (s *scopes) withParent(parnet any) *scopes {
	s.parent = parnet
	return s
}

func (s *scopes) Clone() *scopes {
	return &scopes{
		parent:              s.parent,
		scopes:              cloneSlice(s.scopes),
		withoutGlobalScopes: s.withoutGlobalScopes.Clone(),
	}
}
func (s *scopes) WithScope(scope *Scope) *scopes {
	s.scopes = append(s.scopes, scope)
	return s
}

func (s *scopes) WithoutScope(scope *Scope) *scopes {
	newScopes := make([]*Scope, 0, len(s.scopes))
	for _, sc := range s.scopes {
		if sc.Name != scope.Name {
			newScopes = append(newScopes, sc)
		}
	}
	s.scopes = newScopes
	return s
}

func (s *scopes) allScopes() []*Scope {
	if scoper, ok := s.parent.(Scoper); ok {
		allGlobalScopes := scoper.Scopes()
		globalScopes := make([]*Scope, 0, len(allGlobalScopes))
		for _, scope := range allGlobalScopes {
			if s.withoutGlobalScopes.Has(scope.Name) {
				continue
			}
			globalScopes = append(globalScopes, scope)
		}
		return append(s.scopes, globalScopes...)
	}

	return s.scopes
}

func (b *scopes) WithoutGlobalScope(scope *Scope) *scopes {
	b.withoutGlobalScopes.Add(scope.Name)
	return b
}

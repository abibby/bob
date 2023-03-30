package selects

type Join struct{}

func (j *Join) Join() *Join {
	return j
}

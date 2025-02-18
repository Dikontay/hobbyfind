package read

type Params struct {
	Id string `json:"-" params:"id"`
}

func (p *Params) Validate() error {
	if p.Id == "" {
		
	}
	return nil
}

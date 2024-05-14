package curso

type Curso struct {
	idCurso            string
	titulo             string
	conteudoPragmatico string
	horas              int
	regime             string
}

func (c *Curso) New(id string, t string, cp string, h int, r string) *Curso {
	c.idCurso = id
	c.titulo = t
	c.conteudoPragmatico = cp
	c.horas = h
	c.regime = r
	return c

}

func (c *Curso) GetIdCurso() string{
	return c.idCurso
}
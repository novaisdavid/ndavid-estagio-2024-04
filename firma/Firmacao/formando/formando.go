package formando

type Formando struct {
	idFormando string
	nome       string
	email      string
	telefone   string
}

func (f *Formando) New(idFormando, nome, email, telefone string) *Formando {
	f.idFormando = idFormando
	f.nome = nome
	f.email = email
	f.telefone = telefone
	return f
}
func (f Formando) GetIdFormando() string {
	return f.idFormando
}
func (f Formando) GetNomeFormando() string {
	return f.nome
}

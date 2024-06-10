package repositorio


type Repositorio interface{
	Salvar(dados string)
	LerDados() string

}
package matricula

type Matricula struct {
	IdFormando string
	IdCurso string
}

func (m *Matricula) MatricularFormando(idFormando, idCurso string) {
	m.IdFormando = idFormando
	m.IdCurso = idCurso
	// implementar o save que pode ser implementado na prórpia máquina

}

// fazendo save essa função é substituida
func (m Matricula) MostraEstudaMatriculado(idFormando string) Matricula {
	if m.IdFormando == idFormando {
		return m
	}

	return Matricula{}
}

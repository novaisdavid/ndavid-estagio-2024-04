package funcoes

import (
	//"strings"
)

/*func removeDuplicates(cursos []interface{}) []Curso {
	encountered := map[string]bool{}
	result := []Curso{}

	for _, curso := range cursos {
		if !encountered[curso] {
			encountered[curso] = true
			result = append(result, curso)
		}
	}

	return result
}*/
/*type Dados interface{}

func LerDadosDoArquivo(nomeArquivo string) ([]Dados, error) {
	// Ler os dados do arquivo
	data, err := ioutil.ReadFile(nomeArquivo)
	if err != nil {
		return nil, err
	}

	// Dividir os dados em linhas
	lines := strings.Split(string(data), "\n")

	// Criar uma slice para armazenar os dados
	dados := make([]Dados, 0)

	// Iterar sobre as linhas e criar um item para cada conjunto de dados
	for _, line := range lines {
		if len(line) == 0 {
			continue // ignorar linhas em branco
		}
		item, err := parseItem(line)
		if err != nil {
			fmt.Println("Erro ao fazer o parse do item:", err)
			continue
		}
		dados = append(dados, item)
	}

	return dados, nil
}
/*func RemoveDuplicates(dados string) string {
	
	lines := strings.Split(string(data), "\n")

	// Criar uma slice para armazenar os cursos
	cursos := make([]Curso, 0)

	// Iterar sobre as linhas e criar um Curso para cada conjunto de dados
	for _, line := range lines {
		if len(line) == 0 {
			continue // ignorar linhas em branco
		}
		fields := strings.Fields(line)
		curso := Curso{
			IdCurso:             parseIdCurso(fields[1]),
			Titulo:              parseTitulo(fields[3]),
			ConteudoProgramatico: parseConteudoProgramatico(fields[5:]),
			Horas:               parseHoras(fields),
			Regime:              parseRegime(fields),
		}
		cursos = append(cursos, curso)
	}

	// Exibir os cursos lidos
	for _, curso := range cursos {
		fmt.Printf("Curso: %+v\n", curso)
	}

	/*sliceType := reflect.TypeOf(slice)
	sliceValue := reflect.ValueOf(slice)

	if sliceType.Kind() != reflect.Slice {
		panic("Input is not a slice")
	}

	// Cria um mapa para armazenar os elementos únicos
	encountered := make(map[interface{}]bool)

	// Slice para armazenar os elementos únicos
	resultSlice := reflect.MakeSlice(sliceType, 0, sliceValue.Len())

	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i).Interface()
		if !encountered[element] {
			encountered[element] = true
			resultSlice = reflect.Append(resultSlice, reflect.ValueOf(element))
		}
	}

	return resultSlice.Interface()
}*/
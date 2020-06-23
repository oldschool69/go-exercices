package main

import "fmt"

func main() {
	states := make([]string, 27, 27)
	states = []string{
		"Acre (AC)",
		"Alagoas (AL)",
		"Amapá (AP)",
		"Amazonas (AM)",
		"Bahia (BA)",
		"Ceará (CE)",
		"Distrito Federal (DF)",
		"Espírito Santo (ES)",
		"Goiás (GO)",
		"Maranhão (MA)",
		"Mato Grosso (MT)",
		"Mato Grosso do Sul (MS)",
		"Minas Gerais (MG)",
		"Pará (PA)",
		"Paraíba (PB)",
		"Paraná (PR)",
		"Pernambuco (PE)",
		"Piauí (PI)",
		"Rio de Janeiro (RJ)",
		"Rio Grande do Norte (RN)",
		"Rio Grande do Sul (RS)",
		"Rondônia (RO)",
		"Roraima (RR)",
		"Santa Catarina (SC)",
		"São Paulo (SP)",
		"Sergipe (SE)",
		"Tocantins (TO)",
	}

	fmt.Println("length: ", len(states))
	fmt.Println("capacity: ", cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Printf("from index %v value is %v\n", i, states[i])
	}
}

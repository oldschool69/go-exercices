// Esse pacote fornece funções muito úteis e complexas
package carecao

// Função utilizada para cálcular a média de multiplos números
// fornecidos como entrada
func Average(nums ...float64) float64 {
	var sum float64
	for _,v := range nums {
		sum += v
	}
	return sum / float64(len(nums))
}
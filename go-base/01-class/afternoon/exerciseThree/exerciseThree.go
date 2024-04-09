package exerciseThree

func ExerciseThree(monthNumber uint) string {
	months := [12]string{
		"Janeiro",
		"Fevereiro",
		"Março",
		"Abril",
		"Maio",
		"Junho",
		"Julho",
		"Agosto",
		"Setembro",
		"Outubro",
		"Novembro",
		"Dezembro",
	}

	if monthNumber >= 1 && monthNumber <= 12 {
		return months[monthNumber-1]
	} else {
		return "Mês inválido"
	}
}

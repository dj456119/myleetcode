package myleetcode

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	result := 0
	for i := range energy {
		if initialEnergy > energy[i] && initialExperience > experience[i] {
		} else if initialEnergy > energy[i] && initialExperience <= experience[i] {
			cha := experience[i] - initialExperience + 1
			initialExperience = initialExperience + cha
			result = result + cha
		} else if initialEnergy <= energy[i] && initialExperience > experience[i] {
			cha := energy[i] - initialEnergy + 1
			initialEnergy = initialEnergy + cha
			result = result + cha
		} else {
			cha := experience[i] - initialExperience + 1
			initialExperience = initialExperience + cha
			result = result + cha
			cha = energy[i] - initialEnergy + 1
			initialEnergy = initialEnergy + cha
			result = result + cha
		}
		initialEnergy = initialEnergy - energy[i]
		initialExperience = initialExperience + experience[i]
		//  fmt.Println(initialEnergy, initialExperience)
	}
	return result
}

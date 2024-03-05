package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium); i++ {
		for j := 0; j < len(podium); j++ {
			if len(podium[i]) < len(podium[j]) {
				podium[i], podium[j] = podium[j], podium[i]
			}
		}
	}
	return podium
}

package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium)-1; i++ {
		if podium[i][0] > podium[i+1][0] {
			podium[i], podium[i+1] = podium[i+1], podium[i]
			i = -1
		}
	}
	return podium
}

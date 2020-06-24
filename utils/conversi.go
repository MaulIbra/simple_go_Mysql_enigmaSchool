package utils

func KonversiNilai(nilai int) string {
	if nilai > 90 {
		return "A"
	} else if nilai > 80 && nilai <= 90 {
		return "B"
	} else if nilai > 70 && nilai <= 80 {
		return "C"
	} else if nilai > 60 && nilai <= 60 {
		return "D"
	} else {
		return "E"
	}
}

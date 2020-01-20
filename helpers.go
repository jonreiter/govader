package govader

import "strings"

func inStringSlice(slice []string, theString string) bool {
	if firstIndexOfStringInSlice(slice, theString) == -1 {
		return false
	}
	return true
}

func inStringMap(theMap map[string]float64, theString string) bool {
	if _, ok := theMap[theString]; ok {
		return true
	}
	return false
}

func inStringStringMap(theMap map[string]string, theString string) bool {
	if _, ok := theMap[theString]; ok {
		return true
	}
	return false
}

func firstIndexOfStringInSlice(slice []string, toFind string) int {
	for i, v := range slice {
		if v == toFind {
			return i
		}
	}
	return -1
}

func firstIndexOfFloatInSlice(slice []float64, toFind float64) int {
	for i, v := range slice {
		if v == toFind {
			return i
		}
	}
	return -1
}

func stringSliceToLower(stringsList []string) []string {
	newStrings := make([]string, len(stringsList))
	for i, v := range stringsList {
		newStrings[i] = strings.ToLower(v)
	}
	return newStrings
}

// eof

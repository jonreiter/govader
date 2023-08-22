package govader

// (empirically derived mean sentiment intensity rating increase for booster words)
const bINCR = 0.293
const bDECR = -0.293

// (empirically derived mean sentiment intensity rating increase for using ALLCAPs to emphasize a word)
const cINCR = 0.733
const nSCALAR = -0.74

const alphaDefault = 15

const valenceScalarScale1 = 0.95
const valenceScalarScale2 = 0.9

const epAmplifyScale = 0.292
const qmAmplifyScale = 0.18

const maxEP = 4
const maxQM = 0.96

const negationScale = 1.25
const butScale = 0.5

// TermConstants contains the large list and dictionary constants
type TermConstants struct {
	NegateList        []string
	BoosterDict       map[string]float64
	LadenIdioms       map[string]float64
	SpecialCaseIdioms map[string]float64
	Regex             *PythonesqueRegex
}

// NewTermConstants assembles a fully-populated constant struct
func NewTermConstants() *TermConstants {
	var tc TermConstants
	tc.NegateList = negateList()
	tc.BoosterDict = boosterDict()
	tc.LadenIdioms = sentimentLadenIdioms()
	tc.SpecialCaseIdioms = specialCaseIdioms()
	tc.Regex = NewPythonesqueRegex()
	return &tc
}

func negateList() []string {
	return []string{
		"nemam", "ne mogu", "ne usuđujem se", "neću", "ne može", "nije", "nisu", "možda ne",
		"ne sme", "ne treba", "nikada", "ne", "ni", "nigde", "ne bi trebalo",
		"uh-uh", "uhuh", "nije bilo", "ne bi", "nije bilo", "bez", "ne bih", "retko",
		"nijedno", "nijedna", "nijedan", "ničiji"}
}

func boosterDict() map[string]float64 {
	return map[string]float64{
		"apsolutno": bINCR, "neverovatno": bINCR, "strahovito": bINCR,
		"u potpunosti": bINCR, "prilično": bINCR, "znatno": bINCR,
		"odlučno": bINCR, "duboko": bINCR, "jako": bINCR, "ogroman": bINCR, "ogromno": bINCR,
		"potpuno": bINCR, "posebno": bINCR, "izuzetan": bINCR, "izuzetno": bINCR,
		"ekstremno": bINCR, "vrlo": bINCR, "najviše": bINCR,
		"fantastično": bINCR, "prokletstvo": bINCR,
		"jebeno": bINCR, "jebeni": bINCR, "veoma": bINCR,
		"u velikoj meri": bINCR, "visoko": bINCR, "grdno": bINCR,
		"neverovatan": bINCR, "intenzivno": bINCR, "žestoko": bINCR, "silno": bINCR,
		"glavni": bINCR, "uglavnom": bINCR, "više": bINCR, "većina": bINCR, "naročito": bINCR,
		"čisto": bINCR, "zaista": bINCR, "upadljivo": bINCR,
		"tako": bINCR, "suštinski": bINCR,
		"temeljno": bINCR, "ukupno": bINCR, "strašno": bINCR,
		"neobično": bINCR, "izgovoriti": bINCR,
		"skoro": bDECR, "jedva": bDECR, "grubo": bDECR, "dovoljno": bDECR,
		"kao": bDECR, "nekako": bDECR, "malo": bDECR,
		"u manjem stepenu": bDECR, "mali": bDECR, "marginalni": bDECR, "marginalno": bDECR,
		"povremen": bDECR, "povremeno": bDECR, "delimično": bDECR,
		"oskudna": bDECR, "neznatno": bDECR, "neznatan": bDECR, "donekle": bDECR}
}

func sentimentLadenIdioms() map[string]float64 {
	return map[string]float64{
		"iseći senf": 2, "ruka na usta": -2,
		"ruke iza leđa": -2, "duvati dim": -2, "duvanje dima": -2,
		"ruke u vis": 1, "slomi nogu": 2,
		"kuvanje na gas": 2, "u crnom": 2, "u dugovima": -2,
		"na lopti": 2, "u novčanoj neprilici": -2}
}

func specialCaseIdioms() map[string]float64 {
	return map[string]float64{"sranje": 3, "bomba": 3, "seronja": 1.5,
		"kako da ne": -2, "poljubac smrti": -1.5, "umreti za": 3}
}

/*
// this is unused in the original code, leaving here for consistency
func (tc *TermConstants) sentimentLadenIdiomsCheck(valence float64, sentiTextLower string) float64 {
	idiomsValences := make([]float64, 0)
	for idiom := range tc.LadenIdioms {
		if strings.Contains(sentiTextLower, idiom) {
			thisValence := tc.LadenIdioms[idiom]
			idiomsValences = append(idiomsValences, thisValence)
		}
	}
	if len(idiomsValences) > 0 {
		tot := mat.Sum(mat.NewVecDense(len(idiomsValences), idiomsValences))
		return tot / float64(len(idiomsValences))
	}
	return valence
}
*/

func (tc *TermConstants) specialIdiomsCheck(valence float64, wordsAndEmoticonsLower []string, i int, boosterDict map[string]float64) float64 {
	newValence := valence

	onezero := wordsAndEmoticonsLower[i-1] + " " + wordsAndEmoticonsLower[i-0]
	twoonezero := wordsAndEmoticonsLower[i-2] + " " + wordsAndEmoticonsLower[i-1] + " " + wordsAndEmoticonsLower[i-0]
	twoone := wordsAndEmoticonsLower[i-2] + " " + wordsAndEmoticonsLower[i-1]
	threetwoone := wordsAndEmoticonsLower[i-3] + " " + wordsAndEmoticonsLower[i-2] + " " + wordsAndEmoticonsLower[i-1]
	threetwo := wordsAndEmoticonsLower[i-3] + " " + wordsAndEmoticonsLower[i-2]

	sequences := []string{onezero, twoonezero, twoone, threetwoone, threetwo}

	for _, v := range sequences {
		if inStringMap(tc.SpecialCaseIdioms, v) {
			newValence = tc.SpecialCaseIdioms[v]
			break
		}
	}

	if len(wordsAndEmoticonsLower)-1 > i {
		zeroone := wordsAndEmoticonsLower[i+0] + " " + wordsAndEmoticonsLower[i+1]
		if inStringMap(tc.SpecialCaseIdioms, zeroone) {
			newValence = tc.SpecialCaseIdioms[zeroone]
		}
	}

	if len(wordsAndEmoticonsLower)-1 > (i + 1) {
		zeroonetwo := wordsAndEmoticonsLower[i+0] + " " + wordsAndEmoticonsLower[i+1] + " " + wordsAndEmoticonsLower[i+2]
		if inStringMap(tc.SpecialCaseIdioms, zeroonetwo) {
			newValence = tc.SpecialCaseIdioms[zeroonetwo]
		}
	}

	nGrams := []string{threetwoone, threetwo, twoone}
	for _, nGram := range nGrams {
		if inStringMap(boosterDict, nGram) {
			newValence = newValence + boosterDict[nGram]
		}
	}

	return newValence
}

// eof

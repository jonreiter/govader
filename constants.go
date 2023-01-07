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
		"aint", "arent", "cannot", "cant", "couldnt", "darent", "didnt", "doesnt",
		"ain't", "aren't", "can't", "couldn't", "daren't", "didn't", "doesn't",
		"dont", "hadnt", "hasnt", "havent", "isnt", "mightnt", "mustnt", "neither",
		"don't", "hadn't", "hasn't", "haven't", "isn't", "mightn't", "mustn't",
		"neednt", "needn't", "never", "none", "nope", "nor", "not", "nothing", "nowhere",
		"oughtnt", "shant", "shouldnt", "uhuh", "wasnt", "werent",
		"oughtn't", "shan't", "shouldn't", "uh-uh", "wasn't", "weren't",
		"without", "wont", "wouldnt", "won't", "wouldn't", "rarely", "seldom", "despite"}
}

func boosterDict() map[string]float64 {
	return map[string]float64{
		"absolutely": bINCR, "amazingly": bINCR, "awfully": bINCR,
		"completely": bINCR, "considerable": bINCR, "considerably": bINCR,
		"decidedly": bINCR, "deeply": bINCR, "effing": bINCR, "enormous": bINCR, "enormously": bINCR,
		"entirely": bINCR, "especially": bINCR, "exceptional": bINCR, "exceptionally": bINCR,
		"extreme": bINCR, "extremely": bINCR,
		"fabulously": bINCR, "flipping": bINCR, "flippin": bINCR, "frackin": bINCR, "fracking": bINCR,
		"fricking": bINCR, "frickin": bINCR, "frigging": bINCR, "friggin": bINCR, "fully": bINCR,
		"fuckin": bINCR, "fucking": bINCR, "fuggin": bINCR, "fugging": bINCR,
		"greatly": bINCR, "hella": bINCR, "highly": bINCR, "hugely": bINCR,
		"incredible": bINCR, "incredibly": bINCR, "intensely": bINCR,
		"major": bINCR, "majorly": bINCR, "more": bINCR, "most": bINCR, "particularly": bINCR,
		"purely": bINCR, "quite": bINCR, "really": bINCR, "remarkably": bINCR,
		"so": bINCR, "substantially": bINCR,
		"thoroughly": bINCR, "total": bINCR, "totally": bINCR, "tremendous": bINCR, "tremendously": bINCR,
		"uber": bINCR, "unbelievably": bINCR, "unusually": bINCR, "utter": bINCR, "utterly": bINCR,
		"very":   bINCR,
		"almost": bDECR, "barely": bDECR, "hardly": bDECR, "just enough": bDECR,
		"kind of": bDECR, "kinda": bDECR, "kindof": bDECR, "kind-of": bDECR,
		"less": bDECR, "little": bDECR, "marginal": bDECR, "marginally": bDECR,
		"occasional": bDECR, "occasionally": bDECR, "partly": bDECR,
		"scarce": bDECR, "scarcely": bDECR, "slight": bDECR, "slightly": bDECR, "somewhat": bDECR,
		"sort of": bDECR, "sorta": bDECR, "sortof": bDECR, "sort-of": bDECR}
}

func sentimentLadenIdioms() map[string]float64 {
	return map[string]float64{
		"cut the mustard": 2, "hand to mouth": -2,
		"back handed": -2, "blow smoke": -2, "blowing smoke": -2,
		"upper hand": 1, "break a leg": 2,
		"cooking with gas": 2, "in the black": 2, "in the red": -2,
		"on the ball": 2, "under the weather": -2}
}

func specialCaseIdioms() map[string]float64 {
	return map[string]float64{"the shit": 3, "the bomb": 3, "bad ass": 1.5, "badass": 1.5,
		"yeah right": -2, "kiss of death": -1.5, "to die for": 3}
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

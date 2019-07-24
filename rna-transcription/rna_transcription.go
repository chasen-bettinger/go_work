package strand

// ToRNA takes a DNA strand and converts it to its RNA equivalent
func ToRNA(dna string) string {
	newStrand := ""

	rnaMap := map[string]string {
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
	}

	for _, nucleotide := range dna {
		newStrand += rnaMap[string(nucleotide)]

	}

	return newStrand
}

package moderation

var (
    // The threshold InappropriateLevel where a phrase is considered inappropriate
    InappropriateThreshold int = 1
)

// Analysis contains the result of Analyze()-ing a phrase
type Analysis struct {
    // The total number of inappropriate words, if each word was level 1.
    // Otherwise, the total inappropriate level of all words.
    InappropriateLevel int
}

// IsInappropriate returns true if the analyzed phrase is inappropriate relative to InappropriateThreshold
func (analysis Analysis) IsInappropriate() bool {
    return analysis.InappropriateLevel >= InappropriateThreshold
}

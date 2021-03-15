package db


type Iteration struct {
	iterId          int
	iterCreatorId   string
	iterType        string
	iterAdmin       []string
	iterState       []string
	iterBranch      string
	iterDevActGroup int
	iterPreActGroup int
	iterItgActGroup int
	application     string
	iterDevClc      float64
	iterItgClc      float64
	iterPreClc      float64
	iterDevQs       float64
	iterItgQs       float64
	iterPreQs       float64
	devPr           []int
	itgPr           []int
	prePr           []int
}

func (iter Iteration) insert() bool{
	return false
}
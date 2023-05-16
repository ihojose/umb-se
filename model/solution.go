package model

type Solution struct {
	ID          uint                  `json:"id"`
	Description string                `json:"description"`
	Efficacy    float32               `json:"efficacy"`
	ProblemID   string                `json:"problem_id"`
	HasSolution []SessionHashSolution `json:"has_solution"`
	Problem     Problem
}

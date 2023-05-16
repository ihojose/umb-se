package model

type SessionHashSolution struct {
	SessionID  uint  `json:"session_id"`
	SolutionID uint  `json:"solution_id"`
	ItWorked   int32 `json:"it_worked"`
	Session    Session
	Solution   Solution
}

package voters

import (
	"time"
)

type VoterHistory struct{
	PollId uint `json:"poll_id"`
	VoteId uint `json:"vote_id"`
	VoteDate time.Time `json:"vote_data"`
}

type Voter struct {
	VoterId uint `json:"voter_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	VoteHistory []VoterHistory `json:"vote_history"`
}

type VoterList struct {
	Voters [uint]Voter `json:"voters"`
}

func New() (*VoterList, error){
	vl :=
}
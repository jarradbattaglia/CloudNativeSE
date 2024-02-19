package voters

import (
	"fmt"
	"time"
)

type VoterHistory struct{
	PollId int `json:"poll_id"`
	VoteId int `json:"vote_id"`
	VoteDate time.Time `json:"vote_date"`
}

type Voter struct {
	VoterId int `json:"voter_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	VoteHistory []VoterHistory `json:"vote_history"`
}

type VoterList struct {
	Voters map[uint]Voter `json:"voters"`
}

func (vl *VoterList) Populate() error {
	err := vl.AddVoter(Voter{VoterId: 1, 
		Name: "Test Voter", 
		Email: "test@example.com",
		VoteHistory: make([]VoterHistory, 0)})
	if err != nil {
		return err
	}
	err = vl.AddVoter(Voter{VoterId: 2, 
			Name: "Jarrad Tester", 
			Email: "jarrad@example.com",
			VoteHistory: make([]VoterHistory, 0)})
	if err != nil {
		return err
	}
	return nil
}

func New() (*VoterList, error){
	voterList := &VoterList{ Voters: make(map[uint]Voter)}
	return voterList, nil
}

func (vl *VoterList) GetAllVoters() ([]Voter, error) {
	var voterList []Voter
	for _, voter := range vl.Voters {
		voterList = append(voterList, voter)
	}
	return voterList, nil
}

func (vl *VoterList) GetVoter(voterId int) (Voter, error) {
	if val, ok := vl.Voters[uint(voterId)]; ok {
		return val, nil
	}
	return Voter{}, fmt.Errorf("VoterId %v Not Found", voterId)
}

func (vl *VoterList) AddVoter(v Voter) error {
	if _, ok := vl.Voters[uint(v.VoterId)]; ok {
		return fmt.Errorf("Voter %v already exists", v.VoterId)
	}
	
	vl.Voters[uint(v.VoterId)] = v
	return nil
}

func (vl *VoterList) GetPollsByVoterId(voterId int) ([]VoterHistory, error) {
	if _, ok := vl.Voters[uint(voterId)]; ok {
		return vl.Voters[uint(voterId)].VoteHistory, nil
	}
	return []VoterHistory{}, fmt.Errorf("VoterId %v does not exist", voterId)
}

func (vl *VoterList) AddPollForVoter(voterId int, voterHistory VoterHistory) error {
	if val, ok := vl.Voters[uint(voterId)]; ok {
		_, err := vl.GetPollsByPollId(voterId, voterHistory.PollId)
		if err == nil {
			return fmt.Errorf("poll %v already exists on voterId %v", voterHistory.PollId, voterId)
		}
		val.VoteHistory = append(val.VoteHistory, voterHistory)
		vl.Voters[uint(voterId)] = val
		return nil
	}
	return fmt.Errorf("could not find voterid %v", voterId)
}

func (vl *VoterList) GetPollsByPollId(voterId int, pollId int) (VoterHistory, error) {
	if val, ok := vl.Voters[uint(voterId)]; ok {
		for _, voteHistVal := range val.VoteHistory {
			if voteHistVal.PollId == pollId {
				return voteHistVal, nil
			}
		}
		return VoterHistory{}, fmt.Errorf("PollID %v does not exist in VoterId %v", pollId, voterId)
	}
	return VoterHistory{}, fmt.Errorf("VoterId %v does not exist", voterId)
}

// Delete voterId from voterlist
func (vl *VoterList) DeleteVoter(voterId int) error {
	if _, ok := vl.Voters[uint(voterId)]; ok {
		delete(vl.Voters, uint(voterId))
		return nil
	}
	return fmt.Errorf("could not find voterId %v", voterId)
}

func (vl *VoterList) DeletePollForVoter(voterId int, pollId int) error {
	if val, ok := vl.Voters[uint(voterId)]; ok {
		for iPoll, valPoll := range val.VoteHistory {
			if valPoll.PollId == pollId {
				val.VoteHistory = append(val.VoteHistory[:iPoll], val.VoteHistory[iPoll+1:]...)
				vl.Voters[uint(voterId)] = val
				return nil
			}
		}
		return fmt.Errorf("could not find pollId %v for voterId %v", pollId, voterId)
	}
	return fmt.Errorf("could not find voterId %v", voterId)
}

func (vl *VoterList) UpdateVoter(voter Voter) error {
	if _, ok := vl.Voters[uint(voter.VoterId)]; ok {
		vl.Voters[uint(voter.VoterId)] = voter
		return nil
	}
	return fmt.Errorf("no voter with voterId %v found to update", voter.VoterId)
}

func (vl *VoterList) UpdateVoterPoll(voterId int, voterHistory VoterHistory) error {
	if val, ok := vl.Voters[uint(voterId)]; ok {
		for iPoll, valPoll := range val.VoteHistory {
			if valPoll.PollId == voterHistory.PollId {
				vl.Voters[uint(voterId)].VoteHistory[iPoll] = voterHistory
				return nil
			}
		}
		return fmt.Errorf("could not find pollId %v for voterId %v", voterHistory.PollId, voterId)
	}
	return fmt.Errorf("could not find voterId %v", voterId)
}
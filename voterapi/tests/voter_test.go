package tests

import (
	"fmt"
	"os"
	"testing"
	"time"

	"drexel.edu/voterapi/voters"
	fake "github.com/brianvoe/gofakeit/v6" //aliasing package name
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)	

var (
	BASE_API = "http://localhost:8080"

	cli = resty.New()
)

func TestMain(m *testing.M) {
	//Delete and ignore errors, make sure all reasonable calls will come back
	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("%v/voters/%v", BASE_API, i)
		cli.R().Delete(url)
	}
	code := m.Run()

	//CLEANUP
	for i := 0; i < 10; i++ {
		url := fmt.Sprintf("%v/voters/%v", BASE_API, i)
		cli.R().Delete(url)
	}
	//Now Exit
	os.Exit(code)	
}

func newRandVoter(id int) voters.Voter {
	return voters.Voter{
		VoterId: id,
		Name: fake.Name(),
		Email: fake.Email(),
		VoteHistory: []voters.VoterHistory{
			{
				PollId: 0,
				VoteId: fake.NewCrypto().Rand.Int(),
				VoteDate: fake.Date(),
			},
		},
	}
}

func Test_AddVoters(t *testing.T) {
	numLoad := 5

	for i := 0; i < numLoad; i++ {
		voter := newRandVoter(i)
		url := fmt.Sprintf("%v/voters/%v", BASE_API, i)
		rsp, err := cli.R().SetBody(voter).Post(url)
		assert.Nil(t, err)
		assert.Equal(t, 200, rsp.StatusCode())
	}
}

func Test_GetAllVoters(t *testing.T) {
	var voters []voters.Voter
	url := fmt.Sprintf("%v/voters/", BASE_API)
	rsp, err := cli.R().SetResult(&voters).Get(url)

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())
	assert.Equal(t, 5, len(voters))
}

func Test_AddSpecificVoter(t *testing.T) {
	voter := voters.Voter{VoterId: 6, 
		Name: "Test Voter", 
		Email: "test@example.com",
		VoteHistory: []voters.VoterHistory{
			{
				PollId: 7,
				VoteId: 5,
				VoteDate: time.Now(),
			},
		},
	}
	url := fmt.Sprintf("%v/voters/%v", BASE_API, 6)
	rsp, err := cli.R().SetBody(voter).Post(url)

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())
}

func Test_GetVoteHistory(t *testing.T) {
	var voteHistory []voters.VoterHistory
	url := fmt.Sprintf("%v/voters/6/polls/", BASE_API)
	rsp, err := cli.R().SetResult(&voteHistory).Get(url)

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())
	assert.Equal(t, 1, len(voteHistory))
	assert.Equal(t, 7, voteHistory[0].PollId)
}


func Test_GetSpecificVoter(t *testing.T) {
	var voter voters.Voter
	url := fmt.Sprintf("%v/voters/%v", BASE_API, 6)
	rsp, err := cli.R().SetResult(&voter).Get(url)

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())
	assert.Equal(t, 6, voter.VoterId)
	assert.Equal(t, "Test Voter", voter.Name)
	assert.Equal(t, 1, len(voter.VoteHistory))
}

func Test_GetSpecificVoterOnePoll(t *testing.T) {
	var voteHistory voters.VoterHistory
	url := fmt.Sprintf("%v/voters/%v/polls/%v", BASE_API, 6, 7)
	rsp, err := cli.R().SetResult(&voteHistory).Get(url)

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())
	assert.Equal(t, 7, voteHistory.PollId)
}


func Test_DeleteVoterHistory(t *testing.T) {
	url := fmt.Sprintf("%v/voters/%v/polls/%v", BASE_API, 6, 7)
	rsp, err := cli.R().Delete(url)

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())

	var voter voters.Voter
	url = fmt.Sprintf("%v/voters/%v", BASE_API, 6)
	rsp, err = cli.R().SetResult(&voter).Get(url)
	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())
	assert.Equal(t, 0, len(voter.VoteHistory))
}
func Test_DeleteVoter(t *testing.T) {
	url := fmt.Sprintf("%v/voters/%v", BASE_API, 6)
	rsp, err := cli.R().Delete(url)

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())
}
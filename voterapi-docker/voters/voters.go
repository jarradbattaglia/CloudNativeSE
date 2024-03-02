package voters

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
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

const (
	RedisNilError        = "redis: nil"
	RedisDefaultLocation = "0.0.0.0:6379"
	RedisKeyPrefix       = "voter:"
)

type cache struct {
	cacheClient *redis.Client
	jsonHelper *rejson.Handler
	context context.Context
}

type VoterList struct {
	cache
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
	redisUrl := os.Getenv("REDIS_URL")
	//This handles the default condition
	if redisUrl == "" {
		redisUrl = RedisDefaultLocation
	}
	log.Println("DEBUG:  USING REDIS URL: " + redisUrl)
	return NewWithCacheInstance(redisUrl)
}

func NewWithCacheInstance(location string) (*VoterList, error) {
	//Connect to redis.  Other options can be provided, but the
	//defaults are OK
	client := redis.NewClient(&redis.Options{
		Addr: location,
	})

	//We use this context to coordinate betwen our go code and
	//the redis operaitons
	ctx := context.Background()

	//This is the reccomended way to ensure that our redis connection
	//is working
	err := client.Ping(ctx).Err()
	if err != nil {
		log.Println("Error connecting to redis" + err.Error() + "cache might not be available, continuing...")
	}

	//By default, redis manages keys and values, where the values
	//are either strings, sets, maps, etc.  Redis has an extension
	//module called ReJSON that allows us to store JSON objects
	//however, we need a companion library in order to work with it
	//Below we create an instance of the JSON helper and associate
	//it with our redis connnection
	jsonHelper := rejson.NewReJSONHandler()
	jsonHelper.SetGoRedisClientWithContext(ctx, client)

	//Return a pointer to a new VoterList struct
	return &VoterList{
		cache: cache{
			cacheClient: client,
			jsonHelper:  jsonHelper,
			context:     ctx,
		},
	}, nil
}

// We will use this later, you can ignore for now
func isRedisNilError(err error) bool {
	return errors.Is(err, redis.Nil) || err.Error() == RedisNilError
}

// In redis, our keys will be strings, they will look like
// todo:<number>.  This function will take an integer and
// return a string that can be used as a key in redis
func redisKeyFromId(id int) string {
	return fmt.Sprintf("%s%d", RedisKeyPrefix, id)
}

func (t *VoterList) getVoterFromRedis(key string, item *Voter) error {

	//Lets query redis for the item, note we can return parts of the
	//json structure, the second parameter "." means return the entire
	//json structure
	itemObject, err := t.jsonHelper.JSONGet(key, ".")
	if err != nil {
		return err
	}

	//JSONGet returns an "any" object, or empty interface,
	//we need to convert it to a byte array, which is the
	//underlying type of the object, then we can unmarshal
	//it into our ToDoItem struct
	err = json.Unmarshal(itemObject.([]byte), item)
	if err != nil {
		return err
	}

	return nil
}

func (t *VoterList) getVoterHistoryFromRedis(key string, item *[]VoterHistory) error {

	//Lets query redis for the item, note we can return parts of the
	//json structure, the second parameter "." means return the entire
	//json structure
	itemObject, err := t.jsonHelper.JSONGet(key, ".vote_history.")
	if err != nil {
		return err
	}

	//JSONGet returns an "any" object, or empty interface,
	//we need to convert it to a byte array, which is the
	//underlying type of the object, then we can unmarshal
	//it into our ToDoItem struct
	err = json.Unmarshal(itemObject.([]byte), item)
	if err != nil {
		return err
	}

	return nil
}




func (vl *VoterList) GetAllVoters() ([]Voter, error) {
	var voterList []Voter
	var voter Voter
	pattern := RedisKeyPrefix + "*"
	ks, _ := vl.cacheClient.Keys(vl.context, pattern).Result()
	for _, key := range ks {
		err := vl.getVoterFromRedis(key, &voter)
		if err != nil {
			return nil, err
		}
		voterList = append(voterList, voter)
	}
	return voterList, nil
}

func (vl *VoterList) GetVoter(voterId int) (Voter, error) {
	var voter Voter
	pattern := redisKeyFromId(voterId)
	err := vl.getVoterFromRedis(pattern, &voter)
	if err != nil {
		return Voter{}, fmt.Errorf("VoterId %v Not Found", voterId)
	}
	return voter, nil
}

func (vl *VoterList) AddVoter(v Voter) error {
	pattern := redisKeyFromId(v.VoterId)
	var existingVoter Voter
	if err := vl.getVoterFromRedis(pattern, &existingVoter); err == nil {
		return errors.New("voter already exists")
	}

	if _, err := vl.jsonHelper.JSONSet(pattern, ".", v); err != nil {
		return err
	}
	return nil
}

func (vl *VoterList) GetPollsByVoterId(voterId int) ([]VoterHistory, error) {
	voter, err := vl.GetVoter(voterId)
	if err != nil {
		return []VoterHistory{}, fmt.Errorf("VoterId %v does not exist", voterId)
	}
	return voter.VoteHistory, nil

}

func (vl *VoterList) GetPollsByPollId(voterId int, pollId int) (VoterHistory, error) {

	voter, err := vl.GetVoter(voterId)
	if err != nil {
		return VoterHistory{}, fmt.Errorf("VoterId %v does not exist", voterId)
	}
	for _, voteHistVal := range voter.VoteHistory {
		if voteHistVal.PollId == pollId {
			return voteHistVal, nil
		}
	}
	return VoterHistory{}, fmt.Errorf("PollID %v does not exist in VoterId %v", pollId, voterId)
}

func (vl *VoterList) AddPollForVoter(voterId int, voterHistory VoterHistory) error {
	_, err := vl.GetPollsByPollId(voterId, voterHistory.PollId)
	if err == nil {
		return fmt.Errorf("poll %v already exists on voterId %v", voterHistory.PollId, voterId)
	}
	voter, err := vl.GetVoter(voterId)
	if err != nil {
		return err
	}
	voter.VoteHistory = append(voter.VoteHistory, voterHistory)
	redisKey := redisKeyFromId(voterId)	
	if _, err := vl.jsonHelper.JSONSet(redisKey, ".", voter); err != nil {
		return err
	}
	return nil
	
}



// Delete voterId from voterlist
func (vl *VoterList) DeleteVoter(voterId int) error {
	pattern := redisKeyFromId(voterId)
	numDeleted, err := vl.cacheClient.Del(vl.context, pattern).Result()
	if err != nil {
		return err
	}
	if numDeleted == 0 {
		return errors.New("attempted to delete non-existant item")
	}
	return nil
}

func (vl *VoterList) DeletePollForVoter(voterId int, pollId int) error {

	pattern := redisKeyFromId(voterId)
	polls, err := vl.GetPollsByVoterId(voterId)
	if err != nil {
		return err
	}
	for index, voterHistory := range polls {
		if voterHistory.PollId == pollId {
			voteHistoryPattern := fmt.Sprintf(".vote_history[%d]", index)
			numDeleted, err := vl.cacheClient.JSONDel(vl.context, pattern, voteHistoryPattern).Result()
			if err != nil {
				return err
			}
			if numDeleted == 0 {
				return errors.New("attempted to delete non-existant item")
			}
			return nil
		}
	}
	return errors.New("attempted to delete non-existant poll")
}

func (vl *VoterList) UpdateVoter(voter Voter) error {
	pattern := redisKeyFromId(voter.VoterId)
	var existingVoter Voter
	if err := vl.getVoterFromRedis(pattern, &existingVoter); err != nil {
		return errors.New("voter does not exist to update")
	}

	if _, err := vl.jsonHelper.JSONSet(pattern, ".", voter); err != nil {
		return err
	}
	return nil
}

func (vl *VoterList) UpdateVoterPoll(voterId int, voterHistory VoterHistory) error {
	pattern := redisKeyFromId(voterId)
	var existingVoter Voter
	if err := vl.getVoterFromRedis(pattern, &existingVoter); err != nil {
		return errors.New("voter does not exist to update")
	}

	for iPoll, valPoll := range existingVoter.VoteHistory {
		if valPoll.PollId == voterHistory.PollId {
			existingVoter.VoteHistory[iPoll] = voterHistory				
			if _, err := vl.jsonHelper.JSONSet(pattern, ".", existingVoter); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("could not find pollId %v for voterId %v", voterHistory.PollId, voterId)
}
package api

import (
	"fmt"
	"log"
	"net/http"

	"drexel.edu/voterapi/voters"
	"github.com/gofiber/fiber/v2"
)

// The api package creates and maintains a reference to the data handler
// this is a good design practice
type VoterAPI struct {
	voterList *voters.VoterList
}

func New() (*VoterAPI, error) {
	voterHandler, err := voters.New()
	if err != nil {
		return nil, err
	}

	return &VoterAPI{voterList: voterHandler}, nil
}


// implementation for GET /voters
// returns all voters
func (voterAPI *VoterAPI) GetAllVoters(c *fiber.Ctx) error {

	voterList, err := voterAPI.voterList.GetAllVoters()
	if err != nil {
		log.Println("Error Getting All Voters: ", err)
		return fiber.NewError(http.StatusNotFound,
			"Error Getting All Voters")
	}
	//Note that the database returns a nil slice if there are no items
	//in the database.  We need to convert this to an empty slice
	//so that the JSON marshalling works correctly.  We want to return
	//an empty slice, not a nil slice. This will result in the json being []
	if voterList == nil {
		voterList = make([]voters.Voter, 0)
	}
	return c.JSON(voterList)
}


// implementation for GET /voters/:id
// returns a single todo
func (vAPI *VoterAPI) GetVoterById(c *fiber.Ctx) error {

	//Note go is minimalistic, so we have to get the
	//id parameter using the Param() function, and then
	//convert it to an int64 using the strconv package
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	//Note that ParseInt always returns an int64, so we have to
	//convert it to an int before we can use it.
	voterItem, err := vAPI.voterList.GetVoter(id)
	if err != nil {
		log.Println("Voter not found: ", err)
		return fiber.NewError(http.StatusNotFound)
	}

	//Git will automatically convert the struct to JSON
	//and set the content-type header to application/json
	return c.JSON(voterItem)
}

// implementation of GET /health. It is a good practice to build in a
// health check for your API.  Below the results are just hard coded
// but in a real API you can provide detailed information about the
// health of your API with a Health Check
func (vAPI *VoterAPI) HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"status":             "ok",
			"version":            "1.0.0",
			"uptime":             100,
			"users_processed":    1000,
			"errors_encountered": 10,
		})
}


// implementation for POST /voters/:id
// adds a new voter to the list if it does not exist
func (vAPI *VoterAPI) AddVoter(c *fiber.Ctx) error {
	var voter voters.Voter
	if err := c.BodyParser(&voter); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}
	if err := vAPI.voterList.AddVoter(voter); err != nil {
		log.Println("Error adding item: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}
	return c.JSON(voter)
}

func (vAPI *VoterAPI) GetPollsByVoterId(c *fiber.Ctx) error {
	voterId, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	polls, err := vAPI.voterList.GetPollsByVoterId(voterId)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.JSON(polls)
}

func (vAPI *VoterAPI) AddPollForVoter(c *fiber.Ctx) error {
	var voterHistory voters.VoterHistory
	if err := c.BodyParser(&voterHistory); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}

	voterId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	if err := vAPI.voterList.AddPollForVoter(voterId, voterHistory); err != nil {
		log.Println("Error adding item: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}
	return c.JSON(voterHistory)
}

func (vAPI *VoterAPI) GetPollByPollId(c *fiber.Ctx) (error) {
	voterId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}
	pollId, err := c.ParamsInt("pollid")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	poll, err := vAPI.voterList.GetPollsByPollId(voterId, pollId)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.JSON(poll)
}

func (vAPI *VoterAPI) DeleteVoter(c *fiber.Ctx) error {
	voterId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}
	err = vAPI.voterList.DeleteVoter(voterId)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.SendString(fmt.Sprintf("VoterId %v has been removed", voterId))
}	

func (vAPI *VoterAPI) DeletePollForVoter(c *fiber.Ctx) error {
	voterId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}
	pollId, err := c.ParamsInt("pollid")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	err = vAPI.voterList.DeletePollForVoter(voterId, pollId)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.SendString(fmt.Sprintf("PollId %v has been removed from voter %v", pollId, voterId))
}

func (vAPI *VoterAPI) UpdateVoter(c *fiber.Ctx) error {
	var voter voters.Voter
	if err := c.BodyParser(&voter); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}
	if err := vAPI.voterList.UpdateVoter(voter); err != nil {
		log.Println("Error adding item: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}
	return c.JSON(voter)
}

func (vAPI *VoterAPI) UpdateVoterPoll(c *fiber.Ctx) error {
	var voterHistory voters.VoterHistory
	if err := c.BodyParser(&voterHistory); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}

	voterId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	if err := vAPI.voterList.UpdateVoterPoll(voterId, voterHistory); err != nil {
		log.Println("Error adding item: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}
	return c.JSON(voterHistory)
}

func (vAPI *VoterAPI) Populate(c *fiber.Ctx) error {
	vAPI.voterList.Populate()
	return c.SendString("Successfully Populated")
}
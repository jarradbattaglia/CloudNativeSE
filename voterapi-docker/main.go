package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"drexel.edu/voterapi/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Global variables to hold the command line flags to drive the todo CLI
// application
var (
	hostFlag string
	portFlag uint
)

// processCmdLineFlags parses the command line flags for our CLI
//
func processCmdLineFlags() {

	//Note some networking lingo, some frameworks start the server on localhost
	//this is a local-only interface and is fine for testing but its not accessible
	//from other machines.  To make the server accessible from other machines, we
	//need to listen on an interface, that could be an IP address, but modern
	//cloud servers may have multiple network interfaces for scale.  With TCP/IP
	//the address 0.0.0.0 instructs the network stack to listen on all interfaces
	//We set this up as a flag so that we can overwrite it on the command line if
	//needed
	flag.StringVar(&hostFlag, "h", "0.0.0.0", "Listen on all interfaces")
	flag.UintVar(&portFlag, "p", 8080, "Default Port")

	flag.Parse()
}

// main is the entry point for our todo API application.  It processes
// the command line flags and then uses the db package to perform the
// requested operation
func main() {
	processCmdLineFlags()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())

	apiHandler, err := api.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	app.Get("/voters", apiHandler.GetAllVoters)
	app.Get("/voters/populate", apiHandler.Populate)
	app.Post("/voters/:id<uint>", apiHandler.AddVoter)
	app.Get("/voters/:id<uint>", apiHandler.GetVoterById)
	app.Get("/voters/:id<uint>/polls", apiHandler.GetPollsByVoterId)
	app.Post("/voters/:id<uint>/polls/:pollid<uint>", apiHandler.AddPollForVoter)
	app.Get("/voters/:id<uint>/polls/:pollid<uint>", apiHandler.GetPollByPollId)
	app.Delete("/voters/:id<uint>", apiHandler.DeleteVoter)
	app.Delete("/voters/:id<uint>/polls/:pollid<uint>", apiHandler.DeletePollForVoter)
	app.Put("/voters/:id<uint>", apiHandler.UpdateVoter)
	app.Put("/voters/:id<uint>/polls/:pollid<uint>", apiHandler.UpdateVoterPoll)
	app.Get("/health", apiHandler.HealthCheck)

	serverPath := fmt.Sprintf("%s:%d", hostFlag, portFlag)
	log.Println("Starting server on ", serverPath)
	app.Listen(serverPath)
}
package tests

//Introduction to testing.  Note that testing is built into go and we will be using
//it extensively in this class. Below is a starter for your testing code.  In
//addition to what is built into go, we will be using a few third party packages
//that improve the testing experience.  The first is testify.  This package brings
//asserts to the table, that is much better than directly interacting with the
//testing.T object.  Second is gofakeit.  This package provides a significant number
//of helper functions to generate random data to make testing easier.

import (
	"fmt"
	"os"
	"testing"

	"drexel.edu/todo/db"
	fake "github.com/brianvoe/gofakeit/v6" //aliasing package name
	"github.com/stretchr/testify/assert"
)

// Note the default file path is relative to the test package location.  The
// project has a /tests path where you are at and a /data path where the
// database file sits.  So to get there we need to back up a directory and
// then go into the /data directory.  Thus this is why we are setting the
// default file name to "../data/todo.json"
const (
	DEFAULT_DB_FILE_NAME = "../data/todo.json"
)

var (
	DB *db.ToDo
)

// note init() is a helpful function in golang.  If it exists in a package
// such as we are doing here with the testing package, it will be called
// exactly once.  This is a great place to do setup work for your tests.
func init() {
	//Below we are setting up the gloabal DB variable that we can use in
	//all of our testing functions to make life easier
	testdb, err := db.New(DEFAULT_DB_FILE_NAME)
	if err != nil {
		fmt.Print("ERROR CREATING DB:", err)
		os.Exit(1)
	}

	DB = testdb //setup the global DB variable to support test cases

	//Now lets start with a fresh DB with the sample test data
	testdb.RestoreDB()
}

// Sample Test, will always pass, comparing the second parameter to true, which
// is hard coded as true
func TestTrue(t *testing.T) {
	assert.True(t, true, "True is true!")
}

func TestAddHardCodedItem(t *testing.T) {
	item := db.ToDoItem{
		Id:     999,
		Title:  "This is a test case item",
		IsDone: false,
	}
	t.Log("Testing Adding a Hard Coded Item: ", item)

	//TODO: finish this test, add an item to the database and then
	//check that it was added correctly by looking it back up
	//use assert.NoError() to ensure errors are not returned.
	//explore other useful asserts in the testify package, see
	//https://github.com/stretchr/testify.  Specifically look
	//at things like assert.Equal() and assert.Condition()

	//I will get you started, uncomment the lines below to add to the DB
	//and ensure no errors:
	//---------------------------------------------------------------
	err := DB.AddItem(item)
	assert.NoError(t, err, "Error adding item to DB")

	//TODO: Now finish the test case by looking up the item in the DB
	//and making sure it matches the item that you put in the DB above
	retrievedItem, err := DB.GetItem(item.Id)
	assert.NoError(t, err, "Retrieved item without error")
	assert.Equal(t, retrievedItem.Id,  item.Id, "Item IDs are the same")
	assert.Equal(t, retrievedItem.IsDone,  item.IsDone, "Item IsDone are the same")
	assert.Equal(t, retrievedItem.Title,  item.Title, "Item Titles are the same")
}

func TestAddRandomStructItem(t *testing.T) {
	//You can also use the Stuct() fake function to create a random struct
	//Not going to do anyting
	item := db.ToDoItem{}
	err := fake.Struct(&item)
	t.Log("Testing Adding a Randomly Generated Struct: ", item)

	assert.NoError(t, err, "Created fake item OK")

	//TODO: Complete the test
	addErr := DB.AddItem(item)
	assert.NoError(t, addErr, "Successfully added fake item")

	retrievedItem, err := DB.GetItem(item.Id)
	assert.NoError(t, err, "Retrieved item without error")
	assert.Equal(t, retrievedItem.Id,  item.Id, "Item IDs are the same")
	assert.Equal(t, retrievedItem.IsDone,  item.IsDone, "Item IsDone are the same")
	assert.Equal(t, retrievedItem.Title,  item.Title, "Item Titles are the same")

}

func TestAddRandomItem(t *testing.T) {
	//Lets use the fake helper to create random data for the item
	item := db.ToDoItem{
		Id:     fake.Number(100, 110),
		Title:  fake.JobTitle(),
		IsDone: fake.Bool(),
	}

	t.Log("Testing Adding an Item with Random Fields: ", item)
	
	addErr := DB.AddItem(item)
	assert.NoError(t, addErr, "Successfully added fake item")

	retrievedItem, err := DB.GetItem(item.Id)
	assert.NoError(t, err, "Retrieved random item without error")
	assert.Equal(t, retrievedItem.Id,  item.Id, "Random Item IDs are the same")
	assert.Equal(t, retrievedItem.IsDone,  item.IsDone, "Random Item IsDone are the same")
	assert.Equal(t, retrievedItem.Title,  item.Title, "Random Item Titles are the same")

}

// TODO: Please delete this test from your submission, it does not do anything
// but i wanted to demonstrate how you can starting shelling out your tests
// and then implment them later.  The go testing framework provides a
// Skip() function that just tells the testing framework to skip or ignore
// this test function


//TODO: Create additional tests to showcase the correct operation of your program
//for example getting an item, getting all items, updating items, and so on. Be
//creative here.

func TestGetAllItems(t *testing.T) {
	items, allItemsErr := DB.GetAllItems()
	t.Log("Retrieving all items from database and ensuring it is expected lenght")

	assert.NoError(t, allItemsErr, "Retrieved items successfully")

	assert.NoError(t, allItemsErr, "No errors retrieving all items")
	assert.Equal(t, len(items), 7, "Length of items expected length of 4")
}

func TestRetrieveItem(t *testing.T) {
	retrievedItem, err := DB.GetItem(1)
	t.Log("Testing Retrieving Expected Item: ", retrievedItem)

	assert.NoError(t, err, "Retrieved random item without error")

	assert.Equal(t, retrievedItem.Id,  1, "Item IDs is expected")
	assert.Equal(t, retrievedItem.IsDone,  false, "Item IsDone is expected")
	assert.Equal(t, retrievedItem.Title,  "Learn Go / GoLang", "Item title is expected")

}
func TestUpdateItemToTrue(t *testing.T) {
	item, getItemErr := DB.GetItem(1)
	t.Log("Testing Item Update to IsDone from False to True: ", item)

	assert.NoError(t, getItemErr, "Retrieved item 1 successfully")

	assert.Equal(t, item.IsDone, false, "Item by default is false")
	item.IsDone = true
	DB.UpdateItem(item)
	retrievedItem, getItemAgainErr := DB.GetItem(1)
	assert.NoError(t, getItemAgainErr, "Retrieved Item 1 successfully again")

	assert.Equal(t, retrievedItem.IsDone, true, "Successfully set item 1 to true")
}

func TestChangeItemStatus(t *testing.T) {
	newItem := db.ToDoItem{
		Id:     1000,
		Title:  "TestTitle",
		IsDone: false,
	}
	t.Log("Testing Item Update to IsDone from False to True using function: ", newItem)

	DB.AddItem(newItem)

	item, getItemErr := DB.GetItem(1000)
	assert.NoError(t, getItemErr, "Retrieved item 1000 successfully")
	assert.Equal(t, item.IsDone, false, "Item by default is false as expected")

	DB.ChangeItemDoneStatus(1000, true)
	changedItem, getChangedItemErr := DB.GetItem(1000)
	assert.NoError(t, getChangedItemErr, "Retrieved item 1000 successfully")
	assert.Equal(t, changedItem.IsDone, true, "Item by default is false as expected")

}

func TestDeleteItem(t *testing.T) {
	t.Log("Testing deleting item 1000 from previous step")
	_, err := DB.GetItem(1000)
	assert.NoError(t, err, "Retrieved item successfully")
	DB.DeleteItem(1000)
	_, getItemErr := DB.GetItem(1000)
	assert.Error(t, getItemErr, "Expected to get error as item no longer exists")	
}


# Assignment link

https://github.com/ArchitectingSoftware/Cloud-Native-Software-Engineering/blob/main/assignments/api-design-2.md


# Prompt

Finally, this is what you need to do.

For reference, remember part 1 of this assignment involved reviewing this video - https://www.youtube.com/watch?v=P0a7PwRNLVU.

For this assignment I want you to do some independent research on Hypermedia as the engine of application state (HATEOAS). Here is the wikipedia page - https://en.wikipedia.org/wiki/HATEOAS. You can also use ChatGPT, other google resources, etc. Other keywords that can help you are hypermedia.

After you grasp some of these concepts, please do the following:

Make some proposals to change the APIs above to incorporate hypermedia into the design of our voting application. Note I dont want you creating new APIs (stick to the 3 above), but describe how you would change the API structures to better support a hypermedia driven approach. Note you can add additional fields to support the hypermedia concepts or change existing fields.
Write a brief description of how these APIs would work together a little better using these concepts. Focus on the perspective of the user of your API, how would this better support the interactions between the APIs supporting an application?
Provide at least 2 references that you found helpful during this activity. Note you dont need to formally cite them, just a link would be OK.

# Answer

How I would approach this API using hypermedia links is using links more in the calls as stated in the Google video we've looked at before, but also showing resources that the API can use and move to.

For instance the GET /voters/123 api might return something like:

```
{
  "id": /voters/123
  "FirstName": "name"
  "lastName": "name"
  "vote_history": [list_of_votes_structure (use url for ids of polls) like /polls/456 if they voted here]
  "links": [
      { "ref": "self", "href": "/voters/123",  }
      {"ref": "voter_history", "href" "/voters/123/polls",}
      {"ref": "add_vote", "href": "/votes", } 
      {"ref": "polls", "href": "/polls" }
  ]
}
```

We use the uniform API from the google cloud video to give the client more information everywhere! First the ID needs to print out its own path as thats something we can look up and most hypermedia seems to also have a "self", but also for something like the voters, they may want to get the polls they have voted on and the place where they can vote on a poll.

Next the votes API change would do something similar with its 3 ids that it carries, say i retrieved a list of votes done (`GET /votes`), it might look like the below:

```
[
  { 
    "voteId": "/votes/1"
    "voterId": /voters/123,
    "poll": /polls/2
    "value": 7
    "links": [
      { "rel": "self", "href": /votes/1 }
      { "rel": "voter", "href": /voters/123 }
      { "rel": "poll", "href": /polls/2 }
      {"rel": "register_vote", "href": "/voters/123/polls", "hint": "post"} # would exist if we did a POST
    ]
  }, ... 
]
```

This allows the user to see more information about the relationship of votes, I can find the poll info, voter info, and the link to register the vote with the voter id. Also (if we wanted to expand) we could add a pagination token to GET calls throughout each API, but not POSTs, so we dont need to send back EVERY resource, and can limit or allow the user to go to next page, or add hints to the refernce name (like GET/POST/PATCH) to each link.  This all helps the user use the API intuitevely, but also requires a decent amount of work in backend to know and correlate what the user is doing.

Next up I'll do an example of the Polls API, which really only has POSTs to do (so if I did a GET /polls or POST /polls, i might return:

```
{
  "pollId": /polls/2
  "title": "title",
  "question" :"question",
  "options": [ list of options with an id like /polls/2/option/1234 ]
  "links": [
    { "rel": "self", "href": /polls/2 }
    { 'rel': "get_voters_info", "href": "/voters"}
    { 'rel': "register_vote", "href": "/votes"} 
  ]

}

```

This would tell the user that you can create a vote or get available voters info from each of those endpoints, and maybe can be enhanced with having knowledge of if you just did a GET to do 

Now a lot of this can be improved by reading the database or similar behind the scenes calls to get the state of the application or user in a more traditional programming (to get if voter voted on a poll or not and giving right link, or deleting a poll, etc, which would change some links given in a more complex application).  But at the bare minimum we could provide more info to the user on what calls can be used in relation to what they just called as described above, for instance if I just did a post on a Voter my links might have the self link, a hint to create a vote with POST, and a get of polls that they could vote on.  

References I looked at:

https://www.infoq.com/articles/hypermedia-api-tutorial-part-one/
https://blogs.mulesoft.com/dev-guides/api-design/api-best-practices-hypermedia-part-1/
https://apisyouwonthate.com/blog/common-hypermedia-patterns-with-json-hyper-schema/
https://www.mscharhag.com/api-design/hypermedia-rest

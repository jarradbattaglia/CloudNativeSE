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

For instance the `GET /voters/123` api might return something like:

```
{
  "id": /voters/123
  "FirstName": "name"
  "lastName": "name"
  "vote_history": [list_of_votes_structure (use url for ids of polls) like /polls/456 if they voted here]
  "links": [
      { "ref": "self", "href": "/voters/123",  }
      {"ref": "add_vote", "href" "/voters/123/polls",}
      {"ref": "register_vote", "href": "/votes", } 
      {"ref": "polls", "href": "/polls" }
  ]
}
```

We use the uniform API from the google cloud video to give the client more information about the API at any place, where that info could be used! First the ID needs to print out its own path as thats something we can look up and most hypermedia seems to also have a "self" reference in its links section. But also for something like the voters, they may want to get the polls they have voted on, the place where they can vote on a poll, or add a vote registration (`/votes`). Now also in hypermedia, the state and flow of what the user is doing changes what gets returned, for instance if we deleted a resource, instead of returning self, we could do `add_voter` or something similar (more complex application might do things like this better).

Next the votes API change would do something similar with its 3 ids that it carries, say i retrieved a list of votes registered (`GET /votes`), it might look like the below:

```
[
  { 
    "voteId": "/votes/1"
    "voterId": /voters/123,
    "poll": /polls/2
    "value": 7
    "links": [
      { "rel": "self", "href": /votes/1 }
      { "rel": "voter_info", "href": /voters/123 }
      { "rel": "poll_info", "href": /polls/2 }
      { "rel": "voter_history", "href": /voters/123/polls/2 } # maybe shown if vote is already done?/or did a get
      {"rel": "register_vote", "href": "/voters/123/polls", "hint": "post"} # would exist if we did a POST
    ]
  }, ... 
]
```

This allows the user to see more information about the relationship of votes, I can find the poll info, voter info, and the link to register the vote with the voter id. Also (if we wanted to expand) we could add a pagination token to GET calls throughout each API that does a GET all call, so we dont need to send back EVERY resource, and can limit or allow the user to go to next page.  Also we could add hints to the reference name (like GET/POST/PATCH) to each link and the link reference could be more descriptive like add_vote or get_votes, etc.  Things like this all helps the user use the API intuitevely, but also requires a decent amount of work in backend to know and correlate what the user is doing.  

Next up I'll do an example of the Polls API, so if I did a GET /polls or POST /polls, i might return:

```
{
  "pollId": /polls/2
  "title": "title",
  "question" :"question",
  "options": [ list of options with an id like /polls/2/option/1234 ]
  "links": [
    { "rel": "self", "href": /polls/2 }
    { 'rel': "get_voters_info", "href": "/voters"}
    { 'rel': "get_voted_polls", "href": "/votes?polls=<poll_id>"}
    { 'rel': "register_vote", "href": "/votes"} 
  ]

}

```

This would tell the user that you can create a vote or get available voters info from each of those endpoints.  In a further use of this application doing something like implementing a search function would be able to enhance many of these APIs.  For instance being able to do something like `votes/pollId=<pollId>` to find votes that have used this pollId, would greatly benefit an application like this and the other APIs.

In summary:

I think the features that benefit the most for changing the API, is adding the API url to each Id (or where usable), adding the ability to do a search in GET calls, and obviously adding helpful links to calls to help guide the user (even if mostly static).

The flow of the application would ideally be the same as described in the prompt, but the user wouldn't need to keep checking a Swagger or documentation.  If I call POST to `/voters`, I would create my voter and get a link to create a poll perhaps or create a vote, if I decided to do one of them, for instance i created a poll.  I would then be given the option to add a vote through the links section of the json, when I would POST a new vote i would be given a relation to `register vote` and finally doing the POST to `/voters/123/polls` to mark my vote in the voter history. Also depending on how we design the votes api to lookup polls and voters information, that information could be added to each output as well (search or just a specific endpoint).

Also a lot of this can be improved by reading the database or similar behind the scenes calls to get the state of the application or user in a more traditional programming (to get if voter voted on a poll or not and giving right link, or deleting a poll, etc, which would change some links given in a more complex application).  But at the bare minimum we could provide more info to the user on what calls can be used in relation to what they just called as described above.


References I looked at:

- https://www.infoq.com/articles/hypermedia-api-tutorial-part-one/
- https://blogs.mulesoft.com/dev-guides/api-design/api-best-practices-hypermedia-part-1/
- https://apisyouwonthate.com/blog/common-hypermedia-patterns-with-json-hyper-schema/
- https://www.mscharhag.com/api-design/hypermedia-rest

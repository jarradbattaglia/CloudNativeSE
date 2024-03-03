
# Api-DESIGN PART

Based on: https://github.com/ArchitectingSoftware/Cloud-Native-Software-Engineering/blob/main/assignments/api-design-1.md
# Prompt

Start by watching the following video from the 2018 Google Cloud Next Conference on Designing Quality APIs - https://www.youtube.com/watch?v=P0a7PwRNLVU. This video does a nice job introducing some API best practices that we will be incorporating into our voting application.

The purpose of this assignment is to provide you some background on this topic before we discuss in class. After viewing the video, please submit the following:

No more than a one paragraph (a few sentence summary) of the video highlighting one key API design aspect that you took away from the content
Any questions that you have after watching this video. I don't expect you to understand everything, and this part will help me prepare our in-class discussion on the topic. If you got the general idea, you can just say that you have no questions.


# Answer
1. The video went over some issues with APIs (hard to change, software is hard) and the 2 types of styles (entities or procedure) and he seems to go with entity is the better approach. I think the one api design that I took to heart was about the Uniform API and using links for references, so that the user knows how to use your api better. For instance his example for id from using "12345" as an id to using the url "pets/12345", so now you know what entity/object that API is meant to use and can read and query even further and lessen the confusion on the user and for external links you give the full url and the correct API for that resource.  Also his push for flat APIs makes so much more sense than the indexed/nested approach many apis use, I liked that idea as i can organize on my end the way i actually want instead of how the api wanted it.
2. At the end he talks about "efficient communication between tightly coupled components", why is gRPC the recommended choice here (unsure on specifics of grpc specifically, i know we went over it, probably just need to revisit benefits of it).  


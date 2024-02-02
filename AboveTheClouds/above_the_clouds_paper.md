# Above the clouds paper

# Question 1
## Question 1 Prompt

One of the primary objectives of this paper is to motivate how Cloud Computing can redefine "Computing as a Utility". From a utility, think about things like power and water - you turn the light switch on, and the lights come on without you having to know anything about the complex aspects needed to deliver the power to your house. Same for water, you turn on the sink and safe, clean water flows out of the faucet (at least I hope it does). Using power and water as a conceptual model for a utility, how well do you think this paper motivates RAW COMPUTE ASPECTS of Cloud Computing as a utility? By "raw compute aspects" I mean pure computing resources such as processing (CPU/GPU, etc.), storage, and network.

## Answer for Question 1

So personally I do feel the paper does a good job of defining cloud computing as a utility and giving it a defined explanation of what it deems as cloud computing (Computing resources delivered over the internet, in which the hardware and services are run from somewhere else). 

For instance, if i think of AWSs utility as CPU/MEM/STORAGE/etc, I as a developer can spin up any amount I need almost instantly and scale as much as I need.  But just like a utility it is multiplexed! I dont care who I share with or who I am next to or where the water comes from.  All I care about is I get water (Raw compute).  Which lines up with papers reasoning, 1. "illusion of infinite computing resources" - like my water, I assume it will not run out (only in dire situations, etc)!  2. elimination of up front commitment - I don't prepay for my electricity or water, I get charged what i use just like the resources i would use on AWS (whatever hits my meter/instance).  3. "ability to pay on short term basis" - like my utilities, i can shut off or turn on when needed (spin up/down instances/turn off sink/curbstop/etc).  

But even also as they describe only large companies can be a utility, the resources to do this and make money are only available to companies that have the resources to manage this very large amount of resources and do it at a lower price (because they buy resources in such large quantities).

# Question 2
## Question 2 Prompt

Now address the same question as above from a SOFTWARE ENGINEERING PERSPECTIVE. By a "software engineering perspective" I mean that you are a software engineer responsible for design, delivery and operation of a product or service that will be deployed and made available to your customers from the cloud. From this perspective can you outline one example where the analogy of a pure utility (power and water) still applies, and one example where it doesn't? Hints: Take a careful look at the "Obstacles and Opportunities" section of the paper, or you can do some other online research, for example looking into the AWS "Shared Responsibility Model". You are not limited to these sources when answering this question - feel free to research this question and be creative. Make sure you reference any resource that you used to help answer this question, web links are fine, no need for formal references.

## Question 2 Answer

For the answer of the analogy of how this is similar to water/electricity, I would say if I am creating an application that needs to horizonally scale (i need more cpu for instance).  I can spin up more machines automatically.  That could be me coming home and cooking or watching tv and then when im done doing that I turn off my tv (applications/scale down) and I am charged accordingly.  So I can think of my application as a utility from that perspective alone.  Or for a more specific paper obstacle, availability/outages or breaks in my water/electric line affect me as I do a software developer, when AWS goes down, my application will go down as well and obviously the similar aspect of how i am measured by how much i use per resource (pay per how much volume of the resource I use)!

For an example of where the utility to cloud doesnt match is more about ownership and maybe requirments.  For some things to get access to those services I might have to pay a licensing fee for the OS or application im trying to use.  With water/power, there is nothing on top of that i need to worry about.  I plug something in and it can use that resource directly.  For most things in cloud environments, I'll need an OS and application to do something with that resource and maybe pay more to actually use those resources of Mem/CPU/etc.  Also the comparison of resources itself is different, when I am quoted for using a gallon of water from my utility I get a gallon of water,  where as for Cloud environments I am getting the ability to use X CPU/MEM/etc, but I may not actually be using any or all of it. 

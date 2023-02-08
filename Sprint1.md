# Sprint 1
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------
**USER STORIES:**
  1. As a user of the app, I should be able to run the app and connect to a database, put in an input, and see it displayed on the screen (John)
  2. As an end user, I want to sync my list of contacts so that I can build and strengthen my social connections. (Ria)
  3. As an end user, I want to send my contacts images so that I can have a multimedia rich communication with my contacts. (Ria)
  4. As an end user, I want to be able to set up my profile so that people can recognize me and my personality. (Ria)
  5. As an end user, I want to know the statuses of my contacts so that I can start chatting with them when they are available. (Ria)
  6. As a user who is usually on the go, I would like to have a voice message option so I can quickly send a message without having to type it out (Kevin).
  7. As an end user who values my security, I would like to have my login/account information stored safely without having to worry about it being leaked (Kevin).
  8. As a user who does not want my messages/information stolen, I want to have an authentication system attached to the login to ensure that only I can log in (Kevin).
  9. As a user who is not able to send perfect messages, I would like to be able to edit my messages so my texts are not incorrect or misconstrued (Kevin).
  10. As a user who wants to elevate my conversations with others beyond just text, I would like to have the ability to react to messages with emojis/stickers (Kevin).
  11. As a user who is not familiar with how messaging apps work, I want the layout to be straightforward and not cluttered in terms of presentation (Kevin).
  12. As a user of the app, I want my messages to be saved so that I can start a conversation back up later (John).
  13. As a user, I want to be able to search for past messages and conversations, so that I can easily find information that I need (Rafa).
  14. As a user, I want to be able to create and join groups, so that I can communicate with multiple people at once (Rafa).
  15. As a user, I would like to pin conversations that are of higher priority to me (Rafa).
  16. As a user who uses multiple devices, I would like to be able to use the application on multiple devices (Rafa).
  17. As a user, I want to be able to archive or mute conversations, so that I can declutter my inbox or temporarily silence notifications for a chat (Rafa).
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------
**FRONT-END**

**What issues your team planned to address:**

**USER STORIES/ISSUES TACKLED:**
- "As a user who is not familiar with how messaging apps work, I want the layout to be straightforward and not cluttered in terms of presentation."
- "As an end user, I want to be able to set up my profile so that people can recognize me and my personality."
-  "As a user, I want to be able to search for past messages and conversations, so that I can easily find information that I need"

**PLANNED ISSUES TO FIX FOR SPRINT 1 (SEE ISSUES TAB)**
- Create the chat interface between users.
- Create a panel where users can see all their conversations.
- Being able to have a clutter free layout for the app and allow users to easily navigate the app and a basic login/authentication screen for the app.
- Create an avatar (profile picture feature so users can be easily idenitified)
- Being able to scroll up and down in a conversation to search for a message 
- Have a list of all conversations with other users 
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------
- **Which ones were successfully completed:**
  - The issue that was completed would be setting up the basic layout of the app and having an organized setup. 
  - We were able to develop a basic chat interface on how a conversation panel would look like for our users.
  - We were able to add profile picture to the conversation interface, so users can be easily identified
  - The conversations are scrollable so the user can scroll to see previous messages
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------
 - **Which ones didn't and why?**
  - We weren't able to implement the login screen because we are not yet well-versed with the process. As we progress, we will be able to implement this issue.
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------
- Front-end Video: https://youtu.be/ZyRSMR8j1qI
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------
**BACK-END**

**What issues your team planned to address:**

**USER STORIES/ISSUES TACKLED:**
- "As a user of the app, I should be able to run the app and connect to a database, put in an input, and see it displayed on the screen."
- "As a user who is not able to send perfect messages, I would like to be able to edit my messages so my texts are not incorrect or misconstrued."
- "As a user of the app, I want my messages to be saved so that I can start a conversation back up later."
- "As a user, I want to be able to search for past messages and conversations, so that I can easily find information that I need."

**PLANNED ISSUES TO FIX FOR SPRINT 1 (SEE ISSUES TAB)**
- Develop a Database were we can store user and messages information.
- Make the data in the database persist and have it able to retrieve/edit data.
- Create a restAPI with basic GET, POST, PUT functionaliities.
- Connect the database and API together and allow the API to access/edit the information in the database.
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------
- **Which ones were successfully completed:**
  - We were able to get both a basic database and restAPI created.
  - The database is able to insert information and retrieve it. The information also persits.
  - The API is able to GET, POST, and PUT messaging information.
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------
- **Which ones didn't and why?**
  - We were not able to tackle the issue of connecting the database and API together as we had to switch databases from SQLite to MYSQL. This process took a while to   setup/complete. There were several issues with connecting and understanding the SQL/Go syntax. As a result, we were not able to connect the two files. However, the database file has a rough outline of setting up an HTTP request which will be used with the API.
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------
- Back-end Video: https://www.youtube.com/watch?v=gvnAW3_SR1w
-----------------------------------------------------------------------------------------------------------------------------------------------------------------------

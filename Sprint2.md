# Sprint 2

## Work Completed in Sprint 2:

 ### Frontend:
 - stuff here

<hr>

 ### Backend:
 - Integrated the MySQL database into the API file.
   - Reworked the basic REST request functions (POST, PUT, GET, DELETE) to directly interact with the MySQL database.
     - The REST API is now able to create a message in the database, retrieve it, edit it, and delete it.
 - Added additional functions that were specifically requested by the Frontend:
   - Wrote a function to retrieve **all** message between two users (```GetConversation```).
   - Wrote a function that would search the database for a specific message and would return either exact matches or messages that contain the query string (```searchMessage```).
   - Wrote a function that would delete all messages between two users, effectively clearing a conversation (```deleteConversation```).
 - CORS was implemented into the API to allow the Frontend to send requests to the API without it being blocked.
 - Began adding data validation to functions such as POST.
   - The ```createMessage``` function checks if IDs are four digits and if the message is empty.
 - To prevent confusion and mistakes, more error handling was added to each major step in every function to ensure that it executed correctly. 
   - Likewise, console log messages were put at the end of each function to tell the Frontend if the request was successful or not.
 - Added "reset" functions to help with debugging potential issues:
   - Wrote a function to clear the entire database (```deleteTable```).
   - Wrote a function to retrieve every message in the Messages database (```getAllMessages```).
   - Wrote a function to retrieve every user in the Users database (```getAllUsers```). 
 - Created a new database called ```user_accounts``` which will serve as a way of managing people's accounts on the messaging app.
   - Wrote a function to create a new user in the database (```createUserAccount```).
   - Wrote a function to retrieve a user based on their unique ID (```getUser```).
   - Wrote a function to edit an existing user's conversation list by adding a new ID to it (```addConversation```).

<hr>

### Unit Tests Conducted for Backend
  - Unit tests were conducted for the following functionalities:
     - Creating a message.
     - Retrieving all messages between two users (get a conversation).
     - Searching for a specific message.
     - Editing a message.
     - Deleting a specific message.
     - Deleteing an entire conversation.
     - Creating a user account.
     - Editing a user's conversation list by adding a new ID.
     - Retrieving a user.
  - **The Go Test file can be found [here.]()**

<hr>

### API Documentation
  - The API documentation is too large to fit within Sprint.md, so it is linked in a separate .md file below:
    - **The API documentation can be found [here.]()**

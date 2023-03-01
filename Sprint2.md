# Sprint 2

### [Link to Sprint 2 Video](https://www.youtube.com/watch?v=fxNxrzTqC14)

## Work Completed in Sprint 2:

 ### Frontend:
 - Created functions (using GET & POST request from the API ) that connect to the API.
    - Have ```getMessages()``` --> gets all messages from a conversation, uses GET request that takes in two userIDs as parameters.
    - Have ```sendMessage()``` --> creates a message typed by the user in the Database, uses POST request, adds message to MessageList in the Front at the same time, takes in a message body as parameters.
    - Have ```searchMessage()``` --> gets a list of messages based on the content inside of them, uses GET Request, takes in message content as a parameter.
 - When Message Window is opnened, all mesages between the two users are displayed.
 - The Message-Window simultanisly uploads sent message to datbase while displaying it to the screen.
 - Can search for specific messages and the Message-Window will autoscroll to them.
 - Created a login page with a navigation bar at the top of the screen.
   - The page consists of the elements:
      - Username field
      - Password field
      - Login button
      - In the navigation bar:
        - Home
        - Services
        - Chats
        - Sign up
  - Additionally, an icon was added along with the title of the application to the top-left of the navigation bar.
 - A sign up page was also created which included the form fields of create username, create password, confirm password, and email address.
 
<hr>

 ### Backend:
 - Integrated the MySQL database into the API file.
   - Reworked the basic REST request functions (POST, PUT, GET, DELETE) to directly interact with the MySQL database.
     - The REST API is now able to create a message in the database, retrieve it, edit it, and delete it.
 - Added additional functions that were specifically requested by the Frontend:
   - Wrote a function to retrieve **all** message between two users (```GetConversation```).
   - Wrote a function that would search the database for a specific message and would return either exact matches or messages that contain the query string (```searchMessage```).
   - Wrote a function that would delete all messages between two users, effectively clearing a conversation (```deleteConversation```).
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

 ### Entire Team:
  - Connected the Frontend to the REST API.
  - CORS was implemented into the API to allow the Frontend to send requests to the API without it being blocked.

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
  - **The Go Test file can be found [here.](https://github.com/SWEGroup39/GatorChatApp/blob/main/App_Contents/BackEnd/API/GatorChat_Rest_API_test.go)**
   - _**Note:** This leads to the API folder in the "main" branch. To see the full commit history, visit the working [Back-End branch.](https://github.com/SWEGroup39/GatorChatApp/tree/Back-End-Branch)_

### Cypress Test Conducted for Frontend
- All the front end test cases were created using Cypress
- There are 5 tests in total
- Tests 1-4 deal with the chat message window and Test 5 deals with the Login Page
- There are a mixture of only front-end functionality and API testing cases included
- **Test 1**: Checks if local host for the chat window can be accessed, types in the message "Hello! Cypress" and clicks the send button. This test case only checks the front-end functionality and not the API call. There are designated unit tests for the API.
- **Test 2**: Checks if local host for the chat window can be accessed, types in the search bar "Hello" and clicks the search button. This test case ony checks the front-end functionality. It does not account for the API call.
- **Test 3**: This is a unit test case for the search function where it sends a GET request and expects a response of status 200 (OK). This test case ensures that the API call of GET is properly functioning when the search function is running.
- **Test 4**: This is a unit test case for the send message function where it sends a message of How is the weather along with the message id, receiver id, and sender id.    -Sends a POST request and expects a response where the message is posted to the chat window.
   - This test case ensures that the API call of POST is functioning when the send message function is called
- **Test 5**: Checks if local host for the login page can be accessed, types in a username: harry.k and password: harry and clicks the login button. This test case only checks the front-end functionality and not the API call because that has not been implemented yet.
- In summary, Tests 1, 2, and 5 are check the front-end functionality and Tests 3,4 are unit tests for the chat window functions. 
- All the tests cases ensure that both the front and back ends are correctly connected therefore checks if the app is correctly integrated. 
- These test files can be found in the FRONT-END branch in the folder Cypress Tests
- **Link to Cypress Tests**: <a href="https://github.com/SWEGroup39/GatorChatApp/tree/Front-End-Branch/CypressTesting">Cypress Tests</a>
<hr>

### API Documentation
  - The API documentation is too large to fit within Sprint.md, so it is linked in a separate .md file below:
    - **The API documentation can be found [here.](https://github.com/SWEGroup39/GatorChatApp/blob/main/App_Contents/BackEnd/API/API_Documentation.md)**

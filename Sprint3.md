# Sprint 3

### [Link to Sprint 3 Video]()

## Work Completed in Sprint 3:

 ### Frontend:

 
<hr>

 ### Backend:
 - Added More Functionality to the ”UserAccount” Struct
    - Incorporated a **Delete User** function to allow Users to be deleted.

 - Added More Functionality to the ”UserMessage” Struct
    - Updated the **Edit Message** function to search for the to-be-edited message by using the unique GORM ID instead of nonunique information.
    - Updated the **Delete Specific Message** function to search for the to-be-deleted message by using the unique GORM ID instead of nonunique information.
      - Added additional logic to the **Delete Specific Message** function that will soft delete the most recently deleted message by a user and hard deletes any existing soft deleted messages.
        - This is used in the implementation of a newly added function that allows a User to undo their most recently deleted message.
    - Fixed the **Search Message** function so it will only search for whole words instead of searching for words that match but are within a larger word.
    - 

<hr>

### API Documentation
  - The API documentation is too large to fit within Sprint2.md, so it is linked in a separate .md file below:
    - **The API documentation can be found [here.](https://github.com/SWEGroup39/GatorChatApp/blob/main/App_Contents/BackEnd/API/API_Documentation.md)**

<hr>

 ### Entire Team:
  - 

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
- There are 5 tests in total:
   - Tests 1-4 deal with the chat message window and Test 5 deals with the Login Page
     - There are a mixture of only front-end functionality and API testing cases including:
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

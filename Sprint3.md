# Sprint 3

### [Link to Sprint 3 Video]()

## Work Completed in Sprint 3:

 ### Frontend:
  - Changed the whole color scheme of the application to a gradient blue and orange.
  - Added routing for the entire application. 
    - Routed to the pages that include:
       -Home
       -About Us
       -Login
       -Sign Up
       -Dashboard
       -Contacts
       -Profile
       -Settings
       -Conversations
       -Notifications
   - Additionally the login page and sign up page have form validation 
      - A red line appears underneath the textbox if the field is left empty
      - The login and signup button are both grayed out until all the fields are inputted
      - If the login credentials are wrong, an alert box will pop up saying "Incorrect Username and Password! Please try again."
      - It will then reset the form and allow the user to type his or her credentials again. 
   - The login and sign up page both have easy access links at the bottom
      - For the login it has New User? Sign up
      - For the sign up page it has Already a User? Login
   -The home page has a learn more and get started button
      - The learn more button routes to the about us page
      - The get started button routes to the sign up page
   - Added a gator icon to the navigation bar
   - The navigation bar is exists throughout the different pages of the application
   - Finished the POST method for the login page where each user is authenticated and then redirected to the dashboard.
   - Finished the POST method for the sign up page where after a user is successfully created they are redirected to the login page to log in with their credentials
   - Authentication guards were added to prevent not authenticated users from accessing the dashboard component
         - For example: if the url /dashboard is searched it will not take the user to the dashboard page if they are not logged in. 
         - It will redirect them to the login page
   - Once a user enters their dashboard page they have a left side bar where they can access the Notifications, Profile, Settings, Conversations, and Contacts.
   - They also have a sign out button at the right corner of the navigation bar. 
       - If clicked this will redirect them to the login page.
   -Images were added to some parts of the application to make it more visually appealing. 
      - These were included in the About Us page and the Home page
   
<a id="Cypress"></a>
 ## Cypress Tests
    
 - **Test6LoginPost:** This test types in the user credentials of test@ufl.edu and password of pass and clicks the login button. The user gets redirected to the       dashboard      - component. This test is mainly to test the API and whether a user is authenticated or not.<br>
  
 - **Test7Home:** This test visits the home page of the application and clicks on the about us button which routes the user to the about us page. This test is mainly to test routing.<br>
  
 - **Test8GetStarted:** This test visits the home page and clicks on the Get Started button which gets routed to the sign up page. This test also tests the routing functionality.<br>
  
 - **Test9SignUpPost:** This test creates a user with an email, username, and password and clicks the sign up button. Then the user gets created and the page gets redirected to the login page where the user can log in with the created credentials. This test is mainly to test routing and the POST Api call. <br>

<a href = "https://github.com/SWEGroup39/GatorChatApp/tree/Integrated-Front-End-Branch/Cypress%20Tests">Link to Cypress Tests <a>
<hr>

 ### Backend:
 - **Added More Functionality to the ”UserAccount” Struct**
    - Incorporated a **Delete User** function to allow Users to be deleted.
    - Added an "Email" field to the UserAccount struct to allow users a more logical way of logging in.
    - Updated the request type for "GET" functions that took in a body to be "POST" as Angular does not allow a GET request with request body.
    - Added data validation on the **Create User** function that checks:
      - If IDs are four digit and numerical.
      - The email is a UF email (has @ufl.edu).
      - The email is not already in use.
     - Added a function that locates and returns an available User ID that the Frontend can assign to a user that signs up.
     - Added a function that returns a user's information by passing in the User ID.
     - Updated functions that involve taking in the "Password" field to hash it using SHA-256. This is done for security purposes.

 - **Added More Functionality to the ”UserMessage” Struct**
    - Updated the **Edit Message** function to search for the to-be-edited message by using the unique GORM ID instead of nonunique information.
    - Updated the **Delete Specific Message** function to search for the to-be-deleted message by using the unique GORM ID instead of nonunique information.
      - Added additional logic to the **Delete Specific Message** function that will soft delete the most recently deleted message by a user and hard deletes any existing soft deleted messages.
        - This is used in the implementation of a newly added function that allows a User to undo their most recently deleted message.
       - Added a function that will delete **ALL** soft deleted messages in the messages database (Testing Function).
       - Added a function that will return **ALL** currently soft deleted messages in the messages database (Testing Function).
    - Fixed the **Search Message** function so it will only search for whole words instead of searching for words that match but are within a larger word.
    - Implemented a potential solution to making the **GetConversation** have long polling (to be incorporated into the actual application). This makes the messages appear in real-time.
    - Implemented a function that allows users to update their username.
    - Implemented a WIP function that allows users to update their password (to be completed).
 
 - **Updated the Structure of the Backend Unit Tests**
    - To prevent the Sprint2 issue of potentially having conflicting entries in the database that would mess up the unit tests, we made the following structural changes to ensure that all unit test results will be what is expected.
      - Unit tests for the "UserMessage" Struct no longer have a hardcoded GORM ID but will instead take on the next auto-generated ID. Previously, since a large number was selected for the hardcoded GORM ID, subsequent messages will continue off of that large value and cut off a large chunk of potential GORM ID values.
      - Unit tests for the "UserAccount" Struct now have reserved IDs and messages to prevent confounding users that may mess up some of the unit tests' expected values.
        - User IDs from 9996 to 9999 have now been reserved solely for unit testing purposes. Additionally, unique messages that will not be messages by users of the app have been used as the hardcoded message.

- The API Documentation Was Also Updated for All of the Above Changes.

<hr>

### API Documentation
  - The API documentation is too large to fit within Sprint3.md, so it is linked in a separate .md file below:
    - **The API documentation can be found [here.](https://github.com/SWEGroup39/GatorChatApp/blob/main/App_Contents/BackEnd/API/API_Documentation.md)**

<hr>

 ### Entire Team:
  - Connected Frontend login page with the Backend function to receive and validate a user's login information.
  - Began to integrate long polling into the **getConversation** function which will allow for real-time message updates.
  - Integrated the **deleteSpecificMessage** function into the message window.
  - Used the **getUser** function to retrieve a user from the database, display their personal information (name, conversation list, messages).
 
<hr>

### Unit Tests Conducted for Backend
  - Unit tests were conducted for the following functionalities:
     - Creating a message.
     - Retrieving all messages between two users (get a conversation).
     - Searching for a specific message.
     - Editing a message.
     - Deleting a specific message.
     - Deleteing an entire conversation.
     - Undoing a delete.
     - Creating a user account.
     - Editing a user's conversation list by adding a new ID.
     - Updating a user's username.
     - Updating a user's password (WIP).
     - Retrieving a user (by passing in login credentials).
     - Retrieving a user (by passing in the user's unique ID).
     - Deleting a user.
     - Getting an available user ID.
  - **The Go Test file can be found [here.](https://github.com/SWEGroup39/GatorChatApp/blob/main/App_Contents/BackEnd/API/GatorChat_Rest_API_test.go)**
   - _**Note:** This leads to the API folder in the "main" branch. To see the full commit history, visit the working [Back-End branch.](https://github.com/SWEGroup39/GatorChatApp/tree/Back-End-Branch)_

### Cypress Test Conducted for Frontend
- All the front end test cases were created using Cypress
- There were five tests created in Sprint 2 (tests created in Sprint 3 are found [here](#Cypress):
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

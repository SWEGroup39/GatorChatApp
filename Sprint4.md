# Sprint 4

### [Link to Sprint 4 Video](https://www.youtube.com/watch?v=uT1i3iMAUAQ)

## Work Completed in Sprint 4:

 ### Frontend:
  - 

<hr>

 ### Backend:
 - 

- The API Documentation Was Also Updated for All of the Above Changes.

<hr>

### API Documentation
  - The API documentation is too large to fit within Sprint4.md, so it is linked in a separate .md file below:
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

 ### Cypress Tests
 - **_NOTE: These cypress tests include the past test cases as well. They are an updated verson plus test for additional features. The old cypress tests test the old functionality when the application was not integrated._**

 - **Test5Messages:** This test vists the log in page and manually logs in with the user credentials of test@ufl.edu and pass. This then redirects them to the dashboard where under the conversations tab you have access to the message window. This test then tests the functionality of the send message. Additionally, it tests the functionality of the POST API call of the send message function. 

 - **Test6LoginPost:** This test types in the user credentials of test@ufl.edu and password of pass and clicks the login button. The user gets redirected to the       dashboard component. This test is mainly to test the API and whether a user is authenticated or not.<br>
  
 - **Test7Home:** This test visits the home page of the application and clicks on the about us button which routes the user to the about us page. This test is mainly to test routing.<br>
  
 - **Test8GetStarted:** This test visits the home page and clicks on the Get Started button which gets routed to the sign up page. This test also tests the routing functionality.<br>
  
 - **Test9SignUpPost:** This test creates a user with an email, username, and password and clicks the sign up button. Then the user gets created and the page gets redirected to the login page where the user can log in with the created credentials. This test is mainly to test routing and the POST Api call. <br>

<a href = "https://github.com/SWEGroup39/GatorChatApp/tree/Integrated-Front-End-Branch/Cypress%20Tests">Link to Cypress Tests <a>

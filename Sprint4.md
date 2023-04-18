# Sprint 4

### [Link to Sprint 4 Video](https://www.youtube.com/watch?v=uT1i3iMAUAQ)

## Work Completed in Sprint 4:

 ### Frontend:
  - As long as the user does not click the sign out button or close the browser tab that user session will continue and the user will have access to his or her dashboard.
   - Implemented the use of sessionStorage on the client side to store the user credentials
   - Used information from session storage in the profile, contacts, settings, chat-list, messages, and dashboard components 
  - Added footer a using bootstrap that contains:
     - @2023 GatorChat Inc.
     - Gator icon
     - Home: If this is clicked, the user will be routed to the home page.
     - About us: If this is clicked, the user will be routed to the about us page.
  - Modified the sign up page create user function where it integrates the back end user ID function
   - The back end userID function creates the correct unique id for each user created.
  - Modified the query params so that the url does not display sensitive user credentials
   -Session storage was used to fix or update this approach 
  - Added the phone and full name fields to the sign up page
   -These fields were added for the profile page  
  - Dashboard now has welcome back username message displayed
  - Profile page is implemented to where the profile picture is the capitalized initials of their full name in an orange background with blue font
   - The profile page consists of the Username, Email, Phone number, and Full Name fields
   - It also has a secondary navigation bar which shows the route the user took to get to the page as well as routes back to the dashboard page if clicked
  -Created the settings page with a form field and a  sub navigation bar where users can effectively navigate back to the dashboard. 
   - Has a delete account option that successfully deletes the user account and redirects users back to the login page.
  - Finished the editing functionalities of the settings page.
     - This includes:
      - Edit Username: Has an edit and save button to individually call the functions 
      - Edit Phone Number: Has an edit and save button to individually call the functions 
      - Edit Full Name: Has an edit and save button to individually call the functions 
      - Edit Password: User has to type in the old password and the new/intended password
      - The save changes button underneath saves the password changes
      - The cancel button refreshes the page to prevent the password change from occuring
  - Added a contacts page that displays the username and id of a contact
     - This information is based off of the current conversations array for the logged in or current user
     - Local storage is used to store the contacts for each user
     - A search bar was added for searching a username in the contact page in order to add a contact and its id to the current conversations array.
     - Styled the contact page to where the odd numbered contact has a grey background color
     - A trash can icon was included for the delete functionality
      - If this icon is clicked, the contact is removed from the list and that contact's id is removed from the current conversations array
- Added some routing/easy access headings to the dashboard page
  - This includes:
   - A group heading that when clicked routes to the most recent conversation
   - A group heading that when clicked routes to the contacts page
   - A group heading that when clicked routes to the settings page
   - A group heading that when clicked routes to the profile page
- Users can login with different number of browser tabs and still continue their session without affecting the other logged in user sessions.
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

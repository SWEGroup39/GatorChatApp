<a id="TOC"></a>
# GatorChat API Documentation

### Written by: Kevin Cen and John Struckman

#### This API serves to provide an **abstraction** to the process of manually accessing the database and running MySQL commands to perform queries.
---

## Table of Contents

- [Setting Up Golang](#settingUp)

- [Accessing the GatorChat API](#accessingAPI)

- [Overview of **REST** Functions for Messages](#REST_Messages)

    - [Overview of  **POST** Commands for Messages](#POST_Messages)

    - [Overview of **PUT** Commands for Messages](#PUT_Messages)

    - [Overview of **GET** Commands for Messages](#GET_Messages)

    - [Overview of **DELETE** Commands for Messages](#DELETE_Messages)

- [Overview of **REST** Functions for Users](#REST_Users)

    - [Overview of  **POST** Commands for Users](#POST_Users)

    - [Overview of **PUT** Commands for Users](#PUT_Users)

    - [Overview of **GET** Commands for Users](#GET_Users)

    - [Overview of **DELETE** Commands for Users](#DELETE_Users)

---
<a id="settingUp"></a>

### ➜ Setting Up Golang

#### Please refer to the instructions linked below to properly install Go.

- Golang can be setup by following the instructions listed [here](https://github.com/rb-uf/swe-project/blob/main/go-setup.md).
- It is important to place your Go projects in a valid directory, for example:
    - ```C:\Users\[USER]\go\src\github.com\kevinc3n\API```
        - This ensures that Golang can find all of the packages and can properly run.

- **NOTE:** It is important to install the dependencies/packages required to use the GatorChat API.

- These include:
  - **GORM**
      - GORM is needed for the "gorm.io/driver/mysql" and "gorm.io/gorm" packages.
  - **Gorilla Mux**
    - Gorilla Mux is needed for the "github.com/gorilla/mux" package.
  - **CORS**
    - CORS is needed for the ""github.com/rs/cors" package.

- **Quick Reference**: Use ```go get -u <package>``` in your command line to install a certain package.

---

<a id="accessingAPI"></a>

### ➜ Accessing the GatorChat API

- Once Golang is installed, the GatorChat API can now be opened and run.

- In the [Back-End-Branch](https://github.com/SWEGroup39/GatorChatApp/tree/Back-End-Branch) of the [Github repository](https://github.com/SWEGroup39/GatorChatApp), there is a file named ```GatorChat_Rest_API.go```.
- This file contains the API file that must be run in order to make requests to the API.
- **Pull** the Back-End-Branch into your repository folder (or manually download the file) to access the API.
    - **To pull the branch into your folder through the command line, use the following commands:**
        - _**This assumes that the project has already been forked into a folder on your computer.**_
        - Open the command line/terminal and navigate to your repository folder using the ```cd``` command.
            - For example: ```cd C:\Users\[USER]\Desktop\SWE\GatorChatApp```
        - Next, run the following command to have all of the branch's files be placed into your repository folder.
            - ```git pull origin [BRANCH_NAME] ```
            - In this case, it is ```git pull origin Back-End-Branch ```.
        - The folder should now contain the API file.

- To run the file, open the terminal in your respective IDE and run the following commands:
    - ```go build```
        - Once it has finished, run the command ```go run GatorChat_Rest_API.go```.
        - If running the code in VSCode you can run the command ```./(put the name of the .exe file that was made by go build here)``` instead
- The API should now be running. The localhost port should be active and able to receive requests.
    - In the scenario where the program cannot connect to the database, an error message will appear in the terminal:
        - If the ```user_messages``` database cannot be opened, then "Error: Failed to connect to messages database." will be displayed.
        - If the ```user_accounts``` database cannot be opened, then "Error: Failed to connect to users database." will be displayed.
- **NOTE:** By default, the API is hosted on **port 8080**. This can be changed if the port is already being used.
---

<a id="REST_Messages"></a>

### ➜ Overview of REST Functions for Messages

- The GatorChat API is built upon the **REST** functions **POST**, **PUT**, **GET**, and **DELETE**.

### **For the Messages database:**
- The API supports **POST** to create a message and store it in a messages database for later retrieval. **POST** is also used for functions that require passing in information through the body.
- The API supports **PUT** to update an existing message in the messages database with a new message.
- The API supports **GET** to retrieve messages from the messages database based on certain parameters.
- The API supports **DELETE** to remove messages from the messages database that are no longer needed.
---

<a id="POST_Messages"></a>

### ➜ Overview of  **POST** Command for Messages

- The first **POST** command takes in an input and creates a new message in the messages database. The second and third focus on message retrieval. _While the remaining two options seem more like a "GET", it is not possible to send a GET request with information in the body. Therefore, a POST is used._

### Syntax

- There are currently three POST commands, available:

     - **First Option: Create a Message**: 
        - This **POST** function creates a message object in the messages database and returns the object back to the requester.
        - **Example Syntax:**
        ```http://localhost:8080/api/messages ```

        - For post, the information passed in must be through the request **body**.
        - The required inputs are a unique message ID number, a message string, a sender ID, and a receiver ID.
    
        - **Example Syntax:**
            ```
                {
                    "ID": null,
                    "CreatedAt": null,
                    "UpdatedAt": null,
                    "DeletedAt": null,
                    "message": "Message goes here.",
                    "sender_id": "1234",
                    "receiver_id": "4321"
                }
            ```
    - Input **"null"** for the "CreatedAt", "UpdatedAt", and "DeletedAt" date inputs, these will be automatically filled in.
    - The message ID number cannot be reused unless the previous message with that number was hard-deleted.
    
    - **Second Option: Search for Message in All Conversations**: 
        - This **POST** function returns the message object that matches the specified message, if it exists in the messages database.
        - This function looks for a message across **ALL** conversations.
        - It will find messages that match it exactly or contain the search parameter somewhere within it. It is not case-sensitive.
        - **Example Syntax:**
        ```http://localhost:8080/api/messages/searchAll```
    
    - **Third Option: Search for Message in One Conversation**: 
        - This **POST** function returns the message object that matches the specified message, if it exists in the messages database.
        - This function looks for a message across **ONE** conversation.
        - It will find messages that match it exactly or contain the search parameter somewhere within it. It is not case-sensitive.
        - **Example Syntax:**
        ```http://localhost:8080/api/messages/[FIRST ID]/[SECOND ID]/search```

    - **NOTE:** For the **second** and **third** option, the search message should be placed in the body of the POST request:

        - **Example Syntax:**
            ```
                {
                    "message": "Hi there"
                }
            ```

### Requirements and Error Messages
- A StatusBadRequest error will be returned if the passed-in body cannot be decoded.
- The inputs for the POST command must follow a specific format.
    - The Sender_ID must be **numeric** and **only four digits**.
    - The Receiver_ID must be **numeric** and **only four digits**.
        - If these are not met, an error message will appear that describes what specifically needs to be fixed.
- If the **Sender ID** has any non-numeric characters in it, the "Invalid Sender ID (NOT NUMERIC)" message will be returned.
- If the **Receiver ID** has any non-numeric characters in it, the "Invalid Receiver ID (NOT NUMERIC)" message will be returned.
- If the **Sender ID** is not 4 digits long, the Invalid Sender ID (NOT FOUR DIGITS) message will be returned.
- If the **Receiver ID** is not 4 digits long, the Invalid Receiver ID (NOT FOUR DIGITS) message will be returned.
- If a message is posted with no actual text in the message the "Invalid Message: Messages cannot be empty." message will be returned.
- The **"Search for Message in ALL/ONE Conversation(s)"** functions must have a valid message that exists in the database, or else "No messages found." will be returned.
- An **Internal Server Error** will be returned if there are errors regarding the database connection or the query itself.
- If all requirements are met, either a single user object or a slice of user objects will be returned along with a successful console log message.
---

<a id="PUT_Messages"></a>

### ➜ Overview of  **PUT** Command for Messages

- The **PUT** command takes in an input to edit a message already in the messages database.

### Syntax

- There are currently two **PUT** commands available:

    - **First Option: Edit Message Contents**:
        - This **PUT** function updates an existing message's contents with a new message.
        - **Example Syntax:**
        ```http://localhost:8080/api/messages/[ID] ```

        - The required input is the unique GORM ID that is within a UserMessage struct.

        - The **new** message should be placed in the body of the PUT request.
            - The logic is that, when a user edits a message, that message struct will already be known (through a GET function) and the unique GORM ID can be passed into this function. This ensures that the correct message is beind edited.

            - **Example Syntax:**
                ```
                    {
                        "message": "Updated message"
                    }
                ```
    <a id="Undo"></a>
    - **Second Option: Undo a Deleted Message**:
        - This **PUT** function will undo a user's most recently deleted message. This is accomplished by setting the message's DeletedAt field to be null again.
        - **Example Syntax:**
        ```http://localhost:8080/api/messages/undo/[ID] ```

        - The required input is the user's Sender ID that is within a UserMessage struct.

### Requirements and Error Messages
- For the **Edit Message** function:
    - A StatusBadRequest error will be returned if the passed-in body cannot be decoded.
    - The message with the input Sender ID and Receiver ID must already exist in the messages database to be edited, otherwise an error will be thrown.
        - If the message-to-change was not located, an error message saying "Message not found." will be returned.
    - Otherwise, the updated message object will be returned along with a "Message edited successfully." console log message.
- For the **Undo Deleted Message** function:
    - If the message cannot be located in the database, then "Message not found." will be returned.
    - Otherwise, the message will be reinstated into the database (which can be verified by calling the **Get Conversation** function), the updated message object will be returned, and a successful console log message will be returned.
- An **Internal Server Error** will be returned if there are errors regarding the database connection or the query itself.
---

<a id="GET_Messages"></a>

### ➜ Overview of  **GET** Command for Messages

- The **GET** command returns messages that have been created with the **POST** request.

### Syntax
- There are currently three different **GET** functions available:

    - **First Option: Get Conversation**:
        - This **GET** function returns all messages between the specified sender and receiver IDs.
         - **Example Syntax:**
        ```http://localhost:8080/api/messages/[FIRST ID]/[SECOND ID]```
        - This returns all the messages, in a slice/array, where the first ID was either the sender/receiver and the second ID was either the sender/receiver.

     - **Second Option: Get ALL Messages**: 
        - This **GET** function returns every message in the messages database.
         - **Example Syntax:**
        ```http://localhost:8080/api/messages ```
        - **NOTE:** _This is more of a testing function rather than a function that would be frequently/practically used._

     - **Third Option: Get ALL Deleted Messages**: 
        - This **GET** function returns every soft deleted message in the messages database.
        - **Example Syntax:**
        ```http://localhost:8080/api/deleted ```
        - **NOTE:** _This is considered a testing function and not for Frontend purposes._

### Requirements and Error Messages
- The **"Get Conversation"** function must have a valid conversation that exists in the database, or else "Conversation not found." will be returned.
- If the **"Get ALL Messages/Get ALL Deleted Messages"** function cannot locate any messages, then a message describing how no messages were found will be returned.
- An **Internal Server Error** will be returned if there are errors regarding the database connection or the query itself.
- If all requirements are met, the message(s) will be returned along with a successful console log message.
---

<a id="DELETE_Messages"></a>

### ➜ Overview of  **DELETE** Command for Messages

- The **DELETE** command deletes messages created with the **POST** request from the messages database.

### Syntax
-  There are currently four different **DELETE** functions available:

     - **First Option: Delete a Specific Message**:
        - This **DELETE** function deletes a specified messaged between a sender and receiver, if it exists in the messages database.
         - **Example Syntax:**
        ```http://localhost:8080/api/messages/[ID]```
        - This function takes in the unique GORM ID in a UserMessage struct.
        - **NOTE:** This function soft deletes a message. In other words, the message still exists, but there is a timestamp in its "DeletedAt" property. It will not appear in the normal GET functions or search functions.
            - If a specific user deletes a message (e.g. message A), and then deletes another message (e.g. message B), then message A will be hard deleted and unable to be brought back. This is because **a user is only able to bring back their most recently deleted message.**
                - For information on how to **undo a delete**, [click here](#Undo) to visit the associated PUT function.

     - **Second Option: Delete an Entire Conversation**:
        - This **DELETE** function deletes all messages between a sender and receiver, if they have a current conversation.
         - **Example Syntax:**
        ```http://localhost:8080/api/messages/[FIRST ID]/[SECOND ID]```
        - This function takes in the two IDs of the people whose conversation you want deleted.
        - **NOTE:** This function "hard" deletes the conversation. In other words, this action cannot be reversed and the conversation will be permanently deleted.

     - **Third Option: Delete All Conversations**:
        - This **DELETE** function deletes the entire database of messages.
         - **Example Syntax:**
        ```http://localhost:8080/api/messages/deleteTable```
        - **NOTE:** _This function is used for testing purposes and is most likely not going to be an implemented function in the Frontend._

    - **Fourth Option: Delete All Deleted Messages**:
        - This **DELETE** function deletes all messages that are currently **soft deleted**.
         - **Example Syntax:**
        ```http://localhost:8080/api/messages/deleteDeleted```
        - **NOTE:** _This function is used for testing purposes and is not considered a Frontend feature._

### Requirements and Error Messages
- The **"Delete a Specific Message"** function must have a valid Sender ID, Receiver ID, and message. Otherwise, "Message not found." will be returned. If the message was found and deleted, then "Message deleted successfully." will be returned.
- The **"Delete an Entire Conversation"** function must have a valid Sender ID and Receiver ID. If it does not, "Conversation not found." will be returned. If the messages were found and deleted, then "Conversation deleted successfully." will be returned.
- The **"Delete All Conversations"** function will panic if the table is unable to be truncated/deleted. If it is able to clear the entire table, then "Database deleted successfully." will be returned.
- The **"Delete All Deleted Messages"** function requires there to be at least one deleted message in the database. If there is not, then "Messages not found." will be returned.
- An **Internal Server Error** will be returned if there are errors regarding the database connection or the query itself.
---

<a id="REST_Users"></a>

### ➜ Overview of REST Functions for Users

- The GatorChat API is built upon the **REST** functions **POST**, **PUT**, **GET**, and **DELETE**.

### **For the Messages database:**
- The API supports **POST** to create a user and store it in the users database for later retrieval. **POST** is also used for functions that require passing in information through the body.
- The API supports **PUT** to update an existing user's conversation list in the users database with a new user.
- The API supports **GET** to retrieve a user from the users database.
- The API supports **DELETE** to delete a user from the users database.
---

<a id="POST_Users"></a>

### ➜ Overview of  **POST** Command for Users

- The first **POST** command takes in an input and creates a new user in the users database. The second command is focused on a retrieving a specific user. _The second option appears to be more of a "GET", but a "GET" does not allow for a body request. Therefore a POST is used._

### Syntax

- There are currently two POST commands available:

    - **First Option: Create a User:**
        - This **POST** function creates a user in the users database and returns the user object back to the requester.
        - **Example Syntax:**
        ```http://localhost:8080/api/users ```

        - For post, the information passed in must be through the request **body**.
        - The required inputs are a username, a password, a user ID, an email, and a list of current conversations that the user is in (typically left blank).
        - **NOTE:** The ID value should be manually inserted and must be **a number between 0000 and 9995** (9996 to 9999 are being reserved for the unit tests).
        
            - **Example Syntax:**
                ```
                    {
                        "username": "user",
                        "password": "pass",
                        "user_id": "1234",
                        "email": "example@ufl.edu",
                        "current_conversations": ["4321", "5678"]
                    }
                ```
        - **NOTES:** 
            - Have ```"current_conversations"``` be ```[]``` if you want a user to have no current conversations.
            - The password will be **hashed** using **SHA256** for security purposes. After calling this create user function, any subsequent retrievals of the user will have the hashed password.

    - **Second Option: Get a Specific User**: 
        - This **POST** function returns a singular user from the users database.
        - It will find a user that matches the credentials.
        - **Example Syntax:**
        ```http://localhost:8080/api/users/User```
         - **NOTE:** User in this case is the word "User". In cases where the syntax involves filling in a parameter, brackets ([]) will surround the word.
         - For get, the information passed in must be through the request **body**.
            - The required input is the user's email and password.
            
                - **Example Syntax:**
                    ```
                        {
                            "email": "example@ufl.edu",
                            "password": "pass"
                        }
                    ```

### Requirements and Error Messages
- A StatusBadRequest error will be returned if the passed-in body cannot be decoded.
- The **Create User** function must have:
    - A **unique**, **four-digit ID** that is **less than 9996**.
    - The email must have a University of Florida **domain name** (i.e. it must end with "@ufl.edu").
        - The email must also be unique, in the sense that no other existing accounts currently have that email.
    - Otherwise, a 400 Bad Request will be returned.
- An **Internal Server Error** will be returned if there are errors regarding the database connection or the query itself (e.g. the requested user could not be found in the database).
- If all requirements are met, a user object will be returned along with a successful console log message.
---

<a id="PUT_Users"></a>

### ➜ Overview of  **PUT** Command for Users

- The **PUT** command 

### Syntax

- There are currently three **PUT** commands available:

    - **First Option: Add Conversation**
        - This **PUT** function takes in a user and updates their conversation list by adding in the passed-in user ID.
        - **Example Syntax:** ```http://localhost:8080/api/users/[FIRST ID]/[SECOND ID]```
        - The required inputs are the user's ID (```FIRST ID```) and the ID that you want added to ```FIRST_ID```'s conversation list (```SECOND ID```).
    - **Second Option: Edit Name**
        - This **PUT** function takes in a user and changes their username based on an input
        - **Example Syntax:** ```http://localhost:8080/api/users/updateN/[ID]```
            - The required input is new username.
            
                - **Example Syntax:**
                    ```
                        {
                            "username": "example",
                        }
                    ```
    - **Third Option: Edit Password**
        - This **PUT** funciton takes in a user, changes their password based on an input, and encrypts it again
        - **Example Syntax:** ```http://localhost:8080/api/users/updateP/[ID]```
            - The required input is new password.
            
                - **Example Syntax:**
                    ```
                        {
                            "password": "pass",
                        }
                    ```

### Requirements and Error Messages
- An **Internal Server Error** will be returned if it is unable to locate the passed-in user or if there are errors regarding the database connection.
- If all requirements are met, the updated user object will be returned along with a "ID added successfully." console log message.
---

<a id="GET_Users"></a>

### ➜ Overview of  **GET** Command for Users

- The **GET** command returns information about users that have been created with a **POST** request.

### Syntax
- There are currently three **GET** function available:

    - **First Option: Get All Users**:
        - This **GET** function returns all users in the users database.
         - **Example Syntax:**
        ```http://localhost:8080/api/users```
        - **NOTE:** _It is expected that this function is merely a testing function and will not be implemented in the Frontend._

    - **Second Option: Get Next ID**:
        - This **GET** function returns a valid ID that has not been used inserted yet in the users database.
            - **Example Syntax:**
        ```http://localhost:8080/api/users/nextID```

    - **Third Option: Get User by ID**:
        - This **GET** function returns a user from the users database based on the user's unique ID.
            - **Example Syntax:**
        ```http://localhost:8080/api/users/[ID]```

### Requirements and Error Messages
- A StatusBadRequest error will be returned if the passed-in body cannot be decoded.
- The **"Get All Users"** function must have users that exists in the database, or else "Users not found." will be returned.
- The **"Get Next ID"** function must still have available user IDs in the users database (0000 to 9995), or else "Max number of users reached!" will be returned.
- The **"Get User by ID"** function must have the requested user exist in the users dataabse, or else a StatusBadRequest error will be returned.
- If all requirements are met, the user(s) will be returned along with a successful console log message.
---
<a id="DELETE_Users"></a>

### ➜ Overview of  **DELETE** Command for Users

- The **DELETE** command takes in a user and deletes them from the user database.

### Syntax
- There is currently only one DELETE command, and the syntax is as follows:

- ```http://localhost:8080/users/[USER_ID] ```

### Requirements and Error Messages
- An **Internal Server Error** will be returned if there are errors regarding the database connection or the query itself.
- If all requirements are met, the user will be removed from the database along with a "User deleted successfully." console log message.
---

[Back to top](#TOC)

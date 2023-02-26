<a id="TOC"></a>
# GatorChat API Documentation

### Written by: Kevin Cen and John Struckman
---

## Table of Contents

- [Setting Up Golang](#settingUp)

- [Accessing the GatorChat API](#accessingAPI)
<br>
- [Overview of **REST** Functions for Messages](#REST_Messages)

    - [Overview of  **POST** Commands for Messages](#POST_Messages)

    - [Overview of **PUT** Commands for Messages](#PUT_Messages)

    - [Overview of **GET** Commands for Messages](#GET_Messages)

    - [Overview of **DELETE** Commands for Messages](#DELETE_Messages)
<br>
- [Overview of **REST** Functions for Users](#REST_Users)

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
---

<a id="REST_Messages"></a>

### ➜ Overview of REST Functions for Messages

- The GatorChat API is built upon the **REST** functions **POST**, **PUT**, **GET**, and **DELETE**.

### **For the Messages database:**
- The API supports **POST** to create a message and store it in a database for later retrieval.
- The API supports **PUT** to update an existing message in the database with a new message.
- The API supports **GET** to retrieve messages from the databse based on certain parameters.
- The API supports **DELETE** to remove messages from the databse that are no longer needed.

- This API serves to provide an **abstraction** to the process of manually accessing the database and running MySQL commands to perform the actions listed above.

---

<a id="POST_Messages"></a>

### ➜ Overview of  **POST** Command for Messages

- The **POST** command takes in an input and creates a new message in the database.

### Syntax
- There is currently only one POST command, and the syntax is as follows:

- ```http://localhost:8000/api/messages ```

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

### Requirements and Error Messages
- The inputs for the POST command must follow a specific format.
    - The Sender_ID must be **numeric** and **only four digits**.
    - The Receiver_ID must be **numeric** and **only four digits**.
        - If these are not met, an error message will appear that describes what specifically needs to be fixed.
- If the entry cannot be inserted into the database table, an Internal Server Error will be returned. 
- Otherwise, the newly created message object will be returned along with a "Message created successfully." message.
---

<a id="PUT_Messages"></a>

### ➜ Overview of  **PUT** Command for Messages

- The **PUT** command takes in an input to edit a message already in the database.

### Syntax

- There is currently only one **PUT** command, and the syntax is as follows:

- ``` http://localhost:8000/api/messages/[FIRST ID]/[SECOND ID]/[ORIGINAL MESSAGE] ```

- The required inputs are the Sender ID, the Receiver ID, and the original message that you would like to have changed.

- The **new** message should be placed in the body of the PUT request.

    - **Example Syntax:**
        ```
            {
                "message": "Updated message"
            }
        ```
### Requirements and Error Messages
- The message with the input Sender ID and Receiver ID must already exist in the database to be edited, otherwise an error will be thrown.
    - If the message to edit cannot be found, an Internal Server Error will be returned.
    - If the message-to-change was not located, an error message saying "Message not found." will be returned.
- Otherwise, the updated message object will be returned along with a "Message edited successfully." message.
---

<a id="GET_Messages"></a>

### ➜ Overview of  **GET** Command for Messages

- The **GET** command prints displays messages that have been created with the **POST** request.

### Syntax
- There are currently three different **GET** functions available:

    - **First Option: Get Conversation**:
        - This **GET** function returns all messages between the specified sender and receiver IDs.
         - **Example Syntax:**
        ```http://localhost:8000/api/messages/[FIRST ID]/[SECOND ID] ```
        - This returns all the messages, in a slice/array, where the first ID was either the sender/receiver and the second ID was either the sender/receiver.
    
    - **Second Option: Search for Message**: 
        - This **GET** function returns the message object that matches the specified message, if it exists in the database.
         - **Example Syntax:**
        ```http://localhost:8000/api/messages/[MESSAGE] ```
        - If the message contains spaces, use ```%20``` in place of the space.
        - **NOTE:** _The searching functionality is currently designed to only find messages that match exactly with the input message. In the future, this will be tweaked to find messages that contain the input message._

     - **Third Option: Get ALL Messages**: 
        - This **GET** function returns every message in the database.
         - **Example Syntax:**
        ```http://localhost:8000/api/messages ```
        - **NOTE:** _This is more of a testing function rather than a function that would be frequently/practically used._

### Requirements and Error Messages
- The **"Get Conversation"** function must have a valid Sender and Receiver ID, or else "Messages not found." will be returned.
- The **"Search for Message"** function must have a valid message that exists in the database, or else "Message not found." will be returned.
- If the **"Get ALL Messages"** function cannot locate all of the messages in the database, then an Internal Server Error will be returned.
---

<a id="DELETE_Messages"></a>

### ➜ Overview of  **DELETE** Command for Messages

- The **DELETE** command deletes messages created with the **POST** request from the database.

### Syntax
-  There are currently three different **DELETE** functions available:

     - **First Option: Delete a Specific Message**:
        - This **DELETE** function deletes a specified messaged between a sender and receiver, if it exists in the database.
         - **Example Syntax:**
        ```http://localhost:8000/api/messages/[FIRST ID]/[SECOND ID]/[MESSAGE] ```
        - This function takes in a Sender ID, Receiver ID, and the message in the conversation that you want deleted.
        - If the message contains spaces, use ```%20``` in place of the space.

     - **Second Option: Delete an Entire Conversation**:
        - This **DELETE** function deletes all messages between a sender and receiver, if they have a current conversation.
         - **Example Syntax:**
        ```http://localhost:8000/api/messages/[FIRST ID]/[SECOND ID] ```
        - This function takes in the two IDs of the people whose conversation you want deleted.

     - **Third Option: Delete All Conversations**:
        - This **DELETE** function deletes the entire database of messages.
         - **Example Syntax:**
        ```http://localhost:8000/api/messages/deleteTable```
        - **NOTE:** _This function is used for testing purposes and is most likely not going to be an implemented function in the Frontend._

### Requirements and Error Messages
- The **"Delete a Specific Message"** function must have a valid Sender ID, Receiver ID, and message. Otherwise, an HTTP 404 not found error will be returned. If the message was found and deleted, then "Message deleted successfully." will be returned.
- The **"Delete an Entire Conversation"** function must have a valid Sender ID and Receiver ID. If it does not, an HTTP 404 not found error will be returned. If the messages were found and deleted, then "Messages deleted successfully." will be returned.
- The **"Delete All Conversations"** function will panic if the table is unable to be truncated/deleted. If it is able to clear the entire table, then "Table was deleted." will be returned.
---

<a id="REST_Users"></a>

### ➜ Overview of REST Functions for Users

- The GatorChat API is built upon the **REST** functions **POST**, **PUT**, **GET**, and **DELETE**.

### **For the Users database:**

[Back to top](#TOC)

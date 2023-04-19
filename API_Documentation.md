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
        - This ensures that Golang can find all the packages and can properly run.

- **NOTE:** It is important to install the dependencies/packages required to use the GatorChat API.

- These include:
  - **GORM**
      - GORM is needed for the "gorm.io/driver/mysql" and "gorm.io/gorm" packages.
  - **Gorilla Mux**
    - Gorilla Mux is needed for the "github.com/gorilla/mux" package.
  - **CORS**
    - CORS is needed for the ""github.com/rs/cors" package.
  - **azblob**
    - azblob is a package used for handling images by storing them in a container on the Microsoft Azure account. It is for the "github.com/Azure/azure-storage-blob-go/azblob" package.

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
        - Next, run the following command to have all the branch's files be placed into your repository folder.
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

### ➜ Overview of **POST** Command for Messages

- The first **POST** command takes in an input and creates a new message in the messages database. The second and third focus on message retrieval. _While the remaining two options seem more like a "GET", it is not possible to send a GET request with information in the body. Therefore, a POST is used._

### Syntax

- There are currently four POST commands, available:

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
                    "receiver_id": "4321",
                    "image": null
                }
            ```
    - Input **"null"** for the "CreatedAt", "UpdatedAt", and "DeletedAt" date inputs, these will be automatically filled in.
    - The message ID number cannot be reused unless the previous message with that number was hard-deleted.
    - No image should be passed in this function. If you want to create a message with an image, use the second option below.

    - **Second Option: Create a Message (With an Image)**: 
        - This **POST** function creates a message object in the messages database, with an image attachment stored as a BLOB (Binary Large Object), and returns the object back to the requester.
        - **Example Syntax:**
        ```http://localhost:8080/api/messages/image ```

        - The way of passing in data to this function is **different** for this function than all other functions.
            - Instead of passing in a JSON body, the data must be passed using **form-data**. This involves making key-value pairs where the key is the name of the field and the value is the actual content.
                - For example:
                    ```
                        key: sender_id
                        value: 9999

                        key: receiver_id
                        value: 9998

                        key: message
                        value: Check out this photo:

                        key: image
                        value: image.PNG (File)
                    ```
                - **NOTE:** 
                    - This is not actually the process of passing in through form-data, this is just an example showing what should be placed in the key and what should be placed in the value.
                    - sender_id, receiver_id, and message have a value type of "Text" while image has a value type of "File".
                    - The keys shown above **MUST** be the keys you use when passing with form-data.
                    - All other requirements and errors associated with the original **Create Message** apply here too.
                    - Once the message has been created, the image field will be given a URL that is associated with the BLOB. The URL can be converted into a SAS (Shared Access Signature) URL that allows you to access the image. A GET request can then be called to the SAS URL to retrieve the image (the response type should be BLOB). The function to turn it into an SAS URL can be found [here](#SAS).

    - **Third Option: Search for Message in All Conversations**: 
        - This **POST** function returns the message object that matches the specified message, if it exists in the messages database.
        - This function looks for a message across **ALL** conversations.
        - It will find messages that match it exactly or contain the search parameter somewhere within it. It is not case-sensitive.
        - **Example Syntax:**
        ```http://localhost:8080/api/messages/searchAll```
    
    - **Fourth Option: Search for Message in One Conversation**: 
        - This **POST** function returns the message object that matches the specified message, if it exists in the messages database.
        - This function looks for a message across **ONE** conversation.
        - It will find messages that match it exactly or contain the search parameter somewhere within it. It is not case-sensitive.
        - **Example Syntax:**
        ```http://localhost:8080/api/messages/[FIRST ID]/[SECOND ID]/search```

    - **NOTE:** For the **third** and **fourth** option, the search message should be placed in the body of the POST request:

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
- For the **Create Message with Image** function, an appropriate error message will be returned if:
    - The image cannot be retrieved.
    - The image file is too large (the image size should be 10 megabytes or fewer).
    - A shared key credential could not be made (used to connect to the Azure container).
    - The image could not be uploaded.
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
            - The logic is that, when a user edits a message, that message struct will already be known (through a GET function) and the unique GORM ID can be passed into this function. This ensures that the correct message is being edited.

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
        - **NOTE:** In order for this function call to work as a **PUT**, a request body must be passed in. In this case, you can simply pass in an empty request body.

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
- There are currently five different **GET** functions available:

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
        ```http://localhost:8080/api/messages/deleted ```

        - **NOTE:** _This is considered a testing function and not for Frontend purposes._

    <a id="SAS"></a>
    - **Fourth Option: Get SAS URL**: 
        - This **GET** function takes in the GORM ID of a message and converts the string in the "image" field from a BLOB URL to a SAS URL.
        - **Example Syntax:**
        ```http://localhost:8080/api/messages/getImage/URL/{id}```

        - An **SAS URL** will be returned in the form of a map (key/value pair). The key is "sasUrl" and is a string literal. The value of "sasURL" is a string variable that contains the SAS URL. A GET request can then be made to this SAS URL to retrieve the image (the response type should be BLOB).

    - **Fifth Option: Get Most Recent Conversation**: 
        - This **GET** function returns the user object of whoever the passed-in user talked to last.
         - **Example Syntax:**
        ```http://localhost:8080/api/messages/getRecent/user/{id} ```

        - This will take in the ID in the URL and search for the last message they sent. It will then return the user object of the receiver_id (i.e. the last person they talked to).

### Requirements and Error Messages
- The **"Get Conversation"** function must have a valid conversation that exists in the database, or else "Conversation not found." will be returned.
- If the **"Get ALL Messages, Get ALL Deleted Messages, and Get Most Recent Conversation"** functions cannot locate any messages, then a message describing how no messages were found will be returned.
- For the **Get SAS URL** function, the message with the passed-in GORM ID must exist or an error message saying "Message not found." will be returned.
- An **Internal Server Error** will be returned if there are errors regarding the database connection or the query itself.
    - If the **SAS query parameters** could not be established correctly, then an **Internal Server Error** will be returned.
- If all requirements are met, the message(s), URL, or user will be returned along with a successful console log message.
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

- The first **POST** command takes in an input and creates a new user in the users database. The second command is focused on a retrieving a specific user. _The second option appears to be more of a "GET", but a "GET" does not allow for a body request. Therefore, a POST is used._

### Syntax

- There are currently three POST commands available:

    - **First Option: Create a User:**
        - This **POST** function creates a user in the users database and returns the user object back to the requester.
        - **Example Syntax:**
        ```http://localhost:8080/api/users ```

        - For post, the information passed in must be through the request **body**.
        - The required inputs are a username, a password, a user ID, an email, and a list of current conversations that the user is in (typically left blank).
        - **NOTE:** 
            - The ID value must be **a number between 0000 and 9995** (9996 to 9999 are being reserved for the unit tests).
                - This function will handle finding a valid ID for the new user in the Backend by using an internal function called getNextUserID().
                - Therefore, the User_ID field can be left blank.
                - If you wish to make a user have a specific ID, then you can fill out the User_ID field (it will work assuming that there currently is not a user with that ID).
            - The email should not be **unitTest@ufl.edu** as this has been reserved for unit tests.
            - The phone number should not be **(000) 000-0000** or **(000) 000-0001** as these have been reserved for unit tests.
        
            - **Example Syntax:**
                ```
                {
                    "username": "student",
                    "password": "pass",
                    "user_id": "",
                    "email": "student@ufl.edu",
                    "full_name": "test user",
                    "phone_number": "(123) 456-7890",
                    "current_conversations": []
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

    - **Third Option: Search for a Specific User**: 
        - This **POST** function returns a singular user from the users database.
        - It will find a user that matches the search pattern ```[USERNAME]#[ID]```
        - **Example Syntax:**
        ```http://localhost:8080/api/users/search```

            - The required input is the search pattern, which will be placed in the **username** field.
            
                - **Example Syntax:**
                    ```
                        {
                            "username": "student#1234"
                        }
                    ```

### Requirements and Error Messages
- A StatusBadRequest error will be returned if the passed-in body cannot be decoded.
- The **Create User** function must have:
    - A **unique**, **four-digit ID** that is **less than 9996**.
    - The email must have a University of Florida **domain name** (i.e. it must end with "@ufl.edu").
        - The email must also be **unique**, in the sense that **no other existing accounts** should currently have that email.
    - The phone number must be in the form of **(###) ###-####**.
    - The phone number must be **unique** in the sense that it **does not** already exist in the users database.
        - It should not be **(000) 000-0000** or **(000) 000-0001**.
    - Otherwise, a 400 Bad Request will be returned.
- An **Internal Server Error** will be returned if there are errors regarding the database connection or the query itself (e.g. the requested user could not be found in the database).
- If all requirements are met, a user object will be returned along with a successful console log message.
---

<a id="PUT_Users"></a>

### ➜ Overview of  **PUT** Command for Users

### Syntax

- There are currently four **PUT** commands available:
        
     - **First Option: Edit Username**:
        - This **PUT** function edits a user's username.
         - **Example Syntax:**
            - ```http://localhost:8080/api/users/updateN/[ID]```

            - The required input is the user's new username.
                - **Example Syntax:**
                    ```
                        {
                            "username": "user"
                        }
                    ```

     - **Second Option: Edit Password**:
        - This **PUT** function edits a user's password.
         - **Example Syntax:**
            - ```http://localhost:8080/api/users/updateP/[ID]```

            - The required input is the user's original password and new password.
                - **Example Syntax:**
                    ```
                       {
                            "password": "newPass",
                            "original_pass": "pass"
                        }
                    ```
            - **NOTE:** 
                - For this function, the struct used **must be UserAccountConfirmPass, not UserAccount**. This is because the "original_pass" field is not included in the original **UserAccount** struct.
                - The new password **CANNOT** the same as the original password.

    - **Third Option: Edit Full Name**:
        - This **PUT** function edits a user's full name.
         - **Example Syntax:**
            - ```http://localhost:8080/api/users/updateFN/[ID]```

            - The required input is the user's new full name.
                - **Example Syntax:**
                    ```
                        {
                            "full_name": "Test User"
                        }

    - **Fourth Option: Edit Phone Number**:
        - This **PUT** function edits a user's phone number.
         - **Example Syntax:**
            - ```http://localhost:8080/api/users/updatePN/[ID]```

            - The required input is the user's new username.
                - **Example Syntax:**
                    ```
                        {
                            "phone_number": "(123) 456-7890"
                        }
                    ```
        - **NOTE:** 
            - The updated phone number must be unique in the sense that it **does not** already exist in the database. Otherwise, a 400 Bad Request will be returned. However, a user **CAN** update their phone number to their current phone number. 

### Requirements and Error Messages
- An **Internal Server Error** will be returned if it is unable to locate the passed-in user or if there are errors regarding the database connection.
    - This error will be thrown if the **Edit Username**, **Edit Password**, **Edit Full Name**, or **Edit Phone Number** functions are not able to locate the passed-in user.
        - Otherwise, the username, password, full name, or phone number will be updated and will return the newly updated user object along with a successful console log message.
- For the **Edit Phone Number** function, the phone number **STILL** must follow the format established in the create user function: (###) ###-####. Otherwise, a 400 Bad Request will be returned.
---

<a id="GET_Users"></a>

### ➜ Overview of **GET** Command for Users

- The **GET** command returns information about users that have been created with a **POST** request.

### Syntax
- There are currently three **GET** functions available:

    - **First Option: Edit Current Conversation**:
        - This **GET** function adds an ID to a user's current conversation list.
         - **Example Syntax:**
        - ```http://localhost:8080/api/users/[FIRST ID]/[SECOND ID]```

        - The required inputs are the user's ID (```FIRST ID```) and a second ID (```SECOND ID```).
            - This function will add ```SECOND ID``` into ```FIRST ID```'s conversation list and ```FIRST ID``` into ```SECOND ID```'s conversation list. This is done so when a user starts a conversation with someone else, that other person will also be able to see the conversation.
        - If all requirements are met, the updated user object will be returned along with a "ID added successfully." console log message.
        - **NOTE:** This function is called as a **GET** rather than a **PUT** because it does not involve passing anything into the body. Frameworks such as Angular must have a body for a **PUT**.

    - **Second Option: Get All Users**:
        - This **GET** function returns all users in the users database.
         - **Example Syntax:**
        ```http://localhost:8080/api/users```

        - **NOTE:** _It is expected that this function is merely a testing function and will not be implemented in the Frontend._

    - **Third Option: Get User by ID**:
        - This **GET** function returns a user from the users database based on the user's unique ID.
            - **Example Syntax:**
        ```http://localhost:8080/api/users/[ID]```

### Requirements and Error Messages
- A StatusBadRequest error will be returned if the passed-in body cannot be decoded.
- The **"Get All Users"** function must have users that exist in the database, or else "Users not found." will be returned.
- The **"Get User by ID"** function must have the requested user exist in the users dataabse, or else a StatusBadRequest error will be returned.
- If all requirements are met, the user(s) will be returned along with a successful console log message.
---
<a id="DELETE_Users"></a>

### ➜ Overview of  **DELETE** Command for Users

- The **DELETE** command takes in a user and deletes some aspect of their information.

### Syntax
- There are currently two DELETE functions available:

    - **First Option: Delete a Contact**:
        - This **DELETE** function removes an ID from a user's Current Conversations list.
        - **Example Syntax:**
       ```http://localhost:8080/api/users/removeC/9998/9999```
       
    - **Second Option: Delete a User**:
        - This **DELETE** function takes in a user and deletes them from the user database.
        - **Example Syntax:**
       ```http://localhost:8080/users/[USER_ID] ```

### Requirements and Error Messages
- An **Internal Server Error** will be returned if there are errors regarding the database connection or the query itself.
- If all requirements are met for deleteContact, the contact will be removed from the user's Current Conversation's list along with a "Contact removed successfully." console log message. 
- If all requirements are met for deleteUser, the user will be removed from the database along with a "User deleted successfully." console log message.
---

[Back to top](#TOC)

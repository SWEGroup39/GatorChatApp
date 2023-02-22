# GatorChat API Documentation

### Written by: Kevin Cen and John Struckman
---

## Table of Contents

- [Setting Up Golang](#settingUp)

- [Accessing the GatorChat API](#accessingAPI)

- [Overview of **REST** Functions](#REST)

- [Overview of  **POST** Commands](#POST)

- Overview of **PUT** Commands

- Overview of **GET** Commands

- Overview of **DELETE** Commands

---
<a id="settingUp"></a>

### ➜ Setting Up Golang

##### Please refer to the instructions linked below to properly install Go.

- Golang can be setup by following the instructions listed [here](https://github.com/rb-uf/swe-project/blob/emmett/go-setup.md).
- It is important to place your Go projects in a valid directory, for example:
    - ``` C:\Users\[USER]\go\src\github.com\kevinc3n\API ```
        - This ensures that Golang can find all of the packages and can properly run.

- **NOTE:** It is important to install the packages required to use the GatorChat API.
- These include:
  - **GORM**
      - GORM is needed for the "gorm.io/driver/mysql" and "gorm.io/gorm" packages.
  - **Gorilla Mux**
    - Gorilla Mux is needed for the "github.com/gorilla/mux" package.
- **Quick Reference**: Use ``` go get -u <package> ``` in your command line to install a certain package.
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
            - For example: ``` cd C:\Users\[USER]\Desktop\SWE\GatorChatApp```
        - Next, run the following command to have all of the branch's files be placed into your repository folder.
            - ``` git pull origin [BRANCH_NAME] ```
            - In this case, it is ``` git pull origin Back-End-Branch ```.
        - The folder should now contain the API file.

- To run the file, open the terminal in your respective IDE and run the following commands:
    - ```go build```
        - Once it has finished, run the command ```go run GatorChat_Rest_API.go```.
        - If running the code in VSCode you can run the command ```./(put the name of the .exe file that was made by go build here)``` instead
- The API should now be running. The localhost port should be active and able to receive requests.

---

<a id="REST"></a>

### ➜ Overview of REST Functions

- The GatorChat API is built upon the **REST** functions **POST**, **PUT**, **GET**, and **DELETE**.

- The API supports **POST** to create a message and store it in a database for later retrieval.
- The API supports **PUT** to update an existing message in the database with a new message.
- The API supports **GET** to retrieve messages from the databse based on certain parameters.
- The API supports **DELETE** to remove messages from the databse that are no longer needed.
<br>
- This API serves to provide an **abstraction** to the process of manually accessing the database and running MySQL commands to perform the actions listed above.

---

<a id="POST"></a>

### ➜ Overview of  **POST** Commands

- The **POST** command takes in an input and creates a new message with the data.
- Required inputs are a unique message ID number, a message string, a sender ID, and a receiver ID.
- The message ID number can't be reused unless the previous message with that number was hard-deleted
- Input "null" for the "CreatedAt", "UpdatedAt", and "DeletedAt" date inputs, these will be automatically filled in.

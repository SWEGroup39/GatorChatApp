# README File
## Instructions on How To Run Gator Chat Application (FRONT END)
  - In order to run this application, you first have to download and install angular. [click here](https://github.com/SWEGroup39/GatorChatApp/blob/main/README.md#installing-angular)
  - If you already have angular installed [click here to see how to run Gator Chat](https://github.com/SWEGroup39/GatorChatApp/edit/main/README.md#running-the-gator-chat-application)
  ### Installing Angular
  - Steps include:
    - Install Node/npm on your machine [Click here to see the steps/description](https://github.com/SWEGroup39/GatorChatApp/edit/main/README.md#nodejs-installation)
    - Use and install Angular CLI globally [Click here to see the steps/description](https://github.com/SWEGroup39/GatorChatApp/edit/main/README.md#angular-cli-installation)
    - Run Angular CLI commands [Click here to see the steps/description](https://github.com/SWEGroup39/GatorChatApp/edit/main/README.md#angular-cli-installation)
    - Create an initial workspace for the application [Click here to see the steps/description](https://github.com/SWEGroup39/GatorChatApp/edit/main/README.md#creating-a-test-project-using-angular-cli)
    - Run the Angular application in Browser [Click here to see the steps/description](https://github.com/SWEGroup39/GatorChatApp/edit/main/README.md#running-your-angular-app)
    #### NodeJS installation:
      - Go to the NodeJS website to download it <a href="https://nodejs.org/en">Click here</a>
        - You can install either version based on your needs
        - I have installed the LTS version
      - In order to check if nodeJS has been properly installed run this command in your CLI: ***node -v*** or ***node --version***
      - NodeJS also automatically downloads npm for you. In order to check if npm has been installed run the command: ***npm -v***
    #### Angular CLI installation:
      - To install the Angular CLI on your machine, open the terminal window and run the following command: ***npm install -g @angular/cli***
      - You can verify if this was properly installed by running the command: ***ng version***
    #### Creating a Test Project using Angular CLI  
      - Open the terminal window and type the command: ***ng new hello-world***
        - hello-world is the project name so you can replace it with any name
      - After it is done running, open the directory in any code editor that you like
        - I am using VS code as my IDE
    #### Running your Angular App
      - In order to run your application, first make sure you are in the correct file path. 
        - My file path to run the application is: ***PS C:\Users\Ria Chacko\chatG-app***
      - Then run the command ***ng serve**
        - This will run the application on the default 4200 port number
        - If you want to run the application on a different port number then the command is: ***ng serve --portnumber***
          - The command I use to run the application is: ***ng serve --1655***
        - Once everything is compiled, the terminal will give you a url to access the application. You can either click it or copy it and paste into any browser of your choice
          - The preferred browser for this app is Google Chrome
        - Also adding the CORS policy blocker prevents any CORS errors from occuring 
          - This is a Google Chrome extension that can be downloaded
          
  ### Running the Gator Chat Application
   - First create an angular project with the command: ng new project-name
   - In order to run our application, navigate to the main branch App contents folder, then to the Front-End folder. <a href="https://github.com/SWEGroup39/GatorChatApp/tree/main/App_Contents/FrontEnd">Click Here</a>
     - Copy all of the files and folders **excluding** the Cypress Tests folder and place them into the newly created angular project
     - NOTE: **Before** running the application make sure to run the API file [Click here to see the steps](https://github.com/SWEGroup39/GatorChatApp/edit/main/README.md#instructions-on-how-to-run-api-for-gator-chat-back-end)
     - Then go to the terminal and run the ng serve command to run the application
     - This should take the user to the home page initially.
       - From this point, you have successfully run Gator Chat
       
## Instructions on How To Run API for Gator Chat (BACK END)

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

# README File
## Instructions on How To Run Gator Chat Application (FRONT END)
  - In order to run this application, you first have to download and install angular. <a href="Installing Angular">Click here</a>
  - If you already have angular installed <a href="Running the Gator Chat Application">click here to see how to run Gator Chat</a>
  ### Installing Angular
  - Steps include:
    - Install Node/npm on your machine <a href="NodeJS installation">Click here to see the steps/description</a>
    - Use and install Angular CLI globally <a href="Angular CLI installation">Click here to see the steps/description</a>
    - Run Angular CLI commands <a href="Angular CLI installation">Click here to see the steps/description</a>
    - Create an initial workspace for the application <a href="Running your Angular App">Click here to see the steps/description</a>
    - Run the Angular application in Browser <a href="Running your Angular App">Click here to see the steps/description</a>
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
   - In order to run our application, navigate to the main branch App contents folder, then to the Front-End folder.
     - Copy all of the files and folders **excluding** the Cypress Tests folder and place them into the newly created angular project
     - NOTE: **Before** running the application make sure to run the API file <a href="Instructions on How To Run API for Gator Chat (BACK END)">Click here to see the steps</a>
     - Then go to the terminal and run the ng serve command to run the application
     - This should take the user to the home page initially.
       - From this point, you have successfully run Gator Chat
       
## Instructions on How To Run API for Gator Chat (BACK END) 

- Project Name: GatorChat

- Project Members: Rafael Sevilla, Kevin Cen, John Struckman, Ria Chacko

- Front-end Members: Ria Chacko, Rafael Sevilla

- Back-end Members: Kevin Cen, John Struckman

- Project Description:

  - Create a full-stack application that operates as a general messaging platform for users.

    - Below are a list of general/required features we would like to implement (should be able to implement within our time frame):

      - Login system with authentication. - Kevin
      - Send text messages to other contacts on the application. - Kevin
        - Messages have timestamps associated with them. - Rafael
      - Can send images. - Kevin
      - Able to edit/delete messages. - Kevin
      - Ability to search past messages. - Rafael
      - List of contacts with easy contact import and editing functionality. (Contact Syncing) - Ria
      - Notifications, unread message counts, and/or message states (read/unread). -Ria
      - User presence indication (available, away, offline, time last active). - Ria
      - User Profile - Ria

    - Below are additional features that we would implement if time permits:

      - Have the messaging app be geared towards the computer science community or UF students in general. - Kevin

        - If we went down the CS community track, the features we would add would be:

          - Prioritize the app to be able to share information such as code in a visually appealing way (make it look like an idea). - Kevin
          - Syntax highlighting for code snippets that are shared within the messaging app. - Kevin
          - A built-in code editor for writing/sharing code within the app. - Kevin
          - An ability to create and share diagrams and flowcharts. - Kevin
          - An ability to organize group chats with specific users. - John

      - Other general additional features include:

        - Ability to send voice notes. - Rafael
        - Ability to reply to a message with emojis. - Kevin
        - Message Disappearing - Ria
        - Chatbots - Ria
        - Dark/Light Mode - Ria
        - Read Receipts - Ria
        - Profile Customization - Ria
        - Location tracking and sharing - Ria

///<reference types="cypress"/>
  it('Get - Search', () =>{
    cy.visit('http://localhost:3200')
    cy.get('#chat-searcharea').type('Hello');
    cy.get('#search-button').click();
    
  })


  
  
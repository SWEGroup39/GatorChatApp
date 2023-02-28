///<reference types="cypress"/>
it('Chat App Access', function(){

  cy.visit('http://localhost:3200/')

  cy.get('#chat-textarea').type('Hello! Cypress Testing!')
  cy.get('#send-button').click()
})
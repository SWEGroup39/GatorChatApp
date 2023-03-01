///<reference types="cypress"/>
it('Chat App Access', function(){

  cy.visit('http://localhost:4200/')
  cy.get('#username').type('harry.k')
  cy.get('#password').type('harry')
  cy.get('.login-form > .mdc-button > .mdc-button__label').click()
})
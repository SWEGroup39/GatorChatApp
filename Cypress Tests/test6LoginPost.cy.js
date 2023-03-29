///<reference types="cypress"/>
it('Chat App Access', function(){
  cy.visit('http://localhost:1655/login')
  cy.get('#mat-input-0').type('test@ufl.edu')
  cy.get('#mat-input-1').type('pass')
  cy.get('.login-form > .mdc-button').click()
})
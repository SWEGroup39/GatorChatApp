///<reference types="cypress"/>
it('Chat App Access', function(){
  cy.visit('http://localhost:1655/signup')
  cy.get('#mat-input-0').type('harry@ufl.edu')
  cy.get('#mat-input-1').type('bob')
  cy.get('#mat-input-2').type('janet')
  cy.get('#mat-input-3').type('janet')
  cy.get('.login-form > .mdc-button ').click()

})
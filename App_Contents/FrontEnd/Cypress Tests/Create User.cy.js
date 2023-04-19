///<reference types="cypress"/>
it('Chat App Access', function(){

    cy.visit('http://localhost:4200/signup')
    cy.get('#mat-input-0').type('test8@ufl.edu')
    cy.get('#mat-input-1').type('bobby')
    cy.get('#mat-input-2').type('pass')
    cy.get('#mat-input-3').type('pass')
    cy.get('#mat-input-4').type('bobby gator')
    cy.get('#mat-input-5').type('(231) 234-3245')
    cy.get('.login-form > .mdc-button > .mdc-button__label').click()
    
})
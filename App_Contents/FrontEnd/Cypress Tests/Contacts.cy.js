///<reference types="cypress"/>
it('Chat App Access', function(){

    cy.visit('http://localhost:4200/login')
    cy.get('#mat-input-0').type('user@ufl.edu')
    cy.get('#mat-input-1').type('pass')
    cy.get('.login-form > .mdc-button').click()
    cy.get('.position-sticky > .nav > :nth-child(2) > .nav-link').click()
    cy.wait(10000)
    cy.get('.mdc-button__label').click()

   
})
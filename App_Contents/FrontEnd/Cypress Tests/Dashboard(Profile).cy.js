///<reference types="cypress"/>
it('Chat App Access', function(){

    cy.visit('http://localhost:4200/login')
    cy.get('#mat-input-0').type('user@ufl.edu')
    cy.get('#mat-input-1').type('pass')
    cy.get('.login-form > .mdc-button').click()
    // cy.get('[routerlink="/contacts"] > .d-flex > div > .opacity-75').click()
    cy.get(':nth-child(3) > .nav-link').click()
    cy.get('.mdc-button__label').click()
   
})
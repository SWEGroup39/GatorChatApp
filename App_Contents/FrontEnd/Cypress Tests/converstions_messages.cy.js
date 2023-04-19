///<reference types="cypress"/>
it('Chat App Access', function(){

    cy.visit('http://localhost:4200/login')
    cy.get('#mat-input-0').type('usernew@ufl.edu')
    cy.get('#mat-input-1').type('newPass')
    cy.get('.login-form > .mdc-button').click()
    
    // cy.get('[routerlink="/contacts"] > .d-flex > div > .opacity-75').click()
    cy.get('.position-sticky > .nav > :nth-child(1) > .nav-link').click()
    cy.get(':nth-child(2) > button').click()
    cy.get('#chat-searcharea').click()
    cy.get('#chat-searcharea').type('testing again')
    cy.get('#search-button > img').click()
    // cy.get(':nth-child(12) > .message-container > #message-body')
    // cy.get(':nth-child(12) > .message-container > #message-body')z
    cy.get(':nth-child(12) > .message-container > #message-body').click()
    cy.get('ul > .ng-star-inserted').click()
    cy.get('#edit-text-area').type('hey0')
    cy.get('#save-button').click()
    cy.get(':nth-child(14) > .message-container > #message-body').click()
    cy.wait(7000)
    cy.get('#image-button > img').click()
    cy.wait(4000)
    cy.get('#chat-textarea').click()
    cy.get('#chat-textarea').type('hey')
    cy.get('#send-button').click()

   
})
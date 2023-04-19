///<reference types="cypress"/>
it('Chat App Access', function(){

    cy.visit('http://localhost:4200/login')
    cy.get('#mat-input-0').type('usernew@ufl.edu')
    cy.get('#mat-input-1').type('newPass')
    cy.get('.login-form > .mdc-button').click()
    
    // cy.get('[routerlink="/contacts"] > .d-flex > div > .opacity-75').click()
    cy.get('[routerlink="/settings"] > .d-flex > div > .opacity-75').click()
    cy.get('[style="background-color: rgb(255, 251, 251); height: 100vh; width: 100vw; margin: 0; padding: 0; border: 0; outline: 0;"]').click()
    cy.wait(10000)
    cy.get('.pt-md-0 > [style="width: 50px; height: 30px; background-color: rgb(95, 100, 223); color: white;"]').click()
    cy.wait(10000)
    cy.get('.btn-primary').click()
})
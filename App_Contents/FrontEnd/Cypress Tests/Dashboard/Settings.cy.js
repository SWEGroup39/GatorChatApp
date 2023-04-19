///<reference types="cypress"/>
it('Chat App Access', function(){

    cy.visit('http://localhost:4200/login')
    cy.get('#mat-input-0').type('usernew@ufl.edu')
    cy.get('#mat-input-1').type('newPass')
    cy.get('.login-form > .mdc-button').click()

    // cy.wait(10000)
    // cy.get('[routerlink="/contacts"] > .d-flex > div > .opacity-75').click()
    
    cy.get(':nth-child(4) > .nav-link').click()
    cy.get('[style="float: left;"] > [style="width: 50px; height: 30px; background-color: orange; margin-right: 10px;"] > .mdc-button__label').click()
    
})
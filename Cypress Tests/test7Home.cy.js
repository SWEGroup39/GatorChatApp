///<reference types="cypress"/>
it('Chat App Access', function(){
  cy.visit('http://localhost:1655/home')
  cy.get('.d-grid > [routerlink="/about"]').click()
})
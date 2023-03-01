///<reference types="cypress"/>

it('Get-Search-API', () =>{
  cy.visit('http://localhost:3200/')
  cy.request('GET', 'http://localhost:8080/api/messages/Hello').then((response) => {
    expect(response).to.have.property('status', 200)
  })
})
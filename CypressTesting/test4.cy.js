///<reference types="cypress"/>

it('Get-Send-API', () =>{
  cy.visit('http://localhost:3200/')
  const message = {"messageID" : "87", "message" : "How is the weather?", "sender_id": "1234", "receiver_id": "5678"};
  cy.request('POST', 'http://localhost:8080/api/messages', message)
  .its('body')
  .should('include', {message: 'How is the weather?'})
   
})
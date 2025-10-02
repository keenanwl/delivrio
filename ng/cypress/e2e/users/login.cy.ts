describe('login form', () => {
	it('rejects invalid password for valid/invalid user', () => {

		cy.task('db:seed')
		cy.visit('/');
		cy.get('input')
			.eq(0)
			.type('klinsly@gmail.com')

		cy.get('input')
			.eq(1)
			.type('!123456789!Aa')

		cy.get('button')
		  .contains('Login')
		  .click();

		cy.get('div')
			.contains('Your login credentials were not accepted')
			.should('exist');

		cy.get('input')
		  .eq(0)
		  .clear()
		  .type('not-exists@gmail.com')

		cy.get('button')
		  .contains('Login')
		  .click();

		cy.get('div')
		  .contains('Your login credentials were not accepted')
		  .should('exist');
	})

	it('should login as the credentials are valid (server now() = June 22nd)', () => {
		cy.visit('/');
		cy.get('input')
			.eq(0)
			.type('bb@example.com')

		cy.get('input')
			.eq(1)
			.type("000111kkddmmmaasasÆÆ!`'~^~¨")

		cy.get('button')
		  .contains('Login')
		  .click();

		cy.location().should((loc) => {
			expect(loc.pathname).to.eq('/dashboard')
		});
	})
})

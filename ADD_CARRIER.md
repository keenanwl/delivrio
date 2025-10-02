# Add a new carrier

1. Generate `cairrer<new carrier>`, `cairrerservice<new carrier>`, `cairreradditionalservice<new carrier>` & `shipment<new carrier>`
   2. Add the same edges as an existing carrier
3. Add New `carrierbrand` to `go/seed/seed.go`
4. Update `go/carriers.resolvers.go` -> `CreateCarrierAgreement`
5. Add CRUD to `go/carriers.graphql`
6. Add switch case to `ng/app/settings/carrier-list.ngxs` & `ng/app/settings/carrier-list.component`
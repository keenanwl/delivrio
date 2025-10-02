package main

//go:generate go run -mod=mod generate/entc.go
//go:generate go get github.com/99designs/gqlgen@v0.17.48
//go:generate go run github.com/99designs/gqlgen graphql

////go:generate npm run-script gen:graphql --prefix ../ng
////go:generate npm run-script gen:endpoints --prefix ../ng
//// go run -mod=mod entgo.io/ent/cmd/ent new --template=schema/templates/entinit.tmpl --target schema CarrierServiceBring CarrierAdditionalServiceBring DeliveryOptionBring ShipmentBring ParcelShopBring CarrierBring

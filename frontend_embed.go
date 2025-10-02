package delivrio

import "embed"

//go:embed ng/dist/delivery/browser/*
var FrontendContent embed.FS

//go:embed ng/dist/return-portal/browser/polyfills.js
var FrontendReturnPolyfills string

//go:embed ng/dist/return-portal/browser/main.js
var FrontendReturnMain string

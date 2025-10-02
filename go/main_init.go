package main

import (
	"delivrio.io/go/appconfig"
	"delivrio.io/go/carrierapis/bringapis"
	"delivrio.io/go/carrierapis/daoapis"
	"delivrio.io/go/carrierapis/dfapis"
	"delivrio.io/go/carrierapis/postnordapis"
	"delivrio.io/go/deliverypoints/glsdeliverypoints"
	"delivrio.io/go/deliverypoints/postnorddeliverypoints"
	"delivrio.io/go/endpoints"
	"delivrio.io/go/gengql"
	"delivrio.io/go/mergeutils"
	"delivrio.io/go/schema/hooks/connectionhooks"
	"delivrio.io/go/schema/hooks/dlvshopify"
	"delivrio.io/go/utils"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"log"
	"net/url"
	"time"
)

// TODO: figure out a better way to handle
// static global state. Seems excessive to add it
// to every ctx.
func setConf(conf *appconfig.DelivrioConfig) {
	// Unsure if we can just pass them directly
	// to the package with ldflags?
	appconfig.BuildTime = BuildTime
	appconfig.AppVersion = AppVersion
	appconfig.AppName = conf.ServerID

	connectionhooks.Init(conf)
	mergeutils.Init(conf)
	postnordapis.Init(conf)
	bringapis.Init(conf)
	gengql.Init(conf)

	// Delivery points
	postnorddeliverypoints.Init(conf)
	daoapis.Init(conf)
	dfapis.Init(conf)
	glsdeliverypoints.Init(conf)

	utils.Init(conf)
	baseURL, err := url.Parse(conf.BaseURL)
	if err != nil {
		log.Printf("invalid conf: %v\n", err)
	}
	dlvshopify.BaseURL = baseURL
	endpoints.BaseURL = baseURL
	endpoints.AppConfig = conf
	tokenAuth = jwtauth.New(
		"HS256",
		[]byte(conf.JWTKey),
		nil,
		jwt.WithAcceptableSkew(30*time.Second),
		jwt.WithClock(Clock),
	)
}

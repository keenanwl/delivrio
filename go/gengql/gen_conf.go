package gengql

import "delivrio.io/go/appconfig"

var conf *appconfig.DelivrioConfig
var confSet = false

func Init(c *appconfig.DelivrioConfig) {
	if confSet {
		panic("gengql: may not set config twice")
	}
	conf = c
	confSet = true
}

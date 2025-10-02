# DELIVRIO
## What is DELIVRIO?
DELIVRIO is shipping platform (TMS) for e-commerce focused businesses and 3PLs 
with the goal of providing a highly scalable core system and real-time analytics modules.

## Major Features
* Shopify support: 
  * sync orders & products
  * display highly customizable shipping rates in checkout
* Unified API - don't worry about differences in carrier APIs
  * All labels can be requested as PDF or ZPL
* Returns - easily trackable at SKU level
* Carrier support:
  * PostNord
  * GLS
  * USPS
  * Bring
  * Danske Fragtmænd

Checkout our roadmap to see the detailed feature list and what features are on the way

## Installation
| -                                | Managed | Self-hosted | Desktop |
|----------------------------------|---------|-------------|---------|
| **Tech Skills required**         | No      | Yes         | Basic   |
| **Delivery options in checkout** | Yes     | Yes         | No      |
| **Always up to date**            | Yes     | No          | No      |
| **Sync**                         | Yes     | Yes         | Yes     |
| **Fulfillment**                  | Yes     | Yes         | Yes     |
### DELIVRIO managed cloud (no technical skills required)
* Simply create an account to get started

### Local (requires basic technical skills)
Install DELIVRIO on your personal computer to quickly manage order fulfillment.
* Supports Windows, Mac & Linux
* Get running quickly
* Sync orders and products

(Some features are disabled since your computer is not accessible to the internet)

### Self-hosted (requires technical experience)
Install DELIVRIO on Linux Ubuntu to use all features
* Easy install via package manager
* No External dependencies/telemetry for maximum compliance

## Optional Dependencies
These dependencies are only required if you plan on printing documents
(like packing slips or customs). 
Labels printed as PDF's do not use these dependencies.
### Gotenburg for documents HTML -> PDF converter
Setup an external service by following this guide and
adding access to the DELIVRIO config file.

https://gotenberg.dev/docs/getting-started/installation

Known to work with v6.x

### PDFium for PNG & ZPL labels
Download the webassembly binary (easier install vs added performance of native)

https://github.com/klippa-app/pdfium-cli/releases

Install:
`sudo install pdfium-webassembly-linux-x64 /usr/local/bin/pdfium`

Known to work with v0.9.1


## Print Client
Install the DELIVRIO print client to support efficient scanning & printing
* Supports all install options above
* Supports Windows, Mac & Linux

| -             | Windows              | Mac      | Linux    |
|---------------|----------------------|----------|----------|
| **Network**   | Yes                  | Yes      | Yes      |
| **Local/USB** | X<sup>[183](https://gitlab.com/keenanwl/delivery/-/issues/183)</sup> | Yes/CUPS | Yes/CUPS |

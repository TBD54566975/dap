# Decentralized Agnostic Paytag (DAP) <!-- omit in toc -->

> [!WARNING]
> 🚧 WIP 👷‍♂️


# Table of Contents <!-- omit in toc -->
- [Problem Statement](#problem-statement)
- [Objective](#objective)
- [Requirements](#requirements)
- [TL;DR How it Works](#tldr-how-it-works)
- [DAP](#dap)
  - [Local Handles](#local-handles)
  - [Domain](#domain)
- [Resolution](#resolution)
- [Money Address](#money-address)
    - [Examples](#examples)
      - [USDC on Ethereum](#usdc-on-ethereum)
      - [BTC LNURL](#btc-lnurl)
      - [BTC Address](#btc-address)
      - [KES Mobile Money](#kes-mobile-money)
    - [Currency Specific Parts](#currency-specific-parts)
      - [USDC](#usdc)
        - [Network specific Address](#network-specific-address)
          - [Format](#format)
          - [Examples](#examples-1)
      - [BTC](#btc)
        - [LNURL](#lnurl)
          - [Format](#format-1)
        - [Address](#address)
          - [Format](#format-2)
  - [DID Resource](#did-resource)
    - [Examples](#examples-2)
      - [Example Service Entry](#example-service-entry)
      - [Example DID Document](#example-did-document)
- [DAP Registry](#dap-registry)
  - [Registry HTTP API](#registry-http-api)
    - [CORS Policy](#cors-policy)
    - [Responses](#responses)
    - [Error](#error)
    - [DAP Resolution](#dap-resolution)
      - [Request](#request)
      - [Response](#response)
        - [Errors](#errors)
          - [Not Found](#not-found)
- [Privacy Considerations](#privacy-considerations)
- [FAQ](#faq)
  - [How is this different from UMA (Universal Money Address)?](#how-is-this-different-from-uma-universal-money-address)



# Problem Statement
Countless applications exist worldwide (e.g. CashApp, YellowPay, Bitso, Coinbase, Chipper Cash, Paypal, Venmo etc.) that allow individuals to send and receive money using a variety of payment methods and financial instruments. These apps often provide a P2P function that allows individuals using the same app to send and receive money to/from one another using memorable handles, paytags or aliases. These paytags however are only useful within the context of the app they are created in and cannot be used to send and receive money to/from individuals using other apps or services. Sending the same currency accross different apps or payment networks is a notoriously cumbersome and error-prone process that often involves copy / pasting long, confusing, and nuanced mumbo jumbo. Further, depending on the type of currency being used, the process can be even more complex. For example: 
* sending BTC on Lightning is different than sending BTC on mainnet
* sending USDC on Ethereum has different details and nuances than sending USDC on Stellar.
* sending fiat money to a mobile money wallet is different than sending it to a bank account.

What this leaves us with are many large payment networks that all exist on their own islands. Life is great on the island but getting money onto and off the island is a "yeet and hope for the best" kind of situation. Further, indviduals are left to navigate and understand the nuances of whatever currency they're hoping to use. This is a problem because it makes it difficult for individuals to use the currency of their choice in the way that they want to. Moreover, an inevitable lengthy conversation occurs between two people prior to sending any money, in order to figure out what the hell one another even has or can use to send and receive money. As a result, individuals end up with 18 wallets and a tense negotiation about which one to use everytime they need to send or receive.

# Objective

> [!IMPORTANT]
> Alice should be able to send money to Bob using a memorable handle (aka DAP) regardless of what app or currency is being used. A DAP should be as simple as an email address: `handle@domain` e.g. `moegrammer@didpay.me`. 

The conversation between Alice and Bob should be as simple as:
```
Alice: Yo Bobby boy! Thanks for the coffee. What's your DAP so i can pay you back?
Bob: Anytime. DAP me up at waterboy@cash.app
```

Alice should then be able to pop open whatever app she uses, type in bob's DAP, and send him the money. Bob should then be able to receive the money in whatever app he uses. Hypothetical examples of this are:
* CashApp -> YellowPay (or vice versa)
* Chipper Cash -> Bitso (or vice versa)
* Coinbase -> CashApp (or vice versa)
* CashApp -> Self Custodial Wallet (or vice versa)
* Paper Wallet -> Paper Wallet


More concretely, The objective of this specification is to provide a standardized means to: 
* express a _money address_
* associate any number of money addresses to a resolvable identifier (DID)
* register the identifier with a handle at any/many DAP registries (e.g. registering `did:dht:3456787654323rfvcxaw34tf` with CashApp under the handle `moegrammer` results in `moegrammer@cash.app`)
* resolve a DAP to a DID

# Requirements
* _Any_ currency should be expressable as a _money address_
* _Any_ individual or institution **MUST** be able to run their own DAP Registry
* An individual **MUST** have the ability to custody and control their own resolvable identifier

# TL;DR How it Works

![](./braindump.png)


# DAP

A Decentralized Agnostic Paytag (DAP) is a human-friendly email-like handle that can be shared between individuals for the purposes of sending and recieving money. More specifically, a DAP can be _resolved_ to retrieve information about what currencies and payment networks an individual can receive money on.

A DAP resolves to a Decentralized Identifier ([DID](https://www.w3.org/TR/did-core/)) by the registry associated to the domain portion of the DAP. The resulting DID is then resolved to retrieve all of the [money addresses](#money-address) that have been associated with the respective DID.

> [!IMPORTANT]
> DAPs have nothing to do with the actual mechanics of sending and receiving money. They are simply a means to easily retrieve information about how to send and receive money to/from an individual.

A DAP is composed of a _local handle_ and a _domain_ and is structured as follows:

```
local-handle@domain
```

## Local Handles
A handle is unique to the _domain_ that it is registered at. The handle itself is _NOT_ globally unique. The entire DAP itself however, is globally unique. As such, restrictions on the format are left to the _domain_ that the DAP is registered at. If no pre-existing format is defined, the following is recommended:

- **UTF-8 Encoding**: All characters should be UTF-8 encoded to support internationalization.
- **Character Exclusions**: The local handle must not include [unicode control characters](https://en.wikipedia.org/wiki/Unicode_control_characters) or [punctuation characters](https://en.wikipedia.org/wiki/Template:Punctuation_marks_in_Unicode).
- **Length Restrictions**: The local handle must be between 3 to 30 characters in length.

> [!NOTE] 
> providing the flexibility for domains to define their own formats was decided upon in order to support pre-existing systems that already have their own formats. 

## Domain
The domain portion of a DAP is a string that represents the _domain_ that the DAP is registered at. The domain is used to resolve the DAP to a DID. 

# Resolution

The following steps are taken to resolve a DAP:
1. split the DAP into the local handle and domain using '@' as the delimiter
2. construct a `did:web` DID using the domain as the method specific identifier
3. [Resolve](https://www.w3.org/TR/did-core/#did-resolution) the resulting DID to retrieve the [DID Document](https://www.w3.org/TR/did-core/#dfn-did-documents).
4. Find the `service` of type `dapregistry` in the [DID Document](https://www.w3.org/TR/did-core/#dfn-did-documents)
5. Use the `serviceEndpoint` of the resulting service to construct the following URL `<serviceEndpoint>/daps/<local-handle>` 
6. Make a GET request to the constructed URL
7. The response will contain the DID associated to the DAP
8. [Resolve](https://www.w3.org/TR/did-core/#did-resolution) the DID to retrieve the [DID Document](https://www.w3.org/TR/did-core/#dfn-did-documents)
9. Find all of the `maddr` services in the [DID Document](https://www.w3.org/TR/did-core/#dfn-did-documents)

> [!TIP]
> IDs can be used as part of the DAP protocol to specify a payment address for a DID subject e.g. `moegrammer#munn@didpay.me`

# Money Address

A money address is a [Uniform Resource Name (URN)](https://datatracker.ietf.org/doc/html/rfc8141) that represents a means through which an individual can be payed a specific currency. The URN is structured as follows:

```
urn:<currency_code>:<curr_specific_part>
```

### Examples

#### USDC on Ethereum

`urn:usdc:eth:0x1234567890abcdef1234567890abcdef12345678`

#### BTC LNURL

`urn:btc:lnurl:https://someurl.com`


#### BTC Address

`urn:btc:addr:1LMcKyPmwebfygoeZP8E9jAMS2BcgH3Yip`

#### KES Mobile Money

`urn:kes:momo:mpesa:254712345678`

### Currency Specific Parts

> [!IMPORTANT] 
> from a standards perspective, a Currency Secific Part is an arbitrary string. There are no constraints or limitations. However, in order for Money Addresses to be of any practical use, conventions will need to be established for each currency. A separate registry will be maintained for currency specific parts as these conventions emerge.

This specification proposes Currency Specific Parts for USDC and BTC.

#### USDC

##### Network specific Address

###### Format
`urn:usdc:<network>:<address>`

###### Examples
| URN                                                       | CSP                                              | Description              |
| :-------------------------------------------------------- | :----------------------------------------------- | :----------------------- |
| `urn:usdc:eth:0x1234567890abcdef1234567890abcdef12345678` | `eth:0x1234567890abcdef1234567890abcdef12345678` | USDC address on Ethereum |
| `urn:usdc:xlm:0xff35866aCb80ce4b169d1460cd48108955c1c445` | `xlm:0xff35866aCb80ce4b169d1460cd48108955c1c445` | USDC address on Stellar  |

#### BTC

##### LNURL

###### Format
`urn:btc:lnurl:<url>`

##### Address

###### Format
`urn:btc:addr:<address>`

## DID Resource

A Money address is associated to a DID by representing it as a [Service](https://www.w3.org/TR/did-core/#services) entry on a [DID Document](https://www.w3.org/TR/did-core/#dfn-did-documents).

> [!NOTE]
> [Services](https://www.w3.org/TR/did-core/#services) are used in DID documents to express ways of communicating with the respected DID subject. A service can be any type of service the DID subject wants to advertise, including decentralized identity management services for further discovery, authentication, authorization, or interaction.

Money addresses are represented as a service of type `maddr` in the DID Document. The service is structured as follows:

| field             | value      | description                                                                |
| :---------------- | :--------- | :------------------------------------------------------------------------- |
| `type`            | `maddr`    | **MUST** always be `maddr`                                                 |
| `id`              | `[]string` | an abitrary string that can be used to uniquely identify the money address |
| `serviceEndpoint` | `[]string` | 1 or more payment address URNs                                             |

> [!NOTE]
> Any number of money addresses can be associated with a DID. They can be represented as individual service entries in the DID Document.

### Examples

#### Example Service Entry
```json
{
  "type": "maddr",
  "id": "#some-id",
  "serviceEndpoint": ["urn:<currency_code>:<curr_specific_part>"]
}
```

#### Example DID Document

> [!NOTE]
> Other fields in the DID Document have been omitted for brevity

```json
{
  "@context": "https://www.w3.org/ns/did/v1",
  "id": "did:dht:123456789abcdefghi",
  "service": [
    {
      "type": "maddr",
      "id": "#some-id",
      "serviceEndpoint": ["urn:usdc:eth:0x1234567890abcdef1234567890abcdef12345678"]
    }
  ]
}
```

# DAP Registry

A DAP Registry is a service that is responsible for maintaining a mapping of local handles to DIDs registered at the domain hosting the registry. The service provides two primary operations: 
1. Associate a local handle with a DID (a.k.a DAP Registration)
2. Resolve a DAP to a DID. 

A DAP Registry can be hosted by any individual or institution that controls a domain. In order to provide a DAP Registry a domain **MUST**:
* host a resolvable [DID Document](https://www.w3.org/TR/did-core/#dfn-did-documents) per the [`did:web` specification](https://w3c-ccg.github.io/did-method-web/). 
    * e.g. `did:web:cash.app` resolves to `https://cash.app/.well-known/did.json`
* advertise the DAP Registry as a service in the domain's DID Document e.g.

```json
{
  "id": "did:web:cash.app",
  "service": [
    {
      "type": "dapregistry",
      "serviceEndpoint": ["https://dap.cash.app"],
      "id": "#some-id"
    }
  ]
}
```

> [!NOTE]
> the value of `serviceEndpoint` is the base URL of the domain's DAP Registry HTTP API and can be any URL that the domain owner chooses.

## Registry HTTP API

### CORS Policy
The DAP Registry **MUST** have a [CORS policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) that allows requests from any origin. This is to ensure that a registry can be accessed by any app or service that wishes to resolve a DAP or facilitate registration (if the registry allows it).

### Responses
All responses from the DAP Registry **MUST** be a JSON Object that contains the following entries:

| Field   | Data Type         | Required | Description                                                  |
| :------ | :---------------- | -------- | :----------------------------------------------------------- |
| `data`  | `any`             | N        | set to whatever an endpoint is supposed to return on success |
| `error` | [`Error`](#error) | N        | set if an error occurs                                       |

> [!IMPORTANT]
> `data` and `error` are mutually exclusive. Only one of them should be present in a response.

### Error
An `Error` object is a JSON object that contains the following entries:

| Field     | Data Type | Required | Description                                     |
| :-------- | :-------- | -------- | :---------------------------------------------- |
| `message` | `string`  | Y        | human readable message that describes the error |


The following headers **MUST** be included in every response:

| header         | value              |
| :------------- | :----------------- |
| `Content-Type` | `application/json` |



### DAP Resolution

#### Request
**Method**: `GET`

---

**URL**: `<serviceEndpoint>/daps/<local-handle>`

---

#### Response

| Field   | Data Type | Required | Description                                           |
| :------ | :-------- | :------- | :---------------------------------------------------- |
| `did`   | `string`  | Y        | The DID associated to the DAP provided in the request |
| `proof` | `string`  | N        | Signed payload sent when the DID was registered       |

##### Errors

###### Not Found

**Status Code**: `404`



# Privacy Considerations

> [!WARNING]
> TODO: Fill out


# FAQ

## How is this different from UMA (Universal Money Address)?

> [!WARNING]
> TODO: Fill out
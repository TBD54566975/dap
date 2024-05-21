Decentralized Agnostic Paytag (DAP) 1.0
========================================

**Specification Status:** Draft

**Latest Draft:** https://dap.did.me

**Editors:**
~ [TODO]

**Contributors:**
~ [TODO]

**Participate:**
~ [GitHub repo](https://github.com/TODO)
~ [File a bug](https://github.com/TODO/issues)
~ [Commit history](https://github.com/TODO/commits/main)

## Abstract

Decentralized Agnostic Paytags (DAPs) provide a standardized means to express human-friendly, resolvable payment handles that abstract over various currencies and payment networks. DAPs allow individuals to register memorable handles that resolve to a [Decentralized Identifier](https://www.w3.org/TR/did-core/) containing one or more associated payment addresses. This enables easy cross-app and cross-currency payments using a simple identifier. The syntax and data model used by DAPs is conformant to the [[spec:DID-CORE]] specification.

## Terminology

[[def:Decentralized Identifier, Decentralized Identifier, DID, DIDs, DID Document, DID Documents]]
~ A [W3C specification](https://www.w3.org/TR/did-core/) describing an identifier that enables verifiable, decentralized 
digital identity. A DID identifier is associated with a DID Document containing keys, services, and other data. 

[[def:Decentralized Agnostic Paytag, Decentralized Agnostic Paytag, DAP, DAPs]]
~ A human-friendly, email-like identifier that can be resolved to a [[ref:DID]] containing one or more payment addresses. 
A DAP has the format `<handle>@<domain>`.

[[def:DAP Registry, DAP Registries, Registry, Registries]]
~ A service that maintains a mapping between DAP handles and [[ref:DIDs]] for a given domain. Registries provide APIs
for resolving and optionally registering DAPs.

[[def:Handle, Handles]]
~ The local name portion of a DAP, unique to the domain at which it is registered. 

[[def:Money Address, Money Addresses]]  
~ A [URN](https://datatracker.ietf.org/doc/html/rfc8141) representing a means to send or receive a specific currency,
with the format `urn:<currency>:<curr_specific_part>`.
   
## Conformance
   
The key words "**MUST**", "**MUST NOT**", "**REQUIRED**", "**SHALL**", "**SHALL NOT**", "**SHOULD**", "**SHOULD NOT**",
"**RECOMMENDED**", "**MAY**", and "**OPTIONAL**" in this document are to be interpreted as described in 
[[spec:RFC2119]].

## DAP

A DAP is composed of a handle and a domain, with the following structure:

```
<handle>@<domain>
```  
   
The domain identifies the [[ref:DAP Registry]] where the handle is registered. The handle itself is unique within that
domain.  

The steps to resolve a DAP are:

1. Split the DAP into handle and domain using `@` delimiter. 
2. Construct a `did:web` [[ref:DID]] using the domain.
3. [Resolve](https://www.w3.org/TR/did-core/#did-resolution) the DID to retrieve the DID Document.
4. Find the `dapregistry` service in the DID Document.
5. Use the `serviceEndpoint` to construct the URL `<serviceEndpoint>/daps/<handle>`.
6. Perform a GET request to the URL.
7. The response contains the DID associated with the DAP.  
8. [Resolve](https://www.w3.org/TR/did-core/#did-resolution) the DID to retrieve its DID Document.
9. Find all `maddr` services in the DID Document which contain payment addresses.

### Handle Format

The precise format of DAP handles is determined by the domain where they are registered. If no preexisting format
exists, the following restrictions are ****RECOMMENDED****:

- **UTF-8 Encoding:** Handle characters ****SHOULD**** be UTF-8 encoded.  
- **Character Exclusions:** Handles ****MUST NOT**** include [control characters]
(https://en.wikipedia.org/wiki/Unicode_control_characters) or [punctuation]
(https://en.wikipedia.org/wiki/Template:Punctuation_marks_in_Unicode).
- **Length:** Handles ****SHOULD**** be between 3 and 30 characters.

### Money Address

A money address is a payment-network-specific identifier embedded within a [[ref:DID Document]] that enables 
sending or receiving funds in a particular currency. It is represented by the following 
[URN](https://datatracker.ietf.org/doc/html/rfc8141) structure:

```
urn:<currency>:<curr_specific_part>  
```

The `<currency>` represents a code for the payment type (e.g. `usdc`, `btc`), while the `<curr_specific_part>`
contains currency-specific details like a blockchain address or payment URL. 

Some example money addresses:

- USDC on Ethereum: `urn:usdc:eth:0x1234567890abcdef1234567890abcdef12345678`  
- BTC LNURL: `urn:btc:lnurl:https://someurl.com`
- BTC Address: `urn:btc:addr:1LMcKyPmwebfygoeZP8E9jAMS2BcgH3Yip`

#### Representing in DID Documents

Money addresses are included in [[ref:DID Documents]] as [service](https://www.w3.org/TR/did-core/#services) 
entries with type `maddr`. Each service contains these properties:

| Property         | Value     | Description                                  |
|------------------|-----------|----------------------------------------------|
| `type`           | `maddr`   | Identifies the service as a money address   |
| `id`             | `string`  | Unique identifier for the address           | 
| `serviceEndpoint`| `string[]`| Array of 1+ payment address URNs            |

For example:

```json
{
  "type": "maddr",
  "id": "#btc-1",  
  "serviceEndpoint": ["urn:btc:addr:1LMcKyPmwebfygoeZP8E9jAMS2BcgH3Yip"]
}
```

## DAP Registry

A DAP Registry maintains the mapping between handles and DIDs for a given domain. It provides APIs for resolving
DAPs and optionally for registering new DAPs.  

To host a DAP Registry, a domain ****MUST****: 

1. Host a resolvable `did:web` DID Document at `https://<domain>/.well-known/did.json`.
2. Advertise the DAP Registry in a `dapregistry` service:

```json
{
  "id": "did:web:example.com", 
  "service": [{
    "type": "dapregistry",
    "serviceEndpoint": ["https://dap.example.com"] 
  }]
}
```

### Registry HTTP API

DAP Registries ****MUST**** provide an HTTP API with the following specifications.

#### CORS 

The Registry ****MUST**** have a [CORS policy](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) allowing 
requests from any origin, to enable access by any app that wishes to resolve or register DAPs.

#### Responses

Responses are JSON objects with the following fields:

| Field   | Type              | Required | Description                              |
|---------|-------------------|----------|------------------------------------------|
| `data`  | `any`             | N        | Endpoint-specific response data on success|
| `error` | [`Error`](#error) | N        | Details if an error occurs              |

`data` and `error` are mutually exclusive.

##### Error 
  
An `Error` object has the following properties:  
  
| Field     | Type     | Required | Description                 |
|-----------|----------|----------|-----------------------------|
| `message` | `string` | Y        | Human-readable error details|

Responses ****MUST**** include the `Content-type: application/json` header.   

#### Resolution Endpoint

**URL:** `GET <serviceEndpoint>/daps/<handle>`

Returns the DID currently associated with the given handle.  

**Response:**

| Field    | Type     | Required | Description                              |  
|----------|----------|----------|------------------------------------------|
| `did`    | `string` | Y        | The resolved DID for the requested handle|
| `proof`  | `string` | N        | Signed payload from the DID registration |

**Errors:**

`404 Not Found` if the handle does not exist or has no associated DID.

#### Registration Endpoint

> [!NOTE]
> This endpoint is **OPTIONAL**. Registries may choose whether to allow individuals to register their
own DIDs or to use an external mechanism.  

**URL:** `POST <serviceEndpoint>/daps`

Registers a DAP, associating the given handle with a DID.

**Request:**

| Field       | Type    | Required | Description                                   |
|-------------|---------|----------|-----------------------------------------------|
| `id`        | `string`| Y        | Unique ID for the registration                |
| `domain`    | `string`| Y        | The domain where the handle is being registered|
| `handle`    | `string`| Y        | Desired handle to register                    |
| `did`       | `string`| Y        | DID to associate with the handle              | 
| `signature` | `string`| Y        | JWS signature over the request (see below)    |

The `signature` is a detached JWS as per [[spec:RFC7515]] over the SHA-256 [[spec:RFC6234]] digest of the
canonicalized [[spec:RFC8785]] JSON request payload, with the JWS `kid` header set to the 
[assertion method](https://www.w3.org/TR/did-core/#assertion) verification method ID from the DID being registered.

**Response:**

`HTTP 202 Accepted`

> [!NOTE]
> Registries may redirect with an `HTTP 303 See Other` to an authentication flow before accepting the registration.

#### Metadata Endpoint

**URL:** `GET <serviceEndpoint>/metadata`

Returns metadata about the Registry's capabilities.

**Response:**

| Field          | Type                    | Required | Description                    |
|----------------|-------------------------|----------|--------------------------------|
| `registration` | `RegistrationMetadata`  | Y        | Registration-related metadata  |

##### RegistrationMetadata

| Field                  | Type       | Required | Description                             |
|------------------------|------------|----------|-----------------------------------------|
| `enabled`              | `boolean`  | Y        | Whether the registry allows registrations|  
| `supportedDidMethods`  | `string[]` | N        | DID methods supported for registration  |

## Privacy Considerations

- DAPs and the associated DIDs/payment addresses are public information visible to anyone who can access the
Registry. Avoid including personally identifiable or other sensitive data.

- To mitigate privacy risks, users may choose to create multiple DAPs, each associated with different payment
addresses and used in different contexts.

## Reference

[[spec:DID-CORE]]
~ [Decentralized Identifiers (DIDs) v1.0](https://www.w3.org/TR/did-core/). 
A. Sambra, M. Sporny, D. Longley, D. Reed, M. Sabadello, D. Burnett.
World Wide Web Consortium (W3C), 18 August 2022. W3C Recommendation.

[[spec:RFC2119]]
~ [Key words for use in RFCs to Indicate Requirement Levels](https://www.rfc-editor.org/rfc/rfc2119).
S. Bradner. IETF, March 1997. Best Current Practice.

[[spec:RFC6234]]
~ [US Secure Hash Algorithms (SHA and SHA-based HMAC and HKDF)](https://www.rfc-editor.org/rfc/rfc6234).
D. Eastlake 3rd, T. Hansen. IETF, May 2011. Informational.

[[spec:RFC7515]]
~ [JSON Web Signature (JWS)](https://www.rfc-editor.org/rfc/rfc7515). M. Jones, J. Bradley, N. Sakimura. 
IETF, May 2015. Proposed Standard.

[[spec:RFC8785]]
~ [JSON Canonicalization Scheme (JCS)](https://www.rfc-editor.org/rfc/rfc8785). A. Rundgren, B. Jordan, S. Erdtman.
IETF, June 2020. Informational.
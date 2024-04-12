# Decentralized Agnostic Paytag (DAP) <!-- omit in toc -->

> [!WARNING]
> 🚧 WIP 👷‍♂️

![](./braindump.png)

# Table of Contents <!-- omit in toc -->
- [Problem Statement](#problem-statement)
- [Requirements](#requirements)
- [Money Address](#money-address)
  - [Payment Address](#payment-address)
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
- [Registry](#registry)



# Problem Statement

> [!WARNING]
> TODO: Fill out

# Requirements

> [!WARNING]
> TODO: Fill out

# Money Address

This specification defines a [DID Service](https://www.w3.org/TR/did-core/#services) configuration that can be used to express a money address for a DID subject.

> [!NOTE]
> [Services](https://www.w3.org/TR/did-core/#services) are used in DID documents to express ways of communicating with the respected DID subject. A service can be any type of service the DID subject wants to advertise, including decentralized identity management services for further discovery, authentication, authorization, or interaction.

In order to express a money address, a service object MUST contain the following properties:
| field             | value      | description                                                                |
| :---------------- | :--------- | :------------------------------------------------------------------------- |
| `type`            | `maddr`    | **MUST** always be `maddr`                                                 |
| `id`              | `[]string` | an abitrary string that can be used to uniquely identify the money address |
| `serviceEndpoint` | `[]string` | 1 or more payment address URNs                                             |


> [!TIP]
> IDs can be used as part of the DAP protocol to specify a specific payment address for a DID subject e.g. `moegrammer#munn@didpay.me`

## Payment Address

A payment address is a Uniform Resource Name (URN) that conforms to the following format:

```
urn:<currency_code>:<code_specifc_part>
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

# Registry

> [!WARNING]
> TODO: Fill out
# cfdns

cfdns is a tool that allows the management of Cloudflare DNS records via the API easily within a CLI. It also has the option to set dynamically the public IPv4 of the machine (or the network itself), through detection, for specific DNS records. A system wide config file allows working with different profiles (API key and email) at the same time. 

The tool does __NOT__ cover anything else of the Cloudflare API.

**STATE:** WIP (!)

## Features

* Create, update and delete DNS records for specific Cloudflare zone
* Update the content of type A DNS records automatically to the public IPv4
* Manage multiple Cloudflare profiles locally with different profiles in a config file

## Usage

### CLI

Full usage documentation for the CLI could be found in a separate [folder](docs/cli/cfdns.md).


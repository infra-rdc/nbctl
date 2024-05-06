<div align="center">
  <a href="https://github.com/infra-rdc/nbctl">
    <img src="img/netbox-logo.png" alt="Logo" width="300" height="90">
  </a>

  <h3 align="center">nbctl</h3>

  <p align="center">
    Client for NetBox written in Go.
    <br />
    <a href="https://github.com/infra-rdc/nbctl"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/infra-rdc/nbctl">View Demo</a>
    ·
    <a href="https://github.com/infra-rdc/nbctl">Report Bug</a>
    ·
    <a href="https://github.com/infra-rdc/nbctl">Request Feature</a>
  </p>
</div>

`nbctl` is a command line for interacting with Netbox. For the moment, it only lists items.

## Getting started

```bash
# For insecure usage (dev only)
export NETBOX_HTTP_SCHEME="http"
export NETBOX_HOST="netbox.example.org"
export NETBOX_TOKEN="....."
nbctl help

----
Uncomplicated CLI interaction with Netbox.

Usage:
  nbctl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dcim        Interact with dcim.
  help        Help about any command
  version     Shows the current nbctl version

Flags:
  -h, --help     help for nbctl
  -t, --toggle   Help message for toggle

Use "nbctl [command] --help" for more information about a command.
```

## Usage

### List devices

List the various items of equipment, filtering by equipment, bay, location, type, site, status, etc.

```bash
nbctl dcim devices list [--json|--raw]
```

Example:

```bash
nbctl dcim devices list -k "rack 1"
+---------------+--------------------------------+----------------+--------+----------------+--------------------+----------------+---------+
|     NAME      |              TYPE              |     TENANT     | SERIAL |    LOCATION    |        SITE        |      RACK      | STATUS  |
+---------------+--------------------------------+----------------+--------+----------------+--------------------+----------------+---------+
| Server1       | PowerEdge R430                 | Restos du Cœur |        | Salle serveurs | Data Center        | rack 1         | active  |
| Server2       | PowerEdge R430                 | Restos du Cœur |        | Salle serveurs | Data Center        | rack 1         | active  |
| Web1          | PowerEdge R630                 | Restos du Cœur |        | Salle serveurs | Data Center        | rack 1         | active  |
| Web2          | PowerEdge R630                 | Restos du Cœur |        | Salle serveurs | Data Center        | rack 1         | active  |
```

## License

Distributed under the GPLv3 License. See LICENSE for more information.

## Contact

Julien Briault - [@ju_hnny5](https://twitter.com/ju_hnny5)

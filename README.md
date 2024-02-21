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

## Getting started

```bash
# For insecure usage (dev only)
export NETBOX_HTTP_SCHEME="http"
export NETBOX_HOST="netbox.example.org"
export NETBOX_TOKEN="....."
nbctl help
```

## Usage

### List devices

```bash
nbctl dcim devices list [--jsor|--raw]
```

## License

Distributed under the GPLv3 License. See LICENSE for more information.

## Contact

Julien Briault - @ju_hnny5

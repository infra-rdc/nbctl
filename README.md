# NetBox CLI

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

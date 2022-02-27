# edgi on Proxmox VE

## Quick Start

1. Build Proxmox VE image using [Packer](https://www.packer.io/): 

```
packer build -var-file=vars.json template.json
```

## Notes

Can define IP and other parameter on config/cloud.yml, according to [Configuration Reference](https://edgi.io/cmd/edgi/blob/master/README.md#configuration-reference)

# Storage Architecture

## Current Storage

The U-Server currently uses a single 2TB hard disk as its active storage device.

Because only one healthy storage drive is currently available, RAID is not configured.

## RAID Research

### RAID 0

- Requires at least 2 drives
- Provides increased performance
- Provides no redundancy
- Failure of one drive results in loss of the array
- Not recommended for server data that must be protected

### RAID 1

- Requires at least 2 drives
- Mirrors data between drives
- Provides redundancy if one drive fails
- Usable capacity is approximately equal to one drive
- Suitable for simple server redundancy

### RAID 5

- Requires at least 3 drives
- Uses distributed parity
- Can survive failure of one drive
- Provides more usable capacity than RAID 1
- Rebuild operations can place significant load on remaining drives

### RAID 6

- Requires at least 4 drives
- Uses dual distributed parity
- Can survive failure of two drives
- Provides stronger redundancy than RAID 5
- Requires additional storage overhead

### RAID 10

- Requires at least 4 drives
- Combines mirroring and striping
- Provides strong performance and redundancy
- Usable capacity is approximately 50 percent of raw capacity
- Suitable for high-performance server workloads

## Planned Architecture

For the current homelab, RAID 1 is the preferred initial redundancy configuration once a second suitable storage drive is installed.

RAID 1 provides:

- Simple redundancy
- Straightforward recovery
- Lower drive-count requirements
- Protection against a single drive failure

A future storage expansion may use RAID 10 if four or more matching drives are installed and higher performance is required.

## Current Limitation

RAID configuration is currently blocked because the server does not have enough healthy storage drives installed to create a redundant array.

No existing production filesystem should be converted or repartitioned for RAID until additional storage hardware is installed and all important data is backed up.

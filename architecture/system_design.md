# System Design

## Overview

Ubuntu SQL Server is built around a PostgreSQL database that stores information collected from U-Server.

The Go Agent collects server information and health data. Other applications can then read the stored information and present it in different ways.

## Current Design

```text
                    U-Server
                       |
                  Go Agent
                       |
                       v
                  PostgreSQL
                       |
          +------------+------------+
          |            |            |
          v            v            v
       Django        Python        Rust
      Dashboard      Client        Client

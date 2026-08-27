# Performance

## Overview

This directory contains performance documentation for the Ubuntu SQL Server project.

The goal is to track server performance, document benchmark results, and record tuning changes as the project grows.

Performance work should be based on measured results instead of making changes without knowing whether they are needed.

## Performance Areas

Performance testing may include:

- CPU usage
- Memory usage
- Disk usage
- Disk performance
- PostgreSQL performance
- Go Agent performance
- Django Web Dashboard performance
- Network performance
- Docker performance
- Application response times

## Current Monitoring

The Go Agent currently records basic system health information including:

- 1-minute load average
- Memory usage
- Disk usage
- System uptime

These measurements help identify general server health but are not a replacement for dedicated performance benchmarks.

## Documentation

- `benchmarks.md` - Records performance tests and benchmark results.
- `tuning.md` - Documents performance tuning changes and recommendations.

## Testing Approach

Performance changes should follow a simple process:

1. Measure current performance.
2. Record the results.
3. Identify a possible bottleneck.
4. Make one change at a time.
5. Run the same test again.
6. Compare the results.
7. Keep or reverse the change based on the results.

This helps prevent unnecessary configuration changes and makes performance improvements easier to verify.

## Status

Performance testing and tuning will continue as additional applications and services are added to U-Server.

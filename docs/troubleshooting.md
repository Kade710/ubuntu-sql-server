# Troubleshooting Documentation

## Overview

This document records issues encountered while building and configuring the U-Server environment.

Purpose:

- Document problems and solutions
- Track configuration changes
- Improve repeatability
- Demonstrate troubleshooting workflow


# Issue 1: SSH Installation Package Verification

## Problem

SSH remote access was required to manage the server without a dedicated keyboard and mouse.

## Investigation

The SSH package was installed and the service was tested.

Command:

```bash
systemctl status ssh

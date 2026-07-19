# RFC-0010 — Component Discovery

## Status

Draft

## Summary

This RFC defines how Verso discovers components inside a project.

Discovery must be deterministic.

A project should always resolve to the same component graph regardless of execution environment.

## Initial Discovery Roots

- skills/
- memory/
- workflows/
- templates/

Future RFCs may define additional component roots.
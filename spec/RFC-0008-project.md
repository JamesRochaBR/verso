# RFC-0008 — Project Model

## Status

Draft

## Summary

A Verso Project is the root unit of organization.

It defines which components belong to the project and how they are discovered.

This RFC defines the logical project model.

It does not define any serialization format.

The serialization (TOML, YAML, JSON...) is intentionally specified in a separate RFC.

## Initial Concepts

A project has:

- metadata
- components
- discovery rules
- compatibility information

Future RFCs may extend this model without breaking compatibility.
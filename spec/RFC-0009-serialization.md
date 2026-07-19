# RFC-0009 — Serialization

## Status

Draft

## Summary

This RFC defines how a Verso Project Model is serialized.

The reference serialization format is TOML.

Alternative serialization formats may exist as long as they represent the same logical Project Model.

## Goals

- Human readable
- Git friendly
- Deterministic
- Extensible
- Backward compatible

## Non Goals

This RFC does not define the Project Model itself.

That responsibility belongs to RFC-0008.
# ADR-0000: Why We Keep ADRs

## Context
This project makes many hard-to-reverse design decisions (on-disk formats, concurrency model, crash-recovery semantics). Without a record, the reasoning behind a decision is lost the moment it's made, and future work has no way to tell whether an odd-looking piece of code is deliberate or accidental.

## Options considered
- No record — rely on commit messages and memory.
- A wiki/external doc — decoupled from the code, tends to rot.
- ADRs committed alongside the code — versioned with the code, reviewed in PRs, easy to link from design docs.

## Decision
Record every major design decision as an ADR under `docs/adr/`, using the template in `docs/adr/TEMPLATE.md`, numbered sequentially.

## Consequences
- Extra writing overhead per decision.
- A durable, searchable record of *why*, not just *what* — useful for onboarding, debugging, and interviews.

## Revisit when
N/A — this is a process decision, not a technical one.

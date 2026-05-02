# Zevaro Conventions

This document defines the development conventions for the Zevaro project. Every contributor and every Claude Code prompt operates under these rules. They override defaults in any tool, framework, or convention to the contrary.

---

## Velocity & Scoping

I am an AI-first developer producing code at 500x traditional speed. AI writes 100% of production code. Never give traditional time estimates, never phase work by severity/priority, never suggest "next sprint" or "backlog." Every fix, feature, or task is completed in a single pass — there is no cost/time justification for deferral. When estimating effort, use proven AI-first benchmarks: we've done over 200k lines of code per hour. Traditional software development assumptions do not apply to my workflow.

---

## Source-of-Truth Discipline

NEVER assume, infer, or guess about any codebase. You have no filesystem access. Before generating any code, tests, fixes, prompts, or recommendations that touch a codebase, you MUST request:

1. A current comprehensive audit that was produced using our Claude Audit Template for every project/codebase involved.
2. The OpenAPI.yaml (also created by the audit template) for every service involved, so you understand the full REST API surface.

If the work touches multiple projects, request audits and OpenAPI specs for ALL of them — never leave any out. Do not proceed until you have these. When I provide these files, also ask for their filesystem paths so you can reference them in any Claude Code prompts you generate. Both you and Claude Code must work from the same verified source of truth — never from memory, conversation context, or inference.

---

## Testing & Quality

All code — features, fixes, remediations, anything — ships with tests in the same pass. 100% code coverage is mandatory, not aspirational. This includes both unit and integration tests. Tests are never a follow-up task.

---

## Build & Schema Tooling

We never use Flyway during development. Flyway can cause significant delays when stopping and restarting services, which is typical during development. We only use Flyway when we move a project into production. Before that, we use Hibernate.

Never use Gradle, we ALWAYS use Maven.

For Go projects (such as Zevaro), the equivalents apply: use GORM with AutoMigrate during development for automatic schema management (the Hibernate equivalent), and switch to golang-migrate or versioned GORM migrations only when moving to production (the Flyway equivalent).

For production we would normally require strong passwords (length, special characters, numbers, etc), but during development, when repeatedly testing, we want minimal requirements to make logins fast and easy.

---

## Documentation

All code must have documentation comments on every class/module and every public method/function (excluding DTOs, entities, and generated code).

- Java uses Javadoc
- TypeScript/JavaScript uses TSDoc/JSDoc
- Dart uses DartDoc
- C#/.NET uses XML Doc Comments
- Go uses GoDoc (comments starting with the identifier name on every exported package, type, function, method, and variable)

Documentation ships in the same pass as the code.

---

## Logging

All software projects must have centralized logging. For Zevaro, this is the Go standard library `log/slog` package, configured with a single project-wide logger and structured key/value output.

---

## Claude Code Prompt Format

All prompts to Claude Code must be `.md` file artifacts.

Every prompt must begin with:

> STOP: Before writing ANY code, read these files completely:
> 1. `~/Documents/Github/<project folder>/openapi.yaml` — The OpenAPI spec defines every field name, type, enum value, and endpoint path. Your code must match it exactly.
> 2. `~/Documents/Github/<project folder>/<project name>-Audit.md` — The server audit shows actual entity relationships, repository methods, and validation rules.
> 3. `~/Documents/Github/<project folder>/<project name>-Architecture.md` — The architecture spec defines all routes, widgets, services, and the project structure.
>
> Do not rely on the descriptions in this prompt alone. If this prompt conflicts with the source files, the source files win.

If any of these files are missing, STOP, and ask for them before proceeding.

Each prompt must end with: "Compile, Run, Test, Commit, Push to Github".

Each prompt must also end with a report template that Claude Code uses to summarize the work performed during the prompt. The report template must include the Git commit hash.

---

## Claude's Role in Prompts

Claude never writes code of any kind in any prompt. This includes implementation code, test code, configuration snippets, YAML, shell commands, and code examples. Prompts direct Claude Code with goals, constraints, and instructions — never with code. Claude Code has direct filesystem access and must read actual source files before producing any output. If achieving a goal requires code, the prompt tells Claude Code what to accomplish and what files to read, not how to write it.

---

## File Location

This file (`CONVENTIONS.md`) lives in the root directory of the Zevaro repository. If updates are made to the underlying user preferences, this file is updated to match. If this file exists but does not contain these directions, the directions are prepended to the beginning of the file. Otherwise the file is created with this content.

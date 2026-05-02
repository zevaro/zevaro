## Summary

<!-- One paragraph describing what this PR does and why. Link to the related issue if one exists. -->

Closes #<!-- issue number, or "N/A" -->

---

## Scope

<!-- List every file touched by this PR and the reason. Flag anything touched outside the declared task scope. -->

| File | Reason |
|------|--------|
|      |        |

---

## Checklist

### Code
- [ ] I only modified files within the declared task scope
- [ ] No `// TODO`, `// FIXME`, `panic("not implemented")`, or stubs remain
- [ ] No magic literals — constants live in `internal/constants/` (Go) or `frontend/src/lib/constants.ts`
- [ ] No `fmt.Print*` or `log.Print*` — all output uses `slog`
- [ ] Error messages contain no user content

### Tests
- [ ] All new code paths have unit tests
- [ ] Integration tests added if the change crosses package boundaries
- [ ] `make test` passes locally with no failures
- [ ] Line coverage on authored code is 100%

### Documentation
- [ ] GoDoc comments on every new exported identifier (packages, types, functions, methods, constants)
- [ ] TSDoc comments on every new exported TypeScript class and public method
- [ ] `README.md` updated if the change affects setup, build, or usage

### Build verification
- [ ] `make build` exits 0 with no warnings
- [ ] `make lint` exits 0 (`golangci-lint run` + `eslint + tsc --noEmit`)
- [ ] `make test` exits 0 with all tests passing

### Breaking changes
- [ ] This PR introduces no breaking API changes (or breaking changes are documented below)

---

## Breaking changes (if any)

<!-- Describe any changes to the OpenAI-compatible, Anthropic-compatible, or management API surfaces.
     Breaking changes require updating openapi.yaml and the client integration guides. -->

N/A

---

## Related issues / PRs

<!-- List any related issues or PRs that provide context. -->

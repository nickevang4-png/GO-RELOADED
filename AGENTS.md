# AGENTS.md

## 🧑‍💻 Project: GO-RELOADED

This project is part of a peer-audited program.  
Each student acts as a **developer** and an **auditor**.

---

## 🔍 Audit Rules

- Code must follow Go best practices.
- Tests must pass with \`go test ./...\`.
- Each function should be modular and reusable.
- Check that the program output matches the exercise examples.

---

## ✅ Self-Test Checklist

- [ ] Handles punctuation correctly
- [ ] Converts (hex) and (bin)
- [ ] Applies (up), (low), (cap)
- [ ] Corrects "a" → "an"


## Auditing Steps (quick)

1. Checkout code and run formatting/lint:
   - go fmt ./...
   - go vet ./...
   - (optional) golangci-lint run

2. Run unit tests:
   - go test ./...

3. Run integration/examples:
   - go run main.go sample.txt result.txt
   - cat result.txt
   - Compare with `examples/output.txt` or the examples in README.

4. Verify edge cases:
   - Markers with punctuation: `(cap, 6) ,` etc.
   - Multi-word quotes: `' I am the world '` → `'I am the world'`
   - Punctuation groups: `...` and `!?`

5. Code review checklist:
   - Small, focused functions per file (processor/ or internal/).
   - No duplicated pipeline implementations (pick one canonical Process).
   - Unit tests are table-driven and cover expected behaviors.
   - README explains usage and contains sample inputs → outputs.

---

## Scoring (suggested, for peers)

- Functionality (50%): passes all example transformations and tests.
- Tests & Quality (25%): unit tests cover edge cases; code is formatted and vetted.
- Code Organization (15%): modular, clear packages, single pipeline entrypoint.
- Documentation (10%): README + examples + AGENTS.md present.

---

## Notes

- If you want automated checks, add a GitHub Actions workflow that runs build, test and lint on push/PR.
- Attach failing test cases or diffs when requesting fixes.

Good auditing.
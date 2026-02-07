# MonkeyLines

## What This Is

A web-based interactive experience featuring a three-headed monkey from The Secret of Monkey Island delivering insult sword-fighting exchanges. One side monkey delivers an insult, the center monkey responds with a comeback, complete with lip-sync animation and distinct hooting voices for each head. This is a portfolio piece demonstrating AI-assisted development and design sensibility.

## Core Value

Visitors immediately understand what this is, recognize the Monkey Island nostalgia, and find it delightful to interact with—no confusion, just charm.

## Requirements

### Validated

- ✓ Three-headed monkey visualization with animated overlays — existing
- ✓ Insult/comeback exchange generation from themed content — existing
- ✓ Lip-sync animation synchronized with text display — existing
- ✓ Web Audio API hooting sounds with distinct pitch per monkey — existing
- ✓ Auto-play mode for continuous exchanges — existing
- ✓ Mute/unmute audio control — existing
- ✓ Mobile browser compatibility with audio unlock — existing
- ✓ Stateless Go server with embedded assets — existing

### Active

- [ ] Pixel art button styling (Monkey Island game aesthetic)
- [ ] Reorganized control layout (split up, not single row)
- [ ] Visual hints/labels for each control's purpose
- [ ] Full page Monkey Island theming (background, borders, frames)
- [ ] Styled speech bubble matching game aesthetic
- [ ] Immediate visual communication of purpose
- [ ] Clear cultural references for Monkey Island fans
- [ ] Professional visual polish suitable for portfolio

### Out of Scope

- Additional features beyond UI revamp — focus is visual polish, not new functionality
- Backend changes — Go server and exchange logic stay as-is
- New content themes — existing 8 themes are sufficient
- User accounts or persistence — stateless design is intentional
- Multiple monkey characters — three-headed monkey is the concept

## Context

**Existing Implementation:**
- Single-page Go application with zero external dependencies
- All assets (HTML, CSS, JS, images) embedded at compile time
- Vanilla JavaScript for client-side interactions
- Web Audio API for procedural monkey voices
- Works on modern browsers including mobile

**Visual Reference:**
The Secret of Monkey Island (1990) LucasArts adventure game, specifically:
- Pixel art aesthetic from the original VGA version
- Pirate/Caribbean themed visual elements
- Classic LucasArts UI design language
- Verb coin or SCUMM interface styling

**Portfolio Context:**
This project demonstrates collaboration with AI for UI/UX design and implementation. The goal is showing ability to ship polished, thoughtful work with strong aesthetic cohesion.

## Constraints

- **Tech stack**: Vanilla HTML/CSS/JavaScript only — no frameworks
- **Compatibility**: Must work on modern browsers (desktop + mobile)
- **Embedded assets**: All resources must be embeddable in Go binary
- **Performance**: Keep page load fast, minimal asset size
- **Accessibility**: Controls must be intuitive without instructions

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| Full UI revamp over incremental changes | Portfolio piece needs to demonstrate design vision, not just button polish | — Pending |
| Maintain existing functionality exactly | Core mechanics work well; focus on presentation layer only | — Pending |
| Pixel art aesthetic from original MI | Authenticity matters for nostalgic recognition | — Pending |

---
*Last updated: 2026-02-07 after initialization*

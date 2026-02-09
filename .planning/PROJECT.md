# MonkeyLines

## What This Is

A web-based interactive experience featuring a three-headed monkey from The Secret of Monkey Island delivering insult sword-fighting exchanges with authentic SCUMM-era visual aesthetic. One side monkey delivers an insult, the center monkey responds with a comeback, complete with lip-sync animation, distinct hooting voices for each head, pixel font, MI color palette, and CRT scanline overlay. This is a portfolio piece demonstrating AI-assisted development and design sensibility.

## Core Value

Visitors immediately understand what this is, recognize the Monkey Island nostalgia, and find it delightful to interact with—no confusion, just charm.

## Requirements

### Validated

- ✓ Three-headed monkey visualization with animated overlays — pre-v1.0
- ✓ Insult/comeback exchange generation from themed content — pre-v1.0
- ✓ Lip-sync animation synchronized with text display — pre-v1.0
- ✓ Web Audio API hooting sounds with distinct pitch per monkey — pre-v1.0
- ✓ Auto-play mode for continuous exchanges — pre-v1.0
- ✓ Mute/unmute audio control — pre-v1.0
- ✓ Mobile browser compatibility with audio unlock — pre-v1.0
- ✓ Stateless Go server with embedded assets — pre-v1.0
- ✓ Monkey Island color palette (deep purples, teals, ambers) — v1.0
- ✓ Press Start 2P pixel font with crisp rendering — v1.0
- ✓ SCUMM-style verb bar with text labels — v1.0
- ✓ Scene frame with 9-slice pixel-art borders — v1.0
- ✓ Per-speaker colored dialog text (amber/teal/orange) — v1.0
- ✓ CRT scanline overlay with accessibility support — v1.0
- ✓ Full page Monkey Island theming — v1.0
- ✓ Professional visual polish suitable for portfolio — v1.0

### Active

(No active requirements — next milestone will define new goals)

### Out of Scope

- Additional features beyond UI revamp — focus is visual polish, not new functionality
- Backend changes — Go server and exchange logic stay as-is
- New content themes — existing 8 themes are sufficient
- User accounts or persistence — stateless design is intentional
- Multiple monkey characters — three-headed monkey is the concept

## Context

**Current Codebase (v1.0):**
- Single-page Go application with zero external dependencies
- 526 lines of HTML/CSS/JS (20 files modified in v1.0)
- All assets (HTML, CSS, JS, images, fonts) embedded at compile time
- Vanilla JavaScript for client-side interactions
- Web Audio API for procedural monkey voices
- Press Start 2P pixel font (self-hosted WOFF2)
- MI color palette via CSS custom properties (--mi-*)
- Works on modern browsers including mobile

**Tech Stack:**
- Frontend: Vanilla HTML/CSS/JavaScript
- Backend: Go with embedded assets
- Styling: CSS custom properties, 9-slice borders, pseudo-elements
- Accessibility: prefers-reduced-motion support

**Visual Reference:**
The Secret of Monkey Island (1990) LucasArts adventure game, specifically:
- Pixel art aesthetic from the original VGA version
- Pirate/Caribbean themed visual elements
- Classic LucasArts UI design language
- SCUMM verb interface styling with text labels

**Portfolio Context:**
This project demonstrates collaboration with AI for UI/UX design and implementation. v1.0 achieved portfolio-ready visual polish with cohesive SCUMM-era aesthetic.

**Known Issues:**
None — all v1.0 features working as intended.

## Constraints

- **Tech stack**: Vanilla HTML/CSS/JavaScript only — no frameworks
- **Compatibility**: Must work on modern browsers (desktop + mobile)
- **Embedded assets**: All resources must be embeddable in Go binary
- **Performance**: Keep page load fast, minimal asset size
- **Accessibility**: Controls must be intuitive without instructions

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| Full UI revamp over incremental changes | Portfolio piece needs to demonstrate design vision, not just button polish | ✓ Good - v1.0 achieved cohesive aesthetic |
| Maintain existing functionality exactly | Core mechanics work well; focus on presentation layer only | ✓ Good - zero functionality changes |
| Pixel art aesthetic from original MI | Authenticity matters for nostalgic recognition | ✓ Good - SCUMM-era feel achieved |
| Self-hosted Press Start 2P over Google Fonts CDN | Zero external dependencies for reliability | ✓ Good - consistent with Go embedded assets |
| CSS custom properties on :root for MI palette | Reusable theming pattern across all phases | ✓ Good - enabled consistent color usage |
| Anti-aliasing disabled globally | Authentic pixel art aesthetic | ✓ Good - crisp retro rendering |
| Text labels over pixel art icons for buttons | SCUMM authenticity after icon rendering issues | ✓ Good - more authentic than icons |
| Hybrid approach: speaker icon for mute, text for Fight/Auto | Balance between visual clarity and SCUMM aesthetic | ✓ Good - best of both approaches |
| Inline base64 data URIs for border PNGs | No extra asset files to manage | ✓ Good - clean implementation |
| data-speaker attribute pattern | Per-monkey CSS styling without JS complexity | ✓ Good - scalable pattern |
| Background image speech bubble over CSS border-image | Preserve existing fitText() functionality | ✓ Good - avoided breaking changes |
| 6px scanline pattern (adjusted from 4px) | Better visibility at default zoom based on user feedback | ✓ Good - balanced CRT effect |
| Opacity reduction (not removal) for reduced-motion | Maintains visual consistency while respecting accessibility | ✓ Good - inclusive design |

---
*Last updated: 2026-02-09 after v1.0 milestone*

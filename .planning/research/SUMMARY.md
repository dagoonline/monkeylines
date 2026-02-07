# Project Research Summary

**Project:** MonkeyLines - Monkey Island Aesthetic Revamp
**Domain:** Pixel art retro game UI for web applications
**Researched:** 2026-02-07
**Confidence:** MEDIUM-HIGH

## Executive Summary

MonkeyLines is a single-page web application displaying Monkey Island-style insult sword fighting exchanges. The goal is to transform the current functional-but-generic UI into an authentic LucasArts SCUMM engine aesthetic while maintaining the vanilla HTML/CSS/JS constraint. Research shows this is highly achievable with CSS-based techniques — no build system or frameworks needed.

The recommended approach leverages `image-rendering: pixelated` for crisp pixel art scaling, CSS custom properties for a Monkey Island color palette (deep purples, Caribbean teals, warm ambers), and `border-image` for authentic SCUMM-style UI panels. The critical path is: establish the color palette and pixel font first (foundation), then restyle the verb bar-style button panel, then add decorative framing. This sequence minimizes risk while maximizing visual impact.

The primary risk is **mixed pixel densities** — if new UI elements don't match the existing art's pixel grid, the illusion collapses immediately. Prevention requires establishing a single base pixel grid before creating any assets and auditing all visuals at 400% zoom. Secondary risks include blurry scaling from missing `image-rendering` properties and unreadable text on mobile if pixel fonts are sized incorrectly. Both are preventable with disciplined CSS patterns and responsive testing at 320px minimum width.

## Key Findings

### Recommended Stack

The existing vanilla HTML/CSS/JS approach is exactly correct for this project. CSS alone handles all pixel art rendering needs — no Canvas, WebGL, or frameworks required.

**Core technologies:**
- **`image-rendering: pixelated`**: Critical CSS property that prevents bilinear filtering, keeping pixel art crisp. Already used in project; must extend to all new UI elements.
- **PNG-8 (indexed color)**: Lossless format preserving exact pixel colors. Never use JPEG/WebP for pixel art — lossy compression destroys sharp edges.
- **CSS custom properties**: Monkey Island color palette (16-32 colors) defined as `:root` variables. Zero runtime cost, no build step needed.
- **`border-image` with 9-slice**: Pixel-perfect scalable panels from single sprite. Key technique for SCUMM-style dialog boxes and verb bar panels.
- **Press Start 2P font (self-hosted WOFF2)**: Standard retro pixel font. Deliver via self-hosted WOFF2 (not Google Fonts CDN) to avoid external dependencies.
- **CSS `steps()` animation**: Frame-by-frame sprite animation with no interpolation. Authentic retro feel, native CSS, no JS needed.

**Critical version requirements:**
- `image-rendering: pixelated` — baseline support since January 2020 (all modern browsers)
- WOFF2 — baseline support across all modern browsers
- CSS custom properties — supported everywhere since 2017

### Expected Features

**Must have (table stakes):**
- **Pixel font with antialiasing disabled** — Monospace alone isn't enough; bitmap typography defines the MI aesthetic. Use `-webkit-font-smoothing: none`.
- **Labeled buttons replacing emoji** — Current emoji-only buttons require guessing. Icon+text labels ("Fight!", "Auto", "Mute") match SCUMM verb pattern and work on mobile where tooltips don't exist.
- **Scene frame/border** — Monkey Island always frames the game viewport. A pixel art border or CSS border in MI palette creates "screen within screen" separation.
- **Speaker-colored speech text** — Each character needs distinct text color (e.g., amber for insult, teal for comeback). Currently text lacks color distinction.
- **Page title/tagline** — First-time visitors need to understand "this is a Monkey Island insult exchange" in under 3 seconds. Currently no visible explanation.
- **MI color palette** — Specific palette (deep purples #1a0a2e, teals #2d8b6e, ambers #c8a43c, cream #e8d4a0) applied consistently ties everything together.

**Should have (competitive):**
- **Verb bar panel styling** — Dark background strip below scene with text labels in columns, mimicking SCUMM two-zone layout (scene above, controls below).
- **Pixel art button icons** — Replace emoji with 16x16 or 32x32 pixel art sprites in MI palette. Huge authenticity boost.
- **`border-image` UI panels** — 9-slice panel borders instead of plain CSS borders. Single-asset technique that scales perfectly.
- **Typewriter sound effect** — Subtle key-click per character synced to text reveal, layered with existing hoot system.
- **CRT scanline overlay** — Very subtle repeating-linear-gradient or tiny PNG tile. Must be restrained to avoid hurting readability.

**Defer (v2+):**
- **Idle animations** — Subtle monkey head bob or blinking when not speaking. Requires new sprite frames.
- **Startup text crawl** — "Deep in the Caribbean..." style intro on first load.
- **Complex sprite animations** — Multi-frame walking/gesturing. Current static/toggle approach works fine.

### Architecture Approach

Keep the single-file inline architecture. The Go embed + vanilla constraint is a feature, not a limitation. No build system, no components, no frameworks.

**Major components:**
1. **Page Background** — Themed background color/pattern via CSS. No dependencies, pure visual wrapper.
2. **Scene Frame** — Decorative border via CSS `border-image` or pure CSS. Wraps existing monkey scene without touching JS.
3. **Monkey Scene** — Existing overlay system unchanged (base.png + l/c/r toggles). Just extend `image-rendering: pixelated` to any new elements.
4. **Speech Bubble** — Restyle existing `.speech-text` with pixel-style border and background. Must preserve `fitText()` functionality.
5. **Control Panel** — Restyle existing `.btn-row` as verb bar panel. Replace emoji with labeled buttons or icon+text combos.

**Key patterns:**
- **Individual PNGs, not sprite sheets** — Only 4 images + handful of new UI assets. Individual files simpler to maintain without build tools.
- **Pure CSS borders over images** — Use `box-shadow` layering or `border-image` for pixel borders. Minimal assets, maximum flexibility.
- **Responsive via `clamp()` and aspect-ratio** — Current approach is solid. Don't add breakpoint explosion; single fluid layout works for retro aesthetic.
- **Z-index discipline** — Only needed within `.monkey-scene` (already working). No global z-index war.

### Critical Pitfalls

1. **Mixed pixel densities** — UI elements with visibly different pixel sizes (scene at 4px, border at 2px, text native resolution) destroy retro illusion immediately. Prevention: Establish single base pixel grid before creating any art; audit at 400% zoom to verify all "pixels" are same size.

2. **Blurry pixel art from incorrect scaling** — Missing `image-rendering: pixelated` or CSS transforms causing sub-pixel rendering. Prevention: Apply `image-rendering: pixelated` to all pixel art elements; test on Safari specifically (may need `crisp-edges` fallback); avoid `transform: scale()`, use width/height instead.

3. **Typography breaking illusion** — Pixel fonts at wrong sizes (fractional pixels blur letterforms) or mixing pixel font with antialiased text. Prevention: If using pixel font, serve at exact multiples (8px, 16px, 24px) and apply `-webkit-font-smoothing: none`; never scale with `em`/`vw` units; commit fully to monospace OR pixel font, not both.

4. **Retro aesthetic over usability** — Themed buttons users can't identify as clickable or understand. Prevention: Every interactive element needs visible hover state; ensure touch targets >= 44x44px; five-second test with fresh eyes ("what would you click?").

5. **Non-integer scaling artifacts** — Pixel art scaled to non-multiple sizes yields uneven grid (some pixels 3 CSS px, neighbors 2 CSS px). Prevention: On desktop, constrain containers to integer multiples; on mobile, accept the tradeoff (still better than bilinear blur).

## Implications for Roadmap

Based on research, suggested phase structure follows foundation-first ordering: establish visual language before building components, avoid touching working JS until visual reskin is complete.

### Phase 1: Visual Foundation (Color + Typography)
**Rationale:** Everything else depends on the color palette and font choice. This is pure CSS with zero risk to existing functionality.
**Delivers:** MI color palette as CSS custom properties, pixel font loaded (self-hosted WOFF2), base theme colors applied to body/container.
**Addresses:** Color palette (differentiator), pixel font requirement (table stakes).
**Avoids:** Mixed pixel densities (establish grid), typography breaking illusion (font decision made early).
**Research flag:** Standard patterns, no phase-specific research needed.

### Phase 2: Button Redesign + Control Panel
**Rationale:** Buttons are the primary interaction surface. Needs foundation from Phase 1 (palette + font). Delivers immediate visible impact.
**Delivers:** Replace emoji with labeled buttons using pixel font, style `.btn-row` as SCUMM verb bar panel with dark background.
**Uses:** CSS custom properties (palette), pixel font, `border-image` or pure CSS for panel background.
**Addresses:** Labeled buttons (table stakes), verb bar styling (differentiator).
**Avoids:** Retro aesthetic over usability (labels ensure clarity), touch targets too small (44px minimum enforced).
**Research flag:** Standard patterns, no research needed.

### Phase 3: Scene Frame + Speech Styling
**Rationale:** Depends on Phase 1 palette. Must not break existing `fitText()` logic or speech animation.
**Delivers:** Border/frame around `.monkey-scene` using MI palette, restyle `.speech-text` with speaker colors and pixel-style bubble.
**Addresses:** Scene frame (table stakes), speaker-colored speech (table stakes).
**Avoids:** Dialogue text overflow (test with longest strings), typography issues (font already established in Phase 1).
**Research flag:** Low risk but test responsive behavior — verify `fitText()` still works after CSS changes.

### Phase 4: Polish + Assets
**Rationale:** Everything else is in place. This phase adds pixel art assets and visual effects that enhance but aren't critical.
**Delivers:** Pixel art button icons (replacing text-only buttons from Phase 2), optional CRT scanline overlay, optional typewriter SFX.
**Addresses:** Pixel art icons (differentiator), scanline overlay (differentiator), typewriter sound (differentiator).
**Avoids:** Oversized assets (PNG-8, native resolution exports), CRT effects tanking performance (subtle overlay only, test on throttled CPU).
**Research flag:** Standard asset creation patterns, but verify file sizes and test performance on mobile.

### Phase Ordering Rationale

- **Foundation first (Phase 1):** Color and typography are dependencies for all visual changes. Getting this right early prevents rework.
- **High-impact, low-risk next (Phase 2):** Buttons are most visible interaction. Restyling them doesn't touch animation logic (low risk) but delivers major aesthetic change (high impact).
- **Frame existing features (Phase 3):** Speech and scene are already working. Just wrap/restyle without changing behavior.
- **Polish last (Phase 4):** Pixel art icons and effects are nice-to-have enhancements. If time runs short, Phases 1-3 deliver a complete themed experience.

**Dependency chain:**
```
Phase 1 (foundation) ──> Phase 2 (buttons need palette + font)
                     └──> Phase 3 (frame needs palette)
                     └──> Phase 4 (icons need palette)
```

Phases 2, 3, 4 could theoretically parallelize after Phase 1, but sequential is safer for single developer. Each phase validates the previous (e.g., Phase 3 confirms palette choices from Phase 1 work for borders).

### Research Flags

**Phases likely needing deeper research during planning:**
- None. All phases use well-documented CSS techniques with baseline browser support. The research has covered the domain thoroughly.

**Phases with standard patterns (skip research-phase):**
- **Phase 1:** CSS custom properties and WOFF2 font loading are established patterns. No research needed.
- **Phase 2:** Button restyling and CSS borders are standard. Verify touch target guidelines (44px) but no deep research.
- **Phase 3:** CSS positioning and borders. Already validated in current codebase.
- **Phase 4:** Asset creation (PNG export) and optional audio (existing AudioContext works). Standard practices.

**Validation points (not research, just testing):**
- **Phase 3:** Test `fitText()` still works after speech bubble CSS changes (regression testing, not research).
- **Phase 4:** Performance test CRT overlay on mobile (verify assumption, not research new techniques).

## Confidence Assessment

| Area | Confidence | Notes |
|------|------------|-------|
| Stack | HIGH | All techniques verified via MDN or existing codebase. `image-rendering: pixelated` baseline since 2020, CSS custom properties since 2017. No cutting-edge features. |
| Features | MEDIUM-HIGH | Based on domain knowledge of SCUMM UI patterns and analysis of current codebase. Web search unavailable to verify latest community examples, but patterns are well-established. |
| Architecture | HIGH | Current codebase analysis provides complete picture. Single-file inline approach is proven working. Recommendations don't introduce new paradigms. |
| Pitfalls | MEDIUM-HIGH | Pitfalls based on established pixel art rendering issues (well-documented) and direct codebase analysis. Confidence is medium due to inability to verify latest browser quirks via web search. |

**Overall confidence:** MEDIUM-HIGH

Research is built on solid foundations (baseline CSS features, existing working code, established domain patterns) but lacks validation against latest community discussions or blog posts due to web search unavailability.

### Gaps to Address

- **Font readability on mobile:** Pixel fonts can be hard to read below 14-16px. Research recommends minimum 16px on small screens but exact threshold should be validated during Phase 1 implementation with real device testing.

- **Safari-specific `image-rendering` behavior:** MDN notes Safari historically preferred `crisp-edges` over `pixelated`. Current code only uses `pixelated`. Phase 1 should add `crisp-edges` fallback and test on iOS Safari specifically.

- **Touch target sizing for pixel art buttons:** Research cites 44x44px minimum (Apple HIG) but actual button sizing in current code uses `clamp()`. Phase 2 should measure actual rendered sizes at 320px viewport to verify compliance.

- **`fitText()` interaction with pixel font:** Current code uses monospace + fluid sizing with `fitText()`. If switching to pixel font in Phase 1, verify `fitText()` still produces readable results (pixel fonts at fractional sizes may blur). May need to adjust algorithm to snap to integer multiples.

**How to handle during implementation:**
- Each gap has a specific validation test (device testing, browser testing, size measurement).
- None block starting work; all are "validate assumption during phase" rather than "research alternative approach."
- If a gap proves problematic (e.g., pixel font breaks `fitText()`), fallback is known (keep monospace, defer pixel font to Phase 4 as optional enhancement).

## Sources

### Primary (HIGH confidence)
- **MonkeyLines codebase** (`index.html`, existing images, audio implementation) — Direct analysis, current working implementation.
- **MDN `image-rendering`** — https://developer.mozilla.org/en-US/docs/Web/CSS/image-rendering (verified 2026-02-07, baseline since Jan 2020).
- **MDN `@font-face`** — https://developer.mozilla.org/en-US/docs/Web/CSS/@font-face (verified 2026-02-07, WOFF2 recommended).
- **MDN `border-image`** — https://developer.mozilla.org/en-US/docs/Web/CSS/border-image (baseline, 9-slice pattern).
- **MDN `font-smooth`** — https://developer.mozilla.org/en-US/docs/Web/CSS/font-smooth (verified 2026-02-07, non-standard but widely supported).
- **MDN CSS `steps()`** — https://developer.mozilla.org/en-US/docs/Web/CSS/animation-timing-function (baseline).

### Secondary (MEDIUM confidence)
- **Press Start 2P font** — https://fonts.google.com/specimen/Press+Start+2P (OFL license, standard retro font).
- **Domain knowledge** — LucasArts SCUMM engine UI patterns (Monkey Island 1 & 2, Day of the Tentacle). Well-established game UI patterns, but no specific web examples verified.
- **Pixel art rendering best practices** — Community consensus on `image-rendering`, integer scaling, PNG formats. Established practices but not verified against latest discussions.

### Tertiary (LOW confidence)
- **CRT overlay performance** — Research assumes `repeating-linear-gradient` or tiny PNG tile is performant. Should validate on actual low-end mobile devices during Phase 4.
- **Exact MI color palette values** — Research proposes palette based on visual analysis of existing art. Exact hex values inferred, not extracted from official game assets. Phase 1 should color-pick from base.png to verify accuracy.

---
*Research completed: 2026-02-07*
*Ready for roadmap: yes*

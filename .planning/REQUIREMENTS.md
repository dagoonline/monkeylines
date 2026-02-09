# Requirements: MonkeyLines

**Defined:** 2026-02-07
**Core Value:** Immediate clarity, nostalgic recognition, intuitive interaction

## v1 Requirements

Requirements for Monkey Island-styled UI revamp. Each maps to roadmap phases.

### Visual Foundation

- [x] **VFND-01**: Monkey Island color palette defined as CSS custom properties (deep purples, teals, ambers, cream)
- [x] **VFND-02**: Press Start 2P pixel font loaded (self-hosted WOFF2, antialiasing disabled)
- [x] **VFND-03**: Consistent `image-rendering: pixelated` applied to all pixel art elements
- [x] **VFND-04**: Base theme colors applied to body and container elements

### Button Controls

- [x] **CTRL-01**: Button touch targets meet 44x44px minimum for mobile compatibility
- [x] **CTRL-02**: Verb bar panel styling (dark background strip, SCUMM two-zone layout below scene)
- [x] **CTRL-03**: Pixel art button icons created (16x16 or 32x32 sprites in MI palette) - Hybrid: text labels for Fight/Auto, speaker icon for mute
- [x] **CTRL-04**: Pixel art icons replace emoji buttons for all three controls - Hybrid: speaker icon for mute, text labels for Fight/Auto

### Scene Presentation

- [x] **SCNE-01**: Scene frame/border added around monkey viewport (pixel art or CSS in MI palette)
- [x] **SCNE-02**: Speaker-colored speech text (different color per monkey head: left, center, right)
- [x] **SCNE-03**: 9-slice border-image panels for authentic SCUMM-style scalable borders - Applied to scene frame
- [x] **SCNE-04**: Speech bubble restyled with pixel aesthetic (preserves fitText() functionality) - Used existing background bubble with per-speaker colors

### Polish

- [ ] **PLSH-01**: CRT scanline overlay added (subtle effect, doesn't hurt readability)

## v2 Requirements

Deferred to future release. Tracked but not in current roadmap.

### Onboarding

- **ONBD-01**: Page title/tagline explaining concept to first-time visitors

### Button Enhancements

- **CTRL-05**: Text labels with hover states as fallback/alternative to icons

### Audio Effects

- **AUDI-01**: Typewriter sound effect (key-click per character synced to text reveal)

### Animations

- **ANIM-01**: Idle animations (monkey head bob or blinking when not speaking)
- **ANIM-02**: Startup text crawl ("Deep in the Caribbean..." style intro)
- **ANIM-03**: Complex sprite animations (multi-frame gesturing)

## Out of Scope

Explicitly excluded. Documented to prevent scope creep.

| Feature | Reason |
|---------|--------|
| Additional gameplay features | Focus is UI polish, not new functionality |
| Backend/server changes | Go server and exchange logic work perfectly as-is |
| New content themes | Existing 8 themes provide sufficient variety |
| User accounts or persistence | Stateless design is intentional and appropriate |
| Multiple monkey characters | Three-headed monkey is the core concept |
| Framework migration | Vanilla HTML/CSS/JS constraint is a feature, not limitation |

## Traceability

Which phases cover which requirements. Updated during roadmap creation.

| Requirement | Phase | Status |
|-------------|-------|--------|
| VFND-01 | Phase 1 | Complete |
| VFND-02 | Phase 1 | Complete |
| VFND-03 | Phase 1 | Complete |
| VFND-04 | Phase 1 | Complete |
| CTRL-01 | Phase 2 | Complete |
| CTRL-02 | Phase 2 | Complete |
| CTRL-03 | Phase 2 | Complete |
| CTRL-04 | Phase 2 | Complete |
| SCNE-01 | Phase 3 | Complete |
| SCNE-02 | Phase 3 | Complete |
| SCNE-03 | Phase 3 | Complete |
| SCNE-04 | Phase 3 | Complete |
| PLSH-01 | Phase 4 | Pending |

**Coverage:**
- v1 requirements: 13 total
- Mapped to phases: 13
- Unmapped: 0

---
*Requirements defined: 2026-02-07*
*Last updated: 2026-02-09 after Phase 3 completion*

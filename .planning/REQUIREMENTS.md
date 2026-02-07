# Requirements: MonkeyLines

**Defined:** 2026-02-07
**Core Value:** Immediate clarity, nostalgic recognition, intuitive interaction

## v1 Requirements

Requirements for Monkey Island-styled UI revamp. Each maps to roadmap phases.

### Visual Foundation

- [ ] **VFND-01**: Monkey Island color palette defined as CSS custom properties (deep purples, teals, ambers, cream)
- [ ] **VFND-02**: Press Start 2P pixel font loaded (self-hosted WOFF2, antialiasing disabled)
- [ ] **VFND-03**: Consistent `image-rendering: pixelated` applied to all pixel art elements
- [ ] **VFND-04**: Base theme colors applied to body and container elements

### Button Controls

- [ ] **CTRL-01**: Button touch targets meet 44x44px minimum for mobile compatibility
- [ ] **CTRL-02**: Verb bar panel styling (dark background strip, SCUMM two-zone layout below scene)
- [ ] **CTRL-03**: Pixel art button icons created (16x16 or 32x32 sprites in MI palette)
- [ ] **CTRL-04**: Pixel art icons replace emoji buttons for all three controls

### Scene Presentation

- [ ] **SCNE-01**: Scene frame/border added around monkey viewport (pixel art or CSS in MI palette)
- [ ] **SCNE-02**: Speaker-colored speech text (different color per monkey head: left, center, right)
- [ ] **SCNE-03**: 9-slice border-image panels for authentic SCUMM-style scalable borders
- [ ] **SCNE-04**: Speech bubble restyled with pixel aesthetic (preserves fitText() functionality)

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
| VFND-01 | Phase 1 | Pending |
| VFND-02 | Phase 1 | Pending |
| VFND-03 | Phase 1 | Pending |
| VFND-04 | Phase 1 | Pending |
| CTRL-01 | Phase 2 | Pending |
| CTRL-02 | Phase 2 | Pending |
| CTRL-03 | Phase 2 | Pending |
| CTRL-04 | Phase 2 | Pending |
| SCNE-01 | Phase 3 | Pending |
| SCNE-02 | Phase 3 | Pending |
| SCNE-03 | Phase 3 | Pending |
| SCNE-04 | Phase 3 | Pending |
| PLSH-01 | Phase 4 | Pending |

**Coverage:**
- v1 requirements: 13 total
- Mapped to phases: 13
- Unmapped: 0

---
*Requirements defined: 2026-02-07*
*Last updated: 2026-02-07 after roadmap creation*

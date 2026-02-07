# Roadmap: MonkeyLines

## Overview

Transform MonkeyLines from a functional prototype into a visually authentic Monkey Island experience. Four phases move from visual foundation (palette, font) through interactive controls and scene framing to final polish, each building on the previous. The result is a portfolio-ready piece with cohesive SCUMM-era aesthetic.

## Phases

- [ ] **Phase 1: Visual Foundation** - Establish MI color palette, pixel font, and base theme
- [ ] **Phase 2: Control Panel** - Restyle buttons as SCUMM verb bar with pixel art icons
- [ ] **Phase 3: Scene and Speech** - Frame the monkey viewport and restyle speech bubbles
- [ ] **Phase 4: Polish** - CRT scanline overlay for retro finishing touch

## Phase Details

### Phase 1: Visual Foundation
**Goal**: The page looks and feels like a Monkey Island screen — colors, typography, and base styling establish the retro aesthetic before any component work begins
**Depends on**: Nothing (first phase)
**Requirements**: VFND-01, VFND-02, VFND-03, VFND-04
**Success Criteria** (what must be TRUE):
  1. Page background, text, and containers use deep purples/teals/ambers from the MI palette — no default browser colors remain
  2. All text renders in a pixel font with crisp edges (no antialiasing smoothing visible)
  3. All pixel art elements scale without blurring at any viewport size
**Plans**: TBD

Plans:
- [ ] 01-01: TBD

### Phase 2: Control Panel
**Goal**: Users interact with buttons that look like SCUMM verb bar controls — recognizable, themed, and intuitive on both desktop and mobile
**Depends on**: Phase 1
**Requirements**: CTRL-01, CTRL-02, CTRL-03, CTRL-04
**Success Criteria** (what must be TRUE):
  1. Buttons display pixel art icons instead of emoji, styled in the MI palette
  2. Controls sit in a dark verb bar panel below the scene, visually separated from the viewport
  3. Every button is comfortably tappable on a phone (no mis-taps, no squinting)
**Plans**: TBD

Plans:
- [ ] 02-01: TBD

### Phase 3: Scene and Speech
**Goal**: The monkey scene and speech text feel like a framed game viewport with character-attributed dialog
**Depends on**: Phase 1
**Requirements**: SCNE-01, SCNE-02, SCNE-03, SCNE-04
**Success Criteria** (what must be TRUE):
  1. A visible border or frame surrounds the monkey scene, creating a "screen within screen" effect
  2. Speech text color differs per monkey head so the viewer can tell who is speaking
  3. Speech bubble has a pixel-art styled border that scales correctly and does not break text fitting
  4. Scalable 9-slice borders render without stretching or blurring artifacts
**Plans**: TBD

Plans:
- [ ] 03-01: TBD

### Phase 4: Polish
**Goal**: Subtle retro effects add finishing authenticity without hurting usability
**Depends on**: Phase 3
**Requirements**: PLSH-01
**Success Criteria** (what must be TRUE):
  1. A CRT scanline effect is visible on desktop but does not reduce text readability
  2. The overlay performs smoothly on mobile (no jank or dropped frames)
**Plans**: TBD

Plans:
- [ ] 04-01: TBD

## Progress

**Execution Order:**
Phases execute in numeric order: 1 → 2 → 3 → 4

| Phase | Plans Complete | Status | Completed |
|-------|----------------|--------|-----------|
| 1. Visual Foundation | 0/? | Not started | - |
| 2. Control Panel | 0/? | Not started | - |
| 3. Scene and Speech | 0/? | Not started | - |
| 4. Polish | 0/? | Not started | - |

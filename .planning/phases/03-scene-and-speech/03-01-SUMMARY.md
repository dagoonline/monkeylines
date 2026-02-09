---
phase: 03-scene-and-speech
plan: 01
subsystem: ui
tags: [css, border-image, 9-slice, pixel-art, speech-bubble, data-uri]

requires:
  - phase: 01-visual-foundation
    provides: MI color palette CSS custom properties
provides:
  - SCUMM-style 9-slice pixel border on monkey scene viewport
  - Per-speaker colored dialog text (amber/teal/orange)
  - Pixel-art speech bubble with dark background
affects: [04-final-polish]

tech-stack:
  added: []
  patterns: [inline base64 data URI for pixel-art borders, data-attribute CSS color targeting]

key-files:
  created: []
  modified: [index.html]

key-decisions:
  - "Inline base64 data URIs for border PNGs (no extra asset files)"
  - "Increased speech-text height from 14% to 16% to compensate for border box model"
  - "Default text color changed from --mi-text-dark to --mi-text for speech bubble readability"

patterns-established:
  - "data-speaker attribute pattern for per-monkey CSS styling"
  - "Inline pixel-art PNGs via base64 data URIs for 9-slice borders"

duration: 2min
completed: 2026-02-09
---

# Phase 3 Plan 1: Scene and Speech Summary

**SCUMM-style 9-slice pixel-art frame on monkey viewport with per-speaker colored dialog and pixel-art speech bubble**

## Performance

- **Duration:** 2 min
- **Started:** 2026-02-09T12:21:30Z
- **Completed:** 2026-02-09T12:23:15Z
- **Tasks:** 2
- **Files modified:** 1

## Accomplishments
- Monkey scene framed with decorative 24x24 pixel-art 9-slice border using MI palette
- Speech text now colored per speaking monkey head (amber L, teal C, orange R)
- Speech bubble restyled with 12x12 pixel-art border and dark background for readability

## Task Commits

Each task was committed atomically:

1. **Task 1: Scene frame with pixel-art 9-slice border** - `09de8aa` (feat)
2. **Task 2: Per-speaker text colors and speech bubble restyle** - `a91318e` (feat)

## Files Created/Modified
- `index.html` - Scene frame border, speaker color CSS properties, speech bubble restyle, data-speaker JS attribute

## Decisions Made
- Used inline base64 data URIs for border PNGs to avoid extra asset files
- Increased speech-text height from 14% to 16% to compensate for added 4px border in box model
- Changed default speech text color from --mi-text-dark to --mi-text for readability against dark bubble background

## Deviations from Plan

None - plan executed exactly as written.

## Issues Encountered

None.

## User Setup Required

None - no external service configuration required.

## Next Phase Readiness
- Scene and speech styling complete
- Ready for phase 04 final polish

---
*Phase: 03-scene-and-speech*
*Completed: 2026-02-09*

## Self-Check: PASSED

# Phase 1: Visual Foundation - Context

**Gathered:** 2026-02-07
**Status:** Ready for planning

<domain>
## Phase Boundary

Establish the Monkey Island visual aesthetic through CSS: color palette, pixel font typography, and pixel-perfect rendering. This is the foundation layer that all subsequent visual work depends on. No new functionality—purely CSS theming of existing structure.

</domain>

<decisions>
## Implementation Decisions

### Pixel Rendering Strategy

**Browser compatibility:**
- Modern browsers only (recent Chrome, Firefox, Safari)
- Assume baseline support for `image-rendering: pixelated` (Jan 2020+)
- Skip old browser fallbacks and vendor prefixes

### Claude's Discretion

**Rendering application:**
- Determine which elements get `image-rendering: pixelated` (all pixel art vs selective)
- Balance between element types (existing monkey images vs new UI elements)

**Mobile scaling:**
- Balance pixel consistency vs responsive flexibility
- Handle non-integer scaling tradeoffs (perfect grid vs fluid layouts)

**Pixel grid:**
- Determine if explicit grid system needed (like 4px base unit) or asset-driven approach
- Analyze existing monkey art dimensions to derive natural grid

**Color palette specifics:**
- Extract exact Monkey Island palette colors (research suggests purples, teals, ambers)
- Determine color application strategy (which elements get which colors)
- Define CSS custom properties structure

**Typography approach:**
- Load Press Start 2P pixel font (self-hosted WOFF2)
- Determine sizing strategy (exact multiples, minimum sizes for readability)
- Decide which elements use pixel font vs preserve monospace
- Apply `-webkit-font-smoothing: none` for crisp rendering

**Page layout & theming:**
- Background treatment (solid color, pattern, or subtle texture)
- Container styling within MI aesthetic
- Base theme colors applied to body/container elements

</decisions>

<specifics>
## Specific Ideas

- Research identified deep purples (#1a0a2e), Caribbean teals (#2d8b6e), warm ambers (#c8a43c), cream (#e8d4a0) as MI palette
- Press Start 2P is the recommended pixel font (well-established, OFL license)
- Existing code already uses `image-rendering: pixelated` on monkey scene—extend this pattern
- Must preserve existing responsive behavior (`clamp()`, `aspect-ratio`, single media query)

</specifics>

<deferred>
## Deferred Ideas

None — discussion stayed within phase scope.

</deferred>

---

*Phase: 01-visual-foundation*
*Context gathered: 2026-02-07*

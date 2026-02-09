# Phase 2: Control Panel - Research

**Researched:** 2026-02-09
**Domain:** Pixel art UI, SCUMM verb bar layout, mobile touch targets
**Confidence:** HIGH

## Summary

Phase 2 replaces the current emoji buttons (sword, repeat, mute) with pixel art icons in a SCUMM-style verb bar panel. The existing project is vanilla HTML/CSS/JS with no frameworks and all assets embedded in a Go binary. The MI palette and Press Start 2P font are already established from Phase 1.

The work is entirely frontend: create 16x16 or 32x32 pixel art sprites, build a dark verb bar panel below the monkey scene, and ensure 44x44px minimum touch targets on mobile. No libraries needed -- this is pure CSS layout and hand-crafted pixel art.

**Primary recommendation:** Use inline SVG or small PNG sprites for the three button icons, placed in a CSS-styled dark panel that visually separates from the scene above. Keep it simple -- three buttons in a horizontal bar with generous padding for mobile.

## Standard Stack

No libraries needed. This phase uses:

| Tool | Purpose | Why |
|------|---------|-----|
| Vanilla CSS | Verb bar layout, button styling | Project constraint: no frameworks |
| PNG sprites (or inline SVG) | Button icons at 16x16 or 32x32 | Embeddable in Go binary, pixelated rendering already works |
| `image-rendering: pixelated` | Crisp scaling of small sprites | Already used throughout project |

## Architecture Patterns

### Current Button Structure
```html
<div class="btn-row">
    <button class="dialogue-btn" id="reload-btn">⚔️</button>
    <button class="dialogue-btn" id="auto-btn">🔁</button>
    <button class="dialogue-btn" id="mute-btn">🔇</button>
</div>
```

### Target: SCUMM Verb Bar Layout
The SCUMM engine (Monkey Island 1-2) has a distinctive two-zone layout:
1. **Top zone**: Scene/viewport (already exists as `.monkey-scene`)
2. **Bottom zone**: Dark panel with verb buttons and inventory

The verb bar is a solid dark strip spanning the full width, visually distinct from the scene. In MI1, this was a darker shade with a clear horizontal divider.

### Recommended Structure
```html
<div class="container">
    <div class="monkey-scene">
        <!-- existing scene content -->
    </div>
    <div class="verb-bar">
        <button class="verb-btn" id="reload-btn" aria-label="New exchange">
            <img src="/images/icons/sword.png" alt="" width="32" height="32">
        </button>
        <button class="verb-btn" id="auto-btn" aria-label="Auto-play">
            <img src="/images/icons/repeat.png" alt="" width="32" height="32">
        </button>
        <button class="verb-btn" id="mute-btn" aria-label="Toggle sound">
            <img src="/images/icons/speaker.png" alt="" width="32" height="32">
        </button>
    </div>
</div>
```

### Verb Bar CSS Pattern
```css
.verb-bar {
    background: var(--mi-bg-dark);
    border-top: 3px solid var(--mi-border);
    padding: 12px;
    display: flex;
    justify-content: center;
    gap: 16px;
}

.verb-btn {
    min-width: 44px;
    min-height: 44px;
    padding: 6px 12px;
    background: var(--mi-bg-mid);
    border: 2px solid var(--mi-border);
    cursor: pointer;
    image-rendering: pixelated;
}

.verb-btn img {
    image-rendering: pixelated;
    display: block;
}

.verb-btn:active:not(:disabled) {
    transform: translateY(2px);
}

.verb-btn.active {
    border-color: var(--mi-teal);
}
```

### Anti-Patterns to Avoid
- **Over-engineering the layout**: This is 3 buttons in a row. Flexbox is sufficient, no grid needed.
- **Using CSS background-image for icons**: Inline `<img>` tags are simpler and work with `image-rendering: pixelated` reliably.
- **Rounded corners on verb buttons**: SCUMM buttons were rectangular/boxy. Use small or zero border-radius for authenticity.

## Don't Hand-Roll

| Problem | Don't Build | Use Instead | Why |
|---------|-------------|-------------|-----|
| Pixel art icons | Complex SVG paths | Actual pixel art PNGs drawn at 16x16 or 32x32 | Authenticity; scaled up with `image-rendering: pixelated` |
| Touch target sizing | Custom JS hit area expansion | CSS `min-width`/`min-height` + `padding` | Pure CSS solution, no complexity needed |

## Common Pitfalls

### Pitfall 1: Pixel Art at Wrong Base Size
**What goes wrong:** Creating icons at arbitrary sizes (e.g., 48x48) then scaling down, losing the pixel grid.
**How to avoid:** Draw at exact 16x16 or 32x32. Scale UP only, using `image-rendering: pixelated`. Display at 2x or 4x the native size.

### Pitfall 2: Touch Targets Too Small
**What goes wrong:** Icon is 32x32 but button has no padding, failing the 44x44 minimum.
**How to avoid:** Set `min-width: 44px; min-height: 44px` on the button element. Add padding around the icon. Test on actual phone.
**Warning signs:** Buttons look fine on desktop but require precision tapping on mobile.

### Pitfall 3: Emoji Fallback in Mute Toggle
**What goes wrong:** The mute button currently toggles between emoji characters in JS (`muteBtn.textContent = muted ? '🔇' : '🔈'`). Must update this JS logic to swap icon images instead.
**How to avoid:** Change the JS to toggle `src` attribute or swap CSS class/visibility on two icon images.

### Pitfall 4: Verb Bar Not Full Width on Mobile
**What goes wrong:** Container max-width (700px) leaves the verb bar floating instead of spanning edge-to-edge.
**How to avoid:** The existing `@media (max-aspect-ratio: 1/1)` already removes container max-width on mobile. Verify the verb bar inherits full width correctly.

### Pitfall 5: Active State Visibility
**What goes wrong:** Auto-play active state not visible enough on the pixel art button.
**How to avoid:** Use `var(--mi-teal)` border highlight (already used) plus optionally brighten the icon or add a glow. Keep it consistent with existing `.active` class behavior.

## Code Examples

### Icon Sprite Specifications
- **Sword icon** (new exchange): Crossed swords or single cutlass, 32x32px
- **Repeat icon** (auto-play): Circular arrows, 32x32px
- **Speaker icon** (mute/unmute): Two variants -- speaker with sound waves, speaker with X

All icons should use only MI palette colors:
- `#e8d4a0` (cream) for primary icon fill
- `#c8a43c` (amber) for highlights/accents
- `#2d8b6e` (teal) for active state variants
- Transparent background

### Mute Toggle JS Update
```javascript
// Current:
muteBtn.textContent = muted ? '🔇' : '🔈';

// Replace with:
var muteIcon = muteBtn.querySelector('img');
muteIcon.src = muted ? '/images/icons/speaker-off.png' : '/images/icons/speaker-on.png';
```

### PNG Creation Approach
Create icons at 32x32 in any pixel art editor (Aseprite, Piskel, or even manually). Export as PNG with transparency. Files will be tiny (under 1KB each). Place in `/images/icons/`.

Alternatively, inline the pixel data as base64 data URIs to avoid extra HTTP requests (since Go embeds everything anyway, this is optional).

## Open Questions

1. **Icon style -- outlined or filled?**
   - SCUMM verb icons were simple filled shapes
   - Recommendation: Simple filled silhouettes in cream color, matching the text color
   - This is a design choice the implementer should make based on visual testing

2. **Labels under icons?**
   - SCUMM verbs had text labels ("Open", "Pick up", etc.)
   - With only 3 buttons, labels may not be needed if icons are clear
   - Recommendation: Add small text labels below icons in Press Start 2P at ~8px for clarity, especially since these aren't standard game verbs

## Sources

### Primary (HIGH confidence)
- Existing codebase analysis: `index.html` current button implementation
- Phase 1 context: established MI palette CSS custom properties
- PROJECT.md: confirmed vanilla-only constraint, embedded assets requirement

### Secondary (MEDIUM confidence)
- SCUMM interface reference: well-documented game UI pattern (dark panel, two-zone split)
- 44x44px touch target: Apple HIG and WCAG established standard

## Metadata

**Confidence breakdown:**
- Standard stack: HIGH - no libraries, vanilla CSS/HTML only
- Architecture: HIGH - straightforward CSS layout, well-understood patterns
- Pitfalls: HIGH - based on direct analysis of existing code
- Pixel art specs: MEDIUM - icon design is subjective, specs are recommendations

**Research date:** 2026-02-09
**Valid until:** 2026-03-09 (stable domain, no moving parts)

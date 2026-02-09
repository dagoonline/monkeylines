# Phase 3: Scene and Speech - Research

**Researched:** 2026-02-09
**Domain:** CSS border-image 9-slice, pixel art UI, speech styling
**Confidence:** HIGH

## Summary

This phase is pure CSS/HTML work on an existing vanilla single-page app. The four requirements (scene frame, speaker colors, 9-slice borders, speech bubble restyle) all use well-established CSS techniques. No new libraries needed.

The key technical challenge is CSS `border-image` with 9-slice scaling for SCUMM-style frames. This requires creating a pixel-art border image asset and using `border-image-slice` / `border-image-repeat` to scale it without distortion. The rest is straightforward CSS color changes and styling.

**Primary recommendation:** Use CSS `border-image` with a custom pixel-art PNG for the scene frame and speech bubble. Use CSS custom properties for per-speaker text colors.

## Standard Stack

### Core
| Library | Version | Purpose | Why Standard |
|---------|---------|---------|--------------|
| CSS border-image | native | 9-slice scalable borders | Built into all browsers, no JS needed |
| CSS custom properties | native | Speaker color theming | Already used in project (--mi-*) |

No external libraries needed. This is all vanilla CSS.

### Alternatives Considered
| Instead of | Could Use | Tradeoff |
|------------|-----------|----------|
| border-image PNG | Pure CSS box-shadow pixel borders | More code, harder to maintain, less authentic |
| border-image PNG | SVG borders | Overkill, SVG doesn't render pixel-art well without extra work |

## Architecture Patterns

### Asset Requirements

A 9-slice border image must be created as a PNG. Recommended approach:

```
images/
├── borders/
│   ├── scene-frame.png    # 9-slice source for monkey scene frame
│   └── speech-bubble.png  # 9-slice source for speech bubble
```

The PNG should be a small pixel-art image (e.g., 24x24 or 48x48) with:
- Fixed corner regions (e.g., 8px each corner)
- Repeatable edge regions
- Transparent or filled center

### Pattern 1: CSS border-image 9-Slice
**What:** CSS splits a single image into 9 regions. Corners stay fixed size, edges repeat/stretch.
**When to use:** Any scalable decorative border.
**Example:**
```css
.scene-frame {
    border: 8px solid transparent;
    border-image-source: url('/images/borders/scene-frame.png');
    border-image-slice: 8;       /* pixels from each edge to cut */
    border-image-width: 8px;
    border-image-repeat: round;  /* 'round' adjusts tile count to avoid partial tiles */
    border-image-outset: 0;
    image-rendering: pixelated;  /* critical for pixel art */
}
```

### Pattern 2: Speaker-Colored Text
**What:** Each monkey head gets a distinct text color via CSS custom property.
**Example:**
```css
:root {
    --mi-speaker-l: #e8c84c;   /* left head - amber/yellow */
    --mi-speaker-c: #3aaf8a;   /* center head - teal/green */
    --mi-speaker-r: #c87040;   /* right head - orange/rust */
}

.speech-text[data-speaker="l"] { color: var(--mi-speaker-l); }
.speech-text[data-speaker="c"] { color: var(--mi-speaker-c); }
.speech-text[data-speaker="r"] { color: var(--mi-speaker-r); }
```

JS sets `data-speaker` attribute during typeText():
```javascript
el.setAttribute('data-speaker', monkey);
```

### Pattern 3: Speech Bubble with Preserved fitText()
**What:** Restyle speech bubble with 9-slice border while keeping text fitting.
**Key constraint:** `fitText()` relies on `el.scrollHeight > el.clientHeight`. Adding border/padding changes clientHeight, so padding values must be accounted for.

### Anti-Patterns to Avoid
- **Stretching border edges:** Use `border-image-repeat: round` not `stretch` for pixel art - stretch causes blurring between pixels.
- **Non-integer border sizes:** Pixel art borders must use integer pixel values to avoid sub-pixel rendering artifacts.
- **border-radius with border-image:** `border-radius` is ignored when `border-image` is active. Don't try to combine them.

## Don't Hand-Roll

| Problem | Don't Build | Use Instead | Why |
|---------|-------------|-------------|-----|
| Scalable pixel borders | JS canvas drawing | CSS border-image | Native, performant, zero JS |
| Per-speaker colors | Inline style manipulation | data-attribute + CSS selectors | Cleaner separation of concerns |

## Common Pitfalls

### Pitfall 1: image-rendering on border-image
**What goes wrong:** Border image appears blurry/smoothed despite small pixel source.
**Why it happens:** `image-rendering: pixelated` on the element doesn't always affect border-image in all browsers.
**How to avoid:** Test across browsers. If needed, use a larger source image (2x or 4x the pixel art) so upscaling is less noticeable. The project already disables anti-aliasing globally which helps.
**Warning signs:** Blurry edges on the frame border.

### Pitfall 2: border-image-slice values must match the asset
**What goes wrong:** Corners appear cut off or edges don't tile properly.
**Why it happens:** The slice value must exactly match the corner size in the source PNG.
**How to avoid:** Design the PNG with known corner dimensions first, then set slice to match.

### Pitfall 3: Speech bubble padding affecting fitText()
**What goes wrong:** Text overflows or font size is too small after adding border.
**Why it happens:** `fitText()` compares scrollHeight to clientHeight. border-image with border-width changes the box model.
**How to avoid:** After adding border styling, verify fitText() still works. May need to adjust the height percentage or padding values in `.speech-text`.

### Pitfall 4: border-image-outset can cause overflow
**What goes wrong:** Frame extends outside container, causing scrollbars or clipping.
**Why it happens:** `border-image-outset` pushes the border outside the border box.
**How to avoid:** Keep outset at 0 or account for it in container sizing.

## Code Examples

### Scene Frame Implementation
```css
.monkey-scene {
    position: relative;
    width: 100%;
    aspect-ratio: 1 / 1;
    background: var(--mi-bg-dark);
    /* 9-slice frame */
    border: 8px solid transparent;
    border-image-source: url('/images/borders/scene-frame.png');
    border-image-slice: 8;
    border-image-width: 8px;
    border-image-repeat: round;
    image-rendering: pixelated;
}
```

### Speaker Color JS Integration
```javascript
// In typeText function, set speaker attribute
function typeText(text, monkey) {
    return new Promise(function(resolve) {
        el.setAttribute('data-speaker', monkey);
        // ... rest of existing typeText code
    });
}
```

### Speech Bubble Restyle
```css
.speech-text {
    position: absolute;
    top: 8%;
    left: 24%;
    width: 54%;
    height: 14%;
    padding: 4px 8px;
    font-family: 'Press Start 2P', monospace;
    font-size: clamp(8px, 1.4vw, 12px);
    line-height: 1.2;
    text-align: left;
    overflow: hidden;
    /* Pixel-art speech bubble border */
    border: 4px solid transparent;
    border-image-source: url('/images/borders/speech-bubble.png');
    border-image-slice: 4;
    border-image-width: 4px;
    border-image-repeat: round;
    background: var(--mi-bg-dark);
    background-clip: padding-box;
}
```

## State of the Art

| Old Approach | Current Approach | When Changed | Impact |
|--------------|------------------|--------------|--------|
| CSS box-shadow pixel grids | border-image 9-slice | Stable since CSS3 | Simpler, more performant |
| JS-driven border rendering | Pure CSS border-image | N/A | No JS overhead |

`border-image` has full browser support (97%+ on caniuse). No concerns about compatibility.

## Open Questions

1. **Border asset creation**
   - What we know: Need pixel-art PNGs for scene frame and speech bubble
   - What's unclear: Whether to hand-draw these or generate programmatically
   - Recommendation: Create small (e.g., 24x24) pixel-art PNGs in the MI color palette. Can be done in any pixel editor or even as tiny inline data URIs.

2. **Speaker color choices**
   - What we know: Need 3 distinct colors from MI palette that are readable against the speech bubble background
   - What's unclear: Exact color values for best readability
   - Recommendation: Use existing palette colors (amber for left, teal for center, a new warm tone for right). Test contrast.

3. **Speech bubble background**
   - What we know: Currently speech-text has no visible background, text is dark on the base image
   - What's unclear: Whether speech bubble should have a solid background fill or remain transparent with just a border
   - Recommendation: Add a semi-opaque or solid dark background so colored speaker text is readable. Use `border-image-slice: N fill` (the `fill` keyword fills the center region).

## Sources

### Primary (HIGH confidence)
- [MDN border-image-slice](https://developer.mozilla.org/en-US/docs/Web/CSS/Reference/Properties/border-image-slice)
- Existing project code (index.html) - verified current implementation

### Secondary (MEDIUM confidence)
- [CSS-Tricks border-image-slice](https://css-tricks.com/almanac/properties/b/border-image-slice/)
- [9-slicer tool](https://leanrada.com/9-slicer/) - useful for generating border-image CSS
- [TheLinuxCode border-image guide](https://thelinuxcode.com/css-border-image-property-a-practical-production-ready-guide/)

## Metadata

**Confidence breakdown:**
- Standard stack: HIGH - native CSS, well-documented
- Architecture: HIGH - straightforward CSS patterns
- Pitfalls: HIGH - well-known browser behavior
- Asset creation: MEDIUM - depends on design choices

**Research date:** 2026-02-09
**Valid until:** 2026-06-09 (stable CSS features, long validity)

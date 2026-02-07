# Technology Stack: Pixel Art UI for MonkeyLines

**Project:** MonkeyLines - Monkey Island aesthetic revamp
**Researched:** 2026-02-07
**Constraint:** Vanilla HTML/CSS/JS only (no frameworks)

## Recommended Stack

### CSS Pixel Rendering

| Technique | Purpose | Confidence |
|-----------|---------|------------|
| `image-rendering: pixelated` | Nearest-neighbor scaling for all pixel art assets | HIGH - MDN confirms baseline since Jan 2020 |
| `image-rendering: crisp-edges` | Fallback (Firefox historically preferred this) | HIGH - MDN documented |
| `-webkit-font-smoothing: none` | Disable antialiasing on pixel fonts (WebKit) | HIGH - MDN verified, non-standard but widely supported |
| `font-smooth: never` | Disable antialiasing (non-standard, limited use) | MEDIUM - non-standard per MDN |

**Critical CSS pattern for all pixel art elements:**

```css
.pixel-art {
  image-rendering: pixelated;
  image-rendering: crisp-edges; /* fallback */
  -ms-interpolation-mode: nearest-neighbor; /* legacy IE, drop if not needed */
}

.pixel-text {
  -webkit-font-smoothing: none;
  -moz-osx-font-smoothing: grayscale;
  font-smooth: never;
}
```

**Why this matters:** Without `image-rendering: pixelated`, browsers bilinear-interpolate scaled images, turning crisp pixel art into blurry mush. This is the single most important CSS property for the entire project.

### Asset Format

| Format | Use Case | Why |
|--------|----------|-----|
| **PNG** | All pixel art sprites, UI elements, backgrounds | Lossless compression preserves exact pixel colors. No generation loss. Already used in project. |
| **SVG** | NOT recommended for pixel art | Vector format antithetical to pixel aesthetic. Only acceptable for non-game UI (e.g., GitHub icon already in project). |
| **WebP** | Do NOT use for pixel art | Lossy by default, even lossless WebP can introduce subtle color shifts. PNG is the correct choice. |
| **AVIF** | Do NOT use | Same lossy concerns, worse browser support than PNG |
| **ICO/favicon** | Keep existing | Already present, fine as-is |

**Asset creation guidelines:**
- Author sprites at 1x native pixel resolution (e.g., 320x200 for SCUMM-era scenes)
- Scale up via CSS, never in the image file
- Use indexed color PNGs (palette mode) for smallest file size and color accuracy
- Run through `optipng` or `pngcrush` for lossless compression
- Sprite sheets over individual files when assets share a scene (fewer HTTP requests)

### Pixel Fonts

| Font | Type | License | Why |
|------|------|---------|-----|
| **Press Start 2P** | Bitmap-style, Google Fonts | OFL | The standard pixel font for retro web. 8x8 grid, clean at all multiples. Widely recognized retro aesthetic. |
| **Silkscreen** | Bitmap, 6px base | OFL | Smaller, good for UI labels and secondary text |
| **Custom bitmap font** | Self-authored or extracted | Varies | For maximum Monkey Island authenticity, create a font matching the SCUMM engine dialog font |

**Recommended approach:** Use Press Start 2P as primary via self-hosted WOFF2 (do NOT use Google Fonts CDN -- adds external dependency, privacy concern, and a render-blocking request).

```css
@font-face {
  font-family: 'PixelFont';
  src: url('fonts/press-start-2p.woff2') format('woff2');
  font-display: swap;
}
```

**Why WOFF2:** MDN confirms it is the recommended web font format with best compression. Baseline support across all modern browsers.

**Critical font rendering rules:**
- Set font sizes to exact multiples of the base pixel size (8px, 16px, 24px, 32px for Press Start 2P)
- Non-multiple sizes cause sub-pixel rendering and blurry text
- Use `px` units, not `em`/`rem` for pixel fonts (rems cause fractional sizes)
- Apply `-webkit-font-smoothing: none` to disable antialiasing

### Color Palette

| Approach | Recommendation |
|----------|---------------|
| **Palette source** | Monkey Island 1 (1990) used 256-color VGA palette. Extract key colors from original game screenshots. |
| **CSS implementation** | CSS custom properties (variables) for the full palette |
| **Palette size** | 16-32 named colors maximum for UI consistency |

**Recommended palette structure:**

```css
:root {
  /* Monkey Island inspired - dark purples, teals, warm browns */
  --mi-bg-dark: #1a0a2e;       /* deep purple-black, night sky */
  --mi-bg-mid: #2d1b4e;        /* purple midtone */
  --mi-text-primary: #e8d4a0;  /* warm parchment/dialog text */
  --mi-text-highlight: #ffffff;
  --mi-accent-teal: #2d8b6e;   /* Caribbean teal */
  --mi-accent-gold: #c8a43c;   /* gold accents */
  --mi-ui-border: #5a3a1e;     /* wood/leather border brown */
  --mi-ui-bg: #3a2010;         /* dark wood panel */
  --mi-shadow: #0a0412;        /* near-black shadow */
}
```

**Why CSS variables:** Zero runtime cost, works in vanilla CSS, easy to theme, supported everywhere since 2017.

### Layout and Scaling

| Technique | Purpose | Confidence |
|-----------|---------|------------|
| CSS `aspect-ratio` | Lock game viewport to retro aspect ratios (4:3 or 16:9) | HIGH - already used in project |
| `clamp()` for sizing | Responsive pixel-snapped sizing | HIGH - already used in project |
| Integer scaling via `transform: scale()` | Scale entire game viewport by integer multiples | HIGH |
| CSS Grid/Flexbox | UI layout (verb bar, inventory, dialog) | HIGH |

**Pixel-perfect scaling strategy:**

The SCUMM engine ran at 320x200. For authentic feel:
1. Design UI elements at 320px-wide base resolution
2. Scale up the game container to fill viewport using integer multiples
3. Use `transform: scale(N)` where N is integer, or let `image-rendering: pixelated` handle non-integer CSS scaling

```css
.game-viewport {
  width: 320px;
  height: 200px;
  transform-origin: top left;
  /* JS calculates largest integer scale that fits viewport */
}
```

Alternatively (simpler, what the project already does): use percentage/viewport widths and let `image-rendering: pixelated` handle scaling. This is less "pure" but more practical for responsive design.

**Recommendation:** Keep the current responsive approach (percentage widths + `image-rendering: pixelated`). Integer scaling is purist but creates awkward gaps on non-matching screen sizes. The current project already handles this well.

### UI Component Patterns

| Pattern | CSS Technique | Notes |
|---------|--------------|-------|
| Dialog box | `border-image` with 9-slice PNG | Pixel-perfect scalable borders from a single sprite |
| Buttons | `border-image` or `box-shadow` pixel borders | Replace current rounded borders with pixelated ones |
| Text crawl | Current typewriter approach works | Already implemented, just restyle |
| SCUMM verb bar | CSS Grid, 3x3 or 3x5 grid of verbs | Classic MI interface element |
| Inventory strip | Horizontal scrollable flex container | Bottom-of-screen inventory items |

**`border-image` for pixel art panels (key technique):**

```css
.dialog-box {
  border-image-source: url('ui/panel-border.png');
  border-image-slice: 8 fill;  /* 8px border region */
  border-image-width: 16px;    /* scale up 2x */
  border-image-repeat: round;
  image-rendering: pixelated;
}
```

This lets a single 24x24 PNG (with 8px border edges) scale to any dialog box size while keeping pixel-perfect corners.

### Animation

| Technique | Use Case | Notes |
|-----------|----------|-------|
| CSS `steps()` timing | Sprite sheet animation | `animation-timing-function: steps(N)` for frame-by-frame |
| CSS `visibility` toggle | Mouth open/close (already used) | Keep current approach |
| `requestAnimationFrame` | Complex JS-driven animation | Only if CSS steps insufficient |

```css
.sprite-anim {
  background: url('spritesheet.png');
  animation: walk 0.5s steps(4) infinite;
}

@keyframes walk {
  to { background-position: -128px 0; } /* 4 frames x 32px */
}
```

**Why `steps()`:** Native CSS, no JS needed, gives authentic frame-by-frame feel. No interpolation between frames = pixel-perfect.

## What NOT to Use

| Technology | Why Not |
|------------|---------|
| **Canvas/WebGL** | Overkill. CSS handles pixel art rendering perfectly. Adds complexity, accessibility issues, breaks text selection. |
| **Any JS framework** (React, Vue, etc.) | Project constraint is vanilla. Also unnecessary for this scope. |
| **CSS `filter: pixelate()`** | Does not exist. A common misconception. |
| **SVG for pixel art** | Anti-aliased by nature. Wrong tool entirely. |
| **Google Fonts CDN** | External dependency. Self-host WOFF2 instead. |
| **rem/em for pixel font sizes** | Causes fractional pixel sizes that blur bitmap fonts. Use px. |
| **CSS Houdini / Paint API** | Poor browser support (no Firefox/Safari). Not needed. |
| **Phaser / PixiJS** | Full game engines. Massive overkill for a UI revamp. Breaks vanilla constraint. |
| **WebP/AVIF for sprites** | Lossy formats corrupt pixel art. PNG is correct. |

## Alternatives Considered

| Category | Recommended | Alternative | Why Not |
|----------|-------------|-------------|---------|
| Image format | PNG | WebP lossless | Marginal size savings, broader tooling support for PNG, zero risk of lossy conversion |
| Font delivery | Self-hosted WOFF2 | Google Fonts CDN | External dependency, extra DNS lookup, privacy |
| Scaling | CSS `image-rendering` | Canvas with manual nearest-neighbor | Unnecessary complexity for static/semi-static UI |
| Borders | `border-image` 9-slice | Multiple positioned divs | 9-slice is specifically designed for this; fewer elements, cleaner code |
| Color management | CSS custom properties | Sass variables | Adds build step. Vanilla constraint. CSS vars work natively. |
| Pixel font | Press Start 2P | Dogica, Visitor | Press Start 2P has best readability and widest recognition |

## Installation / Setup

No npm. No build step. File structure:

```
/fonts/
  press-start-2p.woff2      # Self-hosted pixel font
/images/
  base.png                   # Existing monkey scene
  l.png, c.png, r.png        # Existing overlays
  ui/
    panel-border.png          # 9-slice dialog border
    button-border.png         # 9-slice button border
    inventory-bg.png          # Inventory strip background
    icons/                    # Pixel art icons replacing emoji
      sword.png               # Replace sword emoji
      repeat.png              # Replace repeat emoji
      speaker.png             # Replace mute emoji
```

**Asset pipeline (manual, no build tools):**
1. Author pixel art at 1x in any pixel editor (Aseprite, Piskel, GIMP)
2. Export as indexed-color PNG
3. Optimize with `optipng -o7 file.png` (lossless, typically 20-40% smaller)
4. Place in `/images/ui/`

## Mobile Compatibility Notes

- `image-rendering: pixelated` works on iOS Safari 13+, Chrome Android, Samsung Internet -- all current
- Pixel fonts at small sizes (8px) may be unreadable on mobile. Use 16px minimum (2x base) on small screens
- Touch targets: pixel art buttons must still meet 44x44px minimum tap target (Apple HIG)
- `clamp()` already used in project for responsive sizing -- keep this pattern
- High-DPI screens (Retina): `image-rendering: pixelated` handles this correctly; no special treatment needed

## Sources

- MDN `image-rendering`: https://developer.mozilla.org/en-US/docs/Web/CSS/image-rendering (verified 2026-02-07, baseline since Jan 2020)
- MDN `@font-face`: https://developer.mozilla.org/en-US/docs/Web/CSS/@font-face (verified 2026-02-07, WOFF2 recommended)
- MDN `font-smooth`: https://developer.mozilla.org/en-US/docs/Web/CSS/font-smooth (verified 2026-02-07, non-standard)
- Press Start 2P: https://fonts.google.com/specimen/Press+Start+2P (OFL license)
- CSS `border-image`: https://developer.mozilla.org/en-US/docs/Web/CSS/border-image (baseline, 9-slice pattern)
- CSS `steps()`: https://developer.mozilla.org/en-US/docs/Web/CSS/animation-timing-function (baseline)

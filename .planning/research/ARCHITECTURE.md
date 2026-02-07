# Architecture Patterns

**Domain:** Pixel art game UI for single-page web application (Go-embedded)
**Researched:** 2026-02-07

## Current Architecture

MonkeyLines is a single `index.html` with inline CSS and JS, served from a Go binary via `embed.FS`. Images are individual PNGs served from an `/images/` path. The architecture is intentionally simple: one template, one API endpoint, no build step.

### Current Component Map

```
index.html
  +-- inline <style> (reset, layout, buttons, responsive)
  +-- .container
  |     +-- .monkey-scene (aspect-ratio box, relative positioning)
  |     |     +-- base.png (always visible)
  |     |     +-- l.png, c.png, r.png (overlay toggled via .monkey-visible)
  |     |     +-- .speech-text (absolute positioned over scene)
  |     +-- .btn-row (flexbox, 3 emoji buttons)
  +-- inline <script> (IIFE, no modules)
```

## Recommended Architecture for Themed UI

Keep the single-file inline approach. Do not introduce a build system. The constraint (Go embed, no build step) is a feature, not a limitation.

### Component Boundaries

| Component | Responsibility | Depends On |
|-----------|---------------|------------|
| Page Background | Full-viewport themed background (color, pattern, or tiled image) | Nothing |
| Scene Frame | Decorative border/frame around monkey scene | Background |
| Monkey Scene | Existing overlay animation system (unchanged) | Scene Frame (visual only) |
| Speech Bubble | Styled speech area replacing plain `.speech-text` | Monkey Scene (positioning) |
| Control Panel | Pixel art buttons in themed layout | Scene Frame (visual alignment) |
| Footer | GitHub link (existing, minimal change) | Nothing |

### Proposed DOM Structure

```html
<body>  <!-- themed background via CSS -->
  <div class="container">
    <div class="scene-frame">        <!-- NEW: decorative border -->
      <div class="monkey-scene">     <!-- existing, unchanged internally -->
        <img base>
        <img overlays x3>
        <div class="speech-bubble">  <!-- RESTYLE of speech-text -->
      </div>
    </div>
    <div class="control-panel">      <!-- RESTYLE of btn-row -->
      <button class="pixel-btn">     <!-- RESTYLE of dialogue-btn -->
    </div>
  </div>
</body>
```

## Asset Management

### Recommendation: Individual PNGs, Not Sprite Sheets

**Why not sprite sheets:**
- Only 4 existing images + a handful of new decorative assets
- Sprite sheets add complexity (background-position math, maintenance burden)
- No build tool to generate sprite sheets automatically
- Individual files are simpler to update and debug
- Go's `embed.FS` handles multiple files fine

**Why not icon fonts:**
- Pixel art buttons are better as images (precise pixel control)
- Icon fonts require font file generation tooling
- Emoji already works for simple icons; pixel art replacements are images

**Why not inline base64:**
- Current images are served via `/images/` with cache headers (7-day cache)
- Base64 bloats the HTML and defeats caching
- Keep the current approach

### New Assets Needed

| Asset | Format | Purpose | Size Target |
|-------|--------|---------|-------------|
| Background pattern tile | PNG | Repeating background | < 2KB (small tile) |
| Frame border pieces | PNG or pure CSS | 9-slice or CSS border-image | < 1KB each if images |
| Button icons (sword, loop, speaker) | PNG | Replace emoji buttons | 16x16 or 32x32 px |
| Speech bubble tail/corners | CSS or PNG | Styled speech container | Prefer pure CSS |

### Asset Format Guidance

- **PNG-8** for pixel art (indexed color, tiny files, perfect edges)
- **No JPEG** for pixel art (lossy compression destroys sharp pixels)
- **No SVG** for pixel art (vector format, wrong paradigm)
- All assets must use `image-rendering: pixelated` when scaled

## CSS Architecture

### Layer Strategy (No CSS Layers Needed)

The app is small enough that CSS Layers (`@layer`) would be overengineering. Instead, organize the inline `<style>` block by section with comments:

```css
/* === RESET === */
/* === BACKGROUND & THEME === */
/* === SCENE FRAME === */
/* === MONKEY SCENE (existing) === */
/* === SPEECH BUBBLE === */
/* === CONTROL PANEL === */
/* === BUTTONS === */
/* === RESPONSIVE === */
```

### Z-Index Strategy

Only needed within `.monkey-scene` (already working) and for the frame:

| Element | z-index | Notes |
|---------|---------|-------|
| Background / frame | auto (0) | No z-index needed, normal flow |
| base.png | auto | First in DOM = bottom layer |
| Overlay PNGs | auto | Positioned absolute, above base |
| Speech bubble | 1 | Above overlays if they ever overlap |

No global z-index war. Keep it simple.

### Pixel-Perfect Scaling

The critical CSS property is already in use:

```css
image-rendering: pixelated;  /* Chrome, Edge */
image-rendering: crisp-edges; /* Firefox fallback */
```

Apply to ALL pixel art images and any element using `background-image` with pixel art.

For scaling pixel art cleanly:
- Use integer scale factors when possible (2x, 3x, 4x)
- `width: 100%` with `image-rendering: pixelated` works well (already proven in codebase)
- Avoid fractional pixel sizes (use `clamp()` with pixel-snapped values)

### Border/Frame Approaches (Ranked)

**1. CSS `border-image` with 9-slice (Recommended)**
- Single PNG, CSS does the slicing
- Scales naturally, minimal assets
- `border-image-slice` handles corners and edges
- Well-supported across browsers

**2. Pure CSS borders (pixel look)**
- `box-shadow` layering to simulate pixel borders
- No image assets needed
- Limited to simple patterns

**3. CSS `background-image` tiling**
- For repeating patterns on the background
- `background-repeat: repeat` with small tile
- Works perfectly for retro game backgrounds

### Speech Bubble Approach

Pure CSS is best here:
- Background color with hard pixel-style border (no border-radius, or 2-4px radius max)
- Optional: `box-shadow` for pixel drop shadow
- Tail/pointer via CSS pseudo-elements (`::after` with border trick)
- Keeps zero additional assets

## Responsive Strategy

### Current Approach (Keep and Extend)

The existing responsive design is solid:
- `clamp()` for fluid sizing
- `aspect-ratio: 1/1` for scene
- `@media (max-aspect-ratio: 1/1)` for portrait devices (removes padding, goes full-width)

### Extensions for Themed UI

- Background pattern: `background-repeat: repeat` tiles naturally at any viewport
- Frame: `border-image` scales with container (no media queries needed)
- Buttons: Continue using `clamp()` for sizing; replace emoji with `<img>` tags using same sizing
- Speech bubble: Stays percentage-positioned within scene (already responsive)

### No Breakpoint Explosion

Do NOT add traditional breakpoints. The current fluid approach with `clamp()` and aspect-ratio queries is superior for this use case. At most, keep the single existing `@media (max-aspect-ratio: 1/1)` rule.

## Anti-Patterns to Avoid

### Anti-Pattern: Introducing a Build System
**Why bad:** The project's strength is zero build complexity. Go embeds raw files.
**Instead:** Author CSS and JS directly. Use PNG-8 exported from any pixel art tool.

### Anti-Pattern: CSS Framework for Theming
**Why bad:** Tailwind, Bootstrap, etc. are massive overkill for a single-page novelty app.
**Instead:** Hand-written inline CSS, organized by comments.

### Anti-Pattern: Canvas Rendering for UI
**Why bad:** The monkey animation works perfectly with IMG overlays and visibility toggling. Canvas would require rewriting the entire rendering approach.
**Instead:** Keep DOM-based rendering. CSS handles theming, IMG handles pixel art.

### Anti-Pattern: Over-Componentizing
**Why bad:** This is one HTML file. Abstracting into Web Components or a framework adds complexity with zero benefit.
**Instead:** Simple CSS classes, simple DOM manipulation.

## Build Order (Dependencies)

Phase ordering based on what depends on what:

```
1. Background theming (CSS only, no dependencies)
   |
2. Scene frame / border (wraps existing scene, CSS + optional border-image)
   |
3. Speech bubble restyle (CSS rework of .speech-text, inside scene)
   |
4. Button pixel art (new assets + CSS rework of .btn-row / .dialogue-btn)
   |
5. Layout polish (spacing, alignment, responsive tuning)
```

**Rationale:**
- Steps 1-2 are visual wrappers that don't touch existing JS logic
- Step 3 reskins an existing element but must not break text fitting (`fitText()`)
- Step 4 requires new image assets and replaces emoji with `<img>` elements
- Step 5 is final polish once all pieces are in place

Steps 1 and 4 could be parallelized (no dependency), but sequential is simpler for a single developer.

## Sources

- Existing codebase analysis (HIGH confidence)
- CSS `border-image` and `image-rendering: pixelated` are well-established CSS features (HIGH confidence)
- Pixel art scaling best practices are well-documented community consensus (MEDIUM confidence)

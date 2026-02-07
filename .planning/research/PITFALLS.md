# Domain Pitfalls

**Domain:** Pixel art / retro game UI for web applications
**Project:** MonkeyLines (Monkey Island-inspired insult sword fighting)
**Researched:** 2026-02-07

## Critical Pitfalls

Mistakes that destroy the retro illusion or make the project look amateur.

### Pitfall 1: Mixed Pixel Densities

**What goes wrong:** UI elements have visibly different pixel sizes. The scene art has 4px logical pixels, but a border or button uses 2px pixels, and text is at native resolution. The eye immediately registers this as "clip art collage" rather than "retro game."
**Why it happens:** Different assets created at different base resolutions, or mixing pixel art with non-pixel UI elements without thought.
**Consequences:** The entire retro aesthetic collapses. It looks like a cheap mockup, not a coherent design. Portfolio death sentence.
**Prevention:** Establish a single base pixel grid (e.g., all art at 128x128 or 256x256, all UI borders at the same logical pixel size). Every visual element must snap to this grid. Audit by zooming to 400% and checking that all "pixels" are the same physical size.
**Detection:** Screenshot at 2x zoom. If you can see pixels of different sizes, you have this problem.
**Phase relevance:** Asset creation. Must be decided before any new art is made.

### Pitfall 2: Blurry Pixel Art from Incorrect Scaling

**What goes wrong:** Pixel art gets bilinear-filtered into a blurry mess instead of crisp blocky pixels.
**Why it happens:** Missing `image-rendering: pixelated` (or `-webkit-image-rendering: crisp-edges` for Safari), or CSS transforms/animations that trigger sub-pixel rendering, or the container size is not an integer multiple of the source image size.
**Consequences:** The defining visual characteristic of pixel art is destroyed. Looks broken, not retro.
**Prevention:**
- Always set `image-rendering: pixelated` on pixel art images (MonkeyLines already does this correctly).
- Avoid CSS `transform: scale()` on pixel art; use width/height sizing instead.
- Be cautious with `transform: translate()` -- sub-pixel values cause blurring. Use `will-change` or ensure integer pixel positions.
- Test on Safari specifically (it handles `pixelated` differently and may need `crisp-edges`).
**Detection:** View on a high-DPI display. If edges look soft/blurred, scaling is wrong.
**Phase relevance:** CSS implementation. Verify on every target browser.

**Current status in MonkeyLines:** The project already applies `image-rendering: pixelated` to `.monkey-scene img`. Good foundation, but any new UI elements (dialogue boxes, inventory panels, verb bars) must also follow this rule.

### Pitfall 3: Non-Integer Scaling Artifacts

**What goes wrong:** When a pixel art image is scaled to a size that is not an integer multiple of its native resolution, some "pixels" render as 3 CSS pixels wide while neighbors are 2 CSS pixels wide. Creates a wobbly, uneven grid.
**Why it happens:** Responsive/fluid layouts where `width: 100%` on a pixel art image yields arbitrary pixel counts. A 128px-wide image in a 350px container means each source pixel is 2.734 CSS pixels -- guaranteed unevenness.
**Consequences:** Subtle but looks "off" to anyone paying attention. Ruins the crispness that defines good pixel art presentation.
**Prevention:**
- Constrain containers to sizes that are integer multiples of the source art resolution (e.g., 256px, 384px, 512px, 640px for 128px source art).
- Use `max-width` with stepped breakpoints rather than pure fluid width.
- Alternatively, accept the tradeoff on mobile where exact multiples are impossible -- the `pixelated` rendering still looks far better than bilinear, and mobile users are less likely to notice minor unevenness.
**Detection:** Zoom in on pixel art edges. If the grid is uneven (some blocks bigger than others), this is the cause.
**Phase relevance:** Responsive/layout phase. Particularly important if adding new framing elements.

### Pitfall 4: Typography That Breaks the Illusion

**What goes wrong:** Using a bitmap/pixel font that is either too small to read, too large relative to the art, or -- worst of all -- mixing a pixel font with native anti-aliased text.
**Why it happens:** Pixel fonts render poorly at non-native sizes. Or the developer uses a pixel font for headings but browser-default for body text.
**Consequences:** Readability disasters on mobile, or visual inconsistency that screams "unfinished."
**Prevention:**
- If using a pixel font: serve it at its exact designed size (e.g., 8px, 16px, 32px -- integer multiples only). Apply `-webkit-font-smoothing: none` / `font-smooth: never` to disable anti-aliasing.
- If using monospace system fonts (as MonkeyLines currently does): commit fully to this approach. Do not mix with a pixel font halfway. The monospace approach is actually cleaner for readability.
- Never scale a pixel font with `em` or `vw` units -- it will render at fractional sizes and look terrible.
**Detection:** Check text at multiple viewport sizes. If letterforms look inconsistent or blurry at certain sizes, the font scaling is wrong.
**Phase relevance:** Typography decisions in the CSS implementation phase.

**Current status in MonkeyLines:** Uses `font-family: monospace` with `clamp()` for fluid sizing. This works well and avoids pixel font headaches. If switching to a pixel font for theming, be very careful about the tradeoffs above.

## Moderate Pitfalls

Mistakes that cause visual debt or undermine the professional polish.

### Pitfall 5: CRT/Scanline Effects That Tank Performance

**What goes wrong:** Adding scanline overlays, CRT curvature, or chromatic aberration via CSS/canvas that kill frame rate, especially on mobile.
**Why it happens:** CSS `repeating-linear-gradient` scanlines on a large element, or `filter: blur()` for glow effects, trigger expensive repaints. Canvas post-processing shaders are heavy on low-end GPUs.
**Prevention:**
- Scanlines: Use a tiny repeating PNG tile (2px tall, 50% transparent) rather than a CSS gradient. Or bake scanlines directly into the art assets.
- CRT curvature: The current art already has a CRT monitor frame baked into the image. This is the correct approach -- no runtime cost.
- Glow effects: Use `box-shadow` sparingly. Avoid `filter: blur()` on animated elements.
- Test on a throttled CPU (Chrome DevTools > Performance > 4x slowdown).
**Prevention:** Keep post-processing minimal. The art itself should carry the retro feel, not runtime effects.
**Phase relevance:** Any phase adding visual effects. Test performance budget early.

### Pitfall 6: Oversized Assets and Slow Initial Load

**What goes wrong:** Pixel art source files are tiny (128x128) but get exported as large PNGs, or worse, as JPEGs that add compression artifacts to pixel art. Multiple overlay states multiply the payload.
**Why it happens:** Exporting at display resolution instead of native resolution, or using JPEG for pixel art (lossy compression smears clean edges).
**Consequences:** Slow initial load kills the "instant delight" factor critical for a portfolio piece.
**Prevention:**
- Export pixel art as PNG at native resolution (128x128 or 256x256). Let `image-rendering: pixelated` handle upscaling in the browser.
- Never use JPEG for pixel art. The compression artifacts are visible and ugly on hard pixel edges.
- Consider sprite sheets to reduce HTTP requests if adding many animation states.
- Compress PNGs with tools like `oxipng` or `pngquant`.
- Lazy-load overlay states that are not immediately visible.
**Detection:** Check Network tab. If pixel art PNGs are >50KB each at small native resolutions, something is wrong.
**Phase relevance:** Asset creation and optimization phase.

**Current status in MonkeyLines:** Uses PNG for art (good) and JPG for the OG preview image (acceptable for social sharing, not displayed in-app). The overlay pattern (base + l/c/r layers) is smart for animation. Verify native resolutions are small.

### Pitfall 7: Retro Aesthetic Over Usability

**What goes wrong:** Buttons, controls, and interactive elements are themed so heavily that users cannot tell what is clickable or what the controls do.
**Why it happens:** Prioritizing visual consistency over UX. Making buttons look like pixel art inventory items without affordances.
**Consequences:** Users stare at the screen not knowing what to do. For a portfolio piece, this means the viewer bounces before seeing the cool parts.
**Prevention:**
- Every interactive element must have a visible hover/active state.
- Controls should be immediately understandable without instructions. Emoji buttons (as MonkeyLines currently uses) are effective but limit theming potential.
- If replacing emoji with pixel art icons: ensure they have clear borders, hover highlights, and touch targets >= 44x44px.
- Test with someone who has never seen the app. If they cannot figure out the first action in 3 seconds, the UX is broken.
**Detection:** Five-second test: show someone the screen for 5 seconds, ask what they would click.
**Phase relevance:** Every phase that touches interactive elements.

### Pitfall 8: Dialogue Text Rendering Issues

**What goes wrong:** Text overlaid on pixel art scenes either clips, overflows, or is unreadable due to contrast issues at certain viewport sizes.
**Why it happens:** The speech bubble area is a fixed percentage of the image, but text length varies. Font sizing with `clamp()` does not account for all content lengths.
**Consequences:** Truncated dialogue ruins the core experience (insult sword fighting IS the text).
**Prevention:**
- The existing `fitText()` function that shrinks font size to fit is a solid approach.
- Ensure minimum readable font size (do not go below ~12px on mobile).
- Test with the longest possible insult/comeback strings.
- Ensure sufficient contrast between text color and speech bubble background.
**Detection:** Feed in the longest text from the messages database and check all viewport sizes.
**Phase relevance:** Any phase modifying the speech/dialogue system.

## Minor Pitfalls

Annoyances that are fixable but worth avoiding upfront.

### Pitfall 9: Inconsistent Color Palette

**What goes wrong:** New UI elements use colors that are close to but not exactly the same as the existing art palette. Results in a "something is off" feeling.
**Prevention:** Extract the exact color palette from the existing art (the dark purples, teals, warm sand tones visible in base.png). Define these as CSS custom properties. All new elements must use only these colors.
**Phase relevance:** Asset creation, before any new visual elements.

### Pitfall 10: Hover/Focus States That Break Pixel Aesthetic

**What goes wrong:** Default browser focus rings (blue glow) or smooth CSS transitions appear on themed elements, breaking the hard-edge retro feel.
**Prevention:** Replace default focus indicators with pixel-consistent alternatives (e.g., 2px solid borders in palette colors, no border-radius, no box-shadow blur). Use `outline-offset` and solid outlines. If using transitions, make them stepped (instant on/off) rather than smooth.
**Phase relevance:** CSS implementation phase, accessibility review.

### Pitfall 11: Audio That Annoys Instead of Delights

**What goes wrong:** Sound effects are too loud, too frequent, or too long. Users mute and never unmute. Or audio fails silently on mobile.
**Prevention:** Default to muted (MonkeyLines already does this -- good choice). Keep sounds short and quiet. The current monkey "hooting" approach with Web Audio API synthesis is lightweight and avoids asset loading issues.
**Phase relevance:** Any phase modifying audio.

**Current status in MonkeyLines:** Audio is well-handled. Starts muted, uses synthesized sounds (no asset loading), handles iOS AudioContext unlock. Low risk area.

## Mobile-Specific Pitfalls

### Pitfall 12: Touch Targets Too Small

**What goes wrong:** Pixel-art-styled buttons are sized to look proportional to the retro aesthetic but are too small to tap reliably on mobile.
**Prevention:** Minimum 44x44px touch targets regardless of visual size. The visual button can be smaller if the tappable area extends beyond it via padding.
**Phase relevance:** Responsive/mobile phase.

**Current status:** Buttons use `clamp()` for sizing and appear adequately large based on the current CSS.

### Pitfall 13: Viewport Overflow on Mobile

**What goes wrong:** On small/narrow viewports, the pixel art scene plus controls overflow the viewport, causing unwanted scrolling or cropping the art.
**Prevention:** Test at 320px width minimum. Use `min()` and `dvh` units. The current `aspect-ratio: 1/1` approach is good but could overflow on very short landscape viewports.
**Detection:** Test on iPhone SE (375x667) and in landscape orientation.
**Phase relevance:** Responsive phase.

## Phase-Specific Warnings

| Phase Topic | Likely Pitfall | Mitigation |
|-------------|---------------|------------|
| Asset creation | Mixed pixel densities (#1) | Define pixel grid before creating any art |
| Asset creation | Wrong export format (#6) | PNG only, native resolution, verify file sizes |
| Color/theming | Inconsistent palette (#9) | Extract palette from existing art first |
| CSS implementation | Blurry scaling (#2) | Test `image-rendering: pixelated` on all browsers |
| CSS implementation | Typography mismatch (#4) | Commit to monospace OR pixel font, not both |
| CSS implementation | Focus states breaking aesthetic (#10) | Style focus indicators with pixel-consistent approach |
| Responsive/mobile | Non-integer scaling (#3) | Accept tradeoff on mobile, use stepped sizes on desktop |
| Responsive/mobile | Touch targets (#12) | 44px minimum, test on real devices |
| Responsive/mobile | Viewport overflow (#13) | Test 320px width and landscape orientations |
| Visual effects | Performance (#5) | Bake effects into art, avoid runtime CSS filters |
| UX/interaction | Controls unclear (#7) | Five-second test with fresh eyes |
| Dialogue system | Text overflow (#8) | Test with longest strings at all viewport sizes |

## Sources

- Direct analysis of MonkeyLines codebase (index.html, image assets)
- Established CSS pixel art rendering best practices (`image-rendering: pixelated`, integer scaling)
- Web Audio API mobile unlock patterns (verified in existing codebase)
- WCAG touch target guidelines (44x44px minimum)
- Confidence: MEDIUM (based on domain expertise and codebase analysis; web searches for latest browser behavior were not available)

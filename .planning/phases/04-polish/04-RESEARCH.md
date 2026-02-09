# Phase 4: Polish - Research

**Researched:** 2026-02-09
**Domain:** CSS visual effects, retro CRT aesthetics
**Confidence:** HIGH

## Summary

CRT scanline effects can be implemented with CSS-only approaches using pseudo-elements and repeating linear gradients, offering excellent performance on both desktop and mobile. The core technique involves applying a `::before` or `::after` pseudo-element with a repeating gradient pattern that creates horizontal scanlines. For readability and accessibility, scanline opacity should remain subtle (0.15-0.25 range), with considerations for user motion preferences.

The research confirms that CSS gradients outperform both Canvas and SVG pattern approaches for this use case, especially on mobile devices. Key success factors include minimal color stops (2-3 maximum), GPU acceleration via transform properties, and proper layering with `pointer-events: none` to avoid blocking user interactions.

**Primary recommendation:** Implement scanlines using a CSS `::before` pseudo-element on `.monkey-scene` with a 2px or 4px repeating linear gradient pattern, opacity around 0.2, and responsive handling via `prefers-reduced-motion` media query.

## Standard Stack

The established libraries/tools for this domain:

### Core
| Library | Version | Purpose | Why Standard |
|---------|---------|---------|--------------|
| Pure CSS | CSS3 | Scanline rendering | Native browser support, GPU-accelerated, no dependencies |
| repeating-linear-gradient() | CSS3 | Pattern generation | Lightweight, performant, scalable |
| Pseudo-elements (::before/::after) | CSS3 | Overlay layer | Reduces DOM complexity, keeps overlays coupled to parent |

### Supporting
| Library | Version | Purpose | When to Use |
|---------|---------|---------|-------------|
| transform: translateZ(0) | CSS3 | GPU acceleration trigger | When animations are used or on lower-end mobile devices |
| will-change | CSS3 | Browser optimization hint | Sparingly for animated overlays (not needed for static scanlines) |
| @media (prefers-reduced-motion) | CSS3 | Accessibility compliance | Always - respects user motion preferences |

### Alternatives Considered
| Instead of | Could Use | Tradeoff |
|------------|-----------|----------|
| CSS gradients | Canvas overlay | Canvas requires JavaScript, higher computational overhead, battery drain on mobile |
| CSS gradients | SVG pattern via data URI | SVG patterns require paint layer calculations, more complex setup, minimal benefit |
| Pseudo-element | Real DOM element | Extra DOM node, less semantically clean, same performance |

**Installation:**
No external dependencies required - pure CSS3 features with universal browser support.

## Architecture Patterns

### Recommended Implementation Structure

Place scanline overlay as a pseudo-element on the main viewport container:

```css
.monkey-scene {
  position: relative;
  /* existing styles */
}

.monkey-scene::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: repeating-linear-gradient(
    to bottom,
    transparent 0,
    transparent 50%,
    rgba(0, 0, 0, 0.2) 50%,
    rgba(0, 0, 0, 0.2) 100%
  );
  background-size: 100% 4px;
  z-index: 10;
  pointer-events: none;
}
```

### Pattern 1: Static Scanlines with Fixed Pixel Height

**What:** Repeating gradient with fixed background-size creates consistent scanline thickness across all screen sizes.

**When to use:** Default approach - simple, performant, authentic to CRT aesthetic.

**Example:**
```css
/* Source: https://aleclownes.com/2017/02/01/crt-display.html */
/* Adapted for readability focus */
.crt-overlay::before {
  content: " ";
  display: block;
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  background: linear-gradient(
    to bottom,
    rgba(18, 16, 16, 0) 50%,
    rgba(0, 0, 0, 0.25) 50%
  );
  background-size: 100% 4px; /* 2px transparent, 2px scanline */
  z-index: 2;
  pointer-events: none;
}
```

**Key properties:**
- `background-size: 100% 4px` - Creates 4px repeat pattern (half transparent, half scanline)
- `pointer-events: none` - Overlay doesn't block interactions
- `z-index: 2` or higher - Ensures overlay appears above content
- Opacity 0.2-0.25 range balances visibility with readability

### Pattern 2: Responsive Scanlines for Retina Displays

**What:** Use media queries to adjust scanline thickness on high-DPI displays.

**When to use:** When targeting retina displays where 2px scanlines might appear too thin.

**Example:**
```css
.crt-overlay::before {
  background-size: 100% 2px; /* Base size for standard displays */
}

@media (-webkit-min-device-pixel-ratio: 2), (min-resolution: 192dpi) {
  .crt-overlay::before {
    background-size: 100% 4px; /* Thicker scanlines on retina */
  }
}
```

### Pattern 3: Accessibility-Aware Implementation

**What:** Disable or reduce scanline effect for users who prefer reduced motion.

**When to use:** Always - respects WCAG 2.1 accessibility guidelines.

**Example:**
```css
/* Source: https://developer.mozilla.org/en-US/docs/Web/CSS/@media/prefers-reduced-motion */
.crt-overlay::before {
  background: repeating-linear-gradient(...);
  background-size: 100% 4px;
}

@media (prefers-reduced-motion: reduce) {
  .crt-overlay::before {
    opacity: 0.5; /* Reduce intensity */
    /* OR: display: none; to remove completely */
  }
}
```

### Pattern 4: Mobile-Optimized with GPU Acceleration

**What:** Explicitly trigger GPU acceleration for smoother rendering on mobile.

**When to use:** If performance testing reveals jank on lower-end mobile devices.

**Example:**
```css
/* Source: https://www.lexo.ch/blog/2025/01/boost-css-performance-with-will-change-and-transform-translate3d-why-gpu-acceleration-matters/ */
.crt-overlay::before {
  /* ... scanline styles ... */
  transform: translateZ(0); /* Trigger GPU layer */
  backface-visibility: hidden; /* Reduce rendering artifacts */
}

/* Only use will-change if animating */
.crt-overlay.animated::before {
  will-change: background-position;
}
```

**Warning:** Don't use `will-change` for static scanlines - it wastes memory. Only apply when actually animating.

### Anti-Patterns to Avoid

- **Too many color stops:** Gradients with 5+ color stops can degrade mobile performance. Stick to 2-3 stops maximum.
- **Animating gradient colors:** Animating color values forces recalculation every frame. Animate `background-position` instead.
- **Overly small repeat patterns:** Patterns smaller than 2px can cause rendering issues on some devices.
- **High opacity:** Opacity above 0.3 significantly reduces text readability, especially with pixel fonts.
- **Missing pointer-events: none:** Overlay will block click/touch interactions with underlying content.
- **Applying will-change everywhere:** Overuse increases memory consumption and can degrade performance.

## Don't Hand-Roll

Problems that look simple but have existing solutions:

| Problem | Don't Build | Use Instead | Why |
|---------|-------------|-------------|-----|
| Animated scanlines | Custom JavaScript animation loop | CSS `@keyframes` with `background-position` | CSS animations are GPU-accelerated, JavaScript animations can cause jank and battery drain |
| Device pixel ratio detection | JavaScript window.devicePixelRatio checks | CSS `@media (-webkit-min-device-pixel-ratio: 2)` or `@media (min-resolution: 192dpi)` | Native media queries are more reliable and don't require JavaScript |
| Motion preference detection | Custom localStorage preference toggle | CSS `@media (prefers-reduced-motion)` | Respects OS-level accessibility settings, WCAG compliant |
| Canvas-based scanlines | Custom canvas rendering loop | CSS repeating gradients | CSS is declarative, automatically composited, performs better on mobile |

**Key insight:** CSS has evolved to handle these visual effects natively with better performance characteristics than custom JavaScript solutions. The browser's compositor can optimize CSS layers far more effectively than application code can optimize canvas rendering.

## Common Pitfalls

### Pitfall 1: Opacity Too High

**What goes wrong:** Scanlines with opacity above 0.3 severely reduce text readability, especially with small pixel fonts like Press Start 2P.

**Why it happens:** CRT emulator examples often use high opacity (0.5-0.8) for dramatic effect, but these target game graphics, not text-heavy interfaces.

**How to avoid:** Start with opacity 0.15-0.2 and test with actual content. For pixel fonts at small sizes, stay under 0.25.

**Warning signs:** Text appears muddy or hard to read, users squint or zoom in, accessibility tools report low contrast ratios.

### Pitfall 2: Scanlines Block Interactions

**What goes wrong:** Overlay blocks clicks, taps, and hover states on underlying interactive elements.

**Why it happens:** Forgetting to add `pointer-events: none` to the overlay pseudo-element.

**How to avoid:** Always include `pointer-events: none` on scanline overlays. Test touch interactions on actual mobile devices.

**Warning signs:** Buttons don't respond to clicks, hover effects don't trigger, mobile users report unresponsive interface.

### Pitfall 3: Performance Degradation on Mobile

**What goes wrong:** Smooth 60fps on desktop becomes janky 30fps on mobile, especially during scrolling or animations.

**Why it happens:** Complex gradients with too many color stops, or gradients not being composited to GPU layer.

**How to avoid:**
- Use simple 2-color gradients only
- Keep repeat pattern 2px or larger
- Avoid animating gradient colors
- Test on actual mobile devices (not just Chrome DevTools mobile emulation)

**Warning signs:** Dropped frames during scrolling, battery drain, device heating up, laggy touch response.

### Pitfall 4: Ignoring Accessibility Preferences

**What goes wrong:** Users with motion sensitivity or vestibular disorders experience discomfort or nausea from scanline flickering patterns.

**Why it happens:** Not implementing `prefers-reduced-motion` media query, or not understanding that static scanlines can still trigger motion sensitivity in some users.

**How to avoid:**
- Always include `@media (prefers-reduced-motion: reduce)` query
- Either reduce opacity significantly (50%) or remove effect entirely
- Test with motion preferences enabled in OS settings

**Warning signs:** Accessibility audits flag missing motion preferences, users with accessibility needs report discomfort.

### Pitfall 5: Z-Index Stacking Issues

**What goes wrong:** Scanlines appear under content instead of over it, or scanlines cover interactive elements that should be above them.

**Why it happens:** Unclear stacking context or insufficient z-index value.

**How to avoid:**
- Ensure parent element has stacking context (position: relative)
- Set scanline overlay z-index high enough to appear above scene content but below UI controls
- For this project: scene content (z-index: 1), scanlines (z-index: 10), speech text (z-index: 20), controls (natural stacking)

**Warning signs:** Scanlines not visible, or scanlines covering speech bubbles/controls.

### Pitfall 6: Fixed Pixel Sizes on Retina Displays

**What goes wrong:** Scanlines appear too thin or too thick on high-DPI displays.

**Why it happens:** Not accounting for device pixel ratio - CSS pixels are device-independent.

**How to avoid:** Test on retina devices (or use Chrome DevTools DPR emulation). Consider media queries for different pixel densities if needed, but start with 4px default which works well across most displays.

**Warning signs:** Barely visible scanlines on retina MacBooks/iPads, excessively thick scanlines on standard displays.

## Code Examples

Verified patterns from official sources:

### Basic Scanline Overlay

```css
/* Source: https://dev.to/ekeijl/retro-crt-terminal-screen-in-css-js-4afh */
/* Adapted for MonkeyLines context */
.monkey-scene::before {
  content: "";
  display: block;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(
    to bottom,
    transparent 0,
    transparent 50%,
    rgba(0, 0, 0, 0.2) 50%,
    rgba(0, 0, 0, 0.2) 100%
  );
  background-size: 100% 4px;
  z-index: 10;
  pointer-events: none;
}
```

### Accessibility-Enhanced Scanlines

```css
/* Full implementation with accessibility considerations */
.monkey-scene::before {
  content: "";
  display: block;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: linear-gradient(
    to bottom,
    transparent 0,
    transparent 50%,
    rgba(0, 0, 0, 0.2) 50%,
    rgba(0, 0, 0, 0.2) 100%
  );
  background-size: 100% 4px;
  z-index: 10;
  pointer-events: none;
}

/* Respect user motion preferences */
@media (prefers-reduced-motion: reduce) {
  .monkey-scene::before {
    opacity: 0.5; /* Reduce intensity by 50% */
  }
}

/* Optional: Remove on very small screens if it impacts readability */
@media (max-width: 400px) {
  .monkey-scene::before {
    opacity: 0.15; /* Further reduce on small mobile screens */
  }
}
```

### Mobile-Optimized with GPU Acceleration

```css
/* Only add these properties if performance testing reveals issues */
.monkey-scene::before {
  content: "";
  /* ... other scanline properties ... */
  transform: translateZ(0);
  backface-visibility: hidden;
  /* DO NOT use will-change for static scanlines */
}
```

### Alternative: Two-Stage Gradient Pattern

```css
/* Source: https://aleclownes.com/2017/02/01/crt-display.html */
/* More authentic CRT appearance with color channel separation */
.monkey-scene::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background:
    linear-gradient(
      rgba(18, 16, 16, 0) 50%,
      rgba(0, 0, 0, 0.25) 50%
    ),
    linear-gradient(
      90deg,
      rgba(255, 0, 0, 0.06),
      rgba(0, 255, 0, 0.02),
      rgba(0, 0, 255, 0.06)
    );
  background-size: 100% 2px, 3px 100%;
  z-index: 10;
  pointer-events: none;
}
```

**Warning:** The dual-gradient approach (horizontal scanlines + vertical color separation) may impact mobile performance. Start with single gradient and add color separation only if needed and performance testing passes.

## State of the Art

| Old Approach | Current Approach | When Changed | Impact |
|--------------|------------------|--------------|--------|
| Canvas-based CRT filters | CSS repeating gradients | ~2017-2019 | Significantly better mobile performance, declarative, easier to maintain |
| JavaScript animation loops | CSS @keyframes + background-position | ~2016-2018 | GPU acceleration, 60fps on mobile, lower battery consumption |
| Fixed pixel values everywhere | CSS pixels + media queries for DPR | ~2018-2020 | Proper rendering across retina and standard displays |
| Ignore accessibility | prefers-reduced-motion support | 2019-2021 (WCAG 2.1) | WCAG compliance, better user experience for motion-sensitive users |
| will-change on everything | Selective will-change only for animations | ~2020-2022 | Reduced memory consumption, better performance on mobile |

**Deprecated/outdated:**
- **Canvas overlays for simple effects:** CSS compositing is now faster and more efficient for static overlays
- **JavaScript-based scanline rendering:** CSS repeating gradients handle this natively with better performance
- **Vendor prefixes for gradients:** No longer needed for `linear-gradient()` or `repeating-linear-gradient()` as of 2020+
- **opacity-based animation for flicker:** If animating, use transform/position instead for GPU acceleration

## Open Questions

Things that couldn't be fully resolved:

1. **Optimal scanline thickness for Press Start 2P font at various sizes**
   - What we know: 2px-4px is the standard range, 4px works well for most displays
   - What's unclear: Precise optimal value for Press Start 2P at the specific font sizes used in MonkeyLines (8px-12px range)
   - Recommendation: Start with 4px, test with actual content, adjust if readability suffers. Create ticket to gather user feedback if deployed.

2. **Performance impact of dual-gradient approach (scanlines + color separation)**
   - What we know: Multiple gradient layers increase rendering complexity, especially on mobile
   - What's unclear: Specific performance impact on target devices (need real device testing)
   - Recommendation: Start with single gradient scanlines only. Add color separation in future iteration after performance validation.

3. **User preference for effect intensity**
   - What we know: Community recommendations range from 10% to 100% opacity, suggesting highly subjective preference
   - What's unclear: Whether MonkeyLines users will want adjustable intensity or disable option
   - Recommendation: Start with 0.2 opacity (subtle). Consider adding user preference toggle in future iteration based on feedback.

4. **Interaction with existing border-image pixelated border**
   - What we know: `.monkey-scene` uses custom border-image with pixelated rendering
   - What's unclear: Whether scanline overlay will create visual conflicts or artifacts at border edges
   - Recommendation: Test implementation carefully at border intersections. May need to adjust overlay positioning or add edge handling.

## Sources

### Primary (HIGH confidence)
- [Alec Lownes - Using CSS to create a CRT](https://aleclownes.com/2017/02/01/crt-display.html) - Core CSS scanline techniques
- [DEV Community - Retro CRT terminal screen in CSS + JS](https://dev.to/ekeijl/retro-crt-terminal-screen-in-css-js-4afh) - Implementation patterns
- [MDN - prefers-reduced-motion](https://developer.mozilla.org/en-US/docs/Web/CSS/@media/prefers-reduced-motion) - Accessibility requirements
- [MDN - repeating-linear-gradient()](https://developer.mozilla.org/en-US/docs/Web/CSS/gradient/repeating-linear-gradient) - Technical documentation
- [Lexo.ch - CSS GPU Acceleration Guide](https://www.lexo.ch/blog/2025/01/boost-css-performance-with-will-change-and-transform-translate3d-why-gpu-acceleration-matters/) - Performance optimization

### Secondary (MEDIUM confidence)
- [Hoverify - CSS gradient performance](https://tryhoverify.com/blog/i-wish-i-had-known-this-sooner-about-css-gradient-performance/) - Mobile optimization recommendations
- [Elementor - CSS Gradients 2026 Guide](https://elementor.com/blog/css-gradients/) - Current best practices
- [CSS-Tricks - prefers-reduced-motion](https://css-tricks.com/almanac/rules/m/media/prefers-reduced-motion/) - Accessibility implementation
- [Smashing Magazine - GPU Animation: Doing It Right](https://www.smashingmagazine.com/2016/12/gpu-animation-doing-it-right/) - Performance best practices
- [MDN - pointer-events](https://developer.mozilla.org/en-US/docs/Web/CSS/pointer-events) - Overlay interaction handling

### Tertiary (LOW confidence - community discussions)
- [Petrockblock Forums - CRT scanline recommendations](https://www.petrockblock.com/forums/topic/list-of-recommended-shaders-for-raspberry-piretropie-how-to-get-the-crt-look/) - Community opacity preferences
- CodePen examples (various authors) - Implementation inspiration
- WebSearch results on retina display handling - Cross-referenced with MDN documentation

## Metadata

**Confidence breakdown:**
- Standard stack: HIGH - CSS gradients are the established, documented approach with wide adoption
- Architecture: HIGH - Pseudo-element pattern is well-documented with official examples and performance data
- Pitfalls: HIGH - Common issues verified across multiple sources and official documentation
- Opacity values: MEDIUM - Community recommendations vary widely (10%-100%), requires project-specific testing
- Mobile performance: HIGH - Official documentation and performance guides provide clear optimization strategies

**Research date:** 2026-02-09
**Valid until:** ~2026-05-09 (90 days - CSS features are stable, but mobile browser optimizations evolve)

**Notes:**
- No CONTEXT.md found for this phase - full research scope conducted
- Project uses vanilla HTML/CSS/JS (no framework constraints)
- Existing CSS custom properties pattern (--mi-*) should be followed
- Self-hosted font and Go binary embedding favor inline/base64 approaches
- Mobile performance and text readability are explicit project priorities

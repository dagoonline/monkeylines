# Feature Landscape

**Domain:** Retro adventure game UI (LucasArts/Monkey Island style) for web
**Researched:** 2026-02-07
**Confidence:** MEDIUM (based on established domain knowledge; web search unavailable for verification of latest CSS techniques)

## Table Stakes

Features users expect. Missing any of these and the page feels like a generic web app, not a game UI.

| Feature | Why Expected | Complexity | Mobile Notes |
|---------|--------------|------------|--------------|
| **Pixel art / bitmap font** | Monospace alone is not enough. The Monkey Island aesthetic is defined by its bitmap typography (chunky, anti-alias-free). A pixel font like "Press Start 2P" or similar immediately signals "retro game." | Low (Google Fonts or self-hosted WOFF2) | Works fine; use `clamp()` for sizing, minimum 14px for readability |
| **`image-rendering: pixelated`** on all art | Already present. Crucial -- blurry upscaled pixel art breaks the illusion instantly. | Already done | N/A |
| **Dark background with visible border/frame** | Monkey Island scenes are always framed. The game viewport has a clear boundary separating it from the verb/inventory UI. A CSS border or pixel-art frame around the scene gives it "screen within a screen" feel. | Low (CSS border or border-image) | Ensure border doesn't eat too much viewport on small screens |
| **Labeled or icon+text buttons** | Current emoji-only buttons (sword, loop, speaker) require guessing. Table stakes for any UI: controls must be self-explanatory. Pixel art icons with short text labels ("Fight!", "Auto", "Mute") is the pattern MI uses for verbs. | Low-Med (need small pixel icon sprites or pure CSS pixel icons + text) | Touch targets minimum 44x44px; text labels help on mobile where hover tooltips unavailable |
| **Speech text styling** | MI uses colored text floating above characters against the scene background (no bubble). Each character has a distinct text color. The current implementation positions text well but lacks color distinction between insult vs comeback speakers. | Low (CSS color per speaker) | Already responsive via `clamp()` |
| **Clear "what is this" communication** | First-time visitors need to understand in under 3 seconds: this is a Monkey Island insult exchange toy. A title/tagline visible on load is table stakes. Currently the page has no visible title or explanation. | Low (HTML + styled text) | Keep brief; one line max |
| **Visual feedback on interaction** | Button press should feel tactile. Current `:active` translateY is good. Add a brief color flash or border highlight on click to match the snappy feel of SCUMM verb clicks. | Low (CSS transition) | Already works via touch |

## Differentiators

Features that elevate from "themed web page" to "feels like playing a game." Not expected, but create delight.

| Feature | Value Proposition | Complexity | Mobile Notes |
|---------|-------------------|------------|--------------|
| **Pixel art button sprites** | Replace emoji with actual pixel art buttons drawn in MI palette. The MI verb bar uses solid-colored text on dark panels. Custom pixel buttons (a tiny crossed-swords icon, a loop arrow, a speaker) in the MI color palette (teal, amber, cream on dark purple/brown) would be a huge authenticity boost. | Medium (need to create/source small sprite assets or use CSS pixel art via box-shadow technique) | SVG or PNG sprites scale fine; CSS pixel art via box-shadow is resolution-independent |
| **MI color palette** | The original MI uses a very specific palette: deep purples (#1a0a2e), teals (#2a8a6a), warm ambers (#c8a848), cream text (#e8d8a0) on dark backgrounds. Applying this palette consistently across all UI elements ties everything together. | Low (CSS custom properties) | No impact |
| **Scanline or CRT overlay** | Subtle scanline effect (CSS repeating-linear-gradient or pseudo-element) over the scene gives a "playing on an old monitor" vibe. Must be subtle -- too strong and it hurts readability. | Low (CSS only, `pointer-events: none` overlay) | Reduce or disable on small screens where scanlines compress into visual noise |
| **Themed scroll/panel borders** | Instead of plain CSS borders, use `border-image` with a pixel art frame (ornate wood or stone border like MI inventory panel). This single asset change dramatically elevates the feel. | Medium (need border tile asset, 9-slice via `border-image`) | Works natively with `border-image` |
| **Head animation on mouth open/close** | Currently swapping overlay images per speaking monkey. Adding a simple 2-frame mouth animation (open/closed toggle synced to word timing) would add life. Already partially there with word-based overlay toggling. | Low-Med (need alternate sprite frames, JS timing already exists) | No impact |
| **Typewriter sound effect** | A subtle key-click or character-reveal sound synced to text typing. MI has this -- each character appearing makes a tiny blip. The hoot system is charming but a text-reveal sound layered underneath would add polish. | Low (single short audio sample, play per character or per word) | Respect mute state; use existing AudioContext |
| **Verb bar layout** | Arrange buttons in a horizontal bar below the scene mimicking the SCUMM verb panel layout: dark background strip, text labels in columns. Even without full SCUMM verbs, styling the button row as a "panel" with its own background color and border creates the classic two-zone layout (scene above, controls below). | Low (CSS background + padding on btn-row) | Already a horizontal row; just needs panel styling |
| **Idle animation** | When no exchange is happening, subtle idle movement (monkey heads slight bob, or blinking) keeps the scene alive. MI characters always have idle animations. | Medium (additional sprite frames + JS interval) | Battery consideration; use `requestAnimationFrame` with visibility check |
| **Startup text crawl or intro** | A brief "Deep in the Caribbean..." style intro text on first load, then fade to the scene. Sets the mood immediately. | Low-Med (CSS animation sequence) | Keep short, allow skip-tap |

## Anti-Features

Features to explicitly NOT build. Common mistakes when creating retro game UIs for web.

| Anti-Feature | Why Avoid | What to Do Instead |
|--------------|-----------|-------------------|
| **Full SCUMM verb grid** | Tempting to recreate the full "Open / Close / Push / Pull / Give / Use" verb grid. But this app has 3 actions, not 12. A fake verb grid with mostly-dead buttons feels hollow and confusing. | Style the 3 real buttons to evoke the verb bar aesthetic without faking unused verbs |
| **Inventory panel** | MI has an inventory strip. Adding a decorative one with no real items is pure visual noise that confuses users ("what do I do with these?"). | Skip entirely. The scene + controls two-zone layout is enough |
| **Mouse cursor replacement** | Custom pixel cursors are fragile (lag, wrong hotspot, invisible on some browsers, terrible on mobile/touch). High effort, low payoff, actively harmful on touch devices. | Keep native cursor. The pixel art environment provides enough aesthetic |
| **Background music autoplay** | Browsers block it, users hate it, and a looping MIDI-style track gets grating fast on a single-screen toy. The monkey hoots are the right level of audio. | Keep current hoot system. Optionally add subtle SFX (button click, text blip) but no music |
| **Pixel art dithering on photos/gradients** | Some retro UIs try to dither smooth gradients to look 8-bit. This is expensive (canvas manipulation or massive pre-rendered images) and looks muddy on mobile. | Use flat colors from the MI palette. The art assets are already properly pixel-art styled |
| **Overly complex animations** | Parallax scrolling, particle effects, screen transitions with wipes -- these scream "modern dev cosplaying retro" rather than authentic retro. MI animations are simple: 2-4 frame loops, hard cuts between scenes. | Keep animations to 2-3 frame sprite swaps and simple CSS transitions |
| **Tooltip-dependent UI** | Hover tooltips don't exist on mobile. If buttons only make sense with tooltips, the UI is broken for half your users. | Use visible text labels on buttons. If space is tight, use icon+label combo |
| **Responsive breakpoint overhaul** | Don't build separate mobile and desktop layouts. MI had one screen layout. The current single-column responsive approach is correct. | Keep single layout, use `clamp()` for fluid sizing, test on 320px-wide screens |

## Feature Dependencies

```
Pixel font ──> All text styling (speech, buttons, title)
MI color palette ──> Button styling, border styling, speech colors
Button redesign ──> Depends on: pixel font + MI palette + icon assets
Panel/frame border ──> Depends on: MI palette (+ optional border-image asset)
Verb bar layout ──> Depends on: button redesign
CRT overlay ──> Independent (CSS-only, add anytime)
Idle animation ──> Depends on: additional sprite assets
Typewriter SFX ──> Independent (add to existing typing logic)
Title/tagline ──> Depends on: pixel font + MI palette
```

## MVP Recommendation

For the UI redesign MVP, prioritize in this order:

1. **Pixel font + MI color palette** -- Foundation that everything else builds on. Biggest bang for lowest effort.
2. **Title/tagline** -- "Insult Sword Fighting" or similar, one line, immediately communicates purpose.
3. **Labeled buttons with verb-bar panel styling** -- Replace emoji with text labels on a dark panel strip. Even without pixel art icons, text labels in pixel font on a styled panel bar will feel authentic.
4. **Speaker-colored speech text** -- Different colors for insult vs comeback (e.g., amber for side monkey, teal for center monkey).
5. **Scene frame/border** -- CSS border in MI palette colors or simple pixel border-image.

Defer to post-MVP:
- **CRT scanline overlay**: Pure polish, easy to add later, zero dependencies
- **Pixel art icon sprites for buttons**: Text labels work fine initially; icons are a nice-to-have enhancement
- **Idle animations**: Requires new art assets; current static scene is fine
- **Typewriter SFX**: Polish layer; hoots already provide audio character
- **Startup intro text**: Nice but not needed for comprehension if title exists

## Sources

- Domain knowledge of LucasArts SCUMM engine UI patterns (Monkey Island 1 & 2, Day of the Tentacle)
- Direct analysis of current MonkeyLines codebase (`index.html`)
- Confidence: MEDIUM -- recommendations are based on well-established game UI patterns but web search was unavailable to verify latest CSS techniques or community examples

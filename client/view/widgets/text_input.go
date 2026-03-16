package widgets

import (
	"ecs-test/assets/sprites/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"github.com/yohamta/ganim8/v2"
	"image"
	"strings"
)

type TextInput struct {
	Text                    string
	borderSprite            *ganim8.Sprite
	renderer                *etxt.Renderer
	maxLength               int
	scale                   float64
	marginRight, marginLeft int

	keyRepeat  int64
	cursorPos  int
	startIndex int

	hasFocus   bool
	focusGroup *FocusGroup
	masked     bool
}

func NewTextInput(r *etxt.Renderer, maxLength int, scale float64) *TextInput {
	return &TextInput{
		renderer:     r,
		maxLength:    maxLength,
		scale:        scale,
		borderSprite: ui.TextInputSprite,
	}
}

func (t *TextInput) Blur() {
	t.hasFocus = false
}

func (t *TextInput) Focus() {
	t.hasFocus = true
}

func (t *TextInput) WithRightMargin(marginRight int) *TextInput {
	t.marginRight = marginRight
	return t
}

func (t *TextInput) WithLeftMargin(marginLeft int) *TextInput {
	t.marginLeft = marginLeft
	return t
}
func (t *TextInput) Masked() *TextInput {
	t.masked = true
	return t
}

func (t *TextInput) HandleJustPressedMouseButtonLeft(x, y int) bool {
	t.focusGroup.Focused(t)
	t.hasFocus = true
	return false
}

func (t *TextInput) HandleJustReleasedMouseButtonLeft(x, y int) {
}

func (t *TextInput) Update(view *furex.View) {
	sprite := ui.TextInputSprite
	scaledSpriteHeight := float64(sprite.Height()) * t.scale
	view.SetHeight(int(scaledSpriteHeight))

	if !t.hasFocus {
		return
	}

	if len(t.Text) < t.maxLength {
		var input []rune
		input = ebiten.AppendInputChars(input)
		for _, char := range input {
			if char >= 32 && char <= 126 {
				t.Text = t.Text[:t.cursorPos] + string(char) + t.Text[t.cursorPos:] // Insert char at cursor position
				t.cursorPos++
			}
		}

		visibleLen := t.visibleLength(view.Width - t.marginLeft - t.marginRight)
		if t.cursorPos > t.startIndex+visibleLen {
			t.startIndex++
		}
	}

	// Handle backspace key
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) && t.cursorPos > 0 {
		if t.keyRepeat%5 == 0 {
			t.Text = t.Text[:t.cursorPos-1] + t.Text[t.cursorPos:] // Remove char at cursor position
			t.cursorPos--

			// Ensure startIndex stays consistent with cursorPos
			if t.cursorPos < t.startIndex {
				t.startIndex = t.cursorPos
			}
		}
		t.keyRepeat++
	}

	// Handle left arrow key
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && t.cursorPos > 0 {
		if t.keyRepeat%5 == 0 {
			t.cursorPos--
			// Adjust startIndex to ensure the cursor is visible
			if t.cursorPos < t.startIndex {
				t.startIndex = t.cursorPos
			}
		}
		t.keyRepeat++
	}

	// Handle right arrow key
	if ebiten.IsKeyPressed(ebiten.KeyRight) && t.cursorPos < len(t.Text) {
		if t.keyRepeat%5 == 0 {
			t.cursorPos++
			visibleLen := t.visibleLength(view.Width - t.marginLeft - t.marginRight)
			// Ensure startIndex is updated to keep the cursor visible
			if t.cursorPos > t.startIndex+visibleLen {
				t.startIndex++
			}
		}
		t.keyRepeat++
	}

	// Ensure cursor stays in bounds
	if t.cursorPos < t.startIndex {
		t.cursorPos = t.startIndex
	} else if t.cursorPos > len(t.Text) {
		t.cursorPos = len(t.Text)
	}

	//// Reset key repeat counter if no key is pressed
	//if !ebiten.IsKeyPressed(ebiten.KeyBackspace) && !ebiten.IsKeyPressed(ebiten.KeyLeft) && !ebiten.IsKeyPressed(ebiten.KeyRight) {
	//	t.keyRepeat = 0
	//}
}

func (t *TextInput) visibleLength(availableWidth int) int {
	textWidth := 0.0
	visibleChars := 0

	// Use the masked version for width calculation if masked
	for i := t.startIndex; i < len(t.Text); i++ {
		dimensions := t.renderer.Measure(t.maskText(string(t.Text[i])))

		// Check if adding this character would exceed the available width
		if textWidth+dimensions.Width().ToFloat64() > float64(availableWidth) {
			break
		}

		// Only add the character if it fits within available width
		textWidth += dimensions.Width().ToFloat64()
		visibleChars++
	}

	return visibleChars
}

func (t *TextInput) Draw(screen *ebiten.Image, frame image.Rectangle, view *furex.View) {
	t.drawFrame(screen, frame)
	t.drawText(screen, frame, frame.Dx())
}

func (t *TextInput) drawFrame(
	screen *ebiten.Image,
	frame image.Rectangle,
) {
	spriteWidth := float64(t.borderSprite.Width())
	scaledSpriteWidth := spriteWidth * t.scale

	// Position to draw the first frame
	startX := float64(frame.Min.X)
	startY := float64(frame.Min.Y)

	// 1. Draw the first frame (index 0)
	firstFrameOpts := ganim8.DrawOpts(startX, startY, 0, t.scale, t.scale)
	t.borderSprite.Draw(screen, 0, firstFrameOpts)

	// Calculate the total available width for all frames excluding the last frame
	availableWidth := float64(frame.Dx()) - scaledSpriteWidth

	// 2. Calculate position for the middle frame and draw fillers on both sides
	middleFrameX := startX + (availableWidth)/2 // Center the middle frame in the available space

	// Draw left filler (index 1) from startX + scaledSpriteWidth to middleFrameX
	fillerX := startX + scaledSpriteWidth
	for fillerX < middleFrameX {
		fillerOpts := ganim8.DrawOpts(fillerX, startY, 0, t.scale, t.scale)
		t.borderSprite.Draw(screen, 1, fillerOpts)
		fillerX += scaledSpriteWidth
	}

	// 3. Draw the middle frame (index 2)
	middleFrameOpts := ganim8.DrawOpts(middleFrameX, startY, 0, t.scale, t.scale)
	t.borderSprite.Draw(screen, 2, middleFrameOpts)

	// Draw right filler (index 1) from middleFrameX + scaledSpriteWidth to availableWidth
	fillerX = middleFrameX + scaledSpriteWidth
	for fillerX < startX+availableWidth {
		fillerOpts := ganim8.DrawOpts(fillerX, startY, 0, t.scale, t.scale)
		t.borderSprite.Draw(screen, 1, fillerOpts)
		fillerX += scaledSpriteWidth
	}

	// 4. Draw the last frame (index 4)
	lastFrameX := startX + availableWidth // Position of the last frame at the end of available width
	lastFrameOpts := ganim8.DrawOpts(lastFrameX, startY)
	lastFrameOpts.SetScale(t.scale, t.scale)
	t.borderSprite.Draw(screen, 4, lastFrameOpts)
}

func (t *TextInput) drawText(
	screen *ebiten.Image,
	frame image.Rectangle,
	availableWidth int,
) {
	scaledSpriteHeight := float64(t.borderSprite.Height()) * t.scale

	startX := float64(frame.Min.X)
	startY := float64(frame.Min.Y)

	adjustedWidth := availableWidth - (t.marginLeft + t.marginRight)
	visibleLen := t.visibleLength(adjustedWidth)
	textToDisplay := t.Text[t.startIndex : t.startIndex+visibleLen]

	textDimensions := t.renderer.Measure(textToDisplay)
	textYOffset := int(startY + (scaledSpriteHeight / 2) + (textDimensions.Height().ToFloat64() / 4))

	textX := int(startX) + t.marginLeft
	t.renderer.Draw(screen, t.maskText(textToDisplay), textX, textYOffset)

	t.drawCursor(screen, textX, startY, scaledSpriteHeight, availableWidth)
}

func (t *TextInput) drawCursor(
	screen *ebiten.Image,
	textX int,
	startY float64,
	scaledSpriteHeight float64,
	availableWidth int,
) {
	if !t.hasFocus {
		return
	}

	cursorDimensions := t.renderer.Measure("|")

	cursorX := textX
	if len(t.Text) > 0 {
		for i := t.startIndex; i < t.cursorPos && i < len(t.Text); i++ {
			dimensions := t.renderer.Measure(string(t.Text[i]))
			if t.masked {
				dimensions = t.renderer.Measure("*")
			}
			cursorX += int(dimensions.Width().ToFloat64())

			// TODO point precision error?
			if cursorX > textX+availableWidth-t.marginRight {
				cursorX = textX + availableWidth - t.marginRight // Clamp the cursor to the right margin
				break
			}
		}
	}

	cursorHeight := cursorDimensions.Height().ToFloat64()
	cursorYOffset := int(startY + (scaledSpriteHeight / 2) + (cursorHeight / 4))

	t.renderer.Draw(screen, "|", cursorX, cursorYOffset)
}

func (t *TextInput) UsingFocusGroup(group *FocusGroup) *TextInput {
	t.focusGroup = group
	group.AddItem(t)
	return t
}

func (t *TextInput) Clear() {
	t.Text = ""
	t.cursorPos = 0
	t.startIndex = 0
}

func (t *TextInput) maskText(text string) string {
	if !t.masked {
		return text
	}
	return strings.Repeat("*", len(text))
}

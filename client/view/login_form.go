package view

import (
	"ecs-test/assets/fonts"
	"ecs-test/client/view/widgets"
	"github.com/tinne26/etxt"
	"github.com/yohamta/furex/v2"
	"image/color"
	"math"
	"time"
)

type LoginForm struct {
	scale     float64
	renderer  *etxt.Renderer
	loginFunc func(email, password string)

	rootView        *furex.View
	progressView    *furex.View
	loginButtonView *ButtonView

	focusGroup    *widgets.FocusGroup
	emailInput    *widgets.TextInput
	passwordInput *widgets.TextInput

	progressBar    *widgets.ProgressBar
	shaking        bool
	shakeStart     time.Time
	shakeDuration  time.Duration
	shakeIntensity float64
}

func NewLoginForm(scale float64, loginFunc func(email, password string)) *LoginForm {
	r := etxt.NewRenderer()
	r.SetFont(fonts.MainFont)
	r.SetSize(32)
	r.SetColor(color.RGBA{
		R: 217,
		G: 179,
		B: 132,
		A: 255,
	})

	form := &LoginForm{
		scale:     scale,
		renderer:  r,
		loginFunc: loginFunc,
	}
	return form.initView()
}

func (f *LoginForm) initView() *LoginForm {
	f.rootView = &furex.View{
		MarginTop:    100,
		Direction:    furex.Column,
		Height:       250,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyStart,
		AlignItems:   furex.AlignItemEnd,
		Handler:      widgets.NewPanel(f.scale, 400),
	}

	f.initHeader()
	f.initFocusGroup()
	f.initEmailField()
	f.initPasswordField()
	f.initButtons()
	f.initProgressView()
	return f
}

func (f *LoginForm) initHeader() {
	topRow := &furex.View{
		Direction:    furex.Row,
		AlignContent: furex.AlignContentStart,
		Justify:      furex.JustifyStart,
		AlignItems:   furex.AlignItemStart,
		WidthInPct:   100,
	}

	header := &furex.View{
		Handler:      widgets.NewHeader(f.scale),
		MarginLeft:   8,
		MarginRight:  8,
		MarginTop:    8,
		WidthInPct:   100,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemCenter,
	}
	header.AddChild(&furex.View{
		Handler: widgets.
			NewText("Login", f.renderer, f.scale).
			WithVerticalShift(-3),
	})

	topRow.AddChild(header)
	f.rootView.AddChild(topRow)
}

func (f *LoginForm) initFocusGroup() {
	f.focusGroup = widgets.NewFocusGroup()
	f.rootView.AddChild(&furex.View{
		Handler: f.focusGroup,
		Display: furex.DisplayNone,
	})
}

func (f *LoginForm) initEmailField() {
	f.emailInput = widgets.NewTextInput(f.renderer, 40, f.scale).
		WithLeftMargin(10).
		WithRightMargin(10).
		UsingFocusGroup(f.focusGroup)
	f.createFormEntryRow(f.emailInput, "Email")
}

func (f *LoginForm) initPasswordField() {
	f.passwordInput = widgets.NewTextInput(f.renderer, 40, f.scale).
		WithLeftMargin(10).
		WithRightMargin(10).
		UsingFocusGroup(f.focusGroup).
		Masked()
	f.createFormEntryRow(f.passwordInput, "Password")
}

func (f *LoginForm) initButtons() {
	buttonsView := &furex.View{
		Direction:    furex.Row,
		WidthInPct:   100,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemCenter,
		MarginTop:    32,
	}

	f.loginButtonView = NewButtonView(widgets.ButtonVariantGreen, "Login", 2.0).
		OnClick(f.login)

	buttonsView.AddChild(f.loginButtonView.View())
	f.rootView.AddChild(buttonsView)
}

func (f *LoginForm) initProgressView() {
	f.progressView = &furex.View{
		Direction:    furex.Row,
		WidthInPct:   100,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemCenter,
		MarginTop:    32,
		Display:      furex.DisplayNone,
	}

	f.progressBar = widgets.NewProgressBar(0)
	f.progressView.AddChild(f.progressBar.View())

	f.rootView.AddChild(f.progressView)
}

func (f *LoginForm) createFormEntryRow(
	input *widgets.TextInput,
	labelText string,
) {
	row := &furex.View{
		Direction:    furex.Row,
		MarginTop:    20,
		MarginRight:  50,
		AlignContent: furex.AlignContentCenter,
		Justify:      furex.JustifyCenter,
		AlignItems:   furex.AlignItemCenter,
	}

	label := &furex.View{
		Handler: widgets.
			NewText(labelText, f.renderer, f.scale).
			WithColor(color.RGBA{
				R: 48,
				G: 36,
				B: 46,
				A: 255}),
		MarginRight: 10,
	}
	entry := &furex.View{
		Width:   200,
		Handler: input,
	}

	row.AddChild(label)
	row.AddChild(entry)

	f.rootView.AddChild(row)
}

func (f *LoginForm) View() *furex.View {
	return f.rootView
}

func (f *LoginForm) login() {
	if f.loginFunc != nil {
		f.loginFunc(f.emailInput.Text, f.passwordInput.Text)
	}
}

func (f *LoginForm) UpdateProgress(percentage int) {
	f.progressBar.UpdateProgress(percentage)
}

func (f *LoginForm) ShowProgress() {
	f.progressView.Display = furex.DisplayFlex
	f.loginButtonView.Disable()
}

func (f *LoginForm) HideProgress() {
	f.progressView.Display = furex.DisplayNone
	f.loginButtonView.Enable()
}

func (f *LoginForm) Shake(duration time.Duration, intensity float64) {
	f.shaking = true
	f.shakeStart = time.Now()
	f.shakeDuration = duration
	f.shakeIntensity = intensity
}

func (f *LoginForm) UpdateShake() {
	if f.shaking {
		elapsed := time.Since(f.shakeStart)
		if elapsed > f.shakeDuration {
			f.shaking = false
			f.rootView.MarginLeft = 0 // Reset position
		} else {
			// Calculate offset using a sinusoidal function for smooth shake
			shakeOffset := math.Sin(float64(elapsed.Milliseconds())*0.1) * f.shakeIntensity * f.scale
			f.rootView.MarginLeft = int(shakeOffset)
		}
	}
}

func (f *LoginForm) Reset() {
	f.passwordInput.Clear()
	f.emailInput.Clear()
	f.focusGroup.Reset()
}

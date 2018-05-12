package glfw

import (
	"events"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const DefaultFrameRate = 12
const DefaultHeight = 600
const DefaultTitle = "Default Title"
const DefaultWidth = 800

const ResizedEvent = "GlfwWindowResized"

type Option func(win *window)

type WindowHint struct {
	Key   glfw.Hint
	Value int
}

type window struct {
	events.EmitterBase

	width        float64
	height       float64
	hints        []WindowHint
	frameRate    int
	nativeWindow *glfw.Window
	pixelRatio   float64
	title        string
}

func (win *window) OnResize(handler events.EventHandler) events.Unsubscriber {
	return win.On(ResizedEvent, handler)
}

func (win *window) AddHint(h WindowHint) {
	win.hints = append(win.hints, h)
}

func (win *window) BeginFrame() {
	// TODO(lbayes): Make receiver for BgColor on Window
	gl.ClearColor(255, 255, 255, 255)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)

	//gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
	//gl.ClearColor(127, 127, 0, 255)
	// nativeWinWidth, nativeWinHeight := win.nativeWindow.GetSize()

	/*
		fbWidth, fbHeight := win.nativeWindow.GetFramebufferSize()

		w := float64(fbWidth)
		h := float64(fbHeight)

		if w != win.Width() {
			win.SetWidth(w)
		}
		if h != win.Height() {
			win.SetHeight(h)
		}
	*/
}

func (win *window) Close() {
	win.nativeWindow.Destroy()
	glfw.Terminate()
}

func (win *window) ShouldClose() bool {
	if win.nativeWindow != nil {
		return win.nativeWindow.ShouldClose()
	}
	return false
}

func (win *window) EndFrame() {
	gl.Enable(gl.DEPTH_TEST)
	win.nativeWindow.SwapBuffers()
}

func (win *window) PollEvents() {
	glfw.PollEvents()
}

func (win *window) FrameRate() int {
	return win.frameRate
}

func (win *window) Hints() []WindowHint {
	return win.hints
}

func (win *window) Hint(key glfw.Hint) int {
	for _, hint := range win.hints {
		if hint.Key == key {
			return hint.Value
		}
	}
	return -1
}

func (win *window) initPixelRatio() {
	w := win.nativeWindow
	fbWidth, _ := w.GetFramebufferSize()
	winWidth, _ := w.GetSize()

	win.pixelRatio = float64(fbWidth) / float64(winWidth)
}

func (win *window) initGlfw() {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	// Apply glfw window hints
	for _, hint := range win.hints {
		glfw.WindowHint(hint.Key, hint.Value)
	}

	nw, err := glfw.CreateWindow(int(win.Width()), int(win.Height()), win.Title(), nil, nil)
	win.nativeWindow = nw

	if err != nil {
		panic(err)
	}

	win.initPixelRatio()
	win.initResizeHandler()

	win.nativeWindow.MakeContextCurrent()
	// glfw.SwapInterval(0)
}

func (win *window) initResizeHandler() {
	w := win.nativeWindow
	w.SetFramebufferSizeCallback(func(w *glfw.Window, width int, height int) {
		win.resizeHandler(width, height)
	})
}

func (win *window) resizeHandler(width, height int) {
	win.SetWidth(float64(width))
	win.SetHeight(float64(height))
	win.Emit(events.New(ResizedEvent, win, nil))
}

func (win *window) initGl() {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	// Thought we needed this, but it looks like Nanovgo might be handling
	// this behind the scenes?
	// gl.Viewport(0, 0, int32(win.Width()), int32(win.Height()))
}

func (win *window) Init() {
	win.initGlfw()
	win.initGl()
}

func (win *window) PixelRatio() float64 {
	return win.pixelRatio
}

func (win *window) SetWidth(width float64) {
	win.width = width
}

func (win *window) SetHeight(height float64) {
	win.height = height
}

func (win *window) Width() float64 {
	return win.width
}

func (win *window) Height() float64 {
	return win.height
}

func (win *window) SetTitle(title string) {
	win.title = title
}

func (win *window) Title() string {
	return win.title
}

func New(options ...Option) *window {
	defaults := []Option{
		Width(DefaultWidth),
		Height(DefaultHeight),
		Title(DefaultTitle),
		FrameRate(DefaultFrameRate),
		Hint(glfw.Resizable, 1),
		Hint(glfw.Focused, 1),
		Hint(glfw.Visible, 1),
		// Hint(glfw.Floating, 1),
		// Hint(glfw.DoubleBuffer, 1),

		// For some reason, the following configuration breaks:
		//Hint(glfw.ContextVersionMajor, 4),
		//Hint(glfw.ContextVersionMinor, 1),
		// Hint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile),
		// Hint(glfw.OpenGLForwardCompatible, glfw.True),
	}

	win := &window{}
	// Merge and override defaults with provided options.
	options = append(defaults, options...)
	for _, option := range options {
		option(win)
	}
	return win
}

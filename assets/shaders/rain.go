//go:build ignore

//kage:unit pixels

package shaders

// Uniforms for the shader
var (
	Time               float // The time uniform for animating the rain
	ScreenSize         vec2  // The size of the screen to adapt the shader to different resolutions
	DropColor          vec4  // The color of the raindrops
	DropSpeed          float // Speed of the raindrops
	DropLength         float // Length of the raindrops
	DistortionStrength float // Strength of the distortion effect
)

func noise(uv vec2) float {
	// Simple noise function for randomness
	return fract(sin(dot(uv.xy, vec2(12.9898, 78.233))) * 43758.5453)
}

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	uv := texCoord * ScreenSize

	// Animate the rain based on time and speed
	uv.y += Time * DropSpeed

	// Calculate the raindrop effect
	drop := step(fract(uv.y), DropLength) // 1.0 if fract(uv.y) < DropLength, else 0.0
	drops := DropColor * drop

	// Randomize raindrop positions horizontally using noise
	drops.a *= step(noise(uv*10.0), 0.5) // Only display drops with certain randomness

	// Combine the original color with raindrops
	result := mix(color, drops, drops.a)

	// Add a slight distortion effect
	dist := sin(uv.y*50.0+Time*10.0) * DistortionStrength
	result.rgb += dist

	return result
}

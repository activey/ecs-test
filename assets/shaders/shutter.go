//go:build ignore

//kage:unit pixels

package shaders

var Time float
var ScreenSize vec2
var Mode float
var PixelationFactor float

// Shader code
var ShutterSpeed float

func Fragment(dstPos vec4, srcPos vec2, color vec4) vec4 {
	// Normalize screen coordinates (0.0 to 1.0)
	uv := dstPos.xy / ScreenSize

	// Define the center of the screen (0.5, 0.5 in normalized coordinates)
	center := vec2(0.5, 0.5)

	// Calculate aspect ratio correction
	aspectRatio := ScreenSize.x / ScreenSize.y

	// Pixelation: Quantize the UV coordinates for pixelation
	uv = floor(uv*PixelationFactor) / PixelationFactor

	// Adjust the Min-axis for aspect ratio to maintain a circular shape
	adjustedUV := vec2((uv.x-center.x)*aspectRatio+center.x, uv.y)

	// Calculate the distance from the current pixel to the center of the screen
	dist := distance(adjustedUV, center)

	// Radius of the circle, expanding or shrinking based on the mode and time
	var radius float
	if Mode == 0.0 {
		radius = Time
	} else if Mode == 1.0 {
		radius = 1.0 - Time
	}

	// If the distance is greater than the radius, draw black (outside the circle)
	if dist > radius {
		return vec4(0.0, 0.0, 0.0, 1.0) // Black pixel
	}

	// Otherwise, return the original image inside the circle
	return imageSrc0At(srcPos)
}

package media

////////////////////////////////////////////////////////////////////////////////
// TYPES

// SampleFormat specifies the type of a single sample
type SampleFormat uint

// ChannelLayout specifies the layout of channels
type ChannelLayout uint

// AudioFormat specifies the interface for audio format
type AudioFormat struct {
	// Sample rate in Hz
	Rate uint

	// Sample format
	Format SampleFormat

	// Channel layout
	Layout ChannelLayout
}

////////////////////////////////////////////////////////////////////////////////
// INTERFACES

// AudioSamples is a slice of audio samples
type AudioSamples interface {
	// Audio format
	Format() AudioFormat

	// Number of samples in a single channel
	Samples() int

	// Number of audio channels
	Channels() int

	// Linesize
	Linesize() int

	// Alignment
	Align() bool

	// Returns the samples for a specified channel, as array of bytes. For packed
	// audio format, the channel should be 0.
	Bytes(channel int) []byte
}

////////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	SAMPLE_FORMAT_NONE SampleFormat = iota
	SAMPLE_FORMAT_U8                // Byte
	SAMPLE_FORMAT_S16               // Signed 16-bit
	SAMPLE_FORMAT_S32               // Signed 32-bit
	SAMPLE_FORMAT_FLT               // Float 32-bit
	SAMPLE_FORMAT_DBL               // Float 64-bit
	SAMPLE_FORMAT_U8P               // Planar byte
	SAMPLE_FORMAT_S16P              // Planar signed 16-bit
	SAMPLE_FORMAT_S32P              // Planar signed 32-bit
	SAMPLE_FORMAT_FLTP              // Planar float 32-bit
	SAMPLE_FORMAT_DBLP              // Planar float 64-bit
	SAMPLE_FORMAT_S64               // Signed 64-bit
	SAMPLE_FORMAT_S64P              // Planar signed 64-bit
)

const (
	CHANNEL_LAYOUT_NONE ChannelLayout = iota
	CHANNEL_LAYOUT_MONO
	CHANNEL_LAYOUT_STEREO
	CHANNEL_LAYOUT_2POINT1
	CHANNEL_LAYOUT_2_1
	CHANNEL_LAYOUT_SURROUND
	CHANNEL_LAYOUT_3POINT1
	CHANNEL_LAYOUT_4POINT0
	CHANNEL_LAYOUT_4POINT1
	CHANNEL_LAYOUT_2_2
	CHANNEL_LAYOUT_QUAD
	CHANNEL_LAYOUT_5POINT0
	CHANNEL_LAYOUT_5POINT1
	CHANNEL_LAYOUT_5POINT0_BACK
	CHANNEL_LAYOUT_5POINT1_BACK
	CHANNEL_LAYOUT_6POINT0
	CHANNEL_LAYOUT_6POINT0_FRONT
	CHANNEL_LAYOUT_HEXAGONAL
	CHANNEL_LAYOUT_6POINT1
	CHANNEL_LAYOUT_6POINT1_BACK
	CHANNEL_LAYOUT_6POINT1_FRONT
	CHANNEL_LAYOUT_7POINT0
	CHANNEL_LAYOUT_7POINT0_FRONT
	CHANNEL_LAYOUT_7POINT1
	CHANNEL_LAYOUT_7POINT1_WIDE
	CHANNEL_LAYOUT_7POINT1_WIDE_BACK
	CHANNEL_LAYOUT_OCTAGONAL
	CHANNEL_LAYOUT_HEXADECAGONAL
	CHANNEL_LAYOUT_STEREO_DOWNMIX
	CHANNEL_LAYOUT_22POINT2
	CHANNEL_LAYOUT_AMBISONIC_FIRST_ORDER
)

////////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (v SampleFormat) String() string {
	switch v {
	case SAMPLE_FORMAT_NONE:
		return "SAMPLE_FORMAT_NONE"
	case SAMPLE_FORMAT_U8:
		return "SAMPLE_FORMAT_U8"
	case SAMPLE_FORMAT_S16:
		return "SAMPLE_FORMAT_S16"
	case SAMPLE_FORMAT_S32:
		return "SAMPLE_FORMAT_S32"
	case SAMPLE_FORMAT_FLT:
		return "SAMPLE_FORMAT_FLT"
	case SAMPLE_FORMAT_DBL:
		return "SAMPLE_FORMAT_DBL"
	case SAMPLE_FORMAT_U8P:
		return "SAMPLE_FORMAT_U8P"
	case SAMPLE_FORMAT_S16P:
		return "SAMPLE_FORMAT_S16P"
	case SAMPLE_FORMAT_S32P:
		return "SAMPLE_FORMAT_S32P"
	case SAMPLE_FORMAT_FLTP:
		return "SAMPLE_FORMAT_FLTP"
	case SAMPLE_FORMAT_DBLP:
		return "SAMPLE_FORMAT_DBLP"
	case SAMPLE_FORMAT_S64:
		return "SAMPLE_FORMAT_S64"
	case SAMPLE_FORMAT_S64P:
		return "SAMPLE_FORMAT_S64P"
	default:
		return "[?? Invalid SampleFormat value]"
	}
}

func (v ChannelLayout) String() string {
	switch v {
	case CHANNEL_LAYOUT_NONE:
		return "CHANNEL_LAYOUT_NONE"
	case CHANNEL_LAYOUT_MONO:
		return "CHANNEL_LAYOUT_MONO"
	case CHANNEL_LAYOUT_STEREO:
		return "CHANNEL_LAYOUT_STEREO"
	case CHANNEL_LAYOUT_2POINT1:
		return "CHANNEL_LAYOUT_2POINT1"
	case CHANNEL_LAYOUT_2_1:
		return "CHANNEL_LAYOUT_2_1"
	case CHANNEL_LAYOUT_SURROUND:
		return "CHANNEL_LAYOUT_SURROUND"
	case CHANNEL_LAYOUT_3POINT1:
		return "CHANNEL_LAYOUT_3POINT1"
	case CHANNEL_LAYOUT_4POINT0:
		return "CHANNEL_LAYOUT_4POINT0"
	case CHANNEL_LAYOUT_4POINT1:
		return "CHANNEL_LAYOUT_4POINT1"
	case CHANNEL_LAYOUT_2_2:
		return "CHANNEL_LAYOUT_2_2"
	case CHANNEL_LAYOUT_QUAD:
		return "CHANNEL_LAYOUT_QUAD"
	case CHANNEL_LAYOUT_5POINT0:
		return "CHANNEL_LAYOUT_5POINT0"
	case CHANNEL_LAYOUT_5POINT1:
		return "CHANNEL_LAYOUT_5POINT1"
	case CHANNEL_LAYOUT_5POINT0_BACK:
		return "CHANNEL_LAYOUT_5POINT0_BACK"
	case CHANNEL_LAYOUT_5POINT1_BACK:
		return "CHANNEL_LAYOUT_5POINT1_BACK"
	case CHANNEL_LAYOUT_6POINT0:
		return "CHANNEL_LAYOUT_6POINT0"
	case CHANNEL_LAYOUT_6POINT0_FRONT:
		return "CHANNEL_LAYOUT_6POINT0_FRONT"
	case CHANNEL_LAYOUT_HEXAGONAL:
		return "CHANNEL_LAYOUT_HEXAGONAL"
	case CHANNEL_LAYOUT_6POINT1:
		return "CHANNEL_LAYOUT_6POINT1"
	case CHANNEL_LAYOUT_6POINT1_BACK:
		return "CHANNEL_LAYOUT_6POINT1_BACK"
	case CHANNEL_LAYOUT_6POINT1_FRONT:
		return "CHANNEL_LAYOUT_6POINT1_FRONT"
	case CHANNEL_LAYOUT_7POINT0:
		return "CHANNEL_LAYOUT_7POINT0"
	case CHANNEL_LAYOUT_7POINT0_FRONT:
		return "CHANNEL_LAYOUT_7POINT0_FRONT"
	case CHANNEL_LAYOUT_7POINT1:
		return "CHANNEL_LAYOUT_7POINT1"
	case CHANNEL_LAYOUT_7POINT1_WIDE:
		return "CHANNEL_LAYOUT_7POINT1_WIDE"
	case CHANNEL_LAYOUT_7POINT1_WIDE_BACK:
		return "CHANNEL_LAYOUT_7POINT1_WIDE_BACK"
	case CHANNEL_LAYOUT_OCTAGONAL:
		return "CHANNEL_LAYOUT_OCTAGONAL"
	case CHANNEL_LAYOUT_HEXADECAGONAL:
		return "CHANNEL_LAYOUT_HEXADECAGONAL"
	case CHANNEL_LAYOUT_STEREO_DOWNMIX:
		return "CHANNEL_LAYOUT_STEREO_DOWNMIX"
	case CHANNEL_LAYOUT_22POINT2:
		return "CHANNEL_LAYOUT_22POINT2"
	case CHANNEL_LAYOUT_AMBISONIC_FIRST_ORDER:
		return "CHANNEL_LAYOUT_AMBISONIC_FIRST_ORDER"
	default:
		return "[?? Invalid ChannelLayout value]"
	}
}

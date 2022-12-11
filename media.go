package media

import (
	"context"
	"io"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
// TYPES

// MediaFlag is a bitfield of flags for media, including type of media
type MediaFlag uint

// MediaKey is a string which is used for media metadata
type MediaKey string

// DecodeFn is a function which is called for each packet in the media, which
// is associated with a single stream. The function should return an error if
// the decode should be terminated.
type DecodeFn func(context.Context, Packet) error

////////////////////////////////////////////////////////////////////////////////
// INTERFACES

// Manager is an interface to the ffmpeg media library for media manipulation
type Manager interface {
	io.Closer

	// Open media for reading and return it
	OpenFile(path string) (Media, error)

	// Create media for writing and return it
	CreateFile(path string) (Media, error)

	// Create a map of input media. If MediaFlag is MEDIA_FLAG_NONE, then
	// all streams are mapped, or else a combination of MEDIA_FLAG_AUDIO,
	// MEDIA_FLAG_VIDEO, MEDIA_FLAG_SUBTITLE and MEDIA_FLAG_DATA
	// can be used to map specific types of streams.
	Map(Media, MediaFlag) (Map, error)

	// Decode a media file, passing packets to a callback function
	Decode(context.Context, Map, DecodeFn) error

	// Log messages from ffmpeg
	SetDebug(bool)
}

// Map is a mapping of input media, potentially to output media
type Map interface {
	// Return input media
	Input() Media

	// Return streams which are mapped for decoding
	Streams() []Stream
}

// Media is a source or destination of media
type Media interface {
	io.Closer

	// Return best streams for specific types (video, audio, subtitle, data or attachment)
	// or returns empty slice if no streams of that type are in the media file. Only returns
	// one stream of each type.
	//StreamsByType(MediaFlag) []Stream

	// URL for the media
	URL() string

	// Return enumeration of streams
	Streams() []Stream

	// Return media flags for the media
	Flags() MediaFlag

	// Return metadata for the media
	Metadata() Metadata
}

// Stream of data multiplexed in the media
type Stream interface {
	// Return index of stream in the media
	Index() int

	// Return media flags for the stream
	Flags() MediaFlag

	// Return artwork for the stream - if MEDIA_FLAG_ARTWORK is set
	Artwork() []byte
}

// Metadata embedded in the media
type Metadata interface {
	// Return enumeration of keys
	Keys() []MediaKey

	// Return value for key
	Value(MediaKey) any
}

// Packet is a single unit of data in the media
type Packet interface {
	// Flags returns the flags for the packet from the stream
	Flags() MediaFlag

	// Stream returns the stream which the packet belongs to
	Stream() Stream

	// IsKeyFrame returns true if the packet contains a key frame
	IsKeyFrame() bool

	// Pos returns the byte position of the packet in the media
	Pos() int64

	// Duration returns the duration of the packet
	Duration() time.Duration

	// Size of the packet in bytes
	Size() int

	// Bytes returns the raw bytes of the packet
	Bytes() []byte
}

// Codec is an encoder or decoder for a specific media type
type Codec interface {
	// Name returns the unique name for the codec
	Name() string

	// Description returns the long description for the codec
	Description() string

	// Flags for the codec (Audio, Video, Encoder, Decoder)
	Flags() MediaFlag
}

////////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	MEDIA_FLAG_ALBUM             MediaFlag = (1 << iota) // Is part of an album
	MEDIA_FLAG_ALBUM_TRACK                               // Is an album track
	MEDIA_FLAG_ALBUM_COMPILATION                         // Album is a compilation
	MEDIA_FLAG_TVSHOW                                    // Is part of a TV Show
	MEDIA_FLAG_TVSHOW_EPISODE                            // Is a TV Show episode
	MEDIA_FLAG_FILE                                      // Is a file
	MEDIA_FLAG_VIDEO                                     // Contains video
	MEDIA_FLAG_AUDIO                                     // Contains audio
	MEDIA_FLAG_SUBTITLE                                  // Contains subtitles
	MEDIA_FLAG_DATA                                      // Contains data stream
	MEDIA_FLAG_ATTACHMENT                                // Contains attachment
	MEDIA_FLAG_ARTWORK                                   // Contains artwork
	MEDIA_FLAG_CAPTIONS                                  // Contains captions
	MEDIA_FLAG_ENCODER                                   // Is an encoder
	MEDIA_FLAG_DECODER                                   // Is an decoder
	MEDIA_FLAG_NONE              MediaFlag = 0
	MEDIA_FLAG_MAX                         = MEDIA_FLAG_DECODER
)

const (
	MEDIA_KEY_BRAND_MAJOR      MediaKey = "major_brand"       // string
	MEDIA_KEY_BRAND_COMPATIBLE MediaKey = "compatible_brands" // string
	MEDIA_KEY_CREATED          MediaKey = "creation_time"     // time.Time
	MEDIA_KEY_ENCODER          MediaKey = "encoder"           // string
	MEDIA_KEY_ALBUM            MediaKey = "album"             // string
	MEDIA_KEY_ALBUM_ARTIST     MediaKey = "artist"            // string
	MEDIA_KEY_COMMENT          MediaKey = "comment"           // string
	MEDIA_KEY_COMPOSER         MediaKey = "composer"          // string
	MEDIA_KEY_COPYRIGHT        MediaKey = "copyright"         // string
	MEDIA_KEY_YEAR             MediaKey = "date"              // uint
	MEDIA_KEY_DISC             MediaKey = "disc"              // uint
	MEDIA_KEY_ENCODED_BY       MediaKey = "encoded_by"        // string
	MEDIA_KEY_FILENAME         MediaKey = "filename"          // string
	MEDIA_KEY_GENRE            MediaKey = "genre"             // string
	MEDIA_KEY_LANGUAGE         MediaKey = "language"          // string
	MEDIA_KEY_PERFORMER        MediaKey = "performer"         // string
	MEDIA_KEY_PUBLISHER        MediaKey = "publisher"         // string
	MEDIA_KEY_SERVICE_NAME     MediaKey = "service_name"      // string
	MEDIA_KEY_SERVICE_PROVIDER MediaKey = "service_provider"  // string
	MEDIA_KEY_TITLE            MediaKey = "title"             // string
	MEDIA_KEY_TRACK            MediaKey = "track"             // uint
	MEDIA_KEY_VERSION_MAJOR    MediaKey = "major_version"     // string
	MEDIA_KEY_VERSION_MINOR    MediaKey = "minor_version"     // string
	MEDIA_KEY_SHOW             MediaKey = "show"              // string
	MEDIA_KEY_SEASON           MediaKey = "season_number"     // uint
	MEDIA_KEY_EPISODE_SORT     MediaKey = "episode_sort"      // string
	MEDIA_KEY_EPISODE_ID       MediaKey = "episode_id"        // uint
	MEDIA_KEY_COMPILATION      MediaKey = "compilation"       // bool
	MEDIA_KEY_GAPLESS_PLAYBACK MediaKey = "gapless_playback"  // bool
	MEDIA_KEY_ACCOUNT_ID       MediaKey = "account_id"        // string
	MEDIA_KEY_DESCRIPTION      MediaKey = "description"       // string
	MEDIA_KEY_MEDIA_TYPE       MediaKey = "media_type"        // string
	MEDIA_KEY_PURCHASED        MediaKey = "purchase_date"     // time.Time
	MEDIA_KEY_ALBUM_SORT       MediaKey = "sort_album"        // string
	MEDIA_KEY_ARTIST_SORT      MediaKey = "sort_artist"       // string
	MEDIA_KEY_TITLE_SORT       MediaKey = "sort_name"         // string
	MEDIA_KEY_SYNOPSIS         MediaKey = "synopsis"          // string
	MEDIA_KEY_GROUPING         MediaKey = "grouping"          // string
)

////////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (f MediaFlag) String() string {
	if f == MEDIA_FLAG_NONE {
		return f.FlagString()
	}
	str := ""
	for v := MediaFlag(1); v <= MEDIA_FLAG_MAX; v <<= 1 {
		if f&v == v {
			str += "|" + v.FlagString()
		}
	}
	return str[1:]
}

func (f MediaFlag) FlagString() string {
	switch f {
	case MEDIA_FLAG_NONE:
		return "MEDIA_FLAG_NONE"
	case MEDIA_FLAG_ALBUM:
		return "MEDIA_FLAG_ALBUM"
	case MEDIA_FLAG_ALBUM_TRACK:
		return "MEDIA_FLAG_ALBUM_TRACK"
	case MEDIA_FLAG_ALBUM_COMPILATION:
		return "MEDIA_FLAG_ALBUM_COMPILATION"
	case MEDIA_FLAG_TVSHOW:
		return "MEDIA_FLAG_TVSHOW"
	case MEDIA_FLAG_TVSHOW_EPISODE:
		return "MEDIA_FLAG_TVSHOW_EPISODE"
	case MEDIA_FLAG_FILE:
		return "MEDIA_FLAG_FILE"
	case MEDIA_FLAG_VIDEO:
		return "MEDIA_FLAG_VIDEO"
	case MEDIA_FLAG_AUDIO:
		return "MEDIA_FLAG_AUDIO"
	case MEDIA_FLAG_SUBTITLE:
		return "MEDIA_FLAG_SUBTITLE"
	case MEDIA_FLAG_DATA:
		return "MEDIA_FLAG_DATA"
	case MEDIA_FLAG_ATTACHMENT:
		return "MEDIA_FLAG_ATTACHMENT"
	case MEDIA_FLAG_ARTWORK:
		return "MEDIA_FLAG_ARTWORK"
	case MEDIA_FLAG_CAPTIONS:
		return "MEDIA_FLAG_CAPTIONS"
	case MEDIA_FLAG_ENCODER:
		return "MEDIA_FLAG_ENCODER"
	case MEDIA_FLAG_DECODER:
		return "MEDIA_FLAG_DECODER"
	default:
		return "[?? Invalid MediaFlag]"
	}
}

////////////////////////////////////////////////////////////////////////////////
// METHODS

func (f MediaFlag) Is(v MediaFlag) bool {
	return f&v == v
}

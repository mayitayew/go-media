package audio_test

import (
	"io"
	"testing"
	"time"

	// Package imports
	assert "github.com/stretchr/testify/assert"

	// Namespace imports
	. "github.com/mutablelogic/go-media"
	. "github.com/mutablelogic/go-media/pkg/audio"
)

func Test_manager_000(t *testing.T) {
	assert := assert.New(t)
	mgr := New()
	assert.NotNil(mgr)

	// Create an audioframe for input
	in, err := NewAudioFrame(AudioFormat{Rate: 48000, Format: SAMPLE_FORMAT_U8, Layout: CHANNEL_LAYOUT_MONO}, time.Second)
	assert.NoError(err)
	t.Log("in=", in)
	assert.NoError(mgr.Convert(in, AudioFormat{Format: SAMPLE_FORMAT_DBL, Rate: 44100}, func(out AudioFrame) error {
		t.Log("out=", out)
		return io.EOF
	}))

	// Close
	assert.NoError(in.Close())
	assert.NoError(mgr.Close())
}

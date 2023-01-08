package meltysynth_test

import (
	"os"
	"testing"

	"github.com/sinshu/go-meltysynth/meltysynth"
)

const (
	envGS = "MELTYSYNTH_GS" // "GeneralUser GS MuseScore v1.442.sf2"
	envGM = "MELTYSYNTH_GM" // TimGM6mb.sf2
)

func loadGS(t *testing.T) *meltysynth.SoundFont {
	return loadSoundFont(t, envGS)
}

func loadGM(t *testing.T) *meltysynth.SoundFont {
	return loadSoundFont(t, envGM)
}

func loadSoundFont(t *testing.T, env string) *meltysynth.SoundFont {
	p := os.Getenv(env)
	if len(p) == 0 {
		t.Skipf("missing environment variable %q to load soundfont", env)
	}
	f, err := os.Open(p)
	if err != nil {
		t.Fatal(err)
	}
	sf, err := meltysynth.NewSoundFont(f)
	f.Close()
	if err != nil {
		t.Fatal(err)
	}

	return sf
}

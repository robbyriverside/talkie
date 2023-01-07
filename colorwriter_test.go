package talkie_test

import (
	"testing"

	"github.com/robbyriverside/talkie"
	"github.com/stretchr/testify/require"
)

func TestWrite(t *testing.T) {
	cw := talkie.NewColorWriter(talkie.Blue)
	out := []byte("blue out\n")
	n, err := cw.Write(out)
	require.NoError(t, err)

	require.Equal(t, len(out)+cw.WrapSize(), n)
}

func TestErrorWrite(t *testing.T) {
	cw := talkie.NewColorErrorWriter(talkie.Red)
	out := []byte("red out\n")
	n, err := cw.Write(out)
	require.NoError(t, err)

	require.Equal(t, len(out)+cw.WrapSize(), n)
}

func TestPrintln(t *testing.T) {
	cw := talkie.NewColorWriter(talkie.Green)
	out := "green out"
	n, err := cw.Println(out)
	require.NoError(t, err)
	require.Equal(t, len(out)+cw.WrapSize()+1, n)
}

package envfile

import (
	"os"
	"strings"
	"testing"

	"github.com/alecthomas/kong"
	"github.com/stretchr/testify/require"
)

func TestParseENVFileBasic(t *testing.T) {
	var cli struct {
		String string `env:"STRING"`
		Int    int    `env:"INT"`
		Bool   bool   `env:"BOOL"`
	}

	envFile := `STRING=🍕
INT=5
BOOL=true
`

	r, err := ENVFileReader(strings.NewReader(envFile))
	require.NoError(t, err)

	parser := mustNew(t, &cli, kong.Resolvers(r))
	_, err = parser.Parse([]string{})
	require.NoError(t, err)
	require.Equal(t, "🍕", cli.String)
	require.Equal(t, 5, cli.Int)
	require.True(t, cli.Bool)
}

func TestParseENVFileSubstitutions(t *testing.T) {
	var cli struct {
		String  string `env:"STRING"`
		Int     int    `env:"INT"`
		Bool    bool   `env:"BOOL"`
		String2 string `env:"STRING_2"`
	}

	envFile := `STRING=🍕
INT=5
BOOL=true
STRING_2=$STRING
`

	r, err := ENVFileReader(strings.NewReader(envFile))
	require.NoError(t, err)

	parser := mustNew(t, &cli, kong.Resolvers(r))
	_, err = parser.Parse([]string{})
	require.NoError(t, err)
	require.Equal(t, "🍕", cli.String)
	require.Equal(t, 5, cli.Int)
	require.True(t, cli.Bool)
}

func TestPrioritizeEnvVarOverEnvFile(t *testing.T) {
	defer os.Clearenv()

	require.NoError(t, os.Setenv("STRING", "pizza"))

	var cli struct {
		String string `kong:"env=STRING"`
	}

	envFile := `STRING=🍕`

	r, err := ENVFileReader(strings.NewReader(envFile))
	require.NoError(t, err)

	parser := mustNew(t, &cli, kong.Resolvers(r))
	_, err = parser.Parse([]string{})
	require.NoError(t, err)
	require.Equal(t, "pizza", cli.String)
}

func TestPrioritizeDefaultOverEnvFile(t *testing.T) {
	defer os.Clearenv()

	require.NoError(t, os.Setenv("STRING", "pizza"))

	var cli struct {
		String string `kong:"env=STRING,default='pepperoni'"`
	}

	envFile := `STRING=🍕`

	r, err := ENVFileReader(strings.NewReader(envFile))
	require.NoError(t, err)

	parser := mustNew(t, &cli, kong.Resolvers(r))
	_, err = parser.Parse([]string{})
	require.NoError(t, err)
	require.Equal(t, "pizza", cli.String)
}

func mustNew(t *testing.T, cli interface{}, options ...kong.Option) *kong.Kong {
	t.Helper()
	options = append([]kong.Option{
		kong.Name("test"),
		kong.Exit(func(int) {
			t.Helper()
			t.Fatalf("unexpected exit()")
		}),
	}, options...)
	parser, err := kong.New(cli, options...)
	require.NoError(t, err)
	return parser
}

package gunner

import (
	"os"
	"path"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"

	"git.asdf.cafe/abs3nt/gunner/src/yaml"
)

func LoadApp(i any, appName string) {
	yamlDecoder := yaml.New()
	dotenvDecoder := aconfigdotenv.New()

	configDir, err := os.UserConfigDir()
	if err != nil {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		configDir = path.Join(homeDir, ".config")
	}

	filePath := path.Join(configDir, appName)

	loader := aconfig.LoaderFor(i, aconfig.Config{
		AllowUnknownFields: true,
		AllowUnknownEnvs:   true,
		AllowUnknownFlags:  true,
		SkipFlags:          true,
		DontGenerateTags:   true,
		MergeFiles:         true,
		EnvPrefix:          appName,
		FlagPrefix:         appName,
		Files: []string{
			path.Join(filePath, appName+".yml"),
			path.Join(filePath, appName+".yaml"),
			path.Join(filePath, appName+".json"),
			path.Join(filePath, ".env"),
		},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": yamlDecoder,
			".yml":  yamlDecoder,
			".json": yamlDecoder,
			".env":  dotenvDecoder,
		},
	})
	if err := loader.Load(); err != nil {
		panic(err)
	}
}

package cli

import (
	"log"
	"strings"

	"github.com/dpecos/cbox/internal/app/core"
	"github.com/dpecos/cbox/internal/pkg"
	"github.com/dpecos/cbox/pkg/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cboxInstance *models.CBox
)

var rootCmd = &cobra.Command{
	Use:  "cbox",
	Long: pkg.Logo,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if !strings.Contains(err.Error(), "unknown command") {
			log.Fatal(err)
		}
	} else {
		if cboxInstance != nil { // only if the config is initialized
			if err := viper.WriteConfig(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func init() {
	cobra.OnInitialize(func() {
		path := ""
		core.LoadSettings(path)
		cboxInstance = core.Load()
	})
}

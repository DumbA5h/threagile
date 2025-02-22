package threagile

import (
	"github.com/spf13/cobra"
	"github.com/threagile/threagile/pkg/common"
	"github.com/threagile/threagile/pkg/model"
	"github.com/threagile/threagile/pkg/report"
)

func (what *Threagile) initAnalyze() *Threagile {
	analyze := &cobra.Command{
		Use:     common.AnalyzeModelCommand,
		Short:   "Analyze model",
		Aliases: []string{"analyze", "analyse", "run", "analyse-model"},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg := what.readConfig(cmd, what.buildTimestamp)
			commands := what.readCommands()
			progressReporter := common.DefaultProgressReporter{Verbose: cfg.Verbose}

			r, err := model.ReadAndAnalyzeModel(*cfg, progressReporter)
			if err != nil {
				cmd.Printf("Failed to read and analyze model: %v", err)
				return err
			}

			err = report.Generate(cfg, r, commands, progressReporter)
			if err != nil {
				cmd.Printf("Failed to generate reports: %v \n", err)
				return err
			}
			return nil
		},
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	what.rootCmd.AddCommand(analyze)

	return what
}

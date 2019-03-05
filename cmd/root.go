package cmd

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/daniel-pittman/spin/cmd/application"
	"github.com/daniel-pittman/spin/cmd/pipeline"
	"github.com/daniel-pittman/spin/cmd/pipeline-template"
	"github.com/daniel-pittman/spin/version"
)

type RootOptions struct {
	configFile       string
	GateEndpoint     string
	ignoreCertErrors bool
	quiet            bool
	color            bool
	outputFormat     string
}

func Execute(out io.Writer) error {
	cmd := NewCmdRoot(out)
	return cmd.Execute()
}

func NewCmdRoot(out io.Writer) *cobra.Command {
	options := RootOptions{}

	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version.String(),
	}

	cmd.PersistentFlags().StringVar(&options.configFile, "config", "", "path to config file (default $HOME/.spin/config)")
	cmd.PersistentFlags().StringVar(&options.GateEndpoint, "gate-endpoint", "", "Gate (API server) endpoint (default http://localhost:8084)")
	cmd.PersistentFlags().BoolVarP(&options.ignoreCertErrors, "insecure", "k", false, "ignore certificate errors")
	cmd.PersistentFlags().BoolVarP(&options.quiet, "quiet", "q", false, "squelch non-essential output")
	cmd.PersistentFlags().BoolVar(&options.color, "no-color", true, "disable color")
	cmd.PersistentFlags().StringVar(&options.outputFormat, "output", "", "configure output formatting")

	// create subcommands
	cmd.AddCommand(application.NewApplicationCmd(out))
	cmd.AddCommand(pipeline.NewPipelineCmd(out))
	cmd.AddCommand(pipeline_template.NewPipelineTemplateCmd(out))

	return cmd
}

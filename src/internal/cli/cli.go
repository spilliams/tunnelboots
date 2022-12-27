package cli

import (
	"path"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spilliams/terraboots/internal/terraboots"
	"github.com/spilliams/terraboots/pkg/logformatter"
)

var verbose bool
var vertrace bool
var configFile string
var log *logrus.Entry

func init() {
	cobra.OnInitialize(initLogger)
}

func initLogger() {
	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logformatter.PrefixedTextFormatter{
		UseColor: true,
	})
	if verbose {
		logger.SetLevel(logrus.DebugLevel)
	}
	if vertrace {
		logger.SetLevel(logrus.TraceLevel)
	}
	log = logger.WithField("prefix", "main")
}

const commandGroupIDTerraform = "terraform"
const commandGroupIDTerraboots = "terraboots"

var project *terraboots.Project

func NewTerrabootsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "terraboots",
		Short: "A build orchestrator for terraform monorepos",
	}

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "increase log output")
	cmd.PersistentFlags().BoolVar(&vertrace, "vvv", false, "increase log output even more")
	cmd.PersistentFlags().StringVarP(&configFile, "config-file", "c", "terraboots.hcl", "the filename of the project configuration")

	// TODO: version command
	cmd.CompletionOptions.DisableDefaultCmd = true

	cmd.AddGroup(&cobra.Group{ID: commandGroupIDTerraboots, Title: "Working with your terraboots project"})
	cmd.AddGroup(&cobra.Group{ID: commandGroupIDTerraform, Title: "Terraform Commands"})

	// cmd.AddCommand(newTerraformCommand("init"))
	// cmd.AddCommand(newTerraformCommand("plan"))
	// cmd.AddCommand(newTerraformCommand("apply"))
	// cmd.AddCommand(newTerraformCommand("destroy"))
	// cmd.AddCommand(newTerraformCommand("output"))
	// cmd.AddCommand(newTerraformCommand("console"))

	cmd.AddCommand(newScopeCommand())
	cmd.AddCommand(newRootCommand())

	return cmd
}

func bootsbootsPreRunE(cmd *cobra.Command, args []string) error {
	log.Debugf("Using project configuration file: %s", configFile)
	var err error
	project, err = terraboots.ParseProject(configFile, log.Logger)
	if err != nil {
		return err
	}

	rootsDir := path.Join(path.Dir(configFile), project.RootsDir)
	rootsDir, err = filepath.Abs(rootsDir)
	if err != nil {
		return err
	}
	project.RootsDir = rootsDir
	// log.Debugf("Project roots directory: %s", project.RootsDir)
	// log.Debugf("Project scope data files: %s", project.ScopeDataFiles)

	return nil
}

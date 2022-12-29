package cli

import (
	"github.com/spf13/cobra"
)

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "root",
		Aliases: []string{"r"},
		Short:   "Commands relating to root modules",
		GroupID: commandGroupIDTerraboots,

		PersistentPreRunE: bootsbootsPreRunE,
	}

	cmd.AddCommand(newRootBuildCommand())
	// cmd.AddCommand(newRootGenerateCommand())
	// cmd.AddCommand(newRootGraphCommand())
	// cmd.AddCommand(newRootListCommand())

	return cmd
}

func newRootBuildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "build ROOT [SCOPE [SCOPE...]]",
		Aliases: []string{"b"},
		Short:   "Builds the given root and prints the location of the built root to stdout",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			scopes := make([]string, len(args)-1)
			for i := 1; i < len(args); i++ {
				scopes[i-1] = args[i]
			}
			_, err := project.BuildRoot(args[0], scopes)
			return err
		},
	}

	return cmd
}

// func newRootGenerateCommand() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "generate",
// 		Short: "",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			logger.Warn("not yet implemented")
// 			return nil
// 		},
// 	}
// 	return cmd
// }

// func newRootGraphCommand() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "graph",
// 		Short: "",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			logger.Warn("not yet implemented")
// 			return nil
// 		},
// 	}
// 	return cmd
// }

// func newRootListCommand() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "list",
// 		Short: "",
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			logger.Warn("not yet implemented")
// 			return nil
// 		},
// 	}
// 	return cmd
// }
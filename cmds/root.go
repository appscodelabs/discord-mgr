/*
Copyright AppsCode Inc.

Licensed under the Mozilla Public License Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.mozilla.org/en-US/MPL/2.0/

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmds

import (
	"flag"

	"github.com/spf13/cobra"
	v "gomodules.xyz/x/version"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "discord-mgr [command]",
		Short:             `AppsCode Discord Server Manager CLI`,
		DisableAutoGenTag: true,
	}

	flags := rootCmd.PersistentFlags()
	flags.AddGoFlagSet(flag.CommandLine)

	rootCmd.AddCommand(NewCmdAddMember())
	rootCmd.AddCommand(NewCmdAddCompany())
	rootCmd.AddCommand(NewCmdRemoveMember())
	rootCmd.AddCommand(v.NewCmdVersion())
	return rootCmd
}

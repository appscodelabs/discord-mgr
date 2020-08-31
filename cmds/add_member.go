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
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/cobra"
)

/*
discord-mgr add-role --guild=appscode --role=company_a --username=tamal
*/
func NewCmdAddMember() *cobra.Command {
	var (
		guildName string = "appscode"
		roleName  string
		username  string
	)
	cmd := &cobra.Command{
		Use:               "add-role",
		Short:             "Add member to a role",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
			if err != nil {
				panic(err)
			}

			// Open a websocket connection to Discord and begin listening.
			err = dg.Open()
			if err != nil {
				panic(err)
			}
			defer func() {
				_ = dg.Close()
			}()

			return AddUserToRole(dg, guildName, roleName, username)
		},
	}

	cmd.Flags().StringVar(&guildName, "guild", guildName, "Name of guild/server")
	cmd.Flags().StringVar(&roleName, "role", roleName, "Name of role")
	cmd.Flags().StringVar(&username, "username", username, "Usename of the member")
	return cmd
}

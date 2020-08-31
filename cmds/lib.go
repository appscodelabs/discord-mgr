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
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

/*
func main2() {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		panic(err)
	}
	defer dg.Close()

	var g *discordgo.Guild

	for _, guild := range dg.State.Guilds {
		fmt.Println(guild.Name)
		fmt.Println(guild.ID)

		st, err := dg.Guild(guild.ID)
		if err != nil {
			panic(err)
		}
		fmt.Println(st.Name)

		// Get channels for this guild
		channels, _ := dg.GuildChannels(guild.ID)

		for _, c := range channels {
			// Check if channel is a guild text channel and not a voice or DM channel
			if c.Type != discordgo.ChannelTypeGuildText {
				continue
			}
			fmt.Println(c.Name)

			//// Send text message
			//s.ChannelMessageSend(
			//	c.ID,
			//	fmt.Sprintf("testmsg (sorry for spam). Channel name is %q", c.Name),
			//)
		}

		if st.Name == "appscode" {
			g = st
			break
		}
	}

	findMember(dg, g.ID, "tamal")

	// Create Channel
	ch1, err := dg.GuildChannelCreate(g.ID, "company_a", discordgo.ChannelTypeGuildText)
	if err != nil {
		panic(err)
	}
	fmt.Println(ch1.ID)
	fmt.Println(ch1.Name)

		//	455530176357335040
		//	appscode
		//	749747613263462511
		//	company_a
		//
		//Roles:
		//455530176357335040
		//@everyone
		//749725404482240513
		//AppsCodeBots
		//NEW_ROLE 749750688141017178 company_a
		//
		//
		//GuildID=  Nick=  User.ID= 322188333470056448 User.Email=  roles=
		//User.ID= 322188333470056448 User.Email=  Username= tamal
		//GuildID=  Nick=  User.ID= 682479553238663190 User.Email=  roles=
		//User.ID= 682479553238663190 User.Email=  Username= codefx9
		//GuildID=  Nick=  User.ID= 749722091602968586 User.Email=  roles= 749725404482240513
		//User.ID= 749722091602968586 User.Email=  Username= 1gtm

	guildID := g.ID

	// s.User()

	members, err := dg.GuildMembers(guildID, "", 100)
	if err != nil {
		panic(err)
	}
	for _, m := range members {

		m2, err := dg.GuildMember(guildID, m.User.ID)
		if err != nil {
			panic(err)
		}
		fmt.Println("GuildID=", m2.GuildID, "Nick=", m2.Nick, "User.ID=", m2.User.ID, "User.Email=", m2.User.Email, "roles=", strings.Join(m.Roles, "|"))

		u2, err := dg.User(m.User.ID)
		if err != nil {
			panic(err)
		}
		fmt.Println("User.ID=", u2.ID, "User.Email=", u2.Email, "Username=", u2.Username)
	}

	roles, err := dg.GuildRoles(guildID)
	if err != nil {
		panic(err)
	}
	for _, r := range roles {
		fmt.Println(r.ID)
		fmt.Println(r.Name)
	}

	role, err := dg.GuildRoleCreate(guildID)
	if err != nil {
		panic(err)
	}
	role, err = dg.GuildRoleEdit(guildID, role.ID, "company_a", 0x9b59b6, true, 0, true)
	if err != nil {
		panic(err)
	}
	fmt.Println("NEW_ROLE", role.ID, role.Name)

	// s.GuildMemberRoleAdd()

	// s.ChannelPermissionSet()

	err = dg.GuildMemberRoleAdd(guildID, "682479553238663190", "749750688141017178")
	if err != nil {
		panic(err)
	}

	perm := discordgo.PermissionViewChannel |
		discordgo.PermissionSendMessages |
		// discordgo.PermissionSendTTSMessages |
		// discordgo.PermissionManageMessages |
		discordgo.PermissionEmbedLinks |
		discordgo.PermissionAttachFiles |
		discordgo.PermissionReadMessageHistory // |
		// discordgo.PermissionMentionEveryone

	err = dg.ChannelPermissionSet("749747613263462511", "749750688141017178", string(discordgo.AuditLogOptionsTypeRole), perm, 0)
	if err != nil {
		panic(err)
	}
}
*/

func AddCompany(dg *discordgo.Session, guildName, parentChannel, channelName string) error {
	// find channel by name, if found return nil
	// else
	// create a text channel with name under category/parent CLIENTS
	// create a role with name
	// add role to channel with necessary permission

	guild, err := findGuild(dg, guildName)
	if err != nil {
		return err
	}

	// Get channels for this guild
	channels, err := dg.GuildChannels(guild.ID)
	if err != nil {
		return err
	}

	var parent, ch *discordgo.Channel
	for _, c := range channels {
		if c.Type == discordgo.ChannelTypeGuildCategory && strings.EqualFold(c.Name, parentChannel) {
			parent = c
			continue
		}
		if c.Type == discordgo.ChannelTypeGuildText && strings.EqualFold(c.Name, channelName) {
			ch = c
			continue
		}
	}
	if ch != nil {
		return nil // channel already exists
	}
	if parent == nil {
		return fmt.Errorf("parent %s not found", parentChannel)
	}

	// Create Channel
	ch, err = dg.GuildChannelCreateComplex(guild.ID, discordgo.GuildChannelCreateData{
		Name:                 channelName,
		Type:                 discordgo.ChannelTypeGuildText,
		Topic:                "",
		Bitrate:              0,
		UserLimit:            0,
		RateLimitPerUser:     0,
		Position:             0,
		PermissionOverwrites: nil,
		ParentID:             parent.ID,
		NSFW:                 false,
	})
	if err != nil {
		return err
	}

	role, err := dg.GuildRoleCreate(guild.ID)
	if err != nil {
		return err
	}
	role, err = dg.GuildRoleEdit(guild.ID, role.ID, channelName, 0x9b59b6, true, 0, false)
	if err != nil {
		return err
	}

	perm := discordgo.PermissionViewChannel |
		discordgo.PermissionSendMessages |
		discordgo.PermissionEmbedLinks |
		discordgo.PermissionAttachFiles |
		discordgo.PermissionReadMessageHistory |
		discordgo.PermissionMentionEveryone |
		discordgo.PermissionAddReactions |
		discordgo.PermissionUseExternalEmojis

	err = dg.ChannelPermissionSet(ch.ID, role.ID, string(discordgo.AuditLogOptionsTypeRole), perm, 0)
	if err != nil {
		return err
	}
	return nil
}

func AddUserToRole(dg *discordgo.Session, guildName, roleName, username string) error {
	// ensure username is a member of the channel
	// if yes, then add user to the role
	guild, err := findGuild(dg, guildName)
	if err != nil {
		return err
	}
	member, err := findMember(dg, guild.ID, username)
	if err != nil {
		return err
	}
	role, err := findRole(dg, guild.ID, roleName)
	if err != nil {
		return err
	}
	return dg.GuildMemberRoleAdd(guild.ID, member.User.ID, role.ID)
}

func RemoveUserFromRole(dg *discordgo.Session, guildName, roleName, username string) error {
	// ensure username is a member of the channel
	// if yes, then add user to the role
	guild, err := findGuild(dg, guildName)
	if err != nil {
		return err
	}
	member, err := findMember(dg, guild.ID, username)
	if err != nil {
		return err
	}
	role, err := findRole(dg, guild.ID, roleName)
	if err != nil {
		return err
	}
	return dg.GuildMemberRoleRemove(guild.ID, member.User.ID, role.ID)
}

func findGuild(dg *discordgo.Session, guildName string) (*discordgo.Guild, error) {
	for _, guild := range dg.State.Guilds {
		guild, err := dg.Guild(guild.ID)
		if err != nil {
			return nil, err
		}
		if strings.EqualFold(guild.Name, guildName) {
			return guild, nil
		}
	}
	return nil, &ErrNotFound{Type: "Guild", Name: guildName}
}

func findMember(dg *discordgo.Session, guildID, username string) (*discordgo.Member, error) {
	var userID string
	for {
		members, err := dg.GuildMembers(guildID, userID, 500)
		if err != nil {
			return nil, err
		}
		if len(members) == 0 {
			break
		}
		for _, m := range members {
			if strings.EqualFold(m.User.Username, username) {
				return m, nil
			}
		}
		userID = members[len(members)-1].User.ID
	}
	return nil, &ErrNotFound{Type: "Member", Name: guildID + "/" + username}
}

func findRole(dg *discordgo.Session, guildID, role string) (*discordgo.Role, error) {
	roles, err := dg.GuildRoles(guildID)
	if err != nil {
		return nil, err
	}
	for _, r := range roles {
		if strings.EqualFold(r.Name, role) {
			return r, nil
		}
	}
	return nil, &ErrNotFound{Type: "Role", Name: guildID + "/" + role}
}

type ErrNotFound struct {
	Type string
	Name string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("%s %s not found", e.Type, e.Name)
}

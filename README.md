# discord-mgr

## Usage

```console
discord-mgr add-company --guild=appscode --role=client_company
discord-mgr add-role --guild=appscode --role=client_company --username=client_user
discord-mgr remove-role --guild=appscode --role=client_company --username=client_user
```


## API Clients

- https://discordapi.com/unofficial/comparison.html

## Permissions

- [How is the permission hierarchy structured?](https://support.discord.com/hc/en-us/articles/206141927)
- [How do I set up Permissions?](https://support.discord.com/hc/en-us/articles/206029707-How-do-I-set-up-Permissions)
- [How do I set up a Role-Exclusive channel?](https://support.discord.com/hc/en-us/articles/206143877-How-do-I-set-up-a-Role-Exclusive-channel-)
- [How do I set up an announcements channel?](https://support.discord.com/hc/en-us/articles/205369668)

- https://discord.com/developers/docs/topics/permissions#role-object
- https://discord.com/developers/docs/topics/permissions

### Server Permissions

An individual's server-wide permissions are determined by adding up all the allows for roles assigned to that individual along with the allows for @everyone.

### Channel Permissions

Channel permissions start with server permissions as a base. Then, the hierarchy is as follows:

- Apply denies of @everyone on channel
- Apply allows of @everyone on channel
- Sum up all the denies of a member's roles and apply them at once
- Sum up all the allows of a member's roles and apply them at once
- Apply denies for a specific member if they exist
- Apply allows for a specific member if they exist


## Applications

- https://discordapp.com/developers/applications/

## API Token

- [How to Get a Discord Bot Token](https://www.writebots.com/discord-bot-token/)
- [How To Find Your Discord Token (2020)](https://www.youtube.com/watch?v=xuB1WQVM3R8)

**Add Bot**: https://discord.com/api/oauth2/authorize?client_id=749722091602968586&permissions=8&scope=bot

## Examples

- https://github.com/bwmarrin/discordgo/wiki/Awesome-DiscordGo
- https://github.com/bwmarrin/discordgo/tree/master/examples

# Bouquet

Bouquet is a small discord (injected) client.

## Installing



## Development

For local development of bouquet you can use the `./dev` helper script.
The script has the following subcommands (./dev [subcommand]): 

Command | Description
--------|----------------
backup  | Backup the core.asar file in `./`
revert  | Copy the backup back into Discord to revert changes
src     | Extract the current core.asar using the official Electron ASAR CLI tool into `./discord-source/`
cli     | Run the CLI version of the application
clean   | Moves the backup back into discord and removes `./discord-source/`


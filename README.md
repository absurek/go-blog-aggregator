# Gator

A multi user rss aggreagator CLI.

## Dependencies

Gator uses PostgreSQL to store data and is implemented in the Go programming language.

## Installation

You can install gator using `go install github.com/absurek/go-blog-aggregator`

## Configuration

The CLI expects a configuration file in your home directory: `~/.gatorconfig.json`

Example:

```jsonc
{
  "db_url": "postgres://dev:dev@dev-db:5432/devdb?sslmode=disable"
}
```

You will need to set up the config file manually before executing the CLI for the first time.

## Commands

- You can register a user using the `register <username>` command
- You can list all users using the `users` command
- You can log in using the `login <user>` command
- You can add a new rss feed using the `addfeed <name> <url>` command
- You can list all rss feeds using the `feeds` command 
- You can follow a feed using the `follow <url>` command
- You can list which feeds you are currently following using the `follows` command
- You can unfollow a feed using the `unfollow <url>` command
- You can start aggregating using the `agg` command
- You can browse your aggregated posts using the `browse` command 

Example:
```bash
go-blog-aggregator register johndoe
```

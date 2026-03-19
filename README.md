# Orchestra
A simple one binary orchestrator for game server hosting, written in Go with a solidjs frontend.

## Table of Contents
- [Orchestra](#orchestra)
  - [Table of Contents](#table-of-contents)
  - [Project Synopsys](#project-synopsys)
  - [ENV Variables](#env-variables)
    - [ENV Naming Breakdown](#env-naming-breakdown)
    - [ENV Settings](#env-settings)
  - [User Stories (Functionalities)](#user-stories-functionalities)
    - [User](#user)
    - [Admin](#admin)
    - [Service owner](#service-owner)
    - [Other](#other)
  - [Q \& A](#q--a)
  - [European Made](#european-made)

## Project Synopsys


Orchestra is an visual orchestrator that allows you to setup services quickly.

This project is born out a need to host game services and to make it simple for friends and family and to keep it in one binary without the setup hassle.
I would like it to be compliant to Pelican/Pterodactyl's Egg system so you can ingest those files and set up the game servers like a breeze.

TODO: Write more.

INFO: Code organization: https://go.dev/doc/modules/layout

## ENV Variables

### ENV Naming Breakdown

To ensure all env variables have a consistent naming scheme, I enforce the following for Orchestra:

```env
ORCHESTRA_{API|WEB}_{ENV_VAR}
```

### ENV Settings

```env
ORCHESTRA_API_PORT=9810

ORCHESTRA_API_OAUTH_ENABLED=bool
ORCHESTRA_API_OAUTH_SECRET=string
ORCHESTRA_API_OAUTH_AUTH_URL=string
ORCHESTRA_API_OAUTH_TOKEN_URL=string
ORCHESTRA_API_OAUTH_CLIENT_ID=string
ORCHESTRA_API_OAUTH_REDIRECT_URL=string

ORCHESTRA_API_LOCAL_USERS_ENABLED=bool
ORCHESTRA_API_LOCAL_USERS_REGISTRATION=bool

ORCHESTRA_WEB_PORT=9800
```

## User Stories (Functionalities)

All marked user stories are implemented, those that are not are WIP.

### User

- [ ] As a user I want to access my service through a website to be able to see my services.
- [ ] As a user I want to register myself (If OAuth is disabled) through the website to be able to see my services.
- [ ] As a user I want to set 2FA myself through the website to be able to see my services.
- [ ] As a user I want to login to the service to setup a server.
- [ ] As a user I want to login to the service through my OAuth2 provider to setup a server.
- [ ] As a user I want to have a dashboard to see all assigned services that I own.
- [ ] As a user I want to start/stop a service.
- [ ] As a user I want to connect to SSH and browse my containers filesystem to check mods.
- [ ] As a user I want to import a [Pelican/Pterodactyl Egg](https://raw.githubusercontent.com/pelican-eggs/games-steamcmd/refs/heads/main/eco/egg-eco.json) to setup a new service.
- [ ] As a user I want to setup a new service like the above egg setup.
- [ ] As a user I want to export a newly made egg so I can share it with a friend.
- [ ] As a user I want to access my game server through $\$URL/\$USER/\$GAME-UNIQUE-ID$

### Admin
- [ ] As an administrator I want to invite a user to a service if OAuth is disabled.
- [ ] As an administrator I want to recover a user if OAuth is disabled.
- [ ] As an administrator I want to assign a service to a user.
- [ ] As an administrator I want to see who is using the resources on the server.
- [ ] As an administrator I want to set default limits so people can't abuse the service.
- [ ] As an administrator I want to change limits so I can grant extended limits to friends.
- [ ] As an administrator I want to access a service
- [ ] As an administrator I want to disable a service
- [ ] As an administrator I want to remove a service.

### Service owner
- [ ] As the owner of the service I want to host services on Docker/Podman/K3s/K8s.
- [ ] As the owner of the service I want to make Orchestra use terraform to spawn services independently of backend.
- [ ] As the owner of the service I want to select how to login and register, local or OAUTH2.
- [ ] As the owner of the service I want to be able to register an account without OAuth2 authorizor.
- [ ] As the owner of the service I want to restrict sign-ups.
- [ ] As the owner of the service I want to have Orchestra make use of a public or private OAuth2 authorizor.
- [ ] As the owner of the service I want to use Caddy as my reverse proxy.
- [ ] As the owner of the service I want to make it simple to deploy on my own server via a container.
- [ ] As the owner of the service I want to be able to check the logs for any suspicious behaviour with services.
- [ ] As the owner of the service I want to ban access from countries/vpns that are known for hacking services.

### Other
- [ ] As an \{\$ACTOR\} I want \{\$FUNCTION\} because {$REASON}

## Q & A

## European Made

The author is European and living in the European Union. The software was created in Europe.
# piggi
Prompt-base Interactivity based on golang, for GitHub

*Still in pre-alpha...so be aware of using it*


# Overview
There are already several great terminal-based tools for GitHub
- [github/hub](https://github.com/github/hub)
- [donnemartin/gitsome](https://github.com/donnemartin/gitsome)

But I wanted to make something which can interact more simple, using direction key + enter key.

## Detail
This app is based on
> - Golang
> - GraphQL
> - GitHub API v4


# Basic Usage
Get source first, and install dependencies.
```
$ git clone https://github.com/djKooks/piggi.git
...
$ cd piggi
$ dep ensure
...
```

You need to register your information in `user.json` file.
Create `user.json` file, and put data as below:
```json
{
    "token": "your-personal-access-token",
    "id": "your-github-id"
}
```

You can create new personal access token in [here](https://github.com/settings/tokens).

Now you can see user info like:
```
$ go run . user

==================================

 ID: djKooks

 Bio:  your-bio

 Email:  your-email

 Company:  your-company

==================================
```

or your repository list:
```
$ go run . repo

==================================

 Look over Repositories

==================================

Use the arrow keys to navigate: ↓ ↑ → ← 
  [ Repositories ]
  ⛳ djKooks/piggi
    djKooks/djkooks.github.com
    djKooks/jopt
    djKooks/gnoter
↓   ocombe/ocLazyLoad

--------- Repo ----------
Name:               djKooks/piggi
Description:        (WIP!) Prompt-base Interactivity for GitHub
```

## Option
Name | Description
--- | --- 
user | show user information
repo | show your repositories, order by updated date
issue | show your issues, order by updated date
pr | show your pull requests, order by updated date


# Contributing
Welcome for contributions!

It is very unstable, and still lot of things to do:
> - Suggestions which will make it more useful
> - Flexible design
> - Code refactoring
> - Bug fix
> - Document update
> - ...

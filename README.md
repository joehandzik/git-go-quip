# git-go-quip

This is mostly an experimentation project for my own purposes, but someone else may find this useful someday!

The goals are pretty straightforward:

* Create a CLI interface for listing Quip documents and Github repositories
* Allow users to pull Quip documents into Github with a single command

The requirements are simple. You'll need golang, and you'll need two environment variables:

QUIP_TOKEN - an authentication token for interacting with your Quip account, which can be generated at https://quip.com/api/personal-token
GITHUB_TOKEN - an authentication token for interacting with your GitHub account, which can be generated at https://github.com/settings/tokens (I'll update this if I have a clear set of permission requirements, but for now making everything read-only except for creating repositories and committing to repositories seems like the bare minimum)

I'll keep a running update of what works as I remember to.

Works: list-quip
       list-github

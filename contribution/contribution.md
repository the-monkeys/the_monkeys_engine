# Contributing Guidelines
We're glad you're thinking about contributing to The Monkeys. If you think something is missing or could be improved, please open issues and pull requests. If you'd like to help this project grow, we'd love to have you. 

* If you find any issue or bug please create a Github issue or mail us at [mail.themonkeys.life@gmail.com](mail.themonkeys.life@gmail.com). 
* Create branches in your fork, and submit PRs from your forked branch.

# Local Setup Requirement
* Docker
* Golang v1.18.0+
* Protoc compiler
* [migrate](https://github.com/golang-migrate/migrate)

Once you have pull the code run `docker compose up --build` and that will run the development server in the local machine.



# Install linting tool
```
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin vX.Y.Z
```

### Run go lint command
$ `golangci-lint run`



# PR Approval and Merge

We are keeping some checks before merging the PRs to the main branch to maintain the code for a long time and for now we have setup the following rules, it future we may update the rules use some automation and pipeline for code and consistency checks.

* All the PRs need to be approved by [Dave Augustus](https://github.com/daveaugustus) before the merge.
* Code consistency needs to be checked before raising the PR.
* Spelling needs to be checked before the PR.
* The sensitive information like environment variables shouldn't be in the code.
* Linting needs to be checked.


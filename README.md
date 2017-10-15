Go Deploy (WIP)
===
<b>
This software is not ready yet and will do nothing for the moment!

Use this at your own risk. This is created solely for personal use!
</b>

A Continuous Delivery tool written in Go.

## Installation

```bash
$ make
```

## Run

Create a file go-deploy.yml and add your configuration in it

```yaml
build_directory: "~"
projects:
  project_name: 
    vcs: 
      link: "https://github.com/ruudvdd/go-deploy"
    environments:
      production: 
        vcs: 
          branch: "staging"
        build: 
          cmd:
            - "build"
            - "deploy"
```

* `build_directory`: The directory where to clone the projects for preparing the deploy
* `projects`: Add your different project configrations here in the same format as above example with key (project_name) the name of the project
* `vcs.link (project)`: Required on project level. The URL where the project is stored. Currently only git repositories are supported. WIP: Support for private repositories via OAuth
* `environments`: Add your different environment configurations here
* `vcs (environment)`
    * `branch`: The branch you want to deploy and will be checked out during the process
* `build.cmd`: List of bash commands to run to deploy the project. The deploy will stop when one of these commands fail (return code <> 0).

Run

```bash
$ go-deploy <project_name> <environment_name>
```

By default the go-deploy.yml configuration file in the current working directory will be used. To use a configuration file in another directory, you have to specify it with the -config flag

```bash
$ go-deploy -config /path/to/config/go-deploy.yml <project_name> <environment_name>
```

## TODO

* Implement actual deploy process
* Config validation
* Binary for adding projects and environments to a configuration file via the command line
* Support OAuth for Github and Bitbucket
* ...

## License

Copyright (C) 2017 Ruud Van den Dungen

This project is distributed under the MIT license. See the [LICENSE](https://github.com/ruudvdd/go-deploy/blob/master/LICENSE) file for details.
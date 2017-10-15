package config

import (
	"fmt"
	"strings"

	"github.com/olebedev/config"
)

// Config struct with helper funcs to parse the config
type Config struct {
	BuildDirectory string
	Projects       map[string]*Project
}

// AddProject add a project to the config
func (cfg *Config) AddProject(project *Project) {
	if cfg.Projects == nil {
		cfg.Projects = map[string]*Project{}
	}

	cfg.Projects[project.Name] = project
}

// GetProject Get a project by name or an error when it does not exist
func (cfg Config) GetProject(projectName string) (*Project, error) {
	project, ok := cfg.Projects[projectName]

	if !ok {
		return nil, fmt.Errorf("Could not find project with name %q", projectName)
	}

	return project, nil
}

// NewConfig Create a Config message based on a config.Config object
func NewConfig(fileName string) (*Config, error) {
	newCfg := Config{}

	cfg, err := config.ParseYamlFile(fileName)

	if err != nil {
		return nil, err
	}

	// required build_directory
	buildDirectory, err := cfg.String("build_directory")
	if err != nil {
		return nil, err
	}

	newCfg.BuildDirectory = buildDirectory

	cfgProjects := cfg.MapUMap("projects")

	if err != nil {
		return nil, err
	}

	for projectName := range cfgProjects {
		projectCfg, err := cfg.Get("projects." + projectName)

		if err != nil {
			return nil, err
		}

		project := Project{Name: projectName, VcsLink: projectCfg.UString("vcs.link", "")}

		newCfg.AddProject(&project)

		cfgEnvironments := projectCfg.MapUMap("environments")

		for envName := range cfgEnvironments {
			envCfg, err := projectCfg.Get("environments." + envName)

			if err != nil {
				return nil, err
			}

			cfgBuildCmds, err := envCfg.Get("build.cmd")

			buildCmds := []string{}

			if err == nil {
				for _, buildCmd := range cfgBuildCmds.UList("") {
					buildCmds = append(buildCmds, buildCmd.(string))
				}
			}

			env := Environment{
				Name:      envName,
				VcsLink:   envCfg.UString("vcs.link", ""),
				VcsBranch: envCfg.UString("vcs.branch", ""),
				BuildCmd:  buildCmds,
			}

			project.AddEnvironment(&env)
		}
	}

	return &newCfg, nil
}

// Project representation of a project config
type Project struct {
	Name         string
	VcsLink      string
	Environments map[string]*Environment
}

// AddEnvironment add an environment to the project
func (project *Project) AddEnvironment(env *Environment) {
	if project.Environments == nil {
		project.Environments = map[string]*Environment{}
	}

	project.Environments[env.Name] = env
}

// GetEnvironment get an environment by name or return error when not found
func (project *Project) GetEnvironment(envName string) (*Environment, error) {
	env, ok := project.Environments[envName]

	if !ok {
		return nil, fmt.Errorf("Could not find environment in project %q with name %q. Possible environments are %q", project.Name, envName, "["+strings.Join(project.getEnvironmentNames(), ", ")+"]")
	}

	return env, nil
}

// getEnvironmentNames get all environments of a project
func (project Project) getEnvironmentNames() []string {
	keys := []string{}

	for key := range project.Environments {
		keys = append(keys, key)
	}

	return keys
}

// Environment representation of an environment config in a Project
type Environment struct {
	Name      string
	VcsLink   string
	VcsBranch string
	BuildCmd  []string
}

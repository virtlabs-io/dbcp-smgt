package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/virtlabs-io/dbcp-smgt/internal/config"     // Configuration package
	"github.com/virtlabs-io/dbcp-smgt/internal/core"       // Core service manager
	"github.com/virtlabs-io/dbcp-smgt/internal/logging"    // Logger wrapper (e.g., Zap)
	"github.com/virtlabs-io/dbcp-smgt/pkg/types"           // ServiceContext & Service interface definitions
	"github.com/virtlabs-io/dbcp-smgt/services/postgresql" // PostgreSQL service implementation
	"go.uber.org/zap"
)

// GlobalConfig defines the global configuration structure used in main.
type GlobalConfig struct {
	LogLevel    string         `mapstructure:"log_level"`
	Environment string         `mapstructure:"environment"`
	Modules     []ModuleConfig `mapstructure:"modules"`
}

// ModuleConfig defines configuration for an individual module.
type ModuleConfig struct {
	Name       string `mapstructure:"name"`
	Enabled    bool   `mapstructure:"enabled"`
	Path       string `mapstructure:"path"`
	ConfigFile string `mapstructure:"config_file"`
}

// convertModules converts the modules from the config package type to our local ModuleConfig type.
func convertModules(modules []config.ModuleConfig) []ModuleConfig {
	ret := make([]ModuleConfig, len(modules))
	for i, m := range modules {
		ret[i] = ModuleConfig{
			Name:       m.Name,
			Enabled:    m.Enabled,
			Path:       m.Path,
			ConfigFile: m.ConfigFile,
		}
	}
	return ret
}

func main() {
	// Define command-line flags.
	var configFile string
	flag.StringVar(&configFile, "f", "etc/config.yaml", "Path to the global configuration file")
	flag.StringVar(&configFile, "config", "etc/config.yaml", "Path to the global configuration file")
	modulesFlag := flag.String("modules", "", "Comma-separated list of module names to load (overrides config)")
	flag.Parse()

	// Create logger.
	logger := logging.NewLogger()
	defer logger.Sync()

	// Load global configuration from file.
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading global config: %v\n", err)
		os.Exit(1)
	}

	// Convert the loaded configuration into our local GlobalConfig.
	globalCfg := GlobalConfig{
		LogLevel:    cfg.Global.LogLevel,
		Environment: cfg.Global.Environment,
		Modules:     convertModules(cfg.Global.Modules),
	}

	// Process any module override from the command line.
	overrideModules := []string{}
	if *modulesFlag != "" {
		overrideModules = strings.Split(*modulesFlag, ",")
	}

	// Create a ServiceContext.
	svcCtx := &types.ServiceContext{
		Version: "1.0.0", // This could also be configured.
		Config:  map[string]interface{}{"env": globalCfg.Environment},
		Logger:  logger,
		DryRun:  false, // Set true for dry-run behavior.
	}

	// Initialize the ServiceManager.
	manager := core.NewServiceManager(svcCtx)

	// Set up signal handling.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	shutdownChan := make(chan struct{})

	// Goroutine to handle signals.
	go func() {
		for {
			sig := <-sigChan
			switch sig {
			case syscall.SIGHUP:
				logger.Info("Received SIGHUP signal; reloading configuration and modules")
				newCfg, err := config.LoadConfig(configFile)
				if err != nil {
					logger.Error("Failed to reload configuration", zap.Error(err))
				} else {
					// Update global configuration.
					globalCfg = GlobalConfig{
						LogLevel:    newCfg.Global.LogLevel,
						Environment: newCfg.Global.Environment,
						Modules:     convertModules(newCfg.Global.Modules),
					}
					svcCtx.Config["env"] = globalCfg.Environment
					// Here you might unload/load modules as needed based on newCfg.
					logger.Info("Configuration reloaded", zap.Any("globalCfg", globalCfg))
				}
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				logger.Info("Received termination signal", zap.String("signal", sig.String()))
				close(shutdownChan)
				return
			default:
				logger.Info("Received unhandled signal", zap.String("signal", sig.String()))
			}
		}
	}()

	// Start a goroutine to process modules.
	go func() {
		for _, mod := range globalCfg.Modules {
			// Skip disabled modules.
			if !mod.Enabled {
				logger.Info("Module disabled", zap.String("module", mod.Name))
				continue
			}
			// If module override is provided, load only matching modules.
			if len(overrideModules) > 0 {
				found := false
				for _, m := range overrideModules {
					if m == mod.Name {
						found = true
						break
					}
				}
				if !found {
					logger.Info("Module skipped by command-line filter", zap.String("module", mod.Name))
					continue
				}
			}
			// For demonstration, we directly instantiate the PostgreSQL service for the "postgresql" module.
			if strings.ToLower(mod.Name) == "postgresql" {
				svc := &postgresql.PostgreSQLService{
					VersionField: "13.3", // Default or loaded from module-specific config (mod.ConfigFile).
					// Additional fields may be initialized here.
				}
				logger.Info("Executing install for module", zap.String("module", mod.Name))
				if err := manager.ExecuteInstall(context.Background(), svc); err != nil {
					logger.Error("Installation failed for module",
						zap.String("module", mod.Name),
						zap.Error(err))
				}
			} else {
				logger.Info("Module not recognized for dynamic loading in this demo", zap.String("module", mod.Name))
			}
		}
	}()

	// Wait for shutdown signal.
	<-shutdownChan
	logger.Info("Gracefully shutting down...")
	// Perform graceful shutdown tasks: unload modules, cleanup resources, etc.
	time.Sleep(2 * time.Second) // Simulate graceful shutdown.
	logger.Info("Shutdown complete.")
	os.Exit(0)
}

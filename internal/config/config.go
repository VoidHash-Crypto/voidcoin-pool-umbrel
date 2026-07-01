package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

// PoolConfig holds all configurable pool settings
type PoolConfig struct {
	StratumPort   int       `json:"stratum_port"`
	PoolWallet    string    `json:"pool_wallet"`
	PoolName      string    `json:"pool_name"`
	PoolFee       float64   `json:"pool_fee"`
	SoloFee       float64   `json:"solo_fee"`
	MinPayout     float64   `json:"min_payout"`
	CoinbaseTag   string    `json:"coinbase_tag"`
	VardiffMinDiff float64  `json:"vardiff_min_diff"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GetDefaults returns default configuration
func GetDefaults() *PoolConfig {
	return &PoolConfig{
		StratumPort:    3333,
		PoolWallet:     "",
		PoolName:       "My Void Pool",
		PoolFee:        1.0,
		SoloFee:        0.5,
		MinPayout:      5.0,
		CoinbaseTag:    "VoidCoin",
		VardiffMinDiff: 32768,
		UpdatedAt:      time.Now(),
	}
}

// LoadConfig loads configuration from the data directory
func LoadConfig(dataDir string) (*PoolConfig, error) {
	configPath := filepath.Join(dataDir, "config", "pool-config.json")

	data, err := os.ReadFile(configPath)
	if err != nil {
		// Return defaults on any read error (not found, permissions, etc)
		return GetDefaults(), nil
	}

	var cfg PoolConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		// Return defaults if file is corrupted
		return GetDefaults(), nil
	}

	// Apply defaults only for truly missing values (not 0 which is valid)
	defaults := GetDefaults()
	if cfg.StratumPort == 0 {
		cfg.StratumPort = defaults.StratumPort
	}
	if cfg.PoolName == "" {
		cfg.PoolName = defaults.PoolName
	}
	if cfg.CoinbaseTag == "" {
		cfg.CoinbaseTag = defaults.CoinbaseTag
	}
	// Note: PoolFee, SoloFee can be 0 (free pool), don't override
	// MinPayout validated separately (must be >= 0.1)

	return &cfg, nil
}

// SaveConfig saves configuration to the data directory
func SaveConfig(dataDir string, cfg *PoolConfig) error {
	configDir := filepath.Join(dataDir, "config")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	configPath := filepath.Join(configDir, "pool-config.json")

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// ValidateConfig validates the configuration values
func ValidateConfig(cfg *PoolConfig) error {
	if cfg.StratumPort < 1024 || cfg.StratumPort > 65535 {
		return errors.New("stratum port must be between 1024 and 65535")
	}

	if cfg.PoolFee < 0 || cfg.PoolFee > 10 {
		return errors.New("pool fee must be between 0 and 10 percent")
	}

	if cfg.SoloFee < 0 || cfg.SoloFee > 10 {
		return errors.New("solo fee must be between 0 and 10 percent")
	}

	if cfg.MinPayout < 0.1 || cfg.MinPayout > 1000 {
		return errors.New("minimum payout must be between 0.1 and 1000")
	}

	if cfg.PoolName == "" {
		cfg.PoolName = "My Void Pool"
	}

	// Validate CoinbaseTag: 1-20 chars, ASCII printable only
	if cfg.CoinbaseTag != "" {
		if len(cfg.CoinbaseTag) > 20 {
			return errors.New("coinbase tag must be 20 characters or less")
		}
		for _, c := range cfg.CoinbaseTag {
			if c < 32 || c > 126 {
				return errors.New("coinbase tag must contain only ASCII printable characters")
			}
		}
	}

	// Validate VardiffMinDiff: 1024-500000
	if cfg.VardiffMinDiff != 0 && (cfg.VardiffMinDiff < 1024 || cfg.VardiffMinDiff > 500000) {
		return errors.New("vardiff minimum difficulty must be between 1024 and 500000")
	}

	return nil
}

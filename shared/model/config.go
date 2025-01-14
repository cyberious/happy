package model

import "gorm.io/gorm"

type ConfigKey struct {
	Key string `json:"key" validate:"required" gorm:"index:,unique,composite:metadata"`
}

type ConfigValue struct {
	ConfigKey
	Value string `json:"value" validate:"required"`
}

type AppConfigPayload struct {
	AppMetadata
	ConfigValue
} // @Name payload.AppConfig

type AppConfigLookupPayload struct {
	AppMetadata
	ConfigKey
} // @Name payload.AppConfigLookup

type CopyAppConfigPayload struct {
	App
	SrcEnvironment string `json:"source_environment"      validate:"required,valid_env" gorm:"index:,unique,composite:metadata"`
	SrcStack       string `json:"source_stack,omitempty"                                gorm:"default:'';not null;index:,unique,composite:metadata"`
	DstEnvironment string `json:"destination_environment" validate:"required,valid_env_dest" gorm:"index:,unique,composite:metadata"`
	DstStack       string `json:"destination_stack,omitempty"                           gorm:"default:'';not null;index:,unique,composite:metadata"`
	ConfigKey
} // @name payload.CopyAppConfig

type AppConfigDiffPayload struct {
	App
	SrcEnvironment string `json:"source_environment"      validate:"required,valid_env" gorm:"index:,unique,composite:metadata"`
	SrcStack       string `json:"source_stack,omitempty"                                gorm:"default:'';not null;index:,unique,composite:metadata"`
	DstEnvironment string `json:"destination_environment" validate:"required,valid_env_dest" gorm:"index:,unique,composite:metadata"`
	DstStack       string `json:"destination_stack,omitempty"                           gorm:"default:'';not null;index:,unique,composite:metadata"`
} // @name payload.AppConfigDiff

// @Description Object denoting which app config keys are missing from the destination env/stack
type ConfigDiffResponse struct {
	MissingKeys []string `json:"missing_keys" example:"SOME_KEY,ANOTHER_KEY"`
} // @Name response.ConfigDiff

// @Description App config key/value pair with additional metadata
type AppConfig struct {
	gorm.Model `swaggerignore:"true"`
	AppConfigPayload
} // @Name response.AppConfig

// @Description App config key/value pair with additional metadata and "source"
type ResolvedAppConfig struct {
	AppConfig
	Source string `json:"source"`
} // @Name response.ResolvedAppConfig

func NewAppConfigPayload(appName, env, stack, key, value string) *AppConfigPayload {
	return &AppConfigPayload{
		AppMetadata: *NewAppMetadata(appName, env, stack),
		ConfigValue: ConfigValue{
			Value: value,
			ConfigKey: ConfigKey{
				Key: key,
			},
		},
	}
}

func NewAppConfigLookupPayload(appName, env, stack, key string) *AppConfigLookupPayload {
	return &AppConfigLookupPayload{
		AppMetadata: *NewAppMetadata(appName, env, stack),
		ConfigKey: ConfigKey{
			Key: key,
		},
	}
}

func NewCopyAppConfigPayload(appName, srcEnv, srcStack, destEnv, destStack, key string) *CopyAppConfigPayload {
	return &CopyAppConfigPayload{
		App:            App{AppName: appName},
		SrcEnvironment: srcEnv,
		SrcStack:       srcStack,
		DstEnvironment: destEnv,
		DstStack:       destStack,
		ConfigKey: ConfigKey{
			Key: key,
		},
	}
}

func NewAppConfigDiffPayload(appName, srcEnv, srcStack, destEnv, destStack string) *AppConfigDiffPayload {
	return &AppConfigDiffPayload{
		App:            App{AppName: appName},
		SrcEnvironment: srcEnv,
		SrcStack:       srcStack,
		DstEnvironment: destEnv,
		DstStack:       destStack,
	}
}

// @Description App config key/value pair wrapped in "record" key
type WrappedResolvedAppConfig struct {
	Record *ResolvedAppConfig `json:"record"`
} // @Name response.WrappedResolvedAppConfig

// @Description App config key/value pair wrapped in "record" key
type WrappedAppConfig struct {
	Record *AppConfig `json:"record"`
} // @Name response.WrappedAppConfig

type WrappedAppConfigsWithCount struct {
	Records []*AppConfig `json:"records"`
	Count   int          `json:"count" example:"1"`
} // @Name response.WrappedAppConfigsWithCount

type WrappedResolvedAppConfigsWithCount struct {
	Records []*ResolvedAppConfig `json:"records"`
	Count   int                  `json:"count" example:"1"`
} // @Name response.WrappedResolvedAppConfigsWithCount

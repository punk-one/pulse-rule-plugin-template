package main

import sdk "github.com/pulse-rule-engine/pulse-rule-plugin-sdk"

type Plugin struct{}

func (p *Plugin) Info() sdk.PluginInfo {
	return sdk.PluginInfo{
		ID:      "template_plugin",
		Name:    "Template Plugin",
		Version: "0.1.0",
		Type:    sdk.PluginTypeProcessor,
	}
}

func (p *Plugin) Init(cfg sdk.PluginConfig) error { return nil }
func (p *Plugin) Start(ctx sdk.RuntimeContext) error { return nil }
func (p *Plugin) OnEvent(e sdk.Event) error { return nil }
func (p *Plugin) OnEvents(events []sdk.Event) error { return nil }
func (p *Plugin) Stop() error { return nil }

func main() {}

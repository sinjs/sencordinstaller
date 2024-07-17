/*
 * SPDX-License-Identifier: GPL-3.0
 * Vencord Installer, a cross platform gui/cli app for installing Vencord
 * Copyright (c) 2023 Vendicated and Vencord contributors
 */

package main

import (
	"image/color"
	"vencordinstaller/buildinfo"
)

const ReleaseUrl = "https://api.github.com/repos/sinjs/Sencord/releases/latest"
const ReleaseUrlFallback = "https://api.nigga.church/v1/releases/sencord/latest"

// TODO TODO TODO :3333
const InstallerReleaseUrl = "https://api.github.com/repos/Vencord/Installer/releases/latest"
const InstallerReleaseUrlFallback = "https://vencord.dev/releases/installer"

var UserAgent = "VencordInstaller/" + buildinfo.InstallerGitHash + " (https://github.com/Vencord/Installer)"

var (
	DiscordGreen  = color.RGBA{R: 16, G: 138, B: 0, A: 0xFF}
	DiscordRed    = color.RGBA{R: 196, G: 0, B: 95, A: 0xFF}
	DiscordBlue   = color.RGBA{R: 196, G: 0, B: 196, A: 0xFF}
	DiscordYellow = color.RGBA{R: 0xfe, G: 0xe7, B: 0x5c, A: 0xff}
)

var LinuxDiscordNames = []string{
	"Discord",
	"DiscordPTB",
	"DiscordCanary",
	"DiscordDevelopment",
	"discord",
	"discordptb",
	"discordcanary",
	"discorddevelopment",
	"discord-ptb",
	"discord-canary",
	"discord-development",
	// Flatpak
	"com.discordapp.Discord",
	"com.discordapp.DiscordPTB",
	"com.discordapp.DiscordCanary",
	"com.discordapp.DiscordDevelopment",
}

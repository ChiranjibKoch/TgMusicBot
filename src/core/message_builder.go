/*
 * TgMusicBot - Telegram Music Bot
 *  Copyright (c) 2025-2026 Ashok Shau
 *
 *  Licensed under GNU GPL v3
 *  See https://github.com/AshokShau/TgMusicBot
 */

package core

import (
	"ashokshau/tgmusic/src/utils"
	"fmt"
	"html"
	"strconv"
	"strings"
)

// GetPlatformIcon returns a platform-specific emoji icon.
func GetPlatformIcon(platform string) string {
	switch strings.ToLower(platform) {
	case utils.YouTube:
		return "🎥"
	case utils.Spotify:
		return "🎵"
	case utils.JioSaavn:
		return "🎶"
	case utils.Apple:
		return "🍎"
	case utils.SoundCloud:
		return "☁️"
	case utils.Deezer:
		return "🎧"
	case utils.Tidal:
		return "🌊"
	case utils.Twitch, utils.TwitchClip:
		return "📡"
	case utils.Kick, utils.KickClip:
		return "🎮"
	case utils.Telegram:
		return "✈️"
	default:
		return "🎵"
	}
}

// GetPlatformName returns a human-readable platform name.
func GetPlatformName(platform string) string {
	switch strings.ToLower(platform) {
	case utils.YouTube:
		return "YouTube"
	case utils.Spotify:
		return "Spotify"
	case utils.JioSaavn:
		return "JioSaavn"
	case utils.Apple:
		return "Apple Music"
	case utils.SoundCloud:
		return "SoundCloud"
	case utils.Deezer:
		return "Deezer"
	case utils.Tidal:
		return "Tidal"
	case utils.Twitch:
		return "Twitch"
	case utils.TwitchClip:
		return "Twitch Clip"
	case utils.Kick:
		return "Kick"
	case utils.KickClip:
		return "Kick Clip"
	case utils.Telegram:
		return "Telegram"
	case utils.DirectLink:
		return "Direct Link"
	default:
		if platform != "" {
			return platform
		}
		return "Unknown"
	}
}

// FormatViews formats a raw view count string into a human-readable form (e.g. 1.2M, 45K).
func FormatViews(views string) string {
	if views == "" {
		return ""
	}

	n, err := strconv.ParseInt(views, 10, 64)
	if err != nil {
		// already formatted or non-numeric – return as-is
		return views
	}

	switch {
	case n >= 1_000_000_000:
		return fmt.Sprintf("%.1fB", float64(n)/1_000_000_000)
	case n >= 1_000_000:
		return fmt.Sprintf("%.1fM", float64(n)/1_000_000)
	case n >= 1_000:
		return fmt.Sprintf("%.1fK", float64(n)/1_000)
	default:
		return strconv.FormatInt(n, 10)
	}
}

// BuildNowPlayingCaption builds an HTML-formatted caption for a "now playing" message.
func BuildNowPlayingCaption(track *utils.CachedTrack) string {
	icon := GetPlatformIcon(track.Platform)
	escTitle := html.EscapeString(track.Name)
	escUser := html.EscapeString(track.User)

	var sb strings.Builder
	fmt.Fprintf(&sb, "%s <b>%s</b>\n", icon, escTitle)
	sb.WriteString("┌─────────────────\n")

	if track.Channel != "" {
		fmt.Fprintf(&sb, "├─ 🎤 <b>%s</b>\n", html.EscapeString(track.Channel))
	}

	viewsStr := FormatViews(track.Views)
	dur := utils.SecToMin(track.Duration)
	if viewsStr != "" {
		fmt.Fprintf(&sb, "├─ 📊 <b>%s views</b> • <b>%s</b>\n", viewsStr, dur)
	} else {
		fmt.Fprintf(&sb, "├─ 📊 <b>%s</b>\n", dur)
	}

	platformName := GetPlatformName(track.Platform)
	fmt.Fprintf(&sb, "├─ 🔗 <b>Platform:</b> %s\n", html.EscapeString(platformName))
	fmt.Fprintf(&sb, "└─ 👤 <b>Requested by:</b> %s", escUser)

	return sb.String()
}

// BuildQueueCaption builds an HTML-formatted caption for a "added to queue" message.
func BuildQueueCaption(track *utils.CachedTrack, queuePos int) string {
	icon := GetPlatformIcon(track.Platform)
	escTitle := html.EscapeString(track.Name)
	escUser := html.EscapeString(track.User)

	var sb strings.Builder
	fmt.Fprintf(&sb, "📋 <b>Added to Queue: %d</b>\n\n", queuePos)
	fmt.Fprintf(&sb, "%s <b>%s</b>\n", icon, escTitle)
	sb.WriteString("┌─────────────────\n")

	if track.Channel != "" {
		fmt.Fprintf(&sb, "├─ 🎤 <b>%s</b>\n", html.EscapeString(track.Channel))
	}

	viewsStr := FormatViews(track.Views)
	dur := utils.SecToMin(track.Duration)
	if viewsStr != "" {
		fmt.Fprintf(&sb, "├─ 📊 <b>%s views</b> • <b>%s</b>\n", viewsStr, dur)
	} else {
		fmt.Fprintf(&sb, "├─ 📊 <b>%s</b>\n", dur)
	}

	platformName := GetPlatformName(track.Platform)
	fmt.Fprintf(&sb, "├─ 🔗 <b>Platform:</b> %s\n", html.EscapeString(platformName))
	fmt.Fprintf(&sb, "└─ 👤 <b>Requested by:</b> %s", escUser)

	return sb.String()
}

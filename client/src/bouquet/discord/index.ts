import Discord from "./bootstrap/index.ts";
import DiscordOverlay from "./overlay/index.ts";

function discord() {
  const discord = new Discord();

  discord.set_linux_pulse_latency();
}

function discord_overlay() {
  const discord_overlay = new DiscordOverlay();
}

function start_discord(overlay: boolean = false) {
  if (overlay) {
    discord_overlay();
  } else {
    discord();
  }
}

export default start_discord;

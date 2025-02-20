import { app, session, Menu } from "electron";
import global from "../../global/index.ts";
import os, { OSPlatform } from "../../os/index.ts";

import LinuxSetup from "./linux.ts";
import MacOSSetup from "./macos.ts";
import WindowsSetup from "./windows.ts";

type Setup = {
  setup: () => void;
  disabled: string[];
  switch: { [key: string]: string };
};

const osSetup: { [key: string]: Setup } = {
  [OSPlatform.LINUX]: LinuxSetup,
  [OSPlatform.MACOS]: MacOSSetup,
  [OSPlatform.WINDOWS]: WindowsSetup,
};

class Discord {
  disabledFeatures: string[] = [
    "WinRetrieveSuggestionsOnlyOnDemand",
    "HardwareMediaKeyHandling",
    "MediaSessionService",
    "UseEcoQoSForBackgroundProcess",
    "IntensiveWakeUpThrottling",
    "AllowAggressiveThrottlingWithWebSocket",
  ];

  constructor() {
    app.setAboutPanelOptions({
      version: global.buildInfo.version,
      applicationVersion: global.buildInfo.version,
    });
  }

  setup() {
    const system = osSetup[os.os];
    system.setup();
    this.disabledFeatures.push(...system.disabled);
    for (const [key, value] of Object.entries(system.switch)) {
      app.commandLine.appendSwitch(key, value);
    }

    app.commandLine.appendSwitch("autoplay-policy", "no-user-gesture-required");
    app.commandLine.appendSwitch(
      "disable-features",
      this.disabledFeatures.join(","),
    );
  }
}

export default Discord;

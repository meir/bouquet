import os from "../../os/index.ts";

let disabled: string[] = [];

function setup() {
  if (os.major < 24) {
    disabled = [
      "ScreenCaptureKitMac",
      "ScreenCaptureKitMacWindow",
      "ScreenCaptureKitMacScreen",
      "ScreenCaptureKitPickerScreen",
      "ScreenCaptureKitStreamPickerSonoma",
      "WarmScreenCaptureSonoma",
      "UseSCContentSharingPicker",
    ];
  }
}

export default {
  setup,
  disabled,
  switch: {},
};

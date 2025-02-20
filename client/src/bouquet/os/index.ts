import process from "node:process";
import os from "node:os";

enum OSPlatform {
  LINUX = "linux",
  MACOS = "darwin",
  WINDOWS = "win32",
}

class OS {
  os: OSPlatform;
  major: number;
  minor: number;
  patch: number;
  version: string;

  constructor() {
    this.os = process.platform as OSPlatform;
    this.version = os.release();
    const version = this.version.split(".");
    this.major = parseInt(version[0]);
    this.minor = parseInt(version[1]);
    this.patch = parseInt(version[2]);
  }
}

export default new OS();
export { OSPlatform };

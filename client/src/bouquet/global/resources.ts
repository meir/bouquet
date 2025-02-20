import path from "node:path";
import fs from "node:fs";
import process from "node:process";

class Resources {
  resourcePath: string;

  constructor(resourcePath: string) {
    this.resourcePath = resourcePath;
  }

  path(resource: string): string {
    return path.join(this.resourcePath, resource);
  }

  read(resource: string): string {
    return fs.readFileSync(this.path(resource), "utf8");
  }

  readJSON<T>(resource: string): T {
    return JSON.parse(this.read(resource)) as T;
  }
}

export default new Resources(process.resourcesPath);

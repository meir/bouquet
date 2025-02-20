import path from "node:path";
import fs from "node:fs";
import process from "node:process";
import resources from "./resources.ts";
import BuildInfo from "./build_info.ts";

type Global = {
  [key: string]: string | number | object | boolean;
};

declare let global: Global;
declare let __dirname: string;

// Get Version
const VERSION = path.join(__dirname, "../VERSION");
const version = fs
  .readFileSync(VERSION, "utf8")
  .toString()
  .replace(/[^0-9.]+/g, "");

//
export default global = {
  bouquetVersion: version,
  overlay: process.argv.includes("--overlay-host"),
  resources: resources,
  buildInfo: resources.readJSON<BuildInfo>("build_info.json"),
};

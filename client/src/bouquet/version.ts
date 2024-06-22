
// Read the version from the VERSION file, it should be injected in the same folder
declare let __dirname: string;
// @ts-ignore
import path from 'path';
// @ts-ignore
import fs from 'fs';

const VERSION = path.join(__dirname, './VERSION')
const version = fs.readFileSync(VERSION, 'utf8').toString().replace(/[^0-9.]+/g, "");

export default version;

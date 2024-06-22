
// Read the version from the VERSION file, it should be injected in the same folder
declare let __dirname: string;
// @ts-ignore
import path from 'path';
// @ts-ignore
import fs from 'fs';
import window from './window.ts';

// inject will read the contents of the given path of the file (from src/) and inject them into the renderer process window
// this can be used to add additional frontend code
function inject(file: string) {
  const filePath = path.join(__dirname, "../../", file)
  const content = fs.readFileSync(filePath, 'utf8').toString()
  window().webContents.executeJavaScript(content)
}

export default inject;

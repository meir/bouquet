
// @ts-ignore
import { contextBridge } from 'electron';
import native from './native.ts'

// this runs in the renderer process
export default () => {
  contextBridge.exposeInMainWorld("bouquet", native)
}

import { app } from "electron";
import { GPUInfo, GPU, NVIDIA } from "../../os/gpu.ts";

const disable_accelerated_hevc_decode_gpus = [
  new NVIDIA(0x1340),
  new NVIDIA(0x1341),
  new NVIDIA(0x1344),
  new NVIDIA(0x1346),
  new NVIDIA(0x1347),
  new NVIDIA(0x1348),
  new NVIDIA(0x1349),
  new NVIDIA(0x134b),
  new NVIDIA(0x134d),
  new NVIDIA(0x134e),
  new NVIDIA(0x134f),
  new NVIDIA(0x137a),
  new NVIDIA(0x137b),
  new NVIDIA(0x1380),
  new NVIDIA(0x1381),
  new NVIDIA(0x1382),
  new NVIDIA(0x1390),
  new NVIDIA(0x1391),
  new NVIDIA(0x1392),
  new NVIDIA(0x1393),
  new NVIDIA(0x1398),
  new NVIDIA(0x1399),
  new NVIDIA(0x139a),
  new NVIDIA(0x139b),
  new NVIDIA(0x139c),
  new NVIDIA(0x139d),
  new NVIDIA(0x13b0),
  new NVIDIA(0x13b1),
  new NVIDIA(0x13b2),
  new NVIDIA(0x13b3),
  new NVIDIA(0x13b4),
  new NVIDIA(0x13b6),
  new NVIDIA(0x13b9),
  new NVIDIA(0x13ba),
  new NVIDIA(0x13bb),
  new NVIDIA(0x13bc),
  new NVIDIA(0x13c0),
  new NVIDIA(0x13c2),
  new NVIDIA(0x13d7),
  new NVIDIA(0x13d8),
  new NVIDIA(0x13d9),
  new NVIDIA(0x13da),
  new NVIDIA(0x13f0),
  new NVIDIA(0x13f1),
  new NVIDIA(0x13f2),
  new NVIDIA(0x13f3),
  new NVIDIA(0x13f8),
  new NVIDIA(0x13f9),
  new NVIDIA(0x13fa),
  new NVIDIA(0x13fb),
  new NVIDIA(0x1401),
  new NVIDIA(0x1406),
  new NVIDIA(0x1407),
  new NVIDIA(0x1427),
  new NVIDIA(0x1617),
  new NVIDIA(0x1618),
  new NVIDIA(0x1619),
  new NVIDIA(0x161a),
  new NVIDIA(0x1667),
  new NVIDIA(0x174d),
  new NVIDIA(0x174e),
  new NVIDIA(0x179c),
  new NVIDIA(0x17c2),
  new NVIDIA(0x17c8),
  new NVIDIA(0x17f0),
  new NVIDIA(0x17f1),
  new NVIDIA(0x17fd),
];

const switchFlags: { [key: string]: string } = {
  "disable-background-timer-throttling": "1",
};

async function setup() {
  const info = (await app.getGPUInfo("basic")) as GPUInfo;

  for (const gpu of info.gpuDevice) {
    if (disable_accelerated_hevc_decode_gpus.some((g: GPU) => g == gpu)) {
      switchFlags["disable-accelerated-hevc-decode"] = "1";
      break;
    }
  }
}

export default {
  setup,
  switch: switchFlags,
  disabled: [],
};

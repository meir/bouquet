export type GPUInfo = {
  gpuDevice: GPU[];
};

export enum GPUVendor {
  NVIDIA = 0x10de,
  AMD = 0x1002,
  INTEL = 0x8086,
}

export class GPU {
  vendorId: GPUVendor;
  deviceId: number;
  constructor(vendorId: GPUVendor, deviceId: number) {
    this.vendorId = vendorId;
    this.deviceId = deviceId;
  }
}

export class NVIDIA extends GPU {
  constructor(deviceId: number) {
    super(GPUVendor.NVIDIA, deviceId);
  }
}

export class AMD extends GPU {
  constructor(deviceId: number) {
    super(GPUVendor.AMD, deviceId);
  }
}

export class INTEL extends GPU {
  constructor(deviceId: number) {
    super(GPUVendor.INTEL, deviceId);
  }
}

/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "soupyfinance.noodle.bridge";

export interface MsgObserveDeposit {
  creator: string;
  chainId: string;
  depositor: string;
  depositId: string;
  quantity: string;
  asset: string;
}

export interface MsgObserveDepositResponse {}

export interface MsgWithdraw {
  creator: string;
  asset: string;
  quantity: string;
  address: string;
  chainId: string;
}

export interface MsgWithdrawResponse {}

const baseMsgObserveDeposit: object = {
  creator: "",
  chainId: "",
  depositor: "",
  depositId: "",
  quantity: "",
  asset: "",
};

export const MsgObserveDeposit = {
  encode(message: MsgObserveDeposit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.chainId !== "") {
      writer.uint32(18).string(message.chainId);
    }
    if (message.depositor !== "") {
      writer.uint32(26).string(message.depositor);
    }
    if (message.depositId !== "") {
      writer.uint32(34).string(message.depositId);
    }
    if (message.quantity !== "") {
      writer.uint32(42).string(message.quantity);
    }
    if (message.asset !== "") {
      writer.uint32(50).string(message.asset);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgObserveDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgObserveDeposit } as MsgObserveDeposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.chainId = reader.string();
          break;
        case 3:
          message.depositor = reader.string();
          break;
        case 4:
          message.depositId = reader.string();
          break;
        case 5:
          message.quantity = reader.string();
          break;
        case 6:
          message.asset = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgObserveDeposit {
    const message = { ...baseMsgObserveDeposit } as MsgObserveDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = String(object.chainId);
    } else {
      message.chainId = "";
    }
    if (object.depositor !== undefined && object.depositor !== null) {
      message.depositor = String(object.depositor);
    } else {
      message.depositor = "";
    }
    if (object.depositId !== undefined && object.depositId !== null) {
      message.depositId = String(object.depositId);
    } else {
      message.depositId = "";
    }
    if (object.quantity !== undefined && object.quantity !== null) {
      message.quantity = String(object.quantity);
    } else {
      message.quantity = "";
    }
    if (object.asset !== undefined && object.asset !== null) {
      message.asset = String(object.asset);
    } else {
      message.asset = "";
    }
    return message;
  },

  toJSON(message: MsgObserveDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.depositor !== undefined && (obj.depositor = message.depositor);
    message.depositId !== undefined && (obj.depositId = message.depositId);
    message.quantity !== undefined && (obj.quantity = message.quantity);
    message.asset !== undefined && (obj.asset = message.asset);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgObserveDeposit>): MsgObserveDeposit {
    const message = { ...baseMsgObserveDeposit } as MsgObserveDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = object.chainId;
    } else {
      message.chainId = "";
    }
    if (object.depositor !== undefined && object.depositor !== null) {
      message.depositor = object.depositor;
    } else {
      message.depositor = "";
    }
    if (object.depositId !== undefined && object.depositId !== null) {
      message.depositId = object.depositId;
    } else {
      message.depositId = "";
    }
    if (object.quantity !== undefined && object.quantity !== null) {
      message.quantity = object.quantity;
    } else {
      message.quantity = "";
    }
    if (object.asset !== undefined && object.asset !== null) {
      message.asset = object.asset;
    } else {
      message.asset = "";
    }
    return message;
  },
};

const baseMsgObserveDepositResponse: object = {};

export const MsgObserveDepositResponse = {
  encode(
    _: MsgObserveDepositResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgObserveDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgObserveDepositResponse,
    } as MsgObserveDepositResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgObserveDepositResponse {
    const message = {
      ...baseMsgObserveDepositResponse,
    } as MsgObserveDepositResponse;
    return message;
  },

  toJSON(_: MsgObserveDepositResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgObserveDepositResponse>
  ): MsgObserveDepositResponse {
    const message = {
      ...baseMsgObserveDepositResponse,
    } as MsgObserveDepositResponse;
    return message;
  },
};

const baseMsgWithdraw: object = {
  creator: "",
  asset: "",
  quantity: "",
  address: "",
  chainId: "",
};

export const MsgWithdraw = {
  encode(message: MsgWithdraw, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.asset !== "") {
      writer.uint32(18).string(message.asset);
    }
    if (message.quantity !== "") {
      writer.uint32(26).string(message.quantity);
    }
    if (message.address !== "") {
      writer.uint32(34).string(message.address);
    }
    if (message.chainId !== "") {
      writer.uint32(42).string(message.chainId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdraw {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdraw } as MsgWithdraw;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.asset = reader.string();
          break;
        case 3:
          message.quantity = reader.string();
          break;
        case 4:
          message.address = reader.string();
          break;
        case 5:
          message.chainId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdraw {
    const message = { ...baseMsgWithdraw } as MsgWithdraw;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.asset !== undefined && object.asset !== null) {
      message.asset = String(object.asset);
    } else {
      message.asset = "";
    }
    if (object.quantity !== undefined && object.quantity !== null) {
      message.quantity = String(object.quantity);
    } else {
      message.quantity = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = String(object.chainId);
    } else {
      message.chainId = "";
    }
    return message;
  },

  toJSON(message: MsgWithdraw): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.asset !== undefined && (obj.asset = message.asset);
    message.quantity !== undefined && (obj.quantity = message.quantity);
    message.address !== undefined && (obj.address = message.address);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgWithdraw>): MsgWithdraw {
    const message = { ...baseMsgWithdraw } as MsgWithdraw;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.asset !== undefined && object.asset !== null) {
      message.asset = object.asset;
    } else {
      message.asset = "";
    }
    if (object.quantity !== undefined && object.quantity !== null) {
      message.quantity = object.quantity;
    } else {
      message.quantity = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.chainId !== undefined && object.chainId !== null) {
      message.chainId = object.chainId;
    } else {
      message.chainId = "";
    }
    return message;
  },
};

const baseMsgWithdrawResponse: object = {};

export const MsgWithdrawResponse = {
  encode(_: MsgWithdrawResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdrawResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdrawResponse } as MsgWithdrawResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgWithdrawResponse {
    const message = { ...baseMsgWithdrawResponse } as MsgWithdrawResponse;
    return message;
  },

  toJSON(_: MsgWithdrawResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgWithdrawResponse>): MsgWithdrawResponse {
    const message = { ...baseMsgWithdrawResponse } as MsgWithdrawResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  ObserveDeposit(
    request: MsgObserveDeposit
  ): Promise<MsgObserveDepositResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Withdraw(request: MsgWithdraw): Promise<MsgWithdrawResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  ObserveDeposit(
    request: MsgObserveDeposit
  ): Promise<MsgObserveDepositResponse> {
    const data = MsgObserveDeposit.encode(request).finish();
    const promise = this.rpc.request(
      "soupyfinance.noodle.bridge.Msg",
      "ObserveDeposit",
      data
    );
    return promise.then((data) =>
      MsgObserveDepositResponse.decode(new Reader(data))
    );
  }

  Withdraw(request: MsgWithdraw): Promise<MsgWithdrawResponse> {
    const data = MsgWithdraw.encode(request).finish();
    const promise = this.rpc.request(
      "soupyfinance.noodle.bridge.Msg",
      "Withdraw",
      data
    );
    return promise.then((data) => MsgWithdrawResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

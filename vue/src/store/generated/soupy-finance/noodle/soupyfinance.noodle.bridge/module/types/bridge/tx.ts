/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "soupyfinance.noodle.bridge";

export interface MsgObserveDeposit {
  creator: string;
  chainId: string;
  depositor: string;
}

export interface MsgObserveDepositResponse {}

const baseMsgObserveDeposit: object = {
  creator: "",
  chainId: "",
  depositor: "",
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
    return message;
  },

  toJSON(message: MsgObserveDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.chainId !== undefined && (obj.chainId = message.chainId);
    message.depositor !== undefined && (obj.depositor = message.depositor);
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

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  ObserveDeposit(
    request: MsgObserveDeposit
  ): Promise<MsgObserveDepositResponse>;
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

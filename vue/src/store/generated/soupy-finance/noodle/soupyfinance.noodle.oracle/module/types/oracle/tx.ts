/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "soupyfinance.noodle.oracle";

export interface MsgUpdatePrices {
  creator: string;
  data: string;
}

export interface MsgUpdatePricesResponse {}

const baseMsgUpdatePrices: object = { creator: "", data: "" };

export const MsgUpdatePrices = {
  encode(message: MsgUpdatePrices, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.data !== "") {
      writer.uint32(18).string(message.data);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdatePrices {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgUpdatePrices } as MsgUpdatePrices;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.data = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdatePrices {
    const message = { ...baseMsgUpdatePrices } as MsgUpdatePrices;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = String(object.data);
    } else {
      message.data = "";
    }
    return message;
  },

  toJSON(message: MsgUpdatePrices): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.data !== undefined && (obj.data = message.data);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgUpdatePrices>): MsgUpdatePrices {
    const message = { ...baseMsgUpdatePrices } as MsgUpdatePrices;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = "";
    }
    return message;
  },
};

const baseMsgUpdatePricesResponse: object = {};

export const MsgUpdatePricesResponse = {
  encode(_: MsgUpdatePricesResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgUpdatePricesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgUpdatePricesResponse,
    } as MsgUpdatePricesResponse;
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

  fromJSON(_: any): MsgUpdatePricesResponse {
    const message = {
      ...baseMsgUpdatePricesResponse,
    } as MsgUpdatePricesResponse;
    return message;
  },

  toJSON(_: MsgUpdatePricesResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgUpdatePricesResponse>
  ): MsgUpdatePricesResponse {
    const message = {
      ...baseMsgUpdatePricesResponse,
    } as MsgUpdatePricesResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  UpdatePrices(request: MsgUpdatePrices): Promise<MsgUpdatePricesResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  UpdatePrices(request: MsgUpdatePrices): Promise<MsgUpdatePricesResponse> {
    const data = MsgUpdatePrices.encode(request).finish();
    const promise = this.rpc.request(
      "soupyfinance.noodle.oracle.Msg",
      "UpdatePrices",
      data
    );
    return promise.then((data) =>
      MsgUpdatePricesResponse.decode(new Reader(data))
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
